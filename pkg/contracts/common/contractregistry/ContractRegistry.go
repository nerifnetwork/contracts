// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractregistry

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

// ContractRegistryMetaData contains all meta data concerning the ContractRegistry contract.
var ContractRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"ContractAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610696806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063358177731461005c5780633f0ed0df1461008b5780635c211f88146100a05780638c5b8385146100b3578063c4d66de8146100e7575b600080fd5b61006f61006a3660046104f4565b6100fa565b6040516001600160a01b03909116815260200160405180910390f35b61009e610099366004610549565b6101d6565b005b60005461006f906001600160a01b031681565b61006f6100c13660046104f4565b80516020818301810180516001825292820191909301209152546001600160a01b031681565b61009e6100f536600461059b565b610315565b6000806001600160a01b031660018360405161011691906105e3565b908152604051908190036020019020546001600160a01b0316036101a75760405162461bcd60e51b815260206004820152603c60248201527f436f6e747261637452656769737472793a206e6f20616464726573732077617360448201527f20666f756e6420666f722074686520737065636966696564206b65790000000060648201526084015b60405180910390fd5b6001826040516101b791906105e3565b908152604051908190036020019020546001600160a01b031692915050565b60005460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa15801561021f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024391906105ff565b6001600160a01b0316146102995760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604482015260640161019e565b806001836040516102aa91906105e3565b90815260405190819003602001812080546001600160a01b03939093166001600160a01b0319909316929092179091557fa42de6429c1410f4470a8ff5afeeae27c734519ac1693e8eb58798a81715c94790610309908490849061061c565b60405180910390a15050565b600054600160a81b900460ff161580801561033d57506000546001600160a01b90910460ff16105b8061035e5750303b15801561035e5750600054600160a01b900460ff166001145b6103c15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161019e565b6000805460ff60a01b1916600160a01b17905580156103ee576000805460ff60a81b1916600160a81b1790555b600080546001600160a01b0319166001600160a01b038416179055801561044d576000805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610309565b5050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261047857600080fd5b813567ffffffffffffffff8082111561049357610493610451565b604051601f8301601f19908116603f011681019082821181831017156104bb576104bb610451565b816040528381528660208588010111156104d457600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561050657600080fd5b813567ffffffffffffffff81111561051d57600080fd5b61052984828501610467565b949350505050565b6001600160a01b038116811461054657600080fd5b50565b6000806040838503121561055c57600080fd5b823567ffffffffffffffff81111561057357600080fd5b61057f85828601610467565b925050602083013561059081610531565b809150509250929050565b6000602082840312156105ad57600080fd5b81356105b881610531565b9392505050565b60005b838110156105da5781810151838201526020016105c2565b50506000910152565b600082516105f58184602087016105bf565b9190910192915050565b60006020828403121561061157600080fd5b81516105b881610531565b604081526000835180604084015261063b8160608501602088016105bf565b6001600160a01b0393909316602083015250601f91909101601f19160160600191905056fea2646970667358221220d2b5976bc8cd996507c15e405c4d302223b3865101f6cff57279f62e3773672564736f6c63430008120033",
}

// ContractRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractRegistryMetaData.ABI instead.
var ContractRegistryABI = ContractRegistryMetaData.ABI

// ContractRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractRegistryMetaData.Bin instead.
var ContractRegistryBin = ContractRegistryMetaData.Bin

// DeployContractRegistry deploys a new Ethereum contract, binding an instance of ContractRegistry to it.
func DeployContractRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractRegistry, error) {
	parsed, err := ContractRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractRegistry{ContractRegistryCaller: ContractRegistryCaller{contract: contract}, ContractRegistryTransactor: ContractRegistryTransactor{contract: contract}, ContractRegistryFilterer: ContractRegistryFilterer{contract: contract}}, nil
}

// ContractRegistry is an auto generated Go binding around an Ethereum contract.
type ContractRegistry struct {
	ContractRegistryCaller     // Read-only binding to the contract
	ContractRegistryTransactor // Write-only binding to the contract
	ContractRegistryFilterer   // Log filterer for contract events
}

// ContractRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractRegistrySession struct {
	Contract     *ContractRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractRegistryCallerSession struct {
	Contract *ContractRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractRegistryTransactorSession struct {
	Contract     *ContractRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRegistryRaw struct {
	Contract *ContractRegistry // Generic contract binding to access the raw methods on
}

// ContractRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractRegistryCallerRaw struct {
	Contract *ContractRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ContractRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractRegistryTransactorRaw struct {
	Contract *ContractRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractRegistry creates a new instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistry(address common.Address, backend bind.ContractBackend) (*ContractRegistry, error) {
	contract, err := bindContractRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractRegistry{ContractRegistryCaller: ContractRegistryCaller{contract: contract}, ContractRegistryTransactor: ContractRegistryTransactor{contract: contract}, ContractRegistryFilterer: ContractRegistryFilterer{contract: contract}}, nil
}

// NewContractRegistryCaller creates a new read-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryCaller(address common.Address, caller bind.ContractCaller) (*ContractRegistryCaller, error) {
	contract, err := bindContractRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryCaller{contract: contract}, nil
}

// NewContractRegistryTransactor creates a new write-only instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractRegistryTransactor, error) {
	contract, err := bindContractRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryTransactor{contract: contract}, nil
}

// NewContractRegistryFilterer creates a new log filterer instance of ContractRegistry, bound to a specific deployed contract.
func NewContractRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractRegistryFilterer, error) {
	contract, err := bindContractRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractRegistryFilterer{contract: contract}, nil
}

// bindContractRegistry binds a generic wrapper to an already deployed contract.
func bindContractRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.ContractRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.ContractRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractRegistry *ContractRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractRegistry *ContractRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractRegistry.Contract.contract.Transact(opts, method, params...)
}

// Contracts is a free data retrieval call binding the contract method 0x8c5b8385.
//
// Solidity: function contracts(string ) view returns(address)
func (_ContractRegistry *ContractRegistryCaller) Contracts(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "contracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Contracts is a free data retrieval call binding the contract method 0x8c5b8385.
//
// Solidity: function contracts(string ) view returns(address)
func (_ContractRegistry *ContractRegistrySession) Contracts(arg0 string) (common.Address, error) {
	return _ContractRegistry.Contract.Contracts(&_ContractRegistry.CallOpts, arg0)
}

// Contracts is a free data retrieval call binding the contract method 0x8c5b8385.
//
// Solidity: function contracts(string ) view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) Contracts(arg0 string) (common.Address, error) {
	return _ContractRegistry.Contract.Contracts(&_ContractRegistry.CallOpts, arg0)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _key) view returns(address)
func (_ContractRegistry *ContractRegistryCaller) GetContract(opts *bind.CallOpts, _key string) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "getContract", _key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _key) view returns(address)
func (_ContractRegistry *ContractRegistrySession) GetContract(_key string) (common.Address, error) {
	return _ContractRegistry.Contract.GetContract(&_ContractRegistry.CallOpts, _key)
}

// GetContract is a free data retrieval call binding the contract method 0x35817773.
//
// Solidity: function getContract(string _key) view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) GetContract(_key string) (common.Address, error) {
	return _ContractRegistry.Contract.GetContract(&_ContractRegistry.CallOpts, _key)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_ContractRegistry *ContractRegistryCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractRegistry.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_ContractRegistry *ContractRegistrySession) SignerGetter() (common.Address, error) {
	return _ContractRegistry.Contract.SignerGetter(&_ContractRegistry.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_ContractRegistry *ContractRegistryCallerSession) SignerGetter() (common.Address, error) {
	return _ContractRegistry.Contract.SignerGetter(&_ContractRegistry.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signerGetterAddress) returns()
func (_ContractRegistry *ContractRegistryTransactor) Initialize(opts *bind.TransactOpts, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "initialize", _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signerGetterAddress) returns()
func (_ContractRegistry *ContractRegistrySession) Initialize(_signerGetterAddress common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.Initialize(&_ContractRegistry.TransactOpts, _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _signerGetterAddress) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) Initialize(_signerGetterAddress common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.Initialize(&_ContractRegistry.TransactOpts, _signerGetterAddress)
}

// SetContract is a paid mutator transaction binding the contract method 0x3f0ed0df.
//
// Solidity: function setContract(string _key, address _value) returns()
func (_ContractRegistry *ContractRegistryTransactor) SetContract(opts *bind.TransactOpts, _key string, _value common.Address) (*types.Transaction, error) {
	return _ContractRegistry.contract.Transact(opts, "setContract", _key, _value)
}

// SetContract is a paid mutator transaction binding the contract method 0x3f0ed0df.
//
// Solidity: function setContract(string _key, address _value) returns()
func (_ContractRegistry *ContractRegistrySession) SetContract(_key string, _value common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetContract(&_ContractRegistry.TransactOpts, _key, _value)
}

// SetContract is a paid mutator transaction binding the contract method 0x3f0ed0df.
//
// Solidity: function setContract(string _key, address _value) returns()
func (_ContractRegistry *ContractRegistryTransactorSession) SetContract(_key string, _value common.Address) (*types.Transaction, error) {
	return _ContractRegistry.Contract.SetContract(&_ContractRegistry.TransactOpts, _key, _value)
}

// ContractRegistryContractAddressUpdatedIterator is returned from FilterContractAddressUpdated and is used to iterate over the raw logs and unpacked data for ContractAddressUpdated events raised by the ContractRegistry contract.
type ContractRegistryContractAddressUpdatedIterator struct {
	Event *ContractRegistryContractAddressUpdated // Event containing the contract specifics and raw log

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
func (it *ContractRegistryContractAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryContractAddressUpdated)
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
		it.Event = new(ContractRegistryContractAddressUpdated)
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
func (it *ContractRegistryContractAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryContractAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryContractAddressUpdated represents a ContractAddressUpdated event raised by the ContractRegistry contract.
type ContractRegistryContractAddressUpdated struct {
	Key   string
	Value common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractAddressUpdated is a free log retrieval operation binding the contract event 0xa42de6429c1410f4470a8ff5afeeae27c734519ac1693e8eb58798a81715c947.
//
// Solidity: event ContractAddressUpdated(string key, address value)
func (_ContractRegistry *ContractRegistryFilterer) FilterContractAddressUpdated(opts *bind.FilterOpts) (*ContractRegistryContractAddressUpdatedIterator, error) {

	logs, sub, err := _ContractRegistry.contract.FilterLogs(opts, "ContractAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryContractAddressUpdatedIterator{contract: _ContractRegistry.contract, event: "ContractAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchContractAddressUpdated is a free log subscription operation binding the contract event 0xa42de6429c1410f4470a8ff5afeeae27c734519ac1693e8eb58798a81715c947.
//
// Solidity: event ContractAddressUpdated(string key, address value)
func (_ContractRegistry *ContractRegistryFilterer) WatchContractAddressUpdated(opts *bind.WatchOpts, sink chan<- *ContractRegistryContractAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _ContractRegistry.contract.WatchLogs(opts, "ContractAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryContractAddressUpdated)
				if err := _ContractRegistry.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
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

// ParseContractAddressUpdated is a log parse operation binding the contract event 0xa42de6429c1410f4470a8ff5afeeae27c734519ac1693e8eb58798a81715c947.
//
// Solidity: event ContractAddressUpdated(string key, address value)
func (_ContractRegistry *ContractRegistryFilterer) ParseContractAddressUpdated(log types.Log) (*ContractRegistryContractAddressUpdated, error) {
	event := new(ContractRegistryContractAddressUpdated)
	if err := _ContractRegistry.contract.UnpackLog(event, "ContractAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractRegistry contract.
type ContractRegistryInitializedIterator struct {
	Event *ContractRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ContractRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRegistryInitialized)
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
		it.Event = new(ContractRegistryInitialized)
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
func (it *ContractRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRegistryInitialized represents a Initialized event raised by the ContractRegistry contract.
type ContractRegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRegistry *ContractRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractRegistryInitializedIterator, error) {

	logs, sub, err := _ContractRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractRegistryInitializedIterator{contract: _ContractRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContractRegistry *ContractRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRegistryInitialized)
				if err := _ContractRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractRegistry *ContractRegistryFilterer) ParseInitialized(log types.Log) (*ContractRegistryInitialized, error) {
	event := new(ContractRegistryInitialized)
	if err := _ContractRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
