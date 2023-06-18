package logic

import (
	"fmt"
	"math/big"
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Path struct {
	Id                int
	Pools             []types.PoolInterface
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

func (path *Path) UpdatePools(client *ethclient.Client) error {
	for i := 0; i < len(path.Pools); i++ {
		pool := path.Pools[i]
		err := pool.Update(client)
		if err != nil {
			return err
		}
	}
	return nil
}

func (path *Path) SetZeroForOnes() error {
	zeroForOnes := make([]bool, len(path.Pools))
	currToken := path.BaseToken
	for i := 0; i < len(path.Pools); i++ {
		if path.Pools[i].GetTokens()[0] == currToken {
			zeroForOnes[i] = true
			currToken = path.Pools[i].GetTokens()[1]
		} else if path.Pools[i].GetTokens()[1] == currToken {
			zeroForOnes[i] = false
			currToken = path.Pools[i].GetTokens()[0]
		}
	}
	path.ZeroForOnes = zeroForOnes
	fmt.Println("zeroForOnes", zeroForOnes)
	return nil
}

func (path *Path) SetIsV2s() error {
	isV2s := make([]bool, len(path.Pools))
	for i, pool := range path.Pools {
		isV2s[i] = pool.GetType() == "uniswap_v2"
	}
	path.IsV2s = isV2s
	return nil
}

func (path *Path) CalculateRevenue(amountIn *big.Int) (*big.Int, error) {
	return CalculateRevenue(*path, amountIn)
}

func (path *Path) SetBestAmountInAndRevenue(baseFee *big.Int) (*big.Int, *big.Int, error) {
	amountIn, revenue, err := GetBestAmountInAndRevenueForPath(*path, baseFee)
	if err != nil {
		return nil, nil, err
	}

	path.AmountIn = new(big.Int).Set(amountIn)
	path.Revenue = new(big.Int).Set(revenue)
	return amountIn, revenue, nil
}

func (path *Path) SetHasUniswapV3Pools() {
	for _, pool := range path.Pools {
		if pool.GetType() == "uniswap_v3" {
			path.HasUniswapV3Pools = true
			return
		}
	}
}

func (path *Path) SetHasDuplicatePools() {
	for i, pool := range path.Pools {
		for j, pool2 := range path.Pools {
			if i != j && pool.GetAddress() == pool2.GetAddress() {
				path.HasDuplicatePools = true
				return
			}
		}
	}
}
