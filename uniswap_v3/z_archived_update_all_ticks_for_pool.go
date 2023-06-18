package uniswap_v3

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"mev-template-go/pkg/constructor_multicall"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//updates the mutable varaibles of the pools
func getTicksByWordMulticall(pool *Pool, start int, end int, client *ethclient.Client) (*big.Int, error) {

	//get contract bytecode
	contractBytecode, err := constructor_multicall.GetBytecodeStringFromBin(`bin\uniswap_v3\contracts\GetUniswapV3TicksByWords.bin`)
	if err != nil {
		return nil, err
	}

	//add augurements
	Address, _ := abi.NewType("address", "", nil)
	Int, _ := abi.NewType("int256", "", nil)
	argBytes, err := abi.Arguments{{Type: Address}, {Type: Int}, {Type: Int}}.Pack(pool.Address, big.NewInt(int64(start)), big.NewInt(int64(end)))
	if err != nil {
		return nil, err
	}

	fullBytecode := contractBytecode + hex.EncodeToString(argBytes)
	fullBytecodeString, err := hex.DecodeString(fullBytecode)

	//create call msg
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"), // Set the caller address to the zero address
		To:   nil,
		Data: fullBytecodeString,
	}
	fmt.Println("callMsg: ", callMsg)

	//make the eth_call
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("result: ", result)
	fmt.Println("result hex: ", hex.EncodeToString(result))

	//decode result
	//AddressArr, _ = abi.NewType("address[]", "", nil)
	Uint, _ := abi.NewType("uint256", "", nil)
	Uint256Arr, _ := abi.NewType("uint256[]", "", nil)
	Int256Arr, _ := abi.NewType("int256[]", "", nil)
	//Int242DArr, _ := abi.NewType("Int24[][]", "", nil)
	//Int242DArr, _ := abi.NewType("int24[][]", "", nil)
	//Int1282DArr, _ := abi.NewType("int128[][]", "", nil)
	//Uint1282DArr, _ := abi.NewType("uint128[][]", "", nil)
	//Uint2562DArr, _ := abi.NewType("uint256[][]", "", nil)

	// PopulatedTicksTupleArr, err := abi.NewType("tuple[]", "struct PopulatedTick[]", []abi.ArgumentMarshaling{
	// 	{Name: "index", Type: "uint256"},
	// 	{Name: "liquidityGross", Type: "uint256"},
	// 	{Name: "liquidityNet", Type: "uint256"}})
	// if err != nil {
	// 	return nil, err
	// }

	//decode results
	decodedResults, err := abi.Arguments{{Type: Uint}, {Type: Int256Arr}, {Type: Int256Arr}, {Type: Uint256Arr}}.Unpack(result)

	//decodedResults, err := abi.Arguments{{Type: Uint}, {Type: UintArr}, {Type: UintArr}, {Type: UintArr}, {Type: PopulatedTicksTupleArr}}.Unpack(result)
	if err != nil {
		return nil, err
	}
	fmt.Println("decodedResults: ", decodedResults)

	//update pools
	// address[] memory token0Arr = new address[](len);
	// address[] memory token1Arr = new address[](len);
	// uint128[] memory liquidityArr = new uint128[](len);
	// uint160[] memory sqrtPriceX96Arr = new uint160[](len);
	// PopulatedTick[][] memory populatedTicksArr = new PopulatedTick[][](len);
	blockNumber := decodedResults[0].(*big.Int)
	//for each pool make new tickmap and fill all the values

	indexArr := decodedResults[1].([]*big.Int)
	liquidityNetArr := decodedResults[2].([]*big.Int)
	liquidityGrossArr := decodedResults[3].([]*big.Int)

	newTicks := make([]Tick, len(indexArr))
	for j := 0; j < len(indexArr); j++ {
		newTicks[j].Index = int(indexArr[j].Int64())
		newTicks[j].LiquidityNet = liquidityNetArr[j]
		newTicks[j].LiquidityGross = liquidityGrossArr[j]
	}
	pool.Ticks = newTicks

	//pools[i].TickMap = newTickMap
	fmt.Println("pool ticks length", len(pool.Ticks))

	return blockNumber, nil
}
