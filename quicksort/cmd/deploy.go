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
	"strconv"

	"github.com/danielporto/ethereum-smartcontract-snippet/quicksort/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy the smartcontract into the blockchain",
	Long: `Install and initialize a smartcontract that sorts an array of ints in a
Quorum blochckain network.
Example:
./quicksort deploy --host "192.169.10.166" --port 23000 --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy called")
		p, err := strconv.Atoi(port)
		if err != nil {
			LogFatal("Error converting the socket port:", port, err)
		}
		url := "ws://" + host + ":" + strconv.Itoa(p)
		deployQuickSort(key, url)
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

func deployQuickSort(key, url string) (string, uint64, *big.Int) {
	Log("Connecting to ethereum network...")
	conn, err := ethclient.Dial(url)
	if err != nil {
		LogFatal("Failed to connect to ethereum node", err)
	}

	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		LogFatal("Error converting the private key from Hex to ECDSA", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		LogFatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		LogFatal("Error while getting a new nonce", err)
	}
	fmt.Printf("Nonce: %v\n", nonce)

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		LogFatal("Error while trying to get the gas price", err)
	}
	fmt.Println("Gas price in wei (10^-18ether):", gasPrice)

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(trxgaslimit)
	auth.GasPrice = gasPrice

	address, tx, instance, err := contracts.DeployQuickSort(auth, conn)
	if err != nil {
		LogFatal("Error deploying quicksort contract. Check the gaslimit for this transaction (hardcoded):", auth.GasLimit, " err:", err)
	}

	// check receipt
	// if checkReceipt(conn, tx, 3) == false {
	// 	LogFatal("Error: impossible to verify the transaction: ", tx)
	// }

	fmt.Println("Transaction address:", tx.Hash().Hex())
	fmt.Println("Contract address", address.Hex())
	fmt.Println("Contract deployed")

	_ = instance
	return address.Hex(), nonce, gasPrice

}
