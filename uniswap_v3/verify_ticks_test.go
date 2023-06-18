package uniswap_v3

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestVeriftTicks(t *testing.T) {
	//get all filtered pools
	//get all ticks
	//call pool.ticks(index) for each tick
	//compare liquidityNet and liquidityGross

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

	pools := filteredPools //[0:100]

	UpdatePools(pools, client)

	for i := 0; i < len(filteredPools[0].Ticks); i++ {
		result, _ := VerifyTick(filteredPools[0], filteredPools[0].Ticks[i], client)
		assert.Equal(t, true, result)
	}

}

func verifyPool(pool *Pool, tick Tick, client *ethclient.Client) (bool, error) {
	return true, nil
}

func VerifyTick(pool *Pool, tick Tick, client *ethclient.Client) (bool, error) {
	//call pool.ticks(index) for each tick
	//compare liquidityNet and liquidityGross
	//get tick data from pool

	poolAddress := pool.Address
	tickIndex := tick.Index

	ticksSigHash := "f30dba93"

	tickIndexBigInt := big.NewInt(int64(tickIndex))
	fmt.Println("tickIndexBigInt: ", tickIndexBigInt.String())
	fmt.Println("tickIndexBigIntBytes: ", hex.EncodeToString(tickIndexBigInt.Bytes()))

	//tickIndexBytes := padBytesInt(big.NewInt(int64(tickIndex)).Bytes(), 32)

	Int256, _ := abi.NewType("int256", "", nil)

	argBytes, err := abi.Arguments{{Type: Int256}}.Pack(tickIndexBigInt)
	fmt.Println("args: ", argBytes)
	fmt.Println("err: ", err)

	calldataString := ticksSigHash + hex.EncodeToString(argBytes)

	fmt.Println("calldataString: ", calldataString)

	//create callmsg

	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x000000"),
		To:   &poolAddress,
		Data: common.FromHex(calldataString),
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return false, err
	}

	// 	struct Info {
	// 		uint128 liquidityGross;
	// 		int128 liquidityNet;
	// 		uint256 feeGrowthOutside0X128;
	// 		uint256 feeGrowthOutside1X128;
	// 		int56 tickCumulativeOutside;
	// 		uint160 secondsPerLiquidityOutsideX128;
	// 		uint32 secondsOutside;
	// 		bool initialized;
	// }
	//Int256, _ := abi.NewType("int256", "", nil)
	Uint256, _ := abi.NewType("uint256", "", nil)
	Bool, _ := abi.NewType("bool", "", nil)

	decodedResults, err := abi.Arguments{
		{Type: Uint256},
		{Type: Int256},
		{Type: Uint256},
		{Type: Uint256},
		{Type: Int256},
		{Type: Uint256},
		{Type: Uint256},
		{Type: Bool}}.Unpack(result)

	newLiquidityGross := decodedResults[0].(*big.Int)
	newLiquidityNet := decodedResults[1].(*big.Int)
	fmt.Println("result: ", result)

	fmt.Println("poolAddress: ", poolAddress)
	fmt.Println("index: ", tickIndex)
	fmt.Println("tick.LiquidityGross: ", tick.LiquidityGross.String())
	fmt.Println("tick.LiquidityNet: ", tick.LiquidityNet.String())
	fmt.Println("newLiquidityGross: ", newLiquidityGross.String())
	fmt.Println("newLiquidityNet: ", newLiquidityNet.String())

	if tick.LiquidityNet.Cmp(newLiquidityNet) != 0 {
		return false, nil
	}
	if tick.LiquidityGross.Cmp(newLiquidityGross) != 0 {
		return false, nil
	}
	return true, nil
}

//pads bytes to a certain length
func padBytesInt(input []byte, length int) []byte {
	if len(input) >= length {
		return input
	}

	padded := make([]byte, length)

	// Check if the number is negative.
	// If it is, pad with 0xFF bytes; otherwise, pad with 0x00 bytes.
	paddingByte := byte(0x00)
	if input[0]&0x80 == 0x80 {
		paddingByte = byte(0xFF)
	}

	for i := 0; i < length-len(input); i++ {
		padded[i] = paddingByte
	}

	copy(padded[length-len(input):], input)

	return padded
}
