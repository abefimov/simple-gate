package simple_gate

import "github.com/ethereum/go-ethereum/common"

const PayInTopic string = "0x5bdf8f42703e019ab5248b71701aac7fc28b9c4d232dc684e53cbf1eb503ae37"
const PayOutTopic string = "0xafe5bde68f2dccddb6979d04a1dceb6099fe0ceda6bfc2c37e3a40ed8195d90f"
const SuicideTopic string = "0xa1ea9b09ea114021983e9ecf71cf2ffddfd80f5cb4f925e5bf24f9bdb5e55fde"

func PayInTopicHash() common.Hash {
	return common.HexToHash(PayInTopic)
}

func PayOutTopicHash() common.Hash {
	return common.HexToHash(PayOutTopic)
}

func SuicideTopicHash() common.Hash {
	return common.HexToHash(SuicideTopic)
}
