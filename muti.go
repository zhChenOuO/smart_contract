// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Target   common.Address
	CallData []byte
}

// MultCallMetaData contains all meta data concerning the MultCall contract.
var MultCallMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"name\":\"difficulty\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"name\":\"coinbase\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MultCallABI is the input ABI used to generate the binding from.
// Deprecated: Use MultCallMetaData.ABI instead.
var MultCallABI = MultCallMetaData.ABI

// MultCall is an auto generated Go binding around an Ethereum contract.
type MultCall struct {
	MultCallCaller     // Read-only binding to the contract
	MultCallTransactor // Write-only binding to the contract
	MultCallFilterer   // Log filterer for contract events
}

// MultCallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultCallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultCallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultCallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultCallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultCallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultCallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultCallSession struct {
	Contract     *MultCall         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultCallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultCallCallerSession struct {
	Contract *MultCallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MultCallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultCallTransactorSession struct {
	Contract     *MultCallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MultCallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultCallRaw struct {
	Contract *MultCall // Generic contract binding to access the raw methods on
}

// MultCallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultCallCallerRaw struct {
	Contract *MultCallCaller // Generic read-only contract binding to access the raw methods on
}

// MultCallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultCallTransactorRaw struct {
	Contract *MultCallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultCall creates a new instance of MultCall, bound to a specific deployed contract.
func NewMultCall(address common.Address, backend bind.ContractBackend) (*MultCall, error) {
	contract, err := bindMultCall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultCall{MultCallCaller: MultCallCaller{contract: contract}, MultCallTransactor: MultCallTransactor{contract: contract}, MultCallFilterer: MultCallFilterer{contract: contract}}, nil
}

// NewMultCallCaller creates a new read-only instance of MultCall, bound to a specific deployed contract.
func NewMultCallCaller(address common.Address, caller bind.ContractCaller) (*MultCallCaller, error) {
	contract, err := bindMultCall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultCallCaller{contract: contract}, nil
}

// NewMultCallTransactor creates a new write-only instance of MultCall, bound to a specific deployed contract.
func NewMultCallTransactor(address common.Address, transactor bind.ContractTransactor) (*MultCallTransactor, error) {
	contract, err := bindMultCall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultCallTransactor{contract: contract}, nil
}

// NewMultCallFilterer creates a new log filterer instance of MultCall, bound to a specific deployed contract.
func NewMultCallFilterer(address common.Address, filterer bind.ContractFilterer) (*MultCallFilterer, error) {
	contract, err := bindMultCall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultCallFilterer{contract: contract}, nil
}

// bindMultCall binds a generic wrapper to an already deployed contract.
func bindMultCall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultCallABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultCall *MultCallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultCall.Contract.MultCallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultCall *MultCallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultCall.Contract.MultCallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultCall *MultCallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultCall.Contract.MultCallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultCall *MultCallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultCall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultCall *MultCallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultCall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultCall *MultCallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultCall.Contract.contract.Transact(opts, method, params...)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultCall *MultCallCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MultCall.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultCall *MultCallSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MultCall.Contract.GetBlockHash(&_MultCall.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MultCall *MultCallCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MultCall.Contract.GetBlockHash(&_MultCall.CallOpts, blockNumber)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultCall *MultCallCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MultCall.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultCall *MultCallSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MultCall.Contract.GetCurrentBlockCoinbase(&_MultCall.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MultCall *MultCallCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MultCall.Contract.GetCurrentBlockCoinbase(&_MultCall.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultCall *MultCallCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultCall.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultCall *MultCallSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MultCall.Contract.GetCurrentBlockDifficulty(&_MultCall.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MultCall *MultCallCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MultCall.Contract.GetCurrentBlockDifficulty(&_MultCall.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultCall *MultCallCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultCall.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultCall *MultCallSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MultCall.Contract.GetCurrentBlockGasLimit(&_MultCall.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MultCall *MultCallCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MultCall.Contract.GetCurrentBlockGasLimit(&_MultCall.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultCall *MultCallCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultCall.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultCall *MultCallSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MultCall.Contract.GetCurrentBlockTimestamp(&_MultCall.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MultCall *MultCallCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MultCall.Contract.GetCurrentBlockTimestamp(&_MultCall.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultCall *MultCallCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MultCall.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultCall *MultCallSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MultCall.Contract.GetEthBalance(&_MultCall.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MultCall *MultCallCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MultCall.Contract.GetEthBalance(&_MultCall.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultCall *MultCallCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MultCall.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultCall *MultCallSession) GetLastBlockHash() ([32]byte, error) {
	return _MultCall.Contract.GetLastBlockHash(&_MultCall.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MultCall *MultCallCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MultCall.Contract.GetLastBlockHash(&_MultCall.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultCall *MultCallTransactor) Aggregate(opts *bind.TransactOpts, calls []Struct0) (*types.Transaction, error) {
	return _MultCall.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultCall *MultCallSession) Aggregate(calls []Struct0) (*types.Transaction, error) {
	return _MultCall.Contract.Aggregate(&_MultCall.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) returns(uint256 blockNumber, bytes[] returnData)
func (_MultCall *MultCallTransactorSession) Aggregate(calls []Struct0) (*types.Transaction, error) {
	return _MultCall.Contract.Aggregate(&_MultCall.TransactOpts, calls)
}
