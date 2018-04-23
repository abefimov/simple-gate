package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sonm-io/simple-gate"
	"github.com/sonm-io/simple-gate/cmd"
	"log"
	"math/big"
	"os"
)

var (
	HexKey string = ""
)

type PayTrx struct {
	from     common.Address
	txNumber *big.Int
	value    *big.Int
}

func main() {
	log.Println("Starting")

	client, err := ethclient.Dial(cmd.Endpoint)
	if err != nil {
		cmd.HandleError(err, "init live client")
	}

	clientPrivate, err := ethclient.Dial(cmd.EndpointPrivate)
	if err != nil {
		cmd.HandleError(err, "init private client")
	}

	payInMap := make(map[string]PayTrx)
	payOutMap := make(map[string]PayTrx)
	payInNotPayOut := make(map[string]PayTrx)

	// Getting PayIn transactions
	log.Println("Getting PayIn transactions")
	logs, err := clientPrivate.FilterLogs(context.Background(), ethereum.FilterQuery{
		Topics:    [][]common.Hash{{simple_gate.PayInTopicHash()}},
		Addresses: []common.Address{simple_gate.GatekeeperAddr()},
	})

	for _, l := range logs {
		if l.Topics[0].String() == simple_gate.PayOutTopicHash().String() {
			payOutMap[l.Topics[2].Big().String()] = PayTrx{
				from:     common.HexToAddress(l.Topics[1].Hex()),
				txNumber: l.Topics[2].Big(),
				value:    l.Topics[3].Big(),
			}
		}
	}

	// Getting already paid transactions
	log.Println("Getting PayOut transactions")
	logs, err = client.FilterLogs(context.Background(), ethereum.FilterQuery{
		Topics:    [][]common.Hash{{simple_gate.PayOutTopicHash()}},
		Addresses: []common.Address{simple_gate.GatekeeperLiveAddr()},
	})

	for _, l := range logs {
		if l.Topics[0].String() == simple_gate.PayInTopicHash().String() {
			payInMap[l.Topics[2].Big().String()] = PayTrx{
				from:     common.HexToAddress(l.Topics[1].Hex()),
				txNumber: l.Topics[2].Big(),
				value:    l.Topics[3].Big(),
			}
		}
	}

	// Getting already paid transactions
	log.Println("Match transactions")

	var unpaidAmount = 0

	for key := range payInMap {
		_, ok := payOutMap[key]
		if !ok {
			t := payInMap[key]
			payInNotPayOut[key] = t
			unpaidAmount++
		}
	}

	if unpaidAmount == 0 {
		successExit()
	}

	// Sends entire transactions
	key, err := crypto.HexToECDSA(HexKey)
	if err != nil {
		cmd.HandleError(err, "decrypt key")
	}

	gkPrivate, err := cmd.InitPrivateGatekeeper(clientPrivate)
	if err != nil {
		cmd.HandleError(err, "bind private gatekeeper")
	}

	opts := cmd.GetPrivateTxOpts(key, 100000)
	for key := range payInNotPayOut {
		t := payInNotPayOut[key]
		tx, err := gkPrivate.Payout(opts, t.from, t.value, t.txNumber)
		if err != nil {
			cmd.HandleError(err, "send payout transaction")
		}
		log.Println("Fix new payout:")
		log.Printf("To: %v ; Value: %v ; txNum: %v \r\n", t.from.String(), t.value.String(), t.txNumber.String())
		log.Printf("Transaction hash: %v \r\n", tx.Hash().String())
	}

	successExit()
}

func successExit() {
	log.Println("All transations fixed")
	os.Exit(0)
}
