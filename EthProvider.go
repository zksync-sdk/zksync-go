package zksync

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync-go/contracts/ERC20"
	"github.com/zksync-sdk/zksync-go/contracts/ZkSync"
	"math/big"
)

type EthProvider interface {
	ApproveDeposits(token *Token, limit *big.Int, options *GasOptions) (*types.Transaction, error)
	IsDepositApproved(token *Token, userAddress common.Address, threshold *big.Int) (bool, error)
	Deposit(token *Token, amount *big.Int, userAddress common.Address, options *GasOptions) (*types.Transaction, error)
	SetAuthPubkeyHash(pubKeyHash string, zkNonce uint32, options *GasOptions) (*types.Transaction, error)
	IsOnChainAuthPubkeyHashSet(nonce uint32) (bool, error)
	GetBalance() (*big.Int, error)
	GetNonce() (uint64, error)
	FullExit(token *Token, accountId uint32, options *GasOptions) (*types.Transaction, error)
}

type DefaultEthProvider struct {
	client   *ethclient.Client
	contract *ZkSync.ZkSync
	address  common.Address
	auth     *bind.TransactOpts
}

type GasOptions struct {
	GasPrice *big.Int // Gas price to use for the transaction execution (nil = gas price oracle)
	GasLimit uint64   // Gas limit to set for the transaction execution (0 = estimate)
}

func (p *DefaultEthProvider) ApproveDeposits(token *Token, limit *big.Int, options *GasOptions) (*types.Transaction, error) {
	tokenContract, err := ERC20.NewERC20(token.GetAddress(), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load token contract")
	}
	if limit == nil {
		// max approve amount 2^256 - 1
		limit = big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	}
	auth := p.getAuth(options)
	return tokenContract.Approve(auth, p.address, limit)
}

func (p *DefaultEthProvider) IsDepositApproved(token *Token, userAddress common.Address, threshold *big.Int) (bool, error) {
	tokenContract, err := ERC20.NewERC20(token.GetAddress(), p.client)
	if err != nil {
		return false, errors.Wrap(err, "failed to load token contract")
	}
	auth := &bind.CallOpts{}
	allowed, err := tokenContract.Allowance(auth, userAddress, p.address)
	if err != nil {
		return false, errors.Wrap(err, "failed to call Allowance view of token contract")
	}
	if threshold == nil {
		// default threshold 2^255
		threshold = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(255), nil)
	}
	return allowed.Cmp(threshold) >= 0, nil
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

func (p *DefaultEthProvider) Transfer(token *Token, amount *big.Int, recipient common.Address, options *GasOptions) (*types.Transaction, error) {
	if token.IsETH() {
		return nil, errors.New("this method for ERC20 tokens transfering only")
	}
	tokenContract, err := ERC20.NewERC20(token.GetAddress(), p.client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load token contract")
	}
	auth := p.getAuth(options)
	return tokenContract.Transfer(auth, recipient, amount)
}

func (p *DefaultEthProvider) SetAuthPubkeyHash(pubKeyHash string, zkNonce uint32, options *GasOptions) (*types.Transaction, error) {
	auth := p.getAuth(options)
	pkh, err := pkhToBytes(pubKeyHash)
	if err != nil {
		return nil, errors.Wrap(err, "invalid pubKeyHash value")
	}
	return p.contract.SetAuthPubkeyHash(auth, pkh, zkNonce)
}

func (p *DefaultEthProvider) IsOnChainAuthPubkeyHashSet(nonce uint32) (bool, error) {
	opts := &bind.CallOpts{}
	publicKeyHash, err := p.contract.AuthFacts(opts, p.auth.From, nonce)
	if err != nil {
		return false, errors.Wrap(err, "failed to call AuthFacts")
	}
	return publicKeyHash != [32]byte{}, nil
}

func (p *DefaultEthProvider) GetBalance() (*big.Int, error) {
	return p.client.BalanceAt(context.Background(), p.auth.From, nil) // latest
}
func (p *DefaultEthProvider) GetNonce() (uint64, error) {
	return p.client.PendingNonceAt(context.Background(), p.auth.From) // pending
}

func (p *DefaultEthProvider) FullExit(token *Token, accountId uint32, options *GasOptions) (*types.Transaction, error) {
	auth := p.getAuth(options)
	return p.contract.RequestFullExit(auth, accountId, token.GetAddress())
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
