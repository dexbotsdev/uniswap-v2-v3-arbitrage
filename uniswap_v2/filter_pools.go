package uniswap_v2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"mev-template-go/pkg/data"
)

const filterdPoolsFilePath = "filteredUniswapV2Pools.json"

func FilterPoolsAndWriteToFile() ([]Pool, error) {

	//read files
	pools, err := ReadPoolsFromFile()
	if err != nil {
		return nil, err
	}

	fmt.Println("length of unfiltered Uniswap V2 pools: ", len(pools))

	filteredPools, err := FilterPools(pools)
	if err != nil {
		return nil, err
	}

	fmt.Println("length of filtered Uniswap V2 pools: ", len(filteredPools))

	//write to file

	jsonData, err := json.Marshal(filteredPools)
	if err != nil {
		return nil, err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filterdPoolsFilePath, jsonData, 0644)
	if err != nil {
		return nil, err
	}

	return filteredPools, nil
}

func getPoolsWithoutBlacklistedTokens(pools []Pool) []Pool {
	filteredPools := []Pool{}
	for _, pool := range pools {
		//if pool contains a blacklisted token, skip it
		fmt.Println("pool token0: ", pool.Token0.String())
		fmt.Println("pool token1: ", pool.Token1.String())
		if data.IsBlacklisted(pool.Token0.String()) || data.IsBlacklisted(pool.Token1.String()) {
			fmt.Println("skipping pool with blacklisted token")
			continue
		}
		filteredPools = append(filteredPools, pool)
	}

	return filteredPools
}

func getPoolsWithWhitelistedTokens(pools []Pool) []Pool {
	filteredPools := []Pool{}
	for _, pool := range pools {
		//if pool contains a blacklisted token, skip it
		fmt.Println("pool token0: ", pool.Token0.String())
		fmt.Println("pool token1: ", pool.Token1.String())
		if data.IsWhitelisted(pool.Token0.String()) && data.IsWhitelisted(pool.Token1.String()) {
			filteredPools = append(filteredPools, pool)
		}
	}

	return filteredPools
}

func filterOutPoolsWithEmptyReserves(pools []Pool) []Pool {
	filteredPools := []Pool{}
	for _, pool := range pools {
		if pool.Reserve0.Cmp(big.NewInt(0)) == 0 || pool.Reserve1.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		filteredPools = append(filteredPools, pool)
	}

	return filteredPools
}

func filterOutPoolsWithReservesLessThanMin(pools []Pool, minReserves *big.Int) []Pool {
	filteredPools := []Pool{}
	for _, pool := range pools {
		if pool.Reserve0.Cmp(minReserves) == -1 || pool.Reserve1.Cmp(minReserves) == -1 {
			continue
		}
		filteredPools = append(filteredPools, pool)
	}

	return filteredPools
}

func FilterPools(pools []Pool) ([]Pool, error) {
	filteredPools := getPoolsWithoutBlacklistedTokens(pools)
	filteredPools = getPoolsWithWhitelistedTokens(filteredPools)
	//filteredPools = filterOutPoolsWithEmptyReserves(filteredPools)
	filteredPools = filterOutPoolsWithReservesLessThanMin(filteredPools, big.NewInt(10000))

	return filteredPools, nil
}

func ReadPoolsFromFile() ([]Pool, error) {
	// Read the file contents
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a slice of Pool structs
	var pools []Pool
	err = json.Unmarshal(fileContents, &pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

func ReadFilteredPoolsFromFile() ([]Pool, error) {
	// Read the file contents
	fileContents, err := ioutil.ReadFile(filterdPoolsFilePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a slice of Pool structs
	var pools []Pool
	err = json.Unmarshal(fileContents, &pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}
