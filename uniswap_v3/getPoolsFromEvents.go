package uniswap_v3

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getPoolsFromSubGraph() {

}

func GetPoolsFromEvents(client *ethclient.Client) ([]*Pool, error) {

	ctx := context.Background()

	poolCreatedSigHash := "783cca1c0412dd0d695e784568c96da2e9c22ff989357a2e8b1d9b2b4e6b7118"
	uniswapV3FacotryAddress := common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")

	//create filter query with factory address and event name
	filterQuery := ethereum.FilterQuery{
		Addresses: []common.Address{uniswapV3FacotryAddress},
		Topics: [][]common.Hash{
			{
				common.HexToHash(poolCreatedSigHash), //pair created topic
			},
		},
	}

	//get logs from filter query
	logs, err := client.FilterLogs(ctx, filterQuery)
	if err != nil {
		return nil, err
	}
	fmt.Println("logs: ", logs)

	//decode logs and get pool addresses
	//  event PoolCreated(address token0,address token1,uint24 fee,int24 tickSpacing,address pool)

	pools := make([]*Pool, len(logs))
	for _, log := range logs {
		Address, _ := abi.NewType("address", "", nil)
		Uint24, _ := abi.NewType("uint24", "", nil)
		Int24, _ := abi.NewType("int24", "", nil)

		decoded, err := abi.Arguments{{Type: Address}, {Type: Address}, {Type: Uint24}, {Type: Int24}, {Type: Address}}.Unpack(log.Data)
		if err != nil {
			return nil, err
		}

		newPool := Pool{
			Address: decoded[4].(common.Address),
			Token0:  decoded[0].(common.Address),
			Token1:  decoded[1].(common.Address),
			Fee:     decoded[2].(uint32),
		}
		pools = append(pools, &newPool)
		fmt.Println("newPool: ", newPool)
	}

	//decode logs and get pool addresses

	return pools, nil
}
