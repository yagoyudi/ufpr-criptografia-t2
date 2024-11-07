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
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [E] [N]",
	Short: "Encrypt plaintext",
	Long:  "Encrypt plaintext c using public key {e, n}",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		publicKey, err := initializePublicKey(args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}

		plaintext, err := readPlaintextFromStdin()
		if err != nil {
			log.Fatal(err)
		}

		ciphertext := srsa.EncryptBytes(publicKey, []byte(plaintext))
		fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
	},
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

func readPlaintextFromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return plaintext, nil
}
