// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package slashingvoting

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

// ISlashingVotingBaseProposalInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlashingVotingBaseProposalInfo struct {
	Validator       common.Address
	Reason          string
	EpochId         *big.Int
	VotingStartTime *big.Int
	VotingEndTime   *big.Int
}

// ISlashingVotingDetailedProposalInfo is an auto generated low-level Go binding around an user-defined struct.
type ISlashingVotingDetailedProposalInfo struct {
	BaseProposalInfo ISlashingVotingBaseProposalInfo
	IsExecuted       bool
	VotedValidators  []common.Address
}

// SlashingVotingMetaData contains all meta data concerning the SlashingVoting contract.
var SlashingVotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getBaseProposalInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structISlashingVoting.BaseProposalInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getDetailedProposalInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structISlashingVoting.BaseProposalInfo\",\"name\":\"baseProposalInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isExecuted\",\"type\":\"bool\"},{\"internalType\":\"address[]\",\"name\":\"votedValidators\",\"type\":\"address[]\"}],\"internalType\":\"structISlashingVoting.DetailedProposalInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentVotes\",\"type\":\"uint256\"}],\"name\":\"getValidatorsPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"hasPendingSlashingProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_votingThresholdPercentage\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"isProposalExecuted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProposalId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_votingThresholdPercentage\",\"type\":\"uint256\"}],\"name\":\"setVotingThresholdPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingThresholdPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061156d806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806374cb30411161008c578063d00c1aec11610066578063d00c1aec146101e0578063dd0ab13314610207578063e6f7d16814610227578063fe4b84df1461023a57600080fd5b806374cb3041146101a45780638cb941cc146101ad5780638ee7e74c146101c057600080fd5b80633e3b5b19116100c85780633e3b5b191461013d578063602ae84814610165578063680770d714610188578063691304511461019157600080fd5b80630121b93f146100ef5780631333659d146101045780631f4f7d2914610117575b600080fd5b6101026100fd366004610fdf565b61024d565b005b610102610112366004610fdf565b610261565b61012a61012536600461100d565b61026e565b6040519081526020015b60405180910390f35b600080516020611518833981519152546040516001600160a01b039091168152602001610134565b610178610173366004611092565b610519565b6040519015158152602001610134565b61012a60045481565b61010261019f3660046110f6565b610558565b61012a60035481565b6101026101bb366004611092565b6106af565b6101d36101ce366004610fdf565b6106cd565b604051610134919061121f565b6101786101ee366004610fdf565b6000908152600560208190526040909120015460ff1690565b61021a610215366004610fdf565b610722565b604051610134919061129e565b61012a610235366004610fdf565b61080e565b610102610248366004610fdf565b6108a2565b6102556109b0565b61025e81610a81565b50565b610269610d5c565b600455565b60006102786109b0565b600254604051631015428760e21b81526001600160a01b038681166004830152909116906340550a1c90602401602060405180830381865afa1580156102c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102e691906112b1565b6103515760405162461bcd60e51b815260206004820152603160248201527f536c617368696e67566f74696e673a20546172676574206973206e6f7420616e6044820152701030b1ba34bb32903b30b634b230ba37b960791b60648201526084015b60405180910390fd5b61035a84610519565b156103c65760405162461bcd60e51b815260206004820152603660248201527f536c617368696e67566f74696e673a2056616c696461746f7220616c726561646044820152751e481a185cc81c195b991a5b99c81c1c9bdc1bdcd85b60521b6064820152608401610348565b60006003600081546103d7906112e9565b9182905550600081815260056020526040902080546001600160a01b0319166001600160a01b0388161781559091506001810161041585878361138b565b506002546040805163180fd87f60e01b815290516000926001600160a01b03169163180fd87f916004808301926060929190829003018187875af1158015610461573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610485919061144c565b8051600284015560208082018051600386015560408084015160048701556001600160a01b038b16600090815260069093529091208590555190915042106104d0576104d083610a81565b6040516001600160a01b038816815283907fcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b3589060200160405180910390a2509095945050505050565b6001600160a01b0381166000908152600660205260408120548015801590610551575060008181526005602052604090206004015442105b9392505050565b610560610e25565b600082905080600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b0316638e68dce46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105ee91906114a8565b600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b03166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610651573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067591906114a8565b600280546001600160a01b0319166001600160a01b03929092169190911790555033600080516020611518833981519152555050565b5050565b6106b7610e25565b61025e8160008051602061151883398151915255565b6106d5610f80565b6000828152600560205260409081902081516060810190925290806106f985610722565b8152600583015460ff161515602082015260400161071960068401610ea9565b90529392505050565b61072a610fa7565b600082815260056020908152604091829020825160a0810190935280546001600160a01b0316835260018101805491939283019161076790611302565b80601f016020809104026020016040519081016040528092919081815260200182805461079390611302565b80156107e05780601f106107b5576101008083540402835291602001916107e0565b820191906000526020600020905b8154815290600101906020018083116107c357829003601f168201915b5050505050815260200182600201548152602001826003015481526020018260040154815250915050919050565b6002546040805163031147c360e41b815290516000926001600160a01b0316916331147c309160048083019260209291908290030181865afa158015610858573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087c91906114c5565b6108926b033b2e3c9fd0803ce8000000846114de565b61089c91906114f5565b92915050565b600054610100900460ff16158080156108c25750600054600160ff909116105b806108dc5750303b1580156108dc575060005460ff166001145b61093f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610348565b6000805460ff191660011790558015610962576000805461ff0019166101001790555b600482905580156106ab576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b600254604051631015428760e21b81523360048201526001600160a01b03909116906340550a1c90602401602060405180830381865afa1580156109f8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a1c91906112b1565b610a7f5760405162461bcd60e51b815260206004820152602e60248201527f536c617368696e67566f74696e673a204e6f7420616e2061637469766520737960448201526d39ba32b6903b30b634b230ba37b960911b6064820152608401610348565b565b60008181526005602081905260409091209081015460ff1615610afc5760405162461bcd60e51b815260206004820152602d60248201527f536c617368696e67566f74696e673a2050726f706f73616c2068617320616c7260448201526c1958591e48195e1958dd5d1959609a1b6064820152608401610348565b42816003015411158015610b135750806004015442105b610b7d5760405162461bcd60e51b815260206004820152603560248201527f536c617368696e67566f74696e673a20566f74696e67206861736e2774207374604482015274185c9d1959081e595d081bdc88199a5b9a5cda1959605a1b6064820152608401610348565b610b8a6006820133610eb6565b610bea5760405162461bcd60e51b815260206004820152602b60248201527f536c617368696e67566f74696e673a2056616c696461746f722068617320616c60448201526a1c9958591e481d9bdd195960aa1b6064820152608401610348565b600454610bfc61023583600601610ecb565b10610d2b578054600254604051632388fac360e11b81526001600160a01b03928316600482018190529290911690634711f58690602401600060405180830381600087803b158015610c4d57600080fd5b505af1158015610c61573d6000803e3d6000fd5b5050600154604051632388fac360e11b81526001600160a01b0385811660048301529091169250634711f5869150602401600060405180830381600087803b158015610cac57600080fd5b505af1158015610cc0573d6000803e3d6000fd5b5050505060058201805460ff191660011790556001600160a01b0381166000818152600660209081526040808320929092558151868152908101929092527f9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24910160405180910390a1505b604051339083907f34647120636ce03d739347e28e9d1bdd34fd03d066ae60439fb1be2721ad16cd90600090a35050565b60005460408051637ac3c02f60e01b8152905133926201000090046001600160a01b031691637ac3c02f9160048083019260209291908290030181865afa158015610dab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dcf91906114a8565b6001600160a01b031614610a7f5760405162461bcd60e51b815260206004820152601c60248201527f536c617368696e67566f74696e673a204e6f742061207369676e6572000000006044820152606401610348565b6000610e3d6000805160206115188339815191525490565b90506001600160a01b0381161580610e5d57506001600160a01b03811633145b61025e5760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f720000000000006044820152606401610348565b6060600061055183610ed5565b6000610551836001600160a01b038416610f31565b600061089c825490565b606081600001805480602002602001604051908101604052809291908181526020018280548015610f2557602002820191906000526020600020905b815481526020019060010190808311610f11575b50505050509050919050565b6000818152600183016020526040812054610f785750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561089c565b50600061089c565b6040518060600160405280610f93610fa7565b815260006020820152606060409091015290565b6040518060a0016040528060006001600160a01b03168152602001606081526020016000815260200160008152602001600081525090565b600060208284031215610ff157600080fd5b5035919050565b6001600160a01b038116811461025e57600080fd5b60008060006040848603121561102257600080fd5b833561102d81610ff8565b9250602084013567ffffffffffffffff8082111561104a57600080fd5b818601915086601f83011261105e57600080fd5b81358181111561106d57600080fd5b87602082850101111561107f57600080fd5b6020830194508093505050509250925092565b6000602082840312156110a457600080fd5b813561055181610ff8565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156110ee576110ee6110af565b604052919050565b6000806040838503121561110957600080fd5b823561111481610ff8565b915060208381013567ffffffffffffffff8082111561113257600080fd5b818601915086601f83011261114657600080fd5b813581811115611158576111586110af565b61116a601f8201601f191685016110c5565b9150808252878482850101111561118057600080fd5b80848401858401376000848284010152508093505050509250929050565b60018060a01b038151168252600060208083015160a08286015280518060a087015260005b818110156111df5782810184015187820160c0015283016111c3565b50600060c0828801015260408501516040870152606085015160608701526080850151608087015260c0601f19601f830116870101935050505092915050565b60006020808352835160608285015261123b608085018261119e565b858301511515604086810191909152860151858203601f19016060870152805180835290840192506000918401905b808310156112935783516001600160a01b0316825292840192600192909201919084019061126a565b509695505050505050565b602081526000610551602083018461119e565b6000602082840312156112c357600080fd5b8151801515811461055157600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016112fb576112fb6112d3565b5060010190565b600181811c9082168061131657607f821691505b60208210810361133657634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561138657600081815260208120601f850160051c810160208610156113635750805b601f850160051c820191505b818110156113825782815560010161136f565b5050505b505050565b67ffffffffffffffff8311156113a3576113a36110af565b6113b7836113b18354611302565b8361133c565b6000601f8411600181146113eb57600085156113d35750838201355b600019600387901b1c1916600186901b178355611445565b600083815260209020601f19861690835b8281101561141c57868501358255602094850194600190920191016113fc565b50868210156114395760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b60006060828403121561145e57600080fd5b6040516060810181811067ffffffffffffffff82111715611481576114816110af565b80604052508251815260208301516020820152604083015160408201528091505092915050565b6000602082840312156114ba57600080fd5b815161055181610ff8565b6000602082840312156114d757600080fd5b5051919050565b808202811582820484141761089c5761089c6112d3565b60008261151257634e487b7160e01b600052601260045260246000fd5b50049056fe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a26469706673582212209f42722a040b71faf64e37f253fa37231dbcd9d7d7f31490fed9220512819f9f64736f6c63430008120033",
}

// SlashingVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashingVotingMetaData.ABI instead.
var SlashingVotingABI = SlashingVotingMetaData.ABI

// SlashingVotingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlashingVotingMetaData.Bin instead.
var SlashingVotingBin = SlashingVotingMetaData.Bin

// DeploySlashingVoting deploys a new Ethereum contract, binding an instance of SlashingVoting to it.
func DeploySlashingVoting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SlashingVoting, error) {
	parsed, err := SlashingVotingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlashingVotingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SlashingVoting{SlashingVotingCaller: SlashingVotingCaller{contract: contract}, SlashingVotingTransactor: SlashingVotingTransactor{contract: contract}, SlashingVotingFilterer: SlashingVotingFilterer{contract: contract}}, nil
}

// SlashingVoting is an auto generated Go binding around an Ethereum contract.
type SlashingVoting struct {
	SlashingVotingCaller     // Read-only binding to the contract
	SlashingVotingTransactor // Write-only binding to the contract
	SlashingVotingFilterer   // Log filterer for contract events
}

// SlashingVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlashingVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlashingVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlashingVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlashingVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlashingVotingSession struct {
	Contract     *SlashingVoting   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlashingVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlashingVotingCallerSession struct {
	Contract *SlashingVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SlashingVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlashingVotingTransactorSession struct {
	Contract     *SlashingVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SlashingVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlashingVotingRaw struct {
	Contract *SlashingVoting // Generic contract binding to access the raw methods on
}

// SlashingVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlashingVotingCallerRaw struct {
	Contract *SlashingVotingCaller // Generic read-only contract binding to access the raw methods on
}

// SlashingVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlashingVotingTransactorRaw struct {
	Contract *SlashingVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlashingVoting creates a new instance of SlashingVoting, bound to a specific deployed contract.
func NewSlashingVoting(address common.Address, backend bind.ContractBackend) (*SlashingVoting, error) {
	contract, err := bindSlashingVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SlashingVoting{SlashingVotingCaller: SlashingVotingCaller{contract: contract}, SlashingVotingTransactor: SlashingVotingTransactor{contract: contract}, SlashingVotingFilterer: SlashingVotingFilterer{contract: contract}}, nil
}

// NewSlashingVotingCaller creates a new read-only instance of SlashingVoting, bound to a specific deployed contract.
func NewSlashingVotingCaller(address common.Address, caller bind.ContractCaller) (*SlashingVotingCaller, error) {
	contract, err := bindSlashingVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingVotingCaller{contract: contract}, nil
}

// NewSlashingVotingTransactor creates a new write-only instance of SlashingVoting, bound to a specific deployed contract.
func NewSlashingVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*SlashingVotingTransactor, error) {
	contract, err := bindSlashingVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlashingVotingTransactor{contract: contract}, nil
}

// NewSlashingVotingFilterer creates a new log filterer instance of SlashingVoting, bound to a specific deployed contract.
func NewSlashingVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*SlashingVotingFilterer, error) {
	contract, err := bindSlashingVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlashingVotingFilterer{contract: contract}, nil
}

// bindSlashingVoting binds a generic wrapper to an already deployed contract.
func bindSlashingVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlashingVotingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingVoting *SlashingVotingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingVoting.Contract.SlashingVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingVoting *SlashingVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SlashingVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingVoting *SlashingVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SlashingVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SlashingVoting *SlashingVotingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SlashingVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SlashingVoting *SlashingVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SlashingVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SlashingVoting *SlashingVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SlashingVoting.Contract.contract.Transact(opts, method, params...)
}

// GetBaseProposalInfo is a free data retrieval call binding the contract method 0xdd0ab133.
//
// Solidity: function getBaseProposalInfo(uint256 _proposalId) view returns((address,string,uint256,uint256,uint256))
func (_SlashingVoting *SlashingVotingCaller) GetBaseProposalInfo(opts *bind.CallOpts, _proposalId *big.Int) (ISlashingVotingBaseProposalInfo, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "getBaseProposalInfo", _proposalId)

	if err != nil {
		return *new(ISlashingVotingBaseProposalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingVotingBaseProposalInfo)).(*ISlashingVotingBaseProposalInfo)

	return out0, err

}

// GetBaseProposalInfo is a free data retrieval call binding the contract method 0xdd0ab133.
//
// Solidity: function getBaseProposalInfo(uint256 _proposalId) view returns((address,string,uint256,uint256,uint256))
func (_SlashingVoting *SlashingVotingSession) GetBaseProposalInfo(_proposalId *big.Int) (ISlashingVotingBaseProposalInfo, error) {
	return _SlashingVoting.Contract.GetBaseProposalInfo(&_SlashingVoting.CallOpts, _proposalId)
}

// GetBaseProposalInfo is a free data retrieval call binding the contract method 0xdd0ab133.
//
// Solidity: function getBaseProposalInfo(uint256 _proposalId) view returns((address,string,uint256,uint256,uint256))
func (_SlashingVoting *SlashingVotingCallerSession) GetBaseProposalInfo(_proposalId *big.Int) (ISlashingVotingBaseProposalInfo, error) {
	return _SlashingVoting.Contract.GetBaseProposalInfo(&_SlashingVoting.CallOpts, _proposalId)
}

// GetDetailedProposalInfo is a free data retrieval call binding the contract method 0x8ee7e74c.
//
// Solidity: function getDetailedProposalInfo(uint256 _proposalId) view returns(((address,string,uint256,uint256,uint256),bool,address[]))
func (_SlashingVoting *SlashingVotingCaller) GetDetailedProposalInfo(opts *bind.CallOpts, _proposalId *big.Int) (ISlashingVotingDetailedProposalInfo, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "getDetailedProposalInfo", _proposalId)

	if err != nil {
		return *new(ISlashingVotingDetailedProposalInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ISlashingVotingDetailedProposalInfo)).(*ISlashingVotingDetailedProposalInfo)

	return out0, err

}

// GetDetailedProposalInfo is a free data retrieval call binding the contract method 0x8ee7e74c.
//
// Solidity: function getDetailedProposalInfo(uint256 _proposalId) view returns(((address,string,uint256,uint256,uint256),bool,address[]))
func (_SlashingVoting *SlashingVotingSession) GetDetailedProposalInfo(_proposalId *big.Int) (ISlashingVotingDetailedProposalInfo, error) {
	return _SlashingVoting.Contract.GetDetailedProposalInfo(&_SlashingVoting.CallOpts, _proposalId)
}

// GetDetailedProposalInfo is a free data retrieval call binding the contract method 0x8ee7e74c.
//
// Solidity: function getDetailedProposalInfo(uint256 _proposalId) view returns(((address,string,uint256,uint256,uint256),bool,address[]))
func (_SlashingVoting *SlashingVotingCallerSession) GetDetailedProposalInfo(_proposalId *big.Int) (ISlashingVotingDetailedProposalInfo, error) {
	return _SlashingVoting.Contract.GetDetailedProposalInfo(&_SlashingVoting.CallOpts, _proposalId)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_SlashingVoting *SlashingVotingCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_SlashingVoting *SlashingVotingSession) GetInjector() (common.Address, error) {
	return _SlashingVoting.Contract.GetInjector(&_SlashingVoting.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_SlashingVoting *SlashingVotingCallerSession) GetInjector() (common.Address, error) {
	return _SlashingVoting.Contract.GetInjector(&_SlashingVoting.CallOpts)
}

// GetValidatorsPercentage is a free data retrieval call binding the contract method 0xe6f7d168.
//
// Solidity: function getValidatorsPercentage(uint256 _currentVotes) view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) GetValidatorsPercentage(opts *bind.CallOpts, _currentVotes *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "getValidatorsPercentage", _currentVotes)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsPercentage is a free data retrieval call binding the contract method 0xe6f7d168.
//
// Solidity: function getValidatorsPercentage(uint256 _currentVotes) view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) GetValidatorsPercentage(_currentVotes *big.Int) (*big.Int, error) {
	return _SlashingVoting.Contract.GetValidatorsPercentage(&_SlashingVoting.CallOpts, _currentVotes)
}

// GetValidatorsPercentage is a free data retrieval call binding the contract method 0xe6f7d168.
//
// Solidity: function getValidatorsPercentage(uint256 _currentVotes) view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) GetValidatorsPercentage(_currentVotes *big.Int) (*big.Int, error) {
	return _SlashingVoting.Contract.GetValidatorsPercentage(&_SlashingVoting.CallOpts, _currentVotes)
}

// HasPendingSlashingProposal is a free data retrieval call binding the contract method 0x602ae848.
//
// Solidity: function hasPendingSlashingProposal(address _userAddr) view returns(bool)
func (_SlashingVoting *SlashingVotingCaller) HasPendingSlashingProposal(opts *bind.CallOpts, _userAddr common.Address) (bool, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "hasPendingSlashingProposal", _userAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPendingSlashingProposal is a free data retrieval call binding the contract method 0x602ae848.
//
// Solidity: function hasPendingSlashingProposal(address _userAddr) view returns(bool)
func (_SlashingVoting *SlashingVotingSession) HasPendingSlashingProposal(_userAddr common.Address) (bool, error) {
	return _SlashingVoting.Contract.HasPendingSlashingProposal(&_SlashingVoting.CallOpts, _userAddr)
}

// HasPendingSlashingProposal is a free data retrieval call binding the contract method 0x602ae848.
//
// Solidity: function hasPendingSlashingProposal(address _userAddr) view returns(bool)
func (_SlashingVoting *SlashingVotingCallerSession) HasPendingSlashingProposal(_userAddr common.Address) (bool, error) {
	return _SlashingVoting.Contract.HasPendingSlashingProposal(&_SlashingVoting.CallOpts, _userAddr)
}

// IsProposalExecuted is a free data retrieval call binding the contract method 0xd00c1aec.
//
// Solidity: function isProposalExecuted(uint256 _proposalId) view returns(bool)
func (_SlashingVoting *SlashingVotingCaller) IsProposalExecuted(opts *bind.CallOpts, _proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "isProposalExecuted", _proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposalExecuted is a free data retrieval call binding the contract method 0xd00c1aec.
//
// Solidity: function isProposalExecuted(uint256 _proposalId) view returns(bool)
func (_SlashingVoting *SlashingVotingSession) IsProposalExecuted(_proposalId *big.Int) (bool, error) {
	return _SlashingVoting.Contract.IsProposalExecuted(&_SlashingVoting.CallOpts, _proposalId)
}

// IsProposalExecuted is a free data retrieval call binding the contract method 0xd00c1aec.
//
// Solidity: function isProposalExecuted(uint256 _proposalId) view returns(bool)
func (_SlashingVoting *SlashingVotingCallerSession) IsProposalExecuted(_proposalId *big.Int) (bool, error) {
	return _SlashingVoting.Contract.IsProposalExecuted(&_SlashingVoting.CallOpts, _proposalId)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) LastProposalId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "lastProposalId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) LastProposalId() (*big.Int, error) {
	return _SlashingVoting.Contract.LastProposalId(&_SlashingVoting.CallOpts)
}

// LastProposalId is a free data retrieval call binding the contract method 0x74cb3041.
//
// Solidity: function lastProposalId() view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) LastProposalId() (*big.Int, error) {
	return _SlashingVoting.Contract.LastProposalId(&_SlashingVoting.CallOpts)
}

// VotingThresholdPercentage is a free data retrieval call binding the contract method 0x680770d7.
//
// Solidity: function votingThresholdPercentage() view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) VotingThresholdPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "votingThresholdPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingThresholdPercentage is a free data retrieval call binding the contract method 0x680770d7.
//
// Solidity: function votingThresholdPercentage() view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) VotingThresholdPercentage() (*big.Int, error) {
	return _SlashingVoting.Contract.VotingThresholdPercentage(&_SlashingVoting.CallOpts)
}

// VotingThresholdPercentage is a free data retrieval call binding the contract method 0x680770d7.
//
// Solidity: function votingThresholdPercentage() view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) VotingThresholdPercentage() (*big.Int, error) {
	return _SlashingVoting.Contract.VotingThresholdPercentage(&_SlashingVoting.CallOpts)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address _validatorAddr, string _reason) returns(uint256)
func (_SlashingVoting *SlashingVotingTransactor) CreateProposal(opts *bind.TransactOpts, _validatorAddr common.Address, _reason string) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "createProposal", _validatorAddr, _reason)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address _validatorAddr, string _reason) returns(uint256)
func (_SlashingVoting *SlashingVotingSession) CreateProposal(_validatorAddr common.Address, _reason string) (*types.Transaction, error) {
	return _SlashingVoting.Contract.CreateProposal(&_SlashingVoting.TransactOpts, _validatorAddr, _reason)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address _validatorAddr, string _reason) returns(uint256)
func (_SlashingVoting *SlashingVotingTransactorSession) CreateProposal(_validatorAddr common.Address, _reason string) (*types.Transaction, error) {
	return _SlashingVoting.Contract.CreateProposal(&_SlashingVoting.TransactOpts, _validatorAddr, _reason)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _votingThresholdPercentage) returns()
func (_SlashingVoting *SlashingVotingTransactor) Initialize(opts *bind.TransactOpts, _votingThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "initialize", _votingThresholdPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _votingThresholdPercentage) returns()
func (_SlashingVoting *SlashingVotingSession) Initialize(_votingThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _votingThresholdPercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _votingThresholdPercentage) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) Initialize(_votingThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _votingThresholdPercentage)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_SlashingVoting *SlashingVotingTransactor) SetDependencies(opts *bind.TransactOpts, _contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "setDependencies", _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_SlashingVoting *SlashingVotingSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetDependencies(&_SlashingVoting.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetDependencies(&_SlashingVoting.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_SlashingVoting *SlashingVotingTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_SlashingVoting *SlashingVotingSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetInjector(&_SlashingVoting.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetInjector(&_SlashingVoting.TransactOpts, injector_)
}

// SetVotingThresholdPercentage is a paid mutator transaction binding the contract method 0x1333659d.
//
// Solidity: function setVotingThresholdPercentage(uint256 _votingThresholdPercentage) returns()
func (_SlashingVoting *SlashingVotingTransactor) SetVotingThresholdPercentage(opts *bind.TransactOpts, _votingThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "setVotingThresholdPercentage", _votingThresholdPercentage)
}

// SetVotingThresholdPercentage is a paid mutator transaction binding the contract method 0x1333659d.
//
// Solidity: function setVotingThresholdPercentage(uint256 _votingThresholdPercentage) returns()
func (_SlashingVoting *SlashingVotingSession) SetVotingThresholdPercentage(_votingThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetVotingThresholdPercentage(&_SlashingVoting.TransactOpts, _votingThresholdPercentage)
}

// SetVotingThresholdPercentage is a paid mutator transaction binding the contract method 0x1333659d.
//
// Solidity: function setVotingThresholdPercentage(uint256 _votingThresholdPercentage) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) SetVotingThresholdPercentage(_votingThresholdPercentage *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetVotingThresholdPercentage(&_SlashingVoting.TransactOpts, _votingThresholdPercentage)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalId) returns()
func (_SlashingVoting *SlashingVotingTransactor) Vote(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "vote", _proposalId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalId) returns()
func (_SlashingVoting *SlashingVotingSession) Vote(_proposalId *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Vote(&_SlashingVoting.TransactOpts, _proposalId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalId) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) Vote(_proposalId *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Vote(&_SlashingVoting.TransactOpts, _proposalId)
}

// SlashingVotingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SlashingVoting contract.
type SlashingVotingInitializedIterator struct {
	Event *SlashingVotingInitialized // Event containing the contract specifics and raw log

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
func (it *SlashingVotingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingVotingInitialized)
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
		it.Event = new(SlashingVotingInitialized)
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
func (it *SlashingVotingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingVotingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingVotingInitialized represents a Initialized event raised by the SlashingVoting contract.
type SlashingVotingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashingVoting *SlashingVotingFilterer) FilterInitialized(opts *bind.FilterOpts) (*SlashingVotingInitializedIterator, error) {

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SlashingVotingInitializedIterator{contract: _SlashingVoting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SlashingVoting *SlashingVotingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SlashingVotingInitialized) (event.Subscription, error) {

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingVotingInitialized)
				if err := _SlashingVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SlashingVoting *SlashingVotingFilterer) ParseInitialized(log types.Log) (*SlashingVotingInitialized, error) {
	event := new(SlashingVotingInitialized)
	if err := _SlashingVoting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashingVotingProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the SlashingVoting contract.
type SlashingVotingProposalCreatedIterator struct {
	Event *SlashingVotingProposalCreated // Event containing the contract specifics and raw log

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
func (it *SlashingVotingProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingVotingProposalCreated)
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
		it.Event = new(SlashingVotingProposalCreated)
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
func (it *SlashingVotingProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingVotingProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingVotingProposalCreated represents a ProposalCreated event raised by the SlashingVoting contract.
type SlashingVotingProposalCreated struct {
	ProposalId *big.Int
	Validator  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0xcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts, proposalId []*big.Int) (*SlashingVotingProposalCreatedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &SlashingVotingProposalCreatedIterator{contract: _SlashingVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *SlashingVotingProposalCreated, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "ProposalCreated", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingVotingProposalCreated)
				if err := _SlashingVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0xcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358.
//
// Solidity: event ProposalCreated(uint256 indexed proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) ParseProposalCreated(log types.Log) (*SlashingVotingProposalCreated, error) {
	event := new(SlashingVotingProposalCreated)
	if err := _SlashingVoting.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashingVotingProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the SlashingVoting contract.
type SlashingVotingProposalExecutedIterator struct {
	Event *SlashingVotingProposalExecuted // Event containing the contract specifics and raw log

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
func (it *SlashingVotingProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingVotingProposalExecuted)
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
		it.Event = new(SlashingVotingProposalExecuted)
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
func (it *SlashingVotingProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingVotingProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingVotingProposalExecuted represents a ProposalExecuted event raised by the SlashingVoting contract.
type SlashingVotingProposalExecuted struct {
	ProposalId *big.Int
	Validator  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24.
//
// Solidity: event ProposalExecuted(uint256 proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*SlashingVotingProposalExecutedIterator, error) {

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &SlashingVotingProposalExecutedIterator{contract: _SlashingVoting.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24.
//
// Solidity: event ProposalExecuted(uint256 proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *SlashingVotingProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingVotingProposalExecuted)
				if err := _SlashingVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24.
//
// Solidity: event ProposalExecuted(uint256 proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) ParseProposalExecuted(log types.Log) (*SlashingVotingProposalExecuted, error) {
	event := new(SlashingVotingProposalExecuted)
	if err := _SlashingVoting.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashingVotingProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the SlashingVoting contract.
type SlashingVotingProposalVotedIterator struct {
	Event *SlashingVotingProposalVoted // Event containing the contract specifics and raw log

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
func (it *SlashingVotingProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingVotingProposalVoted)
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
		it.Event = new(SlashingVotingProposalVoted)
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
func (it *SlashingVotingProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingVotingProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingVotingProposalVoted represents a ProposalVoted event raised by the SlashingVoting contract.
type SlashingVotingProposalVoted struct {
	ProposalId *big.Int
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0x34647120636ce03d739347e28e9d1bdd34fd03d066ae60439fb1be2721ad16cd.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed voter)
func (_SlashingVoting *SlashingVotingFilterer) FilterProposalVoted(opts *bind.FilterOpts, proposalId []*big.Int, voter []common.Address) (*SlashingVotingProposalVotedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "ProposalVoted", proposalIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &SlashingVotingProposalVotedIterator{contract: _SlashingVoting.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0x34647120636ce03d739347e28e9d1bdd34fd03d066ae60439fb1be2721ad16cd.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed voter)
func (_SlashingVoting *SlashingVotingFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *SlashingVotingProposalVoted, proposalId []*big.Int, voter []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "ProposalVoted", proposalIdRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingVotingProposalVoted)
				if err := _SlashingVoting.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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

// ParseProposalVoted is a log parse operation binding the contract event 0x34647120636ce03d739347e28e9d1bdd34fd03d066ae60439fb1be2721ad16cd.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed voter)
func (_SlashingVoting *SlashingVotingFilterer) ParseProposalVoted(log types.Log) (*SlashingVotingProposalVoted, error) {
	event := new(SlashingVotingProposalVoted)
	if err := _SlashingVoting.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
