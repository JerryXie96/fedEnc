// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	X []byte
	Y []byte
}

// EMABIABI is the input ABI used to generate the binding from.
const EMABIABI = "[{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"x\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"bytes\"}],\"name\":\"ct\",\"type\":\"tuple\"}],\"name\":\"search\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIsPair\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"clearRes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"x\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"bytes\"}],\"name\":\"ct\",\"type\":\"tuple\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EMABI is an auto generated Go binding around an Ethereum contract.
type EMABI struct {
	EMABICaller     // Read-only binding to the contract
	EMABITransactor // Write-only binding to the contract
	EMABIFilterer   // Log filterer for contract events
}

// EMABICaller is an auto generated read-only Go binding around an Ethereum contract.
type EMABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EMABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type EMABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EMABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EMABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EMABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EMABISession struct {
	Contract     *EMABI            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EMABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EMABICallerSession struct {
	Contract *EMABICaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EMABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EMABITransactorSession struct {
	Contract     *EMABITransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EMABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type EMABIRaw struct {
	Contract *EMABI // Generic contract binding to access the raw methods on
}

// EMABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EMABICallerRaw struct {
	Contract *EMABICaller // Generic read-only contract binding to access the raw methods on
}

// EMABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EMABITransactorRaw struct {
	Contract *EMABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewEMABI creates a new instance of EMABI, bound to a specific deployed contract.
func NewEMABI(address common.Address, backend bind.ContractBackend) (*EMABI, error) {
	contract, err := bindEMABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EMABI{EMABICaller: EMABICaller{contract: contract}, EMABITransactor: EMABITransactor{contract: contract}, EMABIFilterer: EMABIFilterer{contract: contract}}, nil
}

// NewEMABICaller creates a new read-only instance of EMABI, bound to a specific deployed contract.
func NewEMABICaller(address common.Address, caller bind.ContractCaller) (*EMABICaller, error) {
	contract, err := bindEMABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EMABICaller{contract: contract}, nil
}

// NewEMABITransactor creates a new write-only instance of EMABI, bound to a specific deployed contract.
func NewEMABITransactor(address common.Address, transactor bind.ContractTransactor) (*EMABITransactor, error) {
	contract, err := bindEMABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EMABITransactor{contract: contract}, nil
}

// NewEMABIFilterer creates a new log filterer instance of EMABI, bound to a specific deployed contract.
func NewEMABIFilterer(address common.Address, filterer bind.ContractFilterer) (*EMABIFilterer, error) {
	contract, err := bindEMABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EMABIFilterer{contract: contract}, nil
}

// bindEMABI binds a generic wrapper to an already deployed contract.
func bindEMABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EMABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EMABI *EMABIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EMABI.Contract.EMABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EMABI *EMABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EMABI.Contract.EMABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EMABI *EMABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EMABI.Contract.EMABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EMABI *EMABICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EMABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EMABI *EMABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EMABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EMABI *EMABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EMABI.Contract.contract.Transact(opts, method, params...)
}

// GetIsPair is a free data retrieval call binding the contract method 0x85c8656b.
//
// Solidity: function getIsPair() constant returns(bool)
func (_EMABI *EMABICaller) GetIsPair(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EMABI.contract.Call(opts, out, "getIsPair")
	return *ret0, err
}

// GetIsPair is a free data retrieval call binding the contract method 0x85c8656b.
//
// Solidity: function getIsPair() constant returns(bool)
func (_EMABI *EMABISession) GetIsPair() (bool, error) {
	return _EMABI.Contract.GetIsPair(&_EMABI.CallOpts)
}

// GetIsPair is a free data retrieval call binding the contract method 0x85c8656b.
//
// Solidity: function getIsPair() constant returns(bool)
func (_EMABI *EMABICallerSession) GetIsPair() (bool, error) {
	return _EMABI.Contract.GetIsPair(&_EMABI.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() constant returns(uint256[])
func (_EMABI *EMABICaller) GetResult(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _EMABI.contract.Call(opts, out, "getResult")
	return *ret0, err
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() constant returns(uint256[])
func (_EMABI *EMABISession) GetResult() ([]*big.Int, error) {
	return _EMABI.Contract.GetResult(&_EMABI.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() constant returns(uint256[])
func (_EMABI *EMABICallerSession) GetResult() ([]*big.Int, error) {
	return _EMABI.Contract.GetResult(&_EMABI.CallOpts)
}

// ClearRes is a paid mutator transaction binding the contract method 0xe1bd81ba.
//
// Solidity: function clearRes() returns()
func (_EMABI *EMABITransactor) ClearRes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EMABI.contract.Transact(opts, "clearRes")
}

// ClearRes is a paid mutator transaction binding the contract method 0xe1bd81ba.
//
// Solidity: function clearRes() returns()
func (_EMABI *EMABISession) ClearRes() (*types.Transaction, error) {
	return _EMABI.Contract.ClearRes(&_EMABI.TransactOpts)
}

// ClearRes is a paid mutator transaction binding the contract method 0xe1bd81ba.
//
// Solidity: function clearRes() returns()
func (_EMABI *EMABITransactorSession) ClearRes() (*types.Transaction, error) {
	return _EMABI.Contract.ClearRes(&_EMABI.TransactOpts)
}

// Search is a paid mutator transaction binding the contract method 0x32fcc7d0.
//
// Solidity: function search(Struct0 ct) returns()
func (_EMABI *EMABITransactor) Search(opts *bind.TransactOpts, ct Struct0) (*types.Transaction, error) {
	return _EMABI.contract.Transact(opts, "search", ct)
}

// Search is a paid mutator transaction binding the contract method 0x32fcc7d0.
//
// Solidity: function search(Struct0 ct) returns()
func (_EMABI *EMABISession) Search(ct Struct0) (*types.Transaction, error) {
	return _EMABI.Contract.Search(&_EMABI.TransactOpts, ct)
}

// Search is a paid mutator transaction binding the contract method 0x32fcc7d0.
//
// Solidity: function search(Struct0 ct) returns()
func (_EMABI *EMABITransactorSession) Search(ct Struct0) (*types.Transaction, error) {
	return _EMABI.Contract.Search(&_EMABI.TransactOpts, ct)
}

// Store is a paid mutator transaction binding the contract method 0xf79a7e06.
//
// Solidity: function store(Struct0 ct, uint256 id) returns()
func (_EMABI *EMABITransactor) Store(opts *bind.TransactOpts, ct Struct0, id *big.Int) (*types.Transaction, error) {
	return _EMABI.contract.Transact(opts, "store", ct, id)
}

// Store is a paid mutator transaction binding the contract method 0xf79a7e06.
//
// Solidity: function store(Struct0 ct, uint256 id) returns()
func (_EMABI *EMABISession) Store(ct Struct0, id *big.Int) (*types.Transaction, error) {
	return _EMABI.Contract.Store(&_EMABI.TransactOpts, ct, id)
}

// Store is a paid mutator transaction binding the contract method 0xf79a7e06.
//
// Solidity: function store(Struct0 ct, uint256 id) returns()
func (_EMABI *EMABITransactorSession) Store(ct Struct0, id *big.Int) (*types.Transaction, error) {
	return _EMABI.Contract.Store(&_EMABI.TransactOpts, ct, id)
}
