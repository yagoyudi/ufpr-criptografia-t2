package main

import (
	"flag"
	"fmt"
	"log"
)

// Maior valor de p e q.
const MaxValuePrimes = 1024

// FindPrimes busca primos (p, q) tal que n=p*q.
func findPrimes(n int) (int, int, error) {
	for p := 2; p < MaxValuePrimes; p++ {
		if n%p == 0 {
			q := n / p
			return p, q, nil
		}
	}
	return 0, 0, fmt.Errorf("erro: findPrimes não encontrou (p, q) válido")
}

// modInverse calcula o inverso multiplicativo de |a| módulo |m| usando o
// algoritmo estendido de Euclides.
func modInverse(a, m int) (int, error) {
	m0 := m
	x0, x1 := 0, 1

	if m == 1 {
		return 0, fmt.Errorf("erro: não existe inverso multiplicativo")
	}

	for a > 1 {
		q := a / m
		m, a = a%m, m
		x0, x1 = x1-q*x0, x0
	}

	// Ajusta x1 para ser positivo
	if x1 < 0 {
		x1 += m0
	}

	return x1, nil
}

// CalculatePrivateKey retorna {d} com base em {e, n}.
func calculatePrivateKey(e, n int) (int, error) {
	if e < 1 {
		return 0, fmt.Errorf("erro: e < 1")
	}
	p, q, err := findPrimes(n)
	if err != nil {
		return 0, err
	}
	phi := (p - 1) * (q - 1)
	if e >= phi {
		return 0, fmt.Errorf("erro: e >= phi(n)")
	}
	d, err := modInverse(e, phi)
	if err != nil {
		return 0, err
	}
	return d, nil
}

// Decrypt descriptografa o texto cifrado c usando a chave privada {d, n}.
func decrypt(c, d, n int) int {
	plaintext := 1
	base := c % n

	for d > 0 {
		if d%2 == 1 {
			plaintext = (plaintext * base) % n
		}
		base = (base * base) % n
		d /= 2
	}

	return plaintext
}

func main() {
	e := flag.Int("e", 0, "Valor de e")
	n := flag.Int("n", 0, "Valor de n")
	c := flag.Int("c", 0, "Valor de c")
	flag.Parse()
	if *e == 0 || *n == 0 || *c == 0 {
		fmt.Println("Uso: go run main.go -e <valor> -n <valor>")
		return
	}

	d, err := calculatePrivateKey(*e, *n)
	if err != nil {
		log.Fatal(err)
	}

	plaintext := decrypt(*c, d, *n)
	fmt.Printf("%d\n", plaintext)
}
