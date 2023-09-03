package handle_new_block

import (
	"fmt"
	"mev-template-go/types"
	"mev-template-go/uniswap_v3"

	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func V3GetAffectedPoolAddresses(poolAddresses []common.Address, ctx context.Context, client *ethclient.Client, header *geth_types.Header) ([]common.Address, error) {
	fmt.Println("getSyncLogs")

	// Create a filter query to get sync events from the specified block range

	//create sync event signature by using keccak256 "Sync(uint112,uint112)"

	filterQuery := ethereum.FilterQuery{
		FromBlock: header.Number,
		Addresses: poolAddresses,
	}

	// Use the filter to get the sync events from the blockchain
	logs, err := client.FilterLogs(ctx, filterQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println("logs: ", logs)

	//get addresses from logs
	addresses := make([]common.Address, 0)
	for _, log := range logs {
		addresses = append(addresses, log.Address)
	}

	return addresses, nil
}
func V3GetPoolFromAddress(address common.Address, pools []*uniswap_v3.Pool) (*uniswap_v3.Pool, error) {
	fmt.Println("getPoolFromAddress")

	for _, pool := range pools {
		if pool.Address == address {
			return pool, nil
		}
	}
	return nil, fmt.Errorf("pool not found")
}
func UniswapV3UpdatePoolsAndGetAffectedAddresses(pools []*uniswap_v3.Pool, header *geth_types.Header, config types.Config) ([]common.Address, error) {
	fmt.Println("UpdateAndGetAffectedPools")
	//Steps
	//get v3 pools
	//use filter logs to get events
	//update pools
	//

	//get relevant sync events
	poolAddresses := make([]common.Address, 0)
	for _, pool := range pools {
		poolAddresses = append(poolAddresses, pool.Address)
	}

	affectedAddresses, err := V3GetAffectedPoolAddresses(poolAddresses, context.Background(), &config.Client, header)
	if err != nil {
		return nil, err
	}

	//get pools from addresses
	affectedPools := make([]*uniswap_v3.Pool, 0)
	for _, address := range affectedAddresses {
		pool, err := V3GetPoolFromAddress(address, pools)
		if err != nil {
			return nil, err
		}
		affectedPools = append(affectedPools, pool)
	}

	//update pools
	for _, pool := range affectedPools {
		err := pool.Update(&config.Client)
		if err != nil {
			return nil, err
		}
	}

	//for each sync event recored addressa and update pool object

	//return the pool addresses
	return affectedAddresses, nil
}
