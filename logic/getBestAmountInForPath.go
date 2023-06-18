package logic

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"mev-template-go/uniswap_v2"
	"os"
)

func GetBestAmountInAndRevenueForPath(path Path, baseFee *big.Int) (*big.Int, *big.Int, error) {
	//if path has uniswap v3 pools or has duplicate pools, run iteratively
	fmt.Println("path.HasUniswapV3Pools: ", path.HasUniswapV3Pools)
	fmt.Println("path.HasDuplicatePools: ", path.HasDuplicatePools)
	//print protocol for all pools
	for i := 0; i < len(path.Pools); i++ {
		fmt.Println("path.Pools[i].Protocol: ", path.Pools[i].GetType())
	}
	if path.HasUniswapV3Pools || path.HasDuplicatePools {
		return GetBestAmountInAndRevenueForPathIteratively(path, baseFee)

	}

	//else run thorugh v2 formula
	bestAmountInTemp, err := GetBestAmountInUsingUniswapV2Formula(path)
	if err != nil {
		return nil, nil, err
	}
	bestAmountIn := new(big.Int).Set(bestAmountInTemp)

	revenueTemp, err := path.CalculateRevenue(bestAmountIn)
	if err != nil {
		fmt.Println("bestAmountIn: ", bestAmountIn.String())
		fmt.Println(path.Pools[0].String())
		return nil, nil, err
	}
	revenue := new(big.Int).Set(revenueTemp)

	return bestAmountIn, revenue, nil
}

func GetBestAmountInAndRevenueForPathIteratively(path Path, baseFee *big.Int) (bestAmountIn *big.Int, revenue *big.Int, err error) {
	//run path through gss

	//define function for gss

	//run gss

	//set lower bound
	//gasFee := 21000 + 100000*len(path.Pools)
	//lowerBound := basefee*gasFee
	gasFee := big.NewInt(int64(21000 + 100000*len(path.Pools)))
	lowerBoundInt := new(big.Int).Mul(baseFee, gasFee)
	lowerBoundFloat := new(big.Float).SetInt(lowerBoundInt)

	//get upperboudn with exponential search
	intF := func(x *big.Int) *big.Int {
		revenueTemp, err := path.CalculateRevenue(x)
		if err != nil {
			//if errors.Is(err, (types.AmountInGreaterThanReserveInError)) {
			// if err == entities.ErrSqrtPriceLimitX96TooHigh {
			// 	return big.NewInt(0)
			// }
			return big.NewInt(0)
			//panic(err)
		}
		return new(big.Int).Set(revenueTemp)
	}
	upperBoundTemp, err := ExponentialSearchForUpperBound(intF, lowerBoundInt, int64(2))
	if err != nil {
		return nil, nil, err
	}

	//if upperbound is greater than uint256 max, return uint256 max
	//fmt.Println("upperBoundTemp: ", upperBoundTemp.String())

	upperBoundFloat := new(big.Float).SetInt(upperBoundTemp)
	fmt.Println("upperBoundFloat: ", upperBoundFloat.String())

	//define function for gss
	floatF := func(x *big.Float) *big.Float {
		newInt := new(big.Int)
		x.Int(newInt)
		revenueTemp, err := path.CalculateRevenue(newInt)
		if err != nil {
			// if err == entities.ErrSqrtPriceLimitX96TooHigh {
			// 	return big.NewFloat(0)
			// }
			return big.NewFloat(0)
		}
		return new(big.Float).SetInt(revenueTemp)
	}

	//run gss
	logger := log.New(os.Stdout, "", 0)
	a, b := BigGss(floatF, lowerBoundFloat, upperBoundFloat, big.NewFloat(1e-6), logger)

	//if a and b are negative, return 0
	if a.Sign() == -1 && b.Sign() == -1 {
		return big.NewInt(0), big.NewInt(0), nil
	}

	//gss returns 2 values, both are candidates for best amount in
	aInt := new(big.Int)
	a.Int(aInt)
	bInt := new(big.Int)
	b.Int(bInt)

	aRevTemp, err := path.CalculateRevenue(aInt)
	if err != nil {
		log.Fatal(err)
	}
	aRev := new(big.Int).Set(aRevTemp)
	bRevTemp, err := path.CalculateRevenue(bInt)
	if err != nil {
		log.Fatal(err)
	}
	bRev := new(big.Int).Set(bRevTemp)

	//return the greater amount
	if aRev.Cmp(bRev) < 0 {
		return bInt, bRev, nil
	}

	return aInt, aRev, nil
}

func GetBestAmountInUsingUniswapV2Formula(path Path) (*big.Int, error) {
	//convert all pools to v2 pools
	v2Pools := make([]uniswap_v2.Pool, len(path.Pools))
	for i, pool := range path.Pools {
		newV2Pool, ok := pool.(*uniswap_v2.Pool)
		if ok {
			v2Pools[i] = *newV2Pool.GetCopy()
		} else {
			return nil, errors.New("failed to convert pool to v2 pool")
		}
	}

	//calculate best amount in
	bestAmountInTemp, err := uniswap_v2.GetBestAmountIn(v2Pools, path.ZeroForOnes)
	if err != nil {
		return nil, err
	}
	bestAmountIn := new(big.Int).Set(bestAmountInTemp)

	return bestAmountIn, nil
}
