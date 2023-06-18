package types

import (
	"crypto/ecdsa"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Config struct {
	Client        ethclient.Client
	ClientWss     ethclient.Client
	RpcClient     rpc.Client
	PrivateKey    *ecdsa.PrivateKey
	WalletAddress common.Address
}

type UniV2Pool struct {
	Address        common.Address
	FactoryAddress common.Address

	Token0 common.Address
	Token1 common.Address

	Reserve0 *big.Int
	Reserve1 *big.Int

	Fees uint16
}

//get token array from pool transactions
type PoolTransaciton struct {
	Pools      []UniV2Pool
	Tokens     []common.Address
	reserves   []*big.Int
	AmountIns  []*big.Int
	AmountOuts []*big.Int
}

type Path struct {
	Pools       []UniV2Pool
	BaseToken   common.Address
	AmountOut   *big.Int
	AmountIn    *big.Int
	Revenue     *big.Int //only used if startToken == endToken
	BlockNumber int
}

type State struct {
	LastCheckedBlockNumber uint64

	Paths []Path
	Pools []UniV2Pool

	PoolToPathsMap map[common.Address][]Path
	TokenToPathMap map[common.Address][]Path
	TokenToPoolMap map[common.Address][]UniV2Pool
}
