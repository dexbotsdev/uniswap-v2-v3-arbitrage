package path

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"mev-template-go/pool_interface_wrapper"

	"github.com/ethereum/go-ethereum/common"
)

// type Path struct {
// 	Pools             []interfaces.PoolInterface
// 	BaseToken         common.Address
// 	AmountOut         *big.Int
// 	AmountIn          *big.Int
// 	Revenue           *big.Int //only used if startToken == endToken
// 	BlockNumber       int
// 	HasDuplicatePools bool
// }
const pathFileName = "paths.json"

type PathWrapper struct {
	Id                int
	Pools             []pool_interface_wrapper.PoolInterfaceWrapper
	BaseToken         common.Address
	AmountOut         *big.Int
	AmountIn          *big.Int
	Revenue           *big.Int //only used if startToken == endToken
	BlockNumber       int
	HasDuplicatePools bool
	HasUniswapV3Pools bool
	ZeroForOnes       []bool
	IsV2s             []bool
}

func WrapPath(path Path) (PathWrapper, error) {
	newPath := PathWrapper{
		Id:                path.Id,
		BaseToken:         path.BaseToken,
		AmountOut:         path.AmountOut,
		AmountIn:          path.AmountIn,
		Revenue:           path.Revenue,
		BlockNumber:       path.BlockNumber,
		HasDuplicatePools: path.HasDuplicatePools,
		HasUniswapV3Pools: path.HasUniswapV3Pools,
		ZeroForOnes:       path.ZeroForOnes,
		IsV2s:             path.IsV2s,
	}
	for _, pool := range path.Pools {
		wrapPoolInterface := pool_interface_wrapper.PoolInterfaceWrapper{pool}
		newPath.Pools = append(newPath.Pools, wrapPoolInterface)
	}
	return newPath, nil
}
func UnwrapPath(path PathWrapper) (Path, error) {
	newPath := Path{
		Id:                path.Id,
		BaseToken:         path.BaseToken,
		AmountOut:         path.AmountOut,
		AmountIn:          path.AmountIn,
		Revenue:           path.Revenue,
		BlockNumber:       path.BlockNumber,
		HasDuplicatePools: path.HasDuplicatePools,
		HasUniswapV3Pools: path.HasUniswapV3Pools,
		ZeroForOnes:       path.ZeroForOnes,
		IsV2s:             path.IsV2s,
	}
	for _, pool := range path.Pools {
		newPath.Pools = append(newPath.Pools, pool.PoolInterface)
	}
	return newPath, nil
}

func ReadPathsFromFile() ([]Path, error) {
	fileContents, err := ioutil.ReadFile(pathFileName)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into a slice of Pool structs
	var wrapperPaths []PathWrapper
	err = json.Unmarshal(fileContents, &wrapperPaths)
	if err != nil {
		return nil, err
	}

	//unwrap paths
	var paths []Path
	for _, wrapperPath := range wrapperPaths {
		path, err := UnwrapPath(wrapperPath)
		if err != nil {
			return nil, err
		}
		paths = append(paths, path)
	}

	return paths, nil
}

func WritePathsToFile(paths []Path) error {
	// Wrap the pools in PoolWrapper structs
	var wrappedPaths []PathWrapper
	for _, path := range paths {
		wrappedPath, err := WrapPath(path)
		if err != nil {
			return err
		}
		wrappedPaths = append(wrappedPaths, wrappedPath)
	}

	jsonData, err := json.Marshal(wrappedPaths)
	if err != nil {
		return err
	}

	// Write the JSON data to the file
	err = ioutil.WriteFile(pathFileName, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
