package uniswap_v3

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCheckPoolsFor0TickSpacing(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	// err = godotenv.Load(".env")
	// fmt.Println("err: ", err)
	// assert.NoError(t, err)

	// client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	// fmt.Println("err: ", err)
	// assert.NoError(t, err)

	pools, err := ReadPoolsFromFile()

	for _, pool := range pools {
		fmt.Println("pool address", pool.Address)
		fmt.Println("tick spacing", pool.TickSpacing)
		assert.Greater(t, pool.TickSpacing, int(0))
	}

	fmt.Println("err: ", err)
}

func TestGetAllPoolsAndWriteToJson(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	pools, err := GetAllPoolsAndWriteToJson(client)

	//check that tickspacing is greater than 0
	for _, pool := range pools {
		assert.Greater(t, pool.TickSpacing, int(0))
	}

	fmt.Println("err: ", err)
	assert.NoError(t, err)
}
