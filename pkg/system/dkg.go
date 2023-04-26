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

// DKGMetaData contains all meta data concerning the DKG contract.
var DKGMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"RoundDataFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"RoundDataProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectiveSigner\",\"type\":\"address\"}],\"name\":\"SignerVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ThresholdSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EVENT_REGISTRY_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlinePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"generations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGenerationsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getRoundBroadcastCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getRoundBroadcastData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumDKG.GenerationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadlinePeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isCurrentValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"isRoundFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastActiveGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_rawData\",\"type\":\"bytes\"}],\"name\":\"roundBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deadlinePeriod\",\"type\":\"uint256\"}],\"name\":\"setDeadlinePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGeneration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"voteSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DKGABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGMetaData.ABI instead.
var DKGABI = DKGMetaData.ABI

// DKG is an auto generated Go binding around an Ethereum contract.
type DKG struct {
	DKGCaller     // Read-only binding to the contract
	DKGTransactor // Write-only binding to the contract
	DKGFilterer   // Log filterer for contract events
}

// DKGCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGSession struct {
	Contract     *DKG              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGCallerSession struct {
	Contract *DKGCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DKGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGTransactorSession struct {
	Contract     *DKGTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGRaw struct {
	Contract *DKG // Generic contract binding to access the raw methods on
}

// DKGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGCallerRaw struct {
	Contract *DKGCaller // Generic read-only contract binding to access the raw methods on
}

// DKGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGTransactorRaw struct {
	Contract *DKGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKG creates a new instance of DKG, bound to a specific deployed contract.
func NewDKG(address common.Address, backend bind.ContractBackend) (*DKG, error) {
	contract, err := bindDKG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKG{DKGCaller: DKGCaller{contract: contract}, DKGTransactor: DKGTransactor{contract: contract}, DKGFilterer: DKGFilterer{contract: contract}}, nil
}

// NewDKGCaller creates a new read-only instance of DKG, bound to a specific deployed contract.
func NewDKGCaller(address common.Address, caller bind.ContractCaller) (*DKGCaller, error) {
	contract, err := bindDKG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGCaller{contract: contract}, nil
}

// NewDKGTransactor creates a new write-only instance of DKG, bound to a specific deployed contract.
func NewDKGTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGTransactor, error) {
	contract, err := bindDKG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGTransactor{contract: contract}, nil
}

// NewDKGFilterer creates a new log filterer instance of DKG, bound to a specific deployed contract.
func NewDKGFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGFilterer, error) {
	contract, err := bindDKG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGFilterer{contract: contract}, nil
}

// bindDKG binds a generic wrapper to an already deployed contract.
func bindDKG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DKGABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.DKGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transact(opts, method, params...)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_DKG *DKGCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_DKG *DKGSession) DKGKEY() (string, error) {
	return _DKG.Contract.DKGKEY(&_DKG.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_DKG *DKGCallerSession) DKGKEY() (string, error) {
	return _DKG.Contract.DKGKEY(&_DKG.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_DKG *DKGCaller) EVENTREGISTRYKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "EVENT_REGISTRY_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_DKG *DKGSession) EVENTREGISTRYKEY() (string, error) {
	return _DKG.Contract.EVENTREGISTRYKEY(&_DKG.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_DKG *DKGCallerSession) EVENTREGISTRYKEY() (string, error) {
	return _DKG.Contract.EVENTREGISTRYKEY(&_DKG.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_DKG *DKGCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_DKG *DKGSession) SLASHINGVOTINGKEY() (string, error) {
	return _DKG.Contract.SLASHINGVOTINGKEY(&_DKG.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_DKG *DKGCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _DKG.Contract.SLASHINGVOTINGKEY(&_DKG.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_DKG *DKGCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_DKG *DKGSession) STAKINGKEY() (string, error) {
	return _DKG.Contract.STAKINGKEY(&_DKG.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_DKG *DKGCallerSession) STAKINGKEY() (string, error) {
	return _DKG.Contract.STAKINGKEY(&_DKG.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_DKG *DKGCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_DKG *DKGSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _DKG.Contract.SUPPORTEDTOKENSKEY(&_DKG.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_DKG *DKGCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _DKG.Contract.SUPPORTEDTOKENSKEY(&_DKG.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_DKG *DKGCaller) VALIDATORREWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_DKG *DKGSession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _DKG.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_DKG.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_DKG *DKGCallerSession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _DKG.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_DKG.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_DKG *DKGCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_DKG *DKGSession) ContractRegistry() (common.Address, error) {
	return _DKG.Contract.ContractRegistry(&_DKG.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_DKG *DKGCallerSession) ContractRegistry() (common.Address, error) {
	return _DKG.Contract.ContractRegistry(&_DKG.CallOpts)
}

// DeadlinePeriod is a free data retrieval call binding the contract method 0x6db24262.
//
// Solidity: function deadlinePeriod() view returns(uint256)
func (_DKG *DKGCaller) DeadlinePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "deadlinePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlinePeriod is a free data retrieval call binding the contract method 0x6db24262.
//
// Solidity: function deadlinePeriod() view returns(uint256)
func (_DKG *DKGSession) DeadlinePeriod() (*big.Int, error) {
	return _DKG.Contract.DeadlinePeriod(&_DKG.CallOpts)
}

// DeadlinePeriod is a free data retrieval call binding the contract method 0x6db24262.
//
// Solidity: function deadlinePeriod() view returns(uint256)
func (_DKG *DKGCallerSession) DeadlinePeriod() (*big.Int, error) {
	return _DKG.Contract.DeadlinePeriod(&_DKG.CallOpts)
}

// Generations is a free data retrieval call binding the contract method 0xad8b0e94.
//
// Solidity: function generations(uint256 ) view returns(address signer, uint256 deadline)
func (_DKG *DKGCaller) Generations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Signer   common.Address
	Deadline *big.Int
}, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "generations", arg0)

	outstruct := new(struct {
		Signer   common.Address
		Deadline *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deadline = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Generations is a free data retrieval call binding the contract method 0xad8b0e94.
//
// Solidity: function generations(uint256 ) view returns(address signer, uint256 deadline)
func (_DKG *DKGSession) Generations(arg0 *big.Int) (struct {
	Signer   common.Address
	Deadline *big.Int
}, error) {
	return _DKG.Contract.Generations(&_DKG.CallOpts, arg0)
}

// Generations is a free data retrieval call binding the contract method 0xad8b0e94.
//
// Solidity: function generations(uint256 ) view returns(address signer, uint256 deadline)
func (_DKG *DKGCallerSession) Generations(arg0 *big.Int) (struct {
	Signer   common.Address
	Deadline *big.Int
}, error) {
	return _DKG.Contract.Generations(&_DKG.CallOpts, arg0)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGCaller) GetCurrentValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getCurrentValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGSession) GetCurrentValidators() ([]common.Address, error) {
	return _DKG.Contract.GetCurrentValidators(&_DKG.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGCallerSession) GetCurrentValidators() ([]common.Address, error) {
	return _DKG.Contract.GetCurrentValidators(&_DKG.CallOpts)
}

// GetGenerationsCount is a free data retrieval call binding the contract method 0x1ea0f036.
//
// Solidity: function getGenerationsCount() view returns(uint256)
func (_DKG *DKGCaller) GetGenerationsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getGenerationsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGenerationsCount is a free data retrieval call binding the contract method 0x1ea0f036.
//
// Solidity: function getGenerationsCount() view returns(uint256)
func (_DKG *DKGSession) GetGenerationsCount() (*big.Int, error) {
	return _DKG.Contract.GetGenerationsCount(&_DKG.CallOpts)
}

// GetGenerationsCount is a free data retrieval call binding the contract method 0x1ea0f036.
//
// Solidity: function getGenerationsCount() view returns(uint256)
func (_DKG *DKGCallerSession) GetGenerationsCount() (*big.Int, error) {
	return _DKG.Contract.GetGenerationsCount(&_DKG.CallOpts)
}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _generation, uint256 _round) view returns(uint256)
func (_DKG *DKGCaller) GetRoundBroadcastCount(opts *bind.CallOpts, _generation *big.Int, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRoundBroadcastCount", _generation, _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _generation, uint256 _round) view returns(uint256)
func (_DKG *DKGSession) GetRoundBroadcastCount(_generation *big.Int, _round *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetRoundBroadcastCount(&_DKG.CallOpts, _generation, _round)
}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _generation, uint256 _round) view returns(uint256)
func (_DKG *DKGCallerSession) GetRoundBroadcastCount(_generation *big.Int, _round *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetRoundBroadcastCount(&_DKG.CallOpts, _generation, _round)
}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _generation, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGCaller) GetRoundBroadcastData(opts *bind.CallOpts, _generation *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRoundBroadcastData", _generation, _round, _validator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _generation, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGSession) GetRoundBroadcastData(_generation *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	return _DKG.Contract.GetRoundBroadcastData(&_DKG.CallOpts, _generation, _round, _validator)
}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _generation, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGCallerSession) GetRoundBroadcastData(_generation *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	return _DKG.Contract.GetRoundBroadcastData(&_DKG.CallOpts, _generation, _round, _validator)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_DKG *DKGCaller) GetSignerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getSignerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_DKG *DKGSession) GetSignerAddress() (common.Address, error) {
	return _DKG.Contract.GetSignerAddress(&_DKG.CallOpts)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_DKG *DKGCallerSession) GetSignerAddress() (common.Address, error) {
	return _DKG.Contract.GetSignerAddress(&_DKG.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _generation) view returns(uint8)
func (_DKG *DKGCaller) GetStatus(opts *bind.CallOpts, _generation *big.Int) (uint8, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getStatus", _generation)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _generation) view returns(uint8)
func (_DKG *DKGSession) GetStatus(_generation *big.Int) (uint8, error) {
	return _DKG.Contract.GetStatus(&_DKG.CallOpts, _generation)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _generation) view returns(uint8)
func (_DKG *DKGCallerSession) GetStatus(_generation *big.Int) (uint8, error) {
	return _DKG.Contract.GetStatus(&_DKG.CallOpts, _generation)
}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _generation) view returns(address[])
func (_DKG *DKGCaller) GetValidators(opts *bind.CallOpts, _generation *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidators", _generation)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _generation) view returns(address[])
func (_DKG *DKGSession) GetValidators(_generation *big.Int) ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts, _generation)
}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _generation) view returns(address[])
func (_DKG *DKGCallerSession) GetValidators(_generation *big.Int) ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts, _generation)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0xcd5fcd15.
//
// Solidity: function getValidatorsCount(uint256 _generation) view returns(uint256)
func (_DKG *DKGCaller) GetValidatorsCount(opts *bind.CallOpts, _generation *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidatorsCount", _generation)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCount is a free data retrieval call binding the contract method 0xcd5fcd15.
//
// Solidity: function getValidatorsCount(uint256 _generation) view returns(uint256)
func (_DKG *DKGSession) GetValidatorsCount(_generation *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetValidatorsCount(&_DKG.CallOpts, _generation)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0xcd5fcd15.
//
// Solidity: function getValidatorsCount(uint256 _generation) view returns(uint256)
func (_DKG *DKGCallerSession) GetValidatorsCount(_generation *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetValidatorsCount(&_DKG.CallOpts, _generation)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address _validator) view returns(bool)
func (_DKG *DKGCaller) IsCurrentValidator(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isCurrentValidator", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address _validator) view returns(bool)
func (_DKG *DKGSession) IsCurrentValidator(_validator common.Address) (bool, error) {
	return _DKG.Contract.IsCurrentValidator(&_DKG.CallOpts, _validator)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address _validator) view returns(bool)
func (_DKG *DKGCallerSession) IsCurrentValidator(_validator common.Address) (bool, error) {
	return _DKG.Contract.IsCurrentValidator(&_DKG.CallOpts, _validator)
}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _generation, uint256 _round) view returns(bool)
func (_DKG *DKGCaller) IsRoundFilled(opts *bind.CallOpts, _generation *big.Int, _round *big.Int) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isRoundFilled", _generation, _round)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _generation, uint256 _round) view returns(bool)
func (_DKG *DKGSession) IsRoundFilled(_generation *big.Int, _round *big.Int) (bool, error) {
	return _DKG.Contract.IsRoundFilled(&_DKG.CallOpts, _generation, _round)
}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _generation, uint256 _round) view returns(bool)
func (_DKG *DKGCallerSession) IsRoundFilled(_generation *big.Int, _round *big.Int) (bool, error) {
	return _DKG.Contract.IsRoundFilled(&_DKG.CallOpts, _generation, _round)
}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 _generation, address _validator) view returns(bool)
func (_DKG *DKGCaller) IsValidator(opts *bind.CallOpts, _generation *big.Int, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isValidator", _generation, _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 _generation, address _validator) view returns(bool)
func (_DKG *DKGSession) IsValidator(_generation *big.Int, _validator common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, _generation, _validator)
}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 _generation, address _validator) view returns(bool)
func (_DKG *DKGCallerSession) IsValidator(_generation *big.Int, _validator common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, _generation, _validator)
}

// LastActiveGeneration is a free data retrieval call binding the contract method 0x11af0a20.
//
// Solidity: function lastActiveGeneration() view returns(uint256)
func (_DKG *DKGCaller) LastActiveGeneration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "lastActiveGeneration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActiveGeneration is a free data retrieval call binding the contract method 0x11af0a20.
//
// Solidity: function lastActiveGeneration() view returns(uint256)
func (_DKG *DKGSession) LastActiveGeneration() (*big.Int, error) {
	return _DKG.Contract.LastActiveGeneration(&_DKG.CallOpts)
}

// LastActiveGeneration is a free data retrieval call binding the contract method 0x11af0a20.
//
// Solidity: function lastActiveGeneration() view returns(uint256)
func (_DKG *DKGCallerSession) LastActiveGeneration() (*big.Int, error) {
	return _DKG.Contract.LastActiveGeneration(&_DKG.CallOpts)
}

// SignerToGeneration is a free data retrieval call binding the contract method 0xad51db6a.
//
// Solidity: function signerToGeneration(address ) view returns(uint256)
func (_DKG *DKGCaller) SignerToGeneration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "signerToGeneration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerToGeneration is a free data retrieval call binding the contract method 0xad51db6a.
//
// Solidity: function signerToGeneration(address ) view returns(uint256)
func (_DKG *DKGSession) SignerToGeneration(arg0 common.Address) (*big.Int, error) {
	return _DKG.Contract.SignerToGeneration(&_DKG.CallOpts, arg0)
}

// SignerToGeneration is a free data retrieval call binding the contract method 0xad51db6a.
//
// Solidity: function signerToGeneration(address ) view returns(uint256)
func (_DKG *DKGCallerSession) SignerToGeneration(arg0 common.Address) (*big.Int, error) {
	return _DKG.Contract.SignerToGeneration(&_DKG.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _contractRegistry, uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactor) Initialize(opts *bind.TransactOpts, _contractRegistry common.Address, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "initialize", _contractRegistry, _deadlinePeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _contractRegistry, uint256 _deadlinePeriod) returns()
func (_DKG *DKGSession) Initialize(_contractRegistry common.Address, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _contractRegistry, _deadlinePeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _contractRegistry, uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactorSession) Initialize(_contractRegistry common.Address, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _contractRegistry, _deadlinePeriod)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _generation, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGTransactor) RoundBroadcast(opts *bind.TransactOpts, _generation *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "roundBroadcast", _generation, _round, _rawData)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _generation, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGSession) RoundBroadcast(_generation *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.Contract.RoundBroadcast(&_DKG.TransactOpts, _generation, _round, _rawData)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _generation, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGTransactorSession) RoundBroadcast(_generation *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.Contract.RoundBroadcast(&_DKG.TransactOpts, _generation, _round, _rawData)
}

// SetDeadlinePeriod is a paid mutator transaction binding the contract method 0x82651c0d.
//
// Solidity: function setDeadlinePeriod(uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactor) SetDeadlinePeriod(opts *bind.TransactOpts, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "setDeadlinePeriod", _deadlinePeriod)
}

// SetDeadlinePeriod is a paid mutator transaction binding the contract method 0x82651c0d.
//
// Solidity: function setDeadlinePeriod(uint256 _deadlinePeriod) returns()
func (_DKG *DKGSession) SetDeadlinePeriod(_deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.SetDeadlinePeriod(&_DKG.TransactOpts, _deadlinePeriod)
}

// SetDeadlinePeriod is a paid mutator transaction binding the contract method 0x82651c0d.
//
// Solidity: function setDeadlinePeriod(uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactorSession) SetDeadlinePeriod(_deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.SetDeadlinePeriod(&_DKG.TransactOpts, _deadlinePeriod)
}

// UpdateGeneration is a paid mutator transaction binding the contract method 0xb32805c3.
//
// Solidity: function updateGeneration() returns()
func (_DKG *DKGTransactor) UpdateGeneration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "updateGeneration")
}

// UpdateGeneration is a paid mutator transaction binding the contract method 0xb32805c3.
//
// Solidity: function updateGeneration() returns()
func (_DKG *DKGSession) UpdateGeneration() (*types.Transaction, error) {
	return _DKG.Contract.UpdateGeneration(&_DKG.TransactOpts)
}

// UpdateGeneration is a paid mutator transaction binding the contract method 0xb32805c3.
//
// Solidity: function updateGeneration() returns()
func (_DKG *DKGTransactorSession) UpdateGeneration() (*types.Transaction, error) {
	return _DKG.Contract.UpdateGeneration(&_DKG.TransactOpts)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _generation, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGTransactor) VoteSigner(opts *bind.TransactOpts, _generation *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "voteSigner", _generation, _signerAddress, _signature)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _generation, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGSession) VoteSigner(_generation *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.Contract.VoteSigner(&_DKG.TransactOpts, _generation, _signerAddress, _signature)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _generation, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGTransactorSession) VoteSigner(_generation *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.Contract.VoteSigner(&_DKG.TransactOpts, _generation, _signerAddress, _signature)
}

// DKGInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DKG contract.
type DKGInitializedIterator struct {
	Event *DKGInitialized // Event containing the contract specifics and raw log

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
func (it *DKGInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGInitialized)
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
		it.Event = new(DKGInitialized)
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
func (it *DKGInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGInitialized represents a Initialized event raised by the DKG contract.
type DKGInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DKG *DKGFilterer) FilterInitialized(opts *bind.FilterOpts) (*DKGInitializedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DKGInitializedIterator{contract: _DKG.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DKG *DKGFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DKGInitialized) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGInitialized)
				if err := _DKG.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DKG *DKGFilterer) ParseInitialized(log types.Log) (*DKGInitialized, error) {
	event := new(DKGInitialized)
	if err := _DKG.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRoundDataFilledIterator is returned from FilterRoundDataFilled and is used to iterate over the raw logs and unpacked data for RoundDataFilled events raised by the DKG contract.
type DKGRoundDataFilledIterator struct {
	Event *DKGRoundDataFilled // Event containing the contract specifics and raw log

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
func (it *DKGRoundDataFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRoundDataFilled)
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
		it.Event = new(DKGRoundDataFilled)
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
func (it *DKGRoundDataFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRoundDataFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRoundDataFilled represents a RoundDataFilled event raised by the DKG contract.
type DKGRoundDataFilled struct {
	Generation *big.Int
	Round      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoundDataFilled is a free log retrieval operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 generation, uint256 round)
func (_DKG *DKGFilterer) FilterRoundDataFilled(opts *bind.FilterOpts) (*DKGRoundDataFilledIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "RoundDataFilled")
	if err != nil {
		return nil, err
	}
	return &DKGRoundDataFilledIterator{contract: _DKG.contract, event: "RoundDataFilled", logs: logs, sub: sub}, nil
}

// WatchRoundDataFilled is a free log subscription operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 generation, uint256 round)
func (_DKG *DKGFilterer) WatchRoundDataFilled(opts *bind.WatchOpts, sink chan<- *DKGRoundDataFilled) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "RoundDataFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRoundDataFilled)
				if err := _DKG.contract.UnpackLog(event, "RoundDataFilled", log); err != nil {
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

// ParseRoundDataFilled is a log parse operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 generation, uint256 round)
func (_DKG *DKGFilterer) ParseRoundDataFilled(log types.Log) (*DKGRoundDataFilled, error) {
	event := new(DKGRoundDataFilled)
	if err := _DKG.contract.UnpackLog(event, "RoundDataFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRoundDataProvidedIterator is returned from FilterRoundDataProvided and is used to iterate over the raw logs and unpacked data for RoundDataProvided events raised by the DKG contract.
type DKGRoundDataProvidedIterator struct {
	Event *DKGRoundDataProvided // Event containing the contract specifics and raw log

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
func (it *DKGRoundDataProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRoundDataProvided)
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
		it.Event = new(DKGRoundDataProvided)
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
func (it *DKGRoundDataProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRoundDataProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRoundDataProvided represents a RoundDataProvided event raised by the DKG contract.
type DKGRoundDataProvided struct {
	Generation *big.Int
	Round      *big.Int
	Validator  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoundDataProvided is a free log retrieval operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 generation, uint256 round, address validator)
func (_DKG *DKGFilterer) FilterRoundDataProvided(opts *bind.FilterOpts) (*DKGRoundDataProvidedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "RoundDataProvided")
	if err != nil {
		return nil, err
	}
	return &DKGRoundDataProvidedIterator{contract: _DKG.contract, event: "RoundDataProvided", logs: logs, sub: sub}, nil
}

// WatchRoundDataProvided is a free log subscription operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 generation, uint256 round, address validator)
func (_DKG *DKGFilterer) WatchRoundDataProvided(opts *bind.WatchOpts, sink chan<- *DKGRoundDataProvided) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "RoundDataProvided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRoundDataProvided)
				if err := _DKG.contract.UnpackLog(event, "RoundDataProvided", log); err != nil {
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

// ParseRoundDataProvided is a log parse operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 generation, uint256 round, address validator)
func (_DKG *DKGFilterer) ParseRoundDataProvided(log types.Log) (*DKGRoundDataProvided, error) {
	event := new(DKGRoundDataProvided)
	if err := _DKG.contract.UnpackLog(event, "RoundDataProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGSignerAddressUpdatedIterator is returned from FilterSignerAddressUpdated and is used to iterate over the raw logs and unpacked data for SignerAddressUpdated events raised by the DKG contract.
type DKGSignerAddressUpdatedIterator struct {
	Event *DKGSignerAddressUpdated // Event containing the contract specifics and raw log

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
func (it *DKGSignerAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGSignerAddressUpdated)
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
		it.Event = new(DKGSignerAddressUpdated)
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
func (it *DKGSignerAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGSignerAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGSignerAddressUpdated represents a SignerAddressUpdated event raised by the DKG contract.
type DKGSignerAddressUpdated struct {
	Generation    *big.Int
	SignerAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignerAddressUpdated is a free log retrieval operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 generation, address signerAddress)
func (_DKG *DKGFilterer) FilterSignerAddressUpdated(opts *bind.FilterOpts) (*DKGSignerAddressUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "SignerAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGSignerAddressUpdatedIterator{contract: _DKG.contract, event: "SignerAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerAddressUpdated is a free log subscription operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 generation, address signerAddress)
func (_DKG *DKGFilterer) WatchSignerAddressUpdated(opts *bind.WatchOpts, sink chan<- *DKGSignerAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "SignerAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGSignerAddressUpdated)
				if err := _DKG.contract.UnpackLog(event, "SignerAddressUpdated", log); err != nil {
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

// ParseSignerAddressUpdated is a log parse operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 generation, address signerAddress)
func (_DKG *DKGFilterer) ParseSignerAddressUpdated(log types.Log) (*DKGSignerAddressUpdated, error) {
	event := new(DKGSignerAddressUpdated)
	if err := _DKG.contract.UnpackLog(event, "SignerAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGSignerVotedIterator is returned from FilterSignerVoted and is used to iterate over the raw logs and unpacked data for SignerVoted events raised by the DKG contract.
type DKGSignerVotedIterator struct {
	Event *DKGSignerVoted // Event containing the contract specifics and raw log

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
func (it *DKGSignerVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGSignerVoted)
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
		it.Event = new(DKGSignerVoted)
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
func (it *DKGSignerVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGSignerVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGSignerVoted represents a SignerVoted event raised by the DKG contract.
type DKGSignerVoted struct {
	Generation       *big.Int
	Validator        common.Address
	CollectiveSigner common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSignerVoted is a free log retrieval operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 generation, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) FilterSignerVoted(opts *bind.FilterOpts) (*DKGSignerVotedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "SignerVoted")
	if err != nil {
		return nil, err
	}
	return &DKGSignerVotedIterator{contract: _DKG.contract, event: "SignerVoted", logs: logs, sub: sub}, nil
}

// WatchSignerVoted is a free log subscription operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 generation, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) WatchSignerVoted(opts *bind.WatchOpts, sink chan<- *DKGSignerVoted) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "SignerVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGSignerVoted)
				if err := _DKG.contract.UnpackLog(event, "SignerVoted", log); err != nil {
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

// ParseSignerVoted is a log parse operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 generation, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) ParseSignerVoted(log types.Log) (*DKGSignerVoted, error) {
	event := new(DKGSignerVoted)
	if err := _DKG.contract.UnpackLog(event, "SignerVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGThresholdSignerUpdatedIterator is returned from FilterThresholdSignerUpdated and is used to iterate over the raw logs and unpacked data for ThresholdSignerUpdated events raised by the DKG contract.
type DKGThresholdSignerUpdatedIterator struct {
	Event *DKGThresholdSignerUpdated // Event containing the contract specifics and raw log

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
func (it *DKGThresholdSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGThresholdSignerUpdated)
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
		it.Event = new(DKGThresholdSignerUpdated)
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
func (it *DKGThresholdSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGThresholdSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGThresholdSignerUpdated represents a ThresholdSignerUpdated event raised by the DKG contract.
type DKGThresholdSignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterThresholdSignerUpdated is a free log retrieval operation binding the contract event 0x8559aaab9220e2f109d51f7d09a3cca9b41451f4698287d982dc4549a70808bf.
//
// Solidity: event ThresholdSignerUpdated(address signer)
func (_DKG *DKGFilterer) FilterThresholdSignerUpdated(opts *bind.FilterOpts) (*DKGThresholdSignerUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ThresholdSignerUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGThresholdSignerUpdatedIterator{contract: _DKG.contract, event: "ThresholdSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdSignerUpdated is a free log subscription operation binding the contract event 0x8559aaab9220e2f109d51f7d09a3cca9b41451f4698287d982dc4549a70808bf.
//
// Solidity: event ThresholdSignerUpdated(address signer)
func (_DKG *DKGFilterer) WatchThresholdSignerUpdated(opts *bind.WatchOpts, sink chan<- *DKGThresholdSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ThresholdSignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGThresholdSignerUpdated)
				if err := _DKG.contract.UnpackLog(event, "ThresholdSignerUpdated", log); err != nil {
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

// ParseThresholdSignerUpdated is a log parse operation binding the contract event 0x8559aaab9220e2f109d51f7d09a3cca9b41451f4698287d982dc4549a70808bf.
//
// Solidity: event ThresholdSignerUpdated(address signer)
func (_DKG *DKGFilterer) ParseThresholdSignerUpdated(log types.Log) (*DKGThresholdSignerUpdated, error) {
	event := new(DKGThresholdSignerUpdated)
	if err := _DKG.contract.UnpackLog(event, "ThresholdSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the DKG contract.
type DKGValidatorsUpdatedIterator struct {
	Event *DKGValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *DKGValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGValidatorsUpdated)
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
		it.Event = new(DKGValidatorsUpdated)
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
func (it *DKGValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGValidatorsUpdated represents a ValidatorsUpdated event raised by the DKG contract.
type DKGValidatorsUpdated struct {
	Generation *big.Int
	Validators []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 generation, address[] validators)
func (_DKG *DKGFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts) (*DKGValidatorsUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGValidatorsUpdatedIterator{contract: _DKG.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 generation, address[] validators)
func (_DKG *DKGFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *DKGValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGValidatorsUpdated)
				if err := _DKG.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
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

// ParseValidatorsUpdated is a log parse operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 generation, address[] validators)
func (_DKG *DKGFilterer) ParseValidatorsUpdated(log types.Log) (*DKGValidatorsUpdated, error) {
	event := new(DKGValidatorsUpdated)
	if err := _DKG.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
