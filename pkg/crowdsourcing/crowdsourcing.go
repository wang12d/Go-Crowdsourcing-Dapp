// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crowdsourcing

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

// CrowdsourcingABI is the input ABI used to generate the binding from.
const CrowdsourcingABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"DataSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_task\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"TaskPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"Tranfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"task\",\"type\":\"address\"}],\"name\":\"JoinCrowdsourcingTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"collater\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"workers_needed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"PublishCrowdsourcingTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"task\",\"type\":\"address\"}],\"name\":\"RemainingWorkers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isok\",\"type\":\"bool\"}],\"name\":\"Rewarding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"task\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SubmitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CrowdsourcingBin is the compiled bytecode used for deploying new contracts.
var CrowdsourcingBin = "0x60806040523480156200001157600080fd5b506040518060400160405280601481526020017f5461736b5375626d697373696f6e5469636b65740000000000000000000000008152506040518060400160405280600381526020017f545354000000000000000000000000000000000000000000000000000000000081525081600390805190602001906200009692919062000104565b508060049080519060200190620000af92919062000104565b506012600560006101000a81548160ff021916908360ff1602179055505050620000e06000620000e660201b60201c565b620001b3565b80600560006101000a81548160ff021916908360ff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200014757805160ff191683800117855562000178565b8280016001018555821562000178579182015b82811115620001775782518255916020019190600101906200015a565b5b5090506200018791906200018b565b5090565b620001b091905b80821115620001ac57600081600090555060010162000192565b5090565b90565b612cae80620001c36000396000f3fe60806040526004361061010d5760003560e01c806342966c681161009557806395d89b411161006457806395d89b4114610a2e578063a457c2d714610abe578063a63ac6c214610b31578063a9059cbb14610b8e578063dd62ed3e14610c015761010e565b806342966c681461084b5780636bdeba391461088657806370a082311461096e57806379cc6790146109d35761010e565b806323b872dd116100dc57806323b872dd146105e7578063275492071461067a5780632be319f914610756578063313ce567146107a757806339509351146107d85761010e565b806306fdde0314610454578063095ea7b3146104e4578063099124511461055757806318160ddd146105bc5761010e565b5b60606040518060400160405280600181526020017f01000000000000000000000000000000000000000000000000000000000000008152509050606060405180604001604052806001815260200160008152509050816040516020018082805190602001908083835b6020831061019a5780518252602082019150602081019050602083039250610177565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001206000366040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012014806102b25750806040516020018082805190602001908083835b602083106102445780518252602082019150602081019050602083039250610221565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060003660405160200180838380828437808301925050509250505060405160208183030381529060405280519060200120145b610307576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180612c2c6028913960400191505060405180910390fd5b816040516020018082805190602001908083835b6020831061033e578051825260208201915060208101905060208303925061031b565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001206000366040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012014156104025734600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550610450565b34600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b5050005b34801561046057600080fd5b50610469610c86565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104a957808201518184015260208101905061048e565b50505050905090810190601f1680156104d65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104f057600080fd5b5061053d6004803603604081101561050757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d28565b604051808215151515815260200191505060405180910390f35b34801561056357600080fd5b506105a66004803603602081101561057a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610d46565b6040518082815260200191505060405180910390f35b3480156105c857600080fd5b506105d1610d8f565b6040518082815260200191505060405180910390f35b3480156105f357600080fd5b506106606004803603606081101561060a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d99565b604051808215151515815260200191505060405180910390f35b34801561068657600080fd5b506107546004803603606081101561069d57600080fd5b810190808035906020019092919080359060200190929190803590602001906401000000008111156106ce57600080fd5b8201836020820111156106e057600080fd5b8035906020019184600183028401116401000000008311171561070257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e72565b005b34801561076257600080fd5b506107a56004803603602081101561077957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611108565b005b3480156107b357600080fd5b506107bc611483565b604051808260ff1660ff16815260200191505060405180910390f35b3480156107e457600080fd5b50610831600480360360408110156107fb57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061149a565b604051808215151515815260200191505060405180910390f35b34801561085757600080fd5b506108846004803603602081101561086e57600080fd5b810190808035906020019092919050505061154d565b005b34801561089257600080fd5b5061096c600480360360408110156108a957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001906401000000008111156108e657600080fd5b8201836020820111156108f857600080fd5b8035906020019184600183028401116401000000008311171561091a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611561565b005b34801561097a57600080fd5b506109bd6004803603602081101561099157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506117b0565b6040518082815260200191505060405180910390f35b3480156109df57600080fd5b50610a2c600480360360408110156109f657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506117f8565b005b348015610a3a57600080fd5b50610a4361185a565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610a83578082015181840152602081019050610a68565b50505050905090810190601f168015610ab05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b348015610aca57600080fd5b50610b1760048036036040811015610ae157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506118fc565b604051808215151515815260200191505060405180910390f35b348015610b3d57600080fd5b50610b8c60048036036040811015610b5457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035151590602001909291905050506119c9565b005b348015610b9a57600080fd5b50610be760048036036040811015610bb157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611d82565b604051808215151515815260200191505060405180910390f35b348015610c0d57600080fd5b50610c7060048036036040811015610c2457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611da0565b6040518082815260200191505060405180910390f35b606060038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d1e5780601f10610cf357610100808354040283529160200191610d1e565b820191906000526020600020905b815481529060010190602001808311610d0157829003601f168201915b5050505050905090565b6000610d3c610d35611e27565b8484611e2f565b6001905092915050565b6000600860008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600254905090565b6000610da6848484612026565b610e6784610db2611e27565b610e6285604051806060016040528060288152602001612b7660289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610e18611e27565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122e79092919063ffffffff16565b611e2f565b600190509392505050565b600033905083600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610f0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180612aaa6028913960400191505060405180910390fd5b6000838581610f1a57fe5b049050610f6d83610f686040518060400160405280600a81526020017f0a526577617264733a2000000000000000000000000000000000000000000000815250610f63856123a7565b612516565b612516565b925084600760008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555080600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555083600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061104e82856125de565b8173ffffffffffffffffffffffffffffffffffffffff167f362295b99791fb7f9cc554e61353656ac0f62436a3a559d02ad33dd59a038de6846040518080602001828103825283818151815260200191508051906020019080838360005b838110156110c75780820151818401526020810190506110ac565b50505050905090810190601f1680156110f45780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050505050565b60003390506000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600081116111aa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612a416024913960400191505060405180910390fd5b600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611281576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180612aaa6028913960400191505060405180910390fd5b6001600b60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600660008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540392505081905550600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000815480929190600190039190505550505050565b6000600560009054906101000a900460ff16905090565b60006115436114a7611e27565b8461153e85600160006114b8611e27565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546127a590919063ffffffff16565b611e2f565b6001905092915050565b61155e611558611e27565b8261282d565b50565b60001515600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514611664576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f596f7520616c7265616479207375626d697474656420796f757220646174610081525060200191505060405180910390fd5b6001600b60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff167f5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3826040518080602001828103825283818151815260200191508051906020019080838360005b83811015611772578082015181840152602081019050611757565b50505050905090810190601f16801561179f5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600061183782604051806060016040528060248152602001612b9e6024913961182886611823611e27565b611da0565b6122e79092919063ffffffff16565b905061184b83611845611e27565b83611e2f565b611855838361282d565b505050565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118f25780601f106118c7576101008083540402835291602001916118f2565b820191906000526020600020905b8154815290600101906020018083116118d557829003601f168201915b5050505050905090565b60006119bf611909611e27565b846119ba85604051806060016040528060258152602001612c546025913960016000611933611e27565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122e79092919063ffffffff16565b611e2f565b6001905092915050565b60011515600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514611aaf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252605c815260200180612ad2605c913960600191505060405180910390fd5b8015611b8d578173ffffffffffffffffffffffffffffffffffffffff166108fc600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054600a60008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054019081150290604051600060405180830381858888f19350505050158015611b7b573d6000803e3d6000fd5b50611b87826001611d82565b50611ca4565b8173ffffffffffffffffffffffffffffffffffffffff166108fc600a60008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f19350505050158015611c12573d6000803e3d6000fd5b503373ffffffffffffffffffffffffffffffffffffffff166108fc600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549081150290604051600060405180830381858888f19350505050158015611c98573d6000803e3d6000fd5b50611ca3600161154d565b5b6000600b60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600a60008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b6000611d96611d8f611e27565b8484612026565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611eb5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180612c086024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611f3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180612b2e6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156120ac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180612be36025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612132576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180612a656023913960400191505060405180910390fd5b61213d8383836129f1565b6121a881604051806060016040528060268152602001612b50602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122e79092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061223b816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546127a590919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290612394576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561235957808201518184015260208101905061233e565b50505050905090810190601f1680156123865780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b60606000606490506060816040519080825280601f01601f1916602001820160405280156123e45781602001600182028036833780820191505090505b50905060008090505b6000851461245f576000600a868161240157fe5b069050600a868161240e57fe5b0495508060300160f81b83838060010194508151811061242a57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350506123ed565b6060816040519080825280601f01601f1916602001820160405280156124945781602001600182028036833780820191505090505b50905060008090505b828110156125095783816001850303815181106124b657fe5b602001015160f81c60f81b8282815181106124cd57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808060010191505061249d565b5080945050505050919050565b606082826040516020018083805190602001908083835b60208310612550578051825260208201915060208101905060208303925061252d565b6001836020036101000a03801982511681845116808217855250505050505090500182805190602001908083835b602083106125a1578051825260208201915060208101905060208303925061257e565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b61268d600083836129f1565b6126a2816002546127a590919063ffffffff16565b6002819055506126f9816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546127a590919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600080828401905083811015612823576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156128b3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180612bc26021913960400191505060405180910390fd5b6128bf826000836129f1565b61292a81604051806060016040528060228152602001612a88602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122e79092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612981816002546129f690919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b505050565b6000612a3883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506122e7565b90509291505056fe546865207461736b206e6f206c6f6e676572206e65656473206d6f726520776f726b657245524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654e6f7420656e6f75676820616d6f756e7420746f207061792074686520636f6c6c61746572616c734f6e6c79207265717565737465722063616e2063616c6c207468697320746f20776f726b6572732077686f207061727469636970616e747320697473207461736b20616e6420686173206e6f74206265656e2072657761726465642e45524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373596f752073686f756c642065697468657220626520776f726b6572206f722072657175657374657245524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212200f57c56680a057b854aa2a90c9f6e94d6a10d547ed8922086d06332019e21b4864736f6c63430006030033"

// DeployCrowdsourcing deploys a new Ethereum contract, binding an instance of Crowdsourcing to it.
func DeployCrowdsourcing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Crowdsourcing, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsourcingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrowdsourcingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Crowdsourcing{CrowdsourcingCaller: CrowdsourcingCaller{contract: contract}, CrowdsourcingTransactor: CrowdsourcingTransactor{contract: contract}, CrowdsourcingFilterer: CrowdsourcingFilterer{contract: contract}}, nil
}

// Crowdsourcing is an auto generated Go binding around an Ethereum contract.
type Crowdsourcing struct {
	CrowdsourcingCaller     // Read-only binding to the contract
	CrowdsourcingTransactor // Write-only binding to the contract
	CrowdsourcingFilterer   // Log filterer for contract events
}

// CrowdsourcingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdsourcingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsourcingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdsourcingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsourcingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdsourcingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsourcingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdsourcingSession struct {
	Contract     *Crowdsourcing    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdsourcingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdsourcingCallerSession struct {
	Contract *CrowdsourcingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CrowdsourcingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdsourcingTransactorSession struct {
	Contract     *CrowdsourcingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CrowdsourcingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdsourcingRaw struct {
	Contract *Crowdsourcing // Generic contract binding to access the raw methods on
}

// CrowdsourcingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdsourcingCallerRaw struct {
	Contract *CrowdsourcingCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdsourcingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdsourcingTransactorRaw struct {
	Contract *CrowdsourcingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdsourcing creates a new instance of Crowdsourcing, bound to a specific deployed contract.
func NewCrowdsourcing(address common.Address, backend bind.ContractBackend) (*Crowdsourcing, error) {
	contract, err := bindCrowdsourcing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crowdsourcing{CrowdsourcingCaller: CrowdsourcingCaller{contract: contract}, CrowdsourcingTransactor: CrowdsourcingTransactor{contract: contract}, CrowdsourcingFilterer: CrowdsourcingFilterer{contract: contract}}, nil
}

// NewCrowdsourcingCaller creates a new read-only instance of Crowdsourcing, bound to a specific deployed contract.
func NewCrowdsourcingCaller(address common.Address, caller bind.ContractCaller) (*CrowdsourcingCaller, error) {
	contract, err := bindCrowdsourcing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingCaller{contract: contract}, nil
}

// NewCrowdsourcingTransactor creates a new write-only instance of Crowdsourcing, bound to a specific deployed contract.
func NewCrowdsourcingTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdsourcingTransactor, error) {
	contract, err := bindCrowdsourcing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingTransactor{contract: contract}, nil
}

// NewCrowdsourcingFilterer creates a new log filterer instance of Crowdsourcing, bound to a specific deployed contract.
func NewCrowdsourcingFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdsourcingFilterer, error) {
	contract, err := bindCrowdsourcing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingFilterer{contract: contract}, nil
}

// bindCrowdsourcing binds a generic wrapper to an already deployed contract.
func bindCrowdsourcing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsourcingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsourcing *CrowdsourcingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crowdsourcing.Contract.CrowdsourcingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsourcing *CrowdsourcingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.CrowdsourcingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsourcing *CrowdsourcingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.CrowdsourcingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsourcing *CrowdsourcingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crowdsourcing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsourcing *CrowdsourcingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsourcing *CrowdsourcingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.contract.Transact(opts, method, params...)
}

// RemainingWorkers is a free data retrieval call binding the contract method 0x09912451.
//
// Solidity: function RemainingWorkers(address task) view returns(uint256 rem)
func (_Crowdsourcing *CrowdsourcingCaller) RemainingWorkers(opts *bind.CallOpts, task common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crowdsourcing.contract.Call(opts, &out, "RemainingWorkers", task)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingWorkers is a free data retrieval call binding the contract method 0x09912451.
//
// Solidity: function RemainingWorkers(address task) view returns(uint256 rem)
func (_Crowdsourcing *CrowdsourcingSession) RemainingWorkers(task common.Address) (*big.Int, error) {
	return _Crowdsourcing.Contract.RemainingWorkers(&_Crowdsourcing.CallOpts, task)
}

// RemainingWorkers is a free data retrieval call binding the contract method 0x09912451.
//
// Solidity: function RemainingWorkers(address task) view returns(uint256 rem)
func (_Crowdsourcing *CrowdsourcingCallerSession) RemainingWorkers(task common.Address) (*big.Int, error) {
	return _Crowdsourcing.Contract.RemainingWorkers(&_Crowdsourcing.CallOpts, task)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Crowdsourcing *CrowdsourcingCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crowdsourcing.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Crowdsourcing *CrowdsourcingSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Crowdsourcing.Contract.Allowance(&_Crowdsourcing.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Crowdsourcing *CrowdsourcingCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Crowdsourcing.Contract.Allowance(&_Crowdsourcing.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Crowdsourcing *CrowdsourcingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Crowdsourcing.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Crowdsourcing *CrowdsourcingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Crowdsourcing.Contract.BalanceOf(&_Crowdsourcing.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Crowdsourcing *CrowdsourcingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Crowdsourcing.Contract.BalanceOf(&_Crowdsourcing.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Crowdsourcing *CrowdsourcingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Crowdsourcing.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Crowdsourcing *CrowdsourcingSession) Decimals() (uint8, error) {
	return _Crowdsourcing.Contract.Decimals(&_Crowdsourcing.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Crowdsourcing *CrowdsourcingCallerSession) Decimals() (uint8, error) {
	return _Crowdsourcing.Contract.Decimals(&_Crowdsourcing.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crowdsourcing *CrowdsourcingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crowdsourcing.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crowdsourcing *CrowdsourcingSession) Name() (string, error) {
	return _Crowdsourcing.Contract.Name(&_Crowdsourcing.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Crowdsourcing *CrowdsourcingCallerSession) Name() (string, error) {
	return _Crowdsourcing.Contract.Name(&_Crowdsourcing.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crowdsourcing *CrowdsourcingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Crowdsourcing.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crowdsourcing *CrowdsourcingSession) Symbol() (string, error) {
	return _Crowdsourcing.Contract.Symbol(&_Crowdsourcing.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Crowdsourcing *CrowdsourcingCallerSession) Symbol() (string, error) {
	return _Crowdsourcing.Contract.Symbol(&_Crowdsourcing.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crowdsourcing *CrowdsourcingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crowdsourcing.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crowdsourcing *CrowdsourcingSession) TotalSupply() (*big.Int, error) {
	return _Crowdsourcing.Contract.TotalSupply(&_Crowdsourcing.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Crowdsourcing *CrowdsourcingCallerSession) TotalSupply() (*big.Int, error) {
	return _Crowdsourcing.Contract.TotalSupply(&_Crowdsourcing.CallOpts)
}

// JoinCrowdsourcingTask is a paid mutator transaction binding the contract method 0x2be319f9.
//
// Solidity: function JoinCrowdsourcingTask(address task) returns()
func (_Crowdsourcing *CrowdsourcingTransactor) JoinCrowdsourcingTask(opts *bind.TransactOpts, task common.Address) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "JoinCrowdsourcingTask", task)
}

// JoinCrowdsourcingTask is a paid mutator transaction binding the contract method 0x2be319f9.
//
// Solidity: function JoinCrowdsourcingTask(address task) returns()
func (_Crowdsourcing *CrowdsourcingSession) JoinCrowdsourcingTask(task common.Address) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.JoinCrowdsourcingTask(&_Crowdsourcing.TransactOpts, task)
}

// JoinCrowdsourcingTask is a paid mutator transaction binding the contract method 0x2be319f9.
//
// Solidity: function JoinCrowdsourcingTask(address task) returns()
func (_Crowdsourcing *CrowdsourcingTransactorSession) JoinCrowdsourcingTask(task common.Address) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.JoinCrowdsourcingTask(&_Crowdsourcing.TransactOpts, task)
}

// PublishCrowdsourcingTask is a paid mutator transaction binding the contract method 0x27549207.
//
// Solidity: function PublishCrowdsourcingTask(uint256 collater, uint256 workers_needed, string description) returns()
func (_Crowdsourcing *CrowdsourcingTransactor) PublishCrowdsourcingTask(opts *bind.TransactOpts, collater *big.Int, workers_needed *big.Int, description string) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "PublishCrowdsourcingTask", collater, workers_needed, description)
}

// PublishCrowdsourcingTask is a paid mutator transaction binding the contract method 0x27549207.
//
// Solidity: function PublishCrowdsourcingTask(uint256 collater, uint256 workers_needed, string description) returns()
func (_Crowdsourcing *CrowdsourcingSession) PublishCrowdsourcingTask(collater *big.Int, workers_needed *big.Int, description string) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.PublishCrowdsourcingTask(&_Crowdsourcing.TransactOpts, collater, workers_needed, description)
}

// PublishCrowdsourcingTask is a paid mutator transaction binding the contract method 0x27549207.
//
// Solidity: function PublishCrowdsourcingTask(uint256 collater, uint256 workers_needed, string description) returns()
func (_Crowdsourcing *CrowdsourcingTransactorSession) PublishCrowdsourcingTask(collater *big.Int, workers_needed *big.Int, description string) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.PublishCrowdsourcingTask(&_Crowdsourcing.TransactOpts, collater, workers_needed, description)
}

// Rewarding is a paid mutator transaction binding the contract method 0xa63ac6c2.
//
// Solidity: function Rewarding(address worker, bool isok) returns()
func (_Crowdsourcing *CrowdsourcingTransactor) Rewarding(opts *bind.TransactOpts, worker common.Address, isok bool) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "Rewarding", worker, isok)
}

// Rewarding is a paid mutator transaction binding the contract method 0xa63ac6c2.
//
// Solidity: function Rewarding(address worker, bool isok) returns()
func (_Crowdsourcing *CrowdsourcingSession) Rewarding(worker common.Address, isok bool) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Rewarding(&_Crowdsourcing.TransactOpts, worker, isok)
}

// Rewarding is a paid mutator transaction binding the contract method 0xa63ac6c2.
//
// Solidity: function Rewarding(address worker, bool isok) returns()
func (_Crowdsourcing *CrowdsourcingTransactorSession) Rewarding(worker common.Address, isok bool) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Rewarding(&_Crowdsourcing.TransactOpts, worker, isok)
}

// SubmitData is a paid mutator transaction binding the contract method 0x6bdeba39.
//
// Solidity: function SubmitData(address task, bytes data) returns()
func (_Crowdsourcing *CrowdsourcingTransactor) SubmitData(opts *bind.TransactOpts, task common.Address, data []byte) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "SubmitData", task, data)
}

// SubmitData is a paid mutator transaction binding the contract method 0x6bdeba39.
//
// Solidity: function SubmitData(address task, bytes data) returns()
func (_Crowdsourcing *CrowdsourcingSession) SubmitData(task common.Address, data []byte) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.SubmitData(&_Crowdsourcing.TransactOpts, task, data)
}

// SubmitData is a paid mutator transaction binding the contract method 0x6bdeba39.
//
// Solidity: function SubmitData(address task, bytes data) returns()
func (_Crowdsourcing *CrowdsourcingTransactorSession) SubmitData(task common.Address, data []byte) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.SubmitData(&_Crowdsourcing.TransactOpts, task, data)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Approve(&_Crowdsourcing.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Approve(&_Crowdsourcing.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Crowdsourcing *CrowdsourcingTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Crowdsourcing *CrowdsourcingSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Burn(&_Crowdsourcing.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Crowdsourcing *CrowdsourcingTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Burn(&_Crowdsourcing.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Crowdsourcing *CrowdsourcingTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Crowdsourcing *CrowdsourcingSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.BurnFrom(&_Crowdsourcing.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Crowdsourcing *CrowdsourcingTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.BurnFrom(&_Crowdsourcing.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Crowdsourcing *CrowdsourcingSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.DecreaseAllowance(&_Crowdsourcing.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.DecreaseAllowance(&_Crowdsourcing.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Crowdsourcing *CrowdsourcingSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.IncreaseAllowance(&_Crowdsourcing.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.IncreaseAllowance(&_Crowdsourcing.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Transfer(&_Crowdsourcing.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Transfer(&_Crowdsourcing.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.TransferFrom(&_Crowdsourcing.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Crowdsourcing *CrowdsourcingTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.TransferFrom(&_Crowdsourcing.TransactOpts, sender, recipient, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Crowdsourcing *CrowdsourcingTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Crowdsourcing.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Crowdsourcing *CrowdsourcingSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Fallback(&_Crowdsourcing.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Crowdsourcing *CrowdsourcingTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Crowdsourcing.Contract.Fallback(&_Crowdsourcing.TransactOpts, calldata)
}

// CrowdsourcingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Crowdsourcing contract.
type CrowdsourcingApprovalIterator struct {
	Event *CrowdsourcingApproval // Event containing the contract specifics and raw log

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
func (it *CrowdsourcingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsourcingApproval)
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
		it.Event = new(CrowdsourcingApproval)
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
func (it *CrowdsourcingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsourcingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsourcingApproval represents a Approval event raised by the Crowdsourcing contract.
type CrowdsourcingApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Crowdsourcing *CrowdsourcingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrowdsourcingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Crowdsourcing.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingApprovalIterator{contract: _Crowdsourcing.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Crowdsourcing *CrowdsourcingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrowdsourcingApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Crowdsourcing.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsourcingApproval)
				if err := _Crowdsourcing.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Crowdsourcing *CrowdsourcingFilterer) ParseApproval(log types.Log) (*CrowdsourcingApproval, error) {
	event := new(CrowdsourcingApproval)
	if err := _Crowdsourcing.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrowdsourcingDataSubmittedIterator is returned from FilterDataSubmitted and is used to iterate over the raw logs and unpacked data for DataSubmitted events raised by the Crowdsourcing contract.
type CrowdsourcingDataSubmittedIterator struct {
	Event *CrowdsourcingDataSubmitted // Event containing the contract specifics and raw log

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
func (it *CrowdsourcingDataSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsourcingDataSubmitted)
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
		it.Event = new(CrowdsourcingDataSubmitted)
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
func (it *CrowdsourcingDataSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsourcingDataSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsourcingDataSubmitted represents a DataSubmitted event raised by the Crowdsourcing contract.
type CrowdsourcingDataSubmitted struct {
	From common.Address
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDataSubmitted is a free log retrieval operation binding the contract event 0x5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3.
//
// Solidity: event DataSubmitted(address indexed _from, bytes data)
func (_Crowdsourcing *CrowdsourcingFilterer) FilterDataSubmitted(opts *bind.FilterOpts, _from []common.Address) (*CrowdsourcingDataSubmittedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Crowdsourcing.contract.FilterLogs(opts, "DataSubmitted", _fromRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingDataSubmittedIterator{contract: _Crowdsourcing.contract, event: "DataSubmitted", logs: logs, sub: sub}, nil
}

// WatchDataSubmitted is a free log subscription operation binding the contract event 0x5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3.
//
// Solidity: event DataSubmitted(address indexed _from, bytes data)
func (_Crowdsourcing *CrowdsourcingFilterer) WatchDataSubmitted(opts *bind.WatchOpts, sink chan<- *CrowdsourcingDataSubmitted, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Crowdsourcing.contract.WatchLogs(opts, "DataSubmitted", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsourcingDataSubmitted)
				if err := _Crowdsourcing.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
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
// Solidity: event DataSubmitted(address indexed _from, bytes data)
func (_Crowdsourcing *CrowdsourcingFilterer) ParseDataSubmitted(log types.Log) (*CrowdsourcingDataSubmitted, error) {
	event := new(CrowdsourcingDataSubmitted)
	if err := _Crowdsourcing.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrowdsourcingTaskPublishedIterator is returned from FilterTaskPublished and is used to iterate over the raw logs and unpacked data for TaskPublished events raised by the Crowdsourcing contract.
type CrowdsourcingTaskPublishedIterator struct {
	Event *CrowdsourcingTaskPublished // Event containing the contract specifics and raw log

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
func (it *CrowdsourcingTaskPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsourcingTaskPublished)
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
		it.Event = new(CrowdsourcingTaskPublished)
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
func (it *CrowdsourcingTaskPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsourcingTaskPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsourcingTaskPublished represents a TaskPublished event raised by the Crowdsourcing contract.
type CrowdsourcingTaskPublished struct {
	Task        common.Address
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskPublished is a free log retrieval operation binding the contract event 0x362295b99791fb7f9cc554e61353656ac0f62436a3a559d02ad33dd59a038de6.
//
// Solidity: event TaskPublished(address indexed _task, string description)
func (_Crowdsourcing *CrowdsourcingFilterer) FilterTaskPublished(opts *bind.FilterOpts, _task []common.Address) (*CrowdsourcingTaskPublishedIterator, error) {

	var _taskRule []interface{}
	for _, _taskItem := range _task {
		_taskRule = append(_taskRule, _taskItem)
	}

	logs, sub, err := _Crowdsourcing.contract.FilterLogs(opts, "TaskPublished", _taskRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingTaskPublishedIterator{contract: _Crowdsourcing.contract, event: "TaskPublished", logs: logs, sub: sub}, nil
}

// WatchTaskPublished is a free log subscription operation binding the contract event 0x362295b99791fb7f9cc554e61353656ac0f62436a3a559d02ad33dd59a038de6.
//
// Solidity: event TaskPublished(address indexed _task, string description)
func (_Crowdsourcing *CrowdsourcingFilterer) WatchTaskPublished(opts *bind.WatchOpts, sink chan<- *CrowdsourcingTaskPublished, _task []common.Address) (event.Subscription, error) {

	var _taskRule []interface{}
	for _, _taskItem := range _task {
		_taskRule = append(_taskRule, _taskItem)
	}

	logs, sub, err := _Crowdsourcing.contract.WatchLogs(opts, "TaskPublished", _taskRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsourcingTaskPublished)
				if err := _Crowdsourcing.contract.UnpackLog(event, "TaskPublished", log); err != nil {
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

// ParseTaskPublished is a log parse operation binding the contract event 0x362295b99791fb7f9cc554e61353656ac0f62436a3a559d02ad33dd59a038de6.
//
// Solidity: event TaskPublished(address indexed _task, string description)
func (_Crowdsourcing *CrowdsourcingFilterer) ParseTaskPublished(log types.Log) (*CrowdsourcingTaskPublished, error) {
	event := new(CrowdsourcingTaskPublished)
	if err := _Crowdsourcing.contract.UnpackLog(event, "TaskPublished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrowdsourcingTranferIterator is returned from FilterTranfer and is used to iterate over the raw logs and unpacked data for Tranfer events raised by the Crowdsourcing contract.
type CrowdsourcingTranferIterator struct {
	Event *CrowdsourcingTranfer // Event containing the contract specifics and raw log

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
func (it *CrowdsourcingTranferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsourcingTranfer)
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
		it.Event = new(CrowdsourcingTranfer)
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
func (it *CrowdsourcingTranferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsourcingTranferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsourcingTranfer represents a Tranfer event raised by the Crowdsourcing contract.
type CrowdsourcingTranfer struct {
	From common.Address
	To   common.Address
	Val  *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTranfer is a free log retrieval operation binding the contract event 0x5225eac2a7facfbeb099c00a4cc7c457701324f1fd84b84b14033f9e911374a4.
//
// Solidity: event Tranfer(address indexed _from, address indexed _to, uint256 _val)
func (_Crowdsourcing *CrowdsourcingFilterer) FilterTranfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*CrowdsourcingTranferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Crowdsourcing.contract.FilterLogs(opts, "Tranfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingTranferIterator{contract: _Crowdsourcing.contract, event: "Tranfer", logs: logs, sub: sub}, nil
}

// WatchTranfer is a free log subscription operation binding the contract event 0x5225eac2a7facfbeb099c00a4cc7c457701324f1fd84b84b14033f9e911374a4.
//
// Solidity: event Tranfer(address indexed _from, address indexed _to, uint256 _val)
func (_Crowdsourcing *CrowdsourcingFilterer) WatchTranfer(opts *bind.WatchOpts, sink chan<- *CrowdsourcingTranfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Crowdsourcing.contract.WatchLogs(opts, "Tranfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsourcingTranfer)
				if err := _Crowdsourcing.contract.UnpackLog(event, "Tranfer", log); err != nil {
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

// ParseTranfer is a log parse operation binding the contract event 0x5225eac2a7facfbeb099c00a4cc7c457701324f1fd84b84b14033f9e911374a4.
//
// Solidity: event Tranfer(address indexed _from, address indexed _to, uint256 _val)
func (_Crowdsourcing *CrowdsourcingFilterer) ParseTranfer(log types.Log) (*CrowdsourcingTranfer, error) {
	event := new(CrowdsourcingTranfer)
	if err := _Crowdsourcing.contract.UnpackLog(event, "Tranfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CrowdsourcingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Crowdsourcing contract.
type CrowdsourcingTransferIterator struct {
	Event *CrowdsourcingTransfer // Event containing the contract specifics and raw log

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
func (it *CrowdsourcingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsourcingTransfer)
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
		it.Event = new(CrowdsourcingTransfer)
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
func (it *CrowdsourcingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsourcingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsourcingTransfer represents a Transfer event raised by the Crowdsourcing contract.
type CrowdsourcingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Crowdsourcing *CrowdsourcingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrowdsourcingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Crowdsourcing.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsourcingTransferIterator{contract: _Crowdsourcing.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Crowdsourcing *CrowdsourcingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrowdsourcingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Crowdsourcing.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsourcingTransfer)
				if err := _Crowdsourcing.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Crowdsourcing *CrowdsourcingFilterer) ParseTransfer(log types.Log) (*CrowdsourcingTransfer, error) {
	event := new(CrowdsourcingTransfer)
	if err := _Crowdsourcing.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
