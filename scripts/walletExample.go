// This script demonstrates simple web3 methods used to...
// ~ generate private, public key, and public address for a wallet

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// create private key used for private and public key generations
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// dump private key to binary in order to get the hexadecimal representation of the private key
	pData := crypto.FromECDSA(pvk)
	privateKey := hexutil.Encode(pData)
	fmt.Println(string(privateKey))

	// generate hexadecimal representation of the public key
	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	publicKey := hexutil.Encode(puData)
	fmt.Println(string(publicKey))

	// get the hexadecimal public address from the public key
	address := crypto.PubkeyToAddress(pvk.PublicKey).Hex()
	fmt.Println(string(address))
}
