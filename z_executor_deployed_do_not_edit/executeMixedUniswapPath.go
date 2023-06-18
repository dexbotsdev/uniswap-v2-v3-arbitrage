package executor

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"mev-template-go/logic"

	"github.com/ethereum/go-ethereum/common"
	//"github.com/umbracle/ethgo/abi"
)

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

var executorAddress = common.HexToAddress("0x5615dEB798BB3E4dFa0139dFa1b3D433Cc23b72f") //TODO

func executeMixedPath(path logic.Path) ([]byte, error) {
	//convert path to payload values

	pathExecutionValues, err := convertPathToPathExecutionValues(path)
	if err != nil {
		return nil, err
	}

	//get payload bytes struct
	flashswapPayloadBytesStruct, err := convertPathExecutionValuesToFlashswapPayloadBytes(pathExecutionValues)
	if err != nil {
		return nil, err
	}

	//create payload
	flashswapPayload := appendFlashswapPayloadBytesToBytesArray(flashswapPayloadBytesStruct)

	calldata, err := createFlashswapCalldata(pathExecutionValues, flashswapPayload)
	if err != nil {
		return nil, err
	}

	return calldata, nil
}
