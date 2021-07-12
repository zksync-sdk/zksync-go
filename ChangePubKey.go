package zksync

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
)

type TransactionTypeChangePubKey struct {
	ChangePubKey string `json:"ChangePubKey"`
}

const (
	TransactionTypeChangePubKeyOnchain TransactionType = "Onchain"
	TransactionTypeChangePubKeyECDSA   TransactionType = "ECDSA"
	TransactionTypeChangePubKeyCREATE2 TransactionType = "CREATE2"
)

type ChangePubKey struct {
	Type        string              `json:"type"`
	AccountId   uint32              `json:"accountId"`
	Account     common.Address      `json:"account"`
	NewPkHash   string              `json:"newPkHash"`
	FeeToken    uint32              `json:"feeToken"`
	Fee         string              `json:"fee"`
	Nonce       uint32              `json:"nonce"`
	Signature   *Signature          `json:"signature"`
	EthAuthData ChangePubKeyVariant `json:"ethAuthData"`
	*TimeRange
}

func (t *ChangePubKey) getType() string {
	return "ChangePubKey"
}

type ChangePubKeyAuthType string

const (
	ChangePubKeyAuthTypeOnchain ChangePubKeyAuthType = `Onchain`
	ChangePubKeyAuthTypeECDSA   ChangePubKeyAuthType = `ECDSA`
	ChangePubKeyAuthTypeCREATE2 ChangePubKeyAuthType = `CREATE2`
)

type ChangePubKeyVariant interface {
	getType() ChangePubKeyAuthType
	getBytes() []byte
}

type ChangePubKeyOnchain struct {
	Type ChangePubKeyAuthType `json:"type"`
}

func (t *ChangePubKeyOnchain) getType() ChangePubKeyAuthType {
	return ChangePubKeyAuthTypeOnchain
}

func (t *ChangePubKeyOnchain) getBytes() []byte {
	return make([]byte, 32)
}

type ChangePubKeyECDSA struct {
	Type         ChangePubKeyAuthType `json:"type"`
	EthSignature string               `json:"ethSignature"`
	BatchHash    string               `json:"batchHash"`
}

func (t *ChangePubKeyECDSA) getType() ChangePubKeyAuthType {
	return ChangePubKeyAuthTypeECDSA
}

func (t *ChangePubKeyECDSA) getBytes() []byte {
	res, _ := hex.DecodeString(t.BatchHash[2:])
	return res
}

type ChangePubKeyCREATE2 struct {
	Type           ChangePubKeyAuthType `json:"type"`
	CreatorAddress string               `json:"creatorAddress"`
	SaltArg        string               `json:"saltArg"`
	CodeHash       string               `json:"codeHash"`
}

func (t *ChangePubKeyCREATE2) getType() ChangePubKeyAuthType {
	return ChangePubKeyAuthTypeCREATE2
}

//func (t *ChangePubKeyCREATE2) getBytes() []byte {
//	return make([]byte, 32)
//}

type Signature struct {
	PubKey    string `json:"pubKey"`
	Signature string `json:"signature"`
}
