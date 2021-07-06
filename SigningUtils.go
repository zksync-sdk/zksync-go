package zksync

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/pkg/errors"
	"math/big"
)

const (
	FeeExponentBitWidth int64 = 5
	FeeMantissaBitWidth int64 = 11
)

func Uint32ToBytes(v uint32) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, v)
	return res
}

func Uint64ToBytes(v uint64) []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, v)
	return res
}

func pkhToBytes(pkh string) ([]byte, error) {
	if pkh[:5] != "sync:" {
		return nil, errors.New("PubKeyHash must start with 'sync:'")
	}
	res, err := hex.DecodeString(pkh[5:])
	if err != nil {
		return nil, err
	}
	if len(res) != 20 {
		return nil, errors.New("pkh must be 20 bytes long")
	}
	return res, nil
}

func packFee(fee *big.Int) ([]byte, error) {
	packedFee, err := integerToDecimalByteArray(fee, FeeExponentBitWidth, FeeMantissaBitWidth, 10)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack fee")
	}
	// check that unpacked fee still has same value
	if unpackedFee, err := decimalByteArrayToInteger(packedFee, FeeExponentBitWidth, FeeMantissaBitWidth, 10); err != nil {
		return nil, errors.Wrap(err, "failed to unpack fee")
	} else if unpackedFee.Cmp(fee) != 0 {
		return nil, errors.New("fee Amount is not packable")
	}
	return packedFee, nil
}

func integerToDecimalByteArray(value *big.Int, expBits, mantissaBits, expBase int64) ([]byte, error) {
	bigExpBase := big.NewInt(expBase)
	// maxExponent = expBase ^ ((2 ^ expBits) - 1)
	maxExpPow := big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(expBits), nil), big.NewInt(1))
	maxExponent := big.NewInt(0).Exp(bigExpBase, maxExpPow, nil)
	// maxMantissa = (2 ^ mantissaBits) - 1
	maxMantissa := big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(mantissaBits), nil), big.NewInt(1))
	// check for max possible value
	if value.Cmp(big.NewInt(0).Mul(maxMantissa, maxExponent)) > 0 {
		return nil, errors.New("Integer is too big")
	}
	exponent := uint64(0)
	mantissa := big.NewInt(0).Set(value)
	for mantissa.Cmp(maxMantissa) > 0 {
		mantissa.Div(mantissa, bigExpBase)
		exponent++
	}

	exponentData := uint64ToBitsLE(exponent, uint(expBits))
	mantissaData := uint64ToBitsLE(mantissa.Uint64(), uint(mantissaBits))
	combined := exponentData.Clone().Append(mantissaData)
	reversed := combined.Reverse()
	bytes, err := reversed.ToBytesBE()
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert bits to bytes BE")
	}
	return bytes, nil
}

func decimalByteArrayToInteger(value []byte, expBits, mantissaBits, expBase int64) (*big.Int, error) {
	if int64(len(value)*8) != expBits+mantissaBits {
		return nil, errors.New("Decimal unpacking, incorrect input length")
	}
	bits := NewBits(uint(expBits + mantissaBits))
	bits.FromBytesBE(value).Reverse()
	exponent := big.NewInt(0)
	expPow2 := big.NewInt(1)
	for i := uint(0); i < uint(expBits); i++ {
		if bits.GetBit(i) {
			exponent.Add(exponent, expPow2)
		}
		expPow2.Mul(expPow2, big.NewInt(2))
	}
	exponent.Exp(big.NewInt(expBase), exponent, nil)

	mantissa := big.NewInt(0)
	mantissaPow2 := big.NewInt(1)
	for i := uint(expBits); i < uint(expBits+mantissaBits); i++ {
		if bits.GetBit(i) {
			mantissa.Add(mantissa, mantissaPow2)
		}
		mantissaPow2.Mul(mantissaPow2, big.NewInt(2))
	}
	return exponent.Mul(exponent, mantissa), nil
}

func uint64ToBitsLE(v uint64, size uint) *Bits {
	res := NewBits(size)
	for i := uint(0); i < size; i++ {
		res.SetBit(i, v&1 == 1)
		v /= 2
	}
	return res
}

func getChangePubKeyData(txData *ChangePubKey) ([]byte, error) {
	buf := bytes.Buffer{}
	pkhBytes, err := pkhToBytes(txData.NewPkHash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get pkh bytes")
	}
	buf.Write(pkhBytes)
	buf.Write(Uint32ToBytes(txData.Nonce))
	buf.Write(Uint32ToBytes(txData.AccountId))
	buf.Write(txData.EthAuthData.getBytes())
	return buf.Bytes(), nil
}
