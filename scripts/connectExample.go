// This script demonstrates simple web3 methods used to...
// ~ connect to a EVM blockchain
// ~ get the blockchains current block number
// ~ check a wallets balance using its public address
// ~ convert a wallets wei balances to eth

package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables
	godotenv.Load()
	NODE_RPC_URI := os.Getenv("NODE_RPC_URI")

	// connect to blockain node rpc
	client, err := ethclient.DialContext(context.Background(), NODE_RPC_URI)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// get current chains most recent block number
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number())

	// check a wallets balance with its address
	addr := "0xBaF6dC2E647aeb6F510f9e318856A1BCd66C5e19"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}

	// convert wallets wei balance to eth
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(value)
	// 1eth = 10^18wei

}
