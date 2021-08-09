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
const CtaskABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workerRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DataSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"TaskPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_worker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_remainingWorkers\",\"type\":\"uint256\"}],\"name\":\"TaskRegistered\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SubmitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addingFinishWorkers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finishedWorkers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingWorkers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requester\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"r\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"name\":\"rewarding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"workerRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CtaskBin is the compiled bytecode used for deploying new contracts.
var CtaskBin = "0x608060405234801561001057600080fd5b50604051610d68380380610d688339818101604052606081101561003357600080fd5b8101908080519060200190929190805190602001909291908051604051939291908464010000000082111561006757600080fd5b8382019150602082018581111561007d57600080fd5b825186600182028301116401000000008211171561009a57600080fd5b8083526020830192505050908051906020019080838360005b838110156100ce5780820151818401526020810190506100b3565b50505050905090810190601f1680156100fb5780820380516001836020036101000a031916815260200191505b5060405250505082600081905550816001819055508060039080519060200190610126929190610180565b5033600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600681905550600154600581905550505050610225565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101c157805160ff19168380011785556101ef565b828001600101855582156101ef579182015b828111156101ee5782518255916020019190600101906101d3565b5b5090506101fc9190610200565b5090565b61022291905b8082111561021e576000816000905550600101610206565b5090565b90565b610b34806102346000396000f3fe6080604052600436106100955760003560e01c8063409ef9e211610059578063409ef9e2146101ef578063b2191781146102b7578063b61e96a5146102e2578063c924298614610339578063f12064fd146103745761009c565b8063075d4782146101545780630e15561a1461016b5780631aa3a00814610196578063269ce69a146101ad57806338b270d6146101c45761009c565b3661009c57005b3373ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610142576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526036815260200180610a7f6036913960400191505060405180910390fd5b34600560008282540192505081905550005b34801561016057600080fd5b5061016961039f565b005b34801561017757600080fd5b506101806105aa565b6040518082815260200191505060405180910390f35b3480156101a257600080fd5b506101ab6105b4565b005b3480156101b957600080fd5b506101c2610716565b005b3480156101d057600080fd5b506101d9610729565b6040518082815260200191505060405180910390f35b3480156101fb57600080fd5b506102b56004803603602081101561021257600080fd5b810190808035906020019064010000000081111561022f57600080fd5b82018360208201111561024157600080fd5b8035906020019184600183028401116401000000008311171561026357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610733565b005b3480156102c357600080fd5b506102cc61091d565b6040518082815260200191505060405180910390f35b3480156102ee57600080fd5b506102f7610926565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034557600080fd5b506103726004803603602081101561035c57600080fd5b8101908080359060200190929190505050610950565b005b34801561038057600080fd5b50610389610963565b6040518082815260200191505060405180910390f35b3373ffffffffffffffffffffffffffffffffffffffff16600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610445576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180610a136028913960400191505060405180910390fd5b60055460015411156104a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526044815260200180610a3b6044913960600191505060405180910390fd5b600554600181905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b6600154600360405180838152602001806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156105995780601f1061056e57610100808354040283529160200191610599565b820191906000526020600020905b81548152906001019060200180831161057c57829003601f168201915b5050935050505060405180910390a2565b6000600154905090565b60006006541161060f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610ab56024913960400191505060405180910390fd5b6000339050600660008154809291906001900391905055506002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b305626006546040518082815260200191505060405180910390a350565b6001600760008282540192505081905550565b6000600654905090565b6000600654600054039050600080905060008090505b828110156107d2573373ffffffffffffffffffffffffffffffffffffffff166002828154811061077557fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156107c557600191506107d2565b8080600101915050610749565b5060088390806001815401808255809150506001900390600052602060002001600090919091909150908051906020019061080e92919061096d565b5080610865576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610ad96026913960400191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3846040518080602001828103825283818151815260200191508051906020019080838360005b838110156108de5780820151818401526020810190506108c3565b50505050905090810190601f16801561090b5780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b60008054905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b8060016000828254039250508190555050565b6000600754905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109ae57805160ff19168380011785556109dc565b828001600101855582156109dc579182015b828111156109db5782518255916020019190600101906109c0565b5b5090506109e991906109ed565b5090565b610a0f91905b80821115610a0b5760008160009055506001016109f3565b5090565b9056fe4e6f7420656e6f75676820616d6f756e7420746f207061792074686520636f6c6c61746572616c73546865206465706f736974696f6e206d757374206c6172676572207468616e207265776172647320746f20776f726b65727320746f20707562756c697368207461736b2e4f6e6c79207265717565737465722063616e2075706c6f6164207472616e73616374696f6e20746f207468697320636f6e7472616374546865207461736b20646f206e6f74206e65656420776f726b65727320616e796d6f72654f6e6c79207265676973746572656420776f726b65722063616e207375626d69742064617461a26469706673582212201f7fb89cdb13028c3a2ce80fc73923e27ff6fb4f0df97caf1efae8159dd886db64736f6c63430006030033"

// DeployCtask deploys a new Ethereum contract, binding an instance of Ctask to it.
func DeployCtask(auth *bind.TransactOpts, backend bind.ContractBackend, workerRequired *big.Int, totalRewards *big.Int, description string) (common.Address, *types.Transaction, *Ctask, error) {
	parsed, err := abi.JSON(strings.NewReader(CtaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CtaskBin), backend, workerRequired, totalRewards, description)
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

// FinishedWorkers is a free data retrieval call binding the contract method 0xf12064fd.
//
// Solidity: function finishedWorkers() view returns(uint256)
func (_Ctask *CtaskCaller) FinishedWorkers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "finishedWorkers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FinishedWorkers is a free data retrieval call binding the contract method 0xf12064fd.
//
// Solidity: function finishedWorkers() view returns(uint256)
func (_Ctask *CtaskSession) FinishedWorkers() (*big.Int, error) {
	return _Ctask.Contract.FinishedWorkers(&_Ctask.CallOpts)
}

// FinishedWorkers is a free data retrieval call binding the contract method 0xf12064fd.
//
// Solidity: function finishedWorkers() view returns(uint256)
func (_Ctask *CtaskCallerSession) FinishedWorkers() (*big.Int, error) {
	return _Ctask.Contract.FinishedWorkers(&_Ctask.CallOpts)
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

// Requester is a free data retrieval call binding the contract method 0xb61e96a5.
//
// Solidity: function requester() view returns(address r)
func (_Ctask *CtaskCaller) Requester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "requester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Requester is a free data retrieval call binding the contract method 0xb61e96a5.
//
// Solidity: function requester() view returns(address r)
func (_Ctask *CtaskSession) Requester() (common.Address, error) {
	return _Ctask.Contract.Requester(&_Ctask.CallOpts)
}

// Requester is a free data retrieval call binding the contract method 0xb61e96a5.
//
// Solidity: function requester() view returns(address r)
func (_Ctask *CtaskCallerSession) Requester() (common.Address, error) {
	return _Ctask.Contract.Requester(&_Ctask.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_Ctask *CtaskCaller) TotalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "totalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_Ctask *CtaskSession) TotalRewards() (*big.Int, error) {
	return _Ctask.Contract.TotalRewards(&_Ctask.CallOpts)
}

// TotalRewards is a free data retrieval call binding the contract method 0x0e15561a.
//
// Solidity: function totalRewards() view returns(uint256)
func (_Ctask *CtaskCallerSession) TotalRewards() (*big.Int, error) {
	return _Ctask.Contract.TotalRewards(&_Ctask.CallOpts)
}

// WorkerRequired is a free data retrieval call binding the contract method 0xb2191781.
//
// Solidity: function workerRequired() view returns(uint256)
func (_Ctask *CtaskCaller) WorkerRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "workerRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkerRequired is a free data retrieval call binding the contract method 0xb2191781.
//
// Solidity: function workerRequired() view returns(uint256)
func (_Ctask *CtaskSession) WorkerRequired() (*big.Int, error) {
	return _Ctask.Contract.WorkerRequired(&_Ctask.CallOpts)
}

// WorkerRequired is a free data retrieval call binding the contract method 0xb2191781.
//
// Solidity: function workerRequired() view returns(uint256)
func (_Ctask *CtaskCallerSession) WorkerRequired() (*big.Int, error) {
	return _Ctask.Contract.WorkerRequired(&_Ctask.CallOpts)
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

// AddingFinishWorkers is a paid mutator transaction binding the contract method 0x269ce69a.
//
// Solidity: function addingFinishWorkers() returns()
func (_Ctask *CtaskTransactor) AddingFinishWorkers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "addingFinishWorkers")
}

// AddingFinishWorkers is a paid mutator transaction binding the contract method 0x269ce69a.
//
// Solidity: function addingFinishWorkers() returns()
func (_Ctask *CtaskSession) AddingFinishWorkers() (*types.Transaction, error) {
	return _Ctask.Contract.AddingFinishWorkers(&_Ctask.TransactOpts)
}

// AddingFinishWorkers is a paid mutator transaction binding the contract method 0x269ce69a.
//
// Solidity: function addingFinishWorkers() returns()
func (_Ctask *CtaskTransactorSession) AddingFinishWorkers() (*types.Transaction, error) {
	return _Ctask.Contract.AddingFinishWorkers(&_Ctask.TransactOpts)
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

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Ctask *CtaskTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Ctask *CtaskSession) Register() (*types.Transaction, error) {
	return _Ctask.Contract.Register(&_Ctask.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Ctask *CtaskTransactorSession) Register() (*types.Transaction, error) {
	return _Ctask.Contract.Register(&_Ctask.TransactOpts)
}

// Rewarding is a paid mutator transaction binding the contract method 0xc9242986.
//
// Solidity: function rewarding(uint256 rewards) returns()
func (_Ctask *CtaskTransactor) Rewarding(opts *bind.TransactOpts, rewards *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "rewarding", rewards)
}

// Rewarding is a paid mutator transaction binding the contract method 0xc9242986.
//
// Solidity: function rewarding(uint256 rewards) returns()
func (_Ctask *CtaskSession) Rewarding(rewards *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Rewarding(&_Ctask.TransactOpts, rewards)
}

// Rewarding is a paid mutator transaction binding the contract method 0xc9242986.
//
// Solidity: function rewarding(uint256 rewards) returns()
func (_Ctask *CtaskTransactorSession) Rewarding(rewards *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Rewarding(&_Ctask.TransactOpts, rewards)
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
