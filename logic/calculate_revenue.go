package logic

import (
	"fmt"
	"math/big"
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
)

//if hasduplicatepools == true; calculate iteratively with temp state
//if hasduplicatepools == false; hasUniswapV3Pools == true; calculate iteratively without temp state

func CalculateRevenue(path Path, amountIn *big.Int) (*big.Int, error) {
	if path.HasDuplicatePools {
		//return nil, fmt.Errorf("calculateRevenue: path has duplicate pools; this is not supported yet")
		return big.NewInt(0), nil
		return path.CalculateRevenueWithDuplicatePools(amountIn)
	}

	return path.CalculateRevenueWithoutDuplicatePools(amountIn)
}
func (path *Path) CalculateRevenueWithoutDuplicatePools(amountIn *big.Int) (*big.Int, error) {
	//fmt.Println("CalculateRevenueWithoutDuplicatePools: ")

	fmt.Println("path zero for ones: ", path.ZeroForOnes)

	currAmount := new(big.Int).Set(amountIn)
	for i := 0; i < len(path.Pools); i++ {
		//print out protocol of pools in path
		fmt.Println("protocol: ", path.Pools[i].GetType())
		currAmountTemp, err := path.Pools[i].GetCopyInterface().GetAmountOut(currAmount, path.ZeroForOnes[i])

		if err != nil {
			fmt.Println("err on pathid: ", path.Id)
			fmt.Println("err on pool: ", path.Pools[i].GetAddress().Hex())
			return nil, err
		}
		currAmount.Set(currAmountTemp)
		fmt.Println("pool: ", i)
		fmt.Println("protocol: ", path.Pools[i].GetType())
		fmt.Println("Pool: ", path.Pools[i].GetAddress().Hex())
		fmt.Println("amountOut: ", currAmount)
	}

	revenue := new(big.Int).Sub(currAmount, amountIn)

	return revenue, nil
}

func (path *Path) CalculateRevenueWithDuplicatePools(amountIn *big.Int) (*big.Int, error) {
	//fmt.Println("CalculateRevenueWithDuplicatePools: ")
	//setup temp state for each pool in path and does not create duplicates for the same pool address
	poolCopies := make([]types.PoolInterface, len(path.Pools))
	poolCopyMap := make(map[common.Address]types.PoolInterface)
	for i, pool := range path.Pools {
		addr := pool.GetAddress()
		if copy, ok := poolCopyMap[addr]; ok {
			// If a copy already exists, use the existing copy
			poolCopies[i] = copy
		} else {
			// If a copy doesn't exist, create a new copy and store it in the map
			newCopy := pool.GetCopyInterface()
			poolCopyMap[addr] = newCopy
			poolCopies[i] = newCopy
		}
	}

	//for each pool in path run get amount out and update temp state
	currAmount := new(big.Int).Set(amountIn)
	for i := 0; i < len(poolCopies); i++ {
		amountOutTemp, err := poolCopies[i].GetCopyInterface().GetAmountOutAndUpdatePool(currAmount, false)
		if err != nil {
			return nil, err
		}
		amountOut := new(big.Int).Set(amountOutTemp)
		currAmount.Set(amountOut)
	}

	revenue := new(big.Int).Sub(currAmount, amountIn)

	return revenue, nil
}
