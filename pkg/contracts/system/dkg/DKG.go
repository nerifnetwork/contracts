// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dkg

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

// DKGMetaData contains all meta data concerning the DKG contract.
var DKGMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"RoundDataFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"RoundDataProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectiveSigner\",\"type\":\"address\"}],\"name\":\"SignerVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ThresholdSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deadlinePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"generations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGenerationsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getRoundBroadcastCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getRoundBroadcastData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumDKG.GenerationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deadlinePeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isCurrentValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"isRoundFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastActiveGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_rawData\",\"type\":\"bytes\"}],\"name\":\"roundBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deadlinePeriod\",\"type\":\"uint256\"}],\"name\":\"setDeadlinePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGeneration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"voteSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506122d2806100206000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806369130451116100c3578063b32805c31161007c578063b32805c3146102fd578063c1718e5314610305578063c5f9dff014610318578063c88bc06714610320578063cd5fcd1514610333578063fe4b84df1461034657600080fd5b806369130451146102695780636db242621461027c57806382651c0d146102855780638cb941cc14610298578063ad51db6a146102ab578063ad8b0e94146102cb57600080fd5b806323f2a73f1161011557806323f2a73f146101cb5780633e3b5b19146101ee578063471f40fb1461020357806350c8548f1461022357806355614fcc146102365780635c622a0e1461024957600080fd5b8063100c11c31461015257806311af0a2014610167578063130b9702146101835780631a296e02146101a35780631ea0f036146101c3575b600080fd5b610165610160366004611c75565b610359565b005b61017060045481565b6040519081526020015b60405180910390f35b610196610191366004611cda565b610715565b60405161017a9190611d37565b6101ab6107f0565b6040516001600160a01b03909116815260200161017a565b600354610170565b6101de6101d9366004611d6a565b610826565b604051901515815260200161017a565b60008051602061227d833981519152546101ab565b610216610211366004611d9a565b610883565b60405161017a9190611df7565b610170610231366004611e0a565b610909565b6101de610244366004611e2c565b610945565b61025c610257366004611d9a565b6109b8565b60405161017a9190611e5f565b610165610277366004611e79565b610a39565b61017060055481565b610165610293366004611d9a565b610b6a565b6101656102a6366004611e2c565b610bec565b6101706102b9366004611e2c565b60026020526000908152604090205481565b6102de6102d9366004611d9a565b610c0d565b604080516001600160a01b03909316835260208301919091520161017a565b610165610c45565b6101de610313366004611e0a565b610fef565b610216611055565b61016561032e366004611ec9565b6110e8565b610170610341366004611d9a565b6114c9565b610165610354366004611d9a565b6114f8565b60035483908110801561039e57506003818154811061037a5761037a611f0c565b600091825260208083203384526003600790930201919091019052604090205460ff165b6103e65760405162461bcd60e51b81526020600482015260146024820152732225a39d103737ba1030903b30b634b230ba37b960611b60448201526064015b60405180910390fd5b836103f2600185611f38565b80158061045757506003828154811061040d5761040d611f0c565b9060005260206000209060070201600101805490506003838154811061043557610435611f0c565b6000918252602080832085845260066007909302019190910190526040902054145b61049f5760405162461bcd60e51b81526020600482015260196024820152781112d1ce881c9bdd5b99081dd85cc81b9bdd08199a5b1b1959603a1b60448201526064016103dd565b8585600382815481106104b4576104b4611f0c565b6000918252602080832084845260066007909302019190910181526040808320338452600101909152902080546104ea90611f4b565b1590506105395760405162461bcd60e51b815260206004820181905260248201527f444b473a20726f756e64206461746120616c72656164792070726f766964656460448201526064016103dd565b876000610545826109b8565b600281111561055657610556611e49565b146105a35760405162461bcd60e51b815260206004820152601d60248201527f444b473a206e6f7420612070656e64696e672067656e65726174696f6e00000060448201526064016103dd565b600389815481106105b6576105b6611f0c565b600091825260208083208b84526006600790930201919091019052604081208054916105e183611f85565b91905055508660038a815481106105fa576105fa611f0c565b600091825260208083208c8452600660079093020191909101815260408083203384526001019091529020906106309082611fed565b50604080518a8152602081018a9052338183015290517fca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c6370096069181900360600190a16003898154811061068357610683611f0c565b90600052602060002090600702016001018054905060038a815481106106ab576106ab611f0c565b600091825260208083208c8452600660079093020191909101905260409020540361070a57604080518a8152602081018a90527fab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9910160405180910390a15b505050505050505050565b60606003848154811061072a5761072a611f0c565b60009182526020808320868452600660079093020191909101815260408083206001600160a01b03861684526001019091529020805461076990611f4b565b80601f016020809104026020016040519081016040528092919081815260200182805461079590611f4b565b80156107e25780601f106107b7576101008083540402835291602001916107e2565b820191906000526020600020905b8154815290600101906020018083116107c557829003601f168201915b505050505090509392505050565b600060036004548154811061080757610807611f0c565b60009182526020909120600790910201546001600160a01b0316919050565b600354600090831015610879576003838154811061084657610846611f0c565b600091825260208083206001600160a01b03861684526003600790930201919091019052604090205460ff16905061087d565b5060005b92915050565b60606003828154811061089857610898611f0c565b90600052602060002090600702016001018054806020026020016040519081016040528092919081815260200182805480156108fd57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116108df575b50505050509050919050565b60006003838154811061091e5761091e611f0c565b60009182526020808320948352600791909102909301600601909252506040902054919050565b600480546040516323f2a73f60e01b8152918201526001600160a01b038216602482015260009030906323f2a73f90604401602060405180830381865afa158015610994573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061087d91906120ad565b6000806001600160a01b0316600383815481106109d7576109d7611f0c565b60009182526020909120600790910201546001600160a01b0316146109fe57506002919050565b4360038381548110610a1257610a12611f0c565b90600052602060002090600702016002015410610a3157506000919050565b506001919050565b610a41611663565b6000829050806001600160a01b0316638e68dce46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a84573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa891906120cf565b600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663fb9d9ac56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3091906120cf565b600180546001600160a01b0319166001600160a01b0392909216919091179055503360008051602061227d833981519152555050565b5050565b600360045481548110610b7f57610b7f611f0c565b60009182526020909120600790910201546001600160a01b03163314610be75760405162461bcd60e51b815260206004820152601860248201527f444b473a206e6f74206120616374697665207369676e6572000000000000000060448201526064016103dd565b600555565b610bf4611663565b610c0a8160008051602061227d83398151915255565b50565b60038181548110610c1d57600080fd5b6000918252602090912060079091020180546002909101546001600160a01b03909116915082565b6003805490600090610c58600184611f38565b81548110610c6857610c68611f0c565b9060005260206000209060070201905060008060008060029054906101000a90046001600160a01b03166001600160a01b031663b7ab4db56040518163ffffffff1660e01b8152600401600060405180830381865afa158015610ccf573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610cf791908101906120ec565b90506000815167ffffffffffffffff811115610d1557610d15611bbe565b604051908082528060200260200182016040528015610d3e578160200160208202803683370190505b50905060005b8251811015610e05576000838281518110610d6157610d61611f0c565b60200260200101519050610d768160016116e7565b80610d875750610d878160026116e7565b15610d925750610df3565b6001600160a01b038116600090815260038801602052604090205460ff16610db957600194505b80838781518110610dcc57610dcc611f0c565b6001600160a01b039092166020928302919091019091015285610dee81611f85565b965050505b80610dfd81611f85565b915050610d44565b5060018501546002851080610e2257508085148015610e22575083155b15610e305750505050505050565b60038054600101815560009081525b85811015610f395760038881548110610e5a57610e5a611f0c565b9060005260206000209060070201600101838281518110610e7d57610e7d611f0c565b60209081029190910181015182546001808201855560009485529290932090920180546001600160a01b0319166001600160a01b0390931692909217909155600380548a908110610ed057610ed0611f0c565b90600052602060002090600702016003016000858481518110610ef557610ef5611f0c565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580610f3181611f85565b915050610e3f565b50600554610f47904361219e565b60038881548110610f5a57610f5a611f0c565b600091825260209091206002600790920201015560048790556040517feadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a9590610fa590899085906121b1565b60405180910390a160408051888152600060208201527fab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9910160405180910390a150505050505050565b60006003838154811061100457611004611f0c565b9060005260206000209060070201600101805490506003848154811061102c5761102c611f0c565b600091825260208083208684526006600790930201919091019052604090205414905092915050565b600380546060919061106990600190611f38565b8154811061107957611079611f0c565b90600052602060002090600702016001018054806020026020016040519081016040528092919081815260200182805480156110de57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116110c0575b5050505050905090565b60035483908110801561112d57506003818154811061110957611109611f0c565b600091825260208083203384526003600790930201919091019052604090205460ff165b6111705760405162461bcd60e51b81526020600482015260146024820152732225a39d103737ba1030903b30b634b230ba37b960611b60448201526064016103dd565b83600380828154811061118557611185611f0c565b906000526020600020906007020160010180549050600383815481106111ad576111ad611f0c565b6000918252602080832085845260066007909302019190910190526040902054146112165760405162461bcd60e51b81526020600482015260196024820152781112d1ce881c9bdd5b99081dd85cc81b9bdd08199a5b1b1959603a1b60448201526064016103dd565b60006003878154811061122b5761122b611f0c565b9060005260206000209060070201905043816002015410156112865760405162461bcd60e51b81526020600482015260146024820152731112d1ce881d9bdd1a5b99c81a5cc8195b99195960621b60448201526064016103dd565b60006112b9866112b36040518060400160405280600681526020016576657269667960d01b815250611762565b9061179d565b9050866001600160a01b0316816001600160a01b03161461131c5760405162461bcd60e51b815260206004820152601960248201527f444b473a207369676e617475726520697320696e76616c69640000000000000060448201526064016103dd565b3360009081526004830160205260409020546001600160a01b0316156113795760405162461bcd60e51b81526020600482015260126024820152711112d1ce88185b1c9958591e481d9bdd195960721b60448201526064016103dd565b336000908152600483016020908152604080832080546001600160a01b0319166001600160a01b038c1690811790915583526005850190915281208054916113c083611f85565b9091555050604080518981523360208201526001600160a01b0389168183015290517f4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec199181900360600190a16001600160a01b0387166000908152600583016020526040812054611432908a906117c1565b83549091506001600160a01b03898116911614158180156114505750805b156114bd5783546001600160a01b0319166001600160a01b038a1690811785556000818152600260209081526040918290208d905581518d8152908101929092527fa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46910160405180910390a15b50505050505050505050565b6000600382815481106114de576114de611f0c565b600091825260209091206001600790920201015492915050565b600054610100900460ff16158080156115185750600054600160ff909116105b806115325750303b158015611532575060005460ff166001145b6115955760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103dd565b6000805460ff1916600117905580156115b8576000805461ff0019166101001790555b6003805460010180825560008281523392916115d6576115d6611f0c565b6000918252602080832060079290920290910180546001600160a01b0319166001600160a01b0394909416939093179092553381526002909152604081205560058290558015610b66576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050565b600061167b60008051602061227d8339815191525490565b90506001600160a01b038116158061169b57506001600160a01b03811633145b610c0a5760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f7200000000000060448201526064016103dd565b600154604051633b4b682b60e21b81526000916001600160a01b03169063ed2da0ac9061171a90869086906004016121d2565b602060405180830381865afa158015611737573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061175b91906120ad565b9392505050565b600061176e8251611800565b826040516020016117809291906121ff565b604051602081830303815290604052805190602001209050919050565b60008060006117ac8585611893565b915091506117b9816118d8565b509392505050565b60006002600384815481106117d8576117d8611f0c565b9060005260206000209060070201600101805490506117f7919061225a565b90911192915050565b6060600061180d83611a22565b600101905060008167ffffffffffffffff81111561182d5761182d611bbe565b6040519080825280601f01601f191660200182016040528015611857576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461186157509392505050565b60008082516041036118c95760208301516040840151606085015160001a6118bd87828585611afa565b945094505050506118d1565b506000905060025b9250929050565b60008160048111156118ec576118ec611e49565b036118f45750565b600181600481111561190857611908611e49565b036119555760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016103dd565b600281600481111561196957611969611e49565b036119b65760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016103dd565b60038160048111156119ca576119ca611e49565b03610c0a5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016103dd565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310611a615772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611a8d576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310611aab57662386f26fc10000830492506010015b6305f5e1008310611ac3576305f5e100830492506008015b6127108310611ad757612710830492506004015b60648310611ae9576064830492506002015b600a831061087d5760010192915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611b315750600090506003611bb5565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611b85573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611bae57600060019250925050611bb5565b9150600090505b94509492505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611bfd57611bfd611bbe565b604052919050565b600082601f830112611c1657600080fd5b813567ffffffffffffffff811115611c3057611c30611bbe565b611c43601f8201601f1916602001611bd4565b818152846020838601011115611c5857600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215611c8a57600080fd5b8335925060208401359150604084013567ffffffffffffffff811115611caf57600080fd5b611cbb86828701611c05565b9150509250925092565b6001600160a01b0381168114610c0a57600080fd5b600080600060608486031215611cef57600080fd5b83359250602084013591506040840135611d0881611cc5565b809150509250925092565b60005b83811015611d2e578181015183820152602001611d16565b50506000910152565b6020815260008251806020840152611d56816040850160208701611d13565b601f01601f19169190910160400192915050565b60008060408385031215611d7d57600080fd5b823591506020830135611d8f81611cc5565b809150509250929050565b600060208284031215611dac57600080fd5b5035919050565b600081518084526020808501945080840160005b83811015611dec5781516001600160a01b031687529582019590820190600101611dc7565b509495945050505050565b60208152600061175b6020830184611db3565b60008060408385031215611e1d57600080fd5b50508035926020909101359150565b600060208284031215611e3e57600080fd5b813561175b81611cc5565b634e487b7160e01b600052602160045260246000fd5b6020810160038310611e7357611e73611e49565b91905290565b60008060408385031215611e8c57600080fd5b8235611e9781611cc5565b9150602083013567ffffffffffffffff811115611eb357600080fd5b611ebf85828601611c05565b9150509250929050565b600080600060608486031215611ede57600080fd5b833592506020840135611ef081611cc5565b9150604084013567ffffffffffffffff811115611caf57600080fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561087d5761087d611f22565b600181811c90821680611f5f57607f821691505b602082108103611f7f57634e487b7160e01b600052602260045260246000fd5b50919050565b600060018201611f9757611f97611f22565b5060010190565b601f821115611fe857600081815260208120601f850160051c81016020861015611fc55750805b601f850160051c820191505b81811015611fe457828155600101611fd1565b5050505b505050565b815167ffffffffffffffff81111561200757612007611bbe565b61201b816120158454611f4b565b84611f9e565b602080601f83116001811461205057600084156120385750858301515b600019600386901b1c1916600185901b178555611fe4565b600085815260208120601f198616915b8281101561207f57888601518255948401946001909101908401612060565b508582101561209d5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000602082840312156120bf57600080fd5b8151801515811461175b57600080fd5b6000602082840312156120e157600080fd5b815161175b81611cc5565b600060208083850312156120ff57600080fd5b825167ffffffffffffffff8082111561211757600080fd5b818501915085601f83011261212b57600080fd5b81518181111561213d5761213d611bbe565b8060051b915061214e848301611bd4565b818152918301840191848101908884111561216857600080fd5b938501935b83851015612192578451925061218283611cc5565b828252938501939085019061216d565b98975050505050505050565b8082018082111561087d5761087d611f22565b8281526040602082015260006121ca6040830184611db3565b949350505050565b6001600160a01b038316815260408101600583106121f2576121f2611e49565b8260208301529392505050565b7f19457468657265756d205369676e6564204d6573736167653a0a00000000000081526000835161223781601a850160208801611d13565b83519083019061224e81601a840160208801611d13565b01601a01949350505050565b60008261227757634e487b7160e01b600052601260045260246000fd5b50049056fe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a26469706673582212209cd27fad1f03081e5e691248c14e0d6916bc0ce3e0bb1c6c96287349f90e7ddc64736f6c63430008120033",
}

// DKGABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGMetaData.ABI instead.
var DKGABI = DKGMetaData.ABI

// DKGBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGMetaData.Bin instead.
var DKGBin = DKGMetaData.Bin

// DeployDKG deploys a new Ethereum contract, binding an instance of DKG to it.
func DeployDKG(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DKG, error) {
	parsed, err := DKGMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DKG{DKGCaller: DKGCaller{contract: contract}, DKGTransactor: DKGTransactor{contract: contract}, DKGFilterer: DKGFilterer{contract: contract}}, nil
}

// DKG is an auto generated Go binding around an Ethereum contract.
type DKG struct {
	DKGCaller     // Read-only binding to the contract
	DKGTransactor // Write-only binding to the contract
	DKGFilterer   // Log filterer for contract events
}

// DKGCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGSession struct {
	Contract     *DKG              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGCallerSession struct {
	Contract *DKGCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DKGTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGTransactorSession struct {
	Contract     *DKGTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGRaw struct {
	Contract *DKG // Generic contract binding to access the raw methods on
}

// DKGCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGCallerRaw struct {
	Contract *DKGCaller // Generic read-only contract binding to access the raw methods on
}

// DKGTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGTransactorRaw struct {
	Contract *DKGTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKG creates a new instance of DKG, bound to a specific deployed contract.
func NewDKG(address common.Address, backend bind.ContractBackend) (*DKG, error) {
	contract, err := bindDKG(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKG{DKGCaller: DKGCaller{contract: contract}, DKGTransactor: DKGTransactor{contract: contract}, DKGFilterer: DKGFilterer{contract: contract}}, nil
}

// NewDKGCaller creates a new read-only instance of DKG, bound to a specific deployed contract.
func NewDKGCaller(address common.Address, caller bind.ContractCaller) (*DKGCaller, error) {
	contract, err := bindDKG(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGCaller{contract: contract}, nil
}

// NewDKGTransactor creates a new write-only instance of DKG, bound to a specific deployed contract.
func NewDKGTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGTransactor, error) {
	contract, err := bindDKG(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGTransactor{contract: contract}, nil
}

// NewDKGFilterer creates a new log filterer instance of DKG, bound to a specific deployed contract.
func NewDKGFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGFilterer, error) {
	contract, err := bindDKG(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGFilterer{contract: contract}, nil
}

// bindDKG binds a generic wrapper to an already deployed contract.
func bindDKG(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DKGMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.DKGCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.DKGTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKG *DKGCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKG.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKG *DKGTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKG *DKGTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKG.Contract.contract.Transact(opts, method, params...)
}

// DeadlinePeriod is a free data retrieval call binding the contract method 0x6db24262.
//
// Solidity: function deadlinePeriod() view returns(uint256)
func (_DKG *DKGCaller) DeadlinePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "deadlinePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeadlinePeriod is a free data retrieval call binding the contract method 0x6db24262.
//
// Solidity: function deadlinePeriod() view returns(uint256)
func (_DKG *DKGSession) DeadlinePeriod() (*big.Int, error) {
	return _DKG.Contract.DeadlinePeriod(&_DKG.CallOpts)
}

// DeadlinePeriod is a free data retrieval call binding the contract method 0x6db24262.
//
// Solidity: function deadlinePeriod() view returns(uint256)
func (_DKG *DKGCallerSession) DeadlinePeriod() (*big.Int, error) {
	return _DKG.Contract.DeadlinePeriod(&_DKG.CallOpts)
}

// Generations is a free data retrieval call binding the contract method 0xad8b0e94.
//
// Solidity: function generations(uint256 ) view returns(address signer, uint256 deadline)
func (_DKG *DKGCaller) Generations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Signer   common.Address
	Deadline *big.Int
}, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "generations", arg0)

	outstruct := new(struct {
		Signer   common.Address
		Deadline *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Deadline = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Generations is a free data retrieval call binding the contract method 0xad8b0e94.
//
// Solidity: function generations(uint256 ) view returns(address signer, uint256 deadline)
func (_DKG *DKGSession) Generations(arg0 *big.Int) (struct {
	Signer   common.Address
	Deadline *big.Int
}, error) {
	return _DKG.Contract.Generations(&_DKG.CallOpts, arg0)
}

// Generations is a free data retrieval call binding the contract method 0xad8b0e94.
//
// Solidity: function generations(uint256 ) view returns(address signer, uint256 deadline)
func (_DKG *DKGCallerSession) Generations(arg0 *big.Int) (struct {
	Signer   common.Address
	Deadline *big.Int
}, error) {
	return _DKG.Contract.Generations(&_DKG.CallOpts, arg0)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGCaller) GetCurrentValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getCurrentValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGSession) GetCurrentValidators() ([]common.Address, error) {
	return _DKG.Contract.GetCurrentValidators(&_DKG.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[])
func (_DKG *DKGCallerSession) GetCurrentValidators() ([]common.Address, error) {
	return _DKG.Contract.GetCurrentValidators(&_DKG.CallOpts)
}

// GetGenerationsCount is a free data retrieval call binding the contract method 0x1ea0f036.
//
// Solidity: function getGenerationsCount() view returns(uint256)
func (_DKG *DKGCaller) GetGenerationsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getGenerationsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGenerationsCount is a free data retrieval call binding the contract method 0x1ea0f036.
//
// Solidity: function getGenerationsCount() view returns(uint256)
func (_DKG *DKGSession) GetGenerationsCount() (*big.Int, error) {
	return _DKG.Contract.GetGenerationsCount(&_DKG.CallOpts)
}

// GetGenerationsCount is a free data retrieval call binding the contract method 0x1ea0f036.
//
// Solidity: function getGenerationsCount() view returns(uint256)
func (_DKG *DKGCallerSession) GetGenerationsCount() (*big.Int, error) {
	return _DKG.Contract.GetGenerationsCount(&_DKG.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_DKG *DKGCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_DKG *DKGSession) GetInjector() (common.Address, error) {
	return _DKG.Contract.GetInjector(&_DKG.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_DKG *DKGCallerSession) GetInjector() (common.Address, error) {
	return _DKG.Contract.GetInjector(&_DKG.CallOpts)
}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _generation, uint256 _round) view returns(uint256)
func (_DKG *DKGCaller) GetRoundBroadcastCount(opts *bind.CallOpts, _generation *big.Int, _round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRoundBroadcastCount", _generation, _round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _generation, uint256 _round) view returns(uint256)
func (_DKG *DKGSession) GetRoundBroadcastCount(_generation *big.Int, _round *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetRoundBroadcastCount(&_DKG.CallOpts, _generation, _round)
}

// GetRoundBroadcastCount is a free data retrieval call binding the contract method 0x50c8548f.
//
// Solidity: function getRoundBroadcastCount(uint256 _generation, uint256 _round) view returns(uint256)
func (_DKG *DKGCallerSession) GetRoundBroadcastCount(_generation *big.Int, _round *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetRoundBroadcastCount(&_DKG.CallOpts, _generation, _round)
}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _generation, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGCaller) GetRoundBroadcastData(opts *bind.CallOpts, _generation *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getRoundBroadcastData", _generation, _round, _validator)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _generation, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGSession) GetRoundBroadcastData(_generation *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	return _DKG.Contract.GetRoundBroadcastData(&_DKG.CallOpts, _generation, _round, _validator)
}

// GetRoundBroadcastData is a free data retrieval call binding the contract method 0x130b9702.
//
// Solidity: function getRoundBroadcastData(uint256 _generation, uint256 _round, address _validator) view returns(bytes)
func (_DKG *DKGCallerSession) GetRoundBroadcastData(_generation *big.Int, _round *big.Int, _validator common.Address) ([]byte, error) {
	return _DKG.Contract.GetRoundBroadcastData(&_DKG.CallOpts, _generation, _round, _validator)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_DKG *DKGCaller) GetSignerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getSignerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_DKG *DKGSession) GetSignerAddress() (common.Address, error) {
	return _DKG.Contract.GetSignerAddress(&_DKG.CallOpts)
}

// GetSignerAddress is a free data retrieval call binding the contract method 0x1a296e02.
//
// Solidity: function getSignerAddress() view returns(address)
func (_DKG *DKGCallerSession) GetSignerAddress() (common.Address, error) {
	return _DKG.Contract.GetSignerAddress(&_DKG.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _generation) view returns(uint8)
func (_DKG *DKGCaller) GetStatus(opts *bind.CallOpts, _generation *big.Int) (uint8, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getStatus", _generation)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _generation) view returns(uint8)
func (_DKG *DKGSession) GetStatus(_generation *big.Int) (uint8, error) {
	return _DKG.Contract.GetStatus(&_DKG.CallOpts, _generation)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 _generation) view returns(uint8)
func (_DKG *DKGCallerSession) GetStatus(_generation *big.Int) (uint8, error) {
	return _DKG.Contract.GetStatus(&_DKG.CallOpts, _generation)
}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _generation) view returns(address[])
func (_DKG *DKGCaller) GetValidators(opts *bind.CallOpts, _generation *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidators", _generation)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _generation) view returns(address[])
func (_DKG *DKGSession) GetValidators(_generation *big.Int) ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts, _generation)
}

// GetValidators is a free data retrieval call binding the contract method 0x471f40fb.
//
// Solidity: function getValidators(uint256 _generation) view returns(address[])
func (_DKG *DKGCallerSession) GetValidators(_generation *big.Int) ([]common.Address, error) {
	return _DKG.Contract.GetValidators(&_DKG.CallOpts, _generation)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0xcd5fcd15.
//
// Solidity: function getValidatorsCount(uint256 _generation) view returns(uint256)
func (_DKG *DKGCaller) GetValidatorsCount(opts *bind.CallOpts, _generation *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidatorsCount", _generation)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetValidatorsCount is a free data retrieval call binding the contract method 0xcd5fcd15.
//
// Solidity: function getValidatorsCount(uint256 _generation) view returns(uint256)
func (_DKG *DKGSession) GetValidatorsCount(_generation *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetValidatorsCount(&_DKG.CallOpts, _generation)
}

// GetValidatorsCount is a free data retrieval call binding the contract method 0xcd5fcd15.
//
// Solidity: function getValidatorsCount(uint256 _generation) view returns(uint256)
func (_DKG *DKGCallerSession) GetValidatorsCount(_generation *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetValidatorsCount(&_DKG.CallOpts, _generation)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address _validator) view returns(bool)
func (_DKG *DKGCaller) IsCurrentValidator(opts *bind.CallOpts, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isCurrentValidator", _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address _validator) view returns(bool)
func (_DKG *DKGSession) IsCurrentValidator(_validator common.Address) (bool, error) {
	return _DKG.Contract.IsCurrentValidator(&_DKG.CallOpts, _validator)
}

// IsCurrentValidator is a free data retrieval call binding the contract method 0x55614fcc.
//
// Solidity: function isCurrentValidator(address _validator) view returns(bool)
func (_DKG *DKGCallerSession) IsCurrentValidator(_validator common.Address) (bool, error) {
	return _DKG.Contract.IsCurrentValidator(&_DKG.CallOpts, _validator)
}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _generation, uint256 _round) view returns(bool)
func (_DKG *DKGCaller) IsRoundFilled(opts *bind.CallOpts, _generation *big.Int, _round *big.Int) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isRoundFilled", _generation, _round)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _generation, uint256 _round) view returns(bool)
func (_DKG *DKGSession) IsRoundFilled(_generation *big.Int, _round *big.Int) (bool, error) {
	return _DKG.Contract.IsRoundFilled(&_DKG.CallOpts, _generation, _round)
}

// IsRoundFilled is a free data retrieval call binding the contract method 0xc1718e53.
//
// Solidity: function isRoundFilled(uint256 _generation, uint256 _round) view returns(bool)
func (_DKG *DKGCallerSession) IsRoundFilled(_generation *big.Int, _round *big.Int) (bool, error) {
	return _DKG.Contract.IsRoundFilled(&_DKG.CallOpts, _generation, _round)
}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 _generation, address _validator) view returns(bool)
func (_DKG *DKGCaller) IsValidator(opts *bind.CallOpts, _generation *big.Int, _validator common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isValidator", _generation, _validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 _generation, address _validator) view returns(bool)
func (_DKG *DKGSession) IsValidator(_generation *big.Int, _validator common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, _generation, _validator)
}

// IsValidator is a free data retrieval call binding the contract method 0x23f2a73f.
//
// Solidity: function isValidator(uint256 _generation, address _validator) view returns(bool)
func (_DKG *DKGCallerSession) IsValidator(_generation *big.Int, _validator common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, _generation, _validator)
}

// LastActiveGeneration is a free data retrieval call binding the contract method 0x11af0a20.
//
// Solidity: function lastActiveGeneration() view returns(uint256)
func (_DKG *DKGCaller) LastActiveGeneration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "lastActiveGeneration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastActiveGeneration is a free data retrieval call binding the contract method 0x11af0a20.
//
// Solidity: function lastActiveGeneration() view returns(uint256)
func (_DKG *DKGSession) LastActiveGeneration() (*big.Int, error) {
	return _DKG.Contract.LastActiveGeneration(&_DKG.CallOpts)
}

// LastActiveGeneration is a free data retrieval call binding the contract method 0x11af0a20.
//
// Solidity: function lastActiveGeneration() view returns(uint256)
func (_DKG *DKGCallerSession) LastActiveGeneration() (*big.Int, error) {
	return _DKG.Contract.LastActiveGeneration(&_DKG.CallOpts)
}

// SignerToGeneration is a free data retrieval call binding the contract method 0xad51db6a.
//
// Solidity: function signerToGeneration(address ) view returns(uint256)
func (_DKG *DKGCaller) SignerToGeneration(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "signerToGeneration", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignerToGeneration is a free data retrieval call binding the contract method 0xad51db6a.
//
// Solidity: function signerToGeneration(address ) view returns(uint256)
func (_DKG *DKGSession) SignerToGeneration(arg0 common.Address) (*big.Int, error) {
	return _DKG.Contract.SignerToGeneration(&_DKG.CallOpts, arg0)
}

// SignerToGeneration is a free data retrieval call binding the contract method 0xad51db6a.
//
// Solidity: function signerToGeneration(address ) view returns(uint256)
func (_DKG *DKGCallerSession) SignerToGeneration(arg0 common.Address) (*big.Int, error) {
	return _DKG.Contract.SignerToGeneration(&_DKG.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactor) Initialize(opts *bind.TransactOpts, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "initialize", _deadlinePeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _deadlinePeriod) returns()
func (_DKG *DKGSession) Initialize(_deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _deadlinePeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactorSession) Initialize(_deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _deadlinePeriod)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _generation, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGTransactor) RoundBroadcast(opts *bind.TransactOpts, _generation *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "roundBroadcast", _generation, _round, _rawData)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _generation, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGSession) RoundBroadcast(_generation *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.Contract.RoundBroadcast(&_DKG.TransactOpts, _generation, _round, _rawData)
}

// RoundBroadcast is a paid mutator transaction binding the contract method 0x100c11c3.
//
// Solidity: function roundBroadcast(uint256 _generation, uint256 _round, bytes _rawData) returns()
func (_DKG *DKGTransactorSession) RoundBroadcast(_generation *big.Int, _round *big.Int, _rawData []byte) (*types.Transaction, error) {
	return _DKG.Contract.RoundBroadcast(&_DKG.TransactOpts, _generation, _round, _rawData)
}

// SetDeadlinePeriod is a paid mutator transaction binding the contract method 0x82651c0d.
//
// Solidity: function setDeadlinePeriod(uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactor) SetDeadlinePeriod(opts *bind.TransactOpts, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "setDeadlinePeriod", _deadlinePeriod)
}

// SetDeadlinePeriod is a paid mutator transaction binding the contract method 0x82651c0d.
//
// Solidity: function setDeadlinePeriod(uint256 _deadlinePeriod) returns()
func (_DKG *DKGSession) SetDeadlinePeriod(_deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.SetDeadlinePeriod(&_DKG.TransactOpts, _deadlinePeriod)
}

// SetDeadlinePeriod is a paid mutator transaction binding the contract method 0x82651c0d.
//
// Solidity: function setDeadlinePeriod(uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactorSession) SetDeadlinePeriod(_deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.SetDeadlinePeriod(&_DKG.TransactOpts, _deadlinePeriod)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_DKG *DKGTransactor) SetDependencies(opts *bind.TransactOpts, _contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "setDependencies", _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_DKG *DKGSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _DKG.Contract.SetDependencies(&_DKG.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_DKG *DKGTransactorSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _DKG.Contract.SetDependencies(&_DKG.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_DKG *DKGTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_DKG *DKGSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _DKG.Contract.SetInjector(&_DKG.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_DKG *DKGTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _DKG.Contract.SetInjector(&_DKG.TransactOpts, injector_)
}

// UpdateGeneration is a paid mutator transaction binding the contract method 0xb32805c3.
//
// Solidity: function updateGeneration() returns()
func (_DKG *DKGTransactor) UpdateGeneration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "updateGeneration")
}

// UpdateGeneration is a paid mutator transaction binding the contract method 0xb32805c3.
//
// Solidity: function updateGeneration() returns()
func (_DKG *DKGSession) UpdateGeneration() (*types.Transaction, error) {
	return _DKG.Contract.UpdateGeneration(&_DKG.TransactOpts)
}

// UpdateGeneration is a paid mutator transaction binding the contract method 0xb32805c3.
//
// Solidity: function updateGeneration() returns()
func (_DKG *DKGTransactorSession) UpdateGeneration() (*types.Transaction, error) {
	return _DKG.Contract.UpdateGeneration(&_DKG.TransactOpts)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _generation, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGTransactor) VoteSigner(opts *bind.TransactOpts, _generation *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "voteSigner", _generation, _signerAddress, _signature)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _generation, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGSession) VoteSigner(_generation *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.Contract.VoteSigner(&_DKG.TransactOpts, _generation, _signerAddress, _signature)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _generation, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGTransactorSession) VoteSigner(_generation *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.Contract.VoteSigner(&_DKG.TransactOpts, _generation, _signerAddress, _signature)
}

// DKGInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DKG contract.
type DKGInitializedIterator struct {
	Event *DKGInitialized // Event containing the contract specifics and raw log

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
func (it *DKGInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGInitialized)
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
		it.Event = new(DKGInitialized)
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
func (it *DKGInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGInitialized represents a Initialized event raised by the DKG contract.
type DKGInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DKG *DKGFilterer) FilterInitialized(opts *bind.FilterOpts) (*DKGInitializedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DKGInitializedIterator{contract: _DKG.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DKG *DKGFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DKGInitialized) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGInitialized)
				if err := _DKG.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DKG *DKGFilterer) ParseInitialized(log types.Log) (*DKGInitialized, error) {
	event := new(DKGInitialized)
	if err := _DKG.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRoundDataFilledIterator is returned from FilterRoundDataFilled and is used to iterate over the raw logs and unpacked data for RoundDataFilled events raised by the DKG contract.
type DKGRoundDataFilledIterator struct {
	Event *DKGRoundDataFilled // Event containing the contract specifics and raw log

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
func (it *DKGRoundDataFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRoundDataFilled)
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
		it.Event = new(DKGRoundDataFilled)
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
func (it *DKGRoundDataFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRoundDataFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRoundDataFilled represents a RoundDataFilled event raised by the DKG contract.
type DKGRoundDataFilled struct {
	Generation *big.Int
	Round      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoundDataFilled is a free log retrieval operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 generation, uint256 round)
func (_DKG *DKGFilterer) FilterRoundDataFilled(opts *bind.FilterOpts) (*DKGRoundDataFilledIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "RoundDataFilled")
	if err != nil {
		return nil, err
	}
	return &DKGRoundDataFilledIterator{contract: _DKG.contract, event: "RoundDataFilled", logs: logs, sub: sub}, nil
}

// WatchRoundDataFilled is a free log subscription operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 generation, uint256 round)
func (_DKG *DKGFilterer) WatchRoundDataFilled(opts *bind.WatchOpts, sink chan<- *DKGRoundDataFilled) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "RoundDataFilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRoundDataFilled)
				if err := _DKG.contract.UnpackLog(event, "RoundDataFilled", log); err != nil {
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

// ParseRoundDataFilled is a log parse operation binding the contract event 0xab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9.
//
// Solidity: event RoundDataFilled(uint256 generation, uint256 round)
func (_DKG *DKGFilterer) ParseRoundDataFilled(log types.Log) (*DKGRoundDataFilled, error) {
	event := new(DKGRoundDataFilled)
	if err := _DKG.contract.UnpackLog(event, "RoundDataFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRoundDataProvidedIterator is returned from FilterRoundDataProvided and is used to iterate over the raw logs and unpacked data for RoundDataProvided events raised by the DKG contract.
type DKGRoundDataProvidedIterator struct {
	Event *DKGRoundDataProvided // Event containing the contract specifics and raw log

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
func (it *DKGRoundDataProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRoundDataProvided)
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
		it.Event = new(DKGRoundDataProvided)
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
func (it *DKGRoundDataProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRoundDataProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRoundDataProvided represents a RoundDataProvided event raised by the DKG contract.
type DKGRoundDataProvided struct {
	Generation *big.Int
	Round      *big.Int
	Validator  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoundDataProvided is a free log retrieval operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 generation, uint256 round, address validator)
func (_DKG *DKGFilterer) FilterRoundDataProvided(opts *bind.FilterOpts) (*DKGRoundDataProvidedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "RoundDataProvided")
	if err != nil {
		return nil, err
	}
	return &DKGRoundDataProvidedIterator{contract: _DKG.contract, event: "RoundDataProvided", logs: logs, sub: sub}, nil
}

// WatchRoundDataProvided is a free log subscription operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 generation, uint256 round, address validator)
func (_DKG *DKGFilterer) WatchRoundDataProvided(opts *bind.WatchOpts, sink chan<- *DKGRoundDataProvided) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "RoundDataProvided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRoundDataProvided)
				if err := _DKG.contract.UnpackLog(event, "RoundDataProvided", log); err != nil {
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

// ParseRoundDataProvided is a log parse operation binding the contract event 0xca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c637009606.
//
// Solidity: event RoundDataProvided(uint256 generation, uint256 round, address validator)
func (_DKG *DKGFilterer) ParseRoundDataProvided(log types.Log) (*DKGRoundDataProvided, error) {
	event := new(DKGRoundDataProvided)
	if err := _DKG.contract.UnpackLog(event, "RoundDataProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGSignerAddressUpdatedIterator is returned from FilterSignerAddressUpdated and is used to iterate over the raw logs and unpacked data for SignerAddressUpdated events raised by the DKG contract.
type DKGSignerAddressUpdatedIterator struct {
	Event *DKGSignerAddressUpdated // Event containing the contract specifics and raw log

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
func (it *DKGSignerAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGSignerAddressUpdated)
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
		it.Event = new(DKGSignerAddressUpdated)
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
func (it *DKGSignerAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGSignerAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGSignerAddressUpdated represents a SignerAddressUpdated event raised by the DKG contract.
type DKGSignerAddressUpdated struct {
	Generation    *big.Int
	SignerAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignerAddressUpdated is a free log retrieval operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 generation, address signerAddress)
func (_DKG *DKGFilterer) FilterSignerAddressUpdated(opts *bind.FilterOpts) (*DKGSignerAddressUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "SignerAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGSignerAddressUpdatedIterator{contract: _DKG.contract, event: "SignerAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerAddressUpdated is a free log subscription operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 generation, address signerAddress)
func (_DKG *DKGFilterer) WatchSignerAddressUpdated(opts *bind.WatchOpts, sink chan<- *DKGSignerAddressUpdated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "SignerAddressUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGSignerAddressUpdated)
				if err := _DKG.contract.UnpackLog(event, "SignerAddressUpdated", log); err != nil {
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

// ParseSignerAddressUpdated is a log parse operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 generation, address signerAddress)
func (_DKG *DKGFilterer) ParseSignerAddressUpdated(log types.Log) (*DKGSignerAddressUpdated, error) {
	event := new(DKGSignerAddressUpdated)
	if err := _DKG.contract.UnpackLog(event, "SignerAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGSignerVotedIterator is returned from FilterSignerVoted and is used to iterate over the raw logs and unpacked data for SignerVoted events raised by the DKG contract.
type DKGSignerVotedIterator struct {
	Event *DKGSignerVoted // Event containing the contract specifics and raw log

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
func (it *DKGSignerVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGSignerVoted)
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
		it.Event = new(DKGSignerVoted)
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
func (it *DKGSignerVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGSignerVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGSignerVoted represents a SignerVoted event raised by the DKG contract.
type DKGSignerVoted struct {
	Generation       *big.Int
	Validator        common.Address
	CollectiveSigner common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSignerVoted is a free log retrieval operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 generation, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) FilterSignerVoted(opts *bind.FilterOpts) (*DKGSignerVotedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "SignerVoted")
	if err != nil {
		return nil, err
	}
	return &DKGSignerVotedIterator{contract: _DKG.contract, event: "SignerVoted", logs: logs, sub: sub}, nil
}

// WatchSignerVoted is a free log subscription operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 generation, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) WatchSignerVoted(opts *bind.WatchOpts, sink chan<- *DKGSignerVoted) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "SignerVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGSignerVoted)
				if err := _DKG.contract.UnpackLog(event, "SignerVoted", log); err != nil {
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

// ParseSignerVoted is a log parse operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 generation, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) ParseSignerVoted(log types.Log) (*DKGSignerVoted, error) {
	event := new(DKGSignerVoted)
	if err := _DKG.contract.UnpackLog(event, "SignerVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGThresholdSignerUpdatedIterator is returned from FilterThresholdSignerUpdated and is used to iterate over the raw logs and unpacked data for ThresholdSignerUpdated events raised by the DKG contract.
type DKGThresholdSignerUpdatedIterator struct {
	Event *DKGThresholdSignerUpdated // Event containing the contract specifics and raw log

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
func (it *DKGThresholdSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGThresholdSignerUpdated)
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
		it.Event = new(DKGThresholdSignerUpdated)
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
func (it *DKGThresholdSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGThresholdSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGThresholdSignerUpdated represents a ThresholdSignerUpdated event raised by the DKG contract.
type DKGThresholdSignerUpdated struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterThresholdSignerUpdated is a free log retrieval operation binding the contract event 0x8559aaab9220e2f109d51f7d09a3cca9b41451f4698287d982dc4549a70808bf.
//
// Solidity: event ThresholdSignerUpdated(address signer)
func (_DKG *DKGFilterer) FilterThresholdSignerUpdated(opts *bind.FilterOpts) (*DKGThresholdSignerUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ThresholdSignerUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGThresholdSignerUpdatedIterator{contract: _DKG.contract, event: "ThresholdSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchThresholdSignerUpdated is a free log subscription operation binding the contract event 0x8559aaab9220e2f109d51f7d09a3cca9b41451f4698287d982dc4549a70808bf.
//
// Solidity: event ThresholdSignerUpdated(address signer)
func (_DKG *DKGFilterer) WatchThresholdSignerUpdated(opts *bind.WatchOpts, sink chan<- *DKGThresholdSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ThresholdSignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGThresholdSignerUpdated)
				if err := _DKG.contract.UnpackLog(event, "ThresholdSignerUpdated", log); err != nil {
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

// ParseThresholdSignerUpdated is a log parse operation binding the contract event 0x8559aaab9220e2f109d51f7d09a3cca9b41451f4698287d982dc4549a70808bf.
//
// Solidity: event ThresholdSignerUpdated(address signer)
func (_DKG *DKGFilterer) ParseThresholdSignerUpdated(log types.Log) (*DKGThresholdSignerUpdated, error) {
	event := new(DKGThresholdSignerUpdated)
	if err := _DKG.contract.UnpackLog(event, "ThresholdSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the DKG contract.
type DKGValidatorsUpdatedIterator struct {
	Event *DKGValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *DKGValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGValidatorsUpdated)
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
		it.Event = new(DKGValidatorsUpdated)
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
func (it *DKGValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGValidatorsUpdated represents a ValidatorsUpdated event raised by the DKG contract.
type DKGValidatorsUpdated struct {
	Generation *big.Int
	Validators []common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 generation, address[] validators)
func (_DKG *DKGFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts) (*DKGValidatorsUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGValidatorsUpdatedIterator{contract: _DKG.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 generation, address[] validators)
func (_DKG *DKGFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *DKGValidatorsUpdated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ValidatorsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGValidatorsUpdated)
				if err := _DKG.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
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

// ParseValidatorsUpdated is a log parse operation binding the contract event 0xeadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a95.
//
// Solidity: event ValidatorsUpdated(uint256 generation, address[] validators)
func (_DKG *DKGFilterer) ParseValidatorsUpdated(log types.Log) (*DKGValidatorsUpdated, error) {
	event := new(DKGValidatorsUpdated)
	if err := _DKG.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
