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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"RoundDataFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"RoundDataProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectiveSigner\",\"type\":\"address\"}],\"name\":\"SignerVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ThresholdSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadlinePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"generations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGenerationsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"getRoundBroadcastCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getRoundBroadcastData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumDKG.GenerationStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"}],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadlinePeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isCurrentValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"}],\"name\":\"isRoundFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastActiveGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_rawData\",\"type\":\"bytes\"}],\"name\":\"roundBroadcast\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deadlinePeriod\",\"type\":\"uint256\"}],\"name\":\"setDeadlinePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerToGeneration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateGeneration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_generation\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"voteSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506122cb806100206000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c806378a5c206116100de578063c1718e5311610097578063cd5fcd1511610071578063cd5fcd15146103c9578063cd6dc687146103dc578063ce119a8a146103ef578063faaa8a641461042b57600080fd5b8063c1718e531461039b578063c5f9dff0146103ae578063c88bc067146103b657600080fd5b806378a5c206146102e657806382651c0d14610315578063abf410e514610328578063ad51db6a14610341578063ad8b0e9414610361578063b32805c31461039357600080fd5b80633a9783f31161014b57806355614fcc1161012557806355614fcc14610284578063561ff9a9146102975780635c622a0e146102bd5780636db24262146102dd57600080fd5b80633a9783f31461022f578063471f40fb1461025157806350c8548f1461027157600080fd5b8063100c11c31461019357806311af0a20146101a8578063130b9702146101c45780631a296e02146101e45780631ea0f0361461020457806323f2a73f1461020c575b600080fd5b6101a66101a1366004611ca7565b610459565b005b6101b160035481565b6040519081526020015b60405180910390f35b6101d76101d2366004611d0c565b610815565b6040516101bb9190611d95565b6101ec6108f0565b6040516001600160a01b0390911681526020016101bb565b6002546101b1565b61021f61021a366004611da8565b610926565b60405190151581526020016101bb565b6101d760405180604001604052806003815260200162646b6760e81b81525081565b61026461025f366004611dd8565b610983565b6040516101bb9190611e35565b6101b161027f366004611e48565b610a09565b61021f610292366004611e6a565b610a45565b6101d7604051806040016040528060078152602001667374616b696e6760c81b81525081565b6102d06102cb366004611dd8565b610abb565b6040516101bb9190611e9d565b6101b160045481565b6101d76040518060400160405280601081526020016f737570706f727465642d746f6b656e7360801b81525081565b6101a6610323366004611dd8565b610b3c565b6000546101ec906201000090046001600160a01b031681565b6101b161034f366004611e6a565b60016020526000908152604090205481565b61037461036f366004611dd8565b610bbe565b604080516001600160a01b0390931683526020830191909152016101bb565b6101a6610bf6565b61021f6103a9366004611e48565b610f93565b610264610ff9565b6101a66103c4366004611eb7565b61108c565b6101b16103d7366004611dd8565b61146e565b6101a66103ea366004611efa565b61149d565b6101d76040518060400160405280601881526020017f7265776172642d646973747269627574696f6e2d706f6f6c000000000000000081525081565b6101d76040518060400160405280600f81526020016e736c617368696e672d766f74696e6760881b81525081565b60025483908110801561049e57506002818154811061047a5761047a611f26565b600091825260208083203384526003600790930201919091019052604090205460ff165b6104e65760405162461bcd60e51b81526020600482015260146024820152732225a39d103737ba1030903b30b634b230ba37b960611b60448201526064015b60405180910390fd5b836104f2600185611f52565b80158061055757506002828154811061050d5761050d611f26565b9060005260206000209060070201600101805490506002838154811061053557610535611f26565b6000918252602080832085845260066007909302019190910190526040902054145b61059f5760405162461bcd60e51b81526020600482015260196024820152781112d1ce881c9bdd5b99081dd85cc81b9bdd08199a5b1b1959603a1b60448201526064016104dd565b8585600282815481106105b4576105b4611f26565b6000918252602080832084845260066007909302019190910181526040808320338452600101909152902080546105ea90611f65565b1590506106395760405162461bcd60e51b815260206004820181905260248201527f444b473a20726f756e64206461746120616c72656164792070726f766964656460448201526064016104dd565b87600061064582610abb565b600281111561065657610656611e87565b146106a35760405162461bcd60e51b815260206004820152601d60248201527f444b473a206e6f7420612070656e64696e672067656e65726174696f6e00000060448201526064016104dd565b600289815481106106b6576106b6611f26565b600091825260208083208b84526006600790930201919091019052604081208054916106e183611f9f565b91905055508660028a815481106106fa576106fa611f26565b600091825260208083208c8452600660079093020191909101815260408083203384526001019091529020906107309082612006565b50604080518a8152602081018a9052338183015290517fca56b6e939787236f062daae635dc1afa2b46ad9a24ad09aa98833c6370096069181900360600190a16002898154811061078357610783611f26565b90600052602060002090600702016001018054905060028a815481106107ab576107ab611f26565b600091825260208083208c8452600660079093020191909101905260409020540361080a57604080518a8152602081018a90527fab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9910160405180910390a15b505050505050505050565b60606002848154811061082a5761082a611f26565b60009182526020808320868452600660079093020191909101815260408083206001600160a01b03861684526001019091529020805461086990611f65565b80601f016020809104026020016040519081016040528092919081815260200182805461089590611f65565b80156108e25780601f106108b7576101008083540402835291602001916108e2565b820191906000526020600020905b8154815290600101906020018083116108c557829003601f168201915b505050505090509392505050565b600060026003548154811061090757610907611f26565b60009182526020909120600790910201546001600160a01b0316919050565b600254600090831015610979576002838154811061094657610946611f26565b600091825260208083206001600160a01b03861684526003600790930201919091019052604090205460ff16905061097d565b5060005b92915050565b60606002828154811061099857610998611f26565b90600052602060002090600702016001018054806020026020016040519081016040528092919081815260200182805480156109fd57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116109df575b50505050509050919050565b600060028381548110610a1e57610a1e611f26565b60009182526020808320948352600791909102909301600601909252506040902054919050565b6003546040516323f2a73f60e01b815260048101919091526001600160a01b038216602482015260009030906323f2a73f90604401602060405180830381865afa158015610a97573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097d91906120c6565b6000806001600160a01b031660028381548110610ada57610ada611f26565b60009182526020909120600790910201546001600160a01b031614610b0157506002919050565b4360028381548110610b1557610b15611f26565b90600052602060002090600702016002015410610b3457506000919050565b506001919050565b600260035481548110610b5157610b51611f26565b60009182526020909120600790910201546001600160a01b03163314610bb95760405162461bcd60e51b815260206004820152601860248201527f444b473a206e6f74206120616374697665207369676e6572000000000000000060448201526064016104dd565b600455565b60028181548110610bce57600080fd5b6000918252602090912060079091020180546002909101546001600160a01b03909116915082565b6002805490600090610c09600184611f52565b81548110610c1957610c19611f26565b906000526020600020906007020190506000806000610c36611625565b6001600160a01b031663b7ab4db56040518163ffffffff1660e01b8152600401600060405180830381865afa158015610c73573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610c9b91908101906120e8565b90506000815167ffffffffffffffff811115610cb957610cb9611bf0565b604051908082528060200260200182016040528015610ce2578160200160208202803683370190505b50905060005b8251811015610da9576000838281518110610d0557610d05611f26565b60200260200101519050610d1a8160016116bb565b80610d2b5750610d2b8160026116bb565b15610d365750610d97565b6001600160a01b038116600090815260038801602052604090205460ff16610d5d57600194505b80838781518110610d7057610d70611f26565b6001600160a01b039092166020928302919091019091015285610d9281611f9f565b965050505b80610da181611f9f565b915050610ce8565b5060018501546002851080610dc657508085148015610dc6575083155b15610dd45750505050505050565b60028054600101815560009081525b85811015610edd5760028881548110610dfe57610dfe611f26565b9060005260206000209060070201600101838281518110610e2157610e21611f26565b60209081029190910181015182546001808201855560009485529290932090920180546001600160a01b0319166001600160a01b0390931692909217909155600280548a908110610e7457610e74611f26565b90600052602060002090600702016003016000858481518110610e9957610e99611f26565b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191691151591909117905580610ed581611f9f565b915050610de3565b50600454610eeb904361219a565b60028881548110610efe57610efe611f26565b600091825260209091206002600790920201015560038790556040517feadf82e9da8b1722bf1769001bdd6d52bb429e0745d9116f69495cedc7db8a9590610f4990899085906121ad565b60405180910390a160408051888152600060208201527fab74ab6fc458020cf5d6116f5c013ebf3c0ad518f10de1391427c225f75db5f9910160405180910390a150505050505050565b600060028381548110610fa857610fa8611f26565b90600052602060002090600702016001018054905060028481548110610fd057610fd0611f26565b600091825260208083208684526006600790930201919091019052604090205414905092915050565b600280546060919061100d90600190611f52565b8154811061101d5761101d611f26565b906000526020600020906007020160010180548060200260200160405190810160405280929190818152602001828054801561108257602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611064575b5050505050905090565b6002548390811080156110d15750600281815481106110ad576110ad611f26565b600091825260208083203384526003600790930201919091019052604090205460ff165b6111145760405162461bcd60e51b81526020600482015260146024820152732225a39d103737ba1030903b30b634b230ba37b960611b60448201526064016104dd565b8360036002828154811061112a5761112a611f26565b9060005260206000209060070201600101805490506002838154811061115257611152611f26565b6000918252602080832085845260066007909302019190910190526040902054146111bb5760405162461bcd60e51b81526020600482015260196024820152781112d1ce881c9bdd5b99081dd85cc81b9bdd08199a5b1b1959603a1b60448201526064016104dd565b6000600287815481106111d0576111d0611f26565b90600052602060002090600702019050438160020154101561122b5760405162461bcd60e51b81526020600482015260146024820152731112d1ce881d9bdd1a5b99c81a5cc8195b99195960621b60448201526064016104dd565b600061125e866112586040518060400160405280600681526020016576657269667960d01b81525061173a565b90611775565b9050866001600160a01b0316816001600160a01b0316146112c15760405162461bcd60e51b815260206004820152601960248201527f444b473a207369676e617475726520697320696e76616c69640000000000000060448201526064016104dd565b3360009081526004830160205260409020546001600160a01b03161561131e5760405162461bcd60e51b81526020600482015260126024820152711112d1ce88185b1c9958591e481d9bdd195960721b60448201526064016104dd565b336000908152600483016020908152604080832080546001600160a01b0319166001600160a01b038c16908117909155835260058501909152812080549161136583611f9f565b9091555050604080518981523360208201526001600160a01b0389168183015290517f4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec199181900360600190a16001600160a01b03871660009081526005830160205260408120546113d7908a90611799565b83549091506001600160a01b03898116911614158180156113f55750805b156114625783546001600160a01b0319166001600160a01b038a1690811785556000818152600160209081526040918290208d905581518d8152908101929092527fa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46910160405180910390a15b50505050505050505050565b60006002828154811061148357611483611f26565b600091825260209091206001600790920201015492915050565b600054610100900460ff16158080156114bd5750600054600160ff909116105b806114d75750303b1580156114d7575060005460ff166001145b61153a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016104dd565b6000805460ff19166001179055801561155d576000805461ff0019166101001790555b60028054600101808255600082815233929161157b5761157b611f26565b60009182526020808320600790920290910180546001600160a01b039485166001600160a01b031990911617905533825260019052604081208190558054918516620100000262010000600160b01b031990921691909117905560048290558015611620576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b6000805460408051808201825260078152667374616b696e6760c81b60208201529051633581777360e01b8152620100009092046001600160a01b03169163358177739161167591600401611d95565b602060405180830381865afa158015611692573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116b691906121ce565b905090565b60006116c56117d7565b6001600160a01b031663ed2da0ac84846040518363ffffffff1660e01b81526004016116f29291906121eb565b602060405180830381865afa15801561170f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061173391906120c6565b9392505050565b6000611746825161182f565b82604051602001611758929190612218565b604051602081830303815290604052805190602001209050919050565b600080600061178485856118c2565b9150915061179181611907565b509392505050565b600060028084815481106117af576117af611f26565b9060005260206000209060070201600101805490506117ce9190612273565b90911192915050565b60008054604080518082018252600f81526e736c617368696e672d766f74696e6760881b60208201529051633581777360e01b8152620100009092046001600160a01b03169163358177739161167591600401611d95565b6060600061183c83611a54565b600101905060008167ffffffffffffffff81111561185c5761185c611bf0565b6040519080825280601f01601f191660200182016040528015611886576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461189057509392505050565b60008082516041036118f85760208301516040840151606085015160001a6118ec87828585611b2c565b94509450505050611900565b506000905060025b9250929050565b600081600481111561191b5761191b611e87565b036119235750565b600181600481111561193757611937611e87565b036119845760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016104dd565b600281600481111561199857611998611e87565b036119e55760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016104dd565b60038160048111156119f9576119f9611e87565b03611a515760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016104dd565b50565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310611a935772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611abf576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310611add57662386f26fc10000830492506010015b6305f5e1008310611af5576305f5e100830492506008015b6127108310611b0957612710830492506004015b60648310611b1b576064830492506002015b600a831061097d5760010192915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611b635750600090506003611be7565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611bb7573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611be057600060019250925050611be7565b9150600090505b94509492505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611c2f57611c2f611bf0565b604052919050565b600082601f830112611c4857600080fd5b813567ffffffffffffffff811115611c6257611c62611bf0565b611c75601f8201601f1916602001611c06565b818152846020838601011115611c8a57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215611cbc57600080fd5b8335925060208401359150604084013567ffffffffffffffff811115611ce157600080fd5b611ced86828701611c37565b9150509250925092565b6001600160a01b0381168114611a5157600080fd5b600080600060608486031215611d2157600080fd5b83359250602084013591506040840135611d3a81611cf7565b809150509250925092565b60005b83811015611d60578181015183820152602001611d48565b50506000910152565b60008151808452611d81816020860160208601611d45565b601f01601f19169290920160200192915050565b6020815260006117336020830184611d69565b60008060408385031215611dbb57600080fd5b823591506020830135611dcd81611cf7565b809150509250929050565b600060208284031215611dea57600080fd5b5035919050565b600081518084526020808501945080840160005b83811015611e2a5781516001600160a01b031687529582019590820190600101611e05565b509495945050505050565b6020815260006117336020830184611df1565b60008060408385031215611e5b57600080fd5b50508035926020909101359150565b600060208284031215611e7c57600080fd5b813561173381611cf7565b634e487b7160e01b600052602160045260246000fd5b6020810160038310611eb157611eb1611e87565b91905290565b600080600060608486031215611ecc57600080fd5b833592506020840135611ede81611cf7565b9150604084013567ffffffffffffffff811115611ce157600080fd5b60008060408385031215611f0d57600080fd5b8235611f1881611cf7565b946020939093013593505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b8181038181111561097d5761097d611f3c565b600181811c90821680611f7957607f821691505b602082108103611f9957634e487b7160e01b600052602260045260246000fd5b50919050565b600060018201611fb157611fb1611f3c565b5060010190565b601f82111561162057600081815260208120601f850160051c81016020861015611fdf5750805b601f850160051c820191505b81811015611ffe57828155600101611feb565b505050505050565b815167ffffffffffffffff81111561202057612020611bf0565b6120348161202e8454611f65565b84611fb8565b602080601f83116001811461206957600084156120515750858301515b600019600386901b1c1916600185901b178555611ffe565b600085815260208120601f198616915b8281101561209857888601518255948401946001909101908401612079565b50858210156120b65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6000602082840312156120d857600080fd5b8151801515811461173357600080fd5b600060208083850312156120fb57600080fd5b825167ffffffffffffffff8082111561211357600080fd5b818501915085601f83011261212757600080fd5b81518181111561213957612139611bf0565b8060051b915061214a848301611c06565b818152918301840191848101908884111561216457600080fd5b938501935b8385101561218e578451925061217e83611cf7565b8282529385019390850190612169565b98975050505050505050565b8082018082111561097d5761097d611f3c565b8281526040602082015260006121c66040830184611df1565b949350505050565b6000602082840312156121e057600080fd5b815161173381611cf7565b6001600160a01b0383168152604081016005831061220b5761220b611e87565b8260208301529392505050565b7f19457468657265756d205369676e6564204d6573736167653a0a00000000000081526000835161225081601a850160208801611d45565b83519083019061226781601a840160208801611d45565b01601a01949350505050565b60008261229057634e487b7160e01b600052601260045260246000fd5b50049056fea264697066735822122073bf9c98b47aa20495db606146d66f2d651a4259909db3a21625a03e8733faba64736f6c63430008120033",
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

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_DKG *DKGCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_DKG *DKGSession) DKGKEY() (string, error) {
	return _DKG.Contract.DKGKEY(&_DKG.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_DKG *DKGCallerSession) DKGKEY() (string, error) {
	return _DKG.Contract.DKGKEY(&_DKG.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_DKG *DKGCaller) REWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_DKG *DKGSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _DKG.Contract.REWARDDISTRIBUTIONPOOLKEY(&_DKG.CallOpts)
}

// REWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0xce119a8a.
//
// Solidity: function REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_DKG *DKGCallerSession) REWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _DKG.Contract.REWARDDISTRIBUTIONPOOLKEY(&_DKG.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_DKG *DKGCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_DKG *DKGSession) SLASHINGVOTINGKEY() (string, error) {
	return _DKG.Contract.SLASHINGVOTINGKEY(&_DKG.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_DKG *DKGCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _DKG.Contract.SLASHINGVOTINGKEY(&_DKG.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_DKG *DKGCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_DKG *DKGSession) STAKINGKEY() (string, error) {
	return _DKG.Contract.STAKINGKEY(&_DKG.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_DKG *DKGCallerSession) STAKINGKEY() (string, error) {
	return _DKG.Contract.STAKINGKEY(&_DKG.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_DKG *DKGCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_DKG *DKGSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _DKG.Contract.SUPPORTEDTOKENSKEY(&_DKG.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_DKG *DKGCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _DKG.Contract.SUPPORTEDTOKENSKEY(&_DKG.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_DKG *DKGCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_DKG *DKGSession) ContractRegistry() (common.Address, error) {
	return _DKG.Contract.ContractRegistry(&_DKG.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_DKG *DKGCallerSession) ContractRegistry() (common.Address, error) {
	return _DKG.Contract.ContractRegistry(&_DKG.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _contractRegistry, uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactor) Initialize(opts *bind.TransactOpts, _contractRegistry common.Address, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "initialize", _contractRegistry, _deadlinePeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _contractRegistry, uint256 _deadlinePeriod) returns()
func (_DKG *DKGSession) Initialize(_contractRegistry common.Address, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _contractRegistry, _deadlinePeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _contractRegistry, uint256 _deadlinePeriod) returns()
func (_DKG *DKGTransactorSession) Initialize(_contractRegistry common.Address, _deadlinePeriod *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _contractRegistry, _deadlinePeriod)
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
