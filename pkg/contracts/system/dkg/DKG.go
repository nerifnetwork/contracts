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

// IDKGDKGEpochInfo is an auto generated low-level Go binding around an user-defined struct.
type IDKGDKGEpochInfo struct {
	EpochId        *big.Int
	EpochStartTime *big.Int
	EpochSigner    common.Address
	EpochStatus    uint8
}

// IDKGValidationData is an auto generated low-level Go binding around an user-defined struct.
type IDKGValidationData struct {
	ValidationTime  *big.Int
	ValidationEpoch *big.Int
}

// IDKGValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IDKGValidatorInfo struct {
	Validator           common.Address
	StartValidationData IDKGValidationData
	EndValidationData   IDKGValidationData
}

// DKGMetaData contains all meta data concerning the DKG contract.
var DKGMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"}],\"name\":\"NewEpochCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startValidationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startValidationEpoch\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectiveSigner\",\"type\":\"address\"}],\"name\":\"SignerVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endValidationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endValidationEpoch\",\"type\":\"uint256\"}],\"name\":\"ValidatorExitAnnounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorToAdd\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorToExit\",\"type\":\"address\"}],\"name\":\"announceValidatorExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dkgGenerationEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_activeValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpochStatus\",\"outputs\":[{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getDKGEpochInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"epochSigner\",\"type\":\"address\"},{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"epochStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIDKG.DKGEpochInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getDKGPeriodEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getEpochEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getEpochStatus\",\"outputs\":[{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddr\",\"type\":\"address\"}],\"name\":\"getSignerVotesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getUpdatedCollectionPeriodEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"validationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.ValidationData\",\"name\":\"startValidationData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"validationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.ValidationData\",\"name\":\"endValidationData\",\"type\":\"tuple\"}],\"internalType\":\"structIDKG.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"getValidatorVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guaranteedWorkingEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updatesCollectionEpochDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dkgGenerationEpochDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_guaranteedWorkingEpochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"isCurrentEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"isDKGGenSuccessful\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"isLastEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorToRemove\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatesCollectionEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"voteSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612250806100206000396000f3fe608060405234801561001057600080fd5b50600436106101fb5760003560e01c806372ef10b21161011a578063c88bc067116100ad578063f072ec081161007c578063f072ec081461043a578063f3513a3714610443578063f7897c071461044b578063f9cf0dd814610485578063facd743b1461049857600080fd5b8063c88bc067146103e1578063ca158348146103f4578063d4ee9f8d14610414578063e80f65341461042757600080fd5b80639de70258116100e95780639de70258146103a9578063a29a839f146103be578063c0474753146103c6578063c3125405146103d957600080fd5b806372ef10b21461035a57806380d85911146103635780638a11d7c9146103765780638cb941cc1461039657600080fd5b80633e3b5b19116101925780634d238c8e116101615780634d238c8e146103235780636913045114610336578063696f2959146103495780636c8fbb641461035257600080fd5b80633e3b5b19146102de57806340550a1c146102f357806340a141ff146103065780634666a62b1461031b57600080fd5b8063214e6a56116101ce578063214e6a561461027357806331147c301461029857806336503f42146102a05780633ae0725a146102b357600080fd5b8063124e3ffa1461020057806315accc92146102305780631a296e02146102505780632020166814610261575b600080fd5b61021361020e366004611da1565b6104ab565b6040516001600160a01b0390911681526020015b60405180910390f35b61024361023e366004611dd1565b6104dd565b6040516102279190611e22565b6005546001600160a01b0316610213565b6004545b604051908152602001610227565b610288610281366004611dd1565b6004541490565b6040519015158152602001610227565b6102656105a0565b6102656102ae366004611dd1565b6105f0565b6102886102c1366004611dd1565b6000908152600860205260409020546001600160a01b0316151590565b6000805160206121fb83398151915254610213565b610288610301366004611e30565b610610565b610319610314366004611e30565b6106df565b005b6102436107c6565b610319610331366004611e30565b6107d8565b610319610344366004611ef0565b6108f6565b61026560035481565b6103196109a6565b61026560015481565b610319610371366004611f40565b610a81565b610389610384366004611e30565b610c59565b6040516102279190611f6c565b6103196103a4366004611e30565b610cc2565b6103b1610ce0565b6040516102279190611faf565b610265610d59565b6102656103d4366004611dd1565b610d86565b610265610d9e565b6103196103ef366004611ffc565b610daa565b610407610402366004611dd1565b61105b565b6040516102279190612055565b610265610422366004611dd1565b6110e5565b610288610435366004611dd1565b6110f3565b61026560025481565b6103b1611105565b610265610459366004611da1565b60008281526008602090815260408083206001600160a01b038516845260020190915290205492915050565b610265610493366004611e30565b611111565b6102886104a6366004611e30565b611275565b60008281526008602090815260408083206001600160a01b038086168552600390910190925290912054165b92915050565b60006004548211156104f157506000919050565b6000828152600860205260409020600101546004548314801561051357504281115b156105215750600192915050565b61052a836110f3565b6105375750600692915050565b426001548261054691906120a3565b915081106105575750600292915050565b426002548261056691906120a3565b915081106105775750600392915050565b426003548261058691906120a3565b915081106105975750600492915050565b50600592915050565b6000806105ad6006611282565b905060005b818110156105eb576105c861030160068361128c565b156105db57826105d7816120b6565b9350505b6105e4816120b6565b90506105b2565b505090565b600180546000838152600860205260408120909201546104d791906120a3565b6001600160a01b0381811660009081526009602090815260408083208151608081018352815481840190815260018301546060830152815282518084018452600283015481526003909201548285015280840191909152805183015184526008909252822054919290918391161515801561068d57508151514210155b905060006106ba8360200151602001516000908152600860205260409020546001600160a01b0316151590565b15806106ca575060208301515142105b90508180156106d65750805b95945050505050565b6106e761129f565b6106f2600682611301565b6107435760405162461bcd60e51b815260206004820152601d60248201527f444b473a2056616c696461746f7220646f6573206e6f7420657869737400000060448201526064015b60405180910390fd5b6001600160a01b0381166000908152600960205260409020600201544210156107ba5760405162461bcd60e51b815260206004820152602360248201527f444b473a2056616c696461746f722063616e27742062652072656d6f766564206044820152621e595d60ea1b606482015260840161073a565b6107c381611323565b50565b60006107d361023e610d59565b905090565b6107e061129f565b6107e981611275565b156108365760405162461bcd60e51b815260206004820152601d60248201527f444b473a2056616c696461746f7220616c726561647920657869737473000000604482015260640161073a565b6000610840611387565b604080518082018252828152815180830183526000198152600060208281018290528084019283526001600160a01b038816825260098152939020915180518355830151600183015551805160028301559091015160039091015590506108a86006836113bf565b50805160208083015160408051938452918301526001600160a01b038416917fcd38d7f8bcd459fe58ab4e4f5ba2ee51418a1a0cf207a6747f61bb7945e62b41910160405180910390a25050565b6108fe6113d4565b6000829050806001600160a01b0316638e68dce46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610941573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096591906120cf565b600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550506109a2336000805160206121fb83398151915255565b5050565b60006109b26006611458565b905060005b81518110156109a2576000600960008484815181106109d8576109d86120ec565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000209050610a24838381518110610a1757610a176120ec565b6020026020010151610610565b15610a3a57610a3581600201611465565b610a43565b610a4381611465565b60028101544210610a7057610a70838381518110610a6357610a636120ec565b6020026020010151611323565b50610a7a816120b6565b90506109b7565b600054610100900460ff1615808015610aa15750600054600160ff909116105b80610abb5750303b158015610abb575060005460ff166001145b610b1e5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161073a565b6000805460ff191660011790558015610b41576000805461ff0019166101001790555b60018490556002839055600382905560048054600091908290610b63906120b6565b918290555060058054336001600160a01b0319918216811790925560008381526008602052604090208054909116821781554260019190910155909150610bac906006906113bf565b506040805160808101825242818301908152606082019390935291825280518082018252600019815260006020828101829052808501928352338252600981529290209251805184558201516001840155518051600284015501516003909101558015610c53576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610c61611d2d565b506001600160a01b031660008181526009602090815260409182902082516060810184529384528251808401845281548152600182015481840152848301528251808401845260028201548152600390910154918101919091529082015290565b610cca6113d4565b6107c3816000805160206121fb83398151915255565b60606000610cec6114b0565b90506000610cfa6006611282565b905060005b81811015610d3f576000610d1460068361128c565b9050610d1f81610610565b15610d2e57610d2e84826114d7565b50610d38816120b6565b9050610cff565b50610d52828051602001516060906104d7565b9250505090565b600454600081815260086020526040902060010154421015610d835780610d7f81612102565b9150505b90565b6000600254610d94836105f0565b6104d791906120a3565b60006107d36006611282565b610db26114ec565b6003610dbd846104dd565b6006811115610dce57610dce611dea565b14610e1b5760405162461bcd60e51b815260206004820181905260248201527f444b473a204e6f74206120444b472067656e65726174696f6e20706572696f64604482015260640161073a565b6000610e4e82610e486040518060400160405280600681526020016576657269667960d01b815250611541565b9061157c565b9050826001600160a01b0316816001600160a01b031614610eb15760405162461bcd60e51b815260206004820152601960248201527f444b473a205369676e617475726520697320696e76616c696400000000000000604482015260640161073a565b600084815260086020908152604080832033845260038101909252909120546001600160a01b031615610f1b5760405162461bcd60e51b81526020600482015260126024820152711112d1ce88105b1c9958591e481d9bdd195960721b604482015260640161073a565b6001600160a01b0384166000908152600282016020526040812080548290610f42906120b6565b918290555033600081815260038501602090815260409182902080546001600160a01b0319166001600160a01b038b1690811790915582518b815291820193909352908101919091529091507f4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec199060600160405180910390a16002610fc56105a0565b610fcf9190612119565b81118015610feb57506005546001600160a01b03868116911614155b156110535781546001600160a01b0386166001600160a01b0319918216811784556005805490921681179091556040805188815260208101929092527fa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46910160405180910390a15b505050505050565b6110836040805160808101825260008082526020820181905291810182905290606082015290565b600082815260086020908152604091829020825160808101845285815260018201549281019290925280546001600160a01b031692820192909252606081016110cb856104dd565b60068111156110dc576110dc611dea565b90529392505050565b6000600354610d9483610d86565b6000816110fe610d59565b1492915050565b60606107d36006611458565b600061111b61129f565b61112482610610565b6111705760405162461bcd60e51b815260206004820152601c60248201527f444b473a2056616c696461746f72206973206e6f742061637469766500000000604482015260640161073a565b6001600160a01b038216600090815260096020526040902060020154600019146111fa5760405162461bcd60e51b815260206004820152603560248201527f444b473a2045786974206f66207468652076616c696461746f722068617320616044820152741b1c9958591e481899595b88185b9b9bdd5b98d959605a1b606482015260840161073a565b6000611204611387565b6001600160a01b03841660008181526009602090815260409182902084516002820181905585830151600390920182905583519081529182015292935090917fd75af7adb28dda89e53e18c225f03e4aee3f0132f19eaff0a642b1f88433baad910160405180910390a25192915050565b60006104d7600683611301565b60006104d7825490565b600061129883836115a0565b9392505050565b6000546201000090046001600160a01b031633146112ff5760405162461bcd60e51b815260206004820152601a60248201527f444b473a204e6f742061207374616b696e672061646472657373000000000000604482015260640161073a565b565b6001600160a01b03811660009081526001830160205260408120541515611298565b61132e6006826115ca565b506001600160a01b0381166000818152600960205260408082208281556001810183905560028101839055600301829055517fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f19190a250565b60408051808201909152600080825260208201526004546113a66115df565b9150808260200151146113bb576113bb6109a6565b5090565b6000611298836001600160a01b0384166116e2565b60006113ec6000805160206121fb8339815191525490565b90506001600160a01b038116158061140c57506001600160a01b03811633145b6107c35760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f72000000000000604482015260640161073a565b6060600061129883611731565b80544210801590611490575060018101546000908152600860205260409020546001600160a01b0316155b156107c357600061149f6115df565b805183556020015160018301555050565b60408051606081018252600060208201818152928201529081526114d261178d565b815290565b81516109a2906001600160a01b0383166117c1565b6114f533610610565b6112ff5760405162461bcd60e51b815260206004820152601c60248201527f444b473a204e6f7420616e206163746976652076616c696461746f7200000000604482015260640161073a565b600061154d8251611812565b8260405160200161155f92919061215f565b604051602081830303815290604052805190602001209050919050565b600080600061158b85856118a5565b91509150611598816118ea565b509392505050565b60008260000182815481106115b7576115b76120ec565b9060005260206000200154905092915050565b6000611298836001600160a01b038416611a34565b604080518082019091526000808252602082015260006115fd610d59565b9050600061160a826104dd565b905081600282600681111561162157611621611dea565b146116be57506004548281036116be5742600583600681111561164657611646611dea565b1461165757611654826110e5565b90505b600460008154611666906120b6565b918290555060008181526008602090815260409182902060010184905581518381529081018490529193507ff27e0ff471920cfae653a7a16cf64784b13fafb16b991f9af7860027c1b475d7910160405180910390a1505b60405180604001604052806116d283610d86565b8152602001919091529392505050565b6000818152600183016020526040812054611729575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556104d7565b5060006104d7565b60608160000180548060200260200160405190810160405280929190818152602001828054801561178157602002820191906000526020600020905b81548152602001906001019080831161176d575b50505050509050919050565b60408051808201825260008082526020820152815160a08101909252906117b5816001611b2e565b60058252602082015290565b60006117cf83602001515190565b83519091506117df8260016120a3565b036117fc5782516117fc9084906117f79060026121ba565b611b4f565b6020928301516001820181529083020190910152565b6060600061181f83611b91565b600101905060008167ffffffffffffffff81111561183f5761183f611e4d565b6040519080825280601f01601f191660200182016040528015611869576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461187357509392505050565b60008082516041036118db5760208301516040840151606085015160001a6118cf87828585611c69565b945094505050506118e3565b506000905060025b9250929050565b60008160048111156118fe576118fe611dea565b036119065750565b600181600481111561191a5761191a611dea565b036119675760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161073a565b600281600481111561197b5761197b611dea565b036119c85760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161073a565b60038160048111156119dc576119dc611dea565b036107c35760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161073a565b60008181526001830160205260408120548015611b1d576000611a586001836121d1565b8554909150600090611a6c906001906121d1565b9050818114611ad1576000866000018281548110611a8c57611a8c6120ec565b9060005260206000200154905080876000018481548110611aaf57611aaf6120ec565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611ae257611ae26121e4565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506104d7565b60009150506104d7565b5092915050565b60005b60208202811015611b4a57600083820152602001611b31565b505050565b604080516020830281019091526020830151805160005b602080830201811015611b83578281015184820152602001611b66565b505050908252602090910152565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310611bd05772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611bfc576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310611c1a57662386f26fc10000830492506010015b6305f5e1008310611c32576305f5e100830492506008015b6127108310611c4657612710830492506004015b60648310611c58576064830492506002015b600a83106104d75760010192915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611ca05750600090506003611d24565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611cf4573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611d1d57600060019250925050611d24565b9150600090505b94509492505050565b604051806060016040528060006001600160a01b03168152602001611d65604051806040016040528060008152602001600081525090565b8152602001611d87604051806040016040528060008152602001600081525090565b905290565b6001600160a01b03811681146107c357600080fd5b60008060408385031215611db457600080fd5b823591506020830135611dc681611d8c565b809150509250929050565b600060208284031215611de357600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b60078110611e1e57634e487b7160e01b600052602160045260246000fd5b9052565b602081016104d78284611e00565b600060208284031215611e4257600080fd5b813561129881611d8c565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611e7457600080fd5b813567ffffffffffffffff80821115611e8f57611e8f611e4d565b604051601f8301601f19908116603f01168101908282118183101715611eb757611eb7611e4d565b81604052838152866020858801011115611ed057600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215611f0357600080fd5b8235611f0e81611d8c565b9150602083013567ffffffffffffffff811115611f2a57600080fd5b611f3685828601611e63565b9150509250929050565b600080600060608486031215611f5557600080fd5b505081359360208301359350604090920135919050565b81516001600160a01b0316815260208083015180518284015290810151604083015260a08201905060408301518051606084015260208101516080840152611b27565b6020808252825182820181905260009190848201906040850190845b81811015611ff05783516001600160a01b031683529284019291840191600101611fcb565b50909695505050505050565b60008060006060848603121561201157600080fd5b83359250602084013561202381611d8c565b9150604084013567ffffffffffffffff81111561203f57600080fd5b61204b86828701611e63565b9150509250925092565b81518152602080830151908201526040808301516001600160a01b0316908201526060808301516080830191611b2790840182611e00565b634e487b7160e01b600052601160045260246000fd5b808201808211156104d7576104d761208d565b6000600182016120c8576120c861208d565b5060010190565b6000602082840312156120e157600080fd5b815161129881611d8c565b634e487b7160e01b600052603260045260246000fd5b6000816121115761211161208d565b506000190190565b60008261213657634e487b7160e01b600052601260045260246000fd5b500490565b60005b8381101561215657818101518382015260200161213e565b50506000910152565b7f19457468657265756d205369676e6564204d6573736167653a0a00000000000081526000835161219781601a85016020880161213b565b8351908301906121ae81601a84016020880161213b565b01601a01949350505050565b80820281158282048414176104d7576104d761208d565b818103818111156104d7576104d761208d565b634e487b7160e01b600052603160045260246000fdfe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a264697066735822122067c22d6d74151d08c4bd0c6c641f34e1d92671695cb62e4c14e0c36bc87a18f264736f6c63430008120033",
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

// DkgGenerationEpochDuration is a free data retrieval call binding the contract method 0xf072ec08.
//
// Solidity: function dkgGenerationEpochDuration() view returns(uint256)
func (_DKG *DKGCaller) DkgGenerationEpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "dkgGenerationEpochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DkgGenerationEpochDuration is a free data retrieval call binding the contract method 0xf072ec08.
//
// Solidity: function dkgGenerationEpochDuration() view returns(uint256)
func (_DKG *DKGSession) DkgGenerationEpochDuration() (*big.Int, error) {
	return _DKG.Contract.DkgGenerationEpochDuration(&_DKG.CallOpts)
}

// DkgGenerationEpochDuration is a free data retrieval call binding the contract method 0xf072ec08.
//
// Solidity: function dkgGenerationEpochDuration() view returns(uint256)
func (_DKG *DKGCallerSession) DkgGenerationEpochDuration() (*big.Int, error) {
	return _DKG.Contract.DkgGenerationEpochDuration(&_DKG.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_DKG *DKGCaller) GetActiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getActiveValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_DKG *DKGSession) GetActiveValidators() ([]common.Address, error) {
	return _DKG.Contract.GetActiveValidators(&_DKG.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_DKG *DKGCallerSession) GetActiveValidators() ([]common.Address, error) {
	return _DKG.Contract.GetActiveValidators(&_DKG.CallOpts)
}

// GetActiveValidatorsCount is a free data retrieval call binding the contract method 0x31147c30.
//
// Solidity: function getActiveValidatorsCount() view returns(uint256 _activeValidatorsCount)
func (_DKG *DKGCaller) GetActiveValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getActiveValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveValidatorsCount is a free data retrieval call binding the contract method 0x31147c30.
//
// Solidity: function getActiveValidatorsCount() view returns(uint256 _activeValidatorsCount)
func (_DKG *DKGSession) GetActiveValidatorsCount() (*big.Int, error) {
	return _DKG.Contract.GetActiveValidatorsCount(&_DKG.CallOpts)
}

// GetActiveValidatorsCount is a free data retrieval call binding the contract method 0x31147c30.
//
// Solidity: function getActiveValidatorsCount() view returns(uint256 _activeValidatorsCount)
func (_DKG *DKGCallerSession) GetActiveValidatorsCount() (*big.Int, error) {
	return _DKG.Contract.GetActiveValidatorsCount(&_DKG.CallOpts)
}

// GetAllValidators is a free data retrieval call binding the contract method 0xf3513a37.
//
// Solidity: function getAllValidators() view returns(address[])
func (_DKG *DKGCaller) GetAllValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getAllValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllValidators is a free data retrieval call binding the contract method 0xf3513a37.
//
// Solidity: function getAllValidators() view returns(address[])
func (_DKG *DKGSession) GetAllValidators() ([]common.Address, error) {
	return _DKG.Contract.GetAllValidators(&_DKG.CallOpts)
}

// GetAllValidators is a free data retrieval call binding the contract method 0xf3513a37.
//
// Solidity: function getAllValidators() view returns(address[])
func (_DKG *DKGCallerSession) GetAllValidators() ([]common.Address, error) {
	return _DKG.Contract.GetAllValidators(&_DKG.CallOpts)
}

// GetAllValidatorsCount is a free data retrieval call binding the contract method 0xc3125405.
//
// Solidity: function getAllValidatorsCount() view returns(uint256)
func (_DKG *DKGCaller) GetAllValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getAllValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllValidatorsCount is a free data retrieval call binding the contract method 0xc3125405.
//
// Solidity: function getAllValidatorsCount() view returns(uint256)
func (_DKG *DKGSession) GetAllValidatorsCount() (*big.Int, error) {
	return _DKG.Contract.GetAllValidatorsCount(&_DKG.CallOpts)
}

// GetAllValidatorsCount is a free data retrieval call binding the contract method 0xc3125405.
//
// Solidity: function getAllValidatorsCount() view returns(uint256)
func (_DKG *DKGCallerSession) GetAllValidatorsCount() (*big.Int, error) {
	return _DKG.Contract.GetAllValidatorsCount(&_DKG.CallOpts)
}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256 _currentEpochId)
func (_DKG *DKGCaller) GetCurrentEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getCurrentEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256 _currentEpochId)
func (_DKG *DKGSession) GetCurrentEpochId() (*big.Int, error) {
	return _DKG.Contract.GetCurrentEpochId(&_DKG.CallOpts)
}

// GetCurrentEpochId is a free data retrieval call binding the contract method 0xa29a839f.
//
// Solidity: function getCurrentEpochId() view returns(uint256 _currentEpochId)
func (_DKG *DKGCallerSession) GetCurrentEpochId() (*big.Int, error) {
	return _DKG.Contract.GetCurrentEpochId(&_DKG.CallOpts)
}

// GetCurrentEpochStatus is a free data retrieval call binding the contract method 0x4666a62b.
//
// Solidity: function getCurrentEpochStatus() view returns(uint8)
func (_DKG *DKGCaller) GetCurrentEpochStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getCurrentEpochStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCurrentEpochStatus is a free data retrieval call binding the contract method 0x4666a62b.
//
// Solidity: function getCurrentEpochStatus() view returns(uint8)
func (_DKG *DKGSession) GetCurrentEpochStatus() (uint8, error) {
	return _DKG.Contract.GetCurrentEpochStatus(&_DKG.CallOpts)
}

// GetCurrentEpochStatus is a free data retrieval call binding the contract method 0x4666a62b.
//
// Solidity: function getCurrentEpochStatus() view returns(uint8)
func (_DKG *DKGCallerSession) GetCurrentEpochStatus() (uint8, error) {
	return _DKG.Contract.GetCurrentEpochStatus(&_DKG.CallOpts)
}

// GetDKGEpochInfo is a free data retrieval call binding the contract method 0xca158348.
//
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns((uint256,uint256,address,uint8))
func (_DKG *DKGCaller) GetDKGEpochInfo(opts *bind.CallOpts, _epochId *big.Int) (IDKGDKGEpochInfo, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getDKGEpochInfo", _epochId)

	if err != nil {
		return *new(IDKGDKGEpochInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDKGDKGEpochInfo)).(*IDKGDKGEpochInfo)

	return out0, err

}

// GetDKGEpochInfo is a free data retrieval call binding the contract method 0xca158348.
//
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns((uint256,uint256,address,uint8))
func (_DKG *DKGSession) GetDKGEpochInfo(_epochId *big.Int) (IDKGDKGEpochInfo, error) {
	return _DKG.Contract.GetDKGEpochInfo(&_DKG.CallOpts, _epochId)
}

// GetDKGEpochInfo is a free data retrieval call binding the contract method 0xca158348.
//
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns((uint256,uint256,address,uint8))
func (_DKG *DKGCallerSession) GetDKGEpochInfo(_epochId *big.Int) (IDKGDKGEpochInfo, error) {
	return _DKG.Contract.GetDKGEpochInfo(&_DKG.CallOpts, _epochId)
}

// GetDKGPeriodEndTime is a free data retrieval call binding the contract method 0xc0474753.
//
// Solidity: function getDKGPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCaller) GetDKGPeriodEndTime(opts *bind.CallOpts, _epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getDKGPeriodEndTime", _epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDKGPeriodEndTime is a free data retrieval call binding the contract method 0xc0474753.
//
// Solidity: function getDKGPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGSession) GetDKGPeriodEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetDKGPeriodEndTime(&_DKG.CallOpts, _epochId)
}

// GetDKGPeriodEndTime is a free data retrieval call binding the contract method 0xc0474753.
//
// Solidity: function getDKGPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCallerSession) GetDKGPeriodEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetDKGPeriodEndTime(&_DKG.CallOpts, _epochId)
}

// GetEpochEndTime is a free data retrieval call binding the contract method 0xd4ee9f8d.
//
// Solidity: function getEpochEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCaller) GetEpochEndTime(opts *bind.CallOpts, _epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getEpochEndTime", _epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochEndTime is a free data retrieval call binding the contract method 0xd4ee9f8d.
//
// Solidity: function getEpochEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGSession) GetEpochEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetEpochEndTime(&_DKG.CallOpts, _epochId)
}

// GetEpochEndTime is a free data retrieval call binding the contract method 0xd4ee9f8d.
//
// Solidity: function getEpochEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCallerSession) GetEpochEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetEpochEndTime(&_DKG.CallOpts, _epochId)
}

// GetEpochStatus is a free data retrieval call binding the contract method 0x15accc92.
//
// Solidity: function getEpochStatus(uint256 _epochId) view returns(uint8)
func (_DKG *DKGCaller) GetEpochStatus(opts *bind.CallOpts, _epochId *big.Int) (uint8, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getEpochStatus", _epochId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetEpochStatus is a free data retrieval call binding the contract method 0x15accc92.
//
// Solidity: function getEpochStatus(uint256 _epochId) view returns(uint8)
func (_DKG *DKGSession) GetEpochStatus(_epochId *big.Int) (uint8, error) {
	return _DKG.Contract.GetEpochStatus(&_DKG.CallOpts, _epochId)
}

// GetEpochStatus is a free data retrieval call binding the contract method 0x15accc92.
//
// Solidity: function getEpochStatus(uint256 _epochId) view returns(uint8)
func (_DKG *DKGCallerSession) GetEpochStatus(_epochId *big.Int) (uint8, error) {
	return _DKG.Contract.GetEpochStatus(&_DKG.CallOpts, _epochId)
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

// GetLastEpochId is a free data retrieval call binding the contract method 0x20201668.
//
// Solidity: function getLastEpochId() view returns(uint256)
func (_DKG *DKGCaller) GetLastEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getLastEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastEpochId is a free data retrieval call binding the contract method 0x20201668.
//
// Solidity: function getLastEpochId() view returns(uint256)
func (_DKG *DKGSession) GetLastEpochId() (*big.Int, error) {
	return _DKG.Contract.GetLastEpochId(&_DKG.CallOpts)
}

// GetLastEpochId is a free data retrieval call binding the contract method 0x20201668.
//
// Solidity: function getLastEpochId() view returns(uint256)
func (_DKG *DKGCallerSession) GetLastEpochId() (*big.Int, error) {
	return _DKG.Contract.GetLastEpochId(&_DKG.CallOpts)
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

// GetSignerVotesCount is a free data retrieval call binding the contract method 0xf7897c07.
//
// Solidity: function getSignerVotesCount(uint256 _epochId, address _signerAddr) view returns(uint256)
func (_DKG *DKGCaller) GetSignerVotesCount(opts *bind.CallOpts, _epochId *big.Int, _signerAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getSignerVotesCount", _epochId, _signerAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignerVotesCount is a free data retrieval call binding the contract method 0xf7897c07.
//
// Solidity: function getSignerVotesCount(uint256 _epochId, address _signerAddr) view returns(uint256)
func (_DKG *DKGSession) GetSignerVotesCount(_epochId *big.Int, _signerAddr common.Address) (*big.Int, error) {
	return _DKG.Contract.GetSignerVotesCount(&_DKG.CallOpts, _epochId, _signerAddr)
}

// GetSignerVotesCount is a free data retrieval call binding the contract method 0xf7897c07.
//
// Solidity: function getSignerVotesCount(uint256 _epochId, address _signerAddr) view returns(uint256)
func (_DKG *DKGCallerSession) GetSignerVotesCount(_epochId *big.Int, _signerAddr common.Address) (*big.Int, error) {
	return _DKG.Contract.GetSignerVotesCount(&_DKG.CallOpts, _epochId, _signerAddr)
}

// GetUpdatedCollectionPeriodEndTime is a free data retrieval call binding the contract method 0x36503f42.
//
// Solidity: function getUpdatedCollectionPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCaller) GetUpdatedCollectionPeriodEndTime(opts *bind.CallOpts, _epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getUpdatedCollectionPeriodEndTime", _epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpdatedCollectionPeriodEndTime is a free data retrieval call binding the contract method 0x36503f42.
//
// Solidity: function getUpdatedCollectionPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGSession) GetUpdatedCollectionPeriodEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetUpdatedCollectionPeriodEndTime(&_DKG.CallOpts, _epochId)
}

// GetUpdatedCollectionPeriodEndTime is a free data retrieval call binding the contract method 0x36503f42.
//
// Solidity: function getUpdatedCollectionPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCallerSession) GetUpdatedCollectionPeriodEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetUpdatedCollectionPeriodEndTime(&_DKG.CallOpts, _epochId)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _validator) view returns((address,(uint256,uint256),(uint256,uint256)))
func (_DKG *DKGCaller) GetValidatorInfo(opts *bind.CallOpts, _validator common.Address) (IDKGValidatorInfo, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidatorInfo", _validator)

	if err != nil {
		return *new(IDKGValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDKGValidatorInfo)).(*IDKGValidatorInfo)

	return out0, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _validator) view returns((address,(uint256,uint256),(uint256,uint256)))
func (_DKG *DKGSession) GetValidatorInfo(_validator common.Address) (IDKGValidatorInfo, error) {
	return _DKG.Contract.GetValidatorInfo(&_DKG.CallOpts, _validator)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _validator) view returns((address,(uint256,uint256),(uint256,uint256)))
func (_DKG *DKGCallerSession) GetValidatorInfo(_validator common.Address) (IDKGValidatorInfo, error) {
	return _DKG.Contract.GetValidatorInfo(&_DKG.CallOpts, _validator)
}

// GetValidatorVote is a free data retrieval call binding the contract method 0x124e3ffa.
//
// Solidity: function getValidatorVote(uint256 _epochId, address _validatorAddr) view returns(address)
func (_DKG *DKGCaller) GetValidatorVote(opts *bind.CallOpts, _epochId *big.Int, _validatorAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getValidatorVote", _epochId, _validatorAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetValidatorVote is a free data retrieval call binding the contract method 0x124e3ffa.
//
// Solidity: function getValidatorVote(uint256 _epochId, address _validatorAddr) view returns(address)
func (_DKG *DKGSession) GetValidatorVote(_epochId *big.Int, _validatorAddr common.Address) (common.Address, error) {
	return _DKG.Contract.GetValidatorVote(&_DKG.CallOpts, _epochId, _validatorAddr)
}

// GetValidatorVote is a free data retrieval call binding the contract method 0x124e3ffa.
//
// Solidity: function getValidatorVote(uint256 _epochId, address _validatorAddr) view returns(address)
func (_DKG *DKGCallerSession) GetValidatorVote(_epochId *big.Int, _validatorAddr common.Address) (common.Address, error) {
	return _DKG.Contract.GetValidatorVote(&_DKG.CallOpts, _epochId, _validatorAddr)
}

// GuaranteedWorkingEpochDuration is a free data retrieval call binding the contract method 0x696f2959.
//
// Solidity: function guaranteedWorkingEpochDuration() view returns(uint256)
func (_DKG *DKGCaller) GuaranteedWorkingEpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "guaranteedWorkingEpochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedWorkingEpochDuration is a free data retrieval call binding the contract method 0x696f2959.
//
// Solidity: function guaranteedWorkingEpochDuration() view returns(uint256)
func (_DKG *DKGSession) GuaranteedWorkingEpochDuration() (*big.Int, error) {
	return _DKG.Contract.GuaranteedWorkingEpochDuration(&_DKG.CallOpts)
}

// GuaranteedWorkingEpochDuration is a free data retrieval call binding the contract method 0x696f2959.
//
// Solidity: function guaranteedWorkingEpochDuration() view returns(uint256)
func (_DKG *DKGCallerSession) GuaranteedWorkingEpochDuration() (*big.Int, error) {
	return _DKG.Contract.GuaranteedWorkingEpochDuration(&_DKG.CallOpts)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddr) view returns(bool)
func (_DKG *DKGCaller) IsActiveValidator(opts *bind.CallOpts, _validatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isActiveValidator", _validatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddr) view returns(bool)
func (_DKG *DKGSession) IsActiveValidator(_validatorAddr common.Address) (bool, error) {
	return _DKG.Contract.IsActiveValidator(&_DKG.CallOpts, _validatorAddr)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address _validatorAddr) view returns(bool)
func (_DKG *DKGCallerSession) IsActiveValidator(_validatorAddr common.Address) (bool, error) {
	return _DKG.Contract.IsActiveValidator(&_DKG.CallOpts, _validatorAddr)
}

// IsCurrentEpoch is a free data retrieval call binding the contract method 0xe80f6534.
//
// Solidity: function isCurrentEpoch(uint256 _epochId) view returns(bool)
func (_DKG *DKGCaller) IsCurrentEpoch(opts *bind.CallOpts, _epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isCurrentEpoch", _epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentEpoch is a free data retrieval call binding the contract method 0xe80f6534.
//
// Solidity: function isCurrentEpoch(uint256 _epochId) view returns(bool)
func (_DKG *DKGSession) IsCurrentEpoch(_epochId *big.Int) (bool, error) {
	return _DKG.Contract.IsCurrentEpoch(&_DKG.CallOpts, _epochId)
}

// IsCurrentEpoch is a free data retrieval call binding the contract method 0xe80f6534.
//
// Solidity: function isCurrentEpoch(uint256 _epochId) view returns(bool)
func (_DKG *DKGCallerSession) IsCurrentEpoch(_epochId *big.Int) (bool, error) {
	return _DKG.Contract.IsCurrentEpoch(&_DKG.CallOpts, _epochId)
}

// IsDKGGenSuccessful is a free data retrieval call binding the contract method 0x3ae0725a.
//
// Solidity: function isDKGGenSuccessful(uint256 _epochId) view returns(bool)
func (_DKG *DKGCaller) IsDKGGenSuccessful(opts *bind.CallOpts, _epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isDKGGenSuccessful", _epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDKGGenSuccessful is a free data retrieval call binding the contract method 0x3ae0725a.
//
// Solidity: function isDKGGenSuccessful(uint256 _epochId) view returns(bool)
func (_DKG *DKGSession) IsDKGGenSuccessful(_epochId *big.Int) (bool, error) {
	return _DKG.Contract.IsDKGGenSuccessful(&_DKG.CallOpts, _epochId)
}

// IsDKGGenSuccessful is a free data retrieval call binding the contract method 0x3ae0725a.
//
// Solidity: function isDKGGenSuccessful(uint256 _epochId) view returns(bool)
func (_DKG *DKGCallerSession) IsDKGGenSuccessful(_epochId *big.Int) (bool, error) {
	return _DKG.Contract.IsDKGGenSuccessful(&_DKG.CallOpts, _epochId)
}

// IsLastEpoch is a free data retrieval call binding the contract method 0x214e6a56.
//
// Solidity: function isLastEpoch(uint256 _epochId) view returns(bool)
func (_DKG *DKGCaller) IsLastEpoch(opts *bind.CallOpts, _epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isLastEpoch", _epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLastEpoch is a free data retrieval call binding the contract method 0x214e6a56.
//
// Solidity: function isLastEpoch(uint256 _epochId) view returns(bool)
func (_DKG *DKGSession) IsLastEpoch(_epochId *big.Int) (bool, error) {
	return _DKG.Contract.IsLastEpoch(&_DKG.CallOpts, _epochId)
}

// IsLastEpoch is a free data retrieval call binding the contract method 0x214e6a56.
//
// Solidity: function isLastEpoch(uint256 _epochId) view returns(bool)
func (_DKG *DKGCallerSession) IsLastEpoch(_epochId *big.Int) (bool, error) {
	return _DKG.Contract.IsLastEpoch(&_DKG.CallOpts, _epochId)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validatorAddr) view returns(bool)
func (_DKG *DKGCaller) IsValidator(opts *bind.CallOpts, _validatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isValidator", _validatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validatorAddr) view returns(bool)
func (_DKG *DKGSession) IsValidator(_validatorAddr common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, _validatorAddr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _validatorAddr) view returns(bool)
func (_DKG *DKGCallerSession) IsValidator(_validatorAddr common.Address) (bool, error) {
	return _DKG.Contract.IsValidator(&_DKG.CallOpts, _validatorAddr)
}

// UpdatesCollectionEpochDuration is a free data retrieval call binding the contract method 0x72ef10b2.
//
// Solidity: function updatesCollectionEpochDuration() view returns(uint256)
func (_DKG *DKGCaller) UpdatesCollectionEpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "updatesCollectionEpochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdatesCollectionEpochDuration is a free data retrieval call binding the contract method 0x72ef10b2.
//
// Solidity: function updatesCollectionEpochDuration() view returns(uint256)
func (_DKG *DKGSession) UpdatesCollectionEpochDuration() (*big.Int, error) {
	return _DKG.Contract.UpdatesCollectionEpochDuration(&_DKG.CallOpts)
}

// UpdatesCollectionEpochDuration is a free data retrieval call binding the contract method 0x72ef10b2.
//
// Solidity: function updatesCollectionEpochDuration() view returns(uint256)
func (_DKG *DKGCallerSession) UpdatesCollectionEpochDuration() (*big.Int, error) {
	return _DKG.Contract.UpdatesCollectionEpochDuration(&_DKG.CallOpts)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validatorToAdd) returns()
func (_DKG *DKGTransactor) AddValidator(opts *bind.TransactOpts, _validatorToAdd common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "addValidator", _validatorToAdd)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validatorToAdd) returns()
func (_DKG *DKGSession) AddValidator(_validatorToAdd common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AddValidator(&_DKG.TransactOpts, _validatorToAdd)
}

// AddValidator is a paid mutator transaction binding the contract method 0x4d238c8e.
//
// Solidity: function addValidator(address _validatorToAdd) returns()
func (_DKG *DKGTransactorSession) AddValidator(_validatorToAdd common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AddValidator(&_DKG.TransactOpts, _validatorToAdd)
}

// AnnounceValidatorExit is a paid mutator transaction binding the contract method 0xf9cf0dd8.
//
// Solidity: function announceValidatorExit(address _validatorToExit) returns(uint256)
func (_DKG *DKGTransactor) AnnounceValidatorExit(opts *bind.TransactOpts, _validatorToExit common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "announceValidatorExit", _validatorToExit)
}

// AnnounceValidatorExit is a paid mutator transaction binding the contract method 0xf9cf0dd8.
//
// Solidity: function announceValidatorExit(address _validatorToExit) returns(uint256)
func (_DKG *DKGSession) AnnounceValidatorExit(_validatorToExit common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AnnounceValidatorExit(&_DKG.TransactOpts, _validatorToExit)
}

// AnnounceValidatorExit is a paid mutator transaction binding the contract method 0xf9cf0dd8.
//
// Solidity: function announceValidatorExit(address _validatorToExit) returns(uint256)
func (_DKG *DKGTransactorSession) AnnounceValidatorExit(_validatorToExit common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AnnounceValidatorExit(&_DKG.TransactOpts, _validatorToExit)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _updatesCollectionEpochDuration, uint256 _dkgGenerationEpochDuration, uint256 _guaranteedWorkingEpochDuration) returns()
func (_DKG *DKGTransactor) Initialize(opts *bind.TransactOpts, _updatesCollectionEpochDuration *big.Int, _dkgGenerationEpochDuration *big.Int, _guaranteedWorkingEpochDuration *big.Int) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "initialize", _updatesCollectionEpochDuration, _dkgGenerationEpochDuration, _guaranteedWorkingEpochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _updatesCollectionEpochDuration, uint256 _dkgGenerationEpochDuration, uint256 _guaranteedWorkingEpochDuration) returns()
func (_DKG *DKGSession) Initialize(_updatesCollectionEpochDuration *big.Int, _dkgGenerationEpochDuration *big.Int, _guaranteedWorkingEpochDuration *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _updatesCollectionEpochDuration, _dkgGenerationEpochDuration, _guaranteedWorkingEpochDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _updatesCollectionEpochDuration, uint256 _dkgGenerationEpochDuration, uint256 _guaranteedWorkingEpochDuration) returns()
func (_DKG *DKGTransactorSession) Initialize(_updatesCollectionEpochDuration *big.Int, _dkgGenerationEpochDuration *big.Int, _guaranteedWorkingEpochDuration *big.Int) (*types.Transaction, error) {
	return _DKG.Contract.Initialize(&_DKG.TransactOpts, _updatesCollectionEpochDuration, _dkgGenerationEpochDuration, _guaranteedWorkingEpochDuration)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorToRemove) returns()
func (_DKG *DKGTransactor) RemoveValidator(opts *bind.TransactOpts, _validatorToRemove common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "removeValidator", _validatorToRemove)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorToRemove) returns()
func (_DKG *DKGSession) RemoveValidator(_validatorToRemove common.Address) (*types.Transaction, error) {
	return _DKG.Contract.RemoveValidator(&_DKG.TransactOpts, _validatorToRemove)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _validatorToRemove) returns()
func (_DKG *DKGTransactorSession) RemoveValidator(_validatorToRemove common.Address) (*types.Transaction, error) {
	return _DKG.Contract.RemoveValidator(&_DKG.TransactOpts, _validatorToRemove)
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

// UpdateAllValidators is a paid mutator transaction binding the contract method 0x6c8fbb64.
//
// Solidity: function updateAllValidators() returns()
func (_DKG *DKGTransactor) UpdateAllValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "updateAllValidators")
}

// UpdateAllValidators is a paid mutator transaction binding the contract method 0x6c8fbb64.
//
// Solidity: function updateAllValidators() returns()
func (_DKG *DKGSession) UpdateAllValidators() (*types.Transaction, error) {
	return _DKG.Contract.UpdateAllValidators(&_DKG.TransactOpts)
}

// UpdateAllValidators is a paid mutator transaction binding the contract method 0x6c8fbb64.
//
// Solidity: function updateAllValidators() returns()
func (_DKG *DKGTransactorSession) UpdateAllValidators() (*types.Transaction, error) {
	return _DKG.Contract.UpdateAllValidators(&_DKG.TransactOpts)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _epochId, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGTransactor) VoteSigner(opts *bind.TransactOpts, _epochId *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "voteSigner", _epochId, _signerAddress, _signature)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _epochId, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGSession) VoteSigner(_epochId *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.Contract.VoteSigner(&_DKG.TransactOpts, _epochId, _signerAddress, _signature)
}

// VoteSigner is a paid mutator transaction binding the contract method 0xc88bc067.
//
// Solidity: function voteSigner(uint256 _epochId, address _signerAddress, bytes _signature) returns()
func (_DKG *DKGTransactorSession) VoteSigner(_epochId *big.Int, _signerAddress common.Address, _signature []byte) (*types.Transaction, error) {
	return _DKG.Contract.VoteSigner(&_DKG.TransactOpts, _epochId, _signerAddress, _signature)
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

// DKGNewEpochCreatedIterator is returned from FilterNewEpochCreated and is used to iterate over the raw logs and unpacked data for NewEpochCreated events raised by the DKG contract.
type DKGNewEpochCreatedIterator struct {
	Event *DKGNewEpochCreated // Event containing the contract specifics and raw log

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
func (it *DKGNewEpochCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGNewEpochCreated)
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
		it.Event = new(DKGNewEpochCreated)
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
func (it *DKGNewEpochCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGNewEpochCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGNewEpochCreated represents a NewEpochCreated event raised by the DKG contract.
type DKGNewEpochCreated struct {
	EpochId        *big.Int
	EpochStartTime *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewEpochCreated is a free log retrieval operation binding the contract event 0xf27e0ff471920cfae653a7a16cf64784b13fafb16b991f9af7860027c1b475d7.
//
// Solidity: event NewEpochCreated(uint256 epochId, uint256 epochStartTime)
func (_DKG *DKGFilterer) FilterNewEpochCreated(opts *bind.FilterOpts) (*DKGNewEpochCreatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "NewEpochCreated")
	if err != nil {
		return nil, err
	}
	return &DKGNewEpochCreatedIterator{contract: _DKG.contract, event: "NewEpochCreated", logs: logs, sub: sub}, nil
}

// WatchNewEpochCreated is a free log subscription operation binding the contract event 0xf27e0ff471920cfae653a7a16cf64784b13fafb16b991f9af7860027c1b475d7.
//
// Solidity: event NewEpochCreated(uint256 epochId, uint256 epochStartTime)
func (_DKG *DKGFilterer) WatchNewEpochCreated(opts *bind.WatchOpts, sink chan<- *DKGNewEpochCreated) (event.Subscription, error) {

	logs, sub, err := _DKG.contract.WatchLogs(opts, "NewEpochCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGNewEpochCreated)
				if err := _DKG.contract.UnpackLog(event, "NewEpochCreated", log); err != nil {
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

// ParseNewEpochCreated is a log parse operation binding the contract event 0xf27e0ff471920cfae653a7a16cf64784b13fafb16b991f9af7860027c1b475d7.
//
// Solidity: event NewEpochCreated(uint256 epochId, uint256 epochStartTime)
func (_DKG *DKGFilterer) ParseNewEpochCreated(log types.Log) (*DKGNewEpochCreated, error) {
	event := new(DKGNewEpochCreated)
	if err := _DKG.contract.UnpackLog(event, "NewEpochCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGNewValidatorAddedIterator is returned from FilterNewValidatorAdded and is used to iterate over the raw logs and unpacked data for NewValidatorAdded events raised by the DKG contract.
type DKGNewValidatorAddedIterator struct {
	Event *DKGNewValidatorAdded // Event containing the contract specifics and raw log

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
func (it *DKGNewValidatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGNewValidatorAdded)
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
		it.Event = new(DKGNewValidatorAdded)
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
func (it *DKGNewValidatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGNewValidatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGNewValidatorAdded represents a NewValidatorAdded event raised by the DKG contract.
type DKGNewValidatorAdded struct {
	ValidatorAddr        common.Address
	StartValidationTime  *big.Int
	StartValidationEpoch *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewValidatorAdded is a free log retrieval operation binding the contract event 0xcd38d7f8bcd459fe58ab4e4f5ba2ee51418a1a0cf207a6747f61bb7945e62b41.
//
// Solidity: event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationTime, uint256 startValidationEpoch)
func (_DKG *DKGFilterer) FilterNewValidatorAdded(opts *bind.FilterOpts, validatorAddr []common.Address) (*DKGNewValidatorAddedIterator, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "NewValidatorAdded", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DKGNewValidatorAddedIterator{contract: _DKG.contract, event: "NewValidatorAdded", logs: logs, sub: sub}, nil
}

// WatchNewValidatorAdded is a free log subscription operation binding the contract event 0xcd38d7f8bcd459fe58ab4e4f5ba2ee51418a1a0cf207a6747f61bb7945e62b41.
//
// Solidity: event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationTime, uint256 startValidationEpoch)
func (_DKG *DKGFilterer) WatchNewValidatorAdded(opts *bind.WatchOpts, sink chan<- *DKGNewValidatorAdded, validatorAddr []common.Address) (event.Subscription, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "NewValidatorAdded", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGNewValidatorAdded)
				if err := _DKG.contract.UnpackLog(event, "NewValidatorAdded", log); err != nil {
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

// ParseNewValidatorAdded is a log parse operation binding the contract event 0xcd38d7f8bcd459fe58ab4e4f5ba2ee51418a1a0cf207a6747f61bb7945e62b41.
//
// Solidity: event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationTime, uint256 startValidationEpoch)
func (_DKG *DKGFilterer) ParseNewValidatorAdded(log types.Log) (*DKGNewValidatorAdded, error) {
	event := new(DKGNewValidatorAdded)
	if err := _DKG.contract.UnpackLog(event, "NewValidatorAdded", log); err != nil {
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
	EpochId       *big.Int
	SignerAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignerAddressUpdated is a free log retrieval operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 epochId, address signerAddress)
func (_DKG *DKGFilterer) FilterSignerAddressUpdated(opts *bind.FilterOpts) (*DKGSignerAddressUpdatedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "SignerAddressUpdated")
	if err != nil {
		return nil, err
	}
	return &DKGSignerAddressUpdatedIterator{contract: _DKG.contract, event: "SignerAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchSignerAddressUpdated is a free log subscription operation binding the contract event 0xa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46.
//
// Solidity: event SignerAddressUpdated(uint256 epochId, address signerAddress)
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
// Solidity: event SignerAddressUpdated(uint256 epochId, address signerAddress)
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
	EpochId          *big.Int
	Validator        common.Address
	CollectiveSigner common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSignerVoted is a free log retrieval operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 epochId, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) FilterSignerVoted(opts *bind.FilterOpts) (*DKGSignerVotedIterator, error) {

	logs, sub, err := _DKG.contract.FilterLogs(opts, "SignerVoted")
	if err != nil {
		return nil, err
	}
	return &DKGSignerVotedIterator{contract: _DKG.contract, event: "SignerVoted", logs: logs, sub: sub}, nil
}

// WatchSignerVoted is a free log subscription operation binding the contract event 0x4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec19.
//
// Solidity: event SignerVoted(uint256 epochId, address validator, address collectiveSigner)
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
// Solidity: event SignerVoted(uint256 epochId, address validator, address collectiveSigner)
func (_DKG *DKGFilterer) ParseSignerVoted(log types.Log) (*DKGSignerVoted, error) {
	event := new(DKGSignerVoted)
	if err := _DKG.contract.UnpackLog(event, "SignerVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGValidatorExitAnnouncedIterator is returned from FilterValidatorExitAnnounced and is used to iterate over the raw logs and unpacked data for ValidatorExitAnnounced events raised by the DKG contract.
type DKGValidatorExitAnnouncedIterator struct {
	Event *DKGValidatorExitAnnounced // Event containing the contract specifics and raw log

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
func (it *DKGValidatorExitAnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGValidatorExitAnnounced)
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
		it.Event = new(DKGValidatorExitAnnounced)
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
func (it *DKGValidatorExitAnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGValidatorExitAnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGValidatorExitAnnounced represents a ValidatorExitAnnounced event raised by the DKG contract.
type DKGValidatorExitAnnounced struct {
	ValidatorAddr      common.Address
	EndValidationTime  *big.Int
	EndValidationEpoch *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorExitAnnounced is a free log retrieval operation binding the contract event 0xd75af7adb28dda89e53e18c225f03e4aee3f0132f19eaff0a642b1f88433baad.
//
// Solidity: event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationTime, uint256 endValidationEpoch)
func (_DKG *DKGFilterer) FilterValidatorExitAnnounced(opts *bind.FilterOpts, validatorAddr []common.Address) (*DKGValidatorExitAnnouncedIterator, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ValidatorExitAnnounced", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DKGValidatorExitAnnouncedIterator{contract: _DKG.contract, event: "ValidatorExitAnnounced", logs: logs, sub: sub}, nil
}

// WatchValidatorExitAnnounced is a free log subscription operation binding the contract event 0xd75af7adb28dda89e53e18c225f03e4aee3f0132f19eaff0a642b1f88433baad.
//
// Solidity: event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationTime, uint256 endValidationEpoch)
func (_DKG *DKGFilterer) WatchValidatorExitAnnounced(opts *bind.WatchOpts, sink chan<- *DKGValidatorExitAnnounced, validatorAddr []common.Address) (event.Subscription, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ValidatorExitAnnounced", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGValidatorExitAnnounced)
				if err := _DKG.contract.UnpackLog(event, "ValidatorExitAnnounced", log); err != nil {
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

// ParseValidatorExitAnnounced is a log parse operation binding the contract event 0xd75af7adb28dda89e53e18c225f03e4aee3f0132f19eaff0a642b1f88433baad.
//
// Solidity: event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationTime, uint256 endValidationEpoch)
func (_DKG *DKGFilterer) ParseValidatorExitAnnounced(log types.Log) (*DKGValidatorExitAnnounced, error) {
	event := new(DKGValidatorExitAnnounced)
	if err := _DKG.contract.UnpackLog(event, "ValidatorExitAnnounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGValidatorRemovedIterator is returned from FilterValidatorRemoved and is used to iterate over the raw logs and unpacked data for ValidatorRemoved events raised by the DKG contract.
type DKGValidatorRemovedIterator struct {
	Event *DKGValidatorRemoved // Event containing the contract specifics and raw log

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
func (it *DKGValidatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGValidatorRemoved)
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
		it.Event = new(DKGValidatorRemoved)
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
func (it *DKGValidatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGValidatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGValidatorRemoved represents a ValidatorRemoved event raised by the DKG contract.
type DKGValidatorRemoved struct {
	ValidatorAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorRemoved is a free log retrieval operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validatorAddr)
func (_DKG *DKGFilterer) FilterValidatorRemoved(opts *bind.FilterOpts, validatorAddr []common.Address) (*DKGValidatorRemovedIterator, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ValidatorRemoved", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DKGValidatorRemovedIterator{contract: _DKG.contract, event: "ValidatorRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorRemoved is a free log subscription operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validatorAddr)
func (_DKG *DKGFilterer) WatchValidatorRemoved(opts *bind.WatchOpts, sink chan<- *DKGValidatorRemoved, validatorAddr []common.Address) (event.Subscription, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ValidatorRemoved", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGValidatorRemoved)
				if err := _DKG.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
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

// ParseValidatorRemoved is a log parse operation binding the contract event 0xe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f1.
//
// Solidity: event ValidatorRemoved(address indexed validatorAddr)
func (_DKG *DKGFilterer) ParseValidatorRemoved(log types.Log) (*DKGValidatorRemoved, error) {
	event := new(DKGValidatorRemoved)
	if err := _DKG.contract.UnpackLog(event, "ValidatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
