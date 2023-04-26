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

// EventRegistryMetaData contains all meta data concerning the EventRegistry contract.
var EventRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_appContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumEventRegistry.EventType\",\"name\":\"_eventType\",\"type\":\"uint8\"}],\"name\":\"EventRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EVENT_REGISTRY_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_appContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorFee\",\"type\":\"uint256\"},{\"internalType\":\"enumEventRegistry.EventType\",\"name\":\"_eventType\",\"type\":\"uint8\"}],\"name\":\"eventKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"eventType\",\"outputs\":[{\"internalType\":\"enumEventRegistry.EventType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_eventKey\",\"type\":\"bytes32\"}],\"name\":\"getEventType\",\"outputs\":[{\"internalType\":\"enumEventRegistry.EventType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorGetterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_appContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_validatorFee\",\"type\":\"uint256\"},{\"internalType\":\"enumEventRegistry.EventType\",\"name\":\"_eventType\",\"type\":\"uint8\"}],\"name\":\"registerEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"registeredEvents\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorGetter\",\"outputs\":[{\"internalType\":\"contractValidatorGetter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EventRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use EventRegistryMetaData.ABI instead.
var EventRegistryABI = EventRegistryMetaData.ABI

// EventRegistry is an auto generated Go binding around an Ethereum contract.
type EventRegistry struct {
	EventRegistryCaller     // Read-only binding to the contract
	EventRegistryTransactor // Write-only binding to the contract
	EventRegistryFilterer   // Log filterer for contract events
}

// EventRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventRegistrySession struct {
	Contract     *EventRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventRegistryCallerSession struct {
	Contract *EventRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EventRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventRegistryTransactorSession struct {
	Contract     *EventRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EventRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventRegistryRaw struct {
	Contract *EventRegistry // Generic contract binding to access the raw methods on
}

// EventRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventRegistryCallerRaw struct {
	Contract *EventRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// EventRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventRegistryTransactorRaw struct {
	Contract *EventRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventRegistry creates a new instance of EventRegistry, bound to a specific deployed contract.
func NewEventRegistry(address common.Address, backend bind.ContractBackend) (*EventRegistry, error) {
	contract, err := bindEventRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventRegistry{EventRegistryCaller: EventRegistryCaller{contract: contract}, EventRegistryTransactor: EventRegistryTransactor{contract: contract}, EventRegistryFilterer: EventRegistryFilterer{contract: contract}}, nil
}

// NewEventRegistryCaller creates a new read-only instance of EventRegistry, bound to a specific deployed contract.
func NewEventRegistryCaller(address common.Address, caller bind.ContractCaller) (*EventRegistryCaller, error) {
	contract, err := bindEventRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventRegistryCaller{contract: contract}, nil
}

// NewEventRegistryTransactor creates a new write-only instance of EventRegistry, bound to a specific deployed contract.
func NewEventRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*EventRegistryTransactor, error) {
	contract, err := bindEventRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventRegistryTransactor{contract: contract}, nil
}

// NewEventRegistryFilterer creates a new log filterer instance of EventRegistry, bound to a specific deployed contract.
func NewEventRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*EventRegistryFilterer, error) {
	contract, err := bindEventRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventRegistryFilterer{contract: contract}, nil
}

// bindEventRegistry binds a generic wrapper to an already deployed contract.
func bindEventRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventRegistry *EventRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventRegistry.Contract.EventRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventRegistry *EventRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventRegistry.Contract.EventRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventRegistry *EventRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventRegistry.Contract.EventRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventRegistry *EventRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventRegistry *EventRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventRegistry *EventRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventRegistry.Contract.contract.Transact(opts, method, params...)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_EventRegistry *EventRegistryCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_EventRegistry *EventRegistrySession) DKGKEY() (string, error) {
	return _EventRegistry.Contract.DKGKEY(&_EventRegistry.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_EventRegistry *EventRegistryCallerSession) DKGKEY() (string, error) {
	return _EventRegistry.Contract.DKGKEY(&_EventRegistry.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_EventRegistry *EventRegistryCaller) EVENTREGISTRYKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "EVENT_REGISTRY_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_EventRegistry *EventRegistrySession) EVENTREGISTRYKEY() (string, error) {
	return _EventRegistry.Contract.EVENTREGISTRYKEY(&_EventRegistry.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_EventRegistry *EventRegistryCallerSession) EVENTREGISTRYKEY() (string, error) {
	return _EventRegistry.Contract.EVENTREGISTRYKEY(&_EventRegistry.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_EventRegistry *EventRegistryCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_EventRegistry *EventRegistrySession) SLASHINGVOTINGKEY() (string, error) {
	return _EventRegistry.Contract.SLASHINGVOTINGKEY(&_EventRegistry.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_EventRegistry *EventRegistryCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _EventRegistry.Contract.SLASHINGVOTINGKEY(&_EventRegistry.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_EventRegistry *EventRegistryCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_EventRegistry *EventRegistrySession) STAKINGKEY() (string, error) {
	return _EventRegistry.Contract.STAKINGKEY(&_EventRegistry.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_EventRegistry *EventRegistryCallerSession) STAKINGKEY() (string, error) {
	return _EventRegistry.Contract.STAKINGKEY(&_EventRegistry.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_EventRegistry *EventRegistryCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_EventRegistry *EventRegistrySession) SUPPORTEDTOKENSKEY() (string, error) {
	return _EventRegistry.Contract.SUPPORTEDTOKENSKEY(&_EventRegistry.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_EventRegistry *EventRegistryCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _EventRegistry.Contract.SUPPORTEDTOKENSKEY(&_EventRegistry.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_EventRegistry *EventRegistryCaller) VALIDATORREWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_EventRegistry *EventRegistrySession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _EventRegistry.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_EventRegistry.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_EventRegistry *EventRegistryCallerSession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _EventRegistry.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_EventRegistry.CallOpts)
}

// EventKey is a free data retrieval call binding the contract method 0x3a4f4e2b.
//
// Solidity: function eventKey(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType) pure returns(bytes32)
func (_EventRegistry *EventRegistryCaller) EventKey(opts *bind.CallOpts, _hash [32]byte, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _validatorFee *big.Int, _eventType uint8) ([32]byte, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "eventKey", _hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EventKey is a free data retrieval call binding the contract method 0x3a4f4e2b.
//
// Solidity: function eventKey(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType) pure returns(bytes32)
func (_EventRegistry *EventRegistrySession) EventKey(_hash [32]byte, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _validatorFee *big.Int, _eventType uint8) ([32]byte, error) {
	return _EventRegistry.Contract.EventKey(&_EventRegistry.CallOpts, _hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType)
}

// EventKey is a free data retrieval call binding the contract method 0x3a4f4e2b.
//
// Solidity: function eventKey(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType) pure returns(bytes32)
func (_EventRegistry *EventRegistryCallerSession) EventKey(_hash [32]byte, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _validatorFee *big.Int, _eventType uint8) ([32]byte, error) {
	return _EventRegistry.Contract.EventKey(&_EventRegistry.CallOpts, _hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType)
}

// EventType is a free data retrieval call binding the contract method 0x20eaf723.
//
// Solidity: function eventType(bytes32 ) view returns(uint8)
func (_EventRegistry *EventRegistryCaller) EventType(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "eventType", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EventType is a free data retrieval call binding the contract method 0x20eaf723.
//
// Solidity: function eventType(bytes32 ) view returns(uint8)
func (_EventRegistry *EventRegistrySession) EventType(arg0 [32]byte) (uint8, error) {
	return _EventRegistry.Contract.EventType(&_EventRegistry.CallOpts, arg0)
}

// EventType is a free data retrieval call binding the contract method 0x20eaf723.
//
// Solidity: function eventType(bytes32 ) view returns(uint8)
func (_EventRegistry *EventRegistryCallerSession) EventType(arg0 [32]byte) (uint8, error) {
	return _EventRegistry.Contract.EventType(&_EventRegistry.CallOpts, arg0)
}

// GetEventType is a free data retrieval call binding the contract method 0x91d04b26.
//
// Solidity: function getEventType(bytes32 _eventKey) view returns(uint8)
func (_EventRegistry *EventRegistryCaller) GetEventType(opts *bind.CallOpts, _eventKey [32]byte) (uint8, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "getEventType", _eventKey)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetEventType is a free data retrieval call binding the contract method 0x91d04b26.
//
// Solidity: function getEventType(bytes32 _eventKey) view returns(uint8)
func (_EventRegistry *EventRegistrySession) GetEventType(_eventKey [32]byte) (uint8, error) {
	return _EventRegistry.Contract.GetEventType(&_EventRegistry.CallOpts, _eventKey)
}

// GetEventType is a free data retrieval call binding the contract method 0x91d04b26.
//
// Solidity: function getEventType(bytes32 _eventKey) view returns(uint8)
func (_EventRegistry *EventRegistryCallerSession) GetEventType(_eventKey [32]byte) (uint8, error) {
	return _EventRegistry.Contract.GetEventType(&_EventRegistry.CallOpts, _eventKey)
}

// RegisteredEvents is a free data retrieval call binding the contract method 0x050b5c5e.
//
// Solidity: function registeredEvents(bytes32 ) view returns(bool)
func (_EventRegistry *EventRegistryCaller) RegisteredEvents(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "registeredEvents", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredEvents is a free data retrieval call binding the contract method 0x050b5c5e.
//
// Solidity: function registeredEvents(bytes32 ) view returns(bool)
func (_EventRegistry *EventRegistrySession) RegisteredEvents(arg0 [32]byte) (bool, error) {
	return _EventRegistry.Contract.RegisteredEvents(&_EventRegistry.CallOpts, arg0)
}

// RegisteredEvents is a free data retrieval call binding the contract method 0x050b5c5e.
//
// Solidity: function registeredEvents(bytes32 ) view returns(bool)
func (_EventRegistry *EventRegistryCallerSession) RegisteredEvents(arg0 [32]byte) (bool, error) {
	return _EventRegistry.Contract.RegisteredEvents(&_EventRegistry.CallOpts, arg0)
}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_EventRegistry *EventRegistryCaller) ValidatorGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventRegistry.contract.Call(opts, &out, "validatorGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_EventRegistry *EventRegistrySession) ValidatorGetter() (common.Address, error) {
	return _EventRegistry.Contract.ValidatorGetter(&_EventRegistry.CallOpts)
}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_EventRegistry *EventRegistryCallerSession) ValidatorGetter() (common.Address, error) {
	return _EventRegistry.Contract.ValidatorGetter(&_EventRegistry.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorGetterAddress) returns()
func (_EventRegistry *EventRegistryTransactor) Initialize(opts *bind.TransactOpts, _validatorGetterAddress common.Address) (*types.Transaction, error) {
	return _EventRegistry.contract.Transact(opts, "initialize", _validatorGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorGetterAddress) returns()
func (_EventRegistry *EventRegistrySession) Initialize(_validatorGetterAddress common.Address) (*types.Transaction, error) {
	return _EventRegistry.Contract.Initialize(&_EventRegistry.TransactOpts, _validatorGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _validatorGetterAddress) returns()
func (_EventRegistry *EventRegistryTransactorSession) Initialize(_validatorGetterAddress common.Address) (*types.Transaction, error) {
	return _EventRegistry.Contract.Initialize(&_EventRegistry.TransactOpts, _validatorGetterAddress)
}

// RegisterEvent is a paid mutator transaction binding the contract method 0xdc69f79b.
//
// Solidity: function registerEvent(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType) returns()
func (_EventRegistry *EventRegistryTransactor) RegisterEvent(opts *bind.TransactOpts, _hash [32]byte, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _validatorFee *big.Int, _eventType uint8) (*types.Transaction, error) {
	return _EventRegistry.contract.Transact(opts, "registerEvent", _hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType)
}

// RegisterEvent is a paid mutator transaction binding the contract method 0xdc69f79b.
//
// Solidity: function registerEvent(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType) returns()
func (_EventRegistry *EventRegistrySession) RegisterEvent(_hash [32]byte, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _validatorFee *big.Int, _eventType uint8) (*types.Transaction, error) {
	return _EventRegistry.Contract.RegisterEvent(&_EventRegistry.TransactOpts, _hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType)
}

// RegisterEvent is a paid mutator transaction binding the contract method 0xdc69f79b.
//
// Solidity: function registerEvent(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType) returns()
func (_EventRegistry *EventRegistryTransactorSession) RegisterEvent(_hash [32]byte, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _validatorFee *big.Int, _eventType uint8) (*types.Transaction, error) {
	return _EventRegistry.Contract.RegisterEvent(&_EventRegistry.TransactOpts, _hash, _appContract, _sourceChain, _destinationChain, _data, _validatorFee, _eventType)
}

// EventRegistryEventRegisteredIterator is returned from FilterEventRegistered and is used to iterate over the raw logs and unpacked data for EventRegistered events raised by the EventRegistry contract.
type EventRegistryEventRegisteredIterator struct {
	Event *EventRegistryEventRegistered // Event containing the contract specifics and raw log

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
func (it *EventRegistryEventRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventRegistryEventRegistered)
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
		it.Event = new(EventRegistryEventRegistered)
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
func (it *EventRegistryEventRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventRegistryEventRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventRegistryEventRegistered represents a EventRegistered event raised by the EventRegistry contract.
type EventRegistryEventRegistered struct {
	Hash             [32]byte
	AppContract      common.Address
	SourceChain      *big.Int
	DestinationChain *big.Int
	Data             []byte
	ValidatorFee     *big.Int
	EventType        uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEventRegistered is a free log retrieval operation binding the contract event 0x7d58b416fcb1356f139a3e315ebd453e89ad07fd73e28b5f5ff4b3bc3229ddc2.
//
// Solidity: event EventRegistered(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType)
func (_EventRegistry *EventRegistryFilterer) FilterEventRegistered(opts *bind.FilterOpts) (*EventRegistryEventRegisteredIterator, error) {

	logs, sub, err := _EventRegistry.contract.FilterLogs(opts, "EventRegistered")
	if err != nil {
		return nil, err
	}
	return &EventRegistryEventRegisteredIterator{contract: _EventRegistry.contract, event: "EventRegistered", logs: logs, sub: sub}, nil
}

// WatchEventRegistered is a free log subscription operation binding the contract event 0x7d58b416fcb1356f139a3e315ebd453e89ad07fd73e28b5f5ff4b3bc3229ddc2.
//
// Solidity: event EventRegistered(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType)
func (_EventRegistry *EventRegistryFilterer) WatchEventRegistered(opts *bind.WatchOpts, sink chan<- *EventRegistryEventRegistered) (event.Subscription, error) {

	logs, sub, err := _EventRegistry.contract.WatchLogs(opts, "EventRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventRegistryEventRegistered)
				if err := _EventRegistry.contract.UnpackLog(event, "EventRegistered", log); err != nil {
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

// ParseEventRegistered is a log parse operation binding the contract event 0x7d58b416fcb1356f139a3e315ebd453e89ad07fd73e28b5f5ff4b3bc3229ddc2.
//
// Solidity: event EventRegistered(bytes32 _hash, address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _validatorFee, uint8 _eventType)
func (_EventRegistry *EventRegistryFilterer) ParseEventRegistered(log types.Log) (*EventRegistryEventRegistered, error) {
	event := new(EventRegistryEventRegistered)
	if err := _EventRegistry.contract.UnpackLog(event, "EventRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the EventRegistry contract.
type EventRegistryInitializedIterator struct {
	Event *EventRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *EventRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventRegistryInitialized)
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
		it.Event = new(EventRegistryInitialized)
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
func (it *EventRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventRegistryInitialized represents a Initialized event raised by the EventRegistry contract.
type EventRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EventRegistry *EventRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*EventRegistryInitializedIterator, error) {

	logs, sub, err := _EventRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EventRegistryInitializedIterator{contract: _EventRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_EventRegistry *EventRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EventRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _EventRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventRegistryInitialized)
				if err := _EventRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_EventRegistry *EventRegistryFilterer) ParseInitialized(log types.Log) (*EventRegistryInitialized, error) {
	event := new(EventRegistryInitialized)
	if err := _EventRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
