package zksync

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"math/big"
)

type Token struct {
	Id       uint32 `json:"id"`
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals uint   `json:"decimals"`
	IsNft    bool   `json:"is_nft"`
}

func CreateETH() *Token {
	return &Token{
		Id:       0,
		Address:  common.Address{}.String(),
		Symbol:   `ETH`,
		Decimals: 18,
	}
}

func (t Token) IsETH() bool {
	return t.Address == common.Address{}.String() && t.Symbol == `ETH`
}

func (t Token) GetAddress() common.Address {
	return common.HexToAddress(t.Address)
}

func (t Token) ToBigFloat(amount *big.Int) *big.Float {
	amountFloat := big.NewFloat(0).SetInt(amount)
	divider := big.NewFloat(0).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(t.Decimals)), nil)) // 10^decimals
	return big.NewFloat(0).SetPrec(t.Decimals*8).Quo(amountFloat, divider)                                   // amount / 10^decimals
}

type Tokens struct {
	Tokens map[string]*Token
}

func (ts *Tokens) GetToken(id string) (*Token, error) {
	if t, ok := ts.Tokens[id]; ok {
		return t, nil
	}
	// suppose id is address
	for _, t := range ts.Tokens {
		if t.Address == id {
			return t, nil
		}
	}
	return nil, errors.New("token not found")
}
