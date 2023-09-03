package path

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestCalculateRevenueWithQuoter(t *testing.T) {
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

	//filter for only has v3
	for i := 0; i < len(paths); i++ {
		if !paths[i].HasUniswapV3Pools {
			//revmoe path
			paths = append(paths[:i], paths[i+1:]...)
		}
	}

	//test only the first 100 paths
	paths = paths[0:100]

	revenues := make([]*big.Int, len(paths))
	revenuesWithQuoter := make([]*big.Int, len(paths))

	//for each path calculate revenue
	for i := 0; i < len(paths); i++ {
		//skip paths with duplicate pools
		if paths[i].HasDuplicatePools {
			continue
		}

		err := paths[i].UpdatePools(client)

		path := paths[i]
		amountIn := big.NewInt(1000000000000000000)
		revenue, err := path.CalculateRevenue(amountIn)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("err: ", err)
		assert.NoError(t, err)

		revenueWithQuoter, err := path.CalculateRevenueWithQuoter(amountIn, client)
		if err != nil {
			panic(err)
		}
		fmt.Println("err: ", err)
		assert.NoError(t, err)

		fmt.Println("revenue: ", revenue)

		fmt.Println("revenueWithQuoter: ", revenueWithQuoter)

		revenues[i] = revenue
		revenuesWithQuoter[i] = revenueWithQuoter
	}

	//print revenue and revenue with quoter
	for i := 0; i < len(revenues); i++ {
		fmt.Println("revenue: ", revenues[i])
		fmt.Println("revenueWithQuoter: ", revenuesWithQuoter[i])
		if revenues[i].Cmp(revenuesWithQuoter[i]) != 0 {
			fmt.Println("revenue and revenueWithQuoter are not equal")
		}
	}

}
