// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lib

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
var CrowdsourcingBin = "0x60806040523480156200001157600080fd5b506040518060400160405280601481526020017f5461736b5375626d697373696f6e5469636b6574000000000000000000000000815250604051806040016040528060038152602001621514d560ea1b81525081600390805190602001906200007c929190620000d3565b50805162000092906004906020840190620000d3565b50506005805460ff1916601217905550620000b760006001600160e01b03620000bd16565b62000178565b6005805460ff191660ff92909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200011657805160ff191683800117855562000146565b8280016001018555821562000146579182015b828111156200014657825182559160200191906001019062000129565b506200015492915062000158565b5090565b6200017591905b808211156200015457600081556001016200015f565b90565b611b4d80620001886000396000f3fe6080604052600436106101095760003560e01c806342966c681161009557806395d89b411161006457806395d89b41146107b5578063a457c2d7146107ca578063a63ac6c214610803578063a9059cbb1461083e578063dd62ed3e1461087757610109565b806342966c681461065c5780636bdeba391461068657806370a082311461074957806379cc67901461077c57610109565b806323b872dd116100dc57806323b872dd146104c157806327549207146105045780632be319f9146105c5578063313ce567146105f8578063395093511461062357610109565b806306fdde0314610390578063095ea7b31461041a578063099124511461046757806318160ddd146104ac575b6040805180820182526001808252600160f81b60208084019182528451808601865292835260008382015293518351939492938593919091019182918083835b602083106101685780518252601f199092019160209182019101610149565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060003660405160200180838380828437808301925050509250505060405160208183030381529060405280519060200120148061027c5750806040516020018082805190602001908083835b6020831061020e5780518252601f1990920191602091820191016101ef565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012060003660405160200180838380828437808301925050509250505060405160208183030381529060405280519060200120145b6102b75760405162461bcd60e51b8152600401808060200182810382526028815260200180611acb6028913960400191505060405180910390fd5b816040516020018082805190602001908083835b602083106102ea5780518252601f1990920191602091820191016102cb565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001206000366040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012014156103765733600090815260076020526040902080543401905561038c565b3360009081526006602052604090208054340190555b5050005b34801561039c57600080fd5b506103a56108b2565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103df5781810151838201526020016103c7565b50505050905090810190601f16801561040c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561042657600080fd5b506104536004803603604081101561043d57600080fd5b506001600160a01b038135169060200135610948565b604080519115158252519081900360200190f35b34801561047357600080fd5b5061049a6004803603602081101561048a57600080fd5b50356001600160a01b0316610965565b60408051918252519081900360200190f35b3480156104b857600080fd5b5061049a610980565b3480156104cd57600080fd5b50610453600480360360608110156104e457600080fd5b506001600160a01b03813581169160208101359091169060400135610986565b34801561051057600080fd5b506105c36004803603606081101561052757600080fd5b81359160208101359181019060608101604082013564010000000081111561054e57600080fd5b82018360208201111561056057600080fd5b8035906020019184600183028401116401000000008311171561058257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610a13945050505050565b005b3480156105d157600080fd5b506105c3600480360360208110156105e857600080fd5b50356001600160a01b0316610b92565b34801561060457600080fd5b5061060d610cc9565b6040805160ff9092168252519081900360200190f35b34801561062f57600080fd5b506104536004803603604081101561064657600080fd5b506001600160a01b038135169060200135610cd2565b34801561066857600080fd5b506105c36004803603602081101561067f57600080fd5b5035610d26565b34801561069257600080fd5b506105c3600480360360408110156106a957600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156106d457600080fd5b8201836020820111156106e657600080fd5b8035906020019184600183028401116401000000008311171561070857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610d3a945050505050565b34801561075557600080fd5b5061049a6004803603602081101561076c57600080fd5b50356001600160a01b0316610e7b565b34801561078857600080fd5b506105c36004803603604081101561079f57600080fd5b506001600160a01b038135169060200135610e96565b3480156107c157600080fd5b506103a5610ef6565b3480156107d657600080fd5b50610453600480360360408110156107ed57600080fd5b506001600160a01b038135169060200135610f57565b34801561080f57600080fd5b506105c36004803603604081101561082657600080fd5b506001600160a01b0381351690602001351515610fc5565b34801561084a57600080fd5b506104536004803603604081101561086157600080fd5b506001600160a01b038135169060200135611155565b34801561088357600080fd5b5061049a6004803603604081101561089a57600080fd5b506001600160a01b0381358116916020013516611169565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561093e5780601f106109135761010080835404028352916020019161093e565b820191906000526020600020905b81548152906001019060200180831161092157829003601f168201915b5050505050905090565b600061095c610955611194565b8484611198565b50600192915050565b6001600160a01b031660009081526008602052604090205490565b60025490565b6000610993848484611284565b610a098461099f611194565b610a0485604051806060016040528060288152602001611a15602891396001600160a01b038a166000908152600160205260408120906109dd611194565b6001600160a01b03168152602081019190915260400160002054919063ffffffff6113eb16565b611198565b5060019392505050565b33600081815260076020526040902054841115610a615760405162461bcd60e51b81526004018080602001828103825260288152602001806119496028913960400191505060405180910390fd5b6000838581610a6c57fe5b049050610aa683610aa16040518060400160405280600a81526020016852932bbb0b932399d160b51b815250610aa185611482565b61157d565b6001600160a01b038316600090815260076020908152604080832080548a9003905560098252808320859055600890915290208590559250610ae88285611638565b816001600160a01b03167f362295b99791fb7f9cc554e61353656ac0f62436a3a559d02ad33dd59a038de6846040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b51578181015183820152602001610b39565b50505050905090810190601f168015610b7e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050505050565b6001600160a01b038116600090815260086020526040902054339080610be95760405162461bcd60e51b81526004018080602001828103825260248152602001806118e06024913960400191505060405180910390fd5b6001600160a01b03808416600090815260096020908152604080832054938616835260069091529020541015610c505760405162461bcd60e51b81526004018080602001828103825260288152602001806119496028913960400191505060405180910390fd5b506001600160a01b039182166000818152600b6020908152604080832094909516808352938152848220805460ff19166001179055828252600981528482208054858452600a83528684208054909101905554938252600681528482208054949094039093559081526008909152208054600019019055565b60055460ff1690565b600061095c610cdf611194565b84610a048560016000610cf0611194565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61173416565b610d37610d31611194565b82611795565b50565b6001600160a01b0382166000908152600b6020908152604080832033845290915290205460ff1615610db3576040805162461bcd60e51b815260206004820152601f60248201527f596f7520616c7265616479207375626d697474656420796f7572206461746100604482015290519081900360640190fd5b6001600160a01b0382166000818152600b602090815260408083203384528252808320805460ff19166001179055805182815285518184015285517f5548223a2b92aad353ef26dad7dcdb43551e721b55887bb3a8b75f52f7bbc3d39487949293849390840192908601918190849084905b83811015610e3d578181015183820152602001610e25565b50505050905090810190601f168015610e6a5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050565b6001600160a01b031660009081526020819052604090205490565b6000610ed382604051806060016040528060248152602001611a3d60249139610ec686610ec1611194565b611169565b919063ffffffff6113eb16565b9050610ee783610ee1611194565b83611198565b610ef18383611795565b505050565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561093e5780601f106109135761010080835404028352916020019161093e565b600061095c610f64611194565b84610a0485604051806060016040528060258152602001611af36025913960016000610f8e611194565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff6113eb16565b336000908152600b602090815260408083206001600160a01b038616845290915290205460ff16151560011461102c5760405162461bcd60e51b815260040180806020018281038252605c815260200180611971605c913960600191505060405180910390fd5b801561109957336000908152600960209081526040808320546001600160a01b038616808552600a9093528184205491519293910180156108fc02929091818181858888f19350505050158015611087573d6000803e3d6000fd5b50611093826001611155565b5061111c565b6001600160a01b0382166000818152600a602052604080822054905181156108fc0292818181858888f193505050501580156110d9573d6000803e3d6000fd5b503360008181526009602052604080822054905181156108fc0292818181858888f19350505050158015611111573d6000803e3d6000fd5b5061111c6001610d26565b50336000908152600b602090815260408083206001600160a01b03949094168352928152828220805460ff19169055600a905290812055565b600061095c611162611194565b8484611284565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166111dd5760405162461bcd60e51b8152600401808060200182810382526024815260200180611aa76024913960400191505060405180910390fd5b6001600160a01b0382166112225760405162461bcd60e51b81526004018080602001828103825260228152602001806119cd6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166112c95760405162461bcd60e51b8152600401808060200182810382526025815260200180611a826025913960400191505060405180910390fd5b6001600160a01b03821661130e5760405162461bcd60e51b81526004018080602001828103825260238152602001806119046023913960400191505060405180910390fd5b611319838383610ef1565b61135c816040518060600160405280602681526020016119ef602691396001600160a01b038616600090815260208190526040902054919063ffffffff6113eb16565b6001600160a01b038085166000908152602081905260408082209390935590841681522054611391908263ffffffff61173416565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561147a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561143f578181015183820152602001611427565b50505050905090810190601f16801561146c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60408051606480825260a082019092526060919082908260208201818038833901905050905060005b84156114f1578151600a808704966001840193919006916030830160f81b91859181106114d457fe5b60200101906001600160f81b031916908160001a905350506114ab565b6060816040519080825280601f01601f19166020018201604052801561151e576020820181803883390190505b50905060005b8281101561157357838160018503038151811061153d57fe5b602001015160f81c60f81b82828151811061155457fe5b60200101906001600160f81b031916908160001a905350600101611524565b5095945050505050565b606082826040516020018083805190602001908083835b602083106115b35780518252601f199092019160209182019101611594565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b602083106115fb5780518252601f1990920191602091820191016115dc565b6001836020036101000a03801982511681845116808217855250505050505090500192505050604051602081830303815290604052905092915050565b6001600160a01b038216611693576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b61169f60008383610ef1565b6002546116b2908263ffffffff61173416565b6002556001600160a01b0382166000908152602081905260409020546116de908263ffffffff61173416565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b60008282018381101561178e576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b0382166117da5760405162461bcd60e51b8152600401808060200182810382526021815260200180611a616021913960400191505060405180910390fd5b6117e682600083610ef1565b61182981604051806060016040528060228152602001611927602291396001600160a01b038516600090815260208190526040902054919063ffffffff6113eb16565b6001600160a01b038316600090815260208190526040902055600254611855908263ffffffff61189d16565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600061178e83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152506113eb56fe546865207461736b206e6f206c6f6e676572206e65656473206d6f726520776f726b657245524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654e6f7420656e6f75676820616d6f756e7420746f207061792074686520636f6c6c61746572616c734f6e6c79207265717565737465722063616e2063616c6c207468697320746f20776f726b6572732077686f207061727469636970616e747320697473207461736b20616e6420686173206e6f74206265656e2072657761726465642e45524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f2061646472657373596f752073686f756c642065697468657220626520776f726b6572206f722072657175657374657245524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220692bd3e361f9c730354c6bdd5bb49e6e7a3cb80c5f422d3402f0a76423af293064736f6c63430006020033"

// DeployCrowdsourcing deploys a new ethereum contract, binding an instance of Crowdsourcing to it.
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

// Crowdsourcing is an auto generated Go binding around an ethereum contract.
type Crowdsourcing struct {
	CrowdsourcingCaller     // Read-only binding to the contract
	CrowdsourcingTransactor // Write-only binding to the contract
	CrowdsourcingFilterer   // Log filterer for contract events
}

// CrowdsourcingCaller is an auto generated read-only Go binding around an ethereum contract.
type CrowdsourcingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsourcingTransactor is an auto generated write-only Go binding around an ethereum contract.
type CrowdsourcingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsourcingFilterer is an auto generated log filtering Go binding around an ethereum contract events.
type CrowdsourcingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsourcingSession is an auto generated Go binding around an ethereum contract,
// with pre-set call and transact options.
type CrowdsourcingSession struct {
	Contract     *Crowdsourcing    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdsourcingCallerSession is an auto generated read-only Go binding around an ethereum contract,
// with pre-set call options.
type CrowdsourcingCallerSession struct {
	Contract *CrowdsourcingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CrowdsourcingTransactorSession is an auto generated write-only Go binding around an ethereum contract,
// with pre-set transact options.
type CrowdsourcingTransactorSession struct {
	Contract     *CrowdsourcingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CrowdsourcingRaw is an auto generated low-level Go binding around an ethereum contract.
type CrowdsourcingRaw struct {
	Contract *Crowdsourcing // Generic contract binding to access the raw methods on
}

// CrowdsourcingCallerRaw is an auto generated low-level read-only Go binding around an ethereum contract.
type CrowdsourcingCallerRaw struct {
	Contract *CrowdsourcingCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdsourcingTransactorRaw is an auto generated low-level write-only Go binding around an ethereum contract.
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
