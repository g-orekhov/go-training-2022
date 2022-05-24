// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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
)

// UsersStorageUser is an auto generated low-level Go binding around an user-defined struct.
type UsersStorageUser struct {
	Id   uint64
	Name string
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"UsersCreate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UsersDelete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UsersGetAll\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structUsersStorage.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UsersGetOne\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structUsersStorage.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060025534801561001557600080fd5b50610a5b806100256000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063161fd7221461005157806362c5aa8c14610066578063bafe4c0c14610084578063caacd139146100a4575b600080fd5b61006461005f3660046107a3565b6100b7565b005b61006e6102d7565b60405161007b919061083e565b60405180910390f35b6100976100923660046107a3565b6103dd565b60405161007b91906108a0565b6100646100b23660046108c9565b61050c565b6001600160401b0381166000908152600160208190526040909120015460ff166100e057600080fd5b6001600160401b038116600090815260016020819052604082205491541115610268576000805481906101159060019061098f565b81548110610125576101256109a6565b60009182526020918290206040805180820190915260029092020180546001600160401b031682526001810180549293919291840191610164906109bc565b80601f0160208091040260200160405190810160405280929190818152602001828054610190906109bc565b80156101dd5780601f106101b2576101008083540402835291602001916101dd565b820191906000526020600020905b8154815290600101906020018083116101c057829003601f168201915b505050505081525050905080600083815481106101fc576101fc6109a6565b6000918252602091829020835160029290920201805467ffffffffffffffff19166001600160401b039092169190911781558282015180519192610248926001850192909101906106cd565b505090516001600160401b03166000908152600160205260409020829055505b6000805480610279576102796109f6565b600082815260208120600260001990930192830201805467ffffffffffffffff19168155906102ab6001830182610751565b50509055506001600160401b03166000908152600160208190526040822091825501805460ff19169055565b60606000805480602002602001604051908101604052809291908181526020016000905b828210156103d4576000848152602090819020604080518082019091526002850290910180546001600160401b031682526001810180549293919291840191610343906109bc565b80601f016020809104026020016040519081016040528092919081815260200182805461036f906109bc565b80156103bc5780601f10610391576101008083540402835291602001916103bc565b820191906000526020600020905b81548152906001019060200180831161039f57829003601f168201915b505050505081525050815260200190600101906102fb565b50505050905090565b6040805180820190915260008152606060208201526001600160401b0382166000908152600160208190526040909120015460ff1661041b57600080fd5b6001600160401b03821660009081526001602052604081205481548110610444576104446109a6565b60009182526020918290206040805180820190915260029092020180546001600160401b031682526001810180549293919291840191610483906109bc565b80601f01602080910402602001604051908101604052809291908181526020018280546104af906109bc565b80156104fc5780601f106104d1576101008083540402835291602001916104fc565b820191906000526020600020905b8154815290600101906020018083116104df57829003601f168201915b5050505050815250509050919050565b60008061051761066a565b90505b6001600160401b0381166000908152600160208190526040909120015460ff16156105685761054761066a565b9050600a821061055657600080fd5b8161056081610a0c565b92505061051a565b604080518082019091526001600160401b03828116825260208083018681526000805460018101825590805284517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5636002909202918201805467ffffffffffffffff1916919095161784559051805161060a937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5649093019291909101906106cd565b5050506040518060400160405280600160008054905061062a919061098f565b8152600160209182018190526001600160401b0393909316600090815283825260409020825181559101519101805460ff19169115159190911790555050565b600280546000918261067b83610a0c565b909155505060025460408051426020808301919091523360601b6bffffffffffffffffffffffff1916828401526054808301949094528251808303909401845260749091019091528151910120919050565b8280546106d9906109bc565b90600052602060002090601f0160209004810192826106fb5760008555610741565b82601f1061071457805160ff1916838001178555610741565b82800160010185558215610741579182015b82811115610741578251825591602001919060010190610726565b5061074d92915061078e565b5090565b50805461075d906109bc565b6000825580601f1061076d575050565b601f01602090049060005260206000209081019061078b919061078e565b50565b5b8082111561074d576000815560010161078f565b6000602082840312156107b557600080fd5b81356001600160401b03811681146107cc57600080fd5b9392505050565b6001600160401b0381511682526000602080830151604082860152805180604087015260005b81811015610815578281018401518782016060015283016107f9565b81811115610827576000606083890101525b50601f01601f191694909401606001949350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561089357603f198886030184526108818583516107d3565b94509285019290850190600101610865565b5092979650505050505050565b6020815260006107cc60208301846107d3565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156108db57600080fd5b81356001600160401b03808211156108f257600080fd5b818401915084601f83011261090657600080fd5b813581811115610918576109186108b3565b604051601f8201601f19908116603f01168101908382118183101715610940576109406108b3565b8160405282815287602084870101111561095957600080fd5b826020860160208301376000928101602001929092525095945050505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156109a1576109a1610979565b500390565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806109d057607f821691505b6020821081036109f057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b600060018201610a1e57610a1e610979565b506001019056fea2646970667358221220c5077286bf67a2d1f4f0937c23f52bd40edb8ce37c9288b6d6ba60ab07692c9364736f6c634300080e0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// UsersGetAll is a free data retrieval call binding the contract method 0x62c5aa8c.
//
// Solidity: function UsersGetAll() view returns((uint64,string)[])
func (_Api *ApiCaller) UsersGetAll(opts *bind.CallOpts) ([]UsersStorageUser, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "UsersGetAll")

	if err != nil {
		return *new([]UsersStorageUser), err
	}

	out0 := *abi.ConvertType(out[0], new([]UsersStorageUser)).(*[]UsersStorageUser)

	return out0, err

}

// UsersGetAll is a free data retrieval call binding the contract method 0x62c5aa8c.
//
// Solidity: function UsersGetAll() view returns((uint64,string)[])
func (_Api *ApiSession) UsersGetAll() ([]UsersStorageUser, error) {
	return _Api.Contract.UsersGetAll(&_Api.CallOpts)
}

// UsersGetAll is a free data retrieval call binding the contract method 0x62c5aa8c.
//
// Solidity: function UsersGetAll() view returns((uint64,string)[])
func (_Api *ApiCallerSession) UsersGetAll() ([]UsersStorageUser, error) {
	return _Api.Contract.UsersGetAll(&_Api.CallOpts)
}

// UsersGetOne is a free data retrieval call binding the contract method 0xbafe4c0c.
//
// Solidity: function UsersGetOne(uint64 id) view returns((uint64,string))
func (_Api *ApiCaller) UsersGetOne(opts *bind.CallOpts, id uint64) (UsersStorageUser, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "UsersGetOne", id)

	if err != nil {
		return *new(UsersStorageUser), err
	}

	out0 := *abi.ConvertType(out[0], new(UsersStorageUser)).(*UsersStorageUser)

	return out0, err

}

// UsersGetOne is a free data retrieval call binding the contract method 0xbafe4c0c.
//
// Solidity: function UsersGetOne(uint64 id) view returns((uint64,string))
func (_Api *ApiSession) UsersGetOne(id uint64) (UsersStorageUser, error) {
	return _Api.Contract.UsersGetOne(&_Api.CallOpts, id)
}

// UsersGetOne is a free data retrieval call binding the contract method 0xbafe4c0c.
//
// Solidity: function UsersGetOne(uint64 id) view returns((uint64,string))
func (_Api *ApiCallerSession) UsersGetOne(id uint64) (UsersStorageUser, error) {
	return _Api.Contract.UsersGetOne(&_Api.CallOpts, id)
}

// UsersCreate is a paid mutator transaction binding the contract method 0xcaacd139.
//
// Solidity: function UsersCreate(string name) returns()
func (_Api *ApiTransactor) UsersCreate(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "UsersCreate", name)
}

// UsersCreate is a paid mutator transaction binding the contract method 0xcaacd139.
//
// Solidity: function UsersCreate(string name) returns()
func (_Api *ApiSession) UsersCreate(name string) (*types.Transaction, error) {
	return _Api.Contract.UsersCreate(&_Api.TransactOpts, name)
}

// UsersCreate is a paid mutator transaction binding the contract method 0xcaacd139.
//
// Solidity: function UsersCreate(string name) returns()
func (_Api *ApiTransactorSession) UsersCreate(name string) (*types.Transaction, error) {
	return _Api.Contract.UsersCreate(&_Api.TransactOpts, name)
}

// UsersDelete is a paid mutator transaction binding the contract method 0x161fd722.
//
// Solidity: function UsersDelete(uint64 id) returns()
func (_Api *ApiTransactor) UsersDelete(opts *bind.TransactOpts, id uint64) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "UsersDelete", id)
}

// UsersDelete is a paid mutator transaction binding the contract method 0x161fd722.
//
// Solidity: function UsersDelete(uint64 id) returns()
func (_Api *ApiSession) UsersDelete(id uint64) (*types.Transaction, error) {
	return _Api.Contract.UsersDelete(&_Api.TransactOpts, id)
}

// UsersDelete is a paid mutator transaction binding the contract method 0x161fd722.
//
// Solidity: function UsersDelete(uint64 id) returns()
func (_Api *ApiTransactorSession) UsersDelete(id uint64) (*types.Transaction, error) {
	return _Api.Contract.UsersDelete(&_Api.TransactOpts, id)
}
