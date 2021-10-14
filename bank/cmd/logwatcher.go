package cmd

import (
	"context"
	"github.com/danielporto/ethereum-smartcontract-snippet/bank/contracts"
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

func watchOperationsExecutedEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.BankMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "OperationsExecuted")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (OperationsExecuted)", err)
	}
	log.Infof("Logging thread: Listening for OperationsExecuted events")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.BankOperationsExecuted)
			if err := bc.UnpackLog(event, "OperationsExecuted", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[BankOperationsExecuted Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}


func watchBalanceTransferredEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.BankMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "BalanceTransferred")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (BalanceTransferred)", err)
	}
	log.Infof("Logging thread: Listening for BalanceTransferred events")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.BankBalanceTransferred)
			if err := bc.UnpackLog(event, "BalanceTransferred", logentry); err != nil {
				log.Error("Error decoding log event:", logentry, err)
			}
			log.Infof("[BalanceTransferred Event]: Value: %v", event.Arg0)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchBalanceReceivedEvents(client *ethclient.Client,  contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.BankMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "BalanceReceived")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (BalanceReceived)", err)
	}
	log.Infof("Logging thread: Listening to BalanceReceived events")

	for {
		select {
		case logentry := <-logs:
			event := new(contracts.BankBalanceReceived)
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
