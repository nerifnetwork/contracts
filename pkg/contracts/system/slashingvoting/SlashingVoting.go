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

// SlashingVotingMetaData contains all meta data concerning the SlashingVoting contract.
var SlashingVotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSlashingReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BannedWithReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"SlashedWithReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSlashingReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"VotedWithReason\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bannedValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bans\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bansByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"bansByReason\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"epochByBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"getBannedValidatorsByReason\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getBansByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorGetterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashingThresold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lashingEpochs\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"isBannedByReason\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"slashingProposalVoteCounts\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochPeriod\",\"type\":\"uint256\"}],\"name\":\"setEpochPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashingEpochs\",\"type\":\"uint256\"}],\"name\":\"setSlashingEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashingThresold\",\"type\":\"uint256\"}],\"name\":\"setSlashingThresold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"shouldShash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractISignerAddress\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashingEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashingThresold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorGetter\",\"outputs\":[{\"internalType\":\"contractValidatorGetter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voteCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"voteWithReason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"votingHashWithReason\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061217a806100206000396000f3fe608060405234801561001057600080fd5b50600436106101ef5760003560e01c80638aaf0dae1161010f578063c4b14d58116100a2578063f84da26e11610071578063f84da26e14610532578063faaa8a6414610545578063fbf9204d14610573578063fce011261461057c57600080fd5b8063c4b14d5814610479578063ce119a8a146104af578063ed2da0ac146104eb578063f03528e7146104fe57600080fd5b8063b5b7a184116100de578063b5b7a18414610437578063bb69ffcd14610440578063be271a0214610453578063c42127b41461046657600080fd5b80638aaf0dae146103d05780639386775a146103e3578063abf410e514610411578063b1a5d12d1461042457600080fd5b806369495ef5116101875780637667180811610156578063766718081461036657806378a5c2061461036e578063807896d51461039d57806381d0e37b146103b057600080fd5b806369495ef51461030d5780636acd4f96146103205780636b5f444c146103335780636e6bb97c1461034657600080fd5b80633a9783f3116101c35780633a9783f31461025a5780633dad9ca914610289578063561ff9a9146102bc5780635c211f88146102e257600080fd5b8062708bb6146101f4578063013cf08b1461021057806303e7f672146102325780631f4f7d2914610247575b600080fd5b6101fd60055481565b6040519081526020015b60405180910390f35b61022361021e3660046118dc565b6105a7565b6040516102079392919061193b565b61024561024036600461199b565b61066f565b005b610245610255366004611a72565b610c75565b61027c60405180604001604052806003815260200162646b6760e81b81525081565b6040516102079190611b1a565b6102ac6102973660046118dc565b60096020526000908152604090205460ff1681565b6040519015158152602001610207565b61027c604051806040016040528060078152602001667374616b696e6760c81b81525081565b6001546102f5906001600160a01b031681565b6040516001600160a01b039091168152602001610207565b6000546102f5906001600160a01b031681565b6101fd61032e36600461199b565b610de1565b6102456103413660046118dc565b610e1b565b610359610354366004611b34565b610eb3565b6040516102079190611b4f565b6101fd610f62565b61027c6040518060400160405280601081526020016f737570706f727465642d746f6b656e7360801b81525081565b6102456103ab3660046118dc565b610f72565b6101fd6103be3660046118dc565b60086020526000908152604090205481565b6102456103de3660046118dc565b6113de565b6102ac6103f1366004611b9c565b600760209081526000928352604080842090915290825290205460ff1681565b6003546102f5906001600160a01b031681565b610245610432366004611bcc565b611476565b6101fd60045481565b6102ac61044e366004611b9c565b611608565b6102f5610461366004611c33565b61169a565b6101fd6104743660046118dc565b6116df565b6101fd610487366004611b9c565b6000918252600b602090815260408084206001600160a01b0393909316845291905290205490565b61027c6040518060400160405280601881526020017f7265776172642d646973747269627574696f6e2d706f6f6c000000000000000081525081565b6102ac6104f9366004611c68565b6116ef565b6102ac61050c366004611c9d565b600a60209081526000938452604080852082529284528284209052825290205460ff1681565b6102456105403660046118dc565b61176e565b61027c6040518060400160405280600f81526020016e736c617368696e672d766f74696e6760881b81525081565b6101fd60065481565b6101fd61058a366004611b9c565b600b60209081526000928352604080842090915290825290205481565b600281815481106105b757600080fd5b6000918252602090912060049091020180546001820180546001600160a01b039092169350906105e690611cdb565b80601f016020809104026020016040519081016040528092919081815260200182805461061290611cdb565b801561065f5780601f106106345761010080835404028352916020019161065f565b820191906000526020600020905b81548152906001019060200180831161064257829003601f168201915b5050505050908060030154905083565b6000546040516310ab556b60e21b81523360048201526001600160a01b03909116906342ad55ac90602401602060405180830381865afa1580156106b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106db9190611d15565b6107005760405162461bcd60e51b81526004016106f790611d37565b60405180910390fd5b600061070a611806565b90506000610716611893565b9050600061072687878787610de1565b6040516310ab556b60e21b81526001600160a01b038981166004830152919250908416906342ad55ac90602401602060405180830381865afa158015610770573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107949190611d15565b15156001146107b55760405162461bcd60e51b81526004016106f790611d6c565b60008181526009602052604090205460ff16156108285760405162461bcd60e51b815260206004820152602b60248201527f536c617368696e67566f74696e673a2076616c696461746f7220697320616c7260448201526a1958591e4818985b9b995960aa1b60648201526084016106f7565b600081815260076020908152604080832033845290915290205460ff16156108b85760405162461bcd60e51b815260206004820152603e60248201527f536c617368696e67566f74696e673a20766f74657220697320616c726561647960448201527f20766f74656420616761696e737420676976656e2076616c696461746f72000060648201526084016106f7565b60008181526007602090815260408083203384528252808320805460ff19166001179055838352600890915281208054916108f283611dd0565b91905055507f42ff2b7c8c611c525511dd04c1ee3cae48f313329b255979b94fc87a0f3a4a2633888860405161092a93929190611e21565b60405180910390a1600061093c610f62565b90506000846001600160a01b031663b7ab4db56040518163ffffffff1660e01b8152600401600060405180830381865afa15801561097e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109a69190810190611e46565b9050600281516109b69190611ef8565b6109c1906001611f1a565b60008481526008602052604090205410610bb8576000838152600960209081526040808320805460ff19166001908117909155858452600a83528184206001600160a01b038e168552909252822090918a6004811115610a2357610a23611de9565b6004811115610a3457610a34611de9565b815260208082019290925260409081016000908120805460ff191694151594909417909355848352600b82528083206001600160a01b038d1684529091528120805491610a8083611dd0565b90915550506000828152600c6020526040812090896004811115610aa657610aa6611de9565b6004811115610ab757610ab7611de9565b815260208082019290925260409081016000908120805460018101825590825292902090910180546001600160a01b0319166001600160a01b038c16179055517f370cc65d87ae599d8b7dd97c0d1f33291921e8cc3652f51fe8db7d974e567c2390610b26908b908b90611f2d565b60405180910390a16001886004811115610b4257610b42611de9565b1480610b5f57506002886004811115610b5d57610b5d611de9565b145b15610bb857836001600160a01b031663b32805c36040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610b9f57600080fd5b505af1158015610bb3573d6000803e3d6000fd5b505050505b610bc2828a611608565b15610c6a57610bcf611806565b60405163c96be4cb60e01b81526001600160a01b038b81166004830152919091169063c96be4cb90602401600060405180830381600087803b158015610c1457600080fd5b505af1158015610c28573d6000803e3d6000fd5b50506040516001600160a01b038c1681527f56848e1c0571e37fab91475a03170418e6a2956e066666c8038484240ea547099250602001905060405180910390a15b505050505050505050565b6000546040516310ab556b60e21b81523360048201526001600160a01b03909116906342ad55ac90602401602060405180830381865afa158015610cbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce19190611d15565b610cfd5760405162461bcd60e51b81526004016106f790611d37565b600280546001810182556000919091526004027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace810180546001600160a01b0385166001600160a01b0319909116178155907f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf01610d7b8382611f99565b50600254600090610d8e90600190612059565b604080518281526001600160a01b03871660208201529192507fcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358910160405180910390a1610ddb81610f72565b50505050565b600084848484604051602001610dfa949392919061206c565b6040516020818303038152906040528051906020012090505b949350505050565b60015460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa158015610e64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8891906120c3565b6001600160a01b031614610eae5760405162461bcd60e51b81526004016106f7906120e0565b600455565b6060600c6000610ec1610f62565b81526020019081526020016000206000836004811115610ee357610ee3611de9565b6004811115610ef457610ef4611de9565b8152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610f5657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610f38575b50505050509050919050565b6000610f6d436116df565b905090565b6000546040516310ab556b60e21b81523360048201526001600160a01b03909116906342ad55ac90602401602060405180830381865afa158015610fba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fde9190611d15565b610ffa5760405162461bcd60e51b81526004016106f790611d37565b6000611004611806565b60025490915082106110685760405162461bcd60e51b815260206004820152602760248201527f536c617368696e67566f74696e673a2070726f706f73616c20646f65736e27746044820152662065786973742160c81b60648201526084016106f7565b60006002838154811061107d5761107d612117565b6000918252602090912060049182020180546040516310ab556b60e21b81526001600160a01b03918216938101939093529092508316906342ad55ac90602401602060405180830381865afa1580156110da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110fe9190611d15565b151560011461111f5760405162461bcd60e51b81526004016106f790611d6c565b6002838154811061113257611132612117565b600091825260208083203384526002600490930201919091019052604090205460ff16156111bd5760405162461bcd60e51b815260206004820152603260248201527f536c617368696e67566f74696e673a20796f7520616c726561647920766f746560448201527119081a5b881d1a1a5cc81c1c9bdc1bdcd85b60721b60648201526084016106f7565b6001600284815481106111d2576111d2612117565b600091825260208083203384526002600490930201820190526040909120805492151560ff199093169290921790915580548490811061121457611214612117565b60009182526020822060036004909202010180549161123283611dd0565b91905055506000826001600160a01b031663b7ab4db56040518163ffffffff1660e01b8152600401600060405180830381865afa158015611277573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261129f9190810190611e46565b9050600281516112af9190611ef8565b6112ba906001611f1a565b600285815481106112cd576112cd612117565b9060005260206000209060040201600301541061138f576112ec611806565b825460405163c96be4cb60e01b81526001600160a01b03918216600482015291169063c96be4cb90602401600060405180830381600087803b15801561133157600080fd5b505af1158015611345573d6000803e3d6000fd5b50508354604080518881526001600160a01b0390921660208301527f9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24935001905060405180910390a15b8154604080518681526001600160a01b0390921660208301523382820152517fd88f7b9f64fb7ba069d57fe9cedb25c7827ee4f7c67c7f0967f6a25bd6d0c53c9181900360600190a150505050565b60015460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa158015611427573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144b91906120c3565b6001600160a01b0316146114715760405162461bcd60e51b81526004016106f7906120e0565b600655565b600154600160a81b900460ff161580801561149c575060018054600160a01b900460ff16105b806114bc5750303b1580156114bc575060018054600160a01b900460ff16145b61151f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106f7565b6001805460ff60a01b1916600160a01b179055801561154c576001805460ff60a81b1916600160a81b1790555b600180546001600160a01b0319166001600160a01b038916179055600080546001600160a01b0319166001600160a01b03881617905561158b85610e1b565b6115948461176e565b61159d836113de565b600380546001600160a01b0319166001600160a01b03841617905580156115ff576001805460ff60a81b191681556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b600060065483101561161c57506000611694565b6000835b60065461162d9086612059565b81111561168b576000818152600b602090815260408083206001600160a01b03881684529091528120549081900361166b5760009350505050611694565b6116758184611f1a565b92505080806116839061212d565b915050611620565b50600554111590505b92915050565b600c60205282600052604060002060205281600052604060002081815481106116c257600080fd5b6000918252602090912001546001600160a01b0316925083915050565b6000600454826116949190611ef8565b6000600a60006116fd610f62565b81526020019081526020016000206000846001600160a01b03166001600160a01b03168152602001908152602001600020600083600481111561174257611742611de9565b600481111561175357611753611de9565b815260208101919091526040016000205460ff169392505050565b60015460408051630d14b70160e11b8152905133926001600160a01b031691631a296e029160048083019260209291908290030181865afa1580156117b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117db91906120c3565b6001600160a01b0316146118015760405162461bcd60e51b81526004016106f7906120e0565b600555565b60035460408051808201825260078152667374616b696e6760c81b60208201529051633581777360e01b81526000926001600160a01b0316916335817773916118529190600401611b1a565b602060405180830381865afa15801561186f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6d91906120c3565b6003805460408051808201825292835262646b6760e81b602084015251633581777360e01b81526000926001600160a01b03909216916335817773916118529190600401611b1a565b6000602082840312156118ee57600080fd5b5035919050565b6000815180845260005b8181101561191b576020818501810151868301820152016118ff565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160a01b038416815260606020820181905260009061195f908301856118f5565b9050826040830152949350505050565b6001600160a01b038116811461198457600080fd5b50565b80356005811061199657600080fd5b919050565b600080600080606085870312156119b157600080fd5b84356119bc8161196f565b93506119ca60208601611987565b9250604085013567ffffffffffffffff808211156119e757600080fd5b818701915087601f8301126119fb57600080fd5b813581811115611a0a57600080fd5b886020828501011115611a1c57600080fd5b95989497505060200194505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611a6a57611a6a611a2b565b604052919050565b60008060408385031215611a8557600080fd5b8235611a908161196f565b915060208381013567ffffffffffffffff80821115611aae57600080fd5b818601915086601f830112611ac257600080fd5b813581811115611ad457611ad4611a2b565b611ae6601f8201601f19168501611a41565b91508082528784828501011115611afc57600080fd5b80848401858401376000848284010152508093505050509250929050565b602081526000611b2d60208301846118f5565b9392505050565b600060208284031215611b4657600080fd5b611b2d82611987565b6020808252825182820181905260009190848201906040850190845b81811015611b905783516001600160a01b031683529284019291840191600101611b6b565b50909695505050505050565b60008060408385031215611baf57600080fd5b823591506020830135611bc18161196f565b809150509250929050565b60008060008060008060c08789031215611be557600080fd5b8635611bf08161196f565b95506020870135611c008161196f565b945060408701359350606087013592506080870135915060a0870135611c258161196f565b809150509295509295509295565b600080600060608486031215611c4857600080fd5b83359250611c5860208501611987565b9150604084013590509250925092565b60008060408385031215611c7b57600080fd5b8235611c868161196f565b9150611c9460208401611987565b90509250929050565b600080600060608486031215611cb257600080fd5b833592506020840135611cc48161196f565b9150611cd260408501611987565b90509250925092565b600181811c90821680611cef57607f821691505b602082108103611d0f57634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215611d2757600080fd5b81518015158114611b2d57600080fd5b6020808252818101527f56616c696461746f724f776e61626c653a206f6e6c792076616c696461746f72604082015260600190565b6020808252602e908201527f536c617368696e67566f74696e673a20746172676574206973206e6f7420616360408201526d3a34bb32903b30b634b230ba37b960911b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600060018201611de257611de2611dba565b5060010190565b634e487b7160e01b600052602160045260246000fd5b60058110611e1d57634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b0384811682528316602082015260608101610e136040830184611dff565b60006020808385031215611e5957600080fd5b825167ffffffffffffffff80821115611e7157600080fd5b818501915085601f830112611e8557600080fd5b815181811115611e9757611e97611a2b565b8060051b9150611ea8848301611a41565b8181529183018401918481019088841115611ec257600080fd5b938501935b83851015611eec5784519250611edc8361196f565b8282529385019390850190611ec7565b98975050505050505050565b600082611f1557634e487b7160e01b600052601260045260246000fd5b500490565b8082018082111561169457611694611dba565b6001600160a01b038316815260408101611b2d6020830184611dff565b601f821115611f9457600081815260208120601f850160051c81016020861015611f715750805b601f850160051c820191505b81811015611f9057828155600101611f7d565b5050505b505050565b815167ffffffffffffffff811115611fb357611fb3611a2b565b611fc781611fc18454611cdb565b84611f4a565b602080601f831160018114611ffc5760008415611fe45750858301515b600019600386901b1c1916600185901b178555611f90565b600085815260208120601f198616915b8281101561202b5788860151825594840194600190910190840161200c565b50858210156120495787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b8181038181111561169457611694611dba565b6bffffffffffffffffffffffff198560601b1681526000600585106120a157634e487b7160e01b600052602160045260246000fd5b8460f81b60148301528284601584013750600091016015019081529392505050565b6000602082840312156120d557600080fd5b8151611b2d8161196f565b6020808252601a908201527f5369676e65724f776e61626c653a206f6e6c79207369676e6572000000000000604082015260600190565b634e487b7160e01b600052603260045260246000fd5b60008161213c5761213c611dba565b50600019019056fea2646970667358221220786aea5c282496022d4e63e8bf0ad45c37baf14db2be4cf070a72406700bbb7c64736f6c63430008120033",
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

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) DKGKEY() (string, error) {
	return _SlashingVoting.Contract.DKGKEY(&_SlashingVoting.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) DKGKEY() (string, error) {
	return _SlashingVoting.Contract.DKGKEY(&_SlashingVoting.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) REWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _SlashingVoting.Contract.REWARDDISTRIBUTIONPOOLKEY(&_SlashingVoting.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _SlashingVoting.Contract.REWARDDISTRIBUTIONPOOLKEY(&_SlashingVoting.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) SLASHINGVOTINGKEY() (string, error) {
	return _SlashingVoting.Contract.SLASHINGVOTINGKEY(&_SlashingVoting.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _SlashingVoting.Contract.SLASHINGVOTINGKEY(&_SlashingVoting.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) STAKINGKEY() (string, error) {
	return _SlashingVoting.Contract.STAKINGKEY(&_SlashingVoting.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) STAKINGKEY() (string, error) {
	return _SlashingVoting.Contract.STAKINGKEY(&_SlashingVoting.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _SlashingVoting.Contract.SUPPORTEDTOKENSKEY(&_SlashingVoting.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _SlashingVoting.Contract.SUPPORTEDTOKENSKEY(&_SlashingVoting.CallOpts)
}

// BannedValidators is a free data retrieval call binding the contract method 0xbe271a02.
//
// Solidity: function bannedValidators(uint256 , uint8 , uint256 ) view returns(address)
func (_SlashingVoting *SlashingVotingCaller) BannedValidators(opts *bind.CallOpts, arg0 *big.Int, arg1 uint8, arg2 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "bannedValidators", arg0, arg1, arg2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BannedValidators is a free data retrieval call binding the contract method 0xbe271a02.
//
// Solidity: function bannedValidators(uint256 , uint8 , uint256 ) view returns(address)
func (_SlashingVoting *SlashingVotingSession) BannedValidators(arg0 *big.Int, arg1 uint8, arg2 *big.Int) (common.Address, error) {
	return _SlashingVoting.Contract.BannedValidators(&_SlashingVoting.CallOpts, arg0, arg1, arg2)
}

// BannedValidators is a free data retrieval call binding the contract method 0xbe271a02.
//
// Solidity: function bannedValidators(uint256 , uint8 , uint256 ) view returns(address)
func (_SlashingVoting *SlashingVotingCallerSession) BannedValidators(arg0 *big.Int, arg1 uint8, arg2 *big.Int) (common.Address, error) {
	return _SlashingVoting.Contract.BannedValidators(&_SlashingVoting.CallOpts, arg0, arg1, arg2)
}

// Bans is a free data retrieval call binding the contract method 0x3dad9ca9.
//
// Solidity: function bans(bytes32 ) view returns(bool)
func (_SlashingVoting *SlashingVotingCaller) Bans(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "bans", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bans is a free data retrieval call binding the contract method 0x3dad9ca9.
//
// Solidity: function bans(bytes32 ) view returns(bool)
func (_SlashingVoting *SlashingVotingSession) Bans(arg0 [32]byte) (bool, error) {
	return _SlashingVoting.Contract.Bans(&_SlashingVoting.CallOpts, arg0)
}

// Bans is a free data retrieval call binding the contract method 0x3dad9ca9.
//
// Solidity: function bans(bytes32 ) view returns(bool)
func (_SlashingVoting *SlashingVotingCallerSession) Bans(arg0 [32]byte) (bool, error) {
	return _SlashingVoting.Contract.Bans(&_SlashingVoting.CallOpts, arg0)
}

// BansByEpoch is a free data retrieval call binding the contract method 0xfce01126.
//
// Solidity: function bansByEpoch(uint256 , address ) view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) BansByEpoch(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "bansByEpoch", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BansByEpoch is a free data retrieval call binding the contract method 0xfce01126.
//
// Solidity: function bansByEpoch(uint256 , address ) view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) BansByEpoch(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SlashingVoting.Contract.BansByEpoch(&_SlashingVoting.CallOpts, arg0, arg1)
}

// BansByEpoch is a free data retrieval call binding the contract method 0xfce01126.
//
// Solidity: function bansByEpoch(uint256 , address ) view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) BansByEpoch(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _SlashingVoting.Contract.BansByEpoch(&_SlashingVoting.CallOpts, arg0, arg1)
}

// BansByReason is a free data retrieval call binding the contract method 0xf03528e7.
//
// Solidity: function bansByReason(uint256 , address , uint8 ) view returns(bool)
func (_SlashingVoting *SlashingVotingCaller) BansByReason(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 uint8) (bool, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "bansByReason", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BansByReason is a free data retrieval call binding the contract method 0xf03528e7.
//
// Solidity: function bansByReason(uint256 , address , uint8 ) view returns(bool)
func (_SlashingVoting *SlashingVotingSession) BansByReason(arg0 *big.Int, arg1 common.Address, arg2 uint8) (bool, error) {
	return _SlashingVoting.Contract.BansByReason(&_SlashingVoting.CallOpts, arg0, arg1, arg2)
}

// BansByReason is a free data retrieval call binding the contract method 0xf03528e7.
//
// Solidity: function bansByReason(uint256 , address , uint8 ) view returns(bool)
func (_SlashingVoting *SlashingVotingCallerSession) BansByReason(arg0 *big.Int, arg1 common.Address, arg2 uint8) (bool, error) {
	return _SlashingVoting.Contract.BansByReason(&_SlashingVoting.CallOpts, arg0, arg1, arg2)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_SlashingVoting *SlashingVotingCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_SlashingVoting *SlashingVotingSession) ContractRegistry() (common.Address, error) {
	return _SlashingVoting.Contract.ContractRegistry(&_SlashingVoting.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_SlashingVoting *SlashingVotingCallerSession) ContractRegistry() (common.Address, error) {
	return _SlashingVoting.Contract.ContractRegistry(&_SlashingVoting.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) CurrentEpoch() (*big.Int, error) {
	return _SlashingVoting.Contract.CurrentEpoch(&_SlashingVoting.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) CurrentEpoch() (*big.Int, error) {
	return _SlashingVoting.Contract.CurrentEpoch(&_SlashingVoting.CallOpts)
}

// EpochByBlock is a free data retrieval call binding the contract method 0xc42127b4.
//
// Solidity: function epochByBlock(uint256 _blockNumber) view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) EpochByBlock(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "epochByBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochByBlock is a free data retrieval call binding the contract method 0xc42127b4.
//
// Solidity: function epochByBlock(uint256 _blockNumber) view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) EpochByBlock(_blockNumber *big.Int) (*big.Int, error) {
	return _SlashingVoting.Contract.EpochByBlock(&_SlashingVoting.CallOpts, _blockNumber)
}

// EpochByBlock is a free data retrieval call binding the contract method 0xc42127b4.
//
// Solidity: function epochByBlock(uint256 _blockNumber) view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) EpochByBlock(_blockNumber *big.Int) (*big.Int, error) {
	return _SlashingVoting.Contract.EpochByBlock(&_SlashingVoting.CallOpts, _blockNumber)
}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) EpochPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "epochPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) EpochPeriod() (*big.Int, error) {
	return _SlashingVoting.Contract.EpochPeriod(&_SlashingVoting.CallOpts)
}

// EpochPeriod is a free data retrieval call binding the contract method 0xb5b7a184.
//
// Solidity: function epochPeriod() view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) EpochPeriod() (*big.Int, error) {
	return _SlashingVoting.Contract.EpochPeriod(&_SlashingVoting.CallOpts)
}

// GetBannedValidatorsByReason is a free data retrieval call binding the contract method 0x6e6bb97c.
//
// Solidity: function getBannedValidatorsByReason(uint8 _reason) view returns(address[])
func (_SlashingVoting *SlashingVotingCaller) GetBannedValidatorsByReason(opts *bind.CallOpts, _reason uint8) ([]common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "getBannedValidatorsByReason", _reason)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBannedValidatorsByReason is a free data retrieval call binding the contract method 0x6e6bb97c.
//
// Solidity: function getBannedValidatorsByReason(uint8 _reason) view returns(address[])
func (_SlashingVoting *SlashingVotingSession) GetBannedValidatorsByReason(_reason uint8) ([]common.Address, error) {
	return _SlashingVoting.Contract.GetBannedValidatorsByReason(&_SlashingVoting.CallOpts, _reason)
}

// GetBannedValidatorsByReason is a free data retrieval call binding the contract method 0x6e6bb97c.
//
// Solidity: function getBannedValidatorsByReason(uint8 _reason) view returns(address[])
func (_SlashingVoting *SlashingVotingCallerSession) GetBannedValidatorsByReason(_reason uint8) ([]common.Address, error) {
	return _SlashingVoting.Contract.GetBannedValidatorsByReason(&_SlashingVoting.CallOpts, _reason)
}

// GetBansByEpoch is a free data retrieval call binding the contract method 0xc4b14d58.
//
// Solidity: function getBansByEpoch(uint256 _epoch, address _validator) view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) GetBansByEpoch(opts *bind.CallOpts, _epoch *big.Int, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "getBansByEpoch", _epoch, _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBansByEpoch is a free data retrieval call binding the contract method 0xc4b14d58.
//
// Solidity: function getBansByEpoch(uint256 _epoch, address _validator) view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) GetBansByEpoch(_epoch *big.Int, _validator common.Address) (*big.Int, error) {
	return _SlashingVoting.Contract.GetBansByEpoch(&_SlashingVoting.CallOpts, _epoch, _validator)
}

// GetBansByEpoch is a free data retrieval call binding the contract method 0xc4b14d58.
//
// Solidity: function getBansByEpoch(uint256 _epoch, address _validator) view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) GetBansByEpoch(_epoch *big.Int, _validator common.Address) (*big.Int, error) {
	return _SlashingVoting.Contract.GetBansByEpoch(&_SlashingVoting.CallOpts, _epoch, _validator)
}

// IsBannedByReason is a free data retrieval call binding the contract method 0xed2da0ac.
//
// Solidity: function isBannedByReason(address _validator, uint8 _reason) view returns(bool)
func (_SlashingVoting *SlashingVotingCaller) IsBannedByReason(opts *bind.CallOpts, _validator common.Address, _reason uint8) (bool, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "isBannedByReason", _validator, _reason)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBannedByReason is a free data retrieval call binding the contract method 0xed2da0ac.
//
// Solidity: function isBannedByReason(address _validator, uint8 _reason) view returns(bool)
func (_SlashingVoting *SlashingVotingSession) IsBannedByReason(_validator common.Address, _reason uint8) (bool, error) {
	return _SlashingVoting.Contract.IsBannedByReason(&_SlashingVoting.CallOpts, _validator, _reason)
}

// IsBannedByReason is a free data retrieval call binding the contract method 0xed2da0ac.
//
// Solidity: function isBannedByReason(address _validator, uint8 _reason) view returns(bool)
func (_SlashingVoting *SlashingVotingCallerSession) IsBannedByReason(_validator common.Address, _reason uint8) (bool, error) {
	return _SlashingVoting.Contract.IsBannedByReason(&_SlashingVoting.CallOpts, _validator, _reason)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address validator, string reason, uint256 slashingProposalVoteCounts)
func (_SlashingVoting *SlashingVotingCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Validator                  common.Address
	Reason                     string
	SlashingProposalVoteCounts *big.Int
}, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Validator                  common.Address
		Reason                     string
		SlashingProposalVoteCounts *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Reason = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SlashingProposalVoteCounts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address validator, string reason, uint256 slashingProposalVoteCounts)
func (_SlashingVoting *SlashingVotingSession) Proposals(arg0 *big.Int) (struct {
	Validator                  common.Address
	Reason                     string
	SlashingProposalVoteCounts *big.Int
}, error) {
	return _SlashingVoting.Contract.Proposals(&_SlashingVoting.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(address validator, string reason, uint256 slashingProposalVoteCounts)
func (_SlashingVoting *SlashingVotingCallerSession) Proposals(arg0 *big.Int) (struct {
	Validator                  common.Address
	Reason                     string
	SlashingProposalVoteCounts *big.Int
}, error) {
	return _SlashingVoting.Contract.Proposals(&_SlashingVoting.CallOpts, arg0)
}

// ShouldShash is a free data retrieval call binding the contract method 0xbb69ffcd.
//
// Solidity: function shouldShash(uint256 _epoch, address _validator) view returns(bool)
func (_SlashingVoting *SlashingVotingCaller) ShouldShash(opts *bind.CallOpts, _epoch *big.Int, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "shouldShash", _epoch, _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldShash is a free data retrieval call binding the contract method 0xbb69ffcd.
//
// Solidity: function shouldShash(uint256 _epoch, address _validator) view returns(bool)
func (_SlashingVoting *SlashingVotingSession) ShouldShash(_epoch *big.Int, _validator common.Address) (bool, error) {
	return _SlashingVoting.Contract.ShouldShash(&_SlashingVoting.CallOpts, _epoch, _validator)
}

// ShouldShash is a free data retrieval call binding the contract method 0xbb69ffcd.
//
// Solidity: function shouldShash(uint256 _epoch, address _validator) view returns(bool)
func (_SlashingVoting *SlashingVotingCallerSession) ShouldShash(_epoch *big.Int, _validator common.Address) (bool, error) {
	return _SlashingVoting.Contract.ShouldShash(&_SlashingVoting.CallOpts, _epoch, _validator)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SlashingVoting *SlashingVotingSession) SignerGetter() (common.Address, error) {
	return _SlashingVoting.Contract.SignerGetter(&_SlashingVoting.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCallerSession) SignerGetter() (common.Address, error) {
	return _SlashingVoting.Contract.SignerGetter(&_SlashingVoting.CallOpts)
}

// SlashingEpochs is a free data retrieval call binding the contract method 0xfbf9204d.
//
// Solidity: function slashingEpochs() view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) SlashingEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "slashingEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingEpochs is a free data retrieval call binding the contract method 0xfbf9204d.
//
// Solidity: function slashingEpochs() view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) SlashingEpochs() (*big.Int, error) {
	return _SlashingVoting.Contract.SlashingEpochs(&_SlashingVoting.CallOpts)
}

// SlashingEpochs is a free data retrieval call binding the contract method 0xfbf9204d.
//
// Solidity: function slashingEpochs() view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) SlashingEpochs() (*big.Int, error) {
	return _SlashingVoting.Contract.SlashingEpochs(&_SlashingVoting.CallOpts)
}

// SlashingThresold is a free data retrieval call binding the contract method 0x00708bb6.
//
// Solidity: function slashingThresold() view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) SlashingThresold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "slashingThresold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashingThresold is a free data retrieval call binding the contract method 0x00708bb6.
//
// Solidity: function slashingThresold() view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) SlashingThresold() (*big.Int, error) {
	return _SlashingVoting.Contract.SlashingThresold(&_SlashingVoting.CallOpts)
}

// SlashingThresold is a free data retrieval call binding the contract method 0x00708bb6.
//
// Solidity: function slashingThresold() view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) SlashingThresold() (*big.Int, error) {
	return _SlashingVoting.Contract.SlashingThresold(&_SlashingVoting.CallOpts)
}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCaller) ValidatorGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "validatorGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_SlashingVoting *SlashingVotingSession) ValidatorGetter() (common.Address, error) {
	return _SlashingVoting.Contract.ValidatorGetter(&_SlashingVoting.CallOpts)
}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCallerSession) ValidatorGetter() (common.Address, error) {
	return _SlashingVoting.Contract.ValidatorGetter(&_SlashingVoting.CallOpts)
}

// VoteCounts is a free data retrieval call binding the contract method 0x81d0e37b.
//
// Solidity: function voteCounts(bytes32 ) view returns(uint256)
func (_SlashingVoting *SlashingVotingCaller) VoteCounts(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "voteCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteCounts is a free data retrieval call binding the contract method 0x81d0e37b.
//
// Solidity: function voteCounts(bytes32 ) view returns(uint256)
func (_SlashingVoting *SlashingVotingSession) VoteCounts(arg0 [32]byte) (*big.Int, error) {
	return _SlashingVoting.Contract.VoteCounts(&_SlashingVoting.CallOpts, arg0)
}

// VoteCounts is a free data retrieval call binding the contract method 0x81d0e37b.
//
// Solidity: function voteCounts(bytes32 ) view returns(uint256)
func (_SlashingVoting *SlashingVotingCallerSession) VoteCounts(arg0 [32]byte) (*big.Int, error) {
	return _SlashingVoting.Contract.VoteCounts(&_SlashingVoting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x9386775a.
//
// Solidity: function votes(bytes32 , address ) view returns(bool)
func (_SlashingVoting *SlashingVotingCaller) Votes(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "votes", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Votes is a free data retrieval call binding the contract method 0x9386775a.
//
// Solidity: function votes(bytes32 , address ) view returns(bool)
func (_SlashingVoting *SlashingVotingSession) Votes(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _SlashingVoting.Contract.Votes(&_SlashingVoting.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0x9386775a.
//
// Solidity: function votes(bytes32 , address ) view returns(bool)
func (_SlashingVoting *SlashingVotingCallerSession) Votes(arg0 [32]byte, arg1 common.Address) (bool, error) {
	return _SlashingVoting.Contract.Votes(&_SlashingVoting.CallOpts, arg0, arg1)
}

// VotingHashWithReason is a free data retrieval call binding the contract method 0x6acd4f96.
//
// Solidity: function votingHashWithReason(address _validator, uint8 _reason, bytes _nonce) pure returns(bytes32)
func (_SlashingVoting *SlashingVotingCaller) VotingHashWithReason(opts *bind.CallOpts, _validator common.Address, _reason uint8, _nonce []byte) ([32]byte, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "votingHashWithReason", _validator, _reason, _nonce)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VotingHashWithReason is a free data retrieval call binding the contract method 0x6acd4f96.
//
// Solidity: function votingHashWithReason(address _validator, uint8 _reason, bytes _nonce) pure returns(bytes32)
func (_SlashingVoting *SlashingVotingSession) VotingHashWithReason(_validator common.Address, _reason uint8, _nonce []byte) ([32]byte, error) {
	return _SlashingVoting.Contract.VotingHashWithReason(&_SlashingVoting.CallOpts, _validator, _reason, _nonce)
}

// VotingHashWithReason is a free data retrieval call binding the contract method 0x6acd4f96.
//
// Solidity: function votingHashWithReason(address _validator, uint8 _reason, bytes _nonce) pure returns(bytes32)
func (_SlashingVoting *SlashingVotingCallerSession) VotingHashWithReason(_validator common.Address, _reason uint8, _nonce []byte) ([32]byte, error) {
	return _SlashingVoting.Contract.VotingHashWithReason(&_SlashingVoting.CallOpts, _validator, _reason, _nonce)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address _validator, string _reason) returns()
func (_SlashingVoting *SlashingVotingTransactor) CreateProposal(opts *bind.TransactOpts, _validator common.Address, _reason string) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "createProposal", _validator, _reason)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address _validator, string _reason) returns()
func (_SlashingVoting *SlashingVotingSession) CreateProposal(_validator common.Address, _reason string) (*types.Transaction, error) {
	return _SlashingVoting.Contract.CreateProposal(&_SlashingVoting.TransactOpts, _validator, _reason)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x1f4f7d29.
//
// Solidity: function createProposal(address _validator, string _reason) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) CreateProposal(_validator common.Address, _reason string) (*types.Transaction, error) {
	return _SlashingVoting.Contract.CreateProposal(&_SlashingVoting.TransactOpts, _validator, _reason)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a5d12d.
//
// Solidity: function initialize(address _signerGetterAddress, address _validatorGetterAddress, uint256 _epochPeriod, uint256 _slashingThresold, uint256 _lashingEpochs, address _contractRegistry) returns()
func (_SlashingVoting *SlashingVotingTransactor) Initialize(opts *bind.TransactOpts, _signerGetterAddress common.Address, _validatorGetterAddress common.Address, _epochPeriod *big.Int, _slashingThresold *big.Int, _lashingEpochs *big.Int, _contractRegistry common.Address) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "initialize", _signerGetterAddress, _validatorGetterAddress, _epochPeriod, _slashingThresold, _lashingEpochs, _contractRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a5d12d.
//
// Solidity: function initialize(address _signerGetterAddress, address _validatorGetterAddress, uint256 _epochPeriod, uint256 _slashingThresold, uint256 _lashingEpochs, address _contractRegistry) returns()
func (_SlashingVoting *SlashingVotingSession) Initialize(_signerGetterAddress common.Address, _validatorGetterAddress common.Address, _epochPeriod *big.Int, _slashingThresold *big.Int, _lashingEpochs *big.Int, _contractRegistry common.Address) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _signerGetterAddress, _validatorGetterAddress, _epochPeriod, _slashingThresold, _lashingEpochs, _contractRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a5d12d.
//
// Solidity: function initialize(address _signerGetterAddress, address _validatorGetterAddress, uint256 _epochPeriod, uint256 _slashingThresold, uint256 _lashingEpochs, address _contractRegistry) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) Initialize(_signerGetterAddress common.Address, _validatorGetterAddress common.Address, _epochPeriod *big.Int, _slashingThresold *big.Int, _lashingEpochs *big.Int, _contractRegistry common.Address) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _signerGetterAddress, _validatorGetterAddress, _epochPeriod, _slashingThresold, _lashingEpochs, _contractRegistry)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _epochPeriod) returns()
func (_SlashingVoting *SlashingVotingTransactor) SetEpochPeriod(opts *bind.TransactOpts, _epochPeriod *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "setEpochPeriod", _epochPeriod)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _epochPeriod) returns()
func (_SlashingVoting *SlashingVotingSession) SetEpochPeriod(_epochPeriod *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetEpochPeriod(&_SlashingVoting.TransactOpts, _epochPeriod)
}

// SetEpochPeriod is a paid mutator transaction binding the contract method 0x6b5f444c.
//
// Solidity: function setEpochPeriod(uint256 _epochPeriod) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) SetEpochPeriod(_epochPeriod *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetEpochPeriod(&_SlashingVoting.TransactOpts, _epochPeriod)
}

// SetSlashingEpochs is a paid mutator transaction binding the contract method 0x8aaf0dae.
//
// Solidity: function setSlashingEpochs(uint256 _slashingEpochs) returns()
func (_SlashingVoting *SlashingVotingTransactor) SetSlashingEpochs(opts *bind.TransactOpts, _slashingEpochs *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "setSlashingEpochs", _slashingEpochs)
}

// SetSlashingEpochs is a paid mutator transaction binding the contract method 0x8aaf0dae.
//
// Solidity: function setSlashingEpochs(uint256 _slashingEpochs) returns()
func (_SlashingVoting *SlashingVotingSession) SetSlashingEpochs(_slashingEpochs *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetSlashingEpochs(&_SlashingVoting.TransactOpts, _slashingEpochs)
}

// SetSlashingEpochs is a paid mutator transaction binding the contract method 0x8aaf0dae.
//
// Solidity: function setSlashingEpochs(uint256 _slashingEpochs) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) SetSlashingEpochs(_slashingEpochs *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetSlashingEpochs(&_SlashingVoting.TransactOpts, _slashingEpochs)
}

// SetSlashingThresold is a paid mutator transaction binding the contract method 0xf84da26e.
//
// Solidity: function setSlashingThresold(uint256 _slashingThresold) returns()
func (_SlashingVoting *SlashingVotingTransactor) SetSlashingThresold(opts *bind.TransactOpts, _slashingThresold *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "setSlashingThresold", _slashingThresold)
}

// SetSlashingThresold is a paid mutator transaction binding the contract method 0xf84da26e.
//
// Solidity: function setSlashingThresold(uint256 _slashingThresold) returns()
func (_SlashingVoting *SlashingVotingSession) SetSlashingThresold(_slashingThresold *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetSlashingThresold(&_SlashingVoting.TransactOpts, _slashingThresold)
}

// SetSlashingThresold is a paid mutator transaction binding the contract method 0xf84da26e.
//
// Solidity: function setSlashingThresold(uint256 _slashingThresold) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) SetSlashingThresold(_slashingThresold *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.SetSlashingThresold(&_SlashingVoting.TransactOpts, _slashingThresold)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x807896d5.
//
// Solidity: function voteProposal(uint256 _proposalId) returns()
func (_SlashingVoting *SlashingVotingTransactor) VoteProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "voteProposal", _proposalId)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x807896d5.
//
// Solidity: function voteProposal(uint256 _proposalId) returns()
func (_SlashingVoting *SlashingVotingSession) VoteProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.VoteProposal(&_SlashingVoting.TransactOpts, _proposalId)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x807896d5.
//
// Solidity: function voteProposal(uint256 _proposalId) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) VoteProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.VoteProposal(&_SlashingVoting.TransactOpts, _proposalId)
}

// VoteWithReason is a paid mutator transaction binding the contract method 0x03e7f672.
//
// Solidity: function voteWithReason(address _validator, uint8 _reason, bytes _nonce) returns()
func (_SlashingVoting *SlashingVotingTransactor) VoteWithReason(opts *bind.TransactOpts, _validator common.Address, _reason uint8, _nonce []byte) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "voteWithReason", _validator, _reason, _nonce)
}

// VoteWithReason is a paid mutator transaction binding the contract method 0x03e7f672.
//
// Solidity: function voteWithReason(address _validator, uint8 _reason, bytes _nonce) returns()
func (_SlashingVoting *SlashingVotingSession) VoteWithReason(_validator common.Address, _reason uint8, _nonce []byte) (*types.Transaction, error) {
	return _SlashingVoting.Contract.VoteWithReason(&_SlashingVoting.TransactOpts, _validator, _reason, _nonce)
}

// VoteWithReason is a paid mutator transaction binding the contract method 0x03e7f672.
//
// Solidity: function voteWithReason(address _validator, uint8 _reason, bytes _nonce) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) VoteWithReason(_validator common.Address, _reason uint8, _nonce []byte) (*types.Transaction, error) {
	return _SlashingVoting.Contract.VoteWithReason(&_SlashingVoting.TransactOpts, _validator, _reason, _nonce)
}

// SlashingVotingBannedWithReasonIterator is returned from FilterBannedWithReason and is used to iterate over the raw logs and unpacked data for BannedWithReason events raised by the SlashingVoting contract.
type SlashingVotingBannedWithReasonIterator struct {
	Event *SlashingVotingBannedWithReason // Event containing the contract specifics and raw log

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
func (it *SlashingVotingBannedWithReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingVotingBannedWithReason)
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
		it.Event = new(SlashingVotingBannedWithReason)
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
func (it *SlashingVotingBannedWithReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingVotingBannedWithReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingVotingBannedWithReason represents a BannedWithReason event raised by the SlashingVoting contract.
type SlashingVotingBannedWithReason struct {
	Validator common.Address
	Reason    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBannedWithReason is a free log retrieval operation binding the contract event 0x370cc65d87ae599d8b7dd97c0d1f33291921e8cc3652f51fe8db7d974e567c23.
//
// Solidity: event BannedWithReason(address validator, uint8 reason)
func (_SlashingVoting *SlashingVotingFilterer) FilterBannedWithReason(opts *bind.FilterOpts) (*SlashingVotingBannedWithReasonIterator, error) {

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "BannedWithReason")
	if err != nil {
		return nil, err
	}
	return &SlashingVotingBannedWithReasonIterator{contract: _SlashingVoting.contract, event: "BannedWithReason", logs: logs, sub: sub}, nil
}

// WatchBannedWithReason is a free log subscription operation binding the contract event 0x370cc65d87ae599d8b7dd97c0d1f33291921e8cc3652f51fe8db7d974e567c23.
//
// Solidity: event BannedWithReason(address validator, uint8 reason)
func (_SlashingVoting *SlashingVotingFilterer) WatchBannedWithReason(opts *bind.WatchOpts, sink chan<- *SlashingVotingBannedWithReason) (event.Subscription, error) {

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "BannedWithReason")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingVotingBannedWithReason)
				if err := _SlashingVoting.contract.UnpackLog(event, "BannedWithReason", log); err != nil {
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

// ParseBannedWithReason is a log parse operation binding the contract event 0x370cc65d87ae599d8b7dd97c0d1f33291921e8cc3652f51fe8db7d974e567c23.
//
// Solidity: event BannedWithReason(address validator, uint8 reason)
func (_SlashingVoting *SlashingVotingFilterer) ParseBannedWithReason(log types.Log) (*SlashingVotingBannedWithReason, error) {
	event := new(SlashingVotingBannedWithReason)
	if err := _SlashingVoting.contract.UnpackLog(event, "BannedWithReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
// Solidity: event ProposalCreated(uint256 proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*SlashingVotingProposalCreatedIterator, error) {

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &SlashingVotingProposalCreatedIterator{contract: _SlashingVoting.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0xcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358.
//
// Solidity: event ProposalCreated(uint256 proposalId, address validator)
func (_SlashingVoting *SlashingVotingFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *SlashingVotingProposalCreated) (event.Subscription, error) {

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "ProposalCreated")
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
// Solidity: event ProposalCreated(uint256 proposalId, address validator)
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
	Validator  common.Address
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0xd88f7b9f64fb7ba069d57fe9cedb25c7827ee4f7c67c7f0967f6a25bd6d0c53c.
//
// Solidity: event ProposalVoted(uint256 proposalId, address validator, address voter)
func (_SlashingVoting *SlashingVotingFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*SlashingVotingProposalVotedIterator, error) {

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &SlashingVotingProposalVotedIterator{contract: _SlashingVoting.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xd88f7b9f64fb7ba069d57fe9cedb25c7827ee4f7c67c7f0967f6a25bd6d0c53c.
//
// Solidity: event ProposalVoted(uint256 proposalId, address validator, address voter)
func (_SlashingVoting *SlashingVotingFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *SlashingVotingProposalVoted) (event.Subscription, error) {

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "ProposalVoted")
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

// ParseProposalVoted is a log parse operation binding the contract event 0xd88f7b9f64fb7ba069d57fe9cedb25c7827ee4f7c67c7f0967f6a25bd6d0c53c.
//
// Solidity: event ProposalVoted(uint256 proposalId, address validator, address voter)
func (_SlashingVoting *SlashingVotingFilterer) ParseProposalVoted(log types.Log) (*SlashingVotingProposalVoted, error) {
	event := new(SlashingVotingProposalVoted)
	if err := _SlashingVoting.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashingVotingSlashedWithReasonIterator is returned from FilterSlashedWithReason and is used to iterate over the raw logs and unpacked data for SlashedWithReason events raised by the SlashingVoting contract.
type SlashingVotingSlashedWithReasonIterator struct {
	Event *SlashingVotingSlashedWithReason // Event containing the contract specifics and raw log

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
func (it *SlashingVotingSlashedWithReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingVotingSlashedWithReason)
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
		it.Event = new(SlashingVotingSlashedWithReason)
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
func (it *SlashingVotingSlashedWithReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingVotingSlashedWithReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingVotingSlashedWithReason represents a SlashedWithReason event raised by the SlashingVoting contract.
type SlashingVotingSlashedWithReason struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSlashedWithReason is a free log retrieval operation binding the contract event 0x56848e1c0571e37fab91475a03170418e6a2956e066666c8038484240ea54709.
//
// Solidity: event SlashedWithReason(address validator)
func (_SlashingVoting *SlashingVotingFilterer) FilterSlashedWithReason(opts *bind.FilterOpts) (*SlashingVotingSlashedWithReasonIterator, error) {

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "SlashedWithReason")
	if err != nil {
		return nil, err
	}
	return &SlashingVotingSlashedWithReasonIterator{contract: _SlashingVoting.contract, event: "SlashedWithReason", logs: logs, sub: sub}, nil
}

// WatchSlashedWithReason is a free log subscription operation binding the contract event 0x56848e1c0571e37fab91475a03170418e6a2956e066666c8038484240ea54709.
//
// Solidity: event SlashedWithReason(address validator)
func (_SlashingVoting *SlashingVotingFilterer) WatchSlashedWithReason(opts *bind.WatchOpts, sink chan<- *SlashingVotingSlashedWithReason) (event.Subscription, error) {

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "SlashedWithReason")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingVotingSlashedWithReason)
				if err := _SlashingVoting.contract.UnpackLog(event, "SlashedWithReason", log); err != nil {
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

// ParseSlashedWithReason is a log parse operation binding the contract event 0x56848e1c0571e37fab91475a03170418e6a2956e066666c8038484240ea54709.
//
// Solidity: event SlashedWithReason(address validator)
func (_SlashingVoting *SlashingVotingFilterer) ParseSlashedWithReason(log types.Log) (*SlashingVotingSlashedWithReason, error) {
	event := new(SlashingVotingSlashedWithReason)
	if err := _SlashingVoting.contract.UnpackLog(event, "SlashedWithReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SlashingVotingVotedWithReasonIterator is returned from FilterVotedWithReason and is used to iterate over the raw logs and unpacked data for VotedWithReason events raised by the SlashingVoting contract.
type SlashingVotingVotedWithReasonIterator struct {
	Event *SlashingVotingVotedWithReason // Event containing the contract specifics and raw log

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
func (it *SlashingVotingVotedWithReasonIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SlashingVotingVotedWithReason)
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
		it.Event = new(SlashingVotingVotedWithReason)
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
func (it *SlashingVotingVotedWithReasonIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SlashingVotingVotedWithReasonIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SlashingVotingVotedWithReason represents a VotedWithReason event raised by the SlashingVoting contract.
type SlashingVotingVotedWithReason struct {
	Voter     common.Address
	Validator common.Address
	Reason    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVotedWithReason is a free log retrieval operation binding the contract event 0x42ff2b7c8c611c525511dd04c1ee3cae48f313329b255979b94fc87a0f3a4a26.
//
// Solidity: event VotedWithReason(address voter, address validator, uint8 reason)
func (_SlashingVoting *SlashingVotingFilterer) FilterVotedWithReason(opts *bind.FilterOpts) (*SlashingVotingVotedWithReasonIterator, error) {

	logs, sub, err := _SlashingVoting.contract.FilterLogs(opts, "VotedWithReason")
	if err != nil {
		return nil, err
	}
	return &SlashingVotingVotedWithReasonIterator{contract: _SlashingVoting.contract, event: "VotedWithReason", logs: logs, sub: sub}, nil
}

// WatchVotedWithReason is a free log subscription operation binding the contract event 0x42ff2b7c8c611c525511dd04c1ee3cae48f313329b255979b94fc87a0f3a4a26.
//
// Solidity: event VotedWithReason(address voter, address validator, uint8 reason)
func (_SlashingVoting *SlashingVotingFilterer) WatchVotedWithReason(opts *bind.WatchOpts, sink chan<- *SlashingVotingVotedWithReason) (event.Subscription, error) {

	logs, sub, err := _SlashingVoting.contract.WatchLogs(opts, "VotedWithReason")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SlashingVotingVotedWithReason)
				if err := _SlashingVoting.contract.UnpackLog(event, "VotedWithReason", log); err != nil {
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

// ParseVotedWithReason is a log parse operation binding the contract event 0x42ff2b7c8c611c525511dd04c1ee3cae48f313329b255979b94fc87a0f3a4a26.
//
// Solidity: event VotedWithReason(address voter, address validator, uint8 reason)
func (_SlashingVoting *SlashingVotingFilterer) ParseVotedWithReason(log types.Log) (*SlashingVotingVotedWithReason, error) {
	event := new(SlashingVotingVotedWithReason)
	if err := _SlashingVoting.contract.UnpackLog(event, "VotedWithReason", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
