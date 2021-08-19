package zksync

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	TransactionTypeMintNFT     TransactionType = "MintNFT"
	TransactionTypeWithdrawNFT TransactionType = "WithdrawNFT"
)

type NFT struct {
	Id             uint32         `json:"id"`
	Symbol         string         `json:"symbol"`
	CreatorId      uint32         `json:"creatorId"`
	ContentHash    common.Hash    `json:"contentHash"`
	CreatorAddress common.Address `json:"creatorAddress"`
	SerialId       uint32         `json:"serialId"`
	Address        string         `json:"address"`
}

func (t *NFT) ToToken() *Token {
	return &Token{
		Id:       t.Id,
		Address:  t.Address,
		Symbol:   t.Symbol,
		Decimals: 0,
		IsNft:    true,
	}
}

type MintNFT struct {
	Type           string         `json:"type"`
	CreatorId      uint32         `json:"creatorId"`
	CreatorAddress common.Address `json:"creatorAddress"`
	ContentHash    common.Hash    `json:"contentHash"`
	Recipient      common.Address `json:"recipient"`
	Fee            string         `json:"fee"`
	FeeToken       uint32         `json:"feeToken"`
	Nonce          uint32         `json:"nonce"`
	Signature      *Signature     `json:"signature"`
}

func (t *MintNFT) getType() string {
	return "MintNFT"
}

type WithdrawNFT struct {
	Type      string         `json:"type"`
	AccountId uint32         `json:"accountId"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	Token     uint32         `json:"token"`
	FeeToken  uint32         `json:"feeToken"`
	Fee       string         `json:"fee"`
	Nonce     uint32         `json:"nonce"`
	Signature *Signature     `json:"signature"`
	*TimeRange
}

func (t *WithdrawNFT) getType() string {
	return "WithdrawNFT"
}
