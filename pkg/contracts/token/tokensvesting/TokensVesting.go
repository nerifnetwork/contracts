// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokensvesting

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

// VestingBaseSchedule is an auto generated low-level Go binding around an user-defined struct.
type VestingBaseSchedule struct {
	SecondsInPeriod   *big.Int
	DurationInPeriods *big.Int
	CliffInPeriods    *big.Int
}

// VestingSchedule is an auto generated low-level Go binding around an user-defined struct.
type VestingSchedule struct {
	ScheduleData VestingBaseSchedule
	Exponent     *big.Int
}

// VestingVestingData is an auto generated low-level Go binding around an user-defined struct.
type VestingVestingData struct {
	VestingStartTime *big.Int
	Beneficiary      common.Address
	VestingToken     common.Address
	VestingAmount    *big.Int
	PaidAmount       *big.Int
	ScheduleId       *big.Int
}

// TokensVestingMetaData contains all meta data concerning the TokensVesting contract.
var TokensVestingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"scheduleId\",\"type\":\"uint256\"}],\"name\":\"ScheduleCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vestingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"VestingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vestingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawnFromVesting\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINEAR_EXPONENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"secondsInPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationInPeriods\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliffInPeriods\",\"type\":\"uint256\"}],\"internalType\":\"structVesting.BaseSchedule\",\"name\":\"_baseSchedule\",\"type\":\"tuple\"}],\"name\":\"createSchedule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"vestingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vestingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vestingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scheduleId\",\"type\":\"uint256\"}],\"internalType\":\"structVesting.VestingData\",\"name\":\"_vesting\",\"type\":\"tuple\"}],\"name\":\"createVesting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scheduleId_\",\"type\":\"uint256\"}],\"name\":\"getSchedule\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"secondsInPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationInPeriods\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliffInPeriods\",\"type\":\"uint256\"}],\"internalType\":\"structVesting.BaseSchedule\",\"name\":\"scheduleData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"}],\"internalType\":\"structVesting.Schedule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vestingId_\",\"type\":\"uint256\"}],\"name\":\"getVestedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vestingId_\",\"type\":\"uint256\"}],\"name\":\"getVesting\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"vestingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vestingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vestingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scheduleId\",\"type\":\"uint256\"}],\"internalType\":\"structVesting.VestingData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary_\",\"type\":\"address\"}],\"name\":\"getVestingIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary_\",\"type\":\"address\"}],\"name\":\"getVestings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"vestingStartTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vestingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vestingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scheduleId\",\"type\":\"uint256\"}],\"internalType\":\"structVesting.VestingData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vestingId_\",\"type\":\"uint256\"}],\"name\":\"getWithdrawableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scheduleId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vestingId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vestingId_\",\"type\":\"uint256\"}],\"name\":\"withdrawFromVesting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061151c806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063c5ca93a711610097578063dc6ec97911610066578063dc6ec97914610225578063e3e690d514610238578063f2fde38b14610258578063fff5197a1461026b57600080fd5b8063c5ca93a7146101b4578063cafeedf6146101f6578063d453bec614610209578063d8e7151e1461021257600080fd5b80638129fc1c116100d35780638129fc1c1461016b578063875f4384146101735780638da5cb5b14610186578063a5dee4eb146101a157600080fd5b806328c802cd14610105578063615155dd14610121578063715018a6146101415780637a0c6dc01461014b575b600080fd5b61010e60015481565b6040519081526020015b60405180910390f35b61013461012f36600461119c565b610273565b60405161011891906111fd565b6101496102e1565b005b61015e610159366004611222565b6102f5565b604051610118919061123d565b61014961043c565b61010e61018136600461119c565b610552565b6038546040516001600160a01b039091168152602001610118565b6101496101af36600461119c565b610578565b6101c76101c236600461119c565b6106f1565b604080518251805182526020808201518184015290830151928201929092529101516060820152608001610118565b61010e61020436600461119c565b61073e565b61010e60025481565b61010e6102203660046112a1565b61075d565b61010e61023336600461130b565b610778565b61024b610246366004611222565b61078b565b604051610118919061139f565b610149610266366004611222565b6107af565b61010e600181565b61027b611118565b50600090815260046020818152604092839020835160c0810185528154815260018201546001600160a01b03908116938201939093526002820154909216938201939093526003830154606082015290820154608082015260059091015460a082015290565b6102e9610825565b6102f3600061087f565b565b6001600160a01b03811660009081526005602052604081206060919061031a906108d1565b90506000815167ffffffffffffffff8111156103385761033861128b565b60405190808252806020026020018201604052801561037157816020015b61035e611118565b8152602001906001900390816103565790505b50905060005b82518110156104345760046000848381518110610396576103966113d7565b6020908102919091018101518252818101929092526040908101600020815160c0810183528154815260018201546001600160a01b039081169482019490945260028201549093169183019190915260038101546060830152600481015460808301526005015460a08201528251839083908110610416576104166113d7565b6020026020010181905250808061042c90611403565b915050610377565b509392505050565b600054610100900460ff161580801561045c5750600054600160ff909116105b806104765750303b158015610476575060005460ff166001145b6104de5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff191660011790558015610501576000805461ff0019166101001790555b6105096108de565b801561054f576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b600081815260046020526040812060058101546105719082904261090d565b9392505050565b600081815260046020526040902060018101546001600160a01b031633146106085760405162461bcd60e51b815260206004820152603d60248201527f56657374696e6757616c6c65743a206f6e6c792062656e65666963696172792060448201527f63616e2077697468647261772066726f6d206869732076657374696e6700000060648201526084016104d5565b806003015481600401541061066a5760405162461bcd60e51b815260206004820152602260248201527f56657374696e6757616c6c65743a206e6f7468696e6720746f20776974686472604482015261617760f01b60648201526084016104d5565b600061067583610552565b60028301546004840180549293506001600160a01b03909116918391906000906106a090849061141c565b909155506106b1905081338461093a565b837fa55ad56cfbc72f702dc24f36e81385216932d5884a4cfad5f032833ee5ae8fa5836040516106e391815260200190565b60405180910390a250505050565b6106f9611160565b50600090815260036020818152604092839020835160a08101855281549481019485526001820154606082015260028201546080820152938452909101549082015290565b60008181526004602052604081206005810154610571908290426109a1565b6000610767610825565b610770826109b7565b90505b919050565b6000610782610825565b61077082610a5f565b6001600160a01b0381166000908152600560205260409020606090610770906108d1565b6107b7610825565b6001600160a01b03811661081c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016104d5565b61054f8161087f565b6038546001600160a01b031633146102f35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016104d5565b603880546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6060600061057183610be6565b600054610100900460ff166109055760405162461bcd60e51b81526004016104d59061142f565b6102f3610c42565b60008360040154610928848660030154876000015486610c72565b610932919061147a565b949350505050565b6040516376d3d25560e01b81526001600160a01b038381166004830152602482018390528416906376d3d25590604401600060405180830381600087803b15801561098457600080fd5b505af1158015610998573d6000803e3d6000fd5b50505050505050565b6000610932838560030154866000015485610c72565b60006109c282610d4f565b60405180604001604052808381526020016001815250600360006001600081546109eb90611403565b9182905550815260208082019290925260409081016000908120845180518255808501516001808401919091559084015160028301559490930151600390930192909255915491517fa845a9f6012186006a4e2f2e7335aa26cb8fc4035d385989bedadf1551b0a4339190a2505060015490565b6000610a6a82610e64565b60a08201516000908152600360205260409020805460018201544291610a8f9161148d565b8451610a9b919061141c565b11610af35760405162461bcd60e51b815260206004820152603460248201526000805160206114c783398151915260448201527374696e6720666f7220612070617374206461746560601b60648201526084016104d5565b6000600260008154610b0490611403565b91829055506020808601516001600160a01b03166000908152600590915260409020909150610b339082610ff6565b5060008181526004602081815260409283902087518155878201516001820180546001600160a01b03199081166001600160a01b039384169081179092558a8701516002850180549092169316928317905560608a0151600384015560808a01519483019490945560a089015160059092019190915583519283529082015282917ff1220882f139b8959ec281facd523bb0ca18a2a254543a2e3c3606855482fdfb910160405180910390a29392505050565b606081600001805480602002602001604051908101604052809291908181526020018280548015610c3657602002820191906000526020600020905b815481526020019060010190808311610c22575b50505050509050919050565b600054610100900460ff16610c695760405162461bcd60e51b81526004016104d59061142f565b6102f33361087f565b60008481526003602052604081208083851115610c90575050610932565b6000610ca18686846000015461100b565b90508160020154811015610cbb5760009350505050610932565b81600101548110610cd157869350505050610932565b6001820154600090610cee6a084595161401484a0000008461148d565b610cf891906114a4565b9050610d136a084595161401484a0000008560030154611030565b88610d22838760030154611030565b610d2c919061148d565b610d3691906114a4565b9450610d4285896110b3565b9998505050505050505050565b60008160200151118015610d635750805115155b610df05760405162461bcd60e51b815260206004820152605260248201527f56657374696e6757616c6c65743a2063616e6e6f74206372656174652073636860448201527f6564756c652077697468207a65726f206475726174696f6e206f72207a65726f606482015271081cd958dbdb991cc81a5b881c195c9a5bd960721b608482015260a4016104d5565b806020015181604001511061054f5760405162461bcd60e51b815260206004820152603460248201527f56657374696e6757616c6c65743a20636c6966662063616e6e6f7420626520676044820152733932b0ba32b9103a3430b710323ab930ba34b7b760611b60648201526084016104d5565b8051610ebb5760405162461bcd60e51b815260206004820152603260248201526000805160206114c783398151915260448201527174696e6720666f72207a65726f2074696d6560701b60648201526084016104d5565b6000816060015111610f1a5760405162461bcd60e51b815260206004820152603460248201526000805160206114c78339815191526044820152731d1a5b99c8199bdc881e995c9bc8185b5bdd5b9d60621b60648201526084016104d5565b60208101516001600160a01b0316610f805760405162461bcd60e51b815260206004820152603560248201526000805160206114c783398151915260448201527474696e6720666f72207a65726f206164647265737360581b60648201526084016104d5565b60408101516001600160a01b031661054f5760405162461bcd60e51b815260206004820152603360248201527f56657374696e6757616c6c65743a2076657374696e6720746f6b656e2063616e6044820152726e6f74206265207a65726f206164647265737360681b60648201526084016104d5565b600061100283836110c9565b90505b92915050565b600083831161101b576000610932565b81611026858561147a565b61093291906114a4565b60006001821615611041578261104e565b6a084595161401484a0000005b90505b60019190911c908115611005576a084595161401484a000000611074848061148d565b61107e91906114a4565b9250816001166001036110ae576a084595161401484a0000006110a1848361148d565b6110ab91906114a4565b90505b611051565b60008183106110c25781611002565b5090919050565b600081815260018301602052604081205461111057508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611005565b506000611005565b6040518060c001604052806000815260200160006001600160a01b0316815260200160006001600160a01b031681526020016000815260200160008152602001600081525090565b604051806040016040528061118f60405180606001604052806000815260200160008152602001600081525090565b8152602001600081525090565b6000602082840312156111ae57600080fd5b5035919050565b80518252602081015160018060a01b0380821660208501528060408401511660408501525050606081015160608301526080810151608083015260a081015160a08301525050565b60c0810161100582846111b5565b80356001600160a01b038116811461077357600080fd5b60006020828403121561123457600080fd5b6110028261120b565b6020808252825182820181905260009190848201906040850190845b8181101561127f5761126c8385516111b5565b9284019260c09290920191600101611259565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b6000606082840312156112b357600080fd5b6040516060810181811067ffffffffffffffff821117156112e457634e487b7160e01b600052604160045260246000fd5b80604052508235815260208301356020820152604083013560408201528091505092915050565b600060c0828403121561131d57600080fd5b60405160c0810181811067ffffffffffffffff8211171561134e57634e487b7160e01b600052604160045260246000fd5b604052823581526113616020840161120b565b60208201526113726040840161120b565b6040820152606083013560608201526080830135608082015260a083013560a08201528091505092915050565b6020808252825182820181905260009190848201906040850190845b8181101561127f578351835292840192918401916001016113bb565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611415576114156113ed565b5060010190565b80820180821115611005576110056113ed565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b81810381811115611005576110056113ed565b8082028115828204841417611005576110056113ed565b6000826114c157634e487b7160e01b600052601260045260246000fd5b50049056fe56657374696e6757616c6c65743a2063616e6e6f742063726561746520766573a2646970667358221220f3b8d5c3aa970ee286c81d40deb4f6884d06d0b32dca75c3eceecef37d87591f64736f6c63430008120033",
}

// TokensVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use TokensVestingMetaData.ABI instead.
var TokensVestingABI = TokensVestingMetaData.ABI

// TokensVestingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokensVestingMetaData.Bin instead.
var TokensVestingBin = TokensVestingMetaData.Bin

// DeployTokensVesting deploys a new Ethereum contract, binding an instance of TokensVesting to it.
func DeployTokensVesting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokensVesting, error) {
	parsed, err := TokensVestingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokensVestingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokensVesting{TokensVestingCaller: TokensVestingCaller{contract: contract}, TokensVestingTransactor: TokensVestingTransactor{contract: contract}, TokensVestingFilterer: TokensVestingFilterer{contract: contract}}, nil
}

// TokensVesting is an auto generated Go binding around an Ethereum contract.
type TokensVesting struct {
	TokensVestingCaller     // Read-only binding to the contract
	TokensVestingTransactor // Write-only binding to the contract
	TokensVestingFilterer   // Log filterer for contract events
}

// TokensVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokensVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokensVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokensVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokensVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokensVestingSession struct {
	Contract     *TokensVesting    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokensVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokensVestingCallerSession struct {
	Contract *TokensVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokensVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokensVestingTransactorSession struct {
	Contract     *TokensVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokensVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokensVestingRaw struct {
	Contract *TokensVesting // Generic contract binding to access the raw methods on
}

// TokensVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokensVestingCallerRaw struct {
	Contract *TokensVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TokensVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokensVestingTransactorRaw struct {
	Contract *TokensVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokensVesting creates a new instance of TokensVesting, bound to a specific deployed contract.
func NewTokensVesting(address common.Address, backend bind.ContractBackend) (*TokensVesting, error) {
	contract, err := bindTokensVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokensVesting{TokensVestingCaller: TokensVestingCaller{contract: contract}, TokensVestingTransactor: TokensVestingTransactor{contract: contract}, TokensVestingFilterer: TokensVestingFilterer{contract: contract}}, nil
}

// NewTokensVestingCaller creates a new read-only instance of TokensVesting, bound to a specific deployed contract.
func NewTokensVestingCaller(address common.Address, caller bind.ContractCaller) (*TokensVestingCaller, error) {
	contract, err := bindTokensVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokensVestingCaller{contract: contract}, nil
}

// NewTokensVestingTransactor creates a new write-only instance of TokensVesting, bound to a specific deployed contract.
func NewTokensVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TokensVestingTransactor, error) {
	contract, err := bindTokensVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokensVestingTransactor{contract: contract}, nil
}

// NewTokensVestingFilterer creates a new log filterer instance of TokensVesting, bound to a specific deployed contract.
func NewTokensVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TokensVestingFilterer, error) {
	contract, err := bindTokensVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokensVestingFilterer{contract: contract}, nil
}

// bindTokensVesting binds a generic wrapper to an already deployed contract.
func bindTokensVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokensVestingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokensVesting *TokensVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokensVesting.Contract.TokensVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokensVesting *TokensVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokensVesting.Contract.TokensVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokensVesting *TokensVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokensVesting.Contract.TokensVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokensVesting *TokensVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokensVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokensVesting *TokensVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokensVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokensVesting *TokensVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokensVesting.Contract.contract.Transact(opts, method, params...)
}

// LINEAREXPONENT is a free data retrieval call binding the contract method 0xfff5197a.
//
// Solidity: function LINEAR_EXPONENT() view returns(uint256)
func (_TokensVesting *TokensVestingCaller) LINEAREXPONENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "LINEAR_EXPONENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LINEAREXPONENT is a free data retrieval call binding the contract method 0xfff5197a.
//
// Solidity: function LINEAR_EXPONENT() view returns(uint256)
func (_TokensVesting *TokensVestingSession) LINEAREXPONENT() (*big.Int, error) {
	return _TokensVesting.Contract.LINEAREXPONENT(&_TokensVesting.CallOpts)
}

// LINEAREXPONENT is a free data retrieval call binding the contract method 0xfff5197a.
//
// Solidity: function LINEAR_EXPONENT() view returns(uint256)
func (_TokensVesting *TokensVestingCallerSession) LINEAREXPONENT() (*big.Int, error) {
	return _TokensVesting.Contract.LINEAREXPONENT(&_TokensVesting.CallOpts)
}

// GetSchedule is a free data retrieval call binding the contract method 0xc5ca93a7.
//
// Solidity: function getSchedule(uint256 scheduleId_) view returns(((uint256,uint256,uint256),uint256))
func (_TokensVesting *TokensVestingCaller) GetSchedule(opts *bind.CallOpts, scheduleId_ *big.Int) (VestingSchedule, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "getSchedule", scheduleId_)

	if err != nil {
		return *new(VestingSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new(VestingSchedule)).(*VestingSchedule)

	return out0, err

}

// GetSchedule is a free data retrieval call binding the contract method 0xc5ca93a7.
//
// Solidity: function getSchedule(uint256 scheduleId_) view returns(((uint256,uint256,uint256),uint256))
func (_TokensVesting *TokensVestingSession) GetSchedule(scheduleId_ *big.Int) (VestingSchedule, error) {
	return _TokensVesting.Contract.GetSchedule(&_TokensVesting.CallOpts, scheduleId_)
}

// GetSchedule is a free data retrieval call binding the contract method 0xc5ca93a7.
//
// Solidity: function getSchedule(uint256 scheduleId_) view returns(((uint256,uint256,uint256),uint256))
func (_TokensVesting *TokensVestingCallerSession) GetSchedule(scheduleId_ *big.Int) (VestingSchedule, error) {
	return _TokensVesting.Contract.GetSchedule(&_TokensVesting.CallOpts, scheduleId_)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xcafeedf6.
//
// Solidity: function getVestedAmount(uint256 vestingId_) view returns(uint256)
func (_TokensVesting *TokensVestingCaller) GetVestedAmount(opts *bind.CallOpts, vestingId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "getVestedAmount", vestingId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestedAmount is a free data retrieval call binding the contract method 0xcafeedf6.
//
// Solidity: function getVestedAmount(uint256 vestingId_) view returns(uint256)
func (_TokensVesting *TokensVestingSession) GetVestedAmount(vestingId_ *big.Int) (*big.Int, error) {
	return _TokensVesting.Contract.GetVestedAmount(&_TokensVesting.CallOpts, vestingId_)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xcafeedf6.
//
// Solidity: function getVestedAmount(uint256 vestingId_) view returns(uint256)
func (_TokensVesting *TokensVestingCallerSession) GetVestedAmount(vestingId_ *big.Int) (*big.Int, error) {
	return _TokensVesting.Contract.GetVestedAmount(&_TokensVesting.CallOpts, vestingId_)
}

// GetVesting is a free data retrieval call binding the contract method 0x615155dd.
//
// Solidity: function getVesting(uint256 vestingId_) view returns((uint256,address,address,uint256,uint256,uint256))
func (_TokensVesting *TokensVestingCaller) GetVesting(opts *bind.CallOpts, vestingId_ *big.Int) (VestingVestingData, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "getVesting", vestingId_)

	if err != nil {
		return *new(VestingVestingData), err
	}

	out0 := *abi.ConvertType(out[0], new(VestingVestingData)).(*VestingVestingData)

	return out0, err

}

// GetVesting is a free data retrieval call binding the contract method 0x615155dd.
//
// Solidity: function getVesting(uint256 vestingId_) view returns((uint256,address,address,uint256,uint256,uint256))
func (_TokensVesting *TokensVestingSession) GetVesting(vestingId_ *big.Int) (VestingVestingData, error) {
	return _TokensVesting.Contract.GetVesting(&_TokensVesting.CallOpts, vestingId_)
}

// GetVesting is a free data retrieval call binding the contract method 0x615155dd.
//
// Solidity: function getVesting(uint256 vestingId_) view returns((uint256,address,address,uint256,uint256,uint256))
func (_TokensVesting *TokensVestingCallerSession) GetVesting(vestingId_ *big.Int) (VestingVestingData, error) {
	return _TokensVesting.Contract.GetVesting(&_TokensVesting.CallOpts, vestingId_)
}

// GetVestingIds is a free data retrieval call binding the contract method 0xe3e690d5.
//
// Solidity: function getVestingIds(address beneficiary_) view returns(uint256[])
func (_TokensVesting *TokensVestingCaller) GetVestingIds(opts *bind.CallOpts, beneficiary_ common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "getVestingIds", beneficiary_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVestingIds is a free data retrieval call binding the contract method 0xe3e690d5.
//
// Solidity: function getVestingIds(address beneficiary_) view returns(uint256[])
func (_TokensVesting *TokensVestingSession) GetVestingIds(beneficiary_ common.Address) ([]*big.Int, error) {
	return _TokensVesting.Contract.GetVestingIds(&_TokensVesting.CallOpts, beneficiary_)
}

// GetVestingIds is a free data retrieval call binding the contract method 0xe3e690d5.
//
// Solidity: function getVestingIds(address beneficiary_) view returns(uint256[])
func (_TokensVesting *TokensVestingCallerSession) GetVestingIds(beneficiary_ common.Address) ([]*big.Int, error) {
	return _TokensVesting.Contract.GetVestingIds(&_TokensVesting.CallOpts, beneficiary_)
}

// GetVestings is a free data retrieval call binding the contract method 0x7a0c6dc0.
//
// Solidity: function getVestings(address beneficiary_) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_TokensVesting *TokensVestingCaller) GetVestings(opts *bind.CallOpts, beneficiary_ common.Address) ([]VestingVestingData, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "getVestings", beneficiary_)

	if err != nil {
		return *new([]VestingVestingData), err
	}

	out0 := *abi.ConvertType(out[0], new([]VestingVestingData)).(*[]VestingVestingData)

	return out0, err

}

// GetVestings is a free data retrieval call binding the contract method 0x7a0c6dc0.
//
// Solidity: function getVestings(address beneficiary_) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_TokensVesting *TokensVestingSession) GetVestings(beneficiary_ common.Address) ([]VestingVestingData, error) {
	return _TokensVesting.Contract.GetVestings(&_TokensVesting.CallOpts, beneficiary_)
}

// GetVestings is a free data retrieval call binding the contract method 0x7a0c6dc0.
//
// Solidity: function getVestings(address beneficiary_) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_TokensVesting *TokensVestingCallerSession) GetVestings(beneficiary_ common.Address) ([]VestingVestingData, error) {
	return _TokensVesting.Contract.GetVestings(&_TokensVesting.CallOpts, beneficiary_)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x875f4384.
//
// Solidity: function getWithdrawableAmount(uint256 vestingId_) view returns(uint256)
func (_TokensVesting *TokensVestingCaller) GetWithdrawableAmount(opts *bind.CallOpts, vestingId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "getWithdrawableAmount", vestingId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x875f4384.
//
// Solidity: function getWithdrawableAmount(uint256 vestingId_) view returns(uint256)
func (_TokensVesting *TokensVestingSession) GetWithdrawableAmount(vestingId_ *big.Int) (*big.Int, error) {
	return _TokensVesting.Contract.GetWithdrawableAmount(&_TokensVesting.CallOpts, vestingId_)
}

// GetWithdrawableAmount is a free data retrieval call binding the contract method 0x875f4384.
//
// Solidity: function getWithdrawableAmount(uint256 vestingId_) view returns(uint256)
func (_TokensVesting *TokensVestingCallerSession) GetWithdrawableAmount(vestingId_ *big.Int) (*big.Int, error) {
	return _TokensVesting.Contract.GetWithdrawableAmount(&_TokensVesting.CallOpts, vestingId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokensVesting *TokensVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokensVesting *TokensVestingSession) Owner() (common.Address, error) {
	return _TokensVesting.Contract.Owner(&_TokensVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokensVesting *TokensVestingCallerSession) Owner() (common.Address, error) {
	return _TokensVesting.Contract.Owner(&_TokensVesting.CallOpts)
}

// ScheduleId is a free data retrieval call binding the contract method 0x28c802cd.
//
// Solidity: function scheduleId() view returns(uint256)
func (_TokensVesting *TokensVestingCaller) ScheduleId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "scheduleId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScheduleId is a free data retrieval call binding the contract method 0x28c802cd.
//
// Solidity: function scheduleId() view returns(uint256)
func (_TokensVesting *TokensVestingSession) ScheduleId() (*big.Int, error) {
	return _TokensVesting.Contract.ScheduleId(&_TokensVesting.CallOpts)
}

// ScheduleId is a free data retrieval call binding the contract method 0x28c802cd.
//
// Solidity: function scheduleId() view returns(uint256)
func (_TokensVesting *TokensVestingCallerSession) ScheduleId() (*big.Int, error) {
	return _TokensVesting.Contract.ScheduleId(&_TokensVesting.CallOpts)
}

// VestingId is a free data retrieval call binding the contract method 0xd453bec6.
//
// Solidity: function vestingId() view returns(uint256)
func (_TokensVesting *TokensVestingCaller) VestingId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokensVesting.contract.Call(opts, &out, "vestingId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingId is a free data retrieval call binding the contract method 0xd453bec6.
//
// Solidity: function vestingId() view returns(uint256)
func (_TokensVesting *TokensVestingSession) VestingId() (*big.Int, error) {
	return _TokensVesting.Contract.VestingId(&_TokensVesting.CallOpts)
}

// VestingId is a free data retrieval call binding the contract method 0xd453bec6.
//
// Solidity: function vestingId() view returns(uint256)
func (_TokensVesting *TokensVestingCallerSession) VestingId() (*big.Int, error) {
	return _TokensVesting.Contract.VestingId(&_TokensVesting.CallOpts)
}

// CreateSchedule is a paid mutator transaction binding the contract method 0xd8e7151e.
//
// Solidity: function createSchedule((uint256,uint256,uint256) _baseSchedule) returns(uint256)
func (_TokensVesting *TokensVestingTransactor) CreateSchedule(opts *bind.TransactOpts, _baseSchedule VestingBaseSchedule) (*types.Transaction, error) {
	return _TokensVesting.contract.Transact(opts, "createSchedule", _baseSchedule)
}

// CreateSchedule is a paid mutator transaction binding the contract method 0xd8e7151e.
//
// Solidity: function createSchedule((uint256,uint256,uint256) _baseSchedule) returns(uint256)
func (_TokensVesting *TokensVestingSession) CreateSchedule(_baseSchedule VestingBaseSchedule) (*types.Transaction, error) {
	return _TokensVesting.Contract.CreateSchedule(&_TokensVesting.TransactOpts, _baseSchedule)
}

// CreateSchedule is a paid mutator transaction binding the contract method 0xd8e7151e.
//
// Solidity: function createSchedule((uint256,uint256,uint256) _baseSchedule) returns(uint256)
func (_TokensVesting *TokensVestingTransactorSession) CreateSchedule(_baseSchedule VestingBaseSchedule) (*types.Transaction, error) {
	return _TokensVesting.Contract.CreateSchedule(&_TokensVesting.TransactOpts, _baseSchedule)
}

// CreateVesting is a paid mutator transaction binding the contract method 0xdc6ec979.
//
// Solidity: function createVesting((uint256,address,address,uint256,uint256,uint256) _vesting) returns(uint256)
func (_TokensVesting *TokensVestingTransactor) CreateVesting(opts *bind.TransactOpts, _vesting VestingVestingData) (*types.Transaction, error) {
	return _TokensVesting.contract.Transact(opts, "createVesting", _vesting)
}

// CreateVesting is a paid mutator transaction binding the contract method 0xdc6ec979.
//
// Solidity: function createVesting((uint256,address,address,uint256,uint256,uint256) _vesting) returns(uint256)
func (_TokensVesting *TokensVestingSession) CreateVesting(_vesting VestingVestingData) (*types.Transaction, error) {
	return _TokensVesting.Contract.CreateVesting(&_TokensVesting.TransactOpts, _vesting)
}

// CreateVesting is a paid mutator transaction binding the contract method 0xdc6ec979.
//
// Solidity: function createVesting((uint256,address,address,uint256,uint256,uint256) _vesting) returns(uint256)
func (_TokensVesting *TokensVestingTransactorSession) CreateVesting(_vesting VestingVestingData) (*types.Transaction, error) {
	return _TokensVesting.Contract.CreateVesting(&_TokensVesting.TransactOpts, _vesting)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TokensVesting *TokensVestingTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokensVesting.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TokensVesting *TokensVestingSession) Initialize() (*types.Transaction, error) {
	return _TokensVesting.Contract.Initialize(&_TokensVesting.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_TokensVesting *TokensVestingTransactorSession) Initialize() (*types.Transaction, error) {
	return _TokensVesting.Contract.Initialize(&_TokensVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokensVesting *TokensVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokensVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokensVesting *TokensVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokensVesting.Contract.RenounceOwnership(&_TokensVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokensVesting *TokensVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokensVesting.Contract.RenounceOwnership(&_TokensVesting.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokensVesting *TokensVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokensVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokensVesting *TokensVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokensVesting.Contract.TransferOwnership(&_TokensVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokensVesting *TokensVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokensVesting.Contract.TransferOwnership(&_TokensVesting.TransactOpts, newOwner)
}

// WithdrawFromVesting is a paid mutator transaction binding the contract method 0xa5dee4eb.
//
// Solidity: function withdrawFromVesting(uint256 vestingId_) returns()
func (_TokensVesting *TokensVestingTransactor) WithdrawFromVesting(opts *bind.TransactOpts, vestingId_ *big.Int) (*types.Transaction, error) {
	return _TokensVesting.contract.Transact(opts, "withdrawFromVesting", vestingId_)
}

// WithdrawFromVesting is a paid mutator transaction binding the contract method 0xa5dee4eb.
//
// Solidity: function withdrawFromVesting(uint256 vestingId_) returns()
func (_TokensVesting *TokensVestingSession) WithdrawFromVesting(vestingId_ *big.Int) (*types.Transaction, error) {
	return _TokensVesting.Contract.WithdrawFromVesting(&_TokensVesting.TransactOpts, vestingId_)
}

// WithdrawFromVesting is a paid mutator transaction binding the contract method 0xa5dee4eb.
//
// Solidity: function withdrawFromVesting(uint256 vestingId_) returns()
func (_TokensVesting *TokensVestingTransactorSession) WithdrawFromVesting(vestingId_ *big.Int) (*types.Transaction, error) {
	return _TokensVesting.Contract.WithdrawFromVesting(&_TokensVesting.TransactOpts, vestingId_)
}

// TokensVestingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokensVesting contract.
type TokensVestingInitializedIterator struct {
	Event *TokensVestingInitialized // Event containing the contract specifics and raw log

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
func (it *TokensVestingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensVestingInitialized)
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
		it.Event = new(TokensVestingInitialized)
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
func (it *TokensVestingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensVestingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensVestingInitialized represents a Initialized event raised by the TokensVesting contract.
type TokensVestingInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokensVesting *TokensVestingFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokensVestingInitializedIterator, error) {

	logs, sub, err := _TokensVesting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokensVestingInitializedIterator{contract: _TokensVesting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokensVesting *TokensVestingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokensVestingInitialized) (event.Subscription, error) {

	logs, sub, err := _TokensVesting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensVestingInitialized)
				if err := _TokensVesting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TokensVesting *TokensVestingFilterer) ParseInitialized(log types.Log) (*TokensVestingInitialized, error) {
	event := new(TokensVestingInitialized)
	if err := _TokensVesting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokensVesting contract.
type TokensVestingOwnershipTransferredIterator struct {
	Event *TokensVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokensVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensVestingOwnershipTransferred)
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
		it.Event = new(TokensVestingOwnershipTransferred)
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
func (it *TokensVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TokensVesting contract.
type TokensVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokensVesting *TokensVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokensVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokensVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokensVestingOwnershipTransferredIterator{contract: _TokensVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokensVesting *TokensVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokensVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokensVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensVestingOwnershipTransferred)
				if err := _TokensVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokensVesting *TokensVestingFilterer) ParseOwnershipTransferred(log types.Log) (*TokensVestingOwnershipTransferred, error) {
	event := new(TokensVestingOwnershipTransferred)
	if err := _TokensVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensVestingScheduleCreatedIterator is returned from FilterScheduleCreated and is used to iterate over the raw logs and unpacked data for ScheduleCreated events raised by the TokensVesting contract.
type TokensVestingScheduleCreatedIterator struct {
	Event *TokensVestingScheduleCreated // Event containing the contract specifics and raw log

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
func (it *TokensVestingScheduleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensVestingScheduleCreated)
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
		it.Event = new(TokensVestingScheduleCreated)
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
func (it *TokensVestingScheduleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensVestingScheduleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensVestingScheduleCreated represents a ScheduleCreated event raised by the TokensVesting contract.
type TokensVestingScheduleCreated struct {
	ScheduleId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScheduleCreated is a free log retrieval operation binding the contract event 0xa845a9f6012186006a4e2f2e7335aa26cb8fc4035d385989bedadf1551b0a433.
//
// Solidity: event ScheduleCreated(uint256 indexed scheduleId)
func (_TokensVesting *TokensVestingFilterer) FilterScheduleCreated(opts *bind.FilterOpts, scheduleId []*big.Int) (*TokensVestingScheduleCreatedIterator, error) {

	var scheduleIdRule []interface{}
	for _, scheduleIdItem := range scheduleId {
		scheduleIdRule = append(scheduleIdRule, scheduleIdItem)
	}

	logs, sub, err := _TokensVesting.contract.FilterLogs(opts, "ScheduleCreated", scheduleIdRule)
	if err != nil {
		return nil, err
	}
	return &TokensVestingScheduleCreatedIterator{contract: _TokensVesting.contract, event: "ScheduleCreated", logs: logs, sub: sub}, nil
}

// WatchScheduleCreated is a free log subscription operation binding the contract event 0xa845a9f6012186006a4e2f2e7335aa26cb8fc4035d385989bedadf1551b0a433.
//
// Solidity: event ScheduleCreated(uint256 indexed scheduleId)
func (_TokensVesting *TokensVestingFilterer) WatchScheduleCreated(opts *bind.WatchOpts, sink chan<- *TokensVestingScheduleCreated, scheduleId []*big.Int) (event.Subscription, error) {

	var scheduleIdRule []interface{}
	for _, scheduleIdItem := range scheduleId {
		scheduleIdRule = append(scheduleIdRule, scheduleIdItem)
	}

	logs, sub, err := _TokensVesting.contract.WatchLogs(opts, "ScheduleCreated", scheduleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensVestingScheduleCreated)
				if err := _TokensVesting.contract.UnpackLog(event, "ScheduleCreated", log); err != nil {
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

// ParseScheduleCreated is a log parse operation binding the contract event 0xa845a9f6012186006a4e2f2e7335aa26cb8fc4035d385989bedadf1551b0a433.
//
// Solidity: event ScheduleCreated(uint256 indexed scheduleId)
func (_TokensVesting *TokensVestingFilterer) ParseScheduleCreated(log types.Log) (*TokensVestingScheduleCreated, error) {
	event := new(TokensVestingScheduleCreated)
	if err := _TokensVesting.contract.UnpackLog(event, "ScheduleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensVestingVestingCreatedIterator is returned from FilterVestingCreated and is used to iterate over the raw logs and unpacked data for VestingCreated events raised by the TokensVesting contract.
type TokensVestingVestingCreatedIterator struct {
	Event *TokensVestingVestingCreated // Event containing the contract specifics and raw log

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
func (it *TokensVestingVestingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensVestingVestingCreated)
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
		it.Event = new(TokensVestingVestingCreated)
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
func (it *TokensVestingVestingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensVestingVestingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensVestingVestingCreated represents a VestingCreated event raised by the TokensVesting contract.
type TokensVestingVestingCreated struct {
	VestingId   *big.Int
	Beneficiary common.Address
	Token       common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterVestingCreated is a free log retrieval operation binding the contract event 0xf1220882f139b8959ec281facd523bb0ca18a2a254543a2e3c3606855482fdfb.
//
// Solidity: event VestingCreated(uint256 indexed vestingId, address beneficiary, address token)
func (_TokensVesting *TokensVestingFilterer) FilterVestingCreated(opts *bind.FilterOpts, vestingId []*big.Int) (*TokensVestingVestingCreatedIterator, error) {

	var vestingIdRule []interface{}
	for _, vestingIdItem := range vestingId {
		vestingIdRule = append(vestingIdRule, vestingIdItem)
	}

	logs, sub, err := _TokensVesting.contract.FilterLogs(opts, "VestingCreated", vestingIdRule)
	if err != nil {
		return nil, err
	}
	return &TokensVestingVestingCreatedIterator{contract: _TokensVesting.contract, event: "VestingCreated", logs: logs, sub: sub}, nil
}

// WatchVestingCreated is a free log subscription operation binding the contract event 0xf1220882f139b8959ec281facd523bb0ca18a2a254543a2e3c3606855482fdfb.
//
// Solidity: event VestingCreated(uint256 indexed vestingId, address beneficiary, address token)
func (_TokensVesting *TokensVestingFilterer) WatchVestingCreated(opts *bind.WatchOpts, sink chan<- *TokensVestingVestingCreated, vestingId []*big.Int) (event.Subscription, error) {

	var vestingIdRule []interface{}
	for _, vestingIdItem := range vestingId {
		vestingIdRule = append(vestingIdRule, vestingIdItem)
	}

	logs, sub, err := _TokensVesting.contract.WatchLogs(opts, "VestingCreated", vestingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensVestingVestingCreated)
				if err := _TokensVesting.contract.UnpackLog(event, "VestingCreated", log); err != nil {
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

// ParseVestingCreated is a log parse operation binding the contract event 0xf1220882f139b8959ec281facd523bb0ca18a2a254543a2e3c3606855482fdfb.
//
// Solidity: event VestingCreated(uint256 indexed vestingId, address beneficiary, address token)
func (_TokensVesting *TokensVestingFilterer) ParseVestingCreated(log types.Log) (*TokensVestingVestingCreated, error) {
	event := new(TokensVestingVestingCreated)
	if err := _TokensVesting.contract.UnpackLog(event, "VestingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokensVestingWithdrawnFromVestingIterator is returned from FilterWithdrawnFromVesting and is used to iterate over the raw logs and unpacked data for WithdrawnFromVesting events raised by the TokensVesting contract.
type TokensVestingWithdrawnFromVestingIterator struct {
	Event *TokensVestingWithdrawnFromVesting // Event containing the contract specifics and raw log

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
func (it *TokensVestingWithdrawnFromVestingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokensVestingWithdrawnFromVesting)
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
		it.Event = new(TokensVestingWithdrawnFromVesting)
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
func (it *TokensVestingWithdrawnFromVestingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokensVestingWithdrawnFromVestingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokensVestingWithdrawnFromVesting represents a WithdrawnFromVesting event raised by the TokensVesting contract.
type TokensVestingWithdrawnFromVesting struct {
	VestingId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFromVesting is a free log retrieval operation binding the contract event 0xa55ad56cfbc72f702dc24f36e81385216932d5884a4cfad5f032833ee5ae8fa5.
//
// Solidity: event WithdrawnFromVesting(uint256 indexed vestingId, uint256 amount)
func (_TokensVesting *TokensVestingFilterer) FilterWithdrawnFromVesting(opts *bind.FilterOpts, vestingId []*big.Int) (*TokensVestingWithdrawnFromVestingIterator, error) {

	var vestingIdRule []interface{}
	for _, vestingIdItem := range vestingId {
		vestingIdRule = append(vestingIdRule, vestingIdItem)
	}

	logs, sub, err := _TokensVesting.contract.FilterLogs(opts, "WithdrawnFromVesting", vestingIdRule)
	if err != nil {
		return nil, err
	}
	return &TokensVestingWithdrawnFromVestingIterator{contract: _TokensVesting.contract, event: "WithdrawnFromVesting", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFromVesting is a free log subscription operation binding the contract event 0xa55ad56cfbc72f702dc24f36e81385216932d5884a4cfad5f032833ee5ae8fa5.
//
// Solidity: event WithdrawnFromVesting(uint256 indexed vestingId, uint256 amount)
func (_TokensVesting *TokensVestingFilterer) WatchWithdrawnFromVesting(opts *bind.WatchOpts, sink chan<- *TokensVestingWithdrawnFromVesting, vestingId []*big.Int) (event.Subscription, error) {

	var vestingIdRule []interface{}
	for _, vestingIdItem := range vestingId {
		vestingIdRule = append(vestingIdRule, vestingIdItem)
	}

	logs, sub, err := _TokensVesting.contract.WatchLogs(opts, "WithdrawnFromVesting", vestingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokensVestingWithdrawnFromVesting)
				if err := _TokensVesting.contract.UnpackLog(event, "WithdrawnFromVesting", log); err != nil {
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

// ParseWithdrawnFromVesting is a log parse operation binding the contract event 0xa55ad56cfbc72f702dc24f36e81385216932d5884a4cfad5f032833ee5ae8fa5.
//
// Solidity: event WithdrawnFromVesting(uint256 indexed vestingId, uint256 amount)
func (_TokensVesting *TokensVestingFilterer) ParseWithdrawnFromVesting(log types.Log) (*TokensVestingWithdrawnFromVesting, error) {
	event := new(TokensVestingWithdrawnFromVesting)
	if err := _TokensVesting.contract.UnpackLog(event, "WithdrawnFromVesting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
