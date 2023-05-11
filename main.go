// This script demonstrates simple web3 methods used to...
// ~ load a wallet and account from a keystore json file using passphrase
// ~ load a TODO list smart contract
// ~ create a transaction with suggested gas prices to call methods on the smart contract
// ~ submit transaction to call smart contract methods

package main

import (
	"context"
	"fmt"
	todo "golang_web3/gen"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
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
	TODO_CONTRACT_ADDRESS := os.Getenv("TODO_CONTRACT_ADDRESS")

	// read wallet keystore JSON
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
	defer client.Close()

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

	// convert contract hex to address
	contractAddress := common.HexToAddress(TODO_CONTRACT_ADDRESS)

	// load web3 contract bindings with contract address
	t, err := todo.NewTodo(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// instantiate a new transaction for trigger functions on the smart contract
	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// set gas limit and gas price
	tx.GasLimit = 3000000
	tx.GasPrice = gas

	// get cli args for triggering smart contract functions
	var cliArgs []string = os.Args
	var method string = cliArgs[1]

	if method == "add" {
		// ======================================================== add task
		var newTask string = cliArgs[2]

		transaction, err := t.Add(tx, newTask)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(transaction.Hash())

	} else if method == "get" {
		// ======================================================== get task
		address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

		task, err := t.Get(&bind.CallOpts{
			From: address,
		}, big.NewInt(0))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(task)

	} else if method == "list" {
		// ======================================================== list tasks
		address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

		tasks, err := t.List(&bind.CallOpts{
			From: address,
		})

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tasks)

	} else if method == "update" {
		// ======================================================== update tasks
		transaction, err := t.Update(tx, big.NewInt(0), "First Task Update")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(transaction.Hash())

	} else if method == "toggle" {
		// ======================================================== toggle tasks
		transaction, err := t.Toggle(tx, big.NewInt(0))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(transaction.Hash())

	} else if method == "remove" {
		// ======================================================== remove tasks
		transaction, err := t.Remove(tx, big.NewInt(0))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(transaction.Hash())

	}
}
