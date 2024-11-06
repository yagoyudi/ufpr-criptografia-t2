// Package srsa is the RSA package, but for small numbers of primes.
package srsa

import "math/big"

const (
	MaxValuePrimes = 1024
)

// If the pkg is for small numbers, then why use big.Int?
// It's because of useful functions like Exp and InverseMod.
type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	N *big.Int
	D *big.Int
}
