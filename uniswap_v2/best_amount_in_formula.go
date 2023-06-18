package uniswap_v2

import (
	"errors"
	"fmt"
	"math/big"
)

func GetBestAmountIn(pathPools []Pool, zeroForOnes []bool) (*big.Int, error) {
	//we dont use virtual reserves to calculate revenue as rounding errors can cause revenue to be inprecise with what will happen on chain
	//we will use the derived best amount in via get amount out/revenue

	//print all pools
	fmt.Println("pools", pathPools)
	for i := 0; i < len(pathPools); i++ {
		fmt.Println("pool", i)
		fmt.Println("pool", pathPools[i].String())
	}

	if hasEmptyReserves(pathPools) {
		return big.NewInt(0), errors.New("path has empty reserves")
	}

	r := big.NewFloat(0.997)

	//calculate virtual reserves
	virtualReserve0Temp, virtualReserve1Temp := calculateVirtualReserves(pathPools, zeroForOnes)
	virtualReserve0 := new(big.Float).Set(virtualReserve0Temp)
	virtualReserve1 := new(big.Float).Set(virtualReserve1Temp)
	fmt.Println("virtualReserve0", virtualReserve0.Text('f', -1))
	fmt.Println("virtualReserve1", virtualReserve1.Text('f', -1))

	// calculate optimum input
	//note: virtualReserve0 < virtualReserve1 will make bestAmountIn positive (ignoring fee)
	//const bestAmountIn = ((virturalReserve0.times(virturalReserve1.times(fee)).sqrt().minus(virturalReserve0))).div(fee)
	//bestAmountIn = ((virturalReserve0*virturalReserve1.times*fee).sqrt - virturalReserve0)/0.997
	//bestAmountIn = ((virturalReserve0*virturalReserve1.times*fee)**0.5 - virturalReserve0)/fee
	//bestAmountIn = ((virturalReserve0*virturalReserve1.times*0.997)**0.5 - virturalReserve0)/0.997
	fmt.Println("calculating bestAmountIn")
	fmt.Println("virtualReserve0", virtualReserve0.String())
	fmt.Println("virtualReserve1", virtualReserve1.String())
	bestAmountIn := new(big.Float).Mul(virtualReserve0, virtualReserve1.Mul(virtualReserve1, r))
	fmt.Println(bestAmountIn.String())
	bestAmountIn.Sqrt(bestAmountIn)
	fmt.Println(bestAmountIn.String())
	bestAmountIn.Sub(bestAmountIn, virtualReserve0) //-E_a
	fmt.Println(bestAmountIn.String())
	//applying 1/fee
	bestAmountIn.Quo(bestAmountIn, r)
	fmt.Println(bestAmountIn.String())

	fmt.Println("virtualReserve0", virtualReserve0.String())
	fmt.Println("virtualReserve1", virtualReserve1.String())
	fmt.Println("bestAmountIn", bestAmountIn.String())
	fmt.Println("((", virtualReserve0.String(), "*", virtualReserve1.String(), "*0.997)**0.5 -", virtualReserve0.String(), ")/0.997")

	// return 0 if bestAmountIn  is neagtive
	if bestAmountIn.Sign() == -1 {
		return big.NewInt(0), nil
	}

	bestAmountInInt, _ := bestAmountIn.Int(new(big.Int))

	fmt.Println("bestAmountIn", bestAmountIn.String())
	return bestAmountInInt, nil
}

func calculateVirtualReserves(pathPools []Pool, zeroForOnes []bool) (*big.Float, *big.Float) {
	reserve0 := new(big.Float)
	reserve1 := new(big.Float)
	reserve2 := new(big.Float)
	reserve3 := new(big.Float)
	virtualReserve0 := new(big.Float)
	virtualReserve1 := new(big.Float)

	//TODO check index of tokens and reserves
	//get flipFlags
	//print tokens and flipflags
	for i := 0; i < len(pathPools); i++ {
		fmt.Println("pool", i)
		fmt.Println("poolAddress", pathPools[i].Address.String())
		fmt.Println("token0", pathPools[i].Token0.String())
		fmt.Println("token1", pathPools[i].Token1.String())
		fmt.Println("reserve0", pathPools[i].Reserve0.String())
		fmt.Println("reserve1", pathPools[i].Reserve1.String())
		fmt.Println("zeroForOnes", zeroForOnes[i])
	}

	if zeroForOnes[0] {
		virtualReserve0.SetInt(pathPools[0].Reserve0)
		virtualReserve1.SetInt(pathPools[0].Reserve1)
	} else {
		virtualReserve0.SetInt(pathPools[0].Reserve1)
		virtualReserve1.SetInt(pathPools[0].Reserve0)
	}
	fmt.Println("virtual reserves", virtualReserve0.String(), virtualReserve1.String())
	fmt.Printf("Type of virtualReserve0 is: %T\n", virtualReserve0)
	fmt.Println("path length", len(pathPools))

	// calculate virtual reserves for whole path
	for i := 1; i < len(pathPools); i++ {
		fmt.Println("pool", i)
		/*
		   reserve0 = virturalReserve0
		   reserve1 = virturalReserve1
		   reserve2 = path.edges[i].getReserve0()
		   reserve3 = path.edges[i].getReserve1()
		   virturalReserve0 = reserve0.times(reserve2).div(reserve2.plus(reserve1.times(fee)))
		   virturalReserve1 = fee.times(reserve1.times(reserve3)).div(reserve2.plus(reserve1.times(fee)))*/
		reserve0.Set(virtualReserve0)
		reserve1.Set(virtualReserve1)
		if zeroForOnes[i] {
			reserve2.SetInt(pathPools[i].Reserve0)
			reserve3.SetInt(pathPools[i].Reserve1)
		} else {
			reserve2.SetInt(pathPools[i].Reserve1)
			reserve3.SetInt(pathPools[i].Reserve0)
		}
		fmt.Println("reserves", reserve0.String(), reserve1.String(), reserve2.String(), reserve3.String())
		//if any of the reserves are 0 return 0,0

		virtualReserve0Temp, virtualReserve1Temp := calculateVirtualReservesHelper(reserve0, reserve1, reserve2, reserve3)
		virtualReserve0.Set(virtualReserve0Temp)
		virtualReserve1.Set(virtualReserve1Temp)
	}

	return virtualReserve0, virtualReserve1
}

func calculateVirtualReservesHelper(
	reserve0 *big.Float,
	reserve1 *big.Float,
	reserve2 *big.Float,
	reserve3 *big.Float,
) (*big.Float, *big.Float) {
	var virtualReserve0 *big.Float
	var virtualReserve1 *big.Float

	// reserve0 = virturalReserve0
	// reserve1 = virturalReserve1
	// reserve2 = path.edges[i].getReserve0()
	// reserve3 = path.edges[i].getReserve1()

	// virturalReserve0 = reserve0.times(reserve2).div(reserve2.plus(reserve1.times(fee)))
	// virturalReserve1 = fee.times(reserve1.times(reserve3)).div(reserve2.plus(reserve1.times(fee)))

	fmt.Println("reserve0", reserve0.String())
	fmt.Println("reserve1", reserve1.String())
	fmt.Println("reserve2", reserve2.String())
	fmt.Println("reserve3", reserve3.String())

	//virtualReserve0 = reserve0.times(reserve2).div(reserve2.plus(reserve1.times(fee)))
	numerator := new(big.Float).Mul(reserve0, reserve2)
	denominator := new(big.Float).Add(reserve2, ApplyFee(reserve1))
	virtualReserve0 = new(big.Float).Quo(numerator, denominator)

	//virtualReserve1 = fee.times(reserve1.times(reserve3)).div(reserve2.plus(reserve1.times(fee)))
	numerator = ApplyFee(new(big.Float).Mul(reserve1, reserve3))
	denominator = new(big.Float).Add(reserve2, ApplyFee(reserve1))
	virtualReserve1 = new(big.Float).Quo(numerator, denominator)

	fmt.Println("after calculation")

	fmt.Println("reserve0", reserve0.String())
	fmt.Println("reserve1", reserve1.String())
	fmt.Println("reserve2", reserve2.String())
	fmt.Println("reserve3", reserve3.String())

	fmt.Println("virtualReserve0", virtualReserve0.String())
	fmt.Println("virtualReserve1", virtualReserve1.String())

	return virtualReserve0, virtualReserve1
}

func hasEmptyReserves(pathPools []Pool) bool {
	for i := 0; i < len(pathPools); i++ {
		if pathPools[i].Reserve0.Sign() == 0 || pathPools[i].Reserve1.Sign() == 0 {
			return true
		}
	}
	return false
}

func ApplyFee(amount *big.Float) *big.Float {
	fee := big.NewFloat(0.997)
	return new(big.Float).Mul(amount, fee)
}
