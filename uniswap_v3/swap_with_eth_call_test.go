package uniswap_v3

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetAmountOutWithEthCall(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//client
	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//get all pools
	pools, err := ReadFilteredPoolsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//for each pool get amount out
	for i := 0; i < len(pools); i++ {
		pool := pools[i]
		amountIn := big.NewInt(1000000)

		//run getamountout
		amountOut, err := pool.GetAmountOut(amountIn, true)
		if err != nil {
			panic(err)
		}

		amountOutWithEthCall, err := GetAmountOutWithEthCall(*pool, amountIn, true, client)
		if err != nil {
			panic(err)
		}

		fmt.Println("err: ", err)

		fmt.Println("amountOut: ", amountOut)
		fmt.Println("amountOutWithEthCall: ", amountOutWithEthCall)

		assert.Equal(t, amountOut, amountOutWithEthCall)

	}
}
