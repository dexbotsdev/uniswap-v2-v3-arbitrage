package uniswap_v2

import (
	"fmt"
	"math/big"
)

//type tempPools struct {
//	address common.Address
//	Token0 common.Address
//	Token1 common.Address
//	Reserve0 *big.Int
//	Reserve1 *big.Int
//}

//calculate revenue by using the get amount out function on each pair
func CalculateRevenue(path Path) *big.Int {
	fmt.Println("FUNCTION CalculateRevenue")
	fmt.Println("path.AmountIn", path.AmountIn)
	currentAmount := new(big.Int).Set(path.AmountIn)
	currToken := path.BaseToken
	for i := 0; i < len(path.Pools); i++ {
		if currToken == path.Pools[i].Token0 { //not flipped
			currentAmountTemp, err := GetAmountOut(currentAmount, path.Pools[i].Reserve0, path.Pools[i].Reserve1)
			if err != nil {
				fmt.Println("error getting amount out: ", err)
			}
			currentAmount.Set(currentAmountTemp)
			currToken = path.Pools[i].Token1
		} else { //flipped
			currentAmountTemp, err := GetAmountOut(currentAmount, path.Pools[i].Reserve1, path.Pools[i].Reserve0)
			if err != nil {
				fmt.Println("error getting amount out: ", err)
			}
			currentAmount.Set(currentAmountTemp)
			currToken = path.Pools[i].Token0
		}
	}
	return new(big.Int).Sub(currentAmount, path.AmountIn)
}
