package uniswap_v3

import (
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/daoleno/uniswapv3-sdk/constants"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/daoleno/uniswapv3-sdk/utils"
	sdk_utils "github.com/daoleno/uniswapv3-sdk/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	// "mev-template-go/uniswapv3-sdk-master/constants"
	// "mev-template-go/uniswapv3-sdk-master/entities"
	// "mev-template-go/uniswapv3-sdk-master/utils"
	// sdk_utils "mev-template-go/uniswapv3-sdk-master/utils"
)

func TestGetAmountOutWithQuoter(t *testing.T) {
	//select pools
	//update pools
	//run get amount out
	//compare with quoter

	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	//get all filtered pools
	filteredPools, err := ReadFilteredPoolsFromFile()
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	pools := filteredPools[0:100]

	// pools = filterOutPoolsWithoutTicks(pools)
	// fmt.Println("err: ", err)
	// assert.NoError(t, err)

	amountIn, _ := new(big.Int).SetString("100000000000000000000", 10)

	for i := 0; i < len(pools); i++ {
		pools[i].Update(client)
		if pools[i].Ticks == nil || len(pools[i].Ticks) == 0 {
			continue
		}

		amountOut, err := pools[i].GetAmountOut(amountIn, true)
		if err != nil {
			fmt.Println("err: ", err)
		}

		amountOutQuoter, err := GetAmountOutWithQuoter(*pools[i], amountIn, true, client)
		if err != nil {
			fmt.Println("err: ", err)
		}

		fmt.Println("amountOut: ", amountOut)
		fmt.Println("amountOutQuoter: ", amountOutQuoter)
	}
}

func TestGetAmountOut(t *testing.T) {
	ticks := []Tick{
		{
			Index:          entities.NearestUsableTick(sdk_utils.MinTick, constants.TickSpacings[constants.FeeLow]),
			LiquidityNet:   big.NewInt(1000000000000000000),
			LiquidityGross: big.NewInt(1000000000000000000),
		},
		{
			Index:          entities.NearestUsableTick(utils.MaxTick, constants.TickSpacings[constants.FeeLow]),
			LiquidityNet:   new(big.Int).Mul(big.NewInt(1000000000000000000), constants.NegativeOne),
			LiquidityGross: big.NewInt(1000000000000000000),
		},
	}

	usdcAddress := common.HexToAddress("0x2791bca1f2de4661ed88a30c99a7a9449aa84174")
	daiAddress := common.HexToAddress("0x8f3cf7ad23cd3cadbd9735aff958023239c6a063")

	pool := Pool{
		Address:        common.HexToAddress("0x1f98431c8ad98523631ae4a59f267346ea31f984"),
		FactoryAddress: common.HexToAddress("0x1f98431c8ad98523631ae4a59f267346ea31f984"),

		Token0: usdcAddress,
		Token1: daiAddress,

		SqrtPriceX96: EncodeSqrtPriceX96(constants.One, constants.One),
		Liquidity:    big.NewInt(1000000000000000000),

		TickCurrent: 0,
		TickSpacing: constants.TickSpacings[constants.FeeLow],
		Ticks:       ticks,

		Fee: 500,
	}

	amountIn := big.NewInt(100)
	amountOut, poolStateAfter, err := GetAmountOutAndUpdatePool(pool, amountIn, true)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("amountOut", amountOut.String())
	fmt.Println("poolStateAfter", poolStateAfter)

	assert.Equal(t, big.NewInt(98), amountOut)

}
func EncodeSqrtPriceX96(amount1 *big.Int, amount0 *big.Int) *big.Int {
	numerator := new(big.Int).Lsh(amount1, 192)
	denominator := amount0
	ratioX192 := new(big.Int).Div(numerator, denominator)
	return new(big.Int).Sqrt(ratioX192)
}
