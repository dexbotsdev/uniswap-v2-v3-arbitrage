package constructor_multicall

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

	//fmt.Println("binData: ", binData)
	//fmt.Println("binData: ", string(binData))

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
