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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"}],\"name\":\"MinimalStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokensSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokensRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"TokensStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensAmount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"WithdrawalPeriodUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addRewardsToStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"announceWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getValidatorStatus\",\"outputs\":[{\"internalType\":\"enumIStaking.ValidatorStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getValidatorsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"enumIStaking.ValidatorStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structIStaking.ValidatorData\",\"name\":\"validatorData\",\"type\":\"tuple\"}],\"internalType\":\"structIStaking.ValidatorInfo[]\",\"name\":\"_validatorsInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getWithdrawalAnnouncement\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"internalType\":\"structIStaking.WithdrawalAnnouncement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"hasWithdrawalAnnouncement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"}],\"name\":\"setMinimalStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_withdrawalPeriod\",\"type\":\"uint256\"}],\"name\":\"setWithdrawalPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sigExpirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"stakeWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611fcb806100206000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c80638cb941cc116100de578063b7ab4db511610097578063c96be4cb11610071578063c96be4cb1461039a578063ecd9ba82146103ad578063ee602ca4146103c0578063f38403281461043057600080fd5b8063b7ab4db514610369578063ba7e31281461037e578063bca7093d1461039157600080fd5b80638cb941cc146102d8578063973b294f146102eb5780639ec41a2d146102fe578063a310624f14610307578063a694fc3a14610343578063ac4fe5ab1461035657600080fd5b806342ad55ac1161014b578063691304511161012557806369130451146102805780637a1ac61e146102935780637a766460146102a65780638b0e9f3f146102cf57600080fd5b806342ad55ac1461023a57806351ed6a301461024d578063621379811461026057600080fd5b80632749824014610193578063282ec26d146101ae57806335c14de1146101ec5780633ccfd60b146101f65780633d6ec65e146101fe5780633e3b5b1914610211575b600080fd5b61019b610443565b6040519081526020015b60405180910390f35b6101dc6101bc366004611abd565b6001600160a01b03166000908152600b6020526040902060010154151590565b60405190151581526020016101a5565b6101f4610454565b005b6101f4610510565b6101f461020c366004611ada565b610692565b600080516020611f76833981519152545b6040516001600160a01b0390911681526020016101a5565b6101dc610248366004611abd565b6106a3565b600454610222906001600160a01b031681565b61027361026e366004611af3565b6106df565b6040516101a59190611b4d565b6101f461028e366004611bd0565b61081f565b6101f46102a1366004611c94565b6109fa565b61019b6102b4366004611abd565b6001600160a01b03166000908152600a602052604090205490565b61019b60075481565b6101f46102e6366004611abd565b610b33565b6101f46102f9366004611ada565b610b51565b61019b60055481565b610336610315366004611abd565b6001600160a01b03166000908152600a602052604090206001015460ff1690565b6040516101a59190611cc9565b6101f4610351366004611ada565b610b62565b6101dc610364366004611abd565b610b93565b610371610b9c565b6040516101a59190611cd7565b6101f461038c366004611ada565b610ba8565b61019b60065481565b6101f46103a8366004611abd565b610d07565b6101f46103bb366004611d24565b610d6c565b6104156103ce366004611abd565b6040805180820190915260008082526020820152506001600160a01b03166000908152600b6020908152604091829020825180840190935280548352600101549082015290565b604080518251815260209283015192810192909252016101a5565b6101f461043e366004611d73565b610e2d565b600061044f6008610ee7565b905090565b61045d33610b93565b156104835760405162461bcd60e51b815260040161047a90611d9f565b60405180910390fd5b336000908152600b60205260409020600101546104b25760405162461bcd60e51b815260040161047a90611dd6565b336000908152600b6020908152604080832080548482556001909101849055600554600a9093529220546104e7908390611e3f565b101580156104fd57506104fb600833610ef7565b155b1561050d5761050d336001610f1c565b50565b61051933610b93565b156105365760405162461bcd60e51b815260040161047a90611d9f565b336000908152600b60205260409020600101546105655760405162461bcd60e51b815260040161047a90611dd6565b600654336000908152600b6020526040902060010154429161058691611e3f565b11156105e25760405162461bcd60e51b815260206004820152602560248201527f5374616b696e673a207769746864726177616c20706572696f64206e6f742070604482015264185cdcd95960da1b606482015260840161047a565b336000908152600b6020908152604080832054600a909252822080549192839261060d908490611e52565b9250508190555080600760008282546106269190611e52565b9091555050336000818152600b602052604081208181556001015560045461065a916001600160a01b03909116908361107f565b60405181815233907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b9060200160405180910390a250565b61069a6110e7565b61050d816111aa565b600060015b6001600160a01b0383166000908152600a602052604090206001015460ff1660028111156106d8576106d8611b15565b1492915050565b606060006106f76106f06008610ee7565b85856111e6565b90506107038482611e52565b67ffffffffffffffff81111561071b5761071b611bba565b60405190808252806020026020018201604052801561075457816020015b610741611a6e565b8152602001906001900390816107395790505b509150835b8181101561081757600061076e60088361120f565b6040805180820182526001600160a01b0383168082526000908152600a6020908152908390208351808501909452805484526001810154949550919381850193929183019060ff1660028111156107c7576107c7611b15565b60028111156107d8576107d8611b15565b9052509052846107e88885611e52565b815181106107f8576107f8611e65565b602002602001018190525050808061080f90611e7b565b915050610759565b505092915050565b61082761121b565b600082905080600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b03166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610890573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b49190611e94565b600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663fb9d9ac56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610918573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061093c9190611e94565b600260006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663046aa0cb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109c49190611e94565b600380546001600160a01b0319166001600160a01b03929092169190911790555033600080516020611f76833981519152555050565b600054610100900460ff1615808015610a1a5750600054600160ff909116105b80610a345750303b158015610a34575060005460ff166001145b610a975760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840161047a565b6000805460ff191660011790558015610aba576000805461ff0019166101001790555b600480546001600160a01b0319166001600160a01b038616179055610ade836111aa565b610ae78261129f565b8015610b2d576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610b3b61121b565b61050d81600080516020611f7683398151915255565b610b596110e7565b61050d8161129f565b610b6b33610b93565b15610b885760405162461bcd60e51b815260040161047a90611d9f565b61050d3333836112d4565b600060026106a8565b606061044f60086114fc565b610bb133610b93565b15610bce5760405162461bcd60e51b815260040161047a90611d9f565b60008111610c305760405162461bcd60e51b815260206004820152602960248201527f5374616b696e673a20616d6f756e74206d7573742062652067726561746572206044820152687468616e207a65726f60b81b606482015260840161047a565b336000908152600a6020526040902054811115610c9b5760405162461bcd60e51b815260206004820152602360248201527f5374616b696e673a20616d6f756e74206d757374206265203c3d20746f207374604482015262616b6560e81b606482015260840161047a565b604080518082018252828152426020808301918252336000908152600b8252848120935184559151600190930192909255600554600a9092529190912054610ce4908390611e52565b108015610cf75750610cf7600833610ef7565b1561050d5761050d336000610f1c565b6002546001600160a01b03163314610d615760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a206e6f74206120736c617368696e6720766f74696e670000604482015260640161047a565b61050d816002610f1c565b610d7533610b93565b15610d925760405162461bcd60e51b815260040161047a90611d9f565b6004805460405163d505accf60e01b81523392810192909252306024830152604482018790526064820186905260ff8516608483015260a4820184905260c482018390526001600160a01b03169063d505accf9060e401600060405180830381600087803b158015610e0357600080fd5b505af1158015610e17573d6000803e3d6000fd5b50505050610e263333876112d4565b5050505050565b6003546001600160a01b03163314610e9d5760405162461bcd60e51b815260206004820152602d60248201527f5374616b696e673a206f6e6c7920526577617264446973747269627574696f6e60448201526c141bdbdb0818dbdb9d1c9858dd609a1b606482015260840161047a565b6001600160a01b0382166000908152600a602052604081208054839290610ec5908490611e3f565b925050819055508060076000828254610ede9190611e3f565b90915550505050565b6000610ef1825490565b92915050565b6001600160a01b038116600090815260018301602052604081205415155b9392505050565b600080826002811115610f3157610f31611b15565b1480610f4e57506002826002811115610f4c57610f4c611b15565b145b15610f6557610f5e600884611509565b9050610f73565b610f7060088461151e565b90505b80610fd35760405162461bcd60e51b815260206004820152602a60248201527f5374616b696e673a204661696c656420746f207570646174652076616c696461604482015269746f722073746174757360b01b606482015260840161047a565b6001600160a01b0383166000908152600a6020526040902060019081018054849260ff199091169083600281111561100d5761100d611b15565b0217905550600160009054906101000a90046001600160a01b03166001600160a01b031663b32805c36040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561106257600080fd5b505af1158015611076573d6000803e3d6000fd5b50505050505050565b6040516001600160a01b0383166024820152604481018290526110e290849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611533565b505050565b60005460408051637ac3c02f60e01b8152905133926201000090046001600160a01b031691637ac3c02f9160048083019260209291908290030181865afa158015611136573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115a9190611e94565b6001600160a01b0316146111a85760405162461bcd60e51b815260206004820152601560248201527429ba30b5b4b7339d102737ba10309039b4b3b732b960591b604482015260640161047a565b565b60058190556040518181527fb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca906020015b60405180910390a150565b60006111f28284611e3f565b9050838111156111ff5750825b80831115610f1557509092915050565b6000610f158383611608565b6000611233600080516020611f768339815191525490565b90506001600160a01b038116158061125357506001600160a01b03811633145b61050d5760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f72000000000000604482015260640161047a565b60068190556040518181527f4157b30b99fc9ed1b54c707f60552f93b22a74b0859af295ae63994b387006ee906020016111db565b600081116113245760405162461bcd60e51b815260206004820152601a60248201527f5374616b696e673a205a65726f207374616b6520616d6f756e74000000000000604482015260640161047a565b600480546040516370a0823160e01b81526001600160a01b03868116938201939093528392909116906370a0823190602401602060405180830381865afa158015611373573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113979190611eb1565b10156113f15760405162461bcd60e51b815260206004820152602360248201527f5374616b696e673a204e6f7420656e6f75676820746f6b656e7320746f207374604482015262616b6560e81b606482015260840161047a565b6001600160a01b0382166000908152600a6020526040812090600182015460ff16600281111561142357611423611b15565b14801561143e5750600554815461143b908490611e3f565b10155b1561145d576001818101805460ff19168217905561145d903390611632565b600454611475906001600160a01b0316853085611708565b818160000160008282546114899190611e3f565b9250508190555081600760008282546114a29190611e3f565b92505081905550826001600160a01b0316846001600160a01b03167f70e256e9264f1aa014ac7f20b4a16618647d26695e23c7181ee67a22c37e7521846040516114ee91815260200190565b60405180910390a350505050565b60606000610f1583611740565b6000610f15836001600160a01b03841661179c565b6000610f15836001600160a01b03841661188f565b6000611588826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166118de9092919063ffffffff16565b90508051600014806115a95750808060200190518101906115a99190611eca565b6110e25760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161047a565b600082600001828154811061161f5761161f611e65565b9060005260206000200154905092915050565b60008161164957611644600884611509565b611654565b61165460088461151e565b9050806116b85760405162461bcd60e51b815260206004820152602c60248201527f5374616b696e673a20496e76616c69642076616c696461746f7220616464726560448201526b737320746f2075706461746560a01b606482015260840161047a565b600160009054906101000a90046001600160a01b03166001600160a01b031663b32805c36040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561106257600080fd5b6040516001600160a01b0380851660248301528316604482015260648101829052610b2d9085906323b872dd60e01b906084016110ab565b60608160000180548060200260200160405190810160405280929190818152602001828054801561179057602002820191906000526020600020905b81548152602001906001019080831161177c575b50505050509050919050565b600081815260018301602052604081205480156118855760006117c0600183611e52565b85549091506000906117d490600190611e52565b90508181146118395760008660000182815481106117f4576117f4611e65565b906000526020600020015490508087600001848154811061181757611817611e65565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061184a5761184a611eec565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610ef1565b6000915050610ef1565b60008181526001830160205260408120546118d657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610ef1565b506000610ef1565b60606118ed84846000856118f5565b949350505050565b6060824710156119565760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161047a565b600080866001600160a01b031685876040516119729190611f26565b60006040518083038185875af1925050503d80600081146119af576040519150601f19603f3d011682016040523d82523d6000602084013e6119b4565b606091505b50915091506119c5878383876119d0565b979650505050505050565b60608315611a3f578251600003611a38576001600160a01b0385163b611a385760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161047a565b50816118ed565b6118ed8383815115611a545781518083602001fd5b8060405162461bcd60e51b815260040161047a9190611f42565b604051806040016040528060006001600160a01b03168152602001611aa3604080518082019091526000808252602082015290565b905290565b6001600160a01b038116811461050d57600080fd5b600060208284031215611acf57600080fd5b8135610f1581611aa8565b600060208284031215611aec57600080fd5b5035919050565b60008060408385031215611b0657600080fd5b50508035926020909101359150565b634e487b7160e01b600052602160045260246000fd5b60038110611b4957634e487b7160e01b600052602160045260246000fd5b9052565b602080825282518282018190526000919060409081850190868401855b82811015611bad57815180516001600160a01b03168552860151805187860152860151611b9986860182611b2b565b506060939093019290850190600101611b6a565b5091979650505050505050565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215611be357600080fd5b8235611bee81611aa8565b9150602083013567ffffffffffffffff80821115611c0b57600080fd5b818501915085601f830112611c1f57600080fd5b813581811115611c3157611c31611bba565b604051601f8201601f19908116603f01168101908382118183101715611c5957611c59611bba565b81604052828152886020848701011115611c7257600080fd5b8260208601602083013760006020848301015280955050505050509250929050565b600080600060608486031215611ca957600080fd5b8335611cb481611aa8565b95602085013595506040909401359392505050565b60208101610ef18284611b2b565b6020808252825182820181905260009190848201906040850190845b81811015611d185783516001600160a01b031683529284019291840191600101611cf3565b50909695505050505050565b600080600080600060a08688031215611d3c57600080fd5b8535945060208601359350604086013560ff81168114611d5b57600080fd5b94979396509394606081013594506080013592915050565b60008060408385031215611d8657600080fd5b8235611d9181611aa8565b946020939093013593505050565b6020808252601d908201527f5374616b696e673a2076616c696461746f7220697320736c6173686564000000604082015260600190565b60208082526033908201527f5374616b696e673a207573657220646f6573206e6f7420686176652077697468604082015272191c985dd85b08185b9b9bdd5b98d95b595b9d606a1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b80820180821115610ef157610ef1611e29565b81810381811115610ef157610ef1611e29565b634e487b7160e01b600052603260045260246000fd5b600060018201611e8d57611e8d611e29565b5060010190565b600060208284031215611ea657600080fd5b8151610f1581611aa8565b600060208284031215611ec357600080fd5b5051919050565b600060208284031215611edc57600080fd5b81518015158114610f1557600080fd5b634e487b7160e01b600052603160045260246000fd5b60005b83811015611f1d578181015183820152602001611f05565b50506000910152565b60008251611f38818460208701611f02565b9190910192915050565b6020815260008251806020840152611f61816040850160208701611f02565b601f01601f1916919091016040019291505056fe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a264697066735822122003336b53b031d51bad853b58cda2dad0622c7b705c06a8f335f0cee91806ace164736f6c63430008120033",
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

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Staking *StakingCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Staking *StakingSession) GetInjector() (common.Address, error) {
	return _Staking.Contract.GetInjector(&_Staking.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_Staking *StakingCallerSession) GetInjector() (common.Address, error) {
	return _Staking.Contract.GetInjector(&_Staking.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _stakeToken, uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _stakeToken common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _stakeToken, _minimalStake, _withdrawalPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _stakeToken, uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_Staking *StakingSession) Initialize(_stakeToken common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _stakeToken, _minimalStake, _withdrawalPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address _stakeToken, uint256 _minimalStake, uint256 _withdrawalPeriod) returns()
func (_Staking *StakingTransactorSession) Initialize(_stakeToken common.Address, _minimalStake *big.Int, _withdrawalPeriod *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _stakeToken, _minimalStake, _withdrawalPeriod)
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

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_Staking *StakingTransactor) SetDependencies(opts *bind.TransactOpts, _contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDependencies", _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_Staking *StakingSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Staking.Contract.SetDependencies(&_Staking.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_Staking *StakingTransactorSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _Staking.Contract.SetDependencies(&_Staking.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Staking *StakingTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Staking *StakingSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetInjector(&_Staking.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_Staking *StakingTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SetInjector(&_Staking.TransactOpts, injector_)
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
