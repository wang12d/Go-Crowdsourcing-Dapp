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
const CplatformABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"node\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CplatformBin is the compiled bytecode used for deploying new contracts.
var CplatformBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600681526020017f5469636b657400000000000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f5454000000000000000000000000000000000000000000000000000000000000815250816003908051906020019062000096929190620003e7565b508060049080519060200190620000af929190620003e7565b506012600560006101000a81548160ff021916908360ff1602179055505050620000e060006200015d60201b60201c565b33600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000157600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff166127106200017b60201b60201c565b62000496565b80600560006101000a81548160ff021916908360ff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156200021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b62000233600083836200035960201b60201c565b6200024f816002546200035e60201b620011d51790919060201c565b600281905550620002ad816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546200035e60201b620011d51790919060201c565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b505050565b600080828401905083811015620003dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200042a57805160ff19168380011785556200045b565b828001600101855582156200045b579182015b828111156200045a5782518255916020019190600101906200043d565b5b5090506200046a91906200046e565b5090565b6200049391905b808211156200048f57600081600090555060010162000475565b5090565b90565b61160e80620004a66000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80634420e48611610097578063a457c2d711610066578063a457c2d7146104ac578063a9059cbb14610512578063dbe55e5614610578578063dd62ed3e146105c2576100f5565b80634420e4861461033f57806370a082311461038357806379cc6790146103db57806395d89b4114610429576100f5565b806323b872dd116100d357806323b872dd14610201578063313ce5671461028757806339509351146102ab57806342966c6814610311576100f5565b806306fdde03146100fa578063095ea7b31461017d57806318160ddd146101e3575b600080fd5b61010261063a565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610142578082015181840152602081019050610127565b50505050905090810190601f16801561016f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c96004803603604081101561019357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506106dc565b604051808215151515815260200191505060405180910390f35b6101eb6106fa565b6040518082815260200191505060405180910390f35b61026d6004803603606081101561021757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610704565b604051808215151515815260200191505060405180910390f35b61028f6107dd565b604051808260ff1660ff16815260200191505060405180910390f35b6102f7600480360360408110156102c157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506107f4565b604051808215151515815260200191505060405180910390f35b61033d6004803603602081101561032757600080fd5b81019080803590602001909291905050506108a7565b005b6103816004803603602081101561035557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506108bb565b005b6103c56004803603602081101561039957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061096d565b6040518082815260200191505060405180910390f35b610427600480360360408110156103f157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506109b5565b005b610431610a17565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610471578082015181840152602081019050610456565b50505050905090810190601f16801561049e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104f8600480360360408110156104c257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ab9565b604051808215151515815260200191505060405180910390f35b61055e6004803603604081101561052857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b86565b604051808215151515815260200191505060405180910390f35b610580610ba4565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610624600480360360408110156105d857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bce565b6040518082815260200191505060405180910390f35b606060038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106d25780601f106106a7576101008083540402835291602001916106d2565b820191906000526020600020905b8154815290600101906020018083116106b557829003601f168201915b5050505050905090565b60006106f06106e9610c55565b8484610c5d565b6001905092915050565b6000600254905090565b6000610711848484610e54565b6107d28461071d610c55565b6107cd856040518060600160405280602881526020016114fe60289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610783610c55565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111159092919063ffffffff16565b610c5d565b600190509392505050565b6000600560009054906101000a900460ff16905090565b600061089d610801610c55565b846108988560016000610812610c55565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111d590919063ffffffff16565b610c5d565b6001905092915050565b6108b86108b2610c55565b8261125d565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561095e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f63616e6e6f74207472616e66657220746f207a65726f2061646472657373000081525060200191505060405180910390fd5b610969816001610b86565b5050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006109f482604051806060016040528060248152602001611526602491396109e5866109e0610c55565b610bce565b6111159092919063ffffffff16565b9050610a0883610a02610c55565b83610c5d565b610a12838361125d565b505050565b606060048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aaf5780601f10610a8457610100808354040283529160200191610aaf565b820191906000526020600020905b815481529060010190602001808311610a9257829003601f168201915b5050505050905090565b6000610b7c610ac6610c55565b84610b77856040518060600160405280602581526020016115b46025913960016000610af0610c55565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111159092919063ffffffff16565b610c5d565b6001905092915050565b6000610b9a610b93610c55565b8484610e54565b6001905092915050565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ce3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806115906024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d69576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806114b66022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610eda576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061156b6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610f60576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001806114716023913960400191505060405180910390fd5b610f6b838383611421565b610fd6816040518060600160405280602681526020016114d8602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111159092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611069816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111d590919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b60008383111582906111c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561118757808201518184015260208101905061116c565b50505050905090810190601f1680156111b45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080828401905083811015611253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156112e3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061154a6021913960400191505060405180910390fd5b6112ef82600083611421565b61135a81604051806060016040528060228152602001611494602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546111159092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506113b18160025461142690919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b505050565b600061146883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611115565b90509291505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220a0dff23a09c38bc14a59ec6e852d80ed2e217dd63362f88d65250775ff11dceb64736f6c63430006030033"

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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cplatform *CplatformCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cplatform *CplatformSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cplatform.Contract.Allowance(&_Cplatform.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cplatform *CplatformCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cplatform.Contract.Allowance(&_Cplatform.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cplatform *CplatformCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cplatform *CplatformSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cplatform.Contract.BalanceOf(&_Cplatform.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cplatform *CplatformCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cplatform.Contract.BalanceOf(&_Cplatform.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cplatform *CplatformCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cplatform *CplatformSession) Decimals() (uint8, error) {
	return _Cplatform.Contract.Decimals(&_Cplatform.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cplatform *CplatformCallerSession) Decimals() (uint8, error) {
	return _Cplatform.Contract.Decimals(&_Cplatform.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cplatform *CplatformCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cplatform *CplatformSession) Name() (string, error) {
	return _Cplatform.Contract.Name(&_Cplatform.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cplatform *CplatformCallerSession) Name() (string, error) {
	return _Cplatform.Contract.Name(&_Cplatform.CallOpts)
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

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cplatform *CplatformCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cplatform *CplatformSession) Symbol() (string, error) {
	return _Cplatform.Contract.Symbol(&_Cplatform.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cplatform *CplatformCallerSession) Symbol() (string, error) {
	return _Cplatform.Contract.Symbol(&_Cplatform.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cplatform *CplatformCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cplatform.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cplatform *CplatformSession) TotalSupply() (*big.Int, error) {
	return _Cplatform.Contract.TotalSupply(&_Cplatform.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cplatform *CplatformCallerSession) TotalSupply() (*big.Int, error) {
	return _Cplatform.Contract.TotalSupply(&_Cplatform.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cplatform *CplatformTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cplatform *CplatformSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.Approve(&_Cplatform.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cplatform *CplatformTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.Approve(&_Cplatform.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Cplatform *CplatformTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Cplatform *CplatformSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.Burn(&_Cplatform.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Cplatform *CplatformTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.Burn(&_Cplatform.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Cplatform *CplatformTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Cplatform *CplatformSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.BurnFrom(&_Cplatform.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Cplatform *CplatformTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.BurnFrom(&_Cplatform.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cplatform *CplatformTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cplatform *CplatformSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.DecreaseAllowance(&_Cplatform.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cplatform *CplatformTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.DecreaseAllowance(&_Cplatform.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cplatform *CplatformTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cplatform *CplatformSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.IncreaseAllowance(&_Cplatform.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cplatform *CplatformTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.IncreaseAllowance(&_Cplatform.TransactOpts, spender, addedValue)
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

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cplatform *CplatformTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cplatform *CplatformSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.Transfer(&_Cplatform.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cplatform *CplatformTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.Transfer(&_Cplatform.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cplatform *CplatformTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cplatform *CplatformSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.TransferFrom(&_Cplatform.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cplatform *CplatformTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cplatform.Contract.TransferFrom(&_Cplatform.TransactOpts, sender, recipient, amount)
}

// CplatformApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cplatform contract.
type CplatformApprovalIterator struct {
	Event *CplatformApproval // Event containing the contract specifics and raw log

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
func (it *CplatformApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CplatformApproval)
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
		it.Event = new(CplatformApproval)
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
func (it *CplatformApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CplatformApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CplatformApproval represents a Approval event raised by the Cplatform contract.
type CplatformApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cplatform *CplatformFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CplatformApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cplatform.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CplatformApprovalIterator{contract: _Cplatform.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cplatform *CplatformFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CplatformApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cplatform.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CplatformApproval)
				if err := _Cplatform.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Cplatform *CplatformFilterer) ParseApproval(log types.Log) (*CplatformApproval, error) {
	event := new(CplatformApproval)
	if err := _Cplatform.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CplatformTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cplatform contract.
type CplatformTransferIterator struct {
	Event *CplatformTransfer // Event containing the contract specifics and raw log

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
func (it *CplatformTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CplatformTransfer)
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
		it.Event = new(CplatformTransfer)
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
func (it *CplatformTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CplatformTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CplatformTransfer represents a Transfer event raised by the Cplatform contract.
type CplatformTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cplatform *CplatformFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CplatformTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cplatform.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CplatformTransferIterator{contract: _Cplatform.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cplatform *CplatformFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CplatformTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cplatform.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CplatformTransfer)
				if err := _Cplatform.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Cplatform *CplatformFilterer) ParseTransfer(log types.Log) (*CplatformTransfer, error) {
	event := new(CplatformTransfer)
	if err := _Cplatform.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
