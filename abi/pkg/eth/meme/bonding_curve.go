// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package meme

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
	_ = abi.ConvertType
)

// BondingCurveMetaData contains all meta data concerning the BondingCurve contract.
var BondingCurveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCost\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"migrator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"migrateToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minOut\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BondingCurveABI is the input ABI used to generate the binding from.
// Deprecated: Use BondingCurveMetaData.ABI instead.
var BondingCurveABI = BondingCurveMetaData.ABI

// BondingCurve is an auto generated Go binding around an Ethereum contract.
type BondingCurve struct {
	BondingCurveCaller     // Read-only binding to the contract
	BondingCurveTransactor // Write-only binding to the contract
	BondingCurveFilterer   // Log filterer for contract events
}

// BondingCurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type BondingCurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingCurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BondingCurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingCurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BondingCurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BondingCurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BondingCurveSession struct {
	Contract     *BondingCurve     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BondingCurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BondingCurveCallerSession struct {
	Contract *BondingCurveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BondingCurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BondingCurveTransactorSession struct {
	Contract     *BondingCurveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BondingCurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type BondingCurveRaw struct {
	Contract *BondingCurve // Generic contract binding to access the raw methods on
}

// BondingCurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BondingCurveCallerRaw struct {
	Contract *BondingCurveCaller // Generic read-only contract binding to access the raw methods on
}

// BondingCurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BondingCurveTransactorRaw struct {
	Contract *BondingCurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBondingCurve creates a new instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurve(address common.Address, backend bind.ContractBackend) (*BondingCurve, error) {
	contract, err := bindBondingCurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BondingCurve{BondingCurveCaller: BondingCurveCaller{contract: contract}, BondingCurveTransactor: BondingCurveTransactor{contract: contract}, BondingCurveFilterer: BondingCurveFilterer{contract: contract}}, nil
}

// NewBondingCurveCaller creates a new read-only instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurveCaller(address common.Address, caller bind.ContractCaller) (*BondingCurveCaller, error) {
	contract, err := bindBondingCurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BondingCurveCaller{contract: contract}, nil
}

// NewBondingCurveTransactor creates a new write-only instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurveTransactor(address common.Address, transactor bind.ContractTransactor) (*BondingCurveTransactor, error) {
	contract, err := bindBondingCurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BondingCurveTransactor{contract: contract}, nil
}

// NewBondingCurveFilterer creates a new log filterer instance of BondingCurve, bound to a specific deployed contract.
func NewBondingCurveFilterer(address common.Address, filterer bind.ContractFilterer) (*BondingCurveFilterer, error) {
	contract, err := bindBondingCurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BondingCurveFilterer{contract: contract}, nil
}

// bindBondingCurve binds a generic wrapper to an already deployed contract.
func bindBondingCurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BondingCurveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingCurve *BondingCurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingCurve.Contract.BondingCurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingCurve *BondingCurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.Contract.BondingCurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingCurve *BondingCurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingCurve.Contract.BondingCurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BondingCurve *BondingCurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BondingCurve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BondingCurve *BondingCurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BondingCurve *BondingCurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BondingCurve.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_BondingCurve *BondingCurveCaller) GetReserves(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _BondingCurve.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_BondingCurve *BondingCurveSession) GetReserves() (*big.Int, *big.Int, error) {
	return _BondingCurve.Contract.GetReserves(&_BondingCurve.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_BondingCurve *BondingCurveCallerSession) GetReserves() (*big.Int, *big.Int, error) {
	return _BondingCurve.Contract.GetReserves(&_BondingCurve.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 amount, uint256 maxCost) returns()
func (_BondingCurve *BondingCurveTransactor) Buy(opts *bind.TransactOpts, amount *big.Int, maxCost *big.Int) (*types.Transaction, error) {
	return _BondingCurve.contract.Transact(opts, "buy", amount, maxCost)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 amount, uint256 maxCost) returns()
func (_BondingCurve *BondingCurveSession) Buy(amount *big.Int, maxCost *big.Int) (*types.Transaction, error) {
	return _BondingCurve.Contract.Buy(&_BondingCurve.TransactOpts, amount, maxCost)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 amount, uint256 maxCost) returns()
func (_BondingCurve *BondingCurveTransactorSession) Buy(amount *big.Int, maxCost *big.Int) (*types.Transaction, error) {
	return _BondingCurve.Contract.Buy(&_BondingCurve.TransactOpts, amount, maxCost)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address token0, address token1, address migrator) returns()
func (_BondingCurve *BondingCurveTransactor) Initialize(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, migrator common.Address) (*types.Transaction, error) {
	return _BondingCurve.contract.Transact(opts, "initialize", token0, token1, migrator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address token0, address token1, address migrator) returns()
func (_BondingCurve *BondingCurveSession) Initialize(token0 common.Address, token1 common.Address, migrator common.Address) (*types.Transaction, error) {
	return _BondingCurve.Contract.Initialize(&_BondingCurve.TransactOpts, token0, token1, migrator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address token0, address token1, address migrator) returns()
func (_BondingCurve *BondingCurveTransactorSession) Initialize(token0 common.Address, token1 common.Address, migrator common.Address) (*types.Transaction, error) {
	return _BondingCurve.Contract.Initialize(&_BondingCurve.TransactOpts, token0, token1, migrator)
}

// MigrateToken is a paid mutator transaction binding the contract method 0xd5482804.
//
// Solidity: function migrateToken() returns(uint256 amount)
func (_BondingCurve *BondingCurveTransactor) MigrateToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.contract.Transact(opts, "migrateToken")
}

// MigrateToken is a paid mutator transaction binding the contract method 0xd5482804.
//
// Solidity: function migrateToken() returns(uint256 amount)
func (_BondingCurve *BondingCurveSession) MigrateToken() (*types.Transaction, error) {
	return _BondingCurve.Contract.MigrateToken(&_BondingCurve.TransactOpts)
}

// MigrateToken is a paid mutator transaction binding the contract method 0xd5482804.
//
// Solidity: function migrateToken() returns(uint256 amount)
func (_BondingCurve *BondingCurveTransactorSession) MigrateToken() (*types.Transaction, error) {
	return _BondingCurve.Contract.MigrateToken(&_BondingCurve.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x67b9a286.
//
// Solidity: function removeLiquidity() returns(uint256 amount0, uint256 amount1)
func (_BondingCurve *BondingCurveTransactor) RemoveLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.contract.Transact(opts, "removeLiquidity")
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x67b9a286.
//
// Solidity: function removeLiquidity() returns(uint256 amount0, uint256 amount1)
func (_BondingCurve *BondingCurveSession) RemoveLiquidity() (*types.Transaction, error) {
	return _BondingCurve.Contract.RemoveLiquidity(&_BondingCurve.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x67b9a286.
//
// Solidity: function removeLiquidity() returns(uint256 amount0, uint256 amount1)
func (_BondingCurve *BondingCurveTransactorSession) RemoveLiquidity() (*types.Transaction, error) {
	return _BondingCurve.Contract.RemoveLiquidity(&_BondingCurve.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 amount, uint256 minOut) returns()
func (_BondingCurve *BondingCurveTransactor) Sell(opts *bind.TransactOpts, amount *big.Int, minOut *big.Int) (*types.Transaction, error) {
	return _BondingCurve.contract.Transact(opts, "sell", amount, minOut)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 amount, uint256 minOut) returns()
func (_BondingCurve *BondingCurveSession) Sell(amount *big.Int, minOut *big.Int) (*types.Transaction, error) {
	return _BondingCurve.Contract.Sell(&_BondingCurve.TransactOpts, amount, minOut)
}

// Sell is a paid mutator transaction binding the contract method 0xd79875eb.
//
// Solidity: function sell(uint256 amount, uint256 minOut) returns()
func (_BondingCurve *BondingCurveTransactorSession) Sell(amount *big.Int, minOut *big.Int) (*types.Transaction, error) {
	return _BondingCurve.Contract.Sell(&_BondingCurve.TransactOpts, amount, minOut)
}

// Token0 is a paid mutator transaction binding the contract method 0x0dfe1681.
//
// Solidity: function token0() returns(address)
func (_BondingCurve *BondingCurveTransactor) Token0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.contract.Transact(opts, "token0")
}

// Token0 is a paid mutator transaction binding the contract method 0x0dfe1681.
//
// Solidity: function token0() returns(address)
func (_BondingCurve *BondingCurveSession) Token0() (*types.Transaction, error) {
	return _BondingCurve.Contract.Token0(&_BondingCurve.TransactOpts)
}

// Token0 is a paid mutator transaction binding the contract method 0x0dfe1681.
//
// Solidity: function token0() returns(address)
func (_BondingCurve *BondingCurveTransactorSession) Token0() (*types.Transaction, error) {
	return _BondingCurve.Contract.Token0(&_BondingCurve.TransactOpts)
}

// Token1 is a paid mutator transaction binding the contract method 0xd21220a7.
//
// Solidity: function token1() returns(address)
func (_BondingCurve *BondingCurveTransactor) Token1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BondingCurve.contract.Transact(opts, "token1")
}

// Token1 is a paid mutator transaction binding the contract method 0xd21220a7.
//
// Solidity: function token1() returns(address)
func (_BondingCurve *BondingCurveSession) Token1() (*types.Transaction, error) {
	return _BondingCurve.Contract.Token1(&_BondingCurve.TransactOpts)
}

// Token1 is a paid mutator transaction binding the contract method 0xd21220a7.
//
// Solidity: function token1() returns(address)
func (_BondingCurve *BondingCurveTransactorSession) Token1() (*types.Transaction, error) {
	return _BondingCurve.Contract.Token1(&_BondingCurve.TransactOpts)
}
