// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

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

// IGatewayConfigInfo is an auto generated low-level Go binding around an user-defined struct.
type IGatewayConfigInfo struct {
	KnownWorkflows         []*big.Int
	KnownWorkflowOwners    []common.Address
	KnownCustomerContracts []common.Address
}

// IGatewayUpdateConfigData is an auto generated low-level Go binding around an user-defined struct.
type IGatewayUpdateConfigData struct {
	UpdateKnownWorkflowsEntries         []IGatewayUpdateKnownWorkflowsEntry
	UpdateKnownWorkflowOwnersEntries    []IGatewayUpdateKnownAddressesEntry
	UpdateKnownCustomerContractsEntries []IGatewayUpdateKnownAddressesEntry
}

// IGatewayUpdateKnownAddressesEntry is an auto generated low-level Go binding around an user-defined struct.
type IGatewayUpdateKnownAddressesEntry struct {
	AddressToUpdate common.Address
	IsAdding        bool
}

// IGatewayUpdateKnownWorkflowsEntry is an auto generated low-level Go binding around an user-defined struct.
type IGatewayUpdateKnownWorkflowsEntry struct {
	KnownWorkflowToUpdate *big.Int
	IsAdding              bool
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getConfigInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"knownWorkflows\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"knownWorkflowOwners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"knownCustomerContracts\",\"type\":\"address[]\"}],\"internalType\":\"structIGateway.ConfigInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gatewayOwnerAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"knownWorkflowToUpdate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"internalType\":\"structIGateway.UpdateKnownWorkflowsEntry[]\",\"name\":\"updateKnownWorkflowsEntries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addressToUpdate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"internalType\":\"structIGateway.UpdateKnownAddressesEntry[]\",\"name\":\"updateKnownWorkflowOwnersEntries\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"addressToUpdate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"internalType\":\"structIGateway.UpdateKnownAddressesEntry[]\",\"name\":\"updateKnownCustomerContractsEntries\",\"type\":\"tuple[]\"}],\"internalType\":\"structIGateway.UpdateConfigData\",\"name\":\"_updateConfigData\",\"type\":\"tuple\"}],\"name\":\"updateConfigData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addressToUpdate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"internalType\":\"structIGateway.UpdateKnownAddressesEntry[]\",\"name\":\"_updateKnownCustomerContractsEntries\",\"type\":\"tuple[]\"}],\"name\":\"updateKnownCustomerContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addressToUpdate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"internalType\":\"structIGateway.UpdateKnownAddressesEntry[]\",\"name\":\"_updateKnownWorkflowOwnersEntries\",\"type\":\"tuple[]\"}],\"name\":\"updateKnownWorkflowOwners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"knownWorkflowToUpdate\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAdding\",\"type\":\"bool\"}],\"internalType\":\"structIGateway.UpdateKnownWorkflowsEntry[]\",\"name\":\"_updateKnownWorkflowsEntries\",\"type\":\"tuple[]\"}],\"name\":\"updateKnownWorkflows\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061116a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063a3c0dd2a11610071578063a3c0dd2a14610143578063a91ee0dc14610156578063cb18c01b14610169578063ec2e65731461017c578063f2fde38b1461018f578063fef81e5b146101a257600080fd5b806315126b94146100b9578063485cc955146100ce578063715018a6146100e157806378a1bf05146100e95780637b103999146101075780638da5cb5b14610132575b600080fd5b6100cc6100c7366004610cd4565b6101b5565b005b6100cc6100dc366004610d26565b6101c9565b6100cc610305565b6100f1610319565b6040516100fe9190610da3565b60405180910390f35b60655461011a906001600160a01b031681565b6040516001600160a01b0390911681526020016100fe565b6033546001600160a01b031661011a565b6100cc610151366004610ebe565b610377565b6100cc610164366004610f73565b6103a8565b6100cc610177366004610f90565b6103d2565b6100cc61018a366004611019565b610511565b6100cc61019d366004610f73565b610524565b6100cc6101b0366004611019565b61059a565b6101bd6105ad565b6101c681610607565b50565b600054610100900460ff16158080156101e95750600054600160ff909116105b806102035750303b158015610203575060005460ff166001145b61026b5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801561028e576000805461ff0019166101001790555b6102966106b6565b606580546001600160a01b0319166001600160a01b0385161790556102ba82610524565b8015610300576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b61030d6105ad565b61031760006106e5565b565b61033d60405180606001604052806060815260200160608152602001606081525090565b60405180606001604052806103526066610737565b81526020016103616068610737565b8152602001610370606a610737565b9052919050565b61037f6105ad565b805161038a90610607565b6103996068826020015161074b565b6101c6606a826040015161074b565b6103b06105ad565b606580546001600160a01b0319166001600160a01b0392909216919091179055565b6065546001600160a01b031633146104385760405162461bcd60e51b815260206004820152602360248201527f476174657761793a206f7065726174696f6e206973206e6f74207065726d69746044820152621d195960ea1b6064820152608401610262565b838361044482826107f4565b6000856001600160a01b0316858560405161046092919061104e565b6000604051808303816000865af19150503d806000811461049d576040519150601f19603f3d011682016040523d82523d6000602084013e6104a2565b606091505b50509050806105085760405162461bcd60e51b815260206004820152602c60248201527f476174657761793a206661696c656420746f206578656375746520637573746f60448201526b1b595c8818dbdb9d1c9858dd60a21b6064820152608401610262565b50505050505050565b6105196105ad565b6101c6606a8261074b565b61052c6105ad565b6001600160a01b0381166105915760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610262565b6101c6816106e5565b6105a26105ad565b6101c660688261074b565b6033546001600160a01b031633146103175760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610262565b60005b81518110156106b2578181815181106106255761062561105e565b6020026020010151602001511561066d5761066782828151811061064b5761064b61105e565b602002602001015160000151606661091f90919063ffffffff16565b506106a0565b61069e8282815181106106825761068261105e565b602002602001015160000151606661093490919063ffffffff16565b505b806106aa8161108a565b91505061060a565b5050565b600054610100900460ff166106dd5760405162461bcd60e51b8152600401610262906110a3565b610317610940565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600061074483610970565b9392505050565b60005b8151811015610300578181815181106107695761076961105e565b602002602001015160200151156107b0576107aa82828151811061078f5761078f61105e565b602002602001015160000151846109cc90919063ffffffff16565b506107e2565b6107e08282815181106107c5576107c561105e565b602002602001015160000151846109e190919063ffffffff16565b505b806107ec8161108a565b91505061074e565b60655460405163d69cd27560e01b8152600481018490526000916001600160a01b03169063d69cd27590602401602060405180830381865afa15801561083e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086291906110ee565b90506000816001600160a01b03166108826033546001600160a01b031690565b6001600160a01b03161490508015816108c357506108a16068836109f6565b806108b257506108b2606685610a18565b806108c357506108c3606a846109f6565b156108cc575060015b806109195760405162461bcd60e51b815260206004820181905260248201527f476174657761793a206973206e6f7420616c6c6f77656420776f726b666c6f776044820152606401610262565b50505050565b600061092b8383610a30565b90505b92915050565b600061092b8383610a7f565b600054610100900460ff166109675760405162461bcd60e51b8152600401610262906110a3565b610317336106e5565b6060816000018054806020026020016040519081016040528092919081815260200182805480156109c057602002820191906000526020600020905b8154815260200190600101908083116109ac575b50505050509050919050565b600061092b836001600160a01b038416610a30565b600061092b836001600160a01b038416610a7f565b6001600160a01b0381166000908152600183016020526040812054151561092b565b6000818152600183016020526040812054151561092b565b6000818152600183016020526040812054610a775750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561092e565b50600061092e565b60008181526001830160205260408120548015610b68576000610aa360018361110b565b8554909150600090610ab79060019061110b565b9050818114610b1c576000866000018281548110610ad757610ad761105e565b9060005260206000200154905080876000018481548110610afa57610afa61105e565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080610b2d57610b2d61111e565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061092e565b600091505061092e565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610bab57610bab610b72565b60405290565b6040516060810167ffffffffffffffff81118282101715610bab57610bab610b72565b604051601f8201601f1916810167ffffffffffffffff81118282101715610bfd57610bfd610b72565b604052919050565b600067ffffffffffffffff821115610c1f57610c1f610b72565b5060051b60200190565b80358015158114610c3957600080fd5b919050565b600082601f830112610c4f57600080fd5b81356020610c64610c5f83610c05565b610bd4565b82815260069290921b84018101918181019086841115610c8357600080fd5b8286015b84811015610cc95760408189031215610ca05760008081fd5b610ca8610b88565b81358152610cb7858301610c29565b81860152835291830191604001610c87565b509695505050505050565b600060208284031215610ce657600080fd5b813567ffffffffffffffff811115610cfd57600080fd5b610d0984828501610c3e565b949350505050565b6001600160a01b03811681146101c657600080fd5b60008060408385031215610d3957600080fd5b8235610d4481610d11565b91506020830135610d5481610d11565b809150509250929050565b600081518084526020808501945080840160005b83811015610d985781516001600160a01b031687529582019590820190600101610d73565b509495945050505050565b6020808252825160608383015280516080840181905260009291820190839060a08601905b80831015610de85783518252928401926001929092019190840190610dc8565b50838701519350601f19925082868203016040870152610e088185610d5f565b93505050604085015181858403016060860152610e258382610d5f565b9695505050505050565b600082601f830112610e4057600080fd5b81356020610e50610c5f83610c05565b82815260069290921b84018101918181019086841115610e6f57600080fd5b8286015b84811015610cc95760408189031215610e8c5760008081fd5b610e94610b88565b8135610e9f81610d11565b8152610eac828601610c29565b81860152835291830191604001610e73565b600060208284031215610ed057600080fd5b813567ffffffffffffffff80821115610ee857600080fd5b9083019060608286031215610efc57600080fd5b610f04610bb1565b823582811115610f1357600080fd5b610f1f87828601610c3e565b825250602083013582811115610f3457600080fd5b610f4087828601610e2f565b602083015250604083013582811115610f5857600080fd5b610f6487828601610e2f565b60408301525095945050505050565b600060208284031215610f8557600080fd5b813561074481610d11565b60008060008060608587031215610fa657600080fd5b843593506020850135610fb881610d11565b9250604085013567ffffffffffffffff80821115610fd557600080fd5b818701915087601f830112610fe957600080fd5b813581811115610ff857600080fd5b88602082850101111561100a57600080fd5b95989497505060200194505050565b60006020828403121561102b57600080fd5b813567ffffffffffffffff81111561104257600080fd5b610d0984828501610e2f565b8183823760009101908152919050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161109c5761109c611074565b5060010190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60006020828403121561110057600080fd5b815161074481610d11565b8181038181111561092e5761092e611074565b634e487b7160e01b600052603160045260246000fdfea264697066735822122062ef5b48169ea5a4f98bc54ac92edd796c512bf7b03e93bc7f9bae645f5832e164736f6c63430008120033",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayMetaData.Bin instead.
var GatewayBin = GatewayMetaData.Bin

// DeployGateway deploys a new Ethereum contract, binding an instance of Gateway to it.
func DeployGateway(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gateway, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// GetConfigInfo is a free data retrieval call binding the contract method 0x78a1bf05.
//
// Solidity: function getConfigInfo() view returns((uint256[],address[],address[]))
func (_Gateway *GatewayCaller) GetConfigInfo(opts *bind.CallOpts) (IGatewayConfigInfo, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getConfigInfo")

	if err != nil {
		return *new(IGatewayConfigInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IGatewayConfigInfo)).(*IGatewayConfigInfo)

	return out0, err

}

// GetConfigInfo is a free data retrieval call binding the contract method 0x78a1bf05.
//
// Solidity: function getConfigInfo() view returns((uint256[],address[],address[]))
func (_Gateway *GatewaySession) GetConfigInfo() (IGatewayConfigInfo, error) {
	return _Gateway.Contract.GetConfigInfo(&_Gateway.CallOpts)
}

// GetConfigInfo is a free data retrieval call binding the contract method 0x78a1bf05.
//
// Solidity: function getConfigInfo() view returns((uint256[],address[],address[]))
func (_Gateway *GatewayCallerSession) GetConfigInfo() (IGatewayConfigInfo, error) {
	return _Gateway.Contract.GetConfigInfo(&_Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewaySession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gateway *GatewayCallerSession) Owner() (common.Address, error) {
	return _Gateway.Contract.Owner(&_Gateway.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Gateway *GatewayCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Gateway *GatewaySession) Registry() (common.Address, error) {
	return _Gateway.Contract.Registry(&_Gateway.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Gateway *GatewayCallerSession) Registry() (common.Address, error) {
	return _Gateway.Contract.Registry(&_Gateway.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _gatewayOwnerAddr) returns()
func (_Gateway *GatewayTransactor) Initialize(opts *bind.TransactOpts, _registryAddr common.Address, _gatewayOwnerAddr common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "initialize", _registryAddr, _gatewayOwnerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _gatewayOwnerAddr) returns()
func (_Gateway *GatewaySession) Initialize(_registryAddr common.Address, _gatewayOwnerAddr common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _registryAddr, _gatewayOwnerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _gatewayOwnerAddr) returns()
func (_Gateway *GatewayTransactorSession) Initialize(_registryAddr common.Address, _gatewayOwnerAddr common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _registryAddr, _gatewayOwnerAddr)
}

// Perform is a paid mutator transaction binding the contract method 0xcb18c01b.
//
// Solidity: function perform(uint256 _id, address _target, bytes _payload) returns()
func (_Gateway *GatewayTransactor) Perform(opts *bind.TransactOpts, _id *big.Int, _target common.Address, _payload []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "perform", _id, _target, _payload)
}

// Perform is a paid mutator transaction binding the contract method 0xcb18c01b.
//
// Solidity: function perform(uint256 _id, address _target, bytes _payload) returns()
func (_Gateway *GatewaySession) Perform(_id *big.Int, _target common.Address, _payload []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Perform(&_Gateway.TransactOpts, _id, _target, _payload)
}

// Perform is a paid mutator transaction binding the contract method 0xcb18c01b.
//
// Solidity: function perform(uint256 _id, address _target, bytes _payload) returns()
func (_Gateway *GatewayTransactorSession) Perform(_id *big.Int, _target common.Address, _payload []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Perform(&_Gateway.TransactOpts, _id, _target, _payload)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gateway *GatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gateway.Contract.RenounceOwnership(&_Gateway.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_Gateway *GatewayTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_Gateway *GatewaySession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetRegistry(&_Gateway.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_Gateway *GatewayTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetRegistry(&_Gateway.TransactOpts, _registry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gateway *GatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.TransferOwnership(&_Gateway.TransactOpts, newOwner)
}

// UpdateConfigData is a paid mutator transaction binding the contract method 0xa3c0dd2a.
//
// Solidity: function updateConfigData(((uint256,bool)[],(address,bool)[],(address,bool)[]) _updateConfigData) returns()
func (_Gateway *GatewayTransactor) UpdateConfigData(opts *bind.TransactOpts, _updateConfigData IGatewayUpdateConfigData) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "updateConfigData", _updateConfigData)
}

// UpdateConfigData is a paid mutator transaction binding the contract method 0xa3c0dd2a.
//
// Solidity: function updateConfigData(((uint256,bool)[],(address,bool)[],(address,bool)[]) _updateConfigData) returns()
func (_Gateway *GatewaySession) UpdateConfigData(_updateConfigData IGatewayUpdateConfigData) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateConfigData(&_Gateway.TransactOpts, _updateConfigData)
}

// UpdateConfigData is a paid mutator transaction binding the contract method 0xa3c0dd2a.
//
// Solidity: function updateConfigData(((uint256,bool)[],(address,bool)[],(address,bool)[]) _updateConfigData) returns()
func (_Gateway *GatewayTransactorSession) UpdateConfigData(_updateConfigData IGatewayUpdateConfigData) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateConfigData(&_Gateway.TransactOpts, _updateConfigData)
}

// UpdateKnownCustomerContracts is a paid mutator transaction binding the contract method 0xec2e6573.
//
// Solidity: function updateKnownCustomerContracts((address,bool)[] _updateKnownCustomerContractsEntries) returns()
func (_Gateway *GatewayTransactor) UpdateKnownCustomerContracts(opts *bind.TransactOpts, _updateKnownCustomerContractsEntries []IGatewayUpdateKnownAddressesEntry) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "updateKnownCustomerContracts", _updateKnownCustomerContractsEntries)
}

// UpdateKnownCustomerContracts is a paid mutator transaction binding the contract method 0xec2e6573.
//
// Solidity: function updateKnownCustomerContracts((address,bool)[] _updateKnownCustomerContractsEntries) returns()
func (_Gateway *GatewaySession) UpdateKnownCustomerContracts(_updateKnownCustomerContractsEntries []IGatewayUpdateKnownAddressesEntry) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateKnownCustomerContracts(&_Gateway.TransactOpts, _updateKnownCustomerContractsEntries)
}

// UpdateKnownCustomerContracts is a paid mutator transaction binding the contract method 0xec2e6573.
//
// Solidity: function updateKnownCustomerContracts((address,bool)[] _updateKnownCustomerContractsEntries) returns()
func (_Gateway *GatewayTransactorSession) UpdateKnownCustomerContracts(_updateKnownCustomerContractsEntries []IGatewayUpdateKnownAddressesEntry) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateKnownCustomerContracts(&_Gateway.TransactOpts, _updateKnownCustomerContractsEntries)
}

// UpdateKnownWorkflowOwners is a paid mutator transaction binding the contract method 0xfef81e5b.
//
// Solidity: function updateKnownWorkflowOwners((address,bool)[] _updateKnownWorkflowOwnersEntries) returns()
func (_Gateway *GatewayTransactor) UpdateKnownWorkflowOwners(opts *bind.TransactOpts, _updateKnownWorkflowOwnersEntries []IGatewayUpdateKnownAddressesEntry) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "updateKnownWorkflowOwners", _updateKnownWorkflowOwnersEntries)
}

// UpdateKnownWorkflowOwners is a paid mutator transaction binding the contract method 0xfef81e5b.
//
// Solidity: function updateKnownWorkflowOwners((address,bool)[] _updateKnownWorkflowOwnersEntries) returns()
func (_Gateway *GatewaySession) UpdateKnownWorkflowOwners(_updateKnownWorkflowOwnersEntries []IGatewayUpdateKnownAddressesEntry) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateKnownWorkflowOwners(&_Gateway.TransactOpts, _updateKnownWorkflowOwnersEntries)
}

// UpdateKnownWorkflowOwners is a paid mutator transaction binding the contract method 0xfef81e5b.
//
// Solidity: function updateKnownWorkflowOwners((address,bool)[] _updateKnownWorkflowOwnersEntries) returns()
func (_Gateway *GatewayTransactorSession) UpdateKnownWorkflowOwners(_updateKnownWorkflowOwnersEntries []IGatewayUpdateKnownAddressesEntry) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateKnownWorkflowOwners(&_Gateway.TransactOpts, _updateKnownWorkflowOwnersEntries)
}

// UpdateKnownWorkflows is a paid mutator transaction binding the contract method 0x15126b94.
//
// Solidity: function updateKnownWorkflows((uint256,bool)[] _updateKnownWorkflowsEntries) returns()
func (_Gateway *GatewayTransactor) UpdateKnownWorkflows(opts *bind.TransactOpts, _updateKnownWorkflowsEntries []IGatewayUpdateKnownWorkflowsEntry) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "updateKnownWorkflows", _updateKnownWorkflowsEntries)
}

// UpdateKnownWorkflows is a paid mutator transaction binding the contract method 0x15126b94.
//
// Solidity: function updateKnownWorkflows((uint256,bool)[] _updateKnownWorkflowsEntries) returns()
func (_Gateway *GatewaySession) UpdateKnownWorkflows(_updateKnownWorkflowsEntries []IGatewayUpdateKnownWorkflowsEntry) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateKnownWorkflows(&_Gateway.TransactOpts, _updateKnownWorkflowsEntries)
}

// UpdateKnownWorkflows is a paid mutator transaction binding the contract method 0x15126b94.
//
// Solidity: function updateKnownWorkflows((uint256,bool)[] _updateKnownWorkflowsEntries) returns()
func (_Gateway *GatewayTransactorSession) UpdateKnownWorkflows(_updateKnownWorkflowsEntries []IGatewayUpdateKnownWorkflowsEntry) (*types.Transaction, error) {
	return _Gateway.Contract.UpdateKnownWorkflows(&_Gateway.TransactOpts, _updateKnownWorkflowsEntries)
}

// GatewayInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Gateway contract.
type GatewayInitializedIterator struct {
	Event *GatewayInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayInitialized)
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
		it.Event = new(GatewayInitialized)
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
func (it *GatewayInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayInitialized represents a Initialized event raised by the Gateway contract.
type GatewayInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gateway *GatewayFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayInitializedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayInitializedIterator{contract: _Gateway.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Gateway *GatewayFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayInitialized) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayInitialized)
				if err := _Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseInitialized(log types.Log) (*GatewayInitialized, error) {
	event := new(GatewayInitialized)
	if err := _Gateway.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gateway contract.
type GatewayOwnershipTransferredIterator struct {
	Event *GatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayOwnershipTransferred)
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
		it.Event = new(GatewayOwnershipTransferred)
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
func (it *GatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayOwnershipTransferred represents a OwnershipTransferred event raised by the Gateway contract.
type GatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayOwnershipTransferredIterator{contract: _Gateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gateway *GatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayOwnershipTransferred)
				if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Gateway *GatewayFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayOwnershipTransferred, error) {
	event := new(GatewayOwnershipTransferred)
	if err := _Gateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
