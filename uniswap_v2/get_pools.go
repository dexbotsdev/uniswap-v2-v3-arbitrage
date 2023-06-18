package uniswap_v2

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const filePath = "uniswapV2Pools.json"

func GetUniswapV2AllPoolLengths(factoryAddress common.Address, client *ethclient.Client) (uint64, error) {
	fmt.Println("GetUniswapV2AllPoolLengths")
	// Create the Uniswap V2 factory contract instance

	uniswapV2FactoryABI := `[
		{
			"inputs": [],
			"name": "allPairsLength",
			"outputs": [
				{
					"internalType": "uint256",
					"name": "",
					"type": "uint256"
				}
			],
			"stateMutability": "view",
			"type": "function",
			"constant": true
		}
	]`

	abi, err := abi.JSON(strings.NewReader(uniswapV2FactoryABI))
	if err != nil {
		return 0, err
	}

	// Encode the function call
	data, err := abi.Pack("allPairsLength")
	if err != nil {
		return 0, err
	}

	// Make the eth_call
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"), // Set the caller address to the zero address
		To:   &factoryAddress,
		Data: data,
	}
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return 0, err
	}

	// Decode the result
	var allPairsLength *big.Int
	decodedResult, err := abi.Unpack("allPairsLength", result)
	if err != nil {
		log.Fatal(err)
	}
	if len(decodedResult) == 0 {
		log.Fatal("empty result")
	}
	allPairsLength = decodedResult[0].(*big.Int)

	// Convert the result to uint64
	return allPairsLength.Uint64(), nil
}
func GetUniswapV2PoolsCall(factoryAddress common.Address, start *big.Int, end *big.Int, client *ethclient.Client) (*big.Int, []Pool, error) {
	fmt.Println("GetUniswapV2PoolsCall")

	contractBytecode, err := GetBytecodeStringFromBin(`bin\pkg\constructor_multicall\contracts\GetUniswapV2Pools.bin`)
	if err != nil {
		return nil, nil, err
	}
	//get the batch of pools
	//create bytecode of the contract

	//get arguements bytes
	Address, _ := abi.NewType("address", "", nil)
	Uint, _ := abi.NewType("uint", "", nil)
	arguementsBytes, err := abi.Arguments{{Type: Address}, {Type: Uint}, {Type: Uint}}.Pack(factoryAddress, start, end)
	if err != nil {
		return nil, nil, err
	}

	arguementsBytesHex := hex.EncodeToString(arguementsBytes)
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
		return nil, nil, err
	}
	fmt.Println("result: ", result)
	fmt.Println("result: ", result)

	//decode the result
	AddressArr, _ := abi.NewType("address[]", "", nil)
	UintArr, _ := abi.NewType("uint256[]", "", nil)
	decodedResults, err := abi.Arguments{{Type: Uint}, {Type: AddressArr}, {Type: AddressArr}, {Type: AddressArr}, {Type: UintArr}, {Type: UintArr}}.Unpack(result)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("decodedResults: ", decodedResults)

	blockNumber := decodedResults[0].(*big.Int)
	addresses := decodedResults[1].([]common.Address)
	token0s := decodedResults[2].([]common.Address)
	token1s := decodedResults[3].([]common.Address)
	reserves0 := decodedResults[4].([]*big.Int)
	reserves1 := decodedResults[5].([]*big.Int)
	//create new univ2pool batch

	pools := make([]Pool, 0)
	for i := 0; i < len(addresses); i++ {
		newPool := Pool{
			Address:        addresses[i],
			FactoryAddress: factoryAddress,
			Token0:         token0s[i],
			Token1:         token1s[i],
			Reserve0:       reserves0[i],
			Reserve1:       reserves1[i],
		}
		fmt.Println("newPool: ", newPool)
		pools = append(pools, newPool)
	}

	fmt.Println("pools: ", pools)
	return blockNumber, pools, nil

}
func GetUniswapV2PoolsFromFactory(factoryAddress common.Address, client *ethclient.Client) ([]Pool, error) {
	fmt.Println("GetUniswapV2PoolsFromFactory")

	//get length of all pools
	allPoolLengths, err := GetUniswapV2AllPoolLengths(factoryAddress, client)
	if err != nil {
		return nil, err
	}
	fmt.Println("allPoolLengths: ", allPoolLengths)

	//empty array of pools
	pools := make([]Pool, 0)

	batchSize := 100

	//for each batch call eth_call on contract constructor with the arguements
	for i := 0; i < int(allPoolLengths); i += batchSize {

		start := big.NewInt(int64(i))
		//if i+batchSize > int(allPoolLengths) then end = allPoolLengths else end = i+batchSize
		end := big.NewInt(int64(i + batchSize))
		if i+batchSize > int(allPoolLengths) {
			end = big.NewInt(int64(allPoolLengths))
		}

		//get the pools from the multicall contract
		_, newPools, err := GetUniswapV2PoolsCall(factoryAddress, start, end, client)
		if err != nil {
			return nil, err
		}

		fmt.Println("newPools: ", newPools)
		pools = append(pools, newPools...)
	}

	fmt.Println("pools: ", pools)
	return pools, nil
}

func GetUniswapV2PoolsFromAllForks(client *ethclient.Client) ([]Pool, error) {
	fmt.Println("GetUniswapV2PoolsFromAllForks")

	uniswapFactoryAddress := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	sushiswapFactoryAddress := common.HexToAddress("0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac")
	//pancakeFactoryAddress := common.HexToAddress("0x1097053Fd2ea711dad45caCcc45EfF7548fCB362")
	//croFactoryAddress := common.HexToAddress("0x9DEB29c9a4c7A88a3C0257393b7f3335338D9A9D")
	//zeusFactoryAddress := common.HexToAddress("0xbdda21dd8da31d5bee0c9bb886c044ebb9b8906a")
	//luaFactoryAddress := common.HexToAddress("0x0388c1e0f210abae597b7de712b9510c6c36c857")

	factoryAddresses := make([]common.Address, 0)
	factoryAddresses = append(factoryAddresses, uniswapFactoryAddress)
	factoryAddresses = append(factoryAddresses, sushiswapFactoryAddress)
	//factoryAddresses = append(factoryAddresses, pancakeFactoryAddress)

	//empty array of pools
	pools := make([]Pool, 0)
	for _, factoryAddress := range factoryAddresses {
		newPools, err := GetUniswapV2PoolsFromFactory(factoryAddress, client)
		if err != nil {
			return nil, err
		}
		//addnew pools to pools
		pools = append(pools, newPools...)
	}
	return pools, nil
}
func GetAllUniswapV2ForkPoolsAndWriteToFile(client *ethclient.Client) ([]Pool, error) {
	fmt.Println("GetAllUniswapV2ForkPoolsAndWriteToFile")

	pools, err := GetUniswapV2PoolsFromAllForks(client)
	if err != nil {
		return nil, err
	}

	//write pools to file
	err = WritePoolsToFile(pools, filePath)
	if err != nil {
		return nil, err
	}

	return pools, nil
}
func WritePoolsToFile(pools []Pool, filePath string) error {
	// Convert the slice of Pool structs to JSON data
	jsonData, err := json.Marshal(pools)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
