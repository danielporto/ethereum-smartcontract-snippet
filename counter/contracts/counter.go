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

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PrintConfirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PrintValue\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0af00bc5": "decrement(uint256,string)",
		"8ada066e": "getCounter()",
		"c87b285a": "increment(uint256,string)",
		"64371977": "set(uint256,string)",
	},
	Bin: "0x60806040526000805534801561001457600080fd5b506102bc806100246000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630af00bc51461005157806364371977146100665780638ada066e14610079578063c87b285a14610093575b600080fd5b61006461005f36600461018c565b6100a6565b005b61006461007436600461018c565b6100fe565b610081610137565b60405190815260200160405180910390f35b6100646100a136600461018c565b61017b565b826000808282546100b7919061021e565b90915550506000546040517fa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811916100f19185918591610235565b60405180910390a1505050565b60008390556040517fa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811906100f190849084908790610235565b60007f25576b06e88ffd030b5861892e5f7f6a72701cd1692179c1acf2f89914204f5f60005460405161016c91815260200190565b60405180910390a15060005490565b826000808282546100b7919061026e565b6000806000604084860312156101a157600080fd5b83359250602084013567ffffffffffffffff808211156101c057600080fd5b818601915086601f8301126101d457600080fd5b8135818111156101e357600080fd5b8760208285010111156101f557600080fd5b6020830194508093505050509250925092565b634e487b7160e01b600052601160045260246000fd5b60008282101561023057610230610208565b500390565b604081528260408201528284606083013760006060848301015260006060601f19601f8601168301019050826020830152949350505050565b6000821982111561028157610281610208565b50019056fea2646970667358221220e9f01cd347fe8cdf1074b8ebcae77691a79e4ab196ea84063b4d49da3022219f64736f6c63430008090033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// Deprecated: Use CounterMetaData.Sigs instead.
// CounterFuncSigs maps the 4-byte function signature to its string representation.
var CounterFuncSigs = CounterMetaData.Sigs

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CounterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// Decrement is a paid mutator transaction binding the contract method 0x0af00bc5.
//
// Solidity: function decrement(uint256 value, string id) returns()
func (_Counter *CounterTransactor) Decrement(opts *bind.TransactOpts, value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrement", value, id)
}

// Decrement is a paid mutator transaction binding the contract method 0x0af00bc5.
//
// Solidity: function decrement(uint256 value, string id) returns()
func (_Counter *CounterSession) Decrement(value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts, value, id)
}

// Decrement is a paid mutator transaction binding the contract method 0x0af00bc5.
//
// Solidity: function decrement(uint256 value, string id) returns()
func (_Counter *CounterTransactorSession) Decrement(value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts, value, id)
}

// GetCounter is a paid mutator transaction binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() returns(uint256)
func (_Counter *CounterTransactor) GetCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "getCounter")
}

// GetCounter is a paid mutator transaction binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() returns(uint256)
func (_Counter *CounterSession) GetCounter() (*types.Transaction, error) {
	return _Counter.Contract.GetCounter(&_Counter.TransactOpts)
}

// GetCounter is a paid mutator transaction binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() returns(uint256)
func (_Counter *CounterTransactorSession) GetCounter() (*types.Transaction, error) {
	return _Counter.Contract.GetCounter(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xc87b285a.
//
// Solidity: function increment(uint256 value, string id) returns()
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts, value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment", value, id)
}

// Increment is a paid mutator transaction binding the contract method 0xc87b285a.
//
// Solidity: function increment(uint256 value, string id) returns()
func (_Counter *CounterSession) Increment(value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts, value, id)
}

// Increment is a paid mutator transaction binding the contract method 0xc87b285a.
//
// Solidity: function increment(uint256 value, string id) returns()
func (_Counter *CounterTransactorSession) Increment(value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts, value, id)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 value, string id) returns()
func (_Counter *CounterTransactor) Set(opts *bind.TransactOpts, value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "set", value, id)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 value, string id) returns()
func (_Counter *CounterSession) Set(value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.Contract.Set(&_Counter.TransactOpts, value, id)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 value, string id) returns()
func (_Counter *CounterTransactorSession) Set(value *big.Int, id string) (*types.Transaction, error) {
	return _Counter.Contract.Set(&_Counter.TransactOpts, value, id)
}

// CounterPrintConfirmationIterator is returned from FilterPrintConfirmation and is used to iterate over the raw logs and unpacked data for PrintConfirmation events raised by the Counter contract.
type CounterPrintConfirmationIterator struct {
	Event *CounterPrintConfirmation // Event containing the contract specifics and raw log

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
func (it *CounterPrintConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterPrintConfirmation)
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
		it.Event = new(CounterPrintConfirmation)
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
func (it *CounterPrintConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterPrintConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterPrintConfirmation represents a PrintConfirmation event raised by the Counter contract.
type CounterPrintConfirmation struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintConfirmation is a free log retrieval operation binding the contract event 0xa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811.
//
// Solidity: event PrintConfirmation(string arg0, uint256 arg1)
func (_Counter *CounterFilterer) FilterPrintConfirmation(opts *bind.FilterOpts) (*CounterPrintConfirmationIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "PrintConfirmation")
	if err != nil {
		return nil, err
	}
	return &CounterPrintConfirmationIterator{contract: _Counter.contract, event: "PrintConfirmation", logs: logs, sub: sub}, nil
}

// WatchPrintConfirmation is a free log subscription operation binding the contract event 0xa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811.
//
// Solidity: event PrintConfirmation(string arg0, uint256 arg1)
func (_Counter *CounterFilterer) WatchPrintConfirmation(opts *bind.WatchOpts, sink chan<- *CounterPrintConfirmation) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "PrintConfirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterPrintConfirmation)
				if err := _Counter.contract.UnpackLog(event, "PrintConfirmation", log); err != nil {
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
func (_Counter *CounterFilterer) ParsePrintConfirmation(log types.Log) (*CounterPrintConfirmation, error) {
	event := new(CounterPrintConfirmation)
	if err := _Counter.contract.UnpackLog(event, "PrintConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterPrintValueIterator is returned from FilterPrintValue and is used to iterate over the raw logs and unpacked data for PrintValue events raised by the Counter contract.
type CounterPrintValueIterator struct {
	Event *CounterPrintValue // Event containing the contract specifics and raw log

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
func (it *CounterPrintValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterPrintValue)
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
		it.Event = new(CounterPrintValue)
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
func (it *CounterPrintValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterPrintValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterPrintValue represents a PrintValue event raised by the Counter contract.
type CounterPrintValue struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintValue is a free log retrieval operation binding the contract event 0x25576b06e88ffd030b5861892e5f7f6a72701cd1692179c1acf2f89914204f5f.
//
// Solidity: event PrintValue(uint256 arg0)
func (_Counter *CounterFilterer) FilterPrintValue(opts *bind.FilterOpts) (*CounterPrintValueIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "PrintValue")
	if err != nil {
		return nil, err
	}
	return &CounterPrintValueIterator{contract: _Counter.contract, event: "PrintValue", logs: logs, sub: sub}, nil
}

// WatchPrintValue is a free log subscription operation binding the contract event 0x25576b06e88ffd030b5861892e5f7f6a72701cd1692179c1acf2f89914204f5f.
//
// Solidity: event PrintValue(uint256 arg0)
func (_Counter *CounterFilterer) WatchPrintValue(opts *bind.WatchOpts, sink chan<- *CounterPrintValue) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "PrintValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterPrintValue)
				if err := _Counter.contract.UnpackLog(event, "PrintValue", log); err != nil {
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

// ParsePrintValue is a log parse operation binding the contract event 0x25576b06e88ffd030b5861892e5f7f6a72701cd1692179c1acf2f89914204f5f.
//
// Solidity: event PrintValue(uint256 arg0)
func (_Counter *CounterFilterer) ParsePrintValue(log types.Log) (*CounterPrintValue, error) {
	event := new(CounterPrintValue)
	if err := _Counter.contract.UnpackLog(event, "PrintValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
