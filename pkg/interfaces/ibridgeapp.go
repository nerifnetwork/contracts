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

// IBridgeAppMetaData contains all meta data concerning the IBridgeApp contract.
var IBridgeAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"bridgeAppAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"revertSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBridgeAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgeAppMetaData.ABI instead.
var IBridgeAppABI = IBridgeAppMetaData.ABI

// IBridgeApp is an auto generated Go binding around an Ethereum contract.
type IBridgeApp struct {
	IBridgeAppCaller     // Read-only binding to the contract
	IBridgeAppTransactor // Write-only binding to the contract
	IBridgeAppFilterer   // Log filterer for contract events
}

// IBridgeAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeAppSession struct {
	Contract     *IBridgeApp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeAppCallerSession struct {
	Contract *IBridgeAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IBridgeAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeAppTransactorSession struct {
	Contract     *IBridgeAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IBridgeAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeAppRaw struct {
	Contract *IBridgeApp // Generic contract binding to access the raw methods on
}

// IBridgeAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeAppCallerRaw struct {
	Contract *IBridgeAppCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeAppTransactorRaw struct {
	Contract *IBridgeAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgeApp creates a new instance of IBridgeApp, bound to a specific deployed contract.
func NewIBridgeApp(address common.Address, backend bind.ContractBackend) (*IBridgeApp, error) {
	contract, err := bindIBridgeApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgeApp{IBridgeAppCaller: IBridgeAppCaller{contract: contract}, IBridgeAppTransactor: IBridgeAppTransactor{contract: contract}, IBridgeAppFilterer: IBridgeAppFilterer{contract: contract}}, nil
}

// NewIBridgeAppCaller creates a new read-only instance of IBridgeApp, bound to a specific deployed contract.
func NewIBridgeAppCaller(address common.Address, caller bind.ContractCaller) (*IBridgeAppCaller, error) {
	contract, err := bindIBridgeApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeAppCaller{contract: contract}, nil
}

// NewIBridgeAppTransactor creates a new write-only instance of IBridgeApp, bound to a specific deployed contract.
func NewIBridgeAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeAppTransactor, error) {
	contract, err := bindIBridgeApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeAppTransactor{contract: contract}, nil
}

// NewIBridgeAppFilterer creates a new log filterer instance of IBridgeApp, bound to a specific deployed contract.
func NewIBridgeAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeAppFilterer, error) {
	contract, err := bindIBridgeApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeAppFilterer{contract: contract}, nil
}

// bindIBridgeApp binds a generic wrapper to an already deployed contract.
func bindIBridgeApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeApp *IBridgeAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeApp.Contract.IBridgeAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeApp *IBridgeAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeApp.Contract.IBridgeAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeApp *IBridgeAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeApp.Contract.IBridgeAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeApp *IBridgeAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeApp *IBridgeAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeApp *IBridgeAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeApp.Contract.contract.Transact(opts, method, params...)
}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_IBridgeApp *IBridgeAppCaller) BridgeAppAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBridgeApp.contract.Call(opts, &out, "bridgeAppAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_IBridgeApp *IBridgeAppSession) BridgeAppAddress() (common.Address, error) {
	return _IBridgeApp.Contract.BridgeAppAddress(&_IBridgeApp.CallOpts)
}

// BridgeAppAddress is a free data retrieval call binding the contract method 0xd7ac8777.
//
// Solidity: function bridgeAppAddress() view returns(address)
func (_IBridgeApp *IBridgeAppCallerSession) BridgeAppAddress() (common.Address, error) {
	return _IBridgeApp.Contract.BridgeAppAddress(&_IBridgeApp.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 _sourceChain, bytes _data) returns()
func (_IBridgeApp *IBridgeAppTransactor) Execute(opts *bind.TransactOpts, _sourceChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IBridgeApp.contract.Transact(opts, "execute", _sourceChain, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 _sourceChain, bytes _data) returns()
func (_IBridgeApp *IBridgeAppSession) Execute(_sourceChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IBridgeApp.Contract.Execute(&_IBridgeApp.TransactOpts, _sourceChain, _data)
}

// Execute is a paid mutator transaction binding the contract method 0x59efcb15.
//
// Solidity: function execute(uint256 _sourceChain, bytes _data) returns()
func (_IBridgeApp *IBridgeAppTransactorSession) Execute(_sourceChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IBridgeApp.Contract.Execute(&_IBridgeApp.TransactOpts, _sourceChain, _data)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 _destinationChain, bytes _data) returns()
func (_IBridgeApp *IBridgeAppTransactor) RevertSend(opts *bind.TransactOpts, _destinationChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IBridgeApp.contract.Transact(opts, "revertSend", _destinationChain, _data)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 _destinationChain, bytes _data) returns()
func (_IBridgeApp *IBridgeAppSession) RevertSend(_destinationChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IBridgeApp.Contract.RevertSend(&_IBridgeApp.TransactOpts, _destinationChain, _data)
}

// RevertSend is a paid mutator transaction binding the contract method 0x0d788db0.
//
// Solidity: function revertSend(uint256 _destinationChain, bytes _data) returns()
func (_IBridgeApp *IBridgeAppTransactorSession) RevertSend(_destinationChain *big.Int, _data []byte) (*types.Transaction, error) {
	return _IBridgeApp.Contract.RevertSend(&_IBridgeApp.TransactOpts, _destinationChain, _data)
}
