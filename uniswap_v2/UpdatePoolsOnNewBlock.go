package uniswap_v2

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const syncEventHash = "1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1"

func UpdatePoolsOnNewBlock(pools []*Pool, ctx context.Context, client *ethclient.Client, header *geth_types.Header) error {
	fmt.Println("Uniswap V3 HandleNewBlock")

	//get affected addresses from logs
	logs, err := getSyncLogs(getPoolAddresses(pools), ctx, client, header)
	if err != nil {
		return err
	}

	//get affected pools from logs
	affectedPoolAddresses := getAffectedPoolAddresses(logs)

	affectedPools := getPoolsFromAddresses(affectedPoolAddresses, pools)

	//update affected pools
	err = UpdatePools(affectedPools, client)
	if err != nil {
		return err
	}

	return nil
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
				common.HexToHash(syncEventHash), // Sync event topic
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

func getPoolsFromAddresses(addresses []common.Address, pools []*Pool) []*Pool {
	var result []*Pool
	for _, address := range addresses {
		for _, pool := range pools {
			if address == pool.Address {
				result = append(result, pool)
			}
		}
	}
	return result
}

func getPoolAddresses(pools []*Pool) []common.Address {
	var poolAddresses []common.Address
	for _, pool := range pools {
		poolAddresses = append(poolAddresses, pool.Address)
	}
	return poolAddresses
}

func addressesToPools(addresses []common.Address, pools []Pool) []Pool {
	var result []Pool
	for _, address := range addresses {
		for _, pool := range pools {
			if address == pool.Address {
				result = append(result, pool)
			}
		}
	}
	return result
}

//get affected pool addresses from logs without duplicates
func getAffectedPoolAddresses(logs []geth_types.Log) []common.Address {
	var poolAddresses []common.Address
	var seen = make(map[common.Address]bool)
	for _, log := range logs {
		if !seen[log.Address] {
			poolAddresses = append(poolAddresses, log.Address)
			seen[log.Address] = true
		}
	}
	return poolAddresses
}
