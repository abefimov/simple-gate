package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sonm-io/simple-gate"
	"github.com/sonm-io/simple-gate/cmd"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
	var valueRaw = os.Args[2]
	var keyRaw = os.Args[1]

	val, err := strconv.Atoi(valueRaw)
	if err != nil {
		cmd.HandleError(err, "convert value")
	}
	value := big.NewInt(int64(val))

	client, err := ethclient.Dial(cmd.Endpoint)
	if err != nil {
		cmd.HandleError(err, "init live client")
	}

	key, err := crypto.HexToECDSA(keyRaw)
	if err != nil {
		cmd.HandleError(err, "decrypt key")
	}

	opts := cmd.GetLiveTxOpts(key, 100000)

	token, err := cmd.InitLiveToken(client)
	if err != nil {
		cmd.HandleError(err, "bind live token")
	}

	address := crypto.PubkeyToAddress(key.PublicKey)

	balance, err := token.BalanceOf(&bind.CallOpts{Pending: true}, address)
	if err != nil {
		cmd.HandleError(err, "getting balance")
	}

	allowance, err := token.Allowance(&bind.CallOpts{Pending: true}, address, simple_gate.GatekeeperLiveAddr())
	if err != nil {
		cmd.HandleError(err, "getting allowance")
	}

	log.Println("Your balance: :", balance)
	log.Println("Your allowance: :", allowance)

	if allowance.Cmp(value) < 0 {
		log.Fatalln("lack of allowance")
	}
	if balance.Cmp(value) < 0 {
		log.Fatalln("lack of balance")
	}

	gk, err := cmd.InitLiveGatekeeper(client)
	if err != nil {
		log.Fatalln(err)
		return
	}

	tx, err := gk.PayIn(opts, value)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(tx.Hash().String())
}
