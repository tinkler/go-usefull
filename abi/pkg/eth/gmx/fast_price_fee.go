// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gmx

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

// FastPriceFeeMetaData contains all meta data concerning the FastPriceFee contract.
var FastPriceFeeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxPriceUpdateDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minBlockInterval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDeviationBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_fastPriceEvents\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"DisableFastPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"EnableFastPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeRefDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFastDelta\",\"type\":\"uint256\"}],\"name\":\"MaxCumulativeDeltaDiffExceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fastPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeRefDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cumulativeFastDelta\",\"type\":\"uint256\"}],\"name\":\"PriceData\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BASIS_POINTS_DIVISOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITMASK_32\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CUMULATIVE_DELTA_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CUMULATIVE_FAST_DELTA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CUMULATIVE_REF_DELTA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PRICE_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REF_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRICE_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableFastPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableFastPriceVoteCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"disableFastPriceVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableFastPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastPriceEvents\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"favorFastPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_refPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_maximise\",\"type\":\"bool\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gov\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAuthorizations\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_updaters\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSpreadEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isUpdater\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxCumulativeDeltaDiffs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDeviationBasisPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPriceUpdateDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTimeDeviation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minAuthorizations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBlockInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"priceData\",\"outputs\":[{\"internalType\":\"uint160\",\"name\":\"refPrice\",\"type\":\"uint160\"},{\"internalType\":\"uint32\",\"name\":\"refTime\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"cumulativeRefDelta\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"cumulativeFastDelta\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDataInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_priceBitArray\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"setCompactedPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_fastPriceEvents\",\"type\":\"address\"}],\"name\":\"setFastPriceEvents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gov\",\"type\":\"address\"}],\"name\":\"setGov\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSpreadEnabled\",\"type\":\"bool\"}],\"name\":\"setIsSpreadEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastUpdatedAt\",\"type\":\"uint256\"}],\"name\":\"setLastUpdatedAt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_maxCumulativeDeltaDiffs\",\"type\":\"uint256[]\"}],\"name\":\"setMaxCumulativeDeltaDiffs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDeviationBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setMaxDeviationBasisPoints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPriceUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"setMaxPriceUpdateDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTimeDeviation\",\"type\":\"uint256\"}],\"name\":\"setMaxTimeDeviation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minAuthorizations\",\"type\":\"uint256\"}],\"name\":\"setMinAuthorizations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minBlockInterval\",\"type\":\"uint256\"}],\"name\":\"setMinBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceDataInterval\",\"type\":\"uint256\"}],\"name\":\"setPriceDataInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceDuration\",\"type\":\"uint256\"}],\"name\":\"setPriceDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"setPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_priceBits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"setPricesWithBits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_positionRouter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceBits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndexForIncreasePositions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endIndexForDecreasePositions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxIncreasePositions\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDecreasePositions\",\"type\":\"uint256\"}],\"name\":\"setPricesWithBitsAndExecute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfChainError\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfChainError\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spreadBasisPointsIfInactive\",\"type\":\"uint256\"}],\"name\":\"setSpreadBasisPointsIfInactive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenManager\",\"type\":\"address\"}],\"name\":\"setTokenManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenPrecisions\",\"type\":\"uint256[]\"}],\"name\":\"setTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isActive\",\"type\":\"bool\"}],\"name\":\"setUpdater\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultPriceFeed\",\"type\":\"address\"}],\"name\":\"setVaultPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadBasisPointsIfChainError\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spreadBasisPointsIfInactive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenPrecisions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultPriceFeed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FastPriceFeeABI is the input ABI used to generate the binding from.
// Deprecated: Use FastPriceFeeMetaData.ABI instead.
var FastPriceFeeABI = FastPriceFeeMetaData.ABI

// FastPriceFee is an auto generated Go binding around an Ethereum contract.
type FastPriceFee struct {
	FastPriceFeeCaller     // Read-only binding to the contract
	FastPriceFeeTransactor // Write-only binding to the contract
	FastPriceFeeFilterer   // Log filterer for contract events
}

// FastPriceFeeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastPriceFeeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceFeeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastPriceFeeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceFeeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastPriceFeeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastPriceFeeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastPriceFeeSession struct {
	Contract     *FastPriceFee     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FastPriceFeeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastPriceFeeCallerSession struct {
	Contract *FastPriceFeeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FastPriceFeeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastPriceFeeTransactorSession struct {
	Contract     *FastPriceFeeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FastPriceFeeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastPriceFeeRaw struct {
	Contract *FastPriceFee // Generic contract binding to access the raw methods on
}

// FastPriceFeeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastPriceFeeCallerRaw struct {
	Contract *FastPriceFeeCaller // Generic read-only contract binding to access the raw methods on
}

// FastPriceFeeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastPriceFeeTransactorRaw struct {
	Contract *FastPriceFeeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastPriceFee creates a new instance of FastPriceFee, bound to a specific deployed contract.
func NewFastPriceFee(address common.Address, backend bind.ContractBackend) (*FastPriceFee, error) {
	contract, err := bindFastPriceFee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastPriceFee{FastPriceFeeCaller: FastPriceFeeCaller{contract: contract}, FastPriceFeeTransactor: FastPriceFeeTransactor{contract: contract}, FastPriceFeeFilterer: FastPriceFeeFilterer{contract: contract}}, nil
}

// NewFastPriceFeeCaller creates a new read-only instance of FastPriceFee, bound to a specific deployed contract.
func NewFastPriceFeeCaller(address common.Address, caller bind.ContractCaller) (*FastPriceFeeCaller, error) {
	contract, err := bindFastPriceFee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastPriceFeeCaller{contract: contract}, nil
}

// NewFastPriceFeeTransactor creates a new write-only instance of FastPriceFee, bound to a specific deployed contract.
func NewFastPriceFeeTransactor(address common.Address, transactor bind.ContractTransactor) (*FastPriceFeeTransactor, error) {
	contract, err := bindFastPriceFee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastPriceFeeTransactor{contract: contract}, nil
}

// NewFastPriceFeeFilterer creates a new log filterer instance of FastPriceFee, bound to a specific deployed contract.
func NewFastPriceFeeFilterer(address common.Address, filterer bind.ContractFilterer) (*FastPriceFeeFilterer, error) {
	contract, err := bindFastPriceFee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastPriceFeeFilterer{contract: contract}, nil
}

// bindFastPriceFee binds a generic wrapper to an already deployed contract.
func bindFastPriceFee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastPriceFeeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastPriceFee *FastPriceFeeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastPriceFee.Contract.FastPriceFeeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastPriceFee *FastPriceFeeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFee.Contract.FastPriceFeeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastPriceFee *FastPriceFeeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastPriceFee.Contract.FastPriceFeeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastPriceFee *FastPriceFeeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastPriceFee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastPriceFee *FastPriceFeeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastPriceFee *FastPriceFeeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastPriceFee.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _FastPriceFee.Contract.BASISPOINTSDIVISOR(&_FastPriceFee.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _FastPriceFee.Contract.BASISPOINTSDIVISOR(&_FastPriceFee.CallOpts)
}

// BITMASK32 is a free data retrieval call binding the contract method 0x807c9782.
//
// Solidity: function BITMASK_32() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) BITMASK32(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "BITMASK_32")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITMASK32 is a free data retrieval call binding the contract method 0x807c9782.
//
// Solidity: function BITMASK_32() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) BITMASK32() (*big.Int, error) {
	return _FastPriceFee.Contract.BITMASK32(&_FastPriceFee.CallOpts)
}

// BITMASK32 is a free data retrieval call binding the contract method 0x807c9782.
//
// Solidity: function BITMASK_32() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) BITMASK32() (*big.Int, error) {
	return _FastPriceFee.Contract.BITMASK32(&_FastPriceFee.CallOpts)
}

// CUMULATIVEDELTAPRECISION is a free data retrieval call binding the contract method 0xa2b47c16.
//
// Solidity: function CUMULATIVE_DELTA_PRECISION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) CUMULATIVEDELTAPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "CUMULATIVE_DELTA_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CUMULATIVEDELTAPRECISION is a free data retrieval call binding the contract method 0xa2b47c16.
//
// Solidity: function CUMULATIVE_DELTA_PRECISION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) CUMULATIVEDELTAPRECISION() (*big.Int, error) {
	return _FastPriceFee.Contract.CUMULATIVEDELTAPRECISION(&_FastPriceFee.CallOpts)
}

// CUMULATIVEDELTAPRECISION is a free data retrieval call binding the contract method 0xa2b47c16.
//
// Solidity: function CUMULATIVE_DELTA_PRECISION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) CUMULATIVEDELTAPRECISION() (*big.Int, error) {
	return _FastPriceFee.Contract.CUMULATIVEDELTAPRECISION(&_FastPriceFee.CallOpts)
}

// MAXCUMULATIVEFASTDELTA is a free data retrieval call binding the contract method 0x4bd66c1c.
//
// Solidity: function MAX_CUMULATIVE_FAST_DELTA() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MAXCUMULATIVEFASTDELTA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "MAX_CUMULATIVE_FAST_DELTA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCUMULATIVEFASTDELTA is a free data retrieval call binding the contract method 0x4bd66c1c.
//
// Solidity: function MAX_CUMULATIVE_FAST_DELTA() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MAXCUMULATIVEFASTDELTA() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXCUMULATIVEFASTDELTA(&_FastPriceFee.CallOpts)
}

// MAXCUMULATIVEFASTDELTA is a free data retrieval call binding the contract method 0x4bd66c1c.
//
// Solidity: function MAX_CUMULATIVE_FAST_DELTA() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MAXCUMULATIVEFASTDELTA() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXCUMULATIVEFASTDELTA(&_FastPriceFee.CallOpts)
}

// MAXCUMULATIVEREFDELTA is a free data retrieval call binding the contract method 0x0604ddea.
//
// Solidity: function MAX_CUMULATIVE_REF_DELTA() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MAXCUMULATIVEREFDELTA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "MAX_CUMULATIVE_REF_DELTA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCUMULATIVEREFDELTA is a free data retrieval call binding the contract method 0x0604ddea.
//
// Solidity: function MAX_CUMULATIVE_REF_DELTA() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MAXCUMULATIVEREFDELTA() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXCUMULATIVEREFDELTA(&_FastPriceFee.CallOpts)
}

// MAXCUMULATIVEREFDELTA is a free data retrieval call binding the contract method 0x0604ddea.
//
// Solidity: function MAX_CUMULATIVE_REF_DELTA() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MAXCUMULATIVEREFDELTA() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXCUMULATIVEREFDELTA(&_FastPriceFee.CallOpts)
}

// MAXPRICEDURATION is a free data retrieval call binding the contract method 0x668d3d65.
//
// Solidity: function MAX_PRICE_DURATION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MAXPRICEDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "MAX_PRICE_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPRICEDURATION is a free data retrieval call binding the contract method 0x668d3d65.
//
// Solidity: function MAX_PRICE_DURATION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MAXPRICEDURATION() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXPRICEDURATION(&_FastPriceFee.CallOpts)
}

// MAXPRICEDURATION is a free data retrieval call binding the contract method 0x668d3d65.
//
// Solidity: function MAX_PRICE_DURATION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MAXPRICEDURATION() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXPRICEDURATION(&_FastPriceFee.CallOpts)
}

// MAXREFPRICE is a free data retrieval call binding the contract method 0xe68a22c0.
//
// Solidity: function MAX_REF_PRICE() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MAXREFPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "MAX_REF_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREFPRICE is a free data retrieval call binding the contract method 0xe68a22c0.
//
// Solidity: function MAX_REF_PRICE() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MAXREFPRICE() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXREFPRICE(&_FastPriceFee.CallOpts)
}

// MAXREFPRICE is a free data retrieval call binding the contract method 0xe68a22c0.
//
// Solidity: function MAX_REF_PRICE() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MAXREFPRICE() (*big.Int, error) {
	return _FastPriceFee.Contract.MAXREFPRICE(&_FastPriceFee.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) PRICEPRECISION() (*big.Int, error) {
	return _FastPriceFee.Contract.PRICEPRECISION(&_FastPriceFee.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _FastPriceFee.Contract.PRICEPRECISION(&_FastPriceFee.CallOpts)
}

// DisableFastPriceVoteCount is a free data retrieval call binding the contract method 0xb0a25666.
//
// Solidity: function disableFastPriceVoteCount() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) DisableFastPriceVoteCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "disableFastPriceVoteCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisableFastPriceVoteCount is a free data retrieval call binding the contract method 0xb0a25666.
//
// Solidity: function disableFastPriceVoteCount() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) DisableFastPriceVoteCount() (*big.Int, error) {
	return _FastPriceFee.Contract.DisableFastPriceVoteCount(&_FastPriceFee.CallOpts)
}

// DisableFastPriceVoteCount is a free data retrieval call binding the contract method 0xb0a25666.
//
// Solidity: function disableFastPriceVoteCount() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) DisableFastPriceVoteCount() (*big.Int, error) {
	return _FastPriceFee.Contract.DisableFastPriceVoteCount(&_FastPriceFee.CallOpts)
}

// DisableFastPriceVotes is a free data retrieval call binding the contract method 0x03b04936.
//
// Solidity: function disableFastPriceVotes(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeCaller) DisableFastPriceVotes(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "disableFastPriceVotes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DisableFastPriceVotes is a free data retrieval call binding the contract method 0x03b04936.
//
// Solidity: function disableFastPriceVotes(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeSession) DisableFastPriceVotes(arg0 common.Address) (bool, error) {
	return _FastPriceFee.Contract.DisableFastPriceVotes(&_FastPriceFee.CallOpts, arg0)
}

// DisableFastPriceVotes is a free data retrieval call binding the contract method 0x03b04936.
//
// Solidity: function disableFastPriceVotes(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeCallerSession) DisableFastPriceVotes(arg0 common.Address) (bool, error) {
	return _FastPriceFee.Contract.DisableFastPriceVotes(&_FastPriceFee.CallOpts, arg0)
}

// FastPriceEvents is a free data retrieval call binding the contract method 0x0e9272ea.
//
// Solidity: function fastPriceEvents() view returns(address)
func (_FastPriceFee *FastPriceFeeCaller) FastPriceEvents(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "fastPriceEvents")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastPriceEvents is a free data retrieval call binding the contract method 0x0e9272ea.
//
// Solidity: function fastPriceEvents() view returns(address)
func (_FastPriceFee *FastPriceFeeSession) FastPriceEvents() (common.Address, error) {
	return _FastPriceFee.Contract.FastPriceEvents(&_FastPriceFee.CallOpts)
}

// FastPriceEvents is a free data retrieval call binding the contract method 0x0e9272ea.
//
// Solidity: function fastPriceEvents() view returns(address)
func (_FastPriceFee *FastPriceFeeCallerSession) FastPriceEvents() (common.Address, error) {
	return _FastPriceFee.Contract.FastPriceEvents(&_FastPriceFee.CallOpts)
}

// FavorFastPrice is a free data retrieval call binding the contract method 0x6c56fd05.
//
// Solidity: function favorFastPrice(address _token) view returns(bool)
func (_FastPriceFee *FastPriceFeeCaller) FavorFastPrice(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "favorFastPrice", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FavorFastPrice is a free data retrieval call binding the contract method 0x6c56fd05.
//
// Solidity: function favorFastPrice(address _token) view returns(bool)
func (_FastPriceFee *FastPriceFeeSession) FavorFastPrice(_token common.Address) (bool, error) {
	return _FastPriceFee.Contract.FavorFastPrice(&_FastPriceFee.CallOpts, _token)
}

// FavorFastPrice is a free data retrieval call binding the contract method 0x6c56fd05.
//
// Solidity: function favorFastPrice(address _token) view returns(bool)
func (_FastPriceFee *FastPriceFeeCallerSession) FavorFastPrice(_token common.Address) (bool, error) {
	return _FastPriceFee.Contract.FavorFastPrice(&_FastPriceFee.CallOpts, _token)
}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _refPrice, bool _maximise) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) GetPrice(opts *bind.CallOpts, _token common.Address, _refPrice *big.Int, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "getPrice", _token, _refPrice, _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _refPrice, bool _maximise) view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) GetPrice(_token common.Address, _refPrice *big.Int, _maximise bool) (*big.Int, error) {
	return _FastPriceFee.Contract.GetPrice(&_FastPriceFee.CallOpts, _token, _refPrice, _maximise)
}

// GetPrice is a free data retrieval call binding the contract method 0x7fece368.
//
// Solidity: function getPrice(address _token, uint256 _refPrice, bool _maximise) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) GetPrice(_token common.Address, _refPrice *big.Int, _maximise bool) (*big.Int, error) {
	return _FastPriceFee.Contract.GetPrice(&_FastPriceFee.CallOpts, _token, _refPrice, _maximise)
}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address _token) view returns(uint256, uint256, uint256, uint256)
func (_FastPriceFee *FastPriceFeeCaller) GetPriceData(opts *bind.CallOpts, _token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "getPriceData", _token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address _token) view returns(uint256, uint256, uint256, uint256)
func (_FastPriceFee *FastPriceFeeSession) GetPriceData(_token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FastPriceFee.Contract.GetPriceData(&_FastPriceFee.CallOpts, _token)
}

// GetPriceData is a free data retrieval call binding the contract method 0x72279ba1.
//
// Solidity: function getPriceData(address _token) view returns(uint256, uint256, uint256, uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) GetPriceData(_token common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FastPriceFee.Contract.GetPriceData(&_FastPriceFee.CallOpts, _token)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceFee *FastPriceFeeCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceFee *FastPriceFeeSession) Gov() (common.Address, error) {
	return _FastPriceFee.Contract.Gov(&_FastPriceFee.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_FastPriceFee *FastPriceFeeCallerSession) Gov() (common.Address, error) {
	return _FastPriceFee.Contract.Gov(&_FastPriceFee.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_FastPriceFee *FastPriceFeeCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_FastPriceFee *FastPriceFeeSession) IsInitialized() (bool, error) {
	return _FastPriceFee.Contract.IsInitialized(&_FastPriceFee.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_FastPriceFee *FastPriceFeeCallerSession) IsInitialized() (bool, error) {
	return _FastPriceFee.Contract.IsInitialized(&_FastPriceFee.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeCaller) IsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "isSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeSession) IsSigner(arg0 common.Address) (bool, error) {
	return _FastPriceFee.Contract.IsSigner(&_FastPriceFee.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeCallerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _FastPriceFee.Contract.IsSigner(&_FastPriceFee.CallOpts, arg0)
}

// IsSpreadEnabled is a free data retrieval call binding the contract method 0x695d4184.
//
// Solidity: function isSpreadEnabled() view returns(bool)
func (_FastPriceFee *FastPriceFeeCaller) IsSpreadEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "isSpreadEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpreadEnabled is a free data retrieval call binding the contract method 0x695d4184.
//
// Solidity: function isSpreadEnabled() view returns(bool)
func (_FastPriceFee *FastPriceFeeSession) IsSpreadEnabled() (bool, error) {
	return _FastPriceFee.Contract.IsSpreadEnabled(&_FastPriceFee.CallOpts)
}

// IsSpreadEnabled is a free data retrieval call binding the contract method 0x695d4184.
//
// Solidity: function isSpreadEnabled() view returns(bool)
func (_FastPriceFee *FastPriceFeeCallerSession) IsSpreadEnabled() (bool, error) {
	return _FastPriceFee.Contract.IsSpreadEnabled(&_FastPriceFee.CallOpts)
}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeCaller) IsUpdater(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "isUpdater", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeSession) IsUpdater(arg0 common.Address) (bool, error) {
	return _FastPriceFee.Contract.IsUpdater(&_FastPriceFee.CallOpts, arg0)
}

// IsUpdater is a free data retrieval call binding the contract method 0x4fdfb086.
//
// Solidity: function isUpdater(address ) view returns(bool)
func (_FastPriceFee *FastPriceFeeCallerSession) IsUpdater(arg0 common.Address) (bool, error) {
	return _FastPriceFee.Contract.IsUpdater(&_FastPriceFee.CallOpts, arg0)
}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) LastUpdatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "lastUpdatedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) LastUpdatedAt() (*big.Int, error) {
	return _FastPriceFee.Contract.LastUpdatedAt(&_FastPriceFee.CallOpts)
}

// LastUpdatedAt is a free data retrieval call binding the contract method 0x54aea127.
//
// Solidity: function lastUpdatedAt() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) LastUpdatedAt() (*big.Int, error) {
	return _FastPriceFee.Contract.LastUpdatedAt(&_FastPriceFee.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) LastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "lastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) LastUpdatedBlock() (*big.Int, error) {
	return _FastPriceFee.Contract.LastUpdatedBlock(&_FastPriceFee.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _FastPriceFee.Contract.LastUpdatedBlock(&_FastPriceFee.CallOpts)
}

// MaxCumulativeDeltaDiffs is a free data retrieval call binding the contract method 0xa3742425.
//
// Solidity: function maxCumulativeDeltaDiffs(address ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MaxCumulativeDeltaDiffs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "maxCumulativeDeltaDiffs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCumulativeDeltaDiffs is a free data retrieval call binding the contract method 0xa3742425.
//
// Solidity: function maxCumulativeDeltaDiffs(address ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MaxCumulativeDeltaDiffs(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFee.Contract.MaxCumulativeDeltaDiffs(&_FastPriceFee.CallOpts, arg0)
}

// MaxCumulativeDeltaDiffs is a free data retrieval call binding the contract method 0xa3742425.
//
// Solidity: function maxCumulativeDeltaDiffs(address ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MaxCumulativeDeltaDiffs(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFee.Contract.MaxCumulativeDeltaDiffs(&_FastPriceFee.CallOpts, arg0)
}

// MaxDeviationBasisPoints is a free data retrieval call binding the contract method 0x715c7536.
//
// Solidity: function maxDeviationBasisPoints() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MaxDeviationBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "maxDeviationBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeviationBasisPoints is a free data retrieval call binding the contract method 0x715c7536.
//
// Solidity: function maxDeviationBasisPoints() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MaxDeviationBasisPoints() (*big.Int, error) {
	return _FastPriceFee.Contract.MaxDeviationBasisPoints(&_FastPriceFee.CallOpts)
}

// MaxDeviationBasisPoints is a free data retrieval call binding the contract method 0x715c7536.
//
// Solidity: function maxDeviationBasisPoints() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MaxDeviationBasisPoints() (*big.Int, error) {
	return _FastPriceFee.Contract.MaxDeviationBasisPoints(&_FastPriceFee.CallOpts)
}

// MaxPriceUpdateDelay is a free data retrieval call binding the contract method 0xe64559ad.
//
// Solidity: function maxPriceUpdateDelay() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MaxPriceUpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "maxPriceUpdateDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPriceUpdateDelay is a free data retrieval call binding the contract method 0xe64559ad.
//
// Solidity: function maxPriceUpdateDelay() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MaxPriceUpdateDelay() (*big.Int, error) {
	return _FastPriceFee.Contract.MaxPriceUpdateDelay(&_FastPriceFee.CallOpts)
}

// MaxPriceUpdateDelay is a free data retrieval call binding the contract method 0xe64559ad.
//
// Solidity: function maxPriceUpdateDelay() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MaxPriceUpdateDelay() (*big.Int, error) {
	return _FastPriceFee.Contract.MaxPriceUpdateDelay(&_FastPriceFee.CallOpts)
}

// MaxTimeDeviation is a free data retrieval call binding the contract method 0x3aa08f86.
//
// Solidity: function maxTimeDeviation() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MaxTimeDeviation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "maxTimeDeviation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTimeDeviation is a free data retrieval call binding the contract method 0x3aa08f86.
//
// Solidity: function maxTimeDeviation() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MaxTimeDeviation() (*big.Int, error) {
	return _FastPriceFee.Contract.MaxTimeDeviation(&_FastPriceFee.CallOpts)
}

// MaxTimeDeviation is a free data retrieval call binding the contract method 0x3aa08f86.
//
// Solidity: function maxTimeDeviation() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MaxTimeDeviation() (*big.Int, error) {
	return _FastPriceFee.Contract.MaxTimeDeviation(&_FastPriceFee.CallOpts)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MinAuthorizations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "minAuthorizations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MinAuthorizations() (*big.Int, error) {
	return _FastPriceFee.Contract.MinAuthorizations(&_FastPriceFee.CallOpts)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MinAuthorizations() (*big.Int, error) {
	return _FastPriceFee.Contract.MinAuthorizations(&_FastPriceFee.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) MinBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "minBlockInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) MinBlockInterval() (*big.Int, error) {
	return _FastPriceFee.Contract.MinBlockInterval(&_FastPriceFee.CallOpts)
}

// MinBlockInterval is a free data retrieval call binding the contract method 0xb3606b56.
//
// Solidity: function minBlockInterval() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) MinBlockInterval() (*big.Int, error) {
	return _FastPriceFee.Contract.MinBlockInterval(&_FastPriceFee.CallOpts)
}

// PriceData is a free data retrieval call binding the contract method 0xcab44b76.
//
// Solidity: function priceData(address ) view returns(uint160 refPrice, uint32 refTime, uint32 cumulativeRefDelta, uint32 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeCaller) PriceData(opts *bind.CallOpts, arg0 common.Address) (struct {
	RefPrice            *big.Int
	RefTime             uint32
	CumulativeRefDelta  uint32
	CumulativeFastDelta uint32
}, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "priceData", arg0)

	outstruct := new(struct {
		RefPrice            *big.Int
		RefTime             uint32
		CumulativeRefDelta  uint32
		CumulativeFastDelta uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RefPrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RefTime = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.CumulativeRefDelta = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.CumulativeFastDelta = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// PriceData is a free data retrieval call binding the contract method 0xcab44b76.
//
// Solidity: function priceData(address ) view returns(uint160 refPrice, uint32 refTime, uint32 cumulativeRefDelta, uint32 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeSession) PriceData(arg0 common.Address) (struct {
	RefPrice            *big.Int
	RefTime             uint32
	CumulativeRefDelta  uint32
	CumulativeFastDelta uint32
}, error) {
	return _FastPriceFee.Contract.PriceData(&_FastPriceFee.CallOpts, arg0)
}

// PriceData is a free data retrieval call binding the contract method 0xcab44b76.
//
// Solidity: function priceData(address ) view returns(uint160 refPrice, uint32 refTime, uint32 cumulativeRefDelta, uint32 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeCallerSession) PriceData(arg0 common.Address) (struct {
	RefPrice            *big.Int
	RefTime             uint32
	CumulativeRefDelta  uint32
	CumulativeFastDelta uint32
}, error) {
	return _FastPriceFee.Contract.PriceData(&_FastPriceFee.CallOpts, arg0)
}

// PriceDataInterval is a free data retrieval call binding the contract method 0xdfb481c9.
//
// Solidity: function priceDataInterval() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) PriceDataInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "priceDataInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDataInterval is a free data retrieval call binding the contract method 0xdfb481c9.
//
// Solidity: function priceDataInterval() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) PriceDataInterval() (*big.Int, error) {
	return _FastPriceFee.Contract.PriceDataInterval(&_FastPriceFee.CallOpts)
}

// PriceDataInterval is a free data retrieval call binding the contract method 0xdfb481c9.
//
// Solidity: function priceDataInterval() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) PriceDataInterval() (*big.Int, error) {
	return _FastPriceFee.Contract.PriceDataInterval(&_FastPriceFee.CallOpts)
}

// PriceDuration is a free data retrieval call binding the contract method 0x03cd2571.
//
// Solidity: function priceDuration() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) PriceDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "priceDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceDuration is a free data retrieval call binding the contract method 0x03cd2571.
//
// Solidity: function priceDuration() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) PriceDuration() (*big.Int, error) {
	return _FastPriceFee.Contract.PriceDuration(&_FastPriceFee.CallOpts)
}

// PriceDuration is a free data retrieval call binding the contract method 0x03cd2571.
//
// Solidity: function priceDuration() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) PriceDuration() (*big.Int, error) {
	return _FastPriceFee.Contract.PriceDuration(&_FastPriceFee.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) Prices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "prices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) Prices(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFee.Contract.Prices(&_FastPriceFee.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) Prices(arg0 common.Address) (*big.Int, error) {
	return _FastPriceFee.Contract.Prices(&_FastPriceFee.CallOpts, arg0)
}

// SpreadBasisPointsIfChainError is a free data retrieval call binding the contract method 0xa6eca896.
//
// Solidity: function spreadBasisPointsIfChainError() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) SpreadBasisPointsIfChainError(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "spreadBasisPointsIfChainError")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadBasisPointsIfChainError is a free data retrieval call binding the contract method 0xa6eca896.
//
// Solidity: function spreadBasisPointsIfChainError() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) SpreadBasisPointsIfChainError() (*big.Int, error) {
	return _FastPriceFee.Contract.SpreadBasisPointsIfChainError(&_FastPriceFee.CallOpts)
}

// SpreadBasisPointsIfChainError is a free data retrieval call binding the contract method 0xa6eca896.
//
// Solidity: function spreadBasisPointsIfChainError() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) SpreadBasisPointsIfChainError() (*big.Int, error) {
	return _FastPriceFee.Contract.SpreadBasisPointsIfChainError(&_FastPriceFee.CallOpts)
}

// SpreadBasisPointsIfInactive is a free data retrieval call binding the contract method 0x74bfed89.
//
// Solidity: function spreadBasisPointsIfInactive() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) SpreadBasisPointsIfInactive(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "spreadBasisPointsIfInactive")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpreadBasisPointsIfInactive is a free data retrieval call binding the contract method 0x74bfed89.
//
// Solidity: function spreadBasisPointsIfInactive() view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) SpreadBasisPointsIfInactive() (*big.Int, error) {
	return _FastPriceFee.Contract.SpreadBasisPointsIfInactive(&_FastPriceFee.CallOpts)
}

// SpreadBasisPointsIfInactive is a free data retrieval call binding the contract method 0x74bfed89.
//
// Solidity: function spreadBasisPointsIfInactive() view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) SpreadBasisPointsIfInactive() (*big.Int, error) {
	return _FastPriceFee.Contract.SpreadBasisPointsIfInactive(&_FastPriceFee.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_FastPriceFee *FastPriceFeeCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_FastPriceFee *FastPriceFeeSession) TokenManager() (common.Address, error) {
	return _FastPriceFee.Contract.TokenManager(&_FastPriceFee.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_FastPriceFee *FastPriceFeeCallerSession) TokenManager() (common.Address, error) {
	return _FastPriceFee.Contract.TokenManager(&_FastPriceFee.CallOpts)
}

// TokenPrecisions is a free data retrieval call binding the contract method 0x4d11fb4a.
//
// Solidity: function tokenPrecisions(uint256 ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCaller) TokenPrecisions(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "tokenPrecisions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPrecisions is a free data retrieval call binding the contract method 0x4d11fb4a.
//
// Solidity: function tokenPrecisions(uint256 ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeSession) TokenPrecisions(arg0 *big.Int) (*big.Int, error) {
	return _FastPriceFee.Contract.TokenPrecisions(&_FastPriceFee.CallOpts, arg0)
}

// TokenPrecisions is a free data retrieval call binding the contract method 0x4d11fb4a.
//
// Solidity: function tokenPrecisions(uint256 ) view returns(uint256)
func (_FastPriceFee *FastPriceFeeCallerSession) TokenPrecisions(arg0 *big.Int) (*big.Int, error) {
	return _FastPriceFee.Contract.TokenPrecisions(&_FastPriceFee.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_FastPriceFee *FastPriceFeeCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_FastPriceFee *FastPriceFeeSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _FastPriceFee.Contract.Tokens(&_FastPriceFee.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_FastPriceFee *FastPriceFeeCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _FastPriceFee.Contract.Tokens(&_FastPriceFee.CallOpts, arg0)
}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_FastPriceFee *FastPriceFeeCaller) VaultPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastPriceFee.contract.Call(opts, &out, "vaultPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_FastPriceFee *FastPriceFeeSession) VaultPriceFeed() (common.Address, error) {
	return _FastPriceFee.Contract.VaultPriceFeed(&_FastPriceFee.CallOpts)
}

// VaultPriceFeed is a free data retrieval call binding the contract method 0xeeaa783a.
//
// Solidity: function vaultPriceFeed() view returns(address)
func (_FastPriceFee *FastPriceFeeCallerSession) VaultPriceFeed() (common.Address, error) {
	return _FastPriceFee.Contract.VaultPriceFeed(&_FastPriceFee.CallOpts)
}

// DisableFastPrice is a paid mutator transaction binding the contract method 0xc84a9124.
//
// Solidity: function disableFastPrice() returns()
func (_FastPriceFee *FastPriceFeeTransactor) DisableFastPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "disableFastPrice")
}

// DisableFastPrice is a paid mutator transaction binding the contract method 0xc84a9124.
//
// Solidity: function disableFastPrice() returns()
func (_FastPriceFee *FastPriceFeeSession) DisableFastPrice() (*types.Transaction, error) {
	return _FastPriceFee.Contract.DisableFastPrice(&_FastPriceFee.TransactOpts)
}

// DisableFastPrice is a paid mutator transaction binding the contract method 0xc84a9124.
//
// Solidity: function disableFastPrice() returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) DisableFastPrice() (*types.Transaction, error) {
	return _FastPriceFee.Contract.DisableFastPrice(&_FastPriceFee.TransactOpts)
}

// EnableFastPrice is a paid mutator transaction binding the contract method 0x6ccd47c4.
//
// Solidity: function enableFastPrice() returns()
func (_FastPriceFee *FastPriceFeeTransactor) EnableFastPrice(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "enableFastPrice")
}

// EnableFastPrice is a paid mutator transaction binding the contract method 0x6ccd47c4.
//
// Solidity: function enableFastPrice() returns()
func (_FastPriceFee *FastPriceFeeSession) EnableFastPrice() (*types.Transaction, error) {
	return _FastPriceFee.Contract.EnableFastPrice(&_FastPriceFee.TransactOpts)
}

// EnableFastPrice is a paid mutator transaction binding the contract method 0x6ccd47c4.
//
// Solidity: function enableFastPrice() returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) EnableFastPrice() (*types.Transaction, error) {
	return _FastPriceFee.Contract.EnableFastPrice(&_FastPriceFee.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbc79c6.
//
// Solidity: function initialize(uint256 _minAuthorizations, address[] _signers, address[] _updaters) returns()
func (_FastPriceFee *FastPriceFeeTransactor) Initialize(opts *bind.TransactOpts, _minAuthorizations *big.Int, _signers []common.Address, _updaters []common.Address) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "initialize", _minAuthorizations, _signers, _updaters)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbc79c6.
//
// Solidity: function initialize(uint256 _minAuthorizations, address[] _signers, address[] _updaters) returns()
func (_FastPriceFee *FastPriceFeeSession) Initialize(_minAuthorizations *big.Int, _signers []common.Address, _updaters []common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.Initialize(&_FastPriceFee.TransactOpts, _minAuthorizations, _signers, _updaters)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbc79c6.
//
// Solidity: function initialize(uint256 _minAuthorizations, address[] _signers, address[] _updaters) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) Initialize(_minAuthorizations *big.Int, _signers []common.Address, _updaters []common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.Initialize(&_FastPriceFee.TransactOpts, _minAuthorizations, _signers, _updaters)
}

// SetCompactedPrices is a paid mutator transaction binding the contract method 0x03f4d7dc.
//
// Solidity: function setCompactedPrices(uint256[] _priceBitArray, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetCompactedPrices(opts *bind.TransactOpts, _priceBitArray []*big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setCompactedPrices", _priceBitArray, _timestamp)
}

// SetCompactedPrices is a paid mutator transaction binding the contract method 0x03f4d7dc.
//
// Solidity: function setCompactedPrices(uint256[] _priceBitArray, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeSession) SetCompactedPrices(_priceBitArray []*big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetCompactedPrices(&_FastPriceFee.TransactOpts, _priceBitArray, _timestamp)
}

// SetCompactedPrices is a paid mutator transaction binding the contract method 0x03f4d7dc.
//
// Solidity: function setCompactedPrices(uint256[] _priceBitArray, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetCompactedPrices(_priceBitArray []*big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetCompactedPrices(&_FastPriceFee.TransactOpts, _priceBitArray, _timestamp)
}

// SetFastPriceEvents is a paid mutator transaction binding the contract method 0x162ac4e0.
//
// Solidity: function setFastPriceEvents(address _fastPriceEvents) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetFastPriceEvents(opts *bind.TransactOpts, _fastPriceEvents common.Address) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setFastPriceEvents", _fastPriceEvents)
}

// SetFastPriceEvents is a paid mutator transaction binding the contract method 0x162ac4e0.
//
// Solidity: function setFastPriceEvents(address _fastPriceEvents) returns()
func (_FastPriceFee *FastPriceFeeSession) SetFastPriceEvents(_fastPriceEvents common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetFastPriceEvents(&_FastPriceFee.TransactOpts, _fastPriceEvents)
}

// SetFastPriceEvents is a paid mutator transaction binding the contract method 0x162ac4e0.
//
// Solidity: function setFastPriceEvents(address _fastPriceEvents) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetFastPriceEvents(_fastPriceEvents common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetFastPriceEvents(&_FastPriceFee.TransactOpts, _fastPriceEvents)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceFee *FastPriceFeeSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetGov(&_FastPriceFee.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetGov(&_FastPriceFee.TransactOpts, _gov)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetIsSpreadEnabled(opts *bind.TransactOpts, _isSpreadEnabled bool) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setIsSpreadEnabled", _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_FastPriceFee *FastPriceFeeSession) SetIsSpreadEnabled(_isSpreadEnabled bool) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetIsSpreadEnabled(&_FastPriceFee.TransactOpts, _isSpreadEnabled)
}

// SetIsSpreadEnabled is a paid mutator transaction binding the contract method 0xce98dfa8.
//
// Solidity: function setIsSpreadEnabled(bool _isSpreadEnabled) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetIsSpreadEnabled(_isSpreadEnabled bool) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetIsSpreadEnabled(&_FastPriceFee.TransactOpts, _isSpreadEnabled)
}

// SetLastUpdatedAt is a paid mutator transaction binding the contract method 0x14dd2dce.
//
// Solidity: function setLastUpdatedAt(uint256 _lastUpdatedAt) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetLastUpdatedAt(opts *bind.TransactOpts, _lastUpdatedAt *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setLastUpdatedAt", _lastUpdatedAt)
}

// SetLastUpdatedAt is a paid mutator transaction binding the contract method 0x14dd2dce.
//
// Solidity: function setLastUpdatedAt(uint256 _lastUpdatedAt) returns()
func (_FastPriceFee *FastPriceFeeSession) SetLastUpdatedAt(_lastUpdatedAt *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetLastUpdatedAt(&_FastPriceFee.TransactOpts, _lastUpdatedAt)
}

// SetLastUpdatedAt is a paid mutator transaction binding the contract method 0x14dd2dce.
//
// Solidity: function setLastUpdatedAt(uint256 _lastUpdatedAt) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetLastUpdatedAt(_lastUpdatedAt *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetLastUpdatedAt(&_FastPriceFee.TransactOpts, _lastUpdatedAt)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetMaxCumulativeDeltaDiffs(opts *bind.TransactOpts, _tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setMaxCumulativeDeltaDiffs", _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_FastPriceFee *FastPriceFeeSession) SetMaxCumulativeDeltaDiffs(_tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxCumulativeDeltaDiffs(&_FastPriceFee.TransactOpts, _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxCumulativeDeltaDiffs is a paid mutator transaction binding the contract method 0x4c0e31c8.
//
// Solidity: function setMaxCumulativeDeltaDiffs(address[] _tokens, uint256[] _maxCumulativeDeltaDiffs) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetMaxCumulativeDeltaDiffs(_tokens []common.Address, _maxCumulativeDeltaDiffs []*big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxCumulativeDeltaDiffs(&_FastPriceFee.TransactOpts, _tokens, _maxCumulativeDeltaDiffs)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetMaxDeviationBasisPoints(opts *bind.TransactOpts, _maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setMaxDeviationBasisPoints", _maxDeviationBasisPoints)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_FastPriceFee *FastPriceFeeSession) SetMaxDeviationBasisPoints(_maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxDeviationBasisPoints(&_FastPriceFee.TransactOpts, _maxDeviationBasisPoints)
}

// SetMaxDeviationBasisPoints is a paid mutator transaction binding the contract method 0x82553aad.
//
// Solidity: function setMaxDeviationBasisPoints(uint256 _maxDeviationBasisPoints) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetMaxDeviationBasisPoints(_maxDeviationBasisPoints *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxDeviationBasisPoints(&_FastPriceFee.TransactOpts, _maxDeviationBasisPoints)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetMaxPriceUpdateDelay(opts *bind.TransactOpts, _maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setMaxPriceUpdateDelay", _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_FastPriceFee *FastPriceFeeSession) SetMaxPriceUpdateDelay(_maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxPriceUpdateDelay(&_FastPriceFee.TransactOpts, _maxPriceUpdateDelay)
}

// SetMaxPriceUpdateDelay is a paid mutator transaction binding the contract method 0x8b7677f4.
//
// Solidity: function setMaxPriceUpdateDelay(uint256 _maxPriceUpdateDelay) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetMaxPriceUpdateDelay(_maxPriceUpdateDelay *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxPriceUpdateDelay(&_FastPriceFee.TransactOpts, _maxPriceUpdateDelay)
}

// SetMaxTimeDeviation is a paid mutator transaction binding the contract method 0x776d16c1.
//
// Solidity: function setMaxTimeDeviation(uint256 _maxTimeDeviation) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetMaxTimeDeviation(opts *bind.TransactOpts, _maxTimeDeviation *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setMaxTimeDeviation", _maxTimeDeviation)
}

// SetMaxTimeDeviation is a paid mutator transaction binding the contract method 0x776d16c1.
//
// Solidity: function setMaxTimeDeviation(uint256 _maxTimeDeviation) returns()
func (_FastPriceFee *FastPriceFeeSession) SetMaxTimeDeviation(_maxTimeDeviation *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxTimeDeviation(&_FastPriceFee.TransactOpts, _maxTimeDeviation)
}

// SetMaxTimeDeviation is a paid mutator transaction binding the contract method 0x776d16c1.
//
// Solidity: function setMaxTimeDeviation(uint256 _maxTimeDeviation) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetMaxTimeDeviation(_maxTimeDeviation *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMaxTimeDeviation(&_FastPriceFee.TransactOpts, _maxTimeDeviation)
}

// SetMinAuthorizations is a paid mutator transaction binding the contract method 0xd925351a.
//
// Solidity: function setMinAuthorizations(uint256 _minAuthorizations) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetMinAuthorizations(opts *bind.TransactOpts, _minAuthorizations *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setMinAuthorizations", _minAuthorizations)
}

// SetMinAuthorizations is a paid mutator transaction binding the contract method 0xd925351a.
//
// Solidity: function setMinAuthorizations(uint256 _minAuthorizations) returns()
func (_FastPriceFee *FastPriceFeeSession) SetMinAuthorizations(_minAuthorizations *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMinAuthorizations(&_FastPriceFee.TransactOpts, _minAuthorizations)
}

// SetMinAuthorizations is a paid mutator transaction binding the contract method 0xd925351a.
//
// Solidity: function setMinAuthorizations(uint256 _minAuthorizations) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetMinAuthorizations(_minAuthorizations *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMinAuthorizations(&_FastPriceFee.TransactOpts, _minAuthorizations)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetMinBlockInterval(opts *bind.TransactOpts, _minBlockInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setMinBlockInterval", _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_FastPriceFee *FastPriceFeeSession) SetMinBlockInterval(_minBlockInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMinBlockInterval(&_FastPriceFee.TransactOpts, _minBlockInterval)
}

// SetMinBlockInterval is a paid mutator transaction binding the contract method 0xd6a153f1.
//
// Solidity: function setMinBlockInterval(uint256 _minBlockInterval) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetMinBlockInterval(_minBlockInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetMinBlockInterval(&_FastPriceFee.TransactOpts, _minBlockInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetPriceDataInterval(opts *bind.TransactOpts, _priceDataInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setPriceDataInterval", _priceDataInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_FastPriceFee *FastPriceFeeSession) SetPriceDataInterval(_priceDataInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPriceDataInterval(&_FastPriceFee.TransactOpts, _priceDataInterval)
}

// SetPriceDataInterval is a paid mutator transaction binding the contract method 0x2e9cd94b.
//
// Solidity: function setPriceDataInterval(uint256 _priceDataInterval) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetPriceDataInterval(_priceDataInterval *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPriceDataInterval(&_FastPriceFee.TransactOpts, _priceDataInterval)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetPriceDuration(opts *bind.TransactOpts, _priceDuration *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setPriceDuration", _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_FastPriceFee *FastPriceFeeSession) SetPriceDuration(_priceDuration *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPriceDuration(&_FastPriceFee.TransactOpts, _priceDuration)
}

// SetPriceDuration is a paid mutator transaction binding the contract method 0x44c23193.
//
// Solidity: function setPriceDuration(uint256 _priceDuration) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetPriceDuration(_priceDuration *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPriceDuration(&_FastPriceFee.TransactOpts, _priceDuration)
}

// SetPrices is a paid mutator transaction binding the contract method 0x782661bc.
//
// Solidity: function setPrices(address[] _tokens, uint256[] _prices, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetPrices(opts *bind.TransactOpts, _tokens []common.Address, _prices []*big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setPrices", _tokens, _prices, _timestamp)
}

// SetPrices is a paid mutator transaction binding the contract method 0x782661bc.
//
// Solidity: function setPrices(address[] _tokens, uint256[] _prices, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeSession) SetPrices(_tokens []common.Address, _prices []*big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPrices(&_FastPriceFee.TransactOpts, _tokens, _prices, _timestamp)
}

// SetPrices is a paid mutator transaction binding the contract method 0x782661bc.
//
// Solidity: function setPrices(address[] _tokens, uint256[] _prices, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetPrices(_tokens []common.Address, _prices []*big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPrices(&_FastPriceFee.TransactOpts, _tokens, _prices, _timestamp)
}

// SetPricesWithBits is a paid mutator transaction binding the contract method 0x17835d1c.
//
// Solidity: function setPricesWithBits(uint256 _priceBits, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetPricesWithBits(opts *bind.TransactOpts, _priceBits *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setPricesWithBits", _priceBits, _timestamp)
}

// SetPricesWithBits is a paid mutator transaction binding the contract method 0x17835d1c.
//
// Solidity: function setPricesWithBits(uint256 _priceBits, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeSession) SetPricesWithBits(_priceBits *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPricesWithBits(&_FastPriceFee.TransactOpts, _priceBits, _timestamp)
}

// SetPricesWithBits is a paid mutator transaction binding the contract method 0x17835d1c.
//
// Solidity: function setPricesWithBits(uint256 _priceBits, uint256 _timestamp) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetPricesWithBits(_priceBits *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPricesWithBits(&_FastPriceFee.TransactOpts, _priceBits, _timestamp)
}

// SetPricesWithBitsAndExecute is a paid mutator transaction binding the contract method 0x32e5f9fa.
//
// Solidity: function setPricesWithBitsAndExecute(address _positionRouter, uint256 _priceBits, uint256 _timestamp, uint256 _endIndexForIncreasePositions, uint256 _endIndexForDecreasePositions, uint256 _maxIncreasePositions, uint256 _maxDecreasePositions) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetPricesWithBitsAndExecute(opts *bind.TransactOpts, _positionRouter common.Address, _priceBits *big.Int, _timestamp *big.Int, _endIndexForIncreasePositions *big.Int, _endIndexForDecreasePositions *big.Int, _maxIncreasePositions *big.Int, _maxDecreasePositions *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setPricesWithBitsAndExecute", _positionRouter, _priceBits, _timestamp, _endIndexForIncreasePositions, _endIndexForDecreasePositions, _maxIncreasePositions, _maxDecreasePositions)
}

// SetPricesWithBitsAndExecute is a paid mutator transaction binding the contract method 0x32e5f9fa.
//
// Solidity: function setPricesWithBitsAndExecute(address _positionRouter, uint256 _priceBits, uint256 _timestamp, uint256 _endIndexForIncreasePositions, uint256 _endIndexForDecreasePositions, uint256 _maxIncreasePositions, uint256 _maxDecreasePositions) returns()
func (_FastPriceFee *FastPriceFeeSession) SetPricesWithBitsAndExecute(_positionRouter common.Address, _priceBits *big.Int, _timestamp *big.Int, _endIndexForIncreasePositions *big.Int, _endIndexForDecreasePositions *big.Int, _maxIncreasePositions *big.Int, _maxDecreasePositions *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPricesWithBitsAndExecute(&_FastPriceFee.TransactOpts, _positionRouter, _priceBits, _timestamp, _endIndexForIncreasePositions, _endIndexForDecreasePositions, _maxIncreasePositions, _maxDecreasePositions)
}

// SetPricesWithBitsAndExecute is a paid mutator transaction binding the contract method 0x32e5f9fa.
//
// Solidity: function setPricesWithBitsAndExecute(address _positionRouter, uint256 _priceBits, uint256 _timestamp, uint256 _endIndexForIncreasePositions, uint256 _endIndexForDecreasePositions, uint256 _maxIncreasePositions, uint256 _maxDecreasePositions) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetPricesWithBitsAndExecute(_positionRouter common.Address, _priceBits *big.Int, _timestamp *big.Int, _endIndexForIncreasePositions *big.Int, _endIndexForDecreasePositions *big.Int, _maxIncreasePositions *big.Int, _maxDecreasePositions *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetPricesWithBitsAndExecute(&_FastPriceFee.TransactOpts, _positionRouter, _priceBits, _timestamp, _endIndexForIncreasePositions, _endIndexForDecreasePositions, _maxIncreasePositions, _maxDecreasePositions)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetSigner(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setSigner", _account, _isActive)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_FastPriceFee *FastPriceFeeSession) SetSigner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetSigner(&_FastPriceFee.TransactOpts, _account, _isActive)
}

// SetSigner is a paid mutator transaction binding the contract method 0x31cb6105.
//
// Solidity: function setSigner(address _account, bool _isActive) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetSigner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetSigner(&_FastPriceFee.TransactOpts, _account, _isActive)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetSpreadBasisPointsIfChainError(opts *bind.TransactOpts, _spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setSpreadBasisPointsIfChainError", _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_FastPriceFee *FastPriceFeeSession) SetSpreadBasisPointsIfChainError(_spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetSpreadBasisPointsIfChainError(&_FastPriceFee.TransactOpts, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfChainError is a paid mutator transaction binding the contract method 0xde0d1b94.
//
// Solidity: function setSpreadBasisPointsIfChainError(uint256 _spreadBasisPointsIfChainError) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetSpreadBasisPointsIfChainError(_spreadBasisPointsIfChainError *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetSpreadBasisPointsIfChainError(&_FastPriceFee.TransactOpts, _spreadBasisPointsIfChainError)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetSpreadBasisPointsIfInactive(opts *bind.TransactOpts, _spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setSpreadBasisPointsIfInactive", _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_FastPriceFee *FastPriceFeeSession) SetSpreadBasisPointsIfInactive(_spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetSpreadBasisPointsIfInactive(&_FastPriceFee.TransactOpts, _spreadBasisPointsIfInactive)
}

// SetSpreadBasisPointsIfInactive is a paid mutator transaction binding the contract method 0xb70c7b70.
//
// Solidity: function setSpreadBasisPointsIfInactive(uint256 _spreadBasisPointsIfInactive) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetSpreadBasisPointsIfInactive(_spreadBasisPointsIfInactive *big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetSpreadBasisPointsIfInactive(&_FastPriceFee.TransactOpts, _spreadBasisPointsIfInactive)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetTokenManager(opts *bind.TransactOpts, _tokenManager common.Address) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setTokenManager", _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_FastPriceFee *FastPriceFeeSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetTokenManager(&_FastPriceFee.TransactOpts, _tokenManager)
}

// SetTokenManager is a paid mutator transaction binding the contract method 0x7cb2b79c.
//
// Solidity: function setTokenManager(address _tokenManager) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetTokenManager(_tokenManager common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetTokenManager(&_FastPriceFee.TransactOpts, _tokenManager)
}

// SetTokens is a paid mutator transaction binding the contract method 0xc8390a48.
//
// Solidity: function setTokens(address[] _tokens, uint256[] _tokenPrecisions) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetTokens(opts *bind.TransactOpts, _tokens []common.Address, _tokenPrecisions []*big.Int) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setTokens", _tokens, _tokenPrecisions)
}

// SetTokens is a paid mutator transaction binding the contract method 0xc8390a48.
//
// Solidity: function setTokens(address[] _tokens, uint256[] _tokenPrecisions) returns()
func (_FastPriceFee *FastPriceFeeSession) SetTokens(_tokens []common.Address, _tokenPrecisions []*big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetTokens(&_FastPriceFee.TransactOpts, _tokens, _tokenPrecisions)
}

// SetTokens is a paid mutator transaction binding the contract method 0xc8390a48.
//
// Solidity: function setTokens(address[] _tokens, uint256[] _tokenPrecisions) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetTokens(_tokens []common.Address, _tokenPrecisions []*big.Int) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetTokens(&_FastPriceFee.TransactOpts, _tokens, _tokenPrecisions)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetUpdater(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setUpdater", _account, _isActive)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_FastPriceFee *FastPriceFeeSession) SetUpdater(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetUpdater(&_FastPriceFee.TransactOpts, _account, _isActive)
}

// SetUpdater is a paid mutator transaction binding the contract method 0x1a153391.
//
// Solidity: function setUpdater(address _account, bool _isActive) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetUpdater(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetUpdater(&_FastPriceFee.TransactOpts, _account, _isActive)
}

// SetVaultPriceFeed is a paid mutator transaction binding the contract method 0x238aafb7.
//
// Solidity: function setVaultPriceFeed(address _vaultPriceFeed) returns()
func (_FastPriceFee *FastPriceFeeTransactor) SetVaultPriceFeed(opts *bind.TransactOpts, _vaultPriceFeed common.Address) (*types.Transaction, error) {
	return _FastPriceFee.contract.Transact(opts, "setVaultPriceFeed", _vaultPriceFeed)
}

// SetVaultPriceFeed is a paid mutator transaction binding the contract method 0x238aafb7.
//
// Solidity: function setVaultPriceFeed(address _vaultPriceFeed) returns()
func (_FastPriceFee *FastPriceFeeSession) SetVaultPriceFeed(_vaultPriceFeed common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetVaultPriceFeed(&_FastPriceFee.TransactOpts, _vaultPriceFeed)
}

// SetVaultPriceFeed is a paid mutator transaction binding the contract method 0x238aafb7.
//
// Solidity: function setVaultPriceFeed(address _vaultPriceFeed) returns()
func (_FastPriceFee *FastPriceFeeTransactorSession) SetVaultPriceFeed(_vaultPriceFeed common.Address) (*types.Transaction, error) {
	return _FastPriceFee.Contract.SetVaultPriceFeed(&_FastPriceFee.TransactOpts, _vaultPriceFeed)
}

// FastPriceFeeDisableFastPriceIterator is returned from FilterDisableFastPrice and is used to iterate over the raw logs and unpacked data for DisableFastPrice events raised by the FastPriceFee contract.
type FastPriceFeeDisableFastPriceIterator struct {
	Event *FastPriceFeeDisableFastPrice // Event containing the contract specifics and raw log

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
func (it *FastPriceFeeDisableFastPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeeDisableFastPrice)
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
		it.Event = new(FastPriceFeeDisableFastPrice)
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
func (it *FastPriceFeeDisableFastPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeeDisableFastPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeeDisableFastPrice represents a DisableFastPrice event raised by the FastPriceFee contract.
type FastPriceFeeDisableFastPrice struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDisableFastPrice is a free log retrieval operation binding the contract event 0x4c0c5fabf50e808e3bc8d19577d305e3a7163eea7e8a74a50caa8896694cd44b.
//
// Solidity: event DisableFastPrice(address signer)
func (_FastPriceFee *FastPriceFeeFilterer) FilterDisableFastPrice(opts *bind.FilterOpts) (*FastPriceFeeDisableFastPriceIterator, error) {

	logs, sub, err := _FastPriceFee.contract.FilterLogs(opts, "DisableFastPrice")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeeDisableFastPriceIterator{contract: _FastPriceFee.contract, event: "DisableFastPrice", logs: logs, sub: sub}, nil
}

// WatchDisableFastPrice is a free log subscription operation binding the contract event 0x4c0c5fabf50e808e3bc8d19577d305e3a7163eea7e8a74a50caa8896694cd44b.
//
// Solidity: event DisableFastPrice(address signer)
func (_FastPriceFee *FastPriceFeeFilterer) WatchDisableFastPrice(opts *bind.WatchOpts, sink chan<- *FastPriceFeeDisableFastPrice) (event.Subscription, error) {

	logs, sub, err := _FastPriceFee.contract.WatchLogs(opts, "DisableFastPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeeDisableFastPrice)
				if err := _FastPriceFee.contract.UnpackLog(event, "DisableFastPrice", log); err != nil {
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

// ParseDisableFastPrice is a log parse operation binding the contract event 0x4c0c5fabf50e808e3bc8d19577d305e3a7163eea7e8a74a50caa8896694cd44b.
//
// Solidity: event DisableFastPrice(address signer)
func (_FastPriceFee *FastPriceFeeFilterer) ParseDisableFastPrice(log types.Log) (*FastPriceFeeDisableFastPrice, error) {
	event := new(FastPriceFeeDisableFastPrice)
	if err := _FastPriceFee.contract.UnpackLog(event, "DisableFastPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastPriceFeeEnableFastPriceIterator is returned from FilterEnableFastPrice and is used to iterate over the raw logs and unpacked data for EnableFastPrice events raised by the FastPriceFee contract.
type FastPriceFeeEnableFastPriceIterator struct {
	Event *FastPriceFeeEnableFastPrice // Event containing the contract specifics and raw log

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
func (it *FastPriceFeeEnableFastPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeeEnableFastPrice)
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
		it.Event = new(FastPriceFeeEnableFastPrice)
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
func (it *FastPriceFeeEnableFastPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeeEnableFastPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeeEnableFastPrice represents a EnableFastPrice event raised by the FastPriceFee contract.
type FastPriceFeeEnableFastPrice struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEnableFastPrice is a free log retrieval operation binding the contract event 0x9fe0c305c33aa92757a537936872a60be0d91549a4303cc99fd8b7fce8a00275.
//
// Solidity: event EnableFastPrice(address signer)
func (_FastPriceFee *FastPriceFeeFilterer) FilterEnableFastPrice(opts *bind.FilterOpts) (*FastPriceFeeEnableFastPriceIterator, error) {

	logs, sub, err := _FastPriceFee.contract.FilterLogs(opts, "EnableFastPrice")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeeEnableFastPriceIterator{contract: _FastPriceFee.contract, event: "EnableFastPrice", logs: logs, sub: sub}, nil
}

// WatchEnableFastPrice is a free log subscription operation binding the contract event 0x9fe0c305c33aa92757a537936872a60be0d91549a4303cc99fd8b7fce8a00275.
//
// Solidity: event EnableFastPrice(address signer)
func (_FastPriceFee *FastPriceFeeFilterer) WatchEnableFastPrice(opts *bind.WatchOpts, sink chan<- *FastPriceFeeEnableFastPrice) (event.Subscription, error) {

	logs, sub, err := _FastPriceFee.contract.WatchLogs(opts, "EnableFastPrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeeEnableFastPrice)
				if err := _FastPriceFee.contract.UnpackLog(event, "EnableFastPrice", log); err != nil {
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

// ParseEnableFastPrice is a log parse operation binding the contract event 0x9fe0c305c33aa92757a537936872a60be0d91549a4303cc99fd8b7fce8a00275.
//
// Solidity: event EnableFastPrice(address signer)
func (_FastPriceFee *FastPriceFeeFilterer) ParseEnableFastPrice(log types.Log) (*FastPriceFeeEnableFastPrice, error) {
	event := new(FastPriceFeeEnableFastPrice)
	if err := _FastPriceFee.contract.UnpackLog(event, "EnableFastPrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastPriceFeeMaxCumulativeDeltaDiffExceededIterator is returned from FilterMaxCumulativeDeltaDiffExceeded and is used to iterate over the raw logs and unpacked data for MaxCumulativeDeltaDiffExceeded events raised by the FastPriceFee contract.
type FastPriceFeeMaxCumulativeDeltaDiffExceededIterator struct {
	Event *FastPriceFeeMaxCumulativeDeltaDiffExceeded // Event containing the contract specifics and raw log

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
func (it *FastPriceFeeMaxCumulativeDeltaDiffExceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeeMaxCumulativeDeltaDiffExceeded)
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
		it.Event = new(FastPriceFeeMaxCumulativeDeltaDiffExceeded)
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
func (it *FastPriceFeeMaxCumulativeDeltaDiffExceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeeMaxCumulativeDeltaDiffExceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeeMaxCumulativeDeltaDiffExceeded represents a MaxCumulativeDeltaDiffExceeded event raised by the FastPriceFee contract.
type FastPriceFeeMaxCumulativeDeltaDiffExceeded struct {
	Token               common.Address
	RefPrice            *big.Int
	FastPrice           *big.Int
	CumulativeRefDelta  *big.Int
	CumulativeFastDelta *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMaxCumulativeDeltaDiffExceeded is a free log retrieval operation binding the contract event 0xe582322b389ad06b2bbf619cd6da3f16a288ec873ea0fa6df4d72f3d9480b447.
//
// Solidity: event MaxCumulativeDeltaDiffExceeded(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeFilterer) FilterMaxCumulativeDeltaDiffExceeded(opts *bind.FilterOpts) (*FastPriceFeeMaxCumulativeDeltaDiffExceededIterator, error) {

	logs, sub, err := _FastPriceFee.contract.FilterLogs(opts, "MaxCumulativeDeltaDiffExceeded")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeeMaxCumulativeDeltaDiffExceededIterator{contract: _FastPriceFee.contract, event: "MaxCumulativeDeltaDiffExceeded", logs: logs, sub: sub}, nil
}

// WatchMaxCumulativeDeltaDiffExceeded is a free log subscription operation binding the contract event 0xe582322b389ad06b2bbf619cd6da3f16a288ec873ea0fa6df4d72f3d9480b447.
//
// Solidity: event MaxCumulativeDeltaDiffExceeded(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeFilterer) WatchMaxCumulativeDeltaDiffExceeded(opts *bind.WatchOpts, sink chan<- *FastPriceFeeMaxCumulativeDeltaDiffExceeded) (event.Subscription, error) {

	logs, sub, err := _FastPriceFee.contract.WatchLogs(opts, "MaxCumulativeDeltaDiffExceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeeMaxCumulativeDeltaDiffExceeded)
				if err := _FastPriceFee.contract.UnpackLog(event, "MaxCumulativeDeltaDiffExceeded", log); err != nil {
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

// ParseMaxCumulativeDeltaDiffExceeded is a log parse operation binding the contract event 0xe582322b389ad06b2bbf619cd6da3f16a288ec873ea0fa6df4d72f3d9480b447.
//
// Solidity: event MaxCumulativeDeltaDiffExceeded(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeFilterer) ParseMaxCumulativeDeltaDiffExceeded(log types.Log) (*FastPriceFeeMaxCumulativeDeltaDiffExceeded, error) {
	event := new(FastPriceFeeMaxCumulativeDeltaDiffExceeded)
	if err := _FastPriceFee.contract.UnpackLog(event, "MaxCumulativeDeltaDiffExceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastPriceFeePriceDataIterator is returned from FilterPriceData and is used to iterate over the raw logs and unpacked data for PriceData events raised by the FastPriceFee contract.
type FastPriceFeePriceDataIterator struct {
	Event *FastPriceFeePriceData // Event containing the contract specifics and raw log

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
func (it *FastPriceFeePriceDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastPriceFeePriceData)
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
		it.Event = new(FastPriceFeePriceData)
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
func (it *FastPriceFeePriceDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastPriceFeePriceDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastPriceFeePriceData represents a PriceData event raised by the FastPriceFee contract.
type FastPriceFeePriceData struct {
	Token               common.Address
	RefPrice            *big.Int
	FastPrice           *big.Int
	CumulativeRefDelta  *big.Int
	CumulativeFastDelta *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPriceData is a free log retrieval operation binding the contract event 0x23b9387f81fca646aac1dc4487ede045c65f5f7445482906565f01e05afdb3a8.
//
// Solidity: event PriceData(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeFilterer) FilterPriceData(opts *bind.FilterOpts) (*FastPriceFeePriceDataIterator, error) {

	logs, sub, err := _FastPriceFee.contract.FilterLogs(opts, "PriceData")
	if err != nil {
		return nil, err
	}
	return &FastPriceFeePriceDataIterator{contract: _FastPriceFee.contract, event: "PriceData", logs: logs, sub: sub}, nil
}

// WatchPriceData is a free log subscription operation binding the contract event 0x23b9387f81fca646aac1dc4487ede045c65f5f7445482906565f01e05afdb3a8.
//
// Solidity: event PriceData(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeFilterer) WatchPriceData(opts *bind.WatchOpts, sink chan<- *FastPriceFeePriceData) (event.Subscription, error) {

	logs, sub, err := _FastPriceFee.contract.WatchLogs(opts, "PriceData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastPriceFeePriceData)
				if err := _FastPriceFee.contract.UnpackLog(event, "PriceData", log); err != nil {
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

// ParsePriceData is a log parse operation binding the contract event 0x23b9387f81fca646aac1dc4487ede045c65f5f7445482906565f01e05afdb3a8.
//
// Solidity: event PriceData(address token, uint256 refPrice, uint256 fastPrice, uint256 cumulativeRefDelta, uint256 cumulativeFastDelta)
func (_FastPriceFee *FastPriceFeeFilterer) ParsePriceData(log types.Log) (*FastPriceFeePriceData, error) {
	event := new(FastPriceFeePriceData)
	if err := _FastPriceFee.contract.UnpackLog(event, "PriceData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
