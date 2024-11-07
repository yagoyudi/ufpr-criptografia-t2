package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "t2",
	Short: "Trabalho 2 de Criptografia",
	Long:  "Dado um texto cifrado com o RSA e sabendo que p e q são números primos menores que 1024 (10 bits), faça um programa que encontre a chave privada d sendo que a chave pública {e , n} é conhecida.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
