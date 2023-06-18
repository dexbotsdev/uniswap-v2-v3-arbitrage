package uniswap_v3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"mev-template-go/pkg/data"
)

const filterdPoolsFilePath = "filteredUniswapV3Pools.json"

func FilterPoolsAndWriteToFile() ([]*Pool, error) {

	//read files
	pools, err := ReadPoolsFromFile()
	if err != nil {
		return nil, err
	}

	fmt.Println("length of unfiltered Uniswap V3 pools: ", len(pools))

	filteredPools, err := FilterPools(pools)
	if err != nil {
		return nil, err
	}

	fmt.Println("length of filtered Uniswap V3 pools: ", len(filteredPools))

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

func ReadFilteredPoolsFromFile() ([]*Pool, error) {
	// Read the file contents
	fileContents, err := ioutil.ReadFile(filterdPoolsFilePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a slice of Pool structs
	var pools []*Pool
	err = json.Unmarshal(fileContents, &pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

func FilterPools(pools []*Pool) ([]*Pool, error) {
	filteredPools := getPoolsWithoutBlacklistedTokens(pools)
	filteredPools = getPoolsWithWhitelistedTokens(filteredPools)
	filteredPools = filterOutPoolsWithoutTicks(filteredPools)
	filteredPools = getPoolsWithLiquidity(filteredPools)

	return filteredPools, nil
}

func getPoolsWithoutBlacklistedTokens(pools []*Pool) []*Pool {
	filteredPools := []*Pool{}
	for _, pool := range pools {
		//if pool contains a blacklisted token, skip it
		if data.IsBlacklisted(pool.Token0.String()) || data.IsBlacklisted(pool.Token1.String()) {
			continue
		}
		filteredPools = append(filteredPools, pool)
	}

	return filteredPools
}
func getPoolsWithLiquidity(pools []*Pool) []*Pool {
	filteredPools := []*Pool{}
	for _, pool := range pools {
		if pool.Liquidity.Cmp(big.NewInt(0)) == 0 {
			continue
		}
		filteredPools = append(filteredPools, pool)
	}

	return filteredPools
}

func getPoolsWithWhitelistedTokens(pools []*Pool) []*Pool {
	filteredPools := []*Pool{}
	for _, pool := range pools {
		//if pool contains a blacklisted token, skip it
		if data.IsWhitelisted(pool.Token0.String()) && data.IsWhitelisted(pool.Token1.String()) {
			filteredPools = append(filteredPools, pool)
		}
	}

	return filteredPools
}

func filterOutPoolsWithoutTicks(pools []*Pool) []*Pool {
	filteredPools := []*Pool{}
	for _, pool := range pools {
		if len(pool.Ticks) == 0 {
			continue
		}
		filteredPools = append(filteredPools, pool)
	}

	return filteredPools
}

func ReadPoolsFromFile() ([]*Pool, error) {
	// Read the file contents
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a slice of Pool structs
	var pools []*Pool
	err = json.Unmarshal(fileContents, &pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}
