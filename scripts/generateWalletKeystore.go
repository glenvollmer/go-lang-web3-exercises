// This script demonstrates simple web3 methods used to...
// ~ create a new wallet as a keystore json file
// ~ set a password for an account on the wallet
// ~ decrypt the keystore using the wallet account password
// ~ display the private key of the wallet keystore as hexadecimal
// ~ display the public key of the wallet keystore as hexadecimal
// ~ displau the public address for the wallet keystore as hexadecimal

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// wallet keystore password
	password := "password"

	// generate new wallet keystore json file with a password

	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// a, err := key.NewAccount(password)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(a.Address)

	// paste path of newly generated keystore wallet file into ReadFile function
	b, err := ioutil.ReadFile("./wallet/newly_generated_wallet_goes_here")
	if err != nil {
		log.Fatal(err)
	}

	// decrypt wallet keys from wallet keystore data with password used to create wallet keystore
	key, err := keystore.DecryptKey(b, password)

	if err != nil {
		log.Fatal(err)
	}

	// read wallets binary private key as hexadecimal and print to the console
	pData := crypto.FromECDSA(key.PrivateKey)
	privateKey := hexutil.Encode(pData)
	fmt.Println(privateKey)

	// read wallets public key as hexadecimal and print to the console
	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	pubKey := hexutil.Encode(pubData)
	fmt.Println(pubKey)

	// get wallets public address as hexadecimal and print to the console
	pubAdd := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	pubAddHex := pubAdd.Hex()
	fmt.Println(pubAddHex)

}
