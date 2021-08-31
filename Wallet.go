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

type Wallet struct {
	accountId  uint32
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

func (w *Wallet) GetAccountId() (uint32, error) {
	return w.accountId, nil
}

func (w *Wallet) GetPubKeyHash() (string, error) {
	return w.pubKeyHash, nil
}

func (w *Wallet) GetAddress() common.Address {
	return w.ethSigner.GetAddress()
}

func (w *Wallet) GetTokens() (*Tokens, error) {
	return w.provider.GetTokens()
}

func (w *Wallet) GetState() (*AccountState, error) {
	state, err := w.provider.GetState(w.ethSigner.GetAddress())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account state")
	}
	return state, nil
}

func (w *Wallet) GetProvider() Provider {
	return w.provider
}

func (w *Wallet) CreateEthereumProvider(client *ethclient.Client) (*DefaultEthProvider, error) {
	contractAddress, err := w.provider.ContractAddress()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contract address")
	}
	contract, err := ZkSync.NewZkSync(contractAddress.GetMainAddress(), client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init ZkSync contract instance")
	}
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain Id")
	}
	auth, err := w.newTransactorWithSigner(w.ethSigner, chainId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init TransactOpts")
	}
	return &DefaultEthProvider{
		client:   client,
		contract: contract,
		address:  contractAddress.GetMainAddress(),
		auth:     auth,
	}, nil
}

func (w *Wallet) newTransactorWithSigner(ethSigner EthSigner, chainID *big.Int) (*bind.TransactOpts, error) {
	if chainID == nil {
		return nil, bind.ErrNoChainID
	}
	keyAddr := ethSigner.GetAddress()
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, bind.ErrNotAuthorized
			}
			signature, err := ethSigner.SignHash(signer.Hash(tx).Bytes())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}, nil
}

func (w *Wallet) SetSigningKey(fee *TransactionFee, nonce uint32, onchainAuth bool, timeRange *TimeRange) (string, error) {
	if w.IsSigningKeySet() {
		return "", errors.New("current signing key is already set")
	}
	if onchainAuth {
		signedTx, err := w.buildSignedChangePubKeyTxOnchain(fee, nonce, timeRange)
		if err != nil {
			return "", errors.Wrap(err, "failed to build signed ChangePubKeyOnchain tx")
		}
		return w.submitSignedTransaction(signedTx.Transaction, nil, false)
	} else {
		signedTx, err := w.buildSignedChangePubKeyTxSigned(fee, nonce, timeRange)
		if err != nil {
			return "", errors.Wrap(err, "failed to build signed ChangePubKeySigned tx")
		}
		return w.submitSignedTransaction(signedTx.Transaction, nil, false)
	}
}

func (w *Wallet) IsSigningKeySet() bool {
	return w.pubKeyHash == w.zkSigner.GetPublicKeyHash()
}

func (w *Wallet) SyncTransfer(to common.Address, amount *big.Int, fee *TransactionFee, nonce uint32, timeRange *TimeRange) (string, error) {
	signedTx, err := w.buildSignedTransferTx(to, amount, fee, nonce, timeRange)
	if err != nil {
		return "", errors.Wrap(err, "failed to build signed Transfer tx")
	}
	return w.submitSignedTransaction(signedTx.Transaction, signedTx.EthereumSignature, false)
}

func (w *Wallet) SyncWithdraw(ethAddress common.Address, amount *big.Int, fee *TransactionFee, nonce uint32, fastProcessing bool, timeRange *TimeRange) (string, error) {
	signedTx, err := w.buildSignedWithdrawTx(ethAddress, fee.FeeToken, amount, fee, nonce, timeRange)
	if err != nil {
		return "", errors.Wrap(err, "failed to build signed Withdraw tx")
	}
	return w.submitSignedTransaction(signedTx.Transaction, signedTx.EthereumSignature, fastProcessing)
}

func (w *Wallet) SyncForcedExit(target common.Address, fee *TransactionFee, nonce uint32, timeRange *TimeRange) (string, error) {
	signedTx, err := w.buildSignedForcedExitTx(target, fee.FeeToken, fee, nonce, timeRange)
	if err != nil {
		return "", errors.Wrap(err, "failed to build signed Withdraw tx")
	}
	return w.submitSignedTransaction(signedTx.Transaction, signedTx.EthereumSignature, false)
}

func (w *Wallet) SyncMintNFT(recipient common.Address, contentHash common.Hash, fee *TransactionFee, nonce uint32) (string, error) {
	signedTx, err := w.buildSignedMintNFTTx(recipient, contentHash, fee.FeeToken, fee, nonce)
	if err != nil {
		return "", errors.Wrap(err, "failed to build signed MintNFT tx")
	}
	return w.submitSignedTransaction(signedTx.Transaction, signedTx.EthereumSignature, false)
}

func (w *Wallet) SyncWithdrawNFT(to common.Address, token *NFT, fee *TransactionFee, nonce uint32, timeRange *TimeRange) (string, error) {
	signedTx, err := w.buildSignedWithdrawNFTTx(to, token, fee.FeeToken, fee, nonce, timeRange)
	if err != nil {
		return "", errors.Wrap(err, "failed to build signed WithdrawNFT tx")
	}
	return w.submitSignedTransaction(signedTx.Transaction, signedTx.EthereumSignature, false)
}

func (w *Wallet) SyncTransferNFT(to common.Address, NFToken *NFT, fee *TransactionFee, nonce uint32, timeRange *TimeRange) ([]string, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	feeToken, err := tokens.GetToken(fee.FeeToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee Token")
	}
	// 1 transfer NFT
	transferTx := &Transfer{
		Type:      "Transfer",
		AccountId: w.accountId,
		From:      w.ethSigner.GetAddress(),
		To:        to,
		Token:     NFToken.ToToken(),
		TokenId:   NFToken.Id,
		Amount:    big.NewInt(1),
		Nonce:     nonce,
		Fee:       big.NewInt(0).String(),
		TimeRange: timeRange,
	}
	transferTx.Signature, err = w.zkSigner.SignTransfer(transferTx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of Transfer")
	}
	// 2 transfer fee
	feeTx := &Transfer{
		Type:      "Transfer",
		AccountId: w.accountId,
		From:      w.ethSigner.GetAddress(),
		To:        w.ethSigner.GetAddress(),
		Token:     feeToken,
		TokenId:   feeToken.Id,
		Amount:    big.NewInt(0),
		Nonce:     nonce + 1,
		Fee:       fee.Fee.String(),
		TimeRange: timeRange,
	}
	feeTx.Signature, err = w.zkSigner.SignTransfer(feeTx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of Transfer")
	}
	// common eth signature
	ethSig, err := w.ethSigner.SignBatch([]ZksTransaction{transferTx, feeTx}, nonce, feeToken, fee.Fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of transactions batch")
	}
	return w.submitSignedTransactionsBatch([]*SignedTransaction{{Transaction: transferTx}, {Transaction: feeTx}}, ethSig)
}

func (w *Wallet) SyncSwap(order1, order2 *Order, amount1, amount2 *big.Int, fee *TransactionFee, nonce uint32) (string, error) {
	signedTx, err := w.buildSignedSwapTx(order1, order2, amount1, amount2, fee, nonce)
	if err != nil {
		return "", errors.Wrap(err, "failed to build signed Swap tx")
	}
	return w.submitMultiSignedTransaction(signedTx.Transaction,
		signedTx.EthereumSignature, order1.EthereumSignature, order2.EthereumSignature)
}

func (w *Wallet) BuildSignedOrder(recipient common.Address, sell, buy *Token, ratio []*big.Int, amount *big.Int,
	nonce uint32, timeRange *TimeRange) (*Order, error) {
	if len(ratio) != 2 {
		return nil, errors.New("invalid ratio")
	}
	order := &Order{
		AccountId:        w.accountId,
		Amount:           amount,
		RecipientAddress: recipient,
		TokenSell:        sell.Id,
		TokenBuy:         buy.Id,
		Ratio:            ratio,
		Nonce:            nonce,
		TimeRange:        timeRange,
	}
	var err error
	order.EthereumSignature, err = w.ethSigner.SignOrder(order, sell, buy)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get eth sign of Order")
	}
	order.Signature, err = w.zkSigner.SignOrder(order)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of Order")
	}
	return order, nil
}

func (w *Wallet) BuildSignedLimitOrder(recipient common.Address, sell, buy *Token, ratio []*big.Int,
	nonce uint32, timeRange *TimeRange) (*Order, error) {
	return w.BuildSignedOrder(recipient, sell, buy, ratio, big.NewInt(0), nonce, timeRange)
}

func (w *Wallet) buildSignedChangePubKeyTxOnchain(fee *TransactionFee, nonce uint32, timeRange *TimeRange) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	token, err := tokens.GetToken(fee.FeeToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee token")
	}
	txData := &ChangePubKey{
		Type:        "ChangePubKey",
		AccountId:   w.accountId,
		Account:     w.ethSigner.GetAddress(),
		NewPkHash:   w.zkSigner.GetPublicKeyHash(),
		Nonce:       nonce,
		FeeToken:    token.Id,
		Fee:         fee.Fee.String(),
		EthAuthData: &ChangePubKeyOnchain{Type: ChangePubKeyAuthTypeOnchain},
		TimeRange:   timeRange,
	}
	txData.Signature, err = w.zkSigner.SignChangePubKey(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of ChangePubKey")
	}
	return &SignedTransaction{
		Transaction: txData,
	}, nil
}

func (w *Wallet) buildSignedChangePubKeyTxSigned(fee *TransactionFee, nonce uint32, timeRange *TimeRange) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	token, err := tokens.GetToken(fee.FeeToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee token")
	}
	txData := &ChangePubKey{
		Type:      "ChangePubKey",
		AccountId: w.accountId,
		Account:   w.ethSigner.GetAddress(),
		NewPkHash: w.zkSigner.GetPublicKeyHash(),
		Nonce:     nonce,
		FeeToken:  token.Id,
		Fee:       fee.Fee.String(),
		TimeRange: timeRange,
	}
	auth, err := w.ethSigner.SignAuth(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get eth sign of auth data")
	}
	txData.EthAuthData = auth
	txData.Signature, err = w.zkSigner.SignChangePubKey(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of ChangePubKey")
	}
	return &SignedTransaction{
		Transaction: txData,
	}, nil
}

func (w *Wallet) buildSignedTransferTx(to common.Address, amount *big.Int, fee *TransactionFee, nonce uint32, timeRange *TimeRange) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	token, err := tokens.GetToken(fee.FeeToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee token")
	}
	txData := &Transfer{
		Type:      "Transfer",
		AccountId: w.accountId,
		From:      w.ethSigner.GetAddress(),
		To:        to,
		Token:     token,
		TokenId:   token.Id,
		Amount:    amount,
		Nonce:     nonce,
		Fee:       fee.Fee.String(),
		TimeRange: timeRange,
	}
	ethSig, err := w.ethSigner.SignTransaction(txData, nonce, token, fee.Fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of transaction")
	}
	txData.Signature, err = w.zkSigner.SignTransfer(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of Transfer")
	}
	return &SignedTransaction{
		Transaction:       txData,
		EthereumSignature: ethSig,
	}, nil
}

func (w *Wallet) buildSignedWithdrawTx(to common.Address, tokenId string, amount *big.Int, fee *TransactionFee, nonce uint32, timeRange *TimeRange) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	token, err := tokens.GetToken(tokenId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee token")
	}
	txData := &Withdraw{
		Type:      "Withdraw",
		AccountId: w.accountId,
		From:      w.ethSigner.GetAddress(),
		To:        to,
		TokenId:   token.Id,
		Amount:    amount,
		Nonce:     nonce,
		Fee:       fee.Fee.String(),
		TimeRange: timeRange,
	}
	ethSig, err := w.ethSigner.SignTransaction(txData, nonce, token, fee.Fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of transaction")
	}
	txData.Signature, err = w.zkSigner.SignWithdraw(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of Withdraw")
	}
	return &SignedTransaction{
		Transaction:       txData,
		EthereumSignature: ethSig,
	}, nil
}

func (w *Wallet) buildSignedForcedExitTx(target common.Address, tokenId string, fee *TransactionFee, nonce uint32, timeRange *TimeRange) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	token, err := tokens.GetToken(tokenId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee token")
	}
	txData := &ForcedExit{
		Type:      "ForcedExit",
		AccountId: w.accountId,
		Target:    target,
		TokenId:   token.Id,
		Nonce:     nonce,
		Fee:       fee.Fee.String(),
		TimeRange: timeRange,
	}
	ethSig, err := w.ethSigner.SignTransaction(txData, nonce, token, fee.Fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of transaction")
	}
	txData.Signature, err = w.zkSigner.SignForcedExit(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of ForcedExit")
	}
	return &SignedTransaction{
		Transaction:       txData,
		EthereumSignature: ethSig,
	}, nil
}

func (w *Wallet) buildSignedMintNFTTx(to common.Address, contentHash common.Hash, feeTokenId string, fee *TransactionFee, nonce uint32) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	feeToken, err := tokens.GetToken(feeTokenId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee Token")
	}
	txData := &MintNFT{
		Type:           "MintNFT",
		CreatorId:      w.accountId,
		CreatorAddress: w.GetAddress(),
		ContentHash:    contentHash,
		Recipient:      to,
		Nonce:          nonce,
		Fee:            fee.Fee.String(),
		FeeToken:       feeToken.Id,
	}
	ethSig, err := w.ethSigner.SignTransaction(txData, nonce, feeToken, fee.Fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of transaction")
	}
	txData.Signature, err = w.zkSigner.SignMintNFT(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of MintNFT")
	}
	return &SignedTransaction{
		Transaction:       txData,
		EthereumSignature: ethSig,
	}, nil
}

func (w *Wallet) buildSignedWithdrawNFTTx(to common.Address, NFToken *NFT, feeTokenId string, fee *TransactionFee, nonce uint32, timeRange *TimeRange) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	feeToken, err := tokens.GetToken(feeTokenId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee Token")
	}
	txData := &WithdrawNFT{
		Type:      "WithdrawNFT",
		AccountId: w.accountId,
		From:      w.GetAddress(),
		To:        to,
		Token:     NFToken.Id,
		Nonce:     nonce,
		Fee:       fee.Fee.String(),
		FeeToken:  feeToken.Id,
		TimeRange: timeRange,
	}
	ethSig, err := w.ethSigner.SignTransaction(txData, nonce, feeToken, fee.Fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of transaction")
	}
	txData.Signature, err = w.zkSigner.SignWithdrawNFT(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of WithdrawNFT")
	}
	return &SignedTransaction{
		Transaction:       txData,
		EthereumSignature: ethSig,
	}, nil
}

func (w *Wallet) buildSignedSwapTx(order1, order2 *Order, amount1, amount2 *big.Int, fee *TransactionFee, nonce uint32) (*SignedTransaction, error) {
	tokens, err := w.provider.GetTokens()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}
	token, err := tokens.GetToken(fee.FeeToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get fee token")
	}
	txData := &Swap{
		Type:             "Swap",
		Orders:           []*Order{order1, order2},
		SubmitterId:      w.accountId,
		SubmitterAddress: w.ethSigner.GetAddress(),
		Amounts:          []*big.Int{amount1, amount2},
		Nonce:            nonce,
		Fee:              fee.Fee.String(),
		FeeToken:         token.Id,
	}
	ethSig, err := w.ethSigner.SignTransaction(txData, nonce, token, fee.Fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get eth sign of swap transaction")
	}
	txData.Signature, err = w.zkSigner.SignSwap(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign of Swap tx")
	}
	return &SignedTransaction{
		Transaction:       txData,
		EthereumSignature: ethSig,
	}, nil
}

func (w *Wallet) submitSignedTransaction(tx ZksTransaction, ethSignature *EthSignature, fastProcessing bool) (string, error) {
	return w.provider.SubmitTx(tx, ethSignature, fastProcessing)
}

func (w *Wallet) submitMultiSignedTransaction(tx ZksTransaction, ethSignatures ...*EthSignature) (string, error) {
	return w.provider.SubmitTxMultiSig(tx, ethSignatures...)
}

func (w *Wallet) submitSignedTransactionsBatch(txs []*SignedTransaction, ethSignature *EthSignature) ([]string, error) {
	return w.provider.SubmitTxsBatch(txs, ethSignature)
}
