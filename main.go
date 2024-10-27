package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
)

// Maior valor de p e q.
const MaxValuePrimes = 1024

// FindPrimes busca primos (p, q) tal que n=p*q.
func findPrimes(n *big.Int) (*big.Int, *big.Int, error) {
	p := big.NewInt(2)
	limit := big.NewInt(MaxValuePrimes)

	for p.Cmp(limit) < 0 {
		q := new(big.Int).Div(n, p)             // q = n / p
		if new(big.Int).Mul(p, q).Cmp(n) == 0 { // Se p * q == n
			return new(big.Int).Set(p), new(big.Int).Set(q), nil
		}
		p.Add(p, big.NewInt(1)) // p++
	}
	return nil, nil, fmt.Errorf("erro: findPrimes não encontrou (p, q) válido")
}

// CalculatePrivateKey retorna {d} com base em {e, n}.
func calculatePrivateKey(e, n *big.Int) (*big.Int, error) {
	one := big.NewInt(1)
	if e.Cmp(one) <= 0 { // e < 1
		return nil, fmt.Errorf("erro: e < 1")
	}
	p, q, err := findPrimes(n)
	if err != nil {
		return nil, err
	}
	phi := new(big.Int).Mul(new(big.Int).Sub(p, one), new(big.Int).Sub(q, one)) // phi(n) = (p - 1) * (q - 1)
	if e.Cmp(phi) >= 0 {                                                        // e >= phi(n)
		return nil, fmt.Errorf("erro: e >= phi(n)")
	}
	d := new(big.Int).ModInverse(e, phi) // d = e^-1 (mod phi(n))
	if d == nil {
		return nil, fmt.Errorf("erro: não existe inverso multiplicativo")
	}
	return d, nil
}

// Decrypt descriptografa o texto cifrado c usando a chave privada {d, n}.
func decrypt(c, d, n *big.Int) *big.Int {
	plaintext := new(big.Int).Exp(c, d, n) // c^d mod n
	return plaintext
}

func main() {
	e := flag.String("e", "", "Valor de e")
	n := flag.String("n", "", "Valor de n")
	c := flag.String("c", "", "Valor de c")
	flag.Parse()
	if *e == "" || *n == "" || *c == "" {
		fmt.Println("Uso: go run main.go -e <valor> -n <valor>")
		return
	}

	// Converte de string para bit.Int
	eInt, _ := new(big.Int).SetString(*e, 10)
	nInt, _ := new(big.Int).SetString(*n, 10)
	cInt, _ := new(big.Int).SetString(*c, 10)

	d, err := calculatePrivateKey(eInt, nInt)
	if err != nil {
		log.Fatal(err)
	}

	plaintext := decrypt(cInt, d, nInt)
	fmt.Printf("%s\n", plaintext.String())
}
