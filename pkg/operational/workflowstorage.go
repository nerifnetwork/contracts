// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operational

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

// Workflow is an auto generated low-level Go binding around an user-defined struct.
type Workflow struct {
	Id         *big.Int
	Owner      common.Address
	Hash       []byte
	Status     uint8
	IsInternal bool
	TotalSpent *big.Int
}

// WorkflowStorageMetaData contains all meta data concerning the WorkflowStorage contract.
var WorkflowStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getWorkflow\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"},{\"internalType\":\"enumWorkflowStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isInternal\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structWorkflow\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWorkflows\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"},{\"internalType\":\"enumWorkflowStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isInternal\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structWorkflow[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getWorkflowsPerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"},{\"internalType\":\"enumWorkflowStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isInternal\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structWorkflow[]\",\"name\":\"_workflows\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"},{\"internalType\":\"enumWorkflowStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isInternal\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structWorkflow\",\"name\":\"workflow\",\"type\":\"tuple\"}],\"name\":\"mustAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"mustRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"hash\",\"type\":\"bytes\"},{\"internalType\":\"enumWorkflowStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isInternal\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structWorkflow\",\"name\":\"workflow\",\"type\":\"tuple\"}],\"name\":\"mustUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WorkflowStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use WorkflowStorageMetaData.ABI instead.
var WorkflowStorageABI = WorkflowStorageMetaData.ABI

// WorkflowStorage is an auto generated Go binding around an Ethereum contract.
type WorkflowStorage struct {
	WorkflowStorageCaller     // Read-only binding to the contract
	WorkflowStorageTransactor // Write-only binding to the contract
	WorkflowStorageFilterer   // Log filterer for contract events
}

// WorkflowStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type WorkflowStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkflowStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WorkflowStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkflowStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WorkflowStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WorkflowStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WorkflowStorageSession struct {
	Contract     *WorkflowStorage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WorkflowStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WorkflowStorageCallerSession struct {
	Contract *WorkflowStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WorkflowStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WorkflowStorageTransactorSession struct {
	Contract     *WorkflowStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WorkflowStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type WorkflowStorageRaw struct {
	Contract *WorkflowStorage // Generic contract binding to access the raw methods on
}

// WorkflowStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WorkflowStorageCallerRaw struct {
	Contract *WorkflowStorageCaller // Generic read-only contract binding to access the raw methods on
}

// WorkflowStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WorkflowStorageTransactorRaw struct {
	Contract *WorkflowStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWorkflowStorage creates a new instance of WorkflowStorage, bound to a specific deployed contract.
func NewWorkflowStorage(address common.Address, backend bind.ContractBackend) (*WorkflowStorage, error) {
	contract, err := bindWorkflowStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WorkflowStorage{WorkflowStorageCaller: WorkflowStorageCaller{contract: contract}, WorkflowStorageTransactor: WorkflowStorageTransactor{contract: contract}, WorkflowStorageFilterer: WorkflowStorageFilterer{contract: contract}}, nil
}

// NewWorkflowStorageCaller creates a new read-only instance of WorkflowStorage, bound to a specific deployed contract.
func NewWorkflowStorageCaller(address common.Address, caller bind.ContractCaller) (*WorkflowStorageCaller, error) {
	contract, err := bindWorkflowStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowStorageCaller{contract: contract}, nil
}

// NewWorkflowStorageTransactor creates a new write-only instance of WorkflowStorage, bound to a specific deployed contract.
func NewWorkflowStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*WorkflowStorageTransactor, error) {
	contract, err := bindWorkflowStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WorkflowStorageTransactor{contract: contract}, nil
}

// NewWorkflowStorageFilterer creates a new log filterer instance of WorkflowStorage, bound to a specific deployed contract.
func NewWorkflowStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*WorkflowStorageFilterer, error) {
	contract, err := bindWorkflowStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WorkflowStorageFilterer{contract: contract}, nil
}

// bindWorkflowStorage binds a generic wrapper to an already deployed contract.
func bindWorkflowStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WorkflowStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkflowStorage *WorkflowStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowStorage.Contract.WorkflowStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkflowStorage *WorkflowStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.WorkflowStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkflowStorage *WorkflowStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.WorkflowStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WorkflowStorage *WorkflowStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WorkflowStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WorkflowStorage *WorkflowStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WorkflowStorage *WorkflowStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0xc34052e0.
//
// Solidity: function contains(uint256 id) view returns(bool)
func (_WorkflowStorage *WorkflowStorageCaller) Contains(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _WorkflowStorage.contract.Call(opts, &out, "contains", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0xc34052e0.
//
// Solidity: function contains(uint256 id) view returns(bool)
func (_WorkflowStorage *WorkflowStorageSession) Contains(id *big.Int) (bool, error) {
	return _WorkflowStorage.Contract.Contains(&_WorkflowStorage.CallOpts, id)
}

// Contains is a free data retrieval call binding the contract method 0xc34052e0.
//
// Solidity: function contains(uint256 id) view returns(bool)
func (_WorkflowStorage *WorkflowStorageCallerSession) Contains(id *big.Int) (bool, error) {
	return _WorkflowStorage.Contract.Contains(&_WorkflowStorage.CallOpts, id)
}

// GetWorkflow is a free data retrieval call binding the contract method 0xeec7b03c.
//
// Solidity: function getWorkflow(uint256 id) view returns((uint256,address,bytes,uint8,bool,uint256))
func (_WorkflowStorage *WorkflowStorageCaller) GetWorkflow(opts *bind.CallOpts, id *big.Int) (Workflow, error) {
	var out []interface{}
	err := _WorkflowStorage.contract.Call(opts, &out, "getWorkflow", id)

	if err != nil {
		return *new(Workflow), err
	}

	out0 := *abi.ConvertType(out[0], new(Workflow)).(*Workflow)

	return out0, err

}

// GetWorkflow is a free data retrieval call binding the contract method 0xeec7b03c.
//
// Solidity: function getWorkflow(uint256 id) view returns((uint256,address,bytes,uint8,bool,uint256))
func (_WorkflowStorage *WorkflowStorageSession) GetWorkflow(id *big.Int) (Workflow, error) {
	return _WorkflowStorage.Contract.GetWorkflow(&_WorkflowStorage.CallOpts, id)
}

// GetWorkflow is a free data retrieval call binding the contract method 0xeec7b03c.
//
// Solidity: function getWorkflow(uint256 id) view returns((uint256,address,bytes,uint8,bool,uint256))
func (_WorkflowStorage *WorkflowStorageCallerSession) GetWorkflow(id *big.Int) (Workflow, error) {
	return _WorkflowStorage.Contract.GetWorkflow(&_WorkflowStorage.CallOpts, id)
}

// GetWorkflows is a free data retrieval call binding the contract method 0x7099588f.
//
// Solidity: function getWorkflows() view returns((uint256,address,bytes,uint8,bool,uint256)[])
func (_WorkflowStorage *WorkflowStorageCaller) GetWorkflows(opts *bind.CallOpts) ([]Workflow, error) {
	var out []interface{}
	err := _WorkflowStorage.contract.Call(opts, &out, "getWorkflows")

	if err != nil {
		return *new([]Workflow), err
	}

	out0 := *abi.ConvertType(out[0], new([]Workflow)).(*[]Workflow)

	return out0, err

}

// GetWorkflows is a free data retrieval call binding the contract method 0x7099588f.
//
// Solidity: function getWorkflows() view returns((uint256,address,bytes,uint8,bool,uint256)[])
func (_WorkflowStorage *WorkflowStorageSession) GetWorkflows() ([]Workflow, error) {
	return _WorkflowStorage.Contract.GetWorkflows(&_WorkflowStorage.CallOpts)
}

// GetWorkflows is a free data retrieval call binding the contract method 0x7099588f.
//
// Solidity: function getWorkflows() view returns((uint256,address,bytes,uint8,bool,uint256)[])
func (_WorkflowStorage *WorkflowStorageCallerSession) GetWorkflows() ([]Workflow, error) {
	return _WorkflowStorage.Contract.GetWorkflows(&_WorkflowStorage.CallOpts)
}

// GetWorkflowsPerAddress is a free data retrieval call binding the contract method 0xf350ea26.
//
// Solidity: function getWorkflowsPerAddress(address owner) view returns(uint256)
func (_WorkflowStorage *WorkflowStorageCaller) GetWorkflowsPerAddress(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowStorage.contract.Call(opts, &out, "getWorkflowsPerAddress", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWorkflowsPerAddress is a free data retrieval call binding the contract method 0xf350ea26.
//
// Solidity: function getWorkflowsPerAddress(address owner) view returns(uint256)
func (_WorkflowStorage *WorkflowStorageSession) GetWorkflowsPerAddress(owner common.Address) (*big.Int, error) {
	return _WorkflowStorage.Contract.GetWorkflowsPerAddress(&_WorkflowStorage.CallOpts, owner)
}

// GetWorkflowsPerAddress is a free data retrieval call binding the contract method 0xf350ea26.
//
// Solidity: function getWorkflowsPerAddress(address owner) view returns(uint256)
func (_WorkflowStorage *WorkflowStorageCallerSession) GetWorkflowsPerAddress(owner common.Address) (*big.Int, error) {
	return _WorkflowStorage.Contract.GetWorkflowsPerAddress(&_WorkflowStorage.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkflowStorage *WorkflowStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkflowStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkflowStorage *WorkflowStorageSession) Owner() (common.Address, error) {
	return _WorkflowStorage.Contract.Owner(&_WorkflowStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WorkflowStorage *WorkflowStorageCallerSession) Owner() (common.Address, error) {
	return _WorkflowStorage.Contract.Owner(&_WorkflowStorage.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WorkflowStorage *WorkflowStorageCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WorkflowStorage.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WorkflowStorage *WorkflowStorageSession) Registry() (common.Address, error) {
	return _WorkflowStorage.Contract.Registry(&_WorkflowStorage.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_WorkflowStorage *WorkflowStorageCallerSession) Registry() (common.Address, error) {
	return _WorkflowStorage.Contract.Registry(&_WorkflowStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_WorkflowStorage *WorkflowStorageCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WorkflowStorage.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_WorkflowStorage *WorkflowStorageSession) Size() (*big.Int, error) {
	return _WorkflowStorage.Contract.Size(&_WorkflowStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_WorkflowStorage *WorkflowStorageCallerSession) Size() (*big.Int, error) {
	return _WorkflowStorage.Contract.Size(&_WorkflowStorage.CallOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_WorkflowStorage *WorkflowStorageTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_WorkflowStorage *WorkflowStorageSession) Clear() (*types.Transaction, error) {
	return _WorkflowStorage.Contract.Clear(&_WorkflowStorage.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_WorkflowStorage *WorkflowStorageTransactorSession) Clear() (*types.Transaction, error) {
	return _WorkflowStorage.Contract.Clear(&_WorkflowStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ef93524.
//
// Solidity: function initialize(address _registry, (uint256,address,bytes,uint8,bool,uint256)[] _workflows) returns()
func (_WorkflowStorage *WorkflowStorageTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _workflows []Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "initialize", _registry, _workflows)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ef93524.
//
// Solidity: function initialize(address _registry, (uint256,address,bytes,uint8,bool,uint256)[] _workflows) returns()
func (_WorkflowStorage *WorkflowStorageSession) Initialize(_registry common.Address, _workflows []Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.Initialize(&_WorkflowStorage.TransactOpts, _registry, _workflows)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ef93524.
//
// Solidity: function initialize(address _registry, (uint256,address,bytes,uint8,bool,uint256)[] _workflows) returns()
func (_WorkflowStorage *WorkflowStorageTransactorSession) Initialize(_registry common.Address, _workflows []Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.Initialize(&_WorkflowStorage.TransactOpts, _registry, _workflows)
}

// MustAdd is a paid mutator transaction binding the contract method 0x4f033ba7.
//
// Solidity: function mustAdd((uint256,address,bytes,uint8,bool,uint256) workflow) returns()
func (_WorkflowStorage *WorkflowStorageTransactor) MustAdd(opts *bind.TransactOpts, workflow Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "mustAdd", workflow)
}

// MustAdd is a paid mutator transaction binding the contract method 0x4f033ba7.
//
// Solidity: function mustAdd((uint256,address,bytes,uint8,bool,uint256) workflow) returns()
func (_WorkflowStorage *WorkflowStorageSession) MustAdd(workflow Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.MustAdd(&_WorkflowStorage.TransactOpts, workflow)
}

// MustAdd is a paid mutator transaction binding the contract method 0x4f033ba7.
//
// Solidity: function mustAdd((uint256,address,bytes,uint8,bool,uint256) workflow) returns()
func (_WorkflowStorage *WorkflowStorageTransactorSession) MustAdd(workflow Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.MustAdd(&_WorkflowStorage.TransactOpts, workflow)
}

// MustRemove is a paid mutator transaction binding the contract method 0xfe14018e.
//
// Solidity: function mustRemove(uint256 id) returns()
func (_WorkflowStorage *WorkflowStorageTransactor) MustRemove(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "mustRemove", id)
}

// MustRemove is a paid mutator transaction binding the contract method 0xfe14018e.
//
// Solidity: function mustRemove(uint256 id) returns()
func (_WorkflowStorage *WorkflowStorageSession) MustRemove(id *big.Int) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.MustRemove(&_WorkflowStorage.TransactOpts, id)
}

// MustRemove is a paid mutator transaction binding the contract method 0xfe14018e.
//
// Solidity: function mustRemove(uint256 id) returns()
func (_WorkflowStorage *WorkflowStorageTransactorSession) MustRemove(id *big.Int) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.MustRemove(&_WorkflowStorage.TransactOpts, id)
}

// MustUpdate is a paid mutator transaction binding the contract method 0x470f9c78.
//
// Solidity: function mustUpdate((uint256,address,bytes,uint8,bool,uint256) workflow) returns()
func (_WorkflowStorage *WorkflowStorageTransactor) MustUpdate(opts *bind.TransactOpts, workflow Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "mustUpdate", workflow)
}

// MustUpdate is a paid mutator transaction binding the contract method 0x470f9c78.
//
// Solidity: function mustUpdate((uint256,address,bytes,uint8,bool,uint256) workflow) returns()
func (_WorkflowStorage *WorkflowStorageSession) MustUpdate(workflow Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.MustUpdate(&_WorkflowStorage.TransactOpts, workflow)
}

// MustUpdate is a paid mutator transaction binding the contract method 0x470f9c78.
//
// Solidity: function mustUpdate((uint256,address,bytes,uint8,bool,uint256) workflow) returns()
func (_WorkflowStorage *WorkflowStorageTransactorSession) MustUpdate(workflow Workflow) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.MustUpdate(&_WorkflowStorage.TransactOpts, workflow)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkflowStorage *WorkflowStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkflowStorage *WorkflowStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkflowStorage.Contract.RenounceOwnership(&_WorkflowStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WorkflowStorage *WorkflowStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WorkflowStorage.Contract.RenounceOwnership(&_WorkflowStorage.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WorkflowStorage *WorkflowStorageTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WorkflowStorage *WorkflowStorageSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.SetRegistry(&_WorkflowStorage.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_WorkflowStorage *WorkflowStorageTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.SetRegistry(&_WorkflowStorage.TransactOpts, _registry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkflowStorage *WorkflowStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WorkflowStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkflowStorage *WorkflowStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.TransferOwnership(&_WorkflowStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WorkflowStorage *WorkflowStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WorkflowStorage.Contract.TransferOwnership(&_WorkflowStorage.TransactOpts, newOwner)
}

// WorkflowStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WorkflowStorage contract.
type WorkflowStorageInitializedIterator struct {
	Event *WorkflowStorageInitialized // Event containing the contract specifics and raw log

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
func (it *WorkflowStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowStorageInitialized)
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
		it.Event = new(WorkflowStorageInitialized)
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
func (it *WorkflowStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkflowStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkflowStorageInitialized represents a Initialized event raised by the WorkflowStorage contract.
type WorkflowStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkflowStorage *WorkflowStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*WorkflowStorageInitializedIterator, error) {

	logs, sub, err := _WorkflowStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WorkflowStorageInitializedIterator{contract: _WorkflowStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WorkflowStorage *WorkflowStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WorkflowStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _WorkflowStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkflowStorageInitialized)
				if err := _WorkflowStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WorkflowStorage *WorkflowStorageFilterer) ParseInitialized(log types.Log) (*WorkflowStorageInitialized, error) {
	event := new(WorkflowStorageInitialized)
	if err := _WorkflowStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WorkflowStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WorkflowStorage contract.
type WorkflowStorageOwnershipTransferredIterator struct {
	Event *WorkflowStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WorkflowStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WorkflowStorageOwnershipTransferred)
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
		it.Event = new(WorkflowStorageOwnershipTransferred)
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
func (it *WorkflowStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WorkflowStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WorkflowStorageOwnershipTransferred represents a OwnershipTransferred event raised by the WorkflowStorage contract.
type WorkflowStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkflowStorage *WorkflowStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WorkflowStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkflowStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WorkflowStorageOwnershipTransferredIterator{contract: _WorkflowStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkflowStorage *WorkflowStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WorkflowStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WorkflowStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WorkflowStorageOwnershipTransferred)
				if err := _WorkflowStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WorkflowStorage *WorkflowStorageFilterer) ParseOwnershipTransferred(log types.Log) (*WorkflowStorageOwnershipTransferred, error) {
	event := new(WorkflowStorageOwnershipTransferred)
	if err := _WorkflowStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
