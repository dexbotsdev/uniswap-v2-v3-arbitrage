package main

import (
	"encoding/json"
	"io/ioutil"
	"mev-template-go/types"
)

func WriteV2PoolsToFile(pools []types.UniV2Pool, filePath string) error {
	// Convert the slice of Pool structs to JSON data
	jsonData, err := json.Marshal(pools)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
func ReadV2PoolsFromFile(filePath string) ([]types.UniV2Pool, error) {
	// Read the file contents
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a slice of Pool structs
	var pools []types.UniV2Pool
	err = json.Unmarshal(fileContents, &pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}
