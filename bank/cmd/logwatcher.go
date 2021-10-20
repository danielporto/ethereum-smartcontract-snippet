package cmd

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/danielporto/ethereum-smartcontract-snippet/bank/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
* functions in this file use the stub generated from the solidity contract
* therefore, if the contract is changed, this file will likely require updates to reflect
* the updated version of datastructures and events.
 */

func watchOperationsExecutedEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.BankMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "OperationsExecuted")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (OperationsExecuted)", err)
	}
	Log("Logging thread: Listening for OperationsExecuted events")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.BankOperationsExecuted)
			if err := bc.UnpackLog(event, "OperationsExecuted", logentry); err != nil {
				LogError("Error decoding log event:", logentry, err)
			}
			Log("[BankOperationsExecuted Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchBalanceTransferredEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.BankMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "BalanceTransferred")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (BalanceTransferred)", err)
	}
	Log("Logging thread: Listening for BalanceTransferred events")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.BankBalanceTransferred)
			if err := bc.UnpackLog(event, "BalanceTransferred", logentry); err != nil {
				LogError("Error decoding log event:", logentry, err)
			}
			Log("[BalanceTransferred Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchBalanceReceivedEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.BankMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "BalanceReceived")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (BalanceReceived)", err)
	}
	Log("Logging thread: Listening to BalanceReceived events")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.BankBalanceReceived)
			if err := bc.UnpackLog(event, "Decrement", logentry); err != nil {
				LogError("Error decoding log event:", logentry, err)
			}
			Log("[Decrement Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}


func watchPrintConfirmation(client *ethclient.Client, contractAddr common.Address, stop chan struct{}, requestTimeMap *sync.Map, stat *StatsStorage) {
	eventName := "PrintConfirmation"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.BankMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events ("+eventName+")", err)
	}
	Log("Logging thread: Listening for "+eventName+" events")

	for {
		select {
		case logentry := <-logs:
			tFin := time.Now().UnixNano() / latency_factor // check latency_factor_unity
			event := new(contracts.BankPrintConfirmation)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				LogError("Error decoding log event:", logentry, err)
			}
			tIni, ok := requestTimeMap.LoadAndDelete(event.Arg0) // the value is no longer needed, release memory
			if !ok {
				// reply was received without a request. breaks causality.
				LogError("Reply was received without a request. breaks causality:"+ "{ \"event\": \"%v\", \"trx_id\": \"%v\", \"trx_number\": \"%v\" }", eventName, event.Arg0, event.Arg1)
			}
			tIniTyped, ok := tIni.(int64)
			if !ok {
				LogError("Unable to convert timestamp type, ignoring entry")
				continue
			}
			latency_time_units := (tFin - tIniTyped)
			LogDebug("{ \"event\":\"%v\", \"trx_id\":\"%v\", \"trx_number\":\"%v\",  \"latency\": %v  } ", event.Arg0, event.Arg0, event.Arg1, latency_time_units)
			stat.StoreLatencySample(latency_time_units)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}