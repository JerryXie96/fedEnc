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
	A []byte
	B []byte
}

// TestABIABI is the input ABI used to generate the binding from.
const TestABIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"bytes\"},{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"compareOri\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"a\",\"type\":\"bytes\"},{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"d\",\"type\":\"tuple\"}],\"name\":\"compare\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"bytes\"},{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"storeOri\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reIsFit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"a\",\"type\":\"bytes\"},{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"d\",\"type\":\"tuple\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TestABI is an auto generated Go binding around an Ethereum contract.
type TestABI struct {
	TestABICaller     // Read-only binding to the contract
	TestABITransactor // Write-only binding to the contract
	TestABIFilterer   // Log filterer for contract events
}

// TestABICaller is an auto generated read-only Go binding around an Ethereum contract.
type TestABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestABISession struct {
	Contract     *TestABI          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestABICallerSession struct {
	Contract *TestABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TestABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestABITransactorSession struct {
	Contract     *TestABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TestABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestABIRaw struct {
	Contract *TestABI // Generic contract binding to access the raw methods on
}

// TestABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestABICallerRaw struct {
	Contract *TestABICaller // Generic read-only contract binding to access the raw methods on
}

// TestABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestABITransactorRaw struct {
	Contract *TestABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestABI creates a new instance of TestABI, bound to a specific deployed contract.
func NewTestABI(address common.Address, backend bind.ContractBackend) (*TestABI, error) {
	contract, err := bindTestABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestABI{TestABICaller: TestABICaller{contract: contract}, TestABITransactor: TestABITransactor{contract: contract}, TestABIFilterer: TestABIFilterer{contract: contract}}, nil
}

// NewTestABICaller creates a new read-only instance of TestABI, bound to a specific deployed contract.
func NewTestABICaller(address common.Address, caller bind.ContractCaller) (*TestABICaller, error) {
	contract, err := bindTestABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestABICaller{contract: contract}, nil
}

// NewTestABITransactor creates a new write-only instance of TestABI, bound to a specific deployed contract.
func NewTestABITransactor(address common.Address, transactor bind.ContractTransactor) (*TestABITransactor, error) {
	contract, err := bindTestABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestABITransactor{contract: contract}, nil
}

// NewTestABIFilterer creates a new log filterer instance of TestABI, bound to a specific deployed contract.
func NewTestABIFilterer(address common.Address, filterer bind.ContractFilterer) (*TestABIFilterer, error) {
	contract, err := bindTestABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestABIFilterer{contract: contract}, nil
}

// bindTestABI binds a generic wrapper to an already deployed contract.
func bindTestABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestABI *TestABIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestABI.Contract.TestABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestABI *TestABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestABI.Contract.TestABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestABI *TestABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestABI.Contract.TestABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestABI *TestABICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TestABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestABI *TestABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestABI *TestABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestABI.Contract.contract.Transact(opts, method, params...)
}

// ReIsFit is a free data retrieval call binding the contract method 0xd05c4601.
//
// Solidity: function reIsFit() constant returns(bool)
func (_TestABI *TestABICaller) ReIsFit(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TestABI.contract.Call(opts, out, "reIsFit")
	return *ret0, err
}

// ReIsFit is a free data retrieval call binding the contract method 0xd05c4601.
//
// Solidity: function reIsFit() constant returns(bool)
func (_TestABI *TestABISession) ReIsFit() (bool, error) {
	return _TestABI.Contract.ReIsFit(&_TestABI.CallOpts)
}

// ReIsFit is a free data retrieval call binding the contract method 0xd05c4601.
//
// Solidity: function reIsFit() constant returns(bool)
func (_TestABI *TestABICallerSession) ReIsFit() (bool, error) {
	return _TestABI.Contract.ReIsFit(&_TestABI.CallOpts)
}

// Compare is a paid mutator transaction binding the contract method 0x43337bd8.
//
// Solidity: function compare(Struct0 d) returns()
func (_TestABI *TestABITransactor) Compare(opts *bind.TransactOpts, d Struct0) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "compare", d)
}

// Compare is a paid mutator transaction binding the contract method 0x43337bd8.
//
// Solidity: function compare(Struct0 d) returns()
func (_TestABI *TestABISession) Compare(d Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Compare(&_TestABI.TransactOpts, d)
}

// Compare is a paid mutator transaction binding the contract method 0x43337bd8.
//
// Solidity: function compare(Struct0 d) returns()
func (_TestABI *TestABITransactorSession) Compare(d Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Compare(&_TestABI.TransactOpts, d)
}

// CompareOri is a paid mutator transaction binding the contract method 0x17305b69.
//
// Solidity: function compareOri(bytes a, bytes b) returns()
func (_TestABI *TestABITransactor) CompareOri(opts *bind.TransactOpts, a []byte, b []byte) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "compareOri", a, b)
}

// CompareOri is a paid mutator transaction binding the contract method 0x17305b69.
//
// Solidity: function compareOri(bytes a, bytes b) returns()
func (_TestABI *TestABISession) CompareOri(a []byte, b []byte) (*types.Transaction, error) {
	return _TestABI.Contract.CompareOri(&_TestABI.TransactOpts, a, b)
}

// CompareOri is a paid mutator transaction binding the contract method 0x17305b69.
//
// Solidity: function compareOri(bytes a, bytes b) returns()
func (_TestABI *TestABITransactorSession) CompareOri(a []byte, b []byte) (*types.Transaction, error) {
	return _TestABI.Contract.CompareOri(&_TestABI.TransactOpts, a, b)
}

// Store is a paid mutator transaction binding the contract method 0xf87027c8.
//
// Solidity: function store(Struct0 d) returns()
func (_TestABI *TestABITransactor) Store(opts *bind.TransactOpts, d Struct0) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "store", d)
}

// Store is a paid mutator transaction binding the contract method 0xf87027c8.
//
// Solidity: function store(Struct0 d) returns()
func (_TestABI *TestABISession) Store(d Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Store(&_TestABI.TransactOpts, d)
}

// Store is a paid mutator transaction binding the contract method 0xf87027c8.
//
// Solidity: function store(Struct0 d) returns()
func (_TestABI *TestABITransactorSession) Store(d Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Store(&_TestABI.TransactOpts, d)
}

// StoreOri is a paid mutator transaction binding the contract method 0x87ff29f4.
//
// Solidity: function storeOri(bytes a, bytes b) returns()
func (_TestABI *TestABITransactor) StoreOri(opts *bind.TransactOpts, a []byte, b []byte) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "storeOri", a, b)
}

// StoreOri is a paid mutator transaction binding the contract method 0x87ff29f4.
//
// Solidity: function storeOri(bytes a, bytes b) returns()
func (_TestABI *TestABISession) StoreOri(a []byte, b []byte) (*types.Transaction, error) {
	return _TestABI.Contract.StoreOri(&_TestABI.TransactOpts, a, b)
}

// StoreOri is a paid mutator transaction binding the contract method 0x87ff29f4.
//
// Solidity: function storeOri(bytes a, bytes b) returns()
func (_TestABI *TestABITransactorSession) StoreOri(a []byte, b []byte) (*types.Transaction, error) {
	return _TestABI.Contract.StoreOri(&_TestABI.TransactOpts, a, b)
}
