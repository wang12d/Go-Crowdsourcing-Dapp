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
const CtaskABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workerRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DataSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"TaskPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_worker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_remainingWorkers\",\"type\":\"uint256\"}],\"name\":\"TaskRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isok\",\"type\":\"bool\"},{\"internalType\":\"contractPlatform\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"Rewarding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SubmitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPlatform\",\"name\":\"p\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingWorkers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CtaskBin is the compiled bytecode used for deploying new contracts.
var CtaskBin = "0x60806040523480156200001157600080fd5b506040516200268f3803806200268f833981810160405260608110156200003757600080fd5b810190808051906020019092919080519060200190929190805160405193929190846401000000008211156200006c57600080fd5b838201915060208201858111156200008357600080fd5b8251866001820283011164010000000082111715620000a157600080fd5b8083526020830192505050908051906020019080838360005b83811015620000d7578082015181840152602081019050620000ba565b50505050905090810190601f168015620001055780820380516001836020036101000a031916815260200191505b506040525050506040518060400160405280601481526020017f5461736b5375626d697373696f6e5469636b65740000000000000000000000008152506040518060400160405280600381526020017f545354000000000000000000000000000000000000000000000000000000000081525081600390805190602001906200019092919062000278565b508060049080519060200190620001a992919062000278565b506012600560006101000a81548160ff021916908360ff160217905550505082600681905550816007819055508060099080519060200190620001ee92919062000278565b5033600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600c819055506000600b819055506200025160006200025a60201b60201c565b50505062000327565b80600560006101000a81548160ff021916908360ff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002bb57805160ff1916838001178555620002ec565b82800160010185558215620002ec579182015b82811115620002eb578251825591602001919060010190620002ce565b5b509050620002fb9190620002ff565b5090565b6200032491905b808211156200032057600081600090555060010162000306565b5090565b90565b61235880620003376000396000f3fe60806040526004361061010d5760003560e01c806342966c681161009557806395d89b411161006457806395d89b4114610687578063a457c2d714610717578063a9059cbb1461078a578063dd62ed3e146107fd578063fd0b7b851461088257610114565b806342966c681461053b5780634420e4861461057657806370a08231146105c757806379cc67901461062c57610114565b806323b872dd116100dc57806323b872dd14610311578063313ce567146103a457806338b270d6146103d55780633950935114610400578063409ef9e21461047357610114565b806306fdde03146101cc578063075d47821461025c578063095ea7b31461027357806318160ddd146102e657610114565b3661011457005b3373ffffffffffffffffffffffffffffffffffffffff16600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806121a26036913960400191505060405180910390fd5b34600b60008282540192505081905550005b3480156101d857600080fd5b506101e16108ff565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610221578082015181840152602081019050610206565b50505050905090810190601f16801561024e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026857600080fd5b506102716109a1565b005b34801561027f57600080fd5b506102cc6004803603604081101561029657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610bbb565b604051808215151515815260200191505060405180910390f35b3480156102f257600080fd5b506102fb610bd9565b6040518082815260200191505060405180910390f35b34801561031d57600080fd5b5061038a6004803603606081101561033457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610be3565b604051808215151515815260200191505060405180910390f35b3480156103b057600080fd5b506103b9610cbc565b604051808260ff1660ff16815260200191505060405180910390f35b3480156103e157600080fd5b506103ea610cd3565b6040518082815260200191505060405180910390f35b34801561040c57600080fd5b506104596004803603604081101561042357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610cdd565b604051808215151515815260200191505060405180910390f35b34801561047f57600080fd5b506105396004803603602081101561049657600080fd5b81019080803590602001906401000000008111156104b357600080fd5b8201836020820111156104c557600080fd5b803590602001918460018302840111640100000000831117156104e757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610d90565b005b34801561054757600080fd5b506105746004803603602081101561055e57600080fd5b8101908080359060200190929190505050610f7a565b005b34801561058257600080fd5b506105c56004803603602081101561059957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f8e565b005b3480156105d357600080fd5b50610616600480360360208110156105ea57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611201565b6040518082815260200191505060405180910390f35b34801561063857600080fd5b506106856004803603604081101561064f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611249565b005b34801561069357600080fd5b5061069c6112ab565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106dc5780820151818401526020810190506106c1565b50505050905090810190601f1680156107095780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561072357600080fd5b506107706004803603604081101561073a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061134d565b604051808215151515815260200191505060405180910390f35b34801561079657600080fd5b506107e3600480360360408110156107ad57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061141a565b604051808215151515815260200191505060405180910390f35b34801561080957600080fd5b5061086c6004803603604081101561082057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611438565b6040518082815260200191505060405180910390f35b34801561088e57600080fd5b506108fd600480360360608110156108a557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803515159060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114bf565b005b606060038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109975780601f1061096c57610100808354040283529160200191610997565b820191906000526020600020905b81548152906001019060200180831161097a57829003601f168201915b5050505050905090565b3373ffffffffffffffffffffffffffffffffffffffff16600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a47576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806120846028913960400191505060405180910390fd5b600b54600654600754021115610aa8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604481526020018061212a6044913960600191505060405180910390fd5b600654600b5481610ab557fe5b04600781905550600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b660075460096040518083815260200180602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610baa5780601f10610b7f57610100808354040283529160200191610baa565b820191906000526020600020905b815481529060010190602001808311610b8d57829003601f168201915b5050935050505060405180910390a2565b6000610bcf610bc861177e565b8484611786565b6001905092915050565b6000600254905090565b6000610bf084848461197d565b610cb184610bfc61177e565b610cac8560405180606001604052806028815260200161222260289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610c6261177e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c3e9092919063ffffffff16565b611786565b600190509392505050565b6000600560009054906101000a900460ff16905090565b6000600c54905090565b6000610d86610cea61177e565b84610d818560016000610cfb61177e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cfe90919063ffffffff16565b611786565b6001905092915050565b6000600c54600654039050600080905060008090505b82811015610e2f573373ffffffffffffffffffffffffffffffffffffffff1660088281548110610dd257fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610e225760019150610e2f565b8080600101915050610da6565b50600d83908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610e6b929190611f99565b5080610ec2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806122d86026913960400191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3846040518080602001828103825283818151815260200191508051906020019080838360005b83811015610f3b578082015181840152602081019050610f20565b50505050905090810190601f168015610f685780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b610f8b610f8561177e565b82611d86565b50565b6000600c5411610fe9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806121d86024913960400191505060405180910390fd5b600033905060008273ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561106d57600080fd5b505afa158015611081573d6000803e3d6000fd5b505050506040513d602081101561109757600080fd5b8101908080519060200190929190505050116110fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603481526020018061216e6034913960400191505060405180910390fd5b600c60008154809291906001900391905055506008819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b30562600c546040518082815260200191505060405180910390a35050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60006112888260405180606001604052806024815260200161224a602491396112798661127461177e565b611438565b611c3e9092919063ffffffff16565b905061129c8361129661177e565b83611786565b6112a68383611d86565b505050565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113435780601f1061131857610100808354040283529160200191611343565b820191906000526020600020905b81548152906001019060200180831161132657829003601f168201915b5050505050905090565b600061141061135a61177e565b8461140b856040518060600160405280602581526020016122fe602591396001600061138461177e565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c3e9092919063ffffffff16565b611786565b6001905092915050565b600061142e61142761177e565b848461197d565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611565576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252605c8152602001806120ac605c913960600191505060405180910390fd5b811561167d578273ffffffffffffffffffffffffffffffffffffffff166108fc6007549081150290604051600060405180830381858888f193505050501580156115b3573d6000803e3d6000fd5b508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8460016040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561163c57600080fd5b505af1158015611650573d6000803e3d6000fd5b505050506040513d602081101561166657600080fd5b810190808051906020019092919050505050611767565b3373ffffffffffffffffffffffffffffffffffffffff166108fc6007549081150290604051600060405180830381858888f193505050501580156116c5573d6000803e3d6000fd5b508073ffffffffffffffffffffffffffffffffffffffff166379cc67908460016040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b15801561174e57600080fd5b505af1158015611762573d6000803e3d6000fd5b505050505b600754600b60008282540392505081905550505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141561180c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806122b46024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611892576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806121086022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611a03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602581526020018061228f6025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611a89576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602381526020018061203f6023913960400191505060405180910390fd5b611a94838383611f4a565b611aff816040518060600160405280602681526020016121fc602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c3e9092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611b92816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611cfe90919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290611ceb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611cb0578082015181840152602081019050611c95565b50505050905090810190601f168015611cdd5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080828401905083811015611d7c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611e0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061226e6021913960400191505060405180910390fd5b611e1882600083611f4a565b611e8381604051806060016040528060228152602001612062602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c3e9092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611eda81600254611f4f90919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b505050565b6000611f9183836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611c3e565b905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611fda57805160ff1916838001178555612008565b82800160010185558215612008579182015b82811115612007578251825591602001919060010190611fec565b5b5090506120159190612019565b5090565b61203b91905b8082111561203757600081600090555060010161201f565b5090565b9056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654e6f7420656e6f75676820616d6f756e7420746f207061792074686520636f6c6c61746572616c734f6e6c79207265717565737465722063616e2063616c6c207468697320746f20776f726b6572732077686f207061727469636970616e747320697473207461736b20616e6420686173206e6f74206265656e2072657761726465642e45524332303a20617070726f766520746f20746865207a65726f2061646472657373546865206465706f736974696f6e206d757374206c6172676572207468616e207265776172647320746f20776f726b65727320746f20707562756c697368207461736b2e54686520776f726b657220646f6573206e6f74206861766520746865207469636b657420746f206a6f696e20746865207461736b4f6e6c79207265717565737465722063616e2075706c6f6164207472616e73616374696f6e20746f207468697320636f6e7472616374546865207461736b20646f206e6f74206e65656420776f726b65727320616e796d6f726545524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573734f6e6c79207265676973746572656420776f726b65722063616e207375626d6974206461746145524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220d3bb083984eeeae4b2c63b22a61d66989026512b1775b8237da938a588a804f864736f6c63430006030033"

// DeployCtask deploys a new Ethereum contract, binding an instance of Ctask to it.
func DeployCtask(auth *bind.TransactOpts, backend bind.ContractBackend, workerRequired *big.Int, rewards *big.Int, description string) (common.Address, *types.Transaction, *Ctask, error) {
	parsed, err := abi.JSON(strings.NewReader(CtaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CtaskBin), backend, workerRequired, rewards, description)
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ctask *CtaskCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ctask *CtaskSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ctask.Contract.Allowance(&_Ctask.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Ctask *CtaskCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Ctask.Contract.Allowance(&_Ctask.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ctask *CtaskCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ctask *CtaskSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ctask.Contract.BalanceOf(&_Ctask.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Ctask *CtaskCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Ctask.Contract.BalanceOf(&_Ctask.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ctask *CtaskCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ctask *CtaskSession) Decimals() (uint8, error) {
	return _Ctask.Contract.Decimals(&_Ctask.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Ctask *CtaskCallerSession) Decimals() (uint8, error) {
	return _Ctask.Contract.Decimals(&_Ctask.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ctask *CtaskCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ctask *CtaskSession) Name() (string, error) {
	return _Ctask.Contract.Name(&_Ctask.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ctask *CtaskCallerSession) Name() (string, error) {
	return _Ctask.Contract.Name(&_Ctask.CallOpts)
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

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ctask *CtaskCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ctask *CtaskSession) Symbol() (string, error) {
	return _Ctask.Contract.Symbol(&_Ctask.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ctask *CtaskCallerSession) Symbol() (string, error) {
	return _Ctask.Contract.Symbol(&_Ctask.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ctask *CtaskCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ctask.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ctask *CtaskSession) TotalSupply() (*big.Int, error) {
	return _Ctask.Contract.TotalSupply(&_Ctask.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Ctask *CtaskCallerSession) TotalSupply() (*big.Int, error) {
	return _Ctask.Contract.TotalSupply(&_Ctask.CallOpts)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ctask *CtaskTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ctask *CtaskSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Approve(&_Ctask.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Ctask *CtaskTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Approve(&_Ctask.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Ctask *CtaskTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Ctask *CtaskSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Burn(&_Ctask.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Ctask *CtaskTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Burn(&_Ctask.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Ctask *CtaskTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Ctask *CtaskSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.BurnFrom(&_Ctask.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Ctask *CtaskTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.BurnFrom(&_Ctask.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ctask *CtaskTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ctask *CtaskSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.DecreaseAllowance(&_Ctask.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Ctask *CtaskTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.DecreaseAllowance(&_Ctask.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ctask *CtaskTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ctask *CtaskSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.IncreaseAllowance(&_Ctask.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Ctask *CtaskTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.IncreaseAllowance(&_Ctask.TransactOpts, spender, addedValue)
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

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ctask *CtaskTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ctask *CtaskSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Transfer(&_Ctask.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Ctask *CtaskTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.Transfer(&_Ctask.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ctask *CtaskTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ctask *CtaskSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.TransferFrom(&_Ctask.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Ctask *CtaskTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ctask.Contract.TransferFrom(&_Ctask.TransactOpts, sender, recipient, amount)
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

// CtaskApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ctask contract.
type CtaskApprovalIterator struct {
	Event *CtaskApproval // Event containing the contract specifics and raw log

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
func (it *CtaskApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtaskApproval)
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
		it.Event = new(CtaskApproval)
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
func (it *CtaskApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtaskApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtaskApproval represents a Approval event raised by the Ctask contract.
type CtaskApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ctask *CtaskFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CtaskApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ctask.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CtaskApprovalIterator{contract: _Ctask.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Ctask *CtaskFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CtaskApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Ctask.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtaskApproval)
				if err := _Ctask.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ctask *CtaskFilterer) ParseApproval(log types.Log) (*CtaskApproval, error) {
	event := new(CtaskApproval)
	if err := _Ctask.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
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

// CtaskTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ctask contract.
type CtaskTransferIterator struct {
	Event *CtaskTransfer // Event containing the contract specifics and raw log

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
func (it *CtaskTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CtaskTransfer)
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
		it.Event = new(CtaskTransfer)
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
func (it *CtaskTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CtaskTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CtaskTransfer represents a Transfer event raised by the Ctask contract.
type CtaskTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ctask *CtaskFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CtaskTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ctask.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CtaskTransferIterator{contract: _Ctask.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Ctask *CtaskFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CtaskTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Ctask.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CtaskTransfer)
				if err := _Ctask.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ctask *CtaskFilterer) ParseTransfer(log types.Log) (*CtaskTransfer, error) {
	event := new(CtaskTransfer)
	if err := _Ctask.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
