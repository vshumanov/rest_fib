package main

import (
	"math/big"
)

func fibbonaci(n *big.Int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := big.NewInt(1); i.Cmp(n) <= 0; i.Add(i, big.NewInt(1)) {
		if i == big.NewInt(1) {
			return b
		}
		if i == big.NewInt(2) {
			return b
		}
		a.Add(a, b)
		a, b = b, a

	}
	return a
}
