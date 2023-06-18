package uniswap_v3

import (
	"math/big"
)

type Tick struct {
	Index          int
	LiquidityNet   *big.Int
	LiquidityGross *big.Int
}

func (tick *Tick) GetCopy() Tick {
	return Tick{
		Index:          tick.Index,
		LiquidityNet:   new(big.Int).Set(tick.LiquidityNet),
		LiquidityGross: new(big.Int).Set(tick.LiquidityGross),
	}
}

func (tick *Tick) Set(newTick Tick) {
	tick.Index = newTick.Index
	tick.LiquidityNet.Set(newTick.LiquidityNet)
	tick.LiquidityGross.Set(newTick.LiquidityGross)
}
