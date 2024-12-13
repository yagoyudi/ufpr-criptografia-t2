package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

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
	Args:  cobra.ExactArgs(3),
	Run:   decryptMain,
}

func decryptMain(cmd *cobra.Command, args []string) {
	privateKey, err := initializePrivateKey(args[0], args[1])
	if err != nil {
		log.Fatal(err)
	}

	ciphertextContent, err := os.ReadFile(args[2])
	if err != nil {
		log.Fatal(err)
	}

	ciphertextEncoded := strings.TrimSpace(string(ciphertextContent))
	ciphertextDecoded, err := base64.StdEncoding.DecodeString(ciphertextEncoded)
	if err != nil {
		log.Fatal(err)
	}

	plaintext := srsa.Decrypt(privateKey, ciphertextDecoded)
	fmt.Printf("%s\n", string(plaintext))
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
