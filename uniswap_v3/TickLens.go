// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap_v3

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PopulatedTick is an auto generated low-level Go binding around an user-defined struct.
type PopulatedTick struct {
	Tick           *big.Int
	LiquidityNet   *big.Int
	LiquidityGross *big.Int
}

// UniswapV3MetaData contains all meta data concerning the UniswapV3 contract.
var UniswapV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"int16\",\"name\":\"tickBitmapIndex\",\"type\":\"int16\"}],\"name\":\"getPopulatedTicksInWord\",\"outputs\":[{\"components\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"internalType\":\"int128\",\"name\":\"liquidityNet\",\"type\":\"int128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityGross\",\"type\":\"uint128\"}],\"internalType\":\"structPopulatedTick[]\",\"name\":\"populatedTicks\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3MetaData.ABI instead.
var UniswapV3ABI = UniswapV3MetaData.ABI

// UniswapV3 is an auto generated Go binding around an Ethereum contract.
type UniswapV3 struct {
	UniswapV3Caller     // Read-only binding to the contract
	UniswapV3Transactor // Write-only binding to the contract
	UniswapV3Filterer   // Log filterer for contract events
}

// UniswapV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3Session struct {
	Contract     *UniswapV3        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3CallerSession struct {
	Contract *UniswapV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UniswapV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3TransactorSession struct {
	Contract     *UniswapV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UniswapV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3Raw struct {
	Contract *UniswapV3 // Generic contract binding to access the raw methods on
}

// UniswapV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3CallerRaw struct {
	Contract *UniswapV3Caller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3TransactorRaw struct {
	Contract *UniswapV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3 creates a new instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3(address common.Address, backend bind.ContractBackend) (*UniswapV3, error) {
	contract, err := bindUniswapV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3{UniswapV3Caller: UniswapV3Caller{contract: contract}, UniswapV3Transactor: UniswapV3Transactor{contract: contract}, UniswapV3Filterer: UniswapV3Filterer{contract: contract}}, nil
}

// NewUniswapV3Caller creates a new read-only instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3Caller(address common.Address, caller bind.ContractCaller) (*UniswapV3Caller, error) {
	contract, err := bindUniswapV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Caller{contract: contract}, nil
}

// NewUniswapV3Transactor creates a new write-only instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3Transactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3Transactor, error) {
	contract, err := bindUniswapV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Transactor{contract: contract}, nil
}

// NewUniswapV3Filterer creates a new log filterer instance of UniswapV3, bound to a specific deployed contract.
func NewUniswapV3Filterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3Filterer, error) {
	contract, err := bindUniswapV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Filterer{contract: contract}, nil
}

// bindUniswapV3 binds a generic wrapper to an already deployed contract.
func bindUniswapV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapV3ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3 *UniswapV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3.Contract.UniswapV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3 *UniswapV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3.Contract.UniswapV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3 *UniswapV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3.Contract.UniswapV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3 *UniswapV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3 *UniswapV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3 *UniswapV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3.Contract.contract.Transact(opts, method, params...)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3 *UniswapV3Caller) GetPopulatedTicksInWord(opts *bind.CallOpts, pool common.Address, tickBitmapIndex int16) ([]PopulatedTick, error) {
	var out []interface{}
	err := _UniswapV3.contract.Call(opts, &out, "getPopulatedTicksInWord", pool, tickBitmapIndex)

	if err != nil {
		return *new([]PopulatedTick), err
	}

	out0 := *abi.ConvertType(out[0], new([]PopulatedTick)).(*[]PopulatedTick)

	return out0, err

}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3 *UniswapV3Session) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]PopulatedTick, error) {
	return _UniswapV3.Contract.GetPopulatedTicksInWord(&_UniswapV3.CallOpts, pool, tickBitmapIndex)
}

// GetPopulatedTicksInWord is a free data retrieval call binding the contract method 0x351fb478.
//
// Solidity: function getPopulatedTicksInWord(address pool, int16 tickBitmapIndex) view returns((int24,int128,uint128)[] populatedTicks)
func (_UniswapV3 *UniswapV3CallerSession) GetPopulatedTicksInWord(pool common.Address, tickBitmapIndex int16) ([]PopulatedTick, error) {
	return _UniswapV3.Contract.GetPopulatedTicksInWord(&_UniswapV3.CallOpts, pool, tickBitmapIndex)
}
