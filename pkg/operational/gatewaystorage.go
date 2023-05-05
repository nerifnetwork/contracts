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

// GatewayStorageGateway is an auto generated low-level Go binding around an user-defined struct.
type GatewayStorageGateway struct {
	Gateway common.Address
	Owner   common.Address
}

// GatewayStorageMetaData contains all meta data concerning the GatewayStorage contract.
var GatewayStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getGateway\",\"outputs\":[{\"internalType\":\"contractIGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGateways\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIGateway\",\"name\":\"gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structGatewayStorage.Gateway[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIGateway\",\"name\":\"gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structGatewayStorage.Gateway[]\",\"name\":\"_gateways\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mustRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"mustSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GatewayStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayStorageMetaData.ABI instead.
var GatewayStorageABI = GatewayStorageMetaData.ABI

// GatewayStorage is an auto generated Go binding around an Ethereum contract.
type GatewayStorage struct {
	GatewayStorageCaller     // Read-only binding to the contract
	GatewayStorageTransactor // Write-only binding to the contract
	GatewayStorageFilterer   // Log filterer for contract events
}

// GatewayStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayStorageSession struct {
	Contract     *GatewayStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayStorageCallerSession struct {
	Contract *GatewayStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GatewayStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayStorageTransactorSession struct {
	Contract     *GatewayStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GatewayStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayStorageRaw struct {
	Contract *GatewayStorage // Generic contract binding to access the raw methods on
}

// GatewayStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayStorageCallerRaw struct {
	Contract *GatewayStorageCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayStorageTransactorRaw struct {
	Contract *GatewayStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayStorage creates a new instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorage(address common.Address, backend bind.ContractBackend) (*GatewayStorage, error) {
	contract, err := bindGatewayStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayStorage{GatewayStorageCaller: GatewayStorageCaller{contract: contract}, GatewayStorageTransactor: GatewayStorageTransactor{contract: contract}, GatewayStorageFilterer: GatewayStorageFilterer{contract: contract}}, nil
}

// NewGatewayStorageCaller creates a new read-only instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorageCaller(address common.Address, caller bind.ContractCaller) (*GatewayStorageCaller, error) {
	contract, err := bindGatewayStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageCaller{contract: contract}, nil
}

// NewGatewayStorageTransactor creates a new write-only instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayStorageTransactor, error) {
	contract, err := bindGatewayStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageTransactor{contract: contract}, nil
}

// NewGatewayStorageFilterer creates a new log filterer instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayStorageFilterer, error) {
	contract, err := bindGatewayStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageFilterer{contract: contract}, nil
}

// bindGatewayStorage binds a generic wrapper to an already deployed contract.
func bindGatewayStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatewayStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayStorage *GatewayStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayStorage.Contract.GatewayStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayStorage *GatewayStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.Contract.GatewayStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayStorage *GatewayStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayStorage.Contract.GatewayStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayStorage *GatewayStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayStorage *GatewayStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayStorage *GatewayStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayStorage.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address owner) view returns(bool)
func (_GatewayStorage *GatewayStorageCaller) Contains(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "contains", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address owner) view returns(bool)
func (_GatewayStorage *GatewayStorageSession) Contains(owner common.Address) (bool, error) {
	return _GatewayStorage.Contract.Contains(&_GatewayStorage.CallOpts, owner)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address owner) view returns(bool)
func (_GatewayStorage *GatewayStorageCallerSession) Contains(owner common.Address) (bool, error) {
	return _GatewayStorage.Contract.Contains(&_GatewayStorage.CallOpts, owner)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address)
func (_GatewayStorage *GatewayStorageCaller) GetGateway(opts *bind.CallOpts, owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "getGateway", owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address)
func (_GatewayStorage *GatewayStorageSession) GetGateway(owner common.Address) (common.Address, error) {
	return _GatewayStorage.Contract.GetGateway(&_GatewayStorage.CallOpts, owner)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address)
func (_GatewayStorage *GatewayStorageCallerSession) GetGateway(owner common.Address) (common.Address, error) {
	return _GatewayStorage.Contract.GetGateway(&_GatewayStorage.CallOpts, owner)
}

// GetGateways is a free data retrieval call binding the contract method 0xd82778ce.
//
// Solidity: function getGateways() view returns((address,address)[])
func (_GatewayStorage *GatewayStorageCaller) GetGateways(opts *bind.CallOpts) ([]GatewayStorageGateway, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "getGateways")

	if err != nil {
		return *new([]GatewayStorageGateway), err
	}

	out0 := *abi.ConvertType(out[0], new([]GatewayStorageGateway)).(*[]GatewayStorageGateway)

	return out0, err

}

// GetGateways is a free data retrieval call binding the contract method 0xd82778ce.
//
// Solidity: function getGateways() view returns((address,address)[])
func (_GatewayStorage *GatewayStorageSession) GetGateways() ([]GatewayStorageGateway, error) {
	return _GatewayStorage.Contract.GetGateways(&_GatewayStorage.CallOpts)
}

// GetGateways is a free data retrieval call binding the contract method 0xd82778ce.
//
// Solidity: function getGateways() view returns((address,address)[])
func (_GatewayStorage *GatewayStorageCallerSession) GetGateways() ([]GatewayStorageGateway, error) {
	return _GatewayStorage.Contract.GetGateways(&_GatewayStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayStorage *GatewayStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayStorage *GatewayStorageSession) Owner() (common.Address, error) {
	return _GatewayStorage.Contract.Owner(&_GatewayStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayStorage *GatewayStorageCallerSession) Owner() (common.Address, error) {
	return _GatewayStorage.Contract.Owner(&_GatewayStorage.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayStorage *GatewayStorageCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayStorage *GatewayStorageSession) Registry() (common.Address, error) {
	return _GatewayStorage.Contract.Registry(&_GatewayStorage.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayStorage *GatewayStorageCallerSession) Registry() (common.Address, error) {
	return _GatewayStorage.Contract.Registry(&_GatewayStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_GatewayStorage *GatewayStorageCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_GatewayStorage *GatewayStorageSession) Size() (*big.Int, error) {
	return _GatewayStorage.Contract.Size(&_GatewayStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_GatewayStorage *GatewayStorageCallerSession) Size() (*big.Int, error) {
	return _GatewayStorage.Contract.Size(&_GatewayStorage.CallOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_GatewayStorage *GatewayStorageTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_GatewayStorage *GatewayStorageSession) Clear() (*types.Transaction, error) {
	return _GatewayStorage.Contract.Clear(&_GatewayStorage.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_GatewayStorage *GatewayStorageTransactorSession) Clear() (*types.Transaction, error) {
	return _GatewayStorage.Contract.Clear(&_GatewayStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x78e70696.
//
// Solidity: function initialize(address _registry, (address,address)[] _gateways) returns()
func (_GatewayStorage *GatewayStorageTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _gateways []GatewayStorageGateway) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "initialize", _registry, _gateways)
}

// Initialize is a paid mutator transaction binding the contract method 0x78e70696.
//
// Solidity: function initialize(address _registry, (address,address)[] _gateways) returns()
func (_GatewayStorage *GatewayStorageSession) Initialize(_registry common.Address, _gateways []GatewayStorageGateway) (*types.Transaction, error) {
	return _GatewayStorage.Contract.Initialize(&_GatewayStorage.TransactOpts, _registry, _gateways)
}

// Initialize is a paid mutator transaction binding the contract method 0x78e70696.
//
// Solidity: function initialize(address _registry, (address,address)[] _gateways) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) Initialize(_registry common.Address, _gateways []GatewayStorageGateway) (*types.Transaction, error) {
	return _GatewayStorage.Contract.Initialize(&_GatewayStorage.TransactOpts, _registry, _gateways)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address owner) returns()
func (_GatewayStorage *GatewayStorageTransactor) MustRemove(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "mustRemove", owner)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address owner) returns()
func (_GatewayStorage *GatewayStorageSession) MustRemove(owner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustRemove(&_GatewayStorage.TransactOpts, owner)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address owner) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) MustRemove(owner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustRemove(&_GatewayStorage.TransactOpts, owner)
}

// MustSet is a paid mutator transaction binding the contract method 0xf7beb8c7.
//
// Solidity: function mustSet(address owner, address gateway) returns()
func (_GatewayStorage *GatewayStorageTransactor) MustSet(opts *bind.TransactOpts, owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "mustSet", owner, gateway)
}

// MustSet is a paid mutator transaction binding the contract method 0xf7beb8c7.
//
// Solidity: function mustSet(address owner, address gateway) returns()
func (_GatewayStorage *GatewayStorageSession) MustSet(owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustSet(&_GatewayStorage.TransactOpts, owner, gateway)
}

// MustSet is a paid mutator transaction binding the contract method 0xf7beb8c7.
//
// Solidity: function mustSet(address owner, address gateway) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) MustSet(owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustSet(&_GatewayStorage.TransactOpts, owner, gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayStorage *GatewayStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayStorage *GatewayStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayStorage.Contract.RenounceOwnership(&_GatewayStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayStorage *GatewayStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayStorage.Contract.RenounceOwnership(&_GatewayStorage.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_GatewayStorage *GatewayStorageTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_GatewayStorage *GatewayStorageSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.SetRegistry(&_GatewayStorage.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.SetRegistry(&_GatewayStorage.TransactOpts, _registry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayStorage *GatewayStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayStorage *GatewayStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.TransferOwnership(&_GatewayStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.TransferOwnership(&_GatewayStorage.TransactOpts, newOwner)
}

// GatewayStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayStorage contract.
type GatewayStorageInitializedIterator struct {
	Event *GatewayStorageInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStorageInitialized)
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
		it.Event = new(GatewayStorageInitialized)
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
func (it *GatewayStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStorageInitialized represents a Initialized event raised by the GatewayStorage contract.
type GatewayStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayStorage *GatewayStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayStorageInitializedIterator, error) {

	logs, sub, err := _GatewayStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayStorageInitializedIterator{contract: _GatewayStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayStorage *GatewayStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStorageInitialized)
				if err := _GatewayStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayStorage *GatewayStorageFilterer) ParseInitialized(log types.Log) (*GatewayStorageInitialized, error) {
	event := new(GatewayStorageInitialized)
	if err := _GatewayStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayStorage contract.
type GatewayStorageOwnershipTransferredIterator struct {
	Event *GatewayStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStorageOwnershipTransferred)
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
		it.Event = new(GatewayStorageOwnershipTransferred)
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
func (it *GatewayStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStorageOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayStorage contract.
type GatewayStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayStorage *GatewayStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageOwnershipTransferredIterator{contract: _GatewayStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayStorage *GatewayStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStorageOwnershipTransferred)
				if err := _GatewayStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GatewayStorage *GatewayStorageFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayStorageOwnershipTransferred, error) {
	event := new(GatewayStorageOwnershipTransferred)
	if err := _GatewayStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
