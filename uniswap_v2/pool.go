package uniswap_v2

import (
	"fmt"
	"math/big"
	"mev-template-go/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Pool struct {
	Address        common.Address
	FactoryAddress common.Address
	Token0         common.Address
	Token1         common.Address
	Reserve0       *big.Int
	Reserve1       *big.Int
	Fee            uint32
}

func (p *Pool) String() string {
	return fmt.Sprintf(
		"Pool{\n  Address: %s,\n  FactoryAddress: %s,\n  Token0: %s,\n  Token1: %s,\n  Reserve0: %s,\n  Reserve1: %s,\n  Fee: %d\n}",
		p.Address.Hex(),
		p.FactoryAddress.Hex(),
		p.Token0.Hex(),
		p.Token1.Hex(),
		p.Reserve0.String(),
		p.Reserve1.String(),
		p.Fee,
	)
}

func (pool *Pool) Update(client *ethclient.Client) error {
	//update reserves
	r0, r1, err := GetReserves(pool.Address, client)
	if err != nil {
		return err
	}
	pool.Reserve0.Set(r0)
	pool.Reserve1.Set(r1)
	return nil
}

func (pool *Pool) GetCopy() *Pool {
	return &Pool{
		Address:        pool.Address,
		FactoryAddress: pool.FactoryAddress,
		Token0:         pool.Token0,
		Token1:         pool.Token1,
		Reserve0:       new(big.Int).Set(pool.Reserve0),
		Reserve1:       new(big.Int).Set(pool.Reserve1),
		Fee:            pool.Fee,
	}
}
func (pool *Pool) GetCopyInterface() types.PoolInterface {
	poolCopy := *pool.GetCopy()
	return &poolCopy
}

func (pool *Pool) GetAmountOutAndUpdatePool(amountIn *big.Int, zeroForOne bool) (*big.Int, error) {

	var reserveIn *big.Int
	var reserveOut *big.Int
	if zeroForOne {
		reserveIn = new(big.Int).Set(pool.Reserve0)
		reserveOut = new(big.Int).Set(pool.Reserve1)
	} else {
		reserveIn = new(big.Int).Set(pool.Reserve1)
		reserveOut = new(big.Int).Set(pool.Reserve0)
	}

	amountOutTemp, err := GetAmountOut(amountIn, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	amountOut := new(big.Int).Set(amountOutTemp)
	//update reserves

	reserveInAfter := new(big.Int).Add(reserveIn, amountIn)
	reserveOutAfter := new(big.Int).Sub(reserveOut, amountOut)

	if zeroForOne {
		pool.Reserve0.Set(reserveInAfter)
		pool.Reserve1.Set(reserveOutAfter)
	} else {
		pool.Reserve0.Set(reserveOutAfter)
		pool.Reserve1.Set(reserveInAfter)
	}

	return amountOut, nil
}

func (pool *Pool) GetAmountOut(amountIn *big.Int, zeroForOne bool) (*big.Int, error) {

	var reserveIn *big.Int
	var reserveOut *big.Int
	if zeroForOne {
		reserveIn = new(big.Int).Set(pool.Reserve0)
		reserveOut = new(big.Int).Set(pool.Reserve1)
	} else {
		reserveIn = new(big.Int).Set(pool.Reserve1)
		reserveOut = new(big.Int).Set(pool.Reserve0)
	}

	amountOutTemp, err := GetAmountOut(amountIn, reserveIn, reserveOut)
	if err != nil {
		return nil, err
	}
	amountOut := new(big.Int).Set(amountOutTemp)

	return amountOut, nil
}

func (pool *Pool) GetTokens() []common.Address {
	return []common.Address{pool.Token0, pool.Token1}
}

func (pool *Pool) GetAddress() common.Address {
	return pool.Address
}

func (pool *Pool) GetFactoryAddress() common.Address {
	return pool.FactoryAddress
}

func (pool *Pool) GetType() string {
	return "uniswap_v2"
}
