package uniswap_v3

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAmountOutWithQuoter(pool Pool, amountIn *big.Int, zeroForOne bool, client *ethclient.Client) (*big.Int, error) {
	//calls exactinput on swap with ethcall to verify the amount out

	if pool.Ticks == nil || len(pool.Ticks) == 0 || pool.Liquidity.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0), nil
	}

	//setup arguements of swap
	sqrtLimitX96 := big.NewInt(0)

	//set sqrtLimitX96
	MIN_SQRT_RATIO, success := new(big.Int).SetString("4295128739", 10)
	if !success {
		return nil, errors.New("could not set min sqrt ratio")
	}
	MAX_SQRT_RATIO, success := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)
	if !success {
		return nil, errors.New("could not set max sqrt ratio")
	}
	if zeroForOne {
		sqrtLimitX96.Add(MIN_SQRT_RATIO, big.NewInt(1))
	} else {
		sqrtLimitX96.Sub(MAX_SQRT_RATIO, big.NewInt(1))
	}

	//set tokenIn and tokenOut
	tokenIn := pool.Token0
	tokenOut := pool.Token1
	if !zeroForOne {
		tokenIn = pool.Token1
		tokenOut = pool.Token0
	}

	fee := new(big.Int).SetUint64(uint64(pool.Fee))

	quoterSigHash := "f7729d43"

	// address tokenIn,
	// address tokenOut,
	// uint24 fee,
	// uint256 amountIn,
	// uint160 sqrtPriceLimitX96

	Uint, _ := abi.NewType("uint256", "", nil)
	Address, _ := abi.NewType("address", "", nil)

	argBytes, err := abi.Arguments{{Type: Address}, {Type: Address}, {Type: Uint}, {Type: Uint}, {Type: Uint}}.Pack(tokenIn, tokenOut, fee, amountIn, sqrtLimitX96)
	if err != nil {
		return nil, err
	}
	//setup call
	callData := append(common.Hex2Bytes(quoterSigHash), argBytes...)

	//fmt.Println("callData: ", hex.EncodeToString(callData))

	//get blocknumber
	// blockNumber, err := client.BlockNumber(context.Background())
	// fmt.Println("blocknumber: ", blockNumber)

	quoterAddress := common.HexToAddress(ContractV3Quoter)

	//setup callmsg
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x00000"),
		To:   &quoterAddress,
		Data: callData,
	}

	//make eth call
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		fmt.Println(pool.String())
		return nil, err
	}

	//unpack result
	decoded, err := abi.Arguments{{Type: Uint}}.UnpackValues(result)
	if err != nil {
		return nil, err
	}

	amountOut := decoded[0].(*big.Int)

	return amountOut, nil
}
