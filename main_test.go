package main

import (
	"math/big"
	"testing"
)

func TestFindPrimes(t *testing.T) {
	n := big.NewInt(187)
	p, q, err := findPrimes(n)
	assertErrorNil(t, err)
	assertEqual(t, p, big.NewInt(11))
	assertEqual(t, q, big.NewInt(17))
}

func TestCalculatePrivateKey(t *testing.T) {
	e := big.NewInt(7)
	n := big.NewInt(187)
	want := big.NewInt(23)
	got, err := calculatePrivateKey(e, n)
	assertErrorNil(t, err)
	assertEqual(t, got, want)
}

func TestDecrypt(t *testing.T) {
	e := big.NewInt(23)
	n := big.NewInt(187)
	c := big.NewInt(11)
	got := decrypt(c, e, n)
	want := big.NewInt(88)
	assertEqual(t, got, want)
}

func assertErrorNil(t *testing.T, actual error) {
	t.Helper()
	if actual != nil {
		t.Errorf("got %v expected nil", actual)
	}
}

func assertEqual(t *testing.T, actual, expected *big.Int) {
	t.Helper()
	if actual.Cmp(expected) != 0 {
		t.Errorf("got %v; want %v", actual.String(), expected.String())
	}
}
