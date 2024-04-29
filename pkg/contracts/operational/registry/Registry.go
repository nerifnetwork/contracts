// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// IRegistryBaseWorkflowInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryBaseWorkflowInfo struct {
	Id    *big.Int
	Owner common.Address
}

// IRegistryDepositAssetInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryDepositAssetInfo struct {
	DepositAssetKey        string
	DepositAssetTotalSpent *big.Int
}

// IRegistryGatewayInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryGatewayInfo struct {
	GatewayOwner common.Address
	Gateway      common.Address
}

// IRegistryRegisterWorkflowInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryRegisterWorkflowInfo struct {
	Id             *big.Int
	WorkflowOwner  common.Address
	RequireGateway bool
	DeployGateway  bool
}

// IRegistryWorkflowInfo is an auto generated low-level Go binding around an user-defined struct.
type IRegistryWorkflowInfo struct {
	BaseInfo          IRegistryBaseWorkflowInfo
	DepositAssetKeys  []string
	DepositAssetsInfo []IRegistryDepositAssetInfo
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Performance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"WorkflowRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deployAndSetGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"}],\"name\":\"getBaseWorkflowInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.BaseWorkflowInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getGatewaysInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gatewayOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.GatewayInfo[]\",\"name\":\"_gatewaysInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalGatewaysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWorkflowsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowDepositAssetKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_depositAssetKeys\",\"type\":\"string[]\"}],\"name\":\"getWorkflowDepositAssetsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"depositAssetTotalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.DepositAssetInfo[]\",\"name\":\"_depositAssetsArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.BaseWorkflowInfo\",\"name\":\"baseInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"depositAssetKeys\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"depositAssetTotalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.DepositAssetInfo[]\",\"name\":\"depositAssetsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structIRegistry.WorkflowInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getWorkflowOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getWorkflowsInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.BaseWorkflowInfo\",\"name\":\"baseInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"depositAssetKeys\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"depositAssetTotalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.DepositAssetInfo[]\",\"name\":\"depositAssetsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structIRegistry.WorkflowInfo[]\",\"name\":\"_workflowsInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_maxWorkflowsPerAccount\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isWorkflowRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxWorkflowsPerAccount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"requireGateway\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deployGateway\",\"type\":\"bool\"}],\"internalType\":\"structIRegistry.RegisterWorkflowInfo[]\",\"name\":\"_registerWorkflowInfoArr\",\"type\":\"tuple[]\"}],\"name\":\"registerWorkflows\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newMaxWorkflowsPerAccount\",\"type\":\"uint16\"}],\"name\":\"setMaxWorkflowsPerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionAmount\",\"type\":\"uint256\"}],\"name\":\"updateWorkflowTotalSpent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"workflowsPerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612248806100206000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063a3190a36116100c3578063c48223191161007c578063c48223191461036b578063c5fd64d814610373578063d69cd2751461037b578063d6bb757b146103a7578063df7459f0146103ba578063e1d1dffe146103da57600080fd5b8063a3190a3614610286578063abad5c1a146102a6578063ae09340f146102ce578063b8f9c0f6146102e1578063bda009fe1461031f578063c17fc3eb1461034b57600080fd5b8063691304511161011557806369130451146101ec5780636b699545146101ff5780638cb941cc1461022d5780638cdfb69a1461024057806390646b4a1461025357806395f5dae41461026657600080fd5b80630bdeec8f14610152578063137509461461017b5780633e3b5b191461019057806344ab987d146101b95780634b322560146101cc575b600080fd5b6101656101603660046118d5565b6103e2565b604051610172919061193e565b60405180910390f35b61018e6101893660046119a0565b610405565b005b6000805160206121f3833981519152545b6040516001600160a01b039091168152602001610172565b61018e6101c73660046119a0565b61052f565b6101df6101da3660046118d5565b610559565b60405161017291906119c4565b61018e6101fa366004611a98565b6105a0565b61021f61020d366004611afc565b60066020526000908152604090205481565b604051908152602001610172565b61018e61023b366004611afc565b6106f4565b61018e61024e366004611b19565b610715565b61018e610261366004611afc565b610c01565b610279610274366004611b8e565b610c0b565b6040516101729190611cbe565b610299610294366004611b8e565b610d1d565b6040516101729190611d13565b6002546102bb90600160a01b900461ffff1681565b60405161ffff9091168152602001610172565b61018e6102dc366004611d82565b610e24565b61030f6102ef3660046118d5565b6000908152600760205260409020600101546001600160a01b0316151590565b6040519015158152602001610172565b6101a161032d366004611afc565b6001600160a01b039081166000908152600860205260409020541690565b61035e610359366004611dd0565b610f03565b6040516101729190611e9f565b61021f611024565b6101a1611035565b6101a16103893660046118d5565b6000908152600760205260409020600101546001600160a01b031690565b61018e6103b5366004611eb2565b611040565b6103cd6103c83660046118d5565b611214565b6040516101729190611f4d565b60035461021f565b60008181526007602052604090206060906103ff90600201611288565b92915050565b600054610100900460ff16158080156104255750600054600160ff909116105b8061043f5750303b15801561043f575060005460ff166001145b6104a75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156104ca576000805461ff0019166101001790555b6002805461ffff60a01b1916600160a01b61ffff851602179055801561052b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b610537611365565b6002805461ffff909216600160a01b0261ffff60a01b19909216919091179055565b604080518082019091526000808252602082015250600090815260076020908152604091829020825180840190935280548352600101546001600160a01b03169082015290565b6105a8611429565b600082905080600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663fabf588e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610612573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106369190611f60565b600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b03166348197f776040518163ffffffff1660e01b8152600401602060405180830381865afa15801561069a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106be9190611f60565b600280546001600160a01b0319166001600160a01b039290921691909117905550336000805160206121f3833981519152555050565b6106fc611429565b610712816000805160206121f383398151915255565b50565b60008060029054906101000a90046001600160a01b03166001600160a01b031663b834f6fb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610769573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078d9190611f8b565b905060008060029054906101000a90046001600160a01b03166001600160a01b0316637ac3c02f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108079190611f60565b90508180156108225750600254600160a01b900461ffff1615155b156108b15760025433600090815260066020526040902054600160a01b90910461ffff1690610852908590611fbe565b11156108b15760405162461bcd60e51b815260206004820152602860248201527f52656769737472793a2072656163686564206d617820776f726b666c6f777320604482015267636170616369747960c01b606482015260840161049e565b60005b83811015610bfa57368585838181106108cf576108cf611fd1565b9050608002019050836108ed576001600160a01b0383163314610909565b336108fe6040830160208401611afc565b6001600160a01b0316145b6109555760405162461bcd60e51b815260206004820181905260248201527f52656769737472793a206e6f7420612073656e646572206f72207369676e6572604482015260640161049e565b80356000908152600760205260409020600101546001600160a01b0316156109cb5760405162461bcd60e51b8152602060048201526024808201527f52656769737472793a20776f726b666c6f7720696420616c72656164792065786044820152636973747360e01b606482015260840161049e565b6109db6060820160408301611fe7565b15610ab25760006008816109f56040850160208601611afc565b6001600160a01b039081168252602082019290925260400160002054169050610a246080830160608401611fe7565b8015610a3757506001600160a01b038116155b15610a5a57610a54610a4f6040840160208501611afc565b6114ad565b50610ab0565b6001600160a01b038116610ab05760405162461bcd60e51b815260206004820152601b60248201527f52656769737472793a2067617465776179206e6f7420666f756e640000000000604482015260640161049e565b505b604051806040016040528082600001358152602001826020016020810190610ada9190611afc565b6001600160a01b03908116909152823560009081526007602090815260408083208551815594820151600190950180546001600160a01b03191695909416949094179092556006929091610b32918501908501611afc565b6001600160a01b0316815260208101919091526040016000908120805491610b5983612004565b90915550506003805460018101825560009190915281357fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101557f756019c5646bbc7e222ed197b047a6cd3c7fe80b10551475f1925d4c9ee4b8ee610bc66040830160208401611afc565b604080516001600160a01b039092168252833560208301520160405180910390a15080610bf281612004565b9150506108b4565b5050505050565b6107123382611529565b60606000610c1f60038054905085856115ba565b9050610c2b848261201d565b67ffffffffffffffff811115610c4357610c436119f9565b604051908082528060200260200182016040528015610ca957816020015b610c966040805160a0810190915260006060820181815260808301919091528190815260200160608152602001606081525090565b815260200190600190039081610c615790505b509150835b81811015610d1557610cdc60038281548110610ccc57610ccc611fd1565b9060005260206000200154611214565b83610ce7878461201d565b81518110610cf757610cf7611fd1565b60200260200101819052508080610d0d90612004565b915050610cae565b505092915050565b60606000610d35610d2e60046115e5565b85856115ba565b9050610d41848261201d565b67ffffffffffffffff811115610d5957610d596119f9565b604051908082528060200260200182016040528015610d9e57816020015b6040805180820190915260008082526020820152815260200190600190039081610d775790505b509150835b81811015610d15576000610db86004836115ef565b6040805180820182526001600160a01b0380841680835260009081526008602090815293902054169181019190915290915084610df5888561201d565b81518110610e0557610e05611fd1565b6020026020010181905250508080610e1c90612004565b915050610da3565b81610e2e816115fb565b6002546001600160a01b03163314610e9a5760405162461bcd60e51b815260206004820152602960248201527f52656769737472793a2073656e646572206973206e6f7420612062696c6c696e604482015268339036b0b730b3b2b960b91b606482015260840161049e565b6000838152600760205260409020610eb5906002018561166c565b50816007600085815260200190815260200160002060040185604051610edb9190612030565b90815260200160405180910390206000828254610ef89190611fbe565b909155505050505050565b6060815167ffffffffffffffff811115610f1f57610f1f6119f9565b604051908082528060200260200182016040528015610f6557816020015b604080518082019091526060815260006020820152815260200190600190039081610f3d5790505b50905060005b825181101561101d576040518060400160405280848381518110610f9157610f91611fd1565b6020026020010151815260200160076000878152602001908152602001600020600401858481518110610fc657610fc6611fd1565b6020026020010151604051610fdb9190612030565b908152602001604051809103902054815250828281518110610fff57610fff611fd1565b6020026020010181905250808061101590612004565b915050610f6b565b5092915050565b600061103060046115e5565b905090565b6000611030336114ad565b611048611365565b84611052816115fb565b6001600160a01b03821630036110b65760405162461bcd60e51b8152602060048201526024808201527f52656769737472793a206f7065726174696f6e206973206e6f74207065726d696044820152631d1d195960e21b606482015260840161049e565b6000868152600760209081526040808320600101546001600160a01b039081168452600890925290912054168061112f5760405162461bcd60e51b815260206004820152601e60248201527f52656769737472793a207a65726f206761746577617920616464726573730000604482015260640161049e565b60006111cd87836040518060400160405280601e81526020017f706572666f726d2875696e743235362c616464726573732c62797465732900008152508b888b8b604051602401611183949392919061204c565b60408051601f19818403018152908290529161119e91612030565b6040519081900390206020820180516001600160e01b03166001600160e01b03199092169190911790526116ce565b604080518a815282151560208201529192507f8e67dfc40b26eb841f1be843b96e6d300ff34a0f3f032496584530cd6544c3fb910160405180910390a15050505050505050565b6112496040805160a0810190915260006060820181815260808301919091528190815260200160608152602001606081525090565b6000611254836103e2565b9050604051806060016040528061126a85610559565b815260200182815260200161127f8584610f03565b90529392505050565b606081600001805480602002602001604051908101604052809291908181526020016000905b8282101561135a5783829060005260206000200180546112cd90612093565b80601f01602080910402602001604051908101604052809291908181526020018280546112f990612093565b80156113465780601f1061131b57610100808354040283529160200191611346565b820191906000526020600020905b81548152906001019060200180831161132957829003601f168201915b5050505050815260200190600101906112ae565b505050509050919050565b60005460408051637ac3c02f60e01b8152905133926201000090046001600160a01b031691637ac3c02f9160048083019260209291908290030181865afa1580156113b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113d89190611f60565b6001600160a01b0316146114275760405162461bcd60e51b81526020600482015260166024820152752932b3b4b9ba393c9d102737ba10309039b4b3b732b960511b604482015260640161049e565b565b60006114416000805160206121f38339815191525490565b90506001600160a01b038116158061146157506001600160a01b03811633145b6107125760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f72000000000000604482015260640161049e565b60015460405163257e6bf960e21b81526001600160a01b03838116600483015260009283929116906395f9afe4906024016020604051808303816000875af11580156114fd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115219190611f60565b90506103ff83825b6001600160a01b038116156115495761154360048361171a565b50611556565b61155460048361172f565b505b6001600160a01b0382811660008181526008602090815260409182902080546001600160a01b031916948616948517905581519283528201929092527f812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba3549101610522565b60006115c68284611fbe565b9050838111156115d35750825b808311156115de5750815b9392505050565b60006103ff825490565b60006115de8383611744565b6000818152600760205260409020600101546001600160a01b03166107125760405162461bcd60e51b815260206004820152602160248201527f52656769737472793a20776f726b666c6f7720646f6573206e6f7420657869736044820152601d60fa1b606482015260840161049e565b6000611678838361176e565b6116c657825460018101845560008481526020902001611698838261211c565b50825460405160018501906116ae908590612030565b908152604051908190036020019020555060016103ff565b5060006103ff565b60005a6113888110156116e057600080fd5b6113888103905084604082048203116116f857600080fd5b50823b61170457600080fd5b60008083516020850160008789f1949350505050565b60006115de836001600160a01b03841661179b565b60006115de836001600160a01b0384166117e2565b600082600001828154811061175b5761175b611fd1565b9060005260206000200154905092915050565b600082600101826040516117829190612030565b9081526040519081900360200190205415159392505050565b60008181526001830160205260408120546116c6575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556103ff565b600081815260018301602052604081205480156118cb57600061180660018361201d565b855490915060009061181a9060019061201d565b905081811461187f57600086600001828154811061183a5761183a611fd1565b906000526020600020015490508087600001848154811061185d5761185d611fd1565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611890576118906121dc565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506103ff565b60009150506103ff565b6000602082840312156118e757600080fd5b5035919050565b60005b838110156119095781810151838201526020016118f1565b50506000910152565b6000815180845261192a8160208601602086016118ee565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561199357603f19888603018452611981858351611912565b94509285019290850190600101611965565b5092979650505050505050565b6000602082840312156119b257600080fd5b813561ffff811681146115de57600080fd5b815181526020808301516001600160a01b031690820152604081016103ff565b6001600160a01b038116811461071257600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611a3857611a386119f9565b604052919050565b600067ffffffffffffffff831115611a5a57611a5a6119f9565b611a6d601f8401601f1916602001611a0f565b9050828152838383011115611a8157600080fd5b828260208301376000602084830101529392505050565b60008060408385031215611aab57600080fd5b8235611ab6816119e4565b9150602083013567ffffffffffffffff811115611ad257600080fd5b8301601f81018513611ae357600080fd5b611af285823560208401611a40565b9150509250929050565b600060208284031215611b0e57600080fd5b81356115de816119e4565b60008060208385031215611b2c57600080fd5b823567ffffffffffffffff80821115611b4457600080fd5b818501915085601f830112611b5857600080fd5b813581811115611b6757600080fd5b8660208260071b8501011115611b7c57600080fd5b60209290920196919550909350505050565b60008060408385031215611ba157600080fd5b50508035926020909101359150565b600081518084526020808501808196508360051b8101915082860160005b85811015611c0e578284038952815160408151818752611bf082880182611912565b92880151968801969096525098850198935090840190600101611bce565b5091979650505050505050565b600060808301611c3f848451805182526020908101516001600160a01b0316910152565b6020808401516080604087015282815180855260a08801915060a08160051b8901019450838301925060005b81811015611c9957609f19898703018352611c87868551611912565b95509284019291840191600101611c6b565b505050505060408301518482036060860152611cb58282611bb0565b95945050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561199357603f19888603018452611d01858351611c1b565b94509285019290850190600101611ce5565b602080825282518282018190526000919060409081850190868401855b82811015611c0e57815180516001600160a01b0390811686529087015116868501529284019290850190600101611d30565b600082601f830112611d7357600080fd5b6115de83833560208501611a40565b600080600060608486031215611d9757600080fd5b833567ffffffffffffffff811115611dae57600080fd5b611dba86828701611d62565b9660208601359650604090950135949350505050565b60008060408385031215611de357600080fd5b8235915060208084013567ffffffffffffffff80821115611e0357600080fd5b818601915086601f830112611e1757600080fd5b813581811115611e2957611e296119f9565b8060051b611e38858201611a0f565b918252838101850191858101908a841115611e5257600080fd5b86860192505b83831015611e8e57823585811115611e705760008081fd5b611e7e8c89838a0101611d62565b8352509186019190860190611e58565b809750505050505050509250929050565b6020815260006115de6020830184611bb0565b600080600080600060808688031215611eca57600080fd5b8535945060208601359350604086013567ffffffffffffffff80821115611ef057600080fd5b818801915088601f830112611f0457600080fd5b813581811115611f1357600080fd5b896020828501011115611f2557600080fd5b6020830195508094505050506060860135611f3f816119e4565b809150509295509295909350565b6020815260006115de6020830184611c1b565b600060208284031215611f7257600080fd5b81516115de816119e4565b801515811461071257600080fd5b600060208284031215611f9d57600080fd5b81516115de81611f7d565b634e487b7160e01b600052601160045260246000fd5b808201808211156103ff576103ff611fa8565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611ff957600080fd5b81356115de81611f7d565b60006001820161201657612016611fa8565b5060010190565b818103818111156103ff576103ff611fa8565b600082516120428184602087016118ee565b9190910192915050565b8481526001600160a01b03841660208201526060604082018190528101829052818360808301376000818301608090810191909152601f909201601f191601019392505050565b600181811c908216806120a757607f821691505b6020821081036120c757634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561211757600081815260208120601f850160051c810160208610156120f45750805b601f850160051c820191505b8181101561211357828155600101612100565b5050505b505050565b815167ffffffffffffffff811115612136576121366119f9565b61214a816121448454612093565b846120cd565b602080601f83116001811461217f57600084156121675750858301515b600019600386901b1c1916600185901b178555612113565b600085815260208120601f198616915b828110156121ae5788860151825594840194600190910190840161218f565b50858210156121cc5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603160045260246000fdfe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a2646970667358221220c3c31b46222f3411ba1cfc7a8eac0a4f69d08e8ebb0f380d3a61a6a84bf9400964736f6c63430008120033",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// GetBaseWorkflowInfo is a free data retrieval call binding the contract method 0x4b322560.
//
// Solidity: function getBaseWorkflowInfo(uint256 _workflowId) view returns((uint256,address))
func (_Registry *RegistryCaller) GetBaseWorkflowInfo(opts *bind.CallOpts, _workflowId *big.Int) (IRegistryBaseWorkflowInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getBaseWorkflowInfo", _workflowId)

	if err != nil {
		return *new(IRegistryBaseWorkflowInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryBaseWorkflowInfo)).(*IRegistryBaseWorkflowInfo)

	return out0, err

}

// GetBaseWorkflowInfo is a free data retrieval call binding the contract method 0x4b322560.
//
// Solidity: function getBaseWorkflowInfo(uint256 _workflowId) view returns((uint256,address))
func (_Registry *RegistrySession) GetBaseWorkflowInfo(_workflowId *big.Int) (IRegistryBaseWorkflowInfo, error) {
	return _Registry.Contract.GetBaseWorkflowInfo(&_Registry.CallOpts, _workflowId)
}

// GetBaseWorkflowInfo is a free data retrieval call binding the contract method 0x4b322560.
//
// Solidity: function getBaseWorkflowInfo(uint256 _workflowId) view returns((uint256,address))
func (_Registry *RegistryCallerSession) GetBaseWorkflowInfo(_workflowId *big.Int) (IRegistryBaseWorkflowInfo, error) {
	return _Registry.Contract.GetBaseWorkflowInfo(&_Registry.CallOpts, _workflowId)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _userAddr) view returns(address)
func (_Registry *RegistryCaller) GetGateway(opts *bind.CallOpts, _userAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getGateway", _userAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _userAddr) view returns(address)
func (_Registry *RegistrySession) GetGateway(_userAddr common.Address) (common.Address, error) {
	return _Registry.Contract.GetGateway(&_Registry.CallOpts, _userAddr)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _userAddr) view returns(address)
func (_Registry *RegistryCallerSession) GetGateway(_userAddr common.Address) (common.Address, error) {
	return _Registry.Contract.GetGateway(&_Registry.CallOpts, _userAddr)
}

// GetGatewaysInfo is a free data retrieval call binding the contract method 0xa3190a36.
//
// Solidity: function getGatewaysInfo(uint256 _offset, uint256 _limit) view returns((address,address)[] _gatewaysInfoArr)
func (_Registry *RegistryCaller) GetGatewaysInfo(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]IRegistryGatewayInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getGatewaysInfo", _offset, _limit)

	if err != nil {
		return *new([]IRegistryGatewayInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRegistryGatewayInfo)).(*[]IRegistryGatewayInfo)

	return out0, err

}

// GetGatewaysInfo is a free data retrieval call binding the contract method 0xa3190a36.
//
// Solidity: function getGatewaysInfo(uint256 _offset, uint256 _limit) view returns((address,address)[] _gatewaysInfoArr)
func (_Registry *RegistrySession) GetGatewaysInfo(_offset *big.Int, _limit *big.Int) ([]IRegistryGatewayInfo, error) {
	return _Registry.Contract.GetGatewaysInfo(&_Registry.CallOpts, _offset, _limit)
}

// GetGatewaysInfo is a free data retrieval call binding the contract method 0xa3190a36.
//
// Solidity: function getGatewaysInfo(uint256 _offset, uint256 _limit) view returns((address,address)[] _gatewaysInfoArr)
func (_Registry *RegistryCallerSession) GetGatewaysInfo(_offset *big.Int, _limit *big.Int) ([]IRegistryGatewayInfo, error) {
	return _Registry.Contract.GetGatewaysInfo(&_Registry.CallOpts, _offset, _limit)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Registry *RegistryCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Registry *RegistrySession) GetInjector() (common.Address, error) {
	return _Registry.Contract.GetInjector(&_Registry.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Registry *RegistryCallerSession) GetInjector() (common.Address, error) {
	return _Registry.Contract.GetInjector(&_Registry.CallOpts)
}

// GetTotalGatewaysCount is a free data retrieval call binding the contract method 0xc4822319.
//
// Solidity: function getTotalGatewaysCount() view returns(uint256)
func (_Registry *RegistryCaller) GetTotalGatewaysCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getTotalGatewaysCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalGatewaysCount is a free data retrieval call binding the contract method 0xc4822319.
//
// Solidity: function getTotalGatewaysCount() view returns(uint256)
func (_Registry *RegistrySession) GetTotalGatewaysCount() (*big.Int, error) {
	return _Registry.Contract.GetTotalGatewaysCount(&_Registry.CallOpts)
}

// GetTotalGatewaysCount is a free data retrieval call binding the contract method 0xc4822319.
//
// Solidity: function getTotalGatewaysCount() view returns(uint256)
func (_Registry *RegistryCallerSession) GetTotalGatewaysCount() (*big.Int, error) {
	return _Registry.Contract.GetTotalGatewaysCount(&_Registry.CallOpts)
}

// GetTotalWorkflowsCount is a free data retrieval call binding the contract method 0xe1d1dffe.
//
// Solidity: function getTotalWorkflowsCount() view returns(uint256)
func (_Registry *RegistryCaller) GetTotalWorkflowsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getTotalWorkflowsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalWorkflowsCount is a free data retrieval call binding the contract method 0xe1d1dffe.
//
// Solidity: function getTotalWorkflowsCount() view returns(uint256)
func (_Registry *RegistrySession) GetTotalWorkflowsCount() (*big.Int, error) {
	return _Registry.Contract.GetTotalWorkflowsCount(&_Registry.CallOpts)
}

// GetTotalWorkflowsCount is a free data retrieval call binding the contract method 0xe1d1dffe.
//
// Solidity: function getTotalWorkflowsCount() view returns(uint256)
func (_Registry *RegistryCallerSession) GetTotalWorkflowsCount() (*big.Int, error) {
	return _Registry.Contract.GetTotalWorkflowsCount(&_Registry.CallOpts)
}

// GetWorkflowDepositAssetKeys is a free data retrieval call binding the contract method 0x0bdeec8f.
//
// Solidity: function getWorkflowDepositAssetKeys(uint256 _workflowId) view returns(string[])
func (_Registry *RegistryCaller) GetWorkflowDepositAssetKeys(opts *bind.CallOpts, _workflowId *big.Int) ([]string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getWorkflowDepositAssetKeys", _workflowId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetWorkflowDepositAssetKeys is a free data retrieval call binding the contract method 0x0bdeec8f.
//
// Solidity: function getWorkflowDepositAssetKeys(uint256 _workflowId) view returns(string[])
func (_Registry *RegistrySession) GetWorkflowDepositAssetKeys(_workflowId *big.Int) ([]string, error) {
	return _Registry.Contract.GetWorkflowDepositAssetKeys(&_Registry.CallOpts, _workflowId)
}

// GetWorkflowDepositAssetKeys is a free data retrieval call binding the contract method 0x0bdeec8f.
//
// Solidity: function getWorkflowDepositAssetKeys(uint256 _workflowId) view returns(string[])
func (_Registry *RegistryCallerSession) GetWorkflowDepositAssetKeys(_workflowId *big.Int) ([]string, error) {
	return _Registry.Contract.GetWorkflowDepositAssetKeys(&_Registry.CallOpts, _workflowId)
}

// GetWorkflowDepositAssetsInfo is a free data retrieval call binding the contract method 0xc17fc3eb.
//
// Solidity: function getWorkflowDepositAssetsInfo(uint256 _workflowId, string[] _depositAssetKeys) view returns((string,uint256)[] _depositAssetsArr)
func (_Registry *RegistryCaller) GetWorkflowDepositAssetsInfo(opts *bind.CallOpts, _workflowId *big.Int, _depositAssetKeys []string) ([]IRegistryDepositAssetInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getWorkflowDepositAssetsInfo", _workflowId, _depositAssetKeys)

	if err != nil {
		return *new([]IRegistryDepositAssetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRegistryDepositAssetInfo)).(*[]IRegistryDepositAssetInfo)

	return out0, err

}

// GetWorkflowDepositAssetsInfo is a free data retrieval call binding the contract method 0xc17fc3eb.
//
// Solidity: function getWorkflowDepositAssetsInfo(uint256 _workflowId, string[] _depositAssetKeys) view returns((string,uint256)[] _depositAssetsArr)
func (_Registry *RegistrySession) GetWorkflowDepositAssetsInfo(_workflowId *big.Int, _depositAssetKeys []string) ([]IRegistryDepositAssetInfo, error) {
	return _Registry.Contract.GetWorkflowDepositAssetsInfo(&_Registry.CallOpts, _workflowId, _depositAssetKeys)
}

// GetWorkflowDepositAssetsInfo is a free data retrieval call binding the contract method 0xc17fc3eb.
//
// Solidity: function getWorkflowDepositAssetsInfo(uint256 _workflowId, string[] _depositAssetKeys) view returns((string,uint256)[] _depositAssetsArr)
func (_Registry *RegistryCallerSession) GetWorkflowDepositAssetsInfo(_workflowId *big.Int, _depositAssetKeys []string) ([]IRegistryDepositAssetInfo, error) {
	return _Registry.Contract.GetWorkflowDepositAssetsInfo(&_Registry.CallOpts, _workflowId, _depositAssetKeys)
}

// GetWorkflowInfo is a free data retrieval call binding the contract method 0xdf7459f0.
//
// Solidity: function getWorkflowInfo(uint256 _workflowId) view returns(((uint256,address),string[],(string,uint256)[]))
func (_Registry *RegistryCaller) GetWorkflowInfo(opts *bind.CallOpts, _workflowId *big.Int) (IRegistryWorkflowInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getWorkflowInfo", _workflowId)

	if err != nil {
		return *new(IRegistryWorkflowInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IRegistryWorkflowInfo)).(*IRegistryWorkflowInfo)

	return out0, err

}

// GetWorkflowInfo is a free data retrieval call binding the contract method 0xdf7459f0.
//
// Solidity: function getWorkflowInfo(uint256 _workflowId) view returns(((uint256,address),string[],(string,uint256)[]))
func (_Registry *RegistrySession) GetWorkflowInfo(_workflowId *big.Int) (IRegistryWorkflowInfo, error) {
	return _Registry.Contract.GetWorkflowInfo(&_Registry.CallOpts, _workflowId)
}

// GetWorkflowInfo is a free data retrieval call binding the contract method 0xdf7459f0.
//
// Solidity: function getWorkflowInfo(uint256 _workflowId) view returns(((uint256,address),string[],(string,uint256)[]))
func (_Registry *RegistryCallerSession) GetWorkflowInfo(_workflowId *big.Int) (IRegistryWorkflowInfo, error) {
	return _Registry.Contract.GetWorkflowInfo(&_Registry.CallOpts, _workflowId)
}

// GetWorkflowOwner is a free data retrieval call binding the contract method 0xd69cd275.
//
// Solidity: function getWorkflowOwner(uint256 _id) view returns(address)
func (_Registry *RegistryCaller) GetWorkflowOwner(opts *bind.CallOpts, _id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getWorkflowOwner", _id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWorkflowOwner is a free data retrieval call binding the contract method 0xd69cd275.
//
// Solidity: function getWorkflowOwner(uint256 _id) view returns(address)
func (_Registry *RegistrySession) GetWorkflowOwner(_id *big.Int) (common.Address, error) {
	return _Registry.Contract.GetWorkflowOwner(&_Registry.CallOpts, _id)
}

// GetWorkflowOwner is a free data retrieval call binding the contract method 0xd69cd275.
//
// Solidity: function getWorkflowOwner(uint256 _id) view returns(address)
func (_Registry *RegistryCallerSession) GetWorkflowOwner(_id *big.Int) (common.Address, error) {
	return _Registry.Contract.GetWorkflowOwner(&_Registry.CallOpts, _id)
}

// GetWorkflowsInfo is a free data retrieval call binding the contract method 0x95f5dae4.
//
// Solidity: function getWorkflowsInfo(uint256 _offset, uint256 _limit) view returns(((uint256,address),string[],(string,uint256)[])[] _workflowsInfoArr)
func (_Registry *RegistryCaller) GetWorkflowsInfo(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]IRegistryWorkflowInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getWorkflowsInfo", _offset, _limit)

	if err != nil {
		return *new([]IRegistryWorkflowInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IRegistryWorkflowInfo)).(*[]IRegistryWorkflowInfo)

	return out0, err

}

// GetWorkflowsInfo is a free data retrieval call binding the contract method 0x95f5dae4.
//
// Solidity: function getWorkflowsInfo(uint256 _offset, uint256 _limit) view returns(((uint256,address),string[],(string,uint256)[])[] _workflowsInfoArr)
func (_Registry *RegistrySession) GetWorkflowsInfo(_offset *big.Int, _limit *big.Int) ([]IRegistryWorkflowInfo, error) {
	return _Registry.Contract.GetWorkflowsInfo(&_Registry.CallOpts, _offset, _limit)
}

// GetWorkflowsInfo is a free data retrieval call binding the contract method 0x95f5dae4.
//
// Solidity: function getWorkflowsInfo(uint256 _offset, uint256 _limit) view returns(((uint256,address),string[],(string,uint256)[])[] _workflowsInfoArr)
func (_Registry *RegistryCallerSession) GetWorkflowsInfo(_offset *big.Int, _limit *big.Int) ([]IRegistryWorkflowInfo, error) {
	return _Registry.Contract.GetWorkflowsInfo(&_Registry.CallOpts, _offset, _limit)
}

// IsWorkflowRegistered is a free data retrieval call binding the contract method 0xb8f9c0f6.
//
// Solidity: function isWorkflowRegistered(uint256 _id) view returns(bool)
func (_Registry *RegistryCaller) IsWorkflowRegistered(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isWorkflowRegistered", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWorkflowRegistered is a free data retrieval call binding the contract method 0xb8f9c0f6.
//
// Solidity: function isWorkflowRegistered(uint256 _id) view returns(bool)
func (_Registry *RegistrySession) IsWorkflowRegistered(_id *big.Int) (bool, error) {
	return _Registry.Contract.IsWorkflowRegistered(&_Registry.CallOpts, _id)
}

// IsWorkflowRegistered is a free data retrieval call binding the contract method 0xb8f9c0f6.
//
// Solidity: function isWorkflowRegistered(uint256 _id) view returns(bool)
func (_Registry *RegistryCallerSession) IsWorkflowRegistered(_id *big.Int) (bool, error) {
	return _Registry.Contract.IsWorkflowRegistered(&_Registry.CallOpts, _id)
}

// MaxWorkflowsPerAccount is a free data retrieval call binding the contract method 0xabad5c1a.
//
// Solidity: function maxWorkflowsPerAccount() view returns(uint16)
func (_Registry *RegistryCaller) MaxWorkflowsPerAccount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "maxWorkflowsPerAccount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxWorkflowsPerAccount is a free data retrieval call binding the contract method 0xabad5c1a.
//
// Solidity: function maxWorkflowsPerAccount() view returns(uint16)
func (_Registry *RegistrySession) MaxWorkflowsPerAccount() (uint16, error) {
	return _Registry.Contract.MaxWorkflowsPerAccount(&_Registry.CallOpts)
}

// MaxWorkflowsPerAccount is a free data retrieval call binding the contract method 0xabad5c1a.
//
// Solidity: function maxWorkflowsPerAccount() view returns(uint16)
func (_Registry *RegistryCallerSession) MaxWorkflowsPerAccount() (uint16, error) {
	return _Registry.Contract.MaxWorkflowsPerAccount(&_Registry.CallOpts)
}

// WorkflowsPerAddress is a free data retrieval call binding the contract method 0x6b699545.
//
// Solidity: function workflowsPerAddress(address ) view returns(uint256)
func (_Registry *RegistryCaller) WorkflowsPerAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "workflowsPerAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkflowsPerAddress is a free data retrieval call binding the contract method 0x6b699545.
//
// Solidity: function workflowsPerAddress(address ) view returns(uint256)
func (_Registry *RegistrySession) WorkflowsPerAddress(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.WorkflowsPerAddress(&_Registry.CallOpts, arg0)
}

// WorkflowsPerAddress is a free data retrieval call binding the contract method 0x6b699545.
//
// Solidity: function workflowsPerAddress(address ) view returns(uint256)
func (_Registry *RegistryCallerSession) WorkflowsPerAddress(arg0 common.Address) (*big.Int, error) {
	return _Registry.Contract.WorkflowsPerAddress(&_Registry.CallOpts, arg0)
}

// DeployAndSetGateway is a paid mutator transaction binding the contract method 0xc5fd64d8.
//
// Solidity: function deployAndSetGateway() returns(address)
func (_Registry *RegistryTransactor) DeployAndSetGateway(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "deployAndSetGateway")
}

// DeployAndSetGateway is a paid mutator transaction binding the contract method 0xc5fd64d8.
//
// Solidity: function deployAndSetGateway() returns(address)
func (_Registry *RegistrySession) DeployAndSetGateway() (*types.Transaction, error) {
	return _Registry.Contract.DeployAndSetGateway(&_Registry.TransactOpts)
}

// DeployAndSetGateway is a paid mutator transaction binding the contract method 0xc5fd64d8.
//
// Solidity: function deployAndSetGateway() returns(address)
func (_Registry *RegistryTransactorSession) DeployAndSetGateway() (*types.Transaction, error) {
	return _Registry.Contract.DeployAndSetGateway(&_Registry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _maxWorkflowsPerAccount) returns()
func (_Registry *RegistryTransactor) Initialize(opts *bind.TransactOpts, _maxWorkflowsPerAccount uint16) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "initialize", _maxWorkflowsPerAccount)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _maxWorkflowsPerAccount) returns()
func (_Registry *RegistrySession) Initialize(_maxWorkflowsPerAccount uint16) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, _maxWorkflowsPerAccount)
}

// Initialize is a paid mutator transaction binding the contract method 0x13750946.
//
// Solidity: function initialize(uint16 _maxWorkflowsPerAccount) returns()
func (_Registry *RegistryTransactorSession) Initialize(_maxWorkflowsPerAccount uint16) (*types.Transaction, error) {
	return _Registry.Contract.Initialize(&_Registry.TransactOpts, _maxWorkflowsPerAccount)
}

// Perform is a paid mutator transaction binding the contract method 0xd6bb757b.
//
// Solidity: function perform(uint256 _workflowId, uint256 _gasAmount, bytes _data, address _target) returns()
func (_Registry *RegistryTransactor) Perform(opts *bind.TransactOpts, _workflowId *big.Int, _gasAmount *big.Int, _data []byte, _target common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "perform", _workflowId, _gasAmount, _data, _target)
}

// Perform is a paid mutator transaction binding the contract method 0xd6bb757b.
//
// Solidity: function perform(uint256 _workflowId, uint256 _gasAmount, bytes _data, address _target) returns()
func (_Registry *RegistrySession) Perform(_workflowId *big.Int, _gasAmount *big.Int, _data []byte, _target common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Perform(&_Registry.TransactOpts, _workflowId, _gasAmount, _data, _target)
}

// Perform is a paid mutator transaction binding the contract method 0xd6bb757b.
//
// Solidity: function perform(uint256 _workflowId, uint256 _gasAmount, bytes _data, address _target) returns()
func (_Registry *RegistryTransactorSession) Perform(_workflowId *big.Int, _gasAmount *big.Int, _data []byte, _target common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Perform(&_Registry.TransactOpts, _workflowId, _gasAmount, _data, _target)
}

// RegisterWorkflows is a paid mutator transaction binding the contract method 0x8cdfb69a.
//
// Solidity: function registerWorkflows((uint256,address,bool,bool)[] _registerWorkflowInfoArr) returns()
func (_Registry *RegistryTransactor) RegisterWorkflows(opts *bind.TransactOpts, _registerWorkflowInfoArr []IRegistryRegisterWorkflowInfo) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "registerWorkflows", _registerWorkflowInfoArr)
}

// RegisterWorkflows is a paid mutator transaction binding the contract method 0x8cdfb69a.
//
// Solidity: function registerWorkflows((uint256,address,bool,bool)[] _registerWorkflowInfoArr) returns()
func (_Registry *RegistrySession) RegisterWorkflows(_registerWorkflowInfoArr []IRegistryRegisterWorkflowInfo) (*types.Transaction, error) {
	return _Registry.Contract.RegisterWorkflows(&_Registry.TransactOpts, _registerWorkflowInfoArr)
}

// RegisterWorkflows is a paid mutator transaction binding the contract method 0x8cdfb69a.
//
// Solidity: function registerWorkflows((uint256,address,bool,bool)[] _registerWorkflowInfoArr) returns()
func (_Registry *RegistryTransactorSession) RegisterWorkflows(_registerWorkflowInfoArr []IRegistryRegisterWorkflowInfo) (*types.Transaction, error) {
	return _Registry.Contract.RegisterWorkflows(&_Registry.TransactOpts, _registerWorkflowInfoArr)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_Registry *RegistryTransactor) SetDependencies(opts *bind.TransactOpts, _contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setDependencies", _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_Registry *RegistrySession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetDependencies(&_Registry.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_Registry *RegistryTransactorSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Registry.Contract.SetDependencies(&_Registry.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address _gateway) returns()
func (_Registry *RegistryTransactor) SetGateway(opts *bind.TransactOpts, _gateway common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setGateway", _gateway)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address _gateway) returns()
func (_Registry *RegistrySession) SetGateway(_gateway common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGateway(&_Registry.TransactOpts, _gateway)
}

// SetGateway is a paid mutator transaction binding the contract method 0x90646b4a.
//
// Solidity: function setGateway(address _gateway) returns()
func (_Registry *RegistryTransactorSession) SetGateway(_gateway common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetGateway(&_Registry.TransactOpts, _gateway)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Registry *RegistryTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Registry *RegistrySession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetInjector(&_Registry.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Registry *RegistryTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetInjector(&_Registry.TransactOpts, injector_)
}

// SetMaxWorkflowsPerAccount is a paid mutator transaction binding the contract method 0x44ab987d.
//
// Solidity: function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) returns()
func (_Registry *RegistryTransactor) SetMaxWorkflowsPerAccount(opts *bind.TransactOpts, _newMaxWorkflowsPerAccount uint16) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setMaxWorkflowsPerAccount", _newMaxWorkflowsPerAccount)
}

// SetMaxWorkflowsPerAccount is a paid mutator transaction binding the contract method 0x44ab987d.
//
// Solidity: function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) returns()
func (_Registry *RegistrySession) SetMaxWorkflowsPerAccount(_newMaxWorkflowsPerAccount uint16) (*types.Transaction, error) {
	return _Registry.Contract.SetMaxWorkflowsPerAccount(&_Registry.TransactOpts, _newMaxWorkflowsPerAccount)
}

// SetMaxWorkflowsPerAccount is a paid mutator transaction binding the contract method 0x44ab987d.
//
// Solidity: function setMaxWorkflowsPerAccount(uint16 _newMaxWorkflowsPerAccount) returns()
func (_Registry *RegistryTransactorSession) SetMaxWorkflowsPerAccount(_newMaxWorkflowsPerAccount uint16) (*types.Transaction, error) {
	return _Registry.Contract.SetMaxWorkflowsPerAccount(&_Registry.TransactOpts, _newMaxWorkflowsPerAccount)
}

// UpdateWorkflowTotalSpent is a paid mutator transaction binding the contract method 0xae09340f.
//
// Solidity: function updateWorkflowTotalSpent(string _depositAssetKey, uint256 _workflowId, uint256 _workflowExecutionAmount) returns()
func (_Registry *RegistryTransactor) UpdateWorkflowTotalSpent(opts *bind.TransactOpts, _depositAssetKey string, _workflowId *big.Int, _workflowExecutionAmount *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateWorkflowTotalSpent", _depositAssetKey, _workflowId, _workflowExecutionAmount)
}

// UpdateWorkflowTotalSpent is a paid mutator transaction binding the contract method 0xae09340f.
//
// Solidity: function updateWorkflowTotalSpent(string _depositAssetKey, uint256 _workflowId, uint256 _workflowExecutionAmount) returns()
func (_Registry *RegistrySession) UpdateWorkflowTotalSpent(_depositAssetKey string, _workflowId *big.Int, _workflowExecutionAmount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UpdateWorkflowTotalSpent(&_Registry.TransactOpts, _depositAssetKey, _workflowId, _workflowExecutionAmount)
}

// UpdateWorkflowTotalSpent is a paid mutator transaction binding the contract method 0xae09340f.
//
// Solidity: function updateWorkflowTotalSpent(string _depositAssetKey, uint256 _workflowId, uint256 _workflowExecutionAmount) returns()
func (_Registry *RegistryTransactorSession) UpdateWorkflowTotalSpent(_depositAssetKey string, _workflowId *big.Int, _workflowExecutionAmount *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.UpdateWorkflowTotalSpent(&_Registry.TransactOpts, _depositAssetKey, _workflowId, _workflowExecutionAmount)
}

// RegistryGatewaySetIterator is returned from FilterGatewaySet and is used to iterate over the raw logs and unpacked data for GatewaySet events raised by the Registry contract.
type RegistryGatewaySetIterator struct {
	Event *RegistryGatewaySet // Event containing the contract specifics and raw log

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
func (it *RegistryGatewaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryGatewaySet)
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
		it.Event = new(RegistryGatewaySet)
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
func (it *RegistryGatewaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryGatewaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryGatewaySet represents a GatewaySet event raised by the Registry contract.
type RegistryGatewaySet struct {
	Owner   common.Address
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewaySet is a free log retrieval operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address owner, address gateway)
func (_Registry *RegistryFilterer) FilterGatewaySet(opts *bind.FilterOpts) (*RegistryGatewaySetIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "GatewaySet")
	if err != nil {
		return nil, err
	}
	return &RegistryGatewaySetIterator{contract: _Registry.contract, event: "GatewaySet", logs: logs, sub: sub}, nil
}

// WatchGatewaySet is a free log subscription operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address owner, address gateway)
func (_Registry *RegistryFilterer) WatchGatewaySet(opts *bind.WatchOpts, sink chan<- *RegistryGatewaySet) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "GatewaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryGatewaySet)
				if err := _Registry.contract.UnpackLog(event, "GatewaySet", log); err != nil {
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

// ParseGatewaySet is a log parse operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address owner, address gateway)
func (_Registry *RegistryFilterer) ParseGatewaySet(log types.Log) (*RegistryGatewaySet, error) {
	event := new(RegistryGatewaySet)
	if err := _Registry.contract.UnpackLog(event, "GatewaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registry contract.
type RegistryInitializedIterator struct {
	Event *RegistryInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryInitialized)
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
		it.Event = new(RegistryInitialized)
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
func (it *RegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryInitialized represents a Initialized event raised by the Registry contract.
type RegistryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryInitializedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryInitializedIterator{contract: _Registry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registry *RegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryInitialized)
				if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registry *RegistryFilterer) ParseInitialized(log types.Log) (*RegistryInitialized, error) {
	event := new(RegistryInitialized)
	if err := _Registry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryPerformanceIterator is returned from FilterPerformance and is used to iterate over the raw logs and unpacked data for Performance events raised by the Registry contract.
type RegistryPerformanceIterator struct {
	Event *RegistryPerformance // Event containing the contract specifics and raw log

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
func (it *RegistryPerformanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryPerformance)
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
		it.Event = new(RegistryPerformance)
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
func (it *RegistryPerformanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryPerformanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryPerformance represents a Performance event raised by the Registry contract.
type RegistryPerformance struct {
	WorkflowId *big.Int
	Success    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPerformance is a free log retrieval operation binding the contract event 0x8e67dfc40b26eb841f1be843b96e6d300ff34a0f3f032496584530cd6544c3fb.
//
// Solidity: event Performance(uint256 workflowId, bool success)
func (_Registry *RegistryFilterer) FilterPerformance(opts *bind.FilterOpts) (*RegistryPerformanceIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Performance")
	if err != nil {
		return nil, err
	}
	return &RegistryPerformanceIterator{contract: _Registry.contract, event: "Performance", logs: logs, sub: sub}, nil
}

// WatchPerformance is a free log subscription operation binding the contract event 0x8e67dfc40b26eb841f1be843b96e6d300ff34a0f3f032496584530cd6544c3fb.
//
// Solidity: event Performance(uint256 workflowId, bool success)
func (_Registry *RegistryFilterer) WatchPerformance(opts *bind.WatchOpts, sink chan<- *RegistryPerformance) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Performance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryPerformance)
				if err := _Registry.contract.UnpackLog(event, "Performance", log); err != nil {
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

// ParsePerformance is a log parse operation binding the contract event 0x8e67dfc40b26eb841f1be843b96e6d300ff34a0f3f032496584530cd6544c3fb.
//
// Solidity: event Performance(uint256 workflowId, bool success)
func (_Registry *RegistryFilterer) ParsePerformance(log types.Log) (*RegistryPerformance, error) {
	event := new(RegistryPerformance)
	if err := _Registry.contract.UnpackLog(event, "Performance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryWorkflowRegisteredIterator is returned from FilterWorkflowRegistered and is used to iterate over the raw logs and unpacked data for WorkflowRegistered events raised by the Registry contract.
type RegistryWorkflowRegisteredIterator struct {
	Event *RegistryWorkflowRegistered // Event containing the contract specifics and raw log

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
func (it *RegistryWorkflowRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryWorkflowRegistered)
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
		it.Event = new(RegistryWorkflowRegistered)
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
func (it *RegistryWorkflowRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryWorkflowRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryWorkflowRegistered represents a WorkflowRegistered event raised by the Registry contract.
type RegistryWorkflowRegistered struct {
	Owner common.Address
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWorkflowRegistered is a free log retrieval operation binding the contract event 0x756019c5646bbc7e222ed197b047a6cd3c7fe80b10551475f1925d4c9ee4b8ee.
//
// Solidity: event WorkflowRegistered(address owner, uint256 id)
func (_Registry *RegistryFilterer) FilterWorkflowRegistered(opts *bind.FilterOpts) (*RegistryWorkflowRegisteredIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "WorkflowRegistered")
	if err != nil {
		return nil, err
	}
	return &RegistryWorkflowRegisteredIterator{contract: _Registry.contract, event: "WorkflowRegistered", logs: logs, sub: sub}, nil
}

// WatchWorkflowRegistered is a free log subscription operation binding the contract event 0x756019c5646bbc7e222ed197b047a6cd3c7fe80b10551475f1925d4c9ee4b8ee.
//
// Solidity: event WorkflowRegistered(address owner, uint256 id)
func (_Registry *RegistryFilterer) WatchWorkflowRegistered(opts *bind.WatchOpts, sink chan<- *RegistryWorkflowRegistered) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "WorkflowRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryWorkflowRegistered)
				if err := _Registry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
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

// ParseWorkflowRegistered is a log parse operation binding the contract event 0x756019c5646bbc7e222ed197b047a6cd3c7fe80b10551475f1925d4c9ee4b8ee.
//
// Solidity: event WorkflowRegistered(address owner, uint256 id)
func (_Registry *RegistryFilterer) ParseWorkflowRegistered(log types.Log) (*RegistryWorkflowRegistered, error) {
	event := new(RegistryWorkflowRegistered)
	if err := _Registry.contract.UnpackLog(event, "WorkflowRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
