package cmd

import (
	"context"
	"encoding/hex"
	"strings"
	"sync"
	"time"

	"github.com/danielporto/ethereum-smartcontract-snippet/quicksort/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

/*
* functions in this file use the stub generated from the solidity contract
* therefore, if the contract is changed, this file will likelly require updates to reflect
* the updated version of datastructures and events.
 */

func watchPrintArrayInfo(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	eventName := "PrintArrayInfo"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events ("+eventName+")", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintArrayInfo)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			//log.Infof("[%v Event]: Size: %v, data hash: 0x%v, Elements: %v", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]), event.Arg2[:])
			log.Infof("{ \"event\": \"%v\", \"size\": \"%v\", \"data_hash\": \"%v\", \"elements\": \"%v\" }",  eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]), event.Arg2[:])

		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchPrintHash(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	eventName := "PrintHash"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events ("+eventName+")", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintHash)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}

			log.Infof("[%v Event]: data hash: 0x%v", eventName, hex.EncodeToString(event.Arg0[:]))
			log.Infof("{ \"event\": \"%v\", \"data_hash\": \"%v\" }",  eventName, hex.EncodeToString(event.Arg0[:]))

		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchPrintSetElementQty(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	eventName := "PrintSetElementQty"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events ("+eventName+")", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintSetElementQty)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			//log.Infof("[%v Event]: Size: %v, data hash: 0x%v, SetEntryIndex: %v, TotalSetEntries %v, Operations: %v", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]), event.Arg2, event.Arg3, event.Arg4)
			log.Infof("{ \"event\": \"%v\", \"size\": \"%v\", \"data_hash\": \"%v\", \"set_entry_index\": \"%v\", \"total_set_entries\": \"%v\", \"operations\": \"%v\" }",  eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]), event.Arg2, event.Arg3, event.Arg4)

		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchPrintArray(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	eventName := "PrintArray"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events ("+eventName+")", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintArray)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			//log.Infof("[%v Event]: Size: %v, Elements: %v", eventName, event.Arg0, event.Arg1[:])
			log.Infof("{ \"event\": \"%v\", \"size\": \"%v\", \"elements\": \"%v\" }", eventName, event.Arg0, event.Arg1[:])

		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchPrintConfirmation(client *ethclient.Client, contractAddr common.Address, stop chan struct{}, reqNanotimeMap *sync.Map, stat *StatsStorage) {
	eventName := "PrintConfirmation"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events ("+eventName+")", err)
	}
	for {
		select {
		case logentry := <-logs:
			tFin := time.Now().UnixNano() // get the timestamp in nanosseconds
			event := new(contracts.QuickSortPrintConfirmation)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			//log.Infof("[%v Event]: trx_id=%v  ; trx_hash: %v", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]))
			//log.Infof("{ \"event\": \"%v\", \"trx_id\": \"%v\", \"trx_hash\": \"%v\" }", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]))
			tIni, ok := reqNanotimeMap.LoadAndDelete(event.Arg0) // the value is no longer needed, release memory
			if !ok {
				// reply was received without a request. breaks causality.
				log.Error("Reply was received without a request. breaks causality:"+ "{ \"event\": \"%v\", \"trx_id\": \"%v\", \"trx_hash\": \"%v\" }", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]))
			}
			tIniTyped, ok := tIni.(int64)
			if !ok {
				log.Error("Unable to convert timestamp type, ignoring entry")
				continue
			}
			lat := tFin - tIniTyped
			log.Infof("{ \"event\": \"%v\", \"trx_id\": \"%v\", \"trx_hash\": \"%v\", latency: %v }", eventName, event.Arg0, hex.EncodeToString(event.Arg1[:]), lat)
			stat.StoreLatencySample(lat)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchPrintConfirmationDebug(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	eventName := "PrintConfirmationDebug"
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, eventName)
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events ("+eventName+")", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintConfirmationDebug)
			if err := bc.UnpackLog(event, eventName, logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("{ \"event\": \"%v\", \"trx_id\": \"%v\", \"size\": \"%v\", \"data_hash\": \"%v\", \"elements\": \"%v\" }",  eventName, event.Arg0, event.Arg1, hex.EncodeToString(event.Arg2[:]), event.Arg2[:])

		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}