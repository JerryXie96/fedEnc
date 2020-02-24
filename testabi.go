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

// TestABIABI is the input ABI used to generate the binding from.
const TestABIABI = "[{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"x\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"bytes\"}],\"name\":\"ct\",\"type\":\"tuple\"}],\"name\":\"search\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"bytes\"}],\"name\":\"searchOri\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"bytes\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"storeOri\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIsPair\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"clearRes\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"x\",\"type\":\"bytes\"},{\"name\":\"y\",\"type\":\"bytes\"}],\"name\":\"ct\",\"type\":\"tuple\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// GetIsPair is a free data retrieval call binding the contract method 0x85c8656b.
//
// Solidity: function getIsPair() constant returns(bytes)
func (_TestABI *TestABICaller) GetIsPair(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _TestABI.contract.Call(opts, out, "getIsPair")
	return *ret0, err
}

// GetIsPair is a free data retrieval call binding the contract method 0x85c8656b.
//
// Solidity: function getIsPair() constant returns(bytes)
func (_TestABI *TestABISession) GetIsPair() ([]byte, error) {
	return _TestABI.Contract.GetIsPair(&_TestABI.CallOpts)
}

// GetIsPair is a free data retrieval call binding the contract method 0x85c8656b.
//
// Solidity: function getIsPair() constant returns(bytes)
func (_TestABI *TestABICallerSession) GetIsPair() ([]byte, error) {
	return _TestABI.Contract.GetIsPair(&_TestABI.CallOpts)
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

// ClearRes is a paid mutator transaction binding the contract method 0xe1bd81ba.
//
// Solidity: function clearRes() returns()
func (_TestABI *TestABITransactor) ClearRes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "clearRes")
}

// ClearRes is a paid mutator transaction binding the contract method 0xe1bd81ba.
//
// Solidity: function clearRes() returns()
func (_TestABI *TestABISession) ClearRes() (*types.Transaction, error) {
	return _TestABI.Contract.ClearRes(&_TestABI.TransactOpts)
}

// ClearRes is a paid mutator transaction binding the contract method 0xe1bd81ba.
//
// Solidity: function clearRes() returns()
func (_TestABI *TestABITransactorSession) ClearRes() (*types.Transaction, error) {
	return _TestABI.Contract.ClearRes(&_TestABI.TransactOpts)
}

// Search is a paid mutator transaction binding the contract method 0x32fcc7d0.
//
// Solidity: function search(Struct0 ct) returns()
func (_TestABI *TestABITransactor) Search(opts *bind.TransactOpts, ct Struct0) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "search", ct)
}

// Search is a paid mutator transaction binding the contract method 0x32fcc7d0.
//
// Solidity: function search(Struct0 ct) returns()
func (_TestABI *TestABISession) Search(ct Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Search(&_TestABI.TransactOpts, ct)
}

// Search is a paid mutator transaction binding the contract method 0x32fcc7d0.
//
// Solidity: function search(Struct0 ct) returns()
func (_TestABI *TestABITransactorSession) Search(ct Struct0) (*types.Transaction, error) {
	return _TestABI.Contract.Search(&_TestABI.TransactOpts, ct)
}

// SearchOri is a paid mutator transaction binding the contract method 0x33c888e1.
//
// Solidity: function searchOri(bytes x, bytes y) returns()
func (_TestABI *TestABITransactor) SearchOri(opts *bind.TransactOpts, x []byte, y []byte) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "searchOri", x, y)
}

// SearchOri is a paid mutator transaction binding the contract method 0x33c888e1.
//
// Solidity: function searchOri(bytes x, bytes y) returns()
func (_TestABI *TestABISession) SearchOri(x []byte, y []byte) (*types.Transaction, error) {
	return _TestABI.Contract.SearchOri(&_TestABI.TransactOpts, x, y)
}

// SearchOri is a paid mutator transaction binding the contract method 0x33c888e1.
//
// Solidity: function searchOri(bytes x, bytes y) returns()
func (_TestABI *TestABITransactorSession) SearchOri(x []byte, y []byte) (*types.Transaction, error) {
	return _TestABI.Contract.SearchOri(&_TestABI.TransactOpts, x, y)
}

// Store is a paid mutator transaction binding the contract method 0xf79a7e06.
//
// Solidity: function store(Struct0 ct, uint256 id) returns()
func (_TestABI *TestABITransactor) Store(opts *bind.TransactOpts, ct Struct0, id *big.Int) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "store", ct, id)
}

// Store is a paid mutator transaction binding the contract method 0xf79a7e06.
//
// Solidity: function store(Struct0 ct, uint256 id) returns()
func (_TestABI *TestABISession) Store(ct Struct0, id *big.Int) (*types.Transaction, error) {
	return _TestABI.Contract.Store(&_TestABI.TransactOpts, ct, id)
}

// Store is a paid mutator transaction binding the contract method 0xf79a7e06.
//
// Solidity: function store(Struct0 ct, uint256 id) returns()
func (_TestABI *TestABITransactorSession) Store(ct Struct0, id *big.Int) (*types.Transaction, error) {
	return _TestABI.Contract.Store(&_TestABI.TransactOpts, ct, id)
}

// StoreOri is a paid mutator transaction binding the contract method 0x536d6cd1.
//
// Solidity: function storeOri(bytes x, bytes y, uint256 id) returns()
func (_TestABI *TestABITransactor) StoreOri(opts *bind.TransactOpts, x []byte, y []byte, id *big.Int) (*types.Transaction, error) {
	return _TestABI.contract.Transact(opts, "storeOri", x, y, id)
}

// StoreOri is a paid mutator transaction binding the contract method 0x536d6cd1.
//
// Solidity: function storeOri(bytes x, bytes y, uint256 id) returns()
func (_TestABI *TestABISession) StoreOri(x []byte, y []byte, id *big.Int) (*types.Transaction, error) {
	return _TestABI.Contract.StoreOri(&_TestABI.TransactOpts, x, y, id)
}

// StoreOri is a paid mutator transaction binding the contract method 0x536d6cd1.
//
// Solidity: function storeOri(bytes x, bytes y, uint256 id) returns()
func (_TestABI *TestABITransactorSession) StoreOri(x []byte, y []byte, id *big.Int) (*types.Transaction, error) {
	return _TestABI.Contract.StoreOri(&_TestABI.TransactOpts, x, y, id)
}
