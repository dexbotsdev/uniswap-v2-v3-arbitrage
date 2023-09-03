package executor

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"mev-template-go/path"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

//"github.com/umbracle/ethgo/abi"

func TestExecuteMixedPathWithV2V2V3V3(t *testing.T) {
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

	//read pools
	paths, err := path.ReadPathsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	profitablePath := paths[95]
	//err = UpdatePools(path.Pools, client)
	//path.SetBestAmountInAndRevenue(baseFee)
	//loop paths until last pool is v3 pool
	for i := 0; i < len(paths); i++ {
		fmt.Println("i: ", i)
		if paths[i].HasDuplicatePools {
			continue
		}

		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" {
			continue
		}

		if len(paths[i].Pools) != 4 {
			continue
		}

		if paths[i].Pools[0].GetType() == "uniswap_v2" && paths[i].Pools[1].GetType() == "uniswap_v2" && paths[i].Pools[2].GetType() == "uniswap_v3" && paths[i].Pools[3].GetType() == "uniswap_v3" {
			err = UpdatePools(paths[i].Pools, client)
			if err != nil {
				fmt.Println("err: ", err)
				t.Fatal(err)
			}
			bestAmountIn, revenue, err := paths[i].SetBestAmountInAndRevenue(baseFee)
			if err != nil {
				fmt.Println("err: ", err)
				t.Fatal(err)
			}
			if paths[i].Revenue.Cmp(big.NewInt(0)) == 1 {
				fmt.Println("bestAmountIn: ", bestAmountIn.String())
				fmt.Println("revenue: ", revenue.String())
				profitablePath = paths[i]
				break
			}
		}
	}

	amountInBefore := new(big.Int).Set(profitablePath.AmountIn)
	revenueBefore := new(big.Int).Set(profitablePath.Revenue)

	//reset revenue and amountIn after update
	err = profitablePath.UpdatePools(client)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}
	bestAmountIn, revenue, err := profitablePath.SetBestAmountInAndRevenue(baseFee)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}

	fmt.Println("bestAmountInBefore: ", amountInBefore.String())
	fmt.Println("revenueBefore: ", revenueBefore.String())
	fmt.Println("bestAmountIn: ", bestAmountIn.String())
	fmt.Println("revenue: ", revenue.String())

	//print quoterer flashswap amountout
	//convert path to path execution values

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, profitablePath.Revenue, true, true)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}

	PrintPathExecutionValues(pathExecutionValues)

	//get payload bytes struct
	flashswapPayloadBytesStruct, err := convertPathExecutionValuesToFlashswapPayloadBytes(pathExecutionValues)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	printFlashswapPayloadBytes(flashswapPayloadBytesStruct)

	//create payload
	flashswapPayload := appendFlashswapPayloadBytesToBytesArray(flashswapPayloadBytesStruct)
	fmt.Println("flashswapPayload: ", hex.EncodeToString(flashswapPayload))

	calldata, err := createFlashswapCalldata(pathExecutionValues, flashswapPayload)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//compare with quoter
	quoterRevenue, err := profitablePath.CalculateRevenueWithQuoter(profitablePath.AmountIn, client)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}
	path.PrintPath(profitablePath)
	fmt.Println("quoterRevenue: ", quoterRevenue.String())
	fmt.Println("path revenue: ", profitablePath.Revenue.String())

	//check if quoter revenue the same as path revenue
	if quoterRevenue.Cmp(profitablePath.Revenue) != 0 {
		fmt.Println("quoterRevenue: ", quoterRevenue.String())
		fmt.Println("path revenue: ", profitablePath.Revenue.String())
		t.Fatal("quoterRevenue != path revenue")
	}

	//print amountIn
	fmt.Println("amountIn: ", profitablePath.AmountIn.String())
	fmt.Println("flashswapAmountOut", pathExecutionValues.V2FlashswapAmountOut.String())
	fmt.Println("revenue: ", pathExecutionValues.Revenue.String())
	fmt.Println("quoterRevenue", quoterRevenue.String())

	//print protocols of each pool
	fmt.Println("flashswapIsV2: ", pathExecutionValues.IsV2Flashswap)
	fmt.Println("calldata: ", hex.EncodeToString(calldata))
	fmt.Println("blockNumber:", header.Number.String())
	fmt.Println("callTarget: ", pathExecutionValues.CallTarget)
	//execute
	//fmt.Println("payloadValues: ", payloadValues)

}

// type PathExecutionValues struct {
// 	FactoryAddress                  common.Address //20 bytes
// 	OtherTokenAddress               common.Address //20 bytes
// 	FeeIndex                        uint           //2 bits each, 2 for each pool; used for v3 callback verification
// 	CoinBaseTransferBool            bool           //1 bit
// 	swapDataCountMinus1             int            //3 bits; max value is 8
// 	MaxIntByteSize                  int            //5 bits; max value is 32
// 	Revenue                         *big.Int       //maxIntByteSize bytes
// 	QuarterBribePercentageOfRevenue int            //8 bits
// 	IsV2s                           []bool         //1 byte; bit each, 1 for each pool
// 	ZerosForOnes                    []bool         //1 byte; 1 bit each, 1 for each pool
// 	SwapDatas                       []SwapData     //max value is 7 as first swap is the flashswap

// }
