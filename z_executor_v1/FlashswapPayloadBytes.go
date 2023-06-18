package executorBackup

import (
	"encoding/hex"
	"fmt"
)

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

//"github.com/umbracle/ethgo/abi"

//modes
//0: v2 flashswap/ coinbase transfer
//1: v2 flashswap/ no coinbase transfer
//2: v3 flashswap/ coinbase transfer
//3: v3 flashswap/ no coinbase transfer

//flow only affects callback verification
//no verification
//v2 verification; no fee
//v3 verification; with fee

//PAYLOAD FORMAT
// startPos = 164 for v2, 196 for v3
// factoryAddress 20 bytes
// otherTokenAddress 20 bytes
// fee 2 bits; only used if flashswap is v3
// coinBaseTransferBool //1 bit
// swapDataCountMinus1 3 bits
// IntByteSize 5 bits
// revenue IntByteSize bytes
// quarterBribePercentageOfRevenue; 8 bits; only if coinBaseTransferBool is true
// isV2s; 1 byte; 1 bit for each swap
// zeroForOnes; 1 byte; 1 bit for each swap
// swapdatas

// SWAPDATA FORMAT
// target 20 bytes
// amountOut IntByteSize bytes; only if isV2 is true
type FlashswapPayloadBytes struct {
	FactoryAddress                          []byte          //20 bytes
	OtherTokenAddress                       []byte          //20 bytes
	FeeIndexAndCoinbaseTransferBool         []byte          //1 byte
	swapDataCountMinus1AndIntByteSizeMinus1 []byte          //3 bits; max value is 8
	Revenue                                 []byte          //IntByteSize bytes
	QuarterBribePercentageOfRevenue         []byte          //8 bits
	SwapDataIsV2s                           []byte          //1 byte; bit each, 1 for each pool
	SwapDataZerosForOnes                    []byte          //1 byte; 1 bit each, 1 for each pool
	SwapDatas                               []SwapDataBytes //max value is 7 as first swap is the flashswap
}
type SwapDataBytes struct {
	Target    []byte
	AmountOut []byte //IntByteSize bytes; only used for v2
}

func convertPathExecutionValuesToFlashswapPayloadBytes(pathExecutionValues PathExecutionValues) (FlashswapPayloadBytes, error) {

	flashSwapFactoryBytes := pathExecutionValues.FlashswapFactoryAddress.Bytes()

	flashswapOtherTokenBytes := pathExecutionValues.FlashswapOtherTokenAddress.Bytes()

	//combine fee and coinbase transfer Bool
	FeeIndexAndCoinbaseTransfer := pathExecutionValues.FlashswapFeeIndex
	if pathExecutionValues.CoinBaseTransferBool {
		FeeIndexAndCoinbaseTransfer += 4 //adds 1 in the 3rd bit position
	}
	FeeIndexAndCoinbaseTransferByte := []byte{byte(FeeIndexAndCoinbaseTransfer)}

	//swap count and max int byte size
	//Format: 3 bits for swap count, 5 bits for max int byte size, left to right
	swapDataCountMinus1AndIntByteSizeMinus1 := pathExecutionValues.swapDataCountMinus1 << 5
	swapDataCountMinus1AndIntByteSizeMinus1 += pathExecutionValues.IntByteSizeMinus1
	swapDataCountMinus1AndInteByteSizeMinus1Byte := []byte{byte(swapDataCountMinus1AndIntByteSizeMinus1)}

	intByteSize := pathExecutionValues.IntByteSizeMinus1 + 1

	//revenue of IntByteSize bytes
	revenueBytes := padBytes(pathExecutionValues.Revenue.Bytes(), intByteSize)

	//bribe percentage of revenue, make it 1 byte
	quarterBribePercentageOfRevenueByte := []byte{byte(pathExecutionValues.QuarterBribePercentageOfRevenue)}

	//isV2s
	swapDataIsV2sByte := byte(0)
	for i, isV2 := range pathExecutionValues.SwapDataIsV2s {
		if isV2 {
			swapDataIsV2sByte += 1 << i
		}
	}

	//zerosForOnes
	zerosForOnesByte := byte(0)
	for i, zeroForOne := range pathExecutionValues.SwapDataZerosForOnes {
		if zeroForOne {
			zerosForOnesByte += 1 << i
		}
	}

	//swapdatas
	swapDataBytesArr := make([]SwapDataBytes, len(pathExecutionValues.SwapDatas))
	for i, _ := range pathExecutionValues.SwapDatas {
		swapDataTargetBytes := pathExecutionValues.SwapDatas[i].Target.Bytes()
		swapDataAmountOutBytes := []byte{}

		if pathExecutionValues.SwapDataIsV2s[i] {
			swapDataAmountOutBytes = padBytes(pathExecutionValues.SwapDatas[i].AmountOut.Bytes(), intByteSize)
		}

		swapDataBytesArr[i] = SwapDataBytes{
			Target:    swapDataTargetBytes,
			AmountOut: swapDataAmountOutBytes,
		}
	}

	//add bytes to struct

	flashswapPayloadBytes := FlashswapPayloadBytes{
		FactoryAddress:                          flashSwapFactoryBytes,
		OtherTokenAddress:                       flashswapOtherTokenBytes,
		FeeIndexAndCoinbaseTransferBool:         FeeIndexAndCoinbaseTransferByte,
		swapDataCountMinus1AndIntByteSizeMinus1: swapDataCountMinus1AndInteByteSizeMinus1Byte,
		Revenue:                                 revenueBytes,
		QuarterBribePercentageOfRevenue:         quarterBribePercentageOfRevenueByte,
		SwapDataIsV2s:                           []byte{swapDataIsV2sByte},
		SwapDataZerosForOnes:                    []byte{zerosForOnesByte},
		SwapDatas:                               swapDataBytesArr,
	}

	return flashswapPayloadBytes, nil
}

func printFlashswapPayloadBytes(flashswapPayloadBytes FlashswapPayloadBytes) {
	fmt.Println("FLASHSWAP PAYLOAD BYTES:")
	fmt.Println("FactoryAddress length: ", len(flashswapPayloadBytes.FactoryAddress), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.FactoryAddress))
	fmt.Println("OtherTokenAddress length: ", len(flashswapPayloadBytes.OtherTokenAddress), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.OtherTokenAddress))
	fmt.Println("FeeIndexAndCoinbaseTransferBool length: ", len(flashswapPayloadBytes.FeeIndexAndCoinbaseTransferBool), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.FeeIndexAndCoinbaseTransferBool))
	fmt.Println("swapDataCountMinus1AndIntByteSizeMinus1 length: ", len(flashswapPayloadBytes.swapDataCountMinus1AndIntByteSizeMinus1), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.swapDataCountMinus1AndIntByteSizeMinus1))
	fmt.Println("Revenue length: ", len(flashswapPayloadBytes.Revenue), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.Revenue))
	fmt.Println("QuarterBribePercentageOfRevenue length: ", len(flashswapPayloadBytes.QuarterBribePercentageOfRevenue), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.QuarterBribePercentageOfRevenue))
	fmt.Println("SwapDataIsV2s length: ", len(flashswapPayloadBytes.SwapDataIsV2s), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.SwapDataIsV2s))
	fmt.Println("SwapDataZerosForOnes length: ", len(flashswapPayloadBytes.SwapDataZerosForOnes), ";bytes: ", hex.EncodeToString(flashswapPayloadBytes.SwapDataZerosForOnes))
	fmt.Println("SwapDatas length: ", len(flashswapPayloadBytes.SwapDatas))
	for i, swapData := range flashswapPayloadBytes.SwapDatas {
		fmt.Println("SwapData ", i, ":")
		fmt.Println("Target length: ", len(swapData.Target), ";bytes: ", hex.EncodeToString(swapData.Target))
		fmt.Println("AmountOut length: ", len(swapData.AmountOut), ";bytes: ", hex.EncodeToString(swapData.AmountOut))
	}
}

func appendFlashswapPayloadBytesToBytesArray(flashswapPayloadBytes FlashswapPayloadBytes) []byte {
	bytesArray := []byte{}

	bytesArray = append(bytesArray, flashswapPayloadBytes.FactoryAddress...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.OtherTokenAddress...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.FeeIndexAndCoinbaseTransferBool...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.swapDataCountMinus1AndIntByteSizeMinus1...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.Revenue...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.QuarterBribePercentageOfRevenue...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.SwapDataIsV2s...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.SwapDataZerosForOnes...)

	for _, swapData := range flashswapPayloadBytes.SwapDatas {
		bytesArray = append(bytesArray, swapData.Target...)
		bytesArray = append(bytesArray, swapData.AmountOut...)
	}

	return bytesArray
}
