package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt [E] [N]",
	Short: "Decrypt ciphertext",
	Long:  "Decrypt ciphertext c using public key {e, n}",
	Run:   decrypt,
}

func decrypt(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		log.Fatal("error: requires 2 arguments")
	}

	e, ok := new(big.Int).SetString(args[0], 10)
	if !ok {
		log.Fatal("error: invalid e")
	}
	n, ok := new(big.Int).SetString(args[1], 10)
	if !ok {
		log.Fatal("error: invalid n")
	}

	d, err := calculatePrivateKey(e, n)
	if err != nil {
		log.Fatal(err)
	}

	privateKey := PrivateKey{
		N: n,
		D: d,
	}

	reader := bufio.NewReader(os.Stdin)
	c, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ciphertext, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		log.Fatal(err)
	}
	plaintext := decryptBytes(&privateKey, ciphertext)
	fmt.Printf("%s\n", string(plaintext))
}

// CalculatePrivateKey retorna {d} com base em {e, n}.
func calculatePrivateKey(e, n *big.Int) (*big.Int, error) {
	one := big.NewInt(1)
	if e.Cmp(one) <= 0 { // E < 1
		return nil, fmt.Errorf("erro: e < 1")
	}
	p, q, err := findPrimes(n)
	if err != nil {
		return nil, err
	}
	// phi(n) = (p - 1) * (q - 1)
	phi := new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one))
	if e.Cmp(phi) >= 0 { // e >= phi(n)
		return nil, fmt.Errorf("erro: e >= phi(n)")
	}
	d := new(big.Int).ModInverse(e, phi) // d = e^-1 (mod phi(n))
	if d == nil {
		return nil, fmt.Errorf("error: no modular multiplicative inverse")
	}
	return d, nil
}

// FindPrimes busca primos (p, q) tal que n=p*q.
func findPrimes(n *big.Int) (*big.Int, *big.Int, error) {
	p := big.NewInt(2)
	limit := big.NewInt(maxValuePrimes)

	for p.Cmp(limit) < 0 {
		q := new(big.Int).Div(n, p)             // q = n / p
		if new(big.Int).Mul(p, q).Cmp(n) == 0 { // Se p * q == n
			return new(big.Int).Set(p), new(big.Int).Set(q), nil
		}
		p.Add(p, big.NewInt(1)) // p++
	}
	return nil, nil, fmt.Errorf("error: couldn't find valid (p, q)")
}

func decryptByte(priv *PrivateKey, b byte) byte {
	c := big.NewInt(int64(b))
	m := new(big.Int).Exp(c, priv.D, priv.N)
	return byte(m.Int64())
}

func decryptBytes(priv *PrivateKey, cipheredBytes []byte) []byte {
	var plainBytes []byte
	for _, cipheredByte := range cipheredBytes {
		plainBytes = append(plainBytes, decryptByte(priv, cipheredByte))
	}
	return plainBytes
}
