package zksync

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/pkg/errors"
	"math/big"
)

const (
	FeeExponentBitWidth int64 = 5
	FeeMantissaBitWidth int64 = 11
)

func Uint32ToBytes(accountId uint32) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, accountId)
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
	// WIP
	return integerToDecimalByteArray(fee, FeeExponentBitWidth, FeeMantissaBitWidth, 10)
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
	mantissa := value
	for mantissa.Cmp(maxMantissa) > 0 {
		mantissa.Div(mantissa, bigExpBase)
		exponent++
	}

	exponentData := uint64ToBitsLE(exponent, uint(expBits))
	mantissaData := uint64ToBitsLE(mantissa.Uint64(), uint(mantissaBits))
	combined := exponentData.Clone().Append(mantissaData)
	reversed := combined.Reverse()
	bytes, _ := reversed.ToBytesBE()
	return reverseByte(bytes), nil
}

func uint64ToBitsLE(v uint64, size uint) *Bits {
	res := NewBits(size)
	for i := uint(0); i < size; i++ {
		res.SetBit(i, v&1 == 1)
		v /= 2
	}
	return res
}

func reverseByte(v []byte) []byte {
	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
		v[i], v[j] = v[j], v[i]
	}
	return v
}
