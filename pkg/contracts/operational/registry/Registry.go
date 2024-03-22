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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"workflowExecutionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"Performance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"WorkflowRegistered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deployAndSetGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"}],\"name\":\"getBaseWorkflowInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.BaseWorkflowInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getGatewaysInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gatewayOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.GatewayInfo[]\",\"name\":\"_gatewaysInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalGatewaysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWorkflowsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowDepositAssetKeys\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"_depositAssetKeys\",\"type\":\"string[]\"}],\"name\":\"getWorkflowDepositAssetsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"depositAssetTotalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.DepositAssetInfo[]\",\"name\":\"_depositAssetsArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"}],\"name\":\"getWorkflowInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.BaseWorkflowInfo\",\"name\":\"baseInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"depositAssetKeys\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"depositAssetTotalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.DepositAssetInfo[]\",\"name\":\"depositAssetsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structIRegistry.WorkflowInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getWorkflowOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getWorkflowsInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"internalType\":\"structIRegistry.BaseWorkflowInfo\",\"name\":\"baseInfo\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"depositAssetKeys\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"depositAssetTotalSpent\",\"type\":\"uint256\"}],\"internalType\":\"structIRegistry.DepositAssetInfo[]\",\"name\":\"depositAssetsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structIRegistry.WorkflowInfo[]\",\"name\":\"_workflowsInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_maxWorkflowsPerAccount\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"isWorkflowExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxWorkflowsPerAccount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"}],\"name\":\"perform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"workflowOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"requireGateway\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"deployGateway\",\"type\":\"bool\"}],\"internalType\":\"structIRegistry.RegisterWorkflowInfo[]\",\"name\":\"_registerWorkflowInfoArr\",\"type\":\"tuple[]\"}],\"name\":\"registerWorkflows\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newMaxWorkflowsPerAccount\",\"type\":\"uint16\"}],\"name\":\"setMaxWorkflowsPerAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_depositAssetKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_workflowId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_workflowExecutionAmount\",\"type\":\"uint256\"}],\"name\":\"updateWorkflowTotalSpent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"workflowsPerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506123fb806100206000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806395f5dae4116100c3578063c17fc3eb1161007c578063c17fc3eb1461035e578063c48223191461037e578063c5fd64d814610386578063d69cd2751461038e578063df7459f0146103ba578063e1d1dffe146103da57600080fd5b806395f5dae4146102795780639b2bfaa314610299578063a3190a36146102d7578063abad5c1a146102f7578063ae09340f1461031f578063bda009fe1461033257600080fd5b8063691304511161011557806369130451146101ec5780636b699545146101ff5780637a03f9b51461022d5780638cb941cc146102405780638cdfb69a1461025357806390646b4a1461026657600080fd5b80630bdeec8f14610152578063137509461461017b5780633e3b5b191461019057806344ab987d146101b95780634b322560146101cc575b600080fd5b610165610160366004611a30565b6103e2565b6040516101729190611a99565b60405180910390f35b61018e610189366004611afb565b610405565b005b6000805160206123a6833981519152545b6040516001600160a01b039091168152602001610172565b61018e6101c7366004611afb565b61052f565b6101df6101da366004611a30565b610559565b6040516101729190611b1f565b61018e6101fa366004611bf3565b6105a0565b61021f61020d366004611c57565b60066020526000908152604090205481565b604051908152602001610172565b61018e61023b366004611c74565b6106f4565b61018e61024e366004611c57565b610a23565b61018e610261366004611d17565b610a44565b61018e610274366004611c57565b610f30565b61028c610287366004611d8c565b610f3a565b6040516101729190611ebc565b6102c76102a7366004611a30565b6000908152600760205260409020600101546001600160a01b0316151590565b6040519015158152602001610172565b6102ea6102e5366004611d8c565b61104c565b6040516101729190611f11565b60025461030c90600160a01b900461ffff1681565b60405161ffff9091168152602001610172565b61018e61032d366004611f80565b611153565b6101a1610340366004611c57565b6001600160a01b039081166000908152600860205260409020541690565b61037161036c366004611fce565b611232565b604051610172919061209d565b61021f611353565b6101a1611364565b6101a161039c366004611a30565b6000908152600760205260409020600101546001600160a01b031690565b6103cd6103c8366004611a30565b61136f565b60405161017291906120b0565b60035461021f565b60008181526007602052604090206060906103ff906002016113e3565b92915050565b600054610100900460ff16158080156104255750600054600160ff909116105b8061043f5750303b15801561043f575060005460ff166001145b6104a75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156104ca576000805461ff0019166101001790555b6002805461ffff60a01b1916600160a01b61ffff851602179055801561052b576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498906020015b60405180910390a15b5050565b6105376114c0565b6002805461ffff909216600160a01b0261ffff60a01b19909216919091179055565b604080518082019091526000808252602082015250600090815260076020908152604091829020825180840190935280548352600101546001600160a01b03169082015290565b6105a8611584565b600082905080600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663fabf588e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610612573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061063691906120c3565b600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b03166348197f776040518163ffffffff1660e01b8152600401602060405180830381865afa15801561069a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106be91906120c3565b600280546001600160a01b0319166001600160a01b039290921691909117905550336000805160206123a6833981519152555050565b6106fc6114c0565b8561070681611608565b600254604051634a686e4960e01b81526004810188905288916001600160a01b031690634a686e4990602401602060405180830381865afa15801561074f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077391906120e0565b1480156107fc57506002546040516307105e9760e41b8152600481018890526001916001600160a01b031690637105e97090602401602060405180830381865afa1580156107c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e9919061210f565b60028111156107fa576107fa6120f9565b145b6108585760405162461bcd60e51b815260206004820152602760248201527f52656769737472793a20696e76616c696420776f726b666c6f772065786563756044820152661d1a5bdb88125160ca1b606482015260840161049e565b6001600160a01b03821630036108bc5760405162461bcd60e51b8152602060048201526024808201527f52656769737472793a206f7065726174696f6e206973206e6f74207065726d696044820152631d1d195960e21b606482015260840161049e565b6000878152600760209081526040808320600101546001600160a01b03908116845260089092529091205416806109355760405162461bcd60e51b815260206004820152601e60248201527f52656769737472793a207a65726f206761746577617920616464726573730000604482015260640161049e565b60006109d387836040518060400160405280601e81526020017f706572666f726d2875696e743235362c616464726573732c62797465732900008152508c888b8b6040516024016109899493929190612130565b60408051601f1981840301815290829052916109a491612177565b6040519081900390206020820180516001600160e01b03166001600160e01b0319909216919091179052611679565b604080518b8152602081018b90528215158183015290519192507fc723c444dde505205b3ec0c789ed1adeade412952dc2caecb0ac55b9668e0105919081900360600190a1505050505050505050565b610a2b611584565b610a41816000805160206123a683398151915255565b50565b60008060029054906101000a90046001600160a01b03166001600160a01b031663b834f6fb6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abc91906121a1565b905060008060029054906101000a90046001600160a01b03166001600160a01b0316637ac3c02f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3691906120c3565b9050818015610b515750600254600160a01b900461ffff1615155b15610be05760025433600090815260066020526040902054600160a01b90910461ffff1690610b819085906121d4565b1115610be05760405162461bcd60e51b815260206004820152602860248201527f52656769737472793a2072656163686564206d617820776f726b666c6f777320604482015267636170616369747960c01b606482015260840161049e565b60005b83811015610f295736858583818110610bfe57610bfe6121e7565b905060800201905083610c1c576001600160a01b0383163314610c38565b33610c2d6040830160208401611c57565b6001600160a01b0316145b610c845760405162461bcd60e51b815260206004820181905260248201527f52656769737472793a206e6f7420612073656e646572206f72207369676e6572604482015260640161049e565b80356000908152600760205260409020600101546001600160a01b031615610cfa5760405162461bcd60e51b8152602060048201526024808201527f52656769737472793a20776f726b666c6f7720696420616c72656164792065786044820152636973747360e01b606482015260840161049e565b610d0a60608201604083016121fd565b15610de1576000600881610d246040850160208601611c57565b6001600160a01b039081168252602082019290925260400160002054169050610d5360808301606084016121fd565b8015610d6657506001600160a01b038116155b15610d8957610d83610d7e6040840160208501611c57565b6116c5565b50610ddf565b6001600160a01b038116610ddf5760405162461bcd60e51b815260206004820152601b60248201527f52656769737472793a2067617465776179206e6f7420666f756e640000000000604482015260640161049e565b505b604051806040016040528082600001358152602001826020016020810190610e099190611c57565b6001600160a01b03908116909152823560009081526007602090815260408083208551815594820151600190950180546001600160a01b03191695909416949094179092556006929091610e61918501908501611c57565b6001600160a01b0316815260208101919091526040016000908120805491610e888361221a565b90915550506003805460018101825560009190915281357fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b909101557f756019c5646bbc7e222ed197b047a6cd3c7fe80b10551475f1925d4c9ee4b8ee610ef56040830160208401611c57565b604080516001600160a01b039092168252833560208301520160405180910390a15080610f218161221a565b915050610be3565b5050505050565b610a413382611741565b60606000610f4e60038054905085856117d2565b9050610f5a8482612233565b67ffffffffffffffff811115610f7257610f72611b54565b604051908082528060200260200182016040528015610fd857816020015b610fc56040805160a0810190915260006060820181815260808301919091528190815260200160608152602001606081525090565b815260200190600190039081610f905790505b509150835b818110156110445761100b60038281548110610ffb57610ffb6121e7565b906000526020600020015461136f565b836110168784612233565b81518110611026576110266121e7565b6020026020010181905250808061103c9061221a565b915050610fdd565b505092915050565b6060600061106461105d60046117fd565b85856117d2565b90506110708482612233565b67ffffffffffffffff81111561108857611088611b54565b6040519080825280602002602001820160405280156110cd57816020015b60408051808201909152600080825260208201528152602001906001900390816110a65790505b509150835b818110156110445760006110e7600483611807565b6040805180820182526001600160a01b03808416808352600090815260086020908152939020541691810191909152909150846111248885612233565b81518110611134576111346121e7565b602002602001018190525050808061114b9061221a565b9150506110d2565b8161115d81611608565b6002546001600160a01b031633146111c95760405162461bcd60e51b815260206004820152602960248201527f52656769737472793a2073656e646572206973206e6f7420612062696c6c696e604482015268339036b0b730b3b2b960b91b606482015260840161049e565b60008381526007602052604090206111e49060020185611813565b5081600760008581526020019081526020016000206004018560405161120a9190612177565b9081526020016040518091039020600082825461122791906121d4565b909155505050505050565b6060815167ffffffffffffffff81111561124e5761124e611b54565b60405190808252806020026020018201604052801561129457816020015b60408051808201909152606081526000602082015281526020019060019003908161126c5790505b50905060005b825181101561134c5760405180604001604052808483815181106112c0576112c06121e7565b60200260200101518152602001600760008781526020019081526020016000206004018584815181106112f5576112f56121e7565b602002602001015160405161130a9190612177565b90815260200160405180910390205481525082828151811061132e5761132e6121e7565b602002602001018190525080806113449061221a565b91505061129a565b5092915050565b600061135f60046117fd565b905090565b600061135f336116c5565b6113a46040805160a0810190915260006060820181815260808301919091528190815260200160608152602001606081525090565b60006113af836103e2565b905060405180606001604052806113c585610559565b81526020018281526020016113da8584611232565b90529392505050565b606081600001805480602002602001604051908101604052809291908181526020016000905b828210156114b557838290600052602060002001805461142890612246565b80601f016020809104026020016040519081016040528092919081815260200182805461145490612246565b80156114a15780601f10611476576101008083540402835291602001916114a1565b820191906000526020600020905b81548152906001019060200180831161148457829003601f168201915b505050505081526020019060010190611409565b505050509050919050565b60005460408051637ac3c02f60e01b8152905133926201000090046001600160a01b031691637ac3c02f9160048083019260209291908290030181865afa15801561150f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153391906120c3565b6001600160a01b0316146115825760405162461bcd60e51b81526020600482015260166024820152752932b3b4b9ba393c9d102737ba10309039b4b3b732b960511b604482015260640161049e565b565b600061159c6000805160206123a68339815191525490565b90506001600160a01b03811615806115bc57506001600160a01b03811633145b610a415760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f72000000000000604482015260640161049e565b6000818152600760205260409020600101546001600160a01b0316610a415760405162461bcd60e51b815260206004820152602160248201527f52656769737472793a20776f726b666c6f7720646f6573206e6f7420657869736044820152601d60fa1b606482015260840161049e565b60005a61138881101561168b57600080fd5b6113888103905084604082048203116116a357600080fd5b50823b6116af57600080fd5b60008083516020850160008789f1949350505050565b60015460405163257e6bf960e21b81526001600160a01b03838116600483015260009283929116906395f9afe4906024016020604051808303816000875af1158015611715573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061173991906120c3565b90506103ff83825b6001600160a01b038116156117615761175b600483611875565b5061176e565b61176c60048361188a565b505b6001600160a01b0382811660008181526008602090815260409182902080546001600160a01b031916948616948517905581519283528201929092527f812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba3549101610522565b60006117de82846121d4565b9050838111156117eb5750825b808311156117f65750815b9392505050565b60006103ff825490565b60006117f6838361189f565b600061181f83836118c9565b61186d5782546001810184556000848152602090200161183f83826122cf565b5082546040516001850190611855908590612177565b908152604051908190036020019020555060016103ff565b5060006103ff565b60006117f6836001600160a01b0384166118f6565b60006117f6836001600160a01b03841661193d565b60008260000182815481106118b6576118b66121e7565b9060005260206000200154905092915050565b600082600101826040516118dd9190612177565b9081526040519081900360200190205415159392505050565b600081815260018301602052604081205461186d575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556103ff565b60008181526001830160205260408120548015611a26576000611961600183612233565b855490915060009061197590600190612233565b90508181146119da576000866000018281548110611995576119956121e7565b90600052602060002001549050808760000184815481106119b8576119b86121e7565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806119eb576119eb61238f565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506103ff565b60009150506103ff565b600060208284031215611a4257600080fd5b5035919050565b60005b83811015611a64578181015183820152602001611a4c565b50506000910152565b60008151808452611a85816020860160208601611a49565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611aee57603f19888603018452611adc858351611a6d565b94509285019290850190600101611ac0565b5092979650505050505050565b600060208284031215611b0d57600080fd5b813561ffff811681146117f657600080fd5b815181526020808301516001600160a01b031690820152604081016103ff565b6001600160a01b0381168114610a4157600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611b9357611b93611b54565b604052919050565b600067ffffffffffffffff831115611bb557611bb5611b54565b611bc8601f8401601f1916602001611b6a565b9050828152838383011115611bdc57600080fd5b828260208301376000602084830101529392505050565b60008060408385031215611c0657600080fd5b8235611c1181611b3f565b9150602083013567ffffffffffffffff811115611c2d57600080fd5b8301601f81018513611c3e57600080fd5b611c4d85823560208401611b9b565b9150509250929050565b600060208284031215611c6957600080fd5b81356117f681611b3f565b60008060008060008060a08789031215611c8d57600080fd5b863595506020870135945060408701359350606087013567ffffffffffffffff80821115611cba57600080fd5b818901915089601f830112611cce57600080fd5b813581811115611cdd57600080fd5b8a6020828501011115611cef57600080fd5b6020830195508094505050506080870135611d0981611b3f565b809150509295509295509295565b60008060208385031215611d2a57600080fd5b823567ffffffffffffffff80821115611d4257600080fd5b818501915085601f830112611d5657600080fd5b813581811115611d6557600080fd5b8660208260071b8501011115611d7a57600080fd5b60209290920196919550909350505050565b60008060408385031215611d9f57600080fd5b50508035926020909101359150565b600081518084526020808501808196508360051b8101915082860160005b85811015611e0c578284038952815160408151818752611dee82880182611a6d565b92880151968801969096525098850198935090840190600101611dcc565b5091979650505050505050565b600060808301611e3d848451805182526020908101516001600160a01b0316910152565b6020808401516080604087015282815180855260a08801915060a08160051b8901019450838301925060005b81811015611e9757609f19898703018352611e85868551611a6d565b95509284019291840191600101611e69565b505050505060408301518482036060860152611eb38282611dae565b95945050505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015611aee57603f19888603018452611eff858351611e19565b94509285019290850190600101611ee3565b602080825282518282018190526000919060409081850190868401855b82811015611e0c57815180516001600160a01b0390811686529087015116868501529284019290850190600101611f2e565b600082601f830112611f7157600080fd5b6117f683833560208501611b9b565b600080600060608486031215611f9557600080fd5b833567ffffffffffffffff811115611fac57600080fd5b611fb886828701611f60565b9660208601359650604090950135949350505050565b60008060408385031215611fe157600080fd5b8235915060208084013567ffffffffffffffff8082111561200157600080fd5b818601915086601f83011261201557600080fd5b81358181111561202757612027611b54565b8060051b612036858201611b6a565b918252838101850191858101908a84111561205057600080fd5b86860192505b8383101561208c5782358581111561206e5760008081fd5b61207c8c89838a0101611f60565b8352509186019190860190612056565b809750505050505050509250929050565b6020815260006117f66020830184611dae565b6020815260006117f66020830184611e19565b6000602082840312156120d557600080fd5b81516117f681611b3f565b6000602082840312156120f257600080fd5b5051919050565b634e487b7160e01b600052602160045260246000fd5b60006020828403121561212157600080fd5b8151600381106117f657600080fd5b8481526001600160a01b03841660208201526060604082018190528101829052818360808301376000818301608090810191909152601f909201601f191601019392505050565b60008251612189818460208701611a49565b9190910192915050565b8015158114610a4157600080fd5b6000602082840312156121b357600080fd5b81516117f681612193565b634e487b7160e01b600052601160045260246000fd5b808201808211156103ff576103ff6121be565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561220f57600080fd5b81356117f681612193565b60006001820161222c5761222c6121be565b5060010190565b818103818111156103ff576103ff6121be565b600181811c9082168061225a57607f821691505b60208210810361227a57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156122ca57600081815260208120601f850160051c810160208610156122a75750805b601f850160051c820191505b818110156122c6578281556001016122b3565b5050505b505050565b815167ffffffffffffffff8111156122e9576122e9611b54565b6122fd816122f78454612246565b84612280565b602080601f831160018114612332576000841561231a5750858301515b600019600386901b1c1916600185901b1785556122c6565b600085815260208120601f198616915b8281101561236157888601518255948401946001909101908401612342565b508582101561237f5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603160045260246000fdfe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a2646970667358221220c61d53ffa37e86e784238e868297de19e2a39e2463606d00ae3f12ddbcc4927764736f6c63430008120033",
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

// IsWorkflowExist is a free data retrieval call binding the contract method 0x9b2bfaa3.
//
// Solidity: function isWorkflowExist(uint256 _id) view returns(bool)
func (_Registry *RegistryCaller) IsWorkflowExist(opts *bind.CallOpts, _id *big.Int) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isWorkflowExist", _id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWorkflowExist is a free data retrieval call binding the contract method 0x9b2bfaa3.
//
// Solidity: function isWorkflowExist(uint256 _id) view returns(bool)
func (_Registry *RegistrySession) IsWorkflowExist(_id *big.Int) (bool, error) {
	return _Registry.Contract.IsWorkflowExist(&_Registry.CallOpts, _id)
}

// IsWorkflowExist is a free data retrieval call binding the contract method 0x9b2bfaa3.
//
// Solidity: function isWorkflowExist(uint256 _id) view returns(bool)
func (_Registry *RegistryCallerSession) IsWorkflowExist(_id *big.Int) (bool, error) {
	return _Registry.Contract.IsWorkflowExist(&_Registry.CallOpts, _id)
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

// Perform is a paid mutator transaction binding the contract method 0x7a03f9b5.
//
// Solidity: function perform(uint256 _workflowId, uint256 _workflowExecutionId, uint256 _gasAmount, bytes _data, address _target) returns()
func (_Registry *RegistryTransactor) Perform(opts *bind.TransactOpts, _workflowId *big.Int, _workflowExecutionId *big.Int, _gasAmount *big.Int, _data []byte, _target common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "perform", _workflowId, _workflowExecutionId, _gasAmount, _data, _target)
}

// Perform is a paid mutator transaction binding the contract method 0x7a03f9b5.
//
// Solidity: function perform(uint256 _workflowId, uint256 _workflowExecutionId, uint256 _gasAmount, bytes _data, address _target) returns()
func (_Registry *RegistrySession) Perform(_workflowId *big.Int, _workflowExecutionId *big.Int, _gasAmount *big.Int, _data []byte, _target common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Perform(&_Registry.TransactOpts, _workflowId, _workflowExecutionId, _gasAmount, _data, _target)
}

// Perform is a paid mutator transaction binding the contract method 0x7a03f9b5.
//
// Solidity: function perform(uint256 _workflowId, uint256 _workflowExecutionId, uint256 _gasAmount, bytes _data, address _target) returns()
func (_Registry *RegistryTransactorSession) Perform(_workflowId *big.Int, _workflowExecutionId *big.Int, _gasAmount *big.Int, _data []byte, _target common.Address) (*types.Transaction, error) {
	return _Registry.Contract.Perform(&_Registry.TransactOpts, _workflowId, _workflowExecutionId, _gasAmount, _data, _target)
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
	WorkflowId          *big.Int
	WorkflowExecutionId *big.Int
	Success             bool
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPerformance is a free log retrieval operation binding the contract event 0xc723c444dde505205b3ec0c789ed1adeade412952dc2caecb0ac55b9668e0105.
//
// Solidity: event Performance(uint256 workflowId, uint256 workflowExecutionId, bool success)
func (_Registry *RegistryFilterer) FilterPerformance(opts *bind.FilterOpts) (*RegistryPerformanceIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Performance")
	if err != nil {
		return nil, err
	}
	return &RegistryPerformanceIterator{contract: _Registry.contract, event: "Performance", logs: logs, sub: sub}, nil
}

// WatchPerformance is a free log subscription operation binding the contract event 0xc723c444dde505205b3ec0c789ed1adeade412952dc2caecb0ac55b9668e0105.
//
// Solidity: event Performance(uint256 workflowId, uint256 workflowExecutionId, bool success)
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

// ParsePerformance is a log parse operation binding the contract event 0xc723c444dde505205b3ec0c789ed1adeade412952dc2caecb0ac55b9668e0105.
//
// Solidity: event Performance(uint256 workflowId, uint256 workflowExecutionId, bool success)
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
