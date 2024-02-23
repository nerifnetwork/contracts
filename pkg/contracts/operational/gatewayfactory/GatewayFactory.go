// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gatewayfactory

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

// GatewayFactoryMetaData contains all meta data concerning the GatewayFactory contract.
var GatewayFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gatewayAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gatewayOwnerAddr\",\"type\":\"address\"}],\"name\":\"GatewayDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gatewayOwner\",\"type\":\"address\"}],\"name\":\"deployGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gatewayBeacon\",\"outputs\":[{\"internalType\":\"contractProxyBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGatewayImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gatewayImpl\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registryAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGatewayImpl\",\"type\":\"address\"}],\"name\":\"setNewImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611531806100206000396000f3fe60806040523480156200001157600080fd5b50600436106200009f5760003560e01c80638da5cb5b116200006e5780638da5cb5b146200010157806395f9afe41462000113578063a65b5873146200012a578063adb3ce92146200013e578063f2fde38b146200015557600080fd5b806341a8328b14620000a4578063485cc95514620000ca57806351da2eaa14620000e3578063715018a614620000f7575b600080fd5b620000ae6200016c565b6040516001600160a01b03909116815260200160405180910390f35b620000e1620000db366004620007c0565b620001e2565b005b606654620000ae906001600160a01b031681565b620000e162000368565b6033546001600160a01b0316620000ae565b620000ae62000124366004620007fe565b62000380565b606554620000ae906001600160a01b031681565b620000e16200014f366004620007fe565b620004fe565b620000e162000166366004620007fe565b62000516565b60655460408051635c60da1b60e01b815290516000926001600160a01b031691635c60da1b9160048083019260209291908290030181865afa158015620001b7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001dd919062000825565b905090565b600054610100900460ff1615808015620002035750600054600160ff909116105b806200021f5750303b1580156200021f575060005460ff166001145b620002885760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015620002ac576000805461ff0019166101001790555b620002b662000592565b604051620002c4906200078e565b604051809103906000f080158015620002e1573d6000803e3d6000fd5b50606580546001600160a01b03199081166001600160a01b0393841617909155606680549091169185169190911790556200031c82620005c6565b801562000363576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b62000372620006ab565b6200037e600062000707565b565b6066546000906001600160a01b03163314620003f25760405162461bcd60e51b815260206004820152602a60248201527f47617465776179466163746f72793a206f7065726174696f6e206973206e6f74604482015269081c195c9b5a5d1d195960b21b60648201526084016200027f565b6065546040516000916001600160a01b03169062000410906200079c565b6001600160a01b039091168152604060208201819052600090820152606001604051809103906000f0801580156200044c573d6000803e3d6000fd5b5060405163485cc95560e01b81523360048201526001600160a01b0385811660248301529192509082169063485cc95590604401600060405180830381600087803b1580156200049b57600080fd5b505af1158015620004b0573d6000803e3d6000fd5b5050604080516001600160a01b038086168252871660208201527f4b209f3d620ba8f4e206ec57dd17cebe32624f42eb77c06bcfd195be78b6c2c2935001905060405180910390a192915050565b62000508620006ab565b6200051381620005c6565b50565b62000520620006ab565b6001600160a01b038116620005875760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016200027f565b620005138162000707565b600054610100900460ff16620005bc5760405162461bcd60e51b81526004016200027f9062000845565b6200037e62000759565b60655460408051635c60da1b60e01b815290516001600160a01b03808516931691635c60da1b9160048083019260209291908290030181865afa15801562000612573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000638919062000825565b6001600160a01b0316146200051357606554604051631b2ce7f360e11b81526001600160a01b03838116600483015290911690633659cfe690602401600060405180830381600087803b1580156200068f57600080fd5b505af1158015620006a4573d6000803e3d6000fd5b5050505050565b6033546001600160a01b031633146200037e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016200027f565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16620007835760405162461bcd60e51b81526004016200027f9062000845565b6200037e3362000707565b61030a806200089183390190565b6109618062000b9b83390190565b6001600160a01b03811681146200051357600080fd5b60008060408385031215620007d457600080fd5b8235620007e181620007aa565b91506020830135620007f381620007aa565b809150509250929050565b6000602082840312156200081157600080fd5b81356200081e81620007aa565b9392505050565b6000602082840312156200083857600080fd5b81516200081e81620007aa565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fe60a060405234801561001057600080fd5b5033806100895760405162461bcd60e51b815260206004820152603360248201527f5065726d616e656e744f776e61626c653a207a65726f2061646472657373206360448201527f616e206e6f7420626520746865206f776e657200000000000000000000000000606482015260840160405180910390fd5b6001600160a01b03166080526080516102546100b6600039600081816086015261016401526102546000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80633659cfe6146100465780635c60da1b1461005b5780638da5cb5b14610084575b600080fd5b6100596100543660046101ee565b6100aa565b005b6000546001600160a01b03165b6040516001600160a01b03909116815260200160405180910390f35b7f0000000000000000000000000000000000000000000000000000000000000000610068565b6100b2610162565b6001600160a01b0381163b61010e5760405162461bcd60e51b815260206004820152601b60248201527f50726f7879426561636f6e3a206e6f74206120636f6e7472616374000000000060448201526064015b60405180910390fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b9060200160405180910390a150565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633146101ec5760405162461bcd60e51b815260206004820152602960248201527f5065726d616e656e744f776e61626c653a2063616c6c6572206973206e6f74206044820152683a34329037bbb732b960b91b6064820152608401610105565b565b60006020828403121561020057600080fd5b81356001600160a01b038116811461021757600080fd5b939250505056fea2646970667358221220e6160f1d424829469272c692f65daad9520615411489d31c616cfaf764211e7f64736f6c6343000812003360806040526040516109613803806109618339810160408190526100229161045f565b818161003082826000610039565b50505050610589565b61004283610104565b6040516001600160a01b038416907f1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e90600090a26000825111806100835750805b156100ff576100fd836001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156100c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100ed919061051f565b836102a760201b61008b1760201c565b505b505050565b610117816102d360201b6100b71760201c565b6101765760405162461bcd60e51b815260206004820152602560248201527f455243313936373a206e657720626561636f6e206973206e6f74206120636f6e6044820152641d1c9858dd60da1b60648201526084015b60405180910390fd5b6101ea816001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101db919061051f565b6102d360201b6100b71760201c565b61024f5760405162461bcd60e51b815260206004820152603060248201527f455243313936373a20626561636f6e20696d706c656d656e746174696f6e206960448201526f1cc81b9bdd08184818dbdb9d1c9858dd60821b606482015260840161016d565b806102867fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d5060001b6102e260201b6100c61760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b60606102cc838360405180606001604052806027815260200161093a602791396102e5565b9392505050565b6001600160a01b03163b151590565b90565b6060600080856001600160a01b031685604051610302919061053a565b600060405180830381855af49150503d806000811461033d576040519150601f19603f3d011682016040523d82523d6000602084013e610342565b606091505b5090925090506103548683838761035e565b9695505050505050565b606083156103cd5782516000036103c6576001600160a01b0385163b6103c65760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161016d565b50816103d7565b6103d783836103df565b949350505050565b8151156103ef5781518083602001fd5b8060405162461bcd60e51b815260040161016d9190610556565b80516001600160a01b038116811461042057600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b8381101561045657818101518382015260200161043e565b50506000910152565b6000806040838503121561047257600080fd5b61047b83610409565b60208401519092506001600160401b038082111561049857600080fd5b818501915085601f8301126104ac57600080fd5b8151818111156104be576104be610425565b604051601f8201601f19908116603f011681019083821181831017156104e6576104e6610425565b816040528281528860208487010111156104ff57600080fd5b61051083602083016020880161043b565b80955050505050509250929050565b60006020828403121561053157600080fd5b6102cc82610409565b6000825161054c81846020870161043b565b9190910192915050565b602081526000825180602084015261057581604085016020870161043b565b601f01601f19169190910160400192915050565b6103a2806105986000396000f3fe6080604052600436106100225760003560e01c80635c60da1b1461003957610031565b366100315761002f61006a565b005b61002f61006a565b34801561004557600080fd5b5061004e61007c565b6040516001600160a01b03909116815260200160405180910390f35b61007a6100756100c9565b61015d565b565b60006100866100c9565b905090565b60606100b0838360405180606001604052806027815260200161034660279139610181565b9392505050565b6001600160a01b03163b151590565b90565b60006100fc7fa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50546001600160a01b031690565b6001600160a01b0316635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610139573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061008691906102a9565b3660008037600080366000845af43d6000803e80801561017c573d6000f35b3d6000fd5b6060600080856001600160a01b03168560405161019e91906102f6565b600060405180830381855af49150503d80600081146101d9576040519150601f19603f3d011682016040523d82523d6000602084013e6101de565b606091505b50915091506101ef868383876101f9565b9695505050505050565b6060831561026d578251600003610266576001600160a01b0385163b6102665760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b5081610277565b610277838361027f565b949350505050565b81511561028f5781518083602001fd5b8060405162461bcd60e51b815260040161025d9190610312565b6000602082840312156102bb57600080fd5b81516001600160a01b03811681146100b057600080fd5b60005b838110156102ed5781810151838201526020016102d5565b50506000910152565b600082516103088184602087016102d2565b9190910192915050565b60208152600082518060208401526103318160408501602087016102d2565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212205de758e8f864b470b45b03032818b972b7f357e4aaf1ff7820bb63597615b25264736f6c63430008120033416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220026a9c1d085ded7a5d08e0a0984d4ffb20b7d325f9fecdb31b3b45a6e84bcf5a64736f6c63430008120033",
}

// GatewayFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayFactoryMetaData.ABI instead.
var GatewayFactoryABI = GatewayFactoryMetaData.ABI

// GatewayFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayFactoryMetaData.Bin instead.
var GatewayFactoryBin = GatewayFactoryMetaData.Bin

// DeployGatewayFactory deploys a new Ethereum contract, binding an instance of GatewayFactory to it.
func DeployGatewayFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayFactory, error) {
	parsed, err := GatewayFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayFactory{GatewayFactoryCaller: GatewayFactoryCaller{contract: contract}, GatewayFactoryTransactor: GatewayFactoryTransactor{contract: contract}, GatewayFactoryFilterer: GatewayFactoryFilterer{contract: contract}}, nil
}

// GatewayFactory is an auto generated Go binding around an Ethereum contract.
type GatewayFactory struct {
	GatewayFactoryCaller     // Read-only binding to the contract
	GatewayFactoryTransactor // Write-only binding to the contract
	GatewayFactoryFilterer   // Log filterer for contract events
}

// GatewayFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayFactorySession struct {
	Contract     *GatewayFactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayFactoryCallerSession struct {
	Contract *GatewayFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GatewayFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayFactoryTransactorSession struct {
	Contract     *GatewayFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GatewayFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayFactoryRaw struct {
	Contract *GatewayFactory // Generic contract binding to access the raw methods on
}

// GatewayFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayFactoryCallerRaw struct {
	Contract *GatewayFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayFactoryTransactorRaw struct {
	Contract *GatewayFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayFactory creates a new instance of GatewayFactory, bound to a specific deployed contract.
func NewGatewayFactory(address common.Address, backend bind.ContractBackend) (*GatewayFactory, error) {
	contract, err := bindGatewayFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayFactory{GatewayFactoryCaller: GatewayFactoryCaller{contract: contract}, GatewayFactoryTransactor: GatewayFactoryTransactor{contract: contract}, GatewayFactoryFilterer: GatewayFactoryFilterer{contract: contract}}, nil
}

// NewGatewayFactoryCaller creates a new read-only instance of GatewayFactory, bound to a specific deployed contract.
func NewGatewayFactoryCaller(address common.Address, caller bind.ContractCaller) (*GatewayFactoryCaller, error) {
	contract, err := bindGatewayFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayFactoryCaller{contract: contract}, nil
}

// NewGatewayFactoryTransactor creates a new write-only instance of GatewayFactory, bound to a specific deployed contract.
func NewGatewayFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayFactoryTransactor, error) {
	contract, err := bindGatewayFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayFactoryTransactor{contract: contract}, nil
}

// NewGatewayFactoryFilterer creates a new log filterer instance of GatewayFactory, bound to a specific deployed contract.
func NewGatewayFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFactoryFilterer, error) {
	contract, err := bindGatewayFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFactoryFilterer{contract: contract}, nil
}

// bindGatewayFactory binds a generic wrapper to an already deployed contract.
func bindGatewayFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GatewayFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayFactory *GatewayFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayFactory.Contract.GatewayFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayFactory *GatewayFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayFactory.Contract.GatewayFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayFactory *GatewayFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayFactory.Contract.GatewayFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayFactory *GatewayFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayFactory *GatewayFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayFactory *GatewayFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayFactory.Contract.contract.Transact(opts, method, params...)
}

// GatewayBeacon is a free data retrieval call binding the contract method 0xa65b5873.
//
// Solidity: function gatewayBeacon() view returns(address)
func (_GatewayFactory *GatewayFactoryCaller) GatewayBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayFactory.contract.Call(opts, &out, "gatewayBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GatewayBeacon is a free data retrieval call binding the contract method 0xa65b5873.
//
// Solidity: function gatewayBeacon() view returns(address)
func (_GatewayFactory *GatewayFactorySession) GatewayBeacon() (common.Address, error) {
	return _GatewayFactory.Contract.GatewayBeacon(&_GatewayFactory.CallOpts)
}

// GatewayBeacon is a free data retrieval call binding the contract method 0xa65b5873.
//
// Solidity: function gatewayBeacon() view returns(address)
func (_GatewayFactory *GatewayFactoryCallerSession) GatewayBeacon() (common.Address, error) {
	return _GatewayFactory.Contract.GatewayBeacon(&_GatewayFactory.CallOpts)
}

// GetGatewayImplementation is a free data retrieval call binding the contract method 0x41a8328b.
//
// Solidity: function getGatewayImplementation() view returns(address)
func (_GatewayFactory *GatewayFactoryCaller) GetGatewayImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayFactory.contract.Call(opts, &out, "getGatewayImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGatewayImplementation is a free data retrieval call binding the contract method 0x41a8328b.
//
// Solidity: function getGatewayImplementation() view returns(address)
func (_GatewayFactory *GatewayFactorySession) GetGatewayImplementation() (common.Address, error) {
	return _GatewayFactory.Contract.GetGatewayImplementation(&_GatewayFactory.CallOpts)
}

// GetGatewayImplementation is a free data retrieval call binding the contract method 0x41a8328b.
//
// Solidity: function getGatewayImplementation() view returns(address)
func (_GatewayFactory *GatewayFactoryCallerSession) GetGatewayImplementation() (common.Address, error) {
	return _GatewayFactory.Contract.GetGatewayImplementation(&_GatewayFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayFactory *GatewayFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayFactory *GatewayFactorySession) Owner() (common.Address, error) {
	return _GatewayFactory.Contract.Owner(&_GatewayFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayFactory *GatewayFactoryCallerSession) Owner() (common.Address, error) {
	return _GatewayFactory.Contract.Owner(&_GatewayFactory.CallOpts)
}

// RegistryAddr is a free data retrieval call binding the contract method 0x51da2eaa.
//
// Solidity: function registryAddr() view returns(address)
func (_GatewayFactory *GatewayFactoryCaller) RegistryAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayFactory.contract.Call(opts, &out, "registryAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryAddr is a free data retrieval call binding the contract method 0x51da2eaa.
//
// Solidity: function registryAddr() view returns(address)
func (_GatewayFactory *GatewayFactorySession) RegistryAddr() (common.Address, error) {
	return _GatewayFactory.Contract.RegistryAddr(&_GatewayFactory.CallOpts)
}

// RegistryAddr is a free data retrieval call binding the contract method 0x51da2eaa.
//
// Solidity: function registryAddr() view returns(address)
func (_GatewayFactory *GatewayFactoryCallerSession) RegistryAddr() (common.Address, error) {
	return _GatewayFactory.Contract.RegistryAddr(&_GatewayFactory.CallOpts)
}

// DeployGateway is a paid mutator transaction binding the contract method 0x95f9afe4.
//
// Solidity: function deployGateway(address _gatewayOwner) returns(address)
func (_GatewayFactory *GatewayFactoryTransactor) DeployGateway(opts *bind.TransactOpts, _gatewayOwner common.Address) (*types.Transaction, error) {
	return _GatewayFactory.contract.Transact(opts, "deployGateway", _gatewayOwner)
}

// DeployGateway is a paid mutator transaction binding the contract method 0x95f9afe4.
//
// Solidity: function deployGateway(address _gatewayOwner) returns(address)
func (_GatewayFactory *GatewayFactorySession) DeployGateway(_gatewayOwner common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.DeployGateway(&_GatewayFactory.TransactOpts, _gatewayOwner)
}

// DeployGateway is a paid mutator transaction binding the contract method 0x95f9afe4.
//
// Solidity: function deployGateway(address _gatewayOwner) returns(address)
func (_GatewayFactory *GatewayFactoryTransactorSession) DeployGateway(_gatewayOwner common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.DeployGateway(&_GatewayFactory.TransactOpts, _gatewayOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _gatewayImpl) returns()
func (_GatewayFactory *GatewayFactoryTransactor) Initialize(opts *bind.TransactOpts, _registryAddr common.Address, _gatewayImpl common.Address) (*types.Transaction, error) {
	return _GatewayFactory.contract.Transact(opts, "initialize", _registryAddr, _gatewayImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _gatewayImpl) returns()
func (_GatewayFactory *GatewayFactorySession) Initialize(_registryAddr common.Address, _gatewayImpl common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.Initialize(&_GatewayFactory.TransactOpts, _registryAddr, _gatewayImpl)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _gatewayImpl) returns()
func (_GatewayFactory *GatewayFactoryTransactorSession) Initialize(_registryAddr common.Address, _gatewayImpl common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.Initialize(&_GatewayFactory.TransactOpts, _registryAddr, _gatewayImpl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayFactory *GatewayFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayFactory *GatewayFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayFactory.Contract.RenounceOwnership(&_GatewayFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayFactory *GatewayFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayFactory.Contract.RenounceOwnership(&_GatewayFactory.TransactOpts)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address _newGatewayImpl) returns()
func (_GatewayFactory *GatewayFactoryTransactor) SetNewImplementation(opts *bind.TransactOpts, _newGatewayImpl common.Address) (*types.Transaction, error) {
	return _GatewayFactory.contract.Transact(opts, "setNewImplementation", _newGatewayImpl)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address _newGatewayImpl) returns()
func (_GatewayFactory *GatewayFactorySession) SetNewImplementation(_newGatewayImpl common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.SetNewImplementation(&_GatewayFactory.TransactOpts, _newGatewayImpl)
}

// SetNewImplementation is a paid mutator transaction binding the contract method 0xadb3ce92.
//
// Solidity: function setNewImplementation(address _newGatewayImpl) returns()
func (_GatewayFactory *GatewayFactoryTransactorSession) SetNewImplementation(_newGatewayImpl common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.SetNewImplementation(&_GatewayFactory.TransactOpts, _newGatewayImpl)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayFactory *GatewayFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayFactory *GatewayFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.TransferOwnership(&_GatewayFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayFactory *GatewayFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayFactory.Contract.TransferOwnership(&_GatewayFactory.TransactOpts, newOwner)
}

// GatewayFactoryGatewayDeployedIterator is returned from FilterGatewayDeployed and is used to iterate over the raw logs and unpacked data for GatewayDeployed events raised by the GatewayFactory contract.
type GatewayFactoryGatewayDeployedIterator struct {
	Event *GatewayFactoryGatewayDeployed // Event containing the contract specifics and raw log

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
func (it *GatewayFactoryGatewayDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayFactoryGatewayDeployed)
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
		it.Event = new(GatewayFactoryGatewayDeployed)
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
func (it *GatewayFactoryGatewayDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayFactoryGatewayDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayFactoryGatewayDeployed represents a GatewayDeployed event raised by the GatewayFactory contract.
type GatewayFactoryGatewayDeployed struct {
	GatewayAddr      common.Address
	GatewayOwnerAddr common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGatewayDeployed is a free log retrieval operation binding the contract event 0x4b209f3d620ba8f4e206ec57dd17cebe32624f42eb77c06bcfd195be78b6c2c2.
//
// Solidity: event GatewayDeployed(address gatewayAddr, address gatewayOwnerAddr)
func (_GatewayFactory *GatewayFactoryFilterer) FilterGatewayDeployed(opts *bind.FilterOpts) (*GatewayFactoryGatewayDeployedIterator, error) {

	logs, sub, err := _GatewayFactory.contract.FilterLogs(opts, "GatewayDeployed")
	if err != nil {
		return nil, err
	}
	return &GatewayFactoryGatewayDeployedIterator{contract: _GatewayFactory.contract, event: "GatewayDeployed", logs: logs, sub: sub}, nil
}

// WatchGatewayDeployed is a free log subscription operation binding the contract event 0x4b209f3d620ba8f4e206ec57dd17cebe32624f42eb77c06bcfd195be78b6c2c2.
//
// Solidity: event GatewayDeployed(address gatewayAddr, address gatewayOwnerAddr)
func (_GatewayFactory *GatewayFactoryFilterer) WatchGatewayDeployed(opts *bind.WatchOpts, sink chan<- *GatewayFactoryGatewayDeployed) (event.Subscription, error) {

	logs, sub, err := _GatewayFactory.contract.WatchLogs(opts, "GatewayDeployed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayFactoryGatewayDeployed)
				if err := _GatewayFactory.contract.UnpackLog(event, "GatewayDeployed", log); err != nil {
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

// ParseGatewayDeployed is a log parse operation binding the contract event 0x4b209f3d620ba8f4e206ec57dd17cebe32624f42eb77c06bcfd195be78b6c2c2.
//
// Solidity: event GatewayDeployed(address gatewayAddr, address gatewayOwnerAddr)
func (_GatewayFactory *GatewayFactoryFilterer) ParseGatewayDeployed(log types.Log) (*GatewayFactoryGatewayDeployed, error) {
	event := new(GatewayFactoryGatewayDeployed)
	if err := _GatewayFactory.contract.UnpackLog(event, "GatewayDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayFactory contract.
type GatewayFactoryInitializedIterator struct {
	Event *GatewayFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayFactoryInitialized)
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
		it.Event = new(GatewayFactoryInitialized)
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
func (it *GatewayFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayFactoryInitialized represents a Initialized event raised by the GatewayFactory contract.
type GatewayFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayFactory *GatewayFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayFactoryInitializedIterator, error) {

	logs, sub, err := _GatewayFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayFactoryInitializedIterator{contract: _GatewayFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayFactory *GatewayFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayFactoryInitialized)
				if err := _GatewayFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayFactory *GatewayFactoryFilterer) ParseInitialized(log types.Log) (*GatewayFactoryInitialized, error) {
	event := new(GatewayFactoryInitialized)
	if err := _GatewayFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayFactory contract.
type GatewayFactoryOwnershipTransferredIterator struct {
	Event *GatewayFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayFactoryOwnershipTransferred)
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
		it.Event = new(GatewayFactoryOwnershipTransferred)
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
func (it *GatewayFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayFactory contract.
type GatewayFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayFactory *GatewayFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayFactoryOwnershipTransferredIterator{contract: _GatewayFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayFactory *GatewayFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayFactoryOwnershipTransferred)
				if err := _GatewayFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GatewayFactory *GatewayFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayFactoryOwnershipTransferred, error) {
	event := new(GatewayFactoryOwnershipTransferred)
	if err := _GatewayFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
