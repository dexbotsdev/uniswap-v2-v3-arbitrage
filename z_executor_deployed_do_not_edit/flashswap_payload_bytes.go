package executor

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

// type PathExecutionValues struct {
// 	//For V2 Flashswap Only
// 	IsV2Flashswap         bool           //1 bit
// 	V2FlashswapTarget     common.Address //20 bytes
// 	V2FlashswapZeroForOne bool           //1 bit
// 	V2FlashswapAmountOut  *big.Int       //IntByteSize bytes
// 	//For V2 Callback Verification Payload Values
// 	V2CallbackV3SwapsDone     bool           //1 bit
// 	V2CallbackFactoryIndex    int            //3 bits
// 	V2CallbackOtherTokenIndex int            //3 bits
// 	V2CallbackFactoryAddress  common.Address //20 bytes

// 	//Payload Values
// 	//v3SwapsCompleted bool //first bit of payload, only set on chain
// 	IntByteSize int //6 bits; max value is 32

// 	SwapsLength   int //4 bits; max value is 8
// 	AmountsLength int //4 bits; max value is 8

// 	CoinBaseTransferBool bool //1 bit
// 	V2SwapDataLength     int  //3 bits; max value is 7
// 	TokensLength         int  //3 bits; max value is 7 since weth is not included

// 	Revenue                         *big.Int //IntByteSize bytes
// 	QuarterBribePercentageOfRevenue int      //8 bits

// 	Tokens      []common.Address //20 bytes each; length = SwapsLength -1
// 	Targets     []common.Address //20 bytes each
// 	Amounts     []*big.Int       //Compressed Int each, amountOut for v2, amountIn for v3
// 	ZeroForOnes []bool           //1 bit each

// 	//V3 Swap Datas
// 	V3SwapDatas []V3SwapData //max value is 7 as first swap is the flashswap

// 	//V2 Swap Datas
// 	V2SwapDatas []V2SwapData //max value is 7 as first swap is the flashswap

// 	CallTarget common.Address //20 bytes
// }

type FlashswapPayloadBytes struct {
	//For V2 Callback Verification Payload Values
	V2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes []byte //1;3;3 bits
	V2CallbackFactoryAddressBytes                               []byte //20 bytes

	//Payload Values; startPos
	IntByteSizeBytes []byte //6 bits

	SwapsLengthAndAmountsLengthBytes []byte //4;4 bits

	CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes []byte //1;3;3 bits

	RevenueBytes []byte //IntByteSize bytes

	BribePercentageOfRevenueMinus745Bytes []byte //8 bits; only if coinBaseTransferBool is true

	TokensBytes      []byte //20 bytes each; pos = startPos + 4 + IntByteSize
	TargetsBytes     []byte //20 bytes each; pos = startPos + 4 + IntByteSize + TokensLength*20
	AmountsBytes     []byte //Compressed Int each, amountOut for v2, amountIn for v3; pos = 4 + IntByteSize + SwapsLength*20 + TokensLength*20
	ZeroForOnesBytes []byte //1 bit each; pos = 4 + IntByteSize + SwapsLength*20 + AmountsLength*IntByteSize + TokensLength*20

	//V3 Swap Datas; pos = 4 + IntByteSize + SwapsLength*20 + AmountsLength*IntByteSize + 1
	V3SwapDataBytesArr []V3SwapDataBytes //max value is 7 as first swap is the flashswap

	//V2 Swap Datas
	V2SwapDataBytesArr []V2SwapDataBytes //max value is 7 as first swap is the flashswap
}

type V3SwapDataBytes struct {
	SwapIndexAndAmountIndexBytes []byte //3;3 bits

	IsFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes []byte //1;1;2;3 bits

	Token0IndexAndToken1IndexBytes []byte //3;3 bits

	FactoryAddressBytes []byte
}

type V2SwapDataBytes struct {
	SwapIndexAndAmountIndexBytes []byte //3;3 bits
}

func convertPathExecutionValuesToFlashswapPayloadBytes(pathExecutionValues PathExecutionValues) (FlashswapPayloadBytes, error) {

	// V2FactoryIndexBytes               []byte //3 bits
	// V2CallbackFactoryAddressBytes    []byte //20 bytes
	// V2CallbackOtherTokenAddressBytes []byte //20 bytes

	v2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes := []byte{}
	v2CallbackFactoryAddressBytes := []byte{}

	if pathExecutionValues.IsV2Flashswap {
		// V2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes []byte //1;3;3 bits
		// V2CallbackFactoryAddressBytes                               []byte //20 bytes
		v2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexValue := pathExecutionValues.V2CallbackFactoryIndex*8 + pathExecutionValues.V2CallbackOtherTokenIndex
		if pathExecutionValues.V2CallbackV3SwapsDone {
			v2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexValue += 64
		}
		v2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes = []byte{byte(v2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexValue)}

		if pathExecutionValues.V2CallbackFactoryIndex == 1 {
			v2CallbackFactoryAddressBytes = pathExecutionValues.V2CallbackFactoryAddress.Bytes()
		}
	}

	// CoinbaseTransferAndIntByteSizeBytes []byte //1;1;6 bits
	// V2SwapDataLengthBytes []byte //3 bits
	// targetsLengthAndAmountsLengthBytes []byte //4;4 bits
	// RevenueBytes                         []byte //IntByteSize bytes
	// QuarterBribePercentageOfRevenueBytes []byte //8 bits
	// targetsBytes []byte //20 bytes each
	// AmountsBytes       []byte //Compressed Int each, amountOut for v2, amountIn for v3
	// ZeroForOnesBytes   []byte //1 bit each
	// V3SwapDatas []V3SwapData //max value is 7 as first swap is the flashswap
	// V2SwapDatas []V2SwapData //max value is 7 as first swap is the flashswap

	intByteSizeBytes := []byte{byte(pathExecutionValues.IntByteSize)}
	intByteSize := pathExecutionValues.IntByteSize

	//SwapsLengthAndAmountsLengthBytes
	swapsLengthAndAmountsLengthBytes := []byte{byte(pathExecutionValues.SwapsLength*16 + pathExecutionValues.AmountsLength)}

	//CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes []byte //1;3;3 bits
	coinbaseTransferAndV2SwapDataLengthAndTokensLengthValue := 0
	if pathExecutionValues.CoinBaseTransferBool {
		coinbaseTransferAndV2SwapDataLengthAndTokensLengthValue += 64
	}
	coinbaseTransferAndV2SwapDataLengthAndTokensLengthValue += pathExecutionValues.V2SwapDataLength*8 + pathExecutionValues.TokensLength
	coinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes := []byte{byte(coinbaseTransferAndV2SwapDataLengthAndTokensLengthValue)}

	//RevenueBytes
	revenueBytes := padBytes(pathExecutionValues.Revenue.Bytes(), intByteSize)

	bribePercentageOfRevenueMinus745Bytes := []byte{byte(pathExecutionValues.BribePercentageOfRevenueMinus745)}

	//tokensBytes
	tokensBytes := []byte{}
	for _, token := range pathExecutionValues.Tokens {
		tokensBytes = append(tokensBytes, token.Bytes()...)
	}

	//targetsBytes
	targetsBytes := []byte{}
	for _, target := range pathExecutionValues.Targets {
		targetsBytes = append(targetsBytes, target.Bytes()...)
	}

	//AmountsBytes
	amountsBytes := []byte{}
	for _, amount := range pathExecutionValues.Amounts {
		amountsBytes = append(amountsBytes, padBytes(amount.Bytes(), intByteSize)...)
	}

	//zerosForOnes
	v2ZerosForOnesByte := byte(0)
	for i, zeroForOne := range pathExecutionValues.ZeroForOnes {
		if zeroForOne {
			v2ZerosForOnesByte += 1 << i
		}
	}

	//v3

	// type V3SwapData struct {
	// 	LastV3Call bool
	// 	FeeIndex   int
	// 	Token0     common.Address
	// 	Token1     common.Address
	// 	AmountIn   *big.Int
	// 	Target     common.Address
	// 	Recipient  common.Address
	// }
	//v3 swap count
	//v3 swap datas
	v3SwapDataBytesArr, err := getV3SwapDataBytesArr(pathExecutionValues)
	if err != nil {
		return FlashswapPayloadBytes{}, err
	}

	v2SwapDataBytesArr, err := getV2SwapDataBytesArr(pathExecutionValues)
	if err != nil {
		return FlashswapPayloadBytes{}, err
	}

	// //For V2 Callback Verification Payload Values
	// V2FactoryIndexBytes               []byte //3 bits
	// V2CallbackFactoryAddressBytes    []byte //20 bytes
	// V2CallbackOtherTokenAddressBytes []byte //20 bytes

	// //Payload Values
	// IntByteSizeBytes []byte //6 bits

	// CoinbaseTransferAndV2SwapDataLengthBytes []byte //1;1;3 bits

	// targetsLengthAndAmountsLengthBytes []byte //4;4 bits

	// RevenueBytes []byte //IntByteSize bytes

	// QuarterBribePercentageOfRevenueBytes []byte //8 bits

	// targetsBytes []byte //20 bytes each
	// AmountsBytes       []byte //Compressed Int each, amountOut for v2, amountIn for v3
	// ZeroForOnesBytes   []byte //1 bit each

	// //V3 Swap Datas
	// V3SwapDataBytesArr []V3SwapData //max value is 7 as first swap is the flashswap

	// //V2 Swap Datas
	// V2SwapDataBytesArr []V2SwapData //max value is 7 as first swap is the flashswap

	flashswapPayloadBytes := FlashswapPayloadBytes{
		V2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes: v2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes,
		V2CallbackFactoryAddressBytes:                               v2CallbackFactoryAddressBytes,
		IntByteSizeBytes:                                            intByteSizeBytes,
		SwapsLengthAndAmountsLengthBytes:                            swapsLengthAndAmountsLengthBytes,
		CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes:     coinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes,
		RevenueBytes:                          revenueBytes,
		BribePercentageOfRevenueMinus745Bytes: bribePercentageOfRevenueMinus745Bytes,
		TokensBytes:                           tokensBytes,
		TargetsBytes:                          targetsBytes,
		AmountsBytes:                          amountsBytes,
		ZeroForOnesBytes:                      []byte{v2ZerosForOnesByte},
		V3SwapDataBytesArr:                    v3SwapDataBytesArr,
		V2SwapDataBytesArr:                    v2SwapDataBytesArr,
	}

	return flashswapPayloadBytes, nil
}

// type V3SwapDataBytes struct {
// 	IsLastV3SwapAndSwapIndexAndAmountIndexBytes []byte //1;3;3 bits

// 	IsToken0WethAndIsToken1WethAndFeeIndexAndFactoryIndexBytes []byte //1;1;2;3 bits

// 	FactoryAddressBytes []byte
// 	Token0Bytes         []byte
// 	Token1Bytes         []byte
// }

// type V2SwapDataBytes struct {
// 	SwapIndexAndAmountIndexBytes []byte //3;3 bits
// }

func printFlashswapPayloadBytes(flashswapPayloadBytes FlashswapPayloadBytes) {
	fmt.Println("***FLASHSWAP PAYLOAD BYTES:***")
	fmt.Println("V2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes:", hex.EncodeToString(flashswapPayloadBytes.V2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes))
	fmt.Println("V2CallbackFactoryAddressBytes:", hex.EncodeToString(flashswapPayloadBytes.V2CallbackFactoryAddressBytes))
	fmt.Println("IntByteSizeBytes:", hex.EncodeToString(flashswapPayloadBytes.IntByteSizeBytes))
	fmt.Println("TargetsLengthAndAmountsLengthBytes:", hex.EncodeToString(flashswapPayloadBytes.SwapsLengthAndAmountsLengthBytes))
	fmt.Println("CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes:", hex.EncodeToString(flashswapPayloadBytes.CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes))
	fmt.Println("RevenueBytes:", hex.EncodeToString(flashswapPayloadBytes.RevenueBytes))
	fmt.Println("BribePercentageOfRevenueMinus745Bytes:", hex.EncodeToString(flashswapPayloadBytes.BribePercentageOfRevenueMinus745Bytes))
	fmt.Println("TokensBytes:", hex.EncodeToString(flashswapPayloadBytes.TokensBytes))
	fmt.Println("TargetsBytes:", hex.EncodeToString(flashswapPayloadBytes.TargetsBytes))
	fmt.Println("AmountsBytes:", hex.EncodeToString(flashswapPayloadBytes.AmountsBytes))
	fmt.Println("ZeroForOnesBytes:", hex.EncodeToString(flashswapPayloadBytes.ZeroForOnesBytes))
	// type V3SwapDataBytes struct {
	// 	V3SwapsRemainingAndTargetsIndexAndFeeIndexBytes []byte //3;3;2 bits

	// 	IsToken0WethAndIsToken1WethAndIsIndexedFactoryAndFactoryIndexAndAmountIndexBytes []byte //1;1;1;2;3 bits

	// 	FactoryAddressBytes []byte
	// 	Token0Bytes         []byte
	// 	Token1Bytes         []byte
	// }
	fmt.Println("V3SwapDataBytesArr:")
	for _, v3SwapDataBytes := range flashswapPayloadBytes.V3SwapDataBytesArr {
		fmt.Println("IsLastV3SwapAndSwapIndexAndAmountIndexBytes:", hex.EncodeToString(v3SwapDataBytes.SwapIndexAndAmountIndexBytes))
		fmt.Println("IsFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes:", hex.EncodeToString(v3SwapDataBytes.IsFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes))
		fmt.Println("Token0IndexAndToken1IndexBytes:", hex.EncodeToString(v3SwapDataBytes.Token0IndexAndToken1IndexBytes))
		fmt.Println("FactoryAddressBytes:", hex.EncodeToString(v3SwapDataBytes.FactoryAddressBytes))

	}
	// type V2SwapDataBytes struct {
	// 	TargetsIndexAndAmountIndexBytes []byte //3;3 bits
	// }
	fmt.Println("V2SwapDataBytesArr:")
	for _, v2SwapDataBytes := range flashswapPayloadBytes.V2SwapDataBytesArr {
		fmt.Println("TargetsIndexAndAmountIndexBytes:", hex.EncodeToString(v2SwapDataBytes.SwapIndexAndAmountIndexBytes))
	}
}

func appendFlashswapPayloadBytesToBytesArray(flashswapPayloadBytes FlashswapPayloadBytes) []byte {
	bytesArray := []byte{}

	bytesArray = append(bytesArray, flashswapPayloadBytes.V2CallbackV3SwapsDoneAndFactoryIndexAndOtherTokenIndexBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.V2CallbackFactoryAddressBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.IntByteSizeBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.SwapsLengthAndAmountsLengthBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.CoinbaseTransferAndV2SwapDataLengthAndTokensLengthBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.RevenueBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.BribePercentageOfRevenueMinus745Bytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.TokensBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.TargetsBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.AmountsBytes...)
	bytesArray = append(bytesArray, flashswapPayloadBytes.ZeroForOnesBytes...)
	for _, v3SwapDataBytes := range flashswapPayloadBytes.V3SwapDataBytesArr {
		bytesArray = append(bytesArray, v3SwapDataBytes.SwapIndexAndAmountIndexBytes...)
		bytesArray = append(bytesArray, v3SwapDataBytes.IsFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes...)
		bytesArray = append(bytesArray, v3SwapDataBytes.Token0IndexAndToken1IndexBytes...)
		bytesArray = append(bytesArray, v3SwapDataBytes.FactoryAddressBytes...)
	}
	for _, v2SwapDataBytes := range flashswapPayloadBytes.V2SwapDataBytesArr {
		bytesArray = append(bytesArray, v2SwapDataBytes.SwapIndexAndAmountIndexBytes...)
	}
	return bytesArray
}
