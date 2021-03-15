package main

import (
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-ethereum/ethereum/ethclient"
)

func main() {
	conn, err = ethclient.Dial("http://localhost:9545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("8bc7767bbe12d4aa2af6722d0ec3d24a32687e602fbba70a0575e3af472bc6a9")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()

}
