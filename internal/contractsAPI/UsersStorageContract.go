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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UserCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UserDeleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"UsersCreate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UsersDelete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UsersGetAll\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structUsersStorage.User[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"}],\"name\":\"UsersGetOne\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"id\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"internalType\":\"structUsersStorage.User\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052600060025534801561001557600080fd5b50610b78806100256000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063161fd7221461005157806362c5aa8c14610066578063bafe4c0c14610084578063caacd139146100a4575b600080fd5b61006461005f36600461087f565b6100b7565b005b61006e610329565b60405161007b919061091a565b60405180910390f35b61009761009236600461087f565b61042f565b60405161007b919061097c565b6100646100b23660046109a5565b610571565b6001600160401b0381166000908152600160208190526040909120015460ff166100fc5760405162461bcd60e51b81526004016100f390610a55565b60405180910390fd5b6001600160401b0381166000908152600160208190526040822054915411156102845760008054819061013190600190610aac565b8154811061014157610141610ac3565b60009182526020918290206040805180820190915260029092020180546001600160401b03168252600181018054929391929184019161018090610ad9565b80601f01602080910402602001604051908101604052809291908181526020018280546101ac90610ad9565b80156101f95780601f106101ce576101008083540402835291602001916101f9565b820191906000526020600020905b8154815290600101906020018083116101dc57829003601f168201915b5050505050815250509050806000838154811061021857610218610ac3565b6000918252602091829020835160029290920201805467ffffffffffffffff19166001600160401b039092169190911781558282015180519192610264926001850192909101906107a9565b505090516001600160401b03166000908152600160205260409020829055505b600080548061029557610295610b13565b600082815260208120600260001990930192830201805467ffffffffffffffff19168155906102c7600183018261082d565b505090556001600160401b038216600081815260016020818152604080842093845592909101805460ff1916905590519182527f226f7feb839e0ecfde3e26fc44b14b87e353b88579c11cedf82dfdc2013aa7f1910160405180910390a15050565b60606000805480602002602001604051908101604052809291908181526020016000905b82821015610426576000848152602090819020604080518082019091526002850290910180546001600160401b03168252600181018054929391929184019161039590610ad9565b80601f01602080910402602001604051908101604052809291908181526020018280546103c190610ad9565b801561040e5780601f106103e35761010080835404028352916020019161040e565b820191906000526020600020905b8154815290600101906020018083116103f157829003601f168201915b5050505050815250508152602001906001019061034d565b50505050905090565b6040805180820190915260008152606060208201526001600160401b0382166000908152600160208190526040909120015460ff166104805760405162461bcd60e51b81526004016100f390610a55565b6001600160401b038216600090815260016020526040812054815481106104a9576104a9610ac3565b60009182526020918290206040805180820190915260029092020180546001600160401b0316825260018101805492939192918401916104e890610ad9565b80601f016020809104026020016040519081016040528092919081815260200182805461051490610ad9565b80156105615780601f1061053657610100808354040283529160200191610561565b820191906000526020600020905b81548152906001019060200180831161054457829003601f168201915b5050505050815250509050919050565b60008061057c610746565b90505b6001600160401b0381166000908152600160208190526040909120015460ff1615610610576105ac610746565b9050600a82106105fe5760405162461bcd60e51b815260206004820152601c60248201527f756e61626c6520746f2067656e657261746520756e697175652049440000000060448201526064016100f3565b8161060881610b29565b92505061057f565b604080518082019091526001600160401b03828116825260208083018681526000805460018101825590805284517f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5636002909202918201805467ffffffffffffffff191691909516178455905180516106b2937f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5649093019291909101906107a9565b505050604051806040016040528060016000805490506106d29190610aac565b8152600160209182018190526001600160401b0384166000818152828452604090819020855181559484015194909201805460ff191694151594909417909355519182527f654a5f371dd267582fdba132709448f256a549360e2ce54ccb3699d3b8ed2394910160405180910390a1505050565b600280546000918261075783610b29565b909155505060025460408051426020808301919091523360601b6bffffffffffffffffffffffff1916828401526054808301949094528251808303909401845260749091019091528151910120919050565b8280546107b590610ad9565b90600052602060002090601f0160209004810192826107d7576000855561081d565b82601f106107f057805160ff191683800117855561081d565b8280016001018555821561081d579182015b8281111561081d578251825591602001919060010190610802565b5061082992915061086a565b5090565b50805461083990610ad9565b6000825580601f10610849575050565b601f016020900490600052602060002090810190610867919061086a565b50565b5b80821115610829576000815560010161086b565b60006020828403121561089157600080fd5b81356001600160401b03811681146108a857600080fd5b9392505050565b6001600160401b0381511682526000602080830151604082860152805180604087015260005b818110156108f1578281018401518782016060015283016108d5565b81811115610903576000606083890101525b50601f01601f191694909401606001949350505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b8281101561096f57603f1988860301845261095d8583516108af565b94509285019290850190600101610941565b5092979650505050505050565b6020815260006108a860208301846108af565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156109b757600080fd5b81356001600160401b03808211156109ce57600080fd5b818401915084601f8301126109e257600080fd5b8135818111156109f4576109f461098f565b604051601f8201601f19908116603f01168101908382118183101715610a1c57610a1c61098f565b81604052828152876020848701011115610a3557600080fd5b826020860160208301376000928101602001929092525095945050505050565b60208082526021908201527f7265636f7264207769746820676976656e204944206973206e6f7420657869736040820152601d60fa1b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b600082821015610abe57610abe610a96565b500390565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680610aed57607f821691505b602082108103610b0d57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603160045260246000fd5b600060018201610b3b57610b3b610a96565b506001019056fea2646970667358221220c0d6aba1f63cf107fa49dcee8fd5462dbc2d7bb14ea8c2e48239929642ee013764736f6c634300080e0033",
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

// ApiUserCreatedIterator is returned from FilterUserCreated and is used to iterate over the raw logs and unpacked data for UserCreated events raised by the Api contract.
type ApiUserCreatedIterator struct {
	Event *ApiUserCreated // Event containing the contract specifics and raw log

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
func (it *ApiUserCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiUserCreated)
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
		it.Event = new(ApiUserCreated)
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
func (it *ApiUserCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiUserCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiUserCreated represents a UserCreated event raised by the Api contract.
type ApiUserCreated struct {
	Id  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUserCreated is a free log retrieval operation binding the contract event 0x654a5f371dd267582fdba132709448f256a549360e2ce54ccb3699d3b8ed2394.
//
// Solidity: event UserCreated(uint64 id)
func (_Api *ApiFilterer) FilterUserCreated(opts *bind.FilterOpts) (*ApiUserCreatedIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "UserCreated")
	if err != nil {
		return nil, err
	}
	return &ApiUserCreatedIterator{contract: _Api.contract, event: "UserCreated", logs: logs, sub: sub}, nil
}

// WatchUserCreated is a free log subscription operation binding the contract event 0x654a5f371dd267582fdba132709448f256a549360e2ce54ccb3699d3b8ed2394.
//
// Solidity: event UserCreated(uint64 id)
func (_Api *ApiFilterer) WatchUserCreated(opts *bind.WatchOpts, sink chan<- *ApiUserCreated) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "UserCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiUserCreated)
				if err := _Api.contract.UnpackLog(event, "UserCreated", log); err != nil {
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

// ParseUserCreated is a log parse operation binding the contract event 0x654a5f371dd267582fdba132709448f256a549360e2ce54ccb3699d3b8ed2394.
//
// Solidity: event UserCreated(uint64 id)
func (_Api *ApiFilterer) ParseUserCreated(log types.Log) (*ApiUserCreated, error) {
	event := new(ApiUserCreated)
	if err := _Api.contract.UnpackLog(event, "UserCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiUserDeletedIterator is returned from FilterUserDeleted and is used to iterate over the raw logs and unpacked data for UserDeleted events raised by the Api contract.
type ApiUserDeletedIterator struct {
	Event *ApiUserDeleted // Event containing the contract specifics and raw log

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
func (it *ApiUserDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiUserDeleted)
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
		it.Event = new(ApiUserDeleted)
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
func (it *ApiUserDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiUserDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiUserDeleted represents a UserDeleted event raised by the Api contract.
type ApiUserDeleted struct {
	Id  uint64
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUserDeleted is a free log retrieval operation binding the contract event 0x226f7feb839e0ecfde3e26fc44b14b87e353b88579c11cedf82dfdc2013aa7f1.
//
// Solidity: event UserDeleted(uint64 id)
func (_Api *ApiFilterer) FilterUserDeleted(opts *bind.FilterOpts) (*ApiUserDeletedIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "UserDeleted")
	if err != nil {
		return nil, err
	}
	return &ApiUserDeletedIterator{contract: _Api.contract, event: "UserDeleted", logs: logs, sub: sub}, nil
}

// WatchUserDeleted is a free log subscription operation binding the contract event 0x226f7feb839e0ecfde3e26fc44b14b87e353b88579c11cedf82dfdc2013aa7f1.
//
// Solidity: event UserDeleted(uint64 id)
func (_Api *ApiFilterer) WatchUserDeleted(opts *bind.WatchOpts, sink chan<- *ApiUserDeleted) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "UserDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiUserDeleted)
				if err := _Api.contract.UnpackLog(event, "UserDeleted", log); err != nil {
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

// ParseUserDeleted is a log parse operation binding the contract event 0x226f7feb839e0ecfde3e26fc44b14b87e353b88579c11cedf82dfdc2013aa7f1.
//
// Solidity: event UserDeleted(uint64 id)
func (_Api *ApiFilterer) ParseUserDeleted(log types.Log) (*ApiUserDeleted, error) {
	event := new(ApiUserDeleted)
	if err := _Api.contract.UnpackLog(event, "UserDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
