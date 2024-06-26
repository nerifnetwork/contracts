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

// IStakingUserStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type IStakingUserStakeInfo struct {
	UserAddr         common.Address
	UserStakedAmount *big.Int
}

// IStakingWithdrawalAnnouncement is an auto generated low-level Go binding around an user-defined struct.
type IStakingWithdrawalAnnouncement struct {
	TokensAmount      *big.Int
	WithdrawalEpochId *big.Int
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimalStake\",\"type\":\"uint256\"}],\"name\":\"MinimalStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokensSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokensRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeAmount\",\"type\":\"uint256\"}],\"name\":\"TokensStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensAmount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalEpochId\",\"type\":\"uint256\"}],\"name\":\"WithdrawalAnnounced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"announceWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getUsersStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"userStakedAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIStaking.UserStakeInfo[]\",\"name\":\"_usersStakeInfoArr\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"getWithdrawalAnnouncement\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalEpochId\",\"type\":\"uint256\"}],\"internalType\":\"structIStaking.WithdrawalAnnouncement\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"hasWithdrawalAnnouncement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_whitelistedUsers\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userAddr\",\"type\":\"address\"}],\"name\":\"isUserWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimalStake\",\"type\":\"uint256\"}],\"name\":\"setMinimalStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"slashValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sigExpirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"stakeWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_usersToUpdate\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_isAdding\",\"type\":\"bool\"}],\"name\":\"updateWhitelistedUsers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612213806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c80637a766460116100ad5780639ec41a2d116100715780639ec41a2d146102a9578063a694fc3a146102b2578063ecd9ba82146102c5578063ee602ca4146102d8578063fb237eb21461034857600080fd5b80637a7664601461022357806386d381f11461025a5780638b0e9f3f1461027a5780638cb941cc146102835780639aebde6e1461029657600080fd5b80634d0e6ca2116100f45780634d0e6ca2146101cd57806350bfc449146101d557806351ed6a30146101ea57806369130451146101fd5780636d91457a1461021057600080fd5b8063282ec26d146101315780633ccfd60b146101745780633d6ec65e1461017e5780633e3b5b19146101915780634711f586146101ba575b600080fd5b61015f61013f366004611c76565b6001600160a01b03166000908152600a6020526040902060010154151590565b60405190151581526020015b60405180910390f35b61017c61035b565b005b61017c61018c366004611c93565b61056d565b6000805160206121be833981519152545b6040516001600160a01b03909116815260200161016b565b61017c6101c8366004611c76565b610581565b61017c610798565b6101dd6109dd565b60405161016b9190611cac565b6004546101a2906001600160a01b031681565b61017c61020b366004611d68565b6109ee565b61017c61021e366004611e47565b610bc9565b61024c610231366004611c76565b6001600160a01b031660009081526009602052604090205490565b60405190815260200161016b565b61026d610268366004611ea3565b610d3a565b60405161016b9190611ec5565b61024c60065481565b61017c610291366004611c76565b610e47565b61017c6102a4366004611f2b565b610e65565b61024c60055481565b61017c6102c0366004611c93565b610ef8565b61017c6102d3366004611f82565b610f13565b61032d6102e6366004611c76565b6040805180820190915260008082526020820152506001600160a01b03166000908152600a6020908152604091829020825180840190935280548352600101549082015290565b6040805182518152602092830151928101929092520161016b565b61015f610356366004611c76565b610fb7565b610363610fca565b336000908152600a60205260409020600101546103e35760405162461bcd60e51b815260206004820152603360248201527f5374616b696e673a205573657220646f6573206e6f7420686176652077697468604482015272191c985dd85b08185b9b9bdd5b98d95b595b9d606a1b60648201526084015b60405180910390fd5b600160009054906101000a90046001600160a01b03166001600160a01b03166337e4e7806040518163ffffffff1660e01b8152600401602060405180830381865afa158015610436573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061045a9190611fd1565b336000908152600a602052604090206001015411156104d15760405162461bcd60e51b815260206004820152602d60248201527f5374616b696e673a205468652074696d6520666f72207769746864726177616c60448201526c20686173206e6f7420636f6d6560981b60648201526084016103da565b336000908152600a602052604081205460068054919283926104f4908490612000565b9091555050336000818152600960209081526040808320839055600a909152812081815560010155600454610535916001600160a01b039091169083611085565b60405181815233907f6352c5382c4a4578e712449ca65e83cdb392d045dfcf1cad9615189db2da244b9060200160405180910390a250565b6105756110e8565b61057e816111b7565b50565b6105896111f2565b600354604080516305e5d06560e01b815290516000926001600160a01b0316916305e5d06591600480830192869291908290030181865afa1580156105d2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526105fa9190810190612037565b6004805460035460405163f2f3a61f60e01b81529394506001600160a01b039182169391169163f2f3a61f91610632918691016120da565b602060405180830381865afa15801561064f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061067391906120ed565b6001600160a01b0316146106d85760405162461bcd60e51b815260206004820152602660248201527f5374616b696e673a205374616b6520746f6b656e206e6f742061204e45524946604482015265103a37b5b2b760d11b60648201526084016103da565b6003546001600160a01b038381166000908152600960205260409081902054905163ddbb4e8560e01b8152919092169163ddbb4e859161071c91859160040161210a565b600060405180830381600087803b15801561073657600080fd5b505af115801561074a573d6000803e3d6000fd5b50506003546001600160a01b0385811660009081526009602052604090205460045461077d955082169350911690611085565b506001600160a01b0316600090815260096020526040812055565b6107a0610fca565b33600090815260096020526040902054806107fd5760405162461bcd60e51b815260206004820152601c60248201527f5374616b696e673a204e6f7468696e6720746f2077697468647261770000000060448201526064016103da565b336000908152600a6020526040902060010154156108775760405162461bcd60e51b815260206004820152603160248201527f5374616b696e673a205573657220616c72656164792068617320776974686472604482015270185dd85b08185b9b9bdd5b98d95b595b9d607a1b60648201526084016103da565b600060055482106108f757600154604051631f39e1bb60e31b81523360048201526001600160a01b039091169063f9cf0dd8906024016020604051808303816000875af11580156108cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108f09190611fd1565b9050610971565b600160009054906101000a90046001600160a01b03166001600160a01b03166337e4e7806040518163ffffffff1660e01b8152600401602060405180830381865afa15801561094a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096e9190611fd1565b90505b6040805180820182528381526020808201848152336000818152600a8452859020935184559051600190930192909255825185815290810184905290917f9af5208f4a6f6c7001c4c7c54434505fccefba7771cf71365559106c39b59d18910160405180910390a25050565b60606109e9600761125b565b905090565b6109f661126f565b600082905080600060026101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b03166266f0a86040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a5f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8391906120ed565b600160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b031663fb9d9ac56040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ae7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0b91906120ed565b600260006101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b03166348197f776040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b6f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9391906120ed565b600380546001600160a01b0319166001600160a01b039290921691909117905550336000805160206121be833981519152555050565b600054610100900460ff1615808015610be95750600054600160ff909116105b80610c035750303b158015610c03575060005460ff166001145b610c665760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016103da565b6000805460ff191660011790558015610c89576000805461ff0019166101001790555b600480546001600160a01b0319166001600160a01b038716179055610cad846111b7565b610ced8383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506007939250506112f39050565b8015610d33576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b60606000610d52610d4b600761133e565b8585611348565b9050610d5e8482612000565b67ffffffffffffffff811115610d7657610d76611cf9565b604051908082528060200260200182016040528015610dbb57816020015b6040805180820190915260008082526020820152815260200190600190039081610d945790505b509150835b81811015610e3f576000610dd5600783611371565b6040805180820182526001600160a01b038316808252600090815260096020908152929020549181019190915290915084610e108885612000565b81518110610e2057610e2061212c565b6020026020010181905250508080610e3790612142565b915050610dc0565b505092915050565b610e4f61126f565b61057e816000805160206121be83398151915255565b610e6d6110e8565b8015610eb857610eb38383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506007939250506112f39050565b505050565b610eb383838080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525060079392505061137d9050565b610f006113c8565b610f08610fca565b61057e33338361141d565b610f1b6113c8565b610f23610fca565b6004805460405163d505accf60e01b81523392810192909252306024830152604482018790526064820186905260ff8516608483015260a4820184905260c482018390526001600160a01b03169063d505accf9060e401600060405180830381600087803b158015610f9457600080fd5b505af1158015610fa8573d6000803e3d6000fd5b50505050610d3333338761141d565b6000610fc46007836117aa565b92915050565b60015460405163ac4fe5ab60e01b81523360048201526001600160a01b039091169063ac4fe5ab90602401602060405180830381865afa158015611012573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611036919061215b565b156110835760405162461bcd60e51b815260206004820152601e60248201527f5374616b696e673a2056616c696461746f722077617320736c6173686564000060448201526064016103da565b565b6040516001600160a01b038316602482015260448101829052610eb390849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b0319909316929092179091526117cc565b600060029054906101000a90046001600160a01b03166001600160a01b0316637ac3c02f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561113b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061115f91906120ed565b6001600160a01b0316336001600160a01b0316146110835760405162461bcd60e51b815260206004820152601560248201527429ba30b5b4b7339d102737ba10309039b4b3b732b960591b60448201526064016103da565b60058190556040518181527fb6b6b85fb975fbed2f174c07b2154f8746c834172eb59cb23a09c7cbe270e4ca9060200160405180910390a150565b6002546001600160a01b031633146110835760405162461bcd60e51b815260206004820152602660248201527f5374616b696e673a204e6f74206120736c617368696e6720766f74696e67206160448201526564647265737360d01b60648201526084016103da565b60606000611268836118a1565b9392505050565b60006112876000805160206121be8339815191525490565b90506001600160a01b03811615806112a757506001600160a01b03811633145b61057e5760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f7200000000000060448201526064016103da565b60005b8151811015610eb35761132b8282815181106113145761131461212c565b6020026020010151846118fd90919063ffffffff16565b508061133681612142565b9150506112f6565b6000610fc4825490565b60006113548284612178565b9050838111156113615750825b8083111561126857509092915050565b60006112688383611912565b60005b8151811015610eb3576113b582828151811061139e5761139e61212c565b60200260200101518461193c90919063ffffffff16565b50806113c081612142565b915050611380565b6113d133610fb7565b6110835760405162461bcd60e51b815260206004820152601f60248201527f5374616b696e673a204e6f7420612077686974656c697374656420757365720060448201526064016103da565b6000811161146d5760405162461bcd60e51b815260206004820152601a60248201527f5374616b696e673a205a65726f207374616b6520616d6f756e7400000000000060448201526064016103da565b600480546040516370a0823160e01b81526001600160a01b03868116938201939093528392909116906370a0823190602401602060405180830381865afa1580156114bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114e09190611fd1565b101561153a5760405162461bcd60e51b815260206004820152602360248201527f5374616b696e673a204e6f7420656e6f75676820746f6b656e7320746f207374604482015262616b6560e81b60648201526084016103da565b60015460405163facd743b60e01b81526001600160a01b0384811660048301529091169063facd743b90602401602060405180830381865afa158015611584573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115a8919061215b565b156116015760405162461bcd60e51b8152602060048201526024808201527f5374616b696e673a205573657220697320616c726561647920612076616c696460448201526330ba37b960e11b60648201526084016103da565b6001600160a01b0382166000908152600a60205260409020600101541561167e5760405162461bcd60e51b815260206004820152602b60248201527f5374616b696e673a2055736572206861732061207769746864726177616c206160448201526a1b9b9bdd5b98d95b595b9d60aa1b60648201526084016103da565b6001600160a01b0382166000908152600960205260408120546116a2908390612178565b9050600554811061170d57600154604051632691c64760e11b81526001600160a01b03858116600483015290911690634d238c8e90602401600060405180830381600087803b1580156116f457600080fd5b505af1158015611708573d6000803e3d6000fd5b505050505b6001600160a01b03831660009081526009602052604081208290556006805484929061173a908490612178565b9091555050600454611757906001600160a01b0316853085611951565b826001600160a01b0316846001600160a01b03167f70e256e9264f1aa014ac7f20b4a16618647d26695e23c7181ee67a22c37e75218460405161179c91815260200190565b60405180910390a350505050565b6001600160a01b03811660009081526001830160205260408120541515611268565b6000611821826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b031661198f9092919063ffffffff16565b9050805160001480611842575080806020019051810190611842919061215b565b610eb35760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016103da565b6060816000018054806020026020016040519081016040528092919081815260200182805480156118f157602002820191906000526020600020905b8154815260200190600101908083116118dd575b50505050509050919050565b6000611268836001600160a01b0384166119a6565b60008260000182815481106119295761192961212c565b9060005260206000200154905092915050565b6000611268836001600160a01b0384166119f5565b6040516001600160a01b03808516602483015283166044820152606481018290526119899085906323b872dd60e01b906084016110b1565b50505050565b606061199e8484600085611ae8565b949350505050565b60008181526001830160205260408120546119ed57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610fc4565b506000610fc4565b60008181526001830160205260408120548015611ade576000611a19600183612000565b8554909150600090611a2d90600190612000565b9050818114611a92576000866000018281548110611a4d57611a4d61212c565b9060005260206000200154905080876000018481548110611a7057611a7061212c565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080611aa357611aa361218b565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610fc4565b6000915050610fc4565b606082471015611b495760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016103da565b600080866001600160a01b03168587604051611b6591906121a1565b60006040518083038185875af1925050503d8060008114611ba2576040519150601f19603f3d011682016040523d82523d6000602084013e611ba7565b606091505b5091509150611bb887838387611bc3565b979650505050505050565b60608315611c32578251600003611c2b576001600160a01b0385163b611c2b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103da565b508161199e565b61199e8383815115611c475781518083602001fd5b8060405162461bcd60e51b81526004016103da91906120da565b6001600160a01b038116811461057e57600080fd5b600060208284031215611c8857600080fd5b813561126881611c61565b600060208284031215611ca557600080fd5b5035919050565b6020808252825182820181905260009190848201906040850190845b81811015611ced5783516001600160a01b031683529284019291840191600101611cc8565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611d3857611d38611cf9565b604052919050565b600067ffffffffffffffff821115611d5a57611d5a611cf9565b50601f01601f191660200190565b60008060408385031215611d7b57600080fd5b8235611d8681611c61565b9150602083013567ffffffffffffffff811115611da257600080fd5b8301601f81018513611db357600080fd5b8035611dc6611dc182611d40565b611d0f565b818152866020838501011115611ddb57600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008083601f840112611e0d57600080fd5b50813567ffffffffffffffff811115611e2557600080fd5b6020830191508360208260051b8501011115611e4057600080fd5b9250929050565b60008060008060608587031215611e5d57600080fd5b8435611e6881611c61565b935060208501359250604085013567ffffffffffffffff811115611e8b57600080fd5b611e9787828801611dfb565b95989497509550505050565b60008060408385031215611eb657600080fd5b50508035926020909101359150565b602080825282518282018190526000919060409081850190868401855b82811015611f1057815180516001600160a01b03168552860151868501529284019290850190600101611ee2565b5091979650505050505050565b801515811461057e57600080fd5b600080600060408486031215611f4057600080fd5b833567ffffffffffffffff811115611f5757600080fd5b611f6386828701611dfb565b9094509250506020840135611f7781611f1d565b809150509250925092565b600080600080600060a08688031215611f9a57600080fd5b8535945060208601359350604086013560ff81168114611fb957600080fd5b94979396509394606081013594506080013592915050565b600060208284031215611fe357600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610fc457610fc4611fea565b60005b8381101561202e578181015183820152602001612016565b50506000910152565b60006020828403121561204957600080fd5b815167ffffffffffffffff81111561206057600080fd5b8201601f8101841361207157600080fd5b805161207f611dc182611d40565b81815285602083850101111561209457600080fd5b6120a5826020830160208601612013565b95945050505050565b600081518084526120c6816020860160208601612013565b601f01601f19169290920160200192915050565b60208152600061126860208301846120ae565b6000602082840312156120ff57600080fd5b815161126881611c61565b60408152600061211d60408301856120ae565b90508260208301529392505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161215457612154611fea565b5060010190565b60006020828403121561216d57600080fd5b815161126881611f1d565b80820180821115610fc457610fc4611fea565b634e487b7160e01b600052603160045260246000fd5b600082516121b3818460208701612013565b919091019291505056fe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a2646970667358221220b7716e59ccf9e425a903268176c38f98238df73eb8b874313e33821f5a7e44ab64736f6c63430008120033",
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
// Solidity: function getStake(address _userAddr) view returns(uint256)
func (_Staking *StakingCaller) GetStake(opts *bind.CallOpts, _userAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStake", _userAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _userAddr) view returns(uint256)
func (_Staking *StakingSession) GetStake(_userAddr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetStake(&_Staking.CallOpts, _userAddr)
}

// GetStake is a free data retrieval call binding the contract method 0x7a766460.
//
// Solidity: function getStake(address _userAddr) view returns(uint256)
func (_Staking *StakingCallerSession) GetStake(_userAddr common.Address) (*big.Int, error) {
	return _Staking.Contract.GetStake(&_Staking.CallOpts, _userAddr)
}

// GetUsersStakeInfo is a free data retrieval call binding the contract method 0x86d381f1.
//
// Solidity: function getUsersStakeInfo(uint256 _offset, uint256 _limit) view returns((address,uint256)[] _usersStakeInfoArr)
func (_Staking *StakingCaller) GetUsersStakeInfo(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]IStakingUserStakeInfo, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getUsersStakeInfo", _offset, _limit)

	if err != nil {
		return *new([]IStakingUserStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakingUserStakeInfo)).(*[]IStakingUserStakeInfo)

	return out0, err

}

// GetUsersStakeInfo is a free data retrieval call binding the contract method 0x86d381f1.
//
// Solidity: function getUsersStakeInfo(uint256 _offset, uint256 _limit) view returns((address,uint256)[] _usersStakeInfoArr)
func (_Staking *StakingSession) GetUsersStakeInfo(_offset *big.Int, _limit *big.Int) ([]IStakingUserStakeInfo, error) {
	return _Staking.Contract.GetUsersStakeInfo(&_Staking.CallOpts, _offset, _limit)
}

// GetUsersStakeInfo is a free data retrieval call binding the contract method 0x86d381f1.
//
// Solidity: function getUsersStakeInfo(uint256 _offset, uint256 _limit) view returns((address,uint256)[] _usersStakeInfoArr)
func (_Staking *StakingCallerSession) GetUsersStakeInfo(_offset *big.Int, _limit *big.Int) ([]IStakingUserStakeInfo, error) {
	return _Staking.Contract.GetUsersStakeInfo(&_Staking.CallOpts, _offset, _limit)
}

// GetWhitelistedUsers is a free data retrieval call binding the contract method 0x50bfc449.
//
// Solidity: function getWhitelistedUsers() view returns(address[])
func (_Staking *StakingCaller) GetWhitelistedUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getWhitelistedUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedUsers is a free data retrieval call binding the contract method 0x50bfc449.
//
// Solidity: function getWhitelistedUsers() view returns(address[])
func (_Staking *StakingSession) GetWhitelistedUsers() ([]common.Address, error) {
	return _Staking.Contract.GetWhitelistedUsers(&_Staking.CallOpts)
}

// GetWhitelistedUsers is a free data retrieval call binding the contract method 0x50bfc449.
//
// Solidity: function getWhitelistedUsers() view returns(address[])
func (_Staking *StakingCallerSession) GetWhitelistedUsers() ([]common.Address, error) {
	return _Staking.Contract.GetWhitelistedUsers(&_Staking.CallOpts)
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

// IsUserWhitelisted is a free data retrieval call binding the contract method 0xfb237eb2.
//
// Solidity: function isUserWhitelisted(address _userAddr) view returns(bool)
func (_Staking *StakingCaller) IsUserWhitelisted(opts *bind.CallOpts, _userAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isUserWhitelisted", _userAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserWhitelisted is a free data retrieval call binding the contract method 0xfb237eb2.
//
// Solidity: function isUserWhitelisted(address _userAddr) view returns(bool)
func (_Staking *StakingSession) IsUserWhitelisted(_userAddr common.Address) (bool, error) {
	return _Staking.Contract.IsUserWhitelisted(&_Staking.CallOpts, _userAddr)
}

// IsUserWhitelisted is a free data retrieval call binding the contract method 0xfb237eb2.
//
// Solidity: function isUserWhitelisted(address _userAddr) view returns(bool)
func (_Staking *StakingCallerSession) IsUserWhitelisted(_userAddr common.Address) (bool, error) {
	return _Staking.Contract.IsUserWhitelisted(&_Staking.CallOpts, _userAddr)
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

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0x4d0e6ca2.
//
// Solidity: function announceWithdrawal() returns()
func (_Staking *StakingTransactor) AnnounceWithdrawal(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "announceWithdrawal")
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0x4d0e6ca2.
//
// Solidity: function announceWithdrawal() returns()
func (_Staking *StakingSession) AnnounceWithdrawal() (*types.Transaction, error) {
	return _Staking.Contract.AnnounceWithdrawal(&_Staking.TransactOpts)
}

// AnnounceWithdrawal is a paid mutator transaction binding the contract method 0x4d0e6ca2.
//
// Solidity: function announceWithdrawal() returns()
func (_Staking *StakingTransactorSession) AnnounceWithdrawal() (*types.Transaction, error) {
	return _Staking.Contract.AnnounceWithdrawal(&_Staking.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x6d91457a.
//
// Solidity: function initialize(address _stakeToken, uint256 _minimalStake, address[] _whitelistedUsers) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, _stakeToken common.Address, _minimalStake *big.Int, _whitelistedUsers []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", _stakeToken, _minimalStake, _whitelistedUsers)
}

// Initialize is a paid mutator transaction binding the contract method 0x6d91457a.
//
// Solidity: function initialize(address _stakeToken, uint256 _minimalStake, address[] _whitelistedUsers) returns()
func (_Staking *StakingSession) Initialize(_stakeToken common.Address, _minimalStake *big.Int, _whitelistedUsers []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _stakeToken, _minimalStake, _whitelistedUsers)
}

// Initialize is a paid mutator transaction binding the contract method 0x6d91457a.
//
// Solidity: function initialize(address _stakeToken, uint256 _minimalStake, address[] _whitelistedUsers) returns()
func (_Staking *StakingTransactorSession) Initialize(_stakeToken common.Address, _minimalStake *big.Int, _whitelistedUsers []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, _stakeToken, _minimalStake, _whitelistedUsers)
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

// SlashValidator is a paid mutator transaction binding the contract method 0x4711f586.
//
// Solidity: function slashValidator(address _validator) returns()
func (_Staking *StakingTransactor) SlashValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "slashValidator", _validator)
}

// SlashValidator is a paid mutator transaction binding the contract method 0x4711f586.
//
// Solidity: function slashValidator(address _validator) returns()
func (_Staking *StakingSession) SlashValidator(_validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SlashValidator(&_Staking.TransactOpts, _validator)
}

// SlashValidator is a paid mutator transaction binding the contract method 0x4711f586.
//
// Solidity: function slashValidator(address _validator) returns()
func (_Staking *StakingTransactorSession) SlashValidator(_validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.SlashValidator(&_Staking.TransactOpts, _validator)
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

// UpdateWhitelistedUsers is a paid mutator transaction binding the contract method 0x9aebde6e.
//
// Solidity: function updateWhitelistedUsers(address[] _usersToUpdate, bool _isAdding) returns()
func (_Staking *StakingTransactor) UpdateWhitelistedUsers(opts *bind.TransactOpts, _usersToUpdate []common.Address, _isAdding bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateWhitelistedUsers", _usersToUpdate, _isAdding)
}

// UpdateWhitelistedUsers is a paid mutator transaction binding the contract method 0x9aebde6e.
//
// Solidity: function updateWhitelistedUsers(address[] _usersToUpdate, bool _isAdding) returns()
func (_Staking *StakingSession) UpdateWhitelistedUsers(_usersToUpdate []common.Address, _isAdding bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateWhitelistedUsers(&_Staking.TransactOpts, _usersToUpdate, _isAdding)
}

// UpdateWhitelistedUsers is a paid mutator transaction binding the contract method 0x9aebde6e.
//
// Solidity: function updateWhitelistedUsers(address[] _usersToUpdate, bool _isAdding) returns()
func (_Staking *StakingTransactorSession) UpdateWhitelistedUsers(_usersToUpdate []common.Address, _isAdding bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateWhitelistedUsers(&_Staking.TransactOpts, _usersToUpdate, _isAdding)
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

// StakingWithdrawalAnnouncedIterator is returned from FilterWithdrawalAnnounced and is used to iterate over the raw logs and unpacked data for WithdrawalAnnounced events raised by the Staking contract.
type StakingWithdrawalAnnouncedIterator struct {
	Event *StakingWithdrawalAnnounced // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawalAnnouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawalAnnounced)
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
		it.Event = new(StakingWithdrawalAnnounced)
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
func (it *StakingWithdrawalAnnouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawalAnnouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawalAnnounced represents a WithdrawalAnnounced event raised by the Staking contract.
type StakingWithdrawalAnnounced struct {
	UserAddr          common.Address
	TokensAmount      *big.Int
	WithdrawalEpochId *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalAnnounced is a free log retrieval operation binding the contract event 0x9af5208f4a6f6c7001c4c7c54434505fccefba7771cf71365559106c39b59d18.
//
// Solidity: event WithdrawalAnnounced(address indexed userAddr, uint256 tokensAmount, uint256 withdrawalEpochId)
func (_Staking *StakingFilterer) FilterWithdrawalAnnounced(opts *bind.FilterOpts, userAddr []common.Address) (*StakingWithdrawalAnnouncedIterator, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WithdrawalAnnounced", userAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawalAnnouncedIterator{contract: _Staking.contract, event: "WithdrawalAnnounced", logs: logs, sub: sub}, nil
}

// WatchWithdrawalAnnounced is a free log subscription operation binding the contract event 0x9af5208f4a6f6c7001c4c7c54434505fccefba7771cf71365559106c39b59d18.
//
// Solidity: event WithdrawalAnnounced(address indexed userAddr, uint256 tokensAmount, uint256 withdrawalEpochId)
func (_Staking *StakingFilterer) WatchWithdrawalAnnounced(opts *bind.WatchOpts, sink chan<- *StakingWithdrawalAnnounced, userAddr []common.Address) (event.Subscription, error) {

	var userAddrRule []interface{}
	for _, userAddrItem := range userAddr {
		userAddrRule = append(userAddrRule, userAddrItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WithdrawalAnnounced", userAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawalAnnounced)
				if err := _Staking.contract.UnpackLog(event, "WithdrawalAnnounced", log); err != nil {
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

// ParseWithdrawalAnnounced is a log parse operation binding the contract event 0x9af5208f4a6f6c7001c4c7c54434505fccefba7771cf71365559106c39b59d18.
//
// Solidity: event WithdrawalAnnounced(address indexed userAddr, uint256 tokensAmount, uint256 withdrawalEpochId)
func (_Staking *StakingFilterer) ParseWithdrawalAnnounced(log types.Log) (*StakingWithdrawalAnnounced, error) {
	event := new(StakingWithdrawalAnnounced)
	if err := _Staking.contract.UnpackLog(event, "WithdrawalAnnounced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
