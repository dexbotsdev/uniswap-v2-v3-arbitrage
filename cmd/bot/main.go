package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"mev-template-go/logic"
	"mev-template-go/recon"
	"mev-template-go/types"

	"context"

	"github.com/ethereum/go-ethereum/common"
	geth_types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	log "github.com/inconshreveable/log15"
	"github.com/joho/godotenv"
)

var config types.Config

// Initialize function to initialize the client, private key, and wallet address
func Initialize() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("BOT_PRIVATE_KEY"))
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
	if err != nil {
		return err
	}

	clientWss, err := ethclient.Dial(os.Getenv("NETWORK_WSS"))
	if err != nil {
		return err
	}

	rpcClient, err := rpc.DialContext(context.Background(), os.Getenv("NETWORK_WSS"))
	if err != nil {
		return err
	}

	config = types.Config{
		Client:        *client,
		ClientWss:     *clientWss,
		RpcClient:     *rpcClient,
		PrivateKey:    privateKey,
		WalletAddress: walletAddress,
	}
	return nil
}

// func temp() {
// 	err := os.Chdir("../")
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}

// 	err = godotenv.Load(".env")
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}

// 	client, err := ethclient.Dial(os.Getenv("NETWORK_RPC"))
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}

// 	err = uniswap_v3.GetAllPoolsAndWriteToJson(client)
// 	if err != nil {
// 		fmt.Println("err: ", err)
// 		return
// 	}
// }
func main() {
	//temp()
	// Initialize the client, private key, and wallet address
	err := Initialize()
	if err != nil {
		log.Error("Initialization failed", "error", err)
	}

	fmt.Println("MEV TEMPLATE IN GO")
	fmt.Println("")

	//Start go routines
	receiveChannel := make(chan *geth_types.Transaction)
	blockChannel := make(chan *geth_types.Header)
	go recon.AlertTransaction(config, map[common.Address]bool{common.HexToAddress(uniV2RouterAddress): true}, receiveChannel)
	go recon.AlertBlocks(config, blockChannel)

	//setup global state
	//State Componenets
	//Pools
	//Paths
	//TokenToPoolMap
	//TokenToPathMap
	//PoolToPathMap

	//setup stat

	//err = SetupState(config)
	//if err != nil {
	//fmt.Println(err)
	//return
	//}

	//update revenue and amountin for all paths in file
	//err = updateAmountInsAndRevenuesInFile("paths.json")

	//read from these files

	//createThese files
	//filteredV2Pools.json
	//paths.json
	//***READ POOLS AND PATHS FROM FILE***
	pools, err := logic.GetFilteredPools()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Pools length ", len(pools))

	paths, err := logic.ReadPathsFromFile()
	fmt.Println("Paths length ", len(paths))

	//***SETUP MAPS***
	poolToPathsMap := make(map[common.Address][]logic.Path)

	//create poolToPathMap and remove duplicates
	for _, path := range paths {
		for _, pool := range path.Pools {
			poolToPathsMap[pool.GetAddress()] = append(poolToPathsMap[pool.GetAddress()], path)
		}
	}
	fmt.Println("poolToPathMap.length: ", len(poolToPathsMap))
	fmt.Println("Paths length: ", len(paths))

	//for all paths, set amount in and revenue to 0
	for _, path := range paths {
		path.AmountIn = big.NewInt(0)
		path.Revenue = big.NewInt(0)
	}

	//make poolAddressToPoolMap
	poolAddressToPoolMap := make(map[common.Address]types.PoolInterface)
	for _, pool := range pools {
		poolAddressToPoolMap[pool.GetAddress()] = pool
	}

	//execute path

	//convert path to pool transction objects
	//executor.ExecuteV2Path(paths[0], config)

	//create state object

	//TODO keep track of last checked block

	//Objects to keep track of
	//last checked block
	//pool objects
	//path objects
	//pool to path map

	for {
		//newTx := <-receiveChannel
		//uniV2.DecodeTransactionInputData(newTx.Data())

		//Steps
		//- check for reserve changes in new block
		//- update all relevant pools
		//- update all relevant paths by recalculating revenue and amountIn
		//- sort paths by revenue
		//- execute path with highest revenue

		header := <-blockChannel
		//check for reserve changes in new block. Runs for all blocks since last checked block
		err = HandleNewBlock(pools, poolToPathsMap, poolAddressToPoolMap, config, header)

		// affectedPoolAddresses, err := UpdatePoolsAndGetAffectedAddresses(state.Pools, header, config)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// fmt.Println("Affected pools length: ", len(affectedPoolAddresses))

		// //update last checked block
		// state.LastCheckedBlockNumber = header.Number.Uint64()

		// //get affect paths and remove duplicates
		// affectedPaths := make([]types.Path, 0)
		// for _, poolAddress := range affectedPoolAddresses {
		// 	affectedPaths = append(affectedPaths, state.PoolToPathsMap[poolAddress]...)
		// }
		// fmt.Println("Affected paths length: ", len(affectedPaths))
		// fmt.Println("affectedPaths", affectedPaths)
		// //remove duplicates
		// affectedPaths = removeDuplicatePaths(affectedPaths)
		// fmt.Println("Affected paths with duplicates removed length: ", len(affectedPaths))

		// //calculate AmountIn and Revenue for all affected paths
		// for i := 0; i < len(affectedPaths); i++ {
		// 	affectedPaths[i].AmountIn = new(big.Int).Set(uniswap_v2.GetBestAmountIn(affectedPaths[i]))
		// 	affectedPaths[i].Revenue = new(big.Int).Set(uniswap_v2.CalculateRevenue(affectedPaths[i]))
		// }

		// //sort pathsWithRevenue by revenue from largest to smallest
		// sort.Slice(affectedPaths, func(i, j int) bool {
		// 	return affectedPaths[i].Revenue.Cmp(affectedPaths[j].Revenue) > 0
		// })

		// //to execute a path
		// //update all reserve values on path
		// //recalculate revenue and amountIn
		// //execute path
		// if len(affectedPaths) > 0 {
		// 	err := executor.ExecuteV2Path(affectedPaths[0], config)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	fmt.Println("path executed")
		// 	return
		// }
	}
}
