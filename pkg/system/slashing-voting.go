// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system

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
)

// SlashingVotingMetaData contains all meta data concerning the SlashingVoting contract.
var SlashingVotingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSlashingReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"BannedWithReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"SlashedWithReason\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumSlashingReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"VotedWithReason\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DKG_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EVENT_REGISTRY_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHING_VOTING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUPPORTED_TOKENS_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bannedValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bans\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bansByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"bansByReason\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractRegistry\",\"outputs\":[{\"internalType\":\"contractContractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"epochByBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"getBannedValidatorsByReason\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getBansByEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signerGetterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorGetterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashingThresold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lashingEpochs\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"isBannedByReason\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"slashingProposalVoteCounts\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epochPeriod\",\"type\":\"uint256\"}],\"name\":\"setEpochPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashingEpochs\",\"type\":\"uint256\"}],\"name\":\"setSlashingEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_slashingThresold\",\"type\":\"uint256\"}],\"name\":\"setSlashingThresold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"shouldShash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signerGetter\",\"outputs\":[{\"internalType\":\"contractSignerGetter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashingEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slashingThresold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorGetter\",\"outputs\":[{\"internalType\":\"contractValidatorGetter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"voteCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"voteWithReason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"enumSlashingReason\",\"name\":\"_reason\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"_nonce\",\"type\":\"bytes\"}],\"name\":\"votingHashWithReason\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// SlashingVotingABI is the input ABI used to generate the binding from.
// Deprecated: Use SlashingVotingMetaData.ABI instead.
var SlashingVotingABI = SlashingVotingMetaData.ABI

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
	parsed, err := abi.JSON(strings.NewReader(SlashingVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) DKGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "DKG_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) DKGKEY() (string, error) {
	return _SlashingVoting.Contract.DKGKEY(&_SlashingVoting.CallOpts)
}

// DKGKEY is a free data retrieval call binding the contract method 0x3a9783f3.
//
// Solidity: function DKG_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) DKGKEY() (string, error) {
	return _SlashingVoting.Contract.DKGKEY(&_SlashingVoting.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) EVENTREGISTRYKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "EVENT_REGISTRY_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) EVENTREGISTRYKEY() (string, error) {
	return _SlashingVoting.Contract.EVENTREGISTRYKEY(&_SlashingVoting.CallOpts)
}

// EVENTREGISTRYKEY is a free data retrieval call binding the contract method 0x3580e192.
//
// Solidity: function EVENT_REGISTRY_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) EVENTREGISTRYKEY() (string, error) {
	return _SlashingVoting.Contract.EVENTREGISTRYKEY(&_SlashingVoting.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) SLASHINGVOTINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "SLASHING_VOTING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) SLASHINGVOTINGKEY() (string, error) {
	return _SlashingVoting.Contract.SLASHINGVOTINGKEY(&_SlashingVoting.CallOpts)
}

// SLASHINGVOTINGKEY is a free data retrieval call binding the contract method 0xfaaa8a64.
//
// Solidity: function SLASHING_VOTING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) SLASHINGVOTINGKEY() (string, error) {
	return _SlashingVoting.Contract.SLASHINGVOTINGKEY(&_SlashingVoting.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) STAKINGKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "STAKING_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) STAKINGKEY() (string, error) {
	return _SlashingVoting.Contract.STAKINGKEY(&_SlashingVoting.CallOpts)
}

// STAKINGKEY is a free data retrieval call binding the contract method 0x561ff9a9.
//
// Solidity: function STAKING_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) STAKINGKEY() (string, error) {
	return _SlashingVoting.Contract.STAKINGKEY(&_SlashingVoting.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) SUPPORTEDTOKENSKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "SUPPORTED_TOKENS_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _SlashingVoting.Contract.SUPPORTEDTOKENSKEY(&_SlashingVoting.CallOpts)
}

// SUPPORTEDTOKENSKEY is a free data retrieval call binding the contract method 0x78a5c206.
//
// Solidity: function SUPPORTED_TOKENS_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) SUPPORTEDTOKENSKEY() (string, error) {
	return _SlashingVoting.Contract.SUPPORTEDTOKENSKEY(&_SlashingVoting.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCaller) VALIDATORREWARDDISTRIBUTIONPOOLKEY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingSession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _SlashingVoting.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_SlashingVoting.CallOpts)
}

// VALIDATORREWARDDISTRIBUTIONPOOLKEY is a free data retrieval call binding the contract method 0x93107614.
//
// Solidity: function VALIDATOR_REWARD_DISTRIBUTION_POOL_KEY() view returns(string)
func (_SlashingVoting *SlashingVotingCallerSession) VALIDATORREWARDDISTRIBUTIONPOOLKEY() (string, error) {
	return _SlashingVoting.Contract.VALIDATORREWARDDISTRIBUTIONPOOLKEY(&_SlashingVoting.CallOpts)
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

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_SlashingVoting *SlashingVotingCaller) ContractRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "contractRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_SlashingVoting *SlashingVotingSession) ContractRegistry() (common.Address, error) {
	return _SlashingVoting.Contract.ContractRegistry(&_SlashingVoting.CallOpts)
}

// ContractRegistry is a free data retrieval call binding the contract method 0xabf410e5.
//
// Solidity: function contractRegistry() view returns(address)
func (_SlashingVoting *SlashingVotingCallerSession) ContractRegistry() (common.Address, error) {
	return _SlashingVoting.Contract.ContractRegistry(&_SlashingVoting.CallOpts)
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

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCaller) SignerGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "signerGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SlashingVoting *SlashingVotingSession) SignerGetter() (common.Address, error) {
	return _SlashingVoting.Contract.SignerGetter(&_SlashingVoting.CallOpts)
}

// SignerGetter is a free data retrieval call binding the contract method 0x5c211f88.
//
// Solidity: function signerGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCallerSession) SignerGetter() (common.Address, error) {
	return _SlashingVoting.Contract.SignerGetter(&_SlashingVoting.CallOpts)
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

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCaller) ValidatorGetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SlashingVoting.contract.Call(opts, &out, "validatorGetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_SlashingVoting *SlashingVotingSession) ValidatorGetter() (common.Address, error) {
	return _SlashingVoting.Contract.ValidatorGetter(&_SlashingVoting.CallOpts)
}

// ValidatorGetter is a free data retrieval call binding the contract method 0x69495ef5.
//
// Solidity: function validatorGetter() view returns(address)
func (_SlashingVoting *SlashingVotingCallerSession) ValidatorGetter() (common.Address, error) {
	return _SlashingVoting.Contract.ValidatorGetter(&_SlashingVoting.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xb1a5d12d.
//
// Solidity: function initialize(address _signerGetterAddress, address _validatorGetterAddress, uint256 _epochPeriod, uint256 _slashingThresold, uint256 _lashingEpochs, address _contractRegistry) returns()
func (_SlashingVoting *SlashingVotingTransactor) Initialize(opts *bind.TransactOpts, _signerGetterAddress common.Address, _validatorGetterAddress common.Address, _epochPeriod *big.Int, _slashingThresold *big.Int, _lashingEpochs *big.Int, _contractRegistry common.Address) (*types.Transaction, error) {
	return _SlashingVoting.contract.Transact(opts, "initialize", _signerGetterAddress, _validatorGetterAddress, _epochPeriod, _slashingThresold, _lashingEpochs, _contractRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a5d12d.
//
// Solidity: function initialize(address _signerGetterAddress, address _validatorGetterAddress, uint256 _epochPeriod, uint256 _slashingThresold, uint256 _lashingEpochs, address _contractRegistry) returns()
func (_SlashingVoting *SlashingVotingSession) Initialize(_signerGetterAddress common.Address, _validatorGetterAddress common.Address, _epochPeriod *big.Int, _slashingThresold *big.Int, _lashingEpochs *big.Int, _contractRegistry common.Address) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _signerGetterAddress, _validatorGetterAddress, _epochPeriod, _slashingThresold, _lashingEpochs, _contractRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0xb1a5d12d.
//
// Solidity: function initialize(address _signerGetterAddress, address _validatorGetterAddress, uint256 _epochPeriod, uint256 _slashingThresold, uint256 _lashingEpochs, address _contractRegistry) returns()
func (_SlashingVoting *SlashingVotingTransactorSession) Initialize(_signerGetterAddress common.Address, _validatorGetterAddress common.Address, _epochPeriod *big.Int, _slashingThresold *big.Int, _lashingEpochs *big.Int, _contractRegistry common.Address) (*types.Transaction, error) {
	return _SlashingVoting.Contract.Initialize(&_SlashingVoting.TransactOpts, _signerGetterAddress, _validatorGetterAddress, _epochPeriod, _slashingThresold, _lashingEpochs, _contractRegistry)
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
