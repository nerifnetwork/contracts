// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// RelayBridgeMetaData contains all meta data concerning the RelayBridge contract.
var RelayBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"appContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorFee\",\"type\":\"uint256\"}],\"name\":\"FailedSend\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"}],\"name\":\"Reverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"appContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destinationChain\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"validatorFee\",\"type\":\"uint256\"}],\"name\":\"Sent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeValidatorFeePool\",\"outputs\":[{\"internalType\":\"contractIBridgeValidatorFeePool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_appContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"dataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_appContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_leader\",\"type\":\"address\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"failed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_appContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sourceChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_validatorFee\",\"type\":\"uint256\"}],\"name\":\"failedSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerStorage\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_bridgeValidatorFeePool\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leaderHistory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaderHistoryLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_appContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_leader\",\"type\":\"address\"}],\"name\":\"revertSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"reverted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_destinationChain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sentData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerStorage\",\"outputs\":[{\"internalType\":\"contractSignerStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RelayBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayBridgeMetaData.ABI instead.
var RelayBridgeABI = RelayBridgeMetaData.ABI

// RelayBridge is an auto generated Go binding around an Ethereum contract.
type RelayBridge struct {
	RelayBridgeCaller     // Read-only binding to the contract
	RelayBridgeTransactor // Write-only binding to the contract
	RelayBridgeFilterer   // Log filterer for contract events
}

// RelayBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayBridgeSession struct {
	Contract     *RelayBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayBridgeCallerSession struct {
	Contract *RelayBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RelayBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayBridgeTransactorSession struct {
	Contract     *RelayBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RelayBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayBridgeRaw struct {
	Contract *RelayBridge // Generic contract binding to access the raw methods on
}

// RelayBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayBridgeCallerRaw struct {
	Contract *RelayBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// RelayBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayBridgeTransactorRaw struct {
	Contract *RelayBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayBridge creates a new instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridge(address common.Address, backend bind.ContractBackend) (*RelayBridge, error) {
	contract, err := bindRelayBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayBridge{RelayBridgeCaller: RelayBridgeCaller{contract: contract}, RelayBridgeTransactor: RelayBridgeTransactor{contract: contract}, RelayBridgeFilterer: RelayBridgeFilterer{contract: contract}}, nil
}

// NewRelayBridgeCaller creates a new read-only instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridgeCaller(address common.Address, caller bind.ContractCaller) (*RelayBridgeCaller, error) {
	contract, err := bindRelayBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayBridgeCaller{contract: contract}, nil
}

// NewRelayBridgeTransactor creates a new write-only instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayBridgeTransactor, error) {
	contract, err := bindRelayBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayBridgeTransactor{contract: contract}, nil
}

// NewRelayBridgeFilterer creates a new log filterer instance of RelayBridge, bound to a specific deployed contract.
func NewRelayBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayBridgeFilterer, error) {
	contract, err := bindRelayBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayBridgeFilterer{contract: contract}, nil
}

// bindRelayBridge binds a generic wrapper to an already deployed contract.
func bindRelayBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayBridge *RelayBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayBridge.Contract.RelayBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayBridge *RelayBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayBridge.Contract.RelayBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayBridge *RelayBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayBridge.Contract.RelayBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayBridge *RelayBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayBridge *RelayBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayBridge *RelayBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeValidatorFeePool is a free data retrieval call binding the contract method 0x6f4b418c.
//
// Solidity: function bridgeValidatorFeePool() view returns(address)
func (_RelayBridge *RelayBridgeCaller) BridgeValidatorFeePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "bridgeValidatorFeePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeValidatorFeePool is a free data retrieval call binding the contract method 0x6f4b418c.
//
// Solidity: function bridgeValidatorFeePool() view returns(address)
func (_RelayBridge *RelayBridgeSession) BridgeValidatorFeePool() (common.Address, error) {
	return _RelayBridge.Contract.BridgeValidatorFeePool(&_RelayBridge.CallOpts)
}

// BridgeValidatorFeePool is a free data retrieval call binding the contract method 0x6f4b418c.
//
// Solidity: function bridgeValidatorFeePool() view returns(address)
func (_RelayBridge *RelayBridgeCallerSession) BridgeValidatorFeePool() (common.Address, error) {
	return _RelayBridge.Contract.BridgeValidatorFeePool(&_RelayBridge.CallOpts)
}

// DataHash is a free data retrieval call binding the contract method 0xc377a06b.
//
// Solidity: function dataHash(address _appContract, uint256 _sourceChain, uint256 _destinationChain, uint256 _gasLimit, bytes _data, uint256 _nonce) pure returns(bytes32)
func (_RelayBridge *RelayBridgeCaller) DataHash(opts *bind.CallOpts, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "dataHash", _appContract, _sourceChain, _destinationChain, _gasLimit, _data, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataHash is a free data retrieval call binding the contract method 0xc377a06b.
//
// Solidity: function dataHash(address _appContract, uint256 _sourceChain, uint256 _destinationChain, uint256 _gasLimit, bytes _data, uint256 _nonce) pure returns(bytes32)
func (_RelayBridge *RelayBridgeSession) DataHash(_appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int) ([32]byte, error) {
	return _RelayBridge.Contract.DataHash(&_RelayBridge.CallOpts, _appContract, _sourceChain, _destinationChain, _gasLimit, _data, _nonce)
}

// DataHash is a free data retrieval call binding the contract method 0xc377a06b.
//
// Solidity: function dataHash(address _appContract, uint256 _sourceChain, uint256 _destinationChain, uint256 _gasLimit, bytes _data, uint256 _nonce) pure returns(bytes32)
func (_RelayBridge *RelayBridgeCallerSession) DataHash(_appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int) ([32]byte, error) {
	return _RelayBridge.Contract.DataHash(&_RelayBridge.CallOpts, _appContract, _sourceChain, _destinationChain, _gasLimit, _data, _nonce)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCaller) Executed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "executed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeSession) Executed(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Executed(&_RelayBridge.CallOpts, arg0)
}

// Executed is a free data retrieval call binding the contract method 0xa9fcfb33.
//
// Solidity: function executed(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCallerSession) Executed(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Executed(&_RelayBridge.CallOpts, arg0)
}

// Failed is a free data retrieval call binding the contract method 0xf183d06b.
//
// Solidity: function failed(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCaller) Failed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "failed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Failed is a free data retrieval call binding the contract method 0xf183d06b.
//
// Solidity: function failed(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeSession) Failed(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Failed(&_RelayBridge.CallOpts, arg0)
}

// Failed is a free data retrieval call binding the contract method 0xf183d06b.
//
// Solidity: function failed(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCallerSession) Failed(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Failed(&_RelayBridge.CallOpts, arg0)
}

// LeaderHistory is a free data retrieval call binding the contract method 0xe10e62ca.
//
// Solidity: function leaderHistory(uint256 ) view returns(address)
func (_RelayBridge *RelayBridgeCaller) LeaderHistory(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "leaderHistory", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LeaderHistory is a free data retrieval call binding the contract method 0xe10e62ca.
//
// Solidity: function leaderHistory(uint256 ) view returns(address)
func (_RelayBridge *RelayBridgeSession) LeaderHistory(arg0 *big.Int) (common.Address, error) {
	return _RelayBridge.Contract.LeaderHistory(&_RelayBridge.CallOpts, arg0)
}

// LeaderHistory is a free data retrieval call binding the contract method 0xe10e62ca.
//
// Solidity: function leaderHistory(uint256 ) view returns(address)
func (_RelayBridge *RelayBridgeCallerSession) LeaderHistory(arg0 *big.Int) (common.Address, error) {
	return _RelayBridge.Contract.LeaderHistory(&_RelayBridge.CallOpts, arg0)
}

// LeaderHistoryLength is a free data retrieval call binding the contract method 0x11393a65.
//
// Solidity: function leaderHistoryLength() view returns(uint256)
func (_RelayBridge *RelayBridgeCaller) LeaderHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "leaderHistoryLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeaderHistoryLength is a free data retrieval call binding the contract method 0x11393a65.
//
// Solidity: function leaderHistoryLength() view returns(uint256)
func (_RelayBridge *RelayBridgeSession) LeaderHistoryLength() (*big.Int, error) {
	return _RelayBridge.Contract.LeaderHistoryLength(&_RelayBridge.CallOpts)
}

// LeaderHistoryLength is a free data retrieval call binding the contract method 0x11393a65.
//
// Solidity: function leaderHistoryLength() view returns(uint256)
func (_RelayBridge *RelayBridgeCallerSession) LeaderHistoryLength() (*big.Int, error) {
	return _RelayBridge.Contract.LeaderHistoryLength(&_RelayBridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_RelayBridge *RelayBridgeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_RelayBridge *RelayBridgeSession) Nonce() (*big.Int, error) {
	return _RelayBridge.Contract.Nonce(&_RelayBridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_RelayBridge *RelayBridgeCallerSession) Nonce() (*big.Int, error) {
	return _RelayBridge.Contract.Nonce(&_RelayBridge.CallOpts)
}

// Reverted is a free data retrieval call binding the contract method 0xabe61ec4.
//
// Solidity: function reverted(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCaller) Reverted(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "reverted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Reverted is a free data retrieval call binding the contract method 0xabe61ec4.
//
// Solidity: function reverted(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeSession) Reverted(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Reverted(&_RelayBridge.CallOpts, arg0)
}

// Reverted is a free data retrieval call binding the contract method 0xabe61ec4.
//
// Solidity: function reverted(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCallerSession) Reverted(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Reverted(&_RelayBridge.CallOpts, arg0)
}

// Sent is a free data retrieval call binding the contract method 0x8ebc5074.
//
// Solidity: function sent(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCaller) Sent(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "sent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Sent is a free data retrieval call binding the contract method 0x8ebc5074.
//
// Solidity: function sent(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeSession) Sent(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Sent(&_RelayBridge.CallOpts, arg0)
}

// Sent is a free data retrieval call binding the contract method 0x8ebc5074.
//
// Solidity: function sent(bytes32 ) view returns(bool)
func (_RelayBridge *RelayBridgeCallerSession) Sent(arg0 [32]byte) (bool, error) {
	return _RelayBridge.Contract.Sent(&_RelayBridge.CallOpts, arg0)
}

// SentData is a free data retrieval call binding the contract method 0xf0b6c9db.
//
// Solidity: function sentData(bytes32 ) view returns(bytes)
func (_RelayBridge *RelayBridgeCaller) SentData(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "sentData", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// SentData is a free data retrieval call binding the contract method 0xf0b6c9db.
//
// Solidity: function sentData(bytes32 ) view returns(bytes)
func (_RelayBridge *RelayBridgeSession) SentData(arg0 [32]byte) ([]byte, error) {
	return _RelayBridge.Contract.SentData(&_RelayBridge.CallOpts, arg0)
}

// SentData is a free data retrieval call binding the contract method 0xf0b6c9db.
//
// Solidity: function sentData(bytes32 ) view returns(bytes)
func (_RelayBridge *RelayBridgeCallerSession) SentData(arg0 [32]byte) ([]byte, error) {
	return _RelayBridge.Contract.SentData(&_RelayBridge.CallOpts, arg0)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_RelayBridge *RelayBridgeCaller) SignerStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RelayBridge.contract.Call(opts, &out, "signerStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_RelayBridge *RelayBridgeSession) SignerStorage() (common.Address, error) {
	return _RelayBridge.Contract.SignerStorage(&_RelayBridge.CallOpts)
}

// SignerStorage is a free data retrieval call binding the contract method 0x449a23c7.
//
// Solidity: function signerStorage() view returns(address)
func (_RelayBridge *RelayBridgeCallerSession) SignerStorage() (common.Address, error) {
	return _RelayBridge.Contract.SignerStorage(&_RelayBridge.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x1842cd77.
//
// Solidity: function execute(address _appContract, uint256 _sourceChain, uint256 _gasLimit, bytes _data, uint256 _nonce, address _leader) returns()
func (_RelayBridge *RelayBridgeTransactor) Execute(opts *bind.TransactOpts, _appContract common.Address, _sourceChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int, _leader common.Address) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "execute", _appContract, _sourceChain, _gasLimit, _data, _nonce, _leader)
}

// Execute is a paid mutator transaction binding the contract method 0x1842cd77.
//
// Solidity: function execute(address _appContract, uint256 _sourceChain, uint256 _gasLimit, bytes _data, uint256 _nonce, address _leader) returns()
func (_RelayBridge *RelayBridgeSession) Execute(_appContract common.Address, _sourceChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int, _leader common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.Execute(&_RelayBridge.TransactOpts, _appContract, _sourceChain, _gasLimit, _data, _nonce, _leader)
}

// Execute is a paid mutator transaction binding the contract method 0x1842cd77.
//
// Solidity: function execute(address _appContract, uint256 _sourceChain, uint256 _gasLimit, bytes _data, uint256 _nonce, address _leader) returns()
func (_RelayBridge *RelayBridgeTransactorSession) Execute(_appContract common.Address, _sourceChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int, _leader common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.Execute(&_RelayBridge.TransactOpts, _appContract, _sourceChain, _gasLimit, _data, _nonce, _leader)
}

// FailedSend is a paid mutator transaction binding the contract method 0x30df45d4.
//
// Solidity: function failedSend(address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _gasLimit, uint256 _nonce, uint256 _validatorFee) returns()
func (_RelayBridge *RelayBridgeTransactor) FailedSend(opts *bind.TransactOpts, _appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _gasLimit *big.Int, _nonce *big.Int, _validatorFee *big.Int) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "failedSend", _appContract, _sourceChain, _destinationChain, _data, _gasLimit, _nonce, _validatorFee)
}

// FailedSend is a paid mutator transaction binding the contract method 0x30df45d4.
//
// Solidity: function failedSend(address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _gasLimit, uint256 _nonce, uint256 _validatorFee) returns()
func (_RelayBridge *RelayBridgeSession) FailedSend(_appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _gasLimit *big.Int, _nonce *big.Int, _validatorFee *big.Int) (*types.Transaction, error) {
	return _RelayBridge.Contract.FailedSend(&_RelayBridge.TransactOpts, _appContract, _sourceChain, _destinationChain, _data, _gasLimit, _nonce, _validatorFee)
}

// FailedSend is a paid mutator transaction binding the contract method 0x30df45d4.
//
// Solidity: function failedSend(address _appContract, uint256 _sourceChain, uint256 _destinationChain, bytes _data, uint256 _gasLimit, uint256 _nonce, uint256 _validatorFee) returns()
func (_RelayBridge *RelayBridgeTransactorSession) FailedSend(_appContract common.Address, _sourceChain *big.Int, _destinationChain *big.Int, _data []byte, _gasLimit *big.Int, _nonce *big.Int, _validatorFee *big.Int) (*types.Transaction, error) {
	return _RelayBridge.Contract.FailedSend(&_RelayBridge.TransactOpts, _appContract, _sourceChain, _destinationChain, _data, _gasLimit, _nonce, _validatorFee)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _signerStorage, address _bridgeValidatorFeePool) returns()
func (_RelayBridge *RelayBridgeTransactor) Initialize(opts *bind.TransactOpts, _signerStorage common.Address, _bridgeValidatorFeePool common.Address) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "initialize", _signerStorage, _bridgeValidatorFeePool)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _signerStorage, address _bridgeValidatorFeePool) returns()
func (_RelayBridge *RelayBridgeSession) Initialize(_signerStorage common.Address, _bridgeValidatorFeePool common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.Initialize(&_RelayBridge.TransactOpts, _signerStorage, _bridgeValidatorFeePool)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _signerStorage, address _bridgeValidatorFeePool) returns()
func (_RelayBridge *RelayBridgeTransactorSession) Initialize(_signerStorage common.Address, _bridgeValidatorFeePool common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.Initialize(&_RelayBridge.TransactOpts, _signerStorage, _bridgeValidatorFeePool)
}

// RevertSend is a paid mutator transaction binding the contract method 0x47176935.
//
// Solidity: function revertSend(address _appContract, uint256 _destinationChain, uint256 _gasLimit, bytes _data, uint256 _nonce, address _leader) returns()
func (_RelayBridge *RelayBridgeTransactor) RevertSend(opts *bind.TransactOpts, _appContract common.Address, _destinationChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int, _leader common.Address) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "revertSend", _appContract, _destinationChain, _gasLimit, _data, _nonce, _leader)
}

// RevertSend is a paid mutator transaction binding the contract method 0x47176935.
//
// Solidity: function revertSend(address _appContract, uint256 _destinationChain, uint256 _gasLimit, bytes _data, uint256 _nonce, address _leader) returns()
func (_RelayBridge *RelayBridgeSession) RevertSend(_appContract common.Address, _destinationChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int, _leader common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.RevertSend(&_RelayBridge.TransactOpts, _appContract, _destinationChain, _gasLimit, _data, _nonce, _leader)
}

// RevertSend is a paid mutator transaction binding the contract method 0x47176935.
//
// Solidity: function revertSend(address _appContract, uint256 _destinationChain, uint256 _gasLimit, bytes _data, uint256 _nonce, address _leader) returns()
func (_RelayBridge *RelayBridgeTransactorSession) RevertSend(_appContract common.Address, _destinationChain *big.Int, _gasLimit *big.Int, _data []byte, _nonce *big.Int, _leader common.Address) (*types.Transaction, error) {
	return _RelayBridge.Contract.RevertSend(&_RelayBridge.TransactOpts, _appContract, _destinationChain, _gasLimit, _data, _nonce, _leader)
}

// Send is a paid mutator transaction binding the contract method 0x2a9bf608.
//
// Solidity: function send(uint256 _destinationChain, uint256 _gasLimit, bytes _data) payable returns()
func (_RelayBridge *RelayBridgeTransactor) Send(opts *bind.TransactOpts, _destinationChain *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _RelayBridge.contract.Transact(opts, "send", _destinationChain, _gasLimit, _data)
}

// Send is a paid mutator transaction binding the contract method 0x2a9bf608.
//
// Solidity: function send(uint256 _destinationChain, uint256 _gasLimit, bytes _data) payable returns()
func (_RelayBridge *RelayBridgeSession) Send(_destinationChain *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _RelayBridge.Contract.Send(&_RelayBridge.TransactOpts, _destinationChain, _gasLimit, _data)
}

// Send is a paid mutator transaction binding the contract method 0x2a9bf608.
//
// Solidity: function send(uint256 _destinationChain, uint256 _gasLimit, bytes _data) payable returns()
func (_RelayBridge *RelayBridgeTransactorSession) Send(_destinationChain *big.Int, _gasLimit *big.Int, _data []byte) (*types.Transaction, error) {
	return _RelayBridge.Contract.Send(&_RelayBridge.TransactOpts, _destinationChain, _gasLimit, _data)
}

// RelayBridgeExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the RelayBridge contract.
type RelayBridgeExecutedIterator struct {
	Event *RelayBridgeExecuted // Event containing the contract specifics and raw log

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
func (it *RelayBridgeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeExecuted)
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
		it.Event = new(RelayBridgeExecuted)
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
func (it *RelayBridgeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeExecuted represents a Executed event raised by the RelayBridge contract.
type RelayBridgeExecuted struct {
	Hash             [32]byte
	SourceChain      *big.Int
	DestinationChain *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xfe95c3832c12555e6bdb94254f0062f3fef9fb21554e39a02734ac465105930b.
//
// Solidity: event Executed(bytes32 hash, uint256 sourceChain, uint256 destinationChain)
func (_RelayBridge *RelayBridgeFilterer) FilterExecuted(opts *bind.FilterOpts) (*RelayBridgeExecutedIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeExecutedIterator{contract: _RelayBridge.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xfe95c3832c12555e6bdb94254f0062f3fef9fb21554e39a02734ac465105930b.
//
// Solidity: event Executed(bytes32 hash, uint256 sourceChain, uint256 destinationChain)
func (_RelayBridge *RelayBridgeFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *RelayBridgeExecuted) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "Executed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeExecuted)
				if err := _RelayBridge.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xfe95c3832c12555e6bdb94254f0062f3fef9fb21554e39a02734ac465105930b.
//
// Solidity: event Executed(bytes32 hash, uint256 sourceChain, uint256 destinationChain)
func (_RelayBridge *RelayBridgeFilterer) ParseExecuted(log types.Log) (*RelayBridgeExecuted, error) {
	event := new(RelayBridgeExecuted)
	if err := _RelayBridge.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayBridgeFailedSendIterator is returned from FilterFailedSend and is used to iterate over the raw logs and unpacked data for FailedSend events raised by the RelayBridge contract.
type RelayBridgeFailedSendIterator struct {
	Event *RelayBridgeFailedSend // Event containing the contract specifics and raw log

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
func (it *RelayBridgeFailedSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeFailedSend)
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
		it.Event = new(RelayBridgeFailedSend)
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
func (it *RelayBridgeFailedSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeFailedSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeFailedSend represents a FailedSend event raised by the RelayBridge contract.
type RelayBridgeFailedSend struct {
	Hash             [32]byte
	AppContract      common.Address
	SourceChain      *big.Int
	DestinationChain *big.Int
	Data             []byte
	GasLimit         *big.Int
	Nonce            *big.Int
	ValidatorFee     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFailedSend is a free log retrieval operation binding the contract event 0x34f4c68c482ca8fdc54ad8c1661ace4051b24dbe309fb9e38ea9f688b3e7c14d.
//
// Solidity: event FailedSend(bytes32 hash, address appContract, uint256 sourceChain, uint256 destinationChain, bytes data, uint256 gasLimit, uint256 nonce, uint256 validatorFee)
func (_RelayBridge *RelayBridgeFilterer) FilterFailedSend(opts *bind.FilterOpts) (*RelayBridgeFailedSendIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "FailedSend")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeFailedSendIterator{contract: _RelayBridge.contract, event: "FailedSend", logs: logs, sub: sub}, nil
}

// WatchFailedSend is a free log subscription operation binding the contract event 0x34f4c68c482ca8fdc54ad8c1661ace4051b24dbe309fb9e38ea9f688b3e7c14d.
//
// Solidity: event FailedSend(bytes32 hash, address appContract, uint256 sourceChain, uint256 destinationChain, bytes data, uint256 gasLimit, uint256 nonce, uint256 validatorFee)
func (_RelayBridge *RelayBridgeFilterer) WatchFailedSend(opts *bind.WatchOpts, sink chan<- *RelayBridgeFailedSend) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "FailedSend")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeFailedSend)
				if err := _RelayBridge.contract.UnpackLog(event, "FailedSend", log); err != nil {
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

// ParseFailedSend is a log parse operation binding the contract event 0x34f4c68c482ca8fdc54ad8c1661ace4051b24dbe309fb9e38ea9f688b3e7c14d.
//
// Solidity: event FailedSend(bytes32 hash, address appContract, uint256 sourceChain, uint256 destinationChain, bytes data, uint256 gasLimit, uint256 nonce, uint256 validatorFee)
func (_RelayBridge *RelayBridgeFilterer) ParseFailedSend(log types.Log) (*RelayBridgeFailedSend, error) {
	event := new(RelayBridgeFailedSend)
	if err := _RelayBridge.contract.UnpackLog(event, "FailedSend", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RelayBridge contract.
type RelayBridgeInitializedIterator struct {
	Event *RelayBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *RelayBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeInitialized)
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
		it.Event = new(RelayBridgeInitialized)
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
func (it *RelayBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeInitialized represents a Initialized event raised by the RelayBridge contract.
type RelayBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayBridge *RelayBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*RelayBridgeInitializedIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeInitializedIterator{contract: _RelayBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RelayBridge *RelayBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RelayBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeInitialized)
				if err := _RelayBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RelayBridge *RelayBridgeFilterer) ParseInitialized(log types.Log) (*RelayBridgeInitialized, error) {
	event := new(RelayBridgeInitialized)
	if err := _RelayBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayBridgeRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the RelayBridge contract.
type RelayBridgeRevertedIterator struct {
	Event *RelayBridgeReverted // Event containing the contract specifics and raw log

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
func (it *RelayBridgeRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeReverted)
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
		it.Event = new(RelayBridgeReverted)
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
func (it *RelayBridgeRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeReverted represents a Reverted event raised by the RelayBridge contract.
type RelayBridgeReverted struct {
	Hash             [32]byte
	SourceChain      *big.Int
	DestinationChain *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xe7cb59aae52f5e18b3e1c1b25deb68a9d250e28c102ecf456c46c058ea7a0425.
//
// Solidity: event Reverted(bytes32 hash, uint256 sourceChain, uint256 destinationChain)
func (_RelayBridge *RelayBridgeFilterer) FilterReverted(opts *bind.FilterOpts) (*RelayBridgeRevertedIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "Reverted")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeRevertedIterator{contract: _RelayBridge.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xe7cb59aae52f5e18b3e1c1b25deb68a9d250e28c102ecf456c46c058ea7a0425.
//
// Solidity: event Reverted(bytes32 hash, uint256 sourceChain, uint256 destinationChain)
func (_RelayBridge *RelayBridgeFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *RelayBridgeReverted) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "Reverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeReverted)
				if err := _RelayBridge.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0xe7cb59aae52f5e18b3e1c1b25deb68a9d250e28c102ecf456c46c058ea7a0425.
//
// Solidity: event Reverted(bytes32 hash, uint256 sourceChain, uint256 destinationChain)
func (_RelayBridge *RelayBridgeFilterer) ParseReverted(log types.Log) (*RelayBridgeReverted, error) {
	event := new(RelayBridgeReverted)
	if err := _RelayBridge.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayBridgeSentIterator is returned from FilterSent and is used to iterate over the raw logs and unpacked data for Sent events raised by the RelayBridge contract.
type RelayBridgeSentIterator struct {
	Event *RelayBridgeSent // Event containing the contract specifics and raw log

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
func (it *RelayBridgeSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayBridgeSent)
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
		it.Event = new(RelayBridgeSent)
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
func (it *RelayBridgeSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayBridgeSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayBridgeSent represents a Sent event raised by the RelayBridge contract.
type RelayBridgeSent struct {
	Hash             [32]byte
	AppContract      common.Address
	SourceChain      *big.Int
	DestinationChain *big.Int
	Data             []byte
	GasLimit         *big.Int
	Nonce            *big.Int
	ValidatorFee     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSent is a free log retrieval operation binding the contract event 0x00682a4fe9b25611837eed848f00ec3ef176dc6f61ae5d49ace0f5385340bd0c.
//
// Solidity: event Sent(bytes32 hash, address appContract, uint256 sourceChain, uint256 destinationChain, bytes data, uint256 gasLimit, uint256 nonce, uint256 validatorFee)
func (_RelayBridge *RelayBridgeFilterer) FilterSent(opts *bind.FilterOpts) (*RelayBridgeSentIterator, error) {

	logs, sub, err := _RelayBridge.contract.FilterLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return &RelayBridgeSentIterator{contract: _RelayBridge.contract, event: "Sent", logs: logs, sub: sub}, nil
}

// WatchSent is a free log subscription operation binding the contract event 0x00682a4fe9b25611837eed848f00ec3ef176dc6f61ae5d49ace0f5385340bd0c.
//
// Solidity: event Sent(bytes32 hash, address appContract, uint256 sourceChain, uint256 destinationChain, bytes data, uint256 gasLimit, uint256 nonce, uint256 validatorFee)
func (_RelayBridge *RelayBridgeFilterer) WatchSent(opts *bind.WatchOpts, sink chan<- *RelayBridgeSent) (event.Subscription, error) {

	logs, sub, err := _RelayBridge.contract.WatchLogs(opts, "Sent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayBridgeSent)
				if err := _RelayBridge.contract.UnpackLog(event, "Sent", log); err != nil {
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

// ParseSent is a log parse operation binding the contract event 0x00682a4fe9b25611837eed848f00ec3ef176dc6f61ae5d49ace0f5385340bd0c.
//
// Solidity: event Sent(bytes32 hash, address appContract, uint256 sourceChain, uint256 destinationChain, bytes data, uint256 gasLimit, uint256 nonce, uint256 validatorFee)
func (_RelayBridge *RelayBridgeFilterer) ParseSent(log types.Log) (*RelayBridgeSent, error) {
	event := new(RelayBridgeSent)
	if err := _RelayBridge.contract.UnpackLog(event, "Sent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
