package cmd

import (
	"context"
	"github.com/danielporto/ethereum-smartcontract-snippet/counter/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
	"sync"
	"time"
)


/*
* functions in this file use the stub generated from the solidity contract
* therefore, if the contract is changed, this file will likely require updates to reflect
* the updated version of datastructures and events.
 */

func watchGetValueEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.CounterABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract: %v", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "GetValue")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (GetValue) %v", err)
	}
	Log("Logging thread: Listening for GetValueEvents")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.CounterGetValue)
			if err := bc.UnpackLog(event, "GetValue", logentry); err != nil {
				LogError("Error decoding log event: %v, %v ", logentry, err)
			}
			Log("[GetValue Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel: %v", err)
		case <-stop:
			return
		}
	}
}


func watchIncrementEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.CounterABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract: %v", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "Increment")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (Increment) %v", err)
	}
	Log("Logging thread: Listening for IncrementEvents")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.CounterIncrement)
			if err := bc.UnpackLog(event, "Increment", logentry); err != nil {
				LogError("Error decoding log event: %v, %v", logentry, err)
			}
			Log("[Increment Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel: %v", err)
		case <-stop:
			return
		}
	}
}

func watchDecrementEvents(client *ethclient.Client,  contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.CounterABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract: %v", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "Decrement")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (Decrement)  %v", err)
	}
	Log("Logging thread: Listening to DecrementEvents")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.CounterDecrement)
			if err := bc.UnpackLog(event, "Decrement", logentry); err != nil {
				LogError("Error decoding log event: %v, %v", logentry, err)
			}
			Log("[Decrement Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel: %v", err)
		case <-stop:
			return
		}
	}
}

func watchPrintConfirmation(client *ethclient.Client, contractAddr common.Address, stop chan struct{}, requestTimeMap *sync.Map, stat *StatsStorage) {
	eventName := "PrintConfirmation"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.CounterMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract: %v", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events ("+eventName+"): %v", err)
	}
	Log("Logging thread: Listening for "+eventName+" events")

	for {
		select {
		case logentry := <-logs:
			tFin := time.Now().UnixNano() / latency_factor // check latency_factor_unity
			event := new(contracts.CounterPrintConfirmation)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				LogError("Error decoding [%v] log event: %v", logentry, err)
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
			LogDebug("{ \"event\":\"%v\", \"trx_id\":\"%v\", \"trx_number\":\"%v\",  \"latency\": %v  } ", eventName, event.Arg0, event.Arg1, latency_time_units)
			stat.StoreLatencySample(latency_time_units)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel: %v", err)
		case <-stop:
			return
		}
	}
}