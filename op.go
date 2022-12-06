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

// OpMetaData contains all meta data concerning the Op contract.
var OpMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oneInchRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_imToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fundsProvider\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_swapFeeTo\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_gasFeeTo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_prev\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_curr\",\"type\":\"bool\"}],\"name\":\"FlipRunning\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_prev\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_curr\",\"type\":\"address\"}],\"name\":\"FundsProvider\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_prev\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_curr\",\"type\":\"address\"}],\"name\":\"GasFeeTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Pull\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isWhitelist\",\"type\":\"bool\"}],\"name\":\"SetWhitelist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"id\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"uniqueId\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"enumOperator.ACTION\",\"name\":\"action\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"retAmt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmt\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_prev\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_curr\",\"type\":\"address\"}],\"name\":\"SwapFeeTo\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_uniqueId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasFeeAmt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"crossSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_uniqueId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_swapFeeAmt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"doSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipRunning\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_uniqueId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_swapFeeAmt\",\"type\":\"uint256\"}],\"name\":\"fromUCross\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasFeeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balances\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundsProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"imToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oneInchRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"pull\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFundsProvider\",\"type\":\"address\"}],\"name\":\"setFundsProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGasFeeTo\",\"type\":\"address\"}],\"name\":\"setGasFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSwapFeeTo\",\"type\":\"address\"}],\"name\":\"setSwapFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrArr\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_flags\",\"type\":\"bool[]\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFeeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_id\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_uniqueId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_amt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasFeeAmt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"toUCross\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"useless\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// OpABI is the input ABI used to generate the binding from.
// Deprecated: Use OpMetaData.ABI instead.
var OpABI = OpMetaData.ABI

// Op is an auto generated Go binding around an Ethereum contract.
type Op struct {
	OpCaller     // Read-only binding to the contract
	OpTransactor // Write-only binding to the contract
	OpFilterer   // Log filterer for contract events
}

// OpCaller is an auto generated read-only Go binding around an Ethereum contract.
type OpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OpSession struct {
	Contract     *Op               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OpCallerSession struct {
	Contract *OpCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OpTransactorSession struct {
	Contract     *OpTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OpRaw is an auto generated low-level Go binding around an Ethereum contract.
type OpRaw struct {
	Contract *Op // Generic contract binding to access the raw methods on
}

// OpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OpCallerRaw struct {
	Contract *OpCaller // Generic read-only contract binding to access the raw methods on
}

// OpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OpTransactorRaw struct {
	Contract *OpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOp creates a new instance of Op, bound to a specific deployed contract.
func NewOp(address common.Address, backend bind.ContractBackend) (*Op, error) {
	contract, err := bindOp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Op{OpCaller: OpCaller{contract: contract}, OpTransactor: OpTransactor{contract: contract}, OpFilterer: OpFilterer{contract: contract}}, nil
}

// NewOpCaller creates a new read-only instance of Op, bound to a specific deployed contract.
func NewOpCaller(address common.Address, caller bind.ContractCaller) (*OpCaller, error) {
	contract, err := bindOp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OpCaller{contract: contract}, nil
}

// NewOpTransactor creates a new write-only instance of Op, bound to a specific deployed contract.
func NewOpTransactor(address common.Address, transactor bind.ContractTransactor) (*OpTransactor, error) {
	contract, err := bindOp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OpTransactor{contract: contract}, nil
}

// NewOpFilterer creates a new log filterer instance of Op, bound to a specific deployed contract.
func NewOpFilterer(address common.Address, filterer bind.ContractFilterer) (*OpFilterer, error) {
	contract, err := bindOp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OpFilterer{contract: contract}, nil
}

// bindOp binds a generic wrapper to an already deployed contract.
func bindOp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OpABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Op *OpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Op.Contract.OpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Op *OpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Op.Contract.OpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Op *OpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Op.Contract.OpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Op *OpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Op.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Op *OpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Op.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Op *OpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Op.Contract.contract.Transact(opts, method, params...)
}

// GasFeeTo is a free data retrieval call binding the contract method 0x142d659d.
//
// Solidity: function gasFeeTo() view returns(address)
func (_Op *OpCaller) GasFeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "gasFeeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasFeeTo is a free data retrieval call binding the contract method 0x142d659d.
//
// Solidity: function gasFeeTo() view returns(address)
func (_Op *OpSession) GasFeeTo() (common.Address, error) {
	return _Op.Contract.GasFeeTo(&_Op.CallOpts)
}

// GasFeeTo is a free data retrieval call binding the contract method 0x142d659d.
//
// Solidity: function gasFeeTo() view returns(address)
func (_Op *OpCallerSession) GasFeeTo() (common.Address, error) {
	return _Op.Contract.GasFeeTo(&_Op.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x0d0f9df4.
//
// Solidity: function getBalance(address[] _tokens) view returns(uint256[] balances)
func (_Op *OpCaller) GetBalance(opts *bind.CallOpts, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "getBalance", _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x0d0f9df4.
//
// Solidity: function getBalance(address[] _tokens) view returns(uint256[] balances)
func (_Op *OpSession) GetBalance(_tokens []common.Address) ([]*big.Int, error) {
	return _Op.Contract.GetBalance(&_Op.CallOpts, _tokens)
}

// GetBalance is a free data retrieval call binding the contract method 0x0d0f9df4.
//
// Solidity: function getBalance(address[] _tokens) view returns(uint256[] balances)
func (_Op *OpCallerSession) GetBalance(_tokens []common.Address) ([]*big.Int, error) {
	return _Op.Contract.GetBalance(&_Op.CallOpts, _tokens)
}

// GetFundsProvider is a free data retrieval call binding the contract method 0xd1d791d8.
//
// Solidity: function getFundsProvider() view returns(address)
func (_Op *OpCaller) GetFundsProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "getFundsProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFundsProvider is a free data retrieval call binding the contract method 0xd1d791d8.
//
// Solidity: function getFundsProvider() view returns(address)
func (_Op *OpSession) GetFundsProvider() (common.Address, error) {
	return _Op.Contract.GetFundsProvider(&_Op.CallOpts)
}

// GetFundsProvider is a free data retrieval call binding the contract method 0xd1d791d8.
//
// Solidity: function getFundsProvider() view returns(address)
func (_Op *OpCallerSession) GetFundsProvider() (common.Address, error) {
	return _Op.Contract.GetFundsProvider(&_Op.CallOpts)
}

// ImToken is a free data retrieval call binding the contract method 0xb937378a.
//
// Solidity: function imToken() view returns(address)
func (_Op *OpCaller) ImToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "imToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ImToken is a free data retrieval call binding the contract method 0xb937378a.
//
// Solidity: function imToken() view returns(address)
func (_Op *OpSession) ImToken() (common.Address, error) {
	return _Op.Contract.ImToken(&_Op.CallOpts)
}

// ImToken is a free data retrieval call binding the contract method 0xb937378a.
//
// Solidity: function imToken() view returns(address)
func (_Op *OpCallerSession) ImToken() (common.Address, error) {
	return _Op.Contract.ImToken(&_Op.CallOpts)
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Op *OpCaller) IsRunning(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "isRunning")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Op *OpSession) IsRunning() (bool, error) {
	return _Op.Contract.IsRunning(&_Op.CallOpts)
}

// IsRunning is a free data retrieval call binding the contract method 0x2014e5d1.
//
// Solidity: function isRunning() view returns(bool)
func (_Op *OpCallerSession) IsRunning() (bool, error) {
	return _Op.Contract.IsRunning(&_Op.CallOpts)
}

// OneInchRouter is a free data retrieval call binding the contract method 0xac3af208.
//
// Solidity: function oneInchRouter() view returns(address)
func (_Op *OpCaller) OneInchRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "oneInchRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OneInchRouter is a free data retrieval call binding the contract method 0xac3af208.
//
// Solidity: function oneInchRouter() view returns(address)
func (_Op *OpSession) OneInchRouter() (common.Address, error) {
	return _Op.Contract.OneInchRouter(&_Op.CallOpts)
}

// OneInchRouter is a free data retrieval call binding the contract method 0xac3af208.
//
// Solidity: function oneInchRouter() view returns(address)
func (_Op *OpCallerSession) OneInchRouter() (common.Address, error) {
	return _Op.Contract.OneInchRouter(&_Op.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Op *OpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Op *OpSession) Owner() (common.Address, error) {
	return _Op.Contract.Owner(&_Op.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Op *OpCallerSession) Owner() (common.Address, error) {
	return _Op.Contract.Owner(&_Op.CallOpts)
}

// SwapFeeTo is a free data retrieval call binding the contract method 0x09221a0c.
//
// Solidity: function swapFeeTo() view returns(address)
func (_Op *OpCaller) SwapFeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "swapFeeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapFeeTo is a free data retrieval call binding the contract method 0x09221a0c.
//
// Solidity: function swapFeeTo() view returns(address)
func (_Op *OpSession) SwapFeeTo() (common.Address, error) {
	return _Op.Contract.SwapFeeTo(&_Op.CallOpts)
}

// SwapFeeTo is a free data retrieval call binding the contract method 0x09221a0c.
//
// Solidity: function swapFeeTo() view returns(address)
func (_Op *OpCallerSession) SwapFeeTo() (common.Address, error) {
	return _Op.Contract.SwapFeeTo(&_Op.CallOpts)
}

// Useless is a free data retrieval call binding the contract method 0xae14ef96.
//
// Solidity: function useless() pure returns(uint256)
func (_Op *OpCaller) Useless(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "useless")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Useless is a free data retrieval call binding the contract method 0xae14ef96.
//
// Solidity: function useless() pure returns(uint256)
func (_Op *OpSession) Useless() (*big.Int, error) {
	return _Op.Contract.Useless(&_Op.CallOpts)
}

// Useless is a free data retrieval call binding the contract method 0xae14ef96.
//
// Solidity: function useless() pure returns(uint256)
func (_Op *OpCallerSession) Useless() (*big.Int, error) {
	return _Op.Contract.Useless(&_Op.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Op *OpCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Op.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Op *OpSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Op.Contract.Whitelist(&_Op.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Op *OpCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Op.Contract.Whitelist(&_Op.CallOpts, arg0)
}

// CrossSwap is a paid mutator transaction binding the contract method 0xa7d4b63e.
//
// Solidity: function crossSwap(bytes _id, bytes _uniqueId, uint256 _gasFeeAmt, bytes _data) returns()
func (_Op *OpTransactor) CrossSwap(opts *bind.TransactOpts, _id []byte, _uniqueId []byte, _gasFeeAmt *big.Int, _data []byte) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "crossSwap", _id, _uniqueId, _gasFeeAmt, _data)
}

// CrossSwap is a paid mutator transaction binding the contract method 0xa7d4b63e.
//
// Solidity: function crossSwap(bytes _id, bytes _uniqueId, uint256 _gasFeeAmt, bytes _data) returns()
func (_Op *OpSession) CrossSwap(_id []byte, _uniqueId []byte, _gasFeeAmt *big.Int, _data []byte) (*types.Transaction, error) {
	return _Op.Contract.CrossSwap(&_Op.TransactOpts, _id, _uniqueId, _gasFeeAmt, _data)
}

// CrossSwap is a paid mutator transaction binding the contract method 0xa7d4b63e.
//
// Solidity: function crossSwap(bytes _id, bytes _uniqueId, uint256 _gasFeeAmt, bytes _data) returns()
func (_Op *OpTransactorSession) CrossSwap(_id []byte, _uniqueId []byte, _gasFeeAmt *big.Int, _data []byte) (*types.Transaction, error) {
	return _Op.Contract.CrossSwap(&_Op.TransactOpts, _id, _uniqueId, _gasFeeAmt, _data)
}

// DoSwap is a paid mutator transaction binding the contract method 0xbd64252c.
//
// Solidity: function doSwap(bytes _id, bytes _uniqueId, uint256 _swapFeeAmt, bytes _data) payable returns()
func (_Op *OpTransactor) DoSwap(opts *bind.TransactOpts, _id []byte, _uniqueId []byte, _swapFeeAmt *big.Int, _data []byte) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "doSwap", _id, _uniqueId, _swapFeeAmt, _data)
}

// DoSwap is a paid mutator transaction binding the contract method 0xbd64252c.
//
// Solidity: function doSwap(bytes _id, bytes _uniqueId, uint256 _swapFeeAmt, bytes _data) payable returns()
func (_Op *OpSession) DoSwap(_id []byte, _uniqueId []byte, _swapFeeAmt *big.Int, _data []byte) (*types.Transaction, error) {
	return _Op.Contract.DoSwap(&_Op.TransactOpts, _id, _uniqueId, _swapFeeAmt, _data)
}

// DoSwap is a paid mutator transaction binding the contract method 0xbd64252c.
//
// Solidity: function doSwap(bytes _id, bytes _uniqueId, uint256 _swapFeeAmt, bytes _data) payable returns()
func (_Op *OpTransactorSession) DoSwap(_id []byte, _uniqueId []byte, _swapFeeAmt *big.Int, _data []byte) (*types.Transaction, error) {
	return _Op.Contract.DoSwap(&_Op.TransactOpts, _id, _uniqueId, _swapFeeAmt, _data)
}

// FlipRunning is a paid mutator transaction binding the contract method 0x8732c858.
//
// Solidity: function flipRunning() returns()
func (_Op *OpTransactor) FlipRunning(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "flipRunning")
}

// FlipRunning is a paid mutator transaction binding the contract method 0x8732c858.
//
// Solidity: function flipRunning() returns()
func (_Op *OpSession) FlipRunning() (*types.Transaction, error) {
	return _Op.Contract.FlipRunning(&_Op.TransactOpts)
}

// FlipRunning is a paid mutator transaction binding the contract method 0x8732c858.
//
// Solidity: function flipRunning() returns()
func (_Op *OpTransactorSession) FlipRunning() (*types.Transaction, error) {
	return _Op.Contract.FlipRunning(&_Op.TransactOpts)
}

// FromUCross is a paid mutator transaction binding the contract method 0xfa37c031.
//
// Solidity: function fromUCross(bytes _id, bytes _uniqueId, uint256 _amt, uint256 _swapFeeAmt) returns()
func (_Op *OpTransactor) FromUCross(opts *bind.TransactOpts, _id []byte, _uniqueId []byte, _amt *big.Int, _swapFeeAmt *big.Int) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "fromUCross", _id, _uniqueId, _amt, _swapFeeAmt)
}

// FromUCross is a paid mutator transaction binding the contract method 0xfa37c031.
//
// Solidity: function fromUCross(bytes _id, bytes _uniqueId, uint256 _amt, uint256 _swapFeeAmt) returns()
func (_Op *OpSession) FromUCross(_id []byte, _uniqueId []byte, _amt *big.Int, _swapFeeAmt *big.Int) (*types.Transaction, error) {
	return _Op.Contract.FromUCross(&_Op.TransactOpts, _id, _uniqueId, _amt, _swapFeeAmt)
}

// FromUCross is a paid mutator transaction binding the contract method 0xfa37c031.
//
// Solidity: function fromUCross(bytes _id, bytes _uniqueId, uint256 _amt, uint256 _swapFeeAmt) returns()
func (_Op *OpTransactorSession) FromUCross(_id []byte, _uniqueId []byte, _amt *big.Int, _swapFeeAmt *big.Int) (*types.Transaction, error) {
	return _Op.Contract.FromUCross(&_Op.TransactOpts, _id, _uniqueId, _amt, _swapFeeAmt)
}

// Pull is a paid mutator transaction binding the contract method 0x07251d88.
//
// Solidity: function pull(address _token, uint256 _amt, address _to) returns(uint256 amt)
func (_Op *OpTransactor) Pull(opts *bind.TransactOpts, _token common.Address, _amt *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "pull", _token, _amt, _to)
}

// Pull is a paid mutator transaction binding the contract method 0x07251d88.
//
// Solidity: function pull(address _token, uint256 _amt, address _to) returns(uint256 amt)
func (_Op *OpSession) Pull(_token common.Address, _amt *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Op.Contract.Pull(&_Op.TransactOpts, _token, _amt, _to)
}

// Pull is a paid mutator transaction binding the contract method 0x07251d88.
//
// Solidity: function pull(address _token, uint256 _amt, address _to) returns(uint256 amt)
func (_Op *OpTransactorSession) Pull(_token common.Address, _amt *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Op.Contract.Pull(&_Op.TransactOpts, _token, _amt, _to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Op *OpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Op *OpSession) RenounceOwnership() (*types.Transaction, error) {
	return _Op.Contract.RenounceOwnership(&_Op.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Op *OpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Op.Contract.RenounceOwnership(&_Op.TransactOpts)
}

// SetFundsProvider is a paid mutator transaction binding the contract method 0xa7ee9c8a.
//
// Solidity: function setFundsProvider(address _newFundsProvider) returns()
func (_Op *OpTransactor) SetFundsProvider(opts *bind.TransactOpts, _newFundsProvider common.Address) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "setFundsProvider", _newFundsProvider)
}

// SetFundsProvider is a paid mutator transaction binding the contract method 0xa7ee9c8a.
//
// Solidity: function setFundsProvider(address _newFundsProvider) returns()
func (_Op *OpSession) SetFundsProvider(_newFundsProvider common.Address) (*types.Transaction, error) {
	return _Op.Contract.SetFundsProvider(&_Op.TransactOpts, _newFundsProvider)
}

// SetFundsProvider is a paid mutator transaction binding the contract method 0xa7ee9c8a.
//
// Solidity: function setFundsProvider(address _newFundsProvider) returns()
func (_Op *OpTransactorSession) SetFundsProvider(_newFundsProvider common.Address) (*types.Transaction, error) {
	return _Op.Contract.SetFundsProvider(&_Op.TransactOpts, _newFundsProvider)
}

// SetGasFeeTo is a paid mutator transaction binding the contract method 0x78e02af1.
//
// Solidity: function setGasFeeTo(address _newGasFeeTo) returns()
func (_Op *OpTransactor) SetGasFeeTo(opts *bind.TransactOpts, _newGasFeeTo common.Address) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "setGasFeeTo", _newGasFeeTo)
}

// SetGasFeeTo is a paid mutator transaction binding the contract method 0x78e02af1.
//
// Solidity: function setGasFeeTo(address _newGasFeeTo) returns()
func (_Op *OpSession) SetGasFeeTo(_newGasFeeTo common.Address) (*types.Transaction, error) {
	return _Op.Contract.SetGasFeeTo(&_Op.TransactOpts, _newGasFeeTo)
}

// SetGasFeeTo is a paid mutator transaction binding the contract method 0x78e02af1.
//
// Solidity: function setGasFeeTo(address _newGasFeeTo) returns()
func (_Op *OpTransactorSession) SetGasFeeTo(_newGasFeeTo common.Address) (*types.Transaction, error) {
	return _Op.Contract.SetGasFeeTo(&_Op.TransactOpts, _newGasFeeTo)
}

// SetSwapFeeTo is a paid mutator transaction binding the contract method 0x656b6b93.
//
// Solidity: function setSwapFeeTo(address _newSwapFeeTo) returns()
func (_Op *OpTransactor) SetSwapFeeTo(opts *bind.TransactOpts, _newSwapFeeTo common.Address) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "setSwapFeeTo", _newSwapFeeTo)
}

// SetSwapFeeTo is a paid mutator transaction binding the contract method 0x656b6b93.
//
// Solidity: function setSwapFeeTo(address _newSwapFeeTo) returns()
func (_Op *OpSession) SetSwapFeeTo(_newSwapFeeTo common.Address) (*types.Transaction, error) {
	return _Op.Contract.SetSwapFeeTo(&_Op.TransactOpts, _newSwapFeeTo)
}

// SetSwapFeeTo is a paid mutator transaction binding the contract method 0x656b6b93.
//
// Solidity: function setSwapFeeTo(address _newSwapFeeTo) returns()
func (_Op *OpTransactorSession) SetSwapFeeTo(_newSwapFeeTo common.Address) (*types.Transaction, error) {
	return _Op.Contract.SetSwapFeeTo(&_Op.TransactOpts, _newSwapFeeTo)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x3b99adf7.
//
// Solidity: function setWhitelist(address[] _addrArr, bool[] _flags) returns()
func (_Op *OpTransactor) SetWhitelist(opts *bind.TransactOpts, _addrArr []common.Address, _flags []bool) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "setWhitelist", _addrArr, _flags)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x3b99adf7.
//
// Solidity: function setWhitelist(address[] _addrArr, bool[] _flags) returns()
func (_Op *OpSession) SetWhitelist(_addrArr []common.Address, _flags []bool) (*types.Transaction, error) {
	return _Op.Contract.SetWhitelist(&_Op.TransactOpts, _addrArr, _flags)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x3b99adf7.
//
// Solidity: function setWhitelist(address[] _addrArr, bool[] _flags) returns()
func (_Op *OpTransactorSession) SetWhitelist(_addrArr []common.Address, _flags []bool) (*types.Transaction, error) {
	return _Op.Contract.SetWhitelist(&_Op.TransactOpts, _addrArr, _flags)
}

// ToUCross is a paid mutator transaction binding the contract method 0xccd2f919.
//
// Solidity: function toUCross(bytes _id, bytes _uniqueId, uint256 _amt, uint256 _gasFeeAmt, address _to) returns()
func (_Op *OpTransactor) ToUCross(opts *bind.TransactOpts, _id []byte, _uniqueId []byte, _amt *big.Int, _gasFeeAmt *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "toUCross", _id, _uniqueId, _amt, _gasFeeAmt, _to)
}

// ToUCross is a paid mutator transaction binding the contract method 0xccd2f919.
//
// Solidity: function toUCross(bytes _id, bytes _uniqueId, uint256 _amt, uint256 _gasFeeAmt, address _to) returns()
func (_Op *OpSession) ToUCross(_id []byte, _uniqueId []byte, _amt *big.Int, _gasFeeAmt *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Op.Contract.ToUCross(&_Op.TransactOpts, _id, _uniqueId, _amt, _gasFeeAmt, _to)
}

// ToUCross is a paid mutator transaction binding the contract method 0xccd2f919.
//
// Solidity: function toUCross(bytes _id, bytes _uniqueId, uint256 _amt, uint256 _gasFeeAmt, address _to) returns()
func (_Op *OpTransactorSession) ToUCross(_id []byte, _uniqueId []byte, _amt *big.Int, _gasFeeAmt *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Op.Contract.ToUCross(&_Op.TransactOpts, _id, _uniqueId, _amt, _gasFeeAmt, _to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Op *OpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Op.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Op *OpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Op.Contract.TransferOwnership(&_Op.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Op *OpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Op.Contract.TransferOwnership(&_Op.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Op *OpTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Op.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Op *OpSession) Receive() (*types.Transaction, error) {
	return _Op.Contract.Receive(&_Op.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Op *OpTransactorSession) Receive() (*types.Transaction, error) {
	return _Op.Contract.Receive(&_Op.TransactOpts)
}

// OpFlipRunningIterator is returned from FilterFlipRunning and is used to iterate over the raw logs and unpacked data for FlipRunning events raised by the Op contract.
type OpFlipRunningIterator struct {
	Event *OpFlipRunning // Event containing the contract specifics and raw log

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
func (it *OpFlipRunningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpFlipRunning)
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
		it.Event = new(OpFlipRunning)
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
func (it *OpFlipRunningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpFlipRunningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpFlipRunning represents a FlipRunning event raised by the Op contract.
type OpFlipRunning struct {
	Prev bool
	Curr bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFlipRunning is a free log retrieval operation binding the contract event 0xa3007e027479d23513afe9254fb76b0cae858019088b1ea41dbc28a9be2e6462.
//
// Solidity: event FlipRunning(bool _prev, bool _curr)
func (_Op *OpFilterer) FilterFlipRunning(opts *bind.FilterOpts) (*OpFlipRunningIterator, error) {

	logs, sub, err := _Op.contract.FilterLogs(opts, "FlipRunning")
	if err != nil {
		return nil, err
	}
	return &OpFlipRunningIterator{contract: _Op.contract, event: "FlipRunning", logs: logs, sub: sub}, nil
}

// WatchFlipRunning is a free log subscription operation binding the contract event 0xa3007e027479d23513afe9254fb76b0cae858019088b1ea41dbc28a9be2e6462.
//
// Solidity: event FlipRunning(bool _prev, bool _curr)
func (_Op *OpFilterer) WatchFlipRunning(opts *bind.WatchOpts, sink chan<- *OpFlipRunning) (event.Subscription, error) {

	logs, sub, err := _Op.contract.WatchLogs(opts, "FlipRunning")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpFlipRunning)
				if err := _Op.contract.UnpackLog(event, "FlipRunning", log); err != nil {
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

// ParseFlipRunning is a log parse operation binding the contract event 0xa3007e027479d23513afe9254fb76b0cae858019088b1ea41dbc28a9be2e6462.
//
// Solidity: event FlipRunning(bool _prev, bool _curr)
func (_Op *OpFilterer) ParseFlipRunning(log types.Log) (*OpFlipRunning, error) {
	event := new(OpFlipRunning)
	if err := _Op.contract.UnpackLog(event, "FlipRunning", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpFundsProviderIterator is returned from FilterFundsProvider and is used to iterate over the raw logs and unpacked data for FundsProvider events raised by the Op contract.
type OpFundsProviderIterator struct {
	Event *OpFundsProvider // Event containing the contract specifics and raw log

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
func (it *OpFundsProviderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpFundsProvider)
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
		it.Event = new(OpFundsProvider)
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
func (it *OpFundsProviderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpFundsProviderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpFundsProvider represents a FundsProvider event raised by the Op contract.
type OpFundsProvider struct {
	Prev common.Address
	Curr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFundsProvider is a free log retrieval operation binding the contract event 0xb122585fc594e9a2e9bff790bf833cde0d387790338016f133a50c4e8eba17b5.
//
// Solidity: event FundsProvider(address _prev, address _curr)
func (_Op *OpFilterer) FilterFundsProvider(opts *bind.FilterOpts) (*OpFundsProviderIterator, error) {

	logs, sub, err := _Op.contract.FilterLogs(opts, "FundsProvider")
	if err != nil {
		return nil, err
	}
	return &OpFundsProviderIterator{contract: _Op.contract, event: "FundsProvider", logs: logs, sub: sub}, nil
}

// WatchFundsProvider is a free log subscription operation binding the contract event 0xb122585fc594e9a2e9bff790bf833cde0d387790338016f133a50c4e8eba17b5.
//
// Solidity: event FundsProvider(address _prev, address _curr)
func (_Op *OpFilterer) WatchFundsProvider(opts *bind.WatchOpts, sink chan<- *OpFundsProvider) (event.Subscription, error) {

	logs, sub, err := _Op.contract.WatchLogs(opts, "FundsProvider")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpFundsProvider)
				if err := _Op.contract.UnpackLog(event, "FundsProvider", log); err != nil {
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

// ParseFundsProvider is a log parse operation binding the contract event 0xb122585fc594e9a2e9bff790bf833cde0d387790338016f133a50c4e8eba17b5.
//
// Solidity: event FundsProvider(address _prev, address _curr)
func (_Op *OpFilterer) ParseFundsProvider(log types.Log) (*OpFundsProvider, error) {
	event := new(OpFundsProvider)
	if err := _Op.contract.UnpackLog(event, "FundsProvider", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpGasFeeToIterator is returned from FilterGasFeeTo and is used to iterate over the raw logs and unpacked data for GasFeeTo events raised by the Op contract.
type OpGasFeeToIterator struct {
	Event *OpGasFeeTo // Event containing the contract specifics and raw log

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
func (it *OpGasFeeToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpGasFeeTo)
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
		it.Event = new(OpGasFeeTo)
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
func (it *OpGasFeeToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpGasFeeToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpGasFeeTo represents a GasFeeTo event raised by the Op contract.
type OpGasFeeTo struct {
	Prev common.Address
	Curr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGasFeeTo is a free log retrieval operation binding the contract event 0x71ee15585b7aa7dffccf24f7114c37e44ee768f843c7b748106a84abe48ecf37.
//
// Solidity: event GasFeeTo(address _prev, address _curr)
func (_Op *OpFilterer) FilterGasFeeTo(opts *bind.FilterOpts) (*OpGasFeeToIterator, error) {

	logs, sub, err := _Op.contract.FilterLogs(opts, "GasFeeTo")
	if err != nil {
		return nil, err
	}
	return &OpGasFeeToIterator{contract: _Op.contract, event: "GasFeeTo", logs: logs, sub: sub}, nil
}

// WatchGasFeeTo is a free log subscription operation binding the contract event 0x71ee15585b7aa7dffccf24f7114c37e44ee768f843c7b748106a84abe48ecf37.
//
// Solidity: event GasFeeTo(address _prev, address _curr)
func (_Op *OpFilterer) WatchGasFeeTo(opts *bind.WatchOpts, sink chan<- *OpGasFeeTo) (event.Subscription, error) {

	logs, sub, err := _Op.contract.WatchLogs(opts, "GasFeeTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpGasFeeTo)
				if err := _Op.contract.UnpackLog(event, "GasFeeTo", log); err != nil {
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

// ParseGasFeeTo is a log parse operation binding the contract event 0x71ee15585b7aa7dffccf24f7114c37e44ee768f843c7b748106a84abe48ecf37.
//
// Solidity: event GasFeeTo(address _prev, address _curr)
func (_Op *OpFilterer) ParseGasFeeTo(log types.Log) (*OpGasFeeTo, error) {
	event := new(OpGasFeeTo)
	if err := _Op.contract.UnpackLog(event, "GasFeeTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Op contract.
type OpOwnershipTransferredIterator struct {
	Event *OpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpOwnershipTransferred)
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
		it.Event = new(OpOwnershipTransferred)
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
func (it *OpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpOwnershipTransferred represents a OwnershipTransferred event raised by the Op contract.
type OpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Op *OpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Op.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OpOwnershipTransferredIterator{contract: _Op.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Op *OpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Op.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpOwnershipTransferred)
				if err := _Op.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Op *OpFilterer) ParseOwnershipTransferred(log types.Log) (*OpOwnershipTransferred, error) {
	event := new(OpOwnershipTransferred)
	if err := _Op.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpPullIterator is returned from FilterPull and is used to iterate over the raw logs and unpacked data for Pull events raised by the Op contract.
type OpPullIterator struct {
	Event *OpPull // Event containing the contract specifics and raw log

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
func (it *OpPullIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpPull)
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
		it.Event = new(OpPull)
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
func (it *OpPullIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpPullIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpPull represents a Pull event raised by the Op contract.
type OpPull struct {
	Token common.Address
	Amt   *big.Int
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPull is a free log retrieval operation binding the contract event 0xe9e024c930150844ce9b8d61c6befed01393ebe790d7eea75fb4070b77aef2df.
//
// Solidity: event Pull(address token, uint256 amt, address to)
func (_Op *OpFilterer) FilterPull(opts *bind.FilterOpts) (*OpPullIterator, error) {

	logs, sub, err := _Op.contract.FilterLogs(opts, "Pull")
	if err != nil {
		return nil, err
	}
	return &OpPullIterator{contract: _Op.contract, event: "Pull", logs: logs, sub: sub}, nil
}

// WatchPull is a free log subscription operation binding the contract event 0xe9e024c930150844ce9b8d61c6befed01393ebe790d7eea75fb4070b77aef2df.
//
// Solidity: event Pull(address token, uint256 amt, address to)
func (_Op *OpFilterer) WatchPull(opts *bind.WatchOpts, sink chan<- *OpPull) (event.Subscription, error) {

	logs, sub, err := _Op.contract.WatchLogs(opts, "Pull")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpPull)
				if err := _Op.contract.UnpackLog(event, "Pull", log); err != nil {
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

// ParsePull is a log parse operation binding the contract event 0xe9e024c930150844ce9b8d61c6befed01393ebe790d7eea75fb4070b77aef2df.
//
// Solidity: event Pull(address token, uint256 amt, address to)
func (_Op *OpFilterer) ParsePull(log types.Log) (*OpPull, error) {
	event := new(OpPull)
	if err := _Op.contract.UnpackLog(event, "Pull", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpSetWhitelistIterator is returned from FilterSetWhitelist and is used to iterate over the raw logs and unpacked data for SetWhitelist events raised by the Op contract.
type OpSetWhitelistIterator struct {
	Event *OpSetWhitelist // Event containing the contract specifics and raw log

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
func (it *OpSetWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpSetWhitelist)
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
		it.Event = new(OpSetWhitelist)
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
func (it *OpSetWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpSetWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpSetWhitelist represents a SetWhitelist event raised by the Op contract.
type OpSetWhitelist struct {
	Addr        common.Address
	IsWhitelist bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetWhitelist is a free log retrieval operation binding the contract event 0xf6019ec0a78d156d249a1ec7579e2321f6ac7521d6e1d2eacf90ba4a184dcceb.
//
// Solidity: event SetWhitelist(address _addr, bool _isWhitelist)
func (_Op *OpFilterer) FilterSetWhitelist(opts *bind.FilterOpts) (*OpSetWhitelistIterator, error) {

	logs, sub, err := _Op.contract.FilterLogs(opts, "SetWhitelist")
	if err != nil {
		return nil, err
	}
	return &OpSetWhitelistIterator{contract: _Op.contract, event: "SetWhitelist", logs: logs, sub: sub}, nil
}

// WatchSetWhitelist is a free log subscription operation binding the contract event 0xf6019ec0a78d156d249a1ec7579e2321f6ac7521d6e1d2eacf90ba4a184dcceb.
//
// Solidity: event SetWhitelist(address _addr, bool _isWhitelist)
func (_Op *OpFilterer) WatchSetWhitelist(opts *bind.WatchOpts, sink chan<- *OpSetWhitelist) (event.Subscription, error) {

	logs, sub, err := _Op.contract.WatchLogs(opts, "SetWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpSetWhitelist)
				if err := _Op.contract.UnpackLog(event, "SetWhitelist", log); err != nil {
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

// ParseSetWhitelist is a log parse operation binding the contract event 0xf6019ec0a78d156d249a1ec7579e2321f6ac7521d6e1d2eacf90ba4a184dcceb.
//
// Solidity: event SetWhitelist(address _addr, bool _isWhitelist)
func (_Op *OpFilterer) ParseSetWhitelist(log types.Log) (*OpSetWhitelist, error) {
	event := new(OpSetWhitelist)
	if err := _Op.contract.UnpackLog(event, "SetWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Op contract.
type OpSwapIterator struct {
	Event *OpSwap // Event containing the contract specifics and raw log

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
func (it *OpSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpSwap)
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
		it.Event = new(OpSwap)
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
func (it *OpSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpSwap represents a Swap event raised by the Op contract.
type OpSwap struct {
	Id        []byte
	UniqueId  []byte
	Action    uint8
	TokenFrom common.Address
	TokenTo   common.Address
	RetAmt    *big.Int
	SrcAmt    *big.Int
	FeeAmt    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xad1cf0cdde6719cd3b92ce2f5107e0e370d036bc43fae4a5c328113ea0477b0e.
//
// Solidity: event Swap(bytes id, bytes uniqueId, uint8 action, address tokenFrom, address tokenTo, uint256 retAmt, uint256 srcAmt, uint256 feeAmt)
func (_Op *OpFilterer) FilterSwap(opts *bind.FilterOpts) (*OpSwapIterator, error) {

	logs, sub, err := _Op.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &OpSwapIterator{contract: _Op.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xad1cf0cdde6719cd3b92ce2f5107e0e370d036bc43fae4a5c328113ea0477b0e.
//
// Solidity: event Swap(bytes id, bytes uniqueId, uint8 action, address tokenFrom, address tokenTo, uint256 retAmt, uint256 srcAmt, uint256 feeAmt)
func (_Op *OpFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *OpSwap) (event.Subscription, error) {

	logs, sub, err := _Op.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpSwap)
				if err := _Op.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xad1cf0cdde6719cd3b92ce2f5107e0e370d036bc43fae4a5c328113ea0477b0e.
//
// Solidity: event Swap(bytes id, bytes uniqueId, uint8 action, address tokenFrom, address tokenTo, uint256 retAmt, uint256 srcAmt, uint256 feeAmt)
func (_Op *OpFilterer) ParseSwap(log types.Log) (*OpSwap, error) {
	event := new(OpSwap)
	if err := _Op.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OpSwapFeeToIterator is returned from FilterSwapFeeTo and is used to iterate over the raw logs and unpacked data for SwapFeeTo events raised by the Op contract.
type OpSwapFeeToIterator struct {
	Event *OpSwapFeeTo // Event containing the contract specifics and raw log

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
func (it *OpSwapFeeToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OpSwapFeeTo)
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
		it.Event = new(OpSwapFeeTo)
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
func (it *OpSwapFeeToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OpSwapFeeToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OpSwapFeeTo represents a SwapFeeTo event raised by the Op contract.
type OpSwapFeeTo struct {
	Prev common.Address
	Curr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSwapFeeTo is a free log retrieval operation binding the contract event 0xb8a7442c5c359723328864f811dcfe0695cac04b76b119c17fea153c4f25166f.
//
// Solidity: event SwapFeeTo(address _prev, address _curr)
func (_Op *OpFilterer) FilterSwapFeeTo(opts *bind.FilterOpts) (*OpSwapFeeToIterator, error) {

	logs, sub, err := _Op.contract.FilterLogs(opts, "SwapFeeTo")
	if err != nil {
		return nil, err
	}
	return &OpSwapFeeToIterator{contract: _Op.contract, event: "SwapFeeTo", logs: logs, sub: sub}, nil
}

// WatchSwapFeeTo is a free log subscription operation binding the contract event 0xb8a7442c5c359723328864f811dcfe0695cac04b76b119c17fea153c4f25166f.
//
// Solidity: event SwapFeeTo(address _prev, address _curr)
func (_Op *OpFilterer) WatchSwapFeeTo(opts *bind.WatchOpts, sink chan<- *OpSwapFeeTo) (event.Subscription, error) {

	logs, sub, err := _Op.contract.WatchLogs(opts, "SwapFeeTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OpSwapFeeTo)
				if err := _Op.contract.UnpackLog(event, "SwapFeeTo", log); err != nil {
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

// ParseSwapFeeTo is a log parse operation binding the contract event 0xb8a7442c5c359723328864f811dcfe0695cac04b76b119c17fea153c4f25166f.
//
// Solidity: event SwapFeeTo(address _prev, address _curr)
func (_Op *OpFilterer) ParseSwapFeeTo(log types.Log) (*OpSwapFeeTo, error) {
	event := new(OpSwapFeeTo)
	if err := _Op.contract.UnpackLog(event, "SwapFeeTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
