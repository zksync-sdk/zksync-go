package zksync

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	zkscrypto "github.com/zksync-sdk/zksync-sdk-go"
	"testing"
)

var (
	err      error
	zkSigner *ZkSigner

	seedHex       = "c639503770c4b61f69185a5491b37caafa95f307ecceff5e1c335108308118b7e3a08276dc65a8794b1b2baa160e1cc8f16dce35706d437c13de26fb3b88303e"
	seed2Hex      = "e60ff3eecbdff618a387ac8d6e5dd335b95424f5a7abbf141a0f749c8e806d97fdbafb3a97bcbff8e6fa23cad70dfd6e31cfec442fb5c06cc766f39370089a98"
	pkHex         = "0422e258833908d5e3fdd12832545e0b121d9c4839cf8e719c8d1c318ddcb5a9"
	expPubKey     = "fd1a3106571e6707280c7358a9bc8961630a809fb70611467a12df2dc823769a"
	toPubKey      = "92401b49f36b6baf1a4d517681afdcfe671e7e0fbdbd709553e0d06a9979e01b"
	expPubKeyHash = "sync:d4ed534abc68bced79f5fc507dfcda2ca814f3dd"

	expMsgZkSignature = "bd808d99e92365d2ccbd35ffb04034638979015328febbfa4ccfec9482a5de16897599bab7844d4e6706dbfc47efd28237d391d04d49650dd54833c62ece8704"
)

func TestNewZkSignerFromSeed(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		seed, _ := hex.DecodeString(seedHex)
		zkSigner, err = NewZkSignerFromSeed(seed)
		require.NoError(t, err)
		assert.NotNil(t, zkSigner)
		assert.IsType(t, &ZkSigner{}, zkSigner)
		pubKey := zkSigner.GetPublicKey()
		assert.EqualValues(t, expPubKey, pubKey)
		pubKeyHash := zkSigner.GetPublicKeyHash()
		assert.EqualValues(t, expPubKeyHash, pubKeyHash)
	})

	t.Run("invalid seed", func(t *testing.T) {
		zkSigner, err := NewZkSignerFromSeed([]byte("invalid seed bytes"))
		require.Error(t, err)
		assert.Nil(t, zkSigner)
	})

}

func TestNewZkSignerFromRawPrivateKey(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pk, _ := hex.DecodeString(pkHex)
		zkSigner, err := NewZkSignerFromRawPrivateKey(pk)
		require.NoError(t, err)
		assert.NotNil(t, zkSigner)
		assert.IsType(t, &ZkSigner{}, zkSigner)
		pubKey := zkSigner.GetPublicKey()
		assert.EqualValues(t, expPubKey, pubKey)
		pubKeyHash := zkSigner.GetPublicKeyHash()
		assert.EqualValues(t, expPubKeyHash, pubKeyHash)
	})

	t.Run("invalid pk", func(t *testing.T) {
		zkSigner, err := NewZkSignerFromRawPrivateKey([]byte("invalid pk bytes"))
		require.Error(t, err)
		assert.Nil(t, zkSigner)
	})

}

func TestNewZkSignerFromEthSigner(t *testing.T) {
	expPubKey := "5c0ab5cac38c912773430be185789edc8b48dc2a1f8f13cbba8dbac1ca589e89"
	expPubKeyHash := "sync:e30e8868fca3e432358420cd22d4a6242ce9377f"

	t.Run("success", func(t *testing.T) {
		ethSigner, err := NewEthSignerFromMnemonic(mnemonic)
		zkSigner, err := NewZkSignerFromEthSigner(ethSigner, ChainIdMainnet)
		require.NoError(t, err)
		assert.NotNil(t, zkSigner)
		assert.IsType(t, &ZkSigner{}, zkSigner)
		pubKey := zkSigner.GetPublicKey()
		assert.EqualValues(t, expPubKey, pubKey)
		pubKeyHash := zkSigner.GetPublicKeyHash()
		assert.EqualValues(t, expPubKeyHash, pubKeyHash)
	})
}

func TestSign(t *testing.T) {
	signature, err := zkSigner.Sign(signMsg)
	require.NoError(t, err)
	assert.NotNil(t, signature)
	assert.IsType(t, &zkscrypto.Signature{}, signature)
	assert.EqualValues(t, expMsgZkSignature, signature.HexString())
}
