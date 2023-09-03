package executor

//changes to make: add bribe amount instead of percentage
//add owner transfer bool

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"math/big"
	"mev-template-go/path"
	//"github.com/umbracle/ethgo/abi"
)

func PathToCalldata(path path.Path, bribe *big.Int, ownerTransfer bool, coinbaseTransfer bool) ([]byte, error) {
	pathExecutionValues, err := convertPathToPathExecutionValues(path, bribe, ownerTransfer, coinbaseTransfer)
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
