package main

import (
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
)

func createTokenToPoolMap(pools []types.UniV2Pool) map[common.Address][]types.UniV2Pool {
	tokenToPoolMap := make(map[common.Address][]types.UniV2Pool)

	// Loop through each pool
	for _, pool := range pools {
		// Add the pool to the tokenToPoolMap for each token
		tokenToPoolMap[pool.Token0] = append(tokenToPoolMap[pool.Token0], pool)
		tokenToPoolMap[pool.Token1] = append(tokenToPoolMap[pool.Token1], pool)
	}

	return tokenToPoolMap
}
