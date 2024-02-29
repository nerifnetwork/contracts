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

// StakingValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type StakingValidatorInfo struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contractRegistry\",\"type\":\"address\"}],\"name\":\"ContractRegistryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"}],\"name\":\"MinimalStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"WithdrawalPeriodUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addRewardsToStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressStorage\",\"outputs\":[{\"internalType\":\"contractAddressStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"announceWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorStorage\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"listValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structStaking.ValidatorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"}],\"name\":\"setMinimalStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawalAnnouncements\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611b40806100206000396000f3fe6080604052600436106101c65760003560e01c80637a766460116100f7578063b7ab4db511610095578063ce119a8a11610064578063ce119a8a1461057c578063e7f166ed146105c0578063f384032814610609578063faaa8a641461062957600080fd5b8063b7ab4db514610504578063ba7e312814610526578063bca7093d14610546578063c96be4cb1461055c57600080fd5b8063973b294f116100d1578063973b294f1461048e5780639ec41a2d146104ae578063abf410e5146104c4578063ac4fe5ab146104e457600080fd5b80637a7664601461042a5780637bc74225146104635780638b0e9f3f1461047857600080fd5b80633ccfd60b11610164578063561ff9a91161013e578063561ff9a9146103635780635c211f8814610396578063612e1488146103ce57806378a5c206146103ee57600080fd5b80633ccfd60b146102fe5780633d6ec65e1461031357806342ad55ac1461033357600080fd5b806327498240116101a0578063274982401461028257806335c14de1146102a55780633a4b66f1146102ba5780633a9783f3146102c257600080fd5b806303b54d52146101d257806316934fc4146101f45780631c78cb381461025557600080fd5b366101cd57005b600080fd5b3480156101de57600080fd5b506101f26101ed366004611687565b610664565b005b34801561020057600080fd5b5061023d61020f3660046116e6565b6004602052600090815260409020805460018201546002909201546001600160a01b03909116919060ff1683565b60405161024c93929190611742565b60405180910390f35b34801561026157600080fd5b5061027561027036600461176e565b6107ef565b60405161024c9190611790565b34801561028e57600080fd5b5061029761094f565b60405190815260200161024c565b3480156102b157600080fd5b506101f26109c2565b6101f2610b0b565b3480156102ce57600080fd5b506102f160405180604001604052806003815260200162646b6760e81b81525081565b60405161024c91906117fc565b34801561030a57600080fd5b506101f2610c41565b34801561031f57600080fd5b506101f261032e36600461184a565b610ea9565b34801561033f57600080fd5b5061035361034e3660046116e6565b610fa8565b604051901515815260200161024c565b34801561036f57600080fd5b506102f1604051806040016040528060078152602001667374616b696e6760c81b81525081565b3480156103a257600080fd5b506000546103b6906001600160a01b031681565b6040516001600160a01b03909116815260200161024c565b3480156103da57600080fd5b506007546103b6906001600160a01b031681565b3480156103fa57600080fd5b506102f16040518060400160405280601081526020016f737570706f727465642d746f6b656e7360801b81525081565b34801561043657600080fd5b506102976104453660046116e6565b6001600160a01b031660009081526004602052604090206001015490565b34801561046f57600080fd5b50600354610297565b34801561048457600080fd5b5061029760035481565b34801561049a57600080fd5b506101f26104a936600461184a565b610fe5565b3480156104ba57600080fd5b5061029760015481565b3480156104d057600080fd5b506006546103b6906001600160a01b031681565b3480156104f057600080fd5b506103536104ff3660046116e6565b6110dd565b34801561051057600080fd5b506105196110e6565b60405161024c9190611863565b34801561053257600080fd5b506101f261054136600461184a565b611158565b34801561055257600080fd5b5061029760025481565b34801561056857600080fd5b506101f26105773660046116e6565b6112db565b34801561058857600080fd5b506102f1604051806040016040528060188152602001771c995dd85c990b591a5cdd1c9a589d5d1a5bdb8b5c1bdbdb60421b81525081565b3480156105cc57600080fd5b506105f46105db3660046116e6565b6005602052600090815260409020805460019091015482565b6040805192835260208301919091520161024c565b34801561061557600080fd5b506101f26106243660046118b0565b611374565b34801561063557600080fd5b506102f16040518060400160405280600f81526020016e736c617368696e672d766f74696e6760881b81525081565b600054600160a81b900460ff161580801561068c57506000546001600160a01b90910460ff16105b806106ad5750303b1580156106ad5750600054600160a01b900460ff166001145b6107155760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff60a01b1916600160a01b1790558015610742576000805460ff60a81b1916600160a81b1790555b600080546001600160a01b0319166001600160a01b03881617905561076685610ea9565b61076f84610fe5565b600680546001600160a01b038086166001600160a01b031992831617909255600780549285169290911691909117905580156107e7576000805460ff60a81b19169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b606060006107fb6110e6565b905060008367ffffffffffffffff811115610818576108186118dc565b60405190808252806020026020018201604052801561086a57816020015b61085760408051606081018252600080825260208201819052909182015290565b8152602001906001900390816108365790505b509050845b6108798587611908565b81101561094457600460008483815181106108965761089661191b565b6020908102919091018101516001600160a01b039081168352828201939093526040918201600020825160608101845281549094168452600181015491840191909152600280820154919284019160ff16908111156108f7576108f761170a565b60028111156109085761090861170a565b905250826109168884611931565b815181106109265761092661191b565b6020026020010181905250808061093c90611944565b91505061086f565b509150505b92915050565b6007546040805163949d225d60e01b815290516000926001600160a01b03169163949d225d9160048083019260209291908290030181865afa158015610999573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bd919061195d565b905090565b33600090815260046020526040902060029081015460ff16818111156109ea576109ea61170a565b03610a075760405162461bcd60e51b815260040161070c90611976565b33600090815260056020526040902054610a5c5760405162461bcd60e51b815260206004820152601660248201527514dd185ada5b99ce881b9bdd08185b9b9bdd5b98d95960521b604482015260640161070c565b336000908152600560205260408120805482825560019091018290559033600090815260046020526040902060029081015460ff1690811115610aa157610aa161170a565b148015610acd57506001805433600090815260046020526040902090910154610aca9083611908565b10155b15610b085733600081815260046020526040902080546001600160a01b03191682178155600201805460ff19166001179055610b089061143f565b50565b33600090815260046020526040902060029081015460ff1681811115610b3357610b3361170a565b03610b505760405162461bcd60e51b815260040161070c90611976565b60003411610b705760405162461bcd60e51b815260040161070c906119ad565b33600090815260046020526040812060029081015460ff1690811115610b9857610b9861170a565b148015610bc457506001805433600090815260046020526040902090910154610bc19034611908565b10155b15610bff5733600081815260046020526040902080546001600160a01b03191682178155600201805460ff19166001179055610bff9061143f565b3360009081526004602052604081206001018054349290610c21908490611908565b925050819055503460036000828254610c3a9190611908565b9091555050565b33600090815260046020526040902060029081015460ff1681811115610c6957610c6961170a565b03610c865760405162461bcd60e51b815260040161070c90611976565b33600090815260056020526040902054610cb25760405162461bcd60e51b815260040161070c906119ad565b600254336000908152600560205260409020600101544291610cd391611908565b1115610d2f5760405162461bcd60e51b815260206004820152602560248201527f5374616b696e673a207769746864726177616c20706572696f64206e6f742070604482015264185cdcd95960da1b606482015260840161070c565b33600090815260056020908152604080832054600490925290912060010154811115610db35760405162461bcd60e51b815260206004820152602d60248201527f5374616b696e673a20616d6f756e74206d757374206265203c3d20746f20766160448201526c6c696461746f72207374616b6560981b606482015260840161070c565b3360009081526004602052604081206001018054839290610dd5908490611931565b925050819055508060036000828254610dee9190611931565b9091555050336000818152600560205260408082208281556001018290555190919061520890849084818181858888f193505050503d8060008114610e4f576040519150601f19603f3d011682016040523d82523d6000602084013e610e54565b606091505b5050905080610ea55760405162461bcd60e51b815260206004820152601860248201527f5374616b696e673a207472616e73666572206661696c65640000000000000000604482015260640161070c565b5050565b60005460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa158015610ef2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f169190611a06565b6001600160a01b031614610f6c5760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604482015260640161070c565b60018190556040518181527fb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca906020015b60405180910390a150565b600060015b6001600160a01b038316600090815260046020526040902060029081015460ff1690811115610fde57610fde61170a565b1492915050565b60005460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa15801561102e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110529190611a06565b6001600160a01b0316146110a85760405162461bcd60e51b815260206004820152601a60248201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604482015260640161070c565b60028190556040518181527f4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee90602001610f9d565b60006002610fad565b600754604080516351cfd60960e11b815290516060926001600160a01b03169163a39fac129160048083019260009291908290030181865afa158015611130573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109bd9190810190611a23565b33600090815260046020526040902060029081015460ff16818111156111805761118061170a565b0361119d5760405162461bcd60e51b815260040161070c90611976565b3360009081526004602052604090206001015481111561120b5760405162461bcd60e51b815260206004820152602360248201527f5374616b696e673a20616d6f756e74206d757374206265203c3d20746f207374604482015262616b6560e81b606482015260840161070c565b336000908152600560209081526040808320848155426001918201558054600490935292209091015461123f908390611931565b1080156112b35750600754604051630bb7c8fd60e31b81523360048201526001600160a01b0390911690635dbe47e890602401602060405180830381865afa15801561128f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b39190611ae8565b15610b0857336000818152600460205260409020600201805460ff19169055610b08906114fa565b6112e3611538565b6001600160a01b0316336001600160a01b0316146113435760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a206e6f74206120736c617368696e6720766f74696e670000604482015260640161070c565b6001600160a01b03811660009081526004602052604090206002908101805460ff19169091179055610b08816114fa565b61137c6115cd565b6001600160a01b0316336001600160a01b0316146113f25760405162461bcd60e51b815260206004820152602d60248201527f5374616b696e673a206f6e6c7920526577617264446973747269627574696f6e60448201526c141bdbdb0818dbdb9d1c9858dd609a1b606482015260840161070c565b6001600160a01b0382166000908152600460205260408120600101805483929061141d908490611908565b9250508190555080600360008282546114369190611908565b90915550505050565b600061144961162a565b600754604051639e9405c360e01b81526001600160a01b038581166004830152929350911690639e9405c3906024015b600060405180830381600087803b15801561149357600080fd5b505af11580156114a7573d6000803e3d6000fd5b50505050806001600160a01b031663b32805c36040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156114e657600080fd5b505af11580156107e7573d6000803e3d6000fd5b600061150461162a565b600754604051636196c02d60e11b81526001600160a01b03858116600483015292935091169063c32d805a90602401611479565b600654604080518082018252600f81526e736c617368696e672d766f74696e6760881b60208201529051633581777360e01b81526000926001600160a01b03169163358177739161158c91906004016117fc565b602060405180830381865afa1580156115a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109bd9190611a06565b60065460408051808201825260188152771c995dd85c990b591a5cdd1c9a589d5d1a5bdb8b5c1bdbdb60421b60208201529051633581777360e01b81526000926001600160a01b03169163358177739161158c91906004016117fc565b6006546040805180820182526003815262646b6760e81b60208201529051633581777360e01b81526000926001600160a01b03169163358177739161158c91906004016117fc565b6001600160a01b0381168114610b0857600080fd5b600080600080600060a0868803121561169f57600080fd5b85356116aa81611672565b9450602086013593506040860135925060608601356116c881611672565b915060808601356116d881611672565b809150509295509295909350565b6000602082840312156116f857600080fd5b813561170381611672565b9392505050565b634e487b7160e01b600052602160045260246000fd5b6003811061173e57634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b038416815260208101839052606081016117666040830184611720565b949350505050565b6000806040838503121561178157600080fd5b50508035926020909101359150565b602080825282518282018190526000919060409081850190868401855b828110156117ef57815180516001600160a01b0316855286810151878601528501516117db86860182611720565b5060609390930192908501906001016117ad565b5091979650505050505050565b600060208083528351808285015260005b818110156118295785810183015185820160400152820161180d565b506000604082860101526040601f19601f8301168501019250505092915050565b60006020828403121561185c57600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b818110156118a45783516001600160a01b03168352928401929184019160010161187f565b50909695505050505050565b600080604083850312156118c357600080fd5b82356118ce81611672565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610949576109496118f2565b634e487b7160e01b600052603260045260246000fd5b81810381811115610949576109496118f2565b600060018201611956576119566118f2565b5060010190565b60006020828403121561196f57600080fd5b5051919050565b6020808252601d908201527f5374616b696e673a2076616c696461746f7220697320736c6173686564000000604082015260600190565b60208082526029908201527f5374616b696e673a20616d6f756e74206d7573742062652067726561746572206040820152687468616e207a65726f60b81b606082015260800190565b8051611a0181611672565b919050565b600060208284031215611a1857600080fd5b815161170381611672565b60006020808385031215611a3657600080fd5b825167ffffffffffffffff80821115611a4e57600080fd5b818501915085601f830112611a6257600080fd5b815181811115611a7457611a746118dc565b8060051b604051601f19603f83011681018181108582111715611a9957611a996118dc565b604052918252848201925083810185019188831115611ab757600080fd5b938501935b82851015611adc57611acd856119f6565b84529385019392850192611abc565b98975050505050505050565b600060208284031215611afa57600080fd5b8151801515811461170357600080fdfea264697066735822122006380d20c01162d62e7ec9282a965088b7d188aa4e4c42a399e4c70d15294ca464736f6c63430008120033",
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

// AddressStorage is a free data retrieval call binding the contract method 0x612e1488.
//
// Solidity: function addressStorage() view returns(address)
func (_Staking *StakingCaller) AddressStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "addressStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressStorage is a free data retrieval call binding the contract method 0x612e1488.
//
// Solidity: function addressStorage() view returns(address)
func (_Staking *StakingSession) AddressStorage() (common.Address, error) {
	return _Staking.Contract.AddressStorage(&_Staking.CallOpts)
}

// AddressStorage is a free data retrieval call binding the contract method 0x612e1488.
//
// Solidity: function addressStorage() view returns(address)
func (_Staking *StakingCallerSession) AddressStorage() (common.Address, error) {
	return _Staking.Contract.AddressStorage(&_Staking.CallOpts)
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

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Staking *StakingCaller) GetTotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getTotalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Staking *StakingSession) GetTotalStake() (*big.Int, error) {
	return _Staking.Contract.GetTotalStake(&_Staking.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x7bc74225.
//
// Solidity: function getTotalStake() view returns(uint256)
func (_Staking *StakingCallerSession) GetTotalStake() (*big.Int, error) {
	return _Staking.Contract.GetTotalStake(&_Staking.CallOpts)
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

// ListValidators is a free data retrieval call binding the contract method 0x1c78cb38.
//
// Solidity: function listValidators(uint256 _offset, uint256 _limit) view returns((address,uint256,uint8)[])
func (_Staking *StakingCaller) ListValidators(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]StakingValidatorInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "listValidators", _offset, _limit)

	if err != nil {
		return *new([]StakingValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingValidatorInfo)).(*[]StakingValidatorInfo)

	return out0, err

}

// ListValidators is a free data retrieval call binding the contract method 0x1c78cb38.
//
// Solidity: function listValidators(uint256 _offset, uint256 _limit) view returns((address,uint256,uint8)[])
func (_Staking *StakingSession) ListValidators(_offset *big.Int, _limit *big.Int) ([]StakingValidatorInfo, error) {
	return _Staking.Contract.ListValidators(&_Staking.CallOpts, _offset, _limit)
}

// ListValidators is a free data retrieval call binding the contract method 0x1c78cb38.
//
// Solidity: function listValidators(uint256 _offset, uint256 _limit) view returns((address,uint256,uint8)[])
func (_Staking *StakingCallerSession) ListValidators(_offset *big.Int, _limit *big.Int) ([]StakingValidatorInfo, error) {
	return _Staking.Contract.ListValidators(&_Staking.CallOpts, _offset, _limit)
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

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_Staking *StakingCaller) Stakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakes", arg0)

	outstruct := new(struct {
		Validator common.Address
		Stake     *big.Int
		Status    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Stake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_Staking *StakingSession) Stakes(arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
}

// Stakes is a free data retrieval call binding the contract method 0x16934fc4.
//
// Solidity: function stakes(address ) view returns(address validator, uint256 stake, uint8 status)
func (_Staking *StakingCallerSession) Stakes(arg0 common.Address) (struct {
	Validator common.Address
	Stake     *big.Int
	Status    uint8
}, error) {
	return _Staking.Contract.Stakes(&_Staking.CallOpts, arg0)
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

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_Staking *StakingCaller) WithdrawalAnnouncements(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawalAnnouncements", arg0)

	outstruct := new(struct {
		Amount *big.Int
		Time   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_Staking *StakingSession) WithdrawalAnnouncements(arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _Staking.Contract.WithdrawalAnnouncements(&_Staking.CallOpts, arg0)
}

// WithdrawalAnnouncements is a free data retrieval call binding the contract method 0xe7f166ed.
//
// Solidity: function withdrawalAnnouncements(address ) view returns(uint256 amount, uint256 time)
func (_Staking *StakingCallerSession) WithdrawalAnnouncements(arg0 common.Address) (struct {
	Amount *big.Int
	Time   *big.Int
}, error) {
	return _Staking.Contract.WithdrawalAnnouncements(&_Staking.CallOpts, arg0)
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

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _signerGetterAddress, uint256 _minimalStake, uint256 _withdrawalPeriod, address _contractRegistry, address _validatorStorage) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _signerGetterAddress common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int, _contractRegistry common.Address, _validatorStorage common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _signerGetterAddress, _minimalStake, _withdrawalPeriod, _contractRegistry, _validatorStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _signerGetterAddress, uint256 _minimalStake, uint256 _withdrawalPeriod, address _contractRegistry, address _validatorStorage) returns()
func (_Staking *StakingSession) Initialize(_signerGetterAddress common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int, _contractRegistry common.Address, _validatorStorage common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _signerGetterAddress, _minimalStake, _withdrawalPeriod, _contractRegistry, _validatorStorage)
}

// Initialize is a paid mutator transaction binding the contract method 0x03b54d52.
//
// Solidity: function initialize(address _signerGetterAddress, uint256 _minimalStake, uint256 _withdrawalPeriod, address _contractRegistry, address _validatorStorage) returns()
func (_Staking *StakingTransactorSession) Initialize(_signerGetterAddress common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int, _contractRegistry common.Address, _validatorStorage common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _signerGetterAddress, _minimalStake, _withdrawalPeriod, _contractRegistry, _validatorStorage)
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

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake")
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x3a4b66f1.
//
// Solidity: function stake() payable returns()
func (_Staking *StakingTransactorSession) Stake() (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts)
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingContractRegistryUpdatedIterator is returned from FilterContractRegistryUpdated and is used to iterate over the raw logs and unpacked data for ContractRegistryUpdated events raised by the Staking contract.
type StakingContractRegistryUpdatedIterator struct {
	Event *StakingContractRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *StakingContractRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractRegistryUpdated)
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
		it.Event = new(StakingContractRegistryUpdated)
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
func (it *StakingContractRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractRegistryUpdated represents a ContractRegistryUpdated event raised by the Staking contract.
type StakingContractRegistryUpdated struct {
	ContractRegistry common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterContractRegistryUpdated is a free log retrieval operation binding the contract event 0xb3a6e7c81ebdb0cf9bf28d5ddf6678ded8d73b44f2eee06ca3974dbb9f41ce7e.
//
// Solidity: event ContractRegistryUpdated(address contractRegistry)
func (_Staking *StakingFilterer) FilterContractRegistryUpdated(opts *bind.FilterOpts) (*StakingContractRegistryUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ContractRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingContractRegistryUpdatedIterator{contract: _Staking.contract, event: "ContractRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchContractRegistryUpdated is a free log subscription operation binding the contract event 0xb3a6e7c81ebdb0cf9bf28d5ddf6678ded8d73b44f2eee06ca3974dbb9f41ce7e.
//
// Solidity: event ContractRegistryUpdated(address contractRegistry)
func (_Staking *StakingFilterer) WatchContractRegistryUpdated(opts *bind.WatchOpts, sink chan<- *StakingContractRegistryUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ContractRegistryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractRegistryUpdated)
				if err := _Staking.contract.UnpackLog(event, "ContractRegistryUpdated", log); err != nil {
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

// ParseContractRegistryUpdated is a log parse operation binding the contract event 0xb3a6e7c81ebdb0cf9bf28d5ddf6678ded8d73b44f2eee06ca3974dbb9f41ce7e.
//
// Solidity: event ContractRegistryUpdated(address contractRegistry)
func (_Staking *StakingFilterer) ParseContractRegistryUpdated(log types.Log) (*StakingContractRegistryUpdated, error) {
	event := new(StakingContractRegistryUpdated)
	if err := _Staking.contract.UnpackLog(event, "ContractRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
