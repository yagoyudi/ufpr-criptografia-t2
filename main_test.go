package main

import (
	"testing"
)

func TestFindPrimes(t *testing.T) {
	n := 187
	p, q, err := findPrimes(n)
	assertErrorNil(t, err)
	assertEqual(t, p, 11)
	assertEqual(t, q, 17)
}

func TestModInverse(t *testing.T) {
	a := 3
	m := 11
	expected := 4
	got, err := modInverse(a, m)
	assertErrorNil(t, err)
	assertEqual(t, got, expected)
}

func TestCalculatePrivateKey(t *testing.T) {
	e := 7
	n := 187
	want := 23
	got, err := calculatePrivateKey(e, n)
	assertErrorNil(t, err)
	assertEqual(t, got, want)
}

func assertErrorNil(t *testing.T, actual error) {
	t.Helper()
	if actual != nil {
		t.Errorf("got %v expected nil", actual)
	}
}

func assertEqual[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %v; want %v", actual, expected)
	}
}
