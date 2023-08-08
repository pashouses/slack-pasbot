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

func TestRandIntDistribution(t *testing.T) {
	length := 10
	repeats := 1000000
	counts := make([]int, length)
	for i := 1; i < repeats; i++ {
		rd := randInt(length)
		if rd < 0 || rd >= length {
			t.Errorf("Random does not meet criteria, max: %d, actual: %d\n", i, rd)
		}
		counts[rd]++
	}
	// Print to console to see distribution
	for i := 0; i < length; i++ {
		percent := float64(counts[i]) / float64(repeats)
		if percent > 0.11 || percent < 0.09 {
			t.Errorf("Random distribution is not uniform, expected: 0.1, actual: %f\n", percent)
		}
	}
}

func TestShuffle(t *testing.T) {
	names := [10]string{"ada", "albert", "george", "john", "marie", "nikola", "rosalind", "srinivasa", "stephen", "tu"}
	for i := 1; i < 20; i++ {
		shuffle(names[:])
	}
}

func TestShuffleDistribution(t *testing.T) {
	repeats := 1000000
	names := [10]string{"ada", "albert", "george", "john", "marie", "nikola", "rosalind", "srinivasa", "stephen", "tu"}
	countMap := map[string]int{}
	for i := 0; i < repeats; i++ {
		n_names := names[:]
		shuffle(n_names)
		if countMap[n_names[0]] == 0 {
			countMap[n_names[0]] = 0
		}
		countMap[n_names[0]]++
	}
	for _, counts := range countMap {
		percent := float64(counts) / float64(repeats)
		if percent > 0.11 || percent < 0.09 {
			t.Errorf("Random distribution is not uniform, expected: 0.1, actual: %f\n", percent)
		}
	}
}
