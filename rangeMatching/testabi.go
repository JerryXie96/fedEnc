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

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Cipher [32][2]Struct1
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	C1 []byte
	C2 []byte
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Cipher [32][2][]byte
	Id     *big.Int
}

// TestABIABI is the input ABI used to generate the binding from.
const TestABIABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"clearResult\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"cipher\",\"type\":\"bytes[2][32]\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"indexItem\",\"type\":\"tuple[]\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"components\":[{\"name\":\"c1\",\"type\":\"bytes\"},{\"name\":\"c2\",\"type\":\"bytes\"}],\"name\":\"cipher\",\"type\":\"tuple[2][32]\"}],\"name\":\"query\",\"type\":\"tuple\"}],\"name\":\"search\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() constant returns(uint256[])
func (_TestABI *TestABICaller) GetResult(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _TestABI.contract.Call(opts, out, "getResult")
	return *ret0, err
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() constant returns(uint256[])
func (_TestABI *TestABISession) GetResult() ([]*big.Int, error) {
	return _TestABI.Contract.GetResult(&_TestABI.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() constant returns(uint256[])
func (_TestABI *TestABICallerSession) GetResult() ([]*big.Int, error) {
	return _TestABI.Contract.GetResult(&_TestABI.CallOpts)
}

// ClearResult is a paid mutator transaction binding the contract method 0x6765350a.
//
// Solidity: function clearResult() returns()
func (_TestABI *TestABITransactor) ClearResult(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "clearResult")
}

// ClearResult is a paid mutator transaction binding the contract method 0x6765350a.
//
// Solidity: function clearResult() returns()
func (_TestABI *TestABISession) ClearResult() (*types.Transaction, error) {
	return _TestABI.Contract.ClearResult(&_TestABI.TransactOpts)
}

// ClearResult is a paid mutator transaction binding the contract method 0x6765350a.
//
// Solidity: function clearResult() returns()
func (_TestABI *TestABITransactorSession) ClearResult() (*types.Transaction, error) {
	return _TestABI.Contract.ClearResult(&_TestABI.TransactOpts)
}

// Search is a paid mutator transaction binding the contract method 0xd0b39c81.
//
// Solidity: function search(Struct2 query) returns()
func (_TestABI *TestABITransactor) Search(opts *bind.TransactOpts, query Struct2) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "search", query)
}

// Search is a paid mutator transaction binding the contract method 0xd0b39c81.
//
// Solidity: function search(Struct2 query) returns()
func (_TestABI *TestABISession) Search(query Struct2) (*types.Transaction, error) {
	return _TestABI.Contract.Search(&_TestABI.TransactOpts, query)
}

// Search is a paid mutator transaction binding the contract method 0xd0b39c81.
//
// Solidity: function search(Struct2 query) returns()
func (_TestABI *TestABITransactorSession) Search(query Struct2) (*types.Transaction, error) {
	return _TestABI.Contract.Search(&_TestABI.TransactOpts, query)
}

// Store is a paid mutator transaction binding the contract method 0xb46c1989.
//
// Solidity: function store([]Struct0 indexItem) returns()
func (_TestABI *TestABITransactor) Store(opts *bind.TransactOpts, indexItem []Struct0) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "store", indexItem)
}

// Store is a paid mutator transaction binding the contract method 0xb46c1989.
//
// Solidity: function store([]Struct0 indexItem) returns()
func (_TestABI *TestABISession) Store(indexItem []Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Store(&_TestABI.TransactOpts, indexItem)
}

// Store is a paid mutator transaction binding the contract method 0xb46c1989.
//
// Solidity: function store([]Struct0 indexItem) returns()
func (_TestABI *TestABITransactorSession) Store(indexItem []Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Store(&_TestABI.TransactOpts, indexItem)
}
