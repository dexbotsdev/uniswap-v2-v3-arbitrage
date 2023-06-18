package constructor_multicall

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetUniswapV2ReservesCall(poolAddresses []common.Address, client *ethclient.Client) (*big.Int, []*big.Int, []*big.Int) {
	//get contract bytecode
	contractBytecode, err := GetBytecodeStringFromBin(`bin\pkg\constructor_multicall\contracts\getUniswapV2Reserves.bin`)
	if err != nil {
		log.Fatal(err)
	}
	//get the batch of pools
	//create bytecode of the contract

	//get arguements bytes
	AddressArr, _ := abi.NewType("address[]", "", nil)
	arguementsBytes, err := abi.Arguments{{Type: AddressArr}}.Pack(poolAddresses)
	if err != nil {
		panic(err)
	}

	arguementsBytesHex := hex.EncodeToString(arguementsBytes)

	//fmt.Println("contractBytecode: ", contractBytecode)
	fmt.Println("arguementsBytes: ", arguementsBytesHex)

	//concatnate bytecode and arguements bytes
	fullBytecode := contractBytecode + arguementsBytesHex
	fullBytecodeString, err := hex.DecodeString(fullBytecode)

	fmt.Println("fullBytecode: ", fullBytecode)
	//create the call message
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"), // Set the caller address to the zero address
		To:   nil,
		Data: fullBytecodeString,
	}

	//make the eth_call
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result: ", result)
	fmt.Println("result: ", result)

	//decode the result
	UintArr, _ := abi.NewType("uint256[]", "", nil)
	Uint, _ := abi.NewType("uint256", "", nil)
	decodedResults, err := abi.Arguments{{Type: Uint}, {Type: UintArr}, {Type: UintArr}}.Unpack(result)
	if err != nil {
		panic(err)
	}
	fmt.Println("decodedResults: ", decodedResults)

	blockNumber := decodedResults[0].(*big.Int)
	reserve0s := decodedResults[1].([]*big.Int)
	reserve1s := decodedResults[2].([]*big.Int)
	//create new univ2pool batch

	fmt.Println("blockNumber: ", blockNumber)
	fmt.Println("reserve0s: ", reserve0s)
	fmt.Println("reserve1s: ", reserve1s)

	//return blockNumber, reserve0s, reserve1s
	return blockNumber, reserve0s, reserve1s
}
