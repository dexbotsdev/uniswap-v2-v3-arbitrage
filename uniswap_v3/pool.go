package uniswap_v3

import (
	"fmt"
	"math/big"
	"mev-template-go/pool_interface"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// type PoolInterface interface {
// 	GetAddress() common.Address
// 	GetFactoryAddress() common.Address
// 	GetTokens() []common.Address
// 	GetAmountOut(amountIn *big.Int, zeroForOne bool) (*big.Int, PoolInterface, error)
// 	GetProtocol() string
// 	GetCopy() PoolInterface
// }

type Pool struct {
	Address        common.Address
	FactoryAddress common.Address

	Token0 common.Address
	Token1 common.Address

	SqrtPriceX96 *big.Int
	Liquidity    *big.Int

	TickCurrent int
	TickSpacing int

	//PopulatedTicks []PopulatedTick
	//TickMap map[int]Tick
	Ticks []Tick

	Fee uint32
}

func (p *Pool) String() string {
	return fmt.Sprintf(
		"Pool{\n  Address: %s,\n  FactoryAddress: %s,\n  Token0: %s,\n  Token1: %s,\n  SqrtPriceX96: %s,\n  Liquidity: %s,\n  TickCurrent: %d,\n  TickSpacing: %d,\n  Fee: %d\n}",
		p.Address.Hex(),
		p.FactoryAddress.Hex(),
		p.Token0.Hex(),
		p.Token1.Hex(),
		p.SqrtPriceX96.String(),
		p.Liquidity.String(),
		p.TickCurrent,
		p.TickSpacing,
		p.Fee,
	)
}

func (pool *Pool) Update(client *ethclient.Client) error {
	err := UpdatePools([]*Pool{pool}, client)
	if err != nil {
		return err
	}

	return UpdateTicks([]*Pool{pool}, client)
}

func (pool *Pool) GetCopyInterface() pool_interface.PoolInterface {
	poolCopy := *pool.GetCopy()
	return &poolCopy
}

func (pool *Pool) GetCopy() *Pool {
	return &Pool{
		Address:        pool.Address,
		FactoryAddress: pool.FactoryAddress,
		Token0:         pool.Token0,
		Token1:         pool.Token1,
		SqrtPriceX96:   new(big.Int).Set(pool.SqrtPriceX96),
		Liquidity:      new(big.Int).Set(pool.Liquidity),
		TickCurrent:    pool.TickCurrent,
		TickSpacing:    pool.TickSpacing,
		//TickMap:        CopyTickMapCopy(pool.TickMap),
		Ticks: pool.GetTicksCopy(),
		Fee:   pool.Fee,
	}
}
func (pool *Pool) GetTicksCopy() []Tick {
	ticksCopy := make([]Tick, len(pool.Ticks))
	for i, tick := range pool.Ticks {
		ticksCopy[i] = tick.GetCopy()
	}
	return ticksCopy
}

func CopyTickMapCopy(tickMap map[int]Tick) map[int]Tick {
	tickMapCopy := make(map[int]Tick)
	for key, value := range tickMap {
		tickMapCopy[key] = value.GetCopy()
	}
	return tickMapCopy
}

// func CreateTickMap(ticks []Tick) map[int]Tick {
// 	tickMap := make(map[int]Tick)
// 	for _, tick := range ticks {
// 		tickMap[tick.Index] = tick
// 	}
// 	return tickMap
// }
func (pool *Pool) GetAmountOutAndUpdatePool(amountIn *big.Int, zeroForOne bool) (*big.Int, error) {
	amountOut, poolStateAfter, err := GetAmountOutAndUpdatePool(*pool, amountIn, zeroForOne)
	if err != nil {
		return nil, err
	}
	//update pool
	pool = &poolStateAfter
	return amountOut, nil
}

func (pool *Pool) GetAmountOut(amountIn *big.Int, zeroForOne bool) (*big.Int, error) {
	amountOut, err := GetAmountOut(*pool, amountIn, zeroForOne)
	if err != nil {
		return nil, err
	}
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
	return "uniswap_v3"
}

func (pool *Pool) UpdateTick(newTick Tick) error {
	//pool.TickMap[newTick.Index] = newTick
	//find tick by index
	for i, tick := range pool.Ticks {
		if tick.Index == newTick.Index {
			pool.Ticks[i] = newTick
		}
	}
	return nil
}
