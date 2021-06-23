// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package task

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

// TaskABI is the input ABI used to generate the binding from.
const TaskABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"workerRequired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"DataSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"TaskPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_worker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_remainingWorkers\",\"type\":\"uint256\"}],\"name\":\"TaskRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"worker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isok\",\"type\":\"bool\"}],\"name\":\"Rewarding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"SubmitData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingWorkers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rem\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// TaskBin is the compiled bytecode used for deploying new contracts.
var TaskBin = "0x60806040523480156200001157600080fd5b506040516200258038038062002580833981810160405260608110156200003757600080fd5b810190808051906020019092919080519060200190929190805160405193929190846401000000008211156200006c57600080fd5b838201915060208201858111156200008357600080fd5b8251866001820283011164010000000082111715620000a157600080fd5b8083526020830192505050908051906020019080838360005b83811015620000d7578082015181840152602081019050620000ba565b50505050905090810190601f168015620001055780820380516001836020036101000a031916815260200191505b506040525050506040518060400160405280601481526020017f5461736b5375626d697373696f6e5469636b65740000000000000000000000008152506040518060400160405280600381526020017f545354000000000000000000000000000000000000000000000000000000000081525081600390805190602001906200019092919062000278565b508060049080519060200190620001a992919062000278565b506012600560006101000a81548160ff021916908360ff160217905550505082600681905550816007819055508060099080519060200190620001ee92919062000278565b5033600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600c819055506000600b819055506200025160006200025a60201b60201c565b50505062000327565b80600560006101000a81548160ff021916908360ff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002bb57805160ff1916838001178555620002ec565b82800160010185558215620002ec579182015b82811115620002eb578251825591602001919060010190620002ce565b5b509050620002fb9190620002ff565b5090565b6200032491905b808211156200032057600081600090555060010162000306565b5090565b90565b61224980620003376000396000f3fe60806040526004361061010d5760003560e01c8063409ef9e21161009557806395d89b411161006457806395d89b411461064d578063a457c2d7146106dd578063a63ac6c214610750578063a9059cbb146107ad578063dd62ed3e1461082057610114565b8063409ef9e21461048a57806342966c681461055257806370a082311461058d57806379cc6790146105f257610114565b80631aa3a008116100dc5780631aa3a0081461031157806323b872dd14610328578063313ce567146103bb57806338b270d6146103ec578063395093511461041757610114565b806306fdde03146101cc578063075d47821461025c578063095ea7b31461027357806318160ddd146102e657610114565b3661011457005b3373ffffffffffffffffffffffffffffffffffffffff16600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001806120936036913960400191505060405180910390fd5b34600b60008282540192505081905550005b3480156101d857600080fd5b506101e16108a5565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610221578082015181840152602081019050610206565b50505050905090810190601f16801561024e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561026857600080fd5b50610271610947565b005b34801561027f57600080fd5b506102cc6004803603604081101561029657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b8f565b604051808215151515815260200191505060405180910390f35b3480156102f257600080fd5b506102fb610bad565b6040518082815260200191505060405180910390f35b34801561031d57600080fd5b50610326610bb7565b005b34801561033457600080fd5b506103a16004803603606081101561034b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d19565b604051808215151515815260200191505060405180910390f35b3480156103c757600080fd5b506103d0610df2565b604051808260ff1660ff16815260200191505060405180910390f35b3480156103f857600080fd5b50610401610e09565b6040518082815260200191505060405180910390f35b34801561042357600080fd5b506104706004803603604081101561043a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e13565b604051808215151515815260200191505060405180910390f35b34801561049657600080fd5b50610550600480360360208110156104ad57600080fd5b81019080803590602001906401000000008111156104ca57600080fd5b8201836020820111156104dc57600080fd5b803590602001918460018302840111640100000000831117156104fe57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610ec6565b005b34801561055e57600080fd5b5061058b6004803603602081101561057557600080fd5b81019080803590602001909291905050506110b0565b005b34801561059957600080fd5b506105dc600480360360208110156105b057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506110c4565b6040518082815260200191505060405180910390f35b3480156105fe57600080fd5b5061064b6004803603604081101561061557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061110c565b005b34801561065957600080fd5b5061066261116e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106a2578082015181840152602081019050610687565b50505050905090810190601f1680156106cf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156106e957600080fd5b506107366004803603604081101561070057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611210565b604051808215151515815260200191505060405180910390f35b34801561075c57600080fd5b506107ab6004803603604081101561077357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035151590602001909291905050506112dd565b005b3480156107b957600080fd5b50610806600480360360408110156107d057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611437565b604051808215151515815260200191505060405180910390f35b34801561082c57600080fd5b5061088f6004803603604081101561084357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611455565b6040518082815260200191505060405180910390f35b606060038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561093d5780601f106109125761010080835404028352916020019161093d565b820191906000526020600020905b81548152906001019060200180831161092057829003601f168201915b5050505050905090565b3373ffffffffffffffffffffffffffffffffffffffff16600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146109ed576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526028815260200180611fa96028913960400191505060405180910390fd5b600b54600654600754021115610a4e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252604481526020018061204f6044913960600191505060405180910390fd5b600654600b5481610a5b57fe5b04600781905550610a90600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166006546114dc565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b660075460096040518083815260200180602001828103825283818154600181600116156101000203166002900481526020019150805460018160011615610100020316600290048015610b7e5780601f10610b5357610100808354040283529160200191610b7e565b820191906000526020600020905b815481529060010190602001808311610b6157829003601f168201915b5050935050505060405180910390a2565b6000610ba3610b9c6116a3565b84846116ab565b6001905092915050565b6000600254905090565b6000600c5411610c12576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806120c96024913960400191505060405180910390fd5b6000339050600c60008154809291906001900391905055506008819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b30562600c546040518082815260200191505060405180910390a350565b6000610d268484846118a2565b610de784610d326116a3565b610de28560405180606001604052806028815260200161211360289139600160008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610d986116a3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611b639092919063ffffffff16565b6116ab565b600190509392505050565b6000600560009054906101000a900460ff16905090565b6000600c54905090565b6000610ebc610e206116a3565b84610eb78560016000610e316116a3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c2390919063ffffffff16565b6116ab565b6001905092915050565b6000600c54600654039050600080905060008090505b82811015610f65573373ffffffffffffffffffffffffffffffffffffffff1660088281548110610f0857fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610f585760019150610f65565b8080600101915050610edc565b50600d83908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610fa1929190611ebe565b5080610ff8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806121c96026913960400191505060405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3846040518080602001828103825283818151815260200191508051906020019080838360005b83811015611071578082015181840152602081019050611056565b50505050905090810190601f16801561109e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a2505050565b6110c16110bb6116a3565b82611cab565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600061114b8260405180606001604052806024815260200161213b6024913961113c866111376116a3565b611455565b611b639092919063ffffffff16565b905061115f836111596116a3565b836116ab565b6111698383611cab565b505050565b606060048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112065780601f106111db57610100808354040283529160200191611206565b820191906000526020600020905b8154815290600101906020018083116111e957829003601f168201915b5050505050905090565b60006112d361121d6116a3565b846112ce856040518060600160405280602581526020016121ef60259139600160006112476116a3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611b639092919063ffffffff16565b6116ab565b6001905092915050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611383576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252605c815260200180611fd1605c913960600191505060405180910390fd5b80156113d7578173ffffffffffffffffffffffffffffffffffffffff166108fc6007549081150290604051600060405180830381858888f193505050501580156113d1573d6000803e3d6000fd5b50611421565b3373ffffffffffffffffffffffffffffffffffffffff166108fc6007549081150290604051600060405180830381858888f1935050505015801561141f573d6000803e3d6000fd5b505b600754600b600082825403925050819055505050565b600061144b6114446116a3565b84846118a2565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141561157f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081525060200191505060405180910390fd5b61158b60008383611e6f565b6115a081600254611c2390919063ffffffff16565b6002819055506115f7816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c2390919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611731576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806121a56024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156117b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061202d6022913960400191505060405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415611928576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806121806025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156119ae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611f646023913960400191505060405180910390fd5b6119b9838383611e6f565b611a24816040518060600160405280602681526020016120ed602691396000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611b639092919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611ab7816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611c2390919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a3505050565b6000838311158290611c10576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611bd5578082015181840152602081019050611bba565b50505050905090810190601f168015611c025780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385039050809150509392505050565b600080828401905083811015611ca1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611d31576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602181526020018061215f6021913960400191505060405180910390fd5b611d3d82600083611e6f565b611da881604051806060016040528060228152602001611f87602291396000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611b639092919063ffffffff16565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611dff81600254611e7490919063ffffffff16565b600281905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35050565b505050565b6000611eb683836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611b63565b905092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10611eff57805160ff1916838001178555611f2d565b82800160010185558215611f2d579182015b82811115611f2c578251825591602001919060010190611f11565b5b509050611f3a9190611f3e565b5090565b611f6091905b80821115611f5c576000816000905550600101611f44565b5090565b9056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654e6f7420656e6f75676820616d6f756e7420746f207061792074686520636f6c6c61746572616c734f6e6c79207265717565737465722063616e2063616c6c207468697320746f20776f726b6572732077686f207061727469636970616e747320697473207461736b20616e6420686173206e6f74206265656e2072657761726465642e45524332303a20617070726f766520746f20746865207a65726f2061646472657373546865206465706f736974696f6e206d757374206c6172676572207468616e207265776172647320746f20776f726b65727320746f20707562756c697368207461736b2e4f6e6c79207265717565737465722063616e2075706c6f6164207472616e73616374696f6e20746f207468697320636f6e7472616374546865207461736b20646f206e6f74206e65656420776f726b65727320616e796d6f726545524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f20616464726573734f6e6c79207265676973746572656420776f726b65722063616e207375626d6974206461746145524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212206946a439ff4eb32bdda99689a3132d09414512396496a9fdfb18fac52ceb4bba64736f6c63430006030033"

// DeployTask deploys a new Ethereum contract, binding an instance of Task to it.
func DeployTask(auth *bind.TransactOpts, backend bind.ContractBackend, workerRequired *big.Int, rewards *big.Int, description string) (common.Address, *types.Transaction, *Task, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TaskBin), backend, workerRequired, rewards, description)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// Task is an auto generated Go binding around an Ethereum contract.
type Task struct {
	TaskCaller     // Read-only binding to the contract
	TaskTransactor // Write-only binding to the contract
	TaskFilterer   // Log filterer for contract events
}

// TaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskSession struct {
	Contract     *Task             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskCallerSession struct {
	Contract *TaskCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskTransactorSession struct {
	Contract     *TaskTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskRaw struct {
	Contract *Task // Generic contract binding to access the raw methods on
}

// TaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskCallerRaw struct {
	Contract *TaskCaller // Generic read-only contract binding to access the raw methods on
}

// TaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskTransactorRaw struct {
	Contract *TaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTask creates a new instance of Task, bound to a specific deployed contract.
func NewTask(address common.Address, backend bind.ContractBackend) (*Task, error) {
	contract, err := bindTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Task{TaskCaller: TaskCaller{contract: contract}, TaskTransactor: TaskTransactor{contract: contract}, TaskFilterer: TaskFilterer{contract: contract}}, nil
}

// NewTaskCaller creates a new read-only instance of Task, bound to a specific deployed contract.
func NewTaskCaller(address common.Address, caller bind.ContractCaller) (*TaskCaller, error) {
	contract, err := bindTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskCaller{contract: contract}, nil
}

// NewTaskTransactor creates a new write-only instance of Task, bound to a specific deployed contract.
func NewTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskTransactor, error) {
	contract, err := bindTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskTransactor{contract: contract}, nil
}

// NewTaskFilterer creates a new log filterer instance of Task, bound to a specific deployed contract.
func NewTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskFilterer, error) {
	contract, err := bindTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskFilterer{contract: contract}, nil
}

// bindTask binds a generic wrapper to an already deployed contract.
func bindTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.TaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.TaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Task *TaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Task.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Task *TaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Task *TaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Task.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Task *TaskCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Task *TaskSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Task.Contract.Allowance(&_Task.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Task *TaskCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Task.Contract.Allowance(&_Task.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Task *TaskCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Task *TaskSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Task *TaskCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Task.Contract.BalanceOf(&_Task.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Task *TaskCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Task *TaskSession) Decimals() (uint8, error) {
	return _Task.Contract.Decimals(&_Task.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Task *TaskCallerSession) Decimals() (uint8, error) {
	return _Task.Contract.Decimals(&_Task.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Task *TaskCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Task *TaskSession) Name() (string, error) {
	return _Task.Contract.Name(&_Task.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Task *TaskCallerSession) Name() (string, error) {
	return _Task.Contract.Name(&_Task.CallOpts)
}

// RemainingWorkers is a free data retrieval call binding the contract method 0x38b270d6.
//
// Solidity: function remainingWorkers() view returns(uint256 rem)
func (_Task *TaskCaller) RemainingWorkers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "remainingWorkers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingWorkers is a free data retrieval call binding the contract method 0x38b270d6.
//
// Solidity: function remainingWorkers() view returns(uint256 rem)
func (_Task *TaskSession) RemainingWorkers() (*big.Int, error) {
	return _Task.Contract.RemainingWorkers(&_Task.CallOpts)
}

// RemainingWorkers is a free data retrieval call binding the contract method 0x38b270d6.
//
// Solidity: function remainingWorkers() view returns(uint256 rem)
func (_Task *TaskCallerSession) RemainingWorkers() (*big.Int, error) {
	return _Task.Contract.RemainingWorkers(&_Task.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Task *TaskCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Task *TaskSession) Symbol() (string, error) {
	return _Task.Contract.Symbol(&_Task.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Task *TaskCallerSession) Symbol() (string, error) {
	return _Task.Contract.Symbol(&_Task.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Task *TaskCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Task.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Task *TaskSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Task *TaskCallerSession) TotalSupply() (*big.Int, error) {
	return _Task.Contract.TotalSupply(&_Task.CallOpts)
}

// Rewarding is a paid mutator transaction binding the contract method 0xa63ac6c2.
//
// Solidity: function Rewarding(address worker, bool isok) returns()
func (_Task *TaskTransactor) Rewarding(opts *bind.TransactOpts, worker common.Address, isok bool) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "Rewarding", worker, isok)
}

// Rewarding is a paid mutator transaction binding the contract method 0xa63ac6c2.
//
// Solidity: function Rewarding(address worker, bool isok) returns()
func (_Task *TaskSession) Rewarding(worker common.Address, isok bool) (*types.Transaction, error) {
	return _Task.Contract.Rewarding(&_Task.TransactOpts, worker, isok)
}

// Rewarding is a paid mutator transaction binding the contract method 0xa63ac6c2.
//
// Solidity: function Rewarding(address worker, bool isok) returns()
func (_Task *TaskTransactorSession) Rewarding(worker common.Address, isok bool) (*types.Transaction, error) {
	return _Task.Contract.Rewarding(&_Task.TransactOpts, worker, isok)
}

// SubmitData is a paid mutator transaction binding the contract method 0x409ef9e2.
//
// Solidity: function SubmitData(bytes data) returns()
func (_Task *TaskTransactor) SubmitData(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "SubmitData", data)
}

// SubmitData is a paid mutator transaction binding the contract method 0x409ef9e2.
//
// Solidity: function SubmitData(bytes data) returns()
func (_Task *TaskSession) SubmitData(data []byte) (*types.Transaction, error) {
	return _Task.Contract.SubmitData(&_Task.TransactOpts, data)
}

// SubmitData is a paid mutator transaction binding the contract method 0x409ef9e2.
//
// Solidity: function SubmitData(bytes data) returns()
func (_Task *TaskTransactorSession) SubmitData(data []byte) (*types.Transaction, error) {
	return _Task.Contract.SubmitData(&_Task.TransactOpts, data)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Task *TaskTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Task *TaskSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Approve(&_Task.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Task *TaskTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Approve(&_Task.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Task *TaskTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Task *TaskSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Burn(&_Task.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Task *TaskTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Burn(&_Task.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Task *TaskTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Task *TaskSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.BurnFrom(&_Task.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Task *TaskTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.BurnFrom(&_Task.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Task *TaskTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Task *TaskSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Task.Contract.DecreaseAllowance(&_Task.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Task *TaskTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Task.Contract.DecreaseAllowance(&_Task.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Task *TaskTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Task *TaskSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Task.Contract.IncreaseAllowance(&_Task.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Task *TaskTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Task.Contract.IncreaseAllowance(&_Task.TransactOpts, spender, addedValue)
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns()
func (_Task *TaskTransactor) Publish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "publish")
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns()
func (_Task *TaskSession) Publish() (*types.Transaction, error) {
	return _Task.Contract.Publish(&_Task.TransactOpts)
}

// Publish is a paid mutator transaction binding the contract method 0x075d4782.
//
// Solidity: function publish() returns()
func (_Task *TaskTransactorSession) Publish() (*types.Transaction, error) {
	return _Task.Contract.Publish(&_Task.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Task *TaskTransactor) Register(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "register")
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Task *TaskSession) Register() (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1aa3a008.
//
// Solidity: function register() returns()
func (_Task *TaskTransactorSession) Register() (*types.Transaction, error) {
	return _Task.Contract.Register(&_Task.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Task *TaskTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Task *TaskSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Task *TaskTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.Transfer(&_Task.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Task *TaskTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Task *TaskSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.TransferFrom(&_Task.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Task *TaskTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Task.Contract.TransferFrom(&_Task.TransactOpts, sender, recipient, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Task *TaskTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Task.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Task *TaskSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Task.Contract.Fallback(&_Task.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Task *TaskTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Task.Contract.Fallback(&_Task.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Task *TaskTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Task.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Task *TaskSession) Receive() (*types.Transaction, error) {
	return _Task.Contract.Receive(&_Task.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Task *TaskTransactorSession) Receive() (*types.Transaction, error) {
	return _Task.Contract.Receive(&_Task.TransactOpts)
}

// TaskApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Task contract.
type TaskApprovalIterator struct {
	Event *TaskApproval // Event containing the contract specifics and raw log

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
func (it *TaskApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskApproval)
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
		it.Event = new(TaskApproval)
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
func (it *TaskApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskApproval represents a Approval event raised by the Task contract.
type TaskApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Task *TaskFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TaskApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TaskApprovalIterator{contract: _Task.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Task *TaskFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TaskApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskApproval)
				if err := _Task.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Task *TaskFilterer) ParseApproval(log types.Log) (*TaskApproval, error) {
	event := new(TaskApproval)
	if err := _Task.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TaskDataSubmittedIterator is returned from FilterDataSubmitted and is used to iterate over the raw logs and unpacked data for DataSubmitted events raised by the Task contract.
type TaskDataSubmittedIterator struct {
	Event *TaskDataSubmitted // Event containing the contract specifics and raw log

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
func (it *TaskDataSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskDataSubmitted)
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
		it.Event = new(TaskDataSubmitted)
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
func (it *TaskDataSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskDataSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskDataSubmitted represents a DataSubmitted event raised by the Task contract.
type TaskDataSubmitted struct {
	From common.Address
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDataSubmitted is a free log retrieval operation binding the contract event 0x5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3.
//
// Solidity: event DataSubmitted(address indexed _from, bytes _data)
func (_Task *TaskFilterer) FilterDataSubmitted(opts *bind.FilterOpts, _from []common.Address) (*TaskDataSubmittedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "DataSubmitted", _fromRule)
	if err != nil {
		return nil, err
	}
	return &TaskDataSubmittedIterator{contract: _Task.contract, event: "DataSubmitted", logs: logs, sub: sub}, nil
}

// WatchDataSubmitted is a free log subscription operation binding the contract event 0x5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d3.
//
// Solidity: event DataSubmitted(address indexed _from, bytes _data)
func (_Task *TaskFilterer) WatchDataSubmitted(opts *bind.WatchOpts, sink chan<- *TaskDataSubmitted, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "DataSubmitted", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskDataSubmitted)
				if err := _Task.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
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
func (_Task *TaskFilterer) ParseDataSubmitted(log types.Log) (*TaskDataSubmitted, error) {
	event := new(TaskDataSubmitted)
	if err := _Task.contract.UnpackLog(event, "DataSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TaskTaskPublishedIterator is returned from FilterTaskPublished and is used to iterate over the raw logs and unpacked data for TaskPublished events raised by the Task contract.
type TaskTaskPublishedIterator struct {
	Event *TaskTaskPublished // Event containing the contract specifics and raw log

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
func (it *TaskTaskPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskPublished)
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
		it.Event = new(TaskTaskPublished)
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
func (it *TaskTaskPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskPublished represents a TaskPublished event raised by the Task contract.
type TaskTaskPublished struct {
	Rewards     *big.Int
	Requester   common.Address
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTaskPublished is a free log retrieval operation binding the contract event 0x31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b6.
//
// Solidity: event TaskPublished(uint256 rewards, address indexed _requester, string _description)
func (_Task *TaskFilterer) FilterTaskPublished(opts *bind.FilterOpts, _requester []common.Address) (*TaskTaskPublishedIterator, error) {

	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskPublished", _requesterRule)
	if err != nil {
		return nil, err
	}
	return &TaskTaskPublishedIterator{contract: _Task.contract, event: "TaskPublished", logs: logs, sub: sub}, nil
}

// WatchTaskPublished is a free log subscription operation binding the contract event 0x31c1351bb52f89f29d764c4520fd6a6f0d51eff05a2067483c9d3c1a2c12b9b6.
//
// Solidity: event TaskPublished(uint256 rewards, address indexed _requester, string _description)
func (_Task *TaskFilterer) WatchTaskPublished(opts *bind.WatchOpts, sink chan<- *TaskTaskPublished, _requester []common.Address) (event.Subscription, error) {

	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskPublished", _requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskPublished)
				if err := _Task.contract.UnpackLog(event, "TaskPublished", log); err != nil {
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
func (_Task *TaskFilterer) ParseTaskPublished(log types.Log) (*TaskTaskPublished, error) {
	event := new(TaskTaskPublished)
	if err := _Task.contract.UnpackLog(event, "TaskPublished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TaskTaskRegisteredIterator is returned from FilterTaskRegistered and is used to iterate over the raw logs and unpacked data for TaskRegistered events raised by the Task contract.
type TaskTaskRegisteredIterator struct {
	Event *TaskTaskRegistered // Event containing the contract specifics and raw log

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
func (it *TaskTaskRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTaskRegistered)
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
		it.Event = new(TaskTaskRegistered)
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
func (it *TaskTaskRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTaskRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTaskRegistered represents a TaskRegistered event raised by the Task contract.
type TaskTaskRegistered struct {
	Worker           common.Address
	Requester        common.Address
	RemainingWorkers *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskRegistered is a free log retrieval operation binding the contract event 0x224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b30562.
//
// Solidity: event TaskRegistered(address indexed _worker, address indexed _requester, uint256 _remainingWorkers)
func (_Task *TaskFilterer) FilterTaskRegistered(opts *bind.FilterOpts, _worker []common.Address, _requester []common.Address) (*TaskTaskRegisteredIterator, error) {

	var _workerRule []interface{}
	for _, _workerItem := range _worker {
		_workerRule = append(_workerRule, _workerItem)
	}
	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "TaskRegistered", _workerRule, _requesterRule)
	if err != nil {
		return nil, err
	}
	return &TaskTaskRegisteredIterator{contract: _Task.contract, event: "TaskRegistered", logs: logs, sub: sub}, nil
}

// WatchTaskRegistered is a free log subscription operation binding the contract event 0x224eb90d6d6f91b35a68921db160c1eb0d0f59578400445d86e5538e37b30562.
//
// Solidity: event TaskRegistered(address indexed _worker, address indexed _requester, uint256 _remainingWorkers)
func (_Task *TaskFilterer) WatchTaskRegistered(opts *bind.WatchOpts, sink chan<- *TaskTaskRegistered, _worker []common.Address, _requester []common.Address) (event.Subscription, error) {

	var _workerRule []interface{}
	for _, _workerItem := range _worker {
		_workerRule = append(_workerRule, _workerItem)
	}
	var _requesterRule []interface{}
	for _, _requesterItem := range _requester {
		_requesterRule = append(_requesterRule, _requesterItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "TaskRegistered", _workerRule, _requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTaskRegistered)
				if err := _Task.contract.UnpackLog(event, "TaskRegistered", log); err != nil {
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
func (_Task *TaskFilterer) ParseTaskRegistered(log types.Log) (*TaskTaskRegistered, error) {
	event := new(TaskTaskRegistered)
	if err := _Task.contract.UnpackLog(event, "TaskRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TaskTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Task contract.
type TaskTransferIterator struct {
	Event *TaskTransfer // Event containing the contract specifics and raw log

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
func (it *TaskTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskTransfer)
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
		it.Event = new(TaskTransfer)
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
func (it *TaskTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskTransfer represents a Transfer event raised by the Task contract.
type TaskTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Task *TaskFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TaskTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Task.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TaskTransferIterator{contract: _Task.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Task *TaskFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TaskTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Task.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskTransfer)
				if err := _Task.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Task *TaskFilterer) ParseTransfer(log types.Log) (*TaskTransfer, error) {
	event := new(TaskTransfer)
	if err := _Task.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
