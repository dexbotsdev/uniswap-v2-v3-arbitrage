package executor

import (
	"errors"
	"fmt"
	"math/big"
	"mev-template-go/logic"
	"mev-template-go/uniswap_v3"

	"github.com/ethereum/go-ethereum/common"
)

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

//"github.com/umbracle/ethgo/abi"

// Define the TokenIndexes struct
type TokenIndexes struct {
	Token0Index int
	Token1Index int
}

func getV2FactoryIndex(factoryAddress common.Address) int {
	//index 1 is not used as it is reserved for a custom address
	if factoryAddress == common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f") {
		return 0
	}
	return 1
}

func getV3Factory(factoryAddress common.Address) int {
	//index 1 is not used as it is reserved for a custom address
	if factoryAddress == common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984") {
		return 0
	}
	return 1
}

func compressTokens(path logic.Path) ([]common.Address, map[common.Address]int) {
	pools := path.Pools
	tokenToIndexMap := make(map[common.Address]int)

	//get all tokens
	allTokens := []common.Address{}
	for i := 0; i < len(pools); i++ {
		allTokens = append(allTokens, pools[i].GetTokens()[0], pools[i].GetTokens()[1])
	}

	//remove duplicatets
	uniqueTokens := []common.Address{}
	for i := 0; i < len(allTokens); i++ {
		isUnique := true
		for j := 0; j < len(uniqueTokens); j++ {
			if allTokens[i] == uniqueTokens[j] {
				isUnique = false
				break
			}
		}
		if isUnique {
			uniqueTokens = append(uniqueTokens, allTokens[i])
		}
	}

	//remove base token
	for i := 0; i < len(uniqueTokens); i++ {
		if uniqueTokens[i] == path.BaseToken {
			uniqueTokens = append(uniqueTokens[:i], uniqueTokens[i+1:]...)
			break
		}
	}

	//create map of token to index
	for i := 0; i < len(uniqueTokens); i++ {
		tokenToIndexMap[uniqueTokens[i]] = i
	}

	//add wethMapping to the end of the map
	tokenToIndexMap[path.BaseToken] = len(uniqueTokens)

	fmt.Println("tokenToIndexMap", tokenToIndexMap)
	fmt.Println("uniqueTokens", uniqueTokens)

	return uniqueTokens, tokenToIndexMap
}

func getV2CallbackV3SwapsDone(path logic.Path) bool {
	//if there is a v3 swap, then v3 swaps are not done. The last swap is ignored since we know its v2.
	for i := 0; i < len(path.Pools)-1; i++ {
		if path.Pools[i].GetType() == "uniswap_v3" {
			return false
		}
	}
	return true
}

func getV2SwapDatas(path logic.Path, amountsIndexMap map[int]int) ([]V2SwapData, error) {
	// type V2SwapData struct {
	// 	swapIndex int //3 bits
	// 	AmountIndex      int //3 bits
	// }
	v2SwapDatas := []V2SwapData{}
	for i := 0; i < len(path.Pools)-1; i++ {
		if path.Pools[i].GetType() == "uniswap_v2" {
			swapIndex := i
			amountIndex := amountsIndexMap[i]

			v2SwapData := V2SwapData{
				SwapIndex:   swapIndex,
				AmountIndex: amountIndex,
			}
			v2SwapDatas = append(v2SwapDatas, v2SwapData)
		}
	}
	return v2SwapDatas, nil
}

func getV3SwapDatas(path logic.Path, targets []common.Address, amountsIndexMap map[int]int, tokenToIndexMap map[common.Address]int) ([]V3SwapData, error) {
	//Overview
	//get reversed indexes
	v3SwapLength := 0
	for i := 0; i < len(path.Pools); i++ {
		if path.Pools[i].GetType() == "uniswap_v3" {
			v3SwapLength++
		}
	}

	firstV3SwapFlag := true

	v3SwapDatas := []V3SwapData{}
	for i := 0; i < len(path.Pools); i++ {
		if path.Pools[i].GetType() == "uniswap_v3" {
			v3SwapsDone := false
			if firstV3SwapFlag {
				v3SwapsDone = true //this is becuase this array will be reversed
				firstV3SwapFlag = false
			}
			swapIndex := i
			feeIndex, err := getFeeIndex(int(path.Pools[i].(*uniswap_v3.Pool).Fee))
			if err != nil {
				return nil, err
			}

			isFlashswap := i == len(path.Pools)-1

			factoryIndex := getV3Factory(path.Pools[i].GetFactoryAddress())
			amountIndex := amountsIndexMap[i]

			factoryAddress := path.Pools[i].GetFactoryAddress()
			Token0Index := tokenToIndexMap[path.Pools[i].GetTokens()[0]]
			Token1Index := tokenToIndexMap[path.Pools[i].GetTokens()[1]]

			v3SwapData := V3SwapData{
				SwapIndex:   swapIndex,
				AmountIndex: amountIndex,

				IsFlashswap:  isFlashswap,
				V3SwapsDone:  v3SwapsDone,
				FeeIndex:     feeIndex,
				FactoryIndex: factoryIndex,

				Token0Index:    Token0Index,
				Token1Index:    Token1Index,
				FactoryAddress: factoryAddress,
			}
			v3SwapDatas = append(v3SwapDatas, v3SwapData)
		}
	}
	//reverse v3SwapDatas
	for i, j := 0, len(v3SwapDatas)-1; i < j; i, j = i+1, j-1 {
		v3SwapDatas[i], v3SwapDatas[j] = v3SwapDatas[j], v3SwapDatas[i]
	}

	return v3SwapDatas, nil
}

func getCompressedAmountsAndIndexMap(path logic.Path, amountOuts []*big.Int, amountIn *big.Int) ([]*big.Int, map[int]int, error) {
	//get uncompressed amounts
	//remove duplicates
	//make sure amountIn is at the first index
	//create bigInt string to compressed index map

	uncompressedAmounts := []*big.Int{}
	for i := 0; i < len(path.Pools); i++ {
		if path.Pools[i].GetType() == "uniswap_v2" {
			uncompressedAmounts = append(uncompressedAmounts, amountOuts[i])
		} else if path.Pools[i].GetType() == "uniswap_v3" {
			if i == 0 { //if first pool, set amount[i] to amountIn
				uncompressedAmounts = append(uncompressedAmounts, amountIn)
			} else {
				uncompressedAmounts = append(uncompressedAmounts, amountOuts[i-1])
			}
		}
	}

	//remove duplicates
	uniqueAmounts := []*big.Int{}
	uniqueAmountsMap := map[string]bool{}
	for i := 0; i < len(uncompressedAmounts); i++ {
		uniqueAmountsMap[uncompressedAmounts[i].String()] = true
	}
	for key, _ := range uniqueAmountsMap {
		uniqueAmount, ok := new(big.Int).SetString(key, 10)
		if !ok {
			return nil, nil, errors.New("could not convert string to big.Int")
		}
		uniqueAmounts = append(uniqueAmounts, uniqueAmount)
	}

	//make sure amountIn is at the first index
	amountInIndex := -1
	for i := 0; i < len(uniqueAmounts); i++ {
		if uniqueAmounts[i].Cmp(amountIn) == 0 {
			uniqueAmounts[0], uniqueAmounts[i] = uniqueAmounts[i], uniqueAmounts[0]
		}
	}
	if amountInIndex == -1 {
		uniqueAmounts = append([]*big.Int{amountIn}, uniqueAmounts...)
	}

	//create mapping from uncompressed index to compressed index
	amountsIndexMap := map[int]int{}
	for i := 0; i < len(uncompressedAmounts); i++ {
		for j := 0; j < len(uniqueAmounts); j++ {
			if uncompressedAmounts[i].Cmp(uniqueAmounts[j]) == 0 {
				amountsIndexMap[i] = j
			}
		}
	}

	return uniqueAmounts, amountsIndexMap, nil
}
