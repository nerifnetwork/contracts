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

// IBridgeMediatorMetaData contains all meta data concerning the IBridgeMediator contract.
var IBridgeMediatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"mediate\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBridgeMediatorABI is the input ABI used to generate the binding from.
// Deprecated: Use IBridgeMediatorMetaData.ABI instead.
var IBridgeMediatorABI = IBridgeMediatorMetaData.ABI

// IBridgeMediator is an auto generated Go binding around an Ethereum contract.
type IBridgeMediator struct {
	IBridgeMediatorCaller     // Read-only binding to the contract
	IBridgeMediatorTransactor // Write-only binding to the contract
	IBridgeMediatorFilterer   // Log filterer for contract events
}

// IBridgeMediatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBridgeMediatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeMediatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBridgeMediatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeMediatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBridgeMediatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBridgeMediatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBridgeMediatorSession struct {
	Contract     *IBridgeMediator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBridgeMediatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBridgeMediatorCallerSession struct {
	Contract *IBridgeMediatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IBridgeMediatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBridgeMediatorTransactorSession struct {
	Contract     *IBridgeMediatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IBridgeMediatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBridgeMediatorRaw struct {
	Contract *IBridgeMediator // Generic contract binding to access the raw methods on
}

// IBridgeMediatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBridgeMediatorCallerRaw struct {
	Contract *IBridgeMediatorCaller // Generic read-only contract binding to access the raw methods on
}

// IBridgeMediatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBridgeMediatorTransactorRaw struct {
	Contract *IBridgeMediatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBridgeMediator creates a new instance of IBridgeMediator, bound to a specific deployed contract.
func NewIBridgeMediator(address common.Address, backend bind.ContractBackend) (*IBridgeMediator, error) {
	contract, err := bindIBridgeMediator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBridgeMediator{IBridgeMediatorCaller: IBridgeMediatorCaller{contract: contract}, IBridgeMediatorTransactor: IBridgeMediatorTransactor{contract: contract}, IBridgeMediatorFilterer: IBridgeMediatorFilterer{contract: contract}}, nil
}

// NewIBridgeMediatorCaller creates a new read-only instance of IBridgeMediator, bound to a specific deployed contract.
func NewIBridgeMediatorCaller(address common.Address, caller bind.ContractCaller) (*IBridgeMediatorCaller, error) {
	contract, err := bindIBridgeMediator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeMediatorCaller{contract: contract}, nil
}

// NewIBridgeMediatorTransactor creates a new write-only instance of IBridgeMediator, bound to a specific deployed contract.
func NewIBridgeMediatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IBridgeMediatorTransactor, error) {
	contract, err := bindIBridgeMediator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBridgeMediatorTransactor{contract: contract}, nil
}

// NewIBridgeMediatorFilterer creates a new log filterer instance of IBridgeMediator, bound to a specific deployed contract.
func NewIBridgeMediatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IBridgeMediatorFilterer, error) {
	contract, err := bindIBridgeMediator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBridgeMediatorFilterer{contract: contract}, nil
}

// bindIBridgeMediator binds a generic wrapper to an already deployed contract.
func bindIBridgeMediator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBridgeMediatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeMediator *IBridgeMediatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeMediator.Contract.IBridgeMediatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeMediator *IBridgeMediatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeMediator.Contract.IBridgeMediatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeMediator *IBridgeMediatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeMediator.Contract.IBridgeMediatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBridgeMediator *IBridgeMediatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBridgeMediator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBridgeMediator *IBridgeMediatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBridgeMediator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBridgeMediator *IBridgeMediatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBridgeMediator.Contract.contract.Transact(opts, method, params...)
}

// Mediate is a free data retrieval call binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 _sourceChain, uint256 _destinationChain, bytes _data) view returns(bytes)
func (_IBridgeMediator *IBridgeMediatorCaller) Mediate(opts *bind.CallOpts, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte) ([]byte, error) {
	var out []interface{}
	err := _IBridgeMediator.contract.Call(opts, &out, "mediate", _sourceChain, _destinationChain, _data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Mediate is a free data retrieval call binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 _sourceChain, uint256 _destinationChain, bytes _data) view returns(bytes)
func (_IBridgeMediator *IBridgeMediatorSession) Mediate(_sourceChain *big.Int, _destinationChain *big.Int, _data []byte) ([]byte, error) {
	return _IBridgeMediator.Contract.Mediate(&_IBridgeMediator.CallOpts, _sourceChain, _destinationChain, _data)
}

// Mediate is a free data retrieval call binding the contract method 0x8b68e0f7.
//
// Solidity: function mediate(uint256 _sourceChain, uint256 _destinationChain, bytes _data) view returns(bytes)
func (_IBridgeMediator *IBridgeMediatorCallerSession) Mediate(_sourceChain *big.Int, _destinationChain *big.Int, _data []byte) ([]byte, error) {
	return _IBridgeMediator.Contract.Mediate(&_IBridgeMediator.CallOpts, _sourceChain, _destinationChain, _data)
}
