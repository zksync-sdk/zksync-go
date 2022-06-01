package zksync

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	ethSigner  *DefaultEthSigner
	mnemonic   = "kick swallow air vanish path pelican author bring group remove space retire retreat denial sphere"
	ethPkHex   = "0945012b971f943073f6e066581f513c8cd81660bb5a64306d5d092b8df9dd3f"
	expAddress = "0xFE64d0cF81848190C18ea5a7ff8d193ac807fd7E"

	mnemonic2 = "timber chronic resemble glide unlock click balance summer beauty cannon intact dwarf cross wrestle super"
	toAddress = "0x09084AAA8814F1781147F8d98798Dce2A86f96A6"

	signMsg          = []byte("sample message")
	expMsgSignature  = []byte{0x6d, 0xa3, 0x47, 0x59, 0xb6, 0xa0, 0xe, 0x2d, 0x64, 0x48, 0xad, 0x2, 0x71, 0xab, 0x16, 0x92, 0xb3, 0x62, 0x2b, 0x6f, 0x8d, 0x5f, 0xba, 0x97, 0xc2, 0x8f, 0xc4, 0xc6, 0x47, 0x7, 0xb7, 0x15, 0x47, 0xe4, 0xb, 0x6d, 0xbc, 0x4a, 0xed, 0x5b, 0x66, 0x98, 0xfc, 0xa2, 0xd3, 0x11, 0x5c, 0xc2, 0x57, 0xd6, 0x58, 0x1d, 0xaf, 0xb4, 0x63, 0x32, 0xcc, 0xfe, 0xbd, 0x24, 0x3f, 0x50, 0x8f, 0x50, 0x1c}
	signHash         = common.BytesToHash([]byte("sample hash"))
	expHashSignature = []byte{0x7c, 0x9a, 0x8, 0x3c, 0x94, 0x28, 0xe9, 0xd2, 0x13, 0x1f, 0x56, 0x20, 0xe9, 0xbf, 0xf1, 0x17, 0x76, 0xf9, 0xeb, 0xae, 0x66, 0x55, 0x77, 0x5f, 0x46, 0x9, 0x97, 0xcd, 0x84, 0x2a, 0x9e, 0x4, 0x55, 0x58, 0xc7, 0xd0, 0x34, 0x9a, 0xe6, 0x5c, 0xed, 0x3, 0x46, 0x5e, 0x43, 0x3b, 0x33, 0xf5, 0xd2, 0xaf, 0x4, 0x63, 0x3b, 0x9b, 0x74, 0x19, 0xf9, 0x2e, 0x78, 0x4f, 0x54, 0x66, 0x6f, 0x95, 0x0}
)

func TestNewEthSignerFromMnemonic(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ethSigner, err = NewEthSignerFromMnemonic(mnemonic)
		require.NoError(t, err)
		assert.NotNil(t, ethSigner)
		assert.IsType(t, &DefaultEthSigner{}, ethSigner)
		address := ethSigner.GetAddress()
		assert.NotNil(t, address)
		assert.IsType(t, common.Address{}, address)
		assert.EqualValues(t, expAddress, address.String())
	})

	t.Run("invalid mnemonic", func(t *testing.T) {
		ethSigner, err := NewEthSignerFromMnemonic("invalid mnemonic")
		require.Error(t, err)
		assert.Nil(t, ethSigner)
	})

}

func TestNewEthSignerFromRawPrivateKey(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		pk, err := hex.DecodeString(ethPkHex)
		require.NoError(t, err)
		ethSigner, err = NewEthSignerFromRawPrivateKey(pk)
		require.NoError(t, err)
		assert.NotNil(t, ethSigner)
		assert.IsType(t, &DefaultEthSigner{}, ethSigner)
		address := ethSigner.GetAddress()
		assert.NotNil(t, address)
		assert.IsType(t, common.Address{}, address)
		assert.EqualValues(t, expAddress, address.String())
	})

	t.Run("invalid raw pk", func(t *testing.T) {
		ethSigner, err := NewEthSignerFromRawPrivateKey([]byte{1, 2, 3})
		require.Error(t, err)
		assert.Nil(t, ethSigner)
	})

}

func TestSignMessage(t *testing.T) {
	signature, err := ethSigner.SignMessage(signMsg)
	require.NoError(t, err)
	assert.NotNil(t, signature)
	assert.EqualValues(t, expMsgSignature, signature)
}

func TestSignHash(t *testing.T) {
	t.Run("main positive flow", func(t *testing.T) {
		signature, err := ethSigner.SignHash(signHash.Bytes())
		require.NoError(t, err)
		assert.NotNil(t, signature)
		assert.EqualValues(t, expHashSignature, signature)
	})
	t.Run("invalid hash", func(t *testing.T) {
		signature, err := ethSigner.SignHash([]byte("invalid hash"))
		require.Error(t, err)
		assert.Nil(t, signature)
	})
}
