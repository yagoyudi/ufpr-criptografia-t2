package cmd

import "math/big"

const (
	maxValuePrimes = 1024 // Maior valor de p e q.
)

type PublicKey struct {
	N *big.Int
	E *big.Int
}

type PrivateKey struct {
	N *big.Int
	D *big.Int
}
