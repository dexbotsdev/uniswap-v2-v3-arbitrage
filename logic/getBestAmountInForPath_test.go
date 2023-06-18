package logic

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGetBestAmountInAndRevenueForV3Path(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	lastetBlock, err := client.BlockByNumber(context.Background(), nil)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	header := lastetBlock.Header()

	baseFee := header.BaseFee

	//get header

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	path := paths[0]

	path.UpdatePools(client)

	fmt.Println("path: ", path)

	//path.SetZeroForOnes()

	fmt.Println("path.ZeroForOnes: ", path.ZeroForOnes)

	//for each path calculate revenue

	bestAmountIn, revenue, err := path.SetBestAmountInAndRevenue(baseFee)
	if err != nil {
		panic(err)
	}
	fmt.Println("err: ", err)
	assert.NoError(t, err)
	fmt.Println("bestAmountIn: ", bestAmountIn)
	fmt.Println("revenue: ", revenue)

}

func TestGetBestAmountInAndRevenueForPath1(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	lastetBlock, err := client.BlockByNumber(context.Background(), nil)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	header := lastetBlock.Header()

	baseFee := header.BaseFee

	//get header

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	path := paths[1]

	path.UpdatePools(client)

	fmt.Println("path: ", path)

	//path.SetZeroForOnes()

	fmt.Println("path.ZeroForOnes: ", path.ZeroForOnes)

	//for each path calculate revenue

	bestAmountIn, revenue, err := path.SetBestAmountInAndRevenue(baseFee)
	if err != nil {
		panic(err)
	}
	fmt.Println("err: ", err)
	assert.NoError(t, err)
	fmt.Println("bestAmountIn: ", bestAmountIn)
	fmt.Println("revenue: ", revenue)

}

func TestGetBestAmountInAndRevenueFor100Paths(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	lastetBlock, err := client.BlockByNumber(context.Background(), nil)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	header := lastetBlock.Header()

	baseFee := header.BaseFee

	//get header

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	paths = paths[:100]

	for i := 0; i < len(paths); i++ {
		fmt.Println(paths[i].HasUniswapV3Pools)
	}

	//for each path calculate revenue
	for i := 0; i < len(paths); i++ {
		bestAmountIn, revenue, err := paths[i].SetBestAmountInAndRevenue(baseFee)
		if err != nil {
			panic(err)
		}
		fmt.Println("err: ", err)
		assert.NoError(t, err)
		fmt.Println("bestAmountIn: ", bestAmountIn)
		fmt.Println("revenue: ", revenue)

	}

	//print all paths, bestAmountIn, revenue
	for i := 0; i < len(paths); i++ {
		fmt.Println("path: ", paths[i])
		fmt.Println("bestAmountIn: ", paths[i].AmountIn)
		fmt.Println("revenue: ", paths[i].Revenue)
	}

}

func TestGetBestAmountInForPath(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	lastetBlock, err := client.BlockByNumber(context.Background(), nil)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	header := lastetBlock.Header()

	baseFee := header.BaseFee

	//get header

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//fitler paths for has uniswap v3
	uniswapV3Paths := make([]Path, 0)
	for i := 0; i < len(paths); i++ {
		if paths[i].HasUniswapV3Pools {
			uniswapV3Paths = append(uniswapV3Paths, paths[i])
		}
	}

	fmt.Println("uniswapV3Paths length ", len(uniswapV3Paths))

	//for each path calculate revenue
	for i := 0; i < len(uniswapV3Paths); i++ {
		bestAmountIn, revenue, err := GetBestAmountInAndRevenueForPath(uniswapV3Paths[i], baseFee)
		if err != nil {
			panic(err)
		}
		fmt.Println("err: ", err)
		assert.NoError(t, err)
		fmt.Println("revenue: ", bestAmountIn)
		fmt.Println("revenue: ", revenue)
	}
}

func TestForHasUniswapV3Pools(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//fitler paths for has uniswap v3
	uniswapV3Paths := make([]Path, 0)
	for i := 0; i < len(paths); i++ {
		if paths[i].HasUniswapV3Pools {
			uniswapV3Paths = append(uniswapV3Paths, paths[i])
		}
	}

	//get length of v2 only paths

	fmt.Println("uniswapV3Paths length ", len(uniswapV3Paths))
}
