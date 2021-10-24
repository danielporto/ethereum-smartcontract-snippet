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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BalanceReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BalanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"OperationsExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PrintConfirmation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"balanceReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"directTransferTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOperations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logTransferOperations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveMoney\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"withdrawMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawMoneyTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"52a90c42": "balanceReceived()",
		"a26e1186": "deposit(string)",
		"a9c73c36": "directTransferTo(address,string)",
		"12065fe0": "getBalance()",
		"f948e50a": "getOperations()",
		"2059ab90": "logTransferOperations()",
		"6d26ec18": "receiveMoney()",
		"f8d708dc": "withdrawMoney(string)",
		"e883c06a": "withdrawMoneyTo(string,address)",
	},
	Bin: "0x60806040526000805534801561001457600080fd5b50610537806100246000396000f3fe6080604052600436106100865760003560e01c8063a26e118611610059578063a26e1186146100e1578063a9c73c36146100f4578063e883c06a14610107578063f8d708dc1461011a578063f948e50a1461013a57600080fd5b806312065fe01461008b5780632059ab90146100ac57806352a90c42146100c35780636d26ec18146100d9575b600080fd5b34801561009757600080fd5b50475b60405190815260200160405180910390f35b3480156100b857600080fd5b506100c161014f565b005b3480156100cf57600080fd5b5061009a60015481565b6100c161018d565b6100c16100ef36600461039d565b6101d4565b6100c16101023660046103fb565b61022c565b6100c161011536600461044e565b6102bb565b34801561012657600080fd5b506100c161013536600461039d565b610326565b34801561014657600080fd5b5060005461009a565b600054604080519182524760208301527fef1c94cbc35308081c9f68ac9090ed940be17995b42a7c9788d6854336e51fb291015b60405180910390a1565b346001600082825461019f91906104a2565b90915550506040513481527fbb27de1f6791c153e5749c6e33c63738bcdf867fb4a778f2e3ea9b01c338eb5f90602001610183565b34600160008282546101e691906104a2565b90915550506000546040517fa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc187878119161022091859185916104c8565b60405180910390a15050565b6040516001600160a01b038416903480156108fc02916000818181858888f19350505050158015610261573d6000803e3d6000fd5b50600160008082825461027491906104a2565b90915550506000546040517fa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc18787811916102ae91859185916104c8565b60405180910390a1505050565b6040516001600160a01b038216903480156108fc02916000818181858888f193505050501580156102f0573d6000803e3d6000fd5b507fa5b236924ed143061aeda12329df0e5b7ecf459631fbbf10a8f604cc1878781183836000546040516102ae939291906104c8565b33806108fc476040518115909202916000818181858888f193505050501580156102f0573d6000803e3d6000fd5b60008083601f84011261036657600080fd5b50813567ffffffffffffffff81111561037e57600080fd5b60208301915083602082850101111561039657600080fd5b9250929050565b600080602083850312156103b057600080fd5b823567ffffffffffffffff8111156103c757600080fd5b6103d385828601610354565b90969095509350505050565b80356001600160a01b03811681146103f657600080fd5b919050565b60008060006040848603121561041057600080fd5b610419846103df565b9250602084013567ffffffffffffffff81111561043557600080fd5b61044186828701610354565b9497909650939450505050565b60008060006040848603121561046357600080fd5b833567ffffffffffffffff81111561047a57600080fd5b61048686828701610354565b90945092506104999050602085016103df565b90509250925092565b600082198211156104c357634e487b7160e01b600052601160045260246000fd5b500190565b604081528260408201528284606083013760006060848301015260006060601f19601f860116830101905082602083015294935050505056fea26469706673582212204404e1b08124663ce9250e650207fe6d3ef715c4cb5d72fb1cafca0c2a96e8d064736f6c63430008090033",
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

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string id) payable returns()
func (_Bank *BankTransactor) Deposit(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "deposit", id)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string id) payable returns()
func (_Bank *BankSession) Deposit(id string) (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts, id)
}

// Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
//
// Solidity: function deposit(string id) payable returns()
func (_Bank *BankTransactorSession) Deposit(id string) (*types.Transaction, error) {
	return _Bank.Contract.Deposit(&_Bank.TransactOpts, id)
}

// DirectTransferTo is a paid mutator transaction binding the contract method 0xa9c73c36.
//
// Solidity: function directTransferTo(address _to, string id) payable returns()
func (_Bank *BankTransactor) DirectTransferTo(opts *bind.TransactOpts, _to common.Address, id string) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "directTransferTo", _to, id)
}

// DirectTransferTo is a paid mutator transaction binding the contract method 0xa9c73c36.
//
// Solidity: function directTransferTo(address _to, string id) payable returns()
func (_Bank *BankSession) DirectTransferTo(_to common.Address, id string) (*types.Transaction, error) {
	return _Bank.Contract.DirectTransferTo(&_Bank.TransactOpts, _to, id)
}

// DirectTransferTo is a paid mutator transaction binding the contract method 0xa9c73c36.
//
// Solidity: function directTransferTo(address _to, string id) payable returns()
func (_Bank *BankTransactorSession) DirectTransferTo(_to common.Address, id string) (*types.Transaction, error) {
	return _Bank.Contract.DirectTransferTo(&_Bank.TransactOpts, _to, id)
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

// WithdrawMoney is a paid mutator transaction binding the contract method 0xf8d708dc.
//
// Solidity: function withdrawMoney(string id) returns()
func (_Bank *BankTransactor) WithdrawMoney(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "withdrawMoney", id)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xf8d708dc.
//
// Solidity: function withdrawMoney(string id) returns()
func (_Bank *BankSession) WithdrawMoney(id string) (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoney(&_Bank.TransactOpts, id)
}

// WithdrawMoney is a paid mutator transaction binding the contract method 0xf8d708dc.
//
// Solidity: function withdrawMoney(string id) returns()
func (_Bank *BankTransactorSession) WithdrawMoney(id string) (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoney(&_Bank.TransactOpts, id)
}

// WithdrawMoneyTo is a paid mutator transaction binding the contract method 0xe883c06a.
//
// Solidity: function withdrawMoneyTo(string id, address _to) payable returns()
func (_Bank *BankTransactor) WithdrawMoneyTo(opts *bind.TransactOpts, id string, _to common.Address) (*types.Transaction, error) {
	return _Bank.contract.Transact(opts, "withdrawMoneyTo", id, _to)
}

// WithdrawMoneyTo is a paid mutator transaction binding the contract method 0xe883c06a.
//
// Solidity: function withdrawMoneyTo(string id, address _to) payable returns()
func (_Bank *BankSession) WithdrawMoneyTo(id string, _to common.Address) (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoneyTo(&_Bank.TransactOpts, id, _to)
}

// WithdrawMoneyTo is a paid mutator transaction binding the contract method 0xe883c06a.
//
// Solidity: function withdrawMoneyTo(string id, address _to) payable returns()
func (_Bank *BankTransactorSession) WithdrawMoneyTo(id string, _to common.Address) (*types.Transaction, error) {
	return _Bank.Contract.WithdrawMoneyTo(&_Bank.TransactOpts, id, _to)
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
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOperationsExecuted is a free log retrieval operation binding the contract event 0xef1c94cbc35308081c9f68ac9090ed940be17995b42a7c9788d6854336e51fb2.
//
// Solidity: event OperationsExecuted(uint256 arg0, uint256 arg1)
func (_Bank *BankFilterer) FilterOperationsExecuted(opts *bind.FilterOpts) (*BankOperationsExecutedIterator, error) {

	logs, sub, err := _Bank.contract.FilterLogs(opts, "OperationsExecuted")
	if err != nil {
		return nil, err
	}
	return &BankOperationsExecutedIterator{contract: _Bank.contract, event: "OperationsExecuted", logs: logs, sub: sub}, nil
}

// WatchOperationsExecuted is a free log subscription operation binding the contract event 0xef1c94cbc35308081c9f68ac9090ed940be17995b42a7c9788d6854336e51fb2.
//
// Solidity: event OperationsExecuted(uint256 arg0, uint256 arg1)
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

// ParseOperationsExecuted is a log parse operation binding the contract event 0xef1c94cbc35308081c9f68ac9090ed940be17995b42a7c9788d6854336e51fb2.
//
// Solidity: event OperationsExecuted(uint256 arg0, uint256 arg1)
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
