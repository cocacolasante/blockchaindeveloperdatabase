// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package credits

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

// CreditsMetaData contains all meta data concerning the Credits contract.
var CreditsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mintToAddress\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numOfTokens\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeemCredits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"usersRedeemed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CreditsABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditsMetaData.ABI instead.
var CreditsABI = CreditsMetaData.ABI

// Credits is an auto generated Go binding around an Ethereum contract.
type Credits struct {
	CreditsCaller     // Read-only binding to the contract
	CreditsTransactor // Write-only binding to the contract
	CreditsFilterer   // Log filterer for contract events
}

// CreditsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditsSession struct {
	Contract     *Credits          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditsCallerSession struct {
	Contract *CreditsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CreditsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditsTransactorSession struct {
	Contract     *CreditsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CreditsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditsRaw struct {
	Contract *Credits // Generic contract binding to access the raw methods on
}

// CreditsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditsCallerRaw struct {
	Contract *CreditsCaller // Generic read-only contract binding to access the raw methods on
}

// CreditsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditsTransactorRaw struct {
	Contract *CreditsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredits creates a new instance of Credits, bound to a specific deployed contract.
func NewCredits(address common.Address, backend bind.ContractBackend) (*Credits, error) {
	contract, err := bindCredits(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Credits{CreditsCaller: CreditsCaller{contract: contract}, CreditsTransactor: CreditsTransactor{contract: contract}, CreditsFilterer: CreditsFilterer{contract: contract}}, nil
}

// NewCreditsCaller creates a new read-only instance of Credits, bound to a specific deployed contract.
func NewCreditsCaller(address common.Address, caller bind.ContractCaller) (*CreditsCaller, error) {
	contract, err := bindCredits(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditsCaller{contract: contract}, nil
}

// NewCreditsTransactor creates a new write-only instance of Credits, bound to a specific deployed contract.
func NewCreditsTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditsTransactor, error) {
	contract, err := bindCredits(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditsTransactor{contract: contract}, nil
}

// NewCreditsFilterer creates a new log filterer instance of Credits, bound to a specific deployed contract.
func NewCreditsFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditsFilterer, error) {
	contract, err := bindCredits(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditsFilterer{contract: contract}, nil
}

// bindCredits binds a generic wrapper to an already deployed contract.
func bindCredits(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CreditsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credits *CreditsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credits.Contract.CreditsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credits *CreditsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credits.Contract.CreditsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credits *CreditsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credits.Contract.CreditsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credits *CreditsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credits.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credits *CreditsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credits.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credits *CreditsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credits.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credits *CreditsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Credits.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credits *CreditsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Credits.Contract.Allowance(&_Credits.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credits *CreditsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Credits.Contract.Allowance(&_Credits.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credits *CreditsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Credits.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credits *CreditsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Credits.Contract.BalanceOf(&_Credits.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credits *CreditsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Credits.Contract.BalanceOf(&_Credits.CallOpts, account)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Credits *CreditsCaller) GetContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Credits.contract.Call(opts, &out, "getContractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Credits *CreditsSession) GetContractBalance() (*big.Int, error) {
	return _Credits.Contract.GetContractBalance(&_Credits.CallOpts)
}

// GetContractBalance is a free data retrieval call binding the contract method 0x6f9fb98a.
//
// Solidity: function getContractBalance() view returns(uint256)
func (_Credits *CreditsCallerSession) GetContractBalance() (*big.Int, error) {
	return _Credits.Contract.GetContractBalance(&_Credits.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credits *CreditsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Credits.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credits *CreditsSession) Owner() (common.Address, error) {
	return _Credits.Contract.Owner(&_Credits.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credits *CreditsCallerSession) Owner() (common.Address, error) {
	return _Credits.Contract.Owner(&_Credits.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256)
func (_Credits *CreditsCaller) TokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Credits.contract.Call(opts, &out, "tokenPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256)
func (_Credits *CreditsSession) TokenPrice() (*big.Int, error) {
	return _Credits.Contract.TokenPrice(&_Credits.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256)
func (_Credits *CreditsCallerSession) TokenPrice() (*big.Int, error) {
	return _Credits.Contract.TokenPrice(&_Credits.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credits *CreditsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Credits.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credits *CreditsSession) TotalSupply() (*big.Int, error) {
	return _Credits.Contract.TotalSupply(&_Credits.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credits *CreditsCallerSession) TotalSupply() (*big.Int, error) {
	return _Credits.Contract.TotalSupply(&_Credits.CallOpts)
}

// UsersRedeemed is a free data retrieval call binding the contract method 0x177cbfa1.
//
// Solidity: function usersRedeemed(address user) view returns(uint256)
func (_Credits *CreditsCaller) UsersRedeemed(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Credits.contract.Call(opts, &out, "usersRedeemed", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsersRedeemed is a free data retrieval call binding the contract method 0x177cbfa1.
//
// Solidity: function usersRedeemed(address user) view returns(uint256)
func (_Credits *CreditsSession) UsersRedeemed(user common.Address) (*big.Int, error) {
	return _Credits.Contract.UsersRedeemed(&_Credits.CallOpts, user)
}

// UsersRedeemed is a free data retrieval call binding the contract method 0x177cbfa1.
//
// Solidity: function usersRedeemed(address user) view returns(uint256)
func (_Credits *CreditsCallerSession) UsersRedeemed(user common.Address) (*big.Int, error) {
	return _Credits.Contract.UsersRedeemed(&_Credits.CallOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Credits *CreditsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Credits *CreditsSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.Approve(&_Credits.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Credits *CreditsTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.Approve(&_Credits.TransactOpts, spender, value)
}

// MintToAddress is a paid mutator transaction binding the contract method 0x21ca4236.
//
// Solidity: function mintToAddress(address target, uint256 amount) payable returns()
func (_Credits *CreditsTransactor) MintToAddress(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credits.contract.Transact(opts, "mintToAddress", target, amount)
}

// MintToAddress is a paid mutator transaction binding the contract method 0x21ca4236.
//
// Solidity: function mintToAddress(address target, uint256 amount) payable returns()
func (_Credits *CreditsSession) MintToAddress(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.MintToAddress(&_Credits.TransactOpts, target, amount)
}

// MintToAddress is a paid mutator transaction binding the contract method 0x21ca4236.
//
// Solidity: function mintToAddress(address target, uint256 amount) payable returns()
func (_Credits *CreditsTransactorSession) MintToAddress(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.MintToAddress(&_Credits.TransactOpts, target, amount)
}

// MintTokens is a paid mutator transaction binding the contract method 0x97304ced.
//
// Solidity: function mintTokens(uint256 numOfTokens) payable returns()
func (_Credits *CreditsTransactor) MintTokens(opts *bind.TransactOpts, numOfTokens *big.Int) (*types.Transaction, error) {
	return _Credits.contract.Transact(opts, "mintTokens", numOfTokens)
}

// MintTokens is a paid mutator transaction binding the contract method 0x97304ced.
//
// Solidity: function mintTokens(uint256 numOfTokens) payable returns()
func (_Credits *CreditsSession) MintTokens(numOfTokens *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.MintTokens(&_Credits.TransactOpts, numOfTokens)
}

// MintTokens is a paid mutator transaction binding the contract method 0x97304ced.
//
// Solidity: function mintTokens(uint256 numOfTokens) payable returns()
func (_Credits *CreditsTransactorSession) MintTokens(numOfTokens *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.MintTokens(&_Credits.TransactOpts, numOfTokens)
}

// RedeemCredits is a paid mutator transaction binding the contract method 0xf1a73c3e.
//
// Solidity: function redeemCredits(address target, uint256 amount) returns()
func (_Credits *CreditsTransactor) RedeemCredits(opts *bind.TransactOpts, target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credits.contract.Transact(opts, "redeemCredits", target, amount)
}

// RedeemCredits is a paid mutator transaction binding the contract method 0xf1a73c3e.
//
// Solidity: function redeemCredits(address target, uint256 amount) returns()
func (_Credits *CreditsSession) RedeemCredits(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.RedeemCredits(&_Credits.TransactOpts, target, amount)
}

// RedeemCredits is a paid mutator transaction binding the contract method 0xf1a73c3e.
//
// Solidity: function redeemCredits(address target, uint256 amount) returns()
func (_Credits *CreditsTransactorSession) RedeemCredits(target common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.RedeemCredits(&_Credits.TransactOpts, target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Credits *CreditsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Credits *CreditsSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.Transfer(&_Credits.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Credits *CreditsTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.Transfer(&_Credits.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Credits *CreditsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Credits *CreditsSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.TransferFrom(&_Credits.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Credits *CreditsTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Credits.Contract.TransferFrom(&_Credits.TransactOpts, from, to, value)
}

// Withdrawl is a paid mutator transaction binding the contract method 0x3aedfb8b.
//
// Solidity: function withdrawl() returns()
func (_Credits *CreditsTransactor) Withdrawl(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credits.contract.Transact(opts, "withdrawl")
}

// Withdrawl is a paid mutator transaction binding the contract method 0x3aedfb8b.
//
// Solidity: function withdrawl() returns()
func (_Credits *CreditsSession) Withdrawl() (*types.Transaction, error) {
	return _Credits.Contract.Withdrawl(&_Credits.TransactOpts)
}

// Withdrawl is a paid mutator transaction binding the contract method 0x3aedfb8b.
//
// Solidity: function withdrawl() returns()
func (_Credits *CreditsTransactorSession) Withdrawl() (*types.Transaction, error) {
	return _Credits.Contract.Withdrawl(&_Credits.TransactOpts)
}

// CreditsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Credits contract.
type CreditsApprovalIterator struct {
	Event *CreditsApproval // Event containing the contract specifics and raw log

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
func (it *CreditsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditsApproval)
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
		it.Event = new(CreditsApproval)
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
func (it *CreditsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditsApproval represents a Approval event raised by the Credits contract.
type CreditsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Credits *CreditsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CreditsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Credits.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CreditsApprovalIterator{contract: _Credits.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Credits *CreditsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CreditsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Credits.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditsApproval)
				if err := _Credits.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Credits *CreditsFilterer) ParseApproval(log types.Log) (*CreditsApproval, error) {
	event := new(CreditsApproval)
	if err := _Credits.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditsTokensRedeemedIterator is returned from FilterTokensRedeemed and is used to iterate over the raw logs and unpacked data for TokensRedeemed events raised by the Credits contract.
type CreditsTokensRedeemedIterator struct {
	Event *CreditsTokensRedeemed // Event containing the contract specifics and raw log

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
func (it *CreditsTokensRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditsTokensRedeemed)
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
		it.Event = new(CreditsTokensRedeemed)
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
func (it *CreditsTokensRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditsTokensRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditsTokensRedeemed represents a TokensRedeemed event raised by the Credits contract.
type CreditsTokensRedeemed struct {
	Redeemer common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokensRedeemed is a free log retrieval operation binding the contract event 0xcdd4c59de26b3f8623e6012ab9bc32a995eda3452d89c941fc8af6b74276ad38.
//
// Solidity: event TokensRedeemed(address indexed redeemer, uint256 indexed amount)
func (_Credits *CreditsFilterer) FilterTokensRedeemed(opts *bind.FilterOpts, redeemer []common.Address, amount []*big.Int) (*CreditsTokensRedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Credits.contract.FilterLogs(opts, "TokensRedeemed", redeemerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &CreditsTokensRedeemedIterator{contract: _Credits.contract, event: "TokensRedeemed", logs: logs, sub: sub}, nil
}

// WatchTokensRedeemed is a free log subscription operation binding the contract event 0xcdd4c59de26b3f8623e6012ab9bc32a995eda3452d89c941fc8af6b74276ad38.
//
// Solidity: event TokensRedeemed(address indexed redeemer, uint256 indexed amount)
func (_Credits *CreditsFilterer) WatchTokensRedeemed(opts *bind.WatchOpts, sink chan<- *CreditsTokensRedeemed, redeemer []common.Address, amount []*big.Int) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Credits.contract.WatchLogs(opts, "TokensRedeemed", redeemerRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditsTokensRedeemed)
				if err := _Credits.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
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

// ParseTokensRedeemed is a log parse operation binding the contract event 0xcdd4c59de26b3f8623e6012ab9bc32a995eda3452d89c941fc8af6b74276ad38.
//
// Solidity: event TokensRedeemed(address indexed redeemer, uint256 indexed amount)
func (_Credits *CreditsFilterer) ParseTokensRedeemed(log types.Log) (*CreditsTokensRedeemed, error) {
	event := new(CreditsTokensRedeemed)
	if err := _Credits.contract.UnpackLog(event, "TokensRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Credits contract.
type CreditsTransferIterator struct {
	Event *CreditsTransfer // Event containing the contract specifics and raw log

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
func (it *CreditsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditsTransfer)
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
		it.Event = new(CreditsTransfer)
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
func (it *CreditsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditsTransfer represents a Transfer event raised by the Credits contract.
type CreditsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Credits *CreditsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CreditsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Credits.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditsTransferIterator{contract: _Credits.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Credits *CreditsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CreditsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Credits.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditsTransfer)
				if err := _Credits.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Credits *CreditsFilterer) ParseTransfer(log types.Log) (*CreditsTransfer, error) {
	event := new(CreditsTransfer)
	if err := _Credits.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
