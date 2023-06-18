package main

import (
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
)

func getPoolAddressesFromPools(pools []types.UniV2Pool) ([]common.Address, error) {
	poolAddresses := make([]common.Address, 0)
	for _, pool := range pools {
		poolAddresses = append(poolAddresses, pool.Address)
	}
	return poolAddresses, nil
}

func removeDuplicatePaths(structs []types.Path) []types.Path {
	encountered := map[*types.Path]bool{}
	result := []types.Path{}

	for _, s := range structs {
		if !encountered[&s] {
			encountered[&s] = true
			result = append(result, s)
		}
	}

	return result
}
