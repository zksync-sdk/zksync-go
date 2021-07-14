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
	Deposit(token *Token, amount *big.Int, userAddress common.Address, options *GasOptions) (*types.Transaction, error)
	SetAuthPubkeyHash(pubKeyHash string, zkNonce uint32, options *GasOptions) (*types.Transaction, error)
	GetBalance() (*big.Int, error)
	GetNonce() (uint64, error)
	Withdraw(token *Token, amount *big.Int, options *GasOptions) (*types.Transaction, error)
	FullExit(token *Token, accountId uint32, options *GasOptions) (*types.Transaction, error)
}

type DefaultEthProvider struct {
	client   *ethclient.Client
	contract *ZkSync.ZkSync
	auth     *bind.TransactOpts
}

type GasOptions struct {
	GasPrice *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit uint64   // Gas limit to set for the transaction execution (0 = estimate)
}

func (p *DefaultEthProvider) Deposit(token *Token, amount *big.Int, userAddress common.Address, options *GasOptions) (*types.Transaction, error) {
	auth := p.getAuth(options)
	auth.Value = amount
	if token.IsETH() {
		return p.contract.DepositETH(auth, userAddress)
	} else {
		return p.contract.DepositERC20(auth, token.GetAddress(), amount, userAddress)
	}
}

func (p *DefaultEthProvider) SetAuthPubkeyHash(pubKeyHash string, zkNonce uint32, options *GasOptions) (*types.Transaction, error) {
	auth := p.getAuth(options)
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

func (p *DefaultEthProvider) Withdraw(token *Token, amount *big.Int, options *GasOptions) (*types.Transaction, error) {
	auth := p.getAuth(options)
	if token.IsETH() {
		return p.contract.WithdrawETH(auth, amount)
	} else {
		return p.contract.WithdrawERC20(auth, token.GetAddress(), amount)
	}
}

func (p *DefaultEthProvider) FullExit(token *Token, accountId uint32, options *GasOptions) (*types.Transaction, error) {
	auth := p.getAuth(options)
	return p.contract.FullExit(auth, accountId, token.GetAddress())
}

// getAuth make a new copy of origin TransactOpts to be used safely for each call
func (p *DefaultEthProvider) getAuth(options *GasOptions) *bind.TransactOpts {
	newAuth := &bind.TransactOpts{
		From:   p.auth.From,
		Signer: p.auth.Signer,
	}
	if options != nil {
		newAuth.GasPrice = options.GasPrice
		newAuth.GasLimit = options.GasLimit
	}
	return newAuth
}
