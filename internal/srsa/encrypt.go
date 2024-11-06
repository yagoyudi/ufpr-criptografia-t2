package srsa

import (
	"math/big"
)

func EncryptBytes(pub *PublicKey, plainBytes []byte) []byte {
	var cipheredBytes []byte
	for _, plainByte := range plainBytes {
		cipheredByte := encryptByte(pub, plainByte)
		cipheredBytes = append(cipheredBytes, cipheredByte)
	}
	return cipheredBytes
}

func encryptByte(pub *PublicKey, b byte) byte {
	m := big.NewInt(int64(b))
	c := new(big.Int).Exp(m, pub.E, pub.N)
	return byte(c.Int64())
}
