// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system

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

// RewardDistributionPoolMetaData contains all meta data concerning the RewardDistributionPool contract.
var RewardDistributionPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollectRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPositions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsOwing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RewardDistributionPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardDistributionPoolMetaData.ABI instead.
var RewardDistributionPoolABI = RewardDistributionPoolMetaData.ABI

// RewardDistributionPool is an auto generated Go binding around an Ethereum contract.
type RewardDistributionPool struct {
	RewardDistributionPoolCaller     // Read-only binding to the contract
	RewardDistributionPoolTransactor // Write-only binding to the contract
	RewardDistributionPoolFilterer   // Log filterer for contract events
}

// RewardDistributionPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardDistributionPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributionPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardDistributionPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributionPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardDistributionPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributionPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardDistributionPoolSession struct {
	Contract     *RewardDistributionPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RewardDistributionPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardDistributionPoolCallerSession struct {
	Contract *RewardDistributionPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// RewardDistributionPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardDistributionPoolTransactorSession struct {
	Contract     *RewardDistributionPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// RewardDistributionPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardDistributionPoolRaw struct {
	Contract *RewardDistributionPool // Generic contract binding to access the raw methods on
}

// RewardDistributionPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardDistributionPoolCallerRaw struct {
	Contract *RewardDistributionPoolCaller // Generic read-only contract binding to access the raw methods on
}

// RewardDistributionPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardDistributionPoolTransactorRaw struct {
	Contract *RewardDistributionPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardDistributionPool creates a new instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPool(address common.Address, backend bind.ContractBackend) (*RewardDistributionPool, error) {
	contract, err := bindRewardDistributionPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPool{RewardDistributionPoolCaller: RewardDistributionPoolCaller{contract: contract}, RewardDistributionPoolTransactor: RewardDistributionPoolTransactor{contract: contract}, RewardDistributionPoolFilterer: RewardDistributionPoolFilterer{contract: contract}}, nil
}

// NewRewardDistributionPoolCaller creates a new read-only instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPoolCaller(address common.Address, caller bind.ContractCaller) (*RewardDistributionPoolCaller, error) {
	contract, err := bindRewardDistributionPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolCaller{contract: contract}, nil
}

// NewRewardDistributionPoolTransactor creates a new write-only instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardDistributionPoolTransactor, error) {
	contract, err := bindRewardDistributionPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolTransactor{contract: contract}, nil
}

// NewRewardDistributionPoolFilterer creates a new log filterer instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardDistributionPoolFilterer, error) {
	contract, err := bindRewardDistributionPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolFilterer{contract: contract}, nil
}

// bindRewardDistributionPool binds a generic wrapper to an already deployed contract.
func bindRewardDistributionPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardDistributionPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributionPool *RewardDistributionPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributionPool.Contract.RewardDistributionPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributionPool *RewardDistributionPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.RewardDistributionPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributionPool *RewardDistributionPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.RewardDistributionPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributionPool *RewardDistributionPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributionPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributionPool *RewardDistributionPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributionPool *RewardDistributionPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.contract.Transact(opts, method, params...)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) DKGKEY() (string, error) {
	return _RewardDistributionPool.Contract.DKGKEY(&_RewardDistributionPool.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) DKGKEY() (string, error) {
	return _RewardDistributionPool.Contract.DKGKEY(&_RewardDistributionPool.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) REWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _RewardDistributionPool.Contract.REWARDDISTRIBUTIONPOOLKEY(&_RewardDistributionPool.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _RewardDistributionPool.Contract.REWARDDISTRIBUTIONPOOLKEY(&_RewardDistributionPool.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) SLASHINGVOTINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.SLASHINGVOTINGKEY(&_RewardDistributionPool.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.SLASHINGVOTINGKEY(&_RewardDistributionPool.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) STAKINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.STAKINGKEY(&_RewardDistributionPool.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) STAKINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.STAKINGKEY(&_RewardDistributionPool.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _RewardDistributionPool.Contract.SUPPORTEDTOKENSKEY(&_RewardDistributionPool.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _RewardDistributionPool.Contract.SUPPORTEDTOKENSKEY(&_RewardDistributionPool.CallOpts)
}

// CollectedRewards is a free data retrieval call binding the contract method 0x8f76c0d8.
//
// Solidity: function collectedRewards() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) CollectedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "collectedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectedRewards is a free data retrieval call binding the contract method 0x8f76c0d8.
//
// Solidity: function collectedRewards() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) CollectedRewards() (*big.Int, error) {
	return _RewardDistributionPool.Contract.CollectedRewards(&_RewardDistributionPool.CallOpts)
}

// CollectedRewards is a free data retrieval call binding the contract method 0x8f76c0d8.
//
// Solidity: function collectedRewards() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) CollectedRewards() (*big.Int, error) {
	return _RewardDistributionPool.Contract.CollectedRewards(&_RewardDistributionPool.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolSession) ContractRegistry() (common.Address, error) {
	return _RewardDistributionPool.Contract.ContractRegistry(&_RewardDistributionPool.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) ContractRegistry() (common.Address, error) {
	return _RewardDistributionPool.Contract.ContractRegistry(&_RewardDistributionPool.CallOpts)
}

// ProvidedStake is a free data retrieval call binding the contract method 0x7cf8b8be.
//
// Solidity: function providedStake() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) ProvidedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "providedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProvidedStake is a free data retrieval call binding the contract method 0x7cf8b8be.
//
// Solidity: function providedStake() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) ProvidedStake() (*big.Int, error) {
	return _RewardDistributionPool.Contract.ProvidedStake(&_RewardDistributionPool.CallOpts)
}

// ProvidedStake is a free data retrieval call binding the contract method 0x7cf8b8be.
//
// Solidity: function providedStake() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) ProvidedStake() (*big.Int, error) {
	return _RewardDistributionPool.Contract.ProvidedStake(&_RewardDistributionPool.CallOpts)
}

// RewardPositions is a free data retrieval call binding the contract method 0xb7b08d4e.
//
// Solidity: function rewardPositions(address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_RewardDistributionPool *RewardDistributionPoolCaller) RewardPositions(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "rewardPositions", arg0)

	outstruct := new(struct {
		Balance          *big.Int
		LastRewardPoints *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardPoints = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardPositions is a free data retrieval call binding the contract method 0xb7b08d4e.
//
// Solidity: function rewardPositions(address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_RewardDistributionPool *RewardDistributionPoolSession) RewardPositions(arg0 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	return _RewardDistributionPool.Contract.RewardPositions(&_RewardDistributionPool.CallOpts, arg0)
}

// RewardPositions is a free data retrieval call binding the contract method 0xb7b08d4e.
//
// Solidity: function rewardPositions(address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) RewardPositions(arg0 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	return _RewardDistributionPool.Contract.RewardPositions(&_RewardDistributionPool.CallOpts, arg0)
}

// RewardsOwing is a free data retrieval call binding the contract method 0xfd4c4aee.
//
// Solidity: function rewardsOwing() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) RewardsOwing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "rewardsOwing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsOwing is a free data retrieval call binding the contract method 0xfd4c4aee.
//
// Solidity: function rewardsOwing() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) RewardsOwing() (*big.Int, error) {
	return _RewardDistributionPool.Contract.RewardsOwing(&_RewardDistributionPool.CallOpts)
}

// RewardsOwing is a free data retrieval call binding the contract method 0xfd4c4aee.
//
// Solidity: function rewardsOwing() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) RewardsOwing() (*big.Int, error) {
	return _RewardDistributionPool.Contract.RewardsOwing(&_RewardDistributionPool.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolSession) SignerGetter() (common.Address, error) {
	return _RewardDistributionPool.Contract.SignerGetter(&_RewardDistributionPool.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) SignerGetter() (common.Address, error) {
	return _RewardDistributionPool.Contract.SignerGetter(&_RewardDistributionPool.CallOpts)
}

// TotalRewardPoints is a free data retrieval call binding the contract method 0xc2e8caf6.
//
// Solidity: function totalRewardPoints() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) TotalRewardPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "totalRewardPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardPoints is a free data retrieval call binding the contract method 0xc2e8caf6.
//
// Solidity: function totalRewardPoints() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) TotalRewardPoints() (*big.Int, error) {
	return _RewardDistributionPool.Contract.TotalRewardPoints(&_RewardDistributionPool.CallOpts)
}

// TotalRewardPoints is a free data retrieval call binding the contract method 0xc2e8caf6.
//
// Solidity: function totalRewardPoints() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) TotalRewardPoints() (*big.Int, error) {
	return _RewardDistributionPool.Contract.TotalRewardPoints(&_RewardDistributionPool.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) ClaimRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ClaimRewards(&_RewardDistributionPool.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ClaimRewards(&_RewardDistributionPool.TransactOpts)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x70bb45b3.
//
// Solidity: function collectRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) CollectRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "collectRewards")
}

// CollectRewards is a paid mutator transaction binding the contract method 0x70bb45b3.
//
// Solidity: function collectRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) CollectRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.CollectRewards(&_RewardDistributionPool.TransactOpts)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x70bb45b3.
//
// Solidity: function collectRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) CollectRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.CollectRewards(&_RewardDistributionPool.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) DistributeRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.DistributeRewards(&_RewardDistributionPool.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.DistributeRewards(&_RewardDistributionPool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _contractRegistry, address _signerGetterAddress) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) Initialize(opts *bind.TransactOpts, _contractRegistry common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "initialize", _contractRegistry, _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _contractRegistry, address _signerGetterAddress) returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) Initialize(_contractRegistry common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Initialize(&_RewardDistributionPool.TransactOpts, _contractRegistry, _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _contractRegistry, address _signerGetterAddress) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) Initialize(_contractRegistry common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Initialize(&_RewardDistributionPool.TransactOpts, _contractRegistry, _signerGetterAddress)
}

// ReinvestRewards is a paid mutator transaction binding the contract method 0xb5f4d38c.
//
// Solidity: function reinvestRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) ReinvestRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "reinvestRewards")
}

// ReinvestRewards is a paid mutator transaction binding the contract method 0xb5f4d38c.
//
// Solidity: function reinvestRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) ReinvestRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ReinvestRewards(&_RewardDistributionPool.TransactOpts)
}

// ReinvestRewards is a paid mutator transaction binding the contract method 0xb5f4d38c.
//
// Solidity: function reinvestRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) ReinvestRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ReinvestRewards(&_RewardDistributionPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) Receive() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Receive(&_RewardDistributionPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Receive(&_RewardDistributionPool.TransactOpts)
}

// RewardDistributionPoolCollectRewardsIterator is returned from FilterCollectRewards and is used to iterate over the raw logs and unpacked data for CollectRewards events raised by the RewardDistributionPool contract.
type RewardDistributionPoolCollectRewardsIterator struct {
	Event *RewardDistributionPoolCollectRewards // Event containing the contract specifics and raw log

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
func (it *RewardDistributionPoolCollectRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDistributionPoolCollectRewards)
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
		it.Event = new(RewardDistributionPoolCollectRewards)
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
func (it *RewardDistributionPoolCollectRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDistributionPoolCollectRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDistributionPoolCollectRewards represents a CollectRewards event raised by the RewardDistributionPool contract.
type RewardDistributionPoolCollectRewards struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectRewards is a free log retrieval operation binding the contract event 0x2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a717.
//
// Solidity: event CollectRewards(address validator, uint256 amount)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) FilterCollectRewards(opts *bind.FilterOpts) (*RewardDistributionPoolCollectRewardsIterator, error) {

	logs, sub, err := _RewardDistributionPool.contract.FilterLogs(opts, "CollectRewards")
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolCollectRewardsIterator{contract: _RewardDistributionPool.contract, event: "CollectRewards", logs: logs, sub: sub}, nil
}

// WatchCollectRewards is a free log subscription operation binding the contract event 0x2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a717.
//
// Solidity: event CollectRewards(address validator, uint256 amount)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) WatchCollectRewards(opts *bind.WatchOpts, sink chan<- *RewardDistributionPoolCollectRewards) (event.Subscription, error) {

	logs, sub, err := _RewardDistributionPool.contract.WatchLogs(opts, "CollectRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDistributionPoolCollectRewards)
				if err := _RewardDistributionPool.contract.UnpackLog(event, "CollectRewards", log); err != nil {
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

// ParseCollectRewards is a log parse operation binding the contract event 0x2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a717.
//
// Solidity: event CollectRewards(address validator, uint256 amount)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) ParseCollectRewards(log types.Log) (*RewardDistributionPoolCollectRewards, error) {
	event := new(RewardDistributionPoolCollectRewards)
	if err := _RewardDistributionPool.contract.UnpackLog(event, "CollectRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardDistributionPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RewardDistributionPool contract.
type RewardDistributionPoolInitializedIterator struct {
	Event *RewardDistributionPoolInitialized // Event containing the contract specifics and raw log

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
func (it *RewardDistributionPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDistributionPoolInitialized)
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
		it.Event = new(RewardDistributionPoolInitialized)
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
func (it *RewardDistributionPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDistributionPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDistributionPoolInitialized represents a Initialized event raised by the RewardDistributionPool contract.
type RewardDistributionPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*RewardDistributionPoolInitializedIterator, error) {

	logs, sub, err := _RewardDistributionPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolInitializedIterator{contract: _RewardDistributionPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RewardDistributionPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _RewardDistributionPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDistributionPoolInitialized)
				if err := _RewardDistributionPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) ParseInitialized(log types.Log) (*RewardDistributionPoolInitialized, error) {
	event := new(RewardDistributionPoolInitialized)
	if err := _RewardDistributionPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
