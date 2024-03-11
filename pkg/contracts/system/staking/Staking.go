// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// IStakingValidatorData is an auto generated low-level Go binding around an user-defined struct.
type IStakingValidatorData struct {
	Stake  *big.Int
	Status uint8
}

// IStakingValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakingValidatorInfo struct {
	ValidatorAddr common.Address
	ValidatorData IStakingValidatorData
}

// IStakingWithdrawalAnnouncement is an auto generated low-level Go binding around an user-defined struct.
type IStakingWithdrawalAnnouncement struct {
	Amount *big.Int
	Time   *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"}],\"name\":\"MinimalStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokensSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokensRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"TokensStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensAmount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"WithdrawalPeriodUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addRewardsToStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"announceWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumIStaking.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getValidatorsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumIStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIStaking.ValidatorData\",\"name\":\"validatorData\",\"type\":\"tuple\"}],\"internalType\":\"structIStaking.ValidatorInfo[]\",\"name\":\"_validatorsInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getWithdrawalAnnouncement\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"internalType\":\"structIStaking.WithdrawalAnnouncement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"hasWithdrawalAnnouncement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"}],\"name\":\"setMinimalStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sigExpirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"stakeWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611f8e806100206000396000f3fe608060405234801561001057600080fd5b50600436106101da5760003560e01c80639ec41a2d11610104578063ba7e3128116100a2578063ecd9ba8211610071578063ecd9ba82146104a3578063ee602ca4146104b6578063f384032814610526578063faaa8a641461053957600080fd5b8063ba7e31281461043d578063bca7093d14610450578063c96be4cb14610459578063ce119a8a1461046c57600080fd5b8063a6b63eb8116100de578063a6b63eb8146103ef578063abf410e514610402578063ac4fe5ab14610415578063b7ab4db51461042857600080fd5b80639ec41a2d14610397578063a310624f146103a0578063a694fc3a146103dc57600080fd5b806351ed6a301161017c57806378a5c2061161014b57806378a5c206146103235780637a766460146103525780638b0e9f3f1461037b578063973b294f1461038457600080fd5b806351ed6a301461029f578063561ff9a9146102ca5780635c211f88146102f0578063621379811461030357600080fd5b80633a9783f3116101b85780633a9783f3146102425780633ccfd60b146102715780633d6ec65e1461027957806342ad55ac1461028c57600080fd5b806327498240146101df578063282ec26d146101fa57806335c14de114610238575b600080fd5b6101e7610567565b6040519081526020015b60405180910390f35b610228610208366004611b56565b6001600160a01b0316600090815260096020526040902060010154151590565b60405190151581526020016101f1565b610240610578565b005b61026460405180604001604052806003815260200162646b6760e81b81525081565b6040516101f19190611b97565b610240610634565b610240610287366004611bca565b6107b6565b61022861029a366004611b56565b610882565b6002546102b2906001600160a01b031681565b6040516001600160a01b0390911681526020016101f1565b610264604051806040016040528060078152602001667374616b696e6760c81b81525081565b6000546102b2906001600160a01b031681565b610316610311366004611be3565b6108be565b6040516101f19190611c3d565b6102646040518060400160405280601081526020016f737570706f727465642d746f6b656e7360801b81525081565b6101e7610360366004611b56565b6001600160a01b031660009081526008602052604090205490565b6101e760055481565b610240610392366004611bca565b6109fe565b6101e760035481565b6103cf6103ae366004611b56565b6001600160a01b031660009081526008602052604090206001015460ff1690565b6040516101f19190611caa565b6102406103ea366004611bca565b610aca565b6102406103fd366004611cb8565b610afb565b6001546102b2906001600160a01b031681565b610228610423366004611b56565b610c73565b610430610c7c565b6040516101f19190611d13565b61024061044b366004611bca565b610c88565b6101e760045481565b610240610467366004611b56565b610de7565b610264604051806040016040528060188152602001771c995dd85c990b591a5cdd1c9a589d5d1a5bdb8b5c1bdbdb60421b81525081565b6102406104b1366004611d60565b610e5a565b61050b6104c4366004611b56565b6040805180820190915260008082526020820152506001600160a01b0316600090815260096020908152604091829020825180840190935280548352600101549082015290565b604080518251815260209283015192810192909252016101f1565b610240610534366004611d97565b610f1a565b6102646040518060400160405280600f81526020016e736c617368696e672d766f74696e6760881b81525081565b60006105736006610fe2565b905090565b61058133610c73565b156105a75760405162461bcd60e51b815260040161059e90611dc3565b60405180910390fd5b336000908152600960205260409020600101546105d65760405162461bcd60e51b815260040161059e90611dfa565b33600090815260096020908152604080832080548482556001909101849055600354600890935292205461060b908390611e63565b10158015610621575061061f600633610ff2565b155b1561063157610631336001611017565b50565b61063d33610c73565b1561065a5760405162461bcd60e51b815260040161059e90611dc3565b336000908152600960205260409020600101546106895760405162461bcd60e51b815260040161059e90611dfa565b6004543360009081526009602052604090206001015442916106aa91611e63565b11156107065760405162461bcd60e51b815260206004820152602560248201527f5374616b696e673a207769746864726177616c20706572696f64206e6f742070604482015264185cdcd95960da1b606482015260840161059e565b3360009081526009602090815260408083205460089092528220805491928392610731908490611e76565b92505081905550806005600082825461074a9190611e76565b90915550503360008181526009602052604081208181556001015560025461077e916001600160a01b03909116908361116c565b60405181815233907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b9060200160405180910390a250565b60005460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa1580156107ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108239190611e89565b6001600160a01b0316146108795760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604482015260640161059e565b610631816111d4565b600060015b6001600160a01b03831660009081526008602052604090206001015460ff1660028111156108b7576108b7611c05565b1492915050565b606060006108d66108cf6006610fe2565b8585611210565b90506108e28482611e76565b67ffffffffffffffff8111156108fa576108fa611ea6565b60405190808252806020026020018201604052801561093357816020015b610920611b07565b8152602001906001900390816109185790505b509150835b818110156109f657600061094d600683611239565b6040805180820182526001600160a01b038316808252600090815260086020908152908390208351808501909452805484526001810154949550919381850193929183019060ff1660028111156109a6576109a6611c05565b60028111156109b7576109b7611c05565b9052509052846109c78885611e76565b815181106109d7576109d7611ebc565b60200260200101819052505080806109ee90611ed2565b915050610938565b505092915050565b60005460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa158015610a47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a6b9190611e89565b6001600160a01b031614610ac15760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604482015260640161059e565b61063181611245565b610ad333610c73565b15610af05760405162461bcd60e51b815260040161059e90611dc3565b61063133338361127a565b600054600160a81b900460ff1615808015610b2357506000546001600160a01b90910460ff16105b80610b445750303b158015610b445750600054600160a01b900460ff166001145b610ba75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161059e565b6000805460ff60a01b1916600160a01b1790558015610bd4576000805460ff60a81b1916600160a81b1790555b600180546001600160a01b038781166001600160a01b031992831617909255600280548784169083161790556000805492891692909116919091179055610c1a836111d4565b610c2382611245565b8015610c6b576000805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b60006002610887565b6060610573600661149d565b610c9133610c73565b15610cae5760405162461bcd60e51b815260040161059e90611dc3565b60008111610d105760405162461bcd60e51b815260206004820152602960248201527f5374616b696e673a20616d6f756e74206d7573742062652067726561746572206044820152687468616e207a65726f60b81b606482015260840161059e565b33600090815260086020526040902054811115610d7b5760405162461bcd60e51b815260206004820152602360248201527f5374616b696e673a20616d6f756e74206d757374206265203c3d20746f207374604482015262616b6560e81b606482015260840161059e565b6040805180820182528281524260208083019182523360009081526009825284812093518455915160019093019290925560035460089092529190912054610dc4908390611e76565b108015610dd75750610dd7600633610ff2565b1561063157610631336000611017565b610def6114aa565b6001600160a01b0316336001600160a01b031614610e4f5760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a206e6f74206120736c617368696e6720766f74696e670000604482015260640161059e565b610631816002611017565b610e6333610c73565b15610e805760405162461bcd60e51b815260040161059e90611dc3565b60025460405163d505accf60e01b8152336004820152306024820152604481018790526064810186905260ff8516608482015260a4810184905260c481018390526001600160a01b039091169063d505accf9060e401600060405180830381600087803b158015610ef057600080fd5b505af1158015610f04573d6000803e3d6000fd5b50505050610f1333338761127a565b5050505050565b610f2261153f565b6001600160a01b0316336001600160a01b031614610f985760405162461bcd60e51b815260206004820152602d60248201527f5374616b696e673a206f6e6c7920526577617264446973747269627574696f6e60448201526c141bdbdb0818dbdb9d1c9858dd609a1b606482015260840161059e565b6001600160a01b03821660009081526008602052604081208054839290610fc0908490611e63565b925050819055508060056000828254610fd99190611e63565b90915550505050565b6000610fec825490565b92915050565b6001600160a01b038116600090815260018301602052604081205415155b9392505050565b60008082600281111561102c5761102c611c05565b14806110495750600282600281111561104757611047611c05565b145b156110605761105960068461159c565b905061106e565b61106b6006846115b1565b90505b806110ce5760405162461bcd60e51b815260206004820152602a60248201527f5374616b696e673a204661696c656420746f207570646174652076616c696461604482015269746f722073746174757360b01b606482015260840161059e565b6001600160a01b038316600090815260086020526040902060019081018054849260ff199091169083600281111561110857611108611c05565b02179055506111156115c6565b6001600160a01b031663b32805c36040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561114f57600080fd5b505af1158015611163573d6000803e3d6000fd5b50505050505050565b6040516001600160a01b0383166024820152604481018290526111cf90849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261160e565b505050565b60038190556040518181527fb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca906020015b60405180910390a150565b600061121c8284611e63565b9050838111156112295750825b8083111561101057509092915050565b600061101083836116e3565b60048190556040518181527f4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee90602001611205565b600081116112ca5760405162461bcd60e51b815260206004820152601a60248201527f5374616b696e673a205a65726f207374616b6520616d6f756e74000000000000604482015260640161059e565b6002546040516370a0823160e01b81526001600160a01b038581166004830152839216906370a0823190602401602060405180830381865afa158015611314573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113389190611eeb565b10156113925760405162461bcd60e51b815260206004820152602360248201527f5374616b696e673a204e6f7420656e6f75676820746f6b656e7320746f207374604482015262616b6560e81b606482015260840161059e565b6001600160a01b038216600090815260086020526040812090600182015460ff1660028111156113c4576113c4611c05565b1480156113df575060035481546113dc908490611e63565b10155b156113fe576001818101805460ff1916821790556113fe90339061170d565b600254611416906001600160a01b031685308561179b565b8181600001600082825461142a9190611e63565b9250508190555081600560008282546114439190611e63565b92505081905550826001600160a01b0316846001600160a01b03167f70e256e9264f1aa014ac7f20b4a16618647d26695e23c7181ee67a22c37e75218460405161148f91815260200190565b60405180910390a350505050565b60606000611010836117d9565b600154604080518082018252600f81526e736c617368696e672d766f74696e6760881b60208201529051633581777360e01b81526000926001600160a01b0316916335817773916114fe9190600401611b97565b602060405180830381865afa15801561151b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105739190611e89565b60015460408051808201825260188152771c995dd85c990b591a5cdd1c9a589d5d1a5bdb8b5c1bdbdb60421b60208201529051633581777360e01b81526000926001600160a01b0316916335817773916114fe9190600401611b97565b6000611010836001600160a01b038416611835565b6000611010836001600160a01b038416611928565b6001546040805180820182526003815262646b6760e81b60208201529051633581777360e01b81526000926001600160a01b0316916335817773916114fe9190600401611b97565b6000611663826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166119779092919063ffffffff16565b90508051600014806116845750808060200190518101906116849190611f04565b6111cf5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161059e565b60008260000182815481106116fa576116fa611ebc565b9060005260206000200154905092915050565b6000816117245761171f60068461159c565b61172f565b61172f6006846115b1565b9050806117935760405162461bcd60e51b815260206004820152602c60248201527f5374616b696e673a20496e76616c69642076616c696461746f7220616464726560448201526b737320746f2075706461746560a01b606482015260840161059e565b6111156115c6565b6040516001600160a01b03808516602483015283166044820152606481018290526117d39085906323b872dd60e01b90608401611198565b50505050565b60608160000180548060200260200160405190810160405280929190818152602001828054801561182957602002820191906000526020600020905b815481526020019060010190808311611815575b50505050509050919050565b6000818152600183016020526040812054801561191e576000611859600183611e76565b855490915060009061186d90600190611e76565b90508181146118d257600086600001828154811061188d5761188d611ebc565b90600052602060002001549050808760000184815481106118b0576118b0611ebc565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806118e3576118e3611f26565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610fec565b6000915050610fec565b600081815260018301602052604081205461196f57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610fec565b506000610fec565b6060611986848460008561198e565b949350505050565b6060824710156119ef5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161059e565b600080866001600160a01b03168587604051611a0b9190611f3c565b60006040518083038185875af1925050503d8060008114611a48576040519150601f19603f3d011682016040523d82523d6000602084013e611a4d565b606091505b5091509150611a5e87838387611a69565b979650505050505050565b60608315611ad8578251600003611ad1576001600160a01b0385163b611ad15760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161059e565b5081611986565b6119868383815115611aed5781518083602001fd5b8060405162461bcd60e51b815260040161059e9190611b97565b604051806040016040528060006001600160a01b03168152602001611b3c604080518082019091526000808252602082015290565b905290565b6001600160a01b038116811461063157600080fd5b600060208284031215611b6857600080fd5b813561101081611b41565b60005b83811015611b8e578181015183820152602001611b76565b50506000910152565b6020815260008251806020840152611bb6816040850160208701611b73565b601f01601f19169190910160400192915050565b600060208284031215611bdc57600080fd5b5035919050565b60008060408385031215611bf657600080fd5b50508035926020909101359150565b634e487b7160e01b600052602160045260246000fd5b60038110611c3957634e487b7160e01b600052602160045260246000fd5b9052565b602080825282518282018190526000919060409081850190868401855b82811015611c9d57815180516001600160a01b03168552860151805187860152860151611c8986860182611c1b565b506060939093019290850190600101611c5a565b5091979650505050505050565b60208101610fec8284611c1b565b600080600080600060a08688031215611cd057600080fd5b8535611cdb81611b41565b94506020860135611ceb81611b41565b93506040860135611cfb81611b41565b94979396509394606081013594506080013592915050565b6020808252825182820181905260009190848201906040850190845b81811015611d545783516001600160a01b031683529284019291840191600101611d2f565b50909695505050505050565b600080600080600060a08688031215611d7857600080fd5b8535945060208601359350604086013560ff81168114611cfb57600080fd5b60008060408385031215611daa57600080fd5b8235611db581611b41565b946020939093013593505050565b6020808252601d908201527f5374616b696e673a2076616c696461746f7220697320736c6173686564000000604082015260600190565b60208082526033908201527f5374616b696e673a207573657220646f6573206e6f7420686176652077697468604082015272191c985dd85b08185b9b9bdd5b98d95b595b9d606a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b80820180821115610fec57610fec611e4d565b81810381811115610fec57610fec611e4d565b600060208284031215611e9b57600080fd5b815161101081611b41565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600060018201611ee457611ee4611e4d565b5060010190565b600060208284031215611efd57600080fd5b5051919050565b600060208284031215611f1657600080fd5b8151801515811461101057600080fd5b634e487b7160e01b600052603160045260246000fd5b60008251611f4e818460208701611b73565b919091019291505056fea26469706673582212208731e27647f3ef28a7644e8fc6a25a6e540ea851e84cb19bdca0e0008a74263964736f6c63430008120033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_Staking *StakingCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_Staking *StakingSession) DKGKEY() (string, error) {
	return _Staking.Contract.DKGKEY(&_Staking.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_Staking *StakingCallerSession) DKGKEY() (string, error) {
	return _Staking.Contract.DKGKEY(&_Staking.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_Staking *StakingCaller) REWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_Staking *StakingSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _Staking.Contract.REWARDDISTRIBUTIONPOOLKEY(&_Staking.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_Staking *StakingCallerSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _Staking.Contract.REWARDDISTRIBUTIONPOOLKEY(&_Staking.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_Staking *StakingCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_Staking *StakingSession) SLASHINGVOTINGKEY() (string, error) {
	return _Staking.Contract.SLASHINGVOTINGKEY(&_Staking.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_Staking *StakingCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _Staking.Contract.SLASHINGVOTINGKEY(&_Staking.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_Staking *StakingCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_Staking *StakingSession) STAKINGKEY() (string, error) {
	return _Staking.Contract.STAKINGKEY(&_Staking.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_Staking *StakingCallerSession) STAKINGKEY() (string, error) {
	return _Staking.Contract.STAKINGKEY(&_Staking.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_Staking *StakingCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_Staking *StakingSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _Staking.Contract.SUPPORTEDTOKENSKEY(&_Staking.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_Staking *StakingCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _Staking.Contract.SUPPORTEDTOKENSKEY(&_Staking.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_Staking *StakingCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_Staking *StakingSession) ContractRegistry() (common.Address, error) {
	return _Staking.Contract.ContractRegistry(&_Staking.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_Staking *StakingCallerSession) ContractRegistry() (common.Address, error) {
	return _Staking.Contract.ContractRegistry(&_Staking.CallOpts)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_Staking *StakingCaller) GetStake(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStake", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_Staking *StakingSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _Staking.Contract.GetStake(&_Staking.CallOpts, _validator)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _validator) view returns(uint256)
func (_Staking *StakingCallerSession) GetStake(_validator common.Address) (*big.Int, error) {
	return _Staking.Contract.GetStake(&_Staking.CallOpts, _validator)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _validator) view returns(uint8)
func (_Staking *StakingCaller) GetValidatorStatus(opts *bind.CallOpts, _validator common.Address) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorStatus", _validator)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _validator) view returns(uint8)
func (_Staking *StakingSession) GetValidatorStatus(_validator common.Address) (uint8, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, _validator)
}

// GetValidatorStatus is a free data retrieval call binding the contract method 0xa310624f.
//
// Solidity: function getValidatorStatus(address _validator) view returns(uint8)
func (_Staking *StakingCallerSession) GetValidatorStatus(_validator common.Address) (uint8, error) {
	return _Staking.Contract.GetValidatorStatus(&_Staking.CallOpts, _validator)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingSession) GetValidators() ([]common.Address, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Staking *StakingCallerSession) GetValidators() ([]common.Address, error) {
	return _Staking.Contract.GetValidators(&_Staking.CallOpts)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Staking *StakingCaller) GetValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Staking *StakingSession) GetValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetValidatorsCount(&_Staking.CallOpts)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0x27498240.
//
// Solidity: function getValidatorsCount() view returns(uint256)
func (_Staking *StakingCallerSession) GetValidatorsCount() (*big.Int, error) {
	return _Staking.Contract.GetValidatorsCount(&_Staking.CallOpts)
}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x62137981.
//
// Solidity: function getValidatorsInfo(uint256 _offset, uint256 _limit) view returns((address,(uint256,uint8))[] _validatorsInfoArr)
func (_Staking *StakingCaller) GetValidatorsInfo(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]IStakingValidatorInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getValidatorsInfo", _offset, _limit)

	if err != nil {
		return *new([]IStakingValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingValidatorInfo)).(*[]IStakingValidatorInfo)

	return out0, err

}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x62137981.
//
// Solidity: function getValidatorsInfo(uint256 _offset, uint256 _limit) view returns((address,(uint256,uint8))[] _validatorsInfoArr)
func (_Staking *StakingSession) GetValidatorsInfo(_offset *big.Int, _limit *big.Int) ([]IStakingValidatorInfo, error) {
	return _Staking.Contract.GetValidatorsInfo(&_Staking.CallOpts, _offset, _limit)
}

// GetValidatorsInfo is a free data retrieval call binding the contract method 0x62137981.
//
// Solidity: function getValidatorsInfo(uint256 _offset, uint256 _limit) view returns((address,(uint256,uint8))[] _validatorsInfoArr)
func (_Staking *StakingCallerSession) GetValidatorsInfo(_offset *big.Int, _limit *big.Int) ([]IStakingValidatorInfo, error) {
	return _Staking.Contract.GetValidatorsInfo(&_Staking.CallOpts, _offset, _limit)
}

// GetWithdrawalAnnouncement is a free data retrieval call binding the contract method 0xee602ca4.
//
// Solidity: function getWithdrawalAnnouncement(address _userAddr) view returns((uint256,uint256))
func (_Staking *StakingCaller) GetWithdrawalAnnouncement(opts *bind.CallOpts, _userAddr common.Address) (IStakingWithdrawalAnnouncement, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getWithdrawalAnnouncement", _userAddr)

	if err != nil {
		return *new(IStakingWithdrawalAnnouncement), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakingWithdrawalAnnouncement)).(*IStakingWithdrawalAnnouncement)

	return out0, err

}

// GetWithdrawalAnnouncement is a free data retrieval call binding the contract method 0xee602ca4.
//
// Solidity: function getWithdrawalAnnouncement(address _userAddr) view returns((uint256,uint256))
func (_Staking *StakingSession) GetWithdrawalAnnouncement(_userAddr common.Address) (IStakingWithdrawalAnnouncement, error) {
	return _Staking.Contract.GetWithdrawalAnnouncement(&_Staking.CallOpts, _userAddr)
}

// GetWithdrawalAnnouncement is a free data retrieval call binding the contract method 0xee602ca4.
//
// Solidity: function getWithdrawalAnnouncement(address _userAddr) view returns((uint256,uint256))
func (_Staking *StakingCallerSession) GetWithdrawalAnnouncement(_userAddr common.Address) (IStakingWithdrawalAnnouncement, error) {
	return _Staking.Contract.GetWithdrawalAnnouncement(&_Staking.CallOpts, _userAddr)
}

// HasWithdrawalAnnouncement is a free data retrieval call binding the contract method 0x282ec26d.
//
// Solidity: function hasWithdrawalAnnouncement(address _userAddr) view returns(bool)
func (_Staking *StakingCaller) HasWithdrawalAnnouncement(opts *bind.CallOpts, _userAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "hasWithdrawalAnnouncement", _userAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasWithdrawalAnnouncement is a free data retrieval call binding the contract method 0x282ec26d.
//
// Solidity: function hasWithdrawalAnnouncement(address _userAddr) view returns(bool)
func (_Staking *StakingSession) HasWithdrawalAnnouncement(_userAddr common.Address) (bool, error) {
	return _Staking.Contract.HasWithdrawalAnnouncement(&_Staking.CallOpts, _userAddr)
}

// HasWithdrawalAnnouncement is a free data retrieval call binding the contract method 0x282ec26d.
//
// Solidity: function hasWithdrawalAnnouncement(address _userAddr) view returns(bool)
func (_Staking *StakingCallerSession) HasWithdrawalAnnouncement(_userAddr common.Address) (bool, error) {
	return _Staking.Contract.HasWithdrawalAnnouncement(&_Staking.CallOpts, _userAddr)
}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address _validator) view returns(bool)
func (_Staking *StakingCaller) IsValidatorActive(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isValidatorActive", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address _validator) view returns(bool)
func (_Staking *StakingSession) IsValidatorActive(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorActive(&_Staking.CallOpts, _validator)
}

// IsValidatorActive is a free data retrieval call binding the contract method 0x42ad55ac.
//
// Solidity: function isValidatorActive(address _validator) view returns(bool)
func (_Staking *StakingCallerSession) IsValidatorActive(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorActive(&_Staking.CallOpts, _validator)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validator) view returns(bool)
func (_Staking *StakingCaller) IsValidatorSlashed(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isValidatorSlashed", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validator) view returns(bool)
func (_Staking *StakingSession) IsValidatorSlashed(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorSlashed(&_Staking.CallOpts, _validator)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validator) view returns(bool)
func (_Staking *StakingCallerSession) IsValidatorSlashed(_validator common.Address) (bool, error) {
	return _Staking.Contract.IsValidatorSlashed(&_Staking.CallOpts, _validator)
}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_Staking *StakingCaller) MinimalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "minimalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_Staking *StakingSession) MinimalStake() (*big.Int, error) {
	return _Staking.Contract.MinimalStake(&_Staking.CallOpts)
}

// MinimalStake is a free data retrieval call binding the contract method 0x9ec41a2d.
//
// Solidity: function minimalStake() view returns(uint256)
func (_Staking *StakingCallerSession) MinimalStake() (*big.Int, error) {
	return _Staking.Contract.MinimalStake(&_Staking.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_Staking *StakingCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_Staking *StakingSession) SignerGetter() (common.Address, error) {
	return _Staking.Contract.SignerGetter(&_Staking.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_Staking *StakingCallerSession) SignerGetter() (common.Address, error) {
	return _Staking.Contract.SignerGetter(&_Staking.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Staking *StakingCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Staking *StakingSession) StakeToken() (common.Address, error) {
	return _Staking.Contract.StakeToken(&_Staking.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Staking *StakingCallerSession) StakeToken() (common.Address, error) {
	return _Staking.Contract.StakeToken(&_Staking.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Staking *StakingCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Staking *StakingSession) TotalStake() (*big.Int, error) {
	return _Staking.Contract.TotalStake(&_Staking.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStake() (*big.Int, error) {
	return _Staking.Contract.TotalStake(&_Staking.CallOpts)
}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_Staking *StakingCaller) WithdrawalPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawalPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_Staking *StakingSession) WithdrawalPeriod() (*big.Int, error) {
	return _Staking.Contract.WithdrawalPeriod(&_Staking.CallOpts)
}

// WithdrawalPeriod is a free data retrieval call binding the contract method 0xbca7093d.
//
// Solidity: function withdrawalPeriod() view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawalPeriod() (*big.Int, error) {
	return _Staking.Contract.WithdrawalPeriod(&_Staking.CallOpts)
}

// AddRewardsToStake is a paid mutator transaction binding the contract method 0xf3840328.
//
// Solidity: function addRewardsToStake(address _validator, uint256 _amount) returns()
func (_Staking *StakingTransactor) AddRewardsToStake(opts *bind.TransactOpts, _validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addRewardsToStake", _validator, _amount)
}

// AddRewardsToStake is a paid mutator transaction binding the contract method 0xf3840328.
//
// Solidity: function addRewardsToStake(address _validator, uint256 _amount) returns()
func (_Staking *StakingSession) AddRewardsToStake(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AddRewardsToStake(&_Staking.TransactOpts, _validator, _amount)
}

// AddRewardsToStake is a paid mutator transaction binding the contract method 0xf3840328.
//
// Solidity: function addRewardsToStake(address _validator, uint256 _amount) returns()
func (_Staking *StakingTransactorSession) AddRewardsToStake(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AddRewardsToStake(&_Staking.TransactOpts, _validator, _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_Staking *StakingTransactor) AnnounceWithdrawal(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "announceWithdrawal", _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_Staking *StakingSession) AnnounceWithdrawal(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AnnounceWithdrawal(&_Staking.TransactOpts, _amount)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0xba7e3128.
//
// Solidity: function announceWithdrawal(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) AnnounceWithdrawal(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.AnnounceWithdrawal(&_Staking.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address _signerGetterAddress, address _contractRegistry, address _stakeToken, uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _signerGetterAddress common.Address, _contractRegistry common.Address, _stakeToken common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _signerGetterAddress, _contractRegistry, _stakeToken, _minimalStake, _withdrawalPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address _signerGetterAddress, address _contractRegistry, address _stakeToken, uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_Staking *StakingSession) Initialize(_signerGetterAddress common.Address, _contractRegistry common.Address, _stakeToken common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _signerGetterAddress, _contractRegistry, _stakeToken, _minimalStake, _withdrawalPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6b63eb8.
//
// Solidity: function initialize(address _signerGetterAddress, address _contractRegistry, address _stakeToken, uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactorSession) Initialize(_signerGetterAddress common.Address, _contractRegistry common.Address, _stakeToken common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _signerGetterAddress, _contractRegistry, _stakeToken, _minimalStake, _withdrawalPeriod)
}

// RevokeWithdrawal is a paid mutator transaction binding the contract method 0x35c14de1.
//
// Solidity: function revokeWithdrawal() returns()
func (_Staking *StakingTransactor) RevokeWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "revokeWithdrawal")
}

// RevokeWithdrawal is a paid mutator transaction binding the contract method 0x35c14de1.
//
// Solidity: function revokeWithdrawal() returns()
func (_Staking *StakingSession) RevokeWithdrawal() (*types.Transaction, error) {
	return _Staking.Contract.RevokeWithdrawal(&_Staking.TransactOpts)
}

// RevokeWithdrawal is a paid mutator transaction binding the contract method 0x35c14de1.
//
// Solidity: function revokeWithdrawal() returns()
func (_Staking *StakingTransactorSession) RevokeWithdrawal() (*types.Transaction, error) {
	return _Staking.Contract.RevokeWithdrawal(&_Staking.TransactOpts)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_Staking *StakingTransactor) SetMinimalStake(opts *bind.TransactOpts, _minimalStake *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setMinimalStake", _minimalStake)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_Staking *StakingSession) SetMinimalStake(_minimalStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinimalStake(&_Staking.TransactOpts, _minimalStake)
}

// SetMinimalStake is a paid mutator transaction binding the contract method 0x3d6ec65e.
//
// Solidity: function setMinimalStake(uint256 _minimalStake) returns()
func (_Staking *StakingTransactorSession) SetMinimalStake(_minimalStake *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetMinimalStake(&_Staking.TransactOpts, _minimalStake)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactor) SetWithdrawalPeriod(opts *bind.TransactOpts, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setWithdrawalPeriod", _withdrawalPeriod)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_Staking *StakingSession) SetWithdrawalPeriod(_withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetWithdrawalPeriod(&_Staking.TransactOpts, _withdrawalPeriod)
}

// SetWithdrawalPeriod is a paid mutator transaction binding the contract method 0x973b294f.
//
// Solidity: function setWithdrawalPeriod(uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactorSession) SetWithdrawalPeriod(_withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetWithdrawalPeriod(&_Staking.TransactOpts, _withdrawalPeriod)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_Staking *StakingTransactor) Slash(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slash", _validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_Staking *StakingSession) Slash(_validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _validator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address _validator) returns()
func (_Staking *StakingTransactorSession) Slash(_validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Slash(&_Staking.TransactOpts, _validator)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, _stakeAmount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) returns()
func (_Staking *StakingSession) Stake(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, _stakeAmount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _stakeAmount) returns()
func (_Staking *StakingTransactorSession) Stake(_stakeAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, _stakeAmount)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 _stakeAmount, uint256 _sigExpirationTime, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Staking *StakingTransactor) StakeWithPermit(opts *bind.TransactOpts, _stakeAmount *big.Int, _sigExpirationTime *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeWithPermit", _stakeAmount, _sigExpirationTime, _v, _r, _s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 _stakeAmount, uint256 _sigExpirationTime, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Staking *StakingSession) StakeWithPermit(_stakeAmount *big.Int, _sigExpirationTime *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.StakeWithPermit(&_Staking.TransactOpts, _stakeAmount, _sigExpirationTime, _v, _r, _s)
}

// StakeWithPermit is a paid mutator transaction binding the contract method 0xecd9ba82.
//
// Solidity: function stakeWithPermit(uint256 _stakeAmount, uint256 _sigExpirationTime, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_Staking *StakingTransactorSession) StakeWithPermit(_stakeAmount *big.Int, _sigExpirationTime *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _Staking.Contract.StakeWithPermit(&_Staking.TransactOpts, _stakeAmount, _sigExpirationTime, _v, _r, _s)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingSession) Withdraw() (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staking *StakingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts)
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingMinimalStakeUpdatedIterator is returned from FilterMinimalStakeUpdated and is used to iterate over the raw logs and unpacked data for MinimalStakeUpdated events raised by the Staking contract.
type StakingMinimalStakeUpdatedIterator struct {
	Event *StakingMinimalStakeUpdated // Event containing the contract specifics and raw log

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
func (it *StakingMinimalStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingMinimalStakeUpdated)
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
		it.Event = new(StakingMinimalStakeUpdated)
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
func (it *StakingMinimalStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingMinimalStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingMinimalStakeUpdated represents a MinimalStakeUpdated event raised by the Staking contract.
type StakingMinimalStakeUpdated struct {
	MinimalStake *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMinimalStakeUpdated is a free log retrieval operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_Staking *StakingFilterer) FilterMinimalStakeUpdated(opts *bind.FilterOpts) (*StakingMinimalStakeUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "MinimalStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingMinimalStakeUpdatedIterator{contract: _Staking.contract, event: "MinimalStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchMinimalStakeUpdated is a free log subscription operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_Staking *StakingFilterer) WatchMinimalStakeUpdated(opts *bind.WatchOpts, sink chan<- *StakingMinimalStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "MinimalStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingMinimalStakeUpdated)
				if err := _Staking.contract.UnpackLog(event, "MinimalStakeUpdated", log); err != nil {
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

// ParseMinimalStakeUpdated is a log parse operation binding the contract event 0xb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca.
//
// Solidity: event MinimalStakeUpdated(uint256 minimalStake)
func (_Staking *StakingFilterer) ParseMinimalStakeUpdated(log types.Log) (*StakingMinimalStakeUpdated, error) {
	event := new(StakingMinimalStakeUpdated)
	if err := _Staking.contract.UnpackLog(event, "MinimalStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTokensStakedIterator is returned from FilterTokensStaked and is used to iterate over the raw logs and unpacked data for TokensStaked events raised by the Staking contract.
type StakingTokensStakedIterator struct {
	Event *StakingTokensStaked // Event containing the contract specifics and raw log

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
func (it *StakingTokensStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTokensStaked)
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
		it.Event = new(StakingTokensStaked)
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
func (it *StakingTokensStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTokensStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTokensStaked represents a TokensStaked event raised by the Staking contract.
type StakingTokensStaked struct {
	TokensSender    common.Address
	TokensRecipient common.Address
	StakeAmount     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokensStaked is a free log retrieval operation binding the contract event 0x70e256e9264f1aa014ac7f20b4a16618647d26695e23c7181ee67a22c37e7521.
//
// Solidity: event TokensStaked(address indexed tokensSender, address indexed tokensRecipient, uint256 stakeAmount)
func (_Staking *StakingFilterer) FilterTokensStaked(opts *bind.FilterOpts, tokensSender []common.Address, tokensRecipient []common.Address) (*StakingTokensStakedIterator, error) {

	var tokensSenderRule []interface{}
	for _, tokensSenderItem := range tokensSender {
		tokensSenderRule = append(tokensSenderRule, tokensSenderItem)
	}
	var tokensRecipientRule []interface{}
	for _, tokensRecipientItem := range tokensRecipient {
		tokensRecipientRule = append(tokensRecipientRule, tokensRecipientItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "TokensStaked", tokensSenderRule, tokensRecipientRule)
	if err != nil {
		return nil, err
	}
	return &StakingTokensStakedIterator{contract: _Staking.contract, event: "TokensStaked", logs: logs, sub: sub}, nil
}

// WatchTokensStaked is a free log subscription operation binding the contract event 0x70e256e9264f1aa014ac7f20b4a16618647d26695e23c7181ee67a22c37e7521.
//
// Solidity: event TokensStaked(address indexed tokensSender, address indexed tokensRecipient, uint256 stakeAmount)
func (_Staking *StakingFilterer) WatchTokensStaked(opts *bind.WatchOpts, sink chan<- *StakingTokensStaked, tokensSender []common.Address, tokensRecipient []common.Address) (event.Subscription, error) {

	var tokensSenderRule []interface{}
	for _, tokensSenderItem := range tokensSender {
		tokensSenderRule = append(tokensSenderRule, tokensSenderItem)
	}
	var tokensRecipientRule []interface{}
	for _, tokensRecipientItem := range tokensRecipient {
		tokensRecipientRule = append(tokensRecipientRule, tokensRecipientItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "TokensStaked", tokensSenderRule, tokensRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTokensStaked)
				if err := _Staking.contract.UnpackLog(event, "TokensStaked", log); err != nil {
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

// ParseTokensStaked is a log parse operation binding the contract event 0x70e256e9264f1aa014ac7f20b4a16618647d26695e23c7181ee67a22c37e7521.
//
// Solidity: event TokensStaked(address indexed tokensSender, address indexed tokensRecipient, uint256 stakeAmount)
func (_Staking *StakingFilterer) ParseTokensStaked(log types.Log) (*StakingTokensStaked, error) {
	event := new(StakingTokensStaked)
	if err := _Staking.contract.UnpackLog(event, "TokensStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the Staking contract.
type StakingTokensWithdrawnIterator struct {
	Event *StakingTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTokensWithdrawn)
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
		it.Event = new(StakingTokensWithdrawn)
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
func (it *StakingTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTokensWithdrawn represents a TokensWithdrawn event raised by the Staking contract.
type StakingTokensWithdrawn struct {
	UserAddr     common.Address
	TokensAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed userAddr, uint256 tokensAmount)
func (_Staking *StakingFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, userAddr []common.Address) (*StakingTokensWithdrawnIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "TokensWithdrawn", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingTokensWithdrawnIterator{contract: _Staking.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed userAddr, uint256 tokensAmount)
func (_Staking *StakingFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingTokensWithdrawn, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "TokensWithdrawn", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTokensWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b.
//
// Solidity: event TokensWithdrawn(address indexed userAddr, uint256 tokensAmount)
func (_Staking *StakingFilterer) ParseTokensWithdrawn(log types.Log) (*StakingTokensWithdrawn, error) {
	event := new(StakingTokensWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawalPeriodUpdatedIterator is returned from FilterWithdrawalPeriodUpdated and is used to iterate over the raw logs and unpacked data for WithdrawalPeriodUpdated events raised by the Staking contract.
type StakingWithdrawalPeriodUpdatedIterator struct {
	Event *StakingWithdrawalPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawalPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawalPeriodUpdated)
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
		it.Event = new(StakingWithdrawalPeriodUpdated)
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
func (it *StakingWithdrawalPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawalPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawalPeriodUpdated represents a WithdrawalPeriodUpdated event raised by the Staking contract.
type StakingWithdrawalPeriodUpdated struct {
	WithdrawalPeriod *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalPeriodUpdated is a free log retrieval operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_Staking *StakingFilterer) FilterWithdrawalPeriodUpdated(opts *bind.FilterOpts) (*StakingWithdrawalPeriodUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawalPeriodUpdatedIterator{contract: _Staking.contract, event: "WithdrawalPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalPeriodUpdated is a free log subscription operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_Staking *StakingFilterer) WatchWithdrawalPeriodUpdated(opts *bind.WatchOpts, sink chan<- *StakingWithdrawalPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WithdrawalPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawalPeriodUpdated)
				if err := _Staking.contract.UnpackLog(event, "WithdrawalPeriodUpdated", log); err != nil {
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

// ParseWithdrawalPeriodUpdated is a log parse operation binding the contract event 0x4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee.
//
// Solidity: event WithdrawalPeriodUpdated(uint256 withdrawalPeriod)
func (_Staking *StakingFilterer) ParseWithdrawalPeriodUpdated(log types.Log) (*StakingWithdrawalPeriodUpdated, error) {
	event := new(StakingWithdrawalPeriodUpdated)
	if err := _Staking.contract.UnpackLog(event, "WithdrawalPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
