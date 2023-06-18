package executor

// type V3SwapDataBytes struct {
// 	SwapIndexAndAmountIndexBytes []byte //3;3 bits

// 	V3SwapsDoneAndFeeIndexAndFactoryIndexBytes []byte //1;2;3 bits

// 	Token0IndexAndToken1IndexBytes []byte //3;3 bits

// 	FactoryAddressBytes []byte
// }

// type V2SwapDataBytes struct {
// 	SwapIndexAndAmountIndexBytes []byte //3;3 bits
// }

// function getV2Factory(uint index) internal returns (address){
// 	if (index ==0){
// 			return 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f;
// 	}
// }
// function getV3Factory(uint index) internal returns (address){
// 	if (index ==0){
// 			return 0x1F98431c8aD98523631AE4a59f267346ea31F984;
// 	}
// }

func getV3SwapDataBytesArr(pathExecutionValues PathExecutionValues) ([]V3SwapDataBytes, error) {
	v3SwapDataBytesArr := make([]V3SwapDataBytes, len(pathExecutionValues.V3SwapDatas))
	for i := 0; i < len(pathExecutionValues.V3SwapDatas); i++ {
		//SwapIndexAndAmountIndexBytes
		swapIndexAndAmountIndexValue := pathExecutionValues.V3SwapDatas[i].SwapIndex*8 + pathExecutionValues.V3SwapDatas[i].AmountIndex
		swapIndexAndAmountIndexBytes := []byte{byte(swapIndexAndAmountIndexValue)}

		//IsFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes []byte //1;1;2;3 bits
		isFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexValue := 0
		if pathExecutionValues.V3SwapDatas[i].IsFlashswap {
			isFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexValue += 64
		}
		if pathExecutionValues.V3SwapDatas[i].V3SwapsDone {
			isFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexValue += 32
		}
		isFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexValue += +pathExecutionValues.V3SwapDatas[i].FeeIndex*8 + pathExecutionValues.V3SwapDatas[i].FactoryIndex
		isFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes := []byte{byte(isFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexValue)}

		//Token0IndexAndToken1IndexBytes []byte //3;3 bits
		token0IndexAndToken1IndexValue := pathExecutionValues.V3SwapDatas[i].Token0Index*8 + pathExecutionValues.V3SwapDatas[i].Token1Index
		token0IndexAndToken1IndexBytes := []byte{byte(token0IndexAndToken1IndexValue)}

		//FactoryAddressBytes
		factoryAddressBytes := []byte{}
		if pathExecutionValues.V3SwapDatas[i].FactoryIndex == 1 {
			factoryAddressBytes = pathExecutionValues.V3SwapDatas[i].FactoryAddress.Bytes()
		}

		v3SwapDataBytesArr[i] = V3SwapDataBytes{
			SwapIndexAndAmountIndexBytes: swapIndexAndAmountIndexBytes,

			IsFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes: isFlashswapAndV3SwapsDoneAndFeeIndexAndFactoryIndexBytes,

			Token0IndexAndToken1IndexBytes: token0IndexAndToken1IndexBytes,

			FactoryAddressBytes: factoryAddressBytes,
		}
	}

	return v3SwapDataBytesArr, nil
}
func getV2SwapDataBytesArr(pathExecutionValues PathExecutionValues) ([]V2SwapDataBytes, error) {
	v2SwapDataBytesArr := make([]V2SwapDataBytes, len(pathExecutionValues.V2SwapDatas))
	for i := 0; i < len(pathExecutionValues.V2SwapDatas); i++ {
		//SwapIndexAndAmountIndexBytes
		SwapIndexAndAmountIndexValue := pathExecutionValues.V2SwapDatas[i].SwapIndex*8 + pathExecutionValues.V2SwapDatas[i].AmountIndex
		SwapIndexAndAmountIndexBytes := []byte{byte(SwapIndexAndAmountIndexValue)}

		v2SwapDataBytesArr[i] = V2SwapDataBytes{
			SwapIndexAndAmountIndexBytes: SwapIndexAndAmountIndexBytes,
		}
	}
	return v2SwapDataBytesArr, nil
}
