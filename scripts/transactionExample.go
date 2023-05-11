// This script demonstrates simple web3 methods used to...
// ~ load two wallets keystore json files
// ~ create a transaction to send eth from wallet one to wallet two

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// helper function for for getting a wallets private from the keystore JSON file
func getWalletPrivateKey(walletKeystorePath string, walletPassword string) *keystore.Key {
	// read Json file for wallet keystore
	b, err := ioutil.ReadFile(walletKeystorePath)
	if err != nil {
		log.Fatal(err)
	}

	// decrypt walelt keystore with wallet passphrase
	key, err := keystore.DecryptKey(b, walletPassword)
	if err != nil {
		log.Fatal(err)
	}

	return key
}

// helped function for getting the public address from a wallets keystore json
func getWalletAddressFromKey(walletKeystorePath string, walletPassword string) common.Address {
	// call the helper function to get the wallets keystore
	key := getWalletPrivateKey(walletKeystorePath, walletPassword)

	// get the public address from the decoded wallet keysore
	pubAdd := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	return pubAdd
}

func main() {
	// load environment variables
	godotenv.Load()
	NODE_RPC_URI := os.Getenv("NODE_RPC_URI")
	WALLET_JSON_ONE := os.Getenv("WALLET_JSON_ONE")
	WALLET_ONE_PASSWORD := os.Getenv("WALLET_ONE_PASSWORD")
	WALLET_JSON_TWO := os.Getenv("WALLET_JSON_TWO")
	WALLET_TWO_PASSWORD := os.Getenv("WALLET_TWO_PASSWORD")

	// connect to the blockchains RPC client
	client, err := ethclient.Dial(NODE_RPC_URI)
	if err != nil {
		log.Fatal(err)
	}

	// get the two wallets public addresses
	addressOne := getWalletAddressFromKey(WALLET_JSON_ONE, WALLET_ONE_PASSWORD)
	addressTwo := getWalletAddressFromKey(WALLET_JSON_TWO, WALLET_TWO_PASSWORD)

	// get wallet ones transaction nonce
	a1Nonce, err := client.PendingNonceAt(context.Background(), addressOne)
	if err != nil {
		log.Fatal(err)
	}

	// get the current balances of both wallets using the public addresses and log them to the console
	b1, err := client.BalanceAt(context.Background(), addressOne, nil)
	b2, err := client.BalanceAt(context.Background(), addressTwo, nil)

	fmt.Println(b1, b2)

	// set the amount and gas for transaction between wallet one and two
	amount := big.NewInt(10000000000000000)
	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// create a new transaction for wallet one to send 1 eth to wallet two
	ptx := types.NewTransaction(a1Nonce, addressTwo, amount, 21000, gas, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// get the private key of wallet one in order to sign the transaction
	key := getWalletPrivateKey(WALLET_JSON_ONE, WALLET_ONE_PASSWORD)

	// sign the transaction between wallet one and two
	tx, err := types.SignTx(ptx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// submit the transaction between wallet one and two
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	// log the transaction hash to the console
	fmt.Println(tx.Hash().Hex())
}
