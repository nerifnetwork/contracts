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
	EpochId               *big.Int
	EpochStartTime        *big.Int
	DkgGenPeriodStartTime *big.Int
	EpochSigner           common.Address
	EpochStatus           uint8
}

// IDKGValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type IDKGValidatorInfo struct {
	Validator              common.Address
	StartValidationEpochId *big.Int
	EndValidationEpochId   *big.Int
}

// DKGMetaData contains all meta data concerning the DKG contract.
var DKGMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"}],\"name\":\"NewEpochCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startValidationEpoch\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectiveSigner\",\"type\":\"address\"}],\"name\":\"SignerVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endValidationEpoch\",\"type\":\"uint256\"}],\"name\":\"ValidatorExitAnnounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIRST_EPOCH_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorToAdd\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"announceValidatorExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_endValidationEpoch\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dkgGenerationEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveEpochStatus\",\"outputs\":[{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_activeValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getDKGEpochInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dkgGenPeriodStartTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"epochSigner\",\"type\":\"address\"},{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"epochStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIDKG.DKGEpochInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getDKGPeriodEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getEpochStatus\",\"outputs\":[{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getGuaranteedWorkingPeriodEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddr\",\"type\":\"address\"}],\"name\":\"getSignerVotesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startValidationEpochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endValidationEpochId\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"getValidatorVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guaranteedWorkingEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updatesCollectionEpochDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dkgGenerationEpochDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_guaranteedWorkingEpochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"slashValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatesCollectionEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"voteSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061214b806100206000396000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80637c34a5251161010f578063c88bc067116100a2578063f7897c0711610071578063f7897c0714610412578063f9cf0dd81461044c578063facd743b1461045f578063fd33434b1461047257600080fd5b8063c88bc067146103ce578063ca158348146103e1578063f072ec0814610401578063f3513a371461040a57600080fd5b80639de70258116100de5780639de7025814610372578063ac4fe5ab14610387578063c0474753146103b3578063c3125405146103c657600080fd5b80637c34a5251461030657806380d859111461030e5780638a11d7c9146103215780638cb941cc1461035f57600080fd5b80633e3b5b1911610187578063691304511161015657806369130451146102d9578063696f2959146102ec5780636c8fbb64146102f557806372ef10b2146102fd57600080fd5b80633e3b5b191461027957806340550a1c1461028e5780634711f586146102b15780634d238c8e146102c657600080fd5b80631a296e02116101c35780631a296e021461025057806331147c301461026157806337e4e780146102695780633e15d5be1461027157600080fd5b8063124e3ffa146101ea57806315accc921461021a578063180fd87f1461023a575b600080fd5b6101fd6101f8366004611cec565b610485565b6040516001600160a01b0390911681526020015b60405180910390f35b61022d610228366004611d1c565b6104b7565b6040516102119190611d6d565b610242610560565b604051908152602001610211565b6006546001600160a01b03166101fd565b61024261058c565b600554610242565b610242600181565b6000805160206120f6833981519152546101fd565b6102a161029c366004611d7b565b6105dc565b6040519015158152602001610211565b6102c46102bf366004611d7b565b610645565b005b6102c46102d4366004611d7b565b610729565b6102c46102e7366004611e3b565b610821565b61024260045481565b6102c4610952565b61024260025481565b61022d6109e8565b6102c461031c366004611e8b565b6109fa565b61033461032f366004611d7b565b610b1d565b6040805182516001600160a01b03168152602080840151908201529181015190820152606001610211565b6102c461036d366004611d7b565b610b85565b61037a610ba6565b6040516102119190611eb7565b6102a1610395366004611d7b565b6001600160a01b03166000908152600b602052604090205460ff1690565b6102426103c1366004611d1c565b610c16565b610242610ca7565b6102c46103dc366004611f04565b610cb3565b6103f46103ef366004611d1c565b610f79565b6040516102119190611f5d565b61024260035481565b61037a610fed565b610242610420366004611cec565b60008281526009602090815260408083206001600160a01b038516845260030190915290205492915050565b61024261045a366004611d7b565b610ff9565b6102a161046d366004611d7b565b61114b565b610242610480366004611d1c565b611158565b60008281526009602090815260408083206001600160a01b038086168552600490910190925290912054165b92915050565b60006005548211156104cb57506000919050565b60055482146104dc57506005919050565b60008281526009602052604090206001015460045442906104fd9083611fb5565b9150811061050e5750600192915050565b60008381526009602052604090206002015480156105565742811115610538575060039392505050565b42600354826105479190611fb5565b10610556575060049392505050565b5060029392505050565b600061056a611178565b6105726111df565b505060055460009081526009602052604090206002015490565b6000806105996007611308565b905060005b818110156105d7576105b461029c600783611312565b156105c757826105c381611fc8565b9350505b6105d081611fc8565b905061059e565b505090565b6001600160a01b0381166000818152600a60209081526040808320815180830183528154815260019091015481840152938352600b90915281205490919060ff1615801561062d5750600554815111155b801561063e57506005548160200151115b9392505050565b61064d611178565b6001600160a01b0381166000908152600b602052604090205460ff16156106c65760405162461bcd60e51b815260206004820152602260248201527f444b473a2056616c696461746f722068617320616c726561647920736c617368604482015261195960f21b60648201526084015b60405180910390fd5b6001600160a01b0381166000908152600b60205260409020805460ff191660011790556106f28161131e565b6040516001600160a01b038216907f1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda90600090a250565b610731611374565b61073a8161114b565b156107875760405162461bcd60e51b815260206004820152601d60248201527f444b473a2056616c696461746f7220616c72656164792065786973747300000060448201526064016106bd565b60006107916111df565b60408051808201825282815260001960208083019182526001600160a01b0387166000908152600a9091529290922090518155905160019091015590506107d96007836113d4565b50816001600160a01b03167f80bd2303d5db0c7f61ca096867fe20736ecab88d458da842d6aedd5170f4a5618260405161081591815260200190565b60405180910390a25050565b6108296113e9565b6000829050806001600160a01b0316638e68dce46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561086c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108909190611fe1565b600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663fb9d9ac56040518163ffffffff1660e01b8152600401602060405180830381865afa1580156108f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109189190611fe1565b600180546001600160a01b0319166001600160a01b039290921691909117905550336000805160206120f6833981519152555050565b5050565b600061095e600761146d565b905060005b815181101561094e57600554600a600084848151811061098557610985611ffe565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060010154116109d8576109d88282815181106109cb576109cb611ffe565b602002602001015161131e565b6109e181611fc8565b9050610963565b60006109f56005546104b7565b905090565b600054610100900460ff1615808015610a1a5750600054600160ff909116105b80610a345750303b158015610a34575060005460ff166001145b610a975760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016106bd565b6000805460ff191660011790558015610aba576000805461ff0019166101001790555b600284905560038390556004829055610ad161147a565b8015610b17576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610b4a604051806060016040528060006001600160a01b0316815260200160008152602001600081525090565b506001600160a01b03166000818152600a60209081526040918290208251606081018452938452805491840191909152600101549082015290565b610b8d6113e9565b610ba3816000805160206120f683398151915255565b50565b60606000610bb26114e5565b90506000610bc06007611308565b905060005b81811015610c05576000610bda600783611312565b9050610be5816105dc565b15610bf457610bf4848261150c565b50610bfe81611fc8565b9050610bc5565b50610c0f82611521565b9250505090565b6000818152600960205260408120600201548103610c885760405162461bcd60e51b815260206004820152602960248201527f444b473a20444b472067656e65726174696f6e20706572696f6420646f6573206044820152681b9bdd08195e1a5cdd60ba1b60648201526084016106bd565b6003546000838152600960205260409020600201546104b19190611fb5565b60006109f56007611308565b610cbb61152f565b6004610cc6846104b7565b6005811115610cd757610cd7611d35565b14610d245760405162461bcd60e51b815260206004820181905260248201527f444b473a204e6f74206120444b472067656e65726174696f6e20706572696f6460448201526064016106bd565b6000610d5782610d516040518060400160405280600681526020016576657269667960d01b815250611584565b906115bf565b9050826001600160a01b0316816001600160a01b031614610dba5760405162461bcd60e51b815260206004820152601960248201527f444b473a205369676e617475726520697320696e76616c69640000000000000060448201526064016106bd565b600084815260096020908152604080832033845260048101909252909120546001600160a01b031615610e245760405162461bcd60e51b81526020600482015260126024820152711112d1ce88105b1c9958591e481d9bdd195960721b60448201526064016106bd565b6001600160a01b0384166000908152600382016020526040812080548290610e4b90611fc8565b918290555033600081815260048501602090815260409182902080546001600160a01b0319166001600160a01b038b1690811790915582518b815291820193909352908101919091529091507f4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec199060600160405180910390a16002610ece61058c565b610ed89190612014565b81118015610ef457506006546001600160a01b03868116911614155b15610f715781546001600160a01b0386166001600160a01b03199182168117845560068054909216179055610f2761147a565b610f2f610952565b604080518781526001600160a01b03871660208201527fa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46910160405180910390a15b505050505050565b610f81611c8f565b600082815260096020908152604091829020825160a08101845285815260018201549281019290925260028101549282019290925281546001600160a01b0316606082015260808101610fd3856104b7565b6005811115610fe457610fe4611d35565b90529392505050565b60606109f5600761146d565b6000611003611374565b61100c826105dc565b6110585760405162461bcd60e51b815260206004820152601c60248201527f444b473a2056616c696461746f72206973206e6f74206163746976650000000060448201526064016106bd565b6001600160a01b0382166000908152600a6020526040902060010154600019146110e25760405162461bcd60e51b815260206004820152603560248201527f444b473a2045786974206f66207468652076616c696461746f722068617320616044820152741b1c9958591e481899595b88185b9b9bdd5b98d959605a1b60648201526084016106bd565b6110ea6111df565b6001600160a01b0383166000818152600a60205260409081902060010183905551919250907f30ac6b6a6a27291e9b642749b60a3e096ce91197734d2a31e39d755a750abf809061113e9084815260200190565b60405180910390a2919050565b60006104b16007836115e3565b60045460008281526009602052604081206001015490916104b191611fb5565b6001546001600160a01b031633146111dd5760405162461bcd60e51b815260206004820152602260248201527f444b473a204e6f74206120736c617368696e6720766f74696e67206164647265604482015261737360f01b60648201526084016106bd565b565b6000806111ea6109e8565b9050600481600581111561120057611200611d35565b036112735760405162461bcd60e51b815260206004820152603f60248201527f444b473a20556e61626c6520746f20636f6c6c6563742075706461746573206460448201527f7572696e672074686520444b472067656e65726174696f6e20706572696f640060648201526084016106bd565b600554600081815260096020526040902060020154158061129b57504261129982610c16565b105b156112ef57600060018360058111156112b6576112b6611d35565b036112c9576112c482611158565b6112cb565b425b9050600254816112db9190611fb5565b600083815260096020526040902060020155505b600181146104b15761130081611fc8565b905080610c0f565b60006104b1825490565b600061063e8383611605565b61132960078261162f565b506001600160a01b0381166000818152600a6020526040808220828155600101829055517fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f19190a250565b6000546201000090046001600160a01b031633146111dd5760405162461bcd60e51b815260206004820152601a60248201527f444b473a204e6f742061207374616b696e67206164647265737300000000000060448201526064016106bd565b600061063e836001600160a01b038416611644565b60006114016000805160206120f68339815191525490565b90506001600160a01b038116158061142157506001600160a01b03811633145b610ba35760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f7200000000000060448201526064016106bd565b6060600061063e83611693565b600060056000815461148b90611fc8565b91829055506000818152600960209081526040918290204260019091018190558251848152918201529192507ff27e0ff471920cfae653a7a16cf64784b13fafb16b991f9af7860027c1b475d7910160405180910390a150565b60408051606081018252600060208201818152928201529081526115076116ef565b815290565b815161094e906001600160a01b038316611723565b8051602001516060906104b1565b611538336105dc565b6111dd5760405162461bcd60e51b815260206004820152601c60248201527f444b473a204e6f7420616e206163746976652076616c696461746f720000000060448201526064016106bd565b60006115908251611774565b826040516020016115a292919061205a565b604051602081830303815290604052805190602001209050919050565b60008060006115ce8585611807565b915091506115db8161184c565b509392505050565b6001600160a01b0381166000908152600183016020526040812054151561063e565b600082600001828154811061161c5761161c611ffe565b9060005260206000200154905092915050565b600061063e836001600160a01b038416611996565b600081815260018301602052604081205461168b575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556104b1565b5060006104b1565b6060816000018054806020026020016040519081016040528092919081815260200182805480156116e357602002820191906000526020600020905b8154815260200190600101908083116116cf575b50505050509050919050565b60408051808201825260008082526020820152815160a0810190925290611717816001611a90565b60058252602082015290565b600061173183602001515190565b8351909150611741826001611fb5565b0361175e57825161175e9084906117599060026120b5565b611ab1565b6020928301516001820181529083020190910152565b6060600061178183611af3565b600101905060008167ffffffffffffffff8111156117a1576117a1611d98565b6040519080825280601f01601f1916602001820160405280156117cb576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a85049450846117d557509392505050565b600080825160410361183d5760208301516040840151606085015160001a61183187828585611bcb565b94509450505050611845565b506000905060025b9250929050565b600081600481111561186057611860611d35565b036118685750565b600181600481111561187c5761187c611d35565b036118c95760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016106bd565b60028160048111156118dd576118dd611d35565b0361192a5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016106bd565b600381600481111561193e5761193e611d35565b03610ba35760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016106bd565b60008181526001830160205260408120548015611a7f5760006119ba6001836120cc565b85549091506000906119ce906001906120cc565b9050818114611a335760008660000182815481106119ee576119ee611ffe565b9060005260206000200154905080876000018481548110611a1157611a11611ffe565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611a4457611a446120df565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506104b1565b60009150506104b1565b5092915050565b60005b60208202811015611aac57600083820152602001611a93565b505050565b604080516020830281019091526020830151805160005b602080830201811015611ae5578281015184820152602001611ac8565b505050908252602090910152565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310611b325772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611b5e576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310611b7c57662386f26fc10000830492506010015b6305f5e1008310611b94576305f5e100830492506008015b6127108310611ba857612710830492506004015b60648310611bba576064830492506002015b600a83106104b15760010192915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611c025750600090506003611c86565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611c56573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611c7f57600060019250925050611c86565b9150600090505b94509492505050565b6040518060a0016040528060008152602001600081526020016000815260200160006001600160a01b0316815260200160006005811115611cd257611cd2611d35565b905290565b6001600160a01b0381168114610ba357600080fd5b60008060408385031215611cff57600080fd5b823591506020830135611d1181611cd7565b809150509250929050565b600060208284031215611d2e57600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b60068110611d6957634e487b7160e01b600052602160045260246000fd5b9052565b602081016104b18284611d4b565b600060208284031215611d8d57600080fd5b813561063e81611cd7565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611dbf57600080fd5b813567ffffffffffffffff80821115611dda57611dda611d98565b604051601f8301601f19908116603f01168101908282118183101715611e0257611e02611d98565b81604052838152866020858801011115611e1b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215611e4e57600080fd5b8235611e5981611cd7565b9150602083013567ffffffffffffffff811115611e7557600080fd5b611e8185828601611dae565b9150509250929050565b600080600060608486031215611ea057600080fd5b505081359360208301359350604090920135919050565b6020808252825182820181905260009190848201906040850190845b81811015611ef85783516001600160a01b031683529284019291840191600101611ed3565b50909695505050505050565b600080600060608486031215611f1957600080fd5b833592506020840135611f2b81611cd7565b9150604084013567ffffffffffffffff811115611f4757600080fd5b611f5386828701611dae565b9150509250925092565b8151815260208083015190820152604080830151908201526060808301516001600160a01b03169082015260808083015160a0830191611a8990840182611d4b565b634e487b7160e01b600052601160045260246000fd5b808201808211156104b1576104b1611f9f565b600060018201611fda57611fda611f9f565b5060010190565b600060208284031215611ff357600080fd5b815161063e81611cd7565b634e487b7160e01b600052603260045260246000fd5b60008261203157634e487b7160e01b600052601260045260246000fd5b500490565b60005b83811015612051578181015183820152602001612039565b50506000910152565b7f19457468657265756d205369676e6564204d6573736167653a0a00000000000081526000835161209281601a850160208801612036565b8351908301906120a981601a840160208801612036565b01601a01949350505050565b80820281158282048414176104b1576104b1611f9f565b818103818111156104b1576104b1611f9f565b634e487b7160e01b600052603160045260246000fdfe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a2646970667358221220dc0a12842164ae11b38687323563c88f5828afd3b492ed7241848322f073793664736f6c63430008120033",
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

// FIRSTEPOCHID is a free data retrieval call binding the contract method 0x3e15d5be.
//
// Solidity: function FIRST_EPOCH_ID() view returns(uint256)
func (_DKG *DKGCaller) FIRSTEPOCHID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "FIRST_EPOCH_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIRSTEPOCHID is a free data retrieval call binding the contract method 0x3e15d5be.
//
// Solidity: function FIRST_EPOCH_ID() view returns(uint256)
func (_DKG *DKGSession) FIRSTEPOCHID() (*big.Int, error) {
	return _DKG.Contract.FIRSTEPOCHID(&_DKG.CallOpts)
}

// FIRSTEPOCHID is a free data retrieval call binding the contract method 0x3e15d5be.
//
// Solidity: function FIRST_EPOCH_ID() view returns(uint256)
func (_DKG *DKGCallerSession) FIRSTEPOCHID() (*big.Int, error) {
	return _DKG.Contract.FIRSTEPOCHID(&_DKG.CallOpts)
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

// GetActiveEpochId is a free data retrieval call binding the contract method 0x37e4e780.
//
// Solidity: function getActiveEpochId() view returns(uint256)
func (_DKG *DKGCaller) GetActiveEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getActiveEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveEpochId is a free data retrieval call binding the contract method 0x37e4e780.
//
// Solidity: function getActiveEpochId() view returns(uint256)
func (_DKG *DKGSession) GetActiveEpochId() (*big.Int, error) {
	return _DKG.Contract.GetActiveEpochId(&_DKG.CallOpts)
}

// GetActiveEpochId is a free data retrieval call binding the contract method 0x37e4e780.
//
// Solidity: function getActiveEpochId() view returns(uint256)
func (_DKG *DKGCallerSession) GetActiveEpochId() (*big.Int, error) {
	return _DKG.Contract.GetActiveEpochId(&_DKG.CallOpts)
}

// GetActiveEpochStatus is a free data retrieval call binding the contract method 0x7c34a525.
//
// Solidity: function getActiveEpochStatus() view returns(uint8)
func (_DKG *DKGCaller) GetActiveEpochStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getActiveEpochStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetActiveEpochStatus is a free data retrieval call binding the contract method 0x7c34a525.
//
// Solidity: function getActiveEpochStatus() view returns(uint8)
func (_DKG *DKGSession) GetActiveEpochStatus() (uint8, error) {
	return _DKG.Contract.GetActiveEpochStatus(&_DKG.CallOpts)
}

// GetActiveEpochStatus is a free data retrieval call binding the contract method 0x7c34a525.
//
// Solidity: function getActiveEpochStatus() view returns(uint8)
func (_DKG *DKGCallerSession) GetActiveEpochStatus() (uint8, error) {
	return _DKG.Contract.GetActiveEpochStatus(&_DKG.CallOpts)
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

// GetDKGEpochInfo is a free data retrieval call binding the contract method 0xca158348.
//
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns((uint256,uint256,uint256,address,uint8))
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
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns((uint256,uint256,uint256,address,uint8))
func (_DKG *DKGSession) GetDKGEpochInfo(_epochId *big.Int) (IDKGDKGEpochInfo, error) {
	return _DKG.Contract.GetDKGEpochInfo(&_DKG.CallOpts, _epochId)
}

// GetDKGEpochInfo is a free data retrieval call binding the contract method 0xca158348.
//
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns((uint256,uint256,uint256,address,uint8))
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

// GetGuaranteedWorkingPeriodEndTime is a free data retrieval call binding the contract method 0xfd33434b.
//
// Solidity: function getGuaranteedWorkingPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCaller) GetGuaranteedWorkingPeriodEndTime(opts *bind.CallOpts, _epochId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getGuaranteedWorkingPeriodEndTime", _epochId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGuaranteedWorkingPeriodEndTime is a free data retrieval call binding the contract method 0xfd33434b.
//
// Solidity: function getGuaranteedWorkingPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGSession) GetGuaranteedWorkingPeriodEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetGuaranteedWorkingPeriodEndTime(&_DKG.CallOpts, _epochId)
}

// GetGuaranteedWorkingPeriodEndTime is a free data retrieval call binding the contract method 0xfd33434b.
//
// Solidity: function getGuaranteedWorkingPeriodEndTime(uint256 _epochId) view returns(uint256)
func (_DKG *DKGCallerSession) GetGuaranteedWorkingPeriodEndTime(_epochId *big.Int) (*big.Int, error) {
	return _DKG.Contract.GetGuaranteedWorkingPeriodEndTime(&_DKG.CallOpts, _epochId)
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

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _validator) view returns((address,uint256,uint256))
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
// Solidity: function getValidatorInfo(address _validator) view returns((address,uint256,uint256))
func (_DKG *DKGSession) GetValidatorInfo(_validator common.Address) (IDKGValidatorInfo, error) {
	return _DKG.Contract.GetValidatorInfo(&_DKG.CallOpts, _validator)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address _validator) view returns((address,uint256,uint256))
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

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validatorAddr) view returns(bool)
func (_DKG *DKGCaller) IsValidatorSlashed(opts *bind.CallOpts, _validatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "isValidatorSlashed", _validatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validatorAddr) view returns(bool)
func (_DKG *DKGSession) IsValidatorSlashed(_validatorAddr common.Address) (bool, error) {
	return _DKG.Contract.IsValidatorSlashed(&_DKG.CallOpts, _validatorAddr)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0xac4fe5ab.
//
// Solidity: function isValidatorSlashed(address _validatorAddr) view returns(bool)
func (_DKG *DKGCallerSession) IsValidatorSlashed(_validatorAddr common.Address) (bool, error) {
	return _DKG.Contract.IsValidatorSlashed(&_DKG.CallOpts, _validatorAddr)
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
// Solidity: function announceValidatorExit(address _validatorAddr) returns(uint256 _endValidationEpoch)
func (_DKG *DKGTransactor) AnnounceValidatorExit(opts *bind.TransactOpts, _validatorAddr common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "announceValidatorExit", _validatorAddr)
}

// AnnounceValidatorExit is a paid mutator transaction binding the contract method 0xf9cf0dd8.
//
// Solidity: function announceValidatorExit(address _validatorAddr) returns(uint256 _endValidationEpoch)
func (_DKG *DKGSession) AnnounceValidatorExit(_validatorAddr common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AnnounceValidatorExit(&_DKG.TransactOpts, _validatorAddr)
}

// AnnounceValidatorExit is a paid mutator transaction binding the contract method 0xf9cf0dd8.
//
// Solidity: function announceValidatorExit(address _validatorAddr) returns(uint256 _endValidationEpoch)
func (_DKG *DKGTransactorSession) AnnounceValidatorExit(_validatorAddr common.Address) (*types.Transaction, error) {
	return _DKG.Contract.AnnounceValidatorExit(&_DKG.TransactOpts, _validatorAddr)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x180fd87f.
//
// Solidity: function createProposal() returns(uint256)
func (_DKG *DKGTransactor) CreateProposal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "createProposal")
}

// CreateProposal is a paid mutator transaction binding the contract method 0x180fd87f.
//
// Solidity: function createProposal() returns(uint256)
func (_DKG *DKGSession) CreateProposal() (*types.Transaction, error) {
	return _DKG.Contract.CreateProposal(&_DKG.TransactOpts)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x180fd87f.
//
// Solidity: function createProposal() returns(uint256)
func (_DKG *DKGTransactorSession) CreateProposal() (*types.Transaction, error) {
	return _DKG.Contract.CreateProposal(&_DKG.TransactOpts)
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

// SlashValidator is a paid mutator transaction binding the contract method 0x4711f586.
//
// Solidity: function slashValidator(address _validatorAddr) returns()
func (_DKG *DKGTransactor) SlashValidator(opts *bind.TransactOpts, _validatorAddr common.Address) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "slashValidator", _validatorAddr)
}

// SlashValidator is a paid mutator transaction binding the contract method 0x4711f586.
//
// Solidity: function slashValidator(address _validatorAddr) returns()
func (_DKG *DKGSession) SlashValidator(_validatorAddr common.Address) (*types.Transaction, error) {
	return _DKG.Contract.SlashValidator(&_DKG.TransactOpts, _validatorAddr)
}

// SlashValidator is a paid mutator transaction binding the contract method 0x4711f586.
//
// Solidity: function slashValidator(address _validatorAddr) returns()
func (_DKG *DKGTransactorSession) SlashValidator(_validatorAddr common.Address) (*types.Transaction, error) {
	return _DKG.Contract.SlashValidator(&_DKG.TransactOpts, _validatorAddr)
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
	StartValidationEpoch *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewValidatorAdded is a free log retrieval operation binding the contract event 0x80bd2303d5db0c7f61ca096867fe20736ecab88d458da842d6aedd5170f4a561.
//
// Solidity: event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationEpoch)
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

// WatchNewValidatorAdded is a free log subscription operation binding the contract event 0x80bd2303d5db0c7f61ca096867fe20736ecab88d458da842d6aedd5170f4a561.
//
// Solidity: event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationEpoch)
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

// ParseNewValidatorAdded is a log parse operation binding the contract event 0x80bd2303d5db0c7f61ca096867fe20736ecab88d458da842d6aedd5170f4a561.
//
// Solidity: event NewValidatorAdded(address indexed validatorAddr, uint256 startValidationEpoch)
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
	EndValidationEpoch *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterValidatorExitAnnounced is a free log retrieval operation binding the contract event 0x30ac6b6a6a27291e9b642749b60a3e096ce91197734d2a31e39d755a750abf80.
//
// Solidity: event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationEpoch)
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

// WatchValidatorExitAnnounced is a free log subscription operation binding the contract event 0x30ac6b6a6a27291e9b642749b60a3e096ce91197734d2a31e39d755a750abf80.
//
// Solidity: event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationEpoch)
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

// ParseValidatorExitAnnounced is a log parse operation binding the contract event 0x30ac6b6a6a27291e9b642749b60a3e096ce91197734d2a31e39d755a750abf80.
//
// Solidity: event ValidatorExitAnnounced(address indexed validatorAddr, uint256 endValidationEpoch)
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

// DKGValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the DKG contract.
type DKGValidatorSlashedIterator struct {
	Event *DKGValidatorSlashed // Event containing the contract specifics and raw log

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
func (it *DKGValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGValidatorSlashed)
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
		it.Event = new(DKGValidatorSlashed)
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
func (it *DKGValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGValidatorSlashed represents a ValidatorSlashed event raised by the DKG contract.
type DKGValidatorSlashed struct {
	ValidatorAddr common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validatorAddr)
func (_DKG *DKGFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validatorAddr []common.Address) (*DKGValidatorSlashedIterator, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.FilterLogs(opts, "ValidatorSlashed", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &DKGValidatorSlashedIterator{contract: _DKG.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validatorAddr)
func (_DKG *DKGFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *DKGValidatorSlashed, validatorAddr []common.Address) (event.Subscription, error) {

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _DKG.contract.WatchLogs(opts, "ValidatorSlashed", validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGValidatorSlashed)
				if err := _DKG.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
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

// ParseValidatorSlashed is a log parse operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validatorAddr)
func (_DKG *DKGFilterer) ParseValidatorSlashed(log types.Log) (*DKGValidatorSlashed, error) {
	event := new(DKGValidatorSlashed)
	if err := _DKG.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
