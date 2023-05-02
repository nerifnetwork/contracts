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

// RegistryConfig is an auto generated low-level Go binding around an user-defined struct.
type RegistryConfig struct {
	PerformanceOverhead         *big.Int
	PerformancePremiumThreshold uint8
	RegistrationOverhead        *big.Int
	CancellationOverhead        *big.Int
	MaxWorkflowsPerAccount      uint16
}

// Workflow is an auto generated low-level Go binding around an user-defined struct.
type Workflow struct {
	Id         *big.Int
	Owner      common.Address
	Hash       []byte
	Status     uint8
	IsInternal bool
	TotalSpent *big.Int
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasUsed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Performance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"WorkflowActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"WorkflowCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"WorkflowPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"WorkflowRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"WorkflowResumed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"activateWorkflow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"cancelWorkflow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"performanceOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"performancePremiumThreshold\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"registrationOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancellationOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"maxWorkflowsPerAccount\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"fundBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getGateway\",\"outputs\":[{\"internalType\":\"contractIGateway\",\"name\":\"gateway\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getWorkflow\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"},{\"internalType\":\"enumWorkflowStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isInternal\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structWorkflow\",\"name\":\"workflow\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_networkAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isMainChain\",\"type\":\"bool\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMainChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkAddress\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"pauseWorkflow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"}],\"name\":\"registerWorkflow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"resumeWorkflow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"performanceOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"performancePremiumThreshold\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"registrationOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cancellationOverhead\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"maxWorkflowsPerAccount\",\"type\":\"uint16\"}],\"internalType\":\"structRegistry.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"withdrawBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 performanceOverhead, uint8 performancePremiumThreshold, uint256 registrationOverhead, uint256 cancellationOverhead, uint16 maxWorkflowsPerAccount)
func (_Registry *RegistryCaller) Config(opts *bind.CallOpts) (struct {
	PerformanceOverhead         *big.Int
	PerformancePremiumThreshold uint8
	RegistrationOverhead        *big.Int
	CancellationOverhead        *big.Int
	MaxWorkflowsPerAccount      uint16
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		PerformanceOverhead         *big.Int
		PerformancePremiumThreshold uint8
		RegistrationOverhead        *big.Int
		CancellationOverhead        *big.Int
		MaxWorkflowsPerAccount      uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PerformanceOverhead = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PerformancePremiumThreshold = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.RegistrationOverhead = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CancellationOverhead = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxWorkflowsPerAccount = *abi.ConvertType(out[4], new(uint16)).(*uint16)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 performanceOverhead, uint8 performancePremiumThreshold, uint256 registrationOverhead, uint256 cancellationOverhead, uint16 maxWorkflowsPerAccount)
func (_Registry *RegistrySession) Config() (struct {
	PerformanceOverhead         *big.Int
	PerformancePremiumThreshold uint8
	RegistrationOverhead        *big.Int
	CancellationOverhead        *big.Int
	MaxWorkflowsPerAccount      uint16
}, error) {
	return _Registry.Contract.Config(&_Registry.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 performanceOverhead, uint8 performancePremiumThreshold, uint256 registrationOverhead, uint256 cancellationOverhead, uint16 maxWorkflowsPerAccount)
func (_Registry *RegistryCallerSession) Config() (struct {
	PerformanceOverhead         *big.Int
	PerformancePremiumThreshold uint8
	RegistrationOverhead        *big.Int
	CancellationOverhead        *big.Int
	MaxWorkflowsPerAccount      uint16
}, error) {
	return _Registry.Contract.Config(&_Registry.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256 balance)
func (_Registry *RegistryCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256 balance)
func (_Registry *RegistrySession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.GetBalance(&_Registry.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256 balance)
func (_Registry *RegistryCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Registry.Contract.GetBalance(&_Registry.CallOpts, addr)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address gateway, uint256 index)
func (_Registry *RegistryCaller) GetGateway(opts *bind.CallOpts, owner common.Address) (struct {
	Gateway common.Address
	Index   *big.Int
}, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getGateway", owner)

	outstruct := new(struct {
		Gateway common.Address
		Index   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gateway = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address gateway, uint256 index)
func (_Registry *RegistrySession) GetGateway(owner common.Address) (struct {
	Gateway common.Address
	Index   *big.Int
}, error) {
	return _Registry.Contract.GetGateway(&_Registry.CallOpts, owner)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address gateway, uint256 index)
func (_Registry *RegistryCallerSession) GetGateway(owner common.Address) (struct {
	Gateway common.Address
	Index   *big.Int
}, error) {
	return _Registry.Contract.GetGateway(&_Registry.CallOpts, owner)
}

// GetWorkflow is a free data retrieval call binding the contract method 0xeec7b03c.
//
// Solidity: function getWorkflow(uint256 id) view returns((uint256,address,bytes,uint8,bool,uint256) workflow)
func (_Registry *RegistryCaller) GetWorkflow(opts *bind.CallOpts, id *big.Int) (Workflow, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getWorkflow", id)

	if err != nil {
		return *new(Workflow), err
	}

	out0 := *abi.ConvertType(out[0], new(Workflow)).(*Workflow)

	return out0, err

}

// GetWorkflow is a free data retrieval call binding the contract method 0xeec7b03c.
//
// Solidity: function getWorkflow(uint256 id) view returns((uint256,address,bytes,uint8,bool,uint256) workflow)
func (_Registry *RegistrySession) GetWorkflow(id *big.Int) (Workflow, error) {
	return _Registry.Contract.GetWorkflow(&_Registry.CallOpts, id)
}

// GetWorkflow is a free data retrieval call binding the contract method 0xeec7b03c.
//
// Solidity: function getWorkflow(uint256 id) view returns((uint256,address,bytes,uint8,bool,uint256) workflow)
func (_Registry *RegistryCallerSession) GetWorkflow(id *big.Int) (Workflow, error) {
	return _Registry.Contract.GetWorkflow(&_Registry.CallOpts, id)
}

// IsMainChain is a free data retrieval call binding the contract method 0xb834f6fb.
//
// Solidity: function isMainChain() view returns(bool)
func (_Registry *RegistryCaller) IsMainChain(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isMainChain")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMainChain is a free data retrieval call binding the contract method 0xb834f6fb.
//
// Solidity: function isMainChain() view returns(bool)
func (_Registry *RegistrySession) IsMainChain() (bool, error) {
	return _Registry.Contract.IsMainChain(&_Registry.CallOpts)
}

// IsMainChain is a free data retrieval call binding the contract method 0xb834f6fb.
//
// Solidity: function isMainChain() view returns(bool)
func (_Registry *RegistryCallerSession) IsMainChain() (bool, error) {
	return _Registry.Contract.IsMainChain(&_Registry.CallOpts)
}

// NetworkAddress is a free data retrieval call binding the contract method 0xd69e1bad.
//
// Solidity: function networkAddress() view returns(address)
func (_Registry *RegistryCaller) NetworkAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "networkAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkAddress is a free data retrieval call binding the contract method 0xd69e1bad.
//
// Solidity: function networkAddress() view returns(address)
func (_Registry *RegistrySession) NetworkAddress() (common.Address, error) {
	return _Registry.Contract.NetworkAddress(&_Registry.CallOpts)
}

// NetworkAddress is a free data retrieval call binding the contract method 0xd69e1bad.
//
// Solidity: function networkAddress() view returns(address)
func (_Registry *RegistryCallerSession) NetworkAddress() (common.Address, error) {
	return _Registry.Contract.NetworkAddress(&_Registry.CallOpts)
}

// NetworkRewards is a free data retrieval call binding the contract method 0x6b5d4206.
//
// Solidity: function networkRewards() view returns(uint256)
func (_Registry *RegistryCaller) NetworkRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "networkRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NetworkRewards is a free data retrieval call binding the contract method 0x6b5d4206.
//
// Solidity: function networkRewards() view returns(uint256)
func (_Registry *RegistrySession) NetworkRewards() (*big.Int, error) {
	return _Registry.Contract.NetworkRewards(&_Registry.CallOpts)
}

// NetworkRewards is a free data retrieval call binding the contract method 0x6b5d4206.
//
// Solidity: function networkRewards() view returns(uint256)
func (_Registry *RegistryCallerSession) NetworkRewards() (*big.Int, error) {
	return _Registry.Contract.NetworkRewards(&_Registry.CallOpts)
}

// ActivateWorkflow is a paid mutator transaction binding the contract method 0x5902c15e.
//
// Solidity: function activateWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactor) ActivateWorkflow(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "activateWorkflow", id)
}

// ActivateWorkflow is a paid mutator transaction binding the contract method 0x5902c15e.
//
// Solidity: function activateWorkflow(uint256 id) returns()
func (_Registry *RegistrySession) ActivateWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ActivateWorkflow(&_Registry.TransactOpts, id)
}

// ActivateWorkflow is a paid mutator transaction binding the contract method 0x5902c15e.
//
// Solidity: function activateWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactorSession) ActivateWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ActivateWorkflow(&_Registry.TransactOpts, id)
}

// CancelWorkflow is a paid mutator transaction binding the contract method 0xd0c81e98.
//
// Solidity: function cancelWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactor) CancelWorkflow(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "cancelWorkflow", id)
}

// CancelWorkflow is a paid mutator transaction binding the contract method 0xd0c81e98.
//
// Solidity: function cancelWorkflow(uint256 id) returns()
func (_Registry *RegistrySession) CancelWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.CancelWorkflow(&_Registry.TransactOpts, id)
}

// CancelWorkflow is a paid mutator transaction binding the contract method 0xd0c81e98.
//
// Solidity: function cancelWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactorSession) CancelWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.CancelWorkflow(&_Registry.TransactOpts, id)
}

// FundBalance is a paid mutator transaction binding the contract method 0x0ad5a865.
//
// Solidity: function fundBalance(address addr) payable returns()
func (_Registry *RegistryTransactor) FundBalance(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "fundBalance", addr)
}

// FundBalance is a paid mutator transaction binding the contract method 0x0ad5a865.
//
// Solidity: function fundBalance(address addr) payable returns()
func (_Registry *RegistrySession) FundBalance(addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.FundBalance(&_Registry.TransactOpts, addr)
}

// FundBalance is a paid mutator transaction binding the contract method 0x0ad5a865.
//
// Solidity: function fundBalance(address addr) payable returns()
func (_Registry *RegistryTransactorSession) FundBalance(addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.FundBalance(&_Registry.TransactOpts, addr)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _networkAddress, bool _isMainChain) returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts, _networkAddress common.Address, _isMainChain bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize", _networkAddress, _isMainChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _networkAddress, bool _isMainChain) returns()
func (_Registry *RegistrySession) Initialize(_networkAddress common.Address, _isMainChain bool) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, _networkAddress, _isMainChain)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address _networkAddress, bool _isMainChain) returns()
func (_Registry *RegistryTransactorSession) Initialize(_networkAddress common.Address, _isMainChain bool) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, _networkAddress, _isMainChain)
}

// PauseWorkflow is a paid mutator transaction binding the contract method 0x0813ce96.
//
// Solidity: function pauseWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactor) PauseWorkflow(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "pauseWorkflow", id)
}

// PauseWorkflow is a paid mutator transaction binding the contract method 0x0813ce96.
//
// Solidity: function pauseWorkflow(uint256 id) returns()
func (_Registry *RegistrySession) PauseWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.PauseWorkflow(&_Registry.TransactOpts, id)
}

// PauseWorkflow is a paid mutator transaction binding the contract method 0x0813ce96.
//
// Solidity: function pauseWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactorSession) PauseWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.PauseWorkflow(&_Registry.TransactOpts, id)
}

// Perform is a paid mutator transaction binding the contract method 0xd6bb757b.
//
// Solidity: function perform(uint256 workflowId, uint256 gasAmount, bytes data, address target) returns()
func (_Registry *RegistryTransactor) Perform(opts *bind.TransactOpts, workflowId *big.Int, gasAmount *big.Int, data []byte, target common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "perform", workflowId, gasAmount, data, target)
}

// Perform is a paid mutator transaction binding the contract method 0xd6bb757b.
//
// Solidity: function perform(uint256 workflowId, uint256 gasAmount, bytes data, address target) returns()
func (_Registry *RegistrySession) Perform(workflowId *big.Int, gasAmount *big.Int, data []byte, target common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Perform(&_Registry.TransactOpts, workflowId, gasAmount, data, target)
}

// Perform is a paid mutator transaction binding the contract method 0xd6bb757b.
//
// Solidity: function perform(uint256 workflowId, uint256 gasAmount, bytes data, address target) returns()
func (_Registry *RegistryTransactorSession) Perform(workflowId *big.Int, gasAmount *big.Int, data []byte, target common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Perform(&_Registry.TransactOpts, workflowId, gasAmount, data, target)
}

// RegisterWorkflow is a paid mutator transaction binding the contract method 0x663d4503.
//
// Solidity: function registerWorkflow(uint256 id, address owner, bytes hash) returns()
func (_Registry *RegistryTransactor) RegisterWorkflow(opts *bind.TransactOpts, id *big.Int, owner common.Address, hash []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerWorkflow", id, owner, hash)
}

// RegisterWorkflow is a paid mutator transaction binding the contract method 0x663d4503.
//
// Solidity: function registerWorkflow(uint256 id, address owner, bytes hash) returns()
func (_Registry *RegistrySession) RegisterWorkflow(id *big.Int, owner common.Address, hash []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterWorkflow(&_Registry.TransactOpts, id, owner, hash)
}

// RegisterWorkflow is a paid mutator transaction binding the contract method 0x663d4503.
//
// Solidity: function registerWorkflow(uint256 id, address owner, bytes hash) returns()
func (_Registry *RegistryTransactorSession) RegisterWorkflow(id *big.Int, owner common.Address, hash []byte) (*types.Transaction, error) {
	return _Registry.Contract.RegisterWorkflow(&_Registry.TransactOpts, id, owner, hash)
}

// ResumeWorkflow is a paid mutator transaction binding the contract method 0xd692157b.
//
// Solidity: function resumeWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactor) ResumeWorkflow(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "resumeWorkflow", id)
}

// ResumeWorkflow is a paid mutator transaction binding the contract method 0xd692157b.
//
// Solidity: function resumeWorkflow(uint256 id) returns()
func (_Registry *RegistrySession) ResumeWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ResumeWorkflow(&_Registry.TransactOpts, id)
}

// ResumeWorkflow is a paid mutator transaction binding the contract method 0xd692157b.
//
// Solidity: function resumeWorkflow(uint256 id) returns()
func (_Registry *RegistryTransactorSession) ResumeWorkflow(id *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ResumeWorkflow(&_Registry.TransactOpts, id)
}

// SetConfig is a paid mutator transaction binding the contract method 0x2ad5d340.
//
// Solidity: function setConfig((uint256,uint8,uint256,uint256,uint16) _config) returns()
func (_Registry *RegistryTransactor) SetConfig(opts *bind.TransactOpts, _config RegistryConfig) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setConfig", _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x2ad5d340.
//
// Solidity: function setConfig((uint256,uint8,uint256,uint256,uint16) _config) returns()
func (_Registry *RegistrySession) SetConfig(_config RegistryConfig) (*types.Transaction, error) {
	return _Registry.Contract.SetConfig(&_Registry.TransactOpts, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x2ad5d340.
//
// Solidity: function setConfig((uint256,uint8,uint256,uint256,uint16) _config) returns()
func (_Registry *RegistryTransactorSession) SetConfig(_config RegistryConfig) (*types.Transaction, error) {
	return _Registry.Contract.SetConfig(&_Registry.TransactOpts, _config)
}

// SetGateway is a paid mutator transaction binding the contract method 0x774c208a.
//
// Solidity: function setGateway(address owner, address gateway) returns()
func (_Registry *RegistryTransactor) SetGateway(opts *bind.TransactOpts, owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setGateway", owner, gateway)
}

// SetGateway is a paid mutator transaction binding the contract method 0x774c208a.
//
// Solidity: function setGateway(address owner, address gateway) returns()
func (_Registry *RegistrySession) SetGateway(owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGateway(&_Registry.TransactOpts, owner, gateway)
}

// SetGateway is a paid mutator transaction binding the contract method 0x774c208a.
//
// Solidity: function setGateway(address owner, address gateway) returns()
func (_Registry *RegistryTransactorSession) SetGateway(owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGateway(&_Registry.TransactOpts, owner, gateway)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x756af45f.
//
// Solidity: function withdrawBalance(address addr) returns()
func (_Registry *RegistryTransactor) WithdrawBalance(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "withdrawBalance", addr)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x756af45f.
//
// Solidity: function withdrawBalance(address addr) returns()
func (_Registry *RegistrySession) WithdrawBalance(addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.WithdrawBalance(&_Registry.TransactOpts, addr)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x756af45f.
//
// Solidity: function withdrawBalance(address addr) returns()
func (_Registry *RegistryTransactorSession) WithdrawBalance(addr common.Address) (*types.Transaction, error) {
	return _Registry.Contract.WithdrawBalance(&_Registry.TransactOpts, addr)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_Registry *RegistryTransactor) WithdrawRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "withdrawRewards")
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_Registry *RegistrySession) WithdrawRewards() (*types.Transaction, error) {
	return _Registry.Contract.WithdrawRewards(&_Registry.TransactOpts)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xc7b8981c.
//
// Solidity: function withdrawRewards() returns()
func (_Registry *RegistryTransactorSession) WithdrawRewards() (*types.Transaction, error) {
	return _Registry.Contract.WithdrawRewards(&_Registry.TransactOpts)
}

// RegistryBalanceFundedIterator is returned from FilterBalanceFunded and is used to iterate over the raw logs and unpacked data for BalanceFunded events raised by the Registry contract.
type RegistryBalanceFundedIterator struct {
	Event *RegistryBalanceFunded // Event containing the contract specifics and raw log

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
func (it *RegistryBalanceFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBalanceFunded)
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
		it.Event = new(RegistryBalanceFunded)
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
func (it *RegistryBalanceFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBalanceFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBalanceFunded represents a BalanceFunded event raised by the Registry contract.
type RegistryBalanceFunded struct {
	WorkflowOwner common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBalanceFunded is a free log retrieval operation binding the contract event 0xe7765cbdc1b48b4b15642a6dadba14a2bfff6298f44152508a14f634fead42c4.
//
// Solidity: event BalanceFunded(address workflowOwner, uint256 amount)
func (_Registry *RegistryFilterer) FilterBalanceFunded(opts *bind.FilterOpts) (*RegistryBalanceFundedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "BalanceFunded")
	if err != nil {
		return nil, err
	}
	return &RegistryBalanceFundedIterator{contract: _Registry.contract, event: "BalanceFunded", logs: logs, sub: sub}, nil
}

// WatchBalanceFunded is a free log subscription operation binding the contract event 0xe7765cbdc1b48b4b15642a6dadba14a2bfff6298f44152508a14f634fead42c4.
//
// Solidity: event BalanceFunded(address workflowOwner, uint256 amount)
func (_Registry *RegistryFilterer) WatchBalanceFunded(opts *bind.WatchOpts, sink chan<- *RegistryBalanceFunded) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "BalanceFunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBalanceFunded)
				if err := _Registry.contract.UnpackLog(event, "BalanceFunded", log); err != nil {
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

// ParseBalanceFunded is a log parse operation binding the contract event 0xe7765cbdc1b48b4b15642a6dadba14a2bfff6298f44152508a14f634fead42c4.
//
// Solidity: event BalanceFunded(address workflowOwner, uint256 amount)
func (_Registry *RegistryFilterer) ParseBalanceFunded(log types.Log) (*RegistryBalanceFunded, error) {
	event := new(RegistryBalanceFunded)
	if err := _Registry.contract.UnpackLog(event, "BalanceFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryBalanceWithdrawnIterator is returned from FilterBalanceWithdrawn and is used to iterate over the raw logs and unpacked data for BalanceWithdrawn events raised by the Registry contract.
type RegistryBalanceWithdrawnIterator struct {
	Event *RegistryBalanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *RegistryBalanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryBalanceWithdrawn)
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
		it.Event = new(RegistryBalanceWithdrawn)
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
func (it *RegistryBalanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryBalanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryBalanceWithdrawn represents a BalanceWithdrawn event raised by the Registry contract.
type RegistryBalanceWithdrawn struct {
	WorkflowOwner common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBalanceWithdrawn is a free log retrieval operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address workflowOwner, uint256 amount)
func (_Registry *RegistryFilterer) FilterBalanceWithdrawn(opts *bind.FilterOpts) (*RegistryBalanceWithdrawnIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "BalanceWithdrawn")
	if err != nil {
		return nil, err
	}
	return &RegistryBalanceWithdrawnIterator{contract: _Registry.contract, event: "BalanceWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBalanceWithdrawn is a free log subscription operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address workflowOwner, uint256 amount)
func (_Registry *RegistryFilterer) WatchBalanceWithdrawn(opts *bind.WatchOpts, sink chan<- *RegistryBalanceWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "BalanceWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryBalanceWithdrawn)
				if err := _Registry.contract.UnpackLog(event, "BalanceWithdrawn", log); err != nil {
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

// ParseBalanceWithdrawn is a log parse operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address workflowOwner, uint256 amount)
func (_Registry *RegistryFilterer) ParseBalanceWithdrawn(log types.Log) (*RegistryBalanceWithdrawn, error) {
	event := new(RegistryBalanceWithdrawn)
	if err := _Registry.contract.UnpackLog(event, "BalanceWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryGatewaySetIterator is returned from FilterGatewaySet and is used to iterate over the raw logs and unpacked data for GatewaySet events raised by the Registry contract.
type RegistryGatewaySetIterator struct {
	Event *RegistryGatewaySet // Event containing the contract specifics and raw log

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
func (it *RegistryGatewaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryGatewaySet)
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
		it.Event = new(RegistryGatewaySet)
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
func (it *RegistryGatewaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryGatewaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryGatewaySet represents a GatewaySet event raised by the Registry contract.
type RegistryGatewaySet struct {
	WorkflowOwner common.Address
	Gateway       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGatewaySet is a free log retrieval operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address workflowOwner, address gateway)
func (_Registry *RegistryFilterer) FilterGatewaySet(opts *bind.FilterOpts) (*RegistryGatewaySetIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "GatewaySet")
	if err != nil {
		return nil, err
	}
	return &RegistryGatewaySetIterator{contract: _Registry.contract, event: "GatewaySet", logs: logs, sub: sub}, nil
}

// WatchGatewaySet is a free log subscription operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address workflowOwner, address gateway)
func (_Registry *RegistryFilterer) WatchGatewaySet(opts *bind.WatchOpts, sink chan<- *RegistryGatewaySet) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "GatewaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryGatewaySet)
				if err := _Registry.contract.UnpackLog(event, "GatewaySet", log); err != nil {
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

// ParseGatewaySet is a log parse operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address workflowOwner, address gateway)
func (_Registry *RegistryFilterer) ParseGatewaySet(log types.Log) (*RegistryGatewaySet, error) {
	event := new(RegistryGatewaySet)
	if err := _Registry.contract.UnpackLog(event, "GatewaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registry contract.
type RegistryInitializedIterator struct {
	Event *RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryInitialized)
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
		it.Event = new(RegistryInitialized)
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
func (it *RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryInitialized represents a Initialized event raised by the Registry contract.
type RegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryInitializedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryInitializedIterator{contract: _Registry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryInitialized)
				if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseInitialized(log types.Log) (*RegistryInitialized, error) {
	event := new(RegistryInitialized)
	if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryPerformanceIterator is returned from FilterPerformance and is used to iterate over the raw logs and unpacked data for Performance events raised by the Registry contract.
type RegistryPerformanceIterator struct {
	Event *RegistryPerformance // Event containing the contract specifics and raw log

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
func (it *RegistryPerformanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPerformance)
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
		it.Event = new(RegistryPerformance)
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
func (it *RegistryPerformanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPerformanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPerformance represents a Performance event raised by the Registry contract.
type RegistryPerformance struct {
	Id      *big.Int
	GasUsed *big.Int
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPerformance is a free log retrieval operation binding the contract event 0xc723c444dde505205b3ec0c789ed1adeade412952dc2caecb0ac55b9668e0105.
//
// Solidity: event Performance(uint256 id, uint256 gasUsed, bool success)
func (_Registry *RegistryFilterer) FilterPerformance(opts *bind.FilterOpts) (*RegistryPerformanceIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Performance")
	if err != nil {
		return nil, err
	}
	return &RegistryPerformanceIterator{contract: _Registry.contract, event: "Performance", logs: logs, sub: sub}, nil
}

// WatchPerformance is a free log subscription operation binding the contract event 0xc723c444dde505205b3ec0c789ed1adeade412952dc2caecb0ac55b9668e0105.
//
// Solidity: event Performance(uint256 id, uint256 gasUsed, bool success)
func (_Registry *RegistryFilterer) WatchPerformance(opts *bind.WatchOpts, sink chan<- *RegistryPerformance) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Performance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPerformance)
				if err := _Registry.contract.UnpackLog(event, "Performance", log); err != nil {
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

// ParsePerformance is a log parse operation binding the contract event 0xc723c444dde505205b3ec0c789ed1adeade412952dc2caecb0ac55b9668e0105.
//
// Solidity: event Performance(uint256 id, uint256 gasUsed, bool success)
func (_Registry *RegistryFilterer) ParsePerformance(log types.Log) (*RegistryPerformance, error) {
	event := new(RegistryPerformance)
	if err := _Registry.contract.UnpackLog(event, "Performance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRewardsWithdrawnIterator is returned from FilterRewardsWithdrawn and is used to iterate over the raw logs and unpacked data for RewardsWithdrawn events raised by the Registry contract.
type RegistryRewardsWithdrawnIterator struct {
	Event *RegistryRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *RegistryRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRewardsWithdrawn)
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
		it.Event = new(RegistryRewardsWithdrawn)
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
func (it *RegistryRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRewardsWithdrawn represents a RewardsWithdrawn event raised by the Registry contract.
type RegistryRewardsWithdrawn struct {
	Addr   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsWithdrawn is a free log retrieval operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address addr, uint256 amount)
func (_Registry *RegistryFilterer) FilterRewardsWithdrawn(opts *bind.FilterOpts) (*RegistryRewardsWithdrawnIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "RewardsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &RegistryRewardsWithdrawnIterator{contract: _Registry.contract, event: "RewardsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRewardsWithdrawn is a free log subscription operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address addr, uint256 amount)
func (_Registry *RegistryFilterer) WatchRewardsWithdrawn(opts *bind.WatchOpts, sink chan<- *RegistryRewardsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "RewardsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRewardsWithdrawn)
				if err := _Registry.contract.UnpackLog(event, "RewardsWithdrawn", log); err != nil {
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

// ParseRewardsWithdrawn is a log parse operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address addr, uint256 amount)
func (_Registry *RegistryFilterer) ParseRewardsWithdrawn(log types.Log) (*RegistryRewardsWithdrawn, error) {
	event := new(RegistryRewardsWithdrawn)
	if err := _Registry.contract.UnpackLog(event, "RewardsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWorkflowActivatedIterator is returned from FilterWorkflowActivated and is used to iterate over the raw logs and unpacked data for WorkflowActivated events raised by the Registry contract.
type RegistryWorkflowActivatedIterator struct {
	Event *RegistryWorkflowActivated // Event containing the contract specifics and raw log

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
func (it *RegistryWorkflowActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWorkflowActivated)
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
		it.Event = new(RegistryWorkflowActivated)
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
func (it *RegistryWorkflowActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWorkflowActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWorkflowActivated represents a WorkflowActivated event raised by the Registry contract.
type RegistryWorkflowActivated struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWorkflowActivated is a free log retrieval operation binding the contract event 0x63c3eef58f270166e6342563b916ac7cff473140da1299c27e9a24353101f24e.
//
// Solidity: event WorkflowActivated(uint256 id)
func (_Registry *RegistryFilterer) FilterWorkflowActivated(opts *bind.FilterOpts) (*RegistryWorkflowActivatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WorkflowActivated")
	if err != nil {
		return nil, err
	}
	return &RegistryWorkflowActivatedIterator{contract: _Registry.contract, event: "WorkflowActivated", logs: logs, sub: sub}, nil
}

// WatchWorkflowActivated is a free log subscription operation binding the contract event 0x63c3eef58f270166e6342563b916ac7cff473140da1299c27e9a24353101f24e.
//
// Solidity: event WorkflowActivated(uint256 id)
func (_Registry *RegistryFilterer) WatchWorkflowActivated(opts *bind.WatchOpts, sink chan<- *RegistryWorkflowActivated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WorkflowActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWorkflowActivated)
				if err := _Registry.contract.UnpackLog(event, "WorkflowActivated", log); err != nil {
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

// ParseWorkflowActivated is a log parse operation binding the contract event 0x63c3eef58f270166e6342563b916ac7cff473140da1299c27e9a24353101f24e.
//
// Solidity: event WorkflowActivated(uint256 id)
func (_Registry *RegistryFilterer) ParseWorkflowActivated(log types.Log) (*RegistryWorkflowActivated, error) {
	event := new(RegistryWorkflowActivated)
	if err := _Registry.contract.UnpackLog(event, "WorkflowActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWorkflowCancelledIterator is returned from FilterWorkflowCancelled and is used to iterate over the raw logs and unpacked data for WorkflowCancelled events raised by the Registry contract.
type RegistryWorkflowCancelledIterator struct {
	Event *RegistryWorkflowCancelled // Event containing the contract specifics and raw log

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
func (it *RegistryWorkflowCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWorkflowCancelled)
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
		it.Event = new(RegistryWorkflowCancelled)
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
func (it *RegistryWorkflowCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWorkflowCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWorkflowCancelled represents a WorkflowCancelled event raised by the Registry contract.
type RegistryWorkflowCancelled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWorkflowCancelled is a free log retrieval operation binding the contract event 0x5f576edaec67ed167152281963c0f265684ba9bf94b02241afbbccdd7c59e099.
//
// Solidity: event WorkflowCancelled(uint256 id)
func (_Registry *RegistryFilterer) FilterWorkflowCancelled(opts *bind.FilterOpts) (*RegistryWorkflowCancelledIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WorkflowCancelled")
	if err != nil {
		return nil, err
	}
	return &RegistryWorkflowCancelledIterator{contract: _Registry.contract, event: "WorkflowCancelled", logs: logs, sub: sub}, nil
}

// WatchWorkflowCancelled is a free log subscription operation binding the contract event 0x5f576edaec67ed167152281963c0f265684ba9bf94b02241afbbccdd7c59e099.
//
// Solidity: event WorkflowCancelled(uint256 id)
func (_Registry *RegistryFilterer) WatchWorkflowCancelled(opts *bind.WatchOpts, sink chan<- *RegistryWorkflowCancelled) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WorkflowCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWorkflowCancelled)
				if err := _Registry.contract.UnpackLog(event, "WorkflowCancelled", log); err != nil {
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

// ParseWorkflowCancelled is a log parse operation binding the contract event 0x5f576edaec67ed167152281963c0f265684ba9bf94b02241afbbccdd7c59e099.
//
// Solidity: event WorkflowCancelled(uint256 id)
func (_Registry *RegistryFilterer) ParseWorkflowCancelled(log types.Log) (*RegistryWorkflowCancelled, error) {
	event := new(RegistryWorkflowCancelled)
	if err := _Registry.contract.UnpackLog(event, "WorkflowCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWorkflowPausedIterator is returned from FilterWorkflowPaused and is used to iterate over the raw logs and unpacked data for WorkflowPaused events raised by the Registry contract.
type RegistryWorkflowPausedIterator struct {
	Event *RegistryWorkflowPaused // Event containing the contract specifics and raw log

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
func (it *RegistryWorkflowPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWorkflowPaused)
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
		it.Event = new(RegistryWorkflowPaused)
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
func (it *RegistryWorkflowPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWorkflowPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWorkflowPaused represents a WorkflowPaused event raised by the Registry contract.
type RegistryWorkflowPaused struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWorkflowPaused is a free log retrieval operation binding the contract event 0xe0b00c447dcbf94bbc6376b9d59fc6d4df9cacab900106a33efaff616af9ccd8.
//
// Solidity: event WorkflowPaused(uint256 id)
func (_Registry *RegistryFilterer) FilterWorkflowPaused(opts *bind.FilterOpts) (*RegistryWorkflowPausedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WorkflowPaused")
	if err != nil {
		return nil, err
	}
	return &RegistryWorkflowPausedIterator{contract: _Registry.contract, event: "WorkflowPaused", logs: logs, sub: sub}, nil
}

// WatchWorkflowPaused is a free log subscription operation binding the contract event 0xe0b00c447dcbf94bbc6376b9d59fc6d4df9cacab900106a33efaff616af9ccd8.
//
// Solidity: event WorkflowPaused(uint256 id)
func (_Registry *RegistryFilterer) WatchWorkflowPaused(opts *bind.WatchOpts, sink chan<- *RegistryWorkflowPaused) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WorkflowPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWorkflowPaused)
				if err := _Registry.contract.UnpackLog(event, "WorkflowPaused", log); err != nil {
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

// ParseWorkflowPaused is a log parse operation binding the contract event 0xe0b00c447dcbf94bbc6376b9d59fc6d4df9cacab900106a33efaff616af9ccd8.
//
// Solidity: event WorkflowPaused(uint256 id)
func (_Registry *RegistryFilterer) ParseWorkflowPaused(log types.Log) (*RegistryWorkflowPaused, error) {
	event := new(RegistryWorkflowPaused)
	if err := _Registry.contract.UnpackLog(event, "WorkflowPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWorkflowRegisteredIterator is returned from FilterWorkflowRegistered and is used to iterate over the raw logs and unpacked data for WorkflowRegistered events raised by the Registry contract.
type RegistryWorkflowRegisteredIterator struct {
	Event *RegistryWorkflowRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryWorkflowRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWorkflowRegistered)
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
		it.Event = new(RegistryWorkflowRegistered)
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
func (it *RegistryWorkflowRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWorkflowRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWorkflowRegistered represents a WorkflowRegistered event raised by the Registry contract.
type RegistryWorkflowRegistered struct {
	Owner common.Address
	Id    *big.Int
	Hash  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWorkflowRegistered is a free log retrieval operation binding the contract event 0x475b9f0c15578f6f25825e8f4794d63cc2b5a664944bc02a1767d37784a2a69c.
//
// Solidity: event WorkflowRegistered(address owner, uint256 id, bytes hash)
func (_Registry *RegistryFilterer) FilterWorkflowRegistered(opts *bind.FilterOpts) (*RegistryWorkflowRegisteredIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WorkflowRegistered")
	if err != nil {
		return nil, err
	}
	return &RegistryWorkflowRegisteredIterator{contract: _Registry.contract, event: "WorkflowRegistered", logs: logs, sub: sub}, nil
}

// WatchWorkflowRegistered is a free log subscription operation binding the contract event 0x475b9f0c15578f6f25825e8f4794d63cc2b5a664944bc02a1767d37784a2a69c.
//
// Solidity: event WorkflowRegistered(address owner, uint256 id, bytes hash)
func (_Registry *RegistryFilterer) WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *RegistryWorkflowRegistered) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WorkflowRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWorkflowRegistered)
				if err := _Registry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
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

// ParseWorkflowRegistered is a log parse operation binding the contract event 0x475b9f0c15578f6f25825e8f4794d63cc2b5a664944bc02a1767d37784a2a69c.
//
// Solidity: event WorkflowRegistered(address owner, uint256 id, bytes hash)
func (_Registry *RegistryFilterer) ParseWorkflowRegistered(log types.Log) (*RegistryWorkflowRegistered, error) {
	event := new(RegistryWorkflowRegistered)
	if err := _Registry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWorkflowResumedIterator is returned from FilterWorkflowResumed and is used to iterate over the raw logs and unpacked data for WorkflowResumed events raised by the Registry contract.
type RegistryWorkflowResumedIterator struct {
	Event *RegistryWorkflowResumed // Event containing the contract specifics and raw log

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
func (it *RegistryWorkflowResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWorkflowResumed)
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
		it.Event = new(RegistryWorkflowResumed)
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
func (it *RegistryWorkflowResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWorkflowResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWorkflowResumed represents a WorkflowResumed event raised by the Registry contract.
type RegistryWorkflowResumed struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWorkflowResumed is a free log retrieval operation binding the contract event 0xf9777a3853ed384d1ed462ecd8a0930375e3a4bbb8329a0a0af79a7ba6b5a2d9.
//
// Solidity: event WorkflowResumed(uint256 id)
func (_Registry *RegistryFilterer) FilterWorkflowResumed(opts *bind.FilterOpts) (*RegistryWorkflowResumedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WorkflowResumed")
	if err != nil {
		return nil, err
	}
	return &RegistryWorkflowResumedIterator{contract: _Registry.contract, event: "WorkflowResumed", logs: logs, sub: sub}, nil
}

// WatchWorkflowResumed is a free log subscription operation binding the contract event 0xf9777a3853ed384d1ed462ecd8a0930375e3a4bbb8329a0a0af79a7ba6b5a2d9.
//
// Solidity: event WorkflowResumed(uint256 id)
func (_Registry *RegistryFilterer) WatchWorkflowResumed(opts *bind.WatchOpts, sink chan<- *RegistryWorkflowResumed) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WorkflowResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWorkflowResumed)
				if err := _Registry.contract.UnpackLog(event, "WorkflowResumed", log); err != nil {
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

// ParseWorkflowResumed is a log parse operation binding the contract event 0xf9777a3853ed384d1ed462ecd8a0930375e3a4bbb8329a0a0af79a7ba6b5a2d9.
//
// Solidity: event WorkflowResumed(uint256 id)
func (_Registry *RegistryFilterer) ParseWorkflowResumed(log types.Log) (*RegistryWorkflowResumed, error) {
	event := new(RegistryWorkflowResumed)
	if err := _Registry.contract.UnpackLog(event, "WorkflowResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
