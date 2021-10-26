// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// KVstoreMetaData contains all meta data concerning the KVstore contract.
var KVstoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PrintConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PrintInserts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"PrintKVAck\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"PrintTotalInserts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"643e1fee": "PrintTotalInserts()",
		"693ec85e": "get(string)",
		"da465d74": "set(string,string,string)",
	},
	Bin: "0x6080604052600060015534801561001557600080fd5b50610584806100256000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063643e1fee14610046578063693ec85e14610050578063da465d7414610079575b600080fd5b61004e61008c565b005b61006361005e366004610342565b6100c9565b60405161007091906103af565b60405180910390f35b61004e6100873660046103e2565b610179565b7f09fa0639e8315051b795e692c47314cab4d06c7baa13ae2a68d0669cea279a526001546040516100bf91815260200190565b60405180910390a1565b60606000826040516100db9190610498565b908152602001604051809103902080546100f4906104b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610120906104b4565b801561016d5780601f106101425761010080835404028352916020019161016d565b820191906000526020600020905b81548152906001019060200180831161015057829003601f168201915b50505050509050919050565b8260008560405161018a9190610498565b908152602001604051809103902090805190602001906101ab929190610206565b5060018060008282546101be91906104ef565b90915550506001546040517fa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811916101f89185918591610515565b60405180910390a150505050565b828054610212906104b4565b90600052602060002090601f016020900481019282610234576000855561027a565b82601f1061024d57805160ff191683800117855561027a565b8280016001018555821561027a579182015b8281111561027a57825182559160200191906001019061025f565b5061028692915061028a565b5090565b5b80821115610286576000815560010161028b565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126102c657600080fd5b813567ffffffffffffffff808211156102e1576102e161029f565b604051601f8301601f19908116603f011681019082821181831017156103095761030961029f565b8160405283815286602085880101111561032257600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561035457600080fd5b813567ffffffffffffffff81111561036b57600080fd5b610377848285016102b5565b949350505050565b60005b8381101561039a578181015183820152602001610382565b838111156103a9576000848401525b50505050565b60208152600082518060208401526103ce81604085016020870161037f565b601f01601f19169190910160400192915050565b600080600080606085870312156103f857600080fd5b843567ffffffffffffffff8082111561041057600080fd5b61041c888389016102b5565b9550602087013591508082111561043257600080fd5b61043e888389016102b5565b9450604087013591508082111561045457600080fd5b818701915087601f83011261046857600080fd5b81358181111561047757600080fd5b88602082850101111561048957600080fd5b95989497505060200194505050565b600082516104aa81846020870161037f565b9190910192915050565b600181811c908216806104c857607f821691505b602082108114156104e957634e487b7160e01b600052602260045260246000fd5b50919050565b6000821982111561051057634e487b7160e01b600052601160045260246000fd5b500190565b604081528260408201528284606083013760006060848301015260006060601f19601f860116830101905082602083015294935050505056fea26469706673582212204c1c0b4379c104e6db788d63a2b6f92627dd5825feb0b86cd7e7a99ad5097dd664736f6c63430008090033",
}

// KVstoreABI is the input ABI used to generate the binding from.
// Deprecated: Use KVstoreMetaData.ABI instead.
var KVstoreABI = KVstoreMetaData.ABI

// Deprecated: Use KVstoreMetaData.Sigs instead.
// KVstoreFuncSigs maps the 4-byte function signature to its string representation.
var KVstoreFuncSigs = KVstoreMetaData.Sigs

// KVstoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KVstoreMetaData.Bin instead.
var KVstoreBin = KVstoreMetaData.Bin

// DeployKVstore deploys a new Ethereum contract, binding an instance of KVstore to it.
func DeployKVstore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KVstore, error) {
	parsed, err := KVstoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KVstoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KVstore{KVstoreCaller: KVstoreCaller{contract: contract}, KVstoreTransactor: KVstoreTransactor{contract: contract}, KVstoreFilterer: KVstoreFilterer{contract: contract}}, nil
}

// KVstore is an auto generated Go binding around an Ethereum contract.
type KVstore struct {
	KVstoreCaller     // Read-only binding to the contract
	KVstoreTransactor // Write-only binding to the contract
	KVstoreFilterer   // Log filterer for contract events
}

// KVstoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type KVstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVstoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KVstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVstoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KVstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KVstoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KVstoreSession struct {
	Contract     *KVstore          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KVstoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KVstoreCallerSession struct {
	Contract *KVstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// KVstoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KVstoreTransactorSession struct {
	Contract     *KVstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KVstoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type KVstoreRaw struct {
	Contract *KVstore // Generic contract binding to access the raw methods on
}

// KVstoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KVstoreCallerRaw struct {
	Contract *KVstoreCaller // Generic read-only contract binding to access the raw methods on
}

// KVstoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KVstoreTransactorRaw struct {
	Contract *KVstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKVstore creates a new instance of KVstore, bound to a specific deployed contract.
func NewKVstore(address common.Address, backend bind.ContractBackend) (*KVstore, error) {
	contract, err := bindKVstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KVstore{KVstoreCaller: KVstoreCaller{contract: contract}, KVstoreTransactor: KVstoreTransactor{contract: contract}, KVstoreFilterer: KVstoreFilterer{contract: contract}}, nil
}

// NewKVstoreCaller creates a new read-only instance of KVstore, bound to a specific deployed contract.
func NewKVstoreCaller(address common.Address, caller bind.ContractCaller) (*KVstoreCaller, error) {
	contract, err := bindKVstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KVstoreCaller{contract: contract}, nil
}

// NewKVstoreTransactor creates a new write-only instance of KVstore, bound to a specific deployed contract.
func NewKVstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*KVstoreTransactor, error) {
	contract, err := bindKVstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KVstoreTransactor{contract: contract}, nil
}

// NewKVstoreFilterer creates a new log filterer instance of KVstore, bound to a specific deployed contract.
func NewKVstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*KVstoreFilterer, error) {
	contract, err := bindKVstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KVstoreFilterer{contract: contract}, nil
}

// bindKVstore binds a generic wrapper to an already deployed contract.
func bindKVstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KVstoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVstore *KVstoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KVstore.Contract.KVstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVstore *KVstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVstore.Contract.KVstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVstore *KVstoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVstore.Contract.KVstoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KVstore *KVstoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KVstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KVstore *KVstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KVstore *KVstoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KVstore.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _key) view returns(string)
func (_KVstore *KVstoreCaller) Get(opts *bind.CallOpts, _key string) (string, error) {
	var out []interface{}
	err := _KVstore.contract.Call(opts, &out, "get", _key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _key) view returns(string)
func (_KVstore *KVstoreSession) Get(_key string) (string, error) {
	return _KVstore.Contract.Get(&_KVstore.CallOpts, _key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string _key) view returns(string)
func (_KVstore *KVstoreCallerSession) Get(_key string) (string, error) {
	return _KVstore.Contract.Get(&_KVstore.CallOpts, _key)
}

// PrintTotalInserts is a paid mutator transaction binding the contract method 0x643e1fee.
//
// Solidity: function PrintTotalInserts() returns()
func (_KVstore *KVstoreTransactor) PrintTotalInserts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KVstore.contract.Transact(opts, "PrintTotalInserts")
}

// PrintTotalInserts is a paid mutator transaction binding the contract method 0x643e1fee.
//
// Solidity: function PrintTotalInserts() returns()
func (_KVstore *KVstoreSession) PrintTotalInserts() (*types.Transaction, error) {
	return _KVstore.Contract.PrintTotalInserts(&_KVstore.TransactOpts)
}

// PrintTotalInserts is a paid mutator transaction binding the contract method 0x643e1fee.
//
// Solidity: function PrintTotalInserts() returns()
func (_KVstore *KVstoreTransactorSession) PrintTotalInserts() (*types.Transaction, error) {
	return _KVstore.Contract.PrintTotalInserts(&_KVstore.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0xda465d74.
//
// Solidity: function set(string _key, string _value, string id) returns()
func (_KVstore *KVstoreTransactor) Set(opts *bind.TransactOpts, _key string, _value string, id string) (*types.Transaction, error) {
	return _KVstore.contract.Transact(opts, "set", _key, _value, id)
}

// Set is a paid mutator transaction binding the contract method 0xda465d74.
//
// Solidity: function set(string _key, string _value, string id) returns()
func (_KVstore *KVstoreSession) Set(_key string, _value string, id string) (*types.Transaction, error) {
	return _KVstore.Contract.Set(&_KVstore.TransactOpts, _key, _value, id)
}

// Set is a paid mutator transaction binding the contract method 0xda465d74.
//
// Solidity: function set(string _key, string _value, string id) returns()
func (_KVstore *KVstoreTransactorSession) Set(_key string, _value string, id string) (*types.Transaction, error) {
	return _KVstore.Contract.Set(&_KVstore.TransactOpts, _key, _value, id)
}

// KVstorePrintConfirmationIterator is returned from FilterPrintConfirmation and is used to iterate over the raw logs and unpacked data for PrintConfirmation events raised by the KVstore contract.
type KVstorePrintConfirmationIterator struct {
	Event *KVstorePrintConfirmation // Event containing the contract specifics and raw log

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
func (it *KVstorePrintConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KVstorePrintConfirmation)
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
		it.Event = new(KVstorePrintConfirmation)
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
func (it *KVstorePrintConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KVstorePrintConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KVstorePrintConfirmation represents a PrintConfirmation event raised by the KVstore contract.
type KVstorePrintConfirmation struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintConfirmation is a free log retrieval operation binding the contract event 0xa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811.
//
// Solidity: event PrintConfirmation(string arg0, uint256 arg1)
func (_KVstore *KVstoreFilterer) FilterPrintConfirmation(opts *bind.FilterOpts) (*KVstorePrintConfirmationIterator, error) {

	logs, sub, err := _KVstore.contract.FilterLogs(opts, "PrintConfirmation")
	if err != nil {
		return nil, err
	}
	return &KVstorePrintConfirmationIterator{contract: _KVstore.contract, event: "PrintConfirmation", logs: logs, sub: sub}, nil
}

// WatchPrintConfirmation is a free log subscription operation binding the contract event 0xa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811.
//
// Solidity: event PrintConfirmation(string arg0, uint256 arg1)
func (_KVstore *KVstoreFilterer) WatchPrintConfirmation(opts *bind.WatchOpts, sink chan<- *KVstorePrintConfirmation) (event.Subscription, error) {

	logs, sub, err := _KVstore.contract.WatchLogs(opts, "PrintConfirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KVstorePrintConfirmation)
				if err := _KVstore.contract.UnpackLog(event, "PrintConfirmation", log); err != nil {
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

// ParsePrintConfirmation is a log parse operation binding the contract event 0xa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811.
//
// Solidity: event PrintConfirmation(string arg0, uint256 arg1)
func (_KVstore *KVstoreFilterer) ParsePrintConfirmation(log types.Log) (*KVstorePrintConfirmation, error) {
	event := new(KVstorePrintConfirmation)
	if err := _KVstore.contract.UnpackLog(event, "PrintConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KVstorePrintInsertsIterator is returned from FilterPrintInserts and is used to iterate over the raw logs and unpacked data for PrintInserts events raised by the KVstore contract.
type KVstorePrintInsertsIterator struct {
	Event *KVstorePrintInserts // Event containing the contract specifics and raw log

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
func (it *KVstorePrintInsertsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KVstorePrintInserts)
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
		it.Event = new(KVstorePrintInserts)
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
func (it *KVstorePrintInsertsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KVstorePrintInsertsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KVstorePrintInserts represents a PrintInserts event raised by the KVstore contract.
type KVstorePrintInserts struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintInserts is a free log retrieval operation binding the contract event 0x09fa0639e8315051b795e692c47314cab4d06c7baa13ae2a68d0669cea279a52.
//
// Solidity: event PrintInserts(uint256 arg0)
func (_KVstore *KVstoreFilterer) FilterPrintInserts(opts *bind.FilterOpts) (*KVstorePrintInsertsIterator, error) {

	logs, sub, err := _KVstore.contract.FilterLogs(opts, "PrintInserts")
	if err != nil {
		return nil, err
	}
	return &KVstorePrintInsertsIterator{contract: _KVstore.contract, event: "PrintInserts", logs: logs, sub: sub}, nil
}

// WatchPrintInserts is a free log subscription operation binding the contract event 0x09fa0639e8315051b795e692c47314cab4d06c7baa13ae2a68d0669cea279a52.
//
// Solidity: event PrintInserts(uint256 arg0)
func (_KVstore *KVstoreFilterer) WatchPrintInserts(opts *bind.WatchOpts, sink chan<- *KVstorePrintInserts) (event.Subscription, error) {

	logs, sub, err := _KVstore.contract.WatchLogs(opts, "PrintInserts")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KVstorePrintInserts)
				if err := _KVstore.contract.UnpackLog(event, "PrintInserts", log); err != nil {
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

// ParsePrintInserts is a log parse operation binding the contract event 0x09fa0639e8315051b795e692c47314cab4d06c7baa13ae2a68d0669cea279a52.
//
// Solidity: event PrintInserts(uint256 arg0)
func (_KVstore *KVstoreFilterer) ParsePrintInserts(log types.Log) (*KVstorePrintInserts, error) {
	event := new(KVstorePrintInserts)
	if err := _KVstore.contract.UnpackLog(event, "PrintInserts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KVstorePrintKVAckIterator is returned from FilterPrintKVAck and is used to iterate over the raw logs and unpacked data for PrintKVAck events raised by the KVstore contract.
type KVstorePrintKVAckIterator struct {
	Event *KVstorePrintKVAck // Event containing the contract specifics and raw log

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
func (it *KVstorePrintKVAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KVstorePrintKVAck)
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
		it.Event = new(KVstorePrintKVAck)
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
func (it *KVstorePrintKVAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KVstorePrintKVAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KVstorePrintKVAck represents a PrintKVAck event raised by the KVstore contract.
type KVstorePrintKVAck struct {
	Arg0 *big.Int
	Arg1 string
	Arg2 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintKVAck is a free log retrieval operation binding the contract event 0xdf4afc806a701f112b8f38868c22361c40e980b57f7c0d0506d26355a45a0c4e.
//
// Solidity: event PrintKVAck(uint256 arg0, string arg1, string arg2)
func (_KVstore *KVstoreFilterer) FilterPrintKVAck(opts *bind.FilterOpts) (*KVstorePrintKVAckIterator, error) {

	logs, sub, err := _KVstore.contract.FilterLogs(opts, "PrintKVAck")
	if err != nil {
		return nil, err
	}
	return &KVstorePrintKVAckIterator{contract: _KVstore.contract, event: "PrintKVAck", logs: logs, sub: sub}, nil
}

// WatchPrintKVAck is a free log subscription operation binding the contract event 0xdf4afc806a701f112b8f38868c22361c40e980b57f7c0d0506d26355a45a0c4e.
//
// Solidity: event PrintKVAck(uint256 arg0, string arg1, string arg2)
func (_KVstore *KVstoreFilterer) WatchPrintKVAck(opts *bind.WatchOpts, sink chan<- *KVstorePrintKVAck) (event.Subscription, error) {

	logs, sub, err := _KVstore.contract.WatchLogs(opts, "PrintKVAck")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KVstorePrintKVAck)
				if err := _KVstore.contract.UnpackLog(event, "PrintKVAck", log); err != nil {
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

// ParsePrintKVAck is a log parse operation binding the contract event 0xdf4afc806a701f112b8f38868c22361c40e980b57f7c0d0506d26355a45a0c4e.
//
// Solidity: event PrintKVAck(uint256 arg0, string arg1, string arg2)
func (_KVstore *KVstoreFilterer) ParsePrintKVAck(log types.Log) (*KVstorePrintKVAck, error) {
	event := new(KVstorePrintKVAck)
	if err := _KVstore.contract.UnpackLog(event, "PrintKVAck", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
