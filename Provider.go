package zksync

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	"math/big"
)

type Provider interface {
	GetTokens() (*Tokens, error)
	UpdateTokenSet() error
	GetTokenPrice(token *Token) (*big.Float, error)
	ContractAddress() (*ContractAddress, error)
	GetState(address common.Address) (*AccountState, error)
	GetTransactionFee(txType TransactionType, address common.Address, token *Token) (*TransactionFeeDetails, error)
	GetTransactionsBatchFee(txTypes []TransactionType, addresses []common.Address, token *Token) (*TransactionFeeDetails, error)
	SubmitTx(signedTx ZksTransaction, ethSignature *EthSignature, fastProcessing bool) (string, error)
	SubmitTxsBatch(signedTxs []*SignedTransaction, ethSignature *EthSignature) ([]string, error)
	GetTransactionDetails(txHash string) (*TransactionDetails, error)
	GetConfirmationsForEthOpAmount() (*big.Int, error)
	GetEthOpInfo(priority uint64) (*EthOpInfo, error)
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
	tokens *Tokens
}

func (p *DefaultProvider) GetTokens() (*Tokens, error) {
	if p.tokens == nil {
		if err := p.UpdateTokenSet(); err != nil {
			return nil, errors.Wrap(err, "failed to get tokens")
		}
	}
	return p.tokens, nil
}

func (p *DefaultProvider) UpdateTokenSet() error {
	res := make(map[string]*Token)
	err := p.client.Call(&res, "tokens")
	if err != nil {
		return errors.Wrap(err, "failed to call `tokens` method")
	}
	p.tokens = &Tokens{
		Tokens: res,
	}
	return nil
}

func (p *DefaultProvider) GetTokenPrice(token *Token) (*big.Float, error) {
	var resp string
	err := p.client.Call(&resp, "get_token_price", token.Symbol)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `get_token_price` method")
	}
	res, ok := big.NewFloat(0).SetString(resp)
	if !ok {
		return nil, errors.Wrap(err, "failed to parse response")
	}
	return res, nil
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

func (p *DefaultProvider) GetTransactionFee(txType TransactionType, address common.Address, token *Token) (*TransactionFeeDetails, error) {
	res := new(TransactionFeeDetails)
	err := p.client.Call(&res, "get_tx_fee", txType.getType(), address.String(), token.Symbol)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `get_tx_fee` method")
	}
	return res, nil
}

func (p *DefaultProvider) GetTransactionsBatchFee(txTypes []TransactionType, addresses []common.Address, token *Token) (*TransactionFeeDetails, error) {
	res := new(TransactionFeeDetails)
	if len(txTypes) != len(addresses) {
		return nil, errors.New("count of Transaction Types and addresses is mismatch")
	}
	txTypesList := make([]string, len(txTypes))
	addressesList := make([]string, len(addresses))
	for i, t := range txTypes {
		if txType, ok := t.getType().(string); ok {
			txTypesList[i] = txType
		} else {
			return nil, errors.New("invalid transaction Type for batch fee request")
		}
		addressesList[i] = addresses[i].String()
	}
	err := p.client.Call(&res, "get_txs_batch_fee_in_wei", txTypesList, addressesList, token.Symbol)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `get_txs_batch_fee_in_wei` method")
	}
	return res, nil
}

func (p *DefaultProvider) GetTransactionDetails(txHash string) (*TransactionDetails, error) {
	res := new(TransactionDetails)
	err := p.client.Call(&res, "tx_info", txHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `tx_info` method")
	}
	return res, nil
}

func (p *DefaultProvider) SubmitTx(signedTx ZksTransaction, ethSignature *EthSignature, fastProcessing bool) (string, error) {
	var res string
	err := p.client.Call(&res, "tx_submit", signedTx, ethSignature, fastProcessing)
	if err != nil {
		return "", errors.Wrap(err, "failed to call `tx_submit` method")
	}
	return res, nil
}

func (p *DefaultProvider) SubmitTxsBatch(signedTxs []*SignedTransaction, ethSignature *EthSignature) ([]string, error) {
	res := make([]string, len(signedTxs))
	signatures := make([]*EthSignature, 0)
	if ethSignature != nil {
		signatures = append(signatures, ethSignature)
	}
	err := p.client.Call(&res, "submit_txs_batch", signedTxs, signatures)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `submit_txs_batch` method")
	}
	return res, nil
}

func (p *DefaultProvider) GetConfirmationsForEthOpAmount() (*big.Int, error) {
	var resp int64
	err := p.client.Call(&resp, "get_confirmations_for_eth_op_amount")
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `get_confirmations_for_eth_op_amount` method")
	}
	return big.NewInt(resp), nil
}

func (p *DefaultProvider) GetEthOpInfo(priority uint64) (*EthOpInfo, error) {
	res := new(EthOpInfo)
	err := p.client.Call(&res, "ethop_info", priority)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call `ethop_info` method")
	}
	return res, nil
}
