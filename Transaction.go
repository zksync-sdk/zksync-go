package zksync

import "github.com/ethereum/go-ethereum/common"

type SignedChangePubKeyTxOnchain struct {
}

type ChangePubKey struct {
	AccountId uint64
	Account   common.Address
	NewPkHash string
	FeeToken  int
	Fee       string
	Nonce     uint64
	Signature *Signature
	//EthAuthData
	TimeRange *TimeRange
}

type Signature struct {
	PubKey    string
	Signature string
}
