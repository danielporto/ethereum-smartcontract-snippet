/*
Copyright © 2021 Daniel Porto <daniel.porto@gmail.com>

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
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/danielporto/ethereum-smartcontract-snippet/quicksort/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"go.uber.org/ratelimit"
)

var operation string
var count int
var maxrate int // using for suggesting a workload value in tps
var checkpoint int
var timeline bool
var totalRequests int64

var requestNanotimeMap sync.Map
var stats StatsStorage

// workloadCmd represents the workload command
var workloadCmd = &cobra.Command{
	Use:   "workload",
	Short: "Deploy and execute contract operations",
	Long: `Simulate a client by first deploying a contract on the network node
Then, call several operations of that contract. `,
	Run: func(cmd *cobra.Command, args []string) {

		Log("Workload operation '%v'", operation)
		Log("Gas price limit for the transaction: %v wei", trxgaslimit)
		Log("Max rate to issue operations (suggested tps) %v tps", maxrate)
		switch operation {
		case "deploy":
			workloadDeploy()
		case "sort":
			Log("Size of the array to sort: %v", payloadsize)
			workloadQuicksort()
		case "mixed":
			Log("Size of the array to sort: %v", payloadsize)
			workloadMixed()
		default:
			LogFatal("Operation not supported: %v", operation)

		}
	},
}

func init() {
	rootCmd.AddCommand(workloadCmd)

	// Here you will define your flags and configuration settings.
	workloadCmd.PersistentFlags().StringVarP(&operation, "operation", "o", "sort", "Issues operations of that type (deploy/sort/mixed) to the blockchain")
	workloadCmd.PersistentFlags().IntVarP(&count, "count", "c", 1, "Number of operations to be issued")
	workloadCmd.PersistentFlags().IntVarP(&duration, "duration", "d", 0, "Duration of the experiment in seconds")
	// this is currently not supporting multi threads. for that we need multiple keys to generate
	// nonces independently per key otherwise the will queue up after a max rate.
	// instead, start multiple processes passing a different key
	//	workloadCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 1, "Number of threads")
	threads=1
	workloadCmd.PersistentFlags().IntVarP(&maxrate, "rate", "r", 10, "Max operations per second for each thread (suggested value)")
	workloadCmd.PersistentFlags().IntVarP(&checkpoint, "checkpoint", "q", 5000, "Print total operations after X operations issued.")
	workloadCmd.PersistentFlags().StringVarP(&client_id, "id", "", "undefined", "client identifier")
	workloadCmd.PersistentFlags().BoolVarP(&timeline, "timeline", "", false, "print a timeline every second")
}

/*
* Generate a sequence of nonces up until a duration or a defined number
 */
func generateNonce(init uint64, count, duration int, nonces chan<- uint64) {
	nonce := init
	if duration > 0 {
		for end := time.Now().Add(time.Second * time.Duration(duration)); ; {
			nonce++
			if time.Now().After(end) {
				break
			}
			nonces <- nonce
		}

	} else {
		for i := 0; i < count; i++ {
			nonce++
			nonces <- nonce
		}
	}
	close(nonces)
}

func generateNonceAtRate(client *ethclient.Client, fromAddress common.Address, count, duration int, nonces chan<- uint64, rate int) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		LogFatal("Impossible to get a nonce for account: %v", err)
	}

	var rl ratelimit.Limiter
	if rate > 0 {
		rl = ratelimit.New(rate) // per second
		Log("[Nonce]: Generate nonces starting from %v at a rate of: %v ops/s", nonce, rate)
	}else{
		rl = ratelimit.NewUnlimited() //disable
		Log("[Nonce]: Generate nonces starting from %v at maximum rate", nonce)
	}

	if duration > 0 {
		lastTimelinePrint := int64(-1)
		for end := time.Now().Add(time.Second * time.Duration(duration)); ; {
			sendInstant := time.Now()
			if sendInstant.After(end) {
				break
			}
			rl.Take()
			nonces <- nonce
			nonce++
			if timeline {
				if sendInstant.UnixNano() -lastTimelinePrint >= 1_000_000_000 {
					lastTimelinePrint = sendInstant.UnixNano()
					Log(stats.ReportStats())
				}
			}
		}

	} else {
		for i := 0; i < count; i++ {
			rl.Take()
			nonces <- nonce
			nonce++

		}
	}
	close(nonces)
}

/*
* Monitor all events of this contract
 */
func watchContractEvents(contractAddr common.Address, done chan struct{}) {
	stop := make(chan struct{})
	defer close(stop)
	p, err := strconv.Atoi(port)
	if err != nil {
		LogFatal("Error converting the socket port[%v]: %v", port, err)
	}
	wsurl := "ws://" + host + ":" + strconv.Itoa(p)

	client, err := ethclient.Dial(wsurl)
	if err != nil {
		LogFatal("Error opening websocket connection to host[%v]: %v", wsurl, err)
	}
	Log("Operations report checkpoint: %v ops", checkpoint)
	Log("Logging thread: Subscribing to contract events: %v", wsurl)
	// initialize the monitor threads for each event
	go watchPrintArray(client, contractAddr, stop)
	go watchPrintArrayInfo(client, contractAddr, stop)
	go watchPrintHash(client, contractAddr, stop)
	go watchPrintSetElementQty(client, contractAddr, stop)
	go watchPrintConfirmation(client, contractAddr, stop, &requestNanotimeMap, &stats)
	go watchPrintConfirmationDebug(client, contractAddr, stop)

	LogDebug("Logging thread: Waiting for logs to be closed")
	// wait for the signal to stop
	<-done
	Log("Logging thread: Teardown subscription to contract events.")

	//stop all 6 logging threads
	<-stop
	<-stop
	<-stop
	<-stop
	<-stop
	<-stop
	Log("Logging thread: All event watchers are closed.")
}

/**********************************************************************************************************************
* The deploy workload issue operations to install the smartcontract on the blockchain.
* This operation can be expensive, be aware of the funds to execute this operation
***********************************************************************************************************************/
func deploy(pk *ecdsa.PrivateKey, c *ethclient.Client, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()
	total_transactions := 0
	Log("Thread %v STARTED - issuing deploy contract transactions", threadid)

	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(trxgaslimit)
		auth.GasPrice = gasPrice

		address, _, _, err := contracts.DeployQuickSort(auth, c)
		if err != nil {
			LogFatal("Error deploying quicksort contract: %v", err)
		}

		total_transactions++
		if total_transactions%checkpoint == 0 {
			Log("Thread %v - deploy transactions issued : %v", threadid, total_transactions)
		}
		LogDebug("Transaction address: %v", address.Hex())
	}
	Log("Thread %v FINISHED - deploy contract transactions issued : %v", threadid, total_transactions)

}

func workloadDeploy() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	Log("Running deploy workload")
	Log("Connecting to ethereum network: %v", url)

	conn, err := ethclient.Dial(url)
	if err != nil {
		LogFatal("Failed to connect to ethereum node: %v", err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		LogFatal("Error converting the private key from Hex to ECDSA: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		LogFatal("Error casting public key to ECDSA")
	}

	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//4. configure gasPrice
	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		LogFatal("Error while trying to get the gas price: %v", err)
	}

	go generateNonceAtRate(conn, fromAddress, count, duration, nonceStream, maxrate)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go deploy(privateKey, conn, gasPrice, nonceStream, &wg, i)
	}
	wg.Wait()
	Log("Experiment FINISHED, wait for the 10 seconds to ensure pending transactions are processed")
	time.Sleep(10 * time.Second)

}

/**********************************************************************************************************************
* The quicksort workload install a single contract and issue operations to sort an array of a given size and store the result
* It also emits a log with a hash of the sorted array. At the end a report of the execution is also emitted as log.
* The logs allow for detecting byzantine faults that could be possible not captured.
***********************************************************************************************************************/
func quicksort(pk *ecdsa.PrivateKey, c *ethclient.Client, instance *contracts.QuickSort, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()
	//total_transactions := 0
	Log("Thread %v STARTED - issuing quicksort transactions", threadid)

	var sort func(opts *bind.TransactOpts, size *big.Int, id string) (*types.Transaction, error)
	// pick the right sort function
	if verbosity == "debug" {
		sort = instance.DebugSort
	} else {
		sort = instance.Sort
	}


	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(trxgaslimit)
		auth.GasPrice = gasPrice

		if nonce%uint64(checkpoint) == 0 {
			Log("Thread %v - quicksort transactions issued : %v", threadid, atomic.LoadInt64(&totalRequests))
			instance.PrintAllData(auth)
			continue // get another nonce
		}

		id := fmt.Sprintf("%v_tx_%v", client_id, nonce)
		tIni_us := time.Now().UnixNano() / latency_factor // get the timestamp in microsseconds
		tx, err := sort(auth, big.NewInt(int64(payloadsize)), id)
		if err != nil {
			LogFatal("Failed to call sort transaction method of quicksort contract. Check the Gas limit [%v] for this transaction. Detail: %v", auth.GasLimit, err)
		}
		requestNanotimeMap.Store(id, tIni_us) //stores the timestamp in which the request was made (this will be updated by the event function)

		atomic.AddInt64(&totalRequests,1)
		LogDebug("nonce %v, tx sent: %s", nonce, tx.Hash().Hex())

	}
}

func workloadQuicksort() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	Log("Running quicksort workload")
	Log("Connecting to ethereum network: %v", url)

	client, err := ethclient.Dial(url)
	if err != nil {
		LogFatal("Failed to connect to ethereum node: %v", err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		LogFatal("Error converting the private key from Hex to ECDSA: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		LogFatal("Error casting public key to ECDSA")
	}

	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//3. configure nonce (prevent replay attacks with a user specific nonce)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		LogFatal("Impossible to get a nonce for account: %v", err)
	}
	//4. configure gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		LogFatal("Unable to get a gas price: %v", err)
	}

	// deploy a new contract
	// 4. setup an authenticated transactor with info from credentials and connection configuration
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)          // not transferring funds, just executing contract operation
	auth.GasLimit = uint64(trxgaslimit) // max value for this transaction
	auth.GasPrice = gasPrice
	contractAddr, _, instance, err := contracts.DeployQuickSort(auth, client)
	if err != nil {
		LogFatal("Impossible to initialize a quicksort contract for this workload: %v", err)
	}
	Log("Wait for the 5 seconds (blocks) while contract is be mined before issuing operations.")
	time.Sleep(5 * time.Second)

	doneWatchingLogs := make(chan struct{})

	if !disable_events {
		go watchContractEvents(contractAddr, doneWatchingLogs)
	} else {
		Log("Event watchers are disabled.")
	}

	// use the contract to run the quicksort workload
	go generateNonceAtRate(client, fromAddress, count, duration, nonceStream, maxrate)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go quicksort(privateKey, client, instance, gasPrice, nonceStream, &wg, i)
	}
	wg.Wait()

	// time's up or #operations has finished
	Log("Experiment FINISHED, wait for the 10 seconds to ensure pending transactions are processed")
	time.Sleep(10 * time.Second)
	lastnonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		LogFatal("Impossible to get a nonce for final report: %v", err)
	}
	auth.Nonce = big.NewInt(int64(lastnonce))
	instance.PrintAllData(auth)
	Log("Wait 10 seconds to receive the final log")
	time.Sleep(10 * time.Second)
	close(doneWatchingLogs)
	Log(stats.ReportStats())
	unconfirmed := int64(0)
	requestNanotimeMap.Range(func(key, value interface{}) bool {
		unconfirmed++
		return true
	})
	Log("Transactions issued : %v, unconfirmed: %v", atomic.LoadInt64(&totalRequests), unconfirmed)

	//wait for logs to be closed
	time.Sleep(10 * time.Second)
	stats.PrintStatsMap()

}

/**********************************************************************************************************************
* The mixed workload alternates with operations that install smartcontracts and also execute sort operations
***********************************************************************************************************************/
func mixed(pk *ecdsa.PrivateKey, c *ethclient.Client, instance *contracts.QuickSort, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()

	Log("Thread %v STARTED - issuing mix deploy/quicksort transactions", threadid)

	var sort func(opts *bind.TransactOpts, size *big.Int, id string) (*types.Transaction, error)
	// pick the right sort function
	if verbosity == "debug" {
		sort = instance.DebugSort
	} else {
		sort = instance.Sort
	}

	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(trxgaslimit)
		auth.GasPrice = gasPrice

		if nonce%uint64(checkpoint) == 0 {
			Log("Thread %v - mix deploy/quicksort transactions issued: %v", threadid, atomic.LoadInt64(&totalRequests))
			instance.PrintAllData(auth)
			continue // get another nonce
		}

		if nonce%2 == 0 {
			address, _, _, err := contracts.DeployQuickSort(auth, c)
			if err != nil {
				LogFatal("Error deploying quicksort contract: %v", err)
			}
			LogDebug("Transaction address: %v", address.Hex())
		} else {
			id := fmt.Sprintf("%v_tx_%v",  client_id,nonce)
			tx, err := sort(auth, big.NewInt(int64(payloadsize)), id)
			if err != nil {
				LogFatal("Failed to call transaction method: ", err)
			}
			LogDebug("nonce %v, tx sent: %s", nonce, tx.Hash().Hex())
		}
		atomic.StoreInt64(&totalRequests, 1)

	}
}

func workloadMixed() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	Log("Running mix deploy/quicksort workload")
	Log("Connecting to ethereum network: %v", url)

	client, err := ethclient.Dial(url)
	if err != nil {
		LogFatal("Failed to connect to ethereum node: %v", err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elliptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		LogFatal("Error converting the private key from Hex to ECDSA: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		LogFatal("Error casting public key to ECDSA")
	}
	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//3. configure nonce (prevent replay attacks with a user specific nonce)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		LogFatal("Impossible to get a nonce for account: %v", err)
	}
	//4. configure gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		LogFatal("Unable to get a gas price: %v", err)
	}

	// deploy a new contract
	// 4. setup an authenticated transactor with info from credentials and connection configuration
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)          // not transferring funds, just executing contract operation
	auth.GasLimit = uint64(trxgaslimit) // max value for this transaction
	auth.GasPrice = gasPrice
	contractAddr, _, instance, err := contracts.DeployQuickSort(auth, client)
	if err != nil {
		LogFatal("Impossible to initialize a quicksort contract for this workload: %v", err)
	}
	Log("Wait for the 5 seconds (blocks) while contract is be mined before issuing operations.")
	time.Sleep(5 * time.Second)

	doneWatchingLogs := make(chan struct{})

	if !disable_events {
		go watchContractEvents(contractAddr, doneWatchingLogs)
	} else {
		Log("Event watchers are disabled.")
	}

	// use the contract to run the increment workload
	go generateNonceAtRate(client, fromAddress, count, duration, nonceStream, maxrate)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go mixed(privateKey, client, instance, gasPrice, nonceStream, &wg, i)
	}
	wg.Wait()

	// time's up or #operations has finished
	Log("Experiment FINISHED, wait for the 10 seconds to ensure pending transactions are processed")
	time.Sleep(10 * time.Second)
	lastnonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		LogFatal("Impossible to get a nonce for final report: %v", err)
	}
	auth.Nonce = big.NewInt(int64(lastnonce))
	instance.PrintAllData(auth)
	Log("Wait 10 seconds to receive the final log")
	time.Sleep(10 * time.Second)
	close(doneWatchingLogs)
	//wait for logs to be closed
	time.Sleep(10 * time.Second)
}
