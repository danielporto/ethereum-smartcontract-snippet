package cmd

import (
	"context"
	"github.com/danielporto/ethereum-smartcontract-snippet/counter/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"strings"
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
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "GetValue")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (GetValue)", err)
	}
	log.Infof("Logging thread: Listening for GetValueEvents")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.CounterGetValue)
			if err := bc.UnpackLog(event, "GetValue", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[GetValue Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}


func watchIncrementEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.CounterABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "Increment")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (Increment)", err)
	}
	log.Infof("Logging thread: Listening for IncrementEvents")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.CounterIncrement)
			if err := bc.UnpackLog(event, "Increment", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[Increment Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchDecrementEvents(client *ethclient.Client,  contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.CounterABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "Decrement")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (Decrement)", err)
	}
	log.Infof("Logging thread: Listening to DecrementEvents")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.CounterDecrement)
			if err := bc.UnpackLog(event, "Decrement", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[Decrement Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}
