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
	UserAddr                    common.Address
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userCurrentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundedAmount\",\"type\":\"uint256\"}],\"name\":\"BalanceFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"}],\"name\":\"ExecutionFundsLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"RewardsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedExecutionAmountFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"name\":\"UserFundsWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_executionAmount\",\"type\":\"uint256\"}],\"name\":\"completeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"}],\"name\":\"fundBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundBalance\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getExecutionWorkflowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getExistingUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalUsersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getUserAvailableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getUserFundsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userFundBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLockedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingWorkflowExecutionIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIBillingManager.UserFundsInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getUsersFundsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userFundBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLockedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingWorkflowExecutionIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIBillingManager.UserFundsInfo[]\",\"name\":\"_usersInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"enumIBillingManager.WorkflowExecutionStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIBillingManager.WorkflowExecutionInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionStatus\",\"outputs\":[{\"internalType\":\"enumIBillingManager.WorkflowExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_executionLockedAmount\",\"type\":\"uint256\"}],\"name\":\"lockExecutionFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextWorkflowExecutionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNetworkRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5060805161237561004c60003960008181610793015281816107d3015281816109e701528181610a270152610aba01526123756000f3fe6080604052600436106101445760003560e01c806352d1902d116100b65780637b1039991161006f5780637b103999146103c257806380017563146103e2578063c4cffb7614610402578063ef4f272d14610422578063f65868121461044f578063f9b41fd51461046f57600080fd5b806352d1902d146102ac5780635c211f88146102c15780636b091226146102ff5780636b5d42061461032c5780636e136778146103425780637105e9701461037b57600080fd5b80633659cfe6116101085780633659cfe6146101f75780633c06794514610217578063485cc9551461021f57806348cf8bab1461023f5780634a686e491461026c5780634f1ef2861461029957600080fd5b80630561010c146101595780630ad5a86514610181578063155dd5ee1461019457806324293941146101b45780632cdf71a7146101ca57600080fd5b366101545761015233610484565b005b600080fd5b34801561016557600080fd5b5061016e610578565b6040519081526020015b60405180910390f35b61015261018f366004611db7565b610589565b3480156101a057600080fd5b506101526101af366004611dd4565b610595565b3480156101c057600080fd5b5061016e60035481565b3480156101d657600080fd5b506101ea6101e5366004611ded565b61068c565b6040516101789190611e80565b34801561020357600080fd5b50610152610212366004611db7565b610789565b610152610865565b34801561022b57600080fd5b5061015261023a366004611ee2565b610870565b34801561024b57600080fd5b5061025f61025a366004611ded565b6109c6565b6040516101789190611f1b565b34801561027857600080fd5b5061016e610287366004611dd4565b60009081526007602052604090205490565b6101526102a7366004611f7e565b6109dd565b3480156102b857600080fd5b5061016e610aad565b3480156102cd57600080fd5b506000546102e7906201000090046001600160a01b031681565b6040516001600160a01b039091168152602001610178565b34801561030b57600080fd5b5061031f61031a366004611db7565b610b60565b6040516101789190612042565b34801561033857600080fd5b5061016e60025481565b34801561034e57600080fd5b506102e761035d366004611dd4565b6000908152600760205260409020600301546001600160a01b031690565b34801561038757600080fd5b506103b5610396366004611dd4565b600090815260076020526040902060030154600160a01b900460ff1690565b604051610178919061208d565b3480156103ce57600080fd5b506001546102e7906001600160a01b031681565b3480156103ee57600080fd5b506101526103fd366004611ded565b610bea565b34801561040e57600080fd5b5061016e61041d366004611db7565b610ed3565b34801561042e57600080fd5b5061044261043d366004611dd4565b610efc565b604051610178919061209b565b34801561045b57600080fd5b5061015261046a366004611ded565b610f8d565b34801561047b57600080fd5b50610152611332565b600034116104e35760405162461bcd60e51b815260206004820152602160248201527f42696c6c696e674d616e616765723a205a65726f2066756e647320746f2061646044820152601960fa1b60648201526084015b60405180910390fd5b6001600160a01b0381166000908152600660205260408120546105079034906120f3565b6001600160a01b0383166000908152600660205260409020819055905061052f6004836114c0565b50604080518281523460208201526001600160a01b038416917f3be325f4f5fcedd7332727df6cd7ca2679dbd8e6a6332f45e0faeeb23b7d54a191015b60405180910390a25050565b600061058460046114d5565b905090565b61059281610484565b50565b8061059f33610ed3565b101561060c5760405162461bcd60e51b815260206004820152603660248201527f42696c6c696e674d616e616765723a204e6f7420656e6f75676820617661696c60448201527561626c652066756e647320746f20776974686472617760501b60648201526084016104da565b33600090815260066020526040812054610627908390612106565b3360009081526006602052604081208290559091508190036106505761064e6004336114df565b505b61065a33836114f4565b60405182815233907f7f9d7c82d89883dcf97d377c4a3098b9a25d87d6d25dc4b289531d7a38460a089060200161056c565b606060006106a461069d60046114d5565b858561160d565b90506106b08482612106565b67ffffffffffffffff8111156106c8576106c8611f68565b60405190808252806020026020018201604052801561072d57816020015b61071a604051806080016040528060006001600160a01b031681526020016000815260200160008152602001606081525090565b8152602001906001900390816106e65790505b509150835b818110156107815761074861031a600483611638565b836107538784612106565b8151811061076357610763612119565b602002602001018190525080806107799061212f565b915050610732565b505092915050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036107d15760405162461bcd60e51b81526004016104da90612148565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661081a6000805160206122f9833981519152546001600160a01b031690565b6001600160a01b0316146108405760405162461bcd60e51b81526004016104da90612194565b61084981611644565b60408051600080825260208201909252610592918391906116dd565b61086e33610484565b565b600054610100900460ff16158080156108905750600054600160ff909116105b806108aa5750303b1580156108aa575060005460ff166001145b61090d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104da565b6000805460ff191660011790558015610930576000805461ff0019166101001790555b600180546001600160a01b0385166001600160a01b031990911617905561097b82600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b80156109c1576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60606109d460048484611848565b90505b92915050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610a255760405162461bcd60e51b81526004016104da90612148565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610a6e6000805160206122f9833981519152546001600160a01b031690565b6001600160a01b031614610a945760405162461bcd60e51b81526004016104da90612194565b610a9d82611644565b610aa9828260016116dd565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610b4d5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016104da565b506000805160206122f983398151915290565b610b94604051806080016040528060006001600160a01b031681526020016000815260200160008152602001606081525090565b6001600160a01b03821660008181526006602090815260409182902082516080810184529384528054918401919091526001810154918301919091529060608101610be16002840161190a565b90529392505050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015610c39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5d91906121e0565b6001600160a01b031614610c835760405162461bcd60e51b81526004016104da906121fd565b600082815260076020526040902060016003820154600160a01b900460ff166002811115610cb357610cb3612055565b14610d195760405162461bcd60e51b815260206004820152603060248201527f42696c6c696e674d616e616765723a204e6f7420612070656e64696e6720776f60448201526f3935b33637bb9032bc32b1baba34b7b760811b60648201526084016104da565b60038101546001600160a01b03166000908152600660205260409020600182015483811015610daf57610d6a848284600101548560000154610d5b9190612106565b610d6591906120f3565b611917565b60408051878152602081018490529081018290529094507faba1d671589f510878d4421f3e9f5574d37ecf6735d3f3df2be569b62b9015d99060600160405180910390a15b60038301805460ff60a01b1916600160a11b17905560028301849055815484908390600090610ddf908490612106565b9250508190555080826001016000828254610dfa9190612106565b90915550610e0d9050600283018661192d565b508360026000828254610e2091906120f3565b90915550506001548354604051632851165f60e21b81526001600160a01b039092169163a144597c91610e60918890600401918252602082015260400190565b600060405180830381600087803b158015610e7a57600080fd5b505af1158015610e8e573d6000803e3d6000fd5b505060408051888152602081018890527f70114d139c12c9f15fb8a74cc0b9d9f2070f048d32cdf038844425444bad7b52935001905060405180910390a15050505050565b6001600160a01b0381166000908152600660205260408120600181015490546109d79190612106565b610f04611d5a565b600082815260076020908152604091829020825160a081018452815481526001820154928101929092526002808201549383019390935260038101546001600160a01b0381166060840152919290916080840191600160a01b90910460ff1690811115610f7357610f73612055565b6002811115610f8457610f84612055565b90525092915050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015610fdc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061100091906121e0565b6001600160a01b0316146110265760405162461bcd60e51b81526004016104da906121fd565b600154604051639b2bfaa360e01b8152600481018490526001600160a01b0390911690639b2bfaa390602401602060405180830381865afa15801561106f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110939190612234565b6110ef5760405162461bcd60e51b815260206004820152602760248201527f42696c6c696e674d616e616765723a20576f726b666c6f7720646f6573206e6f6044820152661d08195e1a5cdd60ca1b60648201526084016104da565b60015460405163d69cd27560e01b8152600481018490526000916001600160a01b03169063d69cd27590602401602060405180830381865afa158015611139573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115d91906121e0565b6001600160a01b03811660009081526006602052604090209091508261118283610ed3565b10156111e15760405162461bcd60e51b815260206004820152602860248201527f42696c6c696e674d616e616765723a204e6f7420656e6f7567682066756e647360448201526720746f206c6f636b60c01b60648201526084016104da565b60038054600091826111f28361212f565b9190505590508382600101600082825461120c91906120f3565b9091555061121f90506002830182611939565b506040518060a0016040528086815260200185815260200160008152602001846001600160a01b031681526020016001600281111561126057611260612055565b9052600082815260076020908152604091829020835181559083015160018201559082015160028083019190915560608301516003830180546001600160a01b039092166001600160a01b031983168117825560808601519391926001600160a81b0319161790600160a01b9084908111156112de576112de612055565b02179055505060408051838152602081018790526001600160a01b038616925087917fc77a14be8053ff3df25436e566ff7bd5d0da84a96e23ced50e403617fab5e97d910160405180910390a35050505050565b600254806113995760405162461bcd60e51b815260206004820152602e60248201527f42696c6c696e674d616e616765723a204e6f206e6574776f726b20726577617260448201526d647320746f20776974686472617760901b60648201526084016104da565b60026000905560008060029054906101000a90046001600160a01b03166001600160a01b0316631a296e026040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113f3573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061141791906121e0565b90506001600160a01b03811661147b5760405162461bcd60e51b815260206004820152602360248201527f42696c6c696e674d616e616765723a205a65726f207369676e6572206164647260448201526265737360e81b60648201526084016104da565b61148581836114f4565b806001600160a01b03167f8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e5031618360405161056c91815260200190565b60006109d4836001600160a01b038416611941565b60006109d7825490565b60006109d4836001600160a01b038416611990565b804710156115445760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016104da565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114611591576040519150601f19603f3d011682016040523d82523d6000602084013e611596565b606091505b50509050806109c15760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016104da565b600061161982846120f3565b9050838111156116265750825b808311156116315750815b9392505050565b60006109d48383611a8a565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015611693573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b791906121e0565b6001600160a01b0316146105925760405162461bcd60e51b81526004016104da906121fd565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615611710576109c183611ab4565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561176a575060408051601f3d908101601f1916820190925261176791810190612256565b60015b6117cd5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016104da565b6000805160206122f9833981519152811461183c5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016104da565b506109c1838383611b50565b6060600061185861069d866114d5565b90506118648482612106565b67ffffffffffffffff81111561187c5761187c611f68565b6040519080825280602002602001820160405280156118a5578160200160208202803683370190505b509150835b81811015611901576118bc8682611638565b836118c78784612106565b815181106118d7576118d7612119565b6001600160a01b0390921660209283029190910190910152806118f98161212f565b9150506118aa565b50509392505050565b6060600061163183611b7b565b600081831061192657816109d4565b5090919050565b60006109d48383611990565b60006109d483835b6000818152600183016020526040812054611988575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556109d7565b5060006109d7565b60008181526001830160205260408120548015611a795760006119b4600183612106565b85549091506000906119c890600190612106565b9050818114611a2d5760008660000182815481106119e8576119e8612119565b9060005260206000200154905080876000018481548110611a0b57611a0b612119565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611a3e57611a3e61226f565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506109d7565b60009150506109d7565b5092915050565b6000826000018281548110611aa157611aa1612119565b9060005260206000200154905092915050565b6001600160a01b0381163b611b215760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016104da565b6000805160206122f983398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b611b5983611bd7565b600082511180611b665750805b156109c157611b758383611c17565b50505050565b606081600001805480602002602001604051908101604052809291908181526020018280548015611bcb57602002820191906000526020600020905b815481526020019060010190808311611bb7575b50505050509050919050565b611be081611ab4565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606109d48383604051806060016040528060278152602001612319602791396060600080856001600160a01b031685604051611c5491906122a9565b600060405180830381855af49150503d8060008114611c8f576040519150601f19603f3d011682016040523d82523d6000602084013e611c94565b606091505b5091509150611ca586838387611caf565b9695505050505050565b60608315611d1e578251600003611d17576001600160a01b0385163b611d175760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016104da565b5081611d28565b611d288383611d30565b949350505050565b815115611d405781518083602001fd5b8060405162461bcd60e51b81526004016104da91906122c5565b6040518060a0016040528060008152602001600081526020016000815260200160006001600160a01b0316815260200160006002811115611d9d57611d9d612055565b905290565b6001600160a01b038116811461059257600080fd5b600060208284031215611dc957600080fd5b813561163181611da2565b600060208284031215611de657600080fd5b5035919050565b60008060408385031215611e0057600080fd5b50508035926020909101359150565b80516001600160a01b0316825260208082015181840152604080830151908401526060808301516080918501829052805191850182905260009290810191839060a08701905b80831015611e755784518252938301936001929092019190830190611e55565b509695505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611ed557603f19888603018452611ec3858351611e0f565b94509285019290850190600101611ea7565b5092979650505050505050565b60008060408385031215611ef557600080fd5b8235611f0081611da2565b91506020830135611f1081611da2565b809150509250929050565b6020808252825182820181905260009190848201906040850190845b81811015611f5c5783516001600160a01b031683529284019291840191600101611f37565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611f9157600080fd5b8235611f9c81611da2565b9150602083013567ffffffffffffffff80821115611fb957600080fd5b818501915085601f830112611fcd57600080fd5b813581811115611fdf57611fdf611f68565b604051601f8201601f19908116603f0116810190838211818310171561200757612007611f68565b8160405282815288602084870101111561202057600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b6020815260006109d46020830184611e0f565b634e487b7160e01b600052602160045260246000fd5b6003811061208957634e487b7160e01b600052602160045260246000fd5b9052565b602081016109d7828461206b565b8151815260208083015190820152604080830151908201526060808301516001600160a01b03169082015260808083015160a0830191611a839084018261206b565b634e487b7160e01b600052601160045260246000fd5b808201808211156109d7576109d76120dd565b818103818111156109d7576109d76120dd565b634e487b7160e01b600052603260045260246000fd5b600060018201612141576121416120dd565b5060010190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6000602082840312156121f257600080fd5b815161163181611da2565b6020808252601a908201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604082015260600190565b60006020828403121561224657600080fd5b8151801515811461163157600080fd5b60006020828403121561226857600080fd5b5051919050565b634e487b7160e01b600052603160045260246000fd5b60005b838110156122a0578181015183820152602001612288565b50506000910152565b600082516122bb818460208701612285565b9190910192915050565b60208152600082518060208401526122e4816040850160208701612285565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220e16d61f03c2c9f6580bbd0e6c24164b91432315874f300059c551725f52154c564736f6c63430008120033",
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

// GetExistingUsers is a free data retrieval call binding the contract method 0x48cf8bab.
//
// Solidity: function getExistingUsers(uint256 _offset, uint256 _limit) view returns(address[])
func (_BillingManager *BillingManagerCaller) GetExistingUsers(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getExistingUsers", _offset, _limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetExistingUsers is a free data retrieval call binding the contract method 0x48cf8bab.
//
// Solidity: function getExistingUsers(uint256 _offset, uint256 _limit) view returns(address[])
func (_BillingManager *BillingManagerSession) GetExistingUsers(_offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	return _BillingManager.Contract.GetExistingUsers(&_BillingManager.CallOpts, _offset, _limit)
}

// GetExistingUsers is a free data retrieval call binding the contract method 0x48cf8bab.
//
// Solidity: function getExistingUsers(uint256 _offset, uint256 _limit) view returns(address[])
func (_BillingManager *BillingManagerCallerSession) GetExistingUsers(_offset *big.Int, _limit *big.Int) ([]common.Address, error) {
	return _BillingManager.Contract.GetExistingUsers(&_BillingManager.CallOpts, _offset, _limit)
}

// GetTotalUsersCount is a free data retrieval call binding the contract method 0x0561010c.
//
// Solidity: function getTotalUsersCount() view returns(uint256)
func (_BillingManager *BillingManagerCaller) GetTotalUsersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getTotalUsersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalUsersCount is a free data retrieval call binding the contract method 0x0561010c.
//
// Solidity: function getTotalUsersCount() view returns(uint256)
func (_BillingManager *BillingManagerSession) GetTotalUsersCount() (*big.Int, error) {
	return _BillingManager.Contract.GetTotalUsersCount(&_BillingManager.CallOpts)
}

// GetTotalUsersCount is a free data retrieval call binding the contract method 0x0561010c.
//
// Solidity: function getTotalUsersCount() view returns(uint256)
func (_BillingManager *BillingManagerCallerSession) GetTotalUsersCount() (*big.Int, error) {
	return _BillingManager.Contract.GetTotalUsersCount(&_BillingManager.CallOpts)
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
// Solidity: function getUserFundsInfo(address _userAddr) view returns((address,uint256,uint256,uint256[]))
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
// Solidity: function getUserFundsInfo(address _userAddr) view returns((address,uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerSession) GetUserFundsInfo(_userAddr common.Address) (IBillingManagerUserFundsInfo, error) {
	return _BillingManager.Contract.GetUserFundsInfo(&_BillingManager.CallOpts, _userAddr)
}

// GetUserFundsInfo is a free data retrieval call binding the contract method 0x6b091226.
//
// Solidity: function getUserFundsInfo(address _userAddr) view returns((address,uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerCallerSession) GetUserFundsInfo(_userAddr common.Address) (IBillingManagerUserFundsInfo, error) {
	return _BillingManager.Contract.GetUserFundsInfo(&_BillingManager.CallOpts, _userAddr)
}

// GetUsersFundsInfo is a free data retrieval call binding the contract method 0x2cdf71a7.
//
// Solidity: function getUsersFundsInfo(uint256 _offset, uint256 _limit) view returns((address,uint256,uint256,uint256[])[] _usersInfoArr)
func (_BillingManager *BillingManagerCaller) GetUsersFundsInfo(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]IBillingManagerUserFundsInfo, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getUsersFundsInfo", _offset, _limit)

	if err != nil {
		return *new([]IBillingManagerUserFundsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBillingManagerUserFundsInfo)).(*[]IBillingManagerUserFundsInfo)

	return out0, err

}

// GetUsersFundsInfo is a free data retrieval call binding the contract method 0x2cdf71a7.
//
// Solidity: function getUsersFundsInfo(uint256 _offset, uint256 _limit) view returns((address,uint256,uint256,uint256[])[] _usersInfoArr)
func (_BillingManager *BillingManagerSession) GetUsersFundsInfo(_offset *big.Int, _limit *big.Int) ([]IBillingManagerUserFundsInfo, error) {
	return _BillingManager.Contract.GetUsersFundsInfo(&_BillingManager.CallOpts, _offset, _limit)
}

// GetUsersFundsInfo is a free data retrieval call binding the contract method 0x2cdf71a7.
//
// Solidity: function getUsersFundsInfo(uint256 _offset, uint256 _limit) view returns((address,uint256,uint256,uint256[])[] _usersInfoArr)
func (_BillingManager *BillingManagerCallerSession) GetUsersFundsInfo(_offset *big.Int, _limit *big.Int) ([]IBillingManagerUserFundsInfo, error) {
	return _BillingManager.Contract.GetUsersFundsInfo(&_BillingManager.CallOpts, _offset, _limit)
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

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BillingManager *BillingManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BillingManager *BillingManagerSession) ProxiableUUID() ([32]byte, error) {
	return _BillingManager.Contract.ProxiableUUID(&_BillingManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_BillingManager *BillingManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _BillingManager.Contract.ProxiableUUID(&_BillingManager.CallOpts)
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

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BillingManager *BillingManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BillingManager *BillingManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BillingManager.Contract.UpgradeTo(&_BillingManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_BillingManager *BillingManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _BillingManager.Contract.UpgradeTo(&_BillingManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BillingManager *BillingManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BillingManager *BillingManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BillingManager.Contract.UpgradeToAndCall(&_BillingManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_BillingManager *BillingManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _BillingManager.Contract.UpgradeToAndCall(&_BillingManager.TransactOpts, newImplementation, data)
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

// BillingManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the BillingManager contract.
type BillingManagerAdminChangedIterator struct {
	Event *BillingManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *BillingManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerAdminChanged)
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
		it.Event = new(BillingManagerAdminChanged)
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
func (it *BillingManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerAdminChanged represents a AdminChanged event raised by the BillingManager contract.
type BillingManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BillingManager *BillingManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*BillingManagerAdminChangedIterator, error) {

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &BillingManagerAdminChangedIterator{contract: _BillingManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BillingManager *BillingManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *BillingManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerAdminChanged)
				if err := _BillingManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_BillingManager *BillingManagerFilterer) ParseAdminChanged(log types.Log) (*BillingManagerAdminChanged, error) {
	event := new(BillingManagerAdminChanged)
	if err := _BillingManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// BillingManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the BillingManager contract.
type BillingManagerBeaconUpgradedIterator struct {
	Event *BillingManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *BillingManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerBeaconUpgraded)
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
		it.Event = new(BillingManagerBeaconUpgraded)
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
func (it *BillingManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerBeaconUpgraded represents a BeaconUpgraded event raised by the BillingManager contract.
type BillingManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BillingManager *BillingManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*BillingManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerBeaconUpgradedIterator{contract: _BillingManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BillingManager *BillingManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *BillingManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerBeaconUpgraded)
				if err := _BillingManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_BillingManager *BillingManagerFilterer) ParseBeaconUpgraded(log types.Log) (*BillingManagerBeaconUpgraded, error) {
	event := new(BillingManagerBeaconUpgraded)
	if err := _BillingManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// BillingManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the BillingManager contract.
type BillingManagerUpgradedIterator struct {
	Event *BillingManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *BillingManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerUpgraded)
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
		it.Event = new(BillingManagerUpgraded)
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
func (it *BillingManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerUpgraded represents a Upgraded event raised by the BillingManager contract.
type BillingManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BillingManager *BillingManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BillingManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerUpgradedIterator{contract: _BillingManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BillingManager *BillingManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BillingManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerUpgraded)
				if err := _BillingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_BillingManager *BillingManagerFilterer) ParseUpgraded(log types.Log) (*BillingManagerUpgraded, error) {
	event := new(BillingManagerUpgraded)
	if err := _BillingManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
