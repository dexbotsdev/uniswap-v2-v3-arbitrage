package uniswap_v2

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGetBestAmountInFromPathUniV2(t *testing.T) {
	// create a test path with known values
	path := Path{
		Pools: []Pool{
			{
				Reserve0: big.NewInt(100),
				Reserve1: big.NewInt(200),
			},
			{
				Reserve0: big.NewInt(150),
				Reserve1: big.NewInt(175),
			},
			{
				Reserve0: big.NewInt(75),
				Reserve1: big.NewInt(225),
			},
		},
	}

	//set zeroForOnes

	// call the function being tested
	amountIn, err := GetBestAmountIn(path.Pools, []bool{true, true, true})
	if err != nil {
		fmt.Println("error getting best amount in: ", err)
	}
	fmt.Println("amountIn", amountIn)

	expectedAmountIn := big.NewInt(31)
	if amountIn.Cmp(expectedAmountIn) != 0 {
		t.Errorf("Expected amount in to be %v but got %v", expectedAmountIn, amountIn)
	}
}

//test vurtual reserve helper function
func TestCalculateVirtualReservesHelper(t *testing.T) {
	reserve0 := big.NewFloat(100)
	reserve1 := big.NewFloat(50)
	reserve2 := big.NewFloat(200)
	reserve3 := big.NewFloat(150)

	virtualReserve0, virtualReserve1 := calculateVirtualReservesHelper(reserve0, reserve1, reserve2, reserve3)

	expectedV0 := big.NewFloat(47)
	expectedV1 := big.NewFloat(74)
	if virtualReserve0.Cmp(expectedV0) != 0 {
		t.Errorf("calculateVirtualReservesHelper failed: expected v0 %v, but got %v", expectedV0, virtualReserve0)
	}

	if virtualReserve1.Cmp(expectedV1) != 0 {
		t.Errorf("calculateVirtualReservesHelper failed: expected v1 %v, but got %v", expectedV1, virtualReserve1)
	}
}
