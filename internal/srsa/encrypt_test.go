package srsa

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptByte(t *testing.T) {
	pub := PublicKey{
		E: big.NewInt(23),
		N: big.NewInt(187),
	}
	got := encryptByte(&pub, byte(88))
	want := byte(11)
	assert.Equal(t, got, want)
}
