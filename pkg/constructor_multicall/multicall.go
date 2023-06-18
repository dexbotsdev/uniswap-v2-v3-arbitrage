package constructor_multicall

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Multicall(targets []common.Address, datas []byte, client *ethclient.Client) ([]byte, error) {
	//check that the number of targets and datas are the same
	if len(targets) != len(datas) {
		return nil, fmt.Errorf("number of targets and datas are not the same")
	}

	//get bytecode of the contract
	contractBytecode, err := GetBytecodeFromBin(`bin\pkg\constructor_multicall\contracts\GetUniswapV2Pools.bin`)
	if err != nil {
		log.Fatal(err)
	}

	//abi pack the targets and datas arrays
	Address, _ := abi.NewType("address", "", nil)
	targetsByteCode, err := abi.Arguments{{Type: Address}}.Pack(targets)
	if err != nil {
		panic(err)
	}

	//abi pack the targets and datas arrays
	Bytes, _ := abi.NewType("bytes", "", nil)
	datasByteCode, err := abi.Arguments{{Type: Bytes}}.Pack(datas)
	if err != nil {
		panic(err)
	}

	fullBytecode := append(contractBytecode, append(targetsByteCode, datasByteCode...)...)

	//create eth msg
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"), // Set the caller address to the zero address
		To:   nil,
		Data: common.Hex2Bytes(string(fullBytecode)),
	}

	//make the eth_call
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result: ", result)
	fmt.Println("result: ", result)

	return fullBytecode, nil
}
