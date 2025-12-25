// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contnet\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112768061005c5f395ff3fe608060405234801561000f575f80fd5b506004361061007b575f3560e01c80638da5cb5b116100595780638da5cb5b146100ea5780639507d39a14610108578063b0c8f9dc14610138578063f745630f146101545761007b565b80630f560cd71461007f5780634cc822151461009d5780638d977672146100b9575b5f80fd5b610087610170565b60405161009491906109b0565b60405180910390f35b6100b760048036038101906100b29190610a14565b6102cc565b005b6100d360048036038101906100ce9190610a14565b61042c565b6040516100e1929190610a96565b60405180910390f35b6100f26104ec565b6040516100ff9190610b03565b60405180910390f35b610122600480360381019061011d9190610a14565b610511565b60405161012f9190610b56565b60405180910390f35b610152600480360381019061014d9190610ca2565b61064b565b005b61016e60048036038101906101699190610ce9565b610719565b005b60603373ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146101ca575f80fd5b5f805480602002602001604051908101604052809291908181526020015f905b828210156102c3578382905f5260205f2090600202016040518060400160405290815f8201805461021a90610d70565b80601f016020809104026020016040519081016040528092919081815260200182805461024690610d70565b80156102915780601f1061026857610100808354040283529160200191610291565b820191905f5260205f20905b81548152906001019060200180831161027457829003601f168201915b50505050508152602001600182015f9054906101000a900460ff161515151581525050815260200190600101906101ea565b50505050905090565b3373ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610324575f80fd5b5f8190505b60015f8054905061033a9190610dcd565b8110156103df575f60018261034f9190610e00565b815481106103605761035f610e33565b5b905f5260205f2090600202015f828154811061037f5761037e610e33565b5b905f5260205f2090600202015f8201815f01908161039d9190611012565b50600182015f9054906101000a900460ff16816001015f6101000a81548160ff02191690831515021790555090505080806103d7906110f7565b915050610329565b505f8054806103f1576103f061111e565b5b600190038181905f5260205f2090600202015f8082015f61041291906107a4565b600182015f6101000a81549060ff02191690555050905550565b5f818154811061043a575f80fd5b905f5260205f2090600202015f91509050805f01805461045990610d70565b80601f016020809104026020016040519081016040528092919081815260200182805461048590610d70565b80156104d05780601f106104a7576101008083540402835291602001916104d0565b820191905f5260205f20905b8154815290600101906020018083116104b357829003601f168201915b505050505090806001015f9054906101000a900460ff16905082565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6105196107e1565b3373ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610571575f80fd5b5f828154811061058457610583610e33565b5b905f5260205f2090600202016040518060400160405290815f820180546105aa90610d70565b80601f01602080910402602001604051908101604052809291908181526020018280546105d690610d70565b80156106215780601f106105f857610100808354040283529160200191610621565b820191905f5260205f20905b81548152906001019060200180831161060457829003601f168201915b50505050508152602001600182015f9054906101000a900460ff1615151515815250509050919050565b3373ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106a3575f80fd5b5f60405180604001604052808381526020015f1515815250908060018154018082558091505060019003905f5260205f2090600202015f909190919091505f820151815f0190816106f4919061114b565b506020820151816001015f6101000a81548160ff021916908315150217905550505050565b3373ffffffffffffffffffffffffffffffffffffffff1660015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610771575f80fd5b805f838154811061078557610784610e33565b5b905f5260205f2090600202015f01908161079f919061114b565b505050565b5080546107b090610d70565b5f825580601f106107c157506107de565b601f0160209004905f5260205f20908101906107dd91906107fc565b5b50565b6040518060400160405280606081526020015f151581525090565b5b80821115610813575f815f9055506001016107fd565b5090565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b8381101561087757808201518184015260208101905061085c565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61089c82610840565b6108a6818561084a565b93506108b681856020860161085a565b6108bf81610882565b840191505092915050565b5f8115159050919050565b6108de816108ca565b82525050565b5f604083015f8301518482035f8601526108fe8282610892565b915050602083015161091360208601826108d5565b508091505092915050565b5f61092983836108e4565b905092915050565b5f602082019050919050565b5f61094782610817565b6109518185610821565b93508360208202850161096385610831565b805f5b8581101561099e578484038952815161097f858261091e565b945061098a83610931565b925060208a01995050600181019050610966565b50829750879550505050505092915050565b5f6020820190508181035f8301526109c8818461093d565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6109f3816109e1565b81146109fd575f80fd5b50565b5f81359050610a0e816109ea565b92915050565b5f60208284031215610a2957610a286109d9565b5b5f610a3684828501610a00565b91505092915050565b5f82825260208201905092915050565b5f610a5982610840565b610a638185610a3f565b9350610a7381856020860161085a565b610a7c81610882565b840191505092915050565b610a90816108ca565b82525050565b5f6040820190508181035f830152610aae8185610a4f565b9050610abd6020830184610a87565b9392505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610aed82610ac4565b9050919050565b610afd81610ae3565b82525050565b5f602082019050610b165f830184610af4565b92915050565b5f604083015f8301518482035f860152610b368282610892565b9150506020830151610b4b60208601826108d5565b508091505092915050565b5f6020820190508181035f830152610b6e8184610b1c565b905092915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610bb482610882565b810181811067ffffffffffffffff82111715610bd357610bd2610b7e565b5b80604052505050565b5f610be56109d0565b9050610bf18282610bab565b919050565b5f67ffffffffffffffff821115610c1057610c0f610b7e565b5b610c1982610882565b9050602081019050919050565b828183375f83830152505050565b5f610c46610c4184610bf6565b610bdc565b905082815260208101848484011115610c6257610c61610b7a565b5b610c6d848285610c26565b509392505050565b5f82601f830112610c8957610c88610b76565b5b8135610c99848260208601610c34565b91505092915050565b5f60208284031215610cb757610cb66109d9565b5b5f82013567ffffffffffffffff811115610cd457610cd36109dd565b5b610ce084828501610c75565b91505092915050565b5f8060408385031215610cff57610cfe6109d9565b5b5f610d0c85828601610a00565b925050602083013567ffffffffffffffff811115610d2d57610d2c6109dd565b5b610d3985828601610c75565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610d8757607f821691505b602082108103610d9a57610d99610d43565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610dd7826109e1565b9150610de2836109e1565b9250828203905081811115610dfa57610df9610da0565b5b92915050565b5f610e0a826109e1565b9150610e15836109e1565b9250828201905080821115610e2d57610e2c610da0565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81549050610e6e81610d70565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610ed17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e96565b610edb8683610e96565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610f16610f11610f0c846109e1565b610ef3565b6109e1565b9050919050565b5f819050919050565b610f2f83610efc565b610f43610f3b82610f1d565b848454610ea2565b825550505050565b5f90565b610f57610f4b565b610f62818484610f26565b505050565b5b81811015610f8557610f7a5f82610f4f565b600181019050610f68565b5050565b601f821115610fca57610f9b81610e75565b610fa484610e87565b81016020851015610fb3578190505b610fc7610fbf85610e87565b830182610f67565b50505b505050565b5f82821c905092915050565b5f610fea5f1984600802610fcf565b1980831691505092915050565b5f6110028383610fdb565b9150826002028217905092915050565b8181036110205750506110f5565b61102982610e60565b67ffffffffffffffff81111561104257611041610b7e565b5b61104c8254610d70565b611057828285610f89565b5f601f831160018114611084575f8415611072578287015490505b61107c8582610ff7565b8655506110ee565b601f19841661109287610e75565b965061109d86610e75565b5f5b828110156110c45784890154825560018201915060018501945060208101905061109f565b868310156110e157848901546110dd601f891682610fdb565b8355505b6001600288020188555050505b5050505050505b565b5f611101826109e1565b91505f820361111357611112610da0565b5b600182039050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b61115482610840565b67ffffffffffffffff81111561116d5761116c610b7e565b5b6111778254610d70565b611182828285610f89565b5f60209050601f8311600181146111b3575f84156111a1578287015190505b6111ab8582610ff7565b865550611212565b601f1984166111c186610e75565b5f5b828110156111e8578489015182556001820191506020850194506020810190506111c3565b868310156112055784890151611201601f891682610fdb565b8355505b6001600288020188555050505b50505050505056fea26469706673582212208010b9a0e85c651117b2e77985f1a2c448ad9f16ff7d5627df00462943520a7f64736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, _id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", _id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(_id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, _id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Todo *TodoCallerSession) Owner() (common.Address, error) {
	return _Todo.Contract.Owner(&_Todo.CallOpts)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string content, bool status)
func (_Todo *TodoCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Content string
	Status  bool
}, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Content string
		Status  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Content = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string content, bool status)
func (_Todo *TodoSession) Tasks(arg0 *big.Int) (struct {
	Content string
	Status  bool
}, error) {
	return _Todo.Contract.Tasks(&_Todo.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(string content, bool status)
func (_Todo *TodoCallerSession) Tasks(arg0 *big.Int) (struct {
	Content string
	Status  bool
}, error) {
	return _Todo.Contract.Tasks(&_Todo.CallOpts, arg0)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _contnet) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, _contnet string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", _contnet)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _contnet) returns()
func (_Todo *TodoSession) Add(_contnet string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _contnet)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string _contnet) returns()
func (_Todo *TodoTransactorSession) Add(_contnet string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, _contnet)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 _id) returns()
func (_Todo *TodoTransactorSession) Remove(_id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, _id)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, _id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}

// Update is a paid mutator transaction binding the contract method 0xf745630f.
//
// Solidity: function update(uint256 _id, string _content) returns()
func (_Todo *TodoTransactorSession) Update(_id *big.Int, _content string) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, _id, _content)
}
