package uniswap_v3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SubgraphPool struct {
	ID      string `json:"id"`
	Token0  tempToken
	Token1  tempToken
	FeeTier string `json:"feeTier"`
}

type tempToken struct {
	ID string `json:"id"`
}

type PoolsData struct {
	Pools []SubgraphPool `json:"pools"`
}

type PoolsResponse struct {
	Data PoolsData `json:"data"`
}

const uniswapV3AddressString = "0x1f98431c8ad98523631ae4a59f267346ea31f984"
const UniswapV3SubgraphURL = "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v3"

const filePath = "uniswapV3Pools.json"

func GetAllPoolsAndWriteToJson(client *ethclient.Client) ([]*Pool, error) {
	pools, err := GetPools(client)
	if err != nil {
		return nil, err
	}
	fmt.Println("Number of Uniswap v3 pools:", len(pools))
	// Convert the slice of Pool structs to JSON data
	jsonData, err := json.Marshal(pools)
	if err != nil {
		return nil, err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

func GetPools(client *ethclient.Client) ([]*Pool, error) {

	//get all pool addresses
	pools, err := GetAllUniswapV3PoolsFromSubGraph()
	if err != nil {
		return nil, err
	}
	fmt.Println("Number of pools:", len(pools))

	//get the rest of the data for the ppol
	err = UpdatePools(pools, client)
	if err != nil {
		return nil, err
	}

	//check for pools with zero tick spacing, return error if any
	for _, pool := range pools {
		if pool.TickSpacing <= 0 {
			fmt.Println("Pool with zero tick spacing:", pool.Address.Hex())
		}
	}

	return pools, nil
}

func GetAllUniswapV3PoolsFromSubGraph() ([]*Pool, error) {
	var allPools []SubgraphPool
	skip := 0
	limit := 100 // Batch size, adjust according to your needs

	for {
		query := fmt.Sprintf(`
			{
				pools(first: %d, skip: %d) {
					id
					token0 {
						id
					}
					token1 {
						id
					}
					feeTier
				}
			}
		`, limit, skip)

		pools, err := fetchPools(query)
		if err != nil {
			return nil, err
		}

		if len(pools) == 0 {
			break
		}

		allPools = append(allPools, pools...)
		skip += limit
	}

	//convert all subgraph pools to pools
	var pools []*Pool
	for _, subgraphPool := range allPools {
		FeeTier, err := strconv.ParseUint(subgraphPool.FeeTier, 10, 16)
		if err != nil {
			return nil, err
		}
		pool := Pool{
			Address:        common.HexToAddress(subgraphPool.ID),
			FactoryAddress: common.HexToAddress(uniswapV3AddressString),
			Token0:         common.HexToAddress(subgraphPool.Token0.ID),
			Token1:         common.HexToAddress(subgraphPool.Token1.ID),
			Fee:            uint32(FeeTier),
			TickSpacing:    int(TickSpacingFromFeeTier(uint32(FeeTier))),
		}
		pools = append(pools, &pool)
	}

	return pools, nil
}

func TickSpacingFromFeeTier(feeTier uint32) int {
	var tickSpacing int

	switch feeTier {
	case 500: // 0.05% fee tier
		tickSpacing = 10
	case 3000: // 0.30% fee tier
		tickSpacing = 60
	case 10000: // 1.00% fee tier
		tickSpacing = 200
	default:
		fmt.Printf("Invalid fee tier: %d\n", feeTier)
		return 0
	}

	return tickSpacing
}

func fetchPools(query string) ([]SubgraphPool, error) {
	reqBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(UniswapV3SubgraphURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var poolsResp PoolsResponse
	if err := json.Unmarshal(body, &poolsResp); err != nil {
		return nil, err
	}

	return poolsResp.Data.Pools, nil
}
