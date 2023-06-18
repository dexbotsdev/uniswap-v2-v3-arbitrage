package logic

import (
	"math/big"
)

func ExponentialSearchForUpperBound(f func(*big.Int) *big.Int, start *big.Int, multiplier int64) (*big.Int, error) {
	left := new(big.Int).Set(start)
	right := new(big.Int).Mul(start, big.NewInt(multiplier))

	for {
		if f(right).Cmp(f(left)) <= 0 {
			break
		}
		left.Set(right)
		right.Mul(right, big.NewInt(multiplier))
	}

	return right, nil
}
