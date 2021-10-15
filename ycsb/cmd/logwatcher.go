// Package cmd
/*
Copyright © 2021 DANIEL PORTO <daniel.porto>

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

	"github.com/danielporto/ethereum-smartcontract-snippet/ycsb/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
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
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintKVAck")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (PrintKVAck)", err)
	}
	log.Infof("Logging thread: Listening for PrintKVAck events")

	for {
		select {
		case logEntry := <-logs:
			event := new(contracts.KVstorePrintKVAck)
			if err := bc.UnpackLog(event, "PrintKVAck", logEntry); err != nil {
				log.Error("Error decoding log event:", logEntry, err)
			}
			log.Infof("[PrintKVAck Event] op(k->v): %v (%v->%v)", event.Arg0, event.Arg1, event.Arg2)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}

func watchPrintInsertsEvents(client *ethclient.Client, contractAddr common.Address, stop chan struct{}) {
	// get a reference to the contract abi
	parsed, err := abi.JSON(strings.NewReader(contracts.KVstoreMetaData.ABI))
	if err != nil {
		log.Fatal("Unable to get the ABI for the contract:", err)
	}

	// get a reference to the *bind.BoundContract of the existing contract
	bc := bind.NewBoundContract(contractAddr, parsed, client, client, client)

	// configure a watcher for PrintHash events
	watchOpts := bind.WatchOpts{Context: context.Background()}
	logs, sub, err := bc.WatchLogs(&watchOpts, "PrintInserts")
	defer sub.Unsubscribe()

	if err != nil {
		log.Fatal("Error subscribing to contract Events (PrintInserts)", err)
	}
	log.Infof("Logging thread: Listening for PrintInserts events")

	for {
		select {
		case logEntry := <-logs:
			event := new(contracts.KVstorePrintInserts)
			if err := bc.UnpackLog(event, "PrintInserts", logEntry); err != nil {
				log.Error("Error decoding log event:", logEntry, err)
			}
			log.Infof("[PrintInserts Event]: %v", event.Arg0)
		case err := <-sub.Err():
			log.Error("Error received in the subscription channel:", err)
		case <-stop:
			return
		}
	}
}
