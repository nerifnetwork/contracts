// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractkeys

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

// ContractKeysMetaData contains all meta data concerning the ContractKeys contract.
var ContractKeysMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractKeysABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractKeysMetaData.ABI instead.
var ContractKeysABI = ContractKeysMetaData.ABI

// ContractKeys is an auto generated Go binding around an Ethereum contract.
type ContractKeys struct {
	ContractKeysCaller     // Read-only binding to the contract
	ContractKeysTransactor // Write-only binding to the contract
	ContractKeysFilterer   // Log filterer for contract events
}

// ContractKeysCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractKeysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKeysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractKeysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKeysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractKeysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractKeysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractKeysSession struct {
	Contract     *ContractKeys     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractKeysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractKeysCallerSession struct {
	Contract *ContractKeysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ContractKeysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractKeysTransactorSession struct {
	Contract     *ContractKeysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractKeysRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractKeysRaw struct {
	Contract *ContractKeys // Generic contract binding to access the raw methods on
}

// ContractKeysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractKeysCallerRaw struct {
	Contract *ContractKeysCaller // Generic read-only contract binding to access the raw methods on
}

// ContractKeysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractKeysTransactorRaw struct {
	Contract *ContractKeysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractKeys creates a new instance of ContractKeys, bound to a specific deployed contract.
func NewContractKeys(address common.Address, backend bind.ContractBackend) (*ContractKeys, error) {
	contract, err := bindContractKeys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractKeys{ContractKeysCaller: ContractKeysCaller{contract: contract}, ContractKeysTransactor: ContractKeysTransactor{contract: contract}, ContractKeysFilterer: ContractKeysFilterer{contract: contract}}, nil
}

// NewContractKeysCaller creates a new read-only instance of ContractKeys, bound to a specific deployed contract.
func NewContractKeysCaller(address common.Address, caller bind.ContractCaller) (*ContractKeysCaller, error) {
	contract, err := bindContractKeys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractKeysCaller{contract: contract}, nil
}

// NewContractKeysTransactor creates a new write-only instance of ContractKeys, bound to a specific deployed contract.
func NewContractKeysTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractKeysTransactor, error) {
	contract, err := bindContractKeys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractKeysTransactor{contract: contract}, nil
}

// NewContractKeysFilterer creates a new log filterer instance of ContractKeys, bound to a specific deployed contract.
func NewContractKeysFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractKeysFilterer, error) {
	contract, err := bindContractKeys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractKeysFilterer{contract: contract}, nil
}

// bindContractKeys binds a generic wrapper to an already deployed contract.
func bindContractKeys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractKeysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractKeys *ContractKeysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractKeys.Contract.ContractKeysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractKeys *ContractKeysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKeys.Contract.ContractKeysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractKeys *ContractKeysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractKeys.Contract.ContractKeysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractKeys *ContractKeysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractKeys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractKeys *ContractKeysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractKeys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractKeys *ContractKeysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractKeys.Contract.contract.Transact(opts, method, params...)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_ContractKeys *ContractKeysCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractKeys.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_ContractKeys *ContractKeysSession) DKGKEY() (string, error) {
	return _ContractKeys.Contract.DKGKEY(&_ContractKeys.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_ContractKeys *ContractKeysCallerSession) DKGKEY() (string, error) {
	return _ContractKeys.Contract.DKGKEY(&_ContractKeys.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_ContractKeys *ContractKeysCaller) REWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractKeys.contract.Call(opts, &out, "REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_ContractKeys *ContractKeysSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _ContractKeys.Contract.REWARDDISTRIBUTIONPOOLKEY(&_ContractKeys.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_ContractKeys *ContractKeysCallerSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _ContractKeys.Contract.REWARDDISTRIBUTIONPOOLKEY(&_ContractKeys.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_ContractKeys *ContractKeysCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractKeys.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_ContractKeys *ContractKeysSession) SLASHINGVOTINGKEY() (string, error) {
	return _ContractKeys.Contract.SLASHINGVOTINGKEY(&_ContractKeys.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_ContractKeys *ContractKeysCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _ContractKeys.Contract.SLASHINGVOTINGKEY(&_ContractKeys.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_ContractKeys *ContractKeysCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractKeys.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_ContractKeys *ContractKeysSession) STAKINGKEY() (string, error) {
	return _ContractKeys.Contract.STAKINGKEY(&_ContractKeys.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_ContractKeys *ContractKeysCallerSession) STAKINGKEY() (string, error) {
	return _ContractKeys.Contract.STAKINGKEY(&_ContractKeys.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_ContractKeys *ContractKeysCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractKeys.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_ContractKeys *ContractKeysSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _ContractKeys.Contract.SUPPORTEDTOKENSKEY(&_ContractKeys.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_ContractKeys *ContractKeysCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _ContractKeys.Contract.SUPPORTEDTOKENSKEY(&_ContractKeys.CallOpts)
}
