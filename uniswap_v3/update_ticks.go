package uniswap_v3

import (
	"context"
	"encoding/hex"
	"math/big"
	"mev-template-go/pkg/constructor_multicall"
	"sort"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func UpdateTicks(pools []*Pool, client *ethclient.Client) error {
	for i := 0; i < len(pools); i++ {
		err := UpdateAllTicksForPool(pools[i], client)
		if err != nil {
			return err
		}
	}

	return nil
}

func UpdateAllTicksForPool(pool *Pool, client *ethclient.Client) error {
	//get tick word bounds
	bottomWord, topWord, err := getTickWordBounds(*pool, client)
	if err != nil {
		return err
	}

	channelSize := topWord - bottomWord + 1

	if channelSize < 0 {
		pool.Ticks = make([]Tick, 0)
		return nil
	}

	//fmt.Println("channelSize: ", channelSize)

	ticksChannel := make(chan []Tick, topWord-bottomWord+1) // Buffer to prevent blocking
	errChannel := make(chan error, topWord-bottomWord+1)    // For error propagation
	var wg sync.WaitGroup                                   // Waitgroup to wait for all goroutines to finish

	for i := bottomWord; i <= topWord; i++ {
		wg.Add(1)
		go func(poolAddress common.Address, tickBitmapIndex int, client *ethclient.Client) {
			defer wg.Done()
			ticks, err := getTicksInWordByTickLens(poolAddress, tickBitmapIndex, client)
			if err != nil {
				errChannel <- err
			} else {
				ticksChannel <- ticks
			}

		}(pool.Address, i, client)
	}

	// Close channels after all goroutines finish
	go func() {
		wg.Wait()
		close(ticksChannel)
		close(errChannel)
	}()

	// Check for errors from goroutines
	for err := range errChannel {
		if err != nil {
			return err
		}
	}

	// Collect all new ticks
	newTicks := make([]Tick, 0)
	for ticks := range ticksChannel {
		newTicks = append(newTicks, ticks...)
	}

	pool.Ticks = newTicks

	//sort ticks
	sort.Slice(pool.Ticks, func(i, j int) bool {
		return pool.Ticks[i].Index < pool.Ticks[j].Index
	})

	return nil
}

func getTickWordBounds(pool Pool, client *ethclient.Client) (int, int, error) {
	//get contract bytecode
	contractBytecode, err := constructor_multicall.GetBytecodeStringFromBin(`bin\uniswap_v3\contracts\GetUniswapV3TickWordBounds.bin`)
	if err != nil {
		return 0, 0, err
	}

	//add augurements
	Address, _ := abi.NewType("address", "", nil)
	argBytes, err := abi.Arguments{{Type: Address}}.Pack(pool.Address)

	fullBytecode := contractBytecode + hex.EncodeToString(argBytes)
	fullBytecodeString, err := hex.DecodeString(fullBytecode)

	//create call msg
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x0000000000000000000000000000000000000000"), // Set the caller address to the zero address
		To:   nil,
		Data: fullBytecodeString,
	}

	//call contract
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return 0, 0, err
	}

	//unpack result
	Int256, _ := abi.NewType("int256", "", nil)
	Uint256, _ := abi.NewType("uint256", "", nil)
	decodedResults, err := abi.Arguments{{Type: Uint256}, {Type: Int256}, {Type: Int256}}.Unpack(result)

	//return results
	start := decodedResults[1].(*big.Int)
	end := decodedResults[2].(*big.Int)

	//cast to int
	bottomWord := int(start.Int64())
	topWord := int(end.Int64())

	return bottomWord, topWord, nil
}

func getTicksInWordByTickLens(poolAddress common.Address, tickBitmapIndex int, client *ethclient.Client) ([]Tick, error) {
	//  function getPopulatedTicksInWord(
	//   address pool,
	//   int16 tickBitmapIndex
	// ) public returns (struct ITickLens.PopulatedTick[] populatedTicks)
	//"351fb478c00f4e654d3c70ba9f3b78cad3b15352bd550410495a69885ff8508f"
	const getPopulatedTicksInWordHash = "351fb478"

	tickLensAddress := common.HexToAddress("0xbfd8137f7d1516D3ea5cA83523914859ec47F573")

	//encode arguments
	Address, _ := abi.NewType("address", "", nil)
	Int256, _ := abi.NewType("int256", "", nil)
	argBytes, err := abi.Arguments{{Type: Address}, {Type: Int256}}.Pack(poolAddress, big.NewInt(int64(tickBitmapIndex)))

	fullBytecode := getPopulatedTicksInWordHash + hex.EncodeToString(argBytes)
	fullBytecodeString, err := hex.DecodeString(fullBytecode)

	//create call msg
	callMsg := ethereum.CallMsg{
		From: common.HexToAddress("0x000000"),
		To:   &tickLensAddress,
		Data: fullBytecodeString,
	}

	//call contract
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}

	//fmt.Println("result", hex.EncodeToString(result))
	// type Tick struct {
	// 	Index          int
	// 	LiquidityNet   *big.Int
	// 	LiquidityGross *big.Int
	// }

	type PopulatedTick struct {
		Tick           *big.Int
		LiquidityNet   *big.Int
		LiquidityGross *big.Int
	}

	tickLensAbi := "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"tickBitmapIndex\",\"type\":\"int16\"}],\"name\":\"getPopulatedTicksInWord\",\"outputs\":[{\"components\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"}],\"internalType\":\"structPopulatedTick[]\",\"name\":\"populatedTicks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

	exp, err := abi.JSON(strings.NewReader(tickLensAbi))
	if err != nil {
		return nil, err
	}

	var populatedTickArr []PopulatedTick

	err = exp.UnpackIntoInterface(&populatedTickArr, "getPopulatedTicksInWord", result)
	if err != nil {
		return nil, err
	}

	newTicks := make([]Tick, 0)
	for _, tick := range populatedTickArr {
		newTick := Tick{
			Index:          int(tick.Tick.Int64()),
			LiquidityNet:   tick.LiquidityNet,
			LiquidityGross: tick.LiquidityGross,
		}
		newTicks = append(newTicks, newTick)
	}

	return newTicks, nil
}

func updateTicksBatched(pools []*Pool, client *ethclient.Client) error {
	batchSize := 1

	for i := 0; i < len(pools); i += batchSize {

		//Set batch end. If current batch size surpasses pools length, set end to pools length.
		end := i + batchSize
		if end > len(pools) {
			end = len(pools)
		}

		batch := pools[i:end]
		//_, err := getTicksMulticall(batch, client)
		for i, _ := range batch {
			err := UpdateAllTicksForPool(pools[i], client)
			if err != nil {
				return err
			}
		}

	}
	//split pools into batches
	return nil
}

//updates the mutable varaibles of the pools
func getTicksMulticall(pools []*Pool, client *ethclient.Client) (*big.Int, error) {

	//get contract bytecode
	contractBytecode, err := constructor_multicall.GetBytecodeStringFromBin(`bin\uniswap_v3\contracts\GetUniswapV3TicksV2.bin`)
	if err != nil {
		return nil, err
	}

	//add augurements
	AddressArr, _ := abi.NewType("address[]", "", nil)
	poolAddresses := make([]common.Address, len(pools))
	for i, pool := range pools {
		poolAddresses[i] = pool.Address
	}
	argBytes, err := abi.Arguments{{Type: AddressArr}}.Pack(poolAddresses)
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

	//make the eth_call
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		return nil, err
	}

	Uint, _ := abi.NewType("uint256", "", nil)
	Uint2562DArr, _ := abi.NewType("uint256[][]", "", nil)
	Int2562DArr, _ := abi.NewType("int256[][]", "", nil)

	//decode results
	decodedResults, err := abi.Arguments{{Type: Uint}, {Type: Int2562DArr}, {Type: Int2562DArr}, {Type: Uint2562DArr}}.Unpack(result)

	//decodedResults, err := abi.Arguments{{Type: Uint}, {Type: UintArr}, {Type: UintArr}, {Type: UintArr}, {Type: PopulatedTicksTupleArr}}.Unpack(result)
	if err != nil {
		return nil, err
	}

	blockNumber := decodedResults[0].(*big.Int)
	for i := 0; i < len(pools); i++ {
		//for each pool make new tickmap and fill all the values

		//newTickMap := make(map[int]Tick)

		indexArr := decodedResults[1].([][]*big.Int)[i]
		liquidityNetArr := decodedResults[2].([][]*big.Int)[i]
		liquidityGrossArr := decodedResults[3].([][]*big.Int)[i]

		for j := 0; j < len(indexArr); j++ {
			index := int(indexArr[j].Int64())
			liquidityNet := new(big.Int).Set(liquidityNetArr[j])
			liquidityGross := new(big.Int).Set(liquidityGrossArr[j])
			newTick := Tick{Index: index, LiquidityNet: liquidityNet, LiquidityGross: liquidityGross}
			pools[i].Ticks = append(pools[i].Ticks, newTick)
			//newTickMap[index] = newTick
		}

		//pools[i].TickMap = newTickMap
	}

	return blockNumber, nil
}
