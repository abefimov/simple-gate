package simple_gate

import "github.com/ethereum/go-ethereum/common"

const GatekeeperAddress string = "0xe5267ea22d91139d906d03575f2e02503af0700e"
const SNMLAddress string = "0xf66d9216f69feda8b1c7a36d0ce0ef586c1b0351"

const GatekeeperLiveAddress string = "0x81a46d5ea60ceb1b1cae6fe536e801e9eceb13db"
const SNMTAddress string = "0x06bda3cf79946e8b32a0bb6a3daa174b577c55b5"

func GatekeeperAddr() common.Address {
	return common.HexToAddress(GatekeeperAddress)
}

func SNMLAddr() common.Address {
	return common.HexToAddress(SNMLAddress)
}

func GatekeeperLiveAddr() common.Address {
	return common.HexToAddress(GatekeeperLiveAddress)
}

func SNMTLiveAddr() common.Address {
	return common.HexToAddress(SNMTAddress)
}
