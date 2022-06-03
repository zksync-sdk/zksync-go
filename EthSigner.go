package zksync

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

type EthSigner interface {
	GetAddress() common.Address
	SignMessage([]byte) ([]byte, error)
	SignHash(msg []byte) ([]byte, error)
	SignAuth(txData *ChangePubKey) (*ChangePubKeyECDSA, error)
	SignTransaction(tx ZksTransaction, nonce uint32, token *Token, fee *big.Int) (*EthSignature, error)
	SignBatch(txs []ZksTransaction, nonce uint32, token *Token, fee *big.Int) (*EthSignature, error)
	SignOrder(order *Order, sell, buy *Token) (*EthSignature, error)
}

type DefaultEthSigner struct {
	pk      *ecdsa.PrivateKey
	address common.Address
}

func NewEthSignerFromMnemonic(mnemonic string) (*DefaultEthSigner, error) {
	return NewEthSignerFromMnemonicAndAccountId(mnemonic, 0)
}

func NewEthSignerFromMnemonicAndAccountId(mnemonic string, accountId uint32) (*DefaultEthSigner, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HD wallet from mnemonic")
	}
	path, err := accounts.ParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%d", accountId))
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse derivation path")
	}
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to derive account from HD wallet")
	}
	pk, err := wallet.PrivateKey(account)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get account's private key from HD wallet")
	}
	pub := pk.Public().(*ecdsa.PublicKey)
	return &DefaultEthSigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
	}, nil
}

func NewEthSignerFromRawPrivateKey(rawPk []byte) (*DefaultEthSigner, error) {
	pk, err := crypto.ToECDSA(rawPk)
	if err != nil {
		return nil, errors.Wrap(err, "invalid raw private key")
	}
	pub := pk.Public().(*ecdsa.PublicKey)
	return &DefaultEthSigner{
		pk:      pk,
		address: crypto.PubkeyToAddress(*pub),
	}, nil
}

func (s *DefaultEthSigner) GetAddress() common.Address {
	return s.address
}

func (s *DefaultEthSigner) SignMessage(msg []byte) ([]byte, error) {
	sig, err := s.SignHash(accounts.TextHash(msg)) // prefixed
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}
	// set recovery byte to 27/28
	if len(sig) == 65 {
		sig[64] += 27
	}
	return sig, nil
}

func (s *DefaultEthSigner) SignHash(msg []byte) ([]byte, error) {
	sig, err := crypto.Sign(msg, s.pk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign hash")
	}
	return sig, nil
}

func (s *DefaultEthSigner) SignAuth(txData *ChangePubKey) (*ChangePubKeyECDSA, error) {
	auth := &ChangePubKeyECDSA{
		Type:         ChangePubKeyAuthTypeECDSA,
		EthSignature: "",
		BatchHash:    "0x" + hex.EncodeToString(make([]byte, 32)),
	}
	txData.EthAuthData = auth
	msg, err := GetChangePubKeyData(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ChangePubKey data for sign")
	}
	sig, err := s.SignMessage(msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign ChangePubKeyECDSA msg")
	}
	auth.EthSignature = "0x" + hex.EncodeToString(sig)
	return auth, nil
}

func (s *DefaultEthSigner) SignTransaction(tx ZksTransaction, nonce uint32, token *Token, fee *big.Int) (*EthSignature, error) {
	msg, err := GetSignMessage(tx, token, fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get sign message for tx")
	}
	msg += "\n" + GetNonceMessagePart(nonce)
	sig, err := s.SignMessage([]byte(msg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign tx")
	}
	return &EthSignature{
		Type:      EthSignatureTypeEth,
		Signature: "0x" + hex.EncodeToString(sig),
	}, nil
}

func (s *DefaultEthSigner) SignBatch(txs []ZksTransaction, nonce uint32, token *Token, fee *big.Int) (*EthSignature, error) {
	batchMsgs := make([]string, 0, len(txs))
	for _, tx := range txs {
		msg, err := GetSignMessage(tx, token, fee)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get sign message for one of txs")
		}
		batchMsgs = append(batchMsgs, msg)
	}
	batchMsg := strings.Join(batchMsgs, "\n")
	batchMsg += "\n" + GetNonceMessagePart(nonce)
	sig, err := s.SignMessage([]byte(batchMsg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign batch of txs")
	}
	return &EthSignature{
		Type:      EthSignatureTypeEth,
		Signature: "0x" + hex.EncodeToString(sig),
	}, nil
}

func (s *DefaultEthSigner) SignOrder(order *Order, sell, buy *Token) (*EthSignature, error) {
	msg, err := GetOrderMessagePart(order.RecipientAddress.String(), order.Amount, sell, buy, order.Ratio)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Order message part")
	}
	msg += "\n" + GetNonceMessagePart(order.Nonce)
	sig, err := s.SignMessage([]byte(msg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign Order")
	}
	return &EthSignature{
		Type:      EthSignatureTypeEth,
		Signature: "0x" + hex.EncodeToString(sig),
	}, nil
}

type EthSignatureType string

const (
	EthSignatureTypeEth     EthSignatureType = "EthereumSignature"
	EthSignatureTypeEIP1271 EthSignatureType = "EIP1271Signature"
)

type EthSignature struct {
	Type      EthSignatureType `json:"type"`
	Signature string           `json:"signature"`
}
