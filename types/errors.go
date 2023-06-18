package types

import (
	"errors"
)

var AmountInGreaterThanReserveInError = errors.New("amountIn is greater than reserveIn")

var UniswapV2BestAmountInHasDuplicatePoolsError = errors.New("Uniswap V2 BestAmountIn: has duplicate pools")
