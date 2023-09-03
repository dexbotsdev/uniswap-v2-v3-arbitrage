package handle_new_block

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestHandleNewBlock(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	// Fetch the block. Here, we're fetching block number 10 as an example.
	blockNumber := big.NewInt(10) // change this to the block number you're interested in
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("Failed to fetch block: %v", err)
	}

	// Print block information
	fmt.Printf("Block number: %d\n", block.Number().Uint64())
	fmt.Printf("Block timestamp: %d\n", block.Time())
	// ... You can access other properties of the block as needed
}
