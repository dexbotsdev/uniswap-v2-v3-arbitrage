package uniswap_v3

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const swapEventHash = "c42079f94a6350d7e6235f29174924f928cc2ac818eb64fed8004e115fbcca67"
const burnEventHash = "0c396cd989a39f4459b5fa1aed6a9a8dcdbc45908acfd67e028cd568da98982c"
const mintEventHash = "7a53080ba414158be7ec69b987b5fb7d07dee101fe85488f0853ae16239d0bde"

func UpdatePoolsOnNewBlock(pools []*Pool, ctx context.Context, client *ethclient.Client, header *geth_types.Header) error {
	fmt.Println("Uniswap V3 HandleNewBlock")

	//get affected addresses from logs
	logs, err := getLogs(getPoolAddresses(pools), ctx, client, header)
	if err != nil {
		return err
	}

	//get affected pools from logs
	affectedPoolAddresses := getAffectedPoolAddresses(logs)

	affectedPools := getPoolsFromAddresses(affectedPoolAddresses, pools)

	//update affected pools

	err = updatePoolsBatched(affectedPools, client)
	if err != nil {
		return err
	}

	return nil
}

//filter logs for Swap Burn Mint events
func getLogs(poolAddresses []common.Address, ctx context.Context, client *ethclient.Client, header *geth_types.Header) ([]geth_types.Log, error) {
	fmt.Println("getSyncLogs")

	// Create a filter query to get sync events from the specified block range

	//create sync event signature by using keccak256 "Sync(uint112,uint112)"

	filterQuery := ethereum.FilterQuery{
		FromBlock: header.Number,
		Addresses: poolAddresses,
		Topics: [][]common.Hash{
			{
				common.HexToHash(swapEventHash),
				common.HexToHash(burnEventHash),
				common.HexToHash(mintEventHash),
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
