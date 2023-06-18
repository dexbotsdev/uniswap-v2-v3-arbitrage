package uniswap_v2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"mev-template-go/types"
	UniV2Router "mev-template-go/uniswap_v2/contracts/uniswap_v2_router"
	"net/http"
	"os"
	"strings"
	"time"

	UniV2Factory "mev-template-go/uniswap_v2/contracts/uniswap_v2_factory"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type UniswapV2 struct {
	Router   UniV2Router.UniV2Router
	RouerAbi abi.ABI
	Factory  UniV2Factory.UniV2Factory
}

func New(config types.Config, routerAddress common.Address, factoryAddress common.Address, pathToAbi string) (*UniswapV2, error) {
	router, err := UniV2Router.NewUniV2Router(routerAddress, &config.Client)
	if err != nil {
		return nil, fmt.Errorf("Error creating new router: %v", err)
	}
	factory, err := UniV2Factory.NewUniV2Factory(factoryAddress, &config.Client)
	if err != nil {
		return nil, fmt.Errorf("Error creating new factory: %v", err)
	}
	routerAbi, err := abi.JSON(strings.NewReader(GetLocalABI(pathToAbi)))
	if err != nil {
		return nil, fmt.Errorf("Error creating new router abi: %v", err)
	}
	return &UniswapV2{Router: *router, RouerAbi: routerAbi, Factory: *factory}, nil
}

func (uniV2 UniswapV2) DecodeTransactionInputData(data []byte) {
	methodSigData := data[:4]
	method, err := uniV2.RouerAbi.MethodById(methodSigData)
	if err != nil {
		fmt.Println("Error getting method by ID: ", err)
		return
	}

	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		fmt.Println("Error unpacking inputs: ", err)
		return
	}

	fmt.Printf("Method Name: %s\n", method.Name)
	fmt.Printf("Method inputs: %v\n", inputsMap)
	fmt.Println("")
}

func GetLocalABI(path string) string {
	abiFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		fmt.Println(err)
	}
	return string(result)
}

type UniV2TempStruct struct {
	Data struct {
		Blocks []struct {
			Number string `json:"number"`
		} `json:"blocks"`
		Pairs []struct {
			ID       string `json:"id"`
			Reserve0 string `json:"reserve0"`
			Reserve1 string `json:"reserve1"`
			Token0   struct {
				ID       string `json:"id"`
				Decimals string `json:"decimals"`
			} `json:"token0"`
			Token1 struct {
				ID       string `json:"id"`
				Decimals string `json:"decimals"`
			} `json:"token1"`
		} `json:"pairs"`
	} `json:"data"`
}

//get uniswap v2 poolsS
func GetUniV2Pools() ([]types.UniV2Pool, error) {
	poolsQuery := `
		{
			pairs(first: 1000 orderBy: trackedReserveETH, orderDirection: desc) 
			{ id, reserve0, reserve1, token0 {id, decimals}, token1{id, decimals} }
		}
	`
	var UniV2Pairs []types.UniV2Pool

	data, err := queryDataV2(poolsQuery)
	if err != nil {
		return nil, err
	}

	var final UniV2TempStruct
	if err := json.Unmarshal(data, &final); err != nil {
		return nil, err
	}

	for _, pair := range final.Data.Pairs {
		newReserve0 := new(big.Int).Set(FloatStringToBigInt(pair.Reserve0, pair.Token0.Decimals))
		newReserve1 := new(big.Int).Set(FloatStringToBigInt(pair.Reserve1, pair.Token1.Decimals))

		pool := types.UniV2Pool{
			Address:        common.HexToAddress(pair.ID),
			FactoryAddress: common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"),
			Token0:         common.HexToAddress(pair.Token0.ID),
			Token1:         common.HexToAddress(pair.Token1.ID),
			Reserve0:       newReserve0,
			Reserve1:       newReserve1,
			Fees:           997,
		}
		UniV2Pairs = append(UniV2Pairs, pool)
	}
	//convert string to bigint

	return UniV2Pairs, nil
}

func queryDataV2(query string) ([]byte, error) {
	data := map[string]string{
		"query": query,
	}
	queryJSON, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2", bytes.NewBuffer(queryJSON))
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

// TODO - only works for 18 decimal tokens
func FloatStringToBigInt(val string, decimals string) *big.Int {
	bigval := new(big.Float)
	bigval.SetString(val)

	bigdec := new(big.Int)
	bigdec.SetString(decimals, 10)

	//set coin to the value of 10^decimals
	coin := new(big.Float)
	coin.SetInt(big.NewInt(10).Exp(big.NewInt(10), bigdec, nil))
	bigval.Mul(bigval, coin)

	//convert bigval to big.Int
	result := new(big.Int)
	bigval.Int(result)

	return result
}
