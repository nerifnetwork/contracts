// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package billingmanager

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

// IBillingManagerUserFundsInfo is an auto generated low-level Go binding around an user-defined struct.
type IBillingManagerUserFundsInfo struct {
	UserFundBalance             *big.Int
	UserLockedBalance           *big.Int
	PendingWorkflowExecutionIds []*big.Int
}

// IBillingManagerWorkflowExecutionInfo is an auto generated low-level Go binding around an user-defined struct.
type IBillingManagerWorkflowExecutionInfo struct {
	WorkflowId            *big.Int
	ExecutionLockedAmount *big.Int
	ExecutionAmount       *big.Int
	WorkflowOwner         common.Address
	Status                uint8
}

// BillingManagerMetaData contains all meta data concerning the BillingManager contract.
var BillingManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userCurrentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundedAmount\",\"type\":\"uint256\"}],\"name\":\"BalanceFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"}],\"name\":\"ExecutionFundsLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"RewardsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedExecutionAmountFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"name\":\"UserFundsWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_executionAmount\",\"type\":\"uint256\"}],\"name\":\"completeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"}],\"name\":\"fundBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getExecutionWorkflowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getUserAvailableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getUserFundsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"userFundBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLockedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingWorkflowExecutionIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIBillingManager.UserFundsInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"enumIBillingManager.WorkflowExecutionStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIBillingManager.WorkflowExecutionInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionStatus\",\"outputs\":[{\"internalType\":\"enumIBillingManager.WorkflowExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_executionLockedAmount\",\"type\":\"uint256\"}],\"name\":\"lockExecutionFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextWorkflowExecutionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNetworkRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061163d806100206000396000f3fe6080604052600436106101025760003560e01c80636b5d420611610095578063800175631161006457806380017563146102e9578063c4cffb7614610309578063ef4f272d14610329578063f658681214610356578063f9b41fd51461037657600080fd5b80636b5d4206146102335780636e136778146102495780637105e970146102825780637b103999146102c957600080fd5b8063485cc955116100d1578063485cc9551461017b5780634a686e491461019b5780635c211f88146101c85780636b0912261461020657600080fd5b80630ad5a86514610117578063155dd5ee1461012a578063242939411461014a5780633c0679451461017357600080fd5b36610112576101103361038b565b005b600080fd5b6101106101253660046113c8565b610474565b34801561013657600080fd5b506101106101453660046113e5565b610480565b34801561015657600080fd5b5061016060035481565b6040519081526020015b60405180910390f35b61011061055e565b34801561018757600080fd5b506101106101963660046113fe565b610569565b3480156101a757600080fd5b506101606101b63660046113e5565b60009081526005602052604090205490565b3480156101d457600080fd5b506000546101ee906201000090046001600160a01b031681565b6040516001600160a01b03909116815260200161016a565b34801561021257600080fd5b506102266102213660046113c8565b6106bf565b60405161016a9190611437565b34801561023f57600080fd5b5061016060025481565b34801561025557600080fd5b506101ee6102643660046113e5565b6000908152600560205260409020600301546001600160a01b031690565b34801561028e57600080fd5b506102bc61029d3660046113e5565b600090815260056020526040902060030154600160a01b900460ff1690565b60405161016a91906114d5565b3480156102d557600080fd5b506001546101ee906001600160a01b031681565b3480156102f557600080fd5b506101106103043660046114e3565b61072f565b34801561031557600080fd5b506101606103243660046113c8565b610a48565b34801561033557600080fd5b506103496103443660046113e5565b610a77565b60405161016a9190611505565b34801561036257600080fd5b506101106103713660046114e3565b610b08565b34801561038257600080fd5b50610110610edd565b600034116103ea5760405162461bcd60e51b815260206004820152602160248201527f42696c6c696e674d616e616765723a205a65726f2066756e647320746f2061646044820152601960fa1b60648201526084015b60405180910390fd5b6001600160a01b03811660009081526004602052604081205461040e90349061155d565b6001600160a01b0383166000818152600460205260409081902083905551919250907f3be325f4f5fcedd7332727df6cd7ca2679dbd8e6a6332f45e0faeeb23b7d54a1906104689084903490918252602082015260400190565b60405180910390a25050565b61047d8161038b565b50565b8061048a33610a48565b10156104f75760405162461bcd60e51b815260206004820152603660248201527f42696c6c696e674d616e616765723a204e6f7420656e6f75676820617661696c60448201527561626c652066756e647320746f20776974686472617760501b60648201526084016103e1565b3360009081526004602052604081208054839290610516908490611570565b909155506105269050338261106b565b60405181815233907f7f9d7c82d89883dcf97d377c4a3098b9a25d87d6d25dc4b289531d7a38460a089060200160405180910390a250565b6105673361038b565b565b600054610100900460ff16158080156105895750600054600160ff909116105b806105a35750303b1580156105a3575060005460ff166001145b6106065760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103e1565b6000805460ff191660011790558015610629576000805461ff0019166101001790555b600180546001600160a01b0385166001600160a01b031990911617905561067482600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b80156106ba576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6106e360405180606001604052806000815260200160008152602001606081525090565b6001600160a01b03821660009081526004602090815260409182902082516060810184528154815260018201549281019290925291810161072660028401611184565b90529392505050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa15801561077e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a29190611583565b6001600160a01b0316146107f85760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e657200000000000060448201526064016103e1565b600082815260056020526040902060016003820154600160a01b900460ff1660028111156108285761082861149d565b1461088e5760405162461bcd60e51b815260206004820152603060248201527f42696c6c696e674d616e616765723a204e6f7420612070656e64696e6720776f60448201526f3935b33637bb9032bc32b1baba34b7b760811b60648201526084016103e1565b60038101546001600160a01b03166000908152600460205260409020600182015483811015610924576108df8482846001015485600001546108d09190611570565b6108da919061155d565b611198565b60408051878152602081018490529081018290529094507faba1d671589f510878d4421f3e9f5574d37ecf6735d3f3df2be569b62b9015d99060600160405180910390a15b60038301805460ff60a01b1916600160a11b17905560028301849055815484908390600090610954908490611570565b925050819055508082600101600082825461096f9190611570565b90915550610982905060028301866111ae565b508360026000828254610995919061155d565b90915550506001548354604051632851165f60e21b81526001600160a01b039092169163a144597c916109d5918890600401918252602082015260400190565b600060405180830381600087803b1580156109ef57600080fd5b505af1158015610a03573d6000803e3d6000fd5b505060408051888152602081018890527f70114d139c12c9f15fb8a74cc0b9d9f2070f048d32cdf038844425444bad7b52935001905060405180910390a15050505050565b6001600160a01b038116600090815260046020526040812060018101549054610a719190611570565b92915050565b610a7f61136b565b600082815260056020908152604091829020825160a081018452815481526001820154928101929092526002808201549383019390935260038101546001600160a01b0381166060840152919290916080840191600160a01b90910460ff1690811115610aee57610aee61149d565b6002811115610aff57610aff61149d565b90525092915050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015610b57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7b9190611583565b6001600160a01b031614610bd15760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e657200000000000060448201526064016103e1565b600154604051639b2bfaa360e01b8152600481018490526001600160a01b0390911690639b2bfaa390602401602060405180830381865afa158015610c1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3e91906115a0565b610c9a5760405162461bcd60e51b815260206004820152602760248201527f42696c6c696e674d616e616765723a20576f726b666c6f7720646f6573206e6f6044820152661d08195e1a5cdd60ca1b60648201526084016103e1565b60015460405163d69cd27560e01b8152600481018490526000916001600160a01b03169063d69cd27590602401602060405180830381865afa158015610ce4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d089190611583565b6001600160a01b038116600090815260046020526040902090915082610d2d83610a48565b1015610d8c5760405162461bcd60e51b815260206004820152602860248201527f42696c6c696e674d616e616765723a204e6f7420656e6f7567682066756e647360448201526720746f206c6f636b60c01b60648201526084016103e1565b6003805460009182610d9d836115c2565b91905055905083826001016000828254610db7919061155d565b90915550610dca905060028301826111ba565b506040518060a0016040528086815260200185815260200160008152602001846001600160a01b0316815260200160016002811115610e0b57610e0b61149d565b9052600082815260056020908152604091829020835181559083015160018201559082015160028083019190915560608301516003830180546001600160a01b039092166001600160a01b031983168117825560808601519391926001600160a81b0319161790600160a01b908490811115610e8957610e8961149d565b02179055505060408051838152602081018790526001600160a01b038616925087917fc77a14be8053ff3df25436e566ff7bd5d0da84a96e23ced50e403617fab5e97d910160405180910390a35050505050565b60025480610f445760405162461bcd60e51b815260206004820152602e60248201527f42696c6c696e674d616e616765723a204e6f206e6574776f726b20726577617260448201526d647320746f20776974686472617760901b60648201526084016103e1565b60026000905560008060029054906101000a90046001600160a01b03166001600160a01b0316631a296e026040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f9e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc29190611583565b90506001600160a01b0381166110265760405162461bcd60e51b815260206004820152602360248201527f42696c6c696e674d616e616765723a205a65726f207369676e6572206164647260448201526265737360e81b60648201526084016103e1565b611030818361106b565b806001600160a01b03167f8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e5031618360405161046891815260200190565b804710156110bb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016103e1565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611108576040519150601f19603f3d011682016040523d82523d6000602084013e61110d565b606091505b50509050806106ba5760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016103e1565b60606000611191836111c6565b9392505050565b60008183106111a75781611191565b5090919050565b60006111918383611222565b6000611191838361131c565b60608160000180548060200260200160405190810160405280929190818152602001828054801561121657602002820191906000526020600020905b815481526020019060010190808311611202575b50505050509050919050565b6000818152600183016020526040812054801561130b576000611246600183611570565b855490915060009061125a90600190611570565b90508181146112bf57600086600001828154811061127a5761127a6115db565b906000526020600020015490508087600001848154811061129d5761129d6115db565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806112d0576112d06115f1565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610a71565b6000915050610a71565b5092915050565b600081815260018301602052604081205461136357508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610a71565b506000610a71565b6040518060a0016040528060008152602001600081526020016000815260200160006001600160a01b03168152602001600060028111156113ae576113ae61149d565b905290565b6001600160a01b038116811461047d57600080fd5b6000602082840312156113da57600080fd5b8135611191816113b3565b6000602082840312156113f757600080fd5b5035919050565b6000806040838503121561141157600080fd5b823561141c816113b3565b9150602083013561142c816113b3565b809150509250929050565b6000602080835260808301845182850152818501516040850152604085015160608086015281815180845260a0870191508483019350600092505b808310156114925783518252928401926001929092019190840190611472565b509695505050505050565b634e487b7160e01b600052602160045260246000fd5b600381106114d157634e487b7160e01b600052602160045260246000fd5b9052565b60208101610a7182846114b3565b600080604083850312156114f657600080fd5b50508035926020909101359150565b8151815260208083015190820152604080830151908201526060808301516001600160a01b03169082015260808083015160a0830191611315908401826114b3565b634e487b7160e01b600052601160045260246000fd5b80820180821115610a7157610a71611547565b81810381811115610a7157610a71611547565b60006020828403121561159557600080fd5b8151611191816113b3565b6000602082840312156115b257600080fd5b8151801515811461119157600080fd5b6000600182016115d4576115d4611547565b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052603160045260246000fdfea26469706673582212204beaba94d4be786e9a07ddb327b3214a3abd48c965f7381d692507cf6b7b5db764736f6c63430008120033",
}

// BillingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BillingManagerMetaData.ABI instead.
var BillingManagerABI = BillingManagerMetaData.ABI

// BillingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BillingManagerMetaData.Bin instead.
var BillingManagerBin = BillingManagerMetaData.Bin

// DeployBillingManager deploys a new Ethereum contract, binding an instance of BillingManager to it.
func DeployBillingManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BillingManager, error) {
	parsed, err := BillingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BillingManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BillingManager{BillingManagerCaller: BillingManagerCaller{contract: contract}, BillingManagerTransactor: BillingManagerTransactor{contract: contract}, BillingManagerFilterer: BillingManagerFilterer{contract: contract}}, nil
}

// BillingManager is an auto generated Go binding around an Ethereum contract.
type BillingManager struct {
	BillingManagerCaller     // Read-only binding to the contract
	BillingManagerTransactor // Write-only binding to the contract
	BillingManagerFilterer   // Log filterer for contract events
}

// BillingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BillingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BillingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BillingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BillingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BillingManagerSession struct {
	Contract     *BillingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BillingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BillingManagerCallerSession struct {
	Contract *BillingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BillingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BillingManagerTransactorSession struct {
	Contract     *BillingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BillingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BillingManagerRaw struct {
	Contract *BillingManager // Generic contract binding to access the raw methods on
}

// BillingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BillingManagerCallerRaw struct {
	Contract *BillingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BillingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BillingManagerTransactorRaw struct {
	Contract *BillingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBillingManager creates a new instance of BillingManager, bound to a specific deployed contract.
func NewBillingManager(address common.Address, backend bind.ContractBackend) (*BillingManager, error) {
	contract, err := bindBillingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BillingManager{BillingManagerCaller: BillingManagerCaller{contract: contract}, BillingManagerTransactor: BillingManagerTransactor{contract: contract}, BillingManagerFilterer: BillingManagerFilterer{contract: contract}}, nil
}

// NewBillingManagerCaller creates a new read-only instance of BillingManager, bound to a specific deployed contract.
func NewBillingManagerCaller(address common.Address, caller bind.ContractCaller) (*BillingManagerCaller, error) {
	contract, err := bindBillingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BillingManagerCaller{contract: contract}, nil
}

// NewBillingManagerTransactor creates a new write-only instance of BillingManager, bound to a specific deployed contract.
func NewBillingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BillingManagerTransactor, error) {
	contract, err := bindBillingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BillingManagerTransactor{contract: contract}, nil
}

// NewBillingManagerFilterer creates a new log filterer instance of BillingManager, bound to a specific deployed contract.
func NewBillingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BillingManagerFilterer, error) {
	contract, err := bindBillingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BillingManagerFilterer{contract: contract}, nil
}

// bindBillingManager binds a generic wrapper to an already deployed contract.
func bindBillingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BillingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BillingManager *BillingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BillingManager.Contract.BillingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BillingManager *BillingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingManager.Contract.BillingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BillingManager *BillingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BillingManager.Contract.BillingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BillingManager *BillingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BillingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BillingManager *BillingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BillingManager *BillingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BillingManager.Contract.contract.Transact(opts, method, params...)
}

// GetExecutionWorkflowId is a free data retrieval call binding the contract method 0x4a686e49.
//
// Solidity: function getExecutionWorkflowId(uint256 _workflowExecutionId) view returns(uint256)
func (_BillingManager *BillingManagerCaller) GetExecutionWorkflowId(opts *bind.CallOpts, _workflowExecutionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getExecutionWorkflowId", _workflowExecutionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionWorkflowId is a free data retrieval call binding the contract method 0x4a686e49.
//
// Solidity: function getExecutionWorkflowId(uint256 _workflowExecutionId) view returns(uint256)
func (_BillingManager *BillingManagerSession) GetExecutionWorkflowId(_workflowExecutionId *big.Int) (*big.Int, error) {
	return _BillingManager.Contract.GetExecutionWorkflowId(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetExecutionWorkflowId is a free data retrieval call binding the contract method 0x4a686e49.
//
// Solidity: function getExecutionWorkflowId(uint256 _workflowExecutionId) view returns(uint256)
func (_BillingManager *BillingManagerCallerSession) GetExecutionWorkflowId(_workflowExecutionId *big.Int) (*big.Int, error) {
	return _BillingManager.Contract.GetExecutionWorkflowId(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetUserAvailableFunds is a free data retrieval call binding the contract method 0xc4cffb76.
//
// Solidity: function getUserAvailableFunds(address _userAddr) view returns(uint256)
func (_BillingManager *BillingManagerCaller) GetUserAvailableFunds(opts *bind.CallOpts, _userAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getUserAvailableFunds", _userAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAvailableFunds is a free data retrieval call binding the contract method 0xc4cffb76.
//
// Solidity: function getUserAvailableFunds(address _userAddr) view returns(uint256)
func (_BillingManager *BillingManagerSession) GetUserAvailableFunds(_userAddr common.Address) (*big.Int, error) {
	return _BillingManager.Contract.GetUserAvailableFunds(&_BillingManager.CallOpts, _userAddr)
}

// GetUserAvailableFunds is a free data retrieval call binding the contract method 0xc4cffb76.
//
// Solidity: function getUserAvailableFunds(address _userAddr) view returns(uint256)
func (_BillingManager *BillingManagerCallerSession) GetUserAvailableFunds(_userAddr common.Address) (*big.Int, error) {
	return _BillingManager.Contract.GetUserAvailableFunds(&_BillingManager.CallOpts, _userAddr)
}

// GetUserFundsInfo is a free data retrieval call binding the contract method 0x6b091226.
//
// Solidity: function getUserFundsInfo(address _userAddr) view returns((uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerCaller) GetUserFundsInfo(opts *bind.CallOpts, _userAddr common.Address) (IBillingManagerUserFundsInfo, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getUserFundsInfo", _userAddr)

	if err != nil {
		return *new(IBillingManagerUserFundsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBillingManagerUserFundsInfo)).(*IBillingManagerUserFundsInfo)

	return out0, err

}

// GetUserFundsInfo is a free data retrieval call binding the contract method 0x6b091226.
//
// Solidity: function getUserFundsInfo(address _userAddr) view returns((uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerSession) GetUserFundsInfo(_userAddr common.Address) (IBillingManagerUserFundsInfo, error) {
	return _BillingManager.Contract.GetUserFundsInfo(&_BillingManager.CallOpts, _userAddr)
}

// GetUserFundsInfo is a free data retrieval call binding the contract method 0x6b091226.
//
// Solidity: function getUserFundsInfo(address _userAddr) view returns((uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerCallerSession) GetUserFundsInfo(_userAddr common.Address) (IBillingManagerUserFundsInfo, error) {
	return _BillingManager.Contract.GetUserFundsInfo(&_BillingManager.CallOpts, _userAddr)
}

// GetWorkflowExecutionInfo is a free data retrieval call binding the contract method 0xef4f272d.
//
// Solidity: function getWorkflowExecutionInfo(uint256 _workflowExecutionId) view returns((uint256,uint256,uint256,address,uint8))
func (_BillingManager *BillingManagerCaller) GetWorkflowExecutionInfo(opts *bind.CallOpts, _workflowExecutionId *big.Int) (IBillingManagerWorkflowExecutionInfo, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getWorkflowExecutionInfo", _workflowExecutionId)

	if err != nil {
		return *new(IBillingManagerWorkflowExecutionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBillingManagerWorkflowExecutionInfo)).(*IBillingManagerWorkflowExecutionInfo)

	return out0, err

}

// GetWorkflowExecutionInfo is a free data retrieval call binding the contract method 0xef4f272d.
//
// Solidity: function getWorkflowExecutionInfo(uint256 _workflowExecutionId) view returns((uint256,uint256,uint256,address,uint8))
func (_BillingManager *BillingManagerSession) GetWorkflowExecutionInfo(_workflowExecutionId *big.Int) (IBillingManagerWorkflowExecutionInfo, error) {
	return _BillingManager.Contract.GetWorkflowExecutionInfo(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetWorkflowExecutionInfo is a free data retrieval call binding the contract method 0xef4f272d.
//
// Solidity: function getWorkflowExecutionInfo(uint256 _workflowExecutionId) view returns((uint256,uint256,uint256,address,uint8))
func (_BillingManager *BillingManagerCallerSession) GetWorkflowExecutionInfo(_workflowExecutionId *big.Int) (IBillingManagerWorkflowExecutionInfo, error) {
	return _BillingManager.Contract.GetWorkflowExecutionInfo(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetWorkflowExecutionOwner is a free data retrieval call binding the contract method 0x6e136778.
//
// Solidity: function getWorkflowExecutionOwner(uint256 _workflowExecutionId) view returns(address)
func (_BillingManager *BillingManagerCaller) GetWorkflowExecutionOwner(opts *bind.CallOpts, _workflowExecutionId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getWorkflowExecutionOwner", _workflowExecutionId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWorkflowExecutionOwner is a free data retrieval call binding the contract method 0x6e136778.
//
// Solidity: function getWorkflowExecutionOwner(uint256 _workflowExecutionId) view returns(address)
func (_BillingManager *BillingManagerSession) GetWorkflowExecutionOwner(_workflowExecutionId *big.Int) (common.Address, error) {
	return _BillingManager.Contract.GetWorkflowExecutionOwner(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetWorkflowExecutionOwner is a free data retrieval call binding the contract method 0x6e136778.
//
// Solidity: function getWorkflowExecutionOwner(uint256 _workflowExecutionId) view returns(address)
func (_BillingManager *BillingManagerCallerSession) GetWorkflowExecutionOwner(_workflowExecutionId *big.Int) (common.Address, error) {
	return _BillingManager.Contract.GetWorkflowExecutionOwner(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetWorkflowExecutionStatus is a free data retrieval call binding the contract method 0x7105e970.
//
// Solidity: function getWorkflowExecutionStatus(uint256 _workflowExecutionId) view returns(uint8)
func (_BillingManager *BillingManagerCaller) GetWorkflowExecutionStatus(opts *bind.CallOpts, _workflowExecutionId *big.Int) (uint8, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getWorkflowExecutionStatus", _workflowExecutionId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWorkflowExecutionStatus is a free data retrieval call binding the contract method 0x7105e970.
//
// Solidity: function getWorkflowExecutionStatus(uint256 _workflowExecutionId) view returns(uint8)
func (_BillingManager *BillingManagerSession) GetWorkflowExecutionStatus(_workflowExecutionId *big.Int) (uint8, error) {
	return _BillingManager.Contract.GetWorkflowExecutionStatus(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetWorkflowExecutionStatus is a free data retrieval call binding the contract method 0x7105e970.
//
// Solidity: function getWorkflowExecutionStatus(uint256 _workflowExecutionId) view returns(uint8)
func (_BillingManager *BillingManagerCallerSession) GetWorkflowExecutionStatus(_workflowExecutionId *big.Int) (uint8, error) {
	return _BillingManager.Contract.GetWorkflowExecutionStatus(&_BillingManager.CallOpts, _workflowExecutionId)
}

// NetworkRewards is a free data retrieval call binding the contract method 0x6b5d4206.
//
// Solidity: function networkRewards() view returns(uint256)
func (_BillingManager *BillingManagerCaller) NetworkRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "networkRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NetworkRewards is a free data retrieval call binding the contract method 0x6b5d4206.
//
// Solidity: function networkRewards() view returns(uint256)
func (_BillingManager *BillingManagerSession) NetworkRewards() (*big.Int, error) {
	return _BillingManager.Contract.NetworkRewards(&_BillingManager.CallOpts)
}

// NetworkRewards is a free data retrieval call binding the contract method 0x6b5d4206.
//
// Solidity: function networkRewards() view returns(uint256)
func (_BillingManager *BillingManagerCallerSession) NetworkRewards() (*big.Int, error) {
	return _BillingManager.Contract.NetworkRewards(&_BillingManager.CallOpts)
}

// NextWorkflowExecutionId is a free data retrieval call binding the contract method 0x24293941.
//
// Solidity: function nextWorkflowExecutionId() view returns(uint256)
func (_BillingManager *BillingManagerCaller) NextWorkflowExecutionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "nextWorkflowExecutionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextWorkflowExecutionId is a free data retrieval call binding the contract method 0x24293941.
//
// Solidity: function nextWorkflowExecutionId() view returns(uint256)
func (_BillingManager *BillingManagerSession) NextWorkflowExecutionId() (*big.Int, error) {
	return _BillingManager.Contract.NextWorkflowExecutionId(&_BillingManager.CallOpts)
}

// NextWorkflowExecutionId is a free data retrieval call binding the contract method 0x24293941.
//
// Solidity: function nextWorkflowExecutionId() view returns(uint256)
func (_BillingManager *BillingManagerCallerSession) NextWorkflowExecutionId() (*big.Int, error) {
	return _BillingManager.Contract.NextWorkflowExecutionId(&_BillingManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BillingManager *BillingManagerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BillingManager *BillingManagerSession) Registry() (common.Address, error) {
	return _BillingManager.Contract.Registry(&_BillingManager.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_BillingManager *BillingManagerCallerSession) Registry() (common.Address, error) {
	return _BillingManager.Contract.Registry(&_BillingManager.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_BillingManager *BillingManagerCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_BillingManager *BillingManagerSession) SignerGetter() (common.Address, error) {
	return _BillingManager.Contract.SignerGetter(&_BillingManager.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_BillingManager *BillingManagerCallerSession) SignerGetter() (common.Address, error) {
	return _BillingManager.Contract.SignerGetter(&_BillingManager.CallOpts)
}

// CompleteExecution is a paid mutator transaction binding the contract method 0x80017563.
//
// Solidity: function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) returns()
func (_BillingManager *BillingManagerTransactor) CompleteExecution(opts *bind.TransactOpts, _workflowExecutionId *big.Int, _executionAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "completeExecution", _workflowExecutionId, _executionAmount)
}

// CompleteExecution is a paid mutator transaction binding the contract method 0x80017563.
//
// Solidity: function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) returns()
func (_BillingManager *BillingManagerSession) CompleteExecution(_workflowExecutionId *big.Int, _executionAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.CompleteExecution(&_BillingManager.TransactOpts, _workflowExecutionId, _executionAmount)
}

// CompleteExecution is a paid mutator transaction binding the contract method 0x80017563.
//
// Solidity: function completeExecution(uint256 _workflowExecutionId, uint256 _executionAmount) returns()
func (_BillingManager *BillingManagerTransactorSession) CompleteExecution(_workflowExecutionId *big.Int, _executionAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.CompleteExecution(&_BillingManager.TransactOpts, _workflowExecutionId, _executionAmount)
}

// FundBalance is a paid mutator transaction binding the contract method 0x0ad5a865.
//
// Solidity: function fundBalance(address _recipientAddr) payable returns()
func (_BillingManager *BillingManagerTransactor) FundBalance(opts *bind.TransactOpts, _recipientAddr common.Address) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "fundBalance", _recipientAddr)
}

// FundBalance is a paid mutator transaction binding the contract method 0x0ad5a865.
//
// Solidity: function fundBalance(address _recipientAddr) payable returns()
func (_BillingManager *BillingManagerSession) FundBalance(_recipientAddr common.Address) (*types.Transaction, error) {
	return _BillingManager.Contract.FundBalance(&_BillingManager.TransactOpts, _recipientAddr)
}

// FundBalance is a paid mutator transaction binding the contract method 0x0ad5a865.
//
// Solidity: function fundBalance(address _recipientAddr) payable returns()
func (_BillingManager *BillingManagerTransactorSession) FundBalance(_recipientAddr common.Address) (*types.Transaction, error) {
	return _BillingManager.Contract.FundBalance(&_BillingManager.TransactOpts, _recipientAddr)
}

// FundBalance0 is a paid mutator transaction binding the contract method 0x3c067945.
//
// Solidity: function fundBalance() payable returns()
func (_BillingManager *BillingManagerTransactor) FundBalance0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "fundBalance0")
}

// FundBalance0 is a paid mutator transaction binding the contract method 0x3c067945.
//
// Solidity: function fundBalance() payable returns()
func (_BillingManager *BillingManagerSession) FundBalance0() (*types.Transaction, error) {
	return _BillingManager.Contract.FundBalance0(&_BillingManager.TransactOpts)
}

// FundBalance0 is a paid mutator transaction binding the contract method 0x3c067945.
//
// Solidity: function fundBalance() payable returns()
func (_BillingManager *BillingManagerTransactorSession) FundBalance0() (*types.Transaction, error) {
	return _BillingManager.Contract.FundBalance0(&_BillingManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _signerGetterAddress) returns()
func (_BillingManager *BillingManagerTransactor) Initialize(opts *bind.TransactOpts, _registryAddr common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "initialize", _registryAddr, _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _signerGetterAddress) returns()
func (_BillingManager *BillingManagerSession) Initialize(_registryAddr common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _BillingManager.Contract.Initialize(&_BillingManager.TransactOpts, _registryAddr, _signerGetterAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _registryAddr, address _signerGetterAddress) returns()
func (_BillingManager *BillingManagerTransactorSession) Initialize(_registryAddr common.Address, _signerGetterAddress common.Address) (*types.Transaction, error) {
	return _BillingManager.Contract.Initialize(&_BillingManager.TransactOpts, _registryAddr, _signerGetterAddress)
}

// LockExecutionFunds is a paid mutator transaction binding the contract method 0xf6586812.
//
// Solidity: function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) returns()
func (_BillingManager *BillingManagerTransactor) LockExecutionFunds(opts *bind.TransactOpts, _workflowId *big.Int, _executionLockedAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "lockExecutionFunds", _workflowId, _executionLockedAmount)
}

// LockExecutionFunds is a paid mutator transaction binding the contract method 0xf6586812.
//
// Solidity: function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) returns()
func (_BillingManager *BillingManagerSession) LockExecutionFunds(_workflowId *big.Int, _executionLockedAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.LockExecutionFunds(&_BillingManager.TransactOpts, _workflowId, _executionLockedAmount)
}

// LockExecutionFunds is a paid mutator transaction binding the contract method 0xf6586812.
//
// Solidity: function lockExecutionFunds(uint256 _workflowId, uint256 _executionLockedAmount) returns()
func (_BillingManager *BillingManagerTransactorSession) LockExecutionFunds(_workflowId *big.Int, _executionLockedAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.LockExecutionFunds(&_BillingManager.TransactOpts, _workflowId, _executionLockedAmount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _amountToWithdraw) returns()
func (_BillingManager *BillingManagerTransactor) WithdrawFunds(opts *bind.TransactOpts, _amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "withdrawFunds", _amountToWithdraw)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _amountToWithdraw) returns()
func (_BillingManager *BillingManagerSession) WithdrawFunds(_amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawFunds(&_BillingManager.TransactOpts, _amountToWithdraw)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x155dd5ee.
//
// Solidity: function withdrawFunds(uint256 _amountToWithdraw) returns()
func (_BillingManager *BillingManagerTransactorSession) WithdrawFunds(_amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawFunds(&_BillingManager.TransactOpts, _amountToWithdraw)
}

// WithdrawNetworkRewards is a paid mutator transaction binding the contract method 0xf9b41fd5.
//
// Solidity: function withdrawNetworkRewards() returns()
func (_BillingManager *BillingManagerTransactor) WithdrawNetworkRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "withdrawNetworkRewards")
}

// WithdrawNetworkRewards is a paid mutator transaction binding the contract method 0xf9b41fd5.
//
// Solidity: function withdrawNetworkRewards() returns()
func (_BillingManager *BillingManagerSession) WithdrawNetworkRewards() (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawNetworkRewards(&_BillingManager.TransactOpts)
}

// WithdrawNetworkRewards is a paid mutator transaction binding the contract method 0xf9b41fd5.
//
// Solidity: function withdrawNetworkRewards() returns()
func (_BillingManager *BillingManagerTransactorSession) WithdrawNetworkRewards() (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawNetworkRewards(&_BillingManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BillingManager *BillingManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BillingManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BillingManager *BillingManagerSession) Receive() (*types.Transaction, error) {
	return _BillingManager.Contract.Receive(&_BillingManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BillingManager *BillingManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _BillingManager.Contract.Receive(&_BillingManager.TransactOpts)
}

// BillingManagerBalanceFundedIterator is returned from FilterBalanceFunded and is used to iterate over the raw logs and unpacked data for BalanceFunded events raised by the BillingManager contract.
type BillingManagerBalanceFundedIterator struct {
	Event *BillingManagerBalanceFunded // Event containing the contract specifics and raw log

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
func (it *BillingManagerBalanceFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerBalanceFunded)
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
		it.Event = new(BillingManagerBalanceFunded)
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
func (it *BillingManagerBalanceFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerBalanceFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerBalanceFunded represents a BalanceFunded event raised by the BillingManager contract.
type BillingManagerBalanceFunded struct {
	UserAddr           common.Address
	UserCurrentBalance *big.Int
	FundedAmount       *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBalanceFunded is a free log retrieval operation binding the contract event 0x3be325f4f5fcedd7332727df6cd7ca2679dbd8e6a6332f45e0faeeb23b7d54a1.
//
// Solidity: event BalanceFunded(address indexed userAddr, uint256 userCurrentBalance, uint256 fundedAmount)
func (_BillingManager *BillingManagerFilterer) FilterBalanceFunded(opts *bind.FilterOpts, userAddr []common.Address) (*BillingManagerBalanceFundedIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "BalanceFunded", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerBalanceFundedIterator{contract: _BillingManager.contract, event: "BalanceFunded", logs: logs, sub: sub}, nil
}

// WatchBalanceFunded is a free log subscription operation binding the contract event 0x3be325f4f5fcedd7332727df6cd7ca2679dbd8e6a6332f45e0faeeb23b7d54a1.
//
// Solidity: event BalanceFunded(address indexed userAddr, uint256 userCurrentBalance, uint256 fundedAmount)
func (_BillingManager *BillingManagerFilterer) WatchBalanceFunded(opts *bind.WatchOpts, sink chan<- *BillingManagerBalanceFunded, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "BalanceFunded", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerBalanceFunded)
				if err := _BillingManager.contract.UnpackLog(event, "BalanceFunded", log); err != nil {
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

// ParseBalanceFunded is a log parse operation binding the contract event 0x3be325f4f5fcedd7332727df6cd7ca2679dbd8e6a6332f45e0faeeb23b7d54a1.
//
// Solidity: event BalanceFunded(address indexed userAddr, uint256 userCurrentBalance, uint256 fundedAmount)
func (_BillingManager *BillingManagerFilterer) ParseBalanceFunded(log types.Log) (*BillingManagerBalanceFunded, error) {
	event := new(BillingManagerBalanceFunded)
	if err := _BillingManager.contract.UnpackLog(event, "BalanceFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerExecutionCompletedIterator is returned from FilterExecutionCompleted and is used to iterate over the raw logs and unpacked data for ExecutionCompleted events raised by the BillingManager contract.
type BillingManagerExecutionCompletedIterator struct {
	Event *BillingManagerExecutionCompleted // Event containing the contract specifics and raw log

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
func (it *BillingManagerExecutionCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerExecutionCompleted)
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
		it.Event = new(BillingManagerExecutionCompleted)
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
func (it *BillingManagerExecutionCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerExecutionCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerExecutionCompleted represents a ExecutionCompleted event raised by the BillingManager contract.
type BillingManagerExecutionCompleted struct {
	WorkflowExecutionId *big.Int
	ExecutionAmount     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterExecutionCompleted is a free log retrieval operation binding the contract event 0x70114d139c12c9f15fb8a74cc0b9d9f2070f048d32cdf038844425444bad7b52.
//
// Solidity: event ExecutionCompleted(uint256 workflowExecutionId, uint256 executionAmount)
func (_BillingManager *BillingManagerFilterer) FilterExecutionCompleted(opts *bind.FilterOpts) (*BillingManagerExecutionCompletedIterator, error) {

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "ExecutionCompleted")
	if err != nil {
		return nil, err
	}
	return &BillingManagerExecutionCompletedIterator{contract: _BillingManager.contract, event: "ExecutionCompleted", logs: logs, sub: sub}, nil
}

// WatchExecutionCompleted is a free log subscription operation binding the contract event 0x70114d139c12c9f15fb8a74cc0b9d9f2070f048d32cdf038844425444bad7b52.
//
// Solidity: event ExecutionCompleted(uint256 workflowExecutionId, uint256 executionAmount)
func (_BillingManager *BillingManagerFilterer) WatchExecutionCompleted(opts *bind.WatchOpts, sink chan<- *BillingManagerExecutionCompleted) (event.Subscription, error) {

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "ExecutionCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerExecutionCompleted)
				if err := _BillingManager.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
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

// ParseExecutionCompleted is a log parse operation binding the contract event 0x70114d139c12c9f15fb8a74cc0b9d9f2070f048d32cdf038844425444bad7b52.
//
// Solidity: event ExecutionCompleted(uint256 workflowExecutionId, uint256 executionAmount)
func (_BillingManager *BillingManagerFilterer) ParseExecutionCompleted(log types.Log) (*BillingManagerExecutionCompleted, error) {
	event := new(BillingManagerExecutionCompleted)
	if err := _BillingManager.contract.UnpackLog(event, "ExecutionCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerExecutionFundsLockedIterator is returned from FilterExecutionFundsLocked and is used to iterate over the raw logs and unpacked data for ExecutionFundsLocked events raised by the BillingManager contract.
type BillingManagerExecutionFundsLockedIterator struct {
	Event *BillingManagerExecutionFundsLocked // Event containing the contract specifics and raw log

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
func (it *BillingManagerExecutionFundsLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerExecutionFundsLocked)
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
		it.Event = new(BillingManagerExecutionFundsLocked)
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
func (it *BillingManagerExecutionFundsLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerExecutionFundsLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerExecutionFundsLocked represents a ExecutionFundsLocked event raised by the BillingManager contract.
type BillingManagerExecutionFundsLocked struct {
	WorkflowId            *big.Int
	UserAddr              common.Address
	WorkflowExecutionId   *big.Int
	ExecutionLockedAmount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecutionFundsLocked is a free log retrieval operation binding the contract event 0xc77a14be8053ff3df25436e566ff7bd5d0da84a96e23ced50e403617fab5e97d.
//
// Solidity: event ExecutionFundsLocked(uint256 indexed workflowId, address indexed userAddr, uint256 workflowExecutionId, uint256 executionLockedAmount)
func (_BillingManager *BillingManagerFilterer) FilterExecutionFundsLocked(opts *bind.FilterOpts, workflowId []*big.Int, userAddr []common.Address) (*BillingManagerExecutionFundsLockedIterator, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "ExecutionFundsLocked", workflowIdRule, userAddrRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerExecutionFundsLockedIterator{contract: _BillingManager.contract, event: "ExecutionFundsLocked", logs: logs, sub: sub}, nil
}

// WatchExecutionFundsLocked is a free log subscription operation binding the contract event 0xc77a14be8053ff3df25436e566ff7bd5d0da84a96e23ced50e403617fab5e97d.
//
// Solidity: event ExecutionFundsLocked(uint256 indexed workflowId, address indexed userAddr, uint256 workflowExecutionId, uint256 executionLockedAmount)
func (_BillingManager *BillingManagerFilterer) WatchExecutionFundsLocked(opts *bind.WatchOpts, sink chan<- *BillingManagerExecutionFundsLocked, workflowId []*big.Int, userAddr []common.Address) (event.Subscription, error) {

	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "ExecutionFundsLocked", workflowIdRule, userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerExecutionFundsLocked)
				if err := _BillingManager.contract.UnpackLog(event, "ExecutionFundsLocked", log); err != nil {
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

// ParseExecutionFundsLocked is a log parse operation binding the contract event 0xc77a14be8053ff3df25436e566ff7bd5d0da84a96e23ced50e403617fab5e97d.
//
// Solidity: event ExecutionFundsLocked(uint256 indexed workflowId, address indexed userAddr, uint256 workflowExecutionId, uint256 executionLockedAmount)
func (_BillingManager *BillingManagerFilterer) ParseExecutionFundsLocked(log types.Log) (*BillingManagerExecutionFundsLocked, error) {
	event := new(BillingManagerExecutionFundsLocked)
	if err := _BillingManager.contract.UnpackLog(event, "ExecutionFundsLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BillingManager contract.
type BillingManagerInitializedIterator struct {
	Event *BillingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *BillingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerInitialized)
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
		it.Event = new(BillingManagerInitialized)
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
func (it *BillingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerInitialized represents a Initialized event raised by the BillingManager contract.
type BillingManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BillingManager *BillingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*BillingManagerInitializedIterator, error) {

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BillingManagerInitializedIterator{contract: _BillingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BillingManager *BillingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BillingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerInitialized)
				if err := _BillingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_BillingManager *BillingManagerFilterer) ParseInitialized(log types.Log) (*BillingManagerInitialized, error) {
	event := new(BillingManagerInitialized)
	if err := _BillingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerRewardsWithdrawnIterator is returned from FilterRewardsWithdrawn and is used to iterate over the raw logs and unpacked data for RewardsWithdrawn events raised by the BillingManager contract.
type BillingManagerRewardsWithdrawnIterator struct {
	Event *BillingManagerRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *BillingManagerRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerRewardsWithdrawn)
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
		it.Event = new(BillingManagerRewardsWithdrawn)
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
func (it *BillingManagerRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerRewardsWithdrawn represents a RewardsWithdrawn event raised by the BillingManager contract.
type BillingManagerRewardsWithdrawn struct {
	RecipientAddr common.Address
	RewardsAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRewardsWithdrawn is a free log retrieval operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address indexed recipientAddr, uint256 rewardsAmount)
func (_BillingManager *BillingManagerFilterer) FilterRewardsWithdrawn(opts *bind.FilterOpts, recipientAddr []common.Address) (*BillingManagerRewardsWithdrawnIterator, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "RewardsWithdrawn", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerRewardsWithdrawnIterator{contract: _BillingManager.contract, event: "RewardsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRewardsWithdrawn is a free log subscription operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address indexed recipientAddr, uint256 rewardsAmount)
func (_BillingManager *BillingManagerFilterer) WatchRewardsWithdrawn(opts *bind.WatchOpts, sink chan<- *BillingManagerRewardsWithdrawn, recipientAddr []common.Address) (event.Subscription, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "RewardsWithdrawn", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerRewardsWithdrawn)
				if err := _BillingManager.contract.UnpackLog(event, "RewardsWithdrawn", log); err != nil {
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

// ParseRewardsWithdrawn is a log parse operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address indexed recipientAddr, uint256 rewardsAmount)
func (_BillingManager *BillingManagerFilterer) ParseRewardsWithdrawn(log types.Log) (*BillingManagerRewardsWithdrawn, error) {
	event := new(BillingManagerRewardsWithdrawn)
	if err := _BillingManager.contract.UnpackLog(event, "RewardsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerUnexpectedExecutionAmountFoundIterator is returned from FilterUnexpectedExecutionAmountFound and is used to iterate over the raw logs and unpacked data for UnexpectedExecutionAmountFound events raised by the BillingManager contract.
type BillingManagerUnexpectedExecutionAmountFoundIterator struct {
	Event *BillingManagerUnexpectedExecutionAmountFound // Event containing the contract specifics and raw log

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
func (it *BillingManagerUnexpectedExecutionAmountFoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerUnexpectedExecutionAmountFound)
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
		it.Event = new(BillingManagerUnexpectedExecutionAmountFound)
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
func (it *BillingManagerUnexpectedExecutionAmountFoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerUnexpectedExecutionAmountFoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerUnexpectedExecutionAmountFound represents a UnexpectedExecutionAmountFound event raised by the BillingManager contract.
type BillingManagerUnexpectedExecutionAmountFound struct {
	WorkflowExecutionId   *big.Int
	ExecutionLockedAmount *big.Int
	ExecutionAmount       *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedExecutionAmountFound is a free log retrieval operation binding the contract event 0xaba1d671589f510878d4421f3e9f5574d37ecf6735d3f3df2be569b62b9015d9.
//
// Solidity: event UnexpectedExecutionAmountFound(uint256 workflowExecutionId, uint256 executionLockedAmount, uint256 executionAmount)
func (_BillingManager *BillingManagerFilterer) FilterUnexpectedExecutionAmountFound(opts *bind.FilterOpts) (*BillingManagerUnexpectedExecutionAmountFoundIterator, error) {

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "UnexpectedExecutionAmountFound")
	if err != nil {
		return nil, err
	}
	return &BillingManagerUnexpectedExecutionAmountFoundIterator{contract: _BillingManager.contract, event: "UnexpectedExecutionAmountFound", logs: logs, sub: sub}, nil
}

// WatchUnexpectedExecutionAmountFound is a free log subscription operation binding the contract event 0xaba1d671589f510878d4421f3e9f5574d37ecf6735d3f3df2be569b62b9015d9.
//
// Solidity: event UnexpectedExecutionAmountFound(uint256 workflowExecutionId, uint256 executionLockedAmount, uint256 executionAmount)
func (_BillingManager *BillingManagerFilterer) WatchUnexpectedExecutionAmountFound(opts *bind.WatchOpts, sink chan<- *BillingManagerUnexpectedExecutionAmountFound) (event.Subscription, error) {

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "UnexpectedExecutionAmountFound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerUnexpectedExecutionAmountFound)
				if err := _BillingManager.contract.UnpackLog(event, "UnexpectedExecutionAmountFound", log); err != nil {
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

// ParseUnexpectedExecutionAmountFound is a log parse operation binding the contract event 0xaba1d671589f510878d4421f3e9f5574d37ecf6735d3f3df2be569b62b9015d9.
//
// Solidity: event UnexpectedExecutionAmountFound(uint256 workflowExecutionId, uint256 executionLockedAmount, uint256 executionAmount)
func (_BillingManager *BillingManagerFilterer) ParseUnexpectedExecutionAmountFound(log types.Log) (*BillingManagerUnexpectedExecutionAmountFound, error) {
	event := new(BillingManagerUnexpectedExecutionAmountFound)
	if err := _BillingManager.contract.UnpackLog(event, "UnexpectedExecutionAmountFound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerUserFundsWithdrawnIterator is returned from FilterUserFundsWithdrawn and is used to iterate over the raw logs and unpacked data for UserFundsWithdrawn events raised by the BillingManager contract.
type BillingManagerUserFundsWithdrawnIterator struct {
	Event *BillingManagerUserFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *BillingManagerUserFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerUserFundsWithdrawn)
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
		it.Event = new(BillingManagerUserFundsWithdrawn)
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
func (it *BillingManagerUserFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerUserFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerUserFundsWithdrawn represents a UserFundsWithdrawn event raised by the BillingManager contract.
type BillingManagerUserFundsWithdrawn struct {
	UserAddr        common.Address
	WithdrawnAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUserFundsWithdrawn is a free log retrieval operation binding the contract event 0x7f9d7c82d89883dcf97d377c4a3098b9a25d87d6d25dc4b289531d7a38460a08.
//
// Solidity: event UserFundsWithdrawn(address indexed userAddr, uint256 withdrawnAmount)
func (_BillingManager *BillingManagerFilterer) FilterUserFundsWithdrawn(opts *bind.FilterOpts, userAddr []common.Address) (*BillingManagerUserFundsWithdrawnIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "UserFundsWithdrawn", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerUserFundsWithdrawnIterator{contract: _BillingManager.contract, event: "UserFundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchUserFundsWithdrawn is a free log subscription operation binding the contract event 0x7f9d7c82d89883dcf97d377c4a3098b9a25d87d6d25dc4b289531d7a38460a08.
//
// Solidity: event UserFundsWithdrawn(address indexed userAddr, uint256 withdrawnAmount)
func (_BillingManager *BillingManagerFilterer) WatchUserFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *BillingManagerUserFundsWithdrawn, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "UserFundsWithdrawn", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerUserFundsWithdrawn)
				if err := _BillingManager.contract.UnpackLog(event, "UserFundsWithdrawn", log); err != nil {
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

// ParseUserFundsWithdrawn is a log parse operation binding the contract event 0x7f9d7c82d89883dcf97d377c4a3098b9a25d87d6d25dc4b289531d7a38460a08.
//
// Solidity: event UserFundsWithdrawn(address indexed userAddr, uint256 withdrawnAmount)
func (_BillingManager *BillingManagerFilterer) ParseUserFundsWithdrawn(log types.Log) (*BillingManagerUserFundsWithdrawn, error) {
	event := new(BillingManagerUserFundsWithdrawn)
	if err := _BillingManager.contract.UnpackLog(event, "UserFundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
