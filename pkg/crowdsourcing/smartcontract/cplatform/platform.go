// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cplatform

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

// CplatformABI is the input ABI used to generate the binding from.
const CplatformABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"increaseReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"reputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CplatformBin is the compiled bytecode used for deploying new contracts.
var CplatformBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610366806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634420e486146100515780634f76dc6114610095578063b9f79451146100d9578063dbe55e5614610131575b600080fd5b6100936004803603602081101561006757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061017b565b005b6100d7600480360360208110156100ab57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061026e565b005b61011b600480360360208110156100ef57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102be565b6040518082815260200191505060405180910390f35b610139610307565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561021e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f63616e6e6f74207472616e66657220746f207a65726f2061646472657373000081525060200191505060405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050565b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690509056fea264697066735822122092ad66a8668b5284fa43e61689c116ed9d687707e347c75f36a02d633f3bc93c64736f6c63430006030033"

// DeployCplatform deploys a new Ethereum contract, binding an instance of Cplatform to it.
func DeployCplatform(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Cplatform, error) {
	parsed, err := abi.JSON(strings.NewReader(CplatformABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CplatformBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cplatform{CplatformCaller: CplatformCaller{contract: contract}, CplatformTransactor: CplatformTransactor{contract: contract}, CplatformFilterer: CplatformFilterer{contract: contract}}, nil
}

// Cplatform is an auto generated Go binding around an Ethereum contract.
type Cplatform struct {
	CplatformCaller     // Read-only binding to the contract
	CplatformTransactor // Write-only binding to the contract
	CplatformFilterer   // Log filterer for contract events
}

// CplatformCaller is an auto generated read-only Go binding around an Ethereum contract.
type CplatformCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CplatformTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CplatformTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CplatformFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CplatformFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CplatformSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CplatformSession struct {
	Contract     *Cplatform        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CplatformCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CplatformCallerSession struct {
	Contract *CplatformCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CplatformTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CplatformTransactorSession struct {
	Contract     *CplatformTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CplatformRaw is an auto generated low-level Go binding around an Ethereum contract.
type CplatformRaw struct {
	Contract *Cplatform // Generic contract binding to access the raw methods on
}

// CplatformCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CplatformCallerRaw struct {
	Contract *CplatformCaller // Generic read-only contract binding to access the raw methods on
}

// CplatformTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CplatformTransactorRaw struct {
	Contract *CplatformTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCplatform creates a new instance of Cplatform, bound to a specific deployed contract.
func NewCplatform(address common.Address, backend bind.ContractBackend) (*Cplatform, error) {
	contract, err := bindCplatform(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cplatform{CplatformCaller: CplatformCaller{contract: contract}, CplatformTransactor: CplatformTransactor{contract: contract}, CplatformFilterer: CplatformFilterer{contract: contract}}, nil
}

// NewCplatformCaller creates a new read-only instance of Cplatform, bound to a specific deployed contract.
func NewCplatformCaller(address common.Address, caller bind.ContractCaller) (*CplatformCaller, error) {
	contract, err := bindCplatform(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CplatformCaller{contract: contract}, nil
}

// NewCplatformTransactor creates a new write-only instance of Cplatform, bound to a specific deployed contract.
func NewCplatformTransactor(address common.Address, transactor bind.ContractTransactor) (*CplatformTransactor, error) {
	contract, err := bindCplatform(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CplatformTransactor{contract: contract}, nil
}

// NewCplatformFilterer creates a new log filterer instance of Cplatform, bound to a specific deployed contract.
func NewCplatformFilterer(address common.Address, filterer bind.ContractFilterer) (*CplatformFilterer, error) {
	contract, err := bindCplatform(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CplatformFilterer{contract: contract}, nil
}

// bindCplatform binds a generic wrapper to an already deployed contract.
func bindCplatform(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CplatformABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cplatform *CplatformRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cplatform.Contract.CplatformCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cplatform *CplatformRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cplatform.Contract.CplatformTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cplatform *CplatformRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cplatform.Contract.CplatformTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cplatform *CplatformCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cplatform.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cplatform *CplatformTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cplatform.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cplatform *CplatformTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cplatform.Contract.contract.Transact(opts, method, params...)
}

// PlatformAddress is a free data retrieval call binding the contract method 0xdbe55e56.
//
// Solidity: function platformAddress() view returns(address)
func (_Cplatform *CplatformCaller) PlatformAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "platformAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PlatformAddress is a free data retrieval call binding the contract method 0xdbe55e56.
//
// Solidity: function platformAddress() view returns(address)
func (_Cplatform *CplatformSession) PlatformAddress() (common.Address, error) {
	return _Cplatform.Contract.PlatformAddress(&_Cplatform.CallOpts)
}

// PlatformAddress is a free data retrieval call binding the contract method 0xdbe55e56.
//
// Solidity: function platformAddress() view returns(address)
func (_Cplatform *CplatformCallerSession) PlatformAddress() (common.Address, error) {
	return _Cplatform.Contract.PlatformAddress(&_Cplatform.CallOpts)
}

// Reputation is a free data retrieval call binding the contract method 0xb9f79451.
//
// Solidity: function reputation(address _arr) view returns(uint256)
func (_Cplatform *CplatformCaller) Reputation(opts *bind.CallOpts, _arr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "reputation", _arr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reputation is a free data retrieval call binding the contract method 0xb9f79451.
//
// Solidity: function reputation(address _arr) view returns(uint256)
func (_Cplatform *CplatformSession) Reputation(_arr common.Address) (*big.Int, error) {
	return _Cplatform.Contract.Reputation(&_Cplatform.CallOpts, _arr)
}

// Reputation is a free data retrieval call binding the contract method 0xb9f79451.
//
// Solidity: function reputation(address _arr) view returns(uint256)
func (_Cplatform *CplatformCallerSession) Reputation(_arr common.Address) (*big.Int, error) {
	return _Cplatform.Contract.Reputation(&_Cplatform.CallOpts, _arr)
}

// IncreaseReputation is a paid mutator transaction binding the contract method 0x4f76dc61.
//
// Solidity: function increaseReputation(address _arr) returns()
func (_Cplatform *CplatformTransactor) IncreaseReputation(opts *bind.TransactOpts, _arr common.Address) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "increaseReputation", _arr)
}

// IncreaseReputation is a paid mutator transaction binding the contract method 0x4f76dc61.
//
// Solidity: function increaseReputation(address _arr) returns()
func (_Cplatform *CplatformSession) IncreaseReputation(_arr common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.IncreaseReputation(&_Cplatform.TransactOpts, _arr)
}

// IncreaseReputation is a paid mutator transaction binding the contract method 0x4f76dc61.
//
// Solidity: function increaseReputation(address _arr) returns()
func (_Cplatform *CplatformTransactorSession) IncreaseReputation(_arr common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.IncreaseReputation(&_Cplatform.TransactOpts, _arr)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address node) returns()
func (_Cplatform *CplatformTransactor) Register(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "register", node)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address node) returns()
func (_Cplatform *CplatformSession) Register(node common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.Register(&_Cplatform.TransactOpts, node)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address node) returns()
func (_Cplatform *CplatformTransactorSession) Register(node common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.Register(&_Cplatform.TransactOpts, node)
}
