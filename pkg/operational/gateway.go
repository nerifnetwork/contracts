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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"knownWorkflows\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"knownWorkflowOwners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"knownCustomerContracts\",\"type\":\"address[]\"}],\"internalType\":\"structGateway.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"knownWorkflows\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"knownWorkflowOwners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"knownCustomerContracts\",\"type\":\"address[]\"}],\"internalType\":\"structGateway.Config\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100ff565b600033905090565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b611e618061010e6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a91ee0dc11610066578063a91ee0dc146100fa578063c3f909d414610116578063c4d66de814610134578063cb18c01b14610150578063f2fde38b1461016c57610093565b8063342b5ece14610098578063715018a6146100b45780637b103999146100be5780638da5cb5b146100dc575b600080fd5b6100b260048036038101906100ad9190610d27565b610188565b005b6100bc610219565b005b6100c66102a1565b6040516100d39190610def565b60405180910390f35b6100e46102c7565b6040516100f19190610e2b565b60405180910390f35b610114600480360381019061010f9190610e72565b6102f0565b005b61011e6103b0565b60405161012b9190611083565b60405180910390f35b61014e60048036038101906101499190610e72565b61053e565b005b61016a60048036038101906101659190611136565b6105d4565b005b61018660048036038101906101819190610e72565b6109f7565b005b610190610aee565b73ffffffffffffffffffffffffffffffffffffffff166101ae6102c7565b73ffffffffffffffffffffffffffffffffffffffff1614610204576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fb90611207565b60405180910390fd5b806002818161021391906117b6565b90505050565b610221610aee565b73ffffffffffffffffffffffffffffffffffffffff1661023f6102c7565b73ffffffffffffffffffffffffffffffffffffffff1614610295576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028c90611207565b60405180910390fd5b61029f6000610af6565b565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102f8610aee565b73ffffffffffffffffffffffffffffffffffffffff166103166102c7565b73ffffffffffffffffffffffffffffffffffffffff161461036c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036390611207565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6103b8610cce565b60026040518060600160405290816000820180548060200260200160405190810160405280929190818152602001828054801561041457602002820191906000526020600020905b815481526020019060010190808311610400575b50505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156104a257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610458575b505050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561053057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116104e6575b505050505081525050905090565b600061054a6001610bbb565b9050801561056e576001600060016101000a81548160ff0219169083151502179055505b610577826102f0565b80156105d05760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516105c7919061180c565b60405180910390a15b5050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610664576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065b90611899565b60405180910390fd5b83836000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eec7b03c846040518263ffffffff1660e01b81526004016106c391906118c8565b600060405180830381865afa1580156106e0573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906107099190611b5d565b90506000816020015173ffffffffffffffffffffffffffffffffffffffff166107306102c7565b73ffffffffffffffffffffffffffffffffffffffff16149050806107f75760005b6002600101805490508110156107f557826020015173ffffffffffffffffffffffffffffffffffffffff166002600101828154811061079357610792611ba6565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036107e257600191506107f5565b80806107ed90611bd5565b915050610751565b505b806108555760005b60026000018054905081101561085357846002600001828154811061082757610826611ba6565b5b9060005260206000200154036108405760019150610853565b808061084b90611bd5565b9150506107ff565b505b806108fd5760005b60028001805490508110156108fb578373ffffffffffffffffffffffffffffffffffffffff1660028001828154811061089957610898611ba6565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036108e857600191506108fb565b80806108f390611bd5565b91505061085d565b505b8061093d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093490611899565b60405180910390fd5b60008773ffffffffffffffffffffffffffffffffffffffff168787604051610966929190611c5c565b6000604051808303816000865af19150503d80600081146109a3576040519150601f19603f3d011682016040523d82523d6000602084013e6109a8565b606091505b50509050806109ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109e390611ce7565b60405180910390fd5b505050505050505050565b6109ff610aee565b73ffffffffffffffffffffffffffffffffffffffff16610a1d6102c7565b73ffffffffffffffffffffffffffffffffffffffff1614610a73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6a90611207565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ae2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad990611d79565b60405180910390fd5b610aeb81610af6565b50565b600033905090565b60008060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008060019054906101000a900460ff1615610c325760018260ff16148015610bea5750610be830610cab565b155b610c29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2090611e0b565b60405180910390fd5b60009050610ca6565b8160ff1660008054906101000a900460ff1660ff1610610c87576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7e90611e0b565b60405180910390fd5b816000806101000a81548160ff021916908360ff160217905550600190505b919050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60405180606001604052806060815260200160608152602001606081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600060608284031215610d1e57610d1d610d03565b5b81905092915050565b600060208284031215610d3d57610d3c610cf9565b5b600082013567ffffffffffffffff811115610d5b57610d5a610cfe565b5b610d6784828501610d08565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610db5610db0610dab84610d70565b610d90565b610d70565b9050919050565b6000610dc782610d9a565b9050919050565b6000610dd982610dbc565b9050919050565b610de981610dce565b82525050565b6000602082019050610e046000830184610de0565b92915050565b6000610e1582610d70565b9050919050565b610e2581610e0a565b82525050565b6000602082019050610e406000830184610e1c565b92915050565b610e4f81610e0a565b8114610e5a57600080fd5b50565b600081359050610e6c81610e46565b92915050565b600060208284031215610e8857610e87610cf9565b5b6000610e9684828501610e5d565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b610ede81610ecb565b82525050565b6000610ef08383610ed5565b60208301905092915050565b6000602082019050919050565b6000610f1482610e9f565b610f1e8185610eaa565b9350610f2983610ebb565b8060005b83811015610f5a578151610f418882610ee4565b9750610f4c83610efc565b925050600181019050610f2d565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610f9c81610e0a565b82525050565b6000610fae8383610f93565b60208301905092915050565b6000602082019050919050565b6000610fd282610f67565b610fdc8185610f72565b9350610fe783610f83565b8060005b83811015611018578151610fff8882610fa2565b975061100a83610fba565b925050600181019050610feb565b5085935050505092915050565b600060608301600083015184820360008601526110428282610f09565b9150506020830151848203602086015261105c8282610fc7565b915050604083015184820360408601526110768282610fc7565b9150508091505092915050565b6000602082019050818103600083015261109d8184611025565b905092915050565b6110ae81610ecb565b81146110b957600080fd5b50565b6000813590506110cb816110a5565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126110f6576110f56110d1565b5b8235905067ffffffffffffffff811115611113576111126110d6565b5b60208301915083600182028301111561112f5761112e6110db565b5b9250929050565b600080600080606085870312156111505761114f610cf9565b5b600061115e878288016110bc565b945050602061116f87828801610e5d565b935050604085013567ffffffffffffffff8111156111905761118f610cfe565b5b61119c878288016110e0565b925092505092959194509250565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006111f16020836111aa565b91506111fc826111bb565b602082019050919050565b60006020820190508181036000830152611220816111e4565b9050919050565b600080fd5b600080fd5b600080fd5b6000808335600160200384360303811261125357611252611227565b5b80840192508235915067ffffffffffffffff8211156112755761127461122c565b5b60208301925060208202360383131561129157611290611231565b5b509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600081549050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008190506001806001038301049050919050565b60008190508160005260206000209050919050565b600082821b905092915050565b6000600883026113697fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261132c565b611373868361132c565b95508019841693508086168417925050509392505050565b60006113a66113a161139c84610ecb565b610d90565b610ecb565b9050919050565b6000819050919050565b6113c08361138b565b6113d46113cc826113ad565b848454611339565b825550505050565b600090565b6113e96113dc565b6113f48184846113b7565b505050565b5b818110156114185761140d6000826113e1565b6001810190506113fa565b5050565b818310156114555761142d82611302565b61143684611302565b61143f83611317565b81810183820161144f81836113f9565b50505050505b505050565b6801000000000000000082111561147457611473611299565b5b61147d816112c8565b82825561148b83828461141c565b505050565b600082905092915050565b600081356114a8816110a5565b80915050919050565b6000819050919050565b6114c58383611490565b67ffffffffffffffff8111156114de576114dd611299565b5b6114e8818361145a565b6114f1836114b1565b6114fa83611317565b6001830460005b818110156115395760006115148561149b565b61151d816113ad565b8092506020870196505050808285015550600181019050611501565b5050505050505050565b61154e8383836114bb565b505050565b600080833560016020038436030381126115705761156f611227565b5b80840192508235915067ffffffffffffffff8211156115925761159161122c565b5b6020830192506020820236038313156115ae576115ad611231565b5b509250929050565b600081549050919050565b60008190506001806001038301049050919050565b60008190508160005260206000209050919050565b5b8181101561160a576115ff6000826113e1565b6001810190506115ec565b5050565b818310156116475761161f826115c1565b611628846115c1565b611631836115d6565b81810183820161164181836115eb565b50505050505b505050565b6801000000000000000082111561166657611665611299565b5b61166f816115b6565b82825561167d83828461160e565b505050565b600082905092915050565b6000813561169a81610e46565b80915050919050565b6000819050919050565b6000819050919050565b6116c18383611682565b67ffffffffffffffff8111156116da576116d9611299565b5b6116e4818361164c565b6116ed836116a3565b6116f6836115d6565b6001830460005b818110156117355760006117108561168d565b611719816116ad565b80925060208701965050508082850155506001810190506116fd565b5050505050505050565b61174a8383836116b7565b505050565b60008101600083016117618185611236565b61176c818386611543565b5050505060018101602083016117828185611553565b61178d81838661173f565b5050505060028101604083016117a38185611553565b6117ae81838661173f565b505050505050565b6117c0828261174f565b5050565b6000819050919050565b600060ff82169050919050565b60006117f66117f16117ec846117c4565b610d90565b6117ce565b9050919050565b611806816117db565b82525050565b600060208201905061182160008301846117fd565b92915050565b7f476174657761793a206f7065726174696f6e206973206e6f74207065726d697460008201527f7465640000000000000000000000000000000000000000000000000000000000602082015250565b60006118836023836111aa565b915061188e82611827565b604082019050919050565b600060208201905081810360008301526118b281611876565b9050919050565b6118c281610ecb565b82525050565b60006020820190506118dd60008301846118b9565b92915050565b600080fd5b6000601f19601f8301169050919050565b611902826118e8565b810181811067ffffffffffffffff8211171561192157611920611299565b5b80604052505050565b6000611934610cef565b905061194082826118f9565b919050565b600080fd5b600081519050611959816110a5565b92915050565b60008151905061196e81610e46565b92915050565b600080fd5b600067ffffffffffffffff82111561199457611993611299565b5b61199d826118e8565b9050602081019050919050565b60005b838110156119c85780820151818401526020810190506119ad565b60008484015250505050565b60006119e76119e284611979565b61192a565b905082815260208101848484011115611a0357611a02611974565b5b611a0e8482856119aa565b509392505050565b600082601f830112611a2b57611a2a6110d1565b5b8151611a3b8482602086016119d4565b91505092915050565b60048110611a5157600080fd5b50565b600081519050611a6381611a44565b92915050565b60008115159050919050565b611a7e81611a69565b8114611a8957600080fd5b50565b600081519050611a9b81611a75565b92915050565b600060c08284031215611ab757611ab66118e3565b5b611ac160c061192a565b90506000611ad18482850161194a565b6000830152506020611ae58482850161195f565b602083015250604082015167ffffffffffffffff811115611b0957611b08611945565b5b611b1584828501611a16565b6040830152506060611b2984828501611a54565b6060830152506080611b3d84828501611a8c565b60808301525060a0611b518482850161194a565b60a08301525092915050565b600060208284031215611b7357611b72610cf9565b5b600082015167ffffffffffffffff811115611b9157611b90610cfe565b5b611b9d84828501611aa1565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000611be082610ecb565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c1257611c116112d3565b5b600182019050919050565b600081905092915050565b82818337600083830152505050565b6000611c438385611c1d565b9350611c50838584611c28565b82840190509392505050565b6000611c69828486611c37565b91508190509392505050565b7f476174657761793a206661696c656420746f206578656375746520637573746f60008201527f6d657220636f6e74726163740000000000000000000000000000000000000000602082015250565b6000611cd1602c836111aa565b9150611cdc82611c75565b604082019050919050565b60006020820190508181036000830152611d0081611cc4565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611d636026836111aa565b9150611d6e82611d07565b604082019050919050565b60006020820190508181036000830152611d9281611d56565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000611df5602e836111aa565b9150611e0082611d99565b604082019050919050565b60006020820190508181036000830152611e2481611de8565b905091905056fea26469706673582212203c7a47d5e8c8dde43cf4748487bbdee25ee8fdeffa86218939fd0b337bd8b2ed64736f6c63430008120033",
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

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Gateway *GatewayTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "initialize", _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Gateway *GatewaySession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _registry)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _registry) returns()
func (_Gateway *GatewayTransactorSession) Initialize(_registry common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Initialize(&_Gateway.TransactOpts, _registry)
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
