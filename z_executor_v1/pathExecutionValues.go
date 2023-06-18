package executorBackup

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"errors"
	"fmt"
	"math/big"
	"mev-template-go/logic"
	"mev-template-go/uniswap_v3"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/umbracle/ethgo/abi"
)

//Load
//shift
//mul

//9+16

//

type PathExecutionValues struct {
	//for v3 only
	FlashswapAmountIn *big.Int       //IntByteSize bytes
	ExecutorAddress   common.Address //20 bytes

	//for v2 only
	FlashswapTarget            common.Address //20 bytes
	FlashswapZeroForOne        bool           //1 bit
	FlashswapAmountOut         *big.Int       //IntByteSize bytes
	FlashswapIsV2              bool
	FlashswapFactoryAddress    common.Address //20 bytes
	FlashswapOtherTokenAddress common.Address //20 bytes

	FlashswapFeeIndex    uint //2 bits each, 2 for each pool; used for v3 callback verification
	CoinBaseTransferBool bool //1 bit
	swapDataCountMinus1  int  //3 bits; max value is 8

	IntByteSizeMinus1               int        //5 bits; max value is 32
	Revenue                         *big.Int   //IntByteSize bytes
	QuarterBribePercentageOfRevenue int        //8 bits
	SwapDataIsV2s                   []bool     //1 byte; bit each, 1 for each pool
	SwapDataZerosForOnes            []bool     //1 byte; 1 bit each, 1 for each pool
	SwapDatas                       []SwapData //max value is 7 as first swap is the flashswap
}
type SwapData struct {
	IsV2      bool
	Target    common.Address
	AmountOut *big.Int //IntByteSize bytes; only used for v2
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

	//flashswap target
	flashswapTarget := path.Pools[len(path.Pools)-1].GetAddress()

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

	//get other token address of last pool
	tokens := path.Pools[len(path.Pools)-1].GetTokens()
	var flashswapOtherTokenAddress common.Address
	if tokens[0] == path.BaseToken {
		flashswapOtherTokenAddress = tokens[1]
	} else {
		flashswapOtherTokenAddress = tokens[0]
	}

	//get flashswapIsV2
	flashswapIsV2 := path.Pools[len(path.Pools)-1].GetType() == "uniswap_v2"

	//get flashswapZeroForOne
	flashswapZeroForOne := path.BaseToken == path.Pools[len(path.Pools)-1].GetTokens()[1]

	//get flashswapAmountIn
	flashswapAmountIn := new(big.Int).Set(amountOuts[len(amountOuts)-2])

	//get flashswapAmountOut
	flashswapAmountOut := new(big.Int).Set(amountOuts[len(amountOuts)-1])

	fmt.Println("flashswapTarget: ", flashswapTarget.String())
	fmt.Println("flashswapAmountIn: ", flashswapAmountIn.String())
	fmt.Println("flashswapAmountOut: ", flashswapAmountOut.String())

	for i := 0; i < len(amountOuts); i++ {
		fmt.Println("amountOuts[i]: ", amountOuts[i].String())
	}

	//get fee of last pool
	feeIndex := uint(0)
	if path.Pools[len(path.Pools)-1].GetType() == "uniswap_v3" {
		fee := path.Pools[len(path.Pools)-1].(*uniswap_v3.Pool).Fee
		if fee == 500 {
			feeIndex = 0
		} else if fee == 3000 {
			feeIndex = 1
		} else if fee == 10000 {
			feeIndex = 2
		} else {
			return PathExecutionValues{}, errors.New("invalid fee")
		}
	}

	//get coinbase transfer bool
	//TODO implement coinbase transfer bool
	coinBaseTransferBool := true

	//get max int byte size
	IntByteSize := 0
	for _, amountOut := range amountOuts {
		if len(amountOut.Bytes()) > IntByteSize {
			IntByteSize = len(amountOut.Bytes())
		}
	}
	if len(path.Revenue.Bytes()) > IntByteSize {
		IntByteSize = len(path.Revenue.Bytes())
	}

	//get quarter bribe percentage of revenue
	//set temporary bribe of revenue
	//1000 is 100%
	quarterBribePercentageOfRevenue := 0

	//make swapDatas
	swapDatas := make([]SwapData, len(path.Pools)-1)
	for i := 0; i < len(path.Pools)-1; i++ {
		swapData := SwapData{
			IsV2:      path.Pools[i].GetType() == "uniswap_v2",
			Target:    path.Pools[i].GetAddress(),
			AmountOut: new(big.Int).Set(amountOuts[i]),
		}
		swapDatas[i] = swapData
	}

	//make payload values
	payloadValues := PathExecutionValues{
		FlashswapTarget:                 flashswapTarget,
		ExecutorAddress:                 executorAddress,
		FlashswapIsV2:                   flashswapIsV2,
		FlashswapZeroForOne:             flashswapZeroForOne,
		FlashswapAmountIn:               new(big.Int).Set(flashswapAmountIn),
		FlashswapAmountOut:              new(big.Int).Set(flashswapAmountOut),
		FlashswapFactoryAddress:         path.Pools[len(path.Pools)-1].GetFactoryAddress(),
		FlashswapOtherTokenAddress:      flashswapOtherTokenAddress,
		FlashswapFeeIndex:               feeIndex,
		CoinBaseTransferBool:            coinBaseTransferBool,
		swapDataCountMinus1:             len(swapDatas) - 1,
		IntByteSizeMinus1:               IntByteSize - 1,
		Revenue:                         new(big.Int).Set(path.Revenue),
		QuarterBribePercentageOfRevenue: quarterBribePercentageOfRevenue,
		SwapDataIsV2s:                   path.IsV2s[0 : len(path.IsV2s)-1],
		SwapDataZerosForOnes:            path.ZeroForOnes[0 : len(path.ZeroForOnes)-1],
		SwapDatas:                       swapDatas,
	}

	return payloadValues, nil
}

func PrintPathExecutionValues(p PathExecutionValues) {
	fmt.Println("PATH EXECUTION VALUES:")
	fmt.Println("FlashswapTarget: ", p.FlashswapTarget)
	fmt.Println("ExecutorAddress: ", p.ExecutorAddress)
	fmt.Println("FlashswapIsV2: ", p.FlashswapIsV2)
	fmt.Println("FlashswapZeroForOne: ", p.FlashswapZeroForOne)
	fmt.Println("FlashswapAmountIn: ", p.FlashswapAmountIn.String())
	fmt.Println("FlashswapAmountOut: ", p.FlashswapAmountOut.String())
	fmt.Println("FlashswapFactoryAddress: ", p.FlashswapFactoryAddress)
	fmt.Println("FlashswapOtherTokenAddress: ", p.FlashswapOtherTokenAddress)
	fmt.Println("FlashswapFeeIndex: ", p.FlashswapFeeIndex)
	fmt.Println("CoinBaseTransferBool: ", p.CoinBaseTransferBool)
	fmt.Println("swapDataCountMinus1: ", p.swapDataCountMinus1)
	fmt.Println("IntByteSizeMinus1: ", p.IntByteSizeMinus1)
	fmt.Println("Revenue: ", p.Revenue.String())
	fmt.Println("QuarterBribePercentageOfRevenue: ", p.QuarterBribePercentageOfRevenue)
	fmt.Println("SwapDataIsV2s: ", p.SwapDataIsV2s)
	fmt.Println("SwapDataZerosForOnes: ", p.SwapDataZerosForOnes)
	for i, swapData := range p.SwapDatas {
		fmt.Println("SwapData", i, ":")
		fmt.Println("IsV2: ", swapData.IsV2)
		fmt.Println("Target: ", swapData.Target)
		fmt.Println("AmountOut: ", swapData.AmountOut.String())
	}
}
