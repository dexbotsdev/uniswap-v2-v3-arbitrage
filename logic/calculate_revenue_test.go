package logic

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCalculateRevenueWithQuoterOnPath6(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//for each path calculate revenue

	path := paths[6]
	amountIn := big.NewInt(1000000000000000000)
	revenue, err := path.CalculateRevenue(amountIn)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("err: ", err)
	assert.NoError(t, err)
	fmt.Println("revenue: ", revenue)

}

func TestCalculateRevenueWIthQuoter(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//setup eth client
	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	paths, err := ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//for each path calculate revenue and calculatre revenue with quoter
	for i := 0; i < len(paths); i++ {
		path := paths[i]
		amountIn := big.NewInt(1000000000000000000)
		revenue, err := path.CalculateRevenue(amountIn)
		if err != nil {
			panic(err)
		}
		fmt.Println("err: ", err)
		assert.NoError(t, err)
		fmt.Println("revenue: ", revenue)

		revenueWithQuoter, err := path.CalculateRevenueWithQuoter(amountIn, client)
		if err != nil {
			panic(err)
		}
		fmt.Println("err: ", err)
		assert.NoError(t, err)
		fmt.Println("revenueWithQuoter: ", revenueWithQuoter)
	}

}
