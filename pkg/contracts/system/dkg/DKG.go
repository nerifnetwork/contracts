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
	MainEpochInfo IDKGMainEpochInfo
	EpochSigner   common.Address
	EpochStatus   uint8
}

// IDKGMainEpochInfo is an auto generated low-level Go binding around an user-defined struct.
type IDKGMainEpochInfo struct {
	EpochId             *big.Int
	EpochStartTime      *big.Int
	DkgGenPeriodEndTime *big.Int
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"}],\"name\":\"NewEpochCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startValidationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startValidationEpoch\",\"type\":\"uint256\"}],\"name\":\"NewValidatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signerAddress\",\"type\":\"address\"}],\"name\":\"SignerAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectiveSigner\",\"type\":\"address\"}],\"name\":\"SignerVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endValidationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endValidationEpoch\",\"type\":\"uint256\"}],\"name\":\"ValidatorExitAnnounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"ValidatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FIRST_EPOCH_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorToAdd\",\"type\":\"address\"}],\"name\":\"addValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorToExit\",\"type\":\"address\"}],\"name\":\"announceValidatorExit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dkgGenPeriodEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.MainEpochInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dkgGenerationEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_activeValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentEpochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEpochStatus\",\"outputs\":[{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getDKGEpochInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dkgGenPeriodEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.MainEpochInfo\",\"name\":\"mainEpochInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"epochSigner\",\"type\":\"address\"},{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"epochStatus\",\"type\":\"uint8\"}],\"internalType\":\"structIDKG.DKGEpochInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getDKGPeriodEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getEpochEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getEpochStatus\",\"outputs\":[{\"internalType\":\"enumIDKG.DKGEpochStatuses\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getMainEpochInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStartTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dkgGenPeriodEndTime\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.MainEpochInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddr\",\"type\":\"address\"}],\"name\":\"getSignerVotesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"getUpdatedCollectionPeriodEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"validationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.ValidationData\",\"name\":\"startValidationData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"validationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validationEpoch\",\"type\":\"uint256\"}],\"internalType\":\"structIDKG.ValidationData\",\"name\":\"endValidationData\",\"type\":\"tuple\"}],\"internalType\":\"structIDKG.ValidatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"getValidatorVote\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guaranteedWorkingEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_updatesCollectionEpochDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dkgGenerationEpochDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_guaranteedWorkingEpochDuration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"isCurrentEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"isDKGGenSuccessful\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"}],\"name\":\"isLastEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorToRemove\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorAddr\",\"type\":\"address\"}],\"name\":\"slashValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAllValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updatesCollectionEpochDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"voteSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612570806100206000396000f3fe608060405234801561001057600080fd5b50600436106102325760003560e01c806372ef10b211610130578063c3125405116100b8578063f072ec081161007c578063f072ec08146104e0578063f3513a37146104e9578063f7897c07146104f1578063f9cf0dd81461052b578063facd743b1461053e57600080fd5b8063c31254051461047f578063c88bc06714610487578063ca1583481461049a578063d4ee9f8d146104ba578063e80f6534146104cd57600080fd5b80639de70258116100ff5780639de7025814610410578063a29a839f14610425578063ac4fe5ab1461042d578063b2d5258c14610459578063c04747531461046c57600080fd5b806372ef10b2146103c157806380d85911146103ca5780638a11d7c9146103dd5780638cb941cc146103fd57600080fd5b80633e15d5be116101be5780634711f586116101825780634711f586146103775780634d238c8e1461038a578063691304511461039d578063696f2959146103b05780636c8fbb64146103b957600080fd5b80633e15d5be1461032a5780633e3b5b191461033257806340550a1c1461034757806340a141ff1461035a5780634666a62b1461036f57600080fd5b8063202016681161020557806320201668146102ad578063214e6a56146102bf57806331147c30146102e457806336503f42146102ec5780633ae0725a146102ff57600080fd5b8063124e3ffa1461023757806315accc9214610267578063180fd87f146102875780631a296e021461029c575b600080fd5b61024a610245366004612098565b610551565b6040516001600160a01b0390911681526020015b60405180910390f35b61027a6102753660046120c8565b610583565b60405161025e9190612119565b61028f61064e565b60405161025e9190612127565b6006546001600160a01b031661024a565b6005545b60405190815260200161025e565b6102d46102cd3660046120c8565b6005541490565b604051901515815260200161025e565b6102b161068b565b6102b16102fa3660046120c8565b6106db565b6102d461030d3660046120c8565b6000908152600960205260409020546001600160a01b0316151590565b6102b1600181565b60008051602061251b8339815191525461024a565b6102d4610355366004612148565b6106fb565b61036d610368366004612148565b610812565b005b61027a6108f9565b61036d610385366004612148565b610906565b61036d610398366004612148565b6109ce565b61036d6103ab366004612208565b610aec565b6102b160045481565b61036d610c1d565b6102b160025481565b61036d6103d8366004612258565b610cf8565b6103f06103eb366004612148565b610e13565b60405161025e9190612284565b61036d61040b366004612148565b610e7c565b610418610e9a565b60405161025e91906122c7565b6102b1610f13565b6102d461043b366004612148565b6001600160a01b03166000908152600b602052604090205460ff1690565b61028f6104673660046120c8565b610f3f565b6102b161047a3660046120c8565b610fa1565b6102b1610fb9565b61036d610495366004612314565b610fc5565b6104ad6104a83660046120c8565b611276565b60405161025e919061236d565b6102b16104c83660046120c8565b6112da565b6102d46104db3660046120c8565b6112e8565b6102b160035481565b6104186112fa565b6102b16104ff366004612098565b60008281526009602090815260408083206001600160a01b038516845260020190915290205492915050565b6102b1610539366004612148565b611306565b6102d461054c366004612148565b61146a565b60008281526009602090815260408083206001600160a01b038086168552600390910190925290912054165b92915050565b600060055482111561059757506000919050565b600082815260096020526040902060010154600554831480156105c15750428111806105c1575080155b156105cf5750600192915050565b6105d8836112e8565b6105e55750600692915050565b42600254826105f491906123c3565b915081106106055750600292915050565b426003548261061491906123c3565b915081106106255750600392915050565b426004548261063491906123c3565b915081106106455750600492915050565b50600592915050565b61067260405180606001604052806000815260200160008152602001600081525090565b61067a611477565b6106856104676114de565b90505b90565b60008061069860076115d1565b905060005b818110156106d6576106b36103556007836115db565b156106c657826106c2816123d6565b9350505b6106cf816123d6565b905061069d565b505090565b600254600082815260096020526040812060010154909161057d916123c3565b6001600160a01b0381166000908152600a60209081526040808320815160808101835281548184019081526001808401546060840152908252835180850190945260028301548452600390920154838501528084019290925281519092015190918391148061078c575081516020015161078c906000908152600960205260409020546001600160a01b0316151590565b801561079a57508151514210155b905060006107c78360200151602001516000908152600960205260409020546001600160a01b0316151590565b15806107d7575060208301515142105b6001600160a01b0386166000908152600b602052604090205490915060ff161580156108005750815b80156108095750805b95945050505050565b61081a6115ee565b61082560078261164e565b6108765760405162461bcd60e51b815260206004820152601d60248201527f444b473a2056616c696461746f7220646f6573206e6f7420657869737400000060448201526064015b60405180910390fd5b6001600160a01b0381166000908152600a60205260409020600201544210156108ed5760405162461bcd60e51b815260206004820152602360248201527f444b473a2056616c696461746f722063616e27742062652072656d6f766564206044820152621e595d60ea1b606482015260840161086d565b6108f681611670565b50565b6000610685610275610f13565b61090e611477565b6001600160a01b0381166000908152600b602052604090205460ff16156109825760405162461bcd60e51b815260206004820152602260248201527f444b473a2056616c696461746f722068617320616c726561647920736c617368604482015261195960f21b606482015260840161086d565b6001600160a01b0381166000818152600b6020526040808220805460ff19166001179055517f1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda9190a250565b6109d66115ee565b6109df8161146a565b15610a2c5760405162461bcd60e51b815260206004820152601d60248201527f444b473a2056616c696461746f7220616c726561647920657869737473000000604482015260640161086d565b6000610a366116d4565b604080518082018252828152815180830183526000198152600060208281018290528084019283526001600160a01b0388168252600a815293902091518051835583015160018301555180516002830155909101516003909101559050610a9e600783611739565b50805160208083015160408051938452918301526001600160a01b038416917fcd38d7f8bcd459fe58ab4e4f5ba2ee51418a1a0cf207a6747f61bb7945e62b41910160405180910390a25050565b610af461174e565b6000829050806001600160a01b0316638e68dce46040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5b91906123ef565b600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663fb9d9ac56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610bbf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610be391906123ef565b600180546001600160a01b0319166001600160a01b0392909216919091179055503360008051602061251b833981519152555050565b5050565b6000610c2960076117d2565b905060005b8151811015610c19576000600a6000848481518110610c4f57610c4f61240c565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000209050610c9b838381518110610c8e57610c8e61240c565b60200260200101516106fb565b15610cb157610cac816002016117df565b610cba565b610cba816117df565b60028101544210610ce757610ce7838381518110610cda57610cda61240c565b6020026020010151611670565b50610cf1816123d6565b9050610c2e565b600054610100900460ff1615808015610d185750600054600160ff909116105b80610d325750303b158015610d32575060005460ff166001145b610d955760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161086d565b6000805460ff191660011790558015610db8576000805461ff0019166101001790555b6002849055600383905560048290558015610e0d576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610e1b611fe1565b506001600160a01b03166000818152600a602090815260409182902082516060810184529384528251808401845281548152600182015481840152848301528251808401845260028201548152600390910154918101919091529082015290565b610e8461174e565b6108f68160008051602061251b83398151915255565b60606000610ea6611832565b90506000610eb460076115d1565b905060005b81811015610ef9576000610ece6007836115db565b9050610ed9816106fb565b15610ee857610ee88482611859565b50610ef2816123d6565b9050610eb9565b50610f0c8280516020015160609061057d565b9250505090565b6005546000818152600960205260409020600101544210156106885780610f3981612422565b91505090565b610f6360405180606001604052806000815260200160008152602001600081525090565b604051806060016040528083815260200160096000858152602001908152602001600020600101548152602001610f9984610fa1565b905292915050565b6000600354610faf836106db565b61057d91906123c3565b600061068560076115d1565b610fcd61186e565b6003610fd884610583565b6006811115610fe957610fe96120e1565b146110365760405162461bcd60e51b815260206004820181905260248201527f444b473a204e6f74206120444b472067656e65726174696f6e20706572696f64604482015260640161086d565b6000611069826110636040518060400160405280600681526020016576657269667960d01b8152506118c3565b906118fe565b9050826001600160a01b0316816001600160a01b0316146110cc5760405162461bcd60e51b815260206004820152601960248201527f444b473a205369676e617475726520697320696e76616c696400000000000000604482015260640161086d565b600084815260096020908152604080832033845260038101909252909120546001600160a01b0316156111365760405162461bcd60e51b81526020600482015260126024820152711112d1ce88105b1c9958591e481d9bdd195960721b604482015260640161086d565b6001600160a01b038416600090815260028201602052604081208054829061115d906123d6565b918290555033600081815260038501602090815260409182902080546001600160a01b0319166001600160a01b038b1690811790915582518b815291820193909352908101919091529091507f4686ba7a5df3cb2d9979ee16ec58c7ea0b92273d836278488fb07c732cf8ec199060600160405180910390a160026111e061068b565b6111ea9190612439565b8111801561120657506006546001600160a01b03868116911614155b1561126e5781546001600160a01b0386166001600160a01b0319918216811784556006805490921681179091556040805188815260208101929092527fa6d3de2b2ccbce3cf736f5f8e515fdad82235d0d1f1ff8a0ad14f98c48af0a46910160405180910390a15b505050505050565b61127e612040565b6000828152600960205260409081902081516060810190925290806112a285610f3f565b815282546001600160a01b031660208201526040016112c085610583565b60068111156112d1576112d16120e1565b90529392505050565b6000600454610faf83610fa1565b6000816112f3610f13565b1492915050565b606061068560076117d2565b60006113106115ee565b611319826106fb565b6113655760405162461bcd60e51b815260206004820152601c60248201527f444b473a2056616c696461746f72206973206e6f742061637469766500000000604482015260640161086d565b6001600160a01b0382166000908152600a6020526040902060020154600019146113ef5760405162461bcd60e51b815260206004820152603560248201527f444b473a2045786974206f66207468652076616c696461746f722068617320616044820152741b1c9958591e481899595b88185b9b9bdd5b98d959605a1b606482015260840161086d565b60006113f96116d4565b6001600160a01b0384166000818152600a602090815260409182902084516002820181905585830151600390920182905583519081529182015292935090917fd75af7adb28dda89e53e18c225f03e4aee3f0132f19eaff0a642b1f88433baad910160405180910390a25192915050565b600061057d60078361164e565b6001546001600160a01b031633146114dc5760405162461bcd60e51b815260206004820152602260248201527f444b473a204e6f74206120736c617368696e6720766f74696e67206164647265604482015261737360f01b606482015260840161086d565b565b6000806114e9610f13565b915081905060006114f982610583565b9050600281600681111561150f5761150f6120e1565b146106d65760055492508183036106d657426005826006811115611535576115356120e1565b1415801561155557506001826006811115611552576115526120e1565b14155b1561156657611563846112da565b90505b600560008154611575906123d6565b918290555060008181526009602090815260409182902060010184905581518381529081018490529195507ff27e0ff471920cfae653a7a16cf64784b13fafb16b991f9af7860027c1b475d7910160405180910390a150505090565b600061057d825490565b60006115e78383611922565b9392505050565b6000546201000090046001600160a01b031633146114dc5760405162461bcd60e51b815260206004820152601a60248201527f444b473a204e6f742061207374616b696e672061646472657373000000000000604482015260640161086d565b6001600160a01b038116600090815260018301602052604081205415156115e7565b61167b60078261194c565b506001600160a01b0381166000818152600a60205260408082208281556001810183905560028101839055600301829055517fe1434e25d6611e0db941968fdc97811c982ac1602e951637d206f5fdda9dd8f19190a250565b604080518082019091526000808252602082015260055460006116f56114de565b90506001811461171d5761170881611961565b9250818360200151146106d6576106d6610c1d565b6040518060400160405280428152602001828152509250505090565b60006115e7836001600160a01b038416611996565b600061176660008051602061251b8339815191525490565b90506001600160a01b038116158061178657506001600160a01b03811633145b6108f65760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f72000000000000604482015260640161086d565b606060006115e7836119e5565b8054421080159061180a575060018101546000908152600960205260409020546001600160a01b0316155b156108f657600061182161181c6114de565b611961565b805183556020015160018301555050565b6040805160608101825260006020820181815292820152908152611854611a41565b815290565b8151610c19906001600160a01b038316611a75565b611877336106fb565b6114dc5760405162461bcd60e51b815260206004820152601c60248201527f444b473a204e6f7420616e206163746976652076616c696461746f7200000000604482015260640161086d565b60006118cf8251611ac6565b826040516020016118e192919061247f565b604051602081830303815290604052805190602001209050919050565b600080600061190d8585611b59565b9150915061191a81611b9e565b509392505050565b60008260000182815481106119395761193961240c565b9060005260206000200154905092915050565b60006115e7836001600160a01b038416611ce8565b6040805180820190915260008082526020820152604051806040016040528061198984610fa1565b8152602001929092525090565b60008181526001830160205260408120546119dd5750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915561057d565b50600061057d565b606081600001805480602002602001604051908101604052809291908181526020018280548015611a3557602002820191906000526020600020905b815481526020019060010190808311611a21575b50505050509050919050565b60408051808201825260008082526020820152815160a0810190925290611a69816001611de2565b60058252602082015290565b6000611a8383602001515190565b8351909150611a938260016123c3565b03611ab0578251611ab0908490611aab9060026124da565b611e03565b6020928301516001820181529083020190910152565b60606000611ad383611e45565b600101905060008167ffffffffffffffff811115611af357611af3612165565b6040519080825280601f01601f191660200182016040528015611b1d576020820181803683370190505b5090508181016020015b600019016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a8504945084611b2757509392505050565b6000808251604103611b8f5760208301516040840151606085015160001a611b8387828585611f1d565b94509450505050611b97565b506000905060025b9250929050565b6000816004811115611bb257611bb26120e1565b03611bba5750565b6001816004811115611bce57611bce6120e1565b03611c1b5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161086d565b6002816004811115611c2f57611c2f6120e1565b03611c7c5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161086d565b6003816004811115611c9057611c906120e1565b036108f65760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161086d565b60008181526001830160205260408120548015611dd1576000611d0c6001836124f1565b8554909150600090611d20906001906124f1565b9050818114611d85576000866000018281548110611d4057611d4061240c565b9060005260206000200154905080876000018481548110611d6357611d6361240c565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611d9657611d96612504565b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505061057d565b600091505061057d565b5092915050565b60005b60208202811015611dfe57600083820152602001611de5565b505050565b604080516020830281019091526020830151805160005b602080830201811015611e37578281015184820152602001611e1a565b505050908252602090910152565b60008072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b8310611e845772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611eb0576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310611ece57662386f26fc10000830492506010015b6305f5e1008310611ee6576305f5e100830492506008015b6127108310611efa57612710830492506004015b60648310611f0c576064830492506002015b600a831061057d5760010192915050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115611f545750600090506003611fd8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611fa8573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116611fd157600060019250925050611fd8565b9150600090505b94509492505050565b604051806060016040528060006001600160a01b03168152602001612019604051806040016040528060008152602001600081525090565b815260200161203b604051806040016040528060008152602001600081525090565b905290565b604051806060016040528061206f60405180606001604052806000815260200160008152602001600081525090565b815260006020820181905260409091015290565b6001600160a01b03811681146108f657600080fd5b600080604083850312156120ab57600080fd5b8235915060208301356120bd81612083565b809150509250929050565b6000602082840312156120da57600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b6007811061211557634e487b7160e01b600052602160045260246000fd5b9052565b6020810161057d82846120f7565b8151815260208083015190820152604080830151908201526060810161057d565b60006020828403121561215a57600080fd5b81356115e781612083565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261218c57600080fd5b813567ffffffffffffffff808211156121a7576121a7612165565b604051601f8301601f19908116603f011681019082821181831017156121cf576121cf612165565b816040528381528660208588010111156121e857600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561221b57600080fd5b823561222681612083565b9150602083013567ffffffffffffffff81111561224257600080fd5b61224e8582860161217b565b9150509250929050565b60008060006060848603121561226d57600080fd5b505081359360208301359350604090920135919050565b81516001600160a01b0316815260208083015180518284015290810151604083015260a08201905060408301518051606084015260208101516080840152611ddb565b6020808252825182820181905260009190848201906040850190845b818110156123085783516001600160a01b0316835292840192918401916001016122e3565b50909695505050505050565b60008060006060848603121561232957600080fd5b83359250602084013561233b81612083565b9150604084013567ffffffffffffffff81111561235757600080fd5b6123638682870161217b565b9150509250925092565b81518051825260208082015181840152604091820151828401528301516001600160a01b0316606083015282015160a0820190611ddb60808401826120f7565b634e487b7160e01b600052601160045260246000fd5b8082018082111561057d5761057d6123ad565b6000600182016123e8576123e86123ad565b5060010190565b60006020828403121561240157600080fd5b81516115e781612083565b634e487b7160e01b600052603260045260246000fd5b600081612431576124316123ad565b506000190190565b60008261245657634e487b7160e01b600052601260045260246000fd5b500490565b60005b8381101561247657818101518382015260200161245e565b50506000910152565b7f19457468657265756d205369676e6564204d6573736167653a0a0000000000008152600083516124b781601a85016020880161245b565b8351908301906124ce81601a84016020880161245b565b01601a01949350505050565b808202811582820484141761057d5761057d6123ad565b8181038181111561057d5761057d6123ad565b634e487b7160e01b600052603160045260246000fdfe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a26469706673582212209590c0c8ff920875d1ec53cf73941f2a67ed5ae3b7239ac28531088fb4795d6464736f6c63430008120033",
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
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns(((uint256,uint256,uint256),address,uint8))
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
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns(((uint256,uint256,uint256),address,uint8))
func (_DKG *DKGSession) GetDKGEpochInfo(_epochId *big.Int) (IDKGDKGEpochInfo, error) {
	return _DKG.Contract.GetDKGEpochInfo(&_DKG.CallOpts, _epochId)
}

// GetDKGEpochInfo is a free data retrieval call binding the contract method 0xca158348.
//
// Solidity: function getDKGEpochInfo(uint256 _epochId) view returns(((uint256,uint256,uint256),address,uint8))
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

// GetMainEpochInfo is a free data retrieval call binding the contract method 0xb2d5258c.
//
// Solidity: function getMainEpochInfo(uint256 _epochId) view returns((uint256,uint256,uint256))
func (_DKG *DKGCaller) GetMainEpochInfo(opts *bind.CallOpts, _epochId *big.Int) (IDKGMainEpochInfo, error) {
	var out []interface{}
	err := _DKG.contract.Call(opts, &out, "getMainEpochInfo", _epochId)

	if err != nil {
		return *new(IDKGMainEpochInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IDKGMainEpochInfo)).(*IDKGMainEpochInfo)

	return out0, err

}

// GetMainEpochInfo is a free data retrieval call binding the contract method 0xb2d5258c.
//
// Solidity: function getMainEpochInfo(uint256 _epochId) view returns((uint256,uint256,uint256))
func (_DKG *DKGSession) GetMainEpochInfo(_epochId *big.Int) (IDKGMainEpochInfo, error) {
	return _DKG.Contract.GetMainEpochInfo(&_DKG.CallOpts, _epochId)
}

// GetMainEpochInfo is a free data retrieval call binding the contract method 0xb2d5258c.
//
// Solidity: function getMainEpochInfo(uint256 _epochId) view returns((uint256,uint256,uint256))
func (_DKG *DKGCallerSession) GetMainEpochInfo(_epochId *big.Int) (IDKGMainEpochInfo, error) {
	return _DKG.Contract.GetMainEpochInfo(&_DKG.CallOpts, _epochId)
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

// CreateProposal is a paid mutator transaction binding the contract method 0x180fd87f.
//
// Solidity: function createProposal() returns((uint256,uint256,uint256))
func (_DKG *DKGTransactor) CreateProposal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKG.contract.Transact(opts, "createProposal")
}

// CreateProposal is a paid mutator transaction binding the contract method 0x180fd87f.
//
// Solidity: function createProposal() returns((uint256,uint256,uint256))
func (_DKG *DKGSession) CreateProposal() (*types.Transaction, error) {
	return _DKG.Contract.CreateProposal(&_DKG.TransactOpts)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x180fd87f.
//
// Solidity: function createProposal() returns((uint256,uint256,uint256))
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
