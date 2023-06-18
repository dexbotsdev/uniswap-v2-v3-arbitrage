package main

import (
	"errors"
	"fmt"
	"math/big"
	executor "mev-template-go/executor_deployed_v2"
	"mev-template-go/logic"
	"mev-template-go/types"
	"mev-template-go/uniswap_v2"
	"mev-template-go/uniswap_v3"
	"sync"

	geth_types "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
)

func HandleNewBlock(pools []types.PoolInterface, poolToPathMap map[common.Address][]logic.Path, poolAddressToPoolMap map[common.Address]types.PoolInterface, config types.Config, header *geth_types.Header) error {
	fmt.Println("HandleNewBlock")
	//get affected pools

	//update pools

	//get affected paths from poolToPathMap

	//calculate amountIn and revenue for all affected paths

	//sort paths by revenue

	//try to execute path with highest revenue

	//get v2 and v3 pools
	v2Pools := make([]*uniswap_v2.Pool, 0)
	v3Pools := make([]*uniswap_v3.Pool, 0)
	for _, pool := range pools {
		if pool.GetType() == "uniswap_v2" {
			v2Pools = append(v2Pools, pool.(*uniswap_v2.Pool))
		} else if pool.GetType() == "uniswap_v3" {
			v3Pools = append(v3Pools, pool.(*uniswap_v3.Pool))
		}
	}

	v2AffectedAddresses, err := UniswapV2UpdatePoolsAndGetAffectedAddresses(v2Pools, header, config)
	if err != nil {
		fmt.Println("UniswapV2UpdatePoolsAndGetAffectedAddresses error: ", err)
		return nil
	}

	v3AffectedAddresses, err := UniswapV3UpdatePoolsAndGetAffectedAddresses(v3Pools, header, config)
	if err != nil {
		fmt.Println("UniswapV3UpdatePoolsAndGetAffectedAddresses error: ", err)
		return nil
	}

	//get affected paths without duplicates
	uniqueAffectedPaths := make([]*logic.Path, 0)
	uniqueAffectedPathMap := make(map[*logic.Path]bool)
	for _, poolAddress := range v2AffectedAddresses {
		for _, path := range poolToPathMap[poolAddress] {
			if _, ok := uniqueAffectedPathMap[&path]; !ok {
				uniqueAffectedPaths = append(uniqueAffectedPaths, &path)
				uniqueAffectedPathMap[&path] = true
			}
		}
	}
	for _, poolAddress := range v3AffectedAddresses {
		for _, path := range poolToPathMap[poolAddress] {
			if _, ok := uniqueAffectedPathMap[&path]; !ok {
				uniqueAffectedPaths = append(uniqueAffectedPaths, &path)
				uniqueAffectedPathMap[&path] = true
			}
		}
	}
	fmt.Println("unique affected paths: ", len(uniqueAffectedPaths))

	//get all pools of affected addresses and call pool.Update()
	pools = make([]types.PoolInterface, 0)
	for _, poolAddress := range v2AffectedAddresses {
		pools = append(pools, poolAddressToPoolMap[poolAddress])
	}
	for _, poolAddress := range v3AffectedAddresses {
		pools = append(pools, poolAddressToPoolMap[poolAddress])
	}

	//update pools via go routine
	var wg sync.WaitGroup
	for i := 0; i < len(pools); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			pools[i].Update(&config.Client)
		}(i)
	}
	wg.Wait()

	fmt.Println("affected pools updated")

	//calculate amountIn and revenue for all affected paths
	for i := 0; i < len(uniqueAffectedPaths); i++ {
		err := handleAffectedPath(*uniqueAffectedPaths[i], header, config)
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("affected paths handled")

	// //sort paths by revenue without function; revenue is *big.int
	// sort.Slice(affectedPaths, func(i, j int) bool {
	// 	return affectedPaths[i].Revenue.Cmp(affectedPaths[j].Revenue) > 0
	// })
	return nil
}

func handleAffectedPath(path logic.Path, header *geth_types.Header, config types.Config) error {
	//update pools of path via go routine

	//
	amountIn, revenue, err := path.SetBestAmountInAndRevenue(header.BaseFee)
	if err != nil {
		return err
	}

	fmt.Println("amountIn: ", amountIn)
	fmt.Println("revenue: ", revenue)

	//if revenue < 0, skip
	if revenue.Cmp(big.NewInt(0)) <= 0 {
		return errors.New("revenue is negative")
	}

	//try to execute path with highest revenue
	err = executor.ExecuteMixedPath(path, config)
	if err != nil {
		return err
	}
	return nil
}
