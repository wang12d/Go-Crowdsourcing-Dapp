// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ctask

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

// CtaskABI is the input ABI used to generate the binding from.
const CtaskABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workerRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reputation\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DataSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"TaskPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_worker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_remainingWorkers\",\"type\":\"uint256\"}],\"name\":\"TaskRegistered\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isok\",\"type\":\"bool\"},{\"internalType\":\"contractPlatform\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"Rewarding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SubmitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPlatform\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingWorkers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CtaskBin is the compiled bytecode used for deploying new contracts.
var CtaskBin = "0x60806040523480156200001157600080fd5b506040516200100738038062001007833981810160405260808110156200003757600080fd5b81019080805190602001909291908051906020019092919080519060200190929190805160405193929190846401000000008211156200007657600080fd5b838201915060208201858111156200008d57600080fd5b8251866001820283011164010000000082111715620000ab57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000e1578082015181840152602081019050620000c4565b50505050905090810190601f1680156200010f5780820380516001836020036101000a031916815260200191505b50604052505050836000819055508260018190555080600390805190602001906200013c9291906200019e565b5033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600681905550600060058190555081600881905550505050506200024d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001e157805160ff191683800117855562000212565b8280016001018555821562000212579182015b8281111562000211578251825591602001919060010190620001f4565b5b50905062000221919062000225565b5090565b6200024a91905b80821115620002465760008160009055506001016200022c565b5090565b90565b610daa806200025d6000396000f3fe60806040526004361061004e5760003560e01c8063075d47821461010d57806338b270d614610124578063409ef9e21461014f5780634420e48614610217578063fd0b7b851461026857610055565b3661005557005b3373ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146100fb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180610cc86036913960400191505060405180910390fd5b34600560008282540192505081905550005b34801561011957600080fd5b506101226102e5565b005b34801561013057600080fd5b506101396104ff565b6040518082815260200191505060405180910390f35b34801561015b57600080fd5b506102156004803603602081101561017257600080fd5b810190808035906020019064010000000081111561018f57600080fd5b8201836020820111156101a157600080fd5b803590602001918460018302840111640100000000831117156101c357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610509565b005b34801561022357600080fd5b506102666004803603602081101561023a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106f3565b005b34801561027457600080fd5b506102e36004803603606081101561028b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610968565b005b3373ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461038b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180610c006028913960400191505060405180910390fd5b6005546000546001540211156103ec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526044815260200180610c846044913960600191505060405180910390fd5b600054600554816103f957fe5b04600181905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b6600154600360405180838152602001806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156104ee5780601f106104c3576101008083540402835291602001916104ee565b820191906000526020600020905b8154815290600101906020018083116104d157829003601f168201915b5050935050505060405180910390a2565b6000600654905090565b6000600654600054039050600080905060008090505b828110156105a8573373ffffffffffffffffffffffffffffffffffffffff166002828154811061054b57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561059b57600191506105a8565b808060010191505061051f565b506007839080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906105e4929190610b5a565b508061063b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610d226026913960400191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3846040518080602001828103825283818151815260200191508051906020019080838360005b838110156106b4578082015181840152602081019050610699565b50505050905090810190601f1680156106e15780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b60006006541161074e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610cfe6024913960400191505060405180910390fd5b6008548173ffffffffffffffffffffffffffffffffffffffff1663b9f79451336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156107ce57600080fd5b505afa1580156107e2573d6000803e3d6000fd5b505050506040513d60208110156107f857600080fd5b81019080805190602001909291905050501015610860576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d815260200180610d48602d913960400191505060405180910390fd5b6000339050600660008154809291906001900391905055506002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b305626006546040518082815260200191505060405180910390a35050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252605c815260200180610c28605c913960600191505060405180910390fd5b8115610af9578273ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f19350505050158015610a5c573d6000803e3d6000fd5b508073ffffffffffffffffffffffffffffffffffffffff16634f76dc61846040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b158015610adc57600080fd5b505af1158015610af0573d6000803e3d6000fd5b50505050610b43565b3373ffffffffffffffffffffffffffffffffffffffff166108fc6001549081150290604051600060405180830381858888f19350505050158015610b41573d6000803e3d6000fd5b505b600154600560008282540392505081905550505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b9b57805160ff1916838001178555610bc9565b82800160010185558215610bc9579182015b82811115610bc8578251825591602001919060010190610bad565b5b509050610bd69190610bda565b5090565b610bfc91905b80821115610bf8576000816000905550600101610be0565b5090565b9056fe4e6f7420656e6f75676820616d6f756e7420746f207061792074686520636f6c6c61746572616c734f6e6c79207265717565737465722063616e2063616c6c207468697320746f20776f726b6572732077686f207061727469636970616e747320697473207461736b20616e6420686173206e6f74206265656e2072657761726465642e546865206465706f736974696f6e206d757374206c6172676572207468616e207265776172647320746f20776f726b65727320746f20707562756c697368207461736b2e4f6e6c79207265717565737465722063616e2075706c6f6164207472616e73616374696f6e20746f207468697320636f6e7472616374546865207461736b20646f206e6f74206e65656420776f726b65727320616e796d6f72654f6e6c79207265676973746572656420776f726b65722063616e207375626d697420646174614e6f7420656e6f7567682072657075746174696f6e20746f207061727469636970616e7420746865207461736ba26469706673582212205342c63178be48ad76834af4368fa82aa78fbdd345e236f2547e3e3e6a2b211a64736f6c63430006030033"

// DeployCtask deploys a new Ethereum contract, binding an instance of Ctask to it.
func DeployCtask(auth *bind.TransactOpts, backend bind.ContractBackend, workerRequired *big.Int, rewards *big.Int, reputation *big.Int, description string) (common.Address, *types.Transaction, *Ctask, error) {
	parsed, err := abi.JSON(strings.NewReader(CtaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CtaskBin), backend, workerRequired, rewards, reputation, description)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ctask{CtaskCaller: CtaskCaller{contract: contract}, CtaskTransactor: CtaskTransactor{contract: contract}, CtaskFilterer: CtaskFilterer{contract: contract}}, nil
}

// Ctask is an auto generated Go binding around an Ethereum contract.
type Ctask struct {
	CtaskCaller     // Read-only binding to the contract
	CtaskTransactor // Write-only binding to the contract
	CtaskFilterer   // Log filterer for contract events
}

// CtaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type CtaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CtaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CtaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CtaskSession struct {
	Contract     *Ctask            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CtaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CtaskCallerSession struct {
	Contract *CtaskCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CtaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CtaskTransactorSession struct {
	Contract     *CtaskTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CtaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type CtaskRaw struct {
	Contract *Ctask // Generic contract binding to access the raw methods on
}

// CtaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CtaskCallerRaw struct {
	Contract *CtaskCaller // Generic read-only contract binding to access the raw methods on
}

// CtaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CtaskTransactorRaw struct {
	Contract *CtaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCtask creates a new instance of Ctask, bound to a specific deployed contract.
func NewCtask(address common.Address, backend bind.ContractBackend) (*Ctask, error) {
	contract, err := bindCtask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ctask{CtaskCaller: CtaskCaller{contract: contract}, CtaskTransactor: CtaskTransactor{contract: contract}, CtaskFilterer: CtaskFilterer{contract: contract}}, nil
}

// NewCtaskCaller creates a new read-only instance of Ctask, bound to a specific deployed contract.
func NewCtaskCaller(address common.Address, caller bind.ContractCaller) (*CtaskCaller, error) {
	contract, err := bindCtask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CtaskCaller{contract: contract}, nil
}

// NewCtaskTransactor creates a new write-only instance of Ctask, bound to a specific deployed contract.
func NewCtaskTransactor(address common.Address, transactor bind.ContractTransactor) (*CtaskTransactor, error) {
	contract, err := bindCtask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CtaskTransactor{contract: contract}, nil
}

// NewCtaskFilterer creates a new log filterer instance of Ctask, bound to a specific deployed contract.
func NewCtaskFilterer(address common.Address, filterer bind.ContractFilterer) (*CtaskFilterer, error) {
	contract, err := bindCtask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CtaskFilterer{contract: contract}, nil
}

// bindCtask binds a generic wrapper to an already deployed contract.
func bindCtask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CtaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ctask *CtaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ctask.Contract.CtaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ctask *CtaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctask.Contract.CtaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ctask *CtaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ctask.Contract.CtaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ctask *CtaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ctask.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ctask *CtaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctask.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ctask *CtaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ctask.Contract.contract.Transact(opts, method, params...)
}

// RemainingWorkers is a free data retrieval call binding the contract method 0x38b270d6.
//
// Solidity: function remainingWorkers() view returns(uint256 rem)
func (_Ctask *CtaskCaller) RemainingWorkers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "remainingWorkers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingWorkers is a free data retrieval call binding the contract method 0x38b270d6.
//
// Solidity: function remainingWorkers() view returns(uint256 rem)
func (_Ctask *CtaskSession) RemainingWorkers() (*big.Int, error) {
	return _Ctask.Contract.RemainingWorkers(&_Ctask.CallOpts)
}

// RemainingWorkers is a free data retrieval call binding the contract method 0x38b270d6.
//
// Solidity: function remainingWorkers() view returns(uint256 rem)
func (_Ctask *CtaskCallerSession) RemainingWorkers() (*big.Int, error) {
	return _Ctask.Contract.RemainingWorkers(&_Ctask.CallOpts)
}

// Rewarding is a paid mutator transaction binding the contract method 0xfd0b7b85.
//
// Solidity: function Rewarding(address worker, bool isok, address p) returns()
func (_Ctask *CtaskTransactor) Rewarding(opts *bind.TransactOpts, worker common.Address, isok bool, p common.Address) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "Rewarding", worker, isok, p)
}

// Rewarding is a paid mutator transaction binding the contract method 0xfd0b7b85.
//
// Solidity: function Rewarding(address worker, bool isok, address p) returns()
func (_Ctask *CtaskSession) Rewarding(worker common.Address, isok bool, p common.Address) (*types.Transaction, error) {
	return _Ctask.Contract.Rewarding(&_Ctask.TransactOpts, worker, isok, p)
}

// Rewarding is a paid mutator transaction binding the contract method 0xfd0b7b85.
//
// Solidity: function Rewarding(address worker, bool isok, address p) returns()
func (_Ctask *CtaskTransactorSession) Rewarding(worker common.Address, isok bool, p common.Address) (*types.Transaction, error) {
	return _Ctask.Contract.Rewarding(&_Ctask.TransactOpts, worker, isok, p)
}

// SubmitData is a paid mutator transaction binding the contract method 0x409ef9e2.
//
// Solidity: function SubmitData(bytes data) returns()
func (_Ctask *CtaskTransactor) SubmitData(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "SubmitData", data)
}

// SubmitData is a paid mutator transaction binding the contract method 0x409ef9e2.
//
// Solidity: function SubmitData(bytes data) returns()
func (_Ctask *CtaskSession) SubmitData(data []byte) (*types.Transaction, error) {
	return _Ctask.Contract.SubmitData(&_Ctask.TransactOpts, data)
}

// SubmitData is a paid mutator transaction binding the contract method 0x409ef9e2.
//
// Solidity: function SubmitData(bytes data) returns()
func (_Ctask *CtaskTransactorSession) SubmitData(data []byte) (*types.Transaction, error) {
	return _Ctask.Contract.SubmitData(&_Ctask.TransactOpts, data)
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns()
func (_Ctask *CtaskTransactor) Publish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "publish")
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns()
func (_Ctask *CtaskSession) Publish() (*types.Transaction, error) {
	return _Ctask.Contract.Publish(&_Ctask.TransactOpts)
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns()
func (_Ctask *CtaskTransactorSession) Publish() (*types.Transaction, error) {
	return _Ctask.Contract.Publish(&_Ctask.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address p) returns()
func (_Ctask *CtaskTransactor) Register(opts *bind.TransactOpts, p common.Address) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "register", p)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address p) returns()
func (_Ctask *CtaskSession) Register(p common.Address) (*types.Transaction, error) {
	return _Ctask.Contract.Register(&_Ctask.TransactOpts, p)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address p) returns()
func (_Ctask *CtaskTransactorSession) Register(p common.Address) (*types.Transaction, error) {
	return _Ctask.Contract.Register(&_Ctask.TransactOpts, p)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ctask *CtaskTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Ctask.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ctask *CtaskSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ctask.Contract.Fallback(&_Ctask.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ctask *CtaskTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ctask.Contract.Fallback(&_Ctask.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ctask *CtaskTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctask.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ctask *CtaskSession) Receive() (*types.Transaction, error) {
	return _Ctask.Contract.Receive(&_Ctask.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ctask *CtaskTransactorSession) Receive() (*types.Transaction, error) {
	return _Ctask.Contract.Receive(&_Ctask.TransactOpts)
}

// CtaskDataSubmittedIterator is returned from FilterDataSubmitted and is used to iterate over the raw logs and unpacked data for DataSubmitted events raised by the Ctask contract.
type CtaskDataSubmittedIterator struct {
	Event *CtaskDataSubmitted // Event containing the contract specifics and raw log

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
func (it *CtaskDataSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtaskDataSubmitted)
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
		it.Event = new(CtaskDataSubmitted)
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
func (it *CtaskDataSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtaskDataSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtaskDataSubmitted represents a DataSubmitted event raised by the Ctask contract.
type CtaskDataSubmitted struct {
	From common.Address
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDataSubmitted is a free log retrieval operation binding the contract event 0x5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3.
//
// Solidity: event DataSubmitted(address indexed _from, bytes _data)
func (_Ctask *CtaskFilterer) FilterDataSubmitted(opts *bind.FilterOpts, _from []common.Address) (*CtaskDataSubmittedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Ctask.contract.FilterLogs(opts, "DataSubmitted", _fromRule)
	if err != nil {
		return nil, err
	}
	return &CtaskDataSubmittedIterator{contract: _Ctask.contract, event: "DataSubmitted", logs: logs, sub: sub}, nil
}

// WatchDataSubmitted is a free log subscription operation binding the contract event 0x5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3.
//
// Solidity: event DataSubmitted(address indexed _from, bytes _data)
func (_Ctask *CtaskFilterer) WatchDataSubmitted(opts *bind.WatchOpts, sink chan<- *CtaskDataSubmitted, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Ctask.contract.WatchLogs(opts, "DataSubmitted", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtaskDataSubmitted)
				if err := _Ctask.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
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

// ParseDataSubmitted is a log parse operation binding the contract event 0x5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3.
//
// Solidity: event DataSubmitted(address indexed _from, bytes _data)
func (_Ctask *CtaskFilterer) ParseDataSubmitted(log types.Log) (*CtaskDataSubmitted, error) {
	event := new(CtaskDataSubmitted)
	if err := _Ctask.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CtaskTaskPublishedIterator is returned from FilterTaskPublished and is used to iterate over the raw logs and unpacked data for TaskPublished events raised by the Ctask contract.
type CtaskTaskPublishedIterator struct {
	Event *CtaskTaskPublished // Event containing the contract specifics and raw log

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
func (it *CtaskTaskPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtaskTaskPublished)
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
		it.Event = new(CtaskTaskPublished)
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
func (it *CtaskTaskPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtaskTaskPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtaskTaskPublished represents a TaskPublished event raised by the Ctask contract.
type CtaskTaskPublished struct {
	Rewards     *big.Int
	Requester   common.Address
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskPublished is a free log retrieval operation binding the contract event 0x31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b6.
//
// Solidity: event TaskPublished(uint256 rewards, address indexed _requester, string _description)
func (_Ctask *CtaskFilterer) FilterTaskPublished(opts *bind.FilterOpts, _requester []common.Address) (*CtaskTaskPublishedIterator, error) {

	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Ctask.contract.FilterLogs(opts, "TaskPublished", _requesterRule)
	if err != nil {
		return nil, err
	}
	return &CtaskTaskPublishedIterator{contract: _Ctask.contract, event: "TaskPublished", logs: logs, sub: sub}, nil
}

// WatchTaskPublished is a free log subscription operation binding the contract event 0x31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b6.
//
// Solidity: event TaskPublished(uint256 rewards, address indexed _requester, string _description)
func (_Ctask *CtaskFilterer) WatchTaskPublished(opts *bind.WatchOpts, sink chan<- *CtaskTaskPublished, _requester []common.Address) (event.Subscription, error) {

	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Ctask.contract.WatchLogs(opts, "TaskPublished", _requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtaskTaskPublished)
				if err := _Ctask.contract.UnpackLog(event, "TaskPublished", log); err != nil {
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

// ParseTaskPublished is a log parse operation binding the contract event 0x31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b6.
//
// Solidity: event TaskPublished(uint256 rewards, address indexed _requester, string _description)
func (_Ctask *CtaskFilterer) ParseTaskPublished(log types.Log) (*CtaskTaskPublished, error) {
	event := new(CtaskTaskPublished)
	if err := _Ctask.contract.UnpackLog(event, "TaskPublished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CtaskTaskRegisteredIterator is returned from FilterTaskRegistered and is used to iterate over the raw logs and unpacked data for TaskRegistered events raised by the Ctask contract.
type CtaskTaskRegisteredIterator struct {
	Event *CtaskTaskRegistered // Event containing the contract specifics and raw log

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
func (it *CtaskTaskRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtaskTaskRegistered)
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
		it.Event = new(CtaskTaskRegistered)
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
func (it *CtaskTaskRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtaskTaskRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtaskTaskRegistered represents a TaskRegistered event raised by the Ctask contract.
type CtaskTaskRegistered struct {
	Worker           common.Address
	Requester        common.Address
	RemainingWorkers *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskRegistered is a free log retrieval operation binding the contract event 0x224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b30562.
//
// Solidity: event TaskRegistered(address indexed _worker, address indexed _requester, uint256 _remainingWorkers)
func (_Ctask *CtaskFilterer) FilterTaskRegistered(opts *bind.FilterOpts, _worker []common.Address, _requester []common.Address) (*CtaskTaskRegisteredIterator, error) {

	var _workerRule []interface{}
	for _, _workerItem := range _worker {
		_workerRule = append(_workerRule, _workerItem)
	}
	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Ctask.contract.FilterLogs(opts, "TaskRegistered", _workerRule, _requesterRule)
	if err != nil {
		return nil, err
	}
	return &CtaskTaskRegisteredIterator{contract: _Ctask.contract, event: "TaskRegistered", logs: logs, sub: sub}, nil
}

// WatchTaskRegistered is a free log subscription operation binding the contract event 0x224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b30562.
//
// Solidity: event TaskRegistered(address indexed _worker, address indexed _requester, uint256 _remainingWorkers)
func (_Ctask *CtaskFilterer) WatchTaskRegistered(opts *bind.WatchOpts, sink chan<- *CtaskTaskRegistered, _worker []common.Address, _requester []common.Address) (event.Subscription, error) {

	var _workerRule []interface{}
	for _, _workerItem := range _worker {
		_workerRule = append(_workerRule, _workerItem)
	}
	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Ctask.contract.WatchLogs(opts, "TaskRegistered", _workerRule, _requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtaskTaskRegistered)
				if err := _Ctask.contract.UnpackLog(event, "TaskRegistered", log); err != nil {
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

// ParseTaskRegistered is a log parse operation binding the contract event 0x224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b30562.
//
// Solidity: event TaskRegistered(address indexed _worker, address indexed _requester, uint256 _remainingWorkers)
func (_Ctask *CtaskFilterer) ParseTaskRegistered(log types.Log) (*CtaskTaskRegistered, error) {
	event := new(CtaskTaskRegistered)
	if err := _Ctask.contract.UnpackLog(event, "TaskRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
