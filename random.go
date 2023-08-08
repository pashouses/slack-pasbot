package main

import (
	"crypto/rand"
	"math/big"
)

func shuffle(names []string) {
	for i := len(names) - 1; i > 0; i-- {
		rd := randInt(i + 1)
		// Swap names
		names[i], names[rd] = names[rd], names[i]
	}
}

func randInt(max int) int {
	rd, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	rdInt := int(rd.Int64())
	return rdInt
}
