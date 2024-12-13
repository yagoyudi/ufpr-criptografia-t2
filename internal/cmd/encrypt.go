package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/criptografia-t2/internal/srsa"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "enc [E] [N]",
	Short: "Encrypt plaintext",
	Long:  "Encrypt plaintext c using public key {e, n}",
	Args:  cobra.ExactArgs(3),
	Run:   encryptMain,
}

func encryptMain(cmd *cobra.Command, args []string) {
	publicKey, err := initializePublicKey(args[0], args[1])
	if err != nil {
		log.Fatal(err)
	}

	plaintext, err := os.ReadFile(args[2])
	if err != nil {
		log.Fatal(err)
	}

	ciphertext := srsa.Encrypt(publicKey, []byte(plaintext))
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
}

func initializePublicKey(eStr, nStr string) (*srsa.PublicKey, error) {
	e, ok := new(big.Int).SetString(eStr, 10)
	if !ok {
		return nil, fmt.Errorf("error: invalid e")
	}
	n, ok := new(big.Int).SetString(nStr, 10)
	if !ok {
		return nil, fmt.Errorf("error: invalid n")
	}

	pub := srsa.PublicKey{
		N: n,
		E: e,
	}
	return &pub, nil
}
