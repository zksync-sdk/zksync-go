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
	return newZkSignerFromPrivateKey(privateKey)
}

func NewZkSignerFromRawPrivateKey(rawPk []byte) (*ZkSigner, error) {
	privateKey, err := zkscrypto.NewPrivateKeyRaw(rawPk)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create private key from raw bytes")
	}
	return newZkSignerFromPrivateKey(privateKey)
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

func newZkSignerFromPrivateKey(privateKey *zkscrypto.PrivateKey) (*ZkSigner, error) {
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

func (s *ZkSigner) SignTransfer(txData *Transfer) (*Signature, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0xff - 0x05)
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(txData.AccountId))
	buf.Write(txData.From[:])
	buf.Write(txData.To[:])
	buf.Write(Uint32ToBytes(txData.Token.Id))
	packedAmount, err := packAmount(txData.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack amount")
	}
	buf.Write(packedAmount)
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
		return nil, errors.Wrap(err, "failed to sign Transfer tx data")
	}
	res := &Signature{
		PubKey:    s.GetPublicKey(),
		Signature: sig.HexString(),
	}
	return res, nil
}

func (s *ZkSigner) SignWithdraw(txData *Withdraw) (*Signature, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0xff - 0x03)
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(txData.AccountId))
	buf.Write(txData.From[:])
	buf.Write(txData.To[:])
	buf.Write(Uint32ToBytes(txData.TokenId))
	amountBytes := txData.Amount.Bytes()
	buf.Write(make([]byte, 16-len(amountBytes))) // total amount slot is 16 bytes BE
	buf.Write(amountBytes)
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
		return nil, errors.Wrap(err, "failed to sign Withdraw tx data")
	}
	res := &Signature{
		PubKey:    s.GetPublicKey(),
		Signature: sig.HexString(),
	}
	return res, nil
}

func (s *ZkSigner) SignForcedExit(txData *ForcedExit) (*Signature, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0xff - 0x08)
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(txData.AccountId))
	buf.Write(txData.Target[:])
	buf.Write(Uint32ToBytes(txData.TokenId))
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
		return nil, errors.Wrap(err, "failed to sign ForcedExit tx data")
	}
	res := &Signature{
		PubKey:    s.GetPublicKey(),
		Signature: sig.HexString(),
	}
	return res, nil
}

func (s *ZkSigner) SignMintNFT(txData *MintNFT) (*Signature, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0xff - 0x09)
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(txData.CreatorId))
	buf.Write(txData.CreatorAddress[:])
	buf.Write(txData.ContentHash.Bytes())
	buf.Write(txData.Recipient[:])
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
	sig, err := s.Sign(buf.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign MintNFT tx data")
	}
	res := &Signature{
		PubKey:    s.GetPublicKey(),
		Signature: sig.HexString(),
	}
	return res, nil
}

func (s *ZkSigner) SignWithdrawNFT(txData *WithdrawNFT) (*Signature, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0xff - 0x0a)
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(txData.AccountId))
	buf.Write(txData.From[:])
	buf.Write(txData.To[:])
	buf.Write(Uint32ToBytes(txData.Token))
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
		return nil, errors.Wrap(err, "failed to sign WithdrawNFT tx data")
	}
	res := &Signature{
		PubKey:    s.GetPublicKey(),
		Signature: sig.HexString(),
	}
	return res, nil
}

func (s *ZkSigner) SignOrder(order *Order) (*Signature, error) {
	message, err := s.getOrderBytes(order)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order bytes")
	}
	sig, err := s.Sign(message)
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign Order data")
	}
	res := &Signature{
		PubKey:    s.GetPublicKey(),
		Signature: sig.HexString(),
	}
	return res, nil
}
func (s *ZkSigner) getOrderBytes(order *Order) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0x6f) // ASCII 'o' in hex for (o)rder
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(order.AccountId))
	buf.Write(order.RecipientAddress[:])
	buf.Write(Uint32ToBytes(order.Nonce))
	buf.Write(Uint32ToBytes(order.TokenSell))
	buf.Write(Uint32ToBytes(order.TokenBuy))
	if len(order.Ratio) != 2 {
		return nil, errors.New("invalid ratio")
	}
	buf.Write(BigIntToBytesBE(order.Ratio[0], 15))
	buf.Write(BigIntToBytesBE(order.Ratio[1], 15))
	packedAmount, err := packAmount(order.Amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack amount")
	}
	buf.Write(packedAmount)
	buf.Write(Uint64ToBytes(order.TimeRange.ValidFrom))
	buf.Write(Uint64ToBytes(order.TimeRange.ValidUntil))
	return buf.Bytes(), nil
}

func (s *ZkSigner) SignSwap(txData *Swap) (*Signature, error) {
	buf := bytes.Buffer{}
	buf.WriteByte(0xff - 0x0b)
	buf.WriteByte(TransactionVersion)
	buf.Write(Uint32ToBytes(txData.SubmitterId))
	buf.Write(txData.SubmitterAddress[:])
	buf.Write(Uint32ToBytes(txData.Nonce))
	if len(txData.Orders) != 2 {
		return nil, errors.New("invalid orders in Swap tx")
	}
	order1, err := s.getOrderBytes(txData.Orders[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order1 bytes")
	}
	order2, err := s.getOrderBytes(txData.Orders[1])
	if err != nil {
		return nil, errors.Wrap(err, "failed to get order2 bytes")
	}
	buf.Write(zkscrypto.ResqueHashOrders(append(order1, order2...)).GetBytes())
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
	if len(txData.Amounts) != 2 {
		return nil, errors.New("invalid amounts in Swap tx")
	}
	packedAmount, err := packAmount(txData.Amounts[0])
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack amount1")
	}
	buf.Write(packedAmount)
	packedAmount, err = packAmount(txData.Amounts[1])
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack amount2")
	}
	buf.Write(packedAmount)

	sig, err := s.Sign(buf.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "failed to sign Swap tx data")
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
