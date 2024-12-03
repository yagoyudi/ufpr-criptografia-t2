package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/criptografia-t2/internal/srsa"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "dec [E] [N]",
	Short: "Decrypt ciphertext",
	Long:  "Decrypt ciphertext c using public key {e, n}",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		privateKey, err := initializePrivateKey(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		ciphertext, err := readCiphertextFromStdin()
		if err != nil {
			log.Fatal(err)
		}

		plaintext := srsa.Decrypt(privateKey, ciphertext)
		fmt.Printf("%s\n", string(plaintext))
	},
}

func initializePrivateKey(eStr, nStr string) (*srsa.PrivateKey, error) {
	e, ok := new(big.Int).SetString(eStr, 10)
	if !ok {
		return nil, fmt.Errorf("error: invalid e")
	}
	n, ok := new(big.Int).SetString(nStr, 10)
	if !ok {
		return nil, fmt.Errorf("error: invalid n")
	}

	d, err := srsa.CalculatePrivateKey(e, n)
	if err != nil {
		return nil, err
	}

	priv := srsa.PrivateKey{
		N: n,
		D: d,
	}
	return &priv, nil
}

func readCiphertextFromStdin() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	encodedCiphertext, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
