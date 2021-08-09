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
const CplatformABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maskedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unmaskedRewards\",\"type\":\"uint256\"},{\"internalType\":\"contractTask\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"Rewarding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"decreaseRepuation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"increaseReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"reputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CplatformBin is the compiled bytecode used for deploying new contracts.
var CplatformBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610aad806100606000396000f3fe6080604052600436106100595760003560e01c806301c70534146100725780634420e486146100f75780634f76dc6114610148578063b9f7945114610199578063dbe55e56146101fe578063e4c5d1371461025557610060565b3661006057005b34600260008282540192505081905550005b34801561007e57600080fd5b506100f56004803603608081101561009557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506102a6565b005b34801561010357600080fd5b506101466004803603602081101561011a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506107cd565b005b34801561015457600080fd5b506101976004803603602081101561016b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108c0565b005b3480156101a557600080fd5b506101e8600480360360208110156101bc57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610910565b6040518082815260200191505060405180910390f35b34801561020a57600080fd5b50610213610959565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561026157600080fd5b506102a46004803603602081101561027857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610982565b005b60008173ffffffffffffffffffffffffffffffffffffffff1663b61e96a56040518163ffffffff1660e01b815260040160206040518083038186803b1580156102ee57600080fd5b505afa158015610302573d6000803e3d6000fd5b505050506040513d602081101561031857600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103af576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252605c815260200180610a1c605c913960600191505060405180910390fd5b600254841115610427576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6f7420656e6f756768207265776172647320746f2061776172642e0000000081525060200191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663269ce69a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561046f57600080fd5b505af1158015610483573d6000803e3d6000fd5b505050506000841061055f578473ffffffffffffffffffffffffffffffffffffffff166108fc859081150290604051600060405180830381858888f193505050501580156104d5573d6000803e3d6000fd5b506104df856108c0565b8173ffffffffffffffffffffffffffffffffffffffff1663c9242986846040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561053257600080fd5b505af1158015610546573d6000803e3d6000fd5b5050505083600260008282540392505081905550610569565b61056885610982565b5b8173ffffffffffffffffffffffffffffffffffffffff1663b21917816040518163ffffffff1660e01b815260040160206040518083038186803b1580156105af57600080fd5b505afa1580156105c3573d6000803e3d6000fd5b505050506040513d60208110156105d957600080fd5b81019080805190602001909291905050508273ffffffffffffffffffffffffffffffffffffffff1663f12064fd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561063057600080fd5b505afa158015610644573d6000803e3d6000fd5b505050506040513d602081101561065a57600080fd5b81019080805190602001909291905050501480156106f9575060008273ffffffffffffffffffffffffffffffffffffffff16630e15561a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156106bb57600080fd5b505afa1580156106cf573d6000803e3d6000fd5b505050506040513d60208110156106e557600080fd5b810190808051906020019092919050505010155b156107c6578073ffffffffffffffffffffffffffffffffffffffff166108fc8373ffffffffffffffffffffffffffffffffffffffff16630e15561a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561075e57600080fd5b505afa158015610772573d6000803e3d6000fd5b505050506040513d602081101561078857600080fd5b81019080805190602001909291905050509081150290604051600060405180830381858888f193505050501580156107c4573d6000803e3d6000fd5b505b5050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610870576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f63616e6e6f74207472616e66657220746f207a65726f2061646472657373000081525060200191505060405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050565b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541115610a185760018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055505b5056fe4f6e6c79207265717565737465722063616e2063616c6c207468697320746f20776f726b6572732077686f207061727469636970616e747320697473207461736b20616e6420686173206e6f74206265656e2072657761726465642ea26469706673582212209177a1a3d1f4d717270194b57ab5c8228da01597b04be89d9bb76d40d198d81864736f6c63430006030033"

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

// Rewarding is a paid mutator transaction binding the contract method 0x01c70534.
//
// Solidity: function Rewarding(address worker, uint256 maskedRewards, uint256 unmaskedRewards, address t) returns()
func (_Cplatform *CplatformTransactor) Rewarding(opts *bind.TransactOpts, worker common.Address, maskedRewards *big.Int, unmaskedRewards *big.Int, t common.Address) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "Rewarding", worker, maskedRewards, unmaskedRewards, t)
}

// Rewarding is a paid mutator transaction binding the contract method 0x01c70534.
//
// Solidity: function Rewarding(address worker, uint256 maskedRewards, uint256 unmaskedRewards, address t) returns()
func (_Cplatform *CplatformSession) Rewarding(worker common.Address, maskedRewards *big.Int, unmaskedRewards *big.Int, t common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.Rewarding(&_Cplatform.TransactOpts, worker, maskedRewards, unmaskedRewards, t)
}

// Rewarding is a paid mutator transaction binding the contract method 0x01c70534.
//
// Solidity: function Rewarding(address worker, uint256 maskedRewards, uint256 unmaskedRewards, address t) returns()
func (_Cplatform *CplatformTransactorSession) Rewarding(worker common.Address, maskedRewards *big.Int, unmaskedRewards *big.Int, t common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.Rewarding(&_Cplatform.TransactOpts, worker, maskedRewards, unmaskedRewards, t)
}

// DecreaseRepuation is a paid mutator transaction binding the contract method 0xe4c5d137.
//
// Solidity: function decreaseRepuation(address _arr) returns()
func (_Cplatform *CplatformTransactor) DecreaseRepuation(opts *bind.TransactOpts, _arr common.Address) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "decreaseRepuation", _arr)
}

// DecreaseRepuation is a paid mutator transaction binding the contract method 0xe4c5d137.
//
// Solidity: function decreaseRepuation(address _arr) returns()
func (_Cplatform *CplatformSession) DecreaseRepuation(_arr common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.DecreaseRepuation(&_Cplatform.TransactOpts, _arr)
}

// DecreaseRepuation is a paid mutator transaction binding the contract method 0xe4c5d137.
//
// Solidity: function decreaseRepuation(address _arr) returns()
func (_Cplatform *CplatformTransactorSession) DecreaseRepuation(_arr common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.DecreaseRepuation(&_Cplatform.TransactOpts, _arr)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Cplatform *CplatformTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Cplatform.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Cplatform *CplatformSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Cplatform.Contract.Fallback(&_Cplatform.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Cplatform *CplatformTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Cplatform.Contract.Fallback(&_Cplatform.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cplatform *CplatformTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cplatform.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cplatform *CplatformSession) Receive() (*types.Transaction, error) {
	return _Cplatform.Contract.Receive(&_Cplatform.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cplatform *CplatformTransactorSession) Receive() (*types.Transaction, error) {
	return _Cplatform.Contract.Receive(&_Cplatform.TransactOpts)
}
