package zksync

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync-go/contracts/ZkSync"
)

type Wallet struct {
	accountId  uint64
	pubKeyHash string
	zkSigner   *ZkSigner
	ethSigner  EthSigner
	provider   Provider
}

func NewWallet(ethSigner EthSigner, zkSigner *ZkSigner, provider Provider) (*Wallet, error) {
	state, err := provider.GetState(ethSigner.GetAddress())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account state")
	}
	return &Wallet{
		accountId:  state.Id,
		pubKeyHash: state.Committed.PubKeyHash,
		zkSigner:   zkSigner,
		ethSigner:  ethSigner,
		provider:   provider,
	}, nil
}

func (w *Wallet) CreateEthereumProvider(client *ethclient.Client) (EthProvider, error) {
	contractAddress, err := w.provider.ContractAddress()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contract address")
	}
	contract, err := ZkSync.NewZkSync(contractAddress.GetMainAddress(), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init ZkSync contract instance")
	}
	pk, err := w.ethSigner.GetPk()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get private key from ethSigner")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain Id")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init TransactOpts")
	}
	return &DefaultEthProvider{
		client:   client,
		contract: contract,
		auth:     auth,
	}, nil
}

func (w *Wallet) SetSigningKey(fee TransactionFee, nonce uint64, onchainAuth bool, timeRange *TimeRange) (string, error) {
	if w.IsSigningKeySet() {
		return "", errors.New("current signing key is already set")
	}
	if onchainAuth {
		signedTx, err := w.buildSignedChangePubKeyTxOnchain(fee, nonce, timeRange)
	} else {

	}
}

func (w *Wallet) IsSigningKeySet() bool {
	return w.pubKeyHash == w.zkSigner.GetPublicKeyHash()
}

func (w *Wallet) SyncTransfer(fee TransactionFee, nonce uint64, timeRange *TimeRange) {

}

func (w *Wallet) buildSignedChangePubKeyTxOnchain(fee TransactionFee, nonce uint64, timeRange *TimeRange) (*SignedChangePubKeyTxOnchain, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	token, err := tokens.GetToken(fee.FeeToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee token")
	}
	changePubKey := &ChangePubKey{
		AccountId: w.accountId,
		Account:   w.ethSigner.GetAddress(),
		NewPkHash: w.zkSigner.GetPublicKeyHash(),
		Nonce:     nonce,
		FeeToken:  token.Id,
		Fee:       fee.Fee.String(),
		//EthAuthData
		TimeRange: timeRange,
	}
}
