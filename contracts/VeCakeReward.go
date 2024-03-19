// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// VeCakeRewardMetaData contains all meta data concerning the VeCakeReward contract.
var VeCakeRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_revenueSharingPools\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"}],\"name\":\"claimMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_revenueSharingPools\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"}],\"name\":\"claimMultipleWithoutProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VeCakeRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use VeCakeRewardMetaData.ABI instead.
var VeCakeRewardABI = VeCakeRewardMetaData.ABI

// VeCakeReward is an auto generated Go binding around an Ethereum contract.
type VeCakeReward struct {
	VeCakeRewardCaller     // Read-only binding to the contract
	VeCakeRewardTransactor // Write-only binding to the contract
	VeCakeRewardFilterer   // Log filterer for contract events
}

// VeCakeRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeCakeRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeCakeRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeCakeRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeCakeRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeCakeRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeCakeRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeCakeRewardSession struct {
	Contract     *VeCakeReward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeCakeRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeCakeRewardCallerSession struct {
	Contract *VeCakeRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VeCakeRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeCakeRewardTransactorSession struct {
	Contract     *VeCakeRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VeCakeRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeCakeRewardRaw struct {
	Contract *VeCakeReward // Generic contract binding to access the raw methods on
}

// VeCakeRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeCakeRewardCallerRaw struct {
	Contract *VeCakeRewardCaller // Generic read-only contract binding to access the raw methods on
}

// VeCakeRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeCakeRewardTransactorRaw struct {
	Contract *VeCakeRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVeCakeReward creates a new instance of VeCakeReward, bound to a specific deployed contract.
func NewVeCakeReward(address common.Address, backend bind.ContractBackend) (*VeCakeReward, error) {
	contract, err := bindVeCakeReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VeCakeReward{VeCakeRewardCaller: VeCakeRewardCaller{contract: contract}, VeCakeRewardTransactor: VeCakeRewardTransactor{contract: contract}, VeCakeRewardFilterer: VeCakeRewardFilterer{contract: contract}}, nil
}

// NewVeCakeRewardCaller creates a new read-only instance of VeCakeReward, bound to a specific deployed contract.
func NewVeCakeRewardCaller(address common.Address, caller bind.ContractCaller) (*VeCakeRewardCaller, error) {
	contract, err := bindVeCakeReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeCakeRewardCaller{contract: contract}, nil
}

// NewVeCakeRewardTransactor creates a new write-only instance of VeCakeReward, bound to a specific deployed contract.
func NewVeCakeRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*VeCakeRewardTransactor, error) {
	contract, err := bindVeCakeReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeCakeRewardTransactor{contract: contract}, nil
}

// NewVeCakeRewardFilterer creates a new log filterer instance of VeCakeReward, bound to a specific deployed contract.
func NewVeCakeRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*VeCakeRewardFilterer, error) {
	contract, err := bindVeCakeReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeCakeRewardFilterer{contract: contract}, nil
}

// bindVeCakeReward binds a generic wrapper to an already deployed contract.
func bindVeCakeReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VeCakeRewardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeCakeReward *VeCakeRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeCakeReward.Contract.VeCakeRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeCakeReward *VeCakeRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeCakeReward.Contract.VeCakeRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeCakeReward *VeCakeRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeCakeReward.Contract.VeCakeRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VeCakeReward *VeCakeRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VeCakeReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VeCakeReward *VeCakeRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VeCakeReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VeCakeReward *VeCakeRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VeCakeReward.Contract.contract.Transact(opts, method, params...)
}

// ClaimMultiple is a paid mutator transaction binding the contract method 0x39184476.
//
// Solidity: function claimMultiple(address[] _revenueSharingPools, address _for) returns()
func (_VeCakeReward *VeCakeRewardTransactor) ClaimMultiple(opts *bind.TransactOpts, _revenueSharingPools []common.Address, _for common.Address) (*types.Transaction, error) {
	return _VeCakeReward.contract.Transact(opts, "claimMultiple", _revenueSharingPools, _for)
}

// ClaimMultiple is a paid mutator transaction binding the contract method 0x39184476.
//
// Solidity: function claimMultiple(address[] _revenueSharingPools, address _for) returns()
func (_VeCakeReward *VeCakeRewardSession) ClaimMultiple(_revenueSharingPools []common.Address, _for common.Address) (*types.Transaction, error) {
	return _VeCakeReward.Contract.ClaimMultiple(&_VeCakeReward.TransactOpts, _revenueSharingPools, _for)
}

// ClaimMultiple is a paid mutator transaction binding the contract method 0x39184476.
//
// Solidity: function claimMultiple(address[] _revenueSharingPools, address _for) returns()
func (_VeCakeReward *VeCakeRewardTransactorSession) ClaimMultiple(_revenueSharingPools []common.Address, _for common.Address) (*types.Transaction, error) {
	return _VeCakeReward.Contract.ClaimMultiple(&_VeCakeReward.TransactOpts, _revenueSharingPools, _for)
}

// ClaimMultipleWithoutProxy is a paid mutator transaction binding the contract method 0xc5bc4a71.
//
// Solidity: function claimMultipleWithoutProxy(address[] _revenueSharingPools, address _for) returns()
func (_VeCakeReward *VeCakeRewardTransactor) ClaimMultipleWithoutProxy(opts *bind.TransactOpts, _revenueSharingPools []common.Address, _for common.Address) (*types.Transaction, error) {
	return _VeCakeReward.contract.Transact(opts, "claimMultipleWithoutProxy", _revenueSharingPools, _for)
}

// ClaimMultipleWithoutProxy is a paid mutator transaction binding the contract method 0xc5bc4a71.
//
// Solidity: function claimMultipleWithoutProxy(address[] _revenueSharingPools, address _for) returns()
func (_VeCakeReward *VeCakeRewardSession) ClaimMultipleWithoutProxy(_revenueSharingPools []common.Address, _for common.Address) (*types.Transaction, error) {
	return _VeCakeReward.Contract.ClaimMultipleWithoutProxy(&_VeCakeReward.TransactOpts, _revenueSharingPools, _for)
}

// ClaimMultipleWithoutProxy is a paid mutator transaction binding the contract method 0xc5bc4a71.
//
// Solidity: function claimMultipleWithoutProxy(address[] _revenueSharingPools, address _for) returns()
func (_VeCakeReward *VeCakeRewardTransactorSession) ClaimMultipleWithoutProxy(_revenueSharingPools []common.Address, _for common.Address) (*types.Transaction, error) {
	return _VeCakeReward.Contract.ClaimMultipleWithoutProxy(&_VeCakeReward.TransactOpts, _revenueSharingPools, _for)
}
