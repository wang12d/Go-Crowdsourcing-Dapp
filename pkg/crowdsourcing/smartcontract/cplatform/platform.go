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
const CplatformABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"worker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"task\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WorkerRewarded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maskedRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unmaskedRewards\",\"type\":\"uint256\"},{\"internalType\":\"contractTask\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"Rewarding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"decreaseRepuation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"increaseReputation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"registerRequester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"taskRequired\",\"type\":\"uint256\"}],\"name\":\"registerWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_arr\",\"type\":\"address\"}],\"name\":\"reputation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CplatformBin is the compiled bytecode used for deploying new contracts.
var CplatformBin = "0x608060405267ffffffffffffffff60065534801561001c57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110e78061006c6000396000f3fe6080604052600436106100745760003560e01c80637902b1551161004e5780637902b155146101b4578063b9f794511461020f578063dbe55e5614610274578063e4c5d137146102cb5761007b565b806301c705341461008d5780632bdde680146101125780634f76dc61146101635761007b565b3661007b57005b34600560008282540192505081905550005b34801561009957600080fd5b50610110600480360360808110156100b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061031c565b005b34801561011e57600080fd5b506101616004803603602081101561013557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610cbb565b005b34801561016f57600080fd5b506101b26004803603602081101561018657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610da5565b005b3480156101c057600080fd5b5061020d600480360360408110156101d757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610df5565b005b34801561021b57600080fd5b5061025e6004803603602081101561023257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f24565b6040518082815260200191505060405180910390f35b34801561028057600080fd5b50610289610f6d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102d757600080fd5b5061031a600480360360208110156102ee57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f96565b005b60008173ffffffffffffffffffffffffffffffffffffffff1663b61e96a56040518163ffffffff1660e01b815260040160206040518083038186803b15801561036457600080fd5b505afa158015610378573d6000803e3d6000fd5b505050506040513d602081101561038e57600080fd5b810190808051906020019092919050505090508073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610425576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252605c815260200180611030605c913960600191505060405180910390fd5b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205411156104fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602681526020018061108c6026913960400191505060405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663269ce69a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561054457600080fd5b505af1158015610558573d6000803e3d6000fd5b5050505060065484600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401816105a857fe5b06600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550600083106107205761064b85610da5565b8173ffffffffffffffffffffffffffffffffffffffff1663c9242986846040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b15801561069e57600080fd5b505af11580156106b2573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167f6befaddae276fe4011115007575545073a68b7ae2a6b76fcf07d58f6af4d3054866040518082815260200191505060405180910390a361072a565b61072985610f96565b5b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205414156109c757600554600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541115610867576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f4e6f7420656e6f756768207265776172647320746f2061776172642e0000000081525060200191505060405180910390fd5b600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546005600082825403925050819055508473ffffffffffffffffffffffffffffffffffffffff166108fc600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f1935050505015801561093b573d6000803e3d6000fd5b506000600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b8173ffffffffffffffffffffffffffffffffffffffff1663b21917816040518163ffffffff1660e01b815260040160206040518083038186803b158015610a0d57600080fd5b505afa158015610a21573d6000803e3d6000fd5b505050506040513d6020811015610a3757600080fd5b81019080805190602001909291905050508273ffffffffffffffffffffffffffffffffffffffff1663f12064fd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610a8e57600080fd5b505afa158015610aa2573d6000803e3d6000fd5b505050506040513d6020811015610ab857600080fd5b8101908080519060200190929190505050148015610b57575060008273ffffffffffffffffffffffffffffffffffffffff16630e15561a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610b1957600080fd5b505afa158015610b2d573d6000803e3d6000fd5b505050506040513d6020811015610b4357600080fd5b810190808051906020019092919050505010155b15610cb4578073ffffffffffffffffffffffffffffffffffffffff166108fc8373ffffffffffffffffffffffffffffffffffffffff16630e15561a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610bbc57600080fd5b505afa158015610bd0573d6000803e3d6000fd5b505050506040513d6020811015610be657600080fd5b81019080805190602001909291905050509081150290604051600060405180830381858888f19350505050158015610c22573d6000803e3d6000fd5b508173ffffffffffffffffffffffffffffffffffffffff16630e15561a6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c6957600080fd5b505afa158015610c7d573d6000803e3d6000fd5b505050506040513d6020811015610c9357600080fd5b81019080805190602001909291905050506005600082825403925050819055505b5050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600f8152602001807f696e76616c69642061646472657373000000000000000000000000000000000081525060200191505060405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555050565b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610e98576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f63616e6e6f74207472616e66657220746f207a65726f2061646472657373000081525060200191505060405180910390fd5b60018060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054111561102c5760018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055505b5056fe4f6e6c79207265717565737465722063616e2063616c6c207468697320746f20776f726b6572732077686f207061727469636970616e747320697473207461736b20616e6420686173206e6f74206265656e2072657761726465642e54686520776f726b657220616c726561646c792066696e697368656420697473207461736b73a26469706673582212201bd02e8bf7ee1a6543d64833ad94178b536e0e735ae9d4ae27ac9afa786ad9b564736f6c63430006030033"

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

// RegisterRequester is a paid mutator transaction binding the contract method 0x2bdde680.
//
// Solidity: function registerRequester(address node) returns()
func (_Cplatform *CplatformTransactor) RegisterRequester(opts *bind.TransactOpts, node common.Address) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "registerRequester", node)
}

// RegisterRequester is a paid mutator transaction binding the contract method 0x2bdde680.
//
// Solidity: function registerRequester(address node) returns()
func (_Cplatform *CplatformSession) RegisterRequester(node common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.RegisterRequester(&_Cplatform.TransactOpts, node)
}

// RegisterRequester is a paid mutator transaction binding the contract method 0x2bdde680.
//
// Solidity: function registerRequester(address node) returns()
func (_Cplatform *CplatformTransactorSession) RegisterRequester(node common.Address) (*types.Transaction, error) {
	return _Cplatform.Contract.RegisterRequester(&_Cplatform.TransactOpts, node)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0x7902b155.
//
// Solidity: function registerWorker(address node, uint256 taskRequired) returns()
func (_Cplatform *CplatformTransactor) RegisterWorker(opts *bind.TransactOpts, node common.Address, taskRequired *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "registerWorker", node, taskRequired)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0x7902b155.
//
// Solidity: function registerWorker(address node, uint256 taskRequired) returns()
func (_Cplatform *CplatformSession) RegisterWorker(node common.Address, taskRequired *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.RegisterWorker(&_Cplatform.TransactOpts, node, taskRequired)
}

// RegisterWorker is a paid mutator transaction binding the contract method 0x7902b155.
//
// Solidity: function registerWorker(address node, uint256 taskRequired) returns()
func (_Cplatform *CplatformTransactorSession) RegisterWorker(node common.Address, taskRequired *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.RegisterWorker(&_Cplatform.TransactOpts, node, taskRequired)
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

// CplatformWorkerRewardedIterator is returned from FilterWorkerRewarded and is used to iterate over the raw logs and unpacked data for WorkerRewarded events raised by the Cplatform contract.
type CplatformWorkerRewardedIterator struct {
	Event *CplatformWorkerRewarded // Event containing the contract specifics and raw log

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
func (it *CplatformWorkerRewardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CplatformWorkerRewarded)
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
		it.Event = new(CplatformWorkerRewarded)
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
func (it *CplatformWorkerRewardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CplatformWorkerRewardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CplatformWorkerRewarded represents a WorkerRewarded event raised by the Cplatform contract.
type CplatformWorkerRewarded struct {
	Worker common.Address
	Task   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWorkerRewarded is a free log retrieval operation binding the contract event 0x6befaddae276fe4011115007575545073a68b7ae2a6b76fcf07d58f6af4d3054.
//
// Solidity: event WorkerRewarded(address indexed worker, address indexed task, uint256 amount)
func (_Cplatform *CplatformFilterer) FilterWorkerRewarded(opts *bind.FilterOpts, worker []common.Address, task []common.Address) (*CplatformWorkerRewardedIterator, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var taskRule []interface{}
	for _, taskItem := range task {
		taskRule = append(taskRule, taskItem)
	}

	logs, sub, err := _Cplatform.contract.FilterLogs(opts, "WorkerRewarded", workerRule, taskRule)
	if err != nil {
		return nil, err
	}
	return &CplatformWorkerRewardedIterator{contract: _Cplatform.contract, event: "WorkerRewarded", logs: logs, sub: sub}, nil
}

// WatchWorkerRewarded is a free log subscription operation binding the contract event 0x6befaddae276fe4011115007575545073a68b7ae2a6b76fcf07d58f6af4d3054.
//
// Solidity: event WorkerRewarded(address indexed worker, address indexed task, uint256 amount)
func (_Cplatform *CplatformFilterer) WatchWorkerRewarded(opts *bind.WatchOpts, sink chan<- *CplatformWorkerRewarded, worker []common.Address, task []common.Address) (event.Subscription, error) {

	var workerRule []interface{}
	for _, workerItem := range worker {
		workerRule = append(workerRule, workerItem)
	}
	var taskRule []interface{}
	for _, taskItem := range task {
		taskRule = append(taskRule, taskItem)
	}

	logs, sub, err := _Cplatform.contract.WatchLogs(opts, "WorkerRewarded", workerRule, taskRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CplatformWorkerRewarded)
				if err := _Cplatform.contract.UnpackLog(event, "WorkerRewarded", log); err != nil {
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

// ParseWorkerRewarded is a log parse operation binding the contract event 0x6befaddae276fe4011115007575545073a68b7ae2a6b76fcf07d58f6af4d3054.
//
// Solidity: event WorkerRewarded(address indexed worker, address indexed task, uint256 amount)
func (_Cplatform *CplatformFilterer) ParseWorkerRewarded(log types.Log) (*CplatformWorkerRewarded, error) {
	event := new(CplatformWorkerRewarded)
	if err := _Cplatform.contract.UnpackLog(event, "WorkerRewarded", log); err != nil {
		return nil, err
	}
	return event, nil
}
