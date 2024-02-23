// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package signerstorage

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
	_ = abi.ConvertType
)

// SignerStorageMetaData contains all meta data concerning the SignerStorage contract.
var SignerStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"SignerUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSigner\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610478806100206000396000f3fe60806040526004361061003f5760003560e01c80631a296e02146100445780635c211f881461007a578063c4d66de8146100a0578063e30081a0146100c2575b600080fd5b34801561005057600080fd5b506001546001600160a01b03165b6040516001600160a01b03909116815260200160405180910390f35b34801561008657600080fd5b5060005461005e906201000090046001600160a01b031681565b3480156100ac57600080fd5b506100c06100bb366004610401565b6100d5565b005b6100c06100d0366004610401565b610218565b600054610100900460ff16158080156100f55750600054600160ff909116105b8061010f5750303b15801561010f575060005460ff166001145b6101775760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801561019a576000805461ff0019166101001790555b6000805462010000600160b01b031916620100003002179055600180546001600160a01b0319166001600160a01b0384161790558015610214576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015610267573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061028b9190610425565b6001600160a01b0316146102e15760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604482015260640161016e565b600180546001600160a01b0319166001600160a01b038316908117909155604051349160009161520890849084818181858888f193505050503d8060008114610346576040519150601f19603f3d011682016040523d82523d6000602084013e61034b565b606091505b50509050806103a85760405162461bcd60e51b8152602060048201526024808201527f5369676e657253746f726167653a207472616e736665722076616c75652066616044820152631a5b195960e21b606482015260840161016e565b6040516001600160a01b03841681527f5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c739060200160405180910390a1505050565b6001600160a01b03811681146103fe57600080fd5b50565b60006020828403121561041357600080fd5b813561041e816103e9565b9392505050565b60006020828403121561043757600080fd5b815161041e816103e956fea264697066735822122035a3037cd19b460e671aad27c075a4693910f6a269a742813d573a17bfe9c11164736f6c63430008120033",
}

// SignerStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SignerStorageMetaData.ABI instead.
var SignerStorageABI = SignerStorageMetaData.ABI

// SignerStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SignerStorageMetaData.Bin instead.
var SignerStorageBin = SignerStorageMetaData.Bin

// DeploySignerStorage deploys a new Ethereum contract, binding an instance of SignerStorage to it.
func DeploySignerStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SignerStorage, error) {
	parsed, err := SignerStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SignerStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SignerStorage{SignerStorageCaller: SignerStorageCaller{contract: contract}, SignerStorageTransactor: SignerStorageTransactor{contract: contract}, SignerStorageFilterer: SignerStorageFilterer{contract: contract}}, nil
}

// SignerStorage is an auto generated Go binding around an Ethereum contract.
type SignerStorage struct {
	SignerStorageCaller     // Read-only binding to the contract
	SignerStorageTransactor // Write-only binding to the contract
	SignerStorageFilterer   // Log filterer for contract events
}

// SignerStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SignerStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SignerStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignerStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignerStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignerStorageSession struct {
	Contract     *SignerStorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SignerStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignerStorageCallerSession struct {
	Contract *SignerStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SignerStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignerStorageTransactorSession struct {
	Contract     *SignerStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SignerStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SignerStorageRaw struct {
	Contract *SignerStorage // Generic contract binding to access the raw methods on
}

// SignerStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignerStorageCallerRaw struct {
	Contract *SignerStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SignerStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignerStorageTransactorRaw struct {
	Contract *SignerStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSignerStorage creates a new instance of SignerStorage, bound to a specific deployed contract.
func NewSignerStorage(address common.Address, backend bind.ContractBackend) (*SignerStorage, error) {
	contract, err := bindSignerStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignerStorage{SignerStorageCaller: SignerStorageCaller{contract: contract}, SignerStorageTransactor: SignerStorageTransactor{contract: contract}, SignerStorageFilterer: SignerStorageFilterer{contract: contract}}, nil
}

// NewSignerStorageCaller creates a new read-only instance of SignerStorage, bound to a specific deployed contract.
func NewSignerStorageCaller(address common.Address, caller bind.ContractCaller) (*SignerStorageCaller, error) {
	contract, err := bindSignerStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignerStorageCaller{contract: contract}, nil
}

// NewSignerStorageTransactor creates a new write-only instance of SignerStorage, bound to a specific deployed contract.
func NewSignerStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SignerStorageTransactor, error) {
	contract, err := bindSignerStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignerStorageTransactor{contract: contract}, nil
}

// NewSignerStorageFilterer creates a new log filterer instance of SignerStorage, bound to a specific deployed contract.
func NewSignerStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SignerStorageFilterer, error) {
	contract, err := bindSignerStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignerStorageFilterer{contract: contract}, nil
}

// bindSignerStorage binds a generic wrapper to an already deployed contract.
func bindSignerStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignerStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignerStorage *SignerStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignerStorage.Contract.SignerStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignerStorage *SignerStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignerStorage.Contract.SignerStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignerStorage *SignerStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignerStorage.Contract.SignerStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignerStorage *SignerStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignerStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignerStorage *SignerStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignerStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignerStorage *SignerStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignerStorage.Contract.contract.Transact(opts, method, params...)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_SignerStorage *SignerStorageCaller) GetSignerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SignerStorage.contract.Call(opts, &out, "getSignerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_SignerStorage *SignerStorageSession) GetSignerAddress() (common.Address, error) {
	return _SignerStorage.Contract.GetSignerAddress(&_SignerStorage.CallOpts)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_SignerStorage *SignerStorageCallerSession) GetSignerAddress() (common.Address, error) {
	return _SignerStorage.Contract.GetSignerAddress(&_SignerStorage.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SignerStorage *SignerStorageCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SignerStorage.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SignerStorage *SignerStorageSession) SignerGetter() (common.Address, error) {
	return _SignerStorage.Contract.SignerGetter(&_SignerStorage.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SignerStorage *SignerStorageCallerSession) SignerGetter() (common.Address, error) {
	return _SignerStorage.Contract.SignerGetter(&_SignerStorage.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signer) returns()
func (_SignerStorage *SignerStorageTransactor) Initialize(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _SignerStorage.contract.Transact(opts, "initialize", _signer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signer) returns()
func (_SignerStorage *SignerStorageSession) Initialize(_signer common.Address) (*types.Transaction, error) {
	return _SignerStorage.Contract.Initialize(&_SignerStorage.TransactOpts, _signer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signer) returns()
func (_SignerStorage *SignerStorageTransactorSession) Initialize(_signer common.Address) (*types.Transaction, error) {
	return _SignerStorage.Contract.Initialize(&_SignerStorage.TransactOpts, _signer)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _newSigner) payable returns()
func (_SignerStorage *SignerStorageTransactor) SetAddress(opts *bind.TransactOpts, _newSigner common.Address) (*types.Transaction, error) {
	return _SignerStorage.contract.Transact(opts, "setAddress", _newSigner)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _newSigner) payable returns()
func (_SignerStorage *SignerStorageSession) SetAddress(_newSigner common.Address) (*types.Transaction, error) {
	return _SignerStorage.Contract.SetAddress(&_SignerStorage.TransactOpts, _newSigner)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _newSigner) payable returns()
func (_SignerStorage *SignerStorageTransactorSession) SetAddress(_newSigner common.Address) (*types.Transaction, error) {
	return _SignerStorage.Contract.SetAddress(&_SignerStorage.TransactOpts, _newSigner)
}

// SignerStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SignerStorage contract.
type SignerStorageInitializedIterator struct {
	Event *SignerStorageInitialized // Event containing the contract specifics and raw log

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
func (it *SignerStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignerStorageInitialized)
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
		it.Event = new(SignerStorageInitialized)
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
func (it *SignerStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignerStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignerStorageInitialized represents a Initialized event raised by the SignerStorage contract.
type SignerStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SignerStorage *SignerStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*SignerStorageInitializedIterator, error) {

	logs, sub, err := _SignerStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SignerStorageInitializedIterator{contract: _SignerStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SignerStorage *SignerStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SignerStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _SignerStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignerStorageInitialized)
				if err := _SignerStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SignerStorage *SignerStorageFilterer) ParseInitialized(log types.Log) (*SignerStorageInitialized, error) {
	event := new(SignerStorageInitialized)
	if err := _SignerStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SignerStorageSignerUpdatedIterator is returned from FilterSignerUpdated and is used to iterate over the raw logs and unpacked data for SignerUpdated events raised by the SignerStorage contract.
type SignerStorageSignerUpdatedIterator struct {
	Event *SignerStorageSignerUpdated // Event containing the contract specifics and raw log

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
func (it *SignerStorageSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignerStorageSignerUpdated)
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
		it.Event = new(SignerStorageSignerUpdated)
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
func (it *SignerStorageSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignerStorageSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignerStorageSignerUpdated represents a SignerUpdated event raised by the SignerStorage contract.
type SignerStorageSignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerUpdated is a free log retrieval operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_SignerStorage *SignerStorageFilterer) FilterSignerUpdated(opts *bind.FilterOpts) (*SignerStorageSignerUpdatedIterator, error) {

	logs, sub, err := _SignerStorage.contract.FilterLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return &SignerStorageSignerUpdatedIterator{contract: _SignerStorage.contract, event: "SignerUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerUpdated is a free log subscription operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_SignerStorage *SignerStorageFilterer) WatchSignerUpdated(opts *bind.WatchOpts, sink chan<- *SignerStorageSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _SignerStorage.contract.WatchLogs(opts, "SignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignerStorageSignerUpdated)
				if err := _SignerStorage.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
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

// ParseSignerUpdated is a log parse operation binding the contract event 0x5553331329228fbd4123164423717a4a7539f6dfa1c3279a923b98fd681a6c73.
//
// Solidity: event SignerUpdated(address signer)
func (_SignerStorage *SignerStorageFilterer) ParseSignerUpdated(log types.Log) (*SignerStorageSignerUpdated, error) {
	event := new(SignerStorageSignerUpdated)
	if err := _SignerStorage.contract.UnpackLog(event, "SignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
