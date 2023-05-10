// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package common

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

// ContractRegistryMetaData contains all meta data concerning the ContractRegistry contract.
var ContractRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"value\",\"type\":\"address\"}],\"name\":\"ContractAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"contracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"getContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_value\",\"type\":\"address\"}],\"name\":\"setContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c93806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063358177731461005c5780633f0ed0df1461008c5780635c211f88146100a85780638c5b8385146100c6578063c4d66de8146100f6575b600080fd5b61007660048036038101906100719190610752565b610112565b60405161008391906107dc565b60405180910390f35b6100a660048036038101906100a19190610823565b610206565b005b6100b061039c565b6040516100bd91906108de565b60405180910390f35b6100e060048036038101906100db9190610752565b6103c0565b6040516100ed91906107dc565b60405180910390f35b610110600480360381019061010b91906108f9565b610409565b005b60008073ffffffffffffffffffffffffffffffffffffffff1660018360405161013b9190610997565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036101c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b790610a31565b60405180910390fd5b6001826040516101d09190610997565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a296e026040518163ffffffff1660e01b8152600401602060405180830381865afa158015610288573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ac9190610a66565b73ffffffffffffffffffffffffffffffffffffffff1614610302576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f990610adf565b60405180910390fd5b806001836040516103139190610997565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fa42de6429c1410f4470a8ff5afeeae27c734519ac1693e8eb58798a81715c9478282604051610390929190610b38565b60405180910390a15050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610415600161049f565b90508015610439576001600060156101000a81548160ff0219169083151502179055505b61044282610592565b801561049b5760008060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516104929190610bb0565b60405180910390a15b5050565b60008060159054906101000a900460ff16156105165760018260ff161480156104ce57506104cc306105d5565b155b61050d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050490610c3d565b60405180910390fd5b6000905061058d565b8160ff16600060149054906101000a900460ff1660ff161061056d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161056490610c3d565b60405180910390fd5b81600060146101000a81548160ff021916908360ff160217905550600190505b919050565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61065f82610616565b810181811067ffffffffffffffff8211171561067e5761067d610627565b5b80604052505050565b60006106916105f8565b905061069d8282610656565b919050565b600067ffffffffffffffff8211156106bd576106bc610627565b5b6106c682610616565b9050602081019050919050565b82818337600083830152505050565b60006106f56106f0846106a2565b610687565b90508281526020810184848401111561071157610710610611565b5b61071c8482856106d3565b509392505050565b600082601f8301126107395761073861060c565b5b81356107498482602086016106e2565b91505092915050565b60006020828403121561076857610767610602565b5b600082013567ffffffffffffffff81111561078657610785610607565b5b61079284828501610724565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006107c68261079b565b9050919050565b6107d6816107bb565b82525050565b60006020820190506107f160008301846107cd565b92915050565b610800816107bb565b811461080b57600080fd5b50565b60008135905061081d816107f7565b92915050565b6000806040838503121561083a57610839610602565b5b600083013567ffffffffffffffff81111561085857610857610607565b5b61086485828601610724565b92505060206108758582860161080e565b9150509250929050565b6000819050919050565b60006108a461089f61089a8461079b565b61087f565b61079b565b9050919050565b60006108b682610889565b9050919050565b60006108c8826108ab565b9050919050565b6108d8816108bd565b82525050565b60006020820190506108f360008301846108cf565b92915050565b60006020828403121561090f5761090e610602565b5b600061091d8482850161080e565b91505092915050565b600081519050919050565b600081905092915050565b60005b8381101561095a57808201518184015260208101905061093f565b60008484015250505050565b600061097182610926565b61097b8185610931565b935061098b81856020860161093c565b80840191505092915050565b60006109a38284610966565b915081905092915050565b600082825260208201905092915050565b7f436f6e747261637452656769737472793a206e6f20616464726573732077617360008201527f20666f756e6420666f722074686520737065636966696564206b657900000000602082015250565b6000610a1b603c836109ae565b9150610a26826109bf565b604082019050919050565b60006020820190508181036000830152610a4a81610a0e565b9050919050565b600081519050610a60816107f7565b92915050565b600060208284031215610a7c57610a7b610602565b5b6000610a8a84828501610a51565b91505092915050565b7f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000600082015250565b6000610ac9601a836109ae565b9150610ad482610a93565b602082019050919050565b60006020820190508181036000830152610af881610abc565b9050919050565b6000610b0a82610926565b610b1481856109ae565b9350610b2481856020860161093c565b610b2d81610616565b840191505092915050565b60006040820190508181036000830152610b528185610aff565b9050610b6160208301846107cd565b9392505050565b6000819050919050565b600060ff82169050919050565b6000610b9a610b95610b9084610b68565b61087f565b610b72565b9050919050565b610baa81610b7f565b82525050565b6000602082019050610bc56000830184610ba1565b92915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000610c27602e836109ae565b9150610c3282610bcb565b604082019050919050565b60006020820190508181036000830152610c5681610c1a565b905091905056fea26469706673582212202226ac45de04800d2dc4136b45db79092a25205afc7421b9175e2148bb3506c964736f6c63430008120033",
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
	parsed, err := abi.JSON(strings.NewReader(ContractRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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
