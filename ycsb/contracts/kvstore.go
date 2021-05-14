// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KVstoreABI is the input ABI used to generate the binding from.
const KVstoreABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KVstoreFuncSigs maps the 4-byte function signature to its string representation.
var KVstoreFuncSigs = map[string]string{
	"693ec85e": "get(string)",
	"e942b516": "set(string,string)",
}

// KVstoreBin is the compiled bytecode used for deploying new contracts.
var KVstoreBin = "0x608060405234801561001057600080fd5b50610422806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063693ec85e1461003b578063e942b51614610064575b600080fd5b61004e610049366004610280565b610079565b60405161005b9190610338565b60405180910390f35b6100776100723660046102bb565b610129565b005b606060008260405161008b919061031c565b908152602001604051809103902080546100a49061039b565b80601f01602080910402602001604051908101604052809291908181526020018280546100d09061039b565b801561011d5780601f106100f25761010080835404028352916020019161011d565b820191906000526020600020905b81548152906001019060200180831161010057829003601f168201915b50505050509050919050565b8060008360405161013a919061031c565b9081526020016040518091039020908051906020019061015b929190610160565b505050565b82805461016c9061039b565b90600052602060002090601f01602090048101928261018e57600085556101d4565b82601f106101a757805160ff19168380011785556101d4565b828001600101855582156101d4579182015b828111156101d45782518255916020019190600101906101b9565b506101e09291506101e4565b5090565b5b808211156101e057600081556001016101e5565b600082601f830112610209578081fd5b813567ffffffffffffffff80821115610224576102246103d6565b604051601f8301601f19908116603f0116810190828211818310171561024c5761024c6103d6565b81604052838152866020858801011115610264578485fd5b8360208701602083013792830160200193909352509392505050565b600060208284031215610291578081fd5b813567ffffffffffffffff8111156102a7578182fd5b6102b3848285016101f9565b949350505050565b600080604083850312156102cd578081fd5b823567ffffffffffffffff808211156102e4578283fd5b6102f0868387016101f9565b93506020850135915080821115610305578283fd5b50610312858286016101f9565b9150509250929050565b6000825161032e81846020870161036b565b9190910192915050565b600060208252825180602084015261035781604085016020870161036b565b601f01601f19169190910160400192915050565b60005b8381101561038657818101518382015260200161036e565b83811115610395576000848401525b50505050565b6002810460018216806103af57607f821691505b602082108114156103d057634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fdfea264697066735822122043235aba1e2c5f299df2458942e6b03d9212c9ad7afda791d4679e2a319d715364736f6c63430008020033"

// DeployKVstore deploys a new Ethereum contract, binding an instance of KVstore to it.
func DeployKVstore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KVstore, error) {
	parsed, err := abi.JSON(strings.NewReader(KVstoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KVstoreBin), backend)
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

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string _key, string _value) returns()
func (_KVstore *KVstoreTransactor) Set(opts *bind.TransactOpts, _key string, _value string) (*types.Transaction, error) {
	return _KVstore.contract.Transact(opts, "set", _key, _value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string _key, string _value) returns()
func (_KVstore *KVstoreSession) Set(_key string, _value string) (*types.Transaction, error) {
	return _KVstore.Contract.Set(&_KVstore.TransactOpts, _key, _value)
}

// Set is a paid mutator transaction binding the contract method 0xe942b516.
//
// Solidity: function set(string _key, string _value) returns()
func (_KVstore *KVstoreTransactorSession) Set(_key string, _value string) (*types.Transaction, error) {
	return _KVstore.Contract.Set(&_KVstore.TransactOpts, _key, _value)
}
