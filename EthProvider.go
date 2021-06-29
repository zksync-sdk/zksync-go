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
}

type DefaultEthProvider struct {
	client   *ethclient.Client
	contract *ZkSync.ZkSync
	auth     *bind.TransactOpts // thread unsafe usage, TODO fix it somehow
}

func (p *DefaultEthProvider) Deposit(token *Token, amount *big.Int, userAddress common.Address) (*types.Transaction, error) {
	nonce, err := p.client.PendingNonceAt(context.Background(), userAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pending nonce")
	}
	gasPrice, err := p.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get suggested gas price")
	}
	p.auth.Nonce = big.NewInt(int64(nonce))
	p.auth.Value = amount
	p.auth.GasLimit = uint64(3000000) // in units
	p.auth.GasPrice = gasPrice
	if token.IsETH() {
		return p.contract.DepositETH(p.auth, userAddress)
	} else {
		return p.contract.DepositERC20(p.auth, token.GetAddress(), amount, userAddress)
	}
}
