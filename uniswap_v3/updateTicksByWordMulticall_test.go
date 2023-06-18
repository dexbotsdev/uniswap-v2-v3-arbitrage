package uniswap_v3

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestUpdateTicksByWordBatched(t *testing.T) {
	err := os.Chdir("../")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	err = godotenv.Load(".env")
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	pool := Pool{
		Address: common.HexToAddress("0x1d42064Fc4Beb5F8aAF85F4617AE8b3b5B8Bd801"),
	}

	err = UpdateAllTicksForPool(&pool, client)
	fmt.Println("err: ", err)
	assert.NoError(t, err)

	fmt.Println("pool ticks length ", len(pool.Ticks))
}

func TestTickDecode(t *testing.T) {
	resultString := "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe16dc0000000000000000000000000000000000000000000000000025e55596b7bee50000000000000000000000000000000000000000000000000025e55596b7bee5"

	//convert hex string to []byte
	resultBytes := common.Hex2Bytes(resultString)

	//unpack result
	// TickTuple, err := abi.NewType("tuple", "PopulatedTick", []abi.ArgumentMarshaling{
	// 	{Name: "Tick", Type: "int256"},
	// 	{Name: "LiquidityNet", Type: "int256"},
	// 	{Name: "LiquidityGross", Type: "uint256"}})
	// fmt.Println("err: ", err)
	TickTuple, err := abi.NewType("(int256,int256,uint256)", "PopulatedTick", []abi.ArgumentMarshaling{
		{Type: "int256"},
		{Type: "int256"},
		{Type: "uint256"}})
	fmt.Println("err: ", err)

	decodedResults, err := abi.Arguments{{Type: TickTuple}}.Unpack(resultBytes)
	fmt.Println("err: ", err)

	type PopulatedTick struct {
		Tick           *big.Int
		LiquidityNet   *big.Int
		LiquidityGross *big.Int
	}

	newTick := decodedResults[0].(PopulatedTick)

	fmt.Println("newTicks: ", newTick)

}

func TestTickDecodeWithABI(t *testing.T) {
	//resultString := "fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe16dc0000000000000000000000000000000000000000000000000025e55596b7bee50000000000000000000000000000000000000000000000000025e55596b7bee5"

	resultString := "00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe16dc0000000000000000000000000000000000000000000000000025e55596b7bee50000000000000000000000000000000000000000000000000025e55596b7bee5fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffde43c00000000000000000000000000000000000000000000001f59aa5af0d690b3a400000000000000000000000000000000000000000000001f59aa5af0d690b3a4"
	resultBytes := common.Hex2Bytes(resultString)

	tickLensAbi := "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"tickBitmapIndex\",\"type\":\"int16\"}],\"name\":\"getPopulatedTicksInWord\",\"outputs\":[{\"components\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"}],\"internalType\":\"structPopulatedTick[]\",\"name\":\"populatedTicks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

	exp, err := abi.JSON(strings.NewReader(tickLensAbi))
	if err != nil {
		t.Fatal(err)
	}

	// decodedResult, err := exp.Unpack("getPopulatedTicksInWord", resultBytes)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	type PopulatedTick struct {
		Tick           int
		LiquidityNet   *big.Int
		LiquidityGross *big.Int
	}

	var tickArr []PopulatedTick

	err = exp.UnpackIntoInterface(&tickArr, "getPopulatedTicksInWord", resultBytes)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("tickArr: ", tickArr)

	// fmt.Println("result: ", decodedResult)

	// newTicks := make([]Tick, len(decodedResult[0].([]struct {
	// 	First  *big.Int
	// 	Second *big.Int
	// 	Third  *big.Int
	// })))

	// fmt.Println("newTicks: ", newTicks)

	fmt.Println("exp: ", exp)
}
