package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// var infuraURI = "https://goerli.infura.io/v3/c15571d3be9c4e6d848ed4dc78e92cab"
var addOne = "6e39ffbd75e3f8bda19c283bc32917578b2abcc0"
var addTwo = "0aa0c3b0ae20292827594333052471197fc07f27"

func testnetTest() {
	client, err := ethclient.Dial(infuraURI)
	if err != nil {
		log.Fatal(err)
	}

	addressOne := common.HexToAddress(addOne)
	addressTwo := common.HexToAddress(addTwo)

	a1Nonce, err := client.PendingNonceAt(context.Background(), addressOne)
	if err != nil {
		log.Fatal(err)
	}
	// b1, err := client.BalanceAt(context.Background(), addressOne, nil)
	// b2, err := client.BalanceAt(context.Background(), addressTwo, nil)

	// fmt.Println(b1, b2)
	amount := big.NewInt(10000000000000000)
	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ptx := types.NewTransaction(a1Nonce, addressTwo, amount, 21000, gas, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./wallet/UTC--2023-02-24T18-43-53.566632000Z--6e39ffbd75e3f8bda19c283bc32917578b2abcc0.json")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	tx, err := types.SignTx(ptx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex())
}
