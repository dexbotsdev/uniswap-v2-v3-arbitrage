package main

import (
	"fmt"
	"math/big"
	"mev-template-go/types"
	"mev-template-go/uniswap_v2"

	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Define the Sync event structure
//event Sync(uint112 reserve0, uint112 reserve1);
type Sync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}

func getSyncLogs(poolAddresses []common.Address, ctx context.Context, client *ethclient.Client, header *geth_types.Header) ([]geth_types.Log, error) {
	fmt.Println("getSyncLogs")

	// Create a filter query to get sync events from the specified block range

	//create sync event signature by using keccak256 "Sync(uint112,uint112)"

	filterQuery := ethereum.FilterQuery{
		FromBlock: header.Number,
		Addresses: poolAddresses,
		Topics: [][]common.Hash{
			{
				common.HexToHash("1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"), // Sync event topic
			},
		},
	}

	// Use the filter to get the sync events from the blockchain
	logs, err := client.FilterLogs(ctx, filterQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println("logs: ", logs)

	return logs, nil
}
func V2GetPoolFromAddress(address common.Address, pools []*uniswap_v2.Pool) (*uniswap_v2.Pool, error) {
	fmt.Println("getPoolFromAddress")

	for _, pool := range pools {
		if pool.Address == address {
			return pool, nil
		}
	}
	return nil, fmt.Errorf("pool not found")
}
func UniswapV2UpdatePoolsAndGetAffectedAddresses(pools []*uniswap_v2.Pool, header *geth_types.Header, config types.Config) ([]common.Address, error) {
	fmt.Println("UpdateAndGetAffectedPools")
	//Steps
	//get pools that are affected by the block
	//update the pool objects

	//get relevant sync events
	poolAddresses := make([]common.Address, 0)
	for _, pool := range pools {
		poolAddresses = append(poolAddresses, pool.Address)
	}

	syncEventLogs, err := getSyncLogs(poolAddresses, context.Background(), &config.Client, header)
	if err != nil {
		return nil, err
	}
	fmt.Println("syncEventLogs: ", syncEventLogs)

	//for each sync event recored addressa and update pool object
	affectedAddresses := make([]common.Address, 0)
	for _, log := range syncEventLogs {
		affectedAddresses = append(affectedAddresses, log.Address)

		affectedPool, err := V2GetPoolFromAddress(log.Address, pools)
		if err != nil {
			return nil, err
		}

		//decode sync events
		Uint112, _ := abi.NewType("uint112", "", nil)
		newReserves, err := abi.Arguments{{Type: Uint112}, {Type: Uint112}}.Unpack(log.Data)
		if err != nil {
			return nil, err
		}

		affectedPool.Reserve0.Set(newReserves[0].(*big.Int))
		affectedPool.Reserve1.Set(newReserves[1].(*big.Int))
	}

	//return the pool addresses
	return affectedAddresses, nil
}
