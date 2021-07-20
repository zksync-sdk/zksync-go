package zksync

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	TransactionTypeTransfer TransactionType = "Transfer"
)

type Transfer struct {
	Type      string         `json:"type"`
	AccountId uint32         `json:"accountId"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Token     *Token         `json:"-"`
	TokenId   uint32         `json:"token"`
	Amount    *big.Int       `json:"amount"`
	Fee       string         `json:"fee"`
	Nonce     uint32         `json:"nonce"`
	Signature *Signature     `json:"signature"`
	*TimeRange
}

func (t *Transfer) getType() string {
	return "Transfer"
}
