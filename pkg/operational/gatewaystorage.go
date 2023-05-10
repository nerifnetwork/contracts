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

// GatewayStorageGateway is an auto generated low-level Go binding around an user-defined struct.
type GatewayStorageGateway struct {
	Gateway common.Address
	Owner   common.Address
}

// GatewayStorageMetaData contains all meta data concerning the GatewayStorage contract.
var GatewayStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"clear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"contains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getGateway\",\"outputs\":[{\"internalType\":\"contractIGateway\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGateways\",\"outputs\":[{\"components\":[{\"internalType\":\"contractIGateway\",\"name\":\"gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structGatewayStorage.Gateway[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIGateway\",\"name\":\"gateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structGatewayStorage.Gateway[]\",\"name\":\"_gateways\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mustRemove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"mustSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"size\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6120158061010d6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063949d225d1161008c578063c32d805a11610066578063c32d805a146101ee578063d82778ce1461020a578063f2fde38b14610228578063f7beb8c714610244576100cf565b8063949d225d14610184578063a91ee0dc146101a2578063bda009fe146101be576100cf565b806352efea6e146100d45780635dbe47e8146100f2578063715018a61461012257806378e706961461012c5780637b103999146101485780638da5cb5b14610166575b600080fd5b6100dc610260565b6040516100e99190611647565b60405180910390f35b61010c600480360381019061010791906116d4565b61039f565b6040516101199190611647565b60405180910390f35b61012a6103ea565b005b610146600480360381019061014191906118ed565b610472565b005b61015061069e565b60405161015d9190611958565b60405180910390f35b61016e6106c4565b60405161017b9190611958565b60405180910390f35b61018c6106ed565b604051610199919061198c565b60405180910390f35b6101bc60048036038101906101b791906116d4565b6106fa565b005b6101d860048036038101906101d391906116d4565b6107ba565b6040516101e59190611a06565b60405180910390f35b610208600480360381019061020391906116d4565b61086b565b005b610212610946565b60405161021f9190611b1d565b60405180910390f35b610242600480360381019061023d91906116d4565b610a51565b005b61025e60048036038101906102599190611b3f565b610b48565b005b600061026a610be6565b73ffffffffffffffffffffffffffffffffffffffff166102886106c4565b73ffffffffffffffffffffffffffffffffffffffff16146102de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d590611bdc565b60405180910390fd5b60005b60038054905081101561038957600260006003838154811061030657610305611bfc565b5b906000526020600020906002020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009055808061038190611c5a565b9150506102e1565b506003600061039891906115a3565b6001905090565b600080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054119050919050565b6103f2610be6565b73ffffffffffffffffffffffffffffffffffffffff166104106106c4565b73ffffffffffffffffffffffffffffffffffffffff1614610466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045d90611bdc565b60405180910390fd5b6104706000610bee565b565b600061047e6001610cb2565b905080156104a2576001600060156101000a81548160ff0219169083151502179055505b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005b825181101561063f57600383828151811061050457610503611bfc565b5b6020026020010151908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600380549050600260008584815181106105e2576105e1611bfc565b5b60200260200101516020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808061063790611c5a565b9150506104e6565b5080156106995760008060156101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516106909190611cea565b60405180910390a15b505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600380549050905090565b610702610be6565b73ffffffffffffffffffffffffffffffffffffffff166107206106c4565b73ffffffffffffffffffffffffffffffffffffffff1614610776576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076d90611bdc565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006107c58261039f565b6107d25760009050610866565b60036001600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546108209190611d05565b8154811061083157610830611bfc565b5b906000526020600020906002020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505b919050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f290611dab565b60405180910390fd5b61090481610da5565b610943576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093a90611e3d565b60405180910390fd5b50565b60606003805480602002602001604051908101604052809291908181526020016000905b82821015610a4857838290600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250508152602001906001019061096a565b50505050905090565b610a59610be6565b73ffffffffffffffffffffffffffffffffffffffff16610a776106c4565b73ffffffffffffffffffffffffffffffffffffffff1614610acd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac490611bdc565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610b3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3390611ecf565b60405180910390fd5b610b4581610bee565b50565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610bd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bcf90611dab565b60405180910390fd5b610be282826111b7565b5050565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008060159054906101000a900460ff1615610d295760018260ff16148015610ce15750610cdf306114ec565b155b610d20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1790611f61565b60405180910390fd5b60009050610da0565b8160ff16600060149054906101000a900460ff1660ff1610610d80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7790611f61565b60405180910390fd5b81600060146101000a81548160ff021916908360ff160217905550600190505b919050565b60003373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610e37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2e90611dab565b60405180910390fd5b610e408261039f565b610e4d57600090506111b2565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060006001600380549050610ea59190611d05565b9050600060038281548110610ebd57610ebc611bfc565b5b90600052602060002090600202016040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050600183610f919190611d05565b82146110a0578260026000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550806003600185610fef9190611d05565b8154811061100057610fff611bfc565b5b906000526020600020906002020160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505b60038054806110b2576110b1611f81565b5b6001900381819060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550509055600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090556111658561150f565b8473ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff16146111aa576111a9816020015161150f565b5b600193505050505b919050565b3373ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611247576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123e90611dab565b60405180910390fd5b6112508261039f565b156113995760405180604001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681525060036001600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546112e49190611d05565b815481106112f5576112f4611bfc565b5b906000526020600020906002020160008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050506114df565b600360405180604001604052808373ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff16815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600380549050600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b6114e88261150f565b5050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060038054905081111561156957611568611fb0565b5b6115728261039f565b1561158d576000811161158857611587611fb0565b5b61159f565b6000811461159e5761159d611fb0565b5b5b5050565b50805460008255600202906000526020600020908101906115c491906115c7565b50565b5b8082111561162857600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506002016115c8565b5090565b60008115159050919050565b6116418161162c565b82525050565b600060208201905061165c6000830184611638565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006116a182611676565b9050919050565b6116b181611696565b81146116bc57600080fd5b50565b6000813590506116ce816116a8565b92915050565b6000602082840312156116ea576116e961166c565b5b60006116f8848285016116bf565b91505092915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61174f82611706565b810181811067ffffffffffffffff8211171561176e5761176d611717565b5b80604052505050565b6000611781611662565b905061178d8282611746565b919050565b600067ffffffffffffffff8211156117ad576117ac611717565b5b602082029050602081019050919050565b600080fd5b600080fd5b60006117d382611696565b9050919050565b6117e3816117c8565b81146117ee57600080fd5b50565b600081359050611800816117da565b92915050565b60006040828403121561181c5761181b6117c3565b5b6118266040611777565b90506000611836848285016117f1565b600083015250602061184a848285016116bf565b60208301525092915050565b600061186961186484611792565b611777565b9050808382526020820190506040840283018581111561188c5761188b6117be565b5b835b818110156118b557806118a18882611806565b84526020840193505060408101905061188e565b5050509392505050565b600082601f8301126118d4576118d3611701565b5b81356118e4848260208601611856565b91505092915050565b600080604083850312156119045761190361166c565b5b6000611912858286016116bf565b925050602083013567ffffffffffffffff81111561193357611932611671565b5b61193f858286016118bf565b9150509250929050565b61195281611696565b82525050565b600060208201905061196d6000830184611949565b92915050565b6000819050919050565b61198681611973565b82525050565b60006020820190506119a1600083018461197d565b92915050565b6000819050919050565b60006119cc6119c76119c284611676565b6119a7565b611676565b9050919050565b60006119de826119b1565b9050919050565b60006119f0826119d3565b9050919050565b611a00816119e5565b82525050565b6000602082019050611a1b60008301846119f7565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611a56816119e5565b82525050565b611a6581611696565b82525050565b604082016000820151611a816000850182611a4d565b506020820151611a946020850182611a5c565b50505050565b6000611aa68383611a6b565b60408301905092915050565b6000602082019050919050565b6000611aca82611a21565b611ad48185611a2c565b9350611adf83611a3d565b8060005b83811015611b10578151611af78882611a9a565b9750611b0283611ab2565b925050600181019050611ae3565b5085935050505092915050565b60006020820190508181036000830152611b378184611abf565b905092915050565b60008060408385031215611b5657611b5561166c565b5b6000611b64858286016116bf565b9250506020611b75858286016116bf565b9150509250929050565b600082825260208201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611bc6602083611b7f565b9150611bd182611b90565b602082019050919050565b60006020820190508181036000830152611bf581611bb9565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611c6582611973565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c9757611c96611c2b565b5b600182019050919050565b6000819050919050565b600060ff82169050919050565b6000611cd4611ccf611cca84611ca2565b6119a7565b611cac565b9050919050565b611ce481611cb9565b82525050565b6000602082019050611cff6000830184611cdb565b92915050565b6000611d1082611973565b9150611d1b83611973565b9250828203905081811115611d3357611d32611c2b565b5b92915050565b7f4761746577617953746f726167653a206f7065726174696f6e206973206e6f7460008201527f207065726d697474656400000000000000000000000000000000000000000000602082015250565b6000611d95602a83611b7f565b9150611da082611d39565b604082019050919050565b60006020820190508181036000830152611dc481611d88565b9050919050565b7f4761746577617953746f726167653a206661696c656420746f2072656d6f766560008201527f2067617465776179000000000000000000000000000000000000000000000000602082015250565b6000611e27602883611b7f565b9150611e3282611dcb565b604082019050919050565b60006020820190508181036000830152611e5681611e1a565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611eb9602683611b7f565b9150611ec482611e5d565b604082019050919050565b60006020820190508181036000830152611ee881611eac565b9050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000611f4b602e83611b7f565b9150611f5682611eef565b604082019050919050565b60006020820190508181036000830152611f7a81611f3e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea26469706673582212209e715101f99cb654397b71215d6ba2243946402ca419b1b5e50a0f71ab1b13aa64736f6c63430008120033",
}

// GatewayStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use GatewayStorageMetaData.ABI instead.
var GatewayStorageABI = GatewayStorageMetaData.ABI

// GatewayStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GatewayStorageMetaData.Bin instead.
var GatewayStorageBin = GatewayStorageMetaData.Bin

// DeployGatewayStorage deploys a new Ethereum contract, binding an instance of GatewayStorage to it.
func DeployGatewayStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GatewayStorage, error) {
	parsed, err := GatewayStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GatewayStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GatewayStorage{GatewayStorageCaller: GatewayStorageCaller{contract: contract}, GatewayStorageTransactor: GatewayStorageTransactor{contract: contract}, GatewayStorageFilterer: GatewayStorageFilterer{contract: contract}}, nil
}

// GatewayStorage is an auto generated Go binding around an Ethereum contract.
type GatewayStorage struct {
	GatewayStorageCaller     // Read-only binding to the contract
	GatewayStorageTransactor // Write-only binding to the contract
	GatewayStorageFilterer   // Log filterer for contract events
}

// GatewayStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewayStorageSession struct {
	Contract     *GatewayStorage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayStorageCallerSession struct {
	Contract *GatewayStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GatewayStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayStorageTransactorSession struct {
	Contract     *GatewayStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GatewayStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayStorageRaw struct {
	Contract *GatewayStorage // Generic contract binding to access the raw methods on
}

// GatewayStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayStorageCallerRaw struct {
	Contract *GatewayStorageCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayStorageTransactorRaw struct {
	Contract *GatewayStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGatewayStorage creates a new instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorage(address common.Address, backend bind.ContractBackend) (*GatewayStorage, error) {
	contract, err := bindGatewayStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GatewayStorage{GatewayStorageCaller: GatewayStorageCaller{contract: contract}, GatewayStorageTransactor: GatewayStorageTransactor{contract: contract}, GatewayStorageFilterer: GatewayStorageFilterer{contract: contract}}, nil
}

// NewGatewayStorageCaller creates a new read-only instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorageCaller(address common.Address, caller bind.ContractCaller) (*GatewayStorageCaller, error) {
	contract, err := bindGatewayStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageCaller{contract: contract}, nil
}

// NewGatewayStorageTransactor creates a new write-only instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayStorageTransactor, error) {
	contract, err := bindGatewayStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageTransactor{contract: contract}, nil
}

// NewGatewayStorageFilterer creates a new log filterer instance of GatewayStorage, bound to a specific deployed contract.
func NewGatewayStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayStorageFilterer, error) {
	contract, err := bindGatewayStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageFilterer{contract: contract}, nil
}

// bindGatewayStorage binds a generic wrapper to an already deployed contract.
func bindGatewayStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatewayStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayStorage *GatewayStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayStorage.Contract.GatewayStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayStorage *GatewayStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.Contract.GatewayStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayStorage *GatewayStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayStorage.Contract.GatewayStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GatewayStorage *GatewayStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GatewayStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GatewayStorage *GatewayStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GatewayStorage *GatewayStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GatewayStorage.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address owner) view returns(bool)
func (_GatewayStorage *GatewayStorageCaller) Contains(opts *bind.CallOpts, owner common.Address) (bool, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "contains", owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address owner) view returns(bool)
func (_GatewayStorage *GatewayStorageSession) Contains(owner common.Address) (bool, error) {
	return _GatewayStorage.Contract.Contains(&_GatewayStorage.CallOpts, owner)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address owner) view returns(bool)
func (_GatewayStorage *GatewayStorageCallerSession) Contains(owner common.Address) (bool, error) {
	return _GatewayStorage.Contract.Contains(&_GatewayStorage.CallOpts, owner)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address)
func (_GatewayStorage *GatewayStorageCaller) GetGateway(opts *bind.CallOpts, owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "getGateway", owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address)
func (_GatewayStorage *GatewayStorageSession) GetGateway(owner common.Address) (common.Address, error) {
	return _GatewayStorage.Contract.GetGateway(&_GatewayStorage.CallOpts, owner)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address owner) view returns(address)
func (_GatewayStorage *GatewayStorageCallerSession) GetGateway(owner common.Address) (common.Address, error) {
	return _GatewayStorage.Contract.GetGateway(&_GatewayStorage.CallOpts, owner)
}

// GetGateways is a free data retrieval call binding the contract method 0xd82778ce.
//
// Solidity: function getGateways() view returns((address,address)[])
func (_GatewayStorage *GatewayStorageCaller) GetGateways(opts *bind.CallOpts) ([]GatewayStorageGateway, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "getGateways")

	if err != nil {
		return *new([]GatewayStorageGateway), err
	}

	out0 := *abi.ConvertType(out[0], new([]GatewayStorageGateway)).(*[]GatewayStorageGateway)

	return out0, err

}

// GetGateways is a free data retrieval call binding the contract method 0xd82778ce.
//
// Solidity: function getGateways() view returns((address,address)[])
func (_GatewayStorage *GatewayStorageSession) GetGateways() ([]GatewayStorageGateway, error) {
	return _GatewayStorage.Contract.GetGateways(&_GatewayStorage.CallOpts)
}

// GetGateways is a free data retrieval call binding the contract method 0xd82778ce.
//
// Solidity: function getGateways() view returns((address,address)[])
func (_GatewayStorage *GatewayStorageCallerSession) GetGateways() ([]GatewayStorageGateway, error) {
	return _GatewayStorage.Contract.GetGateways(&_GatewayStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayStorage *GatewayStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayStorage *GatewayStorageSession) Owner() (common.Address, error) {
	return _GatewayStorage.Contract.Owner(&_GatewayStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GatewayStorage *GatewayStorageCallerSession) Owner() (common.Address, error) {
	return _GatewayStorage.Contract.Owner(&_GatewayStorage.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayStorage *GatewayStorageCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayStorage *GatewayStorageSession) Registry() (common.Address, error) {
	return _GatewayStorage.Contract.Registry(&_GatewayStorage.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_GatewayStorage *GatewayStorageCallerSession) Registry() (common.Address, error) {
	return _GatewayStorage.Contract.Registry(&_GatewayStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_GatewayStorage *GatewayStorageCaller) Size(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GatewayStorage.contract.Call(opts, &out, "size")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_GatewayStorage *GatewayStorageSession) Size() (*big.Int, error) {
	return _GatewayStorage.Contract.Size(&_GatewayStorage.CallOpts)
}

// Size is a free data retrieval call binding the contract method 0x949d225d.
//
// Solidity: function size() view returns(uint256)
func (_GatewayStorage *GatewayStorageCallerSession) Size() (*big.Int, error) {
	return _GatewayStorage.Contract.Size(&_GatewayStorage.CallOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_GatewayStorage *GatewayStorageTransactor) Clear(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "clear")
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_GatewayStorage *GatewayStorageSession) Clear() (*types.Transaction, error) {
	return _GatewayStorage.Contract.Clear(&_GatewayStorage.TransactOpts)
}

// Clear is a paid mutator transaction binding the contract method 0x52efea6e.
//
// Solidity: function clear() returns(bool)
func (_GatewayStorage *GatewayStorageTransactorSession) Clear() (*types.Transaction, error) {
	return _GatewayStorage.Contract.Clear(&_GatewayStorage.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x78e70696.
//
// Solidity: function initialize(address _registry, (address,address)[] _gateways) returns()
func (_GatewayStorage *GatewayStorageTransactor) Initialize(opts *bind.TransactOpts, _registry common.Address, _gateways []GatewayStorageGateway) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "initialize", _registry, _gateways)
}

// Initialize is a paid mutator transaction binding the contract method 0x78e70696.
//
// Solidity: function initialize(address _registry, (address,address)[] _gateways) returns()
func (_GatewayStorage *GatewayStorageSession) Initialize(_registry common.Address, _gateways []GatewayStorageGateway) (*types.Transaction, error) {
	return _GatewayStorage.Contract.Initialize(&_GatewayStorage.TransactOpts, _registry, _gateways)
}

// Initialize is a paid mutator transaction binding the contract method 0x78e70696.
//
// Solidity: function initialize(address _registry, (address,address)[] _gateways) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) Initialize(_registry common.Address, _gateways []GatewayStorageGateway) (*types.Transaction, error) {
	return _GatewayStorage.Contract.Initialize(&_GatewayStorage.TransactOpts, _registry, _gateways)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address owner) returns()
func (_GatewayStorage *GatewayStorageTransactor) MustRemove(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "mustRemove", owner)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address owner) returns()
func (_GatewayStorage *GatewayStorageSession) MustRemove(owner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustRemove(&_GatewayStorage.TransactOpts, owner)
}

// MustRemove is a paid mutator transaction binding the contract method 0xc32d805a.
//
// Solidity: function mustRemove(address owner) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) MustRemove(owner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustRemove(&_GatewayStorage.TransactOpts, owner)
}

// MustSet is a paid mutator transaction binding the contract method 0xf7beb8c7.
//
// Solidity: function mustSet(address owner, address gateway) returns()
func (_GatewayStorage *GatewayStorageTransactor) MustSet(opts *bind.TransactOpts, owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "mustSet", owner, gateway)
}

// MustSet is a paid mutator transaction binding the contract method 0xf7beb8c7.
//
// Solidity: function mustSet(address owner, address gateway) returns()
func (_GatewayStorage *GatewayStorageSession) MustSet(owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustSet(&_GatewayStorage.TransactOpts, owner, gateway)
}

// MustSet is a paid mutator transaction binding the contract method 0xf7beb8c7.
//
// Solidity: function mustSet(address owner, address gateway) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) MustSet(owner common.Address, gateway common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.MustSet(&_GatewayStorage.TransactOpts, owner, gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayStorage *GatewayStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayStorage *GatewayStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayStorage.Contract.RenounceOwnership(&_GatewayStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GatewayStorage *GatewayStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GatewayStorage.Contract.RenounceOwnership(&_GatewayStorage.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_GatewayStorage *GatewayStorageTransactor) SetRegistry(opts *bind.TransactOpts, _registry common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "setRegistry", _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_GatewayStorage *GatewayStorageSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.SetRegistry(&_GatewayStorage.TransactOpts, _registry)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address _registry) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) SetRegistry(_registry common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.SetRegistry(&_GatewayStorage.TransactOpts, _registry)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayStorage *GatewayStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayStorage *GatewayStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.TransferOwnership(&_GatewayStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GatewayStorage *GatewayStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GatewayStorage.Contract.TransferOwnership(&_GatewayStorage.TransactOpts, newOwner)
}

// GatewayStorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GatewayStorage contract.
type GatewayStorageInitializedIterator struct {
	Event *GatewayStorageInitialized // Event containing the contract specifics and raw log

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
func (it *GatewayStorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStorageInitialized)
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
		it.Event = new(GatewayStorageInitialized)
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
func (it *GatewayStorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStorageInitialized represents a Initialized event raised by the GatewayStorage contract.
type GatewayStorageInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayStorage *GatewayStorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*GatewayStorageInitializedIterator, error) {

	logs, sub, err := _GatewayStorage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GatewayStorageInitializedIterator{contract: _GatewayStorage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GatewayStorage *GatewayStorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GatewayStorageInitialized) (event.Subscription, error) {

	logs, sub, err := _GatewayStorage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStorageInitialized)
				if err := _GatewayStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_GatewayStorage *GatewayStorageFilterer) ParseInitialized(log types.Log) (*GatewayStorageInitialized, error) {
	event := new(GatewayStorageInitialized)
	if err := _GatewayStorage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GatewayStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GatewayStorage contract.
type GatewayStorageOwnershipTransferredIterator struct {
	Event *GatewayStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GatewayStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayStorageOwnershipTransferred)
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
		it.Event = new(GatewayStorageOwnershipTransferred)
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
func (it *GatewayStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayStorageOwnershipTransferred represents a OwnershipTransferred event raised by the GatewayStorage contract.
type GatewayStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayStorage *GatewayStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GatewayStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GatewayStorageOwnershipTransferredIterator{contract: _GatewayStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GatewayStorage *GatewayStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GatewayStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GatewayStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayStorageOwnershipTransferred)
				if err := _GatewayStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GatewayStorage *GatewayStorageFilterer) ParseOwnershipTransferred(log types.Log) (*GatewayStorageOwnershipTransferred, error) {
	event := new(GatewayStorageOwnershipTransferred)
	if err := _GatewayStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
