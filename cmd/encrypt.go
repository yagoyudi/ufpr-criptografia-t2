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
		e, ok := new(big.Int).SetString(args[0], 10)
		if !ok {
			log.Fatal("error: invalid e")
		}
		n, ok := new(big.Int).SetString(args[1], 10)
		if !ok {
			log.Fatal("error: invalid n")
		}

		publicKey := srsa.PublicKey{
			N: n,
			E: e,
		}

		reader := bufio.NewReader(os.Stdin)
		plaintext, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error: invalid plaintext")
		}

		ciphertext := srsa.EncryptBytes(&publicKey, []byte(plaintext))
		fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))
	},
}
