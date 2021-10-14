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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Decrement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"GetValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Increment\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3a9ebefd": "decrement(uint256)",
		"8ada066e": "getCounter()",
		"7cf5dab0": "increment(uint256)",
		"60fe47b1": "set(uint256)",
	},
	Bin: "0x60806040526000805534801561001457600080fd5b506101fe806100246000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80633a9ebefd1461005157806360fe47b1146100665780637cf5dab0146100795780638ada066e1461008c575b600080fd5b61006461005f36600461016a565b610094565b005b61006461007436600461016a565b600055565b61006461008736600461016a565b6100e4565b61006461012d565b806000808282546100a5919061019b565b90915550506000546040519081527f32814a5bdfd1b8c3d76c49c54e043d6e8aa93d197a09e16599b567135503f748906020015b60405180910390a150565b806000808282546100f59190610183565b90915550506000546040519081527f51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81906020016100d9565b7ff0282ee95e37b3f7538413a321ff8d12bc72094b9ac8861c5f765f0db377862860005460405161016091815260200190565b60405180910390a1565b60006020828403121561017c57600080fd5b5035919050565b60008219821115610196576101966101b2565b500190565b6000828210156101ad576101ad6101b2565b500390565b634e487b7160e01b600052601160045260246000fdfea264697066735822122052ab3b334c134afb319daa716dac4ffa0d67a40f88789fca02348f6fd9a20f1e64736f6c63430008070033",
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

// Decrement is a paid mutator transaction binding the contract method 0x3a9ebefd.
//
// Solidity: function decrement(uint256 value) returns()
func (_Counter *CounterTransactor) Decrement(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrement", value)
}

// Decrement is a paid mutator transaction binding the contract method 0x3a9ebefd.
//
// Solidity: function decrement(uint256 value) returns()
func (_Counter *CounterSession) Decrement(value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts, value)
}

// Decrement is a paid mutator transaction binding the contract method 0x3a9ebefd.
//
// Solidity: function decrement(uint256 value) returns()
func (_Counter *CounterTransactorSession) Decrement(value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts, value)
}

// GetCounter is a paid mutator transaction binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() returns()
func (_Counter *CounterTransactor) GetCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "getCounter")
}

// GetCounter is a paid mutator transaction binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() returns()
func (_Counter *CounterSession) GetCounter() (*types.Transaction, error) {
	return _Counter.Contract.GetCounter(&_Counter.TransactOpts)
}

// GetCounter is a paid mutator transaction binding the contract method 0x8ada066e.
//
// Solidity: function getCounter() returns()
func (_Counter *CounterTransactorSession) GetCounter() (*types.Transaction, error) {
	return _Counter.Contract.GetCounter(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0x7cf5dab0.
//
// Solidity: function increment(uint256 value) returns()
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment", value)
}

// Increment is a paid mutator transaction binding the contract method 0x7cf5dab0.
//
// Solidity: function increment(uint256 value) returns()
func (_Counter *CounterSession) Increment(value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts, value)
}

// Increment is a paid mutator transaction binding the contract method 0x7cf5dab0.
//
// Solidity: function increment(uint256 value) returns()
func (_Counter *CounterTransactorSession) Increment(value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts, value)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 value) returns()
func (_Counter *CounterTransactor) Set(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "set", value)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 value) returns()
func (_Counter *CounterSession) Set(value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Set(&_Counter.TransactOpts, value)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 value) returns()
func (_Counter *CounterTransactorSession) Set(value *big.Int) (*types.Transaction, error) {
	return _Counter.Contract.Set(&_Counter.TransactOpts, value)
}

// CounterDecrementIterator is returned from FilterDecrement and is used to iterate over the raw logs and unpacked data for Decrement events raised by the Counter contract.
type CounterDecrementIterator struct {
	Event *CounterDecrement // Event containing the contract specifics and raw log

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
func (it *CounterDecrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterDecrement)
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
		it.Event = new(CounterDecrement)
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
func (it *CounterDecrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterDecrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterDecrement represents a Decrement event raised by the Counter contract.
type CounterDecrement struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDecrement is a free log retrieval operation binding the contract event 0x32814a5bdfd1b8c3d76c49c54e043d6e8aa93d197a09e16599b567135503f748.
//
// Solidity: event Decrement(uint256 arg0)
func (_Counter *CounterFilterer) FilterDecrement(opts *bind.FilterOpts) (*CounterDecrementIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "Decrement")
	if err != nil {
		return nil, err
	}
	return &CounterDecrementIterator{contract: _Counter.contract, event: "Decrement", logs: logs, sub: sub}, nil
}

// WatchDecrement is a free log subscription operation binding the contract event 0x32814a5bdfd1b8c3d76c49c54e043d6e8aa93d197a09e16599b567135503f748.
//
// Solidity: event Decrement(uint256 arg0)
func (_Counter *CounterFilterer) WatchDecrement(opts *bind.WatchOpts, sink chan<- *CounterDecrement) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "Decrement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterDecrement)
				if err := _Counter.contract.UnpackLog(event, "Decrement", log); err != nil {
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

// ParseDecrement is a log parse operation binding the contract event 0x32814a5bdfd1b8c3d76c49c54e043d6e8aa93d197a09e16599b567135503f748.
//
// Solidity: event Decrement(uint256 arg0)
func (_Counter *CounterFilterer) ParseDecrement(log types.Log) (*CounterDecrement, error) {
	event := new(CounterDecrement)
	if err := _Counter.contract.UnpackLog(event, "Decrement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterGetValueIterator is returned from FilterGetValue and is used to iterate over the raw logs and unpacked data for GetValue events raised by the Counter contract.
type CounterGetValueIterator struct {
	Event *CounterGetValue // Event containing the contract specifics and raw log

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
func (it *CounterGetValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterGetValue)
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
		it.Event = new(CounterGetValue)
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
func (it *CounterGetValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterGetValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterGetValue represents a GetValue event raised by the Counter contract.
type CounterGetValue struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterGetValue is a free log retrieval operation binding the contract event 0xf0282ee95e37b3f7538413a321ff8d12bc72094b9ac8861c5f765f0db3778628.
//
// Solidity: event GetValue(uint256 arg0)
func (_Counter *CounterFilterer) FilterGetValue(opts *bind.FilterOpts) (*CounterGetValueIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "GetValue")
	if err != nil {
		return nil, err
	}
	return &CounterGetValueIterator{contract: _Counter.contract, event: "GetValue", logs: logs, sub: sub}, nil
}

// WatchGetValue is a free log subscription operation binding the contract event 0xf0282ee95e37b3f7538413a321ff8d12bc72094b9ac8861c5f765f0db3778628.
//
// Solidity: event GetValue(uint256 arg0)
func (_Counter *CounterFilterer) WatchGetValue(opts *bind.WatchOpts, sink chan<- *CounterGetValue) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "GetValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterGetValue)
				if err := _Counter.contract.UnpackLog(event, "GetValue", log); err != nil {
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

// ParseGetValue is a log parse operation binding the contract event 0xf0282ee95e37b3f7538413a321ff8d12bc72094b9ac8861c5f765f0db3778628.
//
// Solidity: event GetValue(uint256 arg0)
func (_Counter *CounterFilterer) ParseGetValue(log types.Log) (*CounterGetValue, error) {
	event := new(CounterGetValue)
	if err := _Counter.contract.UnpackLog(event, "GetValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterIncrementIterator is returned from FilterIncrement and is used to iterate over the raw logs and unpacked data for Increment events raised by the Counter contract.
type CounterIncrementIterator struct {
	Event *CounterIncrement // Event containing the contract specifics and raw log

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
func (it *CounterIncrementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterIncrement)
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
		it.Event = new(CounterIncrement)
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
func (it *CounterIncrementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterIncrementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterIncrement represents a Increment event raised by the Counter contract.
type CounterIncrement struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIncrement is a free log retrieval operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 arg0)
func (_Counter *CounterFilterer) FilterIncrement(opts *bind.FilterOpts) (*CounterIncrementIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return &CounterIncrementIterator{contract: _Counter.contract, event: "Increment", logs: logs, sub: sub}, nil
}

// WatchIncrement is a free log subscription operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 arg0)
func (_Counter *CounterFilterer) WatchIncrement(opts *bind.WatchOpts, sink chan<- *CounterIncrement) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "Increment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterIncrement)
				if err := _Counter.contract.UnpackLog(event, "Increment", log); err != nil {
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

// ParseIncrement is a log parse operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 arg0)
func (_Counter *CounterFilterer) ParseIncrement(log types.Log) (*CounterIncrement, error) {
	event := new(CounterIncrement)
	if err := _Counter.contract.UnpackLog(event, "Increment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
