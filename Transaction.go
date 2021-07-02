package zksync

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SignedChangePubKeyTxOnchain struct {
}

type ChangePubKey struct {
	AccountId   uint32
	Account     common.Address
	NewPkHash   string // TODO []byte
	FeeToken    uint32
	Fee         *big.Int
	Nonce       uint64
	Signature   *Signature
	EthAuthData ChangePubKeyAuthType
	TimeRange   *TimeRange
}

type Signature struct {
	PubKey    string
	Signature string
}

type ChangePubKeyAuthType string

const (
	ChangePubKeyAuthTypeOnchain ChangePubKeyAuthType = `Onchain`
	ChangePubKeyAuthTypeECDSA   ChangePubKeyAuthType = `ECDSA`
	ChangePubKeyAuthTypeCREATE2 ChangePubKeyAuthType = `CREATE2`
)
