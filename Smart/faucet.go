// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Smart

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

// FaucetContractABI is the input ABI used to generate the binding from.
const FaucetContractABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"dotransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalancecontract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"gettodaytoken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"gettransfercount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"gettransfertime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"refreshtodaytoken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usertable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transfercount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transfertime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenaccount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"todaytokenaccount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// FaucetContract is an auto generated Go binding around an Ethereum contract.
type FaucetContract struct {
	FaucetContractCaller     // Read-only binding to the contract
	FaucetContractTransactor // Write-only binding to the contract
	FaucetContractFilterer   // Log filterer for contract events
}

// FaucetContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetContractSession struct {
	Contract     *FaucetContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetContractCallerSession struct {
	Contract *FaucetContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FaucetContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetContractTransactorSession struct {
	Contract     *FaucetContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FaucetContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetContractRaw struct {
	Contract *FaucetContract // Generic contract binding to access the raw methods on
}

// FaucetContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetContractCallerRaw struct {
	Contract *FaucetContractCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetContractTransactorRaw struct {
	Contract *FaucetContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucetContract creates a new instance of FaucetContract, bound to a specific deployed contract.
func NewFaucetContract(address common.Address, backend bind.ContractBackend) (*FaucetContract, error) {
	contract, err := bindFaucetContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FaucetContract{FaucetContractCaller: FaucetContractCaller{contract: contract}, FaucetContractTransactor: FaucetContractTransactor{contract: contract}, FaucetContractFilterer: FaucetContractFilterer{contract: contract}}, nil
}

// NewFaucetContractCaller creates a new read-only instance of FaucetContract, bound to a specific deployed contract.
func NewFaucetContractCaller(address common.Address, caller bind.ContractCaller) (*FaucetContractCaller, error) {
	contract, err := bindFaucetContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetContractCaller{contract: contract}, nil
}

// NewFaucetContractTransactor creates a new write-only instance of FaucetContract, bound to a specific deployed contract.
func NewFaucetContractTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetContractTransactor, error) {
	contract, err := bindFaucetContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetContractTransactor{contract: contract}, nil
}

// NewFaucetContractFilterer creates a new log filterer instance of FaucetContract, bound to a specific deployed contract.
func NewFaucetContractFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetContractFilterer, error) {
	contract, err := bindFaucetContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetContractFilterer{contract: contract}, nil
}

// bindFaucetContract binds a generic wrapper to an already deployed contract.
func bindFaucetContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaucetContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaucetContract *FaucetContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaucetContract.Contract.FaucetContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaucetContract *FaucetContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetContract.Contract.FaucetContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaucetContract *FaucetContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaucetContract.Contract.FaucetContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaucetContract *FaucetContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaucetContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaucetContract *FaucetContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaucetContract *FaucetContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaucetContract.Contract.contract.Transact(opts, method, params...)
}

// GetBalanceUser is a free data retrieval call binding the contract method 0x31dacf77.
//
// Solidity: function getBalanceUser() view returns(uint256)
func (_FaucetContract *FaucetContractCaller) GetBalanceUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaucetContract.contract.Call(opts, &out, "getBalanceUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceUser is a free data retrieval call binding the contract method 0x31dacf77.
//
// Solidity: function getBalanceUser() view returns(uint256)
func (_FaucetContract *FaucetContractSession) GetBalanceUser() (*big.Int, error) {
	return _FaucetContract.Contract.GetBalanceUser(&_FaucetContract.CallOpts)
}

// GetBalanceUser is a free data retrieval call binding the contract method 0x31dacf77.
//
// Solidity: function getBalanceUser() view returns(uint256)
func (_FaucetContract *FaucetContractCallerSession) GetBalanceUser() (*big.Int, error) {
	return _FaucetContract.Contract.GetBalanceUser(&_FaucetContract.CallOpts)
}

// GetBalancecontract is a free data retrieval call binding the contract method 0x62afa143.
//
// Solidity: function getBalancecontract() view returns(uint256)
func (_FaucetContract *FaucetContractCaller) GetBalancecontract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FaucetContract.contract.Call(opts, &out, "getBalancecontract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalancecontract is a free data retrieval call binding the contract method 0x62afa143.
//
// Solidity: function getBalancecontract() view returns(uint256)
func (_FaucetContract *FaucetContractSession) GetBalancecontract() (*big.Int, error) {
	return _FaucetContract.Contract.GetBalancecontract(&_FaucetContract.CallOpts)
}

// GetBalancecontract is a free data retrieval call binding the contract method 0x62afa143.
//
// Solidity: function getBalancecontract() view returns(uint256)
func (_FaucetContract *FaucetContractCallerSession) GetBalancecontract() (*big.Int, error) {
	return _FaucetContract.Contract.GetBalancecontract(&_FaucetContract.CallOpts)
}

// Gettodaytoken is a free data retrieval call binding the contract method 0x1fa51531.
//
// Solidity: function gettodaytoken(address add) view returns(uint256)
func (_FaucetContract *FaucetContractCaller) Gettodaytoken(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaucetContract.contract.Call(opts, &out, "gettodaytoken", add)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettodaytoken is a free data retrieval call binding the contract method 0x1fa51531.
//
// Solidity: function gettodaytoken(address add) view returns(uint256)
func (_FaucetContract *FaucetContractSession) Gettodaytoken(add common.Address) (*big.Int, error) {
	return _FaucetContract.Contract.Gettodaytoken(&_FaucetContract.CallOpts, add)
}

// Gettodaytoken is a free data retrieval call binding the contract method 0x1fa51531.
//
// Solidity: function gettodaytoken(address add) view returns(uint256)
func (_FaucetContract *FaucetContractCallerSession) Gettodaytoken(add common.Address) (*big.Int, error) {
	return _FaucetContract.Contract.Gettodaytoken(&_FaucetContract.CallOpts, add)
}

// Gettransfercount is a free data retrieval call binding the contract method 0xd9359584.
//
// Solidity: function gettransfercount(address add) view returns(uint256)
func (_FaucetContract *FaucetContractCaller) Gettransfercount(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaucetContract.contract.Call(opts, &out, "gettransfercount", add)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettransfercount is a free data retrieval call binding the contract method 0xd9359584.
//
// Solidity: function gettransfercount(address add) view returns(uint256)
func (_FaucetContract *FaucetContractSession) Gettransfercount(add common.Address) (*big.Int, error) {
	return _FaucetContract.Contract.Gettransfercount(&_FaucetContract.CallOpts, add)
}

// Gettransfercount is a free data retrieval call binding the contract method 0xd9359584.
//
// Solidity: function gettransfercount(address add) view returns(uint256)
func (_FaucetContract *FaucetContractCallerSession) Gettransfercount(add common.Address) (*big.Int, error) {
	return _FaucetContract.Contract.Gettransfercount(&_FaucetContract.CallOpts, add)
}

// Gettransfertime is a free data retrieval call binding the contract method 0x3096f21f.
//
// Solidity: function gettransfertime(address add) view returns(uint256)
func (_FaucetContract *FaucetContractCaller) Gettransfertime(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FaucetContract.contract.Call(opts, &out, "gettransfertime", add)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettransfertime is a free data retrieval call binding the contract method 0x3096f21f.
//
// Solidity: function gettransfertime(address add) view returns(uint256)
func (_FaucetContract *FaucetContractSession) Gettransfertime(add common.Address) (*big.Int, error) {
	return _FaucetContract.Contract.Gettransfertime(&_FaucetContract.CallOpts, add)
}

// Gettransfertime is a free data retrieval call binding the contract method 0x3096f21f.
//
// Solidity: function gettransfertime(address add) view returns(uint256)
func (_FaucetContract *FaucetContractCallerSession) Gettransfertime(add common.Address) (*big.Int, error) {
	return _FaucetContract.Contract.Gettransfertime(&_FaucetContract.CallOpts, add)
}

// Usertable is a free data retrieval call binding the contract method 0xc1368e45.
//
// Solidity: function usertable(address ) view returns(address _address, uint256 transfercount, uint256 transfertime, uint256 tokenaccount, uint256 todaytokenaccount)
func (_FaucetContract *FaucetContractCaller) Usertable(opts *bind.CallOpts, arg0 common.Address) (struct {
	Address           common.Address
	Transfercount     *big.Int
	Transfertime      *big.Int
	Tokenaccount      *big.Int
	Todaytokenaccount *big.Int
}, error) {
	var out []interface{}
	err := _FaucetContract.contract.Call(opts, &out, "usertable", arg0)

	outstruct := new(struct {
		Address           common.Address
		Transfercount     *big.Int
		Transfertime      *big.Int
		Tokenaccount      *big.Int
		Todaytokenaccount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Address = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Transfercount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Transfertime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tokenaccount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Todaytokenaccount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Usertable is a free data retrieval call binding the contract method 0xc1368e45.
//
// Solidity: function usertable(address ) view returns(address _address, uint256 transfercount, uint256 transfertime, uint256 tokenaccount, uint256 todaytokenaccount)
func (_FaucetContract *FaucetContractSession) Usertable(arg0 common.Address) (struct {
	Address           common.Address
	Transfercount     *big.Int
	Transfertime      *big.Int
	Tokenaccount      *big.Int
	Todaytokenaccount *big.Int
}, error) {
	return _FaucetContract.Contract.Usertable(&_FaucetContract.CallOpts, arg0)
}

// Usertable is a free data retrieval call binding the contract method 0xc1368e45.
//
// Solidity: function usertable(address ) view returns(address _address, uint256 transfercount, uint256 transfertime, uint256 tokenaccount, uint256 todaytokenaccount)
func (_FaucetContract *FaucetContractCallerSession) Usertable(arg0 common.Address) (struct {
	Address           common.Address
	Transfercount     *big.Int
	Transfertime      *big.Int
	Tokenaccount      *big.Int
	Todaytokenaccount *big.Int
}, error) {
	return _FaucetContract.Contract.Usertable(&_FaucetContract.CallOpts, arg0)
}

// Dotransfer is a paid mutator transaction binding the contract method 0x39e74f87.
//
// Solidity: function dotransfer(address add) returns()
func (_FaucetContract *FaucetContractTransactor) Dotransfer(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _FaucetContract.contract.Transact(opts, "dotransfer", add)
}

// Dotransfer is a paid mutator transaction binding the contract method 0x39e74f87.
//
// Solidity: function dotransfer(address add) returns()
func (_FaucetContract *FaucetContractSession) Dotransfer(add common.Address) (*types.Transaction, error) {
	return _FaucetContract.Contract.Dotransfer(&_FaucetContract.TransactOpts, add)
}

// Dotransfer is a paid mutator transaction binding the contract method 0x39e74f87.
//
// Solidity: function dotransfer(address add) returns()
func (_FaucetContract *FaucetContractTransactorSession) Dotransfer(add common.Address) (*types.Transaction, error) {
	return _FaucetContract.Contract.Dotransfer(&_FaucetContract.TransactOpts, add)
}

// Refreshtodaytoken is a paid mutator transaction binding the contract method 0x8155ea46.
//
// Solidity: function refreshtodaytoken(address add) payable returns()
func (_FaucetContract *FaucetContractTransactor) Refreshtodaytoken(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _FaucetContract.contract.Transact(opts, "refreshtodaytoken", add)
}

// Refreshtodaytoken is a paid mutator transaction binding the contract method 0x8155ea46.
//
// Solidity: function refreshtodaytoken(address add) payable returns()
func (_FaucetContract *FaucetContractSession) Refreshtodaytoken(add common.Address) (*types.Transaction, error) {
	return _FaucetContract.Contract.Refreshtodaytoken(&_FaucetContract.TransactOpts, add)
}

// Refreshtodaytoken is a paid mutator transaction binding the contract method 0x8155ea46.
//
// Solidity: function refreshtodaytoken(address add) payable returns()
func (_FaucetContract *FaucetContractTransactorSession) Refreshtodaytoken(add common.Address) (*types.Transaction, error) {
	return _FaucetContract.Contract.Refreshtodaytoken(&_FaucetContract.TransactOpts, add)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FaucetContract *FaucetContractTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FaucetContract.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FaucetContract *FaucetContractSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FaucetContract.Contract.Fallback(&_FaucetContract.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FaucetContract *FaucetContractTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FaucetContract.Contract.Fallback(&_FaucetContract.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FaucetContract *FaucetContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaucetContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FaucetContract *FaucetContractSession) Receive() (*types.Transaction, error) {
	return _FaucetContract.Contract.Receive(&_FaucetContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FaucetContract *FaucetContractTransactorSession) Receive() (*types.Transaction, error) {
	return _FaucetContract.Contract.Receive(&_FaucetContract.TransactOpts)
}
