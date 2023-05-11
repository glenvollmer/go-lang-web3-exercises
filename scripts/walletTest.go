package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func walletTest() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// private key generation
	pData := crypto.FromECDSA(pvk)
	privateKey := hexutil.Encode(pData)
	fmt.Println(string(privateKey))

	// public key generation
	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	publicKey := hexutil.Encode(puData)
	fmt.Println(string(publicKey))

	// address generation
	address := crypto.PubkeyToAddress(pvk.PublicKey).Hex()
	fmt.Println(string(address))
}
