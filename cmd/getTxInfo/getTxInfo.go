package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sonm-io/simple-gate/cmd"
	"os"
)

func main() {
	txHash := processArgs()

	clientPrivate, err := ethclient.Dial(cmd.EndpointPrivate)
	if err != nil {
		return
	}

	tx, err := clientPrivate.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		cmd.HandleError(err, "get tx receipt")
	}

	if tx.Status == 1 {
		fmt.Printf("Transaction successfull \r\n")
	} else {
		fmt.Println("Transaction failed")
	}

	fmt.Printf("Gas used: %d \r\n", tx.GasUsed)
	fmt.Printf("Cumulative Gas used: %d \r\n", tx.CumulativeGasUsed)
	fmt.Printf("Amount of logs: %d \r\n", len(tx.Logs))

	if tx.ContractAddress.String() != "" {
		fmt.Printf("Contract created: %v \r\n", tx.ContractAddress.String())
	}
}

func processArgs() common.Hash {
	hashRaw := os.Args[1]
	return common.HexToHash(hashRaw)
}
