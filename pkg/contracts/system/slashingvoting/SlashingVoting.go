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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSlashingReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BannedWithReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"SlashedWithReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSlashingReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"VotedWithReason\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bannedValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bans\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bansByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"bansByReason\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"epochByBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"getBannedValidatorsByReason\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getBansByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashingThresold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashingEpochs\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"isBannedByReason\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"slashingProposalVoteCounts\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochPeriod\",\"type\":\"uint256\"}],\"name\":\"setEpochPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashingEpochs\",\"type\":\"uint256\"}],\"name\":\"setSlashingEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashingThresold\",\"type\":\"uint256\"}],\"name\":\"setSlashingThresold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"shouldShash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashingEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashingThresold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voteCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"voteWithReason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"votingHashWithReason\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e22806100206000396000f3fe608060405234801561001057600080fd5b50600436106101a85760003560e01c806381d0e37b116100f9578063c42127b411610097578063f03528e711610071578063f03528e7146103f5578063f84da26e14610429578063fbf9204d1461043c578063fce011261461044557600080fd5b8063c42127b414610399578063c4b14d58146103ac578063ed2da0ac146103e257600080fd5b80639386775a116100d35780639386775a1461033c578063b5b7a1841461036a578063bb69ffcd14610373578063be271a021461038657600080fd5b806381d0e37b146102f65780638aaf0dae146103165780638cb941cc1461032957600080fd5b806369130451116101665780636e6bb97c116101405780636e6bb97c146102a857806376671808146102c8578063807896d5146102d057806380d85911146102e357600080fd5b8063691304511461026f5780636acd4f96146102825780636b5f444c1461029557600080fd5b8062708bb6146101ad578063013cf08b146101c957806303e7f672146101eb5780631f4f7d29146102005780633dad9ca9146102135780633e3b5b1914610246575b600080fd5b6101b660055481565b6040519081526020015b60405180910390f35b6101dc6101d7366004611605565b610470565b6040516101c09392919061161e565b6101fe6101f93660046116ac565b610538565b005b6101fe61020e3660046117db565b610a6a565b610236610221366004611605565b60096020526000908152604090205460ff1681565b60405190151581526020016101c0565b600080516020611dcd833981519152545b6040516001600160a01b0390911681526020016101c0565b6101fe61027d3660046117db565b610b56565b6101b66102903660046116ac565b610ca9565b6101fe6102a3366004611605565b610ce3565b6102bb6102b636600461183f565b610cf0565b6040516101c09190611861565b6101b6610d9f565b6101fe6102de366004611605565b610daf565b6101fe6102f13660046118ae565b611134565b6101b6610304366004611605565b60086020526000908152604090205481565b6101fe610324366004611605565b61124e565b6101fe6103373660046118da565b61125b565b61023661034a3660046118f7565b600760209081526000928352604080842090915290825290205460ff1681565b6101b660045481565b6102366103813660046118f7565b61127c565b610257610394366004611927565b61130e565b6101b66103a7366004611605565b611353565b6101b66103ba3660046118f7565b6000918252600b602090815260408084206001600160a01b0393909316845291905290205490565b6102366103f036600461195c565b611363565b610236610403366004611991565b600a60209081526000938452604080852082529284528284209052825290205460ff1681565b6101fe610437366004611605565b6113e2565b6101b660065481565b6101b66104533660046118f7565b600b60209081526000928352604080842090915290825290205481565b6003818154811061048057600080fd5b6000918252602090912060049091020180546001820180546001600160a01b039092169350906104af906119cf565b80601f01602080910402602001604051908101604052809291908181526020018280546104db906119cf565b80156105285780601f106104fd57610100808354040283529160200191610528565b820191906000526020600020905b81548152906001019060200180831161050b57829003601f168201915b5050505050908060030154905083565b6105406113ef565b600061054e85858585610ca9565b600254604051631015428760e21b81526001600160a01b0388811660048301529293509116906340550a1c90602401602060405180830381865afa15801561059a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105be9190611a09565b15156001146105e85760405162461bcd60e51b81526004016105df90611a2b565b60405180910390fd5b60008181526009602052604090205460ff161561065b5760405162461bcd60e51b815260206004820152602b60248201527f536c617368696e67566f74696e673a2076616c696461746f7220697320616c7260448201526a1958591e4818985b9b995960aa1b60648201526084016105df565b600081815260076020908152604080832033845290915290205460ff16156106eb5760405162461bcd60e51b815260206004820152603e60248201527f536c617368696e67566f74696e673a20766f74657220697320616c726561647960448201527f20766f74656420616761696e737420676976656e2076616c696461746f72000060648201526084016105df565b60008181526007602090815260408083203384528252808320805460ff191660011790558383526008909152812080549161072583611a8f565b91905055507f42ff2b7c8c611c525511dd04c1ee3cae48f313329b255979b94fc87a0f3a4a2633868660405161075d93929190611ae0565b60405180910390a1600061076f610d9f565b90506000600260009054906101000a90046001600160a01b03166001600160a01b0316639de702586040518163ffffffff1660e01b8152600401600060405180830381865afa1580156107c6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107ee9190810190611b05565b9050600281516107fe9190611bb7565b610809906001611bd9565b60008481526008602052604090205410610a15576000838152600960209081526040808320805460ff19166001908117909155858452600a83528184206001600160a01b038c1685529092528220909188600481111561086b5761086b611aa8565b600481111561087c5761087c611aa8565b815260208082019290925260409081016000908120805460ff191694151594909417909355848352600b82528083206001600160a01b038b16845290915281208054916108c883611a8f565b90915550506000828152600c60205260408120908760048111156108ee576108ee611aa8565b60048111156108ff576108ff611aa8565b815260208082019290925260409081016000908120805460018101825590825292902090910180546001600160a01b0319166001600160a01b038a16179055517f370cc65d87ae599d8b7dd97c0d1f33291921e8cc3652f51fe8db7d974e567c239061096e9089908990611bec565b60405180910390a1600186600481111561098a5761098a611aa8565b14806109a7575060028660048111156109a5576109a5611aa8565b145b15610a1557600260009054906101000a90046001600160a01b03166001600160a01b0316636c8fbb646040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156109fc57600080fd5b505af1158015610a10573d6000803e3d6000fd5b505050505b610a1f828861127c565b15610a61576040516001600160a01b03881681527f56848e1c0571e37fab91475a03170418e6a2956e066666c8038484240ea547099060200160405180910390a15b50505050505050565b610a726113ef565b600380546001810182556000919091526004027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810180546001600160a01b0385166001600160a01b0319909116178155907fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c01610af08382611c58565b50600354600090610b0390600190611d18565b604080518281526001600160a01b03871660208201529192507fcd423cc1203c0af96b9b3d68d73b3064a69de2d14450bb7181c5e5df2132b358910160405180910390a1610b5081610daf565b50505050565b610b5e6114b8565b600082905080600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b0316638e68dce46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bc8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bec9190611d2b565b600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b03166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c739190611d2b565b600280546001600160a01b0319166001600160a01b03929092169190911790555033600080516020611dcd833981519152555050565b600084848484604051602001610cc29493929190611d48565b6040516020818303038152906040528051906020012090505b949350505050565b610ceb61153c565b600455565b6060600c6000610cfe610d9f565b81526020019081526020016000206000836004811115610d2057610d20611aa8565b6004811115610d3157610d31611aa8565b8152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610d9357602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610d75575b50505050509050919050565b6000610daa43611353565b905090565b610db76113ef565b6003548110610e185760405162461bcd60e51b815260206004820152602760248201527f536c617368696e67566f74696e673a2070726f706f73616c20646f65736e27746044820152662065786973742160c81b60648201526084016105df565b600060038281548110610e2d57610e2d611d9f565b6000918252602090912060025460049283029091018054604051631015428760e21b81526001600160a01b039182169481019490945290935016906340550a1c90602401602060405180830381865afa158015610e8e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb29190611a09565b1515600114610ed35760405162461bcd60e51b81526004016105df90611a2b565b60038281548110610ee657610ee6611d9f565b600091825260208083203384526002600490930201919091019052604090205460ff1615610f715760405162461bcd60e51b815260206004820152603260248201527f536c617368696e67566f74696e673a20796f7520616c726561647920766f746560448201527119081a5b881d1a1a5cc81c1c9bdc1bdcd85b60721b60648201526084016105df565b600160038381548110610f8657610f86611d9f565b60009182526020808320338452600492909202909101600201905260409020805460ff19169115159190911790556003805483908110610fc857610fc8611d9f565b600091825260208220600360049092020101805491610fe683611a8f565b9091555050600254604080516313bce04b60e31b815290516000926001600160a01b031691639de7025891600480830192869291908290030181865afa158015611034573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261105c9190810190611b05565b90506002815161106c9190611bb7565b611077906001611bd9565b6003848154811061108a5761108a611d9f565b906000526020600020906004020160030154106110e6578154604080518581526001600160a01b0390921660208301527f9c85b616f29fca57a17eafe71cf9ff82ffef41766e2cf01ea7f8f7878dd3ec24910160405180910390a15b8154604080518581526001600160a01b0390921660208301523382820152517fd88f7b9f64fb7ba069d57fe9cedb25c7827ee4f7c67c7f0967f6a25bd6d0c53c9181900360600190a1505050565b600054610100900460ff16158080156111545750600054600160ff909116105b8061116e5750303b15801561116e575060005460ff166001145b6111d15760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105df565b6000805460ff1916600117905580156111f4576000805461ff0019166101001790555b6004849055600583905560068290558015610b50576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b61125661153c565b600655565b6112636114b8565b61127981600080516020611dcd83398151915255565b50565b600060065483101561129057506000611308565b6000835b6006546112a19086611d18565b8111156112ff576000818152600b602090815260408083206001600160a01b0388168452909152812054908190036112df5760009350505050611308565b6112e98184611bd9565b92505080806112f790611db5565b915050611294565b50600554111590505b92915050565b600c602052826000526040600020602052816000526040600020818154811061133657600080fd5b6000918252602090912001546001600160a01b0316925083915050565b6000600454826113089190611bb7565b6000600a6000611371610d9f565b81526020019081526020016000206000846001600160a01b03166001600160a01b0316815260200190815260200160002060008360048111156113b6576113b6611aa8565b60048111156113c7576113c7611aa8565b815260208101919091526040016000205460ff169392505050565b6113ea61153c565b600555565b600254604051631015428760e21b81523360048201526001600160a01b03909116906340550a1c90602401602060405180830381865afa158015611437573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061145b9190611a09565b6114b65760405162461bcd60e51b815260206004820152602660248201527f536c617368696e67566f74696e673a204e6f7420612073797374656d2076616c60448201526534b230ba37b960d11b60648201526084016105df565b565b60006114d0600080516020611dcd8339815191525490565b90506001600160a01b03811615806114f057506001600160a01b03811633145b6112795760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f7200000000000060448201526064016105df565b60005460408051637ac3c02f60e01b8152905133926201000090046001600160a01b031691637ac3c02f9160048083019260209291908290030181865afa15801561158b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115af9190611d2b565b6001600160a01b0316146114b65760405162461bcd60e51b815260206004820152601c60248201527f536c617368696e67566f74696e673a204e6f742061207369676e65720000000060448201526064016105df565b60006020828403121561161757600080fd5b5035919050565b60018060a01b038416815260006020606081840152845180606085015260005b8181101561165a5786810183015185820160800152820161163e565b506000608082860101526080601f19601f83011685010192505050826040830152949350505050565b6001600160a01b038116811461127957600080fd5b8035600581106116a757600080fd5b919050565b600080600080606085870312156116c257600080fd5b84356116cd81611683565b93506116db60208601611698565b9250604085013567ffffffffffffffff808211156116f857600080fd5b818701915087601f83011261170c57600080fd5b81358181111561171b57600080fd5b88602082850101111561172d57600080fd5b95989497505060200194505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561177b5761177b61173c565b604052919050565b600067ffffffffffffffff83111561179d5761179d61173c565b6117b0601f8401601f1916602001611752565b90508281528383830111156117c457600080fd5b828260208301376000602084830101529392505050565b600080604083850312156117ee57600080fd5b82356117f981611683565b9150602083013567ffffffffffffffff81111561181557600080fd5b8301601f8101851361182657600080fd5b61183585823560208401611783565b9150509250929050565b60006020828403121561185157600080fd5b61185a82611698565b9392505050565b6020808252825182820181905260009190848201906040850190845b818110156118a25783516001600160a01b03168352928401929184019160010161187d565b50909695505050505050565b6000806000606084860312156118c357600080fd5b505081359360208301359350604090920135919050565b6000602082840312156118ec57600080fd5b813561185a81611683565b6000806040838503121561190a57600080fd5b82359150602083013561191c81611683565b809150509250929050565b60008060006060848603121561193c57600080fd5b8335925061194c60208501611698565b9150604084013590509250925092565b6000806040838503121561196f57600080fd5b823561197a81611683565b915061198860208401611698565b90509250929050565b6000806000606084860312156119a657600080fd5b8335925060208401356119b881611683565b91506119c660408501611698565b90509250925092565b600181811c908216806119e357607f821691505b602082108103611a0357634e487b7160e01b600052602260045260246000fd5b50919050565b600060208284031215611a1b57600080fd5b8151801515811461185a57600080fd5b6020808252602e908201527f536c617368696e67566f74696e673a20746172676574206973206e6f7420616360408201526d3a34bb32903b30b634b230ba37b960911b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600060018201611aa157611aa1611a79565b5060010190565b634e487b7160e01b600052602160045260246000fd5b60058110611adc57634e487b7160e01b600052602160045260246000fd5b9052565b6001600160a01b0384811682528316602082015260608101610cdb6040830184611abe565b60006020808385031215611b1857600080fd5b825167ffffffffffffffff80821115611b3057600080fd5b818501915085601f830112611b4457600080fd5b815181811115611b5657611b5661173c565b8060051b9150611b67848301611752565b8181529183018401918481019088841115611b8157600080fd5b938501935b83851015611bab5784519250611b9b83611683565b8282529385019390850190611b86565b98975050505050505050565b600082611bd457634e487b7160e01b600052601260045260246000fd5b500490565b8082018082111561130857611308611a79565b6001600160a01b03831681526040810161185a6020830184611abe565b601f821115611c5357600081815260208120601f850160051c81016020861015611c305750805b601f850160051c820191505b81811015611c4f57828155600101611c3c565b5050505b505050565b815167ffffffffffffffff811115611c7257611c7261173c565b611c8681611c8084546119cf565b84611c09565b602080601f831160018114611cbb5760008415611ca35750858301515b600019600386901b1c1916600185901b178555611c4f565b600085815260208120601f198616915b82811015611cea57888601518255948401946001909101908401611ccb565b5085821015611d085787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b8181038181111561130857611308611a79565b600060208284031215611d3d57600080fd5b815161185a81611683565b6bffffffffffffffffffffffff198560601b168152600060058510611d7d57634e487b7160e01b600052602160045260246000fd5b8460f81b60148301528284601584013750600091016015019081529392505050565b634e487b7160e01b600052603260045260246000fd5b600081611dc457611dc4611a79565b50600019019056fe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a2646970667358221220fba291e3bea87ba543d56275415a183d02ab4d49807119169c2bd78b22b5417764736f6c63430008120033",
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

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _epochPeriod, uint256 _slashingThresold, uint256 _slashingEpochs) returns()
func (_SlashingVoting *SlashingVotingTransactor) Initialize(opts *bind.TransactOpts, _epochPeriod *big.Int, _slashingThresold *big.Int, _slashingEpochs *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "initialize", _epochPeriod, _slashingThresold, _slashingEpochs)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _epochPeriod, uint256 _slashingThresold, uint256 _slashingEpochs) returns()
func (_SlashingVoting *SlashingVotingSession) Initialize(_epochPeriod *big.Int, _slashingThresold *big.Int, _slashingEpochs *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _epochPeriod, _slashingThresold, _slashingEpochs)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _epochPeriod, uint256 _slashingThresold, uint256 _slashingEpochs) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) Initialize(_epochPeriod *big.Int, _slashingThresold *big.Int, _slashingEpochs *big.Int) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _epochPeriod, _slashingThresold, _slashingEpochs)
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
