package executor

//payload builder, shared memory transfer of transaction to geth, etc.. should be in here

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"mev-template-go/logic"
	"mev-template-go/types"
	"os"

	geth_types "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
	//"github.com/umbracle/ethgo/abi"
)

//modes
//0: v2 flashswap/ coinbase transfer
//1: v2 flashswap/ no coinbase transfer
//2: v3 flashswap/ coinbase transfer
//3: v3 flashswap/ no coinbase transfer

//flow only affects callback verification
//no verification
//v2 verification; no fee
//v3 verification; with fee

//PAYLOAD FORMAT
// startPos = 164 for v2, 196 for v3
// factoryAddress 20 bytes
// otherTokenAddress 20 bytes
// fee 2 bits; only used if flashswap is v3
// coinBaseTransferBool //1 bit
// swapDataCountMinus1 3 bits
// IntByteSize 5 bits
// revenue IntByteSize bytes
// quarterBribePercentageOfRevenue; 8 bits; only if coinBaseTransferBool is true
// isV2s; 1 byte; 1 bit for each swap
// zeroForOnes; 1 byte; 1 bit for each swap
// swapdatas

// SWAPDATA FORMAT
// target 20 bytes
// amountOut IntByteSize bytes; only if isV2 is true

var executorAddress = common.HexToAddress("0xAF56d20B378CED055E8e384c67214D4B9fD2C9cF") //TODO
var botWallet = common.HexToAddress("0x71296ebC93BB8645Fc0826EAED445e55b0813B41")       //TODO

func ExecuteMixedPath(path logic.Path, config types.Config) error {
	fmt.Println("ExecuteMixedPath")
	//convert path to payload values

	coinbaseTransfer := false
	ownerTransfer := false

	//calulate if coinbase transfer is needed.

	calldata, err := PathToCalldata(path, path.Revenue, ownerTransfer, coinbaseTransfer)
	if err != nil {
		fmt.Println("PathToCalldata error: ", err)
		return err
	}

	pathExecutionValues, err := convertPathToPathExecutionValues(path, path.Revenue, ownerTransfer, coinbaseTransfer)
	if err != nil {
		fmt.Println("convertPathToPathExecutionValues error: ", err)
		return err
	}

	//***ESTIMATE GAS AND UPDATE BRIBE***
	//estimate gas for the flashswap transaction
	gasPrice, err := config.Client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("SuggestGasPrice error: ", err)
		return err
	}
	fmt.Println("gasPrice: ", gasPrice)

	header, err := config.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println("HeaderByNumber error: ", err)
		return err
	}

	//toAddress := path.Pools[len(path.Pools)-1].GetAddress()
	//maxBaseFee = baseFee * 1.125
	maxBaseFee := new(big.Int).Div(new(big.Int).Mul(header.BaseFee, big.NewInt(9)), big.NewInt(8)) // 9/8 = 1.125
	fmt.Println("confit.WalletAddress: ", config.WalletAddress.Hex())
	fmt.Println("pathExecutionValues.CallTarget: ", pathExecutionValues.CallTarget.Hex())
	fmt.Println("callData: ", hexutil.Encode(calldata))
	fmt.Println("block number: ", header.Number.Uint64())

	gasEstimate, err := config.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:      config.WalletAddress,
		To:        &pathExecutionValues.CallTarget,
		GasPrice:  gasPrice,
		GasFeeCap: new(big.Int).Mul(maxBaseFee, big.NewInt(2)), // 2 * maxBaseFee
		Data:      calldata,
	})
	if err != nil {
		fmt.Println("call target: ", pathExecutionValues.CallTarget)
		return err
	}
	gasLimit := gasEstimate //TODO: check if this is correct

	//minGasCost = maxBaseFee * gasLimit
	minGasCost := new(big.Int).Mul(maxBaseFee, big.NewInt(int64(gasLimit)))

	fmt.Println("gasEstimate: ", gasEstimate)
	fmt.Println("gasLimit: ", gasLimit)
	fmt.Println("minGasCost: ", minGasCost)
	fmt.Println("revenue: ", path.Revenue)

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("PROFITABLE PATH BEFORE MIN GAS")

	log.Printf("profitable path: %v\n", path)
	log.Printf("block number: %v\n", header.Number.Uint64())
	log.Printf("path ID: %v\n", path.Id)
	log.Printf("gasEstimate: %v\n", gasEstimate)
	log.Printf("baseFee: %v\n", header.BaseFee)
	log.Printf("revenue: %v\n", path.Revenue)
	log.Printf("minGasCost: %v\n", minGasCost)
	log.Printf("profit: %v\n", new(big.Int).Sub(path.Revenue, minGasCost))
	log.Printf("calldata: %v\n", hexutil.Encode(calldata))
	log.Printf("call target: %v\n", pathExecutionValues.CallTarget.Hex())

	//if minGasCost > revenue, then return
	if minGasCost.Cmp(path.Revenue) > 0 {
		return fmt.Errorf("minGasCost > revenue; unprofitable")
	}

	//log profitable path
	//log block number
	//log path ID
	//log revenue
	//log profit
	//log calldata
	//log call target

	log.SetOutput(file)
	log.Println("PROFITABLE PATH AFTER MIN GAS")
	log.Printf("profitable path: %v\n", path)
	log.Printf("block number: %v\n", header.Number.Uint64())
	log.Printf("path ID: %v\n", path.Id)
	log.Printf("revenue: %v\n", path.Revenue)
	log.Printf("minGasCost: %v\n", minGasCost)
	log.Printf("profit: %v\n", new(big.Int).Sub(path.Revenue, minGasCost))
	log.Printf("calldata: %v\n", hexutil.Encode(calldata))
	log.Printf("call target: %v\n", pathExecutionValues.CallTarget.Hex())

	//***CALCULATE BRIBE***
	profit := new(big.Int).Sub(path.Revenue, minGasCost)                                   //profit = revenue - gasEstimate
	bribePercentage := big.NewInt(999)                                                     //bribePercentage = 99
	bribe := new(big.Int).Div(new(big.Int).Mul(profit, bribePercentage), big.NewInt(1000)) //bribe = profit * bribePercentage / 100

	//print revenue, bribe, block number
	fmt.Println("revenue: ", path.Revenue)
	fmt.Println("bribe: ", bribe)
	fmt.Println("block number: ", header.Number.Uint64())

	//rebuild payload with bribe
	calldata, err = PathToCalldata(path, bribe, true, true)
	if err != nil {
		return err
	}

	//******BUILD TRANSACTION******

	fromAddress := crypto.PubkeyToAddress(config.PrivateKey.PublicKey)

	nonce, err := config.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	transferAmount := big.NewInt(0)

	latestBlockNumber := header.Number.Uint64()

	maxFeePerGas := new(big.Int).Set(maxBaseFee) // Set max fee per gas equal to the max base fee
	maxPriorityFeePerGas := maxFeePerGas
	//set priority fee
	if !coinbaseTransfer { //bribe is sent via priority fee
		//base gas cost = gasLimit * maxFeePerGas
		// total gas cost = base gas cost + bribe
		// gas cost per gas = gasLimit / total gas cost
		//maxPriorityFeePerGas = bribe/ gasLimit
		baseGasCost := new(big.Int).Mul(maxFeePerGas, big.NewInt(int64(gasLimit)))
		totalGasCost := new(big.Int).Add(baseGasCost, bribe)
		maxFeePerGas = new(big.Int).Div(big.NewInt(int64(gasLimit)), totalGasCost)
		maxPriorityFeePerGas = maxFeePerGas
	}

	chainID := big.NewInt(1)

	//tx := gethTypes.NewTransaction(nonce, toAddress, transferAmount, gasLimit, gasPrice, nil)
	tx := geth_types.NewTx(&geth_types.DynamicFeeTx{
		ChainID:    chainID,
		Nonce:      nonce,
		GasTipCap:  maxPriorityFeePerGas,
		GasFeeCap:  maxFeePerGas,
		Gas:        gasLimit,
		To:         &pathExecutionValues.CallTarget,
		Value:      transferAmount,
		Data:       calldata,
		AccessList: nil,
	})
	fmt.Println("tx", tx)

	//******SIGN TRANSACTION******

	signedTx, err := geth_types.SignTx(tx, geth_types.LatestSignerForChainID(chainID), config.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signedTx", signedTx)

	//signedTxs := gethTypes.Transactions{signedTx}

	data, err := signedTx.MarshalBinary()
	if err != nil {
		log.Fatal("Failed to marshal signed transaction:", err)
	}
	hexEncodedTx := hexutil.Encode(data)

	//*******BUILD FLASHBOTS TRANSACTION******
	// Set up the Flashbots client
	flashbotsSigningKey, err := crypto.HexToECDSA("6e1fed8914d24893a3e5cb025683b50c8c17bfaa6acd92835bc8d096e2897002")
	if err != nil {
		log.Fatal(err)
	}
	//flashbotsSigningKey, err := crypto.HexToECDSA(os.Getenv("FLASHBOTS_SIGNING_KEY"))

	rpc := flashbotsrpc.New("https://relay.flashbots.net")

	// Simulate transaction
	callBundleArgs := flashbotsrpc.FlashbotsCallBundleParam{
		Txs:              []string{hexEncodedTx},
		BlockNumber:      fmt.Sprintf("0x%x", latestBlockNumber),
		StateBlockNumber: "latest",
	}
	fmt.Println("callBundleArgs:", callBundleArgs)
	fmt.Println("hexEncodedTx:", hexEncodedTx)

	txJSON, err := signedTx.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("signedTx JSON:", string(txJSON))

	fmt.Println("gasLimit:", gasLimit)
	callResults, err := rpc.FlashbotsCallBundle(flashbotsSigningKey, callBundleArgs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FlashbotsCallBundle: %+v\n", callResults)

	//get user stats
	statsResults, err := rpc.FlashbotsGetUserStats(flashbotsSigningKey, 13281018)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FlashbotsGetUserStats: %+v\n", statsResults)
	return nil

	//send transaction
	sendBundleArgs := flashbotsrpc.FlashbotsSendBundleRequest{
		Txs:         []string{"0x" + hexEncodedTx},
		BlockNumber: fmt.Sprintf("0x%x", 13281018),
	}

	sendResult, err := rpc.FlashbotsSendBundle(flashbotsSigningKey, sendBundleArgs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FlashbotsSendBundle: %+v\n", sendResult)

	//exit program

	return nil

	//***UPDATE THE BRIBE PERCENTAGE IN THE PAYLOAD***

	//estimate gas for the flashswap transaction

	//create the rest of the transactions(pool 1-n)
	//for each
	//create transaction for each depending on which token
	//swap(amountout1, amountout2, to, data)
	//

	//We will do the falshswap transaction for pool0 last, becuase we need to pass in the calldata for the rest of the transactions
}

//function that multiplies big.Int with a decimal
