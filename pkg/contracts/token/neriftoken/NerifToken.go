// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package neriftoken

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

// NerifTokenMetaData contains all meta data concerning the NerifToken contract.
var NerifTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInjector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokensAmount\",\"type\":\"uint256\"}],\"name\":\"ownerMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractsRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"setDependencies\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"injector_\",\"type\":\"address\"}],\"name\":\"setInjector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokensAmount\",\"type\":\"uint256\"}],\"name\":\"vestingMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611a7c806100206000396000f3fe608060405234801561001057600080fd5b50600436106101585760003560e01c806376d3d255116100c3578063a457c2d71161007c578063a457c2d7146102dc578063a9059cbb146102ef578063d505accf14610302578063dd62ed3e14610315578063eedfca5f14610328578063f2fde38b1461033b57600080fd5b806376d3d2551461026f5780637ecebe001461028257806384b0196e146102955780638cb941cc146102b05780638da5cb5b146102c357806395d89b41146102d457600080fd5b8063395093511161011557806339509351146101da5780633e3b5b19146101ed578063484b973c14610216578063691304511461022b57806370a082311461023e578063715018a61461026757600080fd5b806306fdde031461015d578063095ea7b31461017b57806318160ddd1461019e57806323b872dd146101b0578063313ce567146101c35780633644e515146101d2575b600080fd5b61016561034e565b60405161017291906114c9565b60405180910390f35b61018e6101893660046114f8565b6103e0565b6040519015158152602001610172565b6035545b604051908152602001610172565b61018e6101be366004611524565b6103fa565b60405160128152602001610172565b6101a261041e565b61018e6101e83660046114f8565b61042d565b600080516020611a27833981519152545b6040516001600160a01b039091168152602001610172565b6102296102243660046114f8565b61044f565b005b6102296102393660046115f1565b610465565b6101a261024c366004611655565b6001600160a01b031660009081526033602052604090205490565b61022961050a565b61022961027d3660046114f8565b61051e565b6101a2610290366004611655565b610526565b61029d610544565b6040516101729796959493929190611672565b6102296102be366004611655565b6105e7565b60cc546001600160a01b03166101fe565b610165610608565b61018e6102ea3660046114f8565b610617565b61018e6102fd3660046114f8565b610692565b610229610310366004611708565b6106a0565b6101a261032336600461177f565b610804565b6102296103363660046117d8565b61082f565b610229610349366004611655565b610960565b60606036805461035d90611845565b80601f016020809104026020016040519081016040528092919081815260200182805461038990611845565b80156103d65780601f106103ab576101008083540402835291602001916103d6565b820191906000526020600020905b8154815290600101906020018083116103b957829003601f168201915b5050505050905090565b6000336103ee8185856109d6565b60019150505b92915050565b600033610408858285610afa565b610413858585610b6e565b506001949350505050565b6000610428610d19565b905090565b6000336103ee8185856104408383610804565b61044a9190611879565b6109d6565b610457610d23565b6104618282610d7d565b5050565b61046d610e3e565b6000829050806001600160a01b031663d2bea9c66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d4919061189a565b60fe80546001600160a01b0319166001600160a01b03929092169190911790555033600080516020611a27833981519152555050565b610512610d23565b61051c6000610ec2565b565b610457610f14565b6001600160a01b0381166000908152609960205260408120546103f4565b6000606080600080600060606065546000801b1480156105645750606654155b6105ad5760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b6105b5610f6e565b6105bd610f7d565b60408051600080825260208201909252600f60f81b9b939a50919850469750309650945092509050565b6105ef610e3e565b61060581600080516020611a2783398151915255565b50565b60606037805461035d90611845565b600033816106258286610804565b9050838110156106855760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016105a4565b61041382868684036109d6565b6000336103ee818585610b6e565b834211156106f05760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016105a4565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c988888861071f8c610f8c565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e001604051602081830303815290604052805190602001209050600061077a82610fb4565b9050600061078a82878787610fe1565b9050896001600160a01b0316816001600160a01b0316146107ed5760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016105a4565b6107f88a8a8a6109d6565b50505050505050505050565b6001600160a01b03918216600090815260346020908152604080832093909416825291909152205490565b600054610100900460ff161580801561084f5750600054600160ff909116105b806108695750303b158015610869575060005460ff166001145b6108cc5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016105a4565b6000805460ff1916600117905580156108ef576000805461ff0019166101001790555b6108f7611009565b6109018383611038565b61090a83611069565b6109143385610d7d565b801561095a576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050565b610968610d23565b6001600160a01b0381166109cd5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016105a4565b61060581610ec2565b6001600160a01b038316610a385760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016105a4565b6001600160a01b038216610a995760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016105a4565b6001600160a01b0383811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6000610b068484610804565b9050600019811461095a5781811015610b615760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016105a4565b61095a84848484036109d6565b6001600160a01b038316610bd25760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016105a4565b6001600160a01b038216610c345760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016105a4565b6001600160a01b03831660009081526033602052604090205481811015610cac5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016105a4565b6001600160a01b0380851660008181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610d0c9086815260200190565b60405180910390a361095a565b60006104286110b8565b60cc546001600160a01b0316331461051c5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105a4565b6001600160a01b038216610dd35760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016105a4565b8060356000828254610de59190611879565b90915550506001600160a01b0382166000818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6000610e56600080516020611a278339815191525490565b90506001600160a01b0381161580610e7657506001600160a01b03811633145b6106055760405162461bcd60e51b815260206004820152601a60248201527f446570656e64616e743a206e6f7420616e20696e6a6563746f7200000000000060448201526064016105a4565b60cc80546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60fe546001600160a01b0316331461051c5760405162461bcd60e51b815260206004820181905260248201527f4e65726966546f6b656e3a204e6f74206120746f6b656e732076657374696e6760448201526064016105a4565b60606067805461035d90611845565b60606068805461035d90611845565b6001600160a01b03811660009081526099602052604090208054600181018255905b50919050565b60006103f4610fc1610d19565b8360405161190160f01b8152600281019290925260228201526042902090565b6000806000610ff28787878761112c565b91509150610fff816111f0565b5095945050505050565b600054610100900460ff166110305760405162461bcd60e51b81526004016105a4906118b7565b61051c61133a565b600054610100900460ff1661105f5760405162461bcd60e51b81526004016105a4906118b7565b610461828261136a565b600054610100900460ff166110905760405162461bcd60e51b81526004016105a4906118b7565b61060581604051806040016040528060018152602001603160f81b8152506113aa565b505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6110e36113f9565b6110eb611452565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561116357506000905060036111e7565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa1580156111b7573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166111e0576000600192509250506111e7565b9150600090505b94509492505050565b600081600481111561120457611204611902565b0361120c5750565b600181600481111561122057611220611902565b0361126d5760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016105a4565b600281600481111561128157611281611902565b036112ce5760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016105a4565b60038160048111156112e2576112e2611902565b036106055760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b60648201526084016105a4565b600054610100900460ff166113615760405162461bcd60e51b81526004016105a4906118b7565b61051c33610ec2565b600054610100900460ff166113915760405162461bcd60e51b81526004016105a4906118b7565b603661139d8382611966565b5060376110b38282611966565b600054610100900460ff166113d15760405162461bcd60e51b81526004016105a4906118b7565b60676113dd8382611966565b5060686113ea8282611966565b50506000606581905560665550565b600080611404610f6e565b80519091501561141b578051602090910120919050565b606554801561142a5792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b60008061145d610f7d565b805190915015611474578051602090910120919050565b606654801561142a5792915050565b6000815180845260005b818110156114a95760208185018101518683018201520161148d565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006114dc6020830184611483565b9392505050565b6001600160a01b038116811461060557600080fd5b6000806040838503121561150b57600080fd5b8235611516816114e3565b946020939093013593505050565b60008060006060848603121561153957600080fd5b8335611544816114e3565b92506020840135611554816114e3565b929592945050506040919091013590565b634e487b7160e01b600052604160045260246000fd5b600067ffffffffffffffff8084111561159657611596611565565b604051601f8501601f19908116603f011681019082821181831017156115be576115be611565565b816040528093508581528686860111156115d757600080fd5b858560208301376000602087830101525050509392505050565b6000806040838503121561160457600080fd5b823561160f816114e3565b9150602083013567ffffffffffffffff81111561162b57600080fd5b8301601f8101851361163c57600080fd5b61164b8582356020840161157b565b9150509250929050565b60006020828403121561166757600080fd5b81356114dc816114e3565b60ff60f81b881681526000602060e08184015261169260e084018a611483565b83810360408501526116a4818a611483565b606085018990526001600160a01b038816608086015260a0850187905284810360c0860152855180825283870192509083019060005b818110156116f6578351835292840192918401916001016116da565b50909c9b505050505050505050505050565b600080600080600080600060e0888a03121561172357600080fd5b873561172e816114e3565b9650602088013561173e816114e3565b95506040880135945060608801359350608088013560ff8116811461176257600080fd5b9699959850939692959460a0840135945060c09093013592915050565b6000806040838503121561179257600080fd5b823561179d816114e3565b915060208301356117ad816114e3565b809150509250929050565b600082601f8301126117c957600080fd5b6114dc8383356020850161157b565b6000806000606084860312156117ed57600080fd5b83359250602084013567ffffffffffffffff8082111561180c57600080fd5b611818878388016117b8565b9350604086013591508082111561182e57600080fd5b5061183b868287016117b8565b9150509250925092565b600181811c9082168061185957607f821691505b602082108103610fae57634e487b7160e01b600052602260045260246000fd5b808201808211156103f457634e487b7160e01b600052601160045260246000fd5b6000602082840312156118ac57600080fd5b81516114dc816114e3565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b634e487b7160e01b600052602160045260246000fd5b601f8211156110b357600081815260208120601f850160051c8101602086101561193f5750805b601f850160051c820191505b8181101561195e5782815560010161194b565b505050505050565b815167ffffffffffffffff81111561198057611980611565565b6119948161198e8454611845565b84611918565b602080601f8311600181146119c957600084156119b15750858301515b600019600386901b1c1916600185901b17855561195e565b600085815260208120601f198616915b828110156119f8578886015182559484019460019091019084016119d9565b5085821015611a165787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fe3d1f25f1ac447e55e7fec744471c4dab1c6a2b6ffb897825f9ea3d2e8c9be583a2646970667358221220c54c052e5a13b1974be649490cb64a447f51f991a710e349a3f8fd947d39df0364736f6c63430008120033",
}

// NerifTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use NerifTokenMetaData.ABI instead.
var NerifTokenABI = NerifTokenMetaData.ABI

// NerifTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NerifTokenMetaData.Bin instead.
var NerifTokenBin = NerifTokenMetaData.Bin

// DeployNerifToken deploys a new Ethereum contract, binding an instance of NerifToken to it.
func DeployNerifToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NerifToken, error) {
	parsed, err := NerifTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NerifTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NerifToken{NerifTokenCaller: NerifTokenCaller{contract: contract}, NerifTokenTransactor: NerifTokenTransactor{contract: contract}, NerifTokenFilterer: NerifTokenFilterer{contract: contract}}, nil
}

// NerifToken is an auto generated Go binding around an Ethereum contract.
type NerifToken struct {
	NerifTokenCaller     // Read-only binding to the contract
	NerifTokenTransactor // Write-only binding to the contract
	NerifTokenFilterer   // Log filterer for contract events
}

// NerifTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NerifTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NerifTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NerifTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NerifTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NerifTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NerifTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NerifTokenSession struct {
	Contract     *NerifToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NerifTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NerifTokenCallerSession struct {
	Contract *NerifTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NerifTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NerifTokenTransactorSession struct {
	Contract     *NerifTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NerifTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NerifTokenRaw struct {
	Contract *NerifToken // Generic contract binding to access the raw methods on
}

// NerifTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NerifTokenCallerRaw struct {
	Contract *NerifTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NerifTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NerifTokenTransactorRaw struct {
	Contract *NerifTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNerifToken creates a new instance of NerifToken, bound to a specific deployed contract.
func NewNerifToken(address common.Address, backend bind.ContractBackend) (*NerifToken, error) {
	contract, err := bindNerifToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NerifToken{NerifTokenCaller: NerifTokenCaller{contract: contract}, NerifTokenTransactor: NerifTokenTransactor{contract: contract}, NerifTokenFilterer: NerifTokenFilterer{contract: contract}}, nil
}

// NewNerifTokenCaller creates a new read-only instance of NerifToken, bound to a specific deployed contract.
func NewNerifTokenCaller(address common.Address, caller bind.ContractCaller) (*NerifTokenCaller, error) {
	contract, err := bindNerifToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NerifTokenCaller{contract: contract}, nil
}

// NewNerifTokenTransactor creates a new write-only instance of NerifToken, bound to a specific deployed contract.
func NewNerifTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NerifTokenTransactor, error) {
	contract, err := bindNerifToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NerifTokenTransactor{contract: contract}, nil
}

// NewNerifTokenFilterer creates a new log filterer instance of NerifToken, bound to a specific deployed contract.
func NewNerifTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NerifTokenFilterer, error) {
	contract, err := bindNerifToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NerifTokenFilterer{contract: contract}, nil
}

// bindNerifToken binds a generic wrapper to an already deployed contract.
func bindNerifToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NerifTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NerifToken *NerifTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NerifToken.Contract.NerifTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NerifToken *NerifTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NerifToken.Contract.NerifTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NerifToken *NerifTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NerifToken.Contract.NerifTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NerifToken *NerifTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NerifToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NerifToken *NerifTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NerifToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NerifToken *NerifTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NerifToken.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_NerifToken *NerifTokenCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_NerifToken *NerifTokenSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _NerifToken.Contract.DOMAINSEPARATOR(&_NerifToken.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_NerifToken *NerifTokenCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _NerifToken.Contract.DOMAINSEPARATOR(&_NerifToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NerifToken *NerifTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NerifToken *NerifTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NerifToken.Contract.Allowance(&_NerifToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_NerifToken *NerifTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _NerifToken.Contract.Allowance(&_NerifToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NerifToken *NerifTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NerifToken *NerifTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NerifToken.Contract.BalanceOf(&_NerifToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_NerifToken *NerifTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _NerifToken.Contract.BalanceOf(&_NerifToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NerifToken *NerifTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NerifToken *NerifTokenSession) Decimals() (uint8, error) {
	return _NerifToken.Contract.Decimals(&_NerifToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NerifToken *NerifTokenCallerSession) Decimals() (uint8, error) {
	return _NerifToken.Contract.Decimals(&_NerifToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NerifToken *NerifTokenCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NerifToken *NerifTokenSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _NerifToken.Contract.Eip712Domain(&_NerifToken.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_NerifToken *NerifTokenCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _NerifToken.Contract.Eip712Domain(&_NerifToken.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_NerifToken *NerifTokenCaller) GetInjector(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "getInjector")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_NerifToken *NerifTokenSession) GetInjector() (common.Address, error) {
	return _NerifToken.Contract.GetInjector(&_NerifToken.CallOpts)
}

// GetInjector is a free data retrieval call binding the contract method 0x3e3b5b19.
//
// Solidity: function getInjector() view returns(address injector_)
func (_NerifToken *NerifTokenCallerSession) GetInjector() (common.Address, error) {
	return _NerifToken.Contract.GetInjector(&_NerifToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NerifToken *NerifTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NerifToken *NerifTokenSession) Name() (string, error) {
	return _NerifToken.Contract.Name(&_NerifToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NerifToken *NerifTokenCallerSession) Name() (string, error) {
	return _NerifToken.Contract.Name(&_NerifToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_NerifToken *NerifTokenCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_NerifToken *NerifTokenSession) Nonces(owner common.Address) (*big.Int, error) {
	return _NerifToken.Contract.Nonces(&_NerifToken.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_NerifToken *NerifTokenCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _NerifToken.Contract.Nonces(&_NerifToken.CallOpts, owner)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NerifToken *NerifTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NerifToken *NerifTokenSession) Owner() (common.Address, error) {
	return _NerifToken.Contract.Owner(&_NerifToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NerifToken *NerifTokenCallerSession) Owner() (common.Address, error) {
	return _NerifToken.Contract.Owner(&_NerifToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NerifToken *NerifTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NerifToken *NerifTokenSession) Symbol() (string, error) {
	return _NerifToken.Contract.Symbol(&_NerifToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NerifToken *NerifTokenCallerSession) Symbol() (string, error) {
	return _NerifToken.Contract.Symbol(&_NerifToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NerifToken *NerifTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NerifToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NerifToken *NerifTokenSession) TotalSupply() (*big.Int, error) {
	return _NerifToken.Contract.TotalSupply(&_NerifToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NerifToken *NerifTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NerifToken.Contract.TotalSupply(&_NerifToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.Approve(&_NerifToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.Approve(&_NerifToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NerifToken *NerifTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NerifToken *NerifTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.DecreaseAllowance(&_NerifToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_NerifToken *NerifTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.DecreaseAllowance(&_NerifToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NerifToken *NerifTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NerifToken *NerifTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.IncreaseAllowance(&_NerifToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_NerifToken *NerifTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.IncreaseAllowance(&_NerifToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xeedfca5f.
//
// Solidity: function initialize(uint256 _initTokensAmount, string _name, string _symbol) returns()
func (_NerifToken *NerifTokenTransactor) Initialize(opts *bind.TransactOpts, _initTokensAmount *big.Int, _name string, _symbol string) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "initialize", _initTokensAmount, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0xeedfca5f.
//
// Solidity: function initialize(uint256 _initTokensAmount, string _name, string _symbol) returns()
func (_NerifToken *NerifTokenSession) Initialize(_initTokensAmount *big.Int, _name string, _symbol string) (*types.Transaction, error) {
	return _NerifToken.Contract.Initialize(&_NerifToken.TransactOpts, _initTokensAmount, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0xeedfca5f.
//
// Solidity: function initialize(uint256 _initTokensAmount, string _name, string _symbol) returns()
func (_NerifToken *NerifTokenTransactorSession) Initialize(_initTokensAmount *big.Int, _name string, _symbol string) (*types.Transaction, error) {
	return _NerifToken.Contract.Initialize(&_NerifToken.TransactOpts, _initTokensAmount, _name, _symbol)
}

// OwnerMint is a paid mutator transaction binding the contract method 0x484b973c.
//
// Solidity: function ownerMint(address _recipient, uint256 _tokensAmount) returns()
func (_NerifToken *NerifTokenTransactor) OwnerMint(opts *bind.TransactOpts, _recipient common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "ownerMint", _recipient, _tokensAmount)
}

// OwnerMint is a paid mutator transaction binding the contract method 0x484b973c.
//
// Solidity: function ownerMint(address _recipient, uint256 _tokensAmount) returns()
func (_NerifToken *NerifTokenSession) OwnerMint(_recipient common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.OwnerMint(&_NerifToken.TransactOpts, _recipient, _tokensAmount)
}

// OwnerMint is a paid mutator transaction binding the contract method 0x484b973c.
//
// Solidity: function ownerMint(address _recipient, uint256 _tokensAmount) returns()
func (_NerifToken *NerifTokenTransactorSession) OwnerMint(_recipient common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.OwnerMint(&_NerifToken.TransactOpts, _recipient, _tokensAmount)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_NerifToken *NerifTokenTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_NerifToken *NerifTokenSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NerifToken.Contract.Permit(&_NerifToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_NerifToken *NerifTokenTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NerifToken.Contract.Permit(&_NerifToken.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NerifToken *NerifTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NerifToken *NerifTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _NerifToken.Contract.RenounceOwnership(&_NerifToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NerifToken *NerifTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NerifToken.Contract.RenounceOwnership(&_NerifToken.TransactOpts)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_NerifToken *NerifTokenTransactor) SetDependencies(opts *bind.TransactOpts, _contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "setDependencies", _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_NerifToken *NerifTokenSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _NerifToken.Contract.SetDependencies(&_NerifToken.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetDependencies is a paid mutator transaction binding the contract method 0x69130451.
//
// Solidity: function setDependencies(address _contractsRegistryAddr, bytes ) returns()
func (_NerifToken *NerifTokenTransactorSession) SetDependencies(_contractsRegistryAddr common.Address, arg1 []byte) (*types.Transaction, error) {
	return _NerifToken.Contract.SetDependencies(&_NerifToken.TransactOpts, _contractsRegistryAddr, arg1)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_NerifToken *NerifTokenTransactor) SetInjector(opts *bind.TransactOpts, injector_ common.Address) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "setInjector", injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_NerifToken *NerifTokenSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _NerifToken.Contract.SetInjector(&_NerifToken.TransactOpts, injector_)
}

// SetInjector is a paid mutator transaction binding the contract method 0x8cb941cc.
//
// Solidity: function setInjector(address injector_) returns()
func (_NerifToken *NerifTokenTransactorSession) SetInjector(injector_ common.Address) (*types.Transaction, error) {
	return _NerifToken.Contract.SetInjector(&_NerifToken.TransactOpts, injector_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.Transfer(&_NerifToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.Transfer(&_NerifToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.TransferFrom(&_NerifToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_NerifToken *NerifTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.TransferFrom(&_NerifToken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NerifToken *NerifTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NerifToken *NerifTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NerifToken.Contract.TransferOwnership(&_NerifToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NerifToken *NerifTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NerifToken.Contract.TransferOwnership(&_NerifToken.TransactOpts, newOwner)
}

// VestingMint is a paid mutator transaction binding the contract method 0x76d3d255.
//
// Solidity: function vestingMint(address _recipient, uint256 _tokensAmount) returns()
func (_NerifToken *NerifTokenTransactor) VestingMint(opts *bind.TransactOpts, _recipient common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _NerifToken.contract.Transact(opts, "vestingMint", _recipient, _tokensAmount)
}

// VestingMint is a paid mutator transaction binding the contract method 0x76d3d255.
//
// Solidity: function vestingMint(address _recipient, uint256 _tokensAmount) returns()
func (_NerifToken *NerifTokenSession) VestingMint(_recipient common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.VestingMint(&_NerifToken.TransactOpts, _recipient, _tokensAmount)
}

// VestingMint is a paid mutator transaction binding the contract method 0x76d3d255.
//
// Solidity: function vestingMint(address _recipient, uint256 _tokensAmount) returns()
func (_NerifToken *NerifTokenTransactorSession) VestingMint(_recipient common.Address, _tokensAmount *big.Int) (*types.Transaction, error) {
	return _NerifToken.Contract.VestingMint(&_NerifToken.TransactOpts, _recipient, _tokensAmount)
}

// NerifTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NerifToken contract.
type NerifTokenApprovalIterator struct {
	Event *NerifTokenApproval // Event containing the contract specifics and raw log

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
func (it *NerifTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifTokenApproval)
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
		it.Event = new(NerifTokenApproval)
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
func (it *NerifTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifTokenApproval represents a Approval event raised by the NerifToken contract.
type NerifTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NerifToken *NerifTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NerifTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NerifToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NerifTokenApprovalIterator{contract: _NerifToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NerifToken *NerifTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NerifTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _NerifToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifTokenApproval)
				if err := _NerifToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_NerifToken *NerifTokenFilterer) ParseApproval(log types.Log) (*NerifTokenApproval, error) {
	event := new(NerifTokenApproval)
	if err := _NerifToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NerifTokenEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the NerifToken contract.
type NerifTokenEIP712DomainChangedIterator struct {
	Event *NerifTokenEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *NerifTokenEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifTokenEIP712DomainChanged)
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
		it.Event = new(NerifTokenEIP712DomainChanged)
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
func (it *NerifTokenEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifTokenEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifTokenEIP712DomainChanged represents a EIP712DomainChanged event raised by the NerifToken contract.
type NerifTokenEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_NerifToken *NerifTokenFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*NerifTokenEIP712DomainChangedIterator, error) {

	logs, sub, err := _NerifToken.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &NerifTokenEIP712DomainChangedIterator{contract: _NerifToken.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_NerifToken *NerifTokenFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *NerifTokenEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _NerifToken.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifTokenEIP712DomainChanged)
				if err := _NerifToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_NerifToken *NerifTokenFilterer) ParseEIP712DomainChanged(log types.Log) (*NerifTokenEIP712DomainChanged, error) {
	event := new(NerifTokenEIP712DomainChanged)
	if err := _NerifToken.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NerifTokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NerifToken contract.
type NerifTokenInitializedIterator struct {
	Event *NerifTokenInitialized // Event containing the contract specifics and raw log

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
func (it *NerifTokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifTokenInitialized)
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
		it.Event = new(NerifTokenInitialized)
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
func (it *NerifTokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifTokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifTokenInitialized represents a Initialized event raised by the NerifToken contract.
type NerifTokenInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NerifToken *NerifTokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*NerifTokenInitializedIterator, error) {

	logs, sub, err := _NerifToken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NerifTokenInitializedIterator{contract: _NerifToken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_NerifToken *NerifTokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NerifTokenInitialized) (event.Subscription, error) {

	logs, sub, err := _NerifToken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifTokenInitialized)
				if err := _NerifToken.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NerifToken *NerifTokenFilterer) ParseInitialized(log types.Log) (*NerifTokenInitialized, error) {
	event := new(NerifTokenInitialized)
	if err := _NerifToken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NerifTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NerifToken contract.
type NerifTokenOwnershipTransferredIterator struct {
	Event *NerifTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NerifTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifTokenOwnershipTransferred)
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
		it.Event = new(NerifTokenOwnershipTransferred)
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
func (it *NerifTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifTokenOwnershipTransferred represents a OwnershipTransferred event raised by the NerifToken contract.
type NerifTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NerifToken *NerifTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NerifTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NerifToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NerifTokenOwnershipTransferredIterator{contract: _NerifToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NerifToken *NerifTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NerifTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NerifToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifTokenOwnershipTransferred)
				if err := _NerifToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NerifToken *NerifTokenFilterer) ParseOwnershipTransferred(log types.Log) (*NerifTokenOwnershipTransferred, error) {
	event := new(NerifTokenOwnershipTransferred)
	if err := _NerifToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NerifTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NerifToken contract.
type NerifTokenTransferIterator struct {
	Event *NerifTokenTransfer // Event containing the contract specifics and raw log

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
func (it *NerifTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NerifTokenTransfer)
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
		it.Event = new(NerifTokenTransfer)
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
func (it *NerifTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NerifTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NerifTokenTransfer represents a Transfer event raised by the NerifToken contract.
type NerifTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NerifToken *NerifTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NerifTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NerifToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NerifTokenTransferIterator{contract: _NerifToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NerifToken *NerifTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NerifTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NerifToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NerifTokenTransfer)
				if err := _NerifToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_NerifToken *NerifTokenFilterer) ParseTransfer(log types.Log) (*NerifTokenTransfer, error) {
	event := new(NerifTokenTransfer)
	if err := _NerifToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
