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
	//txHash := processArgs()
	txHash := common.HexToHash("0xcbb5174a9a09c63b675fdb6d24d6692a99b43d75364a72e6f7723f00b5985ce3")

	clientPrivate, err := ethclient.Dial(cmd.EndpointPrivate)
	if err != nil {
		return
	}

	t, _, err := clientPrivate.TransactionByHash(context.Background(), txHash)
	if err != nil {
		cmd.HandleError(err, "get tx receipt")
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
	fmt.Printf("Gas: %d \r\n", t.Gas())
	fmt.Printf("GasPrice: %d \r\n", t.GasPrice())
	fmt.Printf("Cost: %d \r\n", t.Cost())
	fmt.Printf("Data: %d \r\n", t.Data())
	fmt.Printf("To: %d \r\n", t.To().String())

	if tx.ContractAddress.String() != "" {
		fmt.Printf("Contract created: %v \r\n", tx.ContractAddress.String())
	}
}

func processArgs() common.Hash {
	hashRaw := os.Args[1]
	return common.HexToHash(hashRaw)
}
