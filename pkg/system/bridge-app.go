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

// BridgeAppMetaData contains all meta data concerning the BridgeApp contract.
var BridgeAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ContractAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mediatorAddress\",\"type\":\"address\"}],\"name\":\"UpdatedMediatorAddress\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setContractAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mediatorAddress\",\"type\":\"address\"}],\"name\":\"setMediator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeAppABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeAppMetaData.ABI instead.
var BridgeAppABI = BridgeAppMetaData.ABI

// BridgeApp is an auto generated Go binding around an Ethereum contract.
type BridgeApp struct {
	BridgeAppCaller     // Read-only binding to the contract
	BridgeAppTransactor // Write-only binding to the contract
	BridgeAppFilterer   // Log filterer for contract events
}

// BridgeAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeAppSession struct {
	Contract     *BridgeApp        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeAppCallerSession struct {
	Contract *BridgeAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BridgeAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeAppTransactorSession struct {
	Contract     *BridgeAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeAppRaw struct {
	Contract *BridgeApp // Generic contract binding to access the raw methods on
}

// BridgeAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeAppCallerRaw struct {
	Contract *BridgeAppCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeAppTransactorRaw struct {
	Contract *BridgeAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeApp creates a new instance of BridgeApp, bound to a specific deployed contract.
func NewBridgeApp(address common.Address, backend bind.ContractBackend) (*BridgeApp, error) {
	contract, err := bindBridgeApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeApp{BridgeAppCaller: BridgeAppCaller{contract: contract}, BridgeAppTransactor: BridgeAppTransactor{contract: contract}, BridgeAppFilterer: BridgeAppFilterer{contract: contract}}, nil
}

// NewBridgeAppCaller creates a new read-only instance of BridgeApp, bound to a specific deployed contract.
func NewBridgeAppCaller(address common.Address, caller bind.ContractCaller) (*BridgeAppCaller, error) {
	contract, err := bindBridgeApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeAppCaller{contract: contract}, nil
}

// NewBridgeAppTransactor creates a new write-only instance of BridgeApp, bound to a specific deployed contract.
func NewBridgeAppTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeAppTransactor, error) {
	contract, err := bindBridgeApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeAppTransactor{contract: contract}, nil
}

// NewBridgeAppFilterer creates a new log filterer instance of BridgeApp, bound to a specific deployed contract.
func NewBridgeAppFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeAppFilterer, error) {
	contract, err := bindBridgeApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeAppFilterer{contract: contract}, nil
}

// bindBridgeApp binds a generic wrapper to an already deployed contract.
func bindBridgeApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeApp *BridgeAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeApp.Contract.BridgeAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeApp *BridgeAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeApp.Contract.BridgeAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeApp *BridgeAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeApp.Contract.BridgeAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeApp *BridgeAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeApp *BridgeAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeApp *BridgeAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeApp.Contract.contract.Transact(opts, method, params...)
}

// ContractAddresses is a free data retrieval call binding the contract method 0x3263b545.
//
// Solidity: function contractAddresses(uint256 ) view returns(address)
func (_BridgeApp *BridgeAppCaller) ContractAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeApp.contract.Call(opts, &out, "contractAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractAddresses is a free data retrieval call binding the contract method 0x3263b545.
//
// Solidity: function contractAddresses(uint256 ) view returns(address)
func (_BridgeApp *BridgeAppSession) ContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _BridgeApp.Contract.ContractAddresses(&_BridgeApp.CallOpts, arg0)
}

// ContractAddresses is a free data retrieval call binding the contract method 0x3263b545.
//
// Solidity: function contractAddresses(uint256 ) view returns(address)
func (_BridgeApp *BridgeAppCallerSession) ContractAddresses(arg0 *big.Int) (common.Address, error) {
	return _BridgeApp.Contract.ContractAddresses(&_BridgeApp.CallOpts, arg0)
}

// MediatorAddress is a free data retrieval call binding the contract method 0x89952022.
//
// Solidity: function mediatorAddress() view returns(address)
func (_BridgeApp *BridgeAppCaller) MediatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeApp.contract.Call(opts, &out, "mediatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MediatorAddress is a free data retrieval call binding the contract method 0x89952022.
//
// Solidity: function mediatorAddress() view returns(address)
func (_BridgeApp *BridgeAppSession) MediatorAddress() (common.Address, error) {
	return _BridgeApp.Contract.MediatorAddress(&_BridgeApp.CallOpts)
}

// MediatorAddress is a free data retrieval call binding the contract method 0x89952022.
//
// Solidity: function mediatorAddress() view returns(address)
func (_BridgeApp *BridgeAppCallerSession) MediatorAddress() (common.Address, error) {
	return _BridgeApp.Contract.MediatorAddress(&_BridgeApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeApp *BridgeAppCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeApp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeApp *BridgeAppSession) Owner() (common.Address, error) {
	return _BridgeApp.Contract.Owner(&_BridgeApp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeApp *BridgeAppCallerSession) Owner() (common.Address, error) {
	return _BridgeApp.Contract.Owner(&_BridgeApp.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeApp *BridgeAppTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeApp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeApp *BridgeAppSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeApp.Contract.RenounceOwnership(&_BridgeApp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BridgeApp *BridgeAppTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BridgeApp.Contract.RenounceOwnership(&_BridgeApp.TransactOpts)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xbb0165b1.
//
// Solidity: function setContractAddress(uint256 _chainId, address _contractAddress) returns()
func (_BridgeApp *BridgeAppTransactor) SetContractAddress(opts *bind.TransactOpts, _chainId *big.Int, _contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeApp.contract.Transact(opts, "setContractAddress", _chainId, _contractAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xbb0165b1.
//
// Solidity: function setContractAddress(uint256 _chainId, address _contractAddress) returns()
func (_BridgeApp *BridgeAppSession) SetContractAddress(_chainId *big.Int, _contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeApp.Contract.SetContractAddress(&_BridgeApp.TransactOpts, _chainId, _contractAddress)
}

// SetContractAddress is a paid mutator transaction binding the contract method 0xbb0165b1.
//
// Solidity: function setContractAddress(uint256 _chainId, address _contractAddress) returns()
func (_BridgeApp *BridgeAppTransactorSession) SetContractAddress(_chainId *big.Int, _contractAddress common.Address) (*types.Transaction, error) {
	return _BridgeApp.Contract.SetContractAddress(&_BridgeApp.TransactOpts, _chainId, _contractAddress)
}

// SetMediator is a paid mutator transaction binding the contract method 0x49e3ec5e.
//
// Solidity: function setMediator(address _mediatorAddress) returns()
func (_BridgeApp *BridgeAppTransactor) SetMediator(opts *bind.TransactOpts, _mediatorAddress common.Address) (*types.Transaction, error) {
	return _BridgeApp.contract.Transact(opts, "setMediator", _mediatorAddress)
}

// SetMediator is a paid mutator transaction binding the contract method 0x49e3ec5e.
//
// Solidity: function setMediator(address _mediatorAddress) returns()
func (_BridgeApp *BridgeAppSession) SetMediator(_mediatorAddress common.Address) (*types.Transaction, error) {
	return _BridgeApp.Contract.SetMediator(&_BridgeApp.TransactOpts, _mediatorAddress)
}

// SetMediator is a paid mutator transaction binding the contract method 0x49e3ec5e.
//
// Solidity: function setMediator(address _mediatorAddress) returns()
func (_BridgeApp *BridgeAppTransactorSession) SetMediator(_mediatorAddress common.Address) (*types.Transaction, error) {
	return _BridgeApp.Contract.SetMediator(&_BridgeApp.TransactOpts, _mediatorAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeApp *BridgeAppTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BridgeApp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeApp *BridgeAppSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeApp.Contract.TransferOwnership(&_BridgeApp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BridgeApp *BridgeAppTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BridgeApp.Contract.TransferOwnership(&_BridgeApp.TransactOpts, newOwner)
}

// BridgeAppContractAddressUpdatedIterator is returned from FilterContractAddressUpdated and is used to iterate over the raw logs and unpacked data for ContractAddressUpdated events raised by the BridgeApp contract.
type BridgeAppContractAddressUpdatedIterator struct {
	Event *BridgeAppContractAddressUpdated // Event containing the contract specifics and raw log

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
func (it *BridgeAppContractAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAppContractAddressUpdated)
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
		it.Event = new(BridgeAppContractAddressUpdated)
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
func (it *BridgeAppContractAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAppContractAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAppContractAddressUpdated represents a ContractAddressUpdated event raised by the BridgeApp contract.
type BridgeAppContractAddressUpdated struct {
	ChainId         *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterContractAddressUpdated is a free log retrieval operation binding the contract event 0xecd0bf4234a8b3e58f584f3531154aab0c59846fe1c9c39f8f246f16150a254e.
//
// Solidity: event ContractAddressUpdated(uint256 chainId, address contractAddress)
func (_BridgeApp *BridgeAppFilterer) FilterContractAddressUpdated(opts *bind.FilterOpts) (*BridgeAppContractAddressUpdatedIterator, error) {

	logs, sub, err := _BridgeApp.contract.FilterLogs(opts, "ContractAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &BridgeAppContractAddressUpdatedIterator{contract: _BridgeApp.contract, event: "ContractAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchContractAddressUpdated is a free log subscription operation binding the contract event 0xecd0bf4234a8b3e58f584f3531154aab0c59846fe1c9c39f8f246f16150a254e.
//
// Solidity: event ContractAddressUpdated(uint256 chainId, address contractAddress)
func (_BridgeApp *BridgeAppFilterer) WatchContractAddressUpdated(opts *bind.WatchOpts, sink chan<- *BridgeAppContractAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _BridgeApp.contract.WatchLogs(opts, "ContractAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAppContractAddressUpdated)
				if err := _BridgeApp.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
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

// ParseContractAddressUpdated is a log parse operation binding the contract event 0xecd0bf4234a8b3e58f584f3531154aab0c59846fe1c9c39f8f246f16150a254e.
//
// Solidity: event ContractAddressUpdated(uint256 chainId, address contractAddress)
func (_BridgeApp *BridgeAppFilterer) ParseContractAddressUpdated(log types.Log) (*BridgeAppContractAddressUpdated, error) {
	event := new(BridgeAppContractAddressUpdated)
	if err := _BridgeApp.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeAppOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BridgeApp contract.
type BridgeAppOwnershipTransferredIterator struct {
	Event *BridgeAppOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BridgeAppOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAppOwnershipTransferred)
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
		it.Event = new(BridgeAppOwnershipTransferred)
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
func (it *BridgeAppOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAppOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAppOwnershipTransferred represents a OwnershipTransferred event raised by the BridgeApp contract.
type BridgeAppOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeApp *BridgeAppFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BridgeAppOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeApp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeAppOwnershipTransferredIterator{contract: _BridgeApp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BridgeApp *BridgeAppFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BridgeAppOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BridgeApp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAppOwnershipTransferred)
				if err := _BridgeApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BridgeApp *BridgeAppFilterer) ParseOwnershipTransferred(log types.Log) (*BridgeAppOwnershipTransferred, error) {
	event := new(BridgeAppOwnershipTransferred)
	if err := _BridgeApp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeAppUpdatedMediatorAddressIterator is returned from FilterUpdatedMediatorAddress and is used to iterate over the raw logs and unpacked data for UpdatedMediatorAddress events raised by the BridgeApp contract.
type BridgeAppUpdatedMediatorAddressIterator struct {
	Event *BridgeAppUpdatedMediatorAddress // Event containing the contract specifics and raw log

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
func (it *BridgeAppUpdatedMediatorAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAppUpdatedMediatorAddress)
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
		it.Event = new(BridgeAppUpdatedMediatorAddress)
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
func (it *BridgeAppUpdatedMediatorAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAppUpdatedMediatorAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAppUpdatedMediatorAddress represents a UpdatedMediatorAddress event raised by the BridgeApp contract.
type BridgeAppUpdatedMediatorAddress struct {
	MediatorAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMediatorAddress is a free log retrieval operation binding the contract event 0xfffe8850461024d4daf3f0dcafc0a3194b70441f38d49592a768132fc11c21c4.
//
// Solidity: event UpdatedMediatorAddress(address mediatorAddress)
func (_BridgeApp *BridgeAppFilterer) FilterUpdatedMediatorAddress(opts *bind.FilterOpts) (*BridgeAppUpdatedMediatorAddressIterator, error) {

	logs, sub, err := _BridgeApp.contract.FilterLogs(opts, "UpdatedMediatorAddress")
	if err != nil {
		return nil, err
	}
	return &BridgeAppUpdatedMediatorAddressIterator{contract: _BridgeApp.contract, event: "UpdatedMediatorAddress", logs: logs, sub: sub}, nil
}

// WatchUpdatedMediatorAddress is a free log subscription operation binding the contract event 0xfffe8850461024d4daf3f0dcafc0a3194b70441f38d49592a768132fc11c21c4.
//
// Solidity: event UpdatedMediatorAddress(address mediatorAddress)
func (_BridgeApp *BridgeAppFilterer) WatchUpdatedMediatorAddress(opts *bind.WatchOpts, sink chan<- *BridgeAppUpdatedMediatorAddress) (event.Subscription, error) {

	logs, sub, err := _BridgeApp.contract.WatchLogs(opts, "UpdatedMediatorAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAppUpdatedMediatorAddress)
				if err := _BridgeApp.contract.UnpackLog(event, "UpdatedMediatorAddress", log); err != nil {
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

// ParseUpdatedMediatorAddress is a log parse operation binding the contract event 0xfffe8850461024d4daf3f0dcafc0a3194b70441f38d49592a768132fc11c21c4.
//
// Solidity: event UpdatedMediatorAddress(address mediatorAddress)
func (_BridgeApp *BridgeAppFilterer) ParseUpdatedMediatorAddress(log types.Log) (*BridgeAppUpdatedMediatorAddress, error) {
	event := new(BridgeAppUpdatedMediatorAddress)
	if err := _BridgeApp.contract.UnpackLog(event, "UpdatedMediatorAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
