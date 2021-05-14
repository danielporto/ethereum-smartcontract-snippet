/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"math/big"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

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

// workloadCmd represents the workload command
var workloadCmd = &cobra.Command{
	Use:   "workload",
	Short: "Deploy and execute contract operations",
	Long: `Simulate a client by first deploying a contract on the network node
Then, call several operations of that contract. `,
	Run: func(cmd *cobra.Command, args []string) {

		log.Infof("Workload operation '%v'", operation)
		log.Infof("Gas price limit for the transaction: %v wei", trxgaslimit)
		log.Infof("Max rate to issue operations (suggested tps) %v tps", maxrate)
		switch operation {
		case "deploy":
			workloadDeploy()
		case "sort":
			log.Infof("Size of the array to sort: %v", payloadsize)
			workloadQuicksort()
		case "mixed":
			log.Infof("Size of the array to sort: %v", payloadsize)
			workloadMixed()
		default:
			log.Fatal("Operation not supported:", operation)

		}
	},
}

func init() {
	rootCmd.AddCommand(workloadCmd)

	// Here you will define your flags and configuration settings.
	workloadCmd.PersistentFlags().StringVarP(&operation, "operation", "o", "sort", "Issues operations of that type (deploy/sort/mixed) to the blockchain")
	workloadCmd.PersistentFlags().IntVarP(&count, "count", "c", 1, "Number of operations to be issued")
	workloadCmd.PersistentFlags().IntVarP(&duration, "duration", "d", 0, "Duration of the experiment in seconds")
	workloadCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 1, "Number of threads")
	workloadCmd.PersistentFlags().IntVarP(&maxrate, "rate", "r", 100, "Max operations per second for each thread (suggested value)")
	workloadCmd.PersistentFlags().IntVarP(&checkpoint, "checkpoint", "q", 5000, "Print total operations after X operations issued.")

}

/*
* Generate a sequence of nonces up until a duration or a defined number
 */
func generateNonce(init uint64, count, duration int, nonces chan<- uint64) {
	nonce := init
	if duration > 0 {
		for end := time.Now().Add(time.Second * time.Duration(duration)); ; {
			if time.Now().After(end) {
				break
			}
			nonces <- nonce
			nonce++
		}
	} else {
		for i := 0; i < count; i++ {
			nonces <- nonce
			nonce++
		}
	}
	close(nonces)
}

func generateNonceAtRate(client *ethclient.Client, fromAddress common.Address, count, duration int, nonces chan<- uint64, rate int) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Impossible to get a nonce for account", err)
	}

	if rate <= 0 {
		log.Fatal("Rate must be greater than 0")
	}

	rl := ratelimit.New(rate) // per second
	log.Infof("[Nonce]: Generate nonces starting from %v at a rate of: %v ops/s", nonce, rate)

	if duration > 0 {
		for end := time.Now().Add(time.Second * time.Duration(duration)); ; {
			if time.Now().After(end) {
				break
			}
			rl.Take()
			nonces <- nonce
			nonce++

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
		log.Fatal("Error converting the socket port:", port, err)
	}
	wsurl := "ws://" + host + ":" + strconv.Itoa(p)

	client, err := ethclient.Dial(wsurl)
	if err != nil {
		log.Fatal("Error opening websocket connection to host.", wsurl, err)
	}
	log.Infof("Operations report checkpoint: %v ops", checkpoint)
	log.Infof("Logging thread: Subscribing to contract events: %v", wsurl)
	// initialize the monitor threads for each event
	go watchArrayEvents(client, contractAddr, stop)
	go watchArrayInfoEvents(client, contractAddr, stop)
	go watchHashEvents(client, contractAddr, stop)
	go watchSetElementQtyEvents(client, contractAddr, stop)

	log.Debug("Logging thread: Waiting for logs to be closed")
	// wait for the signal to stop
	<-done
	log.Info("Logging thread: Teardown subscription to contract events.")

	//stop all 4 logging threads
	<-stop
	<-stop
	<-stop
	<-stop
	log.Info("Logging thread: All event watchers are closed.")
}

/**********************************************************************************************************************
* The deploy workload issue operations to install the smartcontract on the blockchain.
* This operation can be expensive, be aware of the funds to execute this operation
***********************************************************************************************************************/
func deploy(pk *ecdsa.PrivateKey, c *ethclient.Client, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()
	total_transactions := 0
	log.Infof("Thread %v STARTED - issuing deploy contract transactions", threadid)

	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(trxgaslimit)
		auth.GasPrice = gasPrice

		address, _, _, err := contracts.DeployQuickSort(auth, c)
		if err != nil {
			log.Fatal("Error deploying quicksort contract", err)
		}

		total_transactions++
		if total_transactions%checkpoint == 0 {
			log.Infof("Thread %v - deploy transactions issued : %v", threadid, total_transactions)
		}
		log.Debugf("Transaction address: %v\n", address.Hex())
	}
	log.Infof("Thread %v FINISHED - deploy contract transactions issued : %v", threadid, total_transactions)

}

func workloadDeploy() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	log.Println("Running deploy workload")
	log.Println("Connecting to ethereum network...")

	conn, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to ethereum node", err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal("Error converting the private key from Hex to ECDSA", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//4. configure gasPrice
	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error while trying to get the gas price: ", err)
	}

	go generateNonceAtRate(conn, fromAddress, count, duration, nonceStream, maxrate)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go deploy(privateKey, conn, gasPrice, nonceStream, &wg, i)
	}
	wg.Wait()
	log.Info("Experiment FINISHED, wait for the 10 seconds to ensure pending transactions are processed")
	time.Sleep(10 * time.Second)

}

/**********************************************************************************************************************
* The quicksort workload install a single contract and issue operations to sort an array of a given size and store the result
* It also emits a log with a hash of the sorted array. At the end a report of the execution is also emitted as log.
* The logs allow for detecting byzantine faults that could be possible not captured.
***********************************************************************************************************************/
func quicksort(pk *ecdsa.PrivateKey, c *ethclient.Client, instance *contracts.QuickSort, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()
	total_transactions := 0
	log.Infof("Thread %v STARTED - issuing quicksort transactions", threadid)

	var sort func(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error)
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
			log.Infof("Thread %v - quicksort transactions issued : %v", threadid, total_transactions)
			instance.PrintAllData(auth)
			continue // get another nonce
		}

		tx, err := sort(auth, big.NewInt(int64(payloadsize)))
		if err != nil {
			log.Fatal("Failed to call sort transaction method of quicksort contract. Check the gaslimit for this transaction:", auth.GasLimit, " err:", err)
		}


		total_transactions++
		log.Debugf("nonce %v, tx sent: %s\n", nonce, tx.Hash().Hex())

	}
	log.Infof("Thread %v FINISHED - quicksort transactions issued : %v", threadid, total_transactions)

}

func workloadQuicksort() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	log.Info("Running increment workload")
	log.Infof("Connecting to ethereum network: %v", url)

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal("Error converting the private key from Hex to ECDSA", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}
	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//3. configure nonce (prevent replay attacks with a user specific nonce)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Impossible to get a nonce for account", err)
	}
	//4. configure gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Unable to get a gas price", err)
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
		log.Fatal("Impossible to initialize the quicksort contract for this workload.", err)
	}
	log.Info("Wait for the 5 seconds (blocks) while contract is be mined before issuing operations.")
	time.Sleep(5 * time.Second)

	doneWatchingLogs := make(chan struct{})

	if !disable_events {
		go watchContractEvents(contractAddr, doneWatchingLogs)
	} else {
		log.Infof("Event watchers are disabled.")
	}

	// use the contract to run the quicksort workload
	go generateNonceAtRate(client, fromAddress, count, duration, nonceStream, maxrate)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go quicksort(privateKey, client, instance, gasPrice, nonceStream, &wg, i)
	}
	wg.Wait()

	// time's up or #operations has finished
	log.Info("Experiment FINISHED, wait for the 10 seconds to ensure pending transactions are processed")
	time.Sleep(10 * time.Second)
	lastnonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Impossible to get a nonce for final report", err)
	}
	auth.Nonce = big.NewInt(int64(lastnonce))
	instance.PrintAllData(auth)
	log.Info("Wait 10 seconds to receive the final log")
	time.Sleep(10 * time.Second)
	close(doneWatchingLogs)
	//wait for logs to be closed
	time.Sleep(10 * time.Second)
}

/**********************************************************************************************************************
* The mixed workload alternates with operations that install smartcontracts and also execute sort operations
***********************************************************************************************************************/
func mixed(pk *ecdsa.PrivateKey, c *ethclient.Client, instance *contracts.QuickSort, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()
	total_transactions := 0

	log.Infof("Thread %v STARTED - issuing quicksort/deploy transactions", threadid)

	var sort func(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error)
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

		if total_transactions%checkpoint == 0 {
			log.Infof("Thread %v - mix quicksort/deploy transactions issued : %v", threadid, total_transactions)
			instance.PrintAllData(auth)
			total_transactions++
			continue // get another nonce
		}

		if nonce%2 == 0 {
			address, _, _, err := contracts.DeployQuickSort(auth, c)
			if err != nil {
				log.Fatal("Error deploying quicksort", err)
			}
			log.Debugf("Transaction address: %v\n", address.Hex())
		} else {
			tx, err := sort(auth, big.NewInt(int64(payloadsize)))
			if err != nil {
				log.Fatal("Failed to call transaction method: ", err)
			}
			log.Debugf("nonce %v, tx sent: %s\n", nonce, tx.Hash().Hex())
		}
		total_transactions++

	}
	log.Infof("Thread %v FINISHED - mix quicksort/deploy transactions issued : %v", threadid, total_transactions)

}

func workloadMixed() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	log.Info("Running mix quicksort/deploy workload")
	log.Infof("Connecting to ethereum network: %v", url)

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal("Error converting the private key from Hex to ECDSA", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}
	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//3. configure nonce (prevent replay attacks with a user specific nonce)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Impossible to get a nonce for account", err)
	}
	//4. configure gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Unable to get a gas price", err)
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
		log.Fatal("Impossible to initialize a quicksort contract for this workload.", err)
	}
	log.Info("Wait for the 5 seconds (blocks) while contract is be mined before issuing operations.")
	time.Sleep(5 * time.Second)

	doneWatchingLogs := make(chan struct{})

	if !disable_events {
		go watchContractEvents(contractAddr, doneWatchingLogs)
	} else {
		log.Infof("Event watchers are disabled.")
	}

	// use the contract to run the increment workload
	go generateNonceAtRate(client, fromAddress, count, duration, nonceStream, maxrate)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go mixed(privateKey, client, instance, gasPrice, nonceStream, &wg, i)
	}
	wg.Wait()

	// time's up or #operations has finished
	log.Info("Experiment FINISHED, wait for the 10 seconds to ensure pending transactions are processed")
	time.Sleep(10 * time.Second)
	lastnonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Impossible to get a nonce for final report", err)
	}
	auth.Nonce = big.NewInt(int64(lastnonce))
	instance.PrintAllData(auth)
	log.Info("Wait 10 seconds to receive the final log")
	time.Sleep(10 * time.Second)
	close(doneWatchingLogs)
	//wait for logs to be closed
	time.Sleep(10 * time.Second)
}
