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
	"log"
	"math/big"

	"github.com/danielporto/ethereum-smartcontract-snippet/simplestorage/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute smart contract operations",
	Long: `Issues request to the blockchain network with command for 
the simplestorage contract.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
		executeRead()
		// executeWrite()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// Read operations are free of charge and does not require authentication
// thus they are easy to code.
func executeRead() {
	// https://goethereumbook.org/en/smart-contract-load/
	// connects
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	// use the transaction id of the contract deployed to return a reference to the contract
	address := common.HexToAddress("0xF9F093d3bD323968baad4456a8e6BAcB5F0c8A8B")
	instance, err := contracts.NewSimpleStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	// read a value from the blockchain
	counter, err := instance.Get(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current counter value: %v\n", counter)

	// write value to the blockchain
}

// Write operations cost ether. Therefore they require a originator (wallet)
// with a private key that authorize and funds.
func executeWrite() {

	// 1. Initialize a connection
	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA("5be3609f06807dd9c13497acada122f9a660c3ba4abb4d14920fc08b3c39971c")
	if err != nil {
		log.Fatal("Error converting the private key", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: public key is not the type *ecdsa.PublicKey")
	}
	// now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 3. configure nonce (prevent replay attacks with a user specific nonce)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Impossible to get a nonce for acccount", err)
	}
	// 4. configure gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Unable to get a gas price", err)
	}

	// 4. setup an authenticated transactor with info from credentials and connection configuration

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // not transfering funds, just executing contract operation
	auth.GasLimit = uint64(0)  // max value for this transaction
	auth.GasPrice = gasPrice

	// 5. load a smartcontract
	// use the transaction id of the contract deployed to return a reference to the contract
	address := common.HexToAddress("0xF9F093d3bD323968baad4456a8e6BAcB5F0c8A8B")
	instance, err := contracts.NewSimpleStorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.Add(auth, big.NewInt(1))
	if err != nil {
		log.Fatal("Failed to call transaction method", err)
	}
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
