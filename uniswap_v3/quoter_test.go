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

func TestQuoter(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//get all filtered pools
	filteredPools, err := ReadFilteredPoolsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	pools := filteredPools[0:100]

	amountIn, _ := new(big.Int).SetString("100", 10)

	for i := 0; i < len(pools[0].Ticks); i++ {
		amountOut, err := GetAmountOutWithQuoter(*pools[i], amountIn, true, client)
		if err != nil {
			fmt.Println("err: ", err)
		}
		fmt.Println("amountOut: ", amountOut)

	}
}
