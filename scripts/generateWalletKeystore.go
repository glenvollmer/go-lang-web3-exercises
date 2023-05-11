package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func keystoreTest() {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	a, err := key.NewAccount(password)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a.Address)
	// b, err := ioutil.ReadFile("./wallet/UTC--2023-02-24T18-43-53.566632000Z--6e39ffbd75e3f8bda19c283bc32917578b2abcc0.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// key, err := keystore.DecryptKey(b, password)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// pData := crypto.FromECDSA(key.PrivateKey)
	// privateKey := hexutil.Encode(pData)
	// fmt.Println(privateKey)

	// pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	// pubKey := hexutil.Encode(pubData)
	// fmt.Println(pubKey)

	// pubAdd := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	// fmt.Println(pubAdd)

}
