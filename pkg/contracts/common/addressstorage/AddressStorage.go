// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addressstorage

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

// AddressStorageMetaData contains all meta data concerning the AddressStorage contract.
var AddressStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"mustAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"mustRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061001a3361001f565b61006f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610b1b8061007e6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80639e9405c3116100665780639e9405c314610121578063a224cee714610134578063a39fac1214610147578063c32d805a1461015c578063f2fde38b1461016f57600080fd5b806352efea6e146100a35780635dbe47e8146100c0578063715018a6146100eb5780638da5cb5b146100f5578063949d225d14610110575b600080fd5b6100ab610182565b60405190151581526020015b60405180910390f35b6100ab6100ce366004610911565b6001600160a01b0316600090815260016020526040902054151590565b6100f36101fa565b005b6000546040516001600160a01b0390911681526020016100b7565b6002546040519081526020016100b7565b6100f361012f366004610911565b61020e565b6100f3610142366004610949565b610279565b61014f610453565b6040516100b79190610a0e565b6100f361016a366004610911565b6104b5565b6100f361017d366004610911565b61051b565b600061018c610591565b60005b6002548110156101e75760016000600283815481106101b0576101b0610a5b565b60009182526020808320909101546001600160a01b03168352820192909252604001812055806101df81610a87565b91505061018f565b506101f4600260006108c8565b50600190565b610202610591565b61020c60006105eb565b565b6102178161063b565b6102765760405162461bcd60e51b815260206004820152602560248201527f4164647265737353746f726167653a206661696c656420746f20616464206164604482015264647265737360d81b60648201526084015b60405180910390fd5b50565b600054600160a81b900460ff16158080156102a157506000546001600160a01b90910460ff16105b806102c25750303b1580156102c25750600054600160a01b900460ff166001145b6103255760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161026d565b6000805460ff60a01b1916600160a01b1790558015610352576000805460ff60a81b1916600160a81b1790555b60005b825181101561040657600283828151811061037257610372610a5b565b60209081029190910181015182546001808201855560009485529284200180546001600160a01b0319166001600160a01b0390921691909117905560025485519092908690859081106103c7576103c7610a5b565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000208190555080806103fe90610a87565b915050610355565b50801561044f576000805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b606060028054806020026020016040519081016040528092919081815260200182805480156104ab57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161048d575b5050505050905090565b6104be816106db565b6102765760405162461bcd60e51b815260206004820152602860248201527f4164647265737353746f726167653a206661696c656420746f2072656d6f7665604482015267206164647265737360c01b606482015260840161026d565b610523610591565b6001600160a01b0381166105885760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161026d565b610276816105eb565b6000546001600160a01b0316331461020c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161026d565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610645610591565b6001600160a01b0382166000908152600160205260409020541561066b57506000919050565b60028054600180820183557f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910180546001600160a01b0319166001600160a01b038616908117909155915460009283526020919091526040909120556106d282610861565b5060015b919050565b60006106e5610591565b6001600160a01b03821660009081526001602052604090205461070a57506000919050565b6001600160a01b038216600090815260016020819052604082205460025490929161073491610aa0565b905060006002828154811061074b5761074b610a5b565b6000918252602090912001546001600160a01b0316905061076d600184610aa0565b82146107df576001600160a01b038116600090815260016020819052604090912084905581906002906107a09086610aa0565b815481106107b0576107b0610a5b565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60028054806107f0576107f0610ab9565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038716825260019052604081205561083485610861565b846001600160a01b0316816001600160a01b0316146108565761085681610861565b506001949350505050565b6001600160a01b03811660009081526001602052604090205460025481111561088c5761088c610acf565b6001600160a01b038216600090815260016020526040902054156108ba576000811161044f5761044f610acf565b801561044f5761044f610acf565b508054600082559060005260206000209081019061027691905b808211156108f657600081556001016108e2565b5090565b80356001600160a01b03811681146106d657600080fd5b60006020828403121561092357600080fd5b61092c826108fa565b9392505050565b634e487b7160e01b600052604160045260246000fd5b6000602080838503121561095c57600080fd5b823567ffffffffffffffff8082111561097457600080fd5b818501915085601f83011261098857600080fd5b81358181111561099a5761099a610933565b8060051b604051601f19603f830116810181811085821117156109bf576109bf610933565b6040529182528482019250838101850191888311156109dd57600080fd5b938501935b82851015610a02576109f3856108fa565b845293850193928501926109e2565b98975050505050505050565b6020808252825182820181905260009190848201906040850190845b81811015610a4f5783516001600160a01b031683529284019291840191600101610a2a565b50909695505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201610a9957610a99610a71565b5060010190565b81810381811115610ab357610ab3610a71565b92915050565b634e487b7160e01b600052603160045260246000fd5b634e487b7160e01b600052600160045260246000fdfea26469706673582212204247e189ec2c8e62e8db845ae3881105422f14ec535c550b12d5c2e095b00cc164736f6c63430008120033",
}

// AddressStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressStorageMetaData.ABI instead.
var AddressStorageABI = AddressStorageMetaData.ABI

// AddressStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressStorageMetaData.Bin instead.
var AddressStorageBin = AddressStorageMetaData.Bin

// DeployAddressStorage deploys a new Ethereum contract, binding an instance of AddressStorage to it.
func DeployAddressStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressStorage, error) {
	parsed, err := AddressStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressStorage{AddressStorageCaller: AddressStorageCaller{contract: contract}, AddressStorageTransactor: AddressStorageTransactor{contract: contract}, AddressStorageFilterer: AddressStorageFilterer{contract: contract}}, nil
}

// AddressStorage is an auto generated Go binding around an Ethereum contract.
type AddressStorage struct {
	AddressStorageCaller     // Read-only binding to the contract
	AddressStorageTransactor // Write-only binding to the contract
	AddressStorageFilterer   // Log filterer for contract events
}

// AddressStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressStorageSession struct {
	Contract     *AddressStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressStorageCallerSession struct {
	Contract *AddressStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AddressStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressStorageTransactorSession struct {
	Contract     *AddressStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AddressStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressStorageRaw struct {
	Contract *AddressStorage // Generic contract binding to access the raw methods on
}

// AddressStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressStorageCallerRaw struct {
	Contract *AddressStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AddressStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressStorageTransactorRaw struct {
	Contract *AddressStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressStorage creates a new instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorage(address common.Address, backend bind.ContractBackend) (*AddressStorage, error) {
	contract, err := bindAddressStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressStorage{AddressStorageCaller: AddressStorageCaller{contract: contract}, AddressStorageTransactor: AddressStorageTransactor{contract: contract}, AddressStorageFilterer: AddressStorageFilterer{contract: contract}}, nil
}

// NewAddressStorageCaller creates a new read-only instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorageCaller(address common.Address, caller bind.ContractCaller) (*AddressStorageCaller, error) {
	contract, err := bindAddressStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageCaller{contract: contract}, nil
}

// NewAddressStorageTransactor creates a new write-only instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressStorageTransactor, error) {
	contract, err := bindAddressStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressStorageTransactor{contract: contract}, nil
}

// NewAddressStorageFilterer creates a new log filterer instance of AddressStorage, bound to a specific deployed contract.
func NewAddressStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressStorageFilterer, error) {
	contract, err := bindAddressStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressStorageFilterer{contract: contract}, nil
}

// bindAddressStorage binds a generic wrapper to an already deployed contract.
func bindAddressStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorage *AddressStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorage.Contract.AddressStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorage *AddressStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.Contract.AddressStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorage *AddressStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorage.Contract.AddressStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressStorage *AddressStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressStorage *AddressStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressStorage *AddressStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressStorage.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorage *AddressStorageCaller) Contains(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "contains", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorage *AddressStorageSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorage.Contract.Contains(&_AddressStorage.CallOpts, _addr)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address _addr) view returns(bool)
func (_AddressStorage *AddressStorageCallerSession) Contains(_addr common.Address) (bool, error) {
	return _AddressStorage.Contract.Contains(&_AddressStorage.CallOpts, _addr)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorage *AddressStorageCaller) GetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorage *AddressStorageSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorage.Contract.GetAddresses(&_AddressStorage.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressStorage *AddressStorageCallerSession) GetAddresses() ([]common.Address, error) {
	return _AddressStorage.Contract.GetAddresses(&_AddressStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorage *AddressStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorage *AddressStorageSession) Owner() (common.Address, error) {
	return _AddressStorage.Contract.Owner(&_AddressStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AddressStorage *AddressStorageCallerSession) Owner() (common.Address, error) {
	return _AddressStorage.Contract.Owner(&_AddressStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorage *AddressStorageCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AddressStorage.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorage *AddressStorageSession) Size() (*big.Int, error) {
	return _AddressStorage.Contract.Size(&_AddressStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_AddressStorage *AddressStorageCallerSession) Size() (*big.Int, error) {
	return _AddressStorage.Contract.Size(&_AddressStorage.CallOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_AddressStorage *AddressStorageTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_AddressStorage *AddressStorageSession) Clear() (*types.Transaction, error) {
	return _AddressStorage.Contract.Clear(&_AddressStorage.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_AddressStorage *AddressStorageTransactorSession) Clear() (*types.Transaction, error) {
	return _AddressStorage.Contract.Clear(&_AddressStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _addrList) returns()
func (_AddressStorage *AddressStorageTransactor) Initialize(opts *bind.TransactOpts, _addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "initialize", _addrList)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _addrList) returns()
func (_AddressStorage *AddressStorageSession) Initialize(_addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Initialize(&_AddressStorage.TransactOpts, _addrList)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] _addrList) returns()
func (_AddressStorage *AddressStorageTransactorSession) Initialize(_addrList []common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.Initialize(&_AddressStorage.TransactOpts, _addrList)
}

// MustAdd is a paid mutator transaction binding the contract method 0x9e9405c3.
//
// Solidity: function mustAdd(address _addr) returns()
func (_AddressStorage *AddressStorageTransactor) MustAdd(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "mustAdd", _addr)
}

// MustAdd is a paid mutator transaction binding the contract method 0x9e9405c3.
//
// Solidity: function mustAdd(address _addr) returns()
func (_AddressStorage *AddressStorageSession) MustAdd(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustAdd(&_AddressStorage.TransactOpts, _addr)
}

// MustAdd is a paid mutator transaction binding the contract method 0x9e9405c3.
//
// Solidity: function mustAdd(address _addr) returns()
func (_AddressStorage *AddressStorageTransactorSession) MustAdd(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustAdd(&_AddressStorage.TransactOpts, _addr)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address _addr) returns()
func (_AddressStorage *AddressStorageTransactor) MustRemove(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "mustRemove", _addr)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address _addr) returns()
func (_AddressStorage *AddressStorageSession) MustRemove(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustRemove(&_AddressStorage.TransactOpts, _addr)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address _addr) returns()
func (_AddressStorage *AddressStorageTransactorSession) MustRemove(_addr common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.MustRemove(&_AddressStorage.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorage *AddressStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorage *AddressStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorage.Contract.RenounceOwnership(&_AddressStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AddressStorage *AddressStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AddressStorage.Contract.RenounceOwnership(&_AddressStorage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorage *AddressStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorage *AddressStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.TransferOwnership(&_AddressStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AddressStorage *AddressStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressStorage.Contract.TransferOwnership(&_AddressStorage.TransactOpts, newOwner)
}

// AddressStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AddressStorage contract.
type AddressStorageInitializedIterator struct {
	Event *AddressStorageInitialized // Event containing the contract specifics and raw log

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
func (it *AddressStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageInitialized)
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
		it.Event = new(AddressStorageInitialized)
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
func (it *AddressStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageInitialized represents a Initialized event raised by the AddressStorage contract.
type AddressStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AddressStorage *AddressStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*AddressStorageInitializedIterator, error) {

	logs, sub, err := _AddressStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AddressStorageInitializedIterator{contract: _AddressStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AddressStorage *AddressStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AddressStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _AddressStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageInitialized)
				if err := _AddressStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AddressStorage *AddressStorageFilterer) ParseInitialized(log types.Log) (*AddressStorageInitialized, error) {
	event := new(AddressStorageInitialized)
	if err := _AddressStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressStorage contract.
type AddressStorageOwnershipTransferredIterator struct {
	Event *AddressStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressStorageOwnershipTransferred)
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
		it.Event = new(AddressStorageOwnershipTransferred)
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
func (it *AddressStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressStorageOwnershipTransferred represents a OwnershipTransferred event raised by the AddressStorage contract.
type AddressStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorage *AddressStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressStorageOwnershipTransferredIterator{contract: _AddressStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AddressStorage *AddressStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressStorageOwnershipTransferred)
				if err := _AddressStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AddressStorage *AddressStorageFilterer) ParseOwnershipTransferred(log types.Log) (*AddressStorageOwnershipTransferred, error) {
	event := new(AddressStorageOwnershipTransferred)
	if err := _AddressStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
