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

	log "github.com/sirupsen/logrus"
	"math/big"

	"github.com/danielporto/ethereum-smartcontract-snippet/counter/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"
)



// incrementCmd represents the increment command
var incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "Increments a counter in a smartcontract",
	Long: `Connects to a smartcontract in a quorum network and increments a counter`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("increment called")
		Increment(key, host, port)

	},
}

func init() {
	rootCmd.AddCommand(incrementCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incrementCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incrementCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	incrementCmd.PersistentFlags().IntVarP(&amount, "amount", "a", 1, "Amount to add to a counter." )
	incrementCmd.PersistentFlags().StringVarP(&contract, "contract", "c", "", "Contract address on the blockchain." )
	incrementCmd.MarkFlagRequired("contract")
}

func Increment(key, host, port string) *types.Transaction {

	// 1. Initialize a connection
	url := "http://" + host +":"+port
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

	// 4. setup an authenticated transactor with info from credentials and connection configuration
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // not transfering funds, just executing contract operation
	auth.GasLimit = uint64(0)  // max value for this transaction
	auth.GasPrice = gasPrice

	// 5. load a smartcontract
	// use the transaction id of the contract deployed to return a reference to the contract
	address := common.HexToAddress(contract)
	instance, err := contracts.NewCounter(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.Increment(auth, big.NewInt(int64(amount)))
	if err != nil {
		log.Fatal("Failed to call transaction method", err)
	}
	// check receipt
	// if !checkReceipt(client, tx, 3) {
	// 	log.Fatal("Error: impossible to verify the transaction", tx)
	// }
	fmt.Printf("nonce %v, tx sent: %s\n", nonce, tx.Hash().Hex())
	return tx

}
