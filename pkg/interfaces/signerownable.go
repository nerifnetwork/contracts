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

// SignerOwnableMetaData contains all meta data concerning the SignerOwnable contract.
var SignerOwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SignerOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use SignerOwnableMetaData.ABI instead.
var SignerOwnableABI = SignerOwnableMetaData.ABI

// SignerOwnable is an auto generated Go binding around an Ethereum contract.
type SignerOwnable struct {
	SignerOwnableCaller     // Read-only binding to the contract
	SignerOwnableTransactor // Write-only binding to the contract
	SignerOwnableFilterer   // Log filterer for contract events
}

// SignerOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignerOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignerOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignerOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignerOwnableSession struct {
	Contract     *SignerOwnable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignerOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignerOwnableCallerSession struct {
	Contract *SignerOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SignerOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignerOwnableTransactorSession struct {
	Contract     *SignerOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SignerOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignerOwnableRaw struct {
	Contract *SignerOwnable // Generic contract binding to access the raw methods on
}

// SignerOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignerOwnableCallerRaw struct {
	Contract *SignerOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// SignerOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignerOwnableTransactorRaw struct {
	Contract *SignerOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignerOwnable creates a new instance of SignerOwnable, bound to a specific deployed contract.
func NewSignerOwnable(address common.Address, backend bind.ContractBackend) (*SignerOwnable, error) {
	contract, err := bindSignerOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignerOwnable{SignerOwnableCaller: SignerOwnableCaller{contract: contract}, SignerOwnableTransactor: SignerOwnableTransactor{contract: contract}, SignerOwnableFilterer: SignerOwnableFilterer{contract: contract}}, nil
}

// NewSignerOwnableCaller creates a new read-only instance of SignerOwnable, bound to a specific deployed contract.
func NewSignerOwnableCaller(address common.Address, caller bind.ContractCaller) (*SignerOwnableCaller, error) {
	contract, err := bindSignerOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignerOwnableCaller{contract: contract}, nil
}

// NewSignerOwnableTransactor creates a new write-only instance of SignerOwnable, bound to a specific deployed contract.
func NewSignerOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*SignerOwnableTransactor, error) {
	contract, err := bindSignerOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignerOwnableTransactor{contract: contract}, nil
}

// NewSignerOwnableFilterer creates a new log filterer instance of SignerOwnable, bound to a specific deployed contract.
func NewSignerOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*SignerOwnableFilterer, error) {
	contract, err := bindSignerOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignerOwnableFilterer{contract: contract}, nil
}

// bindSignerOwnable binds a generic wrapper to an already deployed contract.
func bindSignerOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SignerOwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignerOwnable *SignerOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignerOwnable.Contract.SignerOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignerOwnable *SignerOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignerOwnable.Contract.SignerOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignerOwnable *SignerOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignerOwnable.Contract.SignerOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignerOwnable *SignerOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignerOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignerOwnable *SignerOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignerOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignerOwnable *SignerOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignerOwnable.Contract.contract.Transact(opts, method, params...)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SignerOwnable *SignerOwnableCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SignerOwnable.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SignerOwnable *SignerOwnableSession) SignerGetter() (common.Address, error) {
	return _SignerOwnable.Contract.SignerGetter(&_SignerOwnable.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SignerOwnable *SignerOwnableCallerSession) SignerGetter() (common.Address, error) {
	return _SignerOwnable.Contract.SignerGetter(&_SignerOwnable.CallOpts)
}
