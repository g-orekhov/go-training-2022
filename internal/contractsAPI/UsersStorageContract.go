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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"UsersCreate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structUsersStorage.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UsersDelete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UsersGetAll\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structUsersStorage.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UsersGetOne\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structUsersStorage.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060025534801561001557600080fd5b50610b44806100256000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063161fd7221461005157806362c5aa8c14610079578063bafe4c0c1461008e578063caacd139146100ae575b600080fd5b61006461005f36600461088c565b6100c1565b60405190151581526020015b60405180910390f35b61008161024b565b6040516100709190610927565b6100a161009c36600461088c565b610351565b6040516100709190610989565b6100a16100bc3660046109b2565b610480565b6001600160401b03811660009081526001602081905260408220015460ff166100e957600080fd5b6001600160401b0382166000908152600160208190526040822054825490929161011291610a78565b8154811061012257610122610a8f565b90600052602060002090600202016000828154811061014357610143610a8f565b600091825260209091208254600290920201805467ffffffffffffffff19166001600160401b03909216919091178155600180830180549183019161018790610aa5565b61019292919061073b565b5090505080600160008084815481106101ad576101ad610a8f565b600091825260208083206002909202909101546001600160401b031683528201929092526040018120919091558054806101e9576101e9610adf565b600082815260208120600260001990930192830201805467ffffffffffffffff191681559061021b60018301826107c6565b5050905550506001600160401b031660009081526001602081905260408220918255908101805460ff1916905590565b60606000805480602002602001604051908101604052809291908181526020016000905b82821015610348576000848152602090819020604080518082019091526002850290910180546001600160401b0316825260018101805492939192918401916102b790610aa5565b80601f01602080910402602001604051908101604052809291908181526020018280546102e390610aa5565b80156103305780601f1061030557610100808354040283529160200191610330565b820191906000526020600020905b81548152906001019060200180831161031357829003601f168201915b5050505050815250508152602001906001019061026f565b50505050905090565b6040805180820190915260008152606060208201526001600160401b0382166000908152600160208190526040909120015460ff1661038f57600080fd5b6001600160401b038216600090815260016020526040812054815481106103b8576103b8610a8f565b60009182526020918290206040805180820190915260029092020180546001600160401b0316825260018101805492939192918401916103f790610aa5565b80601f016020809104026020016040519081016040528092919081815260200182805461042390610aa5565b80156104705780601f1061044557610100808354040283529160200191610470565b820191906000526020600020905b81548152906001019060200180831161045357829003601f168201915b5050505050815250509050919050565b6040805180820190915260008152606060208201526000806104a06106d8565b90505b6001600160401b0381166000908152600160208190526040909120015460ff16156104f1576104d06106d8565b9050600a82106104df57600080fd5b816104e981610af5565b9250506104a3565b604080518082019091526001600160401b03828116825260208083018781526000805460018101825590805284517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5636002909202918201805467ffffffffffffffff19169190951617845590518051610593937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e564909301929190910190610803565b505050604051806040016040528060016000805490506105b39190610a78565b8152600160209182018190526001600160401b038416600090815281835260408120845181559390920151928101805460ff191693151593909317909255805490916105fe91610a78565b8154811061060e5761060e610a8f565b60009182526020918290206040805180820190915260029092020180546001600160401b03168252600181018054929391929184019161064d90610aa5565b80601f016020809104026020016040519081016040528092919081815260200182805461067990610aa5565b80156106c65780601f1061069b576101008083540402835291602001916106c6565b820191906000526020600020905b8154815290600101906020018083116106a957829003601f168201915b50505050508152505092505050919050565b60028054600091826106e983610af5565b909155505060025460408051426020808301919091523360601b6bffffffffffffffffffffffff1916828401526054808301949094528251808303909401845260749091019091528151910120919050565b82805461074790610aa5565b90600052602060002090601f01602090048101928261076957600085556107b6565b82601f1061077a57805485556107b6565b828001600101855582156107b657600052602060002091601f016020900482015b828111156107b657825482559160010191906001019061079b565b506107c2929150610877565b5090565b5080546107d290610aa5565b6000825580601f106107e2575050565b601f0160209004906000526020600020908101906108009190610877565b50565b82805461080f90610aa5565b90600052602060002090601f01602090048101928261083157600085556107b6565b82601f1061084a57805160ff19168380011785556107b6565b828001600101855582156107b6579182015b828111156107b657825182559160200191906001019061085c565b5b808211156107c25760008155600101610878565b60006020828403121561089e57600080fd5b81356001600160401b03811681146108b557600080fd5b9392505050565b6001600160401b0381511682526000602080830151604082860152805180604087015260005b818110156108fe578281018401518782016060015283016108e2565b81811115610910576000606083890101525b50601f01601f191694909401606001949350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561097c57603f1988860301845261096a8583516108bc565b9450928501929085019060010161094e565b5092979650505050505050565b6020815260006108b560208301846108bc565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156109c457600080fd5b81356001600160401b03808211156109db57600080fd5b818401915084601f8301126109ef57600080fd5b813581811115610a0157610a0161099c565b604051601f8201601f19908116603f01168101908382118183101715610a2957610a2961099c565b81604052828152876020848701011115610a4257600080fd5b826020860160208301376000928101602001929092525095945050505050565b634e487b7160e01b600052601160045260246000fd5b600082821015610a8a57610a8a610a62565b500390565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680610ab957607f821691505b602082108103610ad957634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b600060018201610b0757610b07610a62565b506001019056fea2646970667358221220bb653cfb87875b530c6621498c60d5d093ee97b08a1c152a095b30254f6d617664736f6c634300080e0033",
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
// Solidity: function UsersCreate(string name) returns((uint64,string))
func (_Api *ApiTransactor) UsersCreate(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "UsersCreate", name)
}

// UsersCreate is a paid mutator transaction binding the contract method 0xcaacd139.
//
// Solidity: function UsersCreate(string name) returns((uint64,string))
func (_Api *ApiSession) UsersCreate(name string) (*types.Transaction, error) {
	return _Api.Contract.UsersCreate(&_Api.TransactOpts, name)
}

// UsersCreate is a paid mutator transaction binding the contract method 0xcaacd139.
//
// Solidity: function UsersCreate(string name) returns((uint64,string))
func (_Api *ApiTransactorSession) UsersCreate(name string) (*types.Transaction, error) {
	return _Api.Contract.UsersCreate(&_Api.TransactOpts, name)
}

// UsersDelete is a paid mutator transaction binding the contract method 0x161fd722.
//
// Solidity: function UsersDelete(uint64 id) returns(bool success)
func (_Api *ApiTransactor) UsersDelete(opts *bind.TransactOpts, id uint64) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "UsersDelete", id)
}

// UsersDelete is a paid mutator transaction binding the contract method 0x161fd722.
//
// Solidity: function UsersDelete(uint64 id) returns(bool success)
func (_Api *ApiSession) UsersDelete(id uint64) (*types.Transaction, error) {
	return _Api.Contract.UsersDelete(&_Api.TransactOpts, id)
}

// UsersDelete is a paid mutator transaction binding the contract method 0x161fd722.
//
// Solidity: function UsersDelete(uint64 id) returns(bool success)
func (_Api *ApiTransactorSession) UsersDelete(id uint64) (*types.Transaction, error) {
	return _Api.Contract.UsersDelete(&_Api.TransactOpts, id)
}
