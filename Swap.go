package zksync

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Order struct {
	AccountId         uint32         `json:"accountId"`
	RecipientAddress  common.Address `json:"recipient"`
	Nonce             uint32         `json:"nonce"`
	TokenBuy          uint32         `json:"tokenBuy"`
	TokenSell         uint32         `json:"tokenSell"`
	Ratio             []*big.Int     `json:"ratio"`
	Amount            *big.Int       `json:"amount"`
	Signature         *Signature     `json:"signature"`
	EthereumSignature *EthSignature  `json:"ethereumSignature"`
	*TimeRange
}

const (
	TransactionTypeSwap TransactionType = "Swap"
)

type Swap struct {
	Type             string         `json:"type"`
	SubmitterId      uint32         `json:"submitterId"`
	SubmitterAddress common.Address `json:"submitterAddress"`
	Nonce            uint32         `json:"nonce"`
	Orders           []*Order       `json:"orders"`
	Amounts          []*big.Int     `json:"amounts"`
	Fee              string         `json:"fee"`
	FeeToken         uint32         `json:"feeToken"`
	Signature        *Signature     `json:"signature"`
	*TimeRange
}

func (t *Swap) getType() string {
	return "Swap"
}
