// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interfaces

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

// ISignerAddressMetaData contains all meta data concerning the ISignerAddress contract.
var ISignerAddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ISignerAddressABI is the input ABI used to generate the binding from.
// Deprecated: Use ISignerAddressMetaData.ABI instead.
var ISignerAddressABI = ISignerAddressMetaData.ABI

// ISignerAddress is an auto generated Go binding around an Ethereum contract.
type ISignerAddress struct {
	ISignerAddressCaller     // Read-only binding to the contract
	ISignerAddressTransactor // Write-only binding to the contract
	ISignerAddressFilterer   // Log filterer for contract events
}

// ISignerAddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISignerAddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignerAddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISignerAddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignerAddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISignerAddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISignerAddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISignerAddressSession struct {
	Contract     *ISignerAddress   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISignerAddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISignerAddressCallerSession struct {
	Contract *ISignerAddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ISignerAddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISignerAddressTransactorSession struct {
	Contract     *ISignerAddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ISignerAddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISignerAddressRaw struct {
	Contract *ISignerAddress // Generic contract binding to access the raw methods on
}

// ISignerAddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISignerAddressCallerRaw struct {
	Contract *ISignerAddressCaller // Generic read-only contract binding to access the raw methods on
}

// ISignerAddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISignerAddressTransactorRaw struct {
	Contract *ISignerAddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISignerAddress creates a new instance of ISignerAddress, bound to a specific deployed contract.
func NewISignerAddress(address common.Address, backend bind.ContractBackend) (*ISignerAddress, error) {
	contract, err := bindISignerAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISignerAddress{ISignerAddressCaller: ISignerAddressCaller{contract: contract}, ISignerAddressTransactor: ISignerAddressTransactor{contract: contract}, ISignerAddressFilterer: ISignerAddressFilterer{contract: contract}}, nil
}

// NewISignerAddressCaller creates a new read-only instance of ISignerAddress, bound to a specific deployed contract.
func NewISignerAddressCaller(address common.Address, caller bind.ContractCaller) (*ISignerAddressCaller, error) {
	contract, err := bindISignerAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISignerAddressCaller{contract: contract}, nil
}

// NewISignerAddressTransactor creates a new write-only instance of ISignerAddress, bound to a specific deployed contract.
func NewISignerAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*ISignerAddressTransactor, error) {
	contract, err := bindISignerAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISignerAddressTransactor{contract: contract}, nil
}

// NewISignerAddressFilterer creates a new log filterer instance of ISignerAddress, bound to a specific deployed contract.
func NewISignerAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*ISignerAddressFilterer, error) {
	contract, err := bindISignerAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISignerAddressFilterer{contract: contract}, nil
}

// bindISignerAddress binds a generic wrapper to an already deployed contract.
func bindISignerAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISignerAddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignerAddress *ISignerAddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignerAddress.Contract.ISignerAddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignerAddress *ISignerAddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignerAddress.Contract.ISignerAddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignerAddress *ISignerAddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignerAddress.Contract.ISignerAddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISignerAddress *ISignerAddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISignerAddress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISignerAddress *ISignerAddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISignerAddress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISignerAddress *ISignerAddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISignerAddress.Contract.contract.Transact(opts, method, params...)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_ISignerAddress *ISignerAddressCaller) GetSignerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ISignerAddress.contract.Call(opts, &out, "getSignerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_ISignerAddress *ISignerAddressSession) GetSignerAddress() (common.Address, error) {
	return _ISignerAddress.Contract.GetSignerAddress(&_ISignerAddress.CallOpts)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_ISignerAddress *ISignerAddressCallerSession) GetSignerAddress() (common.Address, error) {
	return _ISignerAddress.Contract.GetSignerAddress(&_ISignerAddress.CallOpts)
}
