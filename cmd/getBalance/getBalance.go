package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sonm-io/simple-gate/cmd"
	"os"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
)

func main() {
	addr := common.HexToAddress("0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD")

	clientPrivate, err := ethclient.Dial(cmd.EndpointPrivate)
	if err != nil {
		return
	}

	token, err := cmd.InitSidechainToken(clientPrivate)
	if err != nil {
		cmd.HandleError(err, "bind side token")
	}

	balance, err := token.BalanceOf(&bind.CallOpts{Pending:true}, addr)
	if err != nil {
		cmd.HandleError(err, "getting balance")
	}

	log.Printf("Balance: %v", balance)
}

func processArgs() common.Address {
	hashRaw := os.Args[1]
	return common.HexToAddress(hashRaw)
}
