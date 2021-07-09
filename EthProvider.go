package zksync

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync-go/contracts/ZkSync"
	"math/big"
)

type EthProvider interface {
	Deposit(*Token, *big.Int, common.Address) (*types.Transaction, error)
	SetAuthPubkeyHash(string, uint32) (*types.Transaction, error)
	GetBalance() (*big.Int, error)
}

type DefaultEthProvider struct {
	client   *ethclient.Client
	contract *ZkSync.ZkSync
	auth     *bind.TransactOpts
}

func (p *DefaultEthProvider) Deposit(token *Token, amount *big.Int, userAddress common.Address) (*types.Transaction, error) {
	auth := p.getAuth()
	auth.Value = amount
	if token.IsETH() {
		return p.contract.DepositETH(auth, userAddress)
	} else {
		return p.contract.DepositERC20(auth, token.GetAddress(), amount, userAddress)
	}
}

func (p *DefaultEthProvider) SetAuthPubkeyHash(pubKeyHash string, zkNonce uint32) (*types.Transaction, error) {
	auth := p.getAuth()
	pkh, err := pkhToBytes(pubKeyHash)
	if err != nil {
		return nil, errors.Wrap(err, "invalid pubKeyHash value")
	}
	return p.contract.SetAuthPubkeyHash(auth, pkh, zkNonce)
}

func (p *DefaultEthProvider) GetBalance() (*big.Int, error) {
	return p.client.BalanceAt(context.Background(), p.auth.From, nil) // latest
}
func (p *DefaultEthProvider) GetNonce() (uint64, error) {
	return p.client.PendingNonceAt(context.Background(), p.auth.From) // pending
}

// getAuth make a new copy of origin TransactOpts to be used safely for each call
func (p *DefaultEthProvider) getAuth() *bind.TransactOpts {
	return &bind.TransactOpts{
		From:   p.auth.From,
		Signer: p.auth.Signer,
	}
}
