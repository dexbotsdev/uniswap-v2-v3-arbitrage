package logic

import (
	"fmt"
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
)

func FindAllPathsAndWriteToFile() ([]Path, error) {
	paths, err := FindAllPaths()
	if err != nil {
		return nil, err
	}

	err = WritePathsToFile(paths)
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func createTokenToPoolMap(pools []types.PoolInterface) map[common.Address][]types.PoolInterface {
	tokenToPoolMap := make(map[common.Address][]types.PoolInterface)

	// Loop through each pool
	for _, pool := range pools {
		// Add the pool to the tokenToPoolMap for each token
		tokenToPoolMap[pool.GetTokens()[0]] = append(tokenToPoolMap[pool.GetTokens()[0]], pool)
		tokenToPoolMap[pool.GetTokens()[1]] = append(tokenToPoolMap[pool.GetTokens()[1]], pool)
	}

	return tokenToPoolMap
}

func FindAllPaths() ([]Path, error) {
	weth := common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	startToken := weth
	endToken := weth

	maxDepth := 3 //max depth of 3

	//get all pools
	poolInterfaces, err := GetFilteredPools()
	if err != nil {
		return nil, err
	}

	fmt.Println("Pools length:", len(poolInterfaces))

	//create token to pool map
	tokenToPoolMap := createTokenToPoolMap(poolInterfaces)

	foundPaths := make([]Path, 0) //empty array of paths

	//add the first pool to the path with start token
	startPools := tokenToPoolMap[startToken]
	fmt.Println("startPools", len(startPools))

	//create empty new path
	newPath := Path{Pools: []types.PoolInterface{}}
	err = dfs(newPath, endToken, maxDepth, &foundPaths, tokenToPoolMap, startToken)
	if err != nil {
		return nil, err
	}

	fmt.Println("Found", len(foundPaths), "paths")

	//put id on all paths
	for i := 0; i < len(foundPaths); i++ {
		foundPaths[i].Id = i
	}

	//set hasduplicatepools and hasuniswapv3pools and zero to ones and isv2s
	//set basetoken to weth for all paths
	for i := 0; i < len(foundPaths); i++ {
		foundPaths[i].BaseToken = weth
		foundPaths[i].SetHasDuplicatePools()
		foundPaths[i].SetHasUniswapV3Pools()
		foundPaths[i].SetZeroForOnes()
		foundPaths[i].SetIsV2s()
	}

	// for i := 0; i < len(foundPaths); i++ {
	// 	if foundPaths[i].HasDuplicatePools {
	// 		fmt.Println("Duplicate Pools")
	// 	}
	// 	if foundPaths[i].HasUniswapV3Pools {
	// 		fmt.Println("Uniswap V3 Pools")
	// 	}
	// }

	//write paths to json file
	err = WritePathsToFile(foundPaths)
	if err != nil {
		return nil, err
	}

	return foundPaths, nil
}

func dfs(currentPath Path, endToken common.Address, depth int, foundPaths *[]Path, tokenToPoolMap map[common.Address][]types.PoolInterface, currentToken common.Address) error {
	fmt.Println("DFS", currentToken.Hex(), len(currentPath.Pools))
	//if at end token, return path and current path length is greater than 1
	if currentToken == endToken && len(currentPath.Pools) > 1 {
		*foundPaths = append(*foundPaths, currentPath)
		return nil
	}
	//if at max depth, return paths
	if depth == 0 {
		return nil
	}

	//get neightboors using tokenToPoolMap by searching the for the token thats the current token
	edges := tokenToPoolMap[currentToken]

	//for each neighbor, add it to the path and call dfs
	for _, edge := range edges {
		if edge.GetTokens()[0] == currentToken {
			newPath := Path{
				Pools: make([]types.PoolInterface, len(currentPath.Pools)),
			}
			copy(newPath.Pools, currentPath.Pools)
			newPath.Pools = append(newPath.Pools, edge)
			err := dfs(newPath, endToken, depth-1, foundPaths, tokenToPoolMap, edge.GetTokens()[1])
			if err != nil {
				return err
			}
		}
		if edge.GetTokens()[1] == currentToken {
			newPath := Path{
				Pools: make([]types.PoolInterface, len(currentPath.Pools)),
			}
			copy(newPath.Pools, currentPath.Pools)
			newPath.Pools = append(newPath.Pools, edge)
			err := dfs(newPath, endToken, depth-1, foundPaths, tokenToPoolMap, edge.GetTokens()[0])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func PrintPath(path Path) {
	fmt.Println("Path id", path.Id)
	for i := 0; i < len(path.Pools); i++ {
		fmt.Println("Pool ", i)
		fmt.Println("Pool: ", path.Pools[i].GetAddress().String())
		fmt.Println("Token0: ", path.Pools[i].GetTokens()[0].String())
		fmt.Println("Token1: ", path.Pools[i].GetTokens()[1].String())
	}
}
