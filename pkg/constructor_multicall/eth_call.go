package constructor_multicall

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type EthCallArgs struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

type EthCallResult struct {
	Result string `json:"result"`
	Error  struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func EthCall() {
	// Create a new RPC client
	client, err := rpc.DialHTTP("tcp", os.Getenv("NETWORK_RPC"))
	if err != nil {
		log.Fatal(err)
	}

	// Define the contract address and data
	contractAddress := "CONTRACT_ADDRESS"
	contractData := "FUNCTION_SIGNATURE"

	// Define the call arguments
	args := EthCallArgs{
		To:   contractAddress,
		Data: contractData,
	}

	// Send the eth_call request
	var result EthCallResult
	err = client.Call("eth_call", args, &result)
	if err != nil {
		log.Fatal(err)
	}

	// Check for errors in the response
	if result.Error.Message != "" {
		log.Fatal(result.Error.Message)
	}

	// Print the result
	fmt.Println(result.Result)
}
