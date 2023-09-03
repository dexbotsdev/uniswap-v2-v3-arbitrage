package pool_interface

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//pool interface
type PoolInterface interface {
	GetAddress() common.Address
	GetFactoryAddress() common.Address
	GetCopyInterface() PoolInterface
	GetAmountOutAndUpdatePool(amountIn *big.Int, zeroForOne bool) (*big.Int, error)
	GetAmountOut(amountIn *big.Int, zeroForOne bool) (*big.Int, error)
	GetTokens() []common.Address
	GetType() string
	Update(client *ethclient.Client) error
	String() string
}
