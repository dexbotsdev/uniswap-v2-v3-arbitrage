package executor

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"errors"
	"fmt"
	"math/big"
	"mev-template-go/pool_interface"
	"mev-template-go/types"
	"mev-template-go/uniswap_v2"

	//"github.com/umbracle/ethgo/abi"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getFeeIndex(fee int) (int, error) {
	feeIndex := 0

	if fee == 500 {
		feeIndex = 0
	} else if fee == 3000 {
		feeIndex = 1
	} else if fee == 10000 {
		feeIndex = 2
	} else if fee == 100 {
		feeIndex = 3
	} else {
		return 0, errors.New("invalid fee")
	}

	return feeIndex, nil
}

func getOtherTokenFromPool(pool pool_interface.PoolInterface, knownToken common.Address) common.Address {
	tokens := pool.GetTokens()
	var otherTokenAddress common.Address
	if tokens[0] == knownToken {
		otherTokenAddress = tokens[1]
	} else {
		otherTokenAddress = tokens[0]
	}
	return otherTokenAddress
}

//convert bigint to format: byte for length ... bytes for value
func intToCompressedBytes(bigInt *big.Int) []byte {
	fmt.Println(bigInt)
	intByteSize := len(bigInt.Bytes())
	intByteSizeByte := []byte{byte(intByteSize)}

	intBytes := padBytes(bigInt.Bytes(), intByteSize)
	bytesArray := []byte{}
	bytesArray = append(bytesArray, intByteSizeByte...)
	bytesArray = append(bytesArray, intBytes...)

	return bytesArray
}

//pads bytes to a certain length
func padBytes(input []byte, length int) []byte {
	if len(input) >= length {
		return input
	}

	padded := make([]byte, length)
	copy(padded[length-len(input):], input)

	return padded
}

func getGasOfCalldata(data []byte) uint64 {
	//for each 0 byte, add 4 gas, for each non-0 byte, add 16 gas
	gas := uint64(0)
	for i := 0; i < len(data); i++ {
		if data[i] == 0 {
			gas += 4
		} else {
			gas += 16
		}
	}
	return gas
}

//function that takes an aribtrary amount of big.Ints and returns the max number of bytes to represent the largest number
func getMaxIntByteSize(amounts ...*big.Int) int {
	max := new(big.Int).Set(amounts[0])
	for i := 1; i < len(amounts); i++ {
		if max.Cmp(amounts[i]) == -1 {
			max = new(big.Int).Set(amounts[i])
		}
	}
	return len(max.Bytes())
}

//function that takes amountOuts0 and amountOuts1 and returns the amountOutIndexes as a byte with the leftmost bit being the first index
func getAmountOutIndexes(amountOuts0, amountOuts1 []*big.Int) byte {
	amountOutIndexes := byte(0)
	for i := 0; i < len(amountOuts0); i++ {
		if amountOuts0[i].Cmp(amountOuts1[i]) == -1 {
			amountOutIndexes |= 1 << uint(i)
		}
	}
	return amountOutIndexes
}

func UpdatePools(pools []pool_interface.PoolInterface, client *ethclient.Client) error {
	for i := 0; i < len(pools); i++ {
		pools[i].Update(client)
	}
	return nil
}

func UpdateReservesOfPath(path *types.Path, client *ethclient.Client) {
	fmt.Println("updated reserves of path: ")
	for i := 0; i < len(path.Pools); i++ {
		//get reserves

		fmt.Println("pool: ", i)
		fmt.Println("address: ", path.Pools[i].Address)
		fmt.Println("token0: ", path.Pools[i].Token0)
		fmt.Println("token1: ", path.Pools[i].Token1)
		fmt.Println("old reserve0: ", path.Pools[i].Reserve0.String())
		fmt.Println("old reserve1: ", path.Pools[i].Reserve1.String())

		reserve0, reserve1, err := uniswap_v2.GetReserves(path.Pools[i].Address, client)
		if err != nil {
			fmt.Println("error getting reserves: ", err)
		}
		path.Pools[i].Reserve0 = reserve0
		path.Pools[i].Reserve1 = reserve1

		fmt.Println("new reserve0: ", path.Pools[i].Reserve0.String())
		fmt.Println("new reserve1: ", path.Pools[i].Reserve1.String())

	}
}
