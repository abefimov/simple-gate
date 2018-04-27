package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sonm-io/simple-gate"
	"github.com/sonm-io/simple-gate/cmd"
	"log"
	"math/big"
)

var (
	//HexKey string = ""
	hexKey string = ""
)

func main() {
	clientPrivate, err := ethclient.Dial(cmd.EndpointPrivate)
	if err != nil {
		cmd.HandleError(err, "init live client")
	}

	key, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		cmd.HandleError(err, "decrypt key")
	}

	opts := bind.NewKeyedTransactor(key)
	opts.GasLimit = 100000
	opts.GasPrice = big.NewInt(0)

	token, err := cmd.InitSidechainToken(clientPrivate)
	if err != nil {
		log.Fatalln(err)
		return
	}

	tx, err := token.Approve(opts, common.HexToAddress(simple_gate.GatekeeperSidechainAddress), big.NewInt(1000000000000000))
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(tx.Hash().String())
}
