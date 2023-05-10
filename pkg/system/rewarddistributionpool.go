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

// RewardDistributionPoolMetaData contains all meta data concerning the RewardDistributionPool contract.
var RewardDistributionPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CollectRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reinvestRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardPositions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsOwing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b506118f9806100206000396000f3fe6080604052600436106101025760003560e01c80637cf8b8be11610095578063b7b08d4e11610064578063b7b08d4e146102c0578063c2e8caf6146102fe578063ce119a8a14610329578063faaa8a6414610354578063fd4c4aee1461037f57610109565b80637cf8b8be146102285780638f76c0d814610253578063abf410e51461027e578063b5f4d38c146102a957610109565b80635c211f88116100d15780635c211f88146101a45780636f4a2cd0146101cf57806370bb45b3146101e657806378a5c206146101fd57610109565b8063372500ab1461010e5780633a9783f314610125578063485cc95514610150578063561ff9a91461017957610109565b3661010957005b600080fd5b34801561011a57600080fd5b506101236103aa565b005b34801561013157600080fd5b5061013a61047e565b6040516101479190610f8f565b60405180910390f35b34801561015c57600080fd5b5061017760048036038101906101729190611023565b6104b7565b005b34801561018557600080fd5b5061018e61058f565b60405161019b9190610f8f565b60405180910390f35b3480156101b057600080fd5b506101b96105c8565b6040516101c691906110c2565b60405180910390f35b3480156101db57600080fd5b506101e46105ee565b005b3480156101f257600080fd5b506101fb61070f565b005b34801561020957600080fd5b50610212610722565b60405161021f9190610f8f565b60405180910390f35b34801561023457600080fd5b5061023d61075b565b60405161024a91906110f6565b60405180910390f35b34801561025f57600080fd5b50610268610761565b60405161027591906110f6565b60405180910390f35b34801561028a57600080fd5b50610293610767565b6040516102a09190611132565b60405180910390f35b3480156102b557600080fd5b506102be61078d565b005b3480156102cc57600080fd5b506102e760048036038101906102e2919061114d565b610863565b6040516102f592919061117a565b60405180910390f35b34801561030a57600080fd5b50610313610887565b60405161032091906110f6565b60405180910390f35b34801561033557600080fd5b5061033e61088d565b60405161034b9190610f8f565b60405180910390f35b34801561036057600080fd5b506103696108c6565b6040516103769190610f8f565b60405180910390f35b34801561038b57600080fd5b506103946108ff565b6040516103a191906110f6565b60405180910390f35b60006103b46108ff565b9050600081111561047b5780600260008282546103d191906111d2565b9250508190555080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600082825461042a9190611206565b92505081905550600354600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055505b50565b6040518060400160405280600381526020017f646b67000000000000000000000000000000000000000000000000000000000081525081565b60006104c360016109f9565b905080156104e7576001600060016101000a81548160ff0219169083151502179055505b6104f082610ae9565b82600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550801561058a5760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516105819190611282565b60405180910390a15b505050565b6040518060400160405280600781526020017f7374616b696e670000000000000000000000000000000000000000000000000081525081565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600047905060008111610636576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161062d9061130f565b60405180910390fd5b61063e610b2d565b610646610c32565b73ffffffffffffffffffffffffffffffffffffffff16637bc742256040518163ffffffff1660e01b8152600401602060405180830381865afa158015610690573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b4919061135b565b600481905550600454670de0b6b3a7640000826106d19190611388565b6106db91906113f9565b600360008282546106ec9190611206565b9250508190555080600260008282546107059190611206565b9250508190555050565b6107176103aa565b61072033610d0a565b565b6040518060400160405280601081526020017f737570706f727465642d746f6b656e730000000000000000000000000000000081525081565b60045481565b60025481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6107956103aa565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015490506107ec6107e7610c32565b610d0a565b6107f4610c32565b73ffffffffffffffffffffffffffffffffffffffff1663f384032833836040518363ffffffff1660e01b815260040161082e929190611439565b600060405180830381600087803b15801561084857600080fd5b505af115801561085c573d6000803e3d6000fd5b5050505050565b60056020528060005260406000206000915090508060000154908060010154905082565b60035481565b6040518060400160405280601881526020017f7265776172642d646973747269627574696f6e2d706f6f6c000000000000000081525081565b6040518060400160405280600f81526020017f736c617368696e672d766f74696e67000000000000000000000000000000000081525081565b600080600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015460035461095291906111d2565b9050670de0b6b3a764000081610966610c32565b73ffffffffffffffffffffffffffffffffffffffff16637a766460336040518263ffffffff1660e01b815260040161099e9190611462565b602060405180830381865afa1580156109bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109df919061135b565b6109e99190611388565b6109f391906113f9565b91505090565b60008060019054906101000a900460ff1615610a705760018260ff16148015610a285750610a2630610edc565b155b610a67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5e906114ef565b60405180910390fd5b60009050610ae4565b8160ff1660008054906101000a900460ff1660ff1610610ac5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610abc906114ef565b60405180910390fd5b816000806101000a81548160ff021916908360ff160217905550600190505b919050565b80600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000610b37610c32565b73ffffffffffffffffffffffffffffffffffffffff1663b7ab4db56040518163ffffffff1660e01b8152600401600060405180830381865afa158015610b81573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610baa919061166c565b905060005b8151811015610c2e5760035460056000848481518110610bd257610bd16116b5565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055508080610c26906116e4565b915050610baf565b5050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663358177736040518060400160405280600781526020017f7374616b696e67000000000000000000000000000000000000000000000000008152506040518263ffffffff1660e01b8152600401610cc49190610f8f565b602060405180830381865afa158015610ce1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d05919061172c565b905090565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154905060008111610d94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8b906117cb565b60405180910390fd5b80600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000828254610de691906111d2565b9250508190555060008273ffffffffffffffffffffffffffffffffffffffff168261520890604051610e179061181c565b600060405180830381858888f193505050503d8060008114610e55576040519150601f19603f3d011682016040523d82523d6000602084013e610e5a565b606091505b5050905080610e9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e95906118a3565b60405180910390fd5b7f2a2d1456672a2b5013c6d74f8677f133bbf5f7d5bb6be09231f7814782b9a7173383604051610ecf929190611439565b60405180910390a1505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610f39578082015181840152602081019050610f1e565b60008484015250505050565b6000601f19601f8301169050919050565b6000610f6182610eff565b610f6b8185610f0a565b9350610f7b818560208601610f1b565b610f8481610f45565b840191505092915050565b60006020820190508181036000830152610fa98184610f56565b905092915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ff082610fc5565b9050919050565b61100081610fe5565b811461100b57600080fd5b50565b60008135905061101d81610ff7565b92915050565b6000806040838503121561103a57611039610fbb565b5b60006110488582860161100e565b92505060206110598582860161100e565b9150509250929050565b6000819050919050565b600061108861108361107e84610fc5565b611063565b610fc5565b9050919050565b600061109a8261106d565b9050919050565b60006110ac8261108f565b9050919050565b6110bc816110a1565b82525050565b60006020820190506110d760008301846110b3565b92915050565b6000819050919050565b6110f0816110dd565b82525050565b600060208201905061110b60008301846110e7565b92915050565b600061111c8261108f565b9050919050565b61112c81611111565b82525050565b60006020820190506111476000830184611123565b92915050565b60006020828403121561116357611162610fbb565b5b60006111718482850161100e565b91505092915050565b600060408201905061118f60008301856110e7565b61119c60208301846110e7565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006111dd826110dd565b91506111e8836110dd565b9250828203905081811115611200576111ff6111a3565b5b92915050565b6000611211826110dd565b915061121c836110dd565b9250828201905080821115611234576112336111a3565b5b92915050565b6000819050919050565b600060ff82169050919050565b600061126c6112676112628461123a565b611063565b611244565b9050919050565b61127c81611251565b82525050565b60006020820190506112976000830184611273565b92915050565b7f526577617264446973747269627574696f6e506f6f6c3a20616d6f756e74206d60008201527f7573742062652067726561746572207468616e20300000000000000000000000602082015250565b60006112f9603583610f0a565b91506113048261129d565b604082019050919050565b60006020820190508181036000830152611328816112ec565b9050919050565b611338816110dd565b811461134357600080fd5b50565b6000815190506113558161132f565b92915050565b60006020828403121561137157611370610fbb565b5b600061137f84828501611346565b91505092915050565b6000611393826110dd565b915061139e836110dd565b92508282026113ac816110dd565b915082820484148315176113c3576113c26111a3565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611404826110dd565b915061140f836110dd565b92508261141f5761141e6113ca565b5b828204905092915050565b61143381610fe5565b82525050565b600060408201905061144e600083018561142a565b61145b60208301846110e7565b9392505050565b6000602082019050611477600083018461142a565b92915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60006114d9602e83610f0a565b91506114e48261147d565b604082019050919050565b60006020820190508181036000830152611508816114cc565b9050919050565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61154c82610f45565b810181811067ffffffffffffffff8211171561156b5761156a611514565b5b80604052505050565b600061157e610fb1565b905061158a8282611543565b919050565b600067ffffffffffffffff8211156115aa576115a9611514565b5b602082029050602081019050919050565b600080fd5b6000815190506115cf81610ff7565b92915050565b60006115e86115e38461158f565b611574565b9050808382526020820190506020840283018581111561160b5761160a6115bb565b5b835b81811015611634578061162088826115c0565b84526020840193505060208101905061160d565b5050509392505050565b600082601f8301126116535761165261150f565b5b81516116638482602086016115d5565b91505092915050565b60006020828403121561168257611681610fbb565b5b600082015167ffffffffffffffff8111156116a05761169f610fc0565b5b6116ac8482850161163e565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006116ef826110dd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611721576117206111a3565b5b600182019050919050565b60006020828403121561174257611741610fbb565b5b6000611750848285016115c0565b91505092915050565b7f526577617264446973747269627574696f6e506f6f6c3a20726577617264206d60008201527f7573742062652067726561746572207468616e20300000000000000000000000602082015250565b60006117b5603583610f0a565b91506117c082611759565b604082019050919050565b600060208201905081810360008301526117e4816117a8565b9050919050565b600081905092915050565b50565b60006118066000836117eb565b9150611811826117f6565b600082019050919050565b6000611827826117f9565b9150819050919050565b7f526577617264446973747269627574696f6e506f6f6c3a207472616e7366657260008201527f206661696c656400000000000000000000000000000000000000000000000000602082015250565b600061188d602783610f0a565b915061189882611831565b604082019050919050565b600060208201905081810360008301526118bc81611880565b905091905056fea264697066735822122004244da31e9bb7fef11a6b3f07c4c5834611938c1aa0e0330b75c5199e0a774264736f6c63430008120033",
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
	parsed, err := abi.JSON(strings.NewReader(RewardDistributionPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) DKGKEY() (string, error) {
	return _RewardDistributionPool.Contract.DKGKEY(&_RewardDistributionPool.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) DKGKEY() (string, error) {
	return _RewardDistributionPool.Contract.DKGKEY(&_RewardDistributionPool.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) REWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _RewardDistributionPool.Contract.REWARDDISTRIBUTIONPOOLKEY(&_RewardDistributionPool.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _RewardDistributionPool.Contract.REWARDDISTRIBUTIONPOOLKEY(&_RewardDistributionPool.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) SLASHINGVOTINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.SLASHINGVOTINGKEY(&_RewardDistributionPool.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.SLASHINGVOTINGKEY(&_RewardDistributionPool.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) STAKINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.STAKINGKEY(&_RewardDistributionPool.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) STAKINGKEY() (string, error) {
	return _RewardDistributionPool.Contract.STAKINGKEY(&_RewardDistributionPool.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _RewardDistributionPool.Contract.SUPPORTEDTOKENSKEY(&_RewardDistributionPool.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _RewardDistributionPool.Contract.SUPPORTEDTOKENSKEY(&_RewardDistributionPool.CallOpts)
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

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolSession) ContractRegistry() (common.Address, error) {
	return _RewardDistributionPool.Contract.ContractRegistry(&_RewardDistributionPool.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) ContractRegistry() (common.Address, error) {
	return _RewardDistributionPool.Contract.ContractRegistry(&_RewardDistributionPool.CallOpts)
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

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardDistributionPool.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolSession) SignerGetter() (common.Address, error) {
	return _RewardDistributionPool.Contract.SignerGetter(&_RewardDistributionPool.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_RewardDistributionPool *RewardDistributionPoolCallerSession) SignerGetter() (common.Address, error) {
	return _RewardDistributionPool.Contract.SignerGetter(&_RewardDistributionPool.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _contractRegistry, address _signerGetterAddress) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactor) Initialize(opts *bind.TransactOpts, _contractRegistry common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.contract.Transact(opts, "initialize", _contractRegistry, _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _contractRegistry, address _signerGetterAddress) returns()
func (_RewardDistributionPool *RewardDistributionPoolSession) Initialize(_contractRegistry common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Initialize(&_RewardDistributionPool.TransactOpts, _contractRegistry, _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _contractRegistry, address _signerGetterAddress) returns()
func (_RewardDistributionPool *RewardDistributionPoolTransactorSession) Initialize(_contractRegistry common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _RewardDistributionPool.Contract.Initialize(&_RewardDistributionPool.TransactOpts, _contractRegistry, _signerGetterAddress)
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

// RewardDistributionPoolInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RewardDistributionPool contract.
type RewardDistributionPoolInitializedIterator struct {
	Event *RewardDistributionPoolInitialized // Event containing the contract specifics and raw log

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
func (it *RewardDistributionPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardDistributionPoolInitialized)
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
		it.Event = new(RewardDistributionPoolInitialized)
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
func (it *RewardDistributionPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardDistributionPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardDistributionPoolInitialized represents a Initialized event raised by the RewardDistributionPool contract.
type RewardDistributionPoolInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) FilterInitialized(opts *bind.FilterOpts) (*RewardDistributionPoolInitializedIterator, error) {

	logs, sub, err := _RewardDistributionPool.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RewardDistributionPoolInitializedIterator{contract: _RewardDistributionPool.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RewardDistributionPool *RewardDistributionPoolFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RewardDistributionPoolInitialized) (event.Subscription, error) {

	logs, sub, err := _RewardDistributionPool.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardDistributionPoolInitialized)
				if err := _RewardDistributionPool.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RewardDistributionPool *RewardDistributionPoolFilterer) ParseInitialized(log types.Log) (*RewardDistributionPoolInitialized, error) {
	event := new(RewardDistributionPoolInitialized)
	if err := _RewardDistributionPool.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
