package uniswap_v3

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckTickSpacing(t *testing.T) {
	//check 0x29F0096512B4af1d689c1a11A867A6e707a8DcDe
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	// err = godotenv.Load(".env")
	// fmt.Println("err: ", err)
	// assert.NoError(t, err)

	// client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	// fmt.Println("err: ", err)
	// assert.NoError(t, err)

	pools, err := ReadFilteredPoolsFromFile()

	for _, pool := range pools {
		if pool.TickSpacing <= 0 {
			fmt.Println("pool address", pool.Address)
			fmt.Println("tick spacing", pool.TickSpacing)
		}
	}
}
