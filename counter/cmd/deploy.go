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
	// "time"

	"github.com/danielporto/ethereum-smartcontract-snippet/counter/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the smartcontract into the blockchain",
	Long: `Install and initialize a smartcontract that maintains a counter in a
Quorum blochckain network.
Example:
./counter deploy --url "http://146.193.41.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		deployCounter(key, host, port)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func checkReceipt(ethC *ethclient.Client, tx *types.Transaction, attempts int) bool {
// 	for i := 0; i < attempts; i++ {
// 		time.Sleep(200 * time.Millisecond)
// 		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
// 		r, err := ethC.TransactionReceipt(ctx, tx.Hash())
// 		cancel()
// 		if r == nil || err != nil {
// 			continue
// 		}
// 		if r.Status == 1 {
// 			return true
// 		}
// 		return false
// 	}
// 	return false
// }

func deployCounter(key, host, port string) (string, uint64, *big.Int) {
	log.Println("Connecting to ethereum network...")
	url := "http://" + host + ":" +port
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address, tx, instance, err := contracts.DeployCounter(auth, conn)
	if err != nil {
		log.Fatal("Error deploying simple storage", err)
	}

	// check receipt
	// if checkReceipt(conn, tx, 3) == false {
	// 	log.Fatal("Error: impossible to verify the transaction: ", tx)
	// }

	fmt.Println("Transaction address:", tx.Hash().Hex())
	fmt.Println("Contract address", address.Hex())
	fmt.Println("Contract deployed")

	_ = instance
	return address.Hex(), nonce, gasPrice

}
