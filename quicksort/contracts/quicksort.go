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

// QuickSortMetaData contains all meta data concerning the QuickSort contract.
var QuickSortMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"PrintArray\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"PrintArrayInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"PrintHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PrintSetElementQty\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"debugSort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"printAllData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"sort\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"32a66a63": "debugSort(uint256)",
		"d969f9f4": "printAllData()",
		"fe913865": "sort(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506107e0806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806332a66a6314610046578063d969f9f41461005b578063fe91386514610063575b600080fd5b6100596100543660046105b1565b610076565b005b6100596101f8565b6100596100713660046105b1565b6102a9565b60008167ffffffffffffffff81111561009157610091610794565b6040519080825280602002602001820160405280156100ba578160200160208202803683370190505b50905060005b8151811015610102576100d38184610703565b8282815181106100e5576100e561077e565b6020908102919091010152806100fa81610758565b9150506100c0565b506000816040516020016101169190610605565b6040516020818303038152906040528051906020012090507f3428030cccc788be8ed6d0f32eb4f0fa5fa9958367fe596551a986ddd76dc9038382846040516101619392919061061f565b60405180910390a16101828260006001855161017d9190610703565b6103b7565b6000826040516020016101959190610605565b6040516020818303038152906040528051906020012090506101b78185610538565b7f3428030cccc788be8ed6d0f32eb4f0fa5fa9958367fe596551a986ddd76dc9038482856040516101ea9392919061061f565b60405180910390a150505050565b60005b6002548110156102a657600080600201828154811061021c5761021c61077e565b60009182526020808320919091015480835282825260408084205460025460018552948290205482519182529381018390529081018690526060810193909352608083019190915291507fc2a813325c8b15ae781043d7e241ff901b9a7f7403122547f1e33ee61a8c21139060a00160405180910390a1508061029e81610758565b9150506101fb565b50565b60008167ffffffffffffffff8111156102c4576102c4610794565b6040519080825280602002602001820160405280156102ed578160200160208202803683370190505b50905060005b8151811015610335576103068184610703565b8282815181106103185761031861077e565b60209081029190910101528061032d81610758565b9150506102f3565b5061034a8160006001845161017d9190610703565b60008160405160200161035d9190610605565b60405160208183030381529060405280519060200120905061037f8184610538565b6040518181527fff5f0f4ae79415c7d9c291f9207f5a799f2fcfc9cb2ba600fe443ad1674ae2869060200160405180910390a1505050565b8181808214156103c8575050505050565b60008560026103d787876106c4565b6103e19190610688565b6103eb9087610647565b815181106103fb576103fb61077e565b602002602001015190505b81831361050a575b808684815181106104215761042161077e565b60200260200101511015610441578261043981610738565b93505061040e565b8582815181106104535761045361077e565b6020026020010151811015610474578161046c8161071a565b925050610441565b8183136105055785828151811061048d5761048d61077e565b60200260200101518684815181106104a7576104a761077e565b60200260200101518785815181106104c1576104c161077e565b602002602001018885815181106104da576104da61077e565b602090810291909101019190915252826104f381610738565b93505081806105019061071a565b9250505b610406565b8185121561051d5761051d8686846103b7565b83831215610530576105308684866103b7565b505050505050565b60008281526001602052604090205461058e5760008281526020819052604081208290556002805460018101825591527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace018290555b60008281526001602052604081208054916105a883610758565b91905055505050565b6000602082840312156105c357600080fd5b5035919050565b600081518084526020808501945080840160005b838110156105fa578151875295820195908201906001016105de565b509495945050505050565b60208152600061061860208301846105ca565b9392505050565b83815282602082015260606040820152600061063e60608301846105ca565b95945050505050565b600080821280156001600160ff1b038490038513161561066957610669610768565b600160ff1b839003841281161561068257610682610768565b50500190565b6000826106a557634e487b7160e01b600052601260045260246000fd5b600160ff1b8214600019841416156106bf576106bf610768565b500590565b60008083128015600160ff1b8501841216156106e2576106e2610768565b6001600160ff1b03840183138116156106fd576106fd610768565b50500390565b60008282101561071557610715610768565b500390565b6000600160ff1b82141561073057610730610768565b506000190190565b60006001600160ff1b0382141561075157610751610768565b5060010190565b6000600019821415610751576107515b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220c346608b34b86ac9193f11dda3bf8336b179d743332f63ea309f797c20810e2964736f6c63430008070033",
}

// QuickSortABI is the input ABI used to generate the binding from.
// Deprecated: Use QuickSortMetaData.ABI instead.
var QuickSortABI = QuickSortMetaData.ABI

// Deprecated: Use QuickSortMetaData.Sigs instead.
// QuickSortFuncSigs maps the 4-byte function signature to its string representation.
var QuickSortFuncSigs = QuickSortMetaData.Sigs

// QuickSortBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuickSortMetaData.Bin instead.
var QuickSortBin = QuickSortMetaData.Bin

// DeployQuickSort deploys a new Ethereum contract, binding an instance of QuickSort to it.
func DeployQuickSort(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QuickSort, error) {
	parsed, err := QuickSortMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuickSortBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QuickSort{QuickSortCaller: QuickSortCaller{contract: contract}, QuickSortTransactor: QuickSortTransactor{contract: contract}, QuickSortFilterer: QuickSortFilterer{contract: contract}}, nil
}

// QuickSort is an auto generated Go binding around an Ethereum contract.
type QuickSort struct {
	QuickSortCaller     // Read-only binding to the contract
	QuickSortTransactor // Write-only binding to the contract
	QuickSortFilterer   // Log filterer for contract events
}

// QuickSortCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuickSortCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickSortTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuickSortTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickSortFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuickSortFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuickSortSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuickSortSession struct {
	Contract     *QuickSort        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuickSortCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuickSortCallerSession struct {
	Contract *QuickSortCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// QuickSortTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuickSortTransactorSession struct {
	Contract     *QuickSortTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// QuickSortRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuickSortRaw struct {
	Contract *QuickSort // Generic contract binding to access the raw methods on
}

// QuickSortCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuickSortCallerRaw struct {
	Contract *QuickSortCaller // Generic read-only contract binding to access the raw methods on
}

// QuickSortTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuickSortTransactorRaw struct {
	Contract *QuickSortTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuickSort creates a new instance of QuickSort, bound to a specific deployed contract.
func NewQuickSort(address common.Address, backend bind.ContractBackend) (*QuickSort, error) {
	contract, err := bindQuickSort(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QuickSort{QuickSortCaller: QuickSortCaller{contract: contract}, QuickSortTransactor: QuickSortTransactor{contract: contract}, QuickSortFilterer: QuickSortFilterer{contract: contract}}, nil
}

// NewQuickSortCaller creates a new read-only instance of QuickSort, bound to a specific deployed contract.
func NewQuickSortCaller(address common.Address, caller bind.ContractCaller) (*QuickSortCaller, error) {
	contract, err := bindQuickSort(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuickSortCaller{contract: contract}, nil
}

// NewQuickSortTransactor creates a new write-only instance of QuickSort, bound to a specific deployed contract.
func NewQuickSortTransactor(address common.Address, transactor bind.ContractTransactor) (*QuickSortTransactor, error) {
	contract, err := bindQuickSort(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuickSortTransactor{contract: contract}, nil
}

// NewQuickSortFilterer creates a new log filterer instance of QuickSort, bound to a specific deployed contract.
func NewQuickSortFilterer(address common.Address, filterer bind.ContractFilterer) (*QuickSortFilterer, error) {
	contract, err := bindQuickSort(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuickSortFilterer{contract: contract}, nil
}

// bindQuickSort binds a generic wrapper to an already deployed contract.
func bindQuickSort(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuickSortABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuickSort *QuickSortRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuickSort.Contract.QuickSortCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuickSort *QuickSortRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuickSort.Contract.QuickSortTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuickSort *QuickSortRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuickSort.Contract.QuickSortTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QuickSort *QuickSortCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QuickSort.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QuickSort *QuickSortTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuickSort.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QuickSort *QuickSortTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QuickSort.Contract.contract.Transact(opts, method, params...)
}

// DebugSort is a paid mutator transaction binding the contract method 0x32a66a63.
//
// Solidity: function debugSort(uint256 size) returns()
func (_QuickSort *QuickSortTransactor) DebugSort(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _QuickSort.contract.Transact(opts, "debugSort", size)
}

// DebugSort is a paid mutator transaction binding the contract method 0x32a66a63.
//
// Solidity: function debugSort(uint256 size) returns()
func (_QuickSort *QuickSortSession) DebugSort(size *big.Int) (*types.Transaction, error) {
	return _QuickSort.Contract.DebugSort(&_QuickSort.TransactOpts, size)
}

// DebugSort is a paid mutator transaction binding the contract method 0x32a66a63.
//
// Solidity: function debugSort(uint256 size) returns()
func (_QuickSort *QuickSortTransactorSession) DebugSort(size *big.Int) (*types.Transaction, error) {
	return _QuickSort.Contract.DebugSort(&_QuickSort.TransactOpts, size)
}

// PrintAllData is a paid mutator transaction binding the contract method 0xd969f9f4.
//
// Solidity: function printAllData() returns()
func (_QuickSort *QuickSortTransactor) PrintAllData(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QuickSort.contract.Transact(opts, "printAllData")
}

// PrintAllData is a paid mutator transaction binding the contract method 0xd969f9f4.
//
// Solidity: function printAllData() returns()
func (_QuickSort *QuickSortSession) PrintAllData() (*types.Transaction, error) {
	return _QuickSort.Contract.PrintAllData(&_QuickSort.TransactOpts)
}

// PrintAllData is a paid mutator transaction binding the contract method 0xd969f9f4.
//
// Solidity: function printAllData() returns()
func (_QuickSort *QuickSortTransactorSession) PrintAllData() (*types.Transaction, error) {
	return _QuickSort.Contract.PrintAllData(&_QuickSort.TransactOpts)
}

// Sort is a paid mutator transaction binding the contract method 0xfe913865.
//
// Solidity: function sort(uint256 size) returns()
func (_QuickSort *QuickSortTransactor) Sort(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _QuickSort.contract.Transact(opts, "sort", size)
}

// Sort is a paid mutator transaction binding the contract method 0xfe913865.
//
// Solidity: function sort(uint256 size) returns()
func (_QuickSort *QuickSortSession) Sort(size *big.Int) (*types.Transaction, error) {
	return _QuickSort.Contract.Sort(&_QuickSort.TransactOpts, size)
}

// Sort is a paid mutator transaction binding the contract method 0xfe913865.
//
// Solidity: function sort(uint256 size) returns()
func (_QuickSort *QuickSortTransactorSession) Sort(size *big.Int) (*types.Transaction, error) {
	return _QuickSort.Contract.Sort(&_QuickSort.TransactOpts, size)
}

// QuickSortPrintArrayIterator is returned from FilterPrintArray and is used to iterate over the raw logs and unpacked data for PrintArray events raised by the QuickSort contract.
type QuickSortPrintArrayIterator struct {
	Event *QuickSortPrintArray // Event containing the contract specifics and raw log

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
func (it *QuickSortPrintArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuickSortPrintArray)
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
		it.Event = new(QuickSortPrintArray)
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
func (it *QuickSortPrintArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuickSortPrintArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuickSortPrintArray represents a PrintArray event raised by the QuickSort contract.
type QuickSortPrintArray struct {
	Arg0 *big.Int
	Arg1 []*big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintArray is a free log retrieval operation binding the contract event 0x043daf33342eb58a6acb2c97073a2b497d0fc1af85381229e380d60333ec0424.
//
// Solidity: event PrintArray(uint256 arg0, uint256[] arg1)
func (_QuickSort *QuickSortFilterer) FilterPrintArray(opts *bind.FilterOpts) (*QuickSortPrintArrayIterator, error) {

	logs, sub, err := _QuickSort.contract.FilterLogs(opts, "PrintArray")
	if err != nil {
		return nil, err
	}
	return &QuickSortPrintArrayIterator{contract: _QuickSort.contract, event: "PrintArray", logs: logs, sub: sub}, nil
}

// WatchPrintArray is a free log subscription operation binding the contract event 0x043daf33342eb58a6acb2c97073a2b497d0fc1af85381229e380d60333ec0424.
//
// Solidity: event PrintArray(uint256 arg0, uint256[] arg1)
func (_QuickSort *QuickSortFilterer) WatchPrintArray(opts *bind.WatchOpts, sink chan<- *QuickSortPrintArray) (event.Subscription, error) {

	logs, sub, err := _QuickSort.contract.WatchLogs(opts, "PrintArray")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuickSortPrintArray)
				if err := _QuickSort.contract.UnpackLog(event, "PrintArray", log); err != nil {
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

// ParsePrintArray is a log parse operation binding the contract event 0x043daf33342eb58a6acb2c97073a2b497d0fc1af85381229e380d60333ec0424.
//
// Solidity: event PrintArray(uint256 arg0, uint256[] arg1)
func (_QuickSort *QuickSortFilterer) ParsePrintArray(log types.Log) (*QuickSortPrintArray, error) {
	event := new(QuickSortPrintArray)
	if err := _QuickSort.contract.UnpackLog(event, "PrintArray", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuickSortPrintArrayInfoIterator is returned from FilterPrintArrayInfo and is used to iterate over the raw logs and unpacked data for PrintArrayInfo events raised by the QuickSort contract.
type QuickSortPrintArrayInfoIterator struct {
	Event *QuickSortPrintArrayInfo // Event containing the contract specifics and raw log

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
func (it *QuickSortPrintArrayInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuickSortPrintArrayInfo)
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
		it.Event = new(QuickSortPrintArrayInfo)
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
func (it *QuickSortPrintArrayInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuickSortPrintArrayInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuickSortPrintArrayInfo represents a PrintArrayInfo event raised by the QuickSort contract.
type QuickSortPrintArrayInfo struct {
	Arg0 *big.Int
	Arg1 [32]byte
	Arg2 []*big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintArrayInfo is a free log retrieval operation binding the contract event 0x3428030cccc788be8ed6d0f32eb4f0fa5fa9958367fe596551a986ddd76dc903.
//
// Solidity: event PrintArrayInfo(uint256 arg0, bytes32 arg1, uint256[] arg2)
func (_QuickSort *QuickSortFilterer) FilterPrintArrayInfo(opts *bind.FilterOpts) (*QuickSortPrintArrayInfoIterator, error) {

	logs, sub, err := _QuickSort.contract.FilterLogs(opts, "PrintArrayInfo")
	if err != nil {
		return nil, err
	}
	return &QuickSortPrintArrayInfoIterator{contract: _QuickSort.contract, event: "PrintArrayInfo", logs: logs, sub: sub}, nil
}

// WatchPrintArrayInfo is a free log subscription operation binding the contract event 0x3428030cccc788be8ed6d0f32eb4f0fa5fa9958367fe596551a986ddd76dc903.
//
// Solidity: event PrintArrayInfo(uint256 arg0, bytes32 arg1, uint256[] arg2)
func (_QuickSort *QuickSortFilterer) WatchPrintArrayInfo(opts *bind.WatchOpts, sink chan<- *QuickSortPrintArrayInfo) (event.Subscription, error) {

	logs, sub, err := _QuickSort.contract.WatchLogs(opts, "PrintArrayInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuickSortPrintArrayInfo)
				if err := _QuickSort.contract.UnpackLog(event, "PrintArrayInfo", log); err != nil {
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

// ParsePrintArrayInfo is a log parse operation binding the contract event 0x3428030cccc788be8ed6d0f32eb4f0fa5fa9958367fe596551a986ddd76dc903.
//
// Solidity: event PrintArrayInfo(uint256 arg0, bytes32 arg1, uint256[] arg2)
func (_QuickSort *QuickSortFilterer) ParsePrintArrayInfo(log types.Log) (*QuickSortPrintArrayInfo, error) {
	event := new(QuickSortPrintArrayInfo)
	if err := _QuickSort.contract.UnpackLog(event, "PrintArrayInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuickSortPrintHashIterator is returned from FilterPrintHash and is used to iterate over the raw logs and unpacked data for PrintHash events raised by the QuickSort contract.
type QuickSortPrintHashIterator struct {
	Event *QuickSortPrintHash // Event containing the contract specifics and raw log

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
func (it *QuickSortPrintHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuickSortPrintHash)
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
		it.Event = new(QuickSortPrintHash)
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
func (it *QuickSortPrintHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuickSortPrintHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuickSortPrintHash represents a PrintHash event raised by the QuickSort contract.
type QuickSortPrintHash struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintHash is a free log retrieval operation binding the contract event 0xff5f0f4ae79415c7d9c291f9207f5a799f2fcfc9cb2ba600fe443ad1674ae286.
//
// Solidity: event PrintHash(bytes32 arg0)
func (_QuickSort *QuickSortFilterer) FilterPrintHash(opts *bind.FilterOpts) (*QuickSortPrintHashIterator, error) {

	logs, sub, err := _QuickSort.contract.FilterLogs(opts, "PrintHash")
	if err != nil {
		return nil, err
	}
	return &QuickSortPrintHashIterator{contract: _QuickSort.contract, event: "PrintHash", logs: logs, sub: sub}, nil
}

// WatchPrintHash is a free log subscription operation binding the contract event 0xff5f0f4ae79415c7d9c291f9207f5a799f2fcfc9cb2ba600fe443ad1674ae286.
//
// Solidity: event PrintHash(bytes32 arg0)
func (_QuickSort *QuickSortFilterer) WatchPrintHash(opts *bind.WatchOpts, sink chan<- *QuickSortPrintHash) (event.Subscription, error) {

	logs, sub, err := _QuickSort.contract.WatchLogs(opts, "PrintHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuickSortPrintHash)
				if err := _QuickSort.contract.UnpackLog(event, "PrintHash", log); err != nil {
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

// ParsePrintHash is a log parse operation binding the contract event 0xff5f0f4ae79415c7d9c291f9207f5a799f2fcfc9cb2ba600fe443ad1674ae286.
//
// Solidity: event PrintHash(bytes32 arg0)
func (_QuickSort *QuickSortFilterer) ParsePrintHash(log types.Log) (*QuickSortPrintHash, error) {
	event := new(QuickSortPrintHash)
	if err := _QuickSort.contract.UnpackLog(event, "PrintHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuickSortPrintSetElementQtyIterator is returned from FilterPrintSetElementQty and is used to iterate over the raw logs and unpacked data for PrintSetElementQty events raised by the QuickSort contract.
type QuickSortPrintSetElementQtyIterator struct {
	Event *QuickSortPrintSetElementQty // Event containing the contract specifics and raw log

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
func (it *QuickSortPrintSetElementQtyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuickSortPrintSetElementQty)
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
		it.Event = new(QuickSortPrintSetElementQty)
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
func (it *QuickSortPrintSetElementQtyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuickSortPrintSetElementQtyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuickSortPrintSetElementQty represents a PrintSetElementQty event raised by the QuickSort contract.
type QuickSortPrintSetElementQty struct {
	Arg0 *big.Int
	Arg1 [32]byte
	Arg2 *big.Int
	Arg3 *big.Int
	Arg4 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPrintSetElementQty is a free log retrieval operation binding the contract event 0xc2a813325c8b15ae781043d7e241ff901b9a7f7403122547f1e33ee61a8c2113.
//
// Solidity: event PrintSetElementQty(uint256 arg0, bytes32 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_QuickSort *QuickSortFilterer) FilterPrintSetElementQty(opts *bind.FilterOpts) (*QuickSortPrintSetElementQtyIterator, error) {

	logs, sub, err := _QuickSort.contract.FilterLogs(opts, "PrintSetElementQty")
	if err != nil {
		return nil, err
	}
	return &QuickSortPrintSetElementQtyIterator{contract: _QuickSort.contract, event: "PrintSetElementQty", logs: logs, sub: sub}, nil
}

// WatchPrintSetElementQty is a free log subscription operation binding the contract event 0xc2a813325c8b15ae781043d7e241ff901b9a7f7403122547f1e33ee61a8c2113.
//
// Solidity: event PrintSetElementQty(uint256 arg0, bytes32 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_QuickSort *QuickSortFilterer) WatchPrintSetElementQty(opts *bind.WatchOpts, sink chan<- *QuickSortPrintSetElementQty) (event.Subscription, error) {

	logs, sub, err := _QuickSort.contract.WatchLogs(opts, "PrintSetElementQty")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuickSortPrintSetElementQty)
				if err := _QuickSort.contract.UnpackLog(event, "PrintSetElementQty", log); err != nil {
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

// ParsePrintSetElementQty is a log parse operation binding the contract event 0xc2a813325c8b15ae781043d7e241ff901b9a7f7403122547f1e33ee61a8c2113.
//
// Solidity: event PrintSetElementQty(uint256 arg0, bytes32 arg1, uint256 arg2, uint256 arg3, uint256 arg4)
func (_QuickSort *QuickSortFilterer) ParsePrintSetElementQty(log types.Log) (*QuickSortPrintSetElementQty, error) {
	event := new(QuickSortPrintSetElementQty)
	if err := _QuickSort.contract.UnpackLog(event, "PrintSetElementQty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
