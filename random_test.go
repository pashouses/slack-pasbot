package main

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 1; i < 100; i++ {
		rd := randInt(i)
		if rd < 0 || rd >= i {
			t.Errorf("Random does not meet criteria, max: %d, actual: %d\n", i, rd)
		}
	}
}

func TestShuffle(t *testing.T) {
	names := []string{"Abe", "Einstein", "Marie", "John", "Dick", "Haruka"}
	for i := 1; i < 20; i++ {
		shuffle(names)
	}
}
