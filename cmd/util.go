package cmd

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sonm-io/simple-gate"
	"github.com/sonm-io/simple-gate/api"
	"log"
	"math/big"
	"os"
)

var (
	EndpointPrivate string = "https://private-dev.sonm.io"
	//Endpoint        string = "https://rinkeby.infura.io/00iTrs5PIy0uGODwcsrb"
	Endpoint        string = "http://localhost:8545"
)

func HandleError(err error, message string) {
	log.Fatalf("Error: %v : %v \r\n", message, err)
	os.Exit(1)
}

func InitSidechainToken(client *ethclient.Client) (*api.StandardToken, error) {
	return api.NewStandardToken(simple_gate.SNMSidechainAddr(), client)
}

func InitLiveToken(client *ethclient.Client) (*api.StandardToken, error) {
	return api.NewStandardToken(simple_gate.SNMLiveAddr(), client)
}

func InitPrivateGatekeeper(client *ethclient.Client) (*api.Gatekeeper, error) {
	return api.NewGatekeeper(simple_gate.GatekeeperSidechainAddr(), client)
}

func InitLiveGatekeeper(client *ethclient.Client) (*api.Gatekeeper, error) {
	return api.NewGatekeeper(simple_gate.GatekeeperLiveAddr(), client)
}

func GetPrivateTxOpts(key *ecdsa.PrivateKey, gasLimit uint64) *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(key)
	opts.GasLimit = gasLimit
	opts.GasPrice = big.NewInt(0)
	return opts
}

func GetLiveTxOpts(key *ecdsa.PrivateKey, gasLimit uint64) *bind.TransactOpts {
	opts := bind.NewKeyedTransactor(key)
	opts.GasLimit = gasLimit
	opts.GasPrice = big.NewInt(50000000000)
	return opts
}
