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

// IBillingManagerDepositAssetData is an auto generated low-level Go binding around an user-defined struct.
type IBillingManagerDepositAssetData struct {
	TokenAddr                 common.Address
	WorkflowExecutionDiscount *big.Int
	NetworkRewards            *big.Int
	IsPermitable              bool
	IsEnabled                 bool
}

// IBillingManagerDepositAssetInfo is an auto generated low-level Go binding around an user-defined struct.
type IBillingManagerDepositAssetInfo struct {
	DepositAssetKey  string
	DepositAssetData IBillingManagerDepositAssetData
}

// IBillingManagerUserDepositInfo is an auto generated low-level Go binding around an user-defined struct.
type IBillingManagerUserDepositInfo struct {
	UserAddr                    common.Address
	UserDepositedAmount         *big.Int
	UserLockedAmount            *big.Int
	PendingWorkflowExecutionIds []*big.Int
}

// IBillingManagerWorkflowExecutionInfo is an auto generated low-level Go binding around an user-defined struct.
type IBillingManagerWorkflowExecutionInfo struct {
	DepositAssetKey       string
	WorkflowId            *big.Int
	ExecutionLockedAmount *big.Int
	ExecutionAmount       *big.Int
	WorkflowOwner         common.Address
	Status                uint8
}

// BillingManagerMetaData contains all meta data concerning the BillingManager contract.
var BillingManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokensSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"depositRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"AssetDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workflowExecutionDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkRewards\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPermitable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIBillingManager.DepositAssetData\",\"name\":\"depositAssetData\",\"type\":\"tuple\"}],\"name\":\"DepositAssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"newEnabledStatus\",\"type\":\"bool\"}],\"name\":\"DepositAssetEnabledStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"}],\"name\":\"ExecutionCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"}],\"name\":\"ExecutionFundsLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIBillingManager.NetworkWithdrawReasons\",\"name\":\"withdrawReason\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"NetworkWithdrawCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsAmount\",\"type\":\"uint256\"}],\"name\":\"RewardsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedExecutionAmountFound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"name\":\"UserFundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWorkflowExecutionDiscount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevWorkflowExecutionDiscount\",\"type\":\"uint256\"}],\"name\":\"WorkflowExecutionDiscountUpdated\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workflowExecutionDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkRewards\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPermitable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structIBillingManager.DepositAssetData\",\"name\":\"depositAssetData\",\"type\":\"tuple\"}],\"internalType\":\"structIBillingManager.DepositAssetInfo[]\",\"name\":\"_depositAssetInfoArr\",\"type\":\"tuple[]\"}],\"name\":\"addDepositAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_executionAmount\",\"type\":\"uint256\"}],\"name\":\"completeExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_recipientAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sigExpirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"depositWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_depositAssetKeysArr\",\"type\":\"string[]\"}],\"name\":\"getDepositAssetsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workflowExecutionDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkRewards\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPermitable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structIBillingManager.DepositAssetData\",\"name\":\"depositAssetData\",\"type\":\"tuple\"}],\"internalType\":\"structIBillingManager.DepositAssetInfo[]\",\"name\":\"_depositAssetsInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getExecutionWorkflowId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getExistingUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"getNetworkRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedDepositAssetKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalUsersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"getUserAvailableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getUserDepositAssetKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"getUserDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userDepositedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingWorkflowExecutionIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIBillingManager.UserDepositInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getUsersDepositInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userDepositedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pendingWorkflowExecutionIds\",\"type\":\"uint256[]\"}],\"internalType\":\"structIBillingManager.UserDepositInfo[]\",\"name\":\"_usersInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionLockedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"enumIBillingManager.WorkflowExecutionStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIBillingManager.WorkflowExecutionInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowExecutionStatus\",\"outputs\":[{\"internalType\":\"enumIBillingManager.WorkflowExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"workflowExecutionDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"networkRewards\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPermitable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"internalType\":\"structIBillingManager.DepositAssetData\",\"name\":\"depositAssetData\",\"type\":\"tuple\"}],\"internalType\":\"structIBillingManager.DepositAssetInfo\",\"name\":\"_nativeDepositAssetInfo\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"isDepositAssetEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"isDepositAssetPermitable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"isDepositAssetSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"isNativeDepositAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_executionLockedAmount\",\"type\":\"uint256\"}],\"name\":\"lockExecutionFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeDepositAssetKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountToWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"enumIBillingManager.NetworkWithdrawReasons\",\"name\":\"_withdrawReason\",\"type\":\"uint8\"}],\"name\":\"networkWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextWorkflowExecutionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_newEnabledStatus\",\"type\":\"bool\"}],\"name\":\"updateDepositAssetEnabledStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_newWorkflowExecutionDiscount\",\"type\":\"uint256\"}],\"name\":\"updateWorkflowExecutionDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"withdrawAllFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amountToWithdraw\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"}],\"name\":\"withdrawNetworkRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b506080516149a061004c60003960008181610d9a01528181610dda015281816111d70152818161121701526112aa01526149a06000f3fe6080604052600436106102135760003560e01c8063600fe47d11610118578063a19ddf5b116100a0578063c47102041161006f578063c4710204146107ad578063c66ae1e8146107cd578063ef4f272d146107ed578063f3883e261461081a578063f7fa3edf1461084757600080fd5b8063a19ddf5b1461072b578063a266f96314610758578063afb3f30a1461076d578063c172bfd71461078d57600080fd5b80637b103999116100e75780637b1039991461067e5780637c1c6fba1461069e5780637d7f8384146106be57806380017563146106eb5780639769b4391461070b57600080fd5b8063600fe47d146105be578063606a679f146105de5780636e136778146105fe5780637105e9701461063757600080fd5b806348cf8bab1161019b57806355cbf5a51161016a57806355cbf5a5146105005780635a53669c146105205780635af6b043146105405780635c211f88146105605780635ec26c461461059e57600080fd5b806348cf8bab1461047b5780634a686e49146104a85780634f1ef286146104d857806352d1902d146104eb57600080fd5b80632b888682116101e25780632b888682146103bc57806333c25326146103ec5780633659cfe61461040e5780633666cec51461042e578063379135121461044e57600080fd5b80630561010c1461034b57806316fb02dd1461037357806324293941146103935780632a48ac9a146103a957600080fd5b36610346576002805461022590613a60565b80601f016020809104026020016040519081016040528092919081815260200182805461025190613a60565b801561029e5780601f106102735761010080835404028352916020019161029e565b820191906000526020600020905b81548152906001019060200180831161028157829003601f168201915b50505050506102ac81610867565b610344600280546102bc90613a60565b80601f01602080910402602001604051908101604052809291908181526020018280546102e890613a60565b80156103355780601f1061030a57610100808354040283529160200191610335565b820191906000526020600020905b81548152906001019060200180831161031857829003601f168201915b505050505033333460016108d6565b005b600080fd5b34801561035757600080fd5b506103606109e3565b6040519081526020015b60405180910390f35b34801561037f57600080fd5b5061034461038e366004613b8f565b6109f4565b34801561039f57600080fd5b5061036060035481565b6103446103b7366004613bd8565b610bd3565b3480156103c857600080fd5b506103dc6103d7366004613b8f565b610cce565b604051901515815260200161036a565b3480156103f857600080fd5b50610401610d02565b60405161036a9190613c81565b34801561041a57600080fd5b50610344610429366004613c94565b610d90565b34801561043a57600080fd5b50610344610449366004613cb1565b610e6c565b34801561045a57600080fd5b5061046e610469366004613d8b565b611022565b60405161036a9190613e06565b34801561048757600080fd5b5061049b610496366004613e88565b6111b6565b60405161036a9190613eaa565b3480156104b457600080fd5b506103606104c3366004613ef7565b6000908152600a602052604090206001015490565b6103446104e6366004613f10565b6111cd565b3480156104f757600080fd5b5061036061129d565b34801561050c57600080fd5b5061034461051b366004613f73565b611350565b34801561052c57600080fd5b5061036061053b366004613b8f565b6114c3565b34801561054c57600080fd5b5061036061055b366004613fdb565b6114ee565b34801561056c57600080fd5b50600054610586906201000090046001600160a01b031681565b6040516001600160a01b03909116815260200161036a565b3480156105aa57600080fd5b506103446105b9366004614061565b611544565b3480156105ca57600080fd5b506103446105d93660046140b1565b61198c565b3480156105ea57600080fd5b506103dc6105f9366004613b8f565b611b90565b34801561060a57600080fd5b50610586610619366004613ef7565b6000908152600a60205260409020600401546001600160a01b031690565b34801561064357600080fd5b50610671610652366004613ef7565b6000908152600a6020526040902060040154600160a01b900460ff1690565b60405161036a9190614160565b34801561068a57600080fd5b50600154610586906001600160a01b031681565b3480156106aa57600080fd5b506103446106b9366004613d8b565b611b9d565b3480156106ca57600080fd5b506106de6106d936600461416e565b611c88565b60405161036a919061422c565b3480156106f757600080fd5b50610344610706366004613e88565b611d8b565b34801561071757600080fd5b50610344610726366004614281565b61214f565b34801561073757600080fd5b5061074b610746366004613c94565b6121bf565b60405161036a91906142c5565b34801561076457600080fd5b5061074b6121e3565b34801561077957600080fd5b506103dc610788366004613b8f565b6121ef565b34801561079957600080fd5b506103446107a8366004614328565b612296565b3480156107b957600080fd5b506103dc6107c8366004613b8f565b61248f565b3480156107d957600080fd5b506103446107e8366004614373565b6124bd565b3480156107f957600080fd5b5061080d610808366004613ef7565b612633565b60405161036a91906143be565b34801561082657600080fd5b5061083a610835366004613fdb565b612751565b60405161036a9190614426565b34801561085357600080fd5b50610344610862366004613b8f565b612805565b61087081610cce565b6108d35760405162461bcd60e51b815260206004820152602960248201527f42696c6c696e674d616e616765723a204465706f73697420617373657420697360448201526808191a5cd8589b195960ba1b60648201526084015b60405180910390fd5b50565b600082116109325760405162461bcd60e51b815260206004820152602360248201527f42696c6c696e674d616e616765723a205a65726f206465706f73697420616d6f6044820152621d5b9d60ea1b60648201526084016108ca565b8061096e5761096e84308460088960405161094d9190614439565b908152604051908190036020019020546001600160a01b031692919061286f565b61097b85848460016128da565b846040516109899190614439565b604080519182900382206001600160a01b03808816845286166020840152908201849052907ffbc2fad7e5ffeb083133f666db8691af5364707e2b5bf9d1f6f93834433aeca4906060015b60405180910390a25050505050565b60006109ef6004612a06565b905090565b806109fe81612a10565b6000600883604051610a109190614439565b908152604051908190036020019020600281015490915080610a8b5760405162461bcd60e51b815260206004820152602e60248201527f42696c6c696e674d616e616765723a204e6f206e6574776f726b20726577617260448201526d647320746f20776974686472617760901b60648201526084016108ca565b816002016000905560008060029054906101000a90046001600160a01b03166001600160a01b0316631a296e026040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ae7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0b9190614455565b90506001600160a01b038116610b6f5760405162461bcd60e51b815260206004820152602360248201527f42696c6c696e674d616e616765723a205a65726f207369676e6572206164647260448201526265737360e81b60648201526084016108ca565b610b7a858284612a7a565b806001600160a01b031685604051610b929190614439565b604051908190038120848252907f4da529f8b18684c34deb8621c7d4ab35fe1daea1c726cdb6783c93c095d2651a9060200160405180910390a35050505050565b82610bdd81612a10565b83610be781610867565b6000610bf2866121ef565b90508015610c5857348414610c535760405162461bcd60e51b815260206004820152602160248201527f42696c6c696e674d616e616765723a20496e76616c6964206d73672e76616c756044820152606560f81b60648201526084016108ca565b610cb9565b3415610cb95760405162461bcd60e51b815260206004820152602a60248201527f42696c6c696e674d616e616765723a204e6f742061206e6174697665206465706044820152691bdcda5d08185cdcd95d60b21b60648201526084016108ca565b610cc686338787856108d6565b505050505050565b6000600882604051610ce09190614439565b9081526040519081900360200190206003015460ff6101009091041692915050565b60028054610d0f90613a60565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3b90613a60565b8015610d885780601f10610d5d57610100808354040283529160200191610d88565b820191906000526020600020905b815481529060010190602001808311610d6b57829003601f168201915b505050505081565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610dd85760405162461bcd60e51b81526004016108ca90614472565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610e21600080516020614924833981519152546001600160a01b031690565b6001600160a01b031614610e475760405162461bcd60e51b81526004016108ca906144be565b610e5081612ac7565b604080516000808252602082019092526108d391839190612b60565b86610e7681612a10565b87610e8081610867565b610e89896121ef565b15610efc5760405162461bcd60e51b815260206004820152603d60248201527f42696c6c696e674d616e616765723a20556e61626c6520746f206465706f736960448201527f74206e61746976652063757272656e63792077697468207065726d697400000060648201526084016108ca565b610f058961248f565b610f695760405162461bcd60e51b815260206004820152602f60248201527f42696c6c696e674d616e616765723a204465706f73697420617373657420697360448201526e206e6f74207065726d697461626c6560881b60648201526084016108ca565b600889604051610f799190614439565b9081526040519081900360200181205463d505accf60e01b8252336004830152306024830152604482018990526064820188905260ff8716608483015260a4820186905260c482018590526001600160a01b03169063d505accf9060e401600060405180830381600087803b158015610ff157600080fd5b505af1158015611005573d6000803e3d6000fd5b5050505061101789338a8a60006108d6565b505050505050505050565b6060816001600160401b0381111561103c5761103c613a9a565b60405190808252806020026020018201604052801561107557816020015b6110626139c0565b81526020019060019003908161105a5790505b50905060005b828110156111af5760405180604001604052808585848181106110a0576110a061450a565b90506020028101906110b29190614520565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200160088686858181106111005761110061450a565b90506020028101906111129190614520565b604051611120929190614566565b90815260408051918290036020908101832060a08401835280546001600160a01b031684526001810154918401919091526002810154918301919091526003015460ff808216151560608401526101009091041615156080820152905282518390839081106111915761119161450a565b602002602001018190525080806111a79061458c565b91505061107b565b5092915050565b60606111c460048484612ccb565b90505b92915050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036112155760405162461bcd60e51b81526004016108ca90614472565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031661125e600080516020614924833981519152546001600160a01b031690565b6001600160a01b0316146112845760405162461bcd60e51b81526004016108ca906144be565b61128d82612ac7565b61129982826001612b60565b5050565b6000306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461133d5760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c000000000000000060648201526084016108ca565b5060008051602061492483398151915290565b600054610100900460ff16158080156113705750600054600160ff909116105b8061138a5750303b15801561138a575060005460ff166001145b6113ed5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016108ca565b6000805460ff191660011790558015611410576000805461ff0019166101001790555b600180546001600160a01b0319166001600160a01b0386161790556114358280614520565b6002916114439190836145eb565b50611455611450836146aa565b612d83565b6000805462010000600160b01b031916620100006001600160a01b0386160217905580156114bd576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b60006008826040516114d59190614439565b9081526020016040518091039020600201549050919050565b6001600160a01b038216600090815260096020526040808220905182916002019061151a908590614439565b9081526040519081900360200190206001810154815491925061153c9161475e565b949350505050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015611593573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b79190614455565b6001600160a01b0316146115dd5760405162461bcd60e51b81526004016108ca90614771565b83838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061161f9250839150612a109050565b600154604051639b2bfaa360e01b8152600481018590526001600160a01b0390911690639b2bfaa390602401602060405180830381865afa158015611668573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061168c91906147a8565b6116e85760405162461bcd60e51b815260206004820152602760248201527f42696c6c696e674d616e616765723a20576f726b666c6f7720646f6573206e6f6044820152661d08195e1a5cdd60ca1b60648201526084016108ca565b60015460405163d69cd27560e01b8152600481018590526000916001600160a01b03169063d69cd27590602401602060405180830381865afa158015611732573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117569190614455565b905061179b86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508592508791506130309050565b60038054600091826117ac8361458c565b919050559050600060096000846001600160a01b03166001600160a01b0316815260200190815260200160002060020188886040516117ec929190614566565b908152602001604051809103902090508481600101600082825461181091906147c5565b909155506118239050600282018361309c565b506040518060c0016040528089898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250938552505050602082018990526040820188905260608201526001600160a01b038516608082015260a001600190526000838152600a60205260409020815181906118ac90826147d8565b506020820151600182015560408201516002808301919091556060830151600383015560808301516004830180546001600160a01b039092166001600160a01b031983168117825560a08601519391926001600160a81b0319161790600160a01b90849081111561191f5761191f614128565b0217905550905050826001600160a01b0316868989604051611942929190614566565b60408051918290038220868352602083018a9052917f318669d1aaaf361c152eec849d7bf340dcac03acfb197842a8fd141df0a408fa910160405180910390a45050505050505050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa1580156119db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119ff9190614455565b6001600160a01b031614611a255760405162461bcd60e51b81526004016108ca90614771565b84848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a679250839150612a109050565b611aaa86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508892508791506130309050565b611aec86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525089935088925090506128da565b8260088787604051611aff929190614566565b90815260200160405180910390206002016000828254611b1f91906147c5565b909155508290508015611b3457611b34614128565b846001600160a01b03168787604051611b4e929190614566565b604051908190038120868252907ff913ff1a24ed376db7f7964cb178db8a9ec8c5462a747385218f3610003724d99060200160405180910390a4505050505050565b60006111c76006836130a8565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015611bec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c109190614455565b6001600160a01b031614611c365760405162461bcd60e51b81526004016108ca90614771565b60005b81811015611c8357611c71838383818110611c5657611c5661450a565b9050602002810190611c689190614897565b611450906146aa565b80611c7b8161458c565b915050611c39565b505050565b60606000611ca0611c996004612a06565b85856130d5565b9050611cac848261475e565b6001600160401b03811115611cc357611cc3613a9a565b604051908082528060200260200182016040528015611d2857816020015b611d15604051806080016040528060006001600160a01b031681526020016000815260200160008152602001606081525090565b815260200190600190039081611ce15790505b509150835b81811015611d8257611d49611d43600483613100565b87612751565b83611d54878461475e565b81518110611d6457611d6461450a565b60200260200101819052508080611d7a9061458c565b915050611d2d565b50509392505050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015611dda573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dfe9190614455565b6001600160a01b031614611e245760405162461bcd60e51b81526004016108ca90614771565b6000828152600a6020526040902060016004820154600160a01b900460ff166002811115611e5457611e54614128565b14611eba5760405162461bcd60e51b815260206004820152603060248201527f42696c6c696e674d616e616765723a204e6f7420612070656e64696e6720776f60448201526f3935b33637bb9032bc32b1baba34b7b760811b60648201526084016108ca565b6000816000018054611ecb90613a60565b80601f0160208091040260200160405190810160405280929190818152602001828054611ef790613a60565b8015611f445780601f10611f1957610100808354040283529160200191611f44565b820191906000526020600020905b815481529060010190602001808311611f2757829003601f168201915b50505060048501546001600160a01b03166000908152600960205260408082209051949550909360029091019250611f7e91508490614439565b90815260405190819003602001902060028401549091508481101561200a57611fc5858284600101548560000154611fb6919061475e565b611fc091906147c5565b61310c565b60408051888152602081018490529081018290529095507faba1d671589f510878d4421f3e9f5574d37ecf6735d3f3df2be569b62b9015d99060600160405180910390a15b60048401805460ff60a01b1916600160a11b1790556003840185905560018201805482919060009061203d90849061475e565b9091555061205090506002830187613122565b50600484015461206d9084906001600160a01b03168760006128da565b8460088460405161207e9190614439565b9081526020016040518091039020600201600082825461209e91906147c5565b9091555050600180549085015460405163ae09340f60e01b81526001600160a01b039092169163ae09340f916120db918791908a906004016148ad565b600060405180830381600087803b1580156120f557600080fd5b505af1158015612109573d6000803e3d6000fd5b505060408051898152602081018990527f70114d139c12c9f15fb8a74cc0b9d9f2070f048d32cdf038844425444bad7b52935001905060405180910390a1505050505050565b8161215981612a10565b612164833384613030565b61216f83338461312e565b604051339061217f908590614439565b604051908190038120848252907fe0dd7f2d0bf336d345c1e8051edf7f753d4bfdc13d35cf30351267de2aa905c8906020015b60405180910390a3505050565b6001600160a01b03811660009081526009602052604090206060906111c790613146565b60606109ef6006613146565b60006111c7826002805461220290613a60565b80601f016020809104026020016040519081016040528092919081815260200182805461222e90613a60565b801561227b5780601f106122505761010080835404028352916020019161227b565b820191906000526020600020905b81548152906001019060200180831161225e57829003601f168201915b50505050508051602091820120825192909101919091201490565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa1580156122e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123099190614455565b6001600160a01b03161461232f5760405162461bcd60e51b81526004016108ca90614771565b82828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506123719250839150612a109050565b81151560088585604051612386929190614566565b9081526040519081900360200190206003015460ff610100909104161515036124045760405162461bcd60e51b815260206004820152602a60248201527f42696c6c696e674d616e616765723a20496e76616c6964206e657720656e61626044820152696c65642073746174757360b01b60648201526084016108ca565b8160088585604051612417929190614566565b90815260405190819003602001812060030180549215156101000261ff00199093169290921790915561244d9085908590614566565b6040519081900381208315158252907f838dfb07f2f9eda916131303b1c93fcdb2d1f0cfae1f651eee30c8bd440913569060200160405180910390a250505050565b60006008826040516124a19190614439565b9081526040519081900360200190206003015460ff1692915050565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa15801561250c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125309190614455565b6001600160a01b0316146125565760405162461bcd60e51b81526004016108ca90614771565b82828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506125989250839150612a109050565b6000600885856040516125ac929190614566565b908152602001604051809103902060010154905082600886866040516125d3929190614566565b908152604051908190036020018120600101919091556125f69086908690614566565b6040805191829003822085835260208301849052917f70c21b53e4a29c6a334ae8d8d80a885023e4749ee873d6594b3f6cbb66140b6b91016109d4565b61263b613a16565b6000828152600a602052604090819020815160c0810190925280548290829061266390613a60565b80601f016020809104026020016040519081016040528092919081815260200182805461268f90613a60565b80156126dc5780601f106126b1576101008083540402835291602001916126dc565b820191906000526020600020905b8154815290600101906020018083116126bf57829003601f168201915b50505091835250506001820154602082015260028083015460408301526003830154606083015260048301546001600160a01b038116608084015260a090920191600160a01b900460ff169081111561273757612737614128565b600281111561274857612748614128565b90525092915050565b612785604051806080016040528060006001600160a01b031681526020016000815260200160008152602001606081525090565b6001600160a01b03831660009081526009602052604080822090516002909101906127b1908590614439565b908152602001604051809103902090506040518060800160405280856001600160a01b0316815260200182600001548152602001826001015481526020016127fb83600201613223565b9052949350505050565b8061280f81612a10565b600061281b33846114ee565b905061282883338361312e565b6040513390612838908590614439565b604051908190038120838252907fe0dd7f2d0bf336d345c1e8051edf7f753d4bfdc13d35cf30351267de2aa905c8906020016121b2565b6040516001600160a01b03808516602483015283166044820152606481018290526114bd9085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152613230565b600082116129385760405162461bcd60e51b815260206004820152602560248201527f42696c6c696e674d616e616765723a205a65726f20616d6f756e7420746f2075604482015264706461746560d81b60648201526084016108ca565b6001600160a01b038316600090815260096020526040902081156129ab578281600201866040516129699190614439565b9081526020016040518091039020600001600082825461298991906147c5565b9091555061299990508186613305565b506129a5600485613367565b506129ff565b600081600201866040516129bf9190614439565b9081526020016040518091039020905060008482600001546129e1919061475e565b808355905060008190036129fc576129fa60048761337c565b505b50505b5050505050565b60006111c7825490565b612a1981611b90565b6108d35760405162461bcd60e51b815260206004820152602c60248201527f42696c6c696e674d616e616765723a204465706f73697420617373657420646f60448201526b195cc81b9bdd08195e1a5cdd60a21b60648201526084016108ca565b612a83836121ef565b15612a9257611c838282613391565b611c838282600886604051612aa79190614439565b908152604051908190036020019020546001600160a01b031691906134aa565b60005460408051630d14b70160e11b8152905133926201000090046001600160a01b031691631a296e029160048083019260209291908290030181865afa158015612b16573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b3a9190614455565b6001600160a01b0316146108d35760405162461bcd60e51b81526004016108ca90614771565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff1615612b9357611c83836134da565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612bed575060408051601f3d908101601f19168201909252612bea918101906148d2565b60015b612c505760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b60648201526084016108ca565b6000805160206149248339815191528114612cbf5760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b60648201526084016108ca565b50611c83838383613576565b60606000612cdb611c9986612a06565b9050612ce7848261475e565b6001600160401b03811115612cfe57612cfe613a9a565b604051908082528060200260200182016040528015612d27578160200160208202803683370190505b509150835b81811015611d8257612d3e8682613100565b83612d49878461475e565b81518110612d5957612d5961450a565b6001600160a01b039092166020928302919091019091015280612d7b8161458c565b915050612d2c565b8051612d8e90611b90565b15612df05760405162461bcd60e51b815260206004820152602c60248201527f42696c6c696e674d616e616765723a204465706f73697420617373657420616c60448201526b72656164792065786973747360a01b60648201526084016108ca565b6020810151516001600160a01b0316151580612e1c5750612e1c81600001516002805461220290613a60565b612e845760405162461bcd60e51b815260206004820152603360248201527f42696c6c696e674d616e616765723a20496e76616c6964206465706f73697420604482015272617373657420746f6b656e206164647265737360681b60648201526084016108ca565b805151612ee55760405162461bcd60e51b815260206004820152602960248201527f42696c6c696e674d616e616765723a20496e76616c6964206465706f736974206044820152686173736574206b657960b81b60648201526084016108ca565b60208101516040015115612f565760405162461bcd60e51b815260206004820152603260248201527f42696c6c696e674d616e616765723a20496e76616c696420696e6974206e6574604482015271776f726b20726577617264732076616c756560701b60648201526084016108ca565b60208101518151604051600891612f6c91614439565b90815260408051918290036020908101909220835181546001600160a01b039091166001600160a01b0319909116178155918301516001830155820151600282015560608201516003909101805460809093015115156101000261ff00199215159290921661ffff19909316929092171790558051612fed90600690613305565b50805160208201516040517f310a01f49fbf43988aa0ce4a44296f5fc9f1fe442769d2416f1965d6a1cd236e926130259290916148eb565b60405180910390a150565b8061303b83856114ee565b1015611c835760405162461bcd60e51b815260206004820152602a60248201527f42696c6c696e674d616e616765723a204e6f7420656e6f75676820617661696c60448201526961626c652066756e647360b01b60648201526084016108ca565b60006111c4838361359b565b600082600101826040516130bc9190614439565b9081526040519081900360200190205415159392505050565b60006130e182846147c5565b9050838111156130ee5750825b808311156130f95750815b9392505050565b60006111c483836135e2565b600081831061311b57816111c4565b5090919050565b60006111c4838361360c565b61313b83838360006128da565b611c83838383612a7a565b606081600001805480602002602001604051908101604052809291908181526020016000905b8282101561321857838290600052602060002001805461318b90613a60565b80601f01602080910402602001604051908101604052809291908181526020018280546131b790613a60565b80156132045780601f106131d957610100808354040283529160200191613204565b820191906000526020600020905b8154815290600101906020018083116131e757829003601f168201915b50505050508152602001906001019061316c565b505050509050919050565b606060006130f9836136ff565b6000613285826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661375b9092919063ffffffff16565b90508051600014806132a65750808060200190518101906132a691906147a8565b611c835760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016108ca565b600061331183836130a8565b61335f5782546001810184556000848152602090200161333183826147d8565b5082546040516001850190613347908590614439565b908152604051908190036020019020555060016111c7565b5060006111c7565b60006111c4836001600160a01b03841661359b565b60006111c4836001600160a01b03841661360c565b804710156133e15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016108ca565b6000826001600160a01b03168260405160006040518083038185875af1925050503d806000811461342e576040519150601f19603f3d011682016040523d82523d6000602084013e613433565b606091505b5050905080611c835760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016108ca565b6040516001600160a01b038316602482015260448101829052611c8390849063a9059cbb60e01b906064016128a3565b6001600160a01b0381163b6135475760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016108ca565b60008051602061492483398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b61357f8361376a565b60008251118061358c5750805b15611c83576114bd83836137aa565b600081815260018301602052604081205461335f575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556111c7565b60008260000182815481106135f9576135f961450a565b9060005260206000200154905092915050565b600081815260018301602052604081205480156136f557600061363060018361475e565b85549091506000906136449060019061475e565b90508181146136a95760008660000182815481106136645761366461450a565b90600052602060002001549050808760000184815481106136875761368761450a565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806136ba576136ba61490d565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506111c7565b60009150506111c7565b60608160000180548060200260200160405190810160405280929190818152602001828054801561374f57602002820191906000526020600020905b81548152602001906001019080831161373b575b50505050509050919050565b606061153c84846000856137cf565b613773816134da565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606111c48383604051806060016040528060278152602001614944602791396138aa565b6060824710156138305760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016108ca565b600080866001600160a01b0316858760405161384c9190614439565b60006040518083038185875af1925050503d8060008114613889576040519150601f19603f3d011682016040523d82523d6000602084013e61388e565b606091505b509150915061389f87838387613922565b979650505050505050565b6060600080856001600160a01b0316856040516138c79190614439565b600060405180830381855af49150503d8060008114613902576040519150601f19603f3d011682016040523d82523d6000602084013e613907565b606091505b509150915061391886838387613922565b9695505050505050565b6060831561399157825160000361398a576001600160a01b0385163b61398a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016108ca565b508161153c565b61153c83838151156139a65781518083602001fd5b8060405162461bcd60e51b81526004016108ca9190613c81565b604051806040016040528060608152602001613a116040518060a0016040528060006001600160a01b0316815260200160008152602001600081526020016000151581526020016000151581525090565b905290565b6040518060c001604052806060815260200160008152602001600081526020016000815260200160006001600160a01b0316815260200160006002811115613a1157613a11614128565b600181811c90821680613a7457607f821691505b602082108103613a9457634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b0381118282101715613ad257613ad2613a9a565b60405290565b60405160a081016001600160401b0381118282101715613ad257613ad2613a9a565b60006001600160401b0380841115613b1457613b14613a9a565b604051601f8501601f19908116603f01168101908282118183101715613b3c57613b3c613a9a565b81604052809350858152868686011115613b5557600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112613b8057600080fd5b6111c483833560208501613afa565b600060208284031215613ba157600080fd5b81356001600160401b03811115613bb757600080fd5b61153c84828501613b6f565b6001600160a01b03811681146108d357600080fd5b600080600060608486031215613bed57600080fd5b83356001600160401b03811115613c0357600080fd5b613c0f86828701613b6f565b9350506020840135613c2081613bc3565b929592945050506040919091013590565b60005b83811015613c4c578181015183820152602001613c34565b50506000910152565b60008151808452613c6d816020860160208601613c31565b601f01601f19169290920160200192915050565b6020815260006111c46020830184613c55565b600060208284031215613ca657600080fd5b81356130f981613bc3565b600080600080600080600060e0888a031215613ccc57600080fd5b87356001600160401b03811115613ce257600080fd5b613cee8a828b01613b6f565b9750506020880135613cff81613bc3565b95506040880135945060608801359350608088013560ff81168114613d2357600080fd5b9699959850939692959460a0840135945060c09093013592915050565b60008083601f840112613d5257600080fd5b5081356001600160401b03811115613d6957600080fd5b6020830191508360208260051b8501011115613d8457600080fd5b9250929050565b60008060208385031215613d9e57600080fd5b82356001600160401b03811115613db457600080fd5b613dc085828601613d40565b90969095509350505050565b80516001600160a01b0316825260208082015190830152604080820151908301526060808201511515908301526080908101511515910152565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015613e7b57603f19888603018452815160c08151818852613e5382890182613c55565b915050878201519150613e6888880183613dcc565b9550509285019290850190600101613e2d565b5092979650505050505050565b60008060408385031215613e9b57600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b81811015613eeb5783516001600160a01b031683529284019291840191600101613ec6565b50909695505050505050565b600060208284031215613f0957600080fd5b5035919050565b60008060408385031215613f2357600080fd5b8235613f2e81613bc3565b915060208301356001600160401b03811115613f4957600080fd5b8301601f81018513613f5a57600080fd5b613f6985823560208401613afa565b9150509250929050565b600080600060608486031215613f8857600080fd5b8335613f9381613bc3565b92506020840135613fa381613bc3565b915060408401356001600160401b03811115613fbe57600080fd5b840160c08187031215613fd057600080fd5b809150509250925092565b60008060408385031215613fee57600080fd5b8235613ff981613bc3565b915060208301356001600160401b0381111561401457600080fd5b613f6985828601613b6f565b60008083601f84011261403257600080fd5b5081356001600160401b0381111561404957600080fd5b602083019150836020828501011115613d8457600080fd5b6000806000806060858703121561407757600080fd5b84356001600160401b0381111561408d57600080fd5b61409987828801614020565b90989097506020870135966040013595509350505050565b6000806000806000608086880312156140c957600080fd5b85356001600160401b038111156140df57600080fd5b6140eb88828901614020565b90965094505060208601356140ff81613bc3565b92506040860135915060608601356001811061411a57600080fd5b809150509295509295909350565b634e487b7160e01b600052602160045260246000fd5b6003811061415c57634e487b7160e01b600052602160045260246000fd5b9052565b602081016111c7828461413e565b60008060006060848603121561418357600080fd5b83356001600160401b0381111561419957600080fd5b6141a586828701613b6f565b9660208601359650604090950135949350505050565b80516001600160a01b0316825260208082015181840152604080830151908401526060808301516080918501829052805191850182905260009290810191839060a08701905b808310156142215784518252938301936001929092019190830190614201565b509695505050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015613e7b57603f1988860301845261426f8583516141bb565b94509285019290850190600101614253565b6000806040838503121561429457600080fd5b82356001600160401b038111156142aa57600080fd5b6142b685828601613b6f565b95602094909401359450505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015613e7b57603f19888603018452614308858351613c55565b945092850192908501906001016142ec565b80151581146108d357600080fd5b60008060006040848603121561433d57600080fd5b83356001600160401b0381111561435357600080fd5b61435f86828701614020565b9094509250506020840135613fd08161431a565b60008060006040848603121561438857600080fd5b83356001600160401b0381111561439e57600080fd5b6143aa86828701614020565b909790965060209590950135949350505050565b602081526000825160c060208401526143da60e0840182613c55565b905060208401516040840152604084015160608401526060840151608084015260018060a01b0360808501511660a084015260a084015161441e60c085018261413e565b509392505050565b6020815260006111c460208301846141bb565b6000825161444b818460208701613c31565b9190910192915050565b60006020828403121561446757600080fd5b81516130f981613bc3565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261453757600080fd5b8301803591506001600160401b0382111561455157600080fd5b602001915036819003821315613d8457600080fd5b8183823760009101908152919050565b634e487b7160e01b600052601160045260246000fd5b60006001820161459e5761459e614576565b5060010190565b601f821115611c8357600081815260208120601f850160051c810160208610156145cc5750805b601f850160051c820191505b81811015610cc6578281556001016145d8565b6001600160401b0383111561460257614602613a9a565b614616836146108354613a60565b836145a5565b6000601f84116001811461464a57600085156146325750838201355b600019600387901b1c1916600186901b1783556129ff565b600083815260209020601f19861690835b8281101561467b578685013582556020948501946001909201910161465b565b50868210156146985760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b600081360360c08112156146bd57600080fd5b6146c5613ab0565b83356001600160401b038111156146db57600080fd5b6146e736828701613b6f565b82525060a0601f19830112156146fc57600080fd5b614704613ad8565b9150602084013561471481613bc3565b8252604084810135602084015260608501359083015260808401356147388161431a565b606083015260a084013561474b8161431a565b6080830152602081019190915292915050565b818103818111156111c7576111c7614576565b6020808252601a908201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604082015260600190565b6000602082840312156147ba57600080fd5b81516130f98161431a565b808201808211156111c7576111c7614576565b81516001600160401b038111156147f1576147f1613a9a565b614805816147ff8454613a60565b846145a5565b602080601f83116001811461483a57600084156148225750858301515b600019600386901b1c1916600185901b178555610cc6565b600085815260208120601f198616915b828110156148695788860151825594840194600190910190840161484a565b50858210156148875787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000823560be1983360301811261444b57600080fd5b6060815260006148c06060830186613c55565b60208301949094525060400152919050565b6000602082840312156148e457600080fd5b5051919050565b60c0815260006148fe60c0830185613c55565b90506130f96020830184613dcc565b634e487b7160e01b600052603160045260246000fdfe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212204e1d0faba94b040116f59abec354aaa5819f6c1d0bb5ff98e443f713a02a238a64736f6c63430008120033",
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

// GetDepositAssetsInfo is a free data retrieval call binding the contract method 0x37913512.
//
// Solidity: function getDepositAssetsInfo(string[] _depositAssetKeysArr) view returns((string,(address,uint256,uint256,bool,bool))[] _depositAssetsInfoArr)
func (_BillingManager *BillingManagerCaller) GetDepositAssetsInfo(opts *bind.CallOpts, _depositAssetKeysArr []string) ([]IBillingManagerDepositAssetInfo, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getDepositAssetsInfo", _depositAssetKeysArr)

	if err != nil {
		return *new([]IBillingManagerDepositAssetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBillingManagerDepositAssetInfo)).(*[]IBillingManagerDepositAssetInfo)

	return out0, err

}

// GetDepositAssetsInfo is a free data retrieval call binding the contract method 0x37913512.
//
// Solidity: function getDepositAssetsInfo(string[] _depositAssetKeysArr) view returns((string,(address,uint256,uint256,bool,bool))[] _depositAssetsInfoArr)
func (_BillingManager *BillingManagerSession) GetDepositAssetsInfo(_depositAssetKeysArr []string) ([]IBillingManagerDepositAssetInfo, error) {
	return _BillingManager.Contract.GetDepositAssetsInfo(&_BillingManager.CallOpts, _depositAssetKeysArr)
}

// GetDepositAssetsInfo is a free data retrieval call binding the contract method 0x37913512.
//
// Solidity: function getDepositAssetsInfo(string[] _depositAssetKeysArr) view returns((string,(address,uint256,uint256,bool,bool))[] _depositAssetsInfoArr)
func (_BillingManager *BillingManagerCallerSession) GetDepositAssetsInfo(_depositAssetKeysArr []string) ([]IBillingManagerDepositAssetInfo, error) {
	return _BillingManager.Contract.GetDepositAssetsInfo(&_BillingManager.CallOpts, _depositAssetKeysArr)
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

// GetNetworkRewards is a free data retrieval call binding the contract method 0x5a53669c.
//
// Solidity: function getNetworkRewards(string _depositAssetKey) view returns(uint256)
func (_BillingManager *BillingManagerCaller) GetNetworkRewards(opts *bind.CallOpts, _depositAssetKey string) (*big.Int, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getNetworkRewards", _depositAssetKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNetworkRewards is a free data retrieval call binding the contract method 0x5a53669c.
//
// Solidity: function getNetworkRewards(string _depositAssetKey) view returns(uint256)
func (_BillingManager *BillingManagerSession) GetNetworkRewards(_depositAssetKey string) (*big.Int, error) {
	return _BillingManager.Contract.GetNetworkRewards(&_BillingManager.CallOpts, _depositAssetKey)
}

// GetNetworkRewards is a free data retrieval call binding the contract method 0x5a53669c.
//
// Solidity: function getNetworkRewards(string _depositAssetKey) view returns(uint256)
func (_BillingManager *BillingManagerCallerSession) GetNetworkRewards(_depositAssetKey string) (*big.Int, error) {
	return _BillingManager.Contract.GetNetworkRewards(&_BillingManager.CallOpts, _depositAssetKey)
}

// GetSupportedDepositAssetKeys is a free data retrieval call binding the contract method 0xa266f963.
//
// Solidity: function getSupportedDepositAssetKeys() view returns(string[])
func (_BillingManager *BillingManagerCaller) GetSupportedDepositAssetKeys(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getSupportedDepositAssetKeys")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetSupportedDepositAssetKeys is a free data retrieval call binding the contract method 0xa266f963.
//
// Solidity: function getSupportedDepositAssetKeys() view returns(string[])
func (_BillingManager *BillingManagerSession) GetSupportedDepositAssetKeys() ([]string, error) {
	return _BillingManager.Contract.GetSupportedDepositAssetKeys(&_BillingManager.CallOpts)
}

// GetSupportedDepositAssetKeys is a free data retrieval call binding the contract method 0xa266f963.
//
// Solidity: function getSupportedDepositAssetKeys() view returns(string[])
func (_BillingManager *BillingManagerCallerSession) GetSupportedDepositAssetKeys() ([]string, error) {
	return _BillingManager.Contract.GetSupportedDepositAssetKeys(&_BillingManager.CallOpts)
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

// GetUserAvailableFunds is a free data retrieval call binding the contract method 0x5af6b043.
//
// Solidity: function getUserAvailableFunds(address _userAddr, string _depositAssetKey) view returns(uint256)
func (_BillingManager *BillingManagerCaller) GetUserAvailableFunds(opts *bind.CallOpts, _userAddr common.Address, _depositAssetKey string) (*big.Int, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getUserAvailableFunds", _userAddr, _depositAssetKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserAvailableFunds is a free data retrieval call binding the contract method 0x5af6b043.
//
// Solidity: function getUserAvailableFunds(address _userAddr, string _depositAssetKey) view returns(uint256)
func (_BillingManager *BillingManagerSession) GetUserAvailableFunds(_userAddr common.Address, _depositAssetKey string) (*big.Int, error) {
	return _BillingManager.Contract.GetUserAvailableFunds(&_BillingManager.CallOpts, _userAddr, _depositAssetKey)
}

// GetUserAvailableFunds is a free data retrieval call binding the contract method 0x5af6b043.
//
// Solidity: function getUserAvailableFunds(address _userAddr, string _depositAssetKey) view returns(uint256)
func (_BillingManager *BillingManagerCallerSession) GetUserAvailableFunds(_userAddr common.Address, _depositAssetKey string) (*big.Int, error) {
	return _BillingManager.Contract.GetUserAvailableFunds(&_BillingManager.CallOpts, _userAddr, _depositAssetKey)
}

// GetUserDepositAssetKeys is a free data retrieval call binding the contract method 0xa19ddf5b.
//
// Solidity: function getUserDepositAssetKeys(address _userAddr) view returns(string[])
func (_BillingManager *BillingManagerCaller) GetUserDepositAssetKeys(opts *bind.CallOpts, _userAddr common.Address) ([]string, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getUserDepositAssetKeys", _userAddr)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetUserDepositAssetKeys is a free data retrieval call binding the contract method 0xa19ddf5b.
//
// Solidity: function getUserDepositAssetKeys(address _userAddr) view returns(string[])
func (_BillingManager *BillingManagerSession) GetUserDepositAssetKeys(_userAddr common.Address) ([]string, error) {
	return _BillingManager.Contract.GetUserDepositAssetKeys(&_BillingManager.CallOpts, _userAddr)
}

// GetUserDepositAssetKeys is a free data retrieval call binding the contract method 0xa19ddf5b.
//
// Solidity: function getUserDepositAssetKeys(address _userAddr) view returns(string[])
func (_BillingManager *BillingManagerCallerSession) GetUserDepositAssetKeys(_userAddr common.Address) ([]string, error) {
	return _BillingManager.Contract.GetUserDepositAssetKeys(&_BillingManager.CallOpts, _userAddr)
}

// GetUserDepositInfo is a free data retrieval call binding the contract method 0xf3883e26.
//
// Solidity: function getUserDepositInfo(address _userAddr, string _depositAssetKey) view returns((address,uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerCaller) GetUserDepositInfo(opts *bind.CallOpts, _userAddr common.Address, _depositAssetKey string) (IBillingManagerUserDepositInfo, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getUserDepositInfo", _userAddr, _depositAssetKey)

	if err != nil {
		return *new(IBillingManagerUserDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IBillingManagerUserDepositInfo)).(*IBillingManagerUserDepositInfo)

	return out0, err

}

// GetUserDepositInfo is a free data retrieval call binding the contract method 0xf3883e26.
//
// Solidity: function getUserDepositInfo(address _userAddr, string _depositAssetKey) view returns((address,uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerSession) GetUserDepositInfo(_userAddr common.Address, _depositAssetKey string) (IBillingManagerUserDepositInfo, error) {
	return _BillingManager.Contract.GetUserDepositInfo(&_BillingManager.CallOpts, _userAddr, _depositAssetKey)
}

// GetUserDepositInfo is a free data retrieval call binding the contract method 0xf3883e26.
//
// Solidity: function getUserDepositInfo(address _userAddr, string _depositAssetKey) view returns((address,uint256,uint256,uint256[]))
func (_BillingManager *BillingManagerCallerSession) GetUserDepositInfo(_userAddr common.Address, _depositAssetKey string) (IBillingManagerUserDepositInfo, error) {
	return _BillingManager.Contract.GetUserDepositInfo(&_BillingManager.CallOpts, _userAddr, _depositAssetKey)
}

// GetUsersDepositInfo is a free data retrieval call binding the contract method 0x7d7f8384.
//
// Solidity: function getUsersDepositInfo(string _depositAssetKey, uint256 _offset, uint256 _limit) view returns((address,uint256,uint256,uint256[])[] _usersInfoArr)
func (_BillingManager *BillingManagerCaller) GetUsersDepositInfo(opts *bind.CallOpts, _depositAssetKey string, _offset *big.Int, _limit *big.Int) ([]IBillingManagerUserDepositInfo, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "getUsersDepositInfo", _depositAssetKey, _offset, _limit)

	if err != nil {
		return *new([]IBillingManagerUserDepositInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBillingManagerUserDepositInfo)).(*[]IBillingManagerUserDepositInfo)

	return out0, err

}

// GetUsersDepositInfo is a free data retrieval call binding the contract method 0x7d7f8384.
//
// Solidity: function getUsersDepositInfo(string _depositAssetKey, uint256 _offset, uint256 _limit) view returns((address,uint256,uint256,uint256[])[] _usersInfoArr)
func (_BillingManager *BillingManagerSession) GetUsersDepositInfo(_depositAssetKey string, _offset *big.Int, _limit *big.Int) ([]IBillingManagerUserDepositInfo, error) {
	return _BillingManager.Contract.GetUsersDepositInfo(&_BillingManager.CallOpts, _depositAssetKey, _offset, _limit)
}

// GetUsersDepositInfo is a free data retrieval call binding the contract method 0x7d7f8384.
//
// Solidity: function getUsersDepositInfo(string _depositAssetKey, uint256 _offset, uint256 _limit) view returns((address,uint256,uint256,uint256[])[] _usersInfoArr)
func (_BillingManager *BillingManagerCallerSession) GetUsersDepositInfo(_depositAssetKey string, _offset *big.Int, _limit *big.Int) ([]IBillingManagerUserDepositInfo, error) {
	return _BillingManager.Contract.GetUsersDepositInfo(&_BillingManager.CallOpts, _depositAssetKey, _offset, _limit)
}

// GetWorkflowExecutionInfo is a free data retrieval call binding the contract method 0xef4f272d.
//
// Solidity: function getWorkflowExecutionInfo(uint256 _workflowExecutionId) view returns((string,uint256,uint256,uint256,address,uint8))
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
// Solidity: function getWorkflowExecutionInfo(uint256 _workflowExecutionId) view returns((string,uint256,uint256,uint256,address,uint8))
func (_BillingManager *BillingManagerSession) GetWorkflowExecutionInfo(_workflowExecutionId *big.Int) (IBillingManagerWorkflowExecutionInfo, error) {
	return _BillingManager.Contract.GetWorkflowExecutionInfo(&_BillingManager.CallOpts, _workflowExecutionId)
}

// GetWorkflowExecutionInfo is a free data retrieval call binding the contract method 0xef4f272d.
//
// Solidity: function getWorkflowExecutionInfo(uint256 _workflowExecutionId) view returns((string,uint256,uint256,uint256,address,uint8))
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

// IsDepositAssetEnabled is a free data retrieval call binding the contract method 0x2b888682.
//
// Solidity: function isDepositAssetEnabled(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCaller) IsDepositAssetEnabled(opts *bind.CallOpts, _depositAssetKey string) (bool, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "isDepositAssetEnabled", _depositAssetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositAssetEnabled is a free data retrieval call binding the contract method 0x2b888682.
//
// Solidity: function isDepositAssetEnabled(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerSession) IsDepositAssetEnabled(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsDepositAssetEnabled(&_BillingManager.CallOpts, _depositAssetKey)
}

// IsDepositAssetEnabled is a free data retrieval call binding the contract method 0x2b888682.
//
// Solidity: function isDepositAssetEnabled(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCallerSession) IsDepositAssetEnabled(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsDepositAssetEnabled(&_BillingManager.CallOpts, _depositAssetKey)
}

// IsDepositAssetPermitable is a free data retrieval call binding the contract method 0xc4710204.
//
// Solidity: function isDepositAssetPermitable(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCaller) IsDepositAssetPermitable(opts *bind.CallOpts, _depositAssetKey string) (bool, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "isDepositAssetPermitable", _depositAssetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositAssetPermitable is a free data retrieval call binding the contract method 0xc4710204.
//
// Solidity: function isDepositAssetPermitable(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerSession) IsDepositAssetPermitable(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsDepositAssetPermitable(&_BillingManager.CallOpts, _depositAssetKey)
}

// IsDepositAssetPermitable is a free data retrieval call binding the contract method 0xc4710204.
//
// Solidity: function isDepositAssetPermitable(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCallerSession) IsDepositAssetPermitable(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsDepositAssetPermitable(&_BillingManager.CallOpts, _depositAssetKey)
}

// IsDepositAssetSupported is a free data retrieval call binding the contract method 0x606a679f.
//
// Solidity: function isDepositAssetSupported(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCaller) IsDepositAssetSupported(opts *bind.CallOpts, _depositAssetKey string) (bool, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "isDepositAssetSupported", _depositAssetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositAssetSupported is a free data retrieval call binding the contract method 0x606a679f.
//
// Solidity: function isDepositAssetSupported(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerSession) IsDepositAssetSupported(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsDepositAssetSupported(&_BillingManager.CallOpts, _depositAssetKey)
}

// IsDepositAssetSupported is a free data retrieval call binding the contract method 0x606a679f.
//
// Solidity: function isDepositAssetSupported(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCallerSession) IsDepositAssetSupported(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsDepositAssetSupported(&_BillingManager.CallOpts, _depositAssetKey)
}

// IsNativeDepositAsset is a free data retrieval call binding the contract method 0xafb3f30a.
//
// Solidity: function isNativeDepositAsset(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCaller) IsNativeDepositAsset(opts *bind.CallOpts, _depositAssetKey string) (bool, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "isNativeDepositAsset", _depositAssetKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNativeDepositAsset is a free data retrieval call binding the contract method 0xafb3f30a.
//
// Solidity: function isNativeDepositAsset(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerSession) IsNativeDepositAsset(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsNativeDepositAsset(&_BillingManager.CallOpts, _depositAssetKey)
}

// IsNativeDepositAsset is a free data retrieval call binding the contract method 0xafb3f30a.
//
// Solidity: function isNativeDepositAsset(string _depositAssetKey) view returns(bool)
func (_BillingManager *BillingManagerCallerSession) IsNativeDepositAsset(_depositAssetKey string) (bool, error) {
	return _BillingManager.Contract.IsNativeDepositAsset(&_BillingManager.CallOpts, _depositAssetKey)
}

// NativeDepositAssetKey is a free data retrieval call binding the contract method 0x33c25326.
//
// Solidity: function nativeDepositAssetKey() view returns(string)
func (_BillingManager *BillingManagerCaller) NativeDepositAssetKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BillingManager.contract.Call(opts, &out, "nativeDepositAssetKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NativeDepositAssetKey is a free data retrieval call binding the contract method 0x33c25326.
//
// Solidity: function nativeDepositAssetKey() view returns(string)
func (_BillingManager *BillingManagerSession) NativeDepositAssetKey() (string, error) {
	return _BillingManager.Contract.NativeDepositAssetKey(&_BillingManager.CallOpts)
}

// NativeDepositAssetKey is a free data retrieval call binding the contract method 0x33c25326.
//
// Solidity: function nativeDepositAssetKey() view returns(string)
func (_BillingManager *BillingManagerCallerSession) NativeDepositAssetKey() (string, error) {
	return _BillingManager.Contract.NativeDepositAssetKey(&_BillingManager.CallOpts)
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

// AddDepositAssets is a paid mutator transaction binding the contract method 0x7c1c6fba.
//
// Solidity: function addDepositAssets((string,(address,uint256,uint256,bool,bool))[] _depositAssetInfoArr) returns()
func (_BillingManager *BillingManagerTransactor) AddDepositAssets(opts *bind.TransactOpts, _depositAssetInfoArr []IBillingManagerDepositAssetInfo) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "addDepositAssets", _depositAssetInfoArr)
}

// AddDepositAssets is a paid mutator transaction binding the contract method 0x7c1c6fba.
//
// Solidity: function addDepositAssets((string,(address,uint256,uint256,bool,bool))[] _depositAssetInfoArr) returns()
func (_BillingManager *BillingManagerSession) AddDepositAssets(_depositAssetInfoArr []IBillingManagerDepositAssetInfo) (*types.Transaction, error) {
	return _BillingManager.Contract.AddDepositAssets(&_BillingManager.TransactOpts, _depositAssetInfoArr)
}

// AddDepositAssets is a paid mutator transaction binding the contract method 0x7c1c6fba.
//
// Solidity: function addDepositAssets((string,(address,uint256,uint256,bool,bool))[] _depositAssetInfoArr) returns()
func (_BillingManager *BillingManagerTransactorSession) AddDepositAssets(_depositAssetInfoArr []IBillingManagerDepositAssetInfo) (*types.Transaction, error) {
	return _BillingManager.Contract.AddDepositAssets(&_BillingManager.TransactOpts, _depositAssetInfoArr)
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

// Deposit is a paid mutator transaction binding the contract method 0x2a48ac9a.
//
// Solidity: function deposit(string _depositAssetKey, address _recipientAddr, uint256 _depositAmount) payable returns()
func (_BillingManager *BillingManagerTransactor) Deposit(opts *bind.TransactOpts, _depositAssetKey string, _recipientAddr common.Address, _depositAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "deposit", _depositAssetKey, _recipientAddr, _depositAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2a48ac9a.
//
// Solidity: function deposit(string _depositAssetKey, address _recipientAddr, uint256 _depositAmount) payable returns()
func (_BillingManager *BillingManagerSession) Deposit(_depositAssetKey string, _recipientAddr common.Address, _depositAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.Deposit(&_BillingManager.TransactOpts, _depositAssetKey, _recipientAddr, _depositAmount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2a48ac9a.
//
// Solidity: function deposit(string _depositAssetKey, address _recipientAddr, uint256 _depositAmount) payable returns()
func (_BillingManager *BillingManagerTransactorSession) Deposit(_depositAssetKey string, _recipientAddr common.Address, _depositAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.Deposit(&_BillingManager.TransactOpts, _depositAssetKey, _recipientAddr, _depositAmount)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x3666cec5.
//
// Solidity: function depositWithPermit(string _depositAssetKey, address _recipientAddr, uint256 _depositAmount, uint256 _sigExpirationTime, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_BillingManager *BillingManagerTransactor) DepositWithPermit(opts *bind.TransactOpts, _depositAssetKey string, _recipientAddr common.Address, _depositAmount *big.Int, _sigExpirationTime *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "depositWithPermit", _depositAssetKey, _recipientAddr, _depositAmount, _sigExpirationTime, _v, _r, _s)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x3666cec5.
//
// Solidity: function depositWithPermit(string _depositAssetKey, address _recipientAddr, uint256 _depositAmount, uint256 _sigExpirationTime, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_BillingManager *BillingManagerSession) DepositWithPermit(_depositAssetKey string, _recipientAddr common.Address, _depositAmount *big.Int, _sigExpirationTime *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _BillingManager.Contract.DepositWithPermit(&_BillingManager.TransactOpts, _depositAssetKey, _recipientAddr, _depositAmount, _sigExpirationTime, _v, _r, _s)
}

// DepositWithPermit is a paid mutator transaction binding the contract method 0x3666cec5.
//
// Solidity: function depositWithPermit(string _depositAssetKey, address _recipientAddr, uint256 _depositAmount, uint256 _sigExpirationTime, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_BillingManager *BillingManagerTransactorSession) DepositWithPermit(_depositAssetKey string, _recipientAddr common.Address, _depositAmount *big.Int, _sigExpirationTime *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _BillingManager.Contract.DepositWithPermit(&_BillingManager.TransactOpts, _depositAssetKey, _recipientAddr, _depositAmount, _sigExpirationTime, _v, _r, _s)
}

// Initialize is a paid mutator transaction binding the contract method 0x55cbf5a5.
//
// Solidity: function initialize(address _registryAddr, address _signerGetterAddress, (string,(address,uint256,uint256,bool,bool)) _nativeDepositAssetInfo) returns()
func (_BillingManager *BillingManagerTransactor) Initialize(opts *bind.TransactOpts, _registryAddr common.Address, _signerGetterAddress common.Address, _nativeDepositAssetInfo IBillingManagerDepositAssetInfo) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "initialize", _registryAddr, _signerGetterAddress, _nativeDepositAssetInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x55cbf5a5.
//
// Solidity: function initialize(address _registryAddr, address _signerGetterAddress, (string,(address,uint256,uint256,bool,bool)) _nativeDepositAssetInfo) returns()
func (_BillingManager *BillingManagerSession) Initialize(_registryAddr common.Address, _signerGetterAddress common.Address, _nativeDepositAssetInfo IBillingManagerDepositAssetInfo) (*types.Transaction, error) {
	return _BillingManager.Contract.Initialize(&_BillingManager.TransactOpts, _registryAddr, _signerGetterAddress, _nativeDepositAssetInfo)
}

// Initialize is a paid mutator transaction binding the contract method 0x55cbf5a5.
//
// Solidity: function initialize(address _registryAddr, address _signerGetterAddress, (string,(address,uint256,uint256,bool,bool)) _nativeDepositAssetInfo) returns()
func (_BillingManager *BillingManagerTransactorSession) Initialize(_registryAddr common.Address, _signerGetterAddress common.Address, _nativeDepositAssetInfo IBillingManagerDepositAssetInfo) (*types.Transaction, error) {
	return _BillingManager.Contract.Initialize(&_BillingManager.TransactOpts, _registryAddr, _signerGetterAddress, _nativeDepositAssetInfo)
}

// LockExecutionFunds is a paid mutator transaction binding the contract method 0x5ec26c46.
//
// Solidity: function lockExecutionFunds(string _depositAssetKey, uint256 _workflowId, uint256 _executionLockedAmount) returns()
func (_BillingManager *BillingManagerTransactor) LockExecutionFunds(opts *bind.TransactOpts, _depositAssetKey string, _workflowId *big.Int, _executionLockedAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "lockExecutionFunds", _depositAssetKey, _workflowId, _executionLockedAmount)
}

// LockExecutionFunds is a paid mutator transaction binding the contract method 0x5ec26c46.
//
// Solidity: function lockExecutionFunds(string _depositAssetKey, uint256 _workflowId, uint256 _executionLockedAmount) returns()
func (_BillingManager *BillingManagerSession) LockExecutionFunds(_depositAssetKey string, _workflowId *big.Int, _executionLockedAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.LockExecutionFunds(&_BillingManager.TransactOpts, _depositAssetKey, _workflowId, _executionLockedAmount)
}

// LockExecutionFunds is a paid mutator transaction binding the contract method 0x5ec26c46.
//
// Solidity: function lockExecutionFunds(string _depositAssetKey, uint256 _workflowId, uint256 _executionLockedAmount) returns()
func (_BillingManager *BillingManagerTransactorSession) LockExecutionFunds(_depositAssetKey string, _workflowId *big.Int, _executionLockedAmount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.LockExecutionFunds(&_BillingManager.TransactOpts, _depositAssetKey, _workflowId, _executionLockedAmount)
}

// NetworkWithdraw is a paid mutator transaction binding the contract method 0x600fe47d.
//
// Solidity: function networkWithdraw(string _depositAssetKey, address _userAddr, uint256 _amountToWithdraw, uint8 _withdrawReason) returns()
func (_BillingManager *BillingManagerTransactor) NetworkWithdraw(opts *bind.TransactOpts, _depositAssetKey string, _userAddr common.Address, _amountToWithdraw *big.Int, _withdrawReason uint8) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "networkWithdraw", _depositAssetKey, _userAddr, _amountToWithdraw, _withdrawReason)
}

// NetworkWithdraw is a paid mutator transaction binding the contract method 0x600fe47d.
//
// Solidity: function networkWithdraw(string _depositAssetKey, address _userAddr, uint256 _amountToWithdraw, uint8 _withdrawReason) returns()
func (_BillingManager *BillingManagerSession) NetworkWithdraw(_depositAssetKey string, _userAddr common.Address, _amountToWithdraw *big.Int, _withdrawReason uint8) (*types.Transaction, error) {
	return _BillingManager.Contract.NetworkWithdraw(&_BillingManager.TransactOpts, _depositAssetKey, _userAddr, _amountToWithdraw, _withdrawReason)
}

// NetworkWithdraw is a paid mutator transaction binding the contract method 0x600fe47d.
//
// Solidity: function networkWithdraw(string _depositAssetKey, address _userAddr, uint256 _amountToWithdraw, uint8 _withdrawReason) returns()
func (_BillingManager *BillingManagerTransactorSession) NetworkWithdraw(_depositAssetKey string, _userAddr common.Address, _amountToWithdraw *big.Int, _withdrawReason uint8) (*types.Transaction, error) {
	return _BillingManager.Contract.NetworkWithdraw(&_BillingManager.TransactOpts, _depositAssetKey, _userAddr, _amountToWithdraw, _withdrawReason)
}

// UpdateDepositAssetEnabledStatus is a paid mutator transaction binding the contract method 0xc172bfd7.
//
// Solidity: function updateDepositAssetEnabledStatus(string _depositAssetKey, bool _newEnabledStatus) returns()
func (_BillingManager *BillingManagerTransactor) UpdateDepositAssetEnabledStatus(opts *bind.TransactOpts, _depositAssetKey string, _newEnabledStatus bool) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "updateDepositAssetEnabledStatus", _depositAssetKey, _newEnabledStatus)
}

// UpdateDepositAssetEnabledStatus is a paid mutator transaction binding the contract method 0xc172bfd7.
//
// Solidity: function updateDepositAssetEnabledStatus(string _depositAssetKey, bool _newEnabledStatus) returns()
func (_BillingManager *BillingManagerSession) UpdateDepositAssetEnabledStatus(_depositAssetKey string, _newEnabledStatus bool) (*types.Transaction, error) {
	return _BillingManager.Contract.UpdateDepositAssetEnabledStatus(&_BillingManager.TransactOpts, _depositAssetKey, _newEnabledStatus)
}

// UpdateDepositAssetEnabledStatus is a paid mutator transaction binding the contract method 0xc172bfd7.
//
// Solidity: function updateDepositAssetEnabledStatus(string _depositAssetKey, bool _newEnabledStatus) returns()
func (_BillingManager *BillingManagerTransactorSession) UpdateDepositAssetEnabledStatus(_depositAssetKey string, _newEnabledStatus bool) (*types.Transaction, error) {
	return _BillingManager.Contract.UpdateDepositAssetEnabledStatus(&_BillingManager.TransactOpts, _depositAssetKey, _newEnabledStatus)
}

// UpdateWorkflowExecutionDiscount is a paid mutator transaction binding the contract method 0xc66ae1e8.
//
// Solidity: function updateWorkflowExecutionDiscount(string _depositAssetKey, uint256 _newWorkflowExecutionDiscount) returns()
func (_BillingManager *BillingManagerTransactor) UpdateWorkflowExecutionDiscount(opts *bind.TransactOpts, _depositAssetKey string, _newWorkflowExecutionDiscount *big.Int) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "updateWorkflowExecutionDiscount", _depositAssetKey, _newWorkflowExecutionDiscount)
}

// UpdateWorkflowExecutionDiscount is a paid mutator transaction binding the contract method 0xc66ae1e8.
//
// Solidity: function updateWorkflowExecutionDiscount(string _depositAssetKey, uint256 _newWorkflowExecutionDiscount) returns()
func (_BillingManager *BillingManagerSession) UpdateWorkflowExecutionDiscount(_depositAssetKey string, _newWorkflowExecutionDiscount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.UpdateWorkflowExecutionDiscount(&_BillingManager.TransactOpts, _depositAssetKey, _newWorkflowExecutionDiscount)
}

// UpdateWorkflowExecutionDiscount is a paid mutator transaction binding the contract method 0xc66ae1e8.
//
// Solidity: function updateWorkflowExecutionDiscount(string _depositAssetKey, uint256 _newWorkflowExecutionDiscount) returns()
func (_BillingManager *BillingManagerTransactorSession) UpdateWorkflowExecutionDiscount(_depositAssetKey string, _newWorkflowExecutionDiscount *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.UpdateWorkflowExecutionDiscount(&_BillingManager.TransactOpts, _depositAssetKey, _newWorkflowExecutionDiscount)
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

// WithdrawAllFunds is a paid mutator transaction binding the contract method 0xf7fa3edf.
//
// Solidity: function withdrawAllFunds(string _depositAssetKey) returns()
func (_BillingManager *BillingManagerTransactor) WithdrawAllFunds(opts *bind.TransactOpts, _depositAssetKey string) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "withdrawAllFunds", _depositAssetKey)
}

// WithdrawAllFunds is a paid mutator transaction binding the contract method 0xf7fa3edf.
//
// Solidity: function withdrawAllFunds(string _depositAssetKey) returns()
func (_BillingManager *BillingManagerSession) WithdrawAllFunds(_depositAssetKey string) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawAllFunds(&_BillingManager.TransactOpts, _depositAssetKey)
}

// WithdrawAllFunds is a paid mutator transaction binding the contract method 0xf7fa3edf.
//
// Solidity: function withdrawAllFunds(string _depositAssetKey) returns()
func (_BillingManager *BillingManagerTransactorSession) WithdrawAllFunds(_depositAssetKey string) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawAllFunds(&_BillingManager.TransactOpts, _depositAssetKey)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x9769b439.
//
// Solidity: function withdrawFunds(string _depositAssetKey, uint256 _amountToWithdraw) returns()
func (_BillingManager *BillingManagerTransactor) WithdrawFunds(opts *bind.TransactOpts, _depositAssetKey string, _amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "withdrawFunds", _depositAssetKey, _amountToWithdraw)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x9769b439.
//
// Solidity: function withdrawFunds(string _depositAssetKey, uint256 _amountToWithdraw) returns()
func (_BillingManager *BillingManagerSession) WithdrawFunds(_depositAssetKey string, _amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawFunds(&_BillingManager.TransactOpts, _depositAssetKey, _amountToWithdraw)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x9769b439.
//
// Solidity: function withdrawFunds(string _depositAssetKey, uint256 _amountToWithdraw) returns()
func (_BillingManager *BillingManagerTransactorSession) WithdrawFunds(_depositAssetKey string, _amountToWithdraw *big.Int) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawFunds(&_BillingManager.TransactOpts, _depositAssetKey, _amountToWithdraw)
}

// WithdrawNetworkRewards is a paid mutator transaction binding the contract method 0x16fb02dd.
//
// Solidity: function withdrawNetworkRewards(string _depositAssetKey) returns()
func (_BillingManager *BillingManagerTransactor) WithdrawNetworkRewards(opts *bind.TransactOpts, _depositAssetKey string) (*types.Transaction, error) {
	return _BillingManager.contract.Transact(opts, "withdrawNetworkRewards", _depositAssetKey)
}

// WithdrawNetworkRewards is a paid mutator transaction binding the contract method 0x16fb02dd.
//
// Solidity: function withdrawNetworkRewards(string _depositAssetKey) returns()
func (_BillingManager *BillingManagerSession) WithdrawNetworkRewards(_depositAssetKey string) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawNetworkRewards(&_BillingManager.TransactOpts, _depositAssetKey)
}

// WithdrawNetworkRewards is a paid mutator transaction binding the contract method 0x16fb02dd.
//
// Solidity: function withdrawNetworkRewards(string _depositAssetKey) returns()
func (_BillingManager *BillingManagerTransactorSession) WithdrawNetworkRewards(_depositAssetKey string) (*types.Transaction, error) {
	return _BillingManager.Contract.WithdrawNetworkRewards(&_BillingManager.TransactOpts, _depositAssetKey)
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

// BillingManagerAssetDepositedIterator is returned from FilterAssetDeposited and is used to iterate over the raw logs and unpacked data for AssetDeposited events raised by the BillingManager contract.
type BillingManagerAssetDepositedIterator struct {
	Event *BillingManagerAssetDeposited // Event containing the contract specifics and raw log

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
func (it *BillingManagerAssetDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerAssetDeposited)
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
		it.Event = new(BillingManagerAssetDeposited)
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
func (it *BillingManagerAssetDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerAssetDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerAssetDeposited represents a AssetDeposited event raised by the BillingManager contract.
type BillingManagerAssetDeposited struct {
	DepositAssetKey  common.Hash
	TokensSender     common.Address
	DepositRecipient common.Address
	DepositAmount    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAssetDeposited is a free log retrieval operation binding the contract event 0xfbc2fad7e5ffeb083133f666db8691af5364707e2b5bf9d1f6f93834433aeca4.
//
// Solidity: event AssetDeposited(string indexed depositAssetKey, address tokensSender, address depositRecipient, uint256 depositAmount)
func (_BillingManager *BillingManagerFilterer) FilterAssetDeposited(opts *bind.FilterOpts, depositAssetKey []string) (*BillingManagerAssetDepositedIterator, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "AssetDeposited", depositAssetKeyRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerAssetDepositedIterator{contract: _BillingManager.contract, event: "AssetDeposited", logs: logs, sub: sub}, nil
}

// WatchAssetDeposited is a free log subscription operation binding the contract event 0xfbc2fad7e5ffeb083133f666db8691af5364707e2b5bf9d1f6f93834433aeca4.
//
// Solidity: event AssetDeposited(string indexed depositAssetKey, address tokensSender, address depositRecipient, uint256 depositAmount)
func (_BillingManager *BillingManagerFilterer) WatchAssetDeposited(opts *bind.WatchOpts, sink chan<- *BillingManagerAssetDeposited, depositAssetKey []string) (event.Subscription, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "AssetDeposited", depositAssetKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerAssetDeposited)
				if err := _BillingManager.contract.UnpackLog(event, "AssetDeposited", log); err != nil {
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

// ParseAssetDeposited is a log parse operation binding the contract event 0xfbc2fad7e5ffeb083133f666db8691af5364707e2b5bf9d1f6f93834433aeca4.
//
// Solidity: event AssetDeposited(string indexed depositAssetKey, address tokensSender, address depositRecipient, uint256 depositAmount)
func (_BillingManager *BillingManagerFilterer) ParseAssetDeposited(log types.Log) (*BillingManagerAssetDeposited, error) {
	event := new(BillingManagerAssetDeposited)
	if err := _BillingManager.contract.UnpackLog(event, "AssetDeposited", log); err != nil {
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

// BillingManagerDepositAssetAddedIterator is returned from FilterDepositAssetAdded and is used to iterate over the raw logs and unpacked data for DepositAssetAdded events raised by the BillingManager contract.
type BillingManagerDepositAssetAddedIterator struct {
	Event *BillingManagerDepositAssetAdded // Event containing the contract specifics and raw log

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
func (it *BillingManagerDepositAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerDepositAssetAdded)
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
		it.Event = new(BillingManagerDepositAssetAdded)
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
func (it *BillingManagerDepositAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerDepositAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerDepositAssetAdded represents a DepositAssetAdded event raised by the BillingManager contract.
type BillingManagerDepositAssetAdded struct {
	DepositAssetKey  string
	DepositAssetData IBillingManagerDepositAssetData
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositAssetAdded is a free log retrieval operation binding the contract event 0x310a01f49fbf43988aa0ce4a44296f5fc9f1fe442769d2416f1965d6a1cd236e.
//
// Solidity: event DepositAssetAdded(string depositAssetKey, (address,uint256,uint256,bool,bool) depositAssetData)
func (_BillingManager *BillingManagerFilterer) FilterDepositAssetAdded(opts *bind.FilterOpts) (*BillingManagerDepositAssetAddedIterator, error) {

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "DepositAssetAdded")
	if err != nil {
		return nil, err
	}
	return &BillingManagerDepositAssetAddedIterator{contract: _BillingManager.contract, event: "DepositAssetAdded", logs: logs, sub: sub}, nil
}

// WatchDepositAssetAdded is a free log subscription operation binding the contract event 0x310a01f49fbf43988aa0ce4a44296f5fc9f1fe442769d2416f1965d6a1cd236e.
//
// Solidity: event DepositAssetAdded(string depositAssetKey, (address,uint256,uint256,bool,bool) depositAssetData)
func (_BillingManager *BillingManagerFilterer) WatchDepositAssetAdded(opts *bind.WatchOpts, sink chan<- *BillingManagerDepositAssetAdded) (event.Subscription, error) {

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "DepositAssetAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerDepositAssetAdded)
				if err := _BillingManager.contract.UnpackLog(event, "DepositAssetAdded", log); err != nil {
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

// ParseDepositAssetAdded is a log parse operation binding the contract event 0x310a01f49fbf43988aa0ce4a44296f5fc9f1fe442769d2416f1965d6a1cd236e.
//
// Solidity: event DepositAssetAdded(string depositAssetKey, (address,uint256,uint256,bool,bool) depositAssetData)
func (_BillingManager *BillingManagerFilterer) ParseDepositAssetAdded(log types.Log) (*BillingManagerDepositAssetAdded, error) {
	event := new(BillingManagerDepositAssetAdded)
	if err := _BillingManager.contract.UnpackLog(event, "DepositAssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerDepositAssetEnabledStatusUpdatedIterator is returned from FilterDepositAssetEnabledStatusUpdated and is used to iterate over the raw logs and unpacked data for DepositAssetEnabledStatusUpdated events raised by the BillingManager contract.
type BillingManagerDepositAssetEnabledStatusUpdatedIterator struct {
	Event *BillingManagerDepositAssetEnabledStatusUpdated // Event containing the contract specifics and raw log

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
func (it *BillingManagerDepositAssetEnabledStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerDepositAssetEnabledStatusUpdated)
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
		it.Event = new(BillingManagerDepositAssetEnabledStatusUpdated)
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
func (it *BillingManagerDepositAssetEnabledStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerDepositAssetEnabledStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerDepositAssetEnabledStatusUpdated represents a DepositAssetEnabledStatusUpdated event raised by the BillingManager contract.
type BillingManagerDepositAssetEnabledStatusUpdated struct {
	DepositAssetKey  common.Hash
	NewEnabledStatus bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositAssetEnabledStatusUpdated is a free log retrieval operation binding the contract event 0x838dfb07f2f9eda916131303b1c93fcdb2d1f0cfae1f651eee30c8bd44091356.
//
// Solidity: event DepositAssetEnabledStatusUpdated(string indexed depositAssetKey, bool newEnabledStatus)
func (_BillingManager *BillingManagerFilterer) FilterDepositAssetEnabledStatusUpdated(opts *bind.FilterOpts, depositAssetKey []string) (*BillingManagerDepositAssetEnabledStatusUpdatedIterator, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "DepositAssetEnabledStatusUpdated", depositAssetKeyRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerDepositAssetEnabledStatusUpdatedIterator{contract: _BillingManager.contract, event: "DepositAssetEnabledStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchDepositAssetEnabledStatusUpdated is a free log subscription operation binding the contract event 0x838dfb07f2f9eda916131303b1c93fcdb2d1f0cfae1f651eee30c8bd44091356.
//
// Solidity: event DepositAssetEnabledStatusUpdated(string indexed depositAssetKey, bool newEnabledStatus)
func (_BillingManager *BillingManagerFilterer) WatchDepositAssetEnabledStatusUpdated(opts *bind.WatchOpts, sink chan<- *BillingManagerDepositAssetEnabledStatusUpdated, depositAssetKey []string) (event.Subscription, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "DepositAssetEnabledStatusUpdated", depositAssetKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerDepositAssetEnabledStatusUpdated)
				if err := _BillingManager.contract.UnpackLog(event, "DepositAssetEnabledStatusUpdated", log); err != nil {
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

// ParseDepositAssetEnabledStatusUpdated is a log parse operation binding the contract event 0x838dfb07f2f9eda916131303b1c93fcdb2d1f0cfae1f651eee30c8bd44091356.
//
// Solidity: event DepositAssetEnabledStatusUpdated(string indexed depositAssetKey, bool newEnabledStatus)
func (_BillingManager *BillingManagerFilterer) ParseDepositAssetEnabledStatusUpdated(log types.Log) (*BillingManagerDepositAssetEnabledStatusUpdated, error) {
	event := new(BillingManagerDepositAssetEnabledStatusUpdated)
	if err := _BillingManager.contract.UnpackLog(event, "DepositAssetEnabledStatusUpdated", log); err != nil {
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
	DepositAssetKey       common.Hash
	WorkflowId            *big.Int
	UserAddr              common.Address
	WorkflowExecutionId   *big.Int
	ExecutionLockedAmount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecutionFundsLocked is a free log retrieval operation binding the contract event 0x318669d1aaaf361c152eec849d7bf340dcac03acfb197842a8fd141df0a408fa.
//
// Solidity: event ExecutionFundsLocked(string indexed depositAssetKey, uint256 indexed workflowId, address indexed userAddr, uint256 workflowExecutionId, uint256 executionLockedAmount)
func (_BillingManager *BillingManagerFilterer) FilterExecutionFundsLocked(opts *bind.FilterOpts, depositAssetKey []string, workflowId []*big.Int, userAddr []common.Address) (*BillingManagerExecutionFundsLockedIterator, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "ExecutionFundsLocked", depositAssetKeyRule, workflowIdRule, userAddrRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerExecutionFundsLockedIterator{contract: _BillingManager.contract, event: "ExecutionFundsLocked", logs: logs, sub: sub}, nil
}

// WatchExecutionFundsLocked is a free log subscription operation binding the contract event 0x318669d1aaaf361c152eec849d7bf340dcac03acfb197842a8fd141df0a408fa.
//
// Solidity: event ExecutionFundsLocked(string indexed depositAssetKey, uint256 indexed workflowId, address indexed userAddr, uint256 workflowExecutionId, uint256 executionLockedAmount)
func (_BillingManager *BillingManagerFilterer) WatchExecutionFundsLocked(opts *bind.WatchOpts, sink chan<- *BillingManagerExecutionFundsLocked, depositAssetKey []string, workflowId []*big.Int, userAddr []common.Address) (event.Subscription, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var workflowIdRule []interface{}
	for _, workflowIdItem := range workflowId {
		workflowIdRule = append(workflowIdRule, workflowIdItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "ExecutionFundsLocked", depositAssetKeyRule, workflowIdRule, userAddrRule)
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

// ParseExecutionFundsLocked is a log parse operation binding the contract event 0x318669d1aaaf361c152eec849d7bf340dcac03acfb197842a8fd141df0a408fa.
//
// Solidity: event ExecutionFundsLocked(string indexed depositAssetKey, uint256 indexed workflowId, address indexed userAddr, uint256 workflowExecutionId, uint256 executionLockedAmount)
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

// BillingManagerNetworkWithdrawCompletedIterator is returned from FilterNetworkWithdrawCompleted and is used to iterate over the raw logs and unpacked data for NetworkWithdrawCompleted events raised by the BillingManager contract.
type BillingManagerNetworkWithdrawCompletedIterator struct {
	Event *BillingManagerNetworkWithdrawCompleted // Event containing the contract specifics and raw log

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
func (it *BillingManagerNetworkWithdrawCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerNetworkWithdrawCompleted)
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
		it.Event = new(BillingManagerNetworkWithdrawCompleted)
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
func (it *BillingManagerNetworkWithdrawCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerNetworkWithdrawCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerNetworkWithdrawCompleted represents a NetworkWithdrawCompleted event raised by the BillingManager contract.
type BillingManagerNetworkWithdrawCompleted struct {
	DepositAssetKey  common.Hash
	UserAddr         common.Address
	WithdrawReason   uint8
	AmountToWithdraw *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNetworkWithdrawCompleted is a free log retrieval operation binding the contract event 0xf913ff1a24ed376db7f7964cb178db8a9ec8c5462a747385218f3610003724d9.
//
// Solidity: event NetworkWithdrawCompleted(string indexed depositAssetKey, address indexed userAddr, uint8 indexed withdrawReason, uint256 amountToWithdraw)
func (_BillingManager *BillingManagerFilterer) FilterNetworkWithdrawCompleted(opts *bind.FilterOpts, depositAssetKey []string, userAddr []common.Address, withdrawReason []uint8) (*BillingManagerNetworkWithdrawCompletedIterator, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}
	var withdrawReasonRule []interface{}
	for _, withdrawReasonItem := range withdrawReason {
		withdrawReasonRule = append(withdrawReasonRule, withdrawReasonItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "NetworkWithdrawCompleted", depositAssetKeyRule, userAddrRule, withdrawReasonRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerNetworkWithdrawCompletedIterator{contract: _BillingManager.contract, event: "NetworkWithdrawCompleted", logs: logs, sub: sub}, nil
}

// WatchNetworkWithdrawCompleted is a free log subscription operation binding the contract event 0xf913ff1a24ed376db7f7964cb178db8a9ec8c5462a747385218f3610003724d9.
//
// Solidity: event NetworkWithdrawCompleted(string indexed depositAssetKey, address indexed userAddr, uint8 indexed withdrawReason, uint256 amountToWithdraw)
func (_BillingManager *BillingManagerFilterer) WatchNetworkWithdrawCompleted(opts *bind.WatchOpts, sink chan<- *BillingManagerNetworkWithdrawCompleted, depositAssetKey []string, userAddr []common.Address, withdrawReason []uint8) (event.Subscription, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}
	var withdrawReasonRule []interface{}
	for _, withdrawReasonItem := range withdrawReason {
		withdrawReasonRule = append(withdrawReasonRule, withdrawReasonItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "NetworkWithdrawCompleted", depositAssetKeyRule, userAddrRule, withdrawReasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerNetworkWithdrawCompleted)
				if err := _BillingManager.contract.UnpackLog(event, "NetworkWithdrawCompleted", log); err != nil {
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

// ParseNetworkWithdrawCompleted is a log parse operation binding the contract event 0xf913ff1a24ed376db7f7964cb178db8a9ec8c5462a747385218f3610003724d9.
//
// Solidity: event NetworkWithdrawCompleted(string indexed depositAssetKey, address indexed userAddr, uint8 indexed withdrawReason, uint256 amountToWithdraw)
func (_BillingManager *BillingManagerFilterer) ParseNetworkWithdrawCompleted(log types.Log) (*BillingManagerNetworkWithdrawCompleted, error) {
	event := new(BillingManagerNetworkWithdrawCompleted)
	if err := _BillingManager.contract.UnpackLog(event, "NetworkWithdrawCompleted", log); err != nil {
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
	DepositAssetKey common.Hash
	RecipientAddr   common.Address
	RewardsAmount   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardsWithdrawn is a free log retrieval operation binding the contract event 0x4da529f8b18684c34deb8621c7d4ab35fe1daea1c726cdb6783c93c095d2651a.
//
// Solidity: event RewardsWithdrawn(string indexed depositAssetKey, address indexed recipientAddr, uint256 rewardsAmount)
func (_BillingManager *BillingManagerFilterer) FilterRewardsWithdrawn(opts *bind.FilterOpts, depositAssetKey []string, recipientAddr []common.Address) (*BillingManagerRewardsWithdrawnIterator, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "RewardsWithdrawn", depositAssetKeyRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerRewardsWithdrawnIterator{contract: _BillingManager.contract, event: "RewardsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRewardsWithdrawn is a free log subscription operation binding the contract event 0x4da529f8b18684c34deb8621c7d4ab35fe1daea1c726cdb6783c93c095d2651a.
//
// Solidity: event RewardsWithdrawn(string indexed depositAssetKey, address indexed recipientAddr, uint256 rewardsAmount)
func (_BillingManager *BillingManagerFilterer) WatchRewardsWithdrawn(opts *bind.WatchOpts, sink chan<- *BillingManagerRewardsWithdrawn, depositAssetKey []string, recipientAddr []common.Address) (event.Subscription, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "RewardsWithdrawn", depositAssetKeyRule, recipientAddrRule)
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

// ParseRewardsWithdrawn is a log parse operation binding the contract event 0x4da529f8b18684c34deb8621c7d4ab35fe1daea1c726cdb6783c93c095d2651a.
//
// Solidity: event RewardsWithdrawn(string indexed depositAssetKey, address indexed recipientAddr, uint256 rewardsAmount)
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
	DepositAssetKey common.Hash
	UserAddr        common.Address
	WithdrawnAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUserFundsWithdrawn is a free log retrieval operation binding the contract event 0xe0dd7f2d0bf336d345c1e8051edf7f753d4bfdc13d35cf30351267de2aa905c8.
//
// Solidity: event UserFundsWithdrawn(string indexed depositAssetKey, address indexed userAddr, uint256 withdrawnAmount)
func (_BillingManager *BillingManagerFilterer) FilterUserFundsWithdrawn(opts *bind.FilterOpts, depositAssetKey []string, userAddr []common.Address) (*BillingManagerUserFundsWithdrawnIterator, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "UserFundsWithdrawn", depositAssetKeyRule, userAddrRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerUserFundsWithdrawnIterator{contract: _BillingManager.contract, event: "UserFundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchUserFundsWithdrawn is a free log subscription operation binding the contract event 0xe0dd7f2d0bf336d345c1e8051edf7f753d4bfdc13d35cf30351267de2aa905c8.
//
// Solidity: event UserFundsWithdrawn(string indexed depositAssetKey, address indexed userAddr, uint256 withdrawnAmount)
func (_BillingManager *BillingManagerFilterer) WatchUserFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *BillingManagerUserFundsWithdrawn, depositAssetKey []string, userAddr []common.Address) (event.Subscription, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}
	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "UserFundsWithdrawn", depositAssetKeyRule, userAddrRule)
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

// ParseUserFundsWithdrawn is a log parse operation binding the contract event 0xe0dd7f2d0bf336d345c1e8051edf7f753d4bfdc13d35cf30351267de2aa905c8.
//
// Solidity: event UserFundsWithdrawn(string indexed depositAssetKey, address indexed userAddr, uint256 withdrawnAmount)
func (_BillingManager *BillingManagerFilterer) ParseUserFundsWithdrawn(log types.Log) (*BillingManagerUserFundsWithdrawn, error) {
	event := new(BillingManagerUserFundsWithdrawn)
	if err := _BillingManager.contract.UnpackLog(event, "UserFundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BillingManagerWorkflowExecutionDiscountUpdatedIterator is returned from FilterWorkflowExecutionDiscountUpdated and is used to iterate over the raw logs and unpacked data for WorkflowExecutionDiscountUpdated events raised by the BillingManager contract.
type BillingManagerWorkflowExecutionDiscountUpdatedIterator struct {
	Event *BillingManagerWorkflowExecutionDiscountUpdated // Event containing the contract specifics and raw log

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
func (it *BillingManagerWorkflowExecutionDiscountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BillingManagerWorkflowExecutionDiscountUpdated)
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
		it.Event = new(BillingManagerWorkflowExecutionDiscountUpdated)
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
func (it *BillingManagerWorkflowExecutionDiscountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BillingManagerWorkflowExecutionDiscountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BillingManagerWorkflowExecutionDiscountUpdated represents a WorkflowExecutionDiscountUpdated event raised by the BillingManager contract.
type BillingManagerWorkflowExecutionDiscountUpdated struct {
	DepositAssetKey               common.Hash
	NewWorkflowExecutionDiscount  *big.Int
	PrevWorkflowExecutionDiscount *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterWorkflowExecutionDiscountUpdated is a free log retrieval operation binding the contract event 0x70c21b53e4a29c6a334ae8d8d80a885023e4749ee873d6594b3f6cbb66140b6b.
//
// Solidity: event WorkflowExecutionDiscountUpdated(string indexed depositAssetKey, uint256 newWorkflowExecutionDiscount, uint256 prevWorkflowExecutionDiscount)
func (_BillingManager *BillingManagerFilterer) FilterWorkflowExecutionDiscountUpdated(opts *bind.FilterOpts, depositAssetKey []string) (*BillingManagerWorkflowExecutionDiscountUpdatedIterator, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}

	logs, sub, err := _BillingManager.contract.FilterLogs(opts, "WorkflowExecutionDiscountUpdated", depositAssetKeyRule)
	if err != nil {
		return nil, err
	}
	return &BillingManagerWorkflowExecutionDiscountUpdatedIterator{contract: _BillingManager.contract, event: "WorkflowExecutionDiscountUpdated", logs: logs, sub: sub}, nil
}

// WatchWorkflowExecutionDiscountUpdated is a free log subscription operation binding the contract event 0x70c21b53e4a29c6a334ae8d8d80a885023e4749ee873d6594b3f6cbb66140b6b.
//
// Solidity: event WorkflowExecutionDiscountUpdated(string indexed depositAssetKey, uint256 newWorkflowExecutionDiscount, uint256 prevWorkflowExecutionDiscount)
func (_BillingManager *BillingManagerFilterer) WatchWorkflowExecutionDiscountUpdated(opts *bind.WatchOpts, sink chan<- *BillingManagerWorkflowExecutionDiscountUpdated, depositAssetKey []string) (event.Subscription, error) {

	var depositAssetKeyRule []interface{}
	for _, depositAssetKeyItem := range depositAssetKey {
		depositAssetKeyRule = append(depositAssetKeyRule, depositAssetKeyItem)
	}

	logs, sub, err := _BillingManager.contract.WatchLogs(opts, "WorkflowExecutionDiscountUpdated", depositAssetKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BillingManagerWorkflowExecutionDiscountUpdated)
				if err := _BillingManager.contract.UnpackLog(event, "WorkflowExecutionDiscountUpdated", log); err != nil {
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

// ParseWorkflowExecutionDiscountUpdated is a log parse operation binding the contract event 0x70c21b53e4a29c6a334ae8d8d80a885023e4749ee873d6594b3f6cbb66140b6b.
//
// Solidity: event WorkflowExecutionDiscountUpdated(string indexed depositAssetKey, uint256 newWorkflowExecutionDiscount, uint256 prevWorkflowExecutionDiscount)
func (_BillingManager *BillingManagerFilterer) ParseWorkflowExecutionDiscountUpdated(log types.Log) (*BillingManagerWorkflowExecutionDiscountUpdated, error) {
	event := new(BillingManagerWorkflowExecutionDiscountUpdated)
	if err := _BillingManager.contract.UnpackLog(event, "WorkflowExecutionDiscountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
