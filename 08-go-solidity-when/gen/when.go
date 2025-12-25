// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package when

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WhenMetaData contains all meta data concerning the When contract.
var WhenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040526040518060400160405280600d81526020017f57726170706572204574686572000000000000000000000000000000000000008152505f908161004791906102f3565b506040518060400160405280600481526020017f5748454e000000000000000000000000000000000000000000000000000000008152506001908161008c91906102f3565b50601260025f6101000a81548160ff021916908360ff1602179055503480156100b3575f80fd5b506103c2565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061013457607f821691505b602082108103610147576101466100f0565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101a97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261016e565b6101b3868361016e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101f76101f26101ed846101cb565b6101d4565b6101cb565b9050919050565b5f819050919050565b610210836101dd565b61022461021c826101fe565b84845461017a565b825550505050565b5f90565b61023861022c565b610243818484610207565b505050565b5b818110156102665761025b5f82610230565b600181019050610249565b5050565b601f8211156102ab5761027c8161014d565b6102858461015f565b81016020851015610294578190505b6102a86102a08561015f565b830182610248565b50505b505050565b5f82821c905092915050565b5f6102cb5f19846008026102b0565b1980831691505092915050565b5f6102e383836102bc565b9150826002028217905092915050565b6102fc826100b9565b67ffffffffffffffff811115610315576103146100c3565b5b61031f825461011d565b61032a82828561026a565b5f60209050601f83116001811461035b575f8415610349578287015190505b61035385826102d8565b8655506103ba565b601f1984166103698661014d565b5f5b828110156103905784890151825560018201915060208501945060208101905061036b565b868310156103ad57848901516103a9601f8916826102bc565b8355505b6001600288020188555050505b505050505050565b610f84806103cf5f395ff3fe60806040526004361061009f575f3560e01c8063313ce56711610063578063313ce567146101a657806370a08231146101d057806395d89b411461020c578063a9059cbb14610236578063d0e30db014610272578063dd62ed3e1461027c576100ae565b806306fdde03146100b2578063095ea7b3146100dc57806318160ddd1461011857806323b872dd146101425780632e1a7d4d1461017e576100ae565b366100ae576100ac6102b8565b005b5f80fd5b3480156100bd575f80fd5b506100c661035b565b6040516100d39190610b2d565b60405180910390f35b3480156100e7575f80fd5b5061010260048036038101906100fd9190610bde565b6103e6565b60405161010f9190610c36565b60405180910390f35b348015610123575f80fd5b5061012c6104d3565b6040516101399190610c5e565b60405180910390f35b34801561014d575f80fd5b5061016860048036038101906101639190610c77565b6104da565b6040516101759190610c36565b60405180910390f35b348015610189575f80fd5b506101a4600480360381019061019f9190610cc7565b610826565b005b3480156101b1575f80fd5b506101ba6109bc565b6040516101c79190610d0d565b60405180910390f35b3480156101db575f80fd5b506101f660048036038101906101f19190610d26565b6109ce565b6040516102039190610c5e565b60405180910390f35b348015610217575f80fd5b506102206109e3565b60405161022d9190610b2d565b60405180910390f35b348015610241575f80fd5b5061025c60048036038101906102579190610bde565b610a6f565b6040516102699190610c36565b60405180910390f35b61027a6102b8565b005b348015610287575f80fd5b506102a2600480360381019061029d9190610d51565b610a83565b6040516102af9190610c5e565b60405180910390f35b3460035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546103049190610dbc565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516103519190610c5e565b60405180910390a2565b5f805461036790610e1c565b80601f016020809104026020016040519081016040528092919081815260200182805461039390610e1c565b80156103de5780601f106103b5576101008083540402835291602001916103de565b820191905f5260205f20905b8154815290600101906020018083116103c157829003601f168201915b505050505081565b5f8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104c19190610c5e565b60405180910390a36001905092915050565b5f47905090565b5f8160035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610524575f80fd5b3373ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16141580156105f857507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205414155b15610710578160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610681575f80fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107089190610e4c565b925050819055505b8160035f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825461075c9190610e4c565b925050819055508160035f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546107af9190610dbc565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516108139190610c5e565b60405180910390a3600190509392505050565b8060035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054101561086f575f80fd5b8060035f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546108bb9190610e4c565b925050819055505f3373ffffffffffffffffffffffffffffffffffffffff16826040516108e790610eac565b5f6040518083038185875af1925050503d805f8114610921576040519150601f19603f3d011682016040523d82523d5f602084013e610926565b606091505b505090508061096a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161096190610f0a565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65836040516109b09190610c5e565b60405180910390a25050565b60025f9054906101000a900460ff1681565b6003602052805f5260405f205f915090505481565b600180546109f090610e1c565b80601f0160208091040260200160405190810160405280929190818152602001828054610a1c90610e1c565b8015610a675780601f10610a3e57610100808354040283529160200191610a67565b820191905f5260205f20905b815481529060010190602001808311610a4a57829003601f168201915b505050505081565b5f610a7b3384846104da565b905092915050565b6004602052815f5260405f20602052805f5260405f205f91509150505481565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610ada578082015181840152602081019050610abf565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610aff82610aa3565b610b098185610aad565b9350610b19818560208601610abd565b610b2281610ae5565b840191505092915050565b5f6020820190508181035f830152610b458184610af5565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b7a82610b51565b9050919050565b610b8a81610b70565b8114610b94575f80fd5b50565b5f81359050610ba581610b81565b92915050565b5f819050919050565b610bbd81610bab565b8114610bc7575f80fd5b50565b5f81359050610bd881610bb4565b92915050565b5f8060408385031215610bf457610bf3610b4d565b5b5f610c0185828601610b97565b9250506020610c1285828601610bca565b9150509250929050565b5f8115159050919050565b610c3081610c1c565b82525050565b5f602082019050610c495f830184610c27565b92915050565b610c5881610bab565b82525050565b5f602082019050610c715f830184610c4f565b92915050565b5f805f60608486031215610c8e57610c8d610b4d565b5b5f610c9b86828701610b97565b9350506020610cac86828701610b97565b9250506040610cbd86828701610bca565b9150509250925092565b5f60208284031215610cdc57610cdb610b4d565b5b5f610ce984828501610bca565b91505092915050565b5f60ff82169050919050565b610d0781610cf2565b82525050565b5f602082019050610d205f830184610cfe565b92915050565b5f60208284031215610d3b57610d3a610b4d565b5b5f610d4884828501610b97565b91505092915050565b5f8060408385031215610d6757610d66610b4d565b5b5f610d7485828601610b97565b9250506020610d8585828601610b97565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610dc682610bab565b9150610dd183610bab565b9250828201905080821115610de957610de8610d8f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610e3357607f821691505b602082108103610e4657610e45610def565b5b50919050565b5f610e5682610bab565b9150610e6183610bab565b9250828203905081811115610e7957610e78610d8f565b5b92915050565b5f81905092915050565b50565b5f610e975f83610e7f565b9150610ea282610e89565b5f82019050919050565b5f610eb682610e8c565b9150819050919050565b7f4661696c656420746f2073656e642045746865720000000000000000000000005f82015250565b5f610ef4601483610aad565b9150610eff82610ec0565b602082019050919050565b5f6020820190508181035f830152610f2181610ee8565b905091905056fea2646970667358221220898f2003778370eb51dacceefd28815e91523aece124a36878432a53ad1521b664736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// WhenABI is the input ABI used to generate the binding from.
// Deprecated: Use WhenMetaData.ABI instead.
var WhenABI = WhenMetaData.ABI

// WhenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WhenMetaData.Bin instead.
var WhenBin = WhenMetaData.Bin

// DeployWhen deploys a new Ethereum contract, binding an instance of When to it.
func DeployWhen(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *When, error) {
	parsed, err := WhenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WhenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &When{WhenCaller: WhenCaller{contract: contract}, WhenTransactor: WhenTransactor{contract: contract}, WhenFilterer: WhenFilterer{contract: contract}}, nil
}

// When is an auto generated Go binding around an Ethereum contract.
type When struct {
	WhenCaller     // Read-only binding to the contract
	WhenTransactor // Write-only binding to the contract
	WhenFilterer   // Log filterer for contract events
}

// WhenCaller is an auto generated read-only Go binding around an Ethereum contract.
type WhenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WhenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WhenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WhenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WhenSession struct {
	Contract     *When             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WhenCallerSession struct {
	Contract *WhenCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WhenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WhenTransactorSession struct {
	Contract     *WhenTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WhenRaw is an auto generated low-level Go binding around an Ethereum contract.
type WhenRaw struct {
	Contract *When // Generic contract binding to access the raw methods on
}

// WhenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WhenCallerRaw struct {
	Contract *WhenCaller // Generic read-only contract binding to access the raw methods on
}

// WhenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WhenTransactorRaw struct {
	Contract *WhenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWhen creates a new instance of When, bound to a specific deployed contract.
func NewWhen(address common.Address, backend bind.ContractBackend) (*When, error) {
	contract, err := bindWhen(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &When{WhenCaller: WhenCaller{contract: contract}, WhenTransactor: WhenTransactor{contract: contract}, WhenFilterer: WhenFilterer{contract: contract}}, nil
}

// NewWhenCaller creates a new read-only instance of When, bound to a specific deployed contract.
func NewWhenCaller(address common.Address, caller bind.ContractCaller) (*WhenCaller, error) {
	contract, err := bindWhen(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WhenCaller{contract: contract}, nil
}

// NewWhenTransactor creates a new write-only instance of When, bound to a specific deployed contract.
func NewWhenTransactor(address common.Address, transactor bind.ContractTransactor) (*WhenTransactor, error) {
	contract, err := bindWhen(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WhenTransactor{contract: contract}, nil
}

// NewWhenFilterer creates a new log filterer instance of When, bound to a specific deployed contract.
func NewWhenFilterer(address common.Address, filterer bind.ContractFilterer) (*WhenFilterer, error) {
	contract, err := bindWhen(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WhenFilterer{contract: contract}, nil
}

// bindWhen binds a generic wrapper to an already deployed contract.
func bindWhen(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WhenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_When *WhenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _When.Contract.WhenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_When *WhenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _When.Contract.WhenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_When *WhenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _When.Contract.WhenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_When *WhenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _When.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_When *WhenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _When.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_When *WhenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _When.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_When *WhenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _When.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_When *WhenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _When.Contract.Allowance(&_When.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_When *WhenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _When.Contract.Allowance(&_When.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_When *WhenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _When.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_When *WhenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _When.Contract.BalanceOf(&_When.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_When *WhenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _When.Contract.BalanceOf(&_When.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_When *WhenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _When.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_When *WhenSession) Decimals() (uint8, error) {
	return _When.Contract.Decimals(&_When.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_When *WhenCallerSession) Decimals() (uint8, error) {
	return _When.Contract.Decimals(&_When.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_When *WhenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _When.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_When *WhenSession) Name() (string, error) {
	return _When.Contract.Name(&_When.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_When *WhenCallerSession) Name() (string, error) {
	return _When.Contract.Name(&_When.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_When *WhenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _When.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_When *WhenSession) Symbol() (string, error) {
	return _When.Contract.Symbol(&_When.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_When *WhenCallerSession) Symbol() (string, error) {
	return _When.Contract.Symbol(&_When.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_When *WhenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _When.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_When *WhenSession) TotalSupply() (*big.Int, error) {
	return _When.Contract.TotalSupply(&_When.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_When *WhenCallerSession) TotalSupply() (*big.Int, error) {
	return _When.Contract.TotalSupply(&_When.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_When *WhenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_When *WhenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.Approve(&_When.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_When *WhenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.Approve(&_When.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_When *WhenTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _When.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_When *WhenSession) Deposit() (*types.Transaction, error) {
	return _When.Contract.Deposit(&_When.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_When *WhenTransactorSession) Deposit() (*types.Transaction, error) {
	return _When.Contract.Deposit(&_When.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_When *WhenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_When *WhenSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.Transfer(&_When.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_When *WhenTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.Transfer(&_When.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_When *WhenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_When *WhenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.TransferFrom(&_When.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_When *WhenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.TransferFrom(&_When.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_When *WhenTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _When.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_When *WhenSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.Withdraw(&_When.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_When *WhenTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _When.Contract.Withdraw(&_When.TransactOpts, wad)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_When *WhenTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _When.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_When *WhenSession) Receive() (*types.Transaction, error) {
	return _When.Contract.Receive(&_When.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_When *WhenTransactorSession) Receive() (*types.Transaction, error) {
	return _When.Contract.Receive(&_When.TransactOpts)
}

// WhenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the When contract.
type WhenApprovalIterator struct {
	Event *WhenApproval // Event containing the contract specifics and raw log

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
func (it *WhenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhenApproval)
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
		it.Event = new(WhenApproval)
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
func (it *WhenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhenApproval represents a Approval event raised by the When contract.
type WhenApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_When *WhenFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WhenApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _When.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WhenApprovalIterator{contract: _When.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_When *WhenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WhenApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _When.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhenApproval)
				if err := _When.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_When *WhenFilterer) ParseApproval(log types.Log) (*WhenApproval, error) {
	event := new(WhenApproval)
	if err := _When.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhenDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the When contract.
type WhenDepositIterator struct {
	Event *WhenDeposit // Event containing the contract specifics and raw log

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
func (it *WhenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhenDeposit)
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
		it.Event = new(WhenDeposit)
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
func (it *WhenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhenDeposit represents a Deposit event raised by the When contract.
type WhenDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_When *WhenFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WhenDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _When.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WhenDepositIterator{contract: _When.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_When *WhenFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WhenDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _When.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhenDeposit)
				if err := _When.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_When *WhenFilterer) ParseDeposit(log types.Log) (*WhenDeposit, error) {
	event := new(WhenDeposit)
	if err := _When.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the When contract.
type WhenTransferIterator struct {
	Event *WhenTransfer // Event containing the contract specifics and raw log

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
func (it *WhenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhenTransfer)
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
		it.Event = new(WhenTransfer)
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
func (it *WhenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhenTransfer represents a Transfer event raised by the When contract.
type WhenTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_When *WhenFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WhenTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _When.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WhenTransferIterator{contract: _When.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_When *WhenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WhenTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _When.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhenTransfer)
				if err := _When.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_When *WhenFilterer) ParseTransfer(log types.Log) (*WhenTransfer, error) {
	event := new(WhenTransfer)
	if err := _When.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WhenWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the When contract.
type WhenWithdrawalIterator struct {
	Event *WhenWithdrawal // Event containing the contract specifics and raw log

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
func (it *WhenWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WhenWithdrawal)
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
		it.Event = new(WhenWithdrawal)
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
func (it *WhenWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WhenWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WhenWithdrawal represents a Withdrawal event raised by the When contract.
type WhenWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_When *WhenFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WhenWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _When.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WhenWithdrawalIterator{contract: _When.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_When *WhenFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WhenWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _When.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WhenWithdrawal)
				if err := _When.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_When *WhenFilterer) ParseWithdrawal(log types.Log) (*WhenWithdrawal, error) {
	event := new(WhenWithdrawal)
	if err := _When.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
