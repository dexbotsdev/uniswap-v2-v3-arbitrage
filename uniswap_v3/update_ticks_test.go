package uniswap_v3

import (
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTicksOnPool(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//create test pool
	poolAddress := common.HexToAddress("0x60594a405d53811d3BC4766596EFD80fd545A270")
	pool := Pool{
		Address: poolAddress,
	}

	err = UpdateAllTicksForPool(&pool, client)
	fmt.Println("err: ", err)
	assert.NoError(t, err)
}

func TestUpdateTicksMulticall(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//create test pool
	poolAddress := common.HexToAddress("0x60594a405d53811d3BC4766596EFD80fd545A270")
	pool := Pool{
		Address: poolAddress,
	}

	poolArr := []*Pool{&pool}

	//update ticks
	err = updateTicksBatched(poolArr, client)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

}

func TestUpdateAllTicksMulticall(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//read pools from json
	pools, err := ReadPoolsFromFile()

	//convert to pointers
	poolPointers := []*Pool{}
	for _, pool := range pools {
		poolPointers = append(poolPointers, pool)
	}

	//update ticks
	err = updateTicksBatched(poolPointers, client)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

}
