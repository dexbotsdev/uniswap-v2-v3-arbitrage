package logic

import (
	"mev-template-go/uniswap_v2"
	"mev-template-go/uniswap_v3"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAllPoolsAndWriteToFiles(client *ethclient.Client) error {
	//get all uniswap v2 pools
	_, err := uniswap_v2.GetAllUniswapV2ForkPoolsAndWriteToFile(client)
	if err != nil {
		return err
	}

	//get all uniswap v3 pools
	_, err = uniswap_v3.GetAllPoolsAndWriteToJson(client)
	if err != nil {
		return err
	}

	return nil
}
