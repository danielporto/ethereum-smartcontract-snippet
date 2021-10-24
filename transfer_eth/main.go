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

	client, err := ethclient.Dial("ws://146.193.41.166:23000")
	if err != nil {
		log.Fatalf("Failed to connect to ethereum network: %v ", err)
	}

	// #1 get an account pkey (ganache)
	privateKey, err := crypto.HexToECDSA("1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0")
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

	//value := big.NewInt(5000000000000000000) // value (5 eth) in weight: 5 + 18 zeros.
	value := big.NewInt(1) // value (5 eth) in weight: 5 + 18 zeros.

//	gasLimit := uint64(21000)
	gasLimit := uint64(50000000)

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}
	gasPrice := big.NewInt(0)

	toAddress := common.HexToAddress("0xca843569e3427144cead5e4d5999a3d0ccf92b8e")
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
