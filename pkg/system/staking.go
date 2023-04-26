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

// StakingValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type StakingValidatorInfo struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractRegistry\",\"type\":\"address\"}],\"name\":\"ContractRegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"}],\"name\":\"MinimalStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"WithdrawalPeriodUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EVENT_REGISTRY_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addRewardsToStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressStorage\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"announceWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorStorage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"listValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structStaking.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"}],\"name\":\"setMinimalStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractSignerGetter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawalAnnouncements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_Staking *StakingCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_Staking *StakingSession) DKGKEY() (string, error) {
	return _Staking.Contract.DKGKEY(&_Staking.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_Staking *StakingCallerSession) DKGKEY() (string, error) {
	return _Staking.Contract.DKGKEY(&_Staking.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_Staking *StakingCaller) EVENTREGISTRYKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "EVENT_REGISTRY_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_Staking *StakingSession) EVENTREGISTRYKEY() (string, error) {
	return _Staking.Contract.EVENTREGISTRYKEY(&_Staking.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_Staking *StakingCallerSession) EVENTREGISTRYKEY() (string, error) {
	return _Staking.Contract.EVENTREGISTRYKEY(&_Staking.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_Staking *StakingCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_Staking *StakingSession) SLASHINGVOTINGKEY() (string, error) {
	return _Staking.Contract.SLASHINGVOTINGKEY(&_Staking.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_Staking *StakingCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _Staking.Contract.SLASHINGVOTINGKEY(&_Staking.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_Staking *StakingCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_Staking *StakingSession) STAKINGKEY() (string, error) {
	return _Staking.Contract.STAKINGKEY(&_Staking.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_Staking *StakingCallerSession) STAKINGKEY() (string, error) {
	return _Staking.Contract.STAKINGKEY(&_Staking.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_Staking *StakingCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_Staking *StakingSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _Staking.Contract.SUPPORTEDTOKENSKEY(&_Staking.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_Staking *StakingCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _Staking.Contract.SUPPORTEDTOKENSKEY(&_Staking.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_Staking *StakingCaller) VALIDATORREWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_Staking *StakingSession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _Staking.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_Staking.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_Staking *StakingCallerSession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _Staking.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_Staking.CallOpts)
}

// AddressStorage is a free data retrieval call binding the contract method 0x612e1488.
//
// Solidity: function addressStorage() view returns(address)
func (_Staking *StakingCaller) AddressStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "addressStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressStorage is a free data retrieval call binding the contract method 0x612e1488.
//
// Solidity: function addressStorage() view returns(address)
func (_Staking *StakingSession) AddressStorage() (common.Address, error) {
	return _Staking.Contract.AddressStorage(&_Staking.CallOpts)
}

// AddressStorage is a free data retrieval call binding the contract method 0x612e1488.
//
// Solidity: function addressStorage() view returns(address)
func (_Staking *StakingCallerSession) AddressStorage() (common.Address, error) {
	return _Staking.Contract.AddressStorage(&_Staking.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_Staking *StakingCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_Staking *StakingSession) ContractRegistry() (common.Address, error) {
	return _Staking.Contract.ContractRegistry(&_Staking.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_Staking *StakingCallerSession) ContractRegistry() (common.Address, error) {
	return _Staking.Contract.ContractRegistry(&_Staking.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_Staking *StakingCaller) GetStake(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStake", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_Staking *StakingSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _Staking.Contract.GetStake(&_Staking.CallOpts, _validator)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_Staking *StakingCallerSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _Staking.Contract.GetStake(&_Staking.CallOpts, _validator)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Staking *StakingCaller) GetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getTotalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Staking *StakingSession) GetTotalStake() (*big.Int, error) {
	return _Staking.Contract.GetTotalStake(&_Staking.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Staking *StakingCallerSession) GetTotalStake() (*big.Int, error) {
	return _Staking.Contract.GetTotalStake(&_Staking.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingSession) GetValidators() ([]common.Address, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingCallerSession) GetValidators() ([]common.Address, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Staking *StakingCaller) GetValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Staking *StakingSession) GetValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetValidatorsCount(&_Staking.CallOpts)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Staking *StakingCallerSession) GetValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetValidatorsCount(&_Staking.CallOpts)
}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address _validator) view returns(bool)
func (_Staking *StakingCaller) IsValidatorActive(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isValidatorActive", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address _validator) view returns(bool)
func (_Staking *StakingSession) IsValidatorActive(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorActive(&_Staking.CallOpts, _validator)
}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address _validator) view returns(bool)
func (_Staking *StakingCallerSession) IsValidatorActive(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorActive(&_Staking.CallOpts, _validator)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validator) view returns(bool)
func (_Staking *StakingCaller) IsValidatorSlashed(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isValidatorSlashed", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validator) view returns(bool)
func (_Staking *StakingSession) IsValidatorSlashed(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorSlashed(&_Staking.CallOpts, _validator)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validator) view returns(bool)
func (_Staking *StakingCallerSession) IsValidatorSlashed(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorSlashed(&_Staking.CallOpts, _validator)
}

// ListValidators is a free data retrieval call binding the contract method 0x1c78cb38.
//
// Solidity: function listValidators(uint256 _offset, uint256 _limit) view returns((address,uint256,uint8)[])
func (_Staking *StakingCaller) ListValidators(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]StakingValidatorInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "listValidators", _offset, _limit)

	if err != nil {
		return *new([]StakingValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingValidatorInfo)).(*[]StakingValidatorInfo)

	return out0, err

}

// ListValidators is a free data retrieval call binding the contract method 0x1c78cb38.
//
// Solidity: function listValidators(uint256 _offset, uint256 _limit) view returns((address,uint256,uint8)[])
func (_Staking *StakingSession) ListValidators(_offset *big.Int, _limit *big.Int) ([]StakingValidatorInfo, error) {
	return _Staking.Contract.ListValidators(&_Staking.CallOpts, _offset, _limit)
}

// ListValidators is a free data retrieval call binding the contract method 0x1c78cb38.
//
// Solidity: function listValidators(uint256 _offset, uint256 _limit) view returns((address,uint256,uint8)[])
func (_Staking *StakingCallerSession) ListValidators(_offset *big.Int, _limit *big.Int) ([]StakingValidatorInfo, error) {
	return _Staking.Contract.ListValidators(&_Staking.CallOpts, _offset, _limit)
}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_Staking *StakingCaller) MinimalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "minimalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_Staking *StakingSession) MinimalStake() (*big.Int, error) {
	return _Staking.Contract.MinimalStake(&_Staking.CallOpts)
}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_Staking *StakingCallerSession) MinimalStake() (*big.Int, error) {
	return _Staking.Contract.MinimalStake(&_Staking.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_Staking *StakingCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_Staking *StakingSession) SignerGetter() (common.Address, error) {
	return _Staking.Contract.SignerGetter(&_Staking.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_Staking *StakingCallerSession) SignerGetter() (common.Address, error) {
	return _Staking.Contract.SignerGetter(&_Staking.CallOpts)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_Staking *StakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Validator common.Address
		Stake     *big.Int
		Status    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Stake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_Staking *StakingSession) Stakes(arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_Staking *StakingCallerSession) Stakes(arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Staking *StakingCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Staking *StakingSession) TotalStake() (*big.Int, error) {
	return _Staking.Contract.TotalStake(&_Staking.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStake() (*big.Int, error) {
	return _Staking.Contract.TotalStake(&_Staking.CallOpts)
}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_Staking *StakingCaller) WithdrawalAnnouncements(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawalAnnouncements", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Time   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_Staking *StakingSession) WithdrawalAnnouncements(arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _Staking.Contract.WithdrawalAnnouncements(&_Staking.CallOpts, arg0)
}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_Staking *StakingCallerSession) WithdrawalAnnouncements(arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _Staking.Contract.WithdrawalAnnouncements(&_Staking.CallOpts, arg0)
}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_Staking *StakingCaller) WithdrawalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawalPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_Staking *StakingSession) WithdrawalPeriod() (*big.Int, error) {
	return _Staking.Contract.WithdrawalPeriod(&_Staking.CallOpts)
}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawalPeriod() (*big.Int, error) {
	return _Staking.Contract.WithdrawalPeriod(&_Staking.CallOpts)
}

// AddRewardsToStake is a paid mutator transaction binding the contract method 0xf3840328.
//
// Solidity: function addRewardsToStake(address _validator, uint256 _amount) returns()
func (_Staking *StakingTransactor) AddRewardsToStake(opts *bind.TransactOpts, _validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addRewardsToStake", _validator, _amount)
}

// AddRewardsToStake is a paid mutator transaction binding the contract method 0xf3840328.
//
// Solidity: function addRewardsToStake(address _validator, uint256 _amount) returns()
func (_Staking *StakingSession) AddRewardsToStake(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AddRewardsToStake(&_Staking.TransactOpts, _validator, _amount)
}

// AddRewardsToStake is a paid mutator transaction binding the contract method 0xf3840328.
//
// Solidity: function addRewardsToStake(address _validator, uint256 _amount) returns()
func (_Staking *StakingTransactorSession) AddRewardsToStake(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AddRewardsToStake(&_Staking.TransactOpts, _validator, _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_Staking *StakingTransactor) AnnounceWithdrawal(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "announceWithdrawal", _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_Staking *StakingSession) AnnounceWithdrawal(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AnnounceWithdrawal(&_Staking.TransactOpts, _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) AnnounceWithdrawal(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AnnounceWithdrawal(&_Staking.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _signerGetterAddress, uint256 _minimalStake, uint256 _withdrawalPeriod, address _contractRegistry, address _validatorStorage) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _signerGetterAddress common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int, _contractRegistry common.Address, _validatorStorage common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _signerGetterAddress, _minimalStake, _withdrawalPeriod, _contractRegistry, _validatorStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _signerGetterAddress, uint256 _minimalStake, uint256 _withdrawalPeriod, address _contractRegistry, address _validatorStorage) returns()
func (_Staking *StakingSession) Initialize(_signerGetterAddress common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int, _contractRegistry common.Address, _validatorStorage common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _signerGetterAddress, _minimalStake, _withdrawalPeriod, _contractRegistry, _validatorStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _signerGetterAddress, uint256 _minimalStake, uint256 _withdrawalPeriod, address _contractRegistry, address _validatorStorage) returns()
func (_Staking *StakingTransactorSession) Initialize(_signerGetterAddress common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int, _contractRegistry common.Address, _validatorStorage common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _signerGetterAddress, _minimalStake, _withdrawalPeriod, _contractRegistry, _validatorStorage)
}

// RevokeWithdrawal is a paid mutator transaction binding the contract method 0x35c14de1.
//
// Solidity: function revokeWithdrawal() returns()
func (_Staking *StakingTransactor) RevokeWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "revokeWithdrawal")
}

// RevokeWithdrawal is a paid mutator transaction binding the contract method 0x35c14de1.
//
// Solidity: function revokeWithdrawal() returns()
func (_Staking *StakingSession) RevokeWithdrawal() (*types.Transaction, error) {
	return _Staking.Contract.RevokeWithdrawal(&_Staking.TransactOpts)
}

// RevokeWithdrawal is a paid mutator transaction binding the contract method 0x35c14de1.
//
// Solidity: function revokeWithdrawal() returns()
func (_Staking *StakingTransactorSession) RevokeWithdrawal() (*types.Transaction, error) {
	return _Staking.Contract.RevokeWithdrawal(&_Staking.TransactOpts)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_Staking *StakingTransactor) SetMinimalStake(opts *bind.TransactOpts, _minimalStake *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMinimalStake", _minimalStake)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_Staking *StakingSession) SetMinimalStake(_minimalStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinimalStake(&_Staking.TransactOpts, _minimalStake)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_Staking *StakingTransactorSession) SetMinimalStake(_minimalStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinimalStake(&_Staking.TransactOpts, _minimalStake)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactor) SetWithdrawalPeriod(opts *bind.TransactOpts, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setWithdrawalPeriod", _withdrawalPeriod)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_Staking *StakingSession) SetWithdrawalPeriod(_withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetWithdrawalPeriod(&_Staking.TransactOpts, _withdrawalPeriod)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactorSession) SetWithdrawalPeriod(_withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetWithdrawalPeriod(&_Staking.TransactOpts, _withdrawalPeriod)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", _validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_Staking *StakingSession) Slash(_validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_Staking *StakingTransactorSession) Slash(_validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _validator)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactorSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingSession) Withdraw() (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingContractRegistryUpdatedIterator is returned from FilterContractRegistryUpdated and is used to iterate over the raw logs and unpacked data for ContractRegistryUpdated events raised by the Staking contract.
type StakingContractRegistryUpdatedIterator struct {
	Event *StakingContractRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *StakingContractRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRegistryUpdated)
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
		it.Event = new(StakingContractRegistryUpdated)
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
func (it *StakingContractRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRegistryUpdated represents a ContractRegistryUpdated event raised by the Staking contract.
type StakingContractRegistryUpdated struct {
	ContractRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterContractRegistryUpdated is a free log retrieval operation binding the contract event 0xb3a6e7c81ebdb0cf9bf28d5ddf6678ded8d73b44f2eee06ca3974dbb9f41ce7e.
//
// Solidity: event ContractRegistryUpdated(address contractRegistry)
func (_Staking *StakingFilterer) FilterContractRegistryUpdated(opts *bind.FilterOpts) (*StakingContractRegistryUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ContractRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingContractRegistryUpdatedIterator{contract: _Staking.contract, event: "ContractRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchContractRegistryUpdated is a free log subscription operation binding the contract event 0xb3a6e7c81ebdb0cf9bf28d5ddf6678ded8d73b44f2eee06ca3974dbb9f41ce7e.
//
// Solidity: event ContractRegistryUpdated(address contractRegistry)
func (_Staking *StakingFilterer) WatchContractRegistryUpdated(opts *bind.WatchOpts, sink chan<- *StakingContractRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ContractRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRegistryUpdated)
				if err := _Staking.contract.UnpackLog(event, "ContractRegistryUpdated", log); err != nil {
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

// ParseContractRegistryUpdated is a log parse operation binding the contract event 0xb3a6e7c81ebdb0cf9bf28d5ddf6678ded8d73b44f2eee06ca3974dbb9f41ce7e.
//
// Solidity: event ContractRegistryUpdated(address contractRegistry)
func (_Staking *StakingFilterer) ParseContractRegistryUpdated(log types.Log) (*StakingContractRegistryUpdated, error) {
	event := new(StakingContractRegistryUpdated)
	if err := _Staking.contract.UnpackLog(event, "ContractRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingMinimalStakeUpdatedIterator is returned from FilterMinimalStakeUpdated and is used to iterate over the raw logs and unpacked data for MinimalStakeUpdated events raised by the Staking contract.
type StakingMinimalStakeUpdatedIterator struct {
	Event *StakingMinimalStakeUpdated // Event containing the contract specifics and raw log

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
func (it *StakingMinimalStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingMinimalStakeUpdated)
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
		it.Event = new(StakingMinimalStakeUpdated)
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
func (it *StakingMinimalStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingMinimalStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingMinimalStakeUpdated represents a MinimalStakeUpdated event raised by the Staking contract.
type StakingMinimalStakeUpdated struct {
	MinimalStake *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinimalStakeUpdated is a free log retrieval operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_Staking *StakingFilterer) FilterMinimalStakeUpdated(opts *bind.FilterOpts) (*StakingMinimalStakeUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "MinimalStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingMinimalStakeUpdatedIterator{contract: _Staking.contract, event: "MinimalStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimalStakeUpdated is a free log subscription operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_Staking *StakingFilterer) WatchMinimalStakeUpdated(opts *bind.WatchOpts, sink chan<- *StakingMinimalStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "MinimalStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingMinimalStakeUpdated)
				if err := _Staking.contract.UnpackLog(event, "MinimalStakeUpdated", log); err != nil {
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

// ParseMinimalStakeUpdated is a log parse operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_Staking *StakingFilterer) ParseMinimalStakeUpdated(log types.Log) (*StakingMinimalStakeUpdated, error) {
	event := new(StakingMinimalStakeUpdated)
	if err := _Staking.contract.UnpackLog(event, "MinimalStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawalPeriodUpdatedIterator is returned from FilterWithdrawalPeriodUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalPeriodUpdated events raised by the Staking contract.
type StakingWithdrawalPeriodUpdatedIterator struct {
	Event *StakingWithdrawalPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawalPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawalPeriodUpdated)
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
		it.Event = new(StakingWithdrawalPeriodUpdated)
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
func (it *StakingWithdrawalPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawalPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawalPeriodUpdated represents a WithdrawalPeriodUpdated event raised by the Staking contract.
type StakingWithdrawalPeriodUpdated struct {
	WithdrawalPeriod *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalPeriodUpdated is a free log retrieval operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_Staking *StakingFilterer) FilterWithdrawalPeriodUpdated(opts *bind.FilterOpts) (*StakingWithdrawalPeriodUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawalPeriodUpdatedIterator{contract: _Staking.contract, event: "WithdrawalPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalPeriodUpdated is a free log subscription operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_Staking *StakingFilterer) WatchWithdrawalPeriodUpdated(opts *bind.WatchOpts, sink chan<- *StakingWithdrawalPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawalPeriodUpdated)
				if err := _Staking.contract.UnpackLog(event, "WithdrawalPeriodUpdated", log); err != nil {
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

// ParseWithdrawalPeriodUpdated is a log parse operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_Staking *StakingFilterer) ParseWithdrawalPeriodUpdated(log types.Log) (*StakingWithdrawalPeriodUpdated, error) {
	event := new(StakingWithdrawalPeriodUpdated)
	if err := _Staking.contract.UnpackLog(event, "WithdrawalPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
