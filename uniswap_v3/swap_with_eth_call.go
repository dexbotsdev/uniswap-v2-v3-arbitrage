package uniswap_v3

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func swapWithEthCall(pool Pool, amountIn *big.Int, zeroForOne bool, client *ethclient.Client) (*big.Int, *big.Int, error) {
	//calls exactinput on swap with ethcall to verify the amount out

	poolAddress := pool.Address

	//setup arguements of swap
	recipent := common.HexToAddress("0x0000000")
	sqrtLimitX96 := big.NewInt(0)
	amountSpecified := amountIn
	data := []byte{}

	//set sqrtLimitX96

	MIN_SQRT_RATIO, success := new(big.Int).SetString("4295128739", 10)
	if !success {
		return nil, nil, errors.New("could not set min sqrt ratio")
	}

	MAX_SQRT_RATIO, success := new(big.Int).SetString("1461446703485210103287273052203988822378723970342", 10)
	if !success {
		return nil, nil, errors.New("could not set max sqrt ratio")
	}

	if zeroForOne {
		sqrtLimitX96.Add(MIN_SQRT_RATIO, big.NewInt(1))
	} else {
		sqrtLimitX96.Sub(MAX_SQRT_RATIO, big.NewInt(1))
	}

	Uniswap_V3_SWAP_SIG := "128acb08"

	Int, _ := abi.NewType("int256", "", nil)
	Uint, _ := abi.NewType("uint256", "", nil)
	Address, _ := abi.NewType("address", "", nil)
	Bool, _ := abi.NewType("bool", "", nil)
	Bytes, _ := abi.NewType("bytes", "", nil)

	argBytes, err := abi.Arguments{{Type: Address}, {Type: Bool}, {Type: Int}, {Type: Uint}, {Type: Bytes}}.Pack(recipent, zeroForOne, amountSpecified, sqrtLimitX96, data)
	if err != nil {
		return nil, nil, err
	}

	//setup call
	callData := append(common.Hex2Bytes(Uniswap_V3_SWAP_SIG), argBytes...)

	fmt.Println("callhassig hexto bytes: ", hex.EncodeToString(common.Hex2Bytes(Uniswap_V3_SWAP_SIG)))

	fmt.Println("callData: ", hex.EncodeToString(callData))
	fmt.Println("pool address: ", poolAddress)

	//get blocknumber
	blockNumber, err := client.BlockNumber(context.Background())
	fmt.Println("blocknumber: ", blockNumber)

	//setup callmsg
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x00000"),
		To:   &poolAddress,
		Data: callData,
	}

	//make eth call
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, nil, err
	}

	//unpack result
	decoded, err := abi.Arguments{{Type: Int}, {Type: Int}}.UnpackValues(result)
	if err != nil {
		return nil, nil, err
	}

	amount0 := decoded[0].(*big.Int)
	amount1 := decoded[1].(*big.Int)

	return amount0, amount1, nil

}

func GetAmountOutWithEthCall(pool Pool, amountIn *big.Int, zeroForOne bool, client *ethclient.Client) (*big.Int, error) {
	//check for negative
	amount0, amount1, err := swapWithEthCall(pool, amountIn, zeroForOne, client)
	if err != nil {
		return nil, err
	}

	if zeroForOne {
		return new(big.Int).Mul(amount1, big.NewInt(-1)), nil
	} else {
		return new(big.Int).Mul(amount0, big.NewInt(-1)), nil
	}
}
