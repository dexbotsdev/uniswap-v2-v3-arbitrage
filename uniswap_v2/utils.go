package uniswap_v2

import (
	"fmt"
	"io/ioutil"
	"math/big"

	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

//amountOut = (amountIn * reserveOut * (997 - fee)) / (reserveIn * 1000 + amountIn * (997 - fee))

//create swap call data
func CreateSwapCallData(amount0Out *big.Int, amount1Out *big.Int, toAddress common.Address, data []byte) []byte {
	// convert amounts to byte arrays
	amount0OutBytes := amount0Out.Bytes()
	amount1OutBytes := amount1Out.Bytes()

	// pad byte arrays with leading zeros to 32 bytes
	paddedAmount0OutBytes := common.LeftPadBytes(amount0OutBytes, 32)
	paddedAmount1OutBytes := common.LeftPadBytes(amount1OutBytes, 32)

	// create swap call data
	callData := []byte{}
	callData = append(callData, common.Hex2Bytes("a2e2e6b6")...)               // method signature
	callData = append(callData, paddedAmount0OutBytes...)                      // amount0Out
	callData = append(callData, paddedAmount1OutBytes...)                      // amount1Out
	callData = append(callData, common.LeftPadBytes(toAddress.Bytes(), 32)...) // to address
	callData = append(callData, common.LeftPadBytes(data, 32)...)              // additional data

	return callData
}
func GetReserves(pair common.Address, client *ethclient.Client) (*big.Int, *big.Int, error) {
	//get reserves
	getReservesCallData := []byte{}
	getReservesCallData = append(getReservesCallData, common.Hex2Bytes("0902f1ac")...) // method signature
	//make an eth_call to get the reserves

	//get lastest block number
	blockNumber, err := client.BlockNumber(context.Background())
	blockNumberBigInt := big.NewInt(int64(blockNumber))
	fmt.Println("blockNumber", blockNumber)

	//build call message
	msg := ethereum.CallMsg{
		From:     common.Address{},
		To:       &pair,
		Gas:      0,
		GasPrice: nil,
		Value:    nil,
		Data:     getReservesCallData,
	}

	reserves, err := client.CallContract(context.Background(), msg, blockNumberBigInt)
	if err != nil {
		return nil, nil, err
	}
	//if length is less than 64, then there is an error
	if len(reserves) < 64 {
		return nil, nil, fmt.Errorf("error getting reserves")
	}
	reserve0 := new(big.Int).SetBytes(reserves[0:32])
	reserve1 := new(big.Int).SetBytes(reserves[32:64])

	return reserve0, reserve1, err
}

func SimulateContractDeployment(client *ethclient.Client, bin []byte) ([]byte, error) {
	// Prepare the transaction data
	txData := hexutil.Encode(bin)
	msg := ethereum.CallMsg{
		Data: []byte(txData),
	}

	// Make the eth_call RPC request to the Ethereum client
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call contract: %v", err)
	}

	return result, nil
}

func GetBytecodeStringFromBin(binFilePath string) (string, error) {
	// Read the .bin file into a byte array
	binData, err := GetBytecodeFromBin(binFilePath)
	if err != nil {
		return "", err
	}

	fmt.Println("binData: ", binData)
	fmt.Println("binData: ", string(binData))

	// Convert the byte array to a hex string

	return string(binData), nil
}
func GetBytecodeFromBin(filename string) ([]byte, error) {
	// read the entire contents of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading bytecode file: %v", err)
	}

	// return the contents as a slice of bytes
	return content, nil
}
