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
)

// MyContractMetaData contains all meta data concerning the MyContract contract.
var MyContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"counter\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getYourCounter\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610412806100616000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806333c4b3ef14610051578063c12327b81461006f578063d732d9551461009f578063e8927fbc146100a9575b600080fd5b6100596100b3565b6040516100669190610212565b60405180910390f35b610089600480360381019061008491906101d6565b6100f9565b6040516100969190610212565b60405180910390f35b6100a7610111565b005b6100b1610169565b005b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b60006020528060005260406000206000915090505481565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461016091906102c1565b92505081905550565b60016000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546101b8919061022d565b92505081905550565b6000813590506101d0816103c5565b92915050565b6000602082840312156101ec576101eb6103c0565b5b60006101fa848285016101c1565b91505092915050565b61020c81610367565b82525050565b60006020820190506102276000830184610203565b92915050565b600061023882610367565b915061024383610367565b9250817f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0383136000831215161561027e5761027d610391565b5b817f80000000000000000000000000000000000000000000000000000000000000000383126000831216156102b6576102b5610391565b5b828201905092915050565b60006102cc82610367565b91506102d783610367565b9250827f80000000000000000000000000000000000000000000000000000000000000000182126000841215161561031257610311610391565b5b827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01821360008412161561034a57610349610391565b5b828203905092915050565b600061036082610371565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b6103ce81610355565b81146103d957600080fd5b5056fea2646970667358221220de1f7fa3110ea59960df61dc0437ebddc24790014862795ab5fac32b6931cbd964736f6c63430008070033",
}

// MyContractABI is the input ABI used to generate the binding from.
// Deprecated: Use MyContractMetaData.ABI instead.
var MyContractABI = MyContractMetaData.ABI

// MyContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyContractMetaData.Bin instead.
var MyContractBin = MyContractMetaData.Bin

// DeployMyContract deploys a new Ethereum contract, binding an instance of MyContract to it.
func DeployMyContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyContract, error) {
	parsed, err := MyContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// MyContract is an auto generated Go binding around an Ethereum contract.
type MyContract struct {
	MyContractCaller     // Read-only binding to the contract
	MyContractTransactor // Write-only binding to the contract
	MyContractFilterer   // Log filterer for contract events
}

// MyContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyContractSession struct {
	Contract     *MyContract       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyContractCallerSession struct {
	Contract *MyContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MyContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyContractTransactorSession struct {
	Contract     *MyContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MyContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyContractRaw struct {
	Contract *MyContract // Generic contract binding to access the raw methods on
}

// MyContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyContractCallerRaw struct {
	Contract *MyContractCaller // Generic read-only contract binding to access the raw methods on
}

// MyContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyContractTransactorRaw struct {
	Contract *MyContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyContract creates a new instance of MyContract, bound to a specific deployed contract.
func NewMyContract(address common.Address, backend bind.ContractBackend) (*MyContract, error) {
	contract, err := bindMyContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyContract{MyContractCaller: MyContractCaller{contract: contract}, MyContractTransactor: MyContractTransactor{contract: contract}, MyContractFilterer: MyContractFilterer{contract: contract}}, nil
}

// NewMyContractCaller creates a new read-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractCaller(address common.Address, caller bind.ContractCaller) (*MyContractCaller, error) {
	contract, err := bindMyContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractCaller{contract: contract}, nil
}

// NewMyContractTransactor creates a new write-only instance of MyContract, bound to a specific deployed contract.
func NewMyContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MyContractTransactor, error) {
	contract, err := bindMyContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyContractTransactor{contract: contract}, nil
}

// NewMyContractFilterer creates a new log filterer instance of MyContract, bound to a specific deployed contract.
func NewMyContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MyContractFilterer, error) {
	contract, err := bindMyContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyContractFilterer{contract: contract}, nil
}

// bindMyContract binds a generic wrapper to an already deployed contract.
func bindMyContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.MyContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.MyContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyContract *MyContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyContract *MyContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyContract *MyContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyContract.Contract.contract.Transact(opts, method, params...)
}

// Counter is a free data retrieval call binding the contract method 0xc12327b8.
//
// Solidity: function counter(address ) view returns(int256)
func (_MyContract *MyContractCaller) Counter(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "counter", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0xc12327b8.
//
// Solidity: function counter(address ) view returns(int256)
func (_MyContract *MyContractSession) Counter(arg0 common.Address) (*big.Int, error) {
	return _MyContract.Contract.Counter(&_MyContract.CallOpts, arg0)
}

// Counter is a free data retrieval call binding the contract method 0xc12327b8.
//
// Solidity: function counter(address ) view returns(int256)
func (_MyContract *MyContractCallerSession) Counter(arg0 common.Address) (*big.Int, error) {
	return _MyContract.Contract.Counter(&_MyContract.CallOpts, arg0)
}

// GetYourCounter is a free data retrieval call binding the contract method 0x33c4b3ef.
//
// Solidity: function getYourCounter() view returns(int256)
func (_MyContract *MyContractCaller) GetYourCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyContract.contract.Call(opts, &out, "getYourCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetYourCounter is a free data retrieval call binding the contract method 0x33c4b3ef.
//
// Solidity: function getYourCounter() view returns(int256)
func (_MyContract *MyContractSession) GetYourCounter() (*big.Int, error) {
	return _MyContract.Contract.GetYourCounter(&_MyContract.CallOpts)
}

// GetYourCounter is a free data retrieval call binding the contract method 0x33c4b3ef.
//
// Solidity: function getYourCounter() view returns(int256)
func (_MyContract *MyContractCallerSession) GetYourCounter() (*big.Int, error) {
	return _MyContract.Contract.GetYourCounter(&_MyContract.CallOpts)
}

// Decrease is a paid mutator transaction binding the contract method 0xd732d955.
//
// Solidity: function decrease() returns()
func (_MyContract *MyContractTransactor) Decrease(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "decrease")
}

// Decrease is a paid mutator transaction binding the contract method 0xd732d955.
//
// Solidity: function decrease() returns()
func (_MyContract *MyContractSession) Decrease() (*types.Transaction, error) {
	return _MyContract.Contract.Decrease(&_MyContract.TransactOpts)
}

// Decrease is a paid mutator transaction binding the contract method 0xd732d955.
//
// Solidity: function decrease() returns()
func (_MyContract *MyContractTransactorSession) Decrease() (*types.Transaction, error) {
	return _MyContract.Contract.Decrease(&_MyContract.TransactOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_MyContract *MyContractTransactor) Increase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyContract.contract.Transact(opts, "increase")
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_MyContract *MyContractSession) Increase() (*types.Transaction, error) {
	return _MyContract.Contract.Increase(&_MyContract.TransactOpts)
}

// Increase is a paid mutator transaction binding the contract method 0xe8927fbc.
//
// Solidity: function increase() returns()
func (_MyContract *MyContractTransactorSession) Increase() (*types.Transaction, error) {
	return _MyContract.Contract.Increase(&_MyContract.TransactOpts)
}
