package zksync

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
	"math/big"
)

type EthSigner interface {
	GetAddress() common.Address
	GetPk() (*ecdsa.PrivateKey, error)
	SignMessage([]byte) ([]byte, error)
	SignAuth(txData *ChangePubKey) (*ChangePubKeyECDSA, error)
	SignTransaction(tx ZksTransaction, nonce uint32, token *Token, fee *big.Int) (*EthSignature, error)
}

type DefaultEthSigner struct {
	wallet  *hdwallet.Wallet
	account accounts.Account
}

func NewEthSignerFromMnemonic(mnemonic string) (*DefaultEthSigner, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create HD wallet from mnemonic")
	}
	path, err := accounts.ParseDerivationPath("m/44'/60'/0'/0/0")
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse derivation path")
	}
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, errors.Wrap(err, "failed to derive account from HD wallet")
	}
	return &DefaultEthSigner{
		wallet:  wallet,
		account: account,
	}, nil
}

func (s *DefaultEthSigner) GetAddress() common.Address {
	return s.account.Address
}

func (s *DefaultEthSigner) SignMessage(msg []byte) ([]byte, error) {
	return s.wallet.SignText(s.account, msg) // prefixed
}

func (s *DefaultEthSigner) SignAuth(txData *ChangePubKey) (*ChangePubKeyECDSA, error) {
	auth := &ChangePubKeyECDSA{
		Type:         ChangePubKeyAuthTypeECDSA,
		EthSignature: "",
		BatchHash:    hex.EncodeToString(make([]byte, 32)),
	}
	txData.EthAuthData = auth
	msg, err := getChangePubKeyData(txData)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ChangePubKey data for sign")
	}
	sig, err := s.SignMessage(msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign ChangePubKeyECDSA msg")
	}
	auth.EthSignature = hex.EncodeToString(sig)
	return auth, nil
}

func (s *DefaultEthSigner) SignTransaction(tx ZksTransaction, nonce uint32, token *Token, fee *big.Int) (*EthSignature, error) {
	switch tx.getType() {
	case "ChangePubKey":
		if txData, ok := tx.(*ChangePubKey); ok {
			msg, err := getChangePubKeyData(txData)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get ChangePubKey data for sign")
			}
			sig, err := s.SignMessage(msg)
			if err != nil {
				return nil, errors.Wrap(err, "failed to sign ChangePubKey tx")
			}
			return &EthSignature{
				sigType:   EthSignatureTypeEth,
				signature: hex.EncodeToString(sig),
			}, nil
		}
	}
	return nil, errors.New("unknown tx type")
}

func (s *DefaultEthSigner) GetPk() (*ecdsa.PrivateKey, error) {
	return s.wallet.PrivateKey(s.account)
}

type EthSignatureType string

const (
	EthSignatureTypeEth     EthSignatureType = "EthereumSignature"
	EthSignatureTypeEIP1271 EthSignatureType = "EIP1271Signature"
)

type EthSignature struct {
	sigType   EthSignatureType
	signature string
}
