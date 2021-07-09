package zksync

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync-sdk-go"
	"math/big"
)

const (
	Message                 = "Access zkSync account.\n\nOnly sign this message for a trusted client!"
	TransactionVersion byte = 0x01
)

func NewZkSignerFromSeed(seed []byte) (*ZkSigner, error) {
	privateKey, err := zkscrypto.NewPrivateKey(seed)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create private key")
	}
	publicKey, err := privateKey.PublicKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create public key")
	}
	publicKeyHash, err := publicKey.Hash()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get public key hash")
	}
	return &ZkSigner{
		privateKey:    privateKey,
		publicKey:     publicKey,
		publicKeyHash: publicKeyHash,
	}, nil
}

func NewZkSignerFromRawPrivateKey(rawPk []byte) (*ZkSigner, error) {
	privateKey, err := zkscrypto.NewPrivateKeyRaw(rawPk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create private key from raw bytes")
	}
	publicKey, err := privateKey.PublicKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create public key")
	}
	publicKeyHash, err := publicKey.Hash()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get public key hash")
	}
	return &ZkSigner{
		privateKey:    privateKey,
		publicKey:     publicKey,
		publicKeyHash: publicKeyHash,
	}, nil
}

func NewZkSignerFromEthSigner(es EthSigner, cid ChainId) (*ZkSigner, error) {
	signMsg := Message
	if cid != ChainIdMainnet {
		signMsg = fmt.Sprintf("%s\nChain ID: %d.", Message, cid)
	}
	sig, err := es.SignMessage([]byte(signMsg))
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign special message")
	}
	return NewZkSignerFromSeed(sig)
}

type ZkSigner struct {
	privateKey    *zkscrypto.PrivateKey
	publicKey     *zkscrypto.PublicKey
	publicKeyHash *zkscrypto.PublicKeyHash
}

func (s *ZkSigner) Sign(message []byte) (*zkscrypto.Signature, error) {
	signature, err := s.privateKey.Sign(message)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign message")
	}
	return signature, nil
}

func (s *ZkSigner) SignChangePubKey(txData *ChangePubKey) (*Signature, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0xff - 0x07)
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(txData.AccountId))
	buf.Write(txData.Account[:])
	pkhBytes, err := pkhToBytes(txData.NewPkHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pkh bytes")
	}
	buf.Write(pkhBytes)
	buf.Write(Uint32ToBytes(txData.FeeToken))
	fee, ok := big.NewInt(0).SetString(txData.Fee, 10)
	if !ok {
		return nil, errors.New("failed to convert string fee to big.Int")
	}
	packedFee, err := packFee(fee)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack fee")
	}
	buf.Write(packedFee)
	buf.Write(Uint32ToBytes(txData.Nonce))
	buf.Write(Uint64ToBytes(txData.TimeRange.ValidFrom))
	buf.Write(Uint64ToBytes(txData.TimeRange.ValidUntil))
	sig, err := s.Sign(buf.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign ChangePubKey tx data")
	}
	res := &Signature{
		PubKey:    s.GetPublicKey(),
		Signature: sig.HexString(),
	}
	return res, nil
}

func (s *ZkSigner) GetPublicKeyHash() string {
	return "sync:" + s.publicKeyHash.HexString()
}

func (s *ZkSigner) GetPublicKey() string {
	return s.publicKey.HexString()
}
