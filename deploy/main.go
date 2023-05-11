// This script demonstrates simple web3 methods used to...
// ~ load a wallet and account from a keystore json file using passphrase
// ~ create a transaction with suggested gas prices to deploy the smart contract to the blockchain using smart contract binding

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	todo "golang_web3/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	godotenv.Load()
	NODE_RPC_URI := os.Getenv("NODE_RPC_URI")
	WALLET_JSON := os.Getenv("WALLET_JSON")
	WALLET_PASSWORD := os.Getenv("WALLET_PASSWORD")

	// read wallet keystore from json file
	b, err := ioutil.ReadFile(WALLET_JSON)
	if err != nil {
		log.Fatal(err)
	}

	// decrypt wallet keystore with wallet password
	key, err := keystore.DecryptKey(b, WALLET_PASSWORD)
	if err != nil {
		log.Fatal(err)
	}

	// connect to blockain node rpc
	client, err := ethclient.Dial(NODE_RPC_URI)
	if err != nil {
		log.Fatal(err)
	}

	// user crypto module to get wallet address from wallet keystore public key
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	// get transaction nonce from wallet address
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal(err)
	}

	// get suggested gas price
	gas, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// get blockchain network id
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// create new transaction with wallet
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// set gas, gas limit, and nonce for transaction
	auth.GasPrice = gas
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))

	// deploy smart contract using contract binding
	a, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	// print contract address to the console
	fmt.Println(a.Hex())

	// print transaction hash to the console
	fmt.Println(tx.Hash().Hex())
}
