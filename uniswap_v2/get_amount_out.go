package uniswap_v2

import (
	"math/big"
)

func GetAmountOut(amountIn, reserveIn, reserveOut *big.Int) (*big.Int, error) {

	amountInWithFee := new(big.Int).Mul(amountIn, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(reserveIn, big.NewInt(1000)), amountInWithFee)
	amountOut := new(big.Int).Div(numerator, denominator)

	return amountOut, nil
}
