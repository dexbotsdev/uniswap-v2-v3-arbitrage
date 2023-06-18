package executor

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"errors"
	"fmt"
	"math/big"
	"mev-template-go/logic"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/umbracle/ethgo/abi"
)

type PathExecutionValues struct {
	//For V2 Flashswap Only
	IsV2Flashswap         bool           //1 bit
	V2FlashswapTarget     common.Address //20 bytes
	V2FlashswapZeroForOne bool           //1 bit
	V2FlashswapAmountOut  *big.Int       //IntByteSize bytes
	//For V2 Callback Verification Payload Values
	V2CallbackV3SwapsDone     bool           //1 bit
	V2CallbackFactoryIndex    int            //3 bits
	V2CallbackOtherTokenIndex int            //3 bits
	V2CallbackFactoryAddress  common.Address //20 bytes

	//Payload Values
	//v3SwapsCompleted bool //first bit of payload, only set on chain
	IntByteSize int //6 bits; max value is 32

	SwapsLength   int //4 bits; max value is 8
	AmountsLength int //4 bits; max value is 8

	CoinBaseTransferBool bool //1 bit
	V2SwapDataLength     int  //3 bits; max value is 7
	TokensLength         int  //3 bits; max value is 7 since weth is not included

	Revenue                          *big.Int //IntByteSize bytes
	BribePercentageOfRevenueMinus745 int      //8 bits

	Tokens      []common.Address //20 bytes each; length = SwapsLength -1
	Targets     []common.Address //20 bytes each
	Amounts     []*big.Int       //Compressed Int each, amountOut for v2, amountIn for v3
	ZeroForOnes []bool           //1 bit each

	//V3 Swap Datas
	V3SwapDatas []V3SwapData //max value is 7 as first swap is the flashswap

	//V2 Swap Datas
	V2SwapDatas []V2SwapData //max value is 7 as first swap is the flashswap

	CallTarget common.Address //20 bytes
}

type V3SwapData struct {
	SwapIndex   int // 3 bits
	AmountIndex int //3 bits

	//for callback verification
	IsFlashswap  bool //1 bit
	V3SwapsDone  bool //1 bit
	FeeIndex     int  //2 bits
	FactoryIndex int  //3 bits

	Token0Index    int //3 bits
	Token1Index    int //3 bits
	FactoryAddress common.Address
}

type V2SwapData struct {
	SwapIndex   int //3 bits
	AmountIndex int //3 bits
}

func convertPathToPathExecutionValues(path logic.Path) (PathExecutionValues, error) {
	fmt.Println("***convertPathToPathExecutionValues***")

	if path.AmountIn == nil {
		return PathExecutionValues{}, errors.New("amountIn is nil")
	}
	if path.Revenue == nil {
		return PathExecutionValues{}, errors.New("revenue is nil")
	}
	if path.Revenue.Cmp(big.NewInt(0)) < 0 {
		return PathExecutionValues{}, errors.New("revenue is negative")
	}

	//calculate all amountOuts
	//TODO implement getting amountOuts for duplicate pools
	fmt.Println("amountIn: ", path.AmountIn.String())
	amountOuts := make([]*big.Int, len(path.Pools))
	currAmount := new(big.Int).Set(path.AmountIn)
	for i := 0; i < len(path.Pools); i++ {
		currAmountTemp, err := path.Pools[i].GetAmountOut(currAmount, path.ZeroForOnes[i])
		if err != nil {
			return PathExecutionValues{}, err
		}
		currAmount = new(big.Int).Set(currAmountTemp)
		amountOuts[i] = new(big.Int).Set(currAmount)
		fmt.Println("path.Pools[i].GetAddress(): ", path.Pools[i].GetAddress().String())
		fmt.Println("path.Pools[i].zeroForOne: ", path.ZeroForOnes[i])
		fmt.Println("path.Pools[i].GetType(): ", path.Pools[i].GetType())
		fmt.Println("amountOuts[i]: ", amountOuts[i].String())
	}

	//print out all amountOuts
	fmt.Println("***AMOUNT OUTS***")
	for i := 0; i < len(amountOuts); i++ {
		fmt.Println("amountOuts[i]: ", amountOuts[i].String())
	}

	tokens, tokenToIndexMap := compressTokens(path)

	//***V2 FLASHSWAP VALUES***

	// //For V2 Flashswap Only
	// IsV2Flashswap         bool           //1 bit
	// V2FlashswapTarget     common.Address //20 bytes
	// V2FlashswapZeroForOne bool           //1 bit
	// V2FlashswapAmountOut  *big.Int       //IntByteSize bytes
	// //For V2 Callback Verification Payload Values
	// V2IsFactoryIndexed           bool           //1 bit
	// V2FactoryIndex               int            //3 bits
	// V2FlashswapFactoryAddress    common.Address //20 bytes
	// V2FlashswapOtherTokenAddress common.Address //20 bytes

	var isV2Flashswap bool
	var v2FlashswapTarget common.Address
	var v2FlashswapZeroForOne bool
	var v2FlashswapAmountOut *big.Int
	var v2CallbackV3SwapsDone bool
	var v2CallbackFactoryIndex int
	var v2CallbackFactoryAddress common.Address
	var v2CallbackOtherTokenIndex int

	//check if last pool is v2 or v3
	if path.Pools[len(path.Pools)-1].GetType() == "uniswap_v2" {

		isV2Flashswap = true
		v2FlashswapTarget = path.Pools[len(path.Pools)-1].GetAddress()
		v2FlashswapZeroForOne = path.ZeroForOnes[len(path.ZeroForOnes)-1]
		v2FlashswapAmountOut = amountOuts[len(amountOuts)-1]
		v2CallbackV3SwapsDone = getV2CallbackV3SwapsDone(path)
		v2CallbackFactoryIndex = getV2FactoryIndex(path.Pools[len(path.Pools)-1].GetFactoryAddress())
		v2CallbackFactoryAddress = path.Pools[len(path.Pools)-1].GetFactoryAddress()
		v2CallbackOtherTokenIndex = tokenToIndexMap[getOtherTokenFromPool(path.Pools[len(path.Pools)-1], path.BaseToken)]

	}

	//***PAYLOAD VALUES***

	// //Payload Values
	// IntByteSize          int  //6 bits; max value is 32
	// CoinBaseTransferBool bool //1 bit
	// OwnerTransferBool    bool //1 bit

	// V2SwapDataLength int //3 bits; max value is 8

	// Length int //4 bits; max value is 8
	// AmountsLength       int //4 bits; max value is 8

	// Revenue                         *big.Int //IntByteSize bytes
	// QuarterBribePercentageOfRevenue int      //8 bits

	// PoolAddresses []common.Address //20 bytes each
	// amounts       []*big.Int       //Compressed Int each, amountOut for v2, amountIn for v3
	// zeroForOnes   []bool           //1 bit each

	//amounts
	compressedAmounts, amountsIndexMap, err := getCompressedAmountsAndIndexMap(path, amountOuts, path.AmountIn)
	if err != nil {
		return PathExecutionValues{}, err
	}

	//IntByteSize
	IntByteSize := 0
	for _, amountOut := range amountOuts {
		if len(amountOut.Bytes()) > IntByteSize {
			IntByteSize = len(amountOut.Bytes())
		}
	}
	if len(path.Revenue.Bytes()) > IntByteSize {
		IntByteSize = len(path.Revenue.Bytes())
	}
	if len(path.AmountIn.Bytes()) > IntByteSize {
		IntByteSize = len(path.AmountIn.Bytes())
	}

	//TargetsLength
	swapsLength := len(path.Pools)

	//AmountsLength
	amountsLength := len(compressedAmounts)

	//CoinBaseTransferBool
	//TODO implement coinbase transfer bool
	coinBaseTransferBool := true

	//V2SwapDataLength
	v2SwapDataLength := 0

	//tokensLength
	tokensLength := len(tokens)

	//revenue
	revenue := path.Revenue

	//get quarter bribe percentage of revenue
	//set temporary bribe of revenue
	//1000 is 100%
	bribePercentageOfRevenue := 999
	bribePercentageOfRevenueMinus745 := bribePercentageOfRevenue - 745

	//Targets
	targets := make([]common.Address, len(path.Pools))
	for i := 0; i < len(path.Pools); i++ {
		targets[i] = path.Pools[i].GetAddress()
	}

	//zeroForOnes
	zeroForOnes := path.ZeroForOnes

	//get v3 swap datas
	v3SwapDatas, err := getV3SwapDatas(path, targets, amountsIndexMap, tokenToIndexMap)
	if err != nil {
		return PathExecutionValues{}, err
	}

	//get v2 swap datas
	v2SwapDatas, err := getV2SwapDatas(path, amountsIndexMap)
	if err != nil {
		return PathExecutionValues{}, err
	}

	v2SwapDataLength = len(v2SwapDatas)

	//executorAddress := common.HexToAddress("0x5615dEB798BB3E4dFa0139dFa1b3D433Cc23b72f")
	executorAddress := common.HexToAddress("0x3d7C77070aEB2869C32976978141c1086F64452B")

	callTarget := executorAddress
	if isV2Flashswap {
		callTarget = targets[len(targets)-1]
	}
	// Target     common.Address
	// AmountOut  *big.Int //IntByteSize bytes; only used for v2
	// NextTarget common.Address
	//make swapDatas

	//make payload values
	pathExecutionValues := PathExecutionValues{
		//For V2 Flashswap Only
		IsV2Flashswap:         isV2Flashswap,
		V2FlashswapTarget:     v2FlashswapTarget,
		V2FlashswapZeroForOne: v2FlashswapZeroForOne,
		V2FlashswapAmountOut:  v2FlashswapAmountOut,
		//For V2 Callback Verification Payload Values
		V2CallbackV3SwapsDone:     v2CallbackV3SwapsDone,
		V2CallbackFactoryIndex:    v2CallbackFactoryIndex,
		V2CallbackFactoryAddress:  v2CallbackFactoryAddress,
		V2CallbackOtherTokenIndex: v2CallbackOtherTokenIndex,
		//Payload Values
		IntByteSize:                      IntByteSize,
		SwapsLength:                      swapsLength,
		AmountsLength:                    amountsLength,
		CoinBaseTransferBool:             coinBaseTransferBool,
		V2SwapDataLength:                 v2SwapDataLength,
		TokensLength:                     tokensLength,
		Revenue:                          revenue,
		BribePercentageOfRevenueMinus745: bribePercentageOfRevenueMinus745,
		Tokens:                           tokens,
		Targets:                          targets,
		Amounts:                          compressedAmounts,
		ZeroForOnes:                      zeroForOnes,
		V3SwapDatas:                      v3SwapDatas,
		V2SwapDatas:                      v2SwapDatas,
		CallTarget:                       callTarget,
	}

	return pathExecutionValues, nil
}

// type V3SwapData struct {
// 	IsLastV3Swap bool //1 bits
// 	SwapIndex    int  // 3 bits
// 	AmountIndex  int  //3 bits

// 	//for callback verification
// 	IsToken0Weth bool //1 bit
// 	IsToken1Weth bool //1 bit
// 	FeeIndex     int  //2 bits
// 	FactoryIndex int  //3 bits

// 	FactoryAddress common.Address
// 	Token0         common.Address
// 	Token1         common.Address
// }

// type V2SwapData struct {
// 	SwapIndex   int //3 bits
// 	AmountIndex int //3 bits
// }

func PrintPathExecutionValues(p PathExecutionValues) {
	fmt.Println("***PATH EXECUTION VALUES***")
	fmt.Println("IsV2Flashswap: ", p.IsV2Flashswap)
	fmt.Println("V2FlashswapTarget: ", p.V2FlashswapTarget.Hex())
	fmt.Println("V2FlashswapZeroForOne: ", p.V2FlashswapZeroForOne)
	fmt.Println("V2FlashswapAmountOut: ", p.V2FlashswapAmountOut)
	fmt.Println("V2CallbackV3SwapsDone: ", p.V2CallbackV3SwapsDone)
	fmt.Println("V2CallbackFactoryIndex: ", p.V2CallbackFactoryIndex)
	fmt.Println("V2CallbackFactoryAddress: ", p.V2CallbackFactoryAddress.Hex())
	fmt.Println("V2CallbackOtherTokenIndex: ", p.V2CallbackOtherTokenIndex)
	fmt.Println("IntByteSize: ", p.IntByteSize)
	fmt.Println("SwapsLength: ", p.SwapsLength)
	fmt.Println("AmountsLength: ", p.AmountsLength)
	fmt.Println("CoinBaseTransferBool: ", p.CoinBaseTransferBool)
	fmt.Println("V2SwapDataLength: ", p.V2SwapDataLength)
	fmt.Println("TokensLength: ", p.TokensLength)
	fmt.Println("Revenue: ", p.Revenue)
	fmt.Println("BribePercentageOfRevenueMinus745: ", p.BribePercentageOfRevenueMinus745)
	fmt.Println("Tokens: ", p.Tokens)
	fmt.Println("Targets: ", p.Targets)
	fmt.Println("Amounts: ", p.Amounts)
	fmt.Println("ZeroForOnes: ", p.ZeroForOnes)
	for i := 0; i < len(p.V3SwapDatas); i++ {
		fmt.Println("SwapIndex: ", p.V3SwapDatas[i].SwapIndex)
		fmt.Println("AmountIndex: ", p.V3SwapDatas[i].AmountIndex)
		fmt.Println("IsFlashswap: ", p.V3SwapDatas[i].IsFlashswap)
		fmt.Println("V3SwapsDone: ", p.V3SwapDatas[i].V3SwapsDone)
		fmt.Println("FeeIndex: ", p.V3SwapDatas[i].FeeIndex)
		fmt.Println("FactoryIndex: ", p.V3SwapDatas[i].FactoryIndex)
		fmt.Println("FactoryAddress: ", p.V3SwapDatas[i].FactoryAddress.Hex())
		fmt.Println("Token0Index: ", p.V3SwapDatas[i].Token0Index)
		fmt.Println("Token1Index: ", p.V3SwapDatas[i].Token1Index)
	}
	for i := 0; i < len(p.V2SwapDatas); i++ {
		fmt.Println("SwapIndex: ", p.V2SwapDatas[i].SwapIndex)
		fmt.Println("AmountIndex: ", p.V2SwapDatas[i].AmountIndex)
	}
	fmt.Println("***END PATH EXECUTION VALUES***")
}
