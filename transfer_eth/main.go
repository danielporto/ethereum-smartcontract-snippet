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
	privateKey, err := crypto.HexToECDSA("bb794ec88d7cb0547b4044a01683e765372f605319dc1d2c25c6f905b0044c3c")
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

	toAddress := common.HexToAddress("0xFE293E3Eb66852EB1c3bfd9F6AD4469d463b601e")
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
