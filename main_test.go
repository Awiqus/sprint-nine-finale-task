package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	var tests = []int{-1, 0, 1, 10, SIZE}

	for _, v := range tests {
		n := generateRandomElements(v)
		if v < 1 {
			assert.Nil(t, n)
		} else {
			assert.Len(t, n, v)
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"empty", []int{}, 0},
		{"single elem", []int{199}, 199},
		{"two elem", []int{0, 5}, 5},
		{"normal", []int{5, 6, 9, 37}, 37},
		{"big max", []int{9, 9, 9, 10_000_000, 0, 9, 9}, 10_000_000},
		{"increasing seq", []int{1, 2, 3, 4, 5, 6, 7, 199}, 199},
		{"decreasing seq", []int{88, 77, 66, 55, 44, 33, 22, 11}, 88},
	}

	for _, v := range tests {
		max := maximum(v.input)
		assert.Equal(t, v.expected, max)
	}
}
