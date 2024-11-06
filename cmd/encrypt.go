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
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [E] [N]",
	Short: "Encrypt plaintext",
	Long:  "Encrypt plaintext c using public key {e, n}",
	Run:   encrypt,
}

func encrypt(cmd *cobra.Command, args []string) {
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

	publicKey := PublicKey{
		N: n,
		E: e,
	}

	reader := bufio.NewReader(os.Stdin)
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error: invalid plaintext")
	}

	ciphertext := encryptBytes(&publicKey, []byte(plaintext))
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
}

func encryptBytes(pub *PublicKey, plainBytes []byte) []byte {
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
