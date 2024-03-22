// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rewarddistributionpool

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

// RewardDistributionPoolMetaData contains all meta data concerning the RewardDistributionPool contract.
var RewardDistributionPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollectRewards\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPositions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsOwing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610c0a806100206000396000f3fe6080604052600436106100ab5760003560e01c80638cb941cc116100645780638cb941cc146101765780638f76c0d814610196578063b5f4d38c146101ac578063b7b08d4e146101c1578063c2e8caf61461020a578063fd4c4aee1461022057600080fd5b8063372500ab146100b75780633e3b5b19146100ce57806369130451146101085780636f4a2cd01461012857806370bb45b31461013d5780637cf8b8be1461015257600080fd5b366100b257005b600080fd5b3480156100c357600080fd5b506100cc610235565b005b3480156100da57600080fd5b50600080516020610bb5833981519152546040516001600160a01b0390911681526020015b60405180910390f35b34801561011457600080fd5b506100cc610123366004610956565b61029c565b34801561013457600080fd5b506100cc610345565b34801561014957600080fd5b506100cc610486565b34801561015e57600080fd5b5061016860035481565b6040519081526020016100ff565b34801561018257600080fd5b506100cc6101913660046109fe565b610499565b3480156101a257600080fd5b5061016860015481565b3480156101b857600080fd5b506100cc6104b7565b3480156101cd57600080fd5b506101f56101dc3660046109fe565b6004602052600090815260409020805460019091015482565b604080519283526020830191909152016100ff565b34801561021657600080fd5b5061016860025481565b34801561022c57600080fd5b5061016861054a565b600061023f61054a565b905080156102995780600160008282546102599190610a38565b9091555050336000908152600460205260408120805483929061027d908490610a51565b9091555050600254336000908152600460205260409020600101555b50565b6102a46105fe565b6000829050806001600160a01b0316638e68dce46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b9190610a64565b600080546001600160a01b0319166001600160a01b03929092169190911790555033600080516020610bb5833981519152555050565b5050565b47806103b65760405162461bcd60e51b815260206004820152603560248201527f526577617264446973747269627574696f6e506f6f6c3a20616d6f756e74206d60448201527407573742062652067726561746572207468616e203605c1b60648201526084015b60405180910390fd5b6103be610682565b60008054906101000a90046001600160a01b03166001600160a01b0316638b0e9f3f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561040f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104339190610a81565b600381905561044a670de0b6b3a764000083610a9a565b6104549190610ab1565b600260008282546104659190610a51565b92505081905550806001600082825461047e9190610a51565b909155505050565b61048e610235565b61049733610767565b565b6104a16105fe565b61029981600080516020610bb583398151915255565b6104bf610235565b3360009081526004602052604081205490546104e3906001600160a01b0316610767565b600054604051631e70806560e31b8152336004820152602481018390526001600160a01b039091169063f384032890604401600060405180830381600087803b15801561052f57600080fd5b505af1158015610543573d6000803e3d6000fd5b5050505050565b33600090815260046020526040812060010154600254829161056b91610a38565b6000546040516303d3b32360e51b8152336004820152919250670de0b6b3a76400009183916001600160a01b031690637a76646090602401602060405180830381865afa1580156105c0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105e49190610a81565b6105ee9190610a9a565b6105f89190610ab1565b91505090565b6000610616600080516020610bb58339815191525490565b90506001600160a01b038116158061063657506001600160a01b03811633145b6102995760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f7200000000000060448201526064016103ad565b60008060009054906101000a90046001600160a01b03166001600160a01b031663b7ab4db56040518163ffffffff1660e01b8152600401600060405180830381865afa1580156106d6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106fe9190810190610ad3565b905060005b8151811015610341576002546004600084848151811061072557610725610b85565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060010181905550808061075f90610b9b565b915050610703565b33600090815260046020526040902054806107e25760405162461bcd60e51b815260206004820152603560248201527f526577617264446973747269627574696f6e506f6f6c3a20726577617264206d60448201527407573742062652067726561746572207468616e203605c1b60648201526084016103ad565b3360009081526004602052604081208054839290610801908490610a38565b90915550506040516000906001600160a01b0384169061520890849084818181858888f193505050503d8060008114610856576040519150601f19603f3d011682016040523d82523d6000602084013e61085b565b606091505b50509050806108bc5760405162461bcd60e51b815260206004820152602760248201527f526577617264446973747269627574696f6e506f6f6c3a207472616e736665726044820152660819985a5b195960ca1b60648201526084016103ad565b60408051338152602081018490527f2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a717910160405180910390a1505050565b6001600160a01b038116811461029957600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561094e5761094e61090f565b604052919050565b6000806040838503121561096957600080fd5b8235610974816108fa565b915060208381013567ffffffffffffffff8082111561099257600080fd5b818601915086601f8301126109a657600080fd5b8135818111156109b8576109b861090f565b6109ca601f8201601f19168501610925565b915080825287848285010111156109e057600080fd5b80848401858401376000848284010152508093505050509250929050565b600060208284031215610a1057600080fd5b8135610a1b816108fa565b9392505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610a4b57610a4b610a22565b92915050565b80820180821115610a4b57610a4b610a22565b600060208284031215610a7657600080fd5b8151610a1b816108fa565b600060208284031215610a9357600080fd5b5051919050565b8082028115828204841417610a4b57610a4b610a22565b600082610ace57634e487b7160e01b600052601260045260246000fd5b500490565b60006020808385031215610ae657600080fd5b825167ffffffffffffffff80821115610afe57600080fd5b818501915085601f830112610b1257600080fd5b815181811115610b2457610b2461090f565b8060051b9150610b35848301610925565b8181529183018401918481019088841115610b4f57600080fd5b938501935b83851015610b795784519250610b69836108fa565b8282529385019390850190610b54565b98975050505050505050565b634e487b7160e01b600052603260045260246000fd5b600060018201610bad57610bad610a22565b506001019056fe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a2646970667358221220f690d17b47aa3c561aec08649e22be8a0ca018d10eefe84828c6c9f190c1598964736f6c63430008120033",
}

// RewardDistributionPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardDistributionPoolMetaData.ABI instead.
var RewardDistributionPoolABI = RewardDistributionPoolMetaData.ABI

// RewardDistributionPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RewardDistributionPoolMetaData.Bin instead.
var RewardDistributionPoolBin = RewardDistributionPoolMetaData.Bin

// DeployRewardDistributionPool deploys a new Ethereum contract, binding an instance of RewardDistributionPool to it.
func DeployRewardDistributionPool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RewardDistributionPool, error) {
	parsed, err := RewardDistributionPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RewardDistributionPoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RewardDistributionPool{RewardDistributionPoolCaller: RewardDistributionPoolCaller{contract: contract}, RewardDistributionPoolTransactor: RewardDistributionPoolTransactor{contract: contract}, RewardDistributionPoolFilterer: RewardDistributionPoolFilterer{contract: contract}}, nil
}

// RewardDistributionPool is an auto generated Go binding around an Ethereum contract.
type RewardDistributionPool struct {
	RewardDistributionPoolCaller     // Read-only binding to the contract
	RewardDistributionPoolTransactor // Write-only binding to the contract
	RewardDistributionPoolFilterer   // Log filterer for contract events
}

// RewardDistributionPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardDistributionPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributionPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardDistributionPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributionPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardDistributionPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardDistributionPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardDistributionPoolSession struct {
	Contract     *RewardDistributionPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RewardDistributionPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardDistributionPoolCallerSession struct {
	Contract *RewardDistributionPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// RewardDistributionPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardDistributionPoolTransactorSession struct {
	Contract     *RewardDistributionPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// RewardDistributionPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardDistributionPoolRaw struct {
	Contract *RewardDistributionPool // Generic contract binding to access the raw methods on
}

// RewardDistributionPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardDistributionPoolCallerRaw struct {
	Contract *RewardDistributionPoolCaller // Generic read-only contract binding to access the raw methods on
}

// RewardDistributionPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardDistributionPoolTransactorRaw struct {
	Contract *RewardDistributionPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardDistributionPool creates a new instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPool(address common.Address, backend bind.ContractBackend) (*RewardDistributionPool, error) {
	contract, err := bindRewardDistributionPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPool{RewardDistributionPoolCaller: RewardDistributionPoolCaller{contract: contract}, RewardDistributionPoolTransactor: RewardDistributionPoolTransactor{contract: contract}, RewardDistributionPoolFilterer: RewardDistributionPoolFilterer{contract: contract}}, nil
}

// NewRewardDistributionPoolCaller creates a new read-only instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPoolCaller(address common.Address, caller bind.ContractCaller) (*RewardDistributionPoolCaller, error) {
	contract, err := bindRewardDistributionPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolCaller{contract: contract}, nil
}

// NewRewardDistributionPoolTransactor creates a new write-only instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardDistributionPoolTransactor, error) {
	contract, err := bindRewardDistributionPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolTransactor{contract: contract}, nil
}

// NewRewardDistributionPoolFilterer creates a new log filterer instance of RewardDistributionPool, bound to a specific deployed contract.
func NewRewardDistributionPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardDistributionPoolFilterer, error) {
	contract, err := bindRewardDistributionPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolFilterer{contract: contract}, nil
}

// bindRewardDistributionPool binds a generic wrapper to an already deployed contract.
func bindRewardDistributionPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardDistributionPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributionPool *RewardDistributionPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributionPool.Contract.RewardDistributionPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributionPool *RewardDistributionPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.RewardDistributionPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributionPool *RewardDistributionPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.RewardDistributionPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardDistributionPool *RewardDistributionPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardDistributionPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardDistributionPool *RewardDistributionPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardDistributionPool *RewardDistributionPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.contract.Transact(opts, method, params...)
}

// CollectedRewards is a free data retrieval call binding the contract method 0x8f76c0d8.
//
// Solidity: function collectedRewards() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) CollectedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "collectedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollectedRewards is a free data retrieval call binding the contract method 0x8f76c0d8.
//
// Solidity: function collectedRewards() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) CollectedRewards() (*big.Int, error) {
	return _RewardDistributionPool.Contract.CollectedRewards(&_RewardDistributionPool.CallOpts)
}

// CollectedRewards is a free data retrieval call binding the contract method 0x8f76c0d8.
//
// Solidity: function collectedRewards() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) CollectedRewards() (*big.Int, error) {
	return _RewardDistributionPool.Contract.CollectedRewards(&_RewardDistributionPool.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_RewardDistributionPool *RewardDistributionPoolCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_RewardDistributionPool *RewardDistributionPoolSession) GetInjector() (common.Address, error) {
	return _RewardDistributionPool.Contract.GetInjector(&_RewardDistributionPool.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) GetInjector() (common.Address, error) {
	return _RewardDistributionPool.Contract.GetInjector(&_RewardDistributionPool.CallOpts)
}

// ProvidedStake is a free data retrieval call binding the contract method 0x7cf8b8be.
//
// Solidity: function providedStake() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) ProvidedStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "providedStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProvidedStake is a free data retrieval call binding the contract method 0x7cf8b8be.
//
// Solidity: function providedStake() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) ProvidedStake() (*big.Int, error) {
	return _RewardDistributionPool.Contract.ProvidedStake(&_RewardDistributionPool.CallOpts)
}

// ProvidedStake is a free data retrieval call binding the contract method 0x7cf8b8be.
//
// Solidity: function providedStake() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) ProvidedStake() (*big.Int, error) {
	return _RewardDistributionPool.Contract.ProvidedStake(&_RewardDistributionPool.CallOpts)
}

// RewardPositions is a free data retrieval call binding the contract method 0xb7b08d4e.
//
// Solidity: function rewardPositions(address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_RewardDistributionPool *RewardDistributionPoolCaller) RewardPositions(opts *bind.CallOpts, arg0 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "rewardPositions", arg0)

	outstruct := new(struct {
		Balance          *big.Int
		LastRewardPoints *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastRewardPoints = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardPositions is a free data retrieval call binding the contract method 0xb7b08d4e.
//
// Solidity: function rewardPositions(address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_RewardDistributionPool *RewardDistributionPoolSession) RewardPositions(arg0 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	return _RewardDistributionPool.Contract.RewardPositions(&_RewardDistributionPool.CallOpts, arg0)
}

// RewardPositions is a free data retrieval call binding the contract method 0xb7b08d4e.
//
// Solidity: function rewardPositions(address ) view returns(uint256 balance, uint256 lastRewardPoints)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) RewardPositions(arg0 common.Address) (struct {
	Balance          *big.Int
	LastRewardPoints *big.Int
}, error) {
	return _RewardDistributionPool.Contract.RewardPositions(&_RewardDistributionPool.CallOpts, arg0)
}

// RewardsOwing is a free data retrieval call binding the contract method 0xfd4c4aee.
//
// Solidity: function rewardsOwing() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) RewardsOwing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "rewardsOwing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardsOwing is a free data retrieval call binding the contract method 0xfd4c4aee.
//
// Solidity: function rewardsOwing() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) RewardsOwing() (*big.Int, error) {
	return _RewardDistributionPool.Contract.RewardsOwing(&_RewardDistributionPool.CallOpts)
}

// RewardsOwing is a free data retrieval call binding the contract method 0xfd4c4aee.
//
// Solidity: function rewardsOwing() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) RewardsOwing() (*big.Int, error) {
	return _RewardDistributionPool.Contract.RewardsOwing(&_RewardDistributionPool.CallOpts)
}

// TotalRewardPoints is a free data retrieval call binding the contract method 0xc2e8caf6.
//
// Solidity: function totalRewardPoints() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCaller) TotalRewardPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "totalRewardPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewardPoints is a free data retrieval call binding the contract method 0xc2e8caf6.
//
// Solidity: function totalRewardPoints() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolSession) TotalRewardPoints() (*big.Int, error) {
	return _RewardDistributionPool.Contract.TotalRewardPoints(&_RewardDistributionPool.CallOpts)
}

// TotalRewardPoints is a free data retrieval call binding the contract method 0xc2e8caf6.
//
// Solidity: function totalRewardPoints() view returns(uint256)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) TotalRewardPoints() (*big.Int, error) {
	return _RewardDistributionPool.Contract.TotalRewardPoints(&_RewardDistributionPool.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) ClaimRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "claimRewards")
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) ClaimRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ClaimRewards(&_RewardDistributionPool.TransactOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x372500ab.
//
// Solidity: function claimRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) ClaimRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ClaimRewards(&_RewardDistributionPool.TransactOpts)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x70bb45b3.
//
// Solidity: function collectRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) CollectRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "collectRewards")
}

// CollectRewards is a paid mutator transaction binding the contract method 0x70bb45b3.
//
// Solidity: function collectRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) CollectRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.CollectRewards(&_RewardDistributionPool.TransactOpts)
}

// CollectRewards is a paid mutator transaction binding the contract method 0x70bb45b3.
//
// Solidity: function collectRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) CollectRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.CollectRewards(&_RewardDistributionPool.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) DistributeRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.DistributeRewards(&_RewardDistributionPool.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.DistributeRewards(&_RewardDistributionPool.TransactOpts)
}

// ReinvestRewards is a paid mutator transaction binding the contract method 0xb5f4d38c.
//
// Solidity: function reinvestRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) ReinvestRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "reinvestRewards")
}

// ReinvestRewards is a paid mutator transaction binding the contract method 0xb5f4d38c.
//
// Solidity: function reinvestRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) ReinvestRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ReinvestRewards(&_RewardDistributionPool.TransactOpts)
}

// ReinvestRewards is a paid mutator transaction binding the contract method 0xb5f4d38c.
//
// Solidity: function reinvestRewards() returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) ReinvestRewards() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.ReinvestRewards(&_RewardDistributionPool.TransactOpts)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) SetDependencies(opts *bind.TransactOpts, _contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "setDependencies", _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.SetDependencies(&_RewardDistributionPool.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.SetDependencies(&_RewardDistributionPool.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.SetInjector(&_RewardDistributionPool.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.SetInjector(&_RewardDistributionPool.TransactOpts, injector_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) Receive() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Receive(&_RewardDistributionPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Receive(&_RewardDistributionPool.TransactOpts)
}

// RewardDistributionPoolCollectRewardsIterator is returned from FilterCollectRewards and is used to iterate over the raw logs and unpacked data for CollectRewards events raised by the RewardDistributionPool contract.
type RewardDistributionPoolCollectRewardsIterator struct {
	Event *RewardDistributionPoolCollectRewards // Event containing the contract specifics and raw log

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
func (it *RewardDistributionPoolCollectRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDistributionPoolCollectRewards)
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
		it.Event = new(RewardDistributionPoolCollectRewards)
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
func (it *RewardDistributionPoolCollectRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDistributionPoolCollectRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDistributionPoolCollectRewards represents a CollectRewards event raised by the RewardDistributionPool contract.
type RewardDistributionPoolCollectRewards struct {
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectRewards is a free log retrieval operation binding the contract event 0x2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a717.
//
// Solidity: event CollectRewards(address validator, uint256 amount)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) FilterCollectRewards(opts *bind.FilterOpts) (*RewardDistributionPoolCollectRewardsIterator, error) {

	logs, sub, err := _RewardDistributionPool.contract.FilterLogs(opts, "CollectRewards")
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolCollectRewardsIterator{contract: _RewardDistributionPool.contract, event: "CollectRewards", logs: logs, sub: sub}, nil
}

// WatchCollectRewards is a free log subscription operation binding the contract event 0x2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a717.
//
// Solidity: event CollectRewards(address validator, uint256 amount)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) WatchCollectRewards(opts *bind.WatchOpts, sink chan<- *RewardDistributionPoolCollectRewards) (event.Subscription, error) {

	logs, sub, err := _RewardDistributionPool.contract.WatchLogs(opts, "CollectRewards")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDistributionPoolCollectRewards)
				if err := _RewardDistributionPool.contract.UnpackLog(event, "CollectRewards", log); err != nil {
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

// ParseCollectRewards is a log parse operation binding the contract event 0x2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a717.
//
// Solidity: event CollectRewards(address validator, uint256 amount)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) ParseCollectRewards(log types.Log) (*RewardDistributionPoolCollectRewards, error) {
	event := new(RewardDistributionPoolCollectRewards)
	if err := _RewardDistributionPool.contract.UnpackLog(event, "CollectRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
