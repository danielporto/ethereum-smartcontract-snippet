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

// BankMetaData contains all meta data concerning the Bank contract.
var BankMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BalanceReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BalanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"OperationsExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PrintConfirmation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balanceReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logTransferOperations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"payMoneyTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"transferMoneyTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"52a90c42": "balanceReceived()",
		"12065fe0": "getBalance()",
		"f948e50a": "getOperations()",
		"2059ab90": "logTransferOperations()",
		"82664c61": "payMoneyTo(address)",
		"6d26ec18": "receiveMoney()",
		"ec197ed8": "transferMoneyTo(address,string)",
		"ac446002": "withdrawMoney()",
	},
	Bin: "0x60806040526000805534801561001457600080fd5b5061042d806100246000396000f3fe60806040526004361061007b5760003560e01c806382664c611161004e57806382664c61146100d6578063ac446002146100e9578063ec197ed8146100fe578063f948e50a1461011157600080fd5b806312065fe0146100805780632059ab90146100a157806352a90c42146100b85780636d26ec18146100ce575b600080fd5b34801561008c57600080fd5b50475b60405190815260200160405180910390f35b3480156100ad57600080fd5b506100b6610126565b005b3480156100c457600080fd5b5061008f60015481565b6100b6610163565b6100b66100e43660046102f3565b6101aa565b3480156100f557600080fd5b506100b6610216565b6100b661010c366004610315565b610248565b34801561011d57600080fd5b5060005461008f565b7f45b1c2f5c9e32511e6eb1772825abe442d6058ea66dd4e139a8fcc84b7a3c71260005460405161015991815260200190565b60405180910390a1565b34600160008282546101759190610398565b90915550506040513481527fbb27de1f6791c153e5749c6e33c63738bcdf867fb4a778f2e3ea9b01c338eb5f90602001610159565b6040516001600160a01b038216903480156108fc02916000818181858888f193505050501580156101df573d6000803e3d6000fd5b506040513481527f3760d80235c7b3a7167f5f927afb58bcdcb8e3f99c8571aaf8544cdb70f3b3119060200160405180910390a150565b33806108fc476040518115909202916000818181858888f19350505050158015610244573d6000803e3d6000fd5b5050565b6040516001600160a01b038416903480156108fc02916000818181858888f1935050505015801561027d573d6000803e3d6000fd5b5060016000808282546102909190610398565b90915550506000546040517fa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811916102ca91859185916103be565b60405180910390a1505050565b80356001600160a01b03811681146102ee57600080fd5b919050565b60006020828403121561030557600080fd5b61030e826102d7565b9392505050565b60008060006040848603121561032a57600080fd5b610333846102d7565b9250602084013567ffffffffffffffff8082111561035057600080fd5b818601915086601f83011261036457600080fd5b81358181111561037357600080fd5b87602082850101111561038557600080fd5b6020830194508093505050509250925092565b600082198211156103b957634e487b7160e01b600052601160045260246000fd5b500190565b604081528260408201528284606083013760006060848301015260006060601f19601f860116830101905082602083015294935050505056fea2646970667358221220f0d7fcf9e53483b08a21a4aaff4a1377363a239774e50abdf6ae5b877ad4489c64736f6c63430008090033",
}

// BankABI is the input ABI used to generate the binding from.
// Deprecated: Use BankMetaData.ABI instead.
var BankABI = BankMetaData.ABI

// Deprecated: Use BankMetaData.Sigs instead.
// BankFuncSigs maps the 4-byte function signature to its string representation.
var BankFuncSigs = BankMetaData.Sigs

// BankBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BankMetaData.Bin instead.
var BankBin = BankMetaData.Bin

// DeployBank deploys a new Ethereum contract, binding an instance of Bank to it.
func DeployBank(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bank, error) {
	parsed, err := BankMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BankBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// Bank is an auto generated Go binding around an Ethereum contract.
type Bank struct {
	BankCaller     // Read-only binding to the contract
	BankTransactor // Write-only binding to the contract
	BankFilterer   // Log filterer for contract events
}

// BankCaller is an auto generated read-only Go binding around an Ethereum contract.
type BankCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BankTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BankFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BankSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BankSession struct {
	Contract     *Bank             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BankCallerSession struct {
	Contract *BankCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BankTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BankTransactorSession struct {
	Contract     *BankTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BankRaw is an auto generated low-level Go binding around an Ethereum contract.
type BankRaw struct {
	Contract *Bank // Generic contract binding to access the raw methods on
}

// BankCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BankCallerRaw struct {
	Contract *BankCaller // Generic read-only contract binding to access the raw methods on
}

// BankTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BankTransactorRaw struct {
	Contract *BankTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBank creates a new instance of Bank, bound to a specific deployed contract.
func NewBank(address common.Address, backend bind.ContractBackend) (*Bank, error) {
	contract, err := bindBank(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bank{BankCaller: BankCaller{contract: contract}, BankTransactor: BankTransactor{contract: contract}, BankFilterer: BankFilterer{contract: contract}}, nil
}

// NewBankCaller creates a new read-only instance of Bank, bound to a specific deployed contract.
func NewBankCaller(address common.Address, caller bind.ContractCaller) (*BankCaller, error) {
	contract, err := bindBank(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BankCaller{contract: contract}, nil
}

// NewBankTransactor creates a new write-only instance of Bank, bound to a specific deployed contract.
func NewBankTransactor(address common.Address, transactor bind.ContractTransactor) (*BankTransactor, error) {
	contract, err := bindBank(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BankTransactor{contract: contract}, nil
}

// NewBankFilterer creates a new log filterer instance of Bank, bound to a specific deployed contract.
func NewBankFilterer(address common.Address, filterer bind.ContractFilterer) (*BankFilterer, error) {
	contract, err := bindBank(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BankFilterer{contract: contract}, nil
}

// bindBank binds a generic wrapper to an already deployed contract.
func bindBank(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BankABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.BankCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.BankTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bank *BankCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bank.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bank *BankTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bank *BankTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bank.Contract.contract.Transact(opts, method, params...)
}

// BalanceReceived is a free data retrieval call binding the contract method 0x52a90c42.
//
// Solidity: function balanceReceived() view returns(uint256)
func (_Bank *BankCaller) BalanceReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "balanceReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceReceived is a free data retrieval call binding the contract method 0x52a90c42.
//
// Solidity: function balanceReceived() view returns(uint256)
func (_Bank *BankSession) BalanceReceived() (*big.Int, error) {
	return _Bank.Contract.BalanceReceived(&_Bank.CallOpts)
}

// BalanceReceived is a free data retrieval call binding the contract method 0x52a90c42.
//
// Solidity: function balanceReceived() view returns(uint256)
func (_Bank *BankCallerSession) BalanceReceived() (*big.Int, error) {
	return _Bank.Contract.BalanceReceived(&_Bank.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Bank *BankCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Bank *BankSession) GetBalance() (*big.Int, error) {
	return _Bank.Contract.GetBalance(&_Bank.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Bank *BankCallerSession) GetBalance() (*big.Int, error) {
	return _Bank.Contract.GetBalance(&_Bank.CallOpts)
}

// GetOperations is a free data retrieval call binding the contract method 0xf948e50a.
//
// Solidity: function getOperations() view returns(uint256)
func (_Bank *BankCaller) GetOperations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bank.contract.Call(opts, &out, "getOperations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperations is a free data retrieval call binding the contract method 0xf948e50a.
//
// Solidity: function getOperations() view returns(uint256)
func (_Bank *BankSession) GetOperations() (*big.Int, error) {
	return _Bank.Contract.GetOperations(&_Bank.CallOpts)
}

// GetOperations is a free data retrieval call binding the contract method 0xf948e50a.
//
// Solidity: function getOperations() view returns(uint256)
func (_Bank *BankCallerSession) GetOperations() (*big.Int, error) {
	return _Bank.Contract.GetOperations(&_Bank.CallOpts)
}

// LogTransferOperations is a paid mutator transaction binding the contract method 0x2059ab90.
//
// Solidity: function logTransferOperations() returns()
func (_Bank *BankTransactor) LogTransferOperations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "logTransferOperations")
}

// LogTransferOperations is a paid mutator transaction binding the contract method 0x2059ab90.
//
// Solidity: function logTransferOperations() returns()
func (_Bank *BankSession) LogTransferOperations() (*types.Transaction, error) {
	return _Bank.Contract.LogTransferOperations(&_Bank.TransactOpts)
}

// LogTransferOperations is a paid mutator transaction binding the contract method 0x2059ab90.
//
// Solidity: function logTransferOperations() returns()
func (_Bank *BankTransactorSession) LogTransferOperations() (*types.Transaction, error) {
	return _Bank.Contract.LogTransferOperations(&_Bank.TransactOpts)
}

// PayMoneyTo is a paid mutator transaction binding the contract method 0x82664c61.
//
// Solidity: function payMoneyTo(address _to) payable returns()
func (_Bank *BankTransactor) PayMoneyTo(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "payMoneyTo", _to)
}

// PayMoneyTo is a paid mutator transaction binding the contract method 0x82664c61.
//
// Solidity: function payMoneyTo(address _to) payable returns()
func (_Bank *BankSession) PayMoneyTo(_to common.Address) (*types.Transaction, error) {
	return _Bank.Contract.PayMoneyTo(&_Bank.TransactOpts, _to)
}

// PayMoneyTo is a paid mutator transaction binding the contract method 0x82664c61.
//
// Solidity: function payMoneyTo(address _to) payable returns()
func (_Bank *BankTransactorSession) PayMoneyTo(_to common.Address) (*types.Transaction, error) {
	return _Bank.Contract.PayMoneyTo(&_Bank.TransactOpts, _to)
}

// ReceiveMoney is a paid mutator transaction binding the contract method 0x6d26ec18.
//
// Solidity: function receiveMoney() payable returns()
func (_Bank *BankTransactor) ReceiveMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "receiveMoney")
}

// ReceiveMoney is a paid mutator transaction binding the contract method 0x6d26ec18.
//
// Solidity: function receiveMoney() payable returns()
func (_Bank *BankSession) ReceiveMoney() (*types.Transaction, error) {
	return _Bank.Contract.ReceiveMoney(&_Bank.TransactOpts)
}

// ReceiveMoney is a paid mutator transaction binding the contract method 0x6d26ec18.
//
// Solidity: function receiveMoney() payable returns()
func (_Bank *BankTransactorSession) ReceiveMoney() (*types.Transaction, error) {
	return _Bank.Contract.ReceiveMoney(&_Bank.TransactOpts)
}

// TransferMoneyTo is a paid mutator transaction binding the contract method 0xec197ed8.
//
// Solidity: function transferMoneyTo(address _to, string id) payable returns()
func (_Bank *BankTransactor) TransferMoneyTo(opts *bind.TransactOpts, _to common.Address, id string) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "transferMoneyTo", _to, id)
}

// TransferMoneyTo is a paid mutator transaction binding the contract method 0xec197ed8.
//
// Solidity: function transferMoneyTo(address _to, string id) payable returns()
func (_Bank *BankSession) TransferMoneyTo(_to common.Address, id string) (*types.Transaction, error) {
	return _Bank.Contract.TransferMoneyTo(&_Bank.TransactOpts, _to, id)
}

// TransferMoneyTo is a paid mutator transaction binding the contract method 0xec197ed8.
//
// Solidity: function transferMoneyTo(address _to, string id) payable returns()
func (_Bank *BankTransactorSession) TransferMoneyTo(_to common.Address, id string) (*types.Transaction, error) {
	return _Bank.Contract.TransferMoneyTo(&_Bank.TransactOpts, _to, id)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_Bank *BankTransactor) WithdrawMoney(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "withdrawMoney")
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_Bank *BankSession) WithdrawMoney() (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoney(&_Bank.TransactOpts)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xac446002.
//
// Solidity: function withdrawMoney() returns()
func (_Bank *BankTransactorSession) WithdrawMoney() (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoney(&_Bank.TransactOpts)
}

// BankBalanceReceivedIterator is returned from FilterBalanceReceived and is used to iterate over the raw logs and unpacked data for BalanceReceived events raised by the Bank contract.
type BankBalanceReceivedIterator struct {
	Event *BankBalanceReceived // Event containing the contract specifics and raw log

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
func (it *BankBalanceReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankBalanceReceived)
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
		it.Event = new(BankBalanceReceived)
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
func (it *BankBalanceReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankBalanceReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankBalanceReceived represents a BalanceReceived event raised by the Bank contract.
type BankBalanceReceived struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBalanceReceived is a free log retrieval operation binding the contract event 0xbb27de1f6791c153e5749c6e33c63738bcdf867fb4a778f2e3ea9b01c338eb5f.
//
// Solidity: event BalanceReceived(uint256 arg0)
func (_Bank *BankFilterer) FilterBalanceReceived(opts *bind.FilterOpts) (*BankBalanceReceivedIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "BalanceReceived")
	if err != nil {
		return nil, err
	}
	return &BankBalanceReceivedIterator{contract: _Bank.contract, event: "BalanceReceived", logs: logs, sub: sub}, nil
}

// WatchBalanceReceived is a free log subscription operation binding the contract event 0xbb27de1f6791c153e5749c6e33c63738bcdf867fb4a778f2e3ea9b01c338eb5f.
//
// Solidity: event BalanceReceived(uint256 arg0)
func (_Bank *BankFilterer) WatchBalanceReceived(opts *bind.WatchOpts, sink chan<- *BankBalanceReceived) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "BalanceReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankBalanceReceived)
				if err := _Bank.contract.UnpackLog(event, "BalanceReceived", log); err != nil {
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

// ParseBalanceReceived is a log parse operation binding the contract event 0xbb27de1f6791c153e5749c6e33c63738bcdf867fb4a778f2e3ea9b01c338eb5f.
//
// Solidity: event BalanceReceived(uint256 arg0)
func (_Bank *BankFilterer) ParseBalanceReceived(log types.Log) (*BankBalanceReceived, error) {
	event := new(BankBalanceReceived)
	if err := _Bank.contract.UnpackLog(event, "BalanceReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankBalanceTransferredIterator is returned from FilterBalanceTransferred and is used to iterate over the raw logs and unpacked data for BalanceTransferred events raised by the Bank contract.
type BankBalanceTransferredIterator struct {
	Event *BankBalanceTransferred // Event containing the contract specifics and raw log

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
func (it *BankBalanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankBalanceTransferred)
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
		it.Event = new(BankBalanceTransferred)
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
func (it *BankBalanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankBalanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankBalanceTransferred represents a BalanceTransferred event raised by the Bank contract.
type BankBalanceTransferred struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBalanceTransferred is a free log retrieval operation binding the contract event 0x3760d80235c7b3a7167f5f927afb58bcdcb8e3f99c8571aaf8544cdb70f3b311.
//
// Solidity: event BalanceTransferred(uint256 arg0)
func (_Bank *BankFilterer) FilterBalanceTransferred(opts *bind.FilterOpts) (*BankBalanceTransferredIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "BalanceTransferred")
	if err != nil {
		return nil, err
	}
	return &BankBalanceTransferredIterator{contract: _Bank.contract, event: "BalanceTransferred", logs: logs, sub: sub}, nil
}

// WatchBalanceTransferred is a free log subscription operation binding the contract event 0x3760d80235c7b3a7167f5f927afb58bcdcb8e3f99c8571aaf8544cdb70f3b311.
//
// Solidity: event BalanceTransferred(uint256 arg0)
func (_Bank *BankFilterer) WatchBalanceTransferred(opts *bind.WatchOpts, sink chan<- *BankBalanceTransferred) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "BalanceTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankBalanceTransferred)
				if err := _Bank.contract.UnpackLog(event, "BalanceTransferred", log); err != nil {
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

// ParseBalanceTransferred is a log parse operation binding the contract event 0x3760d80235c7b3a7167f5f927afb58bcdcb8e3f99c8571aaf8544cdb70f3b311.
//
// Solidity: event BalanceTransferred(uint256 arg0)
func (_Bank *BankFilterer) ParseBalanceTransferred(log types.Log) (*BankBalanceTransferred, error) {
	event := new(BankBalanceTransferred)
	if err := _Bank.contract.UnpackLog(event, "BalanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankOperationsExecutedIterator is returned from FilterOperationsExecuted and is used to iterate over the raw logs and unpacked data for OperationsExecuted events raised by the Bank contract.
type BankOperationsExecutedIterator struct {
	Event *BankOperationsExecuted // Event containing the contract specifics and raw log

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
func (it *BankOperationsExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankOperationsExecuted)
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
		it.Event = new(BankOperationsExecuted)
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
func (it *BankOperationsExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankOperationsExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankOperationsExecuted represents a OperationsExecuted event raised by the Bank contract.
type BankOperationsExecuted struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOperationsExecuted is a free log retrieval operation binding the contract event 0x45b1c2f5c9e32511e6eb1772825abe442d6058ea66dd4e139a8fcc84b7a3c712.
//
// Solidity: event OperationsExecuted(uint256 arg0)
func (_Bank *BankFilterer) FilterOperationsExecuted(opts *bind.FilterOpts) (*BankOperationsExecutedIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "OperationsExecuted")
	if err != nil {
		return nil, err
	}
	return &BankOperationsExecutedIterator{contract: _Bank.contract, event: "OperationsExecuted", logs: logs, sub: sub}, nil
}

// WatchOperationsExecuted is a free log subscription operation binding the contract event 0x45b1c2f5c9e32511e6eb1772825abe442d6058ea66dd4e139a8fcc84b7a3c712.
//
// Solidity: event OperationsExecuted(uint256 arg0)
func (_Bank *BankFilterer) WatchOperationsExecuted(opts *bind.WatchOpts, sink chan<- *BankOperationsExecuted) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "OperationsExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankOperationsExecuted)
				if err := _Bank.contract.UnpackLog(event, "OperationsExecuted", log); err != nil {
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

// ParseOperationsExecuted is a log parse operation binding the contract event 0x45b1c2f5c9e32511e6eb1772825abe442d6058ea66dd4e139a8fcc84b7a3c712.
//
// Solidity: event OperationsExecuted(uint256 arg0)
func (_Bank *BankFilterer) ParseOperationsExecuted(log types.Log) (*BankOperationsExecuted, error) {
	event := new(BankOperationsExecuted)
	if err := _Bank.contract.UnpackLog(event, "OperationsExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BankPrintConfirmationIterator is returned from FilterPrintConfirmation and is used to iterate over the raw logs and unpacked data for PrintConfirmation events raised by the Bank contract.
type BankPrintConfirmationIterator struct {
	Event *BankPrintConfirmation // Event containing the contract specifics and raw log

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
func (it *BankPrintConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BankPrintConfirmation)
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
		it.Event = new(BankPrintConfirmation)
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
func (it *BankPrintConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BankPrintConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BankPrintConfirmation represents a PrintConfirmation event raised by the Bank contract.
type BankPrintConfirmation struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintConfirmation is a free log retrieval operation binding the contract event 0xa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811.
//
// Solidity: event PrintConfirmation(string arg0, uint256 arg1)
func (_Bank *BankFilterer) FilterPrintConfirmation(opts *bind.FilterOpts) (*BankPrintConfirmationIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "PrintConfirmation")
	if err != nil {
		return nil, err
	}
	return &BankPrintConfirmationIterator{contract: _Bank.contract, event: "PrintConfirmation", logs: logs, sub: sub}, nil
}

// WatchPrintConfirmation is a free log subscription operation binding the contract event 0xa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811.
//
// Solidity: event PrintConfirmation(string arg0, uint256 arg1)
func (_Bank *BankFilterer) WatchPrintConfirmation(opts *bind.WatchOpts, sink chan<- *BankPrintConfirmation) (event.Subscription, error) {

	logs, sub, err := _Bank.contract.WatchLogs(opts, "PrintConfirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BankPrintConfirmation)
				if err := _Bank.contract.UnpackLog(event, "PrintConfirmation", log); err != nil {
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
func (_Bank *BankFilterer) ParsePrintConfirmation(log types.Log) (*BankPrintConfirmation, error) {
	event := new(BankPrintConfirmation)
	if err := _Bank.contract.UnpackLog(event, "PrintConfirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
