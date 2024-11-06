package srsa

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPrimes(t *testing.T) {
	n := big.NewInt(187)
	p, q, err := findPrimes(n)
	assert.NoError(t, err)
	assert.Equal(t, p, big.NewInt(11))
	assert.Equal(t, q, big.NewInt(17))
}

func TestCalculatePrivateKey(t *testing.T) {
	e := big.NewInt(7)
	n := big.NewInt(187)
	want := big.NewInt(23)
	got, err := CalculatePrivateKey(e, n)
	assert.NoError(t, err)
	assert.Equal(t, got, want)
}

func TestDecryptByte(t *testing.T) {
	e := big.NewInt(23)
	n := big.NewInt(187)
	d, err := CalculatePrivateKey(e, n)
	assert.NoError(t, err)
	priv := PrivateKey{
		N: n,
		D: d,
	}
	got := decryptByte(&priv, byte(11))
	want := byte(88)
	assert.Equal(t, got, want)
}
