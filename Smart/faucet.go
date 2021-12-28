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

// FaucetABI is the input ABI used to generate the binding from.
const FaucetABI = "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"dotransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalanceUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalancecontract\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"gettodaytoken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"gettransfercount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"gettransfertime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"add\",\"type\":\"address\"}],\"name\":\"refreshtodaytoken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usertable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transfercount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"transfertime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenaccount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"todaytokenaccount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// Faucet is an auto generated Go binding around an Ethereum contract.
type Faucet struct {
	FaucetCaller     // Read-only binding to the contract
	FaucetTransactor // Write-only binding to the contract
	FaucetFilterer   // Log filterer for contract events
}

// FaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetSession struct {
	Contract     *Faucet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetCallerSession struct {
	Contract *FaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetTransactorSession struct {
	Contract     *FaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetRaw struct {
	Contract *Faucet // Generic contract binding to access the raw methods on
}

// FaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetCallerRaw struct {
	Contract *FaucetCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetTransactorRaw struct {
	Contract *FaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucet creates a new instance of Faucet, bound to a specific deployed contract.
func NewFaucet(address common.Address, backend bind.ContractBackend) (*Faucet, error) {
	contract, err := bindFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// NewFaucetCaller creates a new read-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetCaller(address common.Address, caller bind.ContractCaller) (*FaucetCaller, error) {
	contract, err := bindFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetCaller{contract: contract}, nil
}

// NewFaucetTransactor creates a new write-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetTransactor, error) {
	contract, err := bindFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTransactor{contract: contract}, nil
}

// NewFaucetFilterer creates a new log filterer instance of Faucet, bound to a specific deployed contract.
func NewFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetFilterer, error) {
	contract, err := bindFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetFilterer{contract: contract}, nil
}

// bindFaucet binds a generic wrapper to an already deployed contract.
func bindFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FaucetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.FaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transact(opts, method, params...)
}

// GetBalanceUser is a free data retrieval call binding the contract method 0x31dacf77.
//
// Solidity: function getBalanceUser() view returns(uint256)
func (_Faucet *FaucetCaller) GetBalanceUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "getBalanceUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceUser is a free data retrieval call binding the contract method 0x31dacf77.
//
// Solidity: function getBalanceUser() view returns(uint256)
func (_Faucet *FaucetSession) GetBalanceUser() (*big.Int, error) {
	return _Faucet.Contract.GetBalanceUser(&_Faucet.CallOpts)
}

// GetBalanceUser is a free data retrieval call binding the contract method 0x31dacf77.
//
// Solidity: function getBalanceUser() view returns(uint256)
func (_Faucet *FaucetCallerSession) GetBalanceUser() (*big.Int, error) {
	return _Faucet.Contract.GetBalanceUser(&_Faucet.CallOpts)
}

// GetBalancecontract is a free data retrieval call binding the contract method 0x62afa143.
//
// Solidity: function getBalancecontract() view returns(uint256)
func (_Faucet *FaucetCaller) GetBalancecontract(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "getBalancecontract")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalancecontract is a free data retrieval call binding the contract method 0x62afa143.
//
// Solidity: function getBalancecontract() view returns(uint256)
func (_Faucet *FaucetSession) GetBalancecontract() (*big.Int, error) {
	return _Faucet.Contract.GetBalancecontract(&_Faucet.CallOpts)
}

// GetBalancecontract is a free data retrieval call binding the contract method 0x62afa143.
//
// Solidity: function getBalancecontract() view returns(uint256)
func (_Faucet *FaucetCallerSession) GetBalancecontract() (*big.Int, error) {
	return _Faucet.Contract.GetBalancecontract(&_Faucet.CallOpts)
}

// Gettodaytoken is a free data retrieval call binding the contract method 0x1fa51531.
//
// Solidity: function gettodaytoken(address add) view returns(uint256)
func (_Faucet *FaucetCaller) Gettodaytoken(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "gettodaytoken", add)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettodaytoken is a free data retrieval call binding the contract method 0x1fa51531.
//
// Solidity: function gettodaytoken(address add) view returns(uint256)
func (_Faucet *FaucetSession) Gettodaytoken(add common.Address) (*big.Int, error) {
	return _Faucet.Contract.Gettodaytoken(&_Faucet.CallOpts, add)
}

// Gettodaytoken is a free data retrieval call binding the contract method 0x1fa51531.
//
// Solidity: function gettodaytoken(address add) view returns(uint256)
func (_Faucet *FaucetCallerSession) Gettodaytoken(add common.Address) (*big.Int, error) {
	return _Faucet.Contract.Gettodaytoken(&_Faucet.CallOpts, add)
}

// Gettransfercount is a free data retrieval call binding the contract method 0xd9359584.
//
// Solidity: function gettransfercount(address add) view returns(uint256)
func (_Faucet *FaucetCaller) Gettransfercount(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "gettransfercount", add)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettransfercount is a free data retrieval call binding the contract method 0xd9359584.
//
// Solidity: function gettransfercount(address add) view returns(uint256)
func (_Faucet *FaucetSession) Gettransfercount(add common.Address) (*big.Int, error) {
	return _Faucet.Contract.Gettransfercount(&_Faucet.CallOpts, add)
}

// Gettransfercount is a free data retrieval call binding the contract method 0xd9359584.
//
// Solidity: function gettransfercount(address add) view returns(uint256)
func (_Faucet *FaucetCallerSession) Gettransfercount(add common.Address) (*big.Int, error) {
	return _Faucet.Contract.Gettransfercount(&_Faucet.CallOpts, add)
}

// Gettransfertime is a free data retrieval call binding the contract method 0x3096f21f.
//
// Solidity: function gettransfertime(address add) view returns(uint256)
func (_Faucet *FaucetCaller) Gettransfertime(opts *bind.CallOpts, add common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "gettransfertime", add)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gettransfertime is a free data retrieval call binding the contract method 0x3096f21f.
//
// Solidity: function gettransfertime(address add) view returns(uint256)
func (_Faucet *FaucetSession) Gettransfertime(add common.Address) (*big.Int, error) {
	return _Faucet.Contract.Gettransfertime(&_Faucet.CallOpts, add)
}

// Gettransfertime is a free data retrieval call binding the contract method 0x3096f21f.
//
// Solidity: function gettransfertime(address add) view returns(uint256)
func (_Faucet *FaucetCallerSession) Gettransfertime(add common.Address) (*big.Int, error) {
	return _Faucet.Contract.Gettransfertime(&_Faucet.CallOpts, add)
}

// Usertable is a free data retrieval call binding the contract method 0xc1368e45.
//
// Solidity: function usertable(address ) view returns(address _address, uint256 transfercount, uint256 transfertime, uint256 tokenaccount, uint256 todaytokenaccount)
func (_Faucet *FaucetCaller) Usertable(opts *bind.CallOpts, arg0 common.Address) (struct {
	Address           common.Address
	Transfercount     *big.Int
	Transfertime      *big.Int
	Tokenaccount      *big.Int
	Todaytokenaccount *big.Int
}, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "usertable", arg0)

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
func (_Faucet *FaucetSession) Usertable(arg0 common.Address) (struct {
	Address           common.Address
	Transfercount     *big.Int
	Transfertime      *big.Int
	Tokenaccount      *big.Int
	Todaytokenaccount *big.Int
}, error) {
	return _Faucet.Contract.Usertable(&_Faucet.CallOpts, arg0)
}

// Usertable is a free data retrieval call binding the contract method 0xc1368e45.
//
// Solidity: function usertable(address ) view returns(address _address, uint256 transfercount, uint256 transfertime, uint256 tokenaccount, uint256 todaytokenaccount)
func (_Faucet *FaucetCallerSession) Usertable(arg0 common.Address) (struct {
	Address           common.Address
	Transfercount     *big.Int
	Transfertime      *big.Int
	Tokenaccount      *big.Int
	Todaytokenaccount *big.Int
}, error) {
	return _Faucet.Contract.Usertable(&_Faucet.CallOpts, arg0)
}

// Dotransfer is a paid mutator transaction binding the contract method 0x39e74f87.
//
// Solidity: function dotransfer(address add) returns()
func (_Faucet *FaucetTransactor) Dotransfer(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "dotransfer", add)
}

// Dotransfer is a paid mutator transaction binding the contract method 0x39e74f87.
//
// Solidity: function dotransfer(address add) returns()
func (_Faucet *FaucetSession) Dotransfer(add common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.Dotransfer(&_Faucet.TransactOpts, add)
}

// Dotransfer is a paid mutator transaction binding the contract method 0x39e74f87.
//
// Solidity: function dotransfer(address add) returns()
func (_Faucet *FaucetTransactorSession) Dotransfer(add common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.Dotransfer(&_Faucet.TransactOpts, add)
}

// Refreshtodaytoken is a paid mutator transaction binding the contract method 0x8155ea46.
//
// Solidity: function refreshtodaytoken(address add) payable returns()
func (_Faucet *FaucetTransactor) Refreshtodaytoken(opts *bind.TransactOpts, add common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "refreshtodaytoken", add)
}

// Refreshtodaytoken is a paid mutator transaction binding the contract method 0x8155ea46.
//
// Solidity: function refreshtodaytoken(address add) payable returns()
func (_Faucet *FaucetSession) Refreshtodaytoken(add common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.Refreshtodaytoken(&_Faucet.TransactOpts, add)
}

// Refreshtodaytoken is a paid mutator transaction binding the contract method 0x8155ea46.
//
// Solidity: function refreshtodaytoken(address add) payable returns()
func (_Faucet *FaucetTransactorSession) Refreshtodaytoken(add common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.Refreshtodaytoken(&_Faucet.TransactOpts, add)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Faucet *FaucetTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Faucet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Faucet *FaucetSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Faucet.Contract.Fallback(&_Faucet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Faucet *FaucetTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Faucet.Contract.Fallback(&_Faucet.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetSession) Receive() (*types.Transaction, error) {
	return _Faucet.Contract.Receive(&_Faucet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetTransactorSession) Receive() (*types.Transaction, error) {
	return _Faucet.Contract.Receive(&_Faucet.TransactOpts)
}
