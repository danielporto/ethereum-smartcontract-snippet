package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("http://localhost:7545")
	if err != nil {
		log.Fatalf("Failed to connect to ethereum network: %v ", err)
	}

	// #1 get an account pkey (ganache)
	privateKey, err := crypto.HexToECDSA("ed94c4cc8cf64ca5c85c838a64bd0ae018299d370bacc2fc204ceee78ff75abe")
	if err != nil {
		log.Fatal(err)
	}

	// #2 extract pubkey from pkey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("Error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(5000000000000000000) // value (5 eth) in weight: 5 + 18 zeros.
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x01345b5946fc518644355EF6dCB4a61BF6EF7955")
	var data []byte //all transactions require data even when is empty
	// now creates a transaction
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	//sign it

	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//send it
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Sent %s wei to %s: %s\n", value.String(), toAddress.Hex(), signedTx.Hash().Hex())
}
