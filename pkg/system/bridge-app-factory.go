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

// BridgeAppFactoryMetaData contains all meta data concerning the BridgeAppFactory contract.
var BridgeAppFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"BridgeAppCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"apps\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createApp\",\"outputs\":[{\"internalType\":\"contractBridgeApp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeAppFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeAppFactoryMetaData.ABI instead.
var BridgeAppFactoryABI = BridgeAppFactoryMetaData.ABI

// BridgeAppFactory is an auto generated Go binding around an Ethereum contract.
type BridgeAppFactory struct {
	BridgeAppFactoryCaller     // Read-only binding to the contract
	BridgeAppFactoryTransactor // Write-only binding to the contract
	BridgeAppFactoryFilterer   // Log filterer for contract events
}

// BridgeAppFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeAppFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAppFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeAppFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAppFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeAppFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeAppFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeAppFactorySession struct {
	Contract     *BridgeAppFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeAppFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeAppFactoryCallerSession struct {
	Contract *BridgeAppFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BridgeAppFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeAppFactoryTransactorSession struct {
	Contract     *BridgeAppFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BridgeAppFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeAppFactoryRaw struct {
	Contract *BridgeAppFactory // Generic contract binding to access the raw methods on
}

// BridgeAppFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeAppFactoryCallerRaw struct {
	Contract *BridgeAppFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeAppFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeAppFactoryTransactorRaw struct {
	Contract *BridgeAppFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeAppFactory creates a new instance of BridgeAppFactory, bound to a specific deployed contract.
func NewBridgeAppFactory(address common.Address, backend bind.ContractBackend) (*BridgeAppFactory, error) {
	contract, err := bindBridgeAppFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeAppFactory{BridgeAppFactoryCaller: BridgeAppFactoryCaller{contract: contract}, BridgeAppFactoryTransactor: BridgeAppFactoryTransactor{contract: contract}, BridgeAppFactoryFilterer: BridgeAppFactoryFilterer{contract: contract}}, nil
}

// NewBridgeAppFactoryCaller creates a new read-only instance of BridgeAppFactory, bound to a specific deployed contract.
func NewBridgeAppFactoryCaller(address common.Address, caller bind.ContractCaller) (*BridgeAppFactoryCaller, error) {
	contract, err := bindBridgeAppFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeAppFactoryCaller{contract: contract}, nil
}

// NewBridgeAppFactoryTransactor creates a new write-only instance of BridgeAppFactory, bound to a specific deployed contract.
func NewBridgeAppFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeAppFactoryTransactor, error) {
	contract, err := bindBridgeAppFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeAppFactoryTransactor{contract: contract}, nil
}

// NewBridgeAppFactoryFilterer creates a new log filterer instance of BridgeAppFactory, bound to a specific deployed contract.
func NewBridgeAppFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeAppFactoryFilterer, error) {
	contract, err := bindBridgeAppFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeAppFactoryFilterer{contract: contract}, nil
}

// bindBridgeAppFactory binds a generic wrapper to an already deployed contract.
func bindBridgeAppFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeAppFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeAppFactory *BridgeAppFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeAppFactory.Contract.BridgeAppFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeAppFactory *BridgeAppFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeAppFactory.Contract.BridgeAppFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeAppFactory *BridgeAppFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeAppFactory.Contract.BridgeAppFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeAppFactory *BridgeAppFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeAppFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeAppFactory *BridgeAppFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeAppFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeAppFactory *BridgeAppFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeAppFactory.Contract.contract.Transact(opts, method, params...)
}

// Apps is a free data retrieval call binding the contract method 0x61acc37e.
//
// Solidity: function apps(uint256 ) view returns(address)
func (_BridgeAppFactory *BridgeAppFactoryCaller) Apps(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BridgeAppFactory.contract.Call(opts, &out, "apps", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Apps is a free data retrieval call binding the contract method 0x61acc37e.
//
// Solidity: function apps(uint256 ) view returns(address)
func (_BridgeAppFactory *BridgeAppFactorySession) Apps(arg0 *big.Int) (common.Address, error) {
	return _BridgeAppFactory.Contract.Apps(&_BridgeAppFactory.CallOpts, arg0)
}

// Apps is a free data retrieval call binding the contract method 0x61acc37e.
//
// Solidity: function apps(uint256 ) view returns(address)
func (_BridgeAppFactory *BridgeAppFactoryCallerSession) Apps(arg0 *big.Int) (common.Address, error) {
	return _BridgeAppFactory.Contract.Apps(&_BridgeAppFactory.CallOpts, arg0)
}

// CreateApp is a paid mutator transaction binding the contract method 0xda38b4aa.
//
// Solidity: function createApp() returns(address)
func (_BridgeAppFactory *BridgeAppFactoryTransactor) CreateApp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeAppFactory.contract.Transact(opts, "createApp")
}

// CreateApp is a paid mutator transaction binding the contract method 0xda38b4aa.
//
// Solidity: function createApp() returns(address)
func (_BridgeAppFactory *BridgeAppFactorySession) CreateApp() (*types.Transaction, error) {
	return _BridgeAppFactory.Contract.CreateApp(&_BridgeAppFactory.TransactOpts)
}

// CreateApp is a paid mutator transaction binding the contract method 0xda38b4aa.
//
// Solidity: function createApp() returns(address)
func (_BridgeAppFactory *BridgeAppFactoryTransactorSession) CreateApp() (*types.Transaction, error) {
	return _BridgeAppFactory.Contract.CreateApp(&_BridgeAppFactory.TransactOpts)
}

// BridgeAppFactoryBridgeAppCreatedIterator is returned from FilterBridgeAppCreated and is used to iterate over the raw logs and unpacked data for BridgeAppCreated events raised by the BridgeAppFactory contract.
type BridgeAppFactoryBridgeAppCreatedIterator struct {
	Event *BridgeAppFactoryBridgeAppCreated // Event containing the contract specifics and raw log

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
func (it *BridgeAppFactoryBridgeAppCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeAppFactoryBridgeAppCreated)
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
		it.Event = new(BridgeAppFactoryBridgeAppCreated)
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
func (it *BridgeAppFactoryBridgeAppCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeAppFactoryBridgeAppCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeAppFactoryBridgeAppCreated represents a BridgeAppCreated event raised by the BridgeAppFactory contract.
type BridgeAppFactoryBridgeAppCreated struct {
	ContractAddress common.Address
	Owner           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeAppCreated is a free log retrieval operation binding the contract event 0x620a6b416f656b1eae0f10bc6d9230e71fa7fb1e1a663d7602b8a0e8a625b770.
//
// Solidity: event BridgeAppCreated(address contractAddress, address owner)
func (_BridgeAppFactory *BridgeAppFactoryFilterer) FilterBridgeAppCreated(opts *bind.FilterOpts) (*BridgeAppFactoryBridgeAppCreatedIterator, error) {

	logs, sub, err := _BridgeAppFactory.contract.FilterLogs(opts, "BridgeAppCreated")
	if err != nil {
		return nil, err
	}
	return &BridgeAppFactoryBridgeAppCreatedIterator{contract: _BridgeAppFactory.contract, event: "BridgeAppCreated", logs: logs, sub: sub}, nil
}

// WatchBridgeAppCreated is a free log subscription operation binding the contract event 0x620a6b416f656b1eae0f10bc6d9230e71fa7fb1e1a663d7602b8a0e8a625b770.
//
// Solidity: event BridgeAppCreated(address contractAddress, address owner)
func (_BridgeAppFactory *BridgeAppFactoryFilterer) WatchBridgeAppCreated(opts *bind.WatchOpts, sink chan<- *BridgeAppFactoryBridgeAppCreated) (event.Subscription, error) {

	logs, sub, err := _BridgeAppFactory.contract.WatchLogs(opts, "BridgeAppCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeAppFactoryBridgeAppCreated)
				if err := _BridgeAppFactory.contract.UnpackLog(event, "BridgeAppCreated", log); err != nil {
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

// ParseBridgeAppCreated is a log parse operation binding the contract event 0x620a6b416f656b1eae0f10bc6d9230e71fa7fb1e1a663d7602b8a0e8a625b770.
//
// Solidity: event BridgeAppCreated(address contractAddress, address owner)
func (_BridgeAppFactory *BridgeAppFactoryFilterer) ParseBridgeAppCreated(log types.Log) (*BridgeAppFactoryBridgeAppCreated, error) {
	event := new(BridgeAppFactoryBridgeAppCreated)
	if err := _BridgeAppFactory.contract.UnpackLog(event, "BridgeAppCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
