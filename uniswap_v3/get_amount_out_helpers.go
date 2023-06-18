package uniswap_v3

import (
	"errors"

	sdk_core_entities "github.com/daoleno/uniswap-sdk-core/entities"
	"github.com/daoleno/uniswapv3-sdk/entities"
	"github.com/ethereum/go-ethereum/common"

	sdk_constants "github.com/daoleno/uniswapv3-sdk/constants"
)

//v3 notes
//each space between tick is treated as its own pool
//each tickSpace has a max value before the tick is crossed

//get amount out steps
//calculate price change
//if

//get amountAmount1Out and returns of tick is crossed

func ToSdkToken(address common.Address) *sdk_core_entities.Token {
	//func NewToken(chainID uint, address common.Address, decimals uint, symbol string, name string) *Token {
	newToken := sdk_core_entities.NewToken(1, address, 6, "", "")
	return newToken
}

func ToSdkFee(fee uint32) (sdk_constants.FeeAmount, error) {
	// FeeLowest FeeAmount = 100
	// FeeLow    FeeAmount = 500
	// FeeMedium FeeAmount = 3000
	// FeeHigh   FeeAmount = 10000

	// FeeMax FeeAmount = 1000000

	if fee == 100 {
		return sdk_constants.FeeLowest, nil
	} else if fee == 500 {
		return sdk_constants.FeeLow, nil
	} else if fee == 3000 {
		return sdk_constants.FeeMedium, nil
	} else if fee == 10000 {
		return sdk_constants.FeeHigh, nil
	} else if fee == 1000000 {
		return sdk_constants.FeeMax, nil
	} else { //error
		return 0, errors.New("invalid fee")
	}
}

func ToSdkTick(tick Tick) entities.Tick {
	return entities.Tick{
		Index:          tick.Index,
		LiquidityGross: tick.LiquidityGross,
		LiquidityNet:   tick.LiquidityNet,
	}
}

func ToSdkTickArr(ticks []Tick) []entities.Tick {
	var sdkTicks []entities.Tick
	for _, tick := range ticks {
		sdkTicks = append(sdkTicks, ToSdkTick(tick))
	}
	return sdkTicks
}
