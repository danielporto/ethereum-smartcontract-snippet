// Package cmd
/*
Copyright © 2021 DANIEL PORTO <daniel.porto@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/danielporto/ethereum-smartcontract-snippet/ycsb/contracts"
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

func watchPrintKVAckEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.KVstoreMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintKVAck")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (PrintKVAck) %v", err)
	}
	Log("Logging thread: Listening for PrintKVAck events")

	for {
		select {
		case logEntry := <-logs:
			event := new(contracts.KVstorePrintKVAck)
			if err := bc.UnpackLog(event, "PrintKVAck", logEntry); err != nil {
				LogError("Error decoding log event:", logEntry, err)
			}
			Log("[PrintKVAck Event] op(k->v): %v (%v->%v)", event.Arg0, event.Arg1, event.Arg2)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel: %v", err)
		case <-stop:
			return
		}
	}
}

func watchPrintInsertsEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.KVstoreMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract: %v", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintInserts")
	defer sub.Unsubscribe()

	if err != nil {
		LogFatal("Error subscribing to contract Events (PrintInserts) %v", err)
	}
	Log("Logging thread: Listening for PrintInserts events")

	for {
		select {
		case logEntry := <-logs:
			event := new(contracts.KVstorePrintInserts)
			if err := bc.UnpackLog(event, "PrintInserts", logEntry); err != nil {
				LogError("Error decoding log event:", logEntry, err)
			}
			Log("[PrintInserts Event]: %v", event.Arg0)
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
	parsed, err := abi.JSON(strings.NewReader(contracts.KVstoreMetaData.ABI))
	if err != nil {
		LogFatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for  events
	watchOpts := bind.WatchOpts{nil, context.Background()}
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
			event := new(contracts.KVstorePrintConfirmation)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				LogError("Error decoding log event:", logentry, err)
			}
			//Log("[%v Event]: trx_id=%v  ; trx_hash: %v", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]))
			//Log("{ \"event\": \"%v\", \"trx_id\": \"%v\", \"trx_hash\": \"%v\" }", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]))
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
			LogDebug("{ \"event\": \"%v\", \"trx_id\": \"%v\", \"trx_number\": \"%v\", \"latency\": %v }", eventName, event.Arg0, event.Arg1, latency_time_units)
			stat.StoreLatencySample(latency_time_units)
		case err := <-sub.Err():
			LogError("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}