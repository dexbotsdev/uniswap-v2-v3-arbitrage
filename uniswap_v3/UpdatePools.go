package uniswap_v3

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"mev-template-go/pkg/constructor_multicall"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func UpdatePools(pools []*Pool, client *ethclient.Client) error {
	return updatePoolsBatched(pools, client)
}

func updatePoolsBatched(pools []*Pool, client *ethclient.Client) error {
	batchSize := 1

	for i := 0; i < len(pools); i += batchSize {

		//Set batch end. If current batch size surpasses pools length, set end to pools length.
		end := i + batchSize
		if end > len(pools) {
			end = len(pools)
		}

		batch := pools[i:end]
		//update all feilds exepct ticks
		_, err := updatePoolsMulticall(batch, client)
		if err != nil {
			return err
		}

		//update ticks
		err = UpdateTicks(batch, client)
		if err != nil {
			return err
		}
	}
	//split pools into batches
	return nil
}

//updates the mutable varaibles of the pools
func updatePoolsMulticall(pools []*Pool, client *ethclient.Client) (*big.Int, error) {

	//get contract bytecode
	contractBytecode, err := constructor_multicall.GetBytecodeStringFromBin(`bin\uniswap_v3\contracts\GetUniswapV3Pools.bin`)
	if err != nil {
		return nil, err
	}

	//add augurements
	AddressArr, _ := abi.NewType("address[]", "", nil)
	poolAddresses := make([]common.Address, len(pools))
	for i, pool := range pools {
		poolAddresses[i] = pool.Address
	}
	argBytes, err := abi.Arguments{{Type: AddressArr}}.Pack(poolAddresses)
	if err != nil {
		return nil, err
	}

	fullBytecode := contractBytecode + hex.EncodeToString(argBytes)
	fullBytecodeString, err := hex.DecodeString(fullBytecode)

	//create call msg
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"), // Set the caller address to the zero address
		To:   nil,
		Data: fullBytecodeString,
	}
	//fmt.Println("callMsg: ", callMsg)

	//make the eth_call
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}
	//fmt.Println("result: ", result)

	//decode result
	//AddressArr, _ = abi.NewType("address[]", "", nil)
	Uint, _ := abi.NewType("uint256", "", nil)
	UintArr, _ := abi.NewType("uint256[]", "", nil)
	IntArr, _ := abi.NewType("int256[]", "", nil)

	//decode results
	decodedResults, err := abi.Arguments{{Type: Uint}, {Type: UintArr}, {Type: UintArr}, {Type: IntArr}, {Type: IntArr}}.Unpack(result)
	if err != nil {
		return nil, err
	}
	//fmt.Println("decodedResults: ", decodedResults)

	//update pool fields
	blockNumber := decodedResults[0].(*big.Int)
	for i := 0; i < len(pools); i++ {
		//update token0
		pools[i].Liquidity = decodedResults[1].([]*big.Int)[i]
		pools[i].SqrtPriceX96 = decodedResults[2].([]*big.Int)[i]
		pools[i].TickCurrent = int(decodedResults[3].([]*big.Int)[i].Int64())
		pools[i].TickSpacing = int(decodedResults[4].([]*big.Int)[i].Int64())

		//if tickspacing is 0 return err
		if pools[i].TickSpacing == 0 {
			return nil, fmt.Errorf("tickSpacing is 0")
		}

		//pools[i].PopulatedTicks = decodedResults[4].([][]PopulatedTick)[i]
	}

	return blockNumber, nil
}
