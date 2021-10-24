/*
Copyright © 2021 DANIEL PORTO <daniel.porto@gmail.com>

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
	"github.com/danielporto/ethereum-smartcontract-snippet/bank/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/spf13/cobra"
	"math/big"
	"math/rand"
	"time"
)

// depositCmd represents the transfer command
var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "transfer funds to a smartcontract",
	Long: `transfer/deposit funds (in wei) to a smartcontract.
bank deposit --host "192.168.10.166" --port 23000 \
    --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0" \
	--contract 0x39345B680c47B03fa5ddd33EDb368bCC550921A3 --money 1000000000000
`,
	Run: func(cmd *cobra.Command, args []string) {
		Log("deposit called")
		Log("Contract: %v", contract)
		Log("money to transfer: %v", money)
		Log("from: %v", key)
		Deposit(key, host, port)
	},
}

//func checkReceipt(ethC *ethclient.Client, tx *types.Transaction, attempts int) bool {
//	for i := 0; i < attempts; i++ {
//		time.Sleep(200 * time.Millisecond)
//		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//		r, err := ethC.TransactionReceipt(ctx, tx.Hash())
//		cancel()
//		if r == nil || err != nil {
//			continue
//		}
//		if r.Status == 1 {
//			return true
//		}
//		return false
//	}
//	return false
//}

func init() {
	rand.Seed(time.Now().UnixNano())
	rootCmd.AddCommand(depositCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// depositCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// depositCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	depositCmd.PersistentFlags().StringVarP(&key, "key", "", "", "key of an account wallet")
	depositCmd.MarkPersistentFlagRequired("key")
	depositCmd.PersistentFlags().Int64VarP(&money, "money", "", 1000000000000, "Amount of money to transfer between accounts")
	depositCmd.PersistentFlags().StringVarP(&client_id, "id", "", "undefined", "client identifier")
	depositCmd.PersistentFlags().StringVarP(&contract, "contract", "", "", "Contract address on the blockchain.")
	depositCmd.MarkFlagRequired("contract")

}

func Deposit(src, host, port string) *types.Transaction {

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	client, err := ethclient.Dial(url)
	if err != nil {
		LogFatal("Failed to connect to ethereum node", err)
	}

	if len(contract) == 0 {
		LogFatal("Contract was not defined")
	}
	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(src)
	if err != nil {
		LogFatal("Error converting the private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		LogFatal("Cannot assert type: public key is not the type *ecdsa.PublicKey")
	}
	//now get the account of that private key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//3. configure nonce (prevent replay attacks with a user specific nonce)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		LogFatal("Impossible to get a nonce for account", err)
	}

	//4. configure gasPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		LogFatal("Unable to get a gas price", err)
	}

	// 4. setup an authenticated transactor with info from credentials and connection configuration
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(money)) // transfering funds
	auth.GasLimit = uint64(trxgaslimit)             // max value for this transaction
	auth.GasPrice = gasPrice

	Log("Gas price: %v", gasPrice)
	Log("Gas limit: %v", auth.GasLimit)

	// 5. load a smartcontract
	// use the transaction id of the contract deployed to return a reference to the contract
	address := common.HexToAddress(contract)
	instance, err := contracts.NewBankTransactor(address, client)
	if err != nil {
		LogFatal("Impossible to initialize the contract for this workload. %v", err)
	}

	id := fmt.Sprintf("%v_tx_%v", client_id, nonce)
	tx, err := instance.Deposit(auth, id)
	if err != nil {
		LogFatal("Failed to call transaction method: %v", err)
	}
	// check receipt
	// if !checkReceipt(client, tx, 3) {
	// 	LogFatal("Error: impossible to verify the transaction", tx)
	// }

	LogDebug("nonce %v, tx sent: %s", nonce, tx.Hash().Hex())
	return tx
}
