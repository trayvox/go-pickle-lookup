// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PickleJar

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PickleJarABI is the input ABI used to generate the binding from.
const PickleJarABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"harvest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"}],\"name\":\"setMin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_timelock\",\"type\":\"address\"}],\"name\":\"setTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PickleJar is an auto generated Go binding around an Ethereum contract.
type PickleJar struct {
	PickleJarCaller     // Read-only binding to the contract
	PickleJarTransactor // Write-only binding to the contract
	PickleJarFilterer   // Log filterer for contract events
}

// PickleJarCaller is an auto generated read-only Go binding around an Ethereum contract.
type PickleJarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PickleJarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PickleJarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PickleJarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PickleJarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PickleJarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PickleJarSession struct {
	Contract     *PickleJar        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PickleJarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PickleJarCallerSession struct {
	Contract *PickleJarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PickleJarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PickleJarTransactorSession struct {
	Contract     *PickleJarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PickleJarRaw is an auto generated low-level Go binding around an Ethereum contract.
type PickleJarRaw struct {
	Contract *PickleJar // Generic contract binding to access the raw methods on
}

// PickleJarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PickleJarCallerRaw struct {
	Contract *PickleJarCaller // Generic read-only contract binding to access the raw methods on
}

// PickleJarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PickleJarTransactorRaw struct {
	Contract *PickleJarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPickleJar creates a new instance of PickleJar, bound to a specific deployed contract.
func NewPickleJar(address common.Address, backend bind.ContractBackend) (*PickleJar, error) {
	contract, err := bindPickleJar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PickleJar{PickleJarCaller: PickleJarCaller{contract: contract}, PickleJarTransactor: PickleJarTransactor{contract: contract}, PickleJarFilterer: PickleJarFilterer{contract: contract}}, nil
}

// NewPickleJarCaller creates a new read-only instance of PickleJar, bound to a specific deployed contract.
func NewPickleJarCaller(address common.Address, caller bind.ContractCaller) (*PickleJarCaller, error) {
	contract, err := bindPickleJar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PickleJarCaller{contract: contract}, nil
}

// NewPickleJarTransactor creates a new write-only instance of PickleJar, bound to a specific deployed contract.
func NewPickleJarTransactor(address common.Address, transactor bind.ContractTransactor) (*PickleJarTransactor, error) {
	contract, err := bindPickleJar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PickleJarTransactor{contract: contract}, nil
}

// NewPickleJarFilterer creates a new log filterer instance of PickleJar, bound to a specific deployed contract.
func NewPickleJarFilterer(address common.Address, filterer bind.ContractFilterer) (*PickleJarFilterer, error) {
	contract, err := bindPickleJar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PickleJarFilterer{contract: contract}, nil
}

// bindPickleJar binds a generic wrapper to an already deployed contract.
func bindPickleJar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PickleJarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PickleJar *PickleJarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PickleJar.Contract.PickleJarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PickleJar *PickleJarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PickleJar.Contract.PickleJarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PickleJar *PickleJarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PickleJar.Contract.PickleJarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PickleJar *PickleJarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PickleJar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PickleJar *PickleJarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PickleJar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PickleJar *PickleJarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PickleJar.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PickleJar *PickleJarCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PickleJar *PickleJarSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PickleJar.Contract.Allowance(&_PickleJar.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PickleJar *PickleJarCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PickleJar.Contract.Allowance(&_PickleJar.CallOpts, owner, spender)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_PickleJar *PickleJarCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_PickleJar *PickleJarSession) Available() (*big.Int, error) {
	return _PickleJar.Contract.Available(&_PickleJar.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_PickleJar *PickleJarCallerSession) Available() (*big.Int, error) {
	return _PickleJar.Contract.Available(&_PickleJar.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_PickleJar *PickleJarCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_PickleJar *PickleJarSession) Balance() (*big.Int, error) {
	return _PickleJar.Contract.Balance(&_PickleJar.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_PickleJar *PickleJarCallerSession) Balance() (*big.Int, error) {
	return _PickleJar.Contract.Balance(&_PickleJar.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PickleJar *PickleJarCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PickleJar *PickleJarSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PickleJar.Contract.BalanceOf(&_PickleJar.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PickleJar *PickleJarCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PickleJar.Contract.BalanceOf(&_PickleJar.CallOpts, account)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PickleJar *PickleJarCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PickleJar *PickleJarSession) Controller() (common.Address, error) {
	return _PickleJar.Contract.Controller(&_PickleJar.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_PickleJar *PickleJarCallerSession) Controller() (common.Address, error) {
	return _PickleJar.Contract.Controller(&_PickleJar.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PickleJar *PickleJarCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PickleJar *PickleJarSession) Decimals() (uint8, error) {
	return _PickleJar.Contract.Decimals(&_PickleJar.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PickleJar *PickleJarCallerSession) Decimals() (uint8, error) {
	return _PickleJar.Contract.Decimals(&_PickleJar.CallOpts)
}

// GetRatio is a free data retrieval call binding the contract method 0xec1ebd7a.
//
// Solidity: function getRatio() view returns(uint256)
func (_PickleJar *PickleJarCaller) GetRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "getRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRatio is a free data retrieval call binding the contract method 0xec1ebd7a.
//
// Solidity: function getRatio() view returns(uint256)
func (_PickleJar *PickleJarSession) GetRatio() (*big.Int, error) {
	return _PickleJar.Contract.GetRatio(&_PickleJar.CallOpts)
}

// GetRatio is a free data retrieval call binding the contract method 0xec1ebd7a.
//
// Solidity: function getRatio() view returns(uint256)
func (_PickleJar *PickleJarCallerSession) GetRatio() (*big.Int, error) {
	return _PickleJar.Contract.GetRatio(&_PickleJar.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PickleJar *PickleJarCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PickleJar *PickleJarSession) Governance() (common.Address, error) {
	return _PickleJar.Contract.Governance(&_PickleJar.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_PickleJar *PickleJarCallerSession) Governance() (common.Address, error) {
	return _PickleJar.Contract.Governance(&_PickleJar.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_PickleJar *PickleJarCaller) Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_PickleJar *PickleJarSession) Max() (*big.Int, error) {
	return _PickleJar.Contract.Max(&_PickleJar.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_PickleJar *PickleJarCallerSession) Max() (*big.Int, error) {
	return _PickleJar.Contract.Max(&_PickleJar.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_PickleJar *PickleJarCaller) Min(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "min")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_PickleJar *PickleJarSession) Min() (*big.Int, error) {
	return _PickleJar.Contract.Min(&_PickleJar.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_PickleJar *PickleJarCallerSession) Min() (*big.Int, error) {
	return _PickleJar.Contract.Min(&_PickleJar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PickleJar *PickleJarCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PickleJar *PickleJarSession) Name() (string, error) {
	return _PickleJar.Contract.Name(&_PickleJar.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PickleJar *PickleJarCallerSession) Name() (string, error) {
	return _PickleJar.Contract.Name(&_PickleJar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PickleJar *PickleJarCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PickleJar *PickleJarSession) Symbol() (string, error) {
	return _PickleJar.Contract.Symbol(&_PickleJar.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PickleJar *PickleJarCallerSession) Symbol() (string, error) {
	return _PickleJar.Contract.Symbol(&_PickleJar.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_PickleJar *PickleJarCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_PickleJar *PickleJarSession) Timelock() (common.Address, error) {
	return _PickleJar.Contract.Timelock(&_PickleJar.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_PickleJar *PickleJarCallerSession) Timelock() (common.Address, error) {
	return _PickleJar.Contract.Timelock(&_PickleJar.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PickleJar *PickleJarCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PickleJar *PickleJarSession) Token() (common.Address, error) {
	return _PickleJar.Contract.Token(&_PickleJar.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_PickleJar *PickleJarCallerSession) Token() (common.Address, error) {
	return _PickleJar.Contract.Token(&_PickleJar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PickleJar *PickleJarCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PickleJar.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PickleJar *PickleJarSession) TotalSupply() (*big.Int, error) {
	return _PickleJar.Contract.TotalSupply(&_PickleJar.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PickleJar *PickleJarCallerSession) TotalSupply() (*big.Int, error) {
	return _PickleJar.Contract.TotalSupply(&_PickleJar.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PickleJar *PickleJarTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PickleJar *PickleJarSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Approve(&_PickleJar.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_PickleJar *PickleJarTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Approve(&_PickleJar.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PickleJar *PickleJarTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PickleJar *PickleJarSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.DecreaseAllowance(&_PickleJar.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_PickleJar *PickleJarTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.DecreaseAllowance(&_PickleJar.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_PickleJar *PickleJarTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_PickleJar *PickleJarSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Deposit(&_PickleJar.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_PickleJar *PickleJarTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Deposit(&_PickleJar.TransactOpts, _amount)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_PickleJar *PickleJarTransactor) DepositAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "depositAll")
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_PickleJar *PickleJarSession) DepositAll() (*types.Transaction, error) {
	return _PickleJar.Contract.DepositAll(&_PickleJar.TransactOpts)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_PickleJar *PickleJarTransactorSession) DepositAll() (*types.Transaction, error) {
	return _PickleJar.Contract.DepositAll(&_PickleJar.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_PickleJar *PickleJarTransactor) Earn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "earn")
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_PickleJar *PickleJarSession) Earn() (*types.Transaction, error) {
	return _PickleJar.Contract.Earn(&_PickleJar.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_PickleJar *PickleJarTransactorSession) Earn() (*types.Transaction, error) {
	return _PickleJar.Contract.Earn(&_PickleJar.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_PickleJar *PickleJarTransactor) Harvest(opts *bind.TransactOpts, reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "harvest", reserve, amount)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_PickleJar *PickleJarSession) Harvest(reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Harvest(&_PickleJar.TransactOpts, reserve, amount)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_PickleJar *PickleJarTransactorSession) Harvest(reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Harvest(&_PickleJar.TransactOpts, reserve, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PickleJar *PickleJarTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PickleJar *PickleJarSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.IncreaseAllowance(&_PickleJar.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_PickleJar *PickleJarTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.IncreaseAllowance(&_PickleJar.TransactOpts, spender, addedValue)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_PickleJar *PickleJarTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_PickleJar *PickleJarSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _PickleJar.Contract.SetController(&_PickleJar.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_PickleJar *PickleJarTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _PickleJar.Contract.SetController(&_PickleJar.TransactOpts, _controller)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PickleJar *PickleJarTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PickleJar *PickleJarSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PickleJar.Contract.SetGovernance(&_PickleJar.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_PickleJar *PickleJarTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _PickleJar.Contract.SetGovernance(&_PickleJar.TransactOpts, _governance)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_PickleJar *PickleJarTransactor) SetMin(opts *bind.TransactOpts, _min *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "setMin", _min)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_PickleJar *PickleJarSession) SetMin(_min *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.SetMin(&_PickleJar.TransactOpts, _min)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_PickleJar *PickleJarTransactorSession) SetMin(_min *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.SetMin(&_PickleJar.TransactOpts, _min)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _timelock) returns()
func (_PickleJar *PickleJarTransactor) SetTimelock(opts *bind.TransactOpts, _timelock common.Address) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "setTimelock", _timelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _timelock) returns()
func (_PickleJar *PickleJarSession) SetTimelock(_timelock common.Address) (*types.Transaction, error) {
	return _PickleJar.Contract.SetTimelock(&_PickleJar.TransactOpts, _timelock)
}

// SetTimelock is a paid mutator transaction binding the contract method 0xbdacb303.
//
// Solidity: function setTimelock(address _timelock) returns()
func (_PickleJar *PickleJarTransactorSession) SetTimelock(_timelock common.Address) (*types.Transaction, error) {
	return _PickleJar.Contract.SetTimelock(&_PickleJar.TransactOpts, _timelock)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PickleJar *PickleJarTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PickleJar *PickleJarSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Transfer(&_PickleJar.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_PickleJar *PickleJarTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Transfer(&_PickleJar.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PickleJar *PickleJarTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PickleJar *PickleJarSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.TransferFrom(&_PickleJar.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_PickleJar *PickleJarTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.TransferFrom(&_PickleJar.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_PickleJar *PickleJarTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_PickleJar *PickleJarSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Withdraw(&_PickleJar.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_PickleJar *PickleJarTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _PickleJar.Contract.Withdraw(&_PickleJar.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_PickleJar *PickleJarTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PickleJar.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_PickleJar *PickleJarSession) WithdrawAll() (*types.Transaction, error) {
	return _PickleJar.Contract.WithdrawAll(&_PickleJar.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_PickleJar *PickleJarTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _PickleJar.Contract.WithdrawAll(&_PickleJar.TransactOpts)
}

// PickleJarApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PickleJar contract.
type PickleJarApprovalIterator struct {
	Event *PickleJarApproval // Event containing the contract specifics and raw log

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
func (it *PickleJarApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PickleJarApproval)
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
		it.Event = new(PickleJarApproval)
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
func (it *PickleJarApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PickleJarApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PickleJarApproval represents a Approval event raised by the PickleJar contract.
type PickleJarApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PickleJar *PickleJarFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PickleJarApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PickleJar.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PickleJarApprovalIterator{contract: _PickleJar.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PickleJar *PickleJarFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PickleJarApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PickleJar.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PickleJarApproval)
				if err := _PickleJar.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_PickleJar *PickleJarFilterer) ParseApproval(log types.Log) (*PickleJarApproval, error) {
	event := new(PickleJarApproval)
	if err := _PickleJar.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PickleJarTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PickleJar contract.
type PickleJarTransferIterator struct {
	Event *PickleJarTransfer // Event containing the contract specifics and raw log

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
func (it *PickleJarTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PickleJarTransfer)
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
		it.Event = new(PickleJarTransfer)
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
func (it *PickleJarTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PickleJarTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PickleJarTransfer represents a Transfer event raised by the PickleJar contract.
type PickleJarTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PickleJar *PickleJarFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PickleJarTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PickleJar.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PickleJarTransferIterator{contract: _PickleJar.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PickleJar *PickleJarFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PickleJarTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PickleJar.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PickleJarTransfer)
				if err := _PickleJar.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_PickleJar *PickleJarFilterer) ParseTransfer(log types.Log) (*PickleJarTransfer, error) {
	event := new(PickleJarTransfer)
	if err := _PickleJar.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
