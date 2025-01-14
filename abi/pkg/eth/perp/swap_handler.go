// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package perp

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

// SwapHandlerMetaData contains all meta data concerning the SwapHandler contract.
var SwapHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDataBase\",\"name\":\"_dataBase\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"DuplicateToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"InvalidIncrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NoTokenAmountIn\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decAmount\",\"type\":\"uint256\"}],\"name\":\"PoolAmountExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bufferAmount\",\"type\":\"uint256\"}],\"name\":\"PoolLessThanBuffer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservedAmount\",\"type\":\"uint256\"}],\"name\":\"ReservedAmountExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"}],\"name\":\"UnauthorizedGov\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeUsd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeTokens\",\"type\":\"uint256\"}],\"name\":\"CollectSwapFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreasePoolAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DecreaseUsdgAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreasePoolAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"IncreaseUsdgAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutAfterFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBasisPoints\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"dataBase\",\"outputs\":[{\"internalType\":\"contractDataBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SwapHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapHandlerMetaData.ABI instead.
var SwapHandlerABI = SwapHandlerMetaData.ABI

// SwapHandler is an auto generated Go binding around an Ethereum contract.
type SwapHandler struct {
	SwapHandlerCaller     // Read-only binding to the contract
	SwapHandlerTransactor // Write-only binding to the contract
	SwapHandlerFilterer   // Log filterer for contract events
}

// SwapHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapHandlerSession struct {
	Contract     *SwapHandler      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapHandlerCallerSession struct {
	Contract *SwapHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SwapHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapHandlerTransactorSession struct {
	Contract     *SwapHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwapHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapHandlerRaw struct {
	Contract *SwapHandler // Generic contract binding to access the raw methods on
}

// SwapHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapHandlerCallerRaw struct {
	Contract *SwapHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// SwapHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapHandlerTransactorRaw struct {
	Contract *SwapHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapHandler creates a new instance of SwapHandler, bound to a specific deployed contract.
func NewSwapHandler(address common.Address, backend bind.ContractBackend) (*SwapHandler, error) {
	contract, err := bindSwapHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapHandler{SwapHandlerCaller: SwapHandlerCaller{contract: contract}, SwapHandlerTransactor: SwapHandlerTransactor{contract: contract}, SwapHandlerFilterer: SwapHandlerFilterer{contract: contract}}, nil
}

// NewSwapHandlerCaller creates a new read-only instance of SwapHandler, bound to a specific deployed contract.
func NewSwapHandlerCaller(address common.Address, caller bind.ContractCaller) (*SwapHandlerCaller, error) {
	contract, err := bindSwapHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapHandlerCaller{contract: contract}, nil
}

// NewSwapHandlerTransactor creates a new write-only instance of SwapHandler, bound to a specific deployed contract.
func NewSwapHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapHandlerTransactor, error) {
	contract, err := bindSwapHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapHandlerTransactor{contract: contract}, nil
}

// NewSwapHandlerFilterer creates a new log filterer instance of SwapHandler, bound to a specific deployed contract.
func NewSwapHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapHandlerFilterer, error) {
	contract, err := bindSwapHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapHandlerFilterer{contract: contract}, nil
}

// bindSwapHandler binds a generic wrapper to an already deployed contract.
func bindSwapHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapHandler *SwapHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapHandler.Contract.SwapHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapHandler *SwapHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapHandler.Contract.SwapHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapHandler *SwapHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapHandler.Contract.SwapHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapHandler *SwapHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapHandler *SwapHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapHandler *SwapHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapHandler.Contract.contract.Transact(opts, method, params...)
}

// DataBase is a free data retrieval call binding the contract method 0x5cb5e41d.
//
// Solidity: function dataBase() view returns(address)
func (_SwapHandler *SwapHandlerCaller) DataBase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapHandler.contract.Call(opts, &out, "dataBase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataBase is a free data retrieval call binding the contract method 0x5cb5e41d.
//
// Solidity: function dataBase() view returns(address)
func (_SwapHandler *SwapHandlerSession) DataBase() (common.Address, error) {
	return _SwapHandler.Contract.DataBase(&_SwapHandler.CallOpts)
}

// DataBase is a free data retrieval call binding the contract method 0x5cb5e41d.
//
// Solidity: function dataBase() view returns(address)
func (_SwapHandler *SwapHandlerCallerSession) DataBase() (common.Address, error) {
	return _SwapHandler.Contract.DataBase(&_SwapHandler.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_SwapHandler *SwapHandlerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapHandler.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_SwapHandler *SwapHandlerSession) Gov() (common.Address, error) {
	return _SwapHandler.Contract.Gov(&_SwapHandler.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_SwapHandler *SwapHandlerCallerSession) Gov() (common.Address, error) {
	return _SwapHandler.Contract.Gov(&_SwapHandler.CallOpts)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_SwapHandler *SwapHandlerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _SwapHandler.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_SwapHandler *SwapHandlerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _SwapHandler.Contract.SetGov(&_SwapHandler.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_SwapHandler *SwapHandlerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _SwapHandler.Contract.SetGov(&_SwapHandler.TransactOpts, _gov)
}

// Swap is a paid mutator transaction binding the contract method 0x8c256d64.
//
// Solidity: function swap(address _poolToken, address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_SwapHandler *SwapHandlerTransactor) Swap(opts *bind.TransactOpts, _poolToken common.Address, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _SwapHandler.contract.Transact(opts, "swap", _poolToken, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x8c256d64.
//
// Solidity: function swap(address _poolToken, address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_SwapHandler *SwapHandlerSession) Swap(_poolToken common.Address, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _SwapHandler.Contract.Swap(&_SwapHandler.TransactOpts, _poolToken, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x8c256d64.
//
// Solidity: function swap(address _poolToken, address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_SwapHandler *SwapHandlerTransactorSession) Swap(_poolToken common.Address, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _SwapHandler.Contract.Swap(&_SwapHandler.TransactOpts, _poolToken, _tokenIn, _tokenOut, _receiver)
}

// SwapHandlerCollectSwapFeesIterator is returned from FilterCollectSwapFees and is used to iterate over the raw logs and unpacked data for CollectSwapFees events raised by the SwapHandler contract.
type SwapHandlerCollectSwapFeesIterator struct {
	Event *SwapHandlerCollectSwapFees // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapHandlerCollectSwapFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapHandlerCollectSwapFees)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapHandlerCollectSwapFees)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapHandlerCollectSwapFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapHandlerCollectSwapFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapHandlerCollectSwapFees represents a CollectSwapFees event raised by the SwapHandler contract.
type SwapHandlerCollectSwapFees struct {
	PoolToken common.Address
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectSwapFees is a free log retrieval operation binding the contract event 0x2dc61bb387e0ef3859238e8604e85e2e597a16b4d456d45832df57701e11d9df.
//
// Solidity: event CollectSwapFees(address poolToken, address token, uint256 feeUsd, uint256 feeTokens)
func (_SwapHandler *SwapHandlerFilterer) FilterCollectSwapFees(opts *bind.FilterOpts) (*SwapHandlerCollectSwapFeesIterator, error) {

	logs, sub, err := _SwapHandler.contract.FilterLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return &SwapHandlerCollectSwapFeesIterator{contract: _SwapHandler.contract, event: "CollectSwapFees", logs: logs, sub: sub}, nil
}

// WatchCollectSwapFees is a free log subscription operation binding the contract event 0x2dc61bb387e0ef3859238e8604e85e2e597a16b4d456d45832df57701e11d9df.
//
// Solidity: event CollectSwapFees(address poolToken, address token, uint256 feeUsd, uint256 feeTokens)
func (_SwapHandler *SwapHandlerFilterer) WatchCollectSwapFees(opts *bind.WatchOpts, sink chan<- *SwapHandlerCollectSwapFees) (event.Subscription, error) {

	logs, sub, err := _SwapHandler.contract.WatchLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapHandlerCollectSwapFees)
				if err := _SwapHandler.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollectSwapFees is a log parse operation binding the contract event 0x2dc61bb387e0ef3859238e8604e85e2e597a16b4d456d45832df57701e11d9df.
//
// Solidity: event CollectSwapFees(address poolToken, address token, uint256 feeUsd, uint256 feeTokens)
func (_SwapHandler *SwapHandlerFilterer) ParseCollectSwapFees(log types.Log) (*SwapHandlerCollectSwapFees, error) {
	event := new(SwapHandlerCollectSwapFees)
	if err := _SwapHandler.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapHandlerDecreasePoolAmountIterator is returned from FilterDecreasePoolAmount and is used to iterate over the raw logs and unpacked data for DecreasePoolAmount events raised by the SwapHandler contract.
type SwapHandlerDecreasePoolAmountIterator struct {
	Event *SwapHandlerDecreasePoolAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapHandlerDecreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapHandlerDecreasePoolAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapHandlerDecreasePoolAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapHandlerDecreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapHandlerDecreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapHandlerDecreasePoolAmount represents a DecreasePoolAmount event raised by the SwapHandler contract.
type SwapHandlerDecreasePoolAmount struct {
	PoolToken common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecreasePoolAmount is a free log retrieval operation binding the contract event 0x668e78c952f6dbefc8e698ded1b5581ae1ba845b454ebd375e07617a28f1cdaa.
//
// Solidity: event DecreasePoolAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) FilterDecreasePoolAmount(opts *bind.FilterOpts) (*SwapHandlerDecreasePoolAmountIterator, error) {

	logs, sub, err := _SwapHandler.contract.FilterLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &SwapHandlerDecreasePoolAmountIterator{contract: _SwapHandler.contract, event: "DecreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchDecreasePoolAmount is a free log subscription operation binding the contract event 0x668e78c952f6dbefc8e698ded1b5581ae1ba845b454ebd375e07617a28f1cdaa.
//
// Solidity: event DecreasePoolAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) WatchDecreasePoolAmount(opts *bind.WatchOpts, sink chan<- *SwapHandlerDecreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _SwapHandler.contract.WatchLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapHandlerDecreasePoolAmount)
				if err := _SwapHandler.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreasePoolAmount is a log parse operation binding the contract event 0x668e78c952f6dbefc8e698ded1b5581ae1ba845b454ebd375e07617a28f1cdaa.
//
// Solidity: event DecreasePoolAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) ParseDecreasePoolAmount(log types.Log) (*SwapHandlerDecreasePoolAmount, error) {
	event := new(SwapHandlerDecreasePoolAmount)
	if err := _SwapHandler.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapHandlerDecreaseUsdgAmountIterator is returned from FilterDecreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for DecreaseUsdgAmount events raised by the SwapHandler contract.
type SwapHandlerDecreaseUsdgAmountIterator struct {
	Event *SwapHandlerDecreaseUsdgAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapHandlerDecreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapHandlerDecreaseUsdgAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapHandlerDecreaseUsdgAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapHandlerDecreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapHandlerDecreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapHandlerDecreaseUsdgAmount represents a DecreaseUsdgAmount event raised by the SwapHandler contract.
type SwapHandlerDecreaseUsdgAmount struct {
	PoolToken common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDecreaseUsdgAmount is a free log retrieval operation binding the contract event 0x0adbf3bcfaf4bcb942ec2d941daf489fc8ca882439500da38271b85e3fdd42bf.
//
// Solidity: event DecreaseUsdgAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) FilterDecreaseUsdgAmount(opts *bind.FilterOpts) (*SwapHandlerDecreaseUsdgAmountIterator, error) {

	logs, sub, err := _SwapHandler.contract.FilterLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &SwapHandlerDecreaseUsdgAmountIterator{contract: _SwapHandler.contract, event: "DecreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseUsdgAmount is a free log subscription operation binding the contract event 0x0adbf3bcfaf4bcb942ec2d941daf489fc8ca882439500da38271b85e3fdd42bf.
//
// Solidity: event DecreaseUsdgAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) WatchDecreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *SwapHandlerDecreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _SwapHandler.contract.WatchLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapHandlerDecreaseUsdgAmount)
				if err := _SwapHandler.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDecreaseUsdgAmount is a log parse operation binding the contract event 0x0adbf3bcfaf4bcb942ec2d941daf489fc8ca882439500da38271b85e3fdd42bf.
//
// Solidity: event DecreaseUsdgAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) ParseDecreaseUsdgAmount(log types.Log) (*SwapHandlerDecreaseUsdgAmount, error) {
	event := new(SwapHandlerDecreaseUsdgAmount)
	if err := _SwapHandler.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapHandlerIncreasePoolAmountIterator is returned from FilterIncreasePoolAmount and is used to iterate over the raw logs and unpacked data for IncreasePoolAmount events raised by the SwapHandler contract.
type SwapHandlerIncreasePoolAmountIterator struct {
	Event *SwapHandlerIncreasePoolAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapHandlerIncreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapHandlerIncreasePoolAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapHandlerIncreasePoolAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapHandlerIncreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapHandlerIncreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapHandlerIncreasePoolAmount represents a IncreasePoolAmount event raised by the SwapHandler contract.
type SwapHandlerIncreasePoolAmount struct {
	PoolToken common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreasePoolAmount is a free log retrieval operation binding the contract event 0x0a102514ed687d8dbf0013c75f431f4941f4c341a1af5c5ccc67007655e0e285.
//
// Solidity: event IncreasePoolAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) FilterIncreasePoolAmount(opts *bind.FilterOpts) (*SwapHandlerIncreasePoolAmountIterator, error) {

	logs, sub, err := _SwapHandler.contract.FilterLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &SwapHandlerIncreasePoolAmountIterator{contract: _SwapHandler.contract, event: "IncreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchIncreasePoolAmount is a free log subscription operation binding the contract event 0x0a102514ed687d8dbf0013c75f431f4941f4c341a1af5c5ccc67007655e0e285.
//
// Solidity: event IncreasePoolAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) WatchIncreasePoolAmount(opts *bind.WatchOpts, sink chan<- *SwapHandlerIncreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _SwapHandler.contract.WatchLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapHandlerIncreasePoolAmount)
				if err := _SwapHandler.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreasePoolAmount is a log parse operation binding the contract event 0x0a102514ed687d8dbf0013c75f431f4941f4c341a1af5c5ccc67007655e0e285.
//
// Solidity: event IncreasePoolAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) ParseIncreasePoolAmount(log types.Log) (*SwapHandlerIncreasePoolAmount, error) {
	event := new(SwapHandlerIncreasePoolAmount)
	if err := _SwapHandler.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapHandlerIncreaseUsdgAmountIterator is returned from FilterIncreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for IncreaseUsdgAmount events raised by the SwapHandler contract.
type SwapHandlerIncreaseUsdgAmountIterator struct {
	Event *SwapHandlerIncreaseUsdgAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapHandlerIncreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapHandlerIncreaseUsdgAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapHandlerIncreaseUsdgAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapHandlerIncreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapHandlerIncreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapHandlerIncreaseUsdgAmount represents a IncreaseUsdgAmount event raised by the SwapHandler contract.
type SwapHandlerIncreaseUsdgAmount struct {
	PoolToken common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIncreaseUsdgAmount is a free log retrieval operation binding the contract event 0x30c9b3edd8bfa73ddd9ee505cf47618bd49feae1fe7cc26b8c6a39de21c18025.
//
// Solidity: event IncreaseUsdgAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) FilterIncreaseUsdgAmount(opts *bind.FilterOpts) (*SwapHandlerIncreaseUsdgAmountIterator, error) {

	logs, sub, err := _SwapHandler.contract.FilterLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &SwapHandlerIncreaseUsdgAmountIterator{contract: _SwapHandler.contract, event: "IncreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseUsdgAmount is a free log subscription operation binding the contract event 0x30c9b3edd8bfa73ddd9ee505cf47618bd49feae1fe7cc26b8c6a39de21c18025.
//
// Solidity: event IncreaseUsdgAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) WatchIncreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *SwapHandlerIncreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _SwapHandler.contract.WatchLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapHandlerIncreaseUsdgAmount)
				if err := _SwapHandler.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIncreaseUsdgAmount is a log parse operation binding the contract event 0x30c9b3edd8bfa73ddd9ee505cf47618bd49feae1fe7cc26b8c6a39de21c18025.
//
// Solidity: event IncreaseUsdgAmount(address poolToken, address token, uint256 amount)
func (_SwapHandler *SwapHandlerFilterer) ParseIncreaseUsdgAmount(log types.Log) (*SwapHandlerIncreaseUsdgAmount, error) {
	event := new(SwapHandlerIncreaseUsdgAmount)
	if err := _SwapHandler.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapHandlerSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SwapHandler contract.
type SwapHandlerSwapIterator struct {
	Event *SwapHandlerSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapHandlerSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapHandlerSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapHandlerSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapHandlerSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapHandlerSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapHandlerSwap represents a Swap event raised by the SwapHandler contract.
type SwapHandlerSwap struct {
	Account            common.Address
	PoolToken          common.Address
	TokenIn            common.Address
	TokenOut           common.Address
	AmountIn           *big.Int
	AmountOut          *big.Int
	AmountOutAfterFees *big.Int
	FeeBasisPoints     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x6950339c7661cca450281e53722525cc136590e622b011d5be7e4c4993685a6c.
//
// Solidity: event Swap(address account, address poolToken, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_SwapHandler *SwapHandlerFilterer) FilterSwap(opts *bind.FilterOpts) (*SwapHandlerSwapIterator, error) {

	logs, sub, err := _SwapHandler.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &SwapHandlerSwapIterator{contract: _SwapHandler.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x6950339c7661cca450281e53722525cc136590e622b011d5be7e4c4993685a6c.
//
// Solidity: event Swap(address account, address poolToken, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_SwapHandler *SwapHandlerFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SwapHandlerSwap) (event.Subscription, error) {

	logs, sub, err := _SwapHandler.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapHandlerSwap)
				if err := _SwapHandler.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0x6950339c7661cca450281e53722525cc136590e622b011d5be7e4c4993685a6c.
//
// Solidity: event Swap(address account, address poolToken, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_SwapHandler *SwapHandlerFilterer) ParseSwap(log types.Log) (*SwapHandlerSwap, error) {
	event := new(SwapHandlerSwap)
	if err := _SwapHandler.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
