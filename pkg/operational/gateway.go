// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operational

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

// GatewayConfig is an auto generated low-level Go binding around an user-defined struct.
type GatewayConfig struct {
	KnownWorkflows         []*big.Int
	KnownWorkflowOwners    []common.Address
	KnownCustomerContracts []common.Address
}

// GatewayMetaData contains all meta data concerning the Gateway contract.
var GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"knownWorkflows\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"knownWorkflowOwners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"knownCustomerContracts\",\"type\":\"address[]\"}],\"internalType\":\"structGateway.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"knownWorkflows\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"knownWorkflowOwners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"knownCustomerContracts\",\"type\":\"address[]\"}],\"internalType\":\"structGateway.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001eb538038062001eb58339818101604052810190620000379190620002a1565b620000576200004b6200006f60201b60201c565b6200007760201b60201c565b62000068816200013b60201b60201c565b5062000356565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6200014b6200006f60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16620001716200020e60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1614620001ca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001c19062000334565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000269826200023c565b9050919050565b6200027b816200025c565b81146200028757600080fd5b50565b6000815190506200029b8162000270565b92915050565b600060208284031215620002ba57620002b962000237565b5b6000620002ca848285016200028a565b91505092915050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006200031c602083620002d3565b91506200032982620002e4565b602082019050919050565b600060208201905081810360008301526200034f816200030d565b9050919050565b611b4f80620003666000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a91ee0dc1161005b578063a91ee0dc146100ef578063c3f909d41461010b578063cb18c01b14610129578063f2fde38b1461014557610088565b8063342b5ece1461008d578063715018a6146100a95780637b103999146100b35780638da5cb5b146100d1575b600080fd5b6100a760048036038101906100a29190610b56565b610161565b005b6100b16101f2565b005b6100bb61027a565b6040516100c89190610c1e565b60405180910390f35b6100d96102a0565b6040516100e69190610c5a565b60405180910390f35b61010960048036038101906101049190610ca1565b6102c9565b005b610113610389565b6040516101209190610eb2565b60405180910390f35b610143600480360381019061013e9190610f65565b610517565b005b61015f600480360381019061015a9190610ca1565b61093a565b005b610169610a31565b73ffffffffffffffffffffffffffffffffffffffff166101876102a0565b73ffffffffffffffffffffffffffffffffffffffff16146101dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d490611036565b60405180910390fd5b80600281816101ec91906115e5565b90505050565b6101fa610a31565b73ffffffffffffffffffffffffffffffffffffffff166102186102a0565b73ffffffffffffffffffffffffffffffffffffffff161461026e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026590611036565b60405180910390fd5b6102786000610a39565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102d1610a31565b73ffffffffffffffffffffffffffffffffffffffff166102ef6102a0565b73ffffffffffffffffffffffffffffffffffffffff1614610345576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033c90611036565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610391610afd565b6002604051806060016040529081600082018054806020026020016040519081016040528092919081815260200182805480156103ed57602002820191906000526020600020905b8154815260200190600101908083116103d9575b505050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561047b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610431575b505050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561050957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116104bf575b505050505081525050905090565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146105a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059e90611665565b60405180910390fd5b83836000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eec7b03c846040518263ffffffff1660e01b81526004016106069190611694565b600060405180830381865afa158015610623573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061064c91906118dd565b90506000816020015173ffffffffffffffffffffffffffffffffffffffff166106736102a0565b73ffffffffffffffffffffffffffffffffffffffff161490508061073a5760005b60026001018054905081101561073857826020015173ffffffffffffffffffffffffffffffffffffffff16600260010182815481106106d6576106d5611926565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036107255760019150610738565b808061073090611955565b915050610694565b505b806107985760005b60026000018054905081101561079657846002600001828154811061076a57610769611926565b5b9060005260206000200154036107835760019150610796565b808061078e90611955565b915050610742565b505b806108405760005b600280018054905081101561083e578373ffffffffffffffffffffffffffffffffffffffff166002800182815481106107dc576107db611926565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361082b576001915061083e565b808061083690611955565b9150506107a0565b505b80610880576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087790611665565b60405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff1687876040516108a99291906119dc565b6000604051808303816000865af19150503d80600081146108e6576040519150601f19603f3d011682016040523d82523d6000602084013e6108eb565b606091505b505090508061092f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092690611a67565b60405180910390fd5b505050505050505050565b610942610a31565b73ffffffffffffffffffffffffffffffffffffffff166109606102a0565b73ffffffffffffffffffffffffffffffffffffffff16146109b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ad90611036565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610a25576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1c90611af9565b60405180910390fd5b610a2e81610a39565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60405180606001604052806060815260200160608152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600060608284031215610b4d57610b4c610b32565b5b81905092915050565b600060208284031215610b6c57610b6b610b28565b5b600082013567ffffffffffffffff811115610b8a57610b89610b2d565b5b610b9684828501610b37565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610be4610bdf610bda84610b9f565b610bbf565b610b9f565b9050919050565b6000610bf682610bc9565b9050919050565b6000610c0882610beb565b9050919050565b610c1881610bfd565b82525050565b6000602082019050610c336000830184610c0f565b92915050565b6000610c4482610b9f565b9050919050565b610c5481610c39565b82525050565b6000602082019050610c6f6000830184610c4b565b92915050565b610c7e81610c39565b8114610c8957600080fd5b50565b600081359050610c9b81610c75565b92915050565b600060208284031215610cb757610cb6610b28565b5b6000610cc584828501610c8c565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b610d0d81610cfa565b82525050565b6000610d1f8383610d04565b60208301905092915050565b6000602082019050919050565b6000610d4382610cce565b610d4d8185610cd9565b9350610d5883610cea565b8060005b83811015610d89578151610d708882610d13565b9750610d7b83610d2b565b925050600181019050610d5c565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610dcb81610c39565b82525050565b6000610ddd8383610dc2565b60208301905092915050565b6000602082019050919050565b6000610e0182610d96565b610e0b8185610da1565b9350610e1683610db2565b8060005b83811015610e47578151610e2e8882610dd1565b9750610e3983610de9565b925050600181019050610e1a565b5085935050505092915050565b60006060830160008301518482036000860152610e718282610d38565b91505060208301518482036020860152610e8b8282610df6565b91505060408301518482036040860152610ea58282610df6565b9150508091505092915050565b60006020820190508181036000830152610ecc8184610e54565b905092915050565b610edd81610cfa565b8114610ee857600080fd5b50565b600081359050610efa81610ed4565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610f2557610f24610f00565b5b8235905067ffffffffffffffff811115610f4257610f41610f05565b5b602083019150836001820283011115610f5e57610f5d610f0a565b5b9250929050565b60008060008060608587031215610f7f57610f7e610b28565b5b6000610f8d87828801610eeb565b9450506020610f9e87828801610c8c565b935050604085013567ffffffffffffffff811115610fbf57610fbe610b2d565b5b610fcb87828801610f0f565b925092505092959194509250565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611020602083610fd9565b915061102b82610fea565b602082019050919050565b6000602082019050818103600083015261104f81611013565b9050919050565b600080fd5b600080fd5b600080fd5b6000808335600160200384360303811261108257611081611056565b5b80840192508235915067ffffffffffffffff8211156110a4576110a361105b565b5b6020830192506020820236038313156110c0576110bf611060565b5b509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081549050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008190506001806001038301049050919050565b60008190508160005260206000209050919050565b600082821b905092915050565b6000600883026111987fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261115b565b6111a2868361115b565b95508019841693508086168417925050509392505050565b60006111d56111d06111cb84610cfa565b610bbf565b610cfa565b9050919050565b6000819050919050565b6111ef836111ba565b6112036111fb826111dc565b848454611168565b825550505050565b600090565b61121861120b565b6112238184846111e6565b505050565b5b818110156112475761123c600082611210565b600181019050611229565b5050565b818310156112845761125c82611131565b61126584611131565b61126e83611146565b81810183820161127e8183611228565b50505050505b505050565b680100000000000000008211156112a3576112a26110c8565b5b6112ac816110f7565b8282556112ba83828461124b565b505050565b600082905092915050565b600081356112d781610ed4565b80915050919050565b6000819050919050565b6112f483836112bf565b67ffffffffffffffff81111561130d5761130c6110c8565b5b6113178183611289565b611320836112e0565b61132983611146565b6001830460005b81811015611368576000611343856112ca565b61134c816111dc565b8092506020870196505050808285015550600181019050611330565b5050505050505050565b61137d8383836112ea565b505050565b6000808335600160200384360303811261139f5761139e611056565b5b80840192508235915067ffffffffffffffff8211156113c1576113c061105b565b5b6020830192506020820236038313156113dd576113dc611060565b5b509250929050565b600081549050919050565b60008190506001806001038301049050919050565b60008190508160005260206000209050919050565b5b818110156114395761142e600082611210565b60018101905061141b565b5050565b818310156114765761144e826113f0565b611457846113f0565b61146083611405565b818101838201611470818361141a565b50505050505b505050565b68010000000000000000821115611495576114946110c8565b5b61149e816113e5565b8282556114ac83828461143d565b505050565b600082905092915050565b600081356114c981610c75565b80915050919050565b6000819050919050565b6000819050919050565b6114f083836114b1565b67ffffffffffffffff811115611509576115086110c8565b5b611513818361147b565b61151c836114d2565b61152583611405565b6001830460005b8181101561156457600061153f856114bc565b611548816114dc565b809250602087019650505080828501555060018101905061152c565b5050505050505050565b6115798383836114e6565b505050565b60008101600083016115908185611065565b61159b818386611372565b5050505060018101602083016115b18185611382565b6115bc81838661156e565b5050505060028101604083016115d28185611382565b6115dd81838661156e565b505050505050565b6115ef828261157e565b5050565b7f476174657761793a206f7065726174696f6e206973206e6f74207065726d697460008201527f7465640000000000000000000000000000000000000000000000000000000000602082015250565b600061164f602383610fd9565b915061165a826115f3565b604082019050919050565b6000602082019050818103600083015261167e81611642565b9050919050565b61168e81610cfa565b82525050565b60006020820190506116a96000830184611685565b92915050565b600080fd5b6000601f19601f8301169050919050565b6116ce826116b4565b810181811067ffffffffffffffff821117156116ed576116ec6110c8565b5b80604052505050565b6000611700610b1e565b905061170c82826116c5565b919050565b600080fd5b60008151905061172581610ed4565b92915050565b60008151905061173a81610c75565b92915050565b600080fd5b600067ffffffffffffffff8211156117605761175f6110c8565b5b611769826116b4565b9050602081019050919050565b60005b83811015611794578082015181840152602081019050611779565b60008484015250505050565b60006117b36117ae84611745565b6116f6565b9050828152602081018484840111156117cf576117ce611740565b5b6117da848285611776565b509392505050565b600082601f8301126117f7576117f6610f00565b5b81516118078482602086016117a0565b91505092915050565b6004811061181d57600080fd5b50565b60008151905061182f81611810565b92915050565b600060a0828403121561184b5761184a6116af565b5b61185560a06116f6565b9050600061186584828501611716565b60008301525060206118798482850161172b565b602083015250604082015167ffffffffffffffff81111561189d5761189c611711565b5b6118a9848285016117e2565b60408301525060606118bd84828501611820565b60608301525060806118d184828501611716565b60808301525092915050565b6000602082840312156118f3576118f2610b28565b5b600082015167ffffffffffffffff81111561191157611910610b2d565b5b61191d84828501611835565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061196082610cfa565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361199257611991611102565b5b600182019050919050565b600081905092915050565b82818337600083830152505050565b60006119c3838561199d565b93506119d08385846119a8565b82840190509392505050565b60006119e98284866119b7565b91508190509392505050565b7f476174657761793a206661696c656420746f206578656375746520637573746f60008201527f6d657220636f6e74726163740000000000000000000000000000000000000000602082015250565b6000611a51602c83610fd9565b9150611a5c826119f5565b604082019050919050565b60006020820190508181036000830152611a8081611a44565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611ae3602683610fd9565b9150611aee82611a87565b604082019050919050565b60006020820190508181036000830152611b1281611ad6565b905091905056fea26469706673582212202d9e997e71b46ff000174d15c4802c4bd583cde05fe798e06314db1e2ac64f4e64736f6c63430008120033",
}

// GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayMetaData.ABI instead.
var GatewayABI = GatewayMetaData.ABI

// GatewayBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayMetaData.Bin instead.
var GatewayBin = GatewayMetaData.Bin

// DeployGateway deploys a new Ethereum contract, binding an instance of Gateway to it.
func DeployGateway(auth *bind.TransactOpts, backend bind.ContractBackend, _registry common.Address) (common.Address, *types.Transaction, *Gateway, error) {
	parsed, err := GatewayMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayBin), backend, _registry)
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
	parsed, err := abi.JSON(strings.NewReader(GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256[],address[],address[]))
func (_Gateway *GatewayCaller) GetConfig(opts *bind.CallOpts) (GatewayConfig, error) {
	var out []interface{}
	err := _Gateway.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(GatewayConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(GatewayConfig)).(*GatewayConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256[],address[],address[]))
func (_Gateway *GatewaySession) GetConfig() (GatewayConfig, error) {
	return _Gateway.Contract.GetConfig(&_Gateway.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint256[],address[],address[]))
func (_Gateway *GatewayCallerSession) GetConfig() (GatewayConfig, error) {
	return _Gateway.Contract.GetConfig(&_Gateway.CallOpts)
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

// Perform is a paid mutator transaction binding the contract method 0xcb18c01b.
//
// Solidity: function perform(uint256 id, address target, bytes payload) returns()
func (_Gateway *GatewayTransactor) Perform(opts *bind.TransactOpts, id *big.Int, target common.Address, payload []byte) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "perform", id, target, payload)
}

// Perform is a paid mutator transaction binding the contract method 0xcb18c01b.
//
// Solidity: function perform(uint256 id, address target, bytes payload) returns()
func (_Gateway *GatewaySession) Perform(id *big.Int, target common.Address, payload []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Perform(&_Gateway.TransactOpts, id, target, payload)
}

// Perform is a paid mutator transaction binding the contract method 0xcb18c01b.
//
// Solidity: function perform(uint256 id, address target, bytes payload) returns()
func (_Gateway *GatewayTransactorSession) Perform(id *big.Int, target common.Address, payload []byte) (*types.Transaction, error) {
	return _Gateway.Contract.Perform(&_Gateway.TransactOpts, id, target, payload)
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

// SetConfig is a paid mutator transaction binding the contract method 0x342b5ece.
//
// Solidity: function setConfig((uint256[],address[],address[]) _config) returns()
func (_Gateway *GatewayTransactor) SetConfig(opts *bind.TransactOpts, _config GatewayConfig) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setConfig", _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x342b5ece.
//
// Solidity: function setConfig((uint256[],address[],address[]) _config) returns()
func (_Gateway *GatewaySession) SetConfig(_config GatewayConfig) (*types.Transaction, error) {
	return _Gateway.Contract.SetConfig(&_Gateway.TransactOpts, _config)
}

// SetConfig is a paid mutator transaction binding the contract method 0x342b5ece.
//
// Solidity: function setConfig((uint256[],address[],address[]) _config) returns()
func (_Gateway *GatewayTransactorSession) SetConfig(_config GatewayConfig) (*types.Transaction, error) {
	return _Gateway.Contract.SetConfig(&_Gateway.TransactOpts, _config)
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
