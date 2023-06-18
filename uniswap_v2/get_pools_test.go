package uniswap_v2

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetAllUniswapV2ForkPoolsAndWriteToFile(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	pools, err := GetAllUniswapV2ForkPoolsAndWriteToFile(client)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//print length of pools
	fmt.Println("len(pools): ", len(pools))
}
