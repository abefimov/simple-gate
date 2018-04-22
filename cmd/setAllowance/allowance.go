package setAllowance

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sonm-io/simple-gate"
	"github.com/sonm-io/simple-gate/api"
	"github.com/sonm-io/simple-gate/cmd"
	"log"
	"math/big"
)

var (
	HexKey string = ""
)

func main() {
	client, err := ethclient.Dial(cmd.Endpoint)
	if err != nil {
		cmd.HandleError(err, "init live client")
	}

	key, err := crypto.HexToECDSA(HexKey)
	if err != nil {
		cmd.HandleError(err, "decrypt key")
	}

	opts := bind.NewKeyedTransactor(key)
	opts.GasLimit = 100000
	opts.GasPrice = big.NewInt(50000000000)

	gk, err := api.NewStandardToken(common.HexToAddress(simple_gate.SNMTAddress), client)
	if err != nil {
		log.Fatalln(err)
		return
	}

	tx, err := gk.Approve(opts, common.HexToAddress(simple_gate.GatekeeperLiveAddress), big.NewInt(1000000000000000))
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(tx.Hash().String())
}
