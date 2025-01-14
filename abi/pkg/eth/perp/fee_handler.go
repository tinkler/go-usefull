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

// FeeHandlerMetaData contains all meta data concerning the FeeHandler contract.
var FeeHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractDataBase\",\"name\":\"_dataBase\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"EmptyAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gov\",\"type\":\"address\"}],\"name\":\"UnauthorizedGov\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ClaimFeesPartner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"poolToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ClaimFeesTeam\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"}],\"name\":\"claimTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataBase\",\"outputs\":[{\"internalType\":\"contractDataBase\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"setPartnerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"setTeamWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolToken\",\"type\":\"address\"}],\"name\":\"teamWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FeeHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeHandlerMetaData.ABI instead.
var FeeHandlerABI = FeeHandlerMetaData.ABI

// FeeHandler is an auto generated Go binding around an Ethereum contract.
type FeeHandler struct {
	FeeHandlerCaller     // Read-only binding to the contract
	FeeHandlerTransactor // Write-only binding to the contract
	FeeHandlerFilterer   // Log filterer for contract events
}

// FeeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeHandlerSession struct {
	Contract     *FeeHandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeHandlerCallerSession struct {
	Contract *FeeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FeeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeHandlerTransactorSession struct {
	Contract     *FeeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FeeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeHandlerRaw struct {
	Contract *FeeHandler // Generic contract binding to access the raw methods on
}

// FeeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeHandlerCallerRaw struct {
	Contract *FeeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// FeeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeHandlerTransactorRaw struct {
	Contract *FeeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeHandler creates a new instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandler(address common.Address, backend bind.ContractBackend) (*FeeHandler, error) {
	contract, err := bindFeeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeHandler{FeeHandlerCaller: FeeHandlerCaller{contract: contract}, FeeHandlerTransactor: FeeHandlerTransactor{contract: contract}, FeeHandlerFilterer: FeeHandlerFilterer{contract: contract}}, nil
}

// NewFeeHandlerCaller creates a new read-only instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerCaller(address common.Address, caller bind.ContractCaller) (*FeeHandlerCaller, error) {
	contract, err := bindFeeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerCaller{contract: contract}, nil
}

// NewFeeHandlerTransactor creates a new write-only instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeHandlerTransactor, error) {
	contract, err := bindFeeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerTransactor{contract: contract}, nil
}

// NewFeeHandlerFilterer creates a new log filterer instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeHandlerFilterer, error) {
	contract, err := bindFeeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerFilterer{contract: contract}, nil
}

// bindFeeHandler binds a generic wrapper to an already deployed contract.
func bindFeeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandler *FeeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandler.Contract.FeeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandler *FeeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.Contract.FeeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandler *FeeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandler.Contract.FeeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandler *FeeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandler *FeeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandler *FeeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandler.Contract.contract.Transact(opts, method, params...)
}

// DataBase is a free data retrieval call binding the contract method 0x5cb5e41d.
//
// Solidity: function dataBase() view returns(address)
func (_FeeHandler *FeeHandlerCaller) DataBase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "dataBase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataBase is a free data retrieval call binding the contract method 0x5cb5e41d.
//
// Solidity: function dataBase() view returns(address)
func (_FeeHandler *FeeHandlerSession) DataBase() (common.Address, error) {
	return _FeeHandler.Contract.DataBase(&_FeeHandler.CallOpts)
}

// DataBase is a free data retrieval call binding the contract method 0x5cb5e41d.
//
// Solidity: function dataBase() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) DataBase() (common.Address, error) {
	return _FeeHandler.Contract.DataBase(&_FeeHandler.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FeeHandler *FeeHandlerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FeeHandler *FeeHandlerSession) Gov() (common.Address, error) {
	return _FeeHandler.Contract.Gov(&_FeeHandler.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) Gov() (common.Address, error) {
	return _FeeHandler.Contract.Gov(&_FeeHandler.CallOpts)
}

// TeamWallet is a free data retrieval call binding the contract method 0xaea9c018.
//
// Solidity: function teamWallet(address _poolToken) view returns(address)
func (_FeeHandler *FeeHandlerCaller) TeamWallet(opts *bind.CallOpts, _poolToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "teamWallet", _poolToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeamWallet is a free data retrieval call binding the contract method 0xaea9c018.
//
// Solidity: function teamWallet(address _poolToken) view returns(address)
func (_FeeHandler *FeeHandlerSession) TeamWallet(_poolToken common.Address) (common.Address, error) {
	return _FeeHandler.Contract.TeamWallet(&_FeeHandler.CallOpts, _poolToken)
}

// TeamWallet is a free data retrieval call binding the contract method 0xaea9c018.
//
// Solidity: function teamWallet(address _poolToken) view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) TeamWallet(_poolToken common.Address) (common.Address, error) {
	return _FeeHandler.Contract.TeamWallet(&_FeeHandler.CallOpts, _poolToken)
}

// ClaimTeam is a paid mutator transaction binding the contract method 0x145628e0.
//
// Solidity: function claimTeam(address _poolToken) returns()
func (_FeeHandler *FeeHandlerTransactor) ClaimTeam(opts *bind.TransactOpts, _poolToken common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "claimTeam", _poolToken)
}

// ClaimTeam is a paid mutator transaction binding the contract method 0x145628e0.
//
// Solidity: function claimTeam(address _poolToken) returns()
func (_FeeHandler *FeeHandlerSession) ClaimTeam(_poolToken common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.ClaimTeam(&_FeeHandler.TransactOpts, _poolToken)
}

// ClaimTeam is a paid mutator transaction binding the contract method 0x145628e0.
//
// Solidity: function claimTeam(address _poolToken) returns()
func (_FeeHandler *FeeHandlerTransactorSession) ClaimTeam(_poolToken common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.ClaimTeam(&_FeeHandler.TransactOpts, _poolToken)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FeeHandler *FeeHandlerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FeeHandler *FeeHandlerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetGov(&_FeeHandler.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetGov(&_FeeHandler.TransactOpts, _gov)
}

// SetPartnerWallet is a paid mutator transaction binding the contract method 0xcf87844f.
//
// Solidity: function setPartnerWallet(address _poolToken, address _wallet) returns()
func (_FeeHandler *FeeHandlerTransactor) SetPartnerWallet(opts *bind.TransactOpts, _poolToken common.Address, _wallet common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setPartnerWallet", _poolToken, _wallet)
}

// SetPartnerWallet is a paid mutator transaction binding the contract method 0xcf87844f.
//
// Solidity: function setPartnerWallet(address _poolToken, address _wallet) returns()
func (_FeeHandler *FeeHandlerSession) SetPartnerWallet(_poolToken common.Address, _wallet common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetPartnerWallet(&_FeeHandler.TransactOpts, _poolToken, _wallet)
}

// SetPartnerWallet is a paid mutator transaction binding the contract method 0xcf87844f.
//
// Solidity: function setPartnerWallet(address _poolToken, address _wallet) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetPartnerWallet(_poolToken common.Address, _wallet common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetPartnerWallet(&_FeeHandler.TransactOpts, _poolToken, _wallet)
}

// SetTeamWallet is a paid mutator transaction binding the contract method 0x67d3eff4.
//
// Solidity: function setTeamWallet(address _poolToken, address _wallet) returns()
func (_FeeHandler *FeeHandlerTransactor) SetTeamWallet(opts *bind.TransactOpts, _poolToken common.Address, _wallet common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setTeamWallet", _poolToken, _wallet)
}

// SetTeamWallet is a paid mutator transaction binding the contract method 0x67d3eff4.
//
// Solidity: function setTeamWallet(address _poolToken, address _wallet) returns()
func (_FeeHandler *FeeHandlerSession) SetTeamWallet(_poolToken common.Address, _wallet common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetTeamWallet(&_FeeHandler.TransactOpts, _poolToken, _wallet)
}

// SetTeamWallet is a paid mutator transaction binding the contract method 0x67d3eff4.
//
// Solidity: function setTeamWallet(address _poolToken, address _wallet) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetTeamWallet(_poolToken common.Address, _wallet common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetTeamWallet(&_FeeHandler.TransactOpts, _poolToken, _wallet)
}

// FeeHandlerClaimFeesPartnerIterator is returned from FilterClaimFeesPartner and is used to iterate over the raw logs and unpacked data for ClaimFeesPartner events raised by the FeeHandler contract.
type FeeHandlerClaimFeesPartnerIterator struct {
	Event *FeeHandlerClaimFeesPartner // Event containing the contract specifics and raw log

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
func (it *FeeHandlerClaimFeesPartnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerClaimFeesPartner)
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
		it.Event = new(FeeHandlerClaimFeesPartner)
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
func (it *FeeHandlerClaimFeesPartnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerClaimFeesPartnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerClaimFeesPartner represents a ClaimFeesPartner event raised by the FeeHandler contract.
type FeeHandlerClaimFeesPartner struct {
	PoolToken common.Address
	Token     common.Address
	Amount    *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimFeesPartner is a free log retrieval operation binding the contract event 0x34bd98703f8669398f7af8ff645a68c0f0e1207639b11aec3761895f3acb7714.
//
// Solidity: event ClaimFeesPartner(address poolToken, address token, uint256 amount, address receiver)
func (_FeeHandler *FeeHandlerFilterer) FilterClaimFeesPartner(opts *bind.FilterOpts) (*FeeHandlerClaimFeesPartnerIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "ClaimFeesPartner")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerClaimFeesPartnerIterator{contract: _FeeHandler.contract, event: "ClaimFeesPartner", logs: logs, sub: sub}, nil
}

// WatchClaimFeesPartner is a free log subscription operation binding the contract event 0x34bd98703f8669398f7af8ff645a68c0f0e1207639b11aec3761895f3acb7714.
//
// Solidity: event ClaimFeesPartner(address poolToken, address token, uint256 amount, address receiver)
func (_FeeHandler *FeeHandlerFilterer) WatchClaimFeesPartner(opts *bind.WatchOpts, sink chan<- *FeeHandlerClaimFeesPartner) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "ClaimFeesPartner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerClaimFeesPartner)
				if err := _FeeHandler.contract.UnpackLog(event, "ClaimFeesPartner", log); err != nil {
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

// ParseClaimFeesPartner is a log parse operation binding the contract event 0x34bd98703f8669398f7af8ff645a68c0f0e1207639b11aec3761895f3acb7714.
//
// Solidity: event ClaimFeesPartner(address poolToken, address token, uint256 amount, address receiver)
func (_FeeHandler *FeeHandlerFilterer) ParseClaimFeesPartner(log types.Log) (*FeeHandlerClaimFeesPartner, error) {
	event := new(FeeHandlerClaimFeesPartner)
	if err := _FeeHandler.contract.UnpackLog(event, "ClaimFeesPartner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerClaimFeesTeamIterator is returned from FilterClaimFeesTeam and is used to iterate over the raw logs and unpacked data for ClaimFeesTeam events raised by the FeeHandler contract.
type FeeHandlerClaimFeesTeamIterator struct {
	Event *FeeHandlerClaimFeesTeam // Event containing the contract specifics and raw log

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
func (it *FeeHandlerClaimFeesTeamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerClaimFeesTeam)
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
		it.Event = new(FeeHandlerClaimFeesTeam)
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
func (it *FeeHandlerClaimFeesTeamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerClaimFeesTeamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerClaimFeesTeam represents a ClaimFeesTeam event raised by the FeeHandler contract.
type FeeHandlerClaimFeesTeam struct {
	PoolToken common.Address
	Token     common.Address
	Amount    *big.Int
	Receiver  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimFeesTeam is a free log retrieval operation binding the contract event 0x967bcdcb4242f9a1ff6f2c4215c1f2037c89a73f37295249aa580785e460dbb9.
//
// Solidity: event ClaimFeesTeam(address poolToken, address token, uint256 amount, address receiver)
func (_FeeHandler *FeeHandlerFilterer) FilterClaimFeesTeam(opts *bind.FilterOpts) (*FeeHandlerClaimFeesTeamIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "ClaimFeesTeam")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerClaimFeesTeamIterator{contract: _FeeHandler.contract, event: "ClaimFeesTeam", logs: logs, sub: sub}, nil
}

// WatchClaimFeesTeam is a free log subscription operation binding the contract event 0x967bcdcb4242f9a1ff6f2c4215c1f2037c89a73f37295249aa580785e460dbb9.
//
// Solidity: event ClaimFeesTeam(address poolToken, address token, uint256 amount, address receiver)
func (_FeeHandler *FeeHandlerFilterer) WatchClaimFeesTeam(opts *bind.WatchOpts, sink chan<- *FeeHandlerClaimFeesTeam) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "ClaimFeesTeam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerClaimFeesTeam)
				if err := _FeeHandler.contract.UnpackLog(event, "ClaimFeesTeam", log); err != nil {
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

// ParseClaimFeesTeam is a log parse operation binding the contract event 0x967bcdcb4242f9a1ff6f2c4215c1f2037c89a73f37295249aa580785e460dbb9.
//
// Solidity: event ClaimFeesTeam(address poolToken, address token, uint256 amount, address receiver)
func (_FeeHandler *FeeHandlerFilterer) ParseClaimFeesTeam(log types.Log) (*FeeHandlerClaimFeesTeam, error) {
	event := new(FeeHandlerClaimFeesTeam)
	if err := _FeeHandler.contract.UnpackLog(event, "ClaimFeesTeam", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
