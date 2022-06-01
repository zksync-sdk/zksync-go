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
	msg, err := getChangePubKeyData(txData)
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
				Type:      EthSignatureTypeEth,
				Signature: "0x" + hex.EncodeToString(sig),
			}, nil
		}
	case "Transfer":
		if txData, ok := tx.(*Transfer); ok {
			var tokenToUse *Token
			if txData.Token != nil {
				tokenToUse = txData.Token
			} else {
				tokenToUse = token
			}
			fee, ok := big.NewInt(0).SetString(txData.Fee, 10)
			if !ok {
				return nil, errors.New("failed to convert string fee to big.Int")
			}
			msg, err := getTransferMessagePart(txData.To.String(), txData.Amount, fee, tokenToUse)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get Transfer message part")
			}
			msg += "\n" + getNonceMessagePart(nonce)
			sig, err := s.SignMessage([]byte(msg))
			if err != nil {
				return nil, errors.Wrap(err, "failed to sign Transfer tx")
			}
			return &EthSignature{
				Type:      EthSignatureTypeEth,
				Signature: "0x" + hex.EncodeToString(sig),
			}, nil
		}
	case "Withdraw":
		if txData, ok := tx.(*Withdraw); ok {
			msg, err := getWithdrawMessagePart(txData.To.String(), txData.Amount, fee, token)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get Withdraw message part")
			}
			msg += "\n" + getNonceMessagePart(nonce)
			sig, err := s.SignMessage([]byte(msg))
			if err != nil {
				return nil, errors.Wrap(err, "failed to sign Withdraw tx")
			}
			return &EthSignature{
				Type:      EthSignatureTypeEth,
				Signature: "0x" + hex.EncodeToString(sig),
			}, nil
		}
	case "ForcedExit":
		if txData, ok := tx.(*ForcedExit); ok {
			msg, err := getForcedExitMessagePart(txData.Target.String(), fee, token)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get ForcedExit message part")
			}
			msg += "\n" + getNonceMessagePart(nonce)
			sig, err := s.SignMessage([]byte(msg))
			if err != nil {
				return nil, errors.Wrap(err, "failed to sign ForcedExit tx")
			}
			return &EthSignature{
				Type:      EthSignatureTypeEth,
				Signature: "0x" + hex.EncodeToString(sig),
			}, nil
		}
	case "MintNFT":
		if txData, ok := tx.(*MintNFT); ok {
			msg, err := getMintNFTMessagePart(txData.ContentHash, txData.Recipient.String(), fee, token)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get MintNFT message part")
			}
			msg += "\n" + getNonceMessagePart(nonce)
			sig, err := s.SignMessage([]byte(msg))
			if err != nil {
				return nil, errors.Wrap(err, "failed to sign MintNFT tx")
			}
			return &EthSignature{
				Type:      EthSignatureTypeEth,
				Signature: "0x" + hex.EncodeToString(sig),
			}, nil
		}
	case "WithdrawNFT":
		if txData, ok := tx.(*WithdrawNFT); ok {
			msg, err := getWithdrawNFTMessagePart(txData.To.String(), txData.Token, fee, token)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get WithdrawNFT message part")
			}
			msg += "\n" + getNonceMessagePart(nonce)
			sig, err := s.SignMessage([]byte(msg))
			if err != nil {
				return nil, errors.Wrap(err, "failed to sign WithdrawNFT tx")
			}
			return &EthSignature{
				Type:      EthSignatureTypeEth,
				Signature: "0x" + hex.EncodeToString(sig),
			}, nil
		}
	case "Swap":
		msg := getSwapMessagePart(token, fee)
		msg += "\n" + getNonceMessagePart(nonce)
		sig, err := s.SignMessage([]byte(msg))
		if err != nil {
			return nil, errors.Wrap(err, "failed to sign Swap tx")
		}
		return &EthSignature{
			Type:      EthSignatureTypeEth,
			Signature: "0x" + hex.EncodeToString(sig),
		}, nil
	}
	return nil, errors.New("unknown tx type")
}

func (s *DefaultEthSigner) SignBatch(txs []ZksTransaction, nonce uint32, token *Token, fee *big.Int) (*EthSignature, error) {
	batchMsgs := make([]string, 0, len(txs))
	for _, tx := range txs {

		switch tx.getType() {
		case "Transfer":
			if txData, ok := tx.(*Transfer); ok {
				var tokenToUse *Token
				if txData.Token != nil {
					tokenToUse = txData.Token
				} else {
					tokenToUse = token
				}
				fee, ok := big.NewInt(0).SetString(txData.Fee, 10)
				if !ok {
					return nil, errors.New("failed to convert string fee to big.Int")
				}
				msg, err := getTransferMessagePart(txData.To.String(), txData.Amount, fee, tokenToUse)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get Transfer message part")
				}
				batchMsgs = append(batchMsgs, msg)
			}
		case "Withdraw":
			if txData, ok := tx.(*Withdraw); ok {
				msg, err := getWithdrawMessagePart(txData.To.String(), txData.Amount, fee, token)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get Withdraw message part")
				}
				batchMsgs = append(batchMsgs, msg)
			}
		case "ForcedExit":
			if txData, ok := tx.(*ForcedExit); ok {
				msg, err := getForcedExitMessagePart(txData.Target.String(), fee, token)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get ForcedExit message part")
				}
				batchMsgs = append(batchMsgs, msg)
			}
		case "MintNFT":
			if txData, ok := tx.(*MintNFT); ok {
				msg, err := getMintNFTMessagePart(txData.ContentHash, txData.Recipient.String(), fee, token)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get MintNFT message part")
				}
				batchMsgs = append(batchMsgs, msg)
			}
		case "WithdrawNFT":
			if txData, ok := tx.(*WithdrawNFT); ok {
				msg, err := getWithdrawNFTMessagePart(txData.To.String(), txData.Token, fee, token)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get WithdrawNFT message part")
				}
				batchMsgs = append(batchMsgs, msg)
			}
		default:
			return nil, errors.New("unknown tx type")
		}
	}
	batchMsg := strings.Join(batchMsgs, "\n")
	batchMsg += "\n" + getNonceMessagePart(nonce)
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
	msg, err := getOrderMessagePart(order.RecipientAddress.String(), order.Amount, sell, buy, order.Ratio)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Order message part")
	}
	msg += "\n" + getNonceMessagePart(order.Nonce)
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
