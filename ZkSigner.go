package zksync

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync-sdk-go"
)

const MESSAGE = "Access zkSync account.\n\nOnly sign this message for a trusted client!"

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
	// TODO
	return nil, errors.New("can't implement due to private zkscrypto.PrivateKey data field")
}

func NewZkSignerFromEthSigner(es EthSigner, cid ChainId) (*ZkSigner, error) {
	signMsg := MESSAGE
	if cid != ChainIdMainnet {
		signMsg = fmt.Sprintf("%s\nChain ID: %d.", MESSAGE, cid)
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

func (s *ZkSigner) SignChangePubKey(message []byte) (*zkscrypto.Signature, error) {
	return nil, nil
}

func (s *ZkSigner) GetPublicKeyHash() string {
	return "sync:" + s.publicKeyHash.HexString()
}
