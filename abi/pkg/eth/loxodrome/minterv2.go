// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package loxodrome

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

// MinterV2MetaData contains all meta data concerning the MinterV2 contract.
var MinterV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__ve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__rewards_distributor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__NFTstakers_rewards\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weekly\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"circulating_supply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"circulating_emission\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_TEAM_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_Loxo\",\"outputs\":[{\"internalType\":\"contractILoxo\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_NFTstakers_rewards\",\"outputs\":[{\"internalType\":\"contractIMasterchef\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_rewards_distributor\",\"outputs\":[{\"internalType\":\"contractIRewardsDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_ve\",\"outputs\":[{\"internalType\":\"contractIVotingEscrow\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_voter\",\"outputs\":[{\"internalType\":\"contractIVoter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"active_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_weeklyMint\",\"type\":\"uint256\"}],\"name\":\"calculate_NFT_Staker_Reward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculate_emission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_weeklyMint\",\"type\":\"uint256\"}],\"name\":\"calculate_rebase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"check\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circulating_emission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"circulating_supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"claimants\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFirstMint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingTeam\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_emission\",\"type\":\"uint256\"}],\"name\":\"setEmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_NFTStakersReward\",\"type\":\"uint256\"}],\"name\":\"setNFTStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__NFTstakers\",\"type\":\"address\"}],\"name\":\"setNFTstakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rebase\",\"type\":\"uint256\"}],\"name\":\"setRebase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_team\",\"type\":\"address\"}],\"name\":\"setTeam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_teamRate\",\"type\":\"uint256\"}],\"name\":\"setTeamRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__voter\",\"type\":\"address\"}],\"name\":\"setVoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"team\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teamRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update_period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weekly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weekly_emission\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MinterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MinterV2MetaData.ABI instead.
var MinterV2ABI = MinterV2MetaData.ABI

// MinterV2 is an auto generated Go binding around an Ethereum contract.
type MinterV2 struct {
	MinterV2Caller     // Read-only binding to the contract
	MinterV2Transactor // Write-only binding to the contract
	MinterV2Filterer   // Log filterer for contract events
}

// MinterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MinterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterV2Session struct {
	Contract     *MinterV2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterV2CallerSession struct {
	Contract *MinterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MinterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterV2TransactorSession struct {
	Contract     *MinterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MinterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MinterV2Raw struct {
	Contract *MinterV2 // Generic contract binding to access the raw methods on
}

// MinterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterV2CallerRaw struct {
	Contract *MinterV2Caller // Generic read-only contract binding to access the raw methods on
}

// MinterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterV2TransactorRaw struct {
	Contract *MinterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterV2 creates a new instance of MinterV2, bound to a specific deployed contract.
func NewMinterV2(address common.Address, backend bind.ContractBackend) (*MinterV2, error) {
	contract, err := bindMinterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterV2{MinterV2Caller: MinterV2Caller{contract: contract}, MinterV2Transactor: MinterV2Transactor{contract: contract}, MinterV2Filterer: MinterV2Filterer{contract: contract}}, nil
}

// NewMinterV2Caller creates a new read-only instance of MinterV2, bound to a specific deployed contract.
func NewMinterV2Caller(address common.Address, caller bind.ContractCaller) (*MinterV2Caller, error) {
	contract, err := bindMinterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterV2Caller{contract: contract}, nil
}

// NewMinterV2Transactor creates a new write-only instance of MinterV2, bound to a specific deployed contract.
func NewMinterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MinterV2Transactor, error) {
	contract, err := bindMinterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterV2Transactor{contract: contract}, nil
}

// NewMinterV2Filterer creates a new log filterer instance of MinterV2, bound to a specific deployed contract.
func NewMinterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MinterV2Filterer, error) {
	contract, err := bindMinterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterV2Filterer{contract: contract}, nil
}

// bindMinterV2 binds a generic wrapper to an already deployed contract.
func bindMinterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MinterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterV2 *MinterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterV2.Contract.MinterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterV2 *MinterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterV2.Contract.MinterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterV2 *MinterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterV2.Contract.MinterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterV2 *MinterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterV2 *MinterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterV2 *MinterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterV2.Contract.contract.Transact(opts, method, params...)
}

// MAXTEAMRATE is a free data retrieval call binding the contract method 0x01c8e6fd.
//
// Solidity: function MAX_TEAM_RATE() view returns(uint256)
func (_MinterV2 *MinterV2Caller) MAXTEAMRATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "MAX_TEAM_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTEAMRATE is a free data retrieval call binding the contract method 0x01c8e6fd.
//
// Solidity: function MAX_TEAM_RATE() view returns(uint256)
func (_MinterV2 *MinterV2Session) MAXTEAMRATE() (*big.Int, error) {
	return _MinterV2.Contract.MAXTEAMRATE(&_MinterV2.CallOpts)
}

// MAXTEAMRATE is a free data retrieval call binding the contract method 0x01c8e6fd.
//
// Solidity: function MAX_TEAM_RATE() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) MAXTEAMRATE() (*big.Int, error) {
	return _MinterV2.Contract.MAXTEAMRATE(&_MinterV2.CallOpts)
}

// Loxo is a free data retrieval call binding the contract method 0x33b875b2.
//
// Solidity: function _Loxo() view returns(address)
func (_MinterV2 *MinterV2Caller) Loxo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "_Loxo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Loxo is a free data retrieval call binding the contract method 0x33b875b2.
//
// Solidity: function _Loxo() view returns(address)
func (_MinterV2 *MinterV2Session) Loxo() (common.Address, error) {
	return _MinterV2.Contract.Loxo(&_MinterV2.CallOpts)
}

// Loxo is a free data retrieval call binding the contract method 0x33b875b2.
//
// Solidity: function _Loxo() view returns(address)
func (_MinterV2 *MinterV2CallerSession) Loxo() (common.Address, error) {
	return _MinterV2.Contract.Loxo(&_MinterV2.CallOpts)
}

// NFTstakersRewards is a free data retrieval call binding the contract method 0x4718d877.
//
// Solidity: function _NFTstakers_rewards() view returns(address)
func (_MinterV2 *MinterV2Caller) NFTstakersRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "_NFTstakers_rewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NFTstakersRewards is a free data retrieval call binding the contract method 0x4718d877.
//
// Solidity: function _NFTstakers_rewards() view returns(address)
func (_MinterV2 *MinterV2Session) NFTstakersRewards() (common.Address, error) {
	return _MinterV2.Contract.NFTstakersRewards(&_MinterV2.CallOpts)
}

// NFTstakersRewards is a free data retrieval call binding the contract method 0x4718d877.
//
// Solidity: function _NFTstakers_rewards() view returns(address)
func (_MinterV2 *MinterV2CallerSession) NFTstakersRewards() (common.Address, error) {
	return _MinterV2.Contract.NFTstakersRewards(&_MinterV2.CallOpts)
}

// RewardsDistributor is a free data retrieval call binding the contract method 0x4b1cd5da.
//
// Solidity: function _rewards_distributor() view returns(address)
func (_MinterV2 *MinterV2Caller) RewardsDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "_rewards_distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsDistributor is a free data retrieval call binding the contract method 0x4b1cd5da.
//
// Solidity: function _rewards_distributor() view returns(address)
func (_MinterV2 *MinterV2Session) RewardsDistributor() (common.Address, error) {
	return _MinterV2.Contract.RewardsDistributor(&_MinterV2.CallOpts)
}

// RewardsDistributor is a free data retrieval call binding the contract method 0x4b1cd5da.
//
// Solidity: function _rewards_distributor() view returns(address)
func (_MinterV2 *MinterV2CallerSession) RewardsDistributor() (common.Address, error) {
	return _MinterV2.Contract.RewardsDistributor(&_MinterV2.CallOpts)
}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_MinterV2 *MinterV2Caller) Ve(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "_ve")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_MinterV2 *MinterV2Session) Ve() (common.Address, error) {
	return _MinterV2.Contract.Ve(&_MinterV2.CallOpts)
}

// Ve is a free data retrieval call binding the contract method 0x8dd598fb.
//
// Solidity: function _ve() view returns(address)
func (_MinterV2 *MinterV2CallerSession) Ve() (common.Address, error) {
	return _MinterV2.Contract.Ve(&_MinterV2.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x3db9b42a.
//
// Solidity: function _voter() view returns(address)
func (_MinterV2 *MinterV2Caller) Voter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "_voter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voter is a free data retrieval call binding the contract method 0x3db9b42a.
//
// Solidity: function _voter() view returns(address)
func (_MinterV2 *MinterV2Session) Voter() (common.Address, error) {
	return _MinterV2.Contract.Voter(&_MinterV2.CallOpts)
}

// Voter is a free data retrieval call binding the contract method 0x3db9b42a.
//
// Solidity: function _voter() view returns(address)
func (_MinterV2 *MinterV2CallerSession) Voter() (common.Address, error) {
	return _MinterV2.Contract.Voter(&_MinterV2.CallOpts)
}

// ActivePeriod is a free data retrieval call binding the contract method 0xd1399608.
//
// Solidity: function active_period() view returns(uint256)
func (_MinterV2 *MinterV2Caller) ActivePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "active_period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActivePeriod is a free data retrieval call binding the contract method 0xd1399608.
//
// Solidity: function active_period() view returns(uint256)
func (_MinterV2 *MinterV2Session) ActivePeriod() (*big.Int, error) {
	return _MinterV2.Contract.ActivePeriod(&_MinterV2.CallOpts)
}

// ActivePeriod is a free data retrieval call binding the contract method 0xd1399608.
//
// Solidity: function active_period() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) ActivePeriod() (*big.Int, error) {
	return _MinterV2.Contract.ActivePeriod(&_MinterV2.CallOpts)
}

// CalculateNFTStakerReward is a free data retrieval call binding the contract method 0xa6e46d98.
//
// Solidity: function calculate_NFT_Staker_Reward(uint256 _weeklyMint) view returns(uint256)
func (_MinterV2 *MinterV2Caller) CalculateNFTStakerReward(opts *bind.CallOpts, _weeklyMint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "calculate_NFT_Staker_Reward", _weeklyMint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateNFTStakerReward is a free data retrieval call binding the contract method 0xa6e46d98.
//
// Solidity: function calculate_NFT_Staker_Reward(uint256 _weeklyMint) view returns(uint256)
func (_MinterV2 *MinterV2Session) CalculateNFTStakerReward(_weeklyMint *big.Int) (*big.Int, error) {
	return _MinterV2.Contract.CalculateNFTStakerReward(&_MinterV2.CallOpts, _weeklyMint)
}

// CalculateNFTStakerReward is a free data retrieval call binding the contract method 0xa6e46d98.
//
// Solidity: function calculate_NFT_Staker_Reward(uint256 _weeklyMint) view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) CalculateNFTStakerReward(_weeklyMint *big.Int) (*big.Int, error) {
	return _MinterV2.Contract.CalculateNFTStakerReward(&_MinterV2.CallOpts, _weeklyMint)
}

// CalculateEmission is a free data retrieval call binding the contract method 0x36d96faf.
//
// Solidity: function calculate_emission() view returns(uint256)
func (_MinterV2 *MinterV2Caller) CalculateEmission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "calculate_emission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateEmission is a free data retrieval call binding the contract method 0x36d96faf.
//
// Solidity: function calculate_emission() view returns(uint256)
func (_MinterV2 *MinterV2Session) CalculateEmission() (*big.Int, error) {
	return _MinterV2.Contract.CalculateEmission(&_MinterV2.CallOpts)
}

// CalculateEmission is a free data retrieval call binding the contract method 0x36d96faf.
//
// Solidity: function calculate_emission() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) CalculateEmission() (*big.Int, error) {
	return _MinterV2.Contract.CalculateEmission(&_MinterV2.CallOpts)
}

// CalculateRebase is a free data retrieval call binding the contract method 0xa129cfbf.
//
// Solidity: function calculate_rebase(uint256 _weeklyMint) view returns(uint256)
func (_MinterV2 *MinterV2Caller) CalculateRebase(opts *bind.CallOpts, _weeklyMint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "calculate_rebase", _weeklyMint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRebase is a free data retrieval call binding the contract method 0xa129cfbf.
//
// Solidity: function calculate_rebase(uint256 _weeklyMint) view returns(uint256)
func (_MinterV2 *MinterV2Session) CalculateRebase(_weeklyMint *big.Int) (*big.Int, error) {
	return _MinterV2.Contract.CalculateRebase(&_MinterV2.CallOpts, _weeklyMint)
}

// CalculateRebase is a free data retrieval call binding the contract method 0xa129cfbf.
//
// Solidity: function calculate_rebase(uint256 _weeklyMint) view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) CalculateRebase(_weeklyMint *big.Int) (*big.Int, error) {
	return _MinterV2.Contract.CalculateRebase(&_MinterV2.CallOpts, _weeklyMint)
}

// Check is a free data retrieval call binding the contract method 0x919840ad.
//
// Solidity: function check() view returns(bool)
func (_MinterV2 *MinterV2Caller) Check(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "check")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Check is a free data retrieval call binding the contract method 0x919840ad.
//
// Solidity: function check() view returns(bool)
func (_MinterV2 *MinterV2Session) Check() (bool, error) {
	return _MinterV2.Contract.Check(&_MinterV2.CallOpts)
}

// Check is a free data retrieval call binding the contract method 0x919840ad.
//
// Solidity: function check() view returns(bool)
func (_MinterV2 *MinterV2CallerSession) Check() (bool, error) {
	return _MinterV2.Contract.Check(&_MinterV2.CallOpts)
}

// CirculatingEmission is a free data retrieval call binding the contract method 0x1eebae80.
//
// Solidity: function circulating_emission() view returns(uint256)
func (_MinterV2 *MinterV2Caller) CirculatingEmission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "circulating_emission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CirculatingEmission is a free data retrieval call binding the contract method 0x1eebae80.
//
// Solidity: function circulating_emission() view returns(uint256)
func (_MinterV2 *MinterV2Session) CirculatingEmission() (*big.Int, error) {
	return _MinterV2.Contract.CirculatingEmission(&_MinterV2.CallOpts)
}

// CirculatingEmission is a free data retrieval call binding the contract method 0x1eebae80.
//
// Solidity: function circulating_emission() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) CirculatingEmission() (*big.Int, error) {
	return _MinterV2.Contract.CirculatingEmission(&_MinterV2.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0xe038c75a.
//
// Solidity: function circulating_supply() view returns(uint256)
func (_MinterV2 *MinterV2Caller) CirculatingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "circulating_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CirculatingSupply is a free data retrieval call binding the contract method 0xe038c75a.
//
// Solidity: function circulating_supply() view returns(uint256)
func (_MinterV2 *MinterV2Session) CirculatingSupply() (*big.Int, error) {
	return _MinterV2.Contract.CirculatingSupply(&_MinterV2.CallOpts)
}

// CirculatingSupply is a free data retrieval call binding the contract method 0xe038c75a.
//
// Solidity: function circulating_supply() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) CirculatingSupply() (*big.Int, error) {
	return _MinterV2.Contract.CirculatingSupply(&_MinterV2.CallOpts)
}

// IsFirstMint is a free data retrieval call binding the contract method 0xa9abf7ec.
//
// Solidity: function isFirstMint() view returns(bool)
func (_MinterV2 *MinterV2Caller) IsFirstMint(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "isFirstMint")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFirstMint is a free data retrieval call binding the contract method 0xa9abf7ec.
//
// Solidity: function isFirstMint() view returns(bool)
func (_MinterV2 *MinterV2Session) IsFirstMint() (bool, error) {
	return _MinterV2.Contract.IsFirstMint(&_MinterV2.CallOpts)
}

// IsFirstMint is a free data retrieval call binding the contract method 0xa9abf7ec.
//
// Solidity: function isFirstMint() view returns(bool)
func (_MinterV2 *MinterV2CallerSession) IsFirstMint() (bool, error) {
	return _MinterV2.Contract.IsFirstMint(&_MinterV2.CallOpts)
}

// PendingTeam is a free data retrieval call binding the contract method 0x59d46ffc.
//
// Solidity: function pendingTeam() view returns(address)
func (_MinterV2 *MinterV2Caller) PendingTeam(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "pendingTeam")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingTeam is a free data retrieval call binding the contract method 0x59d46ffc.
//
// Solidity: function pendingTeam() view returns(address)
func (_MinterV2 *MinterV2Session) PendingTeam() (common.Address, error) {
	return _MinterV2.Contract.PendingTeam(&_MinterV2.CallOpts)
}

// PendingTeam is a free data retrieval call binding the contract method 0x59d46ffc.
//
// Solidity: function pendingTeam() view returns(address)
func (_MinterV2 *MinterV2CallerSession) PendingTeam() (common.Address, error) {
	return _MinterV2.Contract.PendingTeam(&_MinterV2.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint256)
func (_MinterV2 *MinterV2Caller) Period(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint256)
func (_MinterV2 *MinterV2Session) Period() (*big.Int, error) {
	return _MinterV2.Contract.Period(&_MinterV2.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) Period() (*big.Int, error) {
	return _MinterV2.Contract.Period(&_MinterV2.CallOpts)
}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_MinterV2 *MinterV2Caller) Team(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "team")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_MinterV2 *MinterV2Session) Team() (common.Address, error) {
	return _MinterV2.Contract.Team(&_MinterV2.CallOpts)
}

// Team is a free data retrieval call binding the contract method 0x85f2aef2.
//
// Solidity: function team() view returns(address)
func (_MinterV2 *MinterV2CallerSession) Team() (common.Address, error) {
	return _MinterV2.Contract.Team(&_MinterV2.CallOpts)
}

// TeamRate is a free data retrieval call binding the contract method 0x78ef7f02.
//
// Solidity: function teamRate() view returns(uint256)
func (_MinterV2 *MinterV2Caller) TeamRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "teamRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TeamRate is a free data retrieval call binding the contract method 0x78ef7f02.
//
// Solidity: function teamRate() view returns(uint256)
func (_MinterV2 *MinterV2Session) TeamRate() (*big.Int, error) {
	return _MinterV2.Contract.TeamRate(&_MinterV2.CallOpts)
}

// TeamRate is a free data retrieval call binding the contract method 0x78ef7f02.
//
// Solidity: function teamRate() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) TeamRate() (*big.Int, error) {
	return _MinterV2.Contract.TeamRate(&_MinterV2.CallOpts)
}

// Weekly is a free data retrieval call binding the contract method 0x26cfc17b.
//
// Solidity: function weekly() view returns(uint256)
func (_MinterV2 *MinterV2Caller) Weekly(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "weekly")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Weekly is a free data retrieval call binding the contract method 0x26cfc17b.
//
// Solidity: function weekly() view returns(uint256)
func (_MinterV2 *MinterV2Session) Weekly() (*big.Int, error) {
	return _MinterV2.Contract.Weekly(&_MinterV2.CallOpts)
}

// Weekly is a free data retrieval call binding the contract method 0x26cfc17b.
//
// Solidity: function weekly() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) Weekly() (*big.Int, error) {
	return _MinterV2.Contract.Weekly(&_MinterV2.CallOpts)
}

// WeeklyEmission is a free data retrieval call binding the contract method 0xcfc6c8ff.
//
// Solidity: function weekly_emission() view returns(uint256)
func (_MinterV2 *MinterV2Caller) WeeklyEmission(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MinterV2.contract.Call(opts, &out, "weekly_emission")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeeklyEmission is a free data retrieval call binding the contract method 0xcfc6c8ff.
//
// Solidity: function weekly_emission() view returns(uint256)
func (_MinterV2 *MinterV2Session) WeeklyEmission() (*big.Int, error) {
	return _MinterV2.Contract.WeeklyEmission(&_MinterV2.CallOpts)
}

// WeeklyEmission is a free data retrieval call binding the contract method 0xcfc6c8ff.
//
// Solidity: function weekly_emission() view returns(uint256)
func (_MinterV2 *MinterV2CallerSession) WeeklyEmission() (*big.Int, error) {
	return _MinterV2.Contract.WeeklyEmission(&_MinterV2.CallOpts)
}

// AcceptTeam is a paid mutator transaction binding the contract method 0xb5cc143a.
//
// Solidity: function acceptTeam() returns()
func (_MinterV2 *MinterV2Transactor) AcceptTeam(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "acceptTeam")
}

// AcceptTeam is a paid mutator transaction binding the contract method 0xb5cc143a.
//
// Solidity: function acceptTeam() returns()
func (_MinterV2 *MinterV2Session) AcceptTeam() (*types.Transaction, error) {
	return _MinterV2.Contract.AcceptTeam(&_MinterV2.TransactOpts)
}

// AcceptTeam is a paid mutator transaction binding the contract method 0xb5cc143a.
//
// Solidity: function acceptTeam() returns()
func (_MinterV2 *MinterV2TransactorSession) AcceptTeam() (*types.Transaction, error) {
	return _MinterV2.Contract.AcceptTeam(&_MinterV2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x85170813.
//
// Solidity: function initialize(address[] claimants, uint256[] amounts, uint256 max) returns()
func (_MinterV2 *MinterV2Transactor) Initialize(opts *bind.TransactOpts, claimants []common.Address, amounts []*big.Int, max *big.Int) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "initialize", claimants, amounts, max)
}

// Initialize is a paid mutator transaction binding the contract method 0x85170813.
//
// Solidity: function initialize(address[] claimants, uint256[] amounts, uint256 max) returns()
func (_MinterV2 *MinterV2Session) Initialize(claimants []common.Address, amounts []*big.Int, max *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.Initialize(&_MinterV2.TransactOpts, claimants, amounts, max)
}

// Initialize is a paid mutator transaction binding the contract method 0x85170813.
//
// Solidity: function initialize(address[] claimants, uint256[] amounts, uint256 max) returns()
func (_MinterV2 *MinterV2TransactorSession) Initialize(claimants []common.Address, amounts []*big.Int, max *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.Initialize(&_MinterV2.TransactOpts, claimants, amounts, max)
}

// SetEmission is a paid mutator transaction binding the contract method 0xddce102f.
//
// Solidity: function setEmission(uint256 _emission) returns()
func (_MinterV2 *MinterV2Transactor) SetEmission(opts *bind.TransactOpts, _emission *big.Int) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "setEmission", _emission)
}

// SetEmission is a paid mutator transaction binding the contract method 0xddce102f.
//
// Solidity: function setEmission(uint256 _emission) returns()
func (_MinterV2 *MinterV2Session) SetEmission(_emission *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetEmission(&_MinterV2.TransactOpts, _emission)
}

// SetEmission is a paid mutator transaction binding the contract method 0xddce102f.
//
// Solidity: function setEmission(uint256 _emission) returns()
func (_MinterV2 *MinterV2TransactorSession) SetEmission(_emission *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetEmission(&_MinterV2.TransactOpts, _emission)
}

// SetNFTStaker is a paid mutator transaction binding the contract method 0xfde7ef4f.
//
// Solidity: function setNFTStaker(uint256 _NFTStakersReward) returns()
func (_MinterV2 *MinterV2Transactor) SetNFTStaker(opts *bind.TransactOpts, _NFTStakersReward *big.Int) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "setNFTStaker", _NFTStakersReward)
}

// SetNFTStaker is a paid mutator transaction binding the contract method 0xfde7ef4f.
//
// Solidity: function setNFTStaker(uint256 _NFTStakersReward) returns()
func (_MinterV2 *MinterV2Session) SetNFTStaker(_NFTStakersReward *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetNFTStaker(&_MinterV2.TransactOpts, _NFTStakersReward)
}

// SetNFTStaker is a paid mutator transaction binding the contract method 0xfde7ef4f.
//
// Solidity: function setNFTStaker(uint256 _NFTStakersReward) returns()
func (_MinterV2 *MinterV2TransactorSession) SetNFTStaker(_NFTStakersReward *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetNFTStaker(&_MinterV2.TransactOpts, _NFTStakersReward)
}

// SetNFTstakers is a paid mutator transaction binding the contract method 0xd248687f.
//
// Solidity: function setNFTstakers(address __NFTstakers) returns()
func (_MinterV2 *MinterV2Transactor) SetNFTstakers(opts *bind.TransactOpts, __NFTstakers common.Address) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "setNFTstakers", __NFTstakers)
}

// SetNFTstakers is a paid mutator transaction binding the contract method 0xd248687f.
//
// Solidity: function setNFTstakers(address __NFTstakers) returns()
func (_MinterV2 *MinterV2Session) SetNFTstakers(__NFTstakers common.Address) (*types.Transaction, error) {
	return _MinterV2.Contract.SetNFTstakers(&_MinterV2.TransactOpts, __NFTstakers)
}

// SetNFTstakers is a paid mutator transaction binding the contract method 0xd248687f.
//
// Solidity: function setNFTstakers(address __NFTstakers) returns()
func (_MinterV2 *MinterV2TransactorSession) SetNFTstakers(__NFTstakers common.Address) (*types.Transaction, error) {
	return _MinterV2.Contract.SetNFTstakers(&_MinterV2.TransactOpts, __NFTstakers)
}

// SetRebase is a paid mutator transaction binding the contract method 0xde3e3492.
//
// Solidity: function setRebase(uint256 _rebase) returns()
func (_MinterV2 *MinterV2Transactor) SetRebase(opts *bind.TransactOpts, _rebase *big.Int) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "setRebase", _rebase)
}

// SetRebase is a paid mutator transaction binding the contract method 0xde3e3492.
//
// Solidity: function setRebase(uint256 _rebase) returns()
func (_MinterV2 *MinterV2Session) SetRebase(_rebase *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetRebase(&_MinterV2.TransactOpts, _rebase)
}

// SetRebase is a paid mutator transaction binding the contract method 0xde3e3492.
//
// Solidity: function setRebase(uint256 _rebase) returns()
func (_MinterV2 *MinterV2TransactorSession) SetRebase(_rebase *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetRebase(&_MinterV2.TransactOpts, _rebase)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_MinterV2 *MinterV2Transactor) SetTeam(opts *bind.TransactOpts, _team common.Address) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "setTeam", _team)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_MinterV2 *MinterV2Session) SetTeam(_team common.Address) (*types.Transaction, error) {
	return _MinterV2.Contract.SetTeam(&_MinterV2.TransactOpts, _team)
}

// SetTeam is a paid mutator transaction binding the contract method 0x095cf5c6.
//
// Solidity: function setTeam(address _team) returns()
func (_MinterV2 *MinterV2TransactorSession) SetTeam(_team common.Address) (*types.Transaction, error) {
	return _MinterV2.Contract.SetTeam(&_MinterV2.TransactOpts, _team)
}

// SetTeamRate is a paid mutator transaction binding the contract method 0x2e8f7b1f.
//
// Solidity: function setTeamRate(uint256 _teamRate) returns()
func (_MinterV2 *MinterV2Transactor) SetTeamRate(opts *bind.TransactOpts, _teamRate *big.Int) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "setTeamRate", _teamRate)
}

// SetTeamRate is a paid mutator transaction binding the contract method 0x2e8f7b1f.
//
// Solidity: function setTeamRate(uint256 _teamRate) returns()
func (_MinterV2 *MinterV2Session) SetTeamRate(_teamRate *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetTeamRate(&_MinterV2.TransactOpts, _teamRate)
}

// SetTeamRate is a paid mutator transaction binding the contract method 0x2e8f7b1f.
//
// Solidity: function setTeamRate(uint256 _teamRate) returns()
func (_MinterV2 *MinterV2TransactorSession) SetTeamRate(_teamRate *big.Int) (*types.Transaction, error) {
	return _MinterV2.Contract.SetTeamRate(&_MinterV2.TransactOpts, _teamRate)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address __voter) returns()
func (_MinterV2 *MinterV2Transactor) SetVoter(opts *bind.TransactOpts, __voter common.Address) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "setVoter", __voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address __voter) returns()
func (_MinterV2 *MinterV2Session) SetVoter(__voter common.Address) (*types.Transaction, error) {
	return _MinterV2.Contract.SetVoter(&_MinterV2.TransactOpts, __voter)
}

// SetVoter is a paid mutator transaction binding the contract method 0x4bc2a657.
//
// Solidity: function setVoter(address __voter) returns()
func (_MinterV2 *MinterV2TransactorSession) SetVoter(__voter common.Address) (*types.Transaction, error) {
	return _MinterV2.Contract.SetVoter(&_MinterV2.TransactOpts, __voter)
}

// UpdatePeriod is a paid mutator transaction binding the contract method 0xed29fc11.
//
// Solidity: function update_period() returns(uint256)
func (_MinterV2 *MinterV2Transactor) UpdatePeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterV2.contract.Transact(opts, "update_period")
}

// UpdatePeriod is a paid mutator transaction binding the contract method 0xed29fc11.
//
// Solidity: function update_period() returns(uint256)
func (_MinterV2 *MinterV2Session) UpdatePeriod() (*types.Transaction, error) {
	return _MinterV2.Contract.UpdatePeriod(&_MinterV2.TransactOpts)
}

// UpdatePeriod is a paid mutator transaction binding the contract method 0xed29fc11.
//
// Solidity: function update_period() returns(uint256)
func (_MinterV2 *MinterV2TransactorSession) UpdatePeriod() (*types.Transaction, error) {
	return _MinterV2.Contract.UpdatePeriod(&_MinterV2.TransactOpts)
}

// MinterV2MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the MinterV2 contract.
type MinterV2MintIterator struct {
	Event *MinterV2Mint // Event containing the contract specifics and raw log

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
func (it *MinterV2MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterV2Mint)
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
		it.Event = new(MinterV2Mint)
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
func (it *MinterV2MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterV2MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterV2Mint represents a Mint event raised by the MinterV2 contract.
type MinterV2Mint struct {
	Sender              common.Address
	Weekly              *big.Int
	CirculatingSupply   *big.Int
	CirculatingEmission *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed sender, uint256 weekly, uint256 circulating_supply, uint256 circulating_emission)
func (_MinterV2 *MinterV2Filterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*MinterV2MintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MinterV2.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &MinterV2MintIterator{contract: _MinterV2.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed sender, uint256 weekly, uint256 circulating_supply, uint256 circulating_emission)
func (_MinterV2 *MinterV2Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *MinterV2Mint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MinterV2.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterV2Mint)
				if err := _MinterV2.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed sender, uint256 weekly, uint256 circulating_supply, uint256 circulating_emission)
func (_MinterV2 *MinterV2Filterer) ParseMint(log types.Log) (*MinterV2Mint, error) {
	event := new(MinterV2Mint)
	if err := _MinterV2.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
