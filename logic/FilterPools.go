package logic

import (
	"mev-template-go/types"
	"mev-template-go/uniswap_v2"
	"mev-template-go/uniswap_v3"
)

func FilterPoolsAndWriteToFile() error {
	//read files
	_, err := uniswap_v2.FilterPoolsAndWriteToFile()
	if err != nil {
		return err
	}

	_, err = uniswap_v3.FilterPoolsAndWriteToFile()
	if err != nil {
		return err
	}

	return nil
}

func GetFilteredPools() ([]types.PoolInterface, error) {

	uniswapV2Pools, err := uniswap_v2.ReadFilteredPoolsFromFile()
	if err != nil {
		return nil, err
	}

	uniswapV3Pools, err := uniswap_v3.ReadFilteredPoolsFromFile()
	if err != nil {
		return nil, err
	}

	//combine pools
	poolInterfaces := make([]types.PoolInterface, 0)
	for i := 0; i < len(uniswapV2Pools); i++ {
		poolInterfaces = append(poolInterfaces, &uniswapV2Pools[i])
	}

	for i := 0; i < len(uniswapV3Pools); i++ {
		poolInterfaces = append(poolInterfaces, uniswapV3Pools[i])
	}

	return poolInterfaces, nil
}
