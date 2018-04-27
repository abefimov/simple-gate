package main

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	market2 "github.com/sonm-io/core/blockchain/market"
	api2 "github.com/sonm-io/market/api"
)

func main() {
	client, err := ethclient.Dial("https://private-dev.sonm.io")
	if err != nil {
		log.Fatalln(err)
		return
	}

	prv, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(err)
		return
	}


	market, err := api2.NewMarket(common.HexToAddress(market2.MarketAddress), client)
	if err != nil {
		log.Fatalln(err)
		return
	}

	opts := bind.NewKeyedTransactor(prv)
	opts.GasLimit = 5000000
	opts.GasPrice = big.NewInt(0)

	tx, err := market.PlaceOrder(
		opts,
		1,
		common.HexToAddress("0x0"), // counterparty
		big.NewInt(1),              // duration
		big.NewInt(1),              //price
		[3]bool{true, true, true},                          //netflags
		0,                          // identity level
		common.HexToAddress("0x0"), // blacklist
		[32]byte{0},
		[]uint64{1,2,3,4,5,6,7,8,9,10,11,12}, // benchs
	)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("Tx hash: ", tx.Hash().String())

}
