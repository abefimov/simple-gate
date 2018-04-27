package simple_gate

import "github.com/ethereum/go-ethereum/common"

const GatekeeperSidechainAddress string = "0x74cd6dcaae93e61964f60d63925ac3853499b654"
const SNMSidechainAddress string = "0x5db024c6469634f4b307135cb76e8e6806f007b3"

const GatekeeperLiveAddress string = "0x8Ae371b723B2e32333FcfAD7733B6bbd64a4EA6F"
const SNMAddress string = "0x06bda3cf79946e8b32a0bb6a3daa174b577c55b5"

func GatekeeperSidechainAddr() common.Address {
	return common.HexToAddress(GatekeeperSidechainAddress)
}

func SNMSidechainAddr() common.Address {
	return common.HexToAddress(SNMSidechainAddress)
}

func GatekeeperLiveAddr() common.Address {
	return common.HexToAddress(GatekeeperLiveAddress)
}

func SNMLiveAddr() common.Address {
	return common.HexToAddress(SNMAddress)
}
