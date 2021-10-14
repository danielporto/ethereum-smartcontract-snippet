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
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/danielporto/ethereum-smartcontract-snippet/bank/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
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
Then, call several operations of that contract. 
If no origin or target is defined, a random is selected amongst the list of keys.
rate is in ops/second.
checkpoint is in number of operations.
money is in wei.

bank workload -o [transfer|deploy|mixed] \
    [--count 1 | --duration 0 ] \
	[--threads 1]	\
	[--rate 10] \
	[--checkpoint 5000] \
	--host "192.168.10.166" --port 23000 \
    --keys "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0,9bdd6a2e7cc1ca4a4019029df3834d2633ea6e14034d6dcc3b944396fe13a08b" \
    --wallets "0xed9d02e382b34818e88b88a309c7fe71e65f419d,0xca843569e3427144cead5e4d5999a3d0ccf92b8e" \ 
	[--money 1000000000000] \
    [--origin 0]  [--target 1]
`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Infof("Workload operation '%v'", operation)
		log.Infof("Gas price limit for the transaction: %v wei", trxgaslimit)
		log.Infof("Max rate to issue operations (suggested tps) %v tps", maxrate)
		log.Infoln("workload called")

		keys = strings.Split(keys_str,",")
		wallets = strings.Split(wallets_str,",")

		// if origin == -1 select a key randomly
		// if target == -1 select a target randomly, target cannot be origin
		if origin == -1 {
			origin = rand.Intn(len(keys))
		}
		if target == -1 {
			target = origin
			for ; target == origin ;target = rand.Intn(len(wallets)) {
				// just loop for a random num until is different from self
			}
		}
		log.Debugf("Origin: %v key: %v\n", origin, keys[origin])
		log.Debugf("Target:%v wallet:%v\n", target, wallets[target])

		switch operation {
			case "transfer":
				workloadTransfer()
			case "deploy":
				workloadDeploy()
			case "mixed":
				workloadMixed()
			default:
				log.Fatal("Operation not supported:", operation)
		}
	},
}

func init() {
	rootCmd.AddCommand(workloadCmd)

	// Here you will define your flags and configuration settings.
	workloadCmd.PersistentFlags().StringVarP(&operation, "operation", "", "transfer", "Issues operations of that type (deploy/transfer/mixed) to the blockchain")
	workloadCmd.PersistentFlags().IntVarP(&count, "count", "", 1, "Number of operations to be issued")
	workloadCmd.PersistentFlags().IntVarP(&duration, "duration", "", 0, "Duration of the experiment in seconds")
	workloadCmd.PersistentFlags().IntVarP(&threads, "threads", "", 1, "Number of threads")
	workloadCmd.PersistentFlags().IntVarP(&maxrate, "rate", "", 10, "Max operations per second for each thread (suggested value)")
	workloadCmd.PersistentFlags().IntVarP(&checkpoint, "checkpoint", "", 5000, "Print total operations after X operations issued.")
	workloadCmd.PersistentFlags().StringVarP(&keys_str, "keys", "", "", "list of keys of a account wallets")
	workloadCmd.MarkPersistentFlagRequired("keys")
	workloadCmd.PersistentFlags().StringVarP(&wallets_str, "wallets", "", "", "list of addresses of account wallets")
	workloadCmd.MarkPersistentFlagRequired("wallets")
	workloadCmd.PersistentFlags().IntVarP(&origin, "origin", "", -1, "index of source keys list (-1 selects a random origin key)")
	workloadCmd.PersistentFlags().IntVarP(&target, "target", "", -1, "index of destination address (-1 will choose a random target wallet)")
	workloadCmd.PersistentFlags().IntVarP(&money, "money", "", 1000000000000, "Amount of money to transfer between accounts")
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
	go watchOperationsExecutedEvents(client, contractAddr, stop)
	go watchBalanceTransferredEvents(client, contractAddr, stop)
	go watchBalanceReceivedEvents(client, contractAddr, stop)

	log.Debug("Logging thread: Waiting for logs to be closed")
	// wait for the signal to stop
	<-done
	log.Info("Logging thread: Teardown subscription to contract events.")

	//stop all 3 logging threads
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

		address, _, _, err := contracts.DeployBank(auth, c)
		if err != nil {
			log.Fatal("Error deploying simple storage", err)
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
	log.Info("Running deploy workload")
	log.Infof("Connecting to ethereum network: %v", url)

	conn, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to ethereum node", err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(keys[origin])
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
		log.Fatal("Error while trying to get the gas price", err)
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
* The transfer workload install a single contract and issue operations to transfer a counter
* It also emits a log with a value of the counter.
* The logs allow for detecting byzantine faults that could be possible not captured.
***********************************************************************************************************************/
func transfer(pk *ecdsa.PrivateKey, instance *contracts.Bank, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()
	total_transactions := 0
	log.Infof("Thread %v STARTED - issuing transfer transactions", threadid)

	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(int64(money))
		auth.GasLimit = uint64(trxgaslimit)
		auth.GasPrice = gasPrice

		if nonce%uint64(checkpoint) == 0 {
			log.Infof("Thread %v - counter transactions issued : %v", threadid, total_transactions)
			instance.LogTransferOperations(auth)
			continue // get another nonce
		} 
		
		tx, err := instance.TransferMoneyTo(auth, common.HexToAddress(wallets[target]))
		if err != nil {
			log.Fatal("Failed to call transfer transaction method of counter contract. Check the gaslimit for this transaction:", auth.GasLimit, " err:", err)
		}

		total_transactions++
		log.Debugf("nonce %v, tx sent: %s", nonce, tx.Hash().Hex())

	}
	log.Infof("Thread %v FINISHED - transfer transactions issued : %v", threadid, total_transactions)

}

func workloadTransfer() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	log.Info("Running transfer workload")
	log.Infof("Connecting to ethereum network: %v", url)

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to ethereum node", err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(keys[origin])
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
	auth.Value = big.NewInt(0) // not transferring funds, just deploying contract operation
	auth.GasLimit = uint64(trxgaslimit)   // max value for this transaction
	auth.GasPrice = gasPrice
	contractAddr, _, instance, err := contracts.DeployBank(auth, client)
	if err != nil {
		log.Fatal("Impossible to initialize a counter for this workload. ", err)
	}
	log.Info("Wait for the 5 seconds (blocks) while contract is be mined before issuing operations.")
	time.Sleep(5 * time.Second)

	doneWatchingLogs := make(chan struct{})

	if !disable_events {
		go watchContractEvents(contractAddr, doneWatchingLogs)
	} else {
		log.Infof("Event watchers are disabled.")
	}

	// use the contract to run the transfer workload
	go generateNonceAtRate(client, fromAddress, count, duration, nonceStream, maxrate)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go transfer(privateKey, instance, gasPrice, nonceStream, &wg, i)
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
	instance.LogTransferOperations(auth)
	log.Info("Wait 10 seconds to receive the final log")
	time.Sleep(10 * time.Second)
	close(doneWatchingLogs)
	//wait for logs to be closed
	time.Sleep(10 * time.Second)
}

/**********************************************************************************************************************
* The mixed workload alternates with operations that install smartcontracts and also execute counter transfer ops
***********************************************************************************************************************/
func mixed(pk *ecdsa.PrivateKey, c *ethclient.Client, instance *contracts.Bank, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup, threadid int) {
	defer wg.Done()
	total_transactions := 0
	log.Infof("Thread %v STARTED - issuing mix deploy/transfer transactions", threadid)

	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(trxgaslimit)
		auth.GasPrice = gasPrice

		if nonce%uint64(checkpoint) == 0 {
			log.Infof("Thread %v - deploy/transfer transactions issued : %v", threadid, total_transactions)
			instance.LogTransferOperations(auth)
			total_transactions++
			continue // get another nonce
		}

		if nonce%2 == 0 {
			address, _, _, err := contracts.DeployBank(auth, c)
			if err != nil {
				log.Fatal("Error deploying counter", err)
			}
			log.Debugf("Transaction address: %v\n", address.Hex())
		} else {
			tx, err := instance.TransferMoneyTo(auth, common.HexToAddress(wallets[target]))
			if err != nil {
				log.Fatal("Failed to call transaction method: ", err)
			}
			log.Debugf("nonce %v, tx sent: %s\n", nonce, tx.Hash().Hex())
		}
		total_transactions++

	}
	log.Infof("Thread %v FINISHED - mix transfer/deploy transactions issued : %v", threadid, total_transactions)

}

func workloadMixed() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	log.Info("Running mix deploy/transfer workload")
	log.Infof("Connecting to ethereum network: %v", url)

	client, err := ethclient.Dial(url)
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
	contractAddr, _, instance, err := contracts.DeployBank(auth, client)
	if err != nil {
		log.Fatal("Impossible to initialize a counter contract for this workload.", err)
	}
	log.Info("Wait for the 5 seconds (blocks) while contract is be mined before issuing operations.")
	time.Sleep(5 * time.Second)

	doneWatchingLogs := make(chan struct{})

	if !disable_events {
		go watchContractEvents(contractAddr, doneWatchingLogs)
	} else {
		log.Infof("Event watchers are disabled.")
	}

	// use the contract to run the transfer workload
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
	instance.LogTransferOperations(auth)
	log.Info("Wait 10 seconds to receive the final log")
	time.Sleep(10 * time.Second)
	close(doneWatchingLogs)
	//wait for logs to be closed
	time.Sleep(10 * time.Second)
}
