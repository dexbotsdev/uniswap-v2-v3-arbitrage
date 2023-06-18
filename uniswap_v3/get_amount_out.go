package uniswap_v3

import (
	"fmt"
	"math/big"

	sdk_core_entities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/ethereum/go-ethereum/common"
)

//v3 notes
//each space between tick is treated as its own pool
//each tickSpace has a max value before the tick is crossed

//get amount out steps
//calculate price change
//if

//get amountAmount1Out and returns of tick is crossed
func GetAmountOut(
	pool Pool,
	amountIn *big.Int,
	zeroForOne bool,
) (result *big.Int, returnErr error) {
	//if panic, return 0
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
			result = big.NewInt(0)
			returnErr = nil
		}
	}()
	//Steps
	//track current amountIn and amountOut
	//while
	//calculate price change
	//if price change is within tick calculate amountout normally
	//if price change is outside of tick calculate amountout and check if tick is crossed

	tickListDataProvider, err := entities.NewTickListDataProvider(ToSdkTickArr(pool.Ticks), pool.TickSpacing)
	if err != nil {
		return nil, err
	}

	sdkFee, err := ToSdkFee(pool.Fee)
	if err != nil {
		return nil, err
	}
	sdkPool, err := entities.NewPool(
		ToSdkToken(pool.Token0),
		ToSdkToken(pool.Token1),
		sdkFee,
		pool.SqrtPriceX96,
		pool.Liquidity,
		pool.TickCurrent,
		tickListDataProvider,
	)
	//func (p *Pool) GetOutputAmount(inputAmount *entities.CurrencyAmount, sqrtPriceLimitX96 *big.Int) (*entities.CurrencyAmount, *Pool, error) {
	tokenIn := pool.Token0
	if !zeroForOne {
		tokenIn = pool.Token1
	}
	sdkInputCurrencyAmount := sdk_core_entities.CurrencyAmount{
		Fraction:     sdk_core_entities.NewFraction(amountIn, big.NewInt(1)),
		Currency:     ToSdkToken(tokenIn),
		DecimalScale: big.NewInt(0),
	}

	if pool.Address == common.HexToAddress("0x59f2044fc191f758fD3680478514244ACF253D48") {
		fmt.Println("pool address: ", pool.Address)
	}

	amountOutTemp, _, err := sdkPool.GetOutputAmount(&sdkInputCurrencyAmount, nil)
	if err != nil {
		if err == entities.ErrSqrtPriceLimitX96TooHigh {
			return big.NewInt(0), nil
		}
		fmt.Println("pool address: ", pool.Address)
		return nil, err
	}
	amountOut := new(big.Int).Set(amountOutTemp.Fraction.Numerator)

	return amountOut, nil
}

func GetAmountOutAndUpdatePool(
	pool Pool,
	amountIn *big.Int,
	zeroForOne bool,
) (*big.Int, Pool, error) {
	//Steps
	//track current amountIn and amountOut
	//while
	//calculate price change
	//if price change is within tick calculate amountout normally
	//if price change is outside of tick calculate amountout and check if tick is crossed

	tickListDataProvider, err := entities.NewTickListDataProvider(ToSdkTickArr(pool.Ticks), pool.TickSpacing)
	if err != nil {
		return nil, Pool{}, err
	}

	sdkFee, err := ToSdkFee(pool.Fee)
	if err != nil {
		return nil, Pool{}, err
	}
	sdkPool, err := entities.NewPool(
		ToSdkToken(pool.Token0),
		ToSdkToken(pool.Token1),
		sdkFee,
		pool.SqrtPriceX96,
		pool.Liquidity,
		pool.TickCurrent,
		tickListDataProvider,
	)
	//func (p *Pool) GetOutputAmount(inputAmount *entities.CurrencyAmount, sqrtPriceLimitX96 *big.Int) (*entities.CurrencyAmount, *Pool, error) {
	tokenIn := pool.Token0
	if !zeroForOne {
		tokenIn = pool.Token1
	}
	sdkInputCurrencyAmount := sdk_core_entities.CurrencyAmount{
		Fraction:     sdk_core_entities.NewFraction(amountIn, big.NewInt(1)),
		Currency:     ToSdkToken(tokenIn),
		DecimalScale: big.NewInt(0),
	}

	amountOutTemp, sdkPoolAfter, err := sdkPool.GetOutputAmount(&sdkInputCurrencyAmount, nil)
	amountOut := new(big.Int).Set(amountOutTemp.Fraction.Numerator)

	if err != nil {
		return nil, Pool{}, err
	}

	//convert sdk pool state to our pool state

	newPoolState := Pool{
		Address:        pool.Address,
		FactoryAddress: pool.FactoryAddress,
		Token0:         pool.Token0,
		Token1:         pool.Token1,
		SqrtPriceX96:   sdkPoolAfter.SqrtRatioX96,
		Liquidity:      sdkPoolAfter.Liquidity,
		TickCurrent:    sdkPoolAfter.TickCurrent,
		TickSpacing:    pool.TickSpacing,
		//TickMap:        CopyTickMapCopy(pool.TickMap),
		Ticks: pool.GetTicksCopy(),
		Fee:   pool.Fee,
	}

	//update all ticks from before tick to after tick
	// tickBefore := pool.TickCurrent
	// tickAfter := sdkPoolAfter.TickCurrent

	// tickStart := tickBefore
	// tickEnd := tickAfter
	// if tickBefore > tickAfter {
	// 	tickStart = tickAfter
	// 	tickEnd = tickBefore
	// }

	// //update the ticks in pool
	// for i := tickStart; i <= tickEnd; i += pool.TickSpacing {
	// 	//for each traversed tick, find the tick in the pool and update it
	// 	newTick := Tick{
	// 		Index:          i,
	// 		LiquidityGross: sdkPoolAfter.TickDataProvider.GetTick(i).LiquidityGross,
	// 		LiquidityNet:   sdkPoolAfter.TickDataProvider.GetTick(i).LiquidityNet,
	// 	}
	// 	pool.UpdateTick(newTick)
	// }

	return amountOut, newPoolState, nil
}
