package main

import (
	"crypto/rand"
	"math/big"
)

func shuffle(names []string) {
	var tmp string
	for i := 0; i < len(names); i++ {
		rd := randInt(len(names))
		// Swap names
		tmp = names[i]
		names[i] = names[rd]
		names[rd] = tmp
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
