package zksync

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

type Provider interface {
	GetTokens() (*Tokens, error)
	ContractAddress() (*ContractAddress, error)
	GetState(address common.Address) (*AccountState, error)
}

func NewDefaultProvider(rawUrl string) (*DefaultProvider, error) {
	client, err := rpc.Dial(rawUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial client")
	}
	return &DefaultProvider{
		client: client,
	}, nil
}

func NewDefaultProviderFor(cid ChainId) (*DefaultProvider, error) {
	var rawUrl string
	switch cid {
	case ChainIdMainnet:
		rawUrl = `https://api.zksync.io/jsrpc`
	case ChainIdRinkeby:
		rawUrl = `https://rinkeby-api.zksync.io/jsrpc`
	case ChainIdRopsten:
		rawUrl = `https://ropsten-api.zksync.io/jsrpc`
	case ChainIdLocalhost:
		rawUrl = `http://127.0.0.1:3030`
	}
	return NewDefaultProvider(rawUrl)
}

type DefaultProvider struct {
	client *rpc.Client
}

func (p *DefaultProvider) GetTokens() (*Tokens, error) {
	res := make(map[string]*Token)
	err := p.client.Call(&res, "tokens")
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `tokens` method")
	}
	return &Tokens{
		Tokens: res,
	}, nil
}

func (p *DefaultProvider) ContractAddress() (*ContractAddress, error) {
	res := new(ContractAddress)
	err := p.client.Call(&res, "contract_address")
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `contract_address` method")
	}
	return res, nil
}

func (p *DefaultProvider) GetState(address common.Address) (*AccountState, error) {
	res := new(AccountState)
	err := p.client.Call(&res, "account_info", address.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `account_info` method")
	}
	return res, nil
}

func (p *DefaultProvider) GetTransactionFee(txType string, address common.Address, token *Token) (*TransactionFeeDetails, error) {
	res := new(TransactionFeeDetails)
	err := p.client.Call(&res, "get_tx_fee", txType, address.String(), token.Symbol)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `get_tx_fee` method")
	}
	return res, nil
}

func (p *DefaultProvider) SubmitTx(txType string, address common.Address, token *Token) (*TransactionFeeDetails, error) {
	res := new(TransactionFeeDetails)
	err := p.client.Call(&res, "tx_submit", txType, address.String(), token.Symbol)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `get_tx_fee` method")
	}
	return res, nil
}
