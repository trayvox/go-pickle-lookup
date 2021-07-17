// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Gauge

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

// GaugeABI is the input ABI used to generate the binding from.
const GaugeABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DILL\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISTRIBUTION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PICKLE\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREASURY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"derivedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"derivedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"derivedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Gauge is an auto generated Go binding around an Ethereum contract.
type Gauge struct {
	GaugeCaller     // Read-only binding to the contract
	GaugeTransactor // Write-only binding to the contract
	GaugeFilterer   // Log filterer for contract events
}

// GaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type GaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GaugeSession struct {
	Contract     *Gauge            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GaugeCallerSession struct {
	Contract *GaugeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GaugeTransactorSession struct {
	Contract     *GaugeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type GaugeRaw struct {
	Contract *Gauge // Generic contract binding to access the raw methods on
}

// GaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GaugeCallerRaw struct {
	Contract *GaugeCaller // Generic read-only contract binding to access the raw methods on
}

// GaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GaugeTransactorRaw struct {
	Contract *GaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGauge creates a new instance of Gauge, bound to a specific deployed contract.
func NewGauge(address common.Address, backend bind.ContractBackend) (*Gauge, error) {
	contract, err := bindGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gauge{GaugeCaller: GaugeCaller{contract: contract}, GaugeTransactor: GaugeTransactor{contract: contract}, GaugeFilterer: GaugeFilterer{contract: contract}}, nil
}

// NewGaugeCaller creates a new read-only instance of Gauge, bound to a specific deployed contract.
func NewGaugeCaller(address common.Address, caller bind.ContractCaller) (*GaugeCaller, error) {
	contract, err := bindGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeCaller{contract: contract}, nil
}

// NewGaugeTransactor creates a new write-only instance of Gauge, bound to a specific deployed contract.
func NewGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*GaugeTransactor, error) {
	contract, err := bindGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GaugeTransactor{contract: contract}, nil
}

// NewGaugeFilterer creates a new log filterer instance of Gauge, bound to a specific deployed contract.
func NewGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*GaugeFilterer, error) {
	contract, err := bindGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GaugeFilterer{contract: contract}, nil
}

// bindGauge binds a generic wrapper to an already deployed contract.
func bindGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GaugeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gauge *GaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gauge.Contract.GaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gauge *GaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.Contract.GaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gauge *GaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gauge.Contract.GaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gauge *GaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gauge *GaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gauge *GaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gauge.Contract.contract.Transact(opts, method, params...)
}

// DILL is a free data retrieval call binding the contract method 0x4eb1fbeb.
//
// Solidity: function DILL() view returns(address)
func (_Gauge *GaugeCaller) DILL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "DILL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DILL is a free data retrieval call binding the contract method 0x4eb1fbeb.
//
// Solidity: function DILL() view returns(address)
func (_Gauge *GaugeSession) DILL() (common.Address, error) {
	return _Gauge.Contract.DILL(&_Gauge.CallOpts)
}

// DILL is a free data retrieval call binding the contract method 0x4eb1fbeb.
//
// Solidity: function DILL() view returns(address)
func (_Gauge *GaugeCallerSession) DILL() (common.Address, error) {
	return _Gauge.Contract.DILL(&_Gauge.CallOpts)
}

// DISTRIBUTION is a free data retrieval call binding the contract method 0x7c91e4eb.
//
// Solidity: function DISTRIBUTION() view returns(address)
func (_Gauge *GaugeCaller) DISTRIBUTION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "DISTRIBUTION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DISTRIBUTION is a free data retrieval call binding the contract method 0x7c91e4eb.
//
// Solidity: function DISTRIBUTION() view returns(address)
func (_Gauge *GaugeSession) DISTRIBUTION() (common.Address, error) {
	return _Gauge.Contract.DISTRIBUTION(&_Gauge.CallOpts)
}

// DISTRIBUTION is a free data retrieval call binding the contract method 0x7c91e4eb.
//
// Solidity: function DISTRIBUTION() view returns(address)
func (_Gauge *GaugeCallerSession) DISTRIBUTION() (common.Address, error) {
	return _Gauge.Contract.DISTRIBUTION(&_Gauge.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Gauge *GaugeCaller) DURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Gauge *GaugeSession) DURATION() (*big.Int, error) {
	return _Gauge.Contract.DURATION(&_Gauge.CallOpts)
}

// DURATION is a free data retrieval call binding the contract method 0x1be05289.
//
// Solidity: function DURATION() view returns(uint256)
func (_Gauge *GaugeCallerSession) DURATION() (*big.Int, error) {
	return _Gauge.Contract.DURATION(&_Gauge.CallOpts)
}

// PICKLE is a free data retrieval call binding the contract method 0xc19048b2.
//
// Solidity: function PICKLE() view returns(address)
func (_Gauge *GaugeCaller) PICKLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "PICKLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PICKLE is a free data retrieval call binding the contract method 0xc19048b2.
//
// Solidity: function PICKLE() view returns(address)
func (_Gauge *GaugeSession) PICKLE() (common.Address, error) {
	return _Gauge.Contract.PICKLE(&_Gauge.CallOpts)
}

// PICKLE is a free data retrieval call binding the contract method 0xc19048b2.
//
// Solidity: function PICKLE() view returns(address)
func (_Gauge *GaugeCallerSession) PICKLE() (common.Address, error) {
	return _Gauge.Contract.PICKLE(&_Gauge.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Gauge *GaugeCaller) TOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Gauge *GaugeSession) TOKEN() (common.Address, error) {
	return _Gauge.Contract.TOKEN(&_Gauge.CallOpts)
}

// TOKEN is a free data retrieval call binding the contract method 0x82bfefc8.
//
// Solidity: function TOKEN() view returns(address)
func (_Gauge *GaugeCallerSession) TOKEN() (common.Address, error) {
	return _Gauge.Contract.TOKEN(&_Gauge.CallOpts)
}

// TREASURY is a free data retrieval call binding the contract method 0x2d2c5565.
//
// Solidity: function TREASURY() view returns(address)
func (_Gauge *GaugeCaller) TREASURY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "TREASURY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TREASURY is a free data retrieval call binding the contract method 0x2d2c5565.
//
// Solidity: function TREASURY() view returns(address)
func (_Gauge *GaugeSession) TREASURY() (common.Address, error) {
	return _Gauge.Contract.TREASURY(&_Gauge.CallOpts)
}

// TREASURY is a free data retrieval call binding the contract method 0x2d2c5565.
//
// Solidity: function TREASURY() view returns(address)
func (_Gauge *GaugeCallerSession) TREASURY() (common.Address, error) {
	return _Gauge.Contract.TREASURY(&_Gauge.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gauge *GaugeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gauge *GaugeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.BalanceOf(&_Gauge.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gauge *GaugeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.BalanceOf(&_Gauge.CallOpts, account)
}

// DerivedBalance is a free data retrieval call binding the contract method 0xd35e2544.
//
// Solidity: function derivedBalance(address account) view returns(uint256)
func (_Gauge *GaugeCaller) DerivedBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "derivedBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DerivedBalance is a free data retrieval call binding the contract method 0xd35e2544.
//
// Solidity: function derivedBalance(address account) view returns(uint256)
func (_Gauge *GaugeSession) DerivedBalance(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalance(&_Gauge.CallOpts, account)
}

// DerivedBalance is a free data retrieval call binding the contract method 0xd35e2544.
//
// Solidity: function derivedBalance(address account) view returns(uint256)
func (_Gauge *GaugeCallerSession) DerivedBalance(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalance(&_Gauge.CallOpts, account)
}

// DerivedBalances is a free data retrieval call binding the contract method 0x63fb415b.
//
// Solidity: function derivedBalances(address ) view returns(uint256)
func (_Gauge *GaugeCaller) DerivedBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "derivedBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DerivedBalances is a free data retrieval call binding the contract method 0x63fb415b.
//
// Solidity: function derivedBalances(address ) view returns(uint256)
func (_Gauge *GaugeSession) DerivedBalances(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalances(&_Gauge.CallOpts, arg0)
}

// DerivedBalances is a free data retrieval call binding the contract method 0x63fb415b.
//
// Solidity: function derivedBalances(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) DerivedBalances(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.DerivedBalances(&_Gauge.CallOpts, arg0)
}

// DerivedSupply is a free data retrieval call binding the contract method 0xd7da4bb0.
//
// Solidity: function derivedSupply() view returns(uint256)
func (_Gauge *GaugeCaller) DerivedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "derivedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DerivedSupply is a free data retrieval call binding the contract method 0xd7da4bb0.
//
// Solidity: function derivedSupply() view returns(uint256)
func (_Gauge *GaugeSession) DerivedSupply() (*big.Int, error) {
	return _Gauge.Contract.DerivedSupply(&_Gauge.CallOpts)
}

// DerivedSupply is a free data retrieval call binding the contract method 0xd7da4bb0.
//
// Solidity: function derivedSupply() view returns(uint256)
func (_Gauge *GaugeCallerSession) DerivedSupply() (*big.Int, error) {
	return _Gauge.Contract.DerivedSupply(&_Gauge.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Gauge *GaugeCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Gauge *GaugeSession) Earned(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.Earned(&_Gauge.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_Gauge *GaugeCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _Gauge.Contract.Earned(&_Gauge.CallOpts, account)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Gauge *GaugeCaller) GetRewardForDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "getRewardForDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Gauge *GaugeSession) GetRewardForDuration() (*big.Int, error) {
	return _Gauge.Contract.GetRewardForDuration(&_Gauge.CallOpts)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0x1c1f78eb.
//
// Solidity: function getRewardForDuration() view returns(uint256)
func (_Gauge *GaugeCallerSession) GetRewardForDuration() (*big.Int, error) {
	return _Gauge.Contract.GetRewardForDuration(&_Gauge.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Gauge *GaugeCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Gauge *GaugeSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Gauge.Contract.LastTimeRewardApplicable(&_Gauge.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_Gauge *GaugeCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _Gauge.Contract.LastTimeRewardApplicable(&_Gauge.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Gauge *GaugeCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Gauge *GaugeSession) LastUpdateTime() (*big.Int, error) {
	return _Gauge.Contract.LastUpdateTime(&_Gauge.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_Gauge *GaugeCallerSession) LastUpdateTime() (*big.Int, error) {
	return _Gauge.Contract.LastUpdateTime(&_Gauge.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Gauge *GaugeCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Gauge *GaugeSession) PeriodFinish() (*big.Int, error) {
	return _Gauge.Contract.PeriodFinish(&_Gauge.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_Gauge *GaugeCallerSession) PeriodFinish() (*big.Int, error) {
	return _Gauge.Contract.PeriodFinish(&_Gauge.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Gauge *GaugeCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Gauge *GaugeSession) RewardPerToken() (*big.Int, error) {
	return _Gauge.Contract.RewardPerToken(&_Gauge.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardPerToken() (*big.Int, error) {
	return _Gauge.Contract.RewardPerToken(&_Gauge.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Gauge *GaugeCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Gauge *GaugeSession) RewardPerTokenStored() (*big.Int, error) {
	return _Gauge.Contract.RewardPerTokenStored(&_Gauge.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _Gauge.Contract.RewardPerTokenStored(&_Gauge.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Gauge *GaugeCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Gauge *GaugeSession) RewardRate() (*big.Int, error) {
	return _Gauge.Contract.RewardRate(&_Gauge.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_Gauge *GaugeCallerSession) RewardRate() (*big.Int, error) {
	return _Gauge.Contract.RewardRate(&_Gauge.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Gauge *GaugeCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Gauge *GaugeSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.Rewards(&_Gauge.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.Rewards(&_Gauge.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeSession) TotalSupply() (*big.Int, error) {
	return _Gauge.Contract.TotalSupply(&_Gauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gauge *GaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _Gauge.Contract.TotalSupply(&_Gauge.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Gauge *GaugeCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gauge.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Gauge *GaugeSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.UserRewardPerTokenPaid(&_Gauge.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_Gauge *GaugeCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _Gauge.Contract.UserRewardPerTokenPaid(&_Gauge.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Gauge *GaugeTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Gauge *GaugeSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit(&_Gauge.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Gauge *GaugeTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Deposit(&_Gauge.TransactOpts, amount)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Gauge *GaugeTransactor) DepositAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "depositAll")
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Gauge *GaugeSession) DepositAll() (*types.Transaction, error) {
	return _Gauge.Contract.DepositAll(&_Gauge.TransactOpts)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Gauge *GaugeTransactorSession) DepositAll() (*types.Transaction, error) {
	return _Gauge.Contract.DepositAll(&_Gauge.TransactOpts)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (_Gauge *GaugeTransactor) DepositFor(opts *bind.TransactOpts, amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "depositFor", amount, account)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (_Gauge *GaugeSession) DepositFor(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.DepositFor(&_Gauge.TransactOpts, amount, account)
}

// DepositFor is a paid mutator transaction binding the contract method 0x36efd16f.
//
// Solidity: function depositFor(uint256 amount, address account) returns()
func (_Gauge *GaugeTransactorSession) DepositFor(amount *big.Int, account common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.DepositFor(&_Gauge.TransactOpts, amount, account)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Gauge *GaugeTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Gauge *GaugeSession) Exit() (*types.Transaction, error) {
	return _Gauge.Contract.Exit(&_Gauge.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Gauge *GaugeTransactorSession) Exit() (*types.Transaction, error) {
	return _Gauge.Contract.Exit(&_Gauge.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Gauge *GaugeTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Gauge *GaugeSession) GetReward() (*types.Transaction, error) {
	return _Gauge.Contract.GetReward(&_Gauge.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Gauge *GaugeTransactorSession) GetReward() (*types.Transaction, error) {
	return _Gauge.Contract.GetReward(&_Gauge.TransactOpts)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address account) returns()
func (_Gauge *GaugeTransactor) Kick(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "kick", account)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address account) returns()
func (_Gauge *GaugeSession) Kick(account common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.Kick(&_Gauge.TransactOpts, account)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address account) returns()
func (_Gauge *GaugeTransactorSession) Kick(account common.Address) (*types.Transaction, error) {
	return _Gauge.Contract.Kick(&_Gauge.TransactOpts, account)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Gauge *GaugeTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "notifyRewardAmount", reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Gauge *GaugeSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.NotifyRewardAmount(&_Gauge.TransactOpts, reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x3c6b16ab.
//
// Solidity: function notifyRewardAmount(uint256 reward) returns()
func (_Gauge *GaugeTransactorSession) NotifyRewardAmount(reward *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.NotifyRewardAmount(&_Gauge.TransactOpts, reward)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gauge *GaugeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gauge *GaugeSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Withdraw(&_Gauge.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Gauge *GaugeTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Gauge.Contract.Withdraw(&_Gauge.TransactOpts, amount)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Gauge *GaugeTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gauge.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Gauge *GaugeSession) WithdrawAll() (*types.Transaction, error) {
	return _Gauge.Contract.WithdrawAll(&_Gauge.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Gauge *GaugeTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Gauge.Contract.WithdrawAll(&_Gauge.TransactOpts)
}

// GaugeRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Gauge contract.
type GaugeRewardAddedIterator struct {
	Event *GaugeRewardAdded // Event containing the contract specifics and raw log

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
func (it *GaugeRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeRewardAdded)
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
		it.Event = new(GaugeRewardAdded)
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
func (it *GaugeRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeRewardAdded represents a RewardAdded event raised by the Gauge contract.
type GaugeRewardAdded struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Gauge *GaugeFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*GaugeRewardAddedIterator, error) {

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &GaugeRewardAddedIterator{contract: _Gauge.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Gauge *GaugeFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *GaugeRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeRewardAdded)
				if err := _Gauge.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xde88a922e0d3b88b24e9623efeb464919c6bf9f66857a65e2bfcf2ce87a9433d.
//
// Solidity: event RewardAdded(uint256 reward)
func (_Gauge *GaugeFilterer) ParseRewardAdded(log types.Log) (*GaugeRewardAdded, error) {
	event := new(GaugeRewardAdded)
	if err := _Gauge.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the Gauge contract.
type GaugeRewardPaidIterator struct {
	Event *GaugeRewardPaid // Event containing the contract specifics and raw log

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
func (it *GaugeRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeRewardPaid)
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
		it.Event = new(GaugeRewardPaid)
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
func (it *GaugeRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeRewardPaid represents a RewardPaid event raised by the Gauge contract.
type GaugeRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Gauge *GaugeFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*GaugeRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &GaugeRewardPaidIterator{contract: _Gauge.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Gauge *GaugeFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *GaugeRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeRewardPaid)
				if err := _Gauge.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_Gauge *GaugeFilterer) ParseRewardPaid(log types.Log) (*GaugeRewardPaid, error) {
	event := new(GaugeRewardPaid)
	if err := _Gauge.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Gauge contract.
type GaugeStakedIterator struct {
	Event *GaugeStaked // Event containing the contract specifics and raw log

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
func (it *GaugeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeStaked)
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
		it.Event = new(GaugeStaked)
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
func (it *GaugeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeStaked represents a Staked event raised by the Gauge contract.
type GaugeStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Gauge *GaugeFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*GaugeStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &GaugeStakedIterator{contract: _Gauge.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Gauge *GaugeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *GaugeStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeStaked)
				if err := _Gauge.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Gauge *GaugeFilterer) ParseStaked(log types.Log) (*GaugeStaked, error) {
	event := new(GaugeStaked)
	if err := _Gauge.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GaugeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Gauge contract.
type GaugeWithdrawnIterator struct {
	Event *GaugeWithdrawn // Event containing the contract specifics and raw log

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
func (it *GaugeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GaugeWithdrawn)
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
		it.Event = new(GaugeWithdrawn)
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
func (it *GaugeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GaugeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GaugeWithdrawn represents a Withdrawn event raised by the Gauge contract.
type GaugeWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Gauge *GaugeFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*GaugeWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gauge.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &GaugeWithdrawnIterator{contract: _Gauge.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Gauge *GaugeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GaugeWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Gauge.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GaugeWithdrawn)
				if err := _Gauge.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Gauge *GaugeFilterer) ParseWithdrawn(log types.Log) (*GaugeWithdrawn, error) {
	event := new(GaugeWithdrawn)
	if err := _Gauge.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
