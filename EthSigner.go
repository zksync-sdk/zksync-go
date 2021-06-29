package zksync

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
)

type EthSigner interface {
	GetAddress() common.Address
	GetPk() (*ecdsa.PrivateKey, error)
	SignMessage([]byte) ([]byte, error)
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
	account, err := wallet.Derive(path, false)
	if err != nil {
		return nil, errors.Wrap(err, "failed to derive account from HD wallet")
	}
	return &DefaultEthSigner{
		wallet:  wallet,
		account: account,
	}, nil
}

func (s *DefaultEthSigner) GetAddress() string {
	return s.account.Address.String()
}

func (s *DefaultEthSigner) SignMessage(msg []byte) ([]byte, error) {
	return s.wallet.SignText(s.account, msg)
}

func (s *DefaultEthSigner) GetPk() (*ecdsa.PrivateKey, error) {
	return s.wallet.PrivateKey(s.account)
}
