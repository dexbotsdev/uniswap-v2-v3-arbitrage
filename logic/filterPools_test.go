package logic

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterPoolsAndWriteToFile(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)
	//read files
	err = FilterPoolsAndWriteToFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)
}

func TestGetFileteredPools(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	pools, err := GetFilteredPools()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//print all pools
	for _, pool := range pools {
		fmt.Println("pool address", pool.GetAddress().Hex())
		fmt.Println("pool type", pool.GetType())
		fmt.Println("pool token0", pool.GetTokens()[0].Hex())
		fmt.Println("pool token1", pool.GetTokens()[1].Hex())
		if pool.GetType() == "uniswap_v3" {
			fmt.Println("v3 pool")
		}
	}

	fmt.Println("len(pools): ", len(pools))
}
