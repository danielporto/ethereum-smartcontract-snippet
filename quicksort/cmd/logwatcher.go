package cmd

import (
	"context"
	"encoding/hex"
	"strings"

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

func watchArrayInfoEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintArrayInfo")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (PrintArrayInfo)", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintArrayInfo)
			if err := bc.UnpackLog(event, "PrintArrayInfo", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[PrintArrayInfo Event]: Size: %v, data hash: 0x%v, Elements: %v", event.Arg0, hex.EncodeToString(event.Arg1[:]), event.Arg2[:])
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchHashEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintHash")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (PrintHash)", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintHash)
			if err := bc.UnpackLog(event, "PrintHash", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}

			log.Infof("[PrintHash Event]: data hash: 0x%v", hex.EncodeToString(event.Arg0[:]))
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchSetElementQtyEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintSetElementQty")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (PrintSetElementQty)", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintSetElementQty)
			if err := bc.UnpackLog(event, "PrintSetElementQty", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[PrintSetElementQty Event]: Size: %v, data hash: 0x%v, SetEntryIndex: %v, TotalSetEntries %v, Operations: %v", event.Arg0, hex.EncodeToString(event.Arg1[:]), event.Arg2, event.Arg3, event.Arg4)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchArrayEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.QuickSortABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)
	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{nil, context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintArray")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (PrintArray)", err)
	}
	for {
		select {
		case logentry := <-logs:
			event := new(contracts.QuickSortPrintArray)
			if err := bc.UnpackLog(event, "PrintArray", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[PrintArray Event]: Size: %v, Elements: %v", event.Arg0, event.Arg1[:])
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}
