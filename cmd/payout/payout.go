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

	clientPrivate, err := ethclient.Dial(cmd.EndpointPrivate)
	if err != nil {
		cmd.HandleError(err, "init live client")
	}

	key, err := crypto.HexToECDSA(keyRaw)
	if err != nil {
		cmd.HandleError(err, "decrypt key")
	}

	opts := cmd.GetPrivateTxOpts(key, 100000)

	token, err := cmd.InitSidechainToken(clientPrivate)
	if err != nil {
		cmd.HandleError(err, "bind live token")
	}

	address := crypto.PubkeyToAddress(key.PublicKey)
	log.Println("Address: ", address.String())

	balance, err := token.BalanceOf(&bind.CallOpts{Pending: true}, address)
	if err != nil {
		cmd.HandleError(err, "getting balance")
	}

	allowance, err := token.Allowance(&bind.CallOpts{Pending: true}, address, simple_gate.GatekeeperSidechainAddr())
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

	gkPrivate, err := cmd.InitPrivateGatekeeper(clientPrivate)
	if err != nil {
		log.Fatalln(err)
		return
	}

	tx, err := gkPrivate.PayIn(opts, value)
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(tx.Hash().String())
}
