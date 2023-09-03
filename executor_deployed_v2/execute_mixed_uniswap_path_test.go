package executor

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"mev-template-go/path"
	"mev-template-go/uniswap_v3"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

//"github.com/umbracle/ethgo/abi"
func TestExecuteMixedPathWithV2V2V2Sushiswap(t *testing.T) {
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
		if paths[i].HasDuplicatePools {
			continue
		}

		//if last swap factory is not 0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac, continue
		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" {
			continue
		}

		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" && paths[i].Pools[0].GetType() == "uniswap_v2" && paths[i].Pools[len(paths[i].Pools)-2].GetType() == "uniswap_v2" && paths[i].Pools[len(paths[i].Pools)-1].GetType() == "uniswap_v2" && paths[i].HasDuplicatePools == false {
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

	//bribe = revenue * 0.9
	bribe := new(big.Int).Mul(revenue, big.NewInt(9))
	bribe = new(big.Int).Div(bribe, big.NewInt(10))

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, bribe, false, true)
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

func TestExecuteMixedPathWithV2V2V3(t *testing.T) {
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
		if paths[i].HasDuplicatePools {
			continue
		}

		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" && paths[i].Pools[0].GetType() == "uniswap_v2" && paths[i].Pools[len(paths[i].Pools)-2].GetType() == "uniswap_v2" && paths[i].Pools[len(paths[i].Pools)-1].GetType() == "uniswap_v3" && paths[i].HasDuplicatePools == false {
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

	//bribe = revenue * 0.9
	bribe := new(big.Int).Mul(revenue, big.NewInt(9))
	bribe = new(big.Int).Div(bribe, big.NewInt(10))

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, bribe, false, true)
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
	fmt.Println("bribe: ", bribe.String())
	fmt.Println("profit: ", new(big.Int).Sub(pathExecutionValues.Revenue, bribe).String())
	//print protocols of each pool
	fmt.Println("flashswapIsV2: ", pathExecutionValues.IsV2Flashswap)
	fmt.Println("calldata: ", hex.EncodeToString(calldata))
	fmt.Println("blockNumber:", header.Number.String())
	fmt.Println("callTarget: ", pathExecutionValues.CallTarget)
	//execute
	//fmt.Println("payloadValues: ", payloadValues)

}

func TestExecuteMixedPathWithV3V2V3(t *testing.T) {
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
		if paths[i].HasDuplicatePools {
			continue
		}

		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" && paths[i].Pools[0].GetType() == "uniswap_v3" && paths[i].Pools[len(paths[i].Pools)-2].GetType() == "uniswap_v2" && paths[i].Pools[len(paths[i].Pools)-1].GetType() == "uniswap_v3" {
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

	//bribe = revenue * 0.9
	bribe := new(big.Int).Mul(revenue, big.NewInt(9))
	bribe = new(big.Int).Div(bribe, big.NewInt(10))

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, bribe, true, true)
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
	fmt.Println("bribe: ", bribe.String())
	fmt.Println("profit: ", new(big.Int).Sub(pathExecutionValues.Revenue, bribe).String())

	//print protocols of each pool
	fmt.Println("flashswapIsV2: ", pathExecutionValues.IsV2Flashswap)
	fmt.Println("calldata: ", hex.EncodeToString(calldata))
	fmt.Println("blockNumber:", header.Number.String())
	fmt.Println("callTarget: ", pathExecutionValues.CallTarget)
	//execute
	//fmt.Println("payloadValues: ", payloadValues)

}

func TestExecuteMixedPathWithV3V3V3(t *testing.T) {
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
		if paths[i].HasDuplicatePools {
			continue
		}

		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" && paths[i].Pools[len(paths[i].Pools)-1].GetType() == "uniswap_v3" && paths[i].Pools[len(paths[i].Pools)-2].GetType() == "uniswap_v3" && paths[i].Pools[0].GetType() == "uniswap_v3" && paths[i].HasDuplicatePools == false {
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

	//bribe = revenue * 0.9
	bribe := new(big.Int).Mul(revenue, big.NewInt(9))
	bribe = new(big.Int).Div(bribe, big.NewInt(10))

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, bribe, true, false)
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
	fmt.Println("bribe: ", bribe.String())
	fmt.Println("profit: ", new(big.Int).Sub(pathExecutionValues.Revenue, bribe).String())

	//print protocols of each pool
	fmt.Println("flashswapIsV2: ", pathExecutionValues.IsV2Flashswap)
	fmt.Println("calldata: ", hex.EncodeToString(calldata))
	fmt.Println("blockNumber:", header.Number.String())
	fmt.Println("callTarget: ", pathExecutionValues.CallTarget)
	fmt.Println("Bribe:", bribe)
	fmt.Println("revenue:", revenue)
	//execute
	//fmt.Println("payloadValues: ", payloadValues)

}

func TestExecuteMixedPathWithV2V2V2(t *testing.T) {
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
		if paths[i].HasDuplicatePools {
			continue
		}

		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" && paths[i].Pools[len(paths[i].Pools)-1].GetType() == "uniswap_v2" && paths[i].Pools[len(paths[i].Pools)-2].GetType() == "uniswap_v2" && paths[i].Pools[0].GetType() == "uniswap_v2" && paths[i].HasDuplicatePools == false {
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
	//bribe = revenue * 0.9
	bribe := new(big.Int).Mul(revenue, big.NewInt(9))
	bribe = new(big.Int).Div(bribe, big.NewInt(10))

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, bribe, true, true)
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
	fmt.Println("bribe: ", bribe.String())
	fmt.Println("profit: ", new(big.Int).Sub(pathExecutionValues.Revenue, bribe).String())

	//print protocols of each pool
	fmt.Println("flashswapIsV2: ", pathExecutionValues.IsV2Flashswap)
	fmt.Println("calldata: ", hex.EncodeToString(calldata))
	fmt.Println("blockNumber:", header.Number.String())
	fmt.Println("callTarget: ", pathExecutionValues.CallTarget)
	//execute
	//fmt.Println("payloadValues: ", payloadValues)

}

func TestExecuteMixedPathWithV2V3V2(t *testing.T) {
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
		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xBDdA21dD8DA31D5bEe0c9bB886C044EBb9b8906a" && paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" && paths[i].Pools[0].GetType() == "uniswap_v2" && paths[i].Pools[1].GetType() == "uniswap_v3" && paths[i].Pools[2].GetType() == "uniswap_v2" && paths[i].HasDuplicatePools == false {
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

	//bribe = revenue * 0.9
	bribe := new(big.Int).Mul(revenue, big.NewInt(9))
	bribe = new(big.Int).Div(bribe, big.NewInt(10))

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, bribe, true, true)
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

	//PRINT RESULTS

	//compare with quoter
	quoterRevenue, err := profitablePath.CalculateRevenueWithQuoter(profitablePath.AmountIn, client)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}
	fmt.Println("quoterRevenue: ", quoterRevenue.String())

	path.PrintPath(profitablePath)

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

func TestExecuteMixedPathWithV3V3V2(t *testing.T) {
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
	for i := 0; i < len(paths); i++ {
		if len(paths[i].Pools) < 3 {
			continue
		}
		if paths[i].Pools[len(paths[i].Pools)-1].GetFactoryAddress().Hex() != "0xC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac" && paths[i].Pools[0].GetType() == "uniswap_v3" && paths[i].Pools[1].GetType() == "uniswap_v3" && paths[i].Pools[2].GetType() == "uniswap_v2" && paths[i].HasDuplicatePools == false {
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

	//PRINT RESULTS

	//compare with quoter
	quoterRevenue, err := profitablePath.CalculateRevenueWithQuoter(profitablePath.AmountIn, client)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}
	fmt.Println("quoterRevenue: ", quoterRevenue.String())

	path.PrintPath(profitablePath)

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

func TestExecuteMixedPathWithV3(t *testing.T) {
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

	//set amountInFor the first path
	//set revenue
	//set amountIn
	//find path with positive revenue
	// for i := 0; i < len(paths); i++ {
	// 	paths[i].SetBestAmountInAndRevenue(baseFee)
	// 	if paths[i].Revenue.Cmp(big.NewInt(0)) == 1 {
	// 		profitablePath = paths[i]
	// 		break
	// 	}
	// }
	profitablePath := paths[95]
	//err = UpdatePools(path.Pools, client)
	//path.SetBestAmountInAndRevenue(baseFee)
	//loop paths until last pool is v3 pool
	for i := 0; i < len(paths); i++ {
		if paths[i].Pools[len(paths[i].Pools)-1].GetType() == "uniswap_v3" && paths[i].Pools[len(paths[i].Pools)-2].GetType() == "uniswap_v3" && paths[i].Pools[0].GetType() == "uniswap_v3" && paths[i].HasDuplicatePools == false {
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

	//PRINT RESULTS

	//compare with quoter
	quoterRevenue, err := profitablePath.CalculateRevenueWithQuoter(profitablePath.AmountIn, client)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}
	fmt.Println("quoterRevenue: ", quoterRevenue.String())

	path.PrintPath(profitablePath)

	fmt.Println("path revenue: ", profitablePath.Revenue.String())

	//check if quoter revenue the same as path revenue
	if quoterRevenue.Cmp(profitablePath.Revenue) != 0 {
		fmt.Println("quoterRevenue: ", quoterRevenue.String())
		fmt.Println("path revenue: ", profitablePath.Revenue.String())
		t.Fatal("quoterRevenue != path revenue")
	}

	V3FlashswapPool, ok := profitablePath.Pools[len(profitablePath.Pools)-1].(*uniswap_v3.Pool)
	if !ok {
		t.Fatal("pool is not v3 pool")
	}
	quoterFlashswapAmountOut, err := uniswap_v3.GetAmountOutWithQuoter(*V3FlashswapPool, profitablePath.AmountIn, pathExecutionValues.V2FlashswapZeroForOne, client)
	if err != nil {
		fmt.Println("err: ", err)
		t.Fatal(err)
	}
	fmt.Println("quoterFlashswapAmountOut: ", quoterFlashswapAmountOut.String())

	//print amountIn
	fmt.Println("amountIn: ", profitablePath.AmountIn.String())

	// //print swapdata amountouts
	// for i := 0; i < len(pathExecutionValues.V2SwapDatas); i++ {
	// 	fmt.Println("amountOuts: ", pathExecutionValues.V2SwapDatas[i].AmountOut.String())
	// }

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

func TestExecuteMixedPath(t *testing.T) {
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

	//set amountInFor the first path
	//set revenue
	//set amountIn
	//find path with positive revenue
	profitablePath := paths[1]
	// for i := 0; i < len(paths); i++ {
	// 	paths[i].SetBestAmountInAndRevenue(baseFee)
	// 	if paths[i].Revenue.Cmp(big.NewInt(0)) == 1 {
	// 		profitablePath = paths[i]
	// 		break
	// 	}
	// }

	//check if v3 path
	if profitablePath.HasUniswapV3Pools {
		fmt.Println("v3 path")
	}

	//update profitable path
	err = UpdatePools(profitablePath.Pools, client)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//reset revenue and amountIn after update
	profitablePath.SetBestAmountInAndRevenue(baseFee)
	fmt.Println("path bestAmountIn: ", profitablePath.AmountIn.String())
	fmt.Println("path revenue: ", profitablePath.Revenue.String())

	path.PrintPath(profitablePath)

	//convert path to path execution values

	pathExecutionValues, err := convertPathToPathExecutionValues(profitablePath, profitablePath.Revenue, true, true)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

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

	fmt.Println("calldata: ", hex.EncodeToString(calldata))
	fmt.Println("blockNumber:", header.Number.String())
	fmt.Println("flashswapTarget:", profitablePath.Pools[len(profitablePath.Pools)-1].GetAddress().Hex())

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
