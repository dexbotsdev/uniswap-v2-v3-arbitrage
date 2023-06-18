package uniswap_v2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Path struct {
	Pools       []Pool
	BaseToken   common.Address
	AmountOut   *big.Int
	AmountIn    *big.Int
	Revenue     *big.Int //only used if startToken == endToken
	BlockNumber int
}
