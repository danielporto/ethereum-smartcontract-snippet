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
	"fmt"
	"math/big"
	"math/rand"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/danielporto/ethereum-smartcontract-snippet/ycsb/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var operation string
var count int

var dbkeysize int
var dbvaluesize int

// workloadCmd represents the workload command
var workloadCmd = &cobra.Command{
	Use:   "workload",
	Short: "Deploy and execute contract operations",
	Long: `Simulate a client by first deploying a contract on the network node
Then, call several operations of that contract. `,
	Run: func(cmd *cobra.Command, args []string) {

		switch operation {
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
	rand.Seed(time.Now().UnixNano())
	rootCmd.AddCommand(workloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	workloadCmd.PersistentFlags().StringVarP(&operation, "operation", "o", "deploy", "Issues operations of that type (deploy/mixed) to the blockchain")
	workloadCmd.PersistentFlags().IntVarP(&dbkeysize, "dbkeysize", "e", 1, "size of the entry string")
	workloadCmd.PersistentFlags().IntVarP(&dbvaluesize, "dbvaluesize", "v", 1, "size of the value string")
	workloadCmd.PersistentFlags().IntVarP(&duration, "duration", "d", 0, "Duration of the experiment in seconds")
	workloadCmd.PersistentFlags().IntVarP(&threads, "threads", "t", 1, "Number of threads")

}

// generate a random string of sizeX
func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//
// generate a sequence of nonces up until a duration or a defined number
//
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

func deploy(pk *ecdsa.PrivateKey, c *ethclient.Client, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup) {
	defer wg.Done()

	total_transactions := 0
	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		auth.GasPrice = gasPrice
		address, tx, _, err := contracts.DeployKVstore(auth, c)
		if err != nil {
			log.Fatal("Error deploying KVstore contract", err)
		}
		if total_transactions%10000 == 0 {
			log.Infof("Transactions sent by this thread: %v", total_transactions)
		}

		total_transactions++

		log.Debugf("Transaction address: %v\n", address.Hex())
		log.Debugf("Contract address %v", tx.Hash().Hex())
		log.Debugln("Contract deployed")

	}
	log.Infof("Total transactions sent by this thread: %v", total_transactions)

}

func workloadDeploy() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	log.Println("Running deploy workload")
	log.Println("Connecting to ethereum network...")
	conn, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal("Failed to connect to ethereum node", err)
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatal("Error converting the private key from Hex to ECDSA", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Error while getting a new nonce", err)
	}
	fmt.Printf("Nonce: %v\n", nonce)

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error while trying to get the gas price", err)
	}
	fmt.Println("Gas price:", gasPrice)

	go generateNonce(nonce, count, duration, nonceStream)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go deploy(privateKey, conn, gasPrice, nonceStream, &wg)
	}
	wg.Wait()

}

func mixed(pk *ecdsa.PrivateKey, c *ethclient.Client, instance *contracts.KVstore, gasPrice *big.Int, nonces <-chan uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	total_transactions := 0

	for nonce := range nonces {
		auth := bind.NewKeyedTransactor(pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)
		auth.GasLimit = uint64(300000)
		auth.GasPrice = gasPrice
		if nonce%2 == 0 {
			address, tx, _, err := contracts.DeployKVstore(auth, c)
			if err != nil {
				log.Fatal("Error deploying simple storage", err)
			}
			log.Debugf("Transaction address: %v\n", address.Hex())
			log.Debugf("Contract address %v", tx.Hash().Hex())
			log.Debugln("Contract deployed")
		} else {
			// TODO generate random string
			tx, err := instance.Set(auth, randStringBytes(dbkeysize), randStringBytes(dbvaluesize))
			if err != nil {
				log.Fatal("Failed to call transaction method: ", err)
			}
			log.Debugf("nonce %v, tx sent: %s\n", nonce, tx.Hash().Hex())
		}
		total_transactions++
		if total_transactions%10000 == 0 {
			log.Infof("Transactions sent by this thread: %v", total_transactions)
		}
	}
	log.Infof("Total transactions sent by this thread: %v", total_transactions)

}

func workloadMixed() {
	nonceStream := make(chan uint64)
	wg := sync.WaitGroup{}

	// 1. Initialize a connection
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
		log.Fatal("Error converting the private key", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: public key is not the type *ecdsa.PublicKey")
	}
	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//3. configure nonce (prevent replay attacks with a user specific nonce)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Impossible to get a nonce for acccount", err)
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
	auth.Value = big.NewInt(0) // not transfering funds, just executing contract operation
	auth.GasLimit = uint64(0)  // max value for this transaction
	auth.GasPrice = gasPrice
	_, _, instance, err := contracts.DeployKVstore(auth, client)
	if err != nil {
		log.Fatal("Impossible to initialize a counter for this workload.", err)
	}
	fmt.Println("wait for the 5 blocks for the contract to be mined")
	time.Sleep(5 * time.Second)

	// use the contract to run the increment workload
	go generateNonce(nonce, count, duration, nonceStream)
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go mixed(privateKey, client, instance, gasPrice, nonceStream, &wg)
	}
	wg.Wait()

	fmt.Println("Finishing experiment, wait for the 5 blocks to be mined (ensure pending transactions are mined)")
	time.Sleep(5 * time.Second)
}
