package logic

import (
	"fmt"
	"math/big"
	"mev-template-go/uniswap_v3"

	"github.com/ethereum/go-ethereum/ethclient"
)

//if hasduplicatepools == true; calculate iteratively with temp state
//if hasduplicatepools == false; hasUniswapV3Pools == true; calculate iteratively without temp state

func (path *Path) CalculateRevenueWithQuoter(amountIn *big.Int, client *ethclient.Client) (*big.Int, error) {
	//fmt.Println("CalculateRevenueWithoutDuplicatePools: ")

	fmt.Println("path zero for ones: ", path.ZeroForOnes)

	currAmount := new(big.Int).Set(amountIn)
	for i := 0; i < len(path.Pools); i++ {
		//print out protocol of pools in path
		fmt.Println("protocol: ", path.Pools[i].GetType())

		if path.Pools[i].GetType() == "uniswap_v2" {
			
			currAmountTemp, err := path.Pools[i].GetCopyInterface().GetAmountOut(currAmount, path.ZeroForOnes[i])
			if err != nil {
				return nil, err
			}
			currAmount.Set(currAmountTemp)

		} else {

			v3Pool, ok := path.Pools[i].GetCopyInterface().(*uniswap_v3.Pool)
			if !ok {
				return nil, fmt.Errorf("pool is not uniswap_v3.Pool")
			}
			currAmountTemp, err := uniswap_v3.GetAmountOutWithQuoter(*v3Pool, currAmount, path.ZeroForOnes[i], client)
			if err != nil {
				fmt.Println("err on pathid: ", path.Id)
				fmt.Println("err on pool: ", path.Pools[i].GetAddress().Hex())
				return nil, err
			}
			currAmount.Set(currAmountTemp)

		}

		//if v2 use v2 getamountout

		//if v3 use v3 getamountoutWithQuoter
		fmt.Println("pool: ", i)
		fmt.Println("protocol: ", path.Pools[i].GetType())
		fmt.Println("Pool: ", path.Pools[i].GetAddress().Hex())
		fmt.Println("amountOut: ", currAmount)
	}

	revenue := new(big.Int).Sub(currAmount, amountIn)

	return revenue, nil
}
