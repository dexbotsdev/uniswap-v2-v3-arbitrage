package executor

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	//"github.com/umbracle/ethgo/abi"
	UniV2Pair "mev-template-go/uniswap_v2/contracts/uniswap_v2_pair"
)

type FlashswapCall struct {
	Address  common.Address
	Calldata []byte
}

type V3Calldata struct { //pool is called via executor contract
	target     common.Address
	zeroForOne bool
	amountIn   *big.Int
	payload    []byte
}

type V2Calldata struct { //pool is called directly
	amountOut0 *big.Int
	amountOut1 *big.Int
	to         common.Address
	payload    []byte
}

// type PathExecutionValues struct {
// 	//V2 flashswap values a
// 	IsV2Flashswap         bool           //1 bit
// 	V2FlashswapTarget     common.Address //20 bytes
// 	V2FlashswapZeroForOne bool           //1 bit
// 	V2FlashswapAmountOut  *big.Int       //IntByteSize bytes

// 	//for V3 Initial call
// 	V3FlashswapTarget     common.Address //20 bytes
// 	V3FlashswapZeroForOne bool           //1 bit
// 	V3FlashswapAmountIn   *big.Int       //IntByteSize bytes

func createFlashswapCalldata(pathExecutionValues PathExecutionValues, payload []byte) ([]byte, error) {
	fmt.Println("***createFlashswapCalldata***")
	//TODO
	// bytes4 internal constant Uniswap_V2_SWAP_SIG = 0x022c0d9f;
	// bytes4 internal constant Uniswap_V3_SWAP_SIG = 0x128acb08;
	// uniswapV2SwapSig := "0x022c0d9f"
	// uniswapV3SwapSig := "0x128acb08"
	if pathExecutionValues.IsV2Flashswap {
		// 	IsV2Flashswap         bool           //1 bit
		// 	V2FlashswapTarget     common.Address //20 bytes
		// 	V2FlashswapZeroForOne bool           //1 bit
		// 	V2FlashswapAmountOut  *big.Int       //IntByteSize bytes
		//	Payload

		var amountOut0 *big.Int
		var amountOut1 *big.Int

		if pathExecutionValues.V2FlashswapZeroForOne {
			amountOut0 = big.NewInt(0)
			amountOut1 = new(big.Int).Set(pathExecutionValues.V2FlashswapAmountOut)
		} else {
			amountOut0 = new(big.Int).Set(pathExecutionValues.V2FlashswapAmountOut)
			amountOut1 = big.NewInt(0)
		}

		fmt.Println("flashswap amountOut0", amountOut0)
		fmt.Println("flashswap amountOut1", amountOut1)

		flashswapCalldata, err := GetUniswapV2SwapCallData(
			amountOut0,
			amountOut1,
			executorAddress,
			payload,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println("flashswapCalldata", flashswapCalldata)
		return flashswapCalldata, nil

	} else {

		return payload, nil
		// 	//print argurements
		// 	fmt.Println("flashswap executorAddress", executorAddress)
		// 	fmt.Println("flashswap executorAddressString", executorAddress)
		// 	fmt.Println("flashswap zeroForOne", pathExecutionValues.V3FlashswapZeroForOne)
		// 	fmt.Println("flashswap amountIn", pathExecutionValues.V3FlashswapAmountIn)

		// 	// //if first swap is V3, set target to executor address, else set to pool address
		// 	// if pathExecutionValues.SwapDataIsV2s[0] == false {

		// 	flashswapCalldata, err := GetUniswapV3ExecutorCalldata(
		// 		pathExecutionValues.V3FlashswapTarget,
		// 		pathExecutionValues.V3FlashswapZeroForOne,
		// 		pathExecutionValues.V3FlashswapAmountIn,
		// 		payload,
		// 	)
		// 	if err != nil {
		// 		return nil, err
		// 	}

		// 	fmt.Println("flashswapCalldata", hex.EncodeToString(flashswapCalldata))
		// 	return flashswapCalldata, nil
	}

}

func GetUniswapV3ExecutorCalldata(
	target common.Address,
	zeroForOne bool,
	amountIn *big.Int,
	payload []byte,
) ([]byte, error) {

	// //for V3 Initial call
	// V3FlashswapTarget     common.Address //20 bytes
	// V3FlashswapZeroForOne bool           //1 bit
	// V3FlashswapAmountIn   *big.Int       //Compressed Int

	fmt.Println("***GetUniswapV3ExecutorCalldata***")
	fmt.Println("target", target)
	fmt.Println("zeroForOne", zeroForOne)
	fmt.Println("amountIn", amountIn)
	fmt.Println("payload", hex.EncodeToString(payload))

	flashswapTargetBytes := target.Bytes()
	zeroForOneBytes := []byte{boolToByte(zeroForOne)}
	amountSpecifiedBytes := intToCompressedBytes(amountIn)

	fmt.Println("flashswapTargetBytes", hex.EncodeToString(flashswapTargetBytes))
	fmt.Println("zeroForOneBytes", hex.EncodeToString(zeroForOneBytes))
	fmt.Println("amountSpecifiedBytes(compressed bytes)", hex.EncodeToString(amountSpecifiedBytes))

	//append together
	calldata := append(flashswapTargetBytes, zeroForOneBytes...)
	calldata = append(calldata, amountSpecifiedBytes...)
	calldata = append(calldata, payload...)

	//

	return calldata, nil
}

func GetUniswapV2SwapCallData(
	amountOut0 *big.Int,
	amountOut1 *big.Int,
	recipient common.Address,
	data []byte,
) ([]byte, error) {

	fmt.Println("***GetUniswapV2SwapCallData***")
	fmt.Println("amountOut0", amountOut0)
	fmt.Println("amountOut1", amountOut1)
	fmt.Println("recipient", recipient)
	fmt.Println("data", hex.EncodeToString(data))

	pairAbi, err := abi.JSON(strings.NewReader(UniV2Pair.UniV2PairMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return pairAbi.Pack("swap", amountOut0, amountOut1, recipient, data)

}
func boolToByte(b bool) byte {
	var bitSet byte
	if b {
		bitSet = 1
	}
	return bitSet
}
