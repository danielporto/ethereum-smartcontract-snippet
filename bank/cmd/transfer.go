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
	"strings"
	"time"
)

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "transfer funds between accounts",
	Long: `Transfer funds (in wei) between accounts.
bank transfer --host "192.168.10.166" --port 23000 \
    --keys "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0,9bdd6a2e7cc1ca4a4019029df3834d2633ea6e14034d6dcc3b944396fe13a08b" \
    --wallets "0xed9d02e382b34818e88b88a309c7fe71e65f419d,0xca843569e3427144cead5e4d5999a3d0ccf92b8e" \ 
	--money 1000000000000
    --origin 0  --target 1
`,
	Run: func(cmd *cobra.Command, args []string) {
		Log("transfer called")

		keys = strings.Split(keys_str, ",")
		wallets = strings.Split(wallets_str, ",")

		// if origin == -1 select a key randomly
		// if target == -1 select a target randomly, target cannot be origin

		if origin == -1 {
			origin = rand.Intn(len(keys))
		}

		if target == -1 {
			target = origin
			for ; target == origin; target = rand.Intn(len(wallets)) {
				// just loop for a random num until is different from self
			}
		}

		LogDebug("Origin: %v key: %v\n", origin, keys[origin])
		LogDebug("Target:%v wallet:%v\n", target, wallets[target])

		Transfer(keys[origin], wallets[target], host, port)
	},
}

func checkReceipt(ethC *ethclient.Client, tx *types.Transaction, attempts int) bool {
	for i := 0; i < attempts; i++ {
		time.Sleep(200 * time.Millisecond)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		r, err := ethC.TransactionReceipt(ctx, tx.Hash())
		cancel()
		if r == nil || err != nil {
			continue
		}
		if r.Status == 1 {
			return true
		}
		return false
	}
	return false
}

func init() {
	rand.Seed(time.Now().UnixNano())
	rootCmd.AddCommand(transferCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	transferCmd.PersistentFlags().StringVarP(&keys_str, "keys", "", "", "list of keys of a account wallets")
	transferCmd.MarkPersistentFlagRequired("keys")
	transferCmd.PersistentFlags().StringVarP(&wallets_str, "wallets", "", "", "list of addresses of account wallets")
	transferCmd.MarkPersistentFlagRequired("wallets")
	transferCmd.PersistentFlags().IntVarP(&origin, "origin", "", -1, "index of source keys list (-1 selects a random origin key)")
	transferCmd.PersistentFlags().IntVarP(&target, "target", "", -1, "index of destination address (-1 will choose a random target wallet)")
	transferCmd.PersistentFlags().IntVarP(&money, "money", "", 1000000000000, "Amount of money to transfer between accounts")
	transferCmd.PersistentFlags().StringVarP(&contract, "contract", "", "", "Contract address on the blockchain.")
	transferCmd.PersistentFlags().StringVarP(&client_id, "id", "", "undefined", "client identifier")
	transferCmd.MarkFlagRequired("contract")

}

func Transfer(src, dst, host, port string) *types.Transaction {

	// 1. Initialize a connection
	url := "ws://" + host + ":" + port
	client, err := ethclient.Dial(url)
	if err != nil {
		LogFatal("Failed to connect to ethereum node", err)
	}

	// 2. Load credentials
	// get credentials to write
	// from ganache the private key is in the key icon of any account/wallet
	// ECDSA (elyptic curve DSA is the standard used by ethereum)
	privateKey, err := crypto.HexToECDSA(src)
	if err != nil {
		LogFatal("Error converting the private key", err)
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
		LogFatal("Impossible to get a nonce for acccount", err)
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
	auth.GasLimit = uint64(0)             // max value for this transaction
	auth.GasPrice = gasPrice

	// 5. load a smartcontract
	// use the transaction id of the contract deployed to return a reference to the contract
	address := common.HexToAddress(contract)
	instance, err := contracts.NewBankTransactor(address, client)
	if err != nil {
		LogFatal("Impossible to initialize the contract for this workload.", err)
	}

	id := fmt.Sprintf("%v_tx_%v", client_id, nonce)
	tx, err := instance.TransferMoneyTo(auth, common.HexToAddress(dst),id)
	if err != nil {
		LogFatal("Failed to call transaction method", err)
	}
	// check receipt
	// if !checkReceipt(client, tx, 3) {
	// 	LogFatal("Error: impossible to verify the transaction", tx)
	// }

	fmt.Printf("nonce %v, tx sent: %s\n", nonce, tx.Hash().Hex())
	return tx
}
