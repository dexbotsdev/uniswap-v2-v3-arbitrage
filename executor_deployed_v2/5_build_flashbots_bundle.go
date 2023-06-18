package executor

//changes to make: add bribe amount instead of percentage
//add owner transfer bool

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"mev-template-go/logic"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/umbracle/ethgo/abi"
)

func BuildFlashbotsBundle(path logic.Path, ownerTransfer bool, coinbaseTransfer bool, client *ethclient.Client) ([]byte, error) {

	calldata, err := PathToCalldata(path, path.Revenue, ownerTransfer, coinbaseTransfer)
	if err != nil {
		return nil, err
	}

	//***ESTIMATE GAS AND UPDATE BRIBE***
	//estimate gas for the flashswap transaction
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	fmt.Println("gasPrice: ", gasPrice)

	toAddress := path.Pools[len(path.Pools)-1].GetAddress()

	gasEstimate, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     executorAddress,
		To:       &toAddress,
		GasPrice: gasPrice,
		Data:     calldata,
	})
	gasLimit := gasEstimate //TODO: check if this is correct

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	//maxBaseFee = baseFee * 1.125
	maxBaseFee := new(big.Int).Div(new(big.Int).Mul(header.BaseFee, big.NewInt(9)), big.NewInt(8)) // 9/8 = 1.125

	//minGasCost = maxBaseFee * gasLimit
	minGasCost := new(big.Int).Mul(maxBaseFee, big.NewInt(int64(gasLimit)))

	//if minGasCost > revenue, then return
	if minGasCost.Cmp(path.Revenue) > 0 {
		return nil, fmt.Errorf("minGasCost > revenue; unprofitable")
	}

	//***CALCULATE BRIBE***
	profit := new(big.Int).Sub(path.Revenue, minGasCost)                                   //profit = revenue - gasEstimate
	bribePercentage := big.NewInt(999)                                                     //bribePercentage = 99
	bribe := new(big.Int).Div(new(big.Int).Mul(profit, bribePercentage), big.NewInt(1000)) //bribe = profit * bribePercentage / 100

	//rebuild payload with bribe
	calldata, err = PathToCalldata(path, bribe, true, true)
	if err != nil {
		return nil, err
	}

	//

	return calldata, nil
}
