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
	Use:   "decrypt [E] [N]",
	Short: "Decrypt ciphertext",
	Long:  "Decrypt ciphertext c using public key {e, n}",
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

		d, err := srsa.CalculatePrivateKey(e, n)
		if err != nil {
			log.Fatal(err)
		}

		privateKey := srsa.PrivateKey{
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
		plaintext := srsa.DecryptBytes(&privateKey, ciphertext)
		fmt.Printf("%s\n", string(plaintext))
	},
}
