package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	var tests = []int{-1, 0, 10, SIZE}

	for _, v := range tests {
		if v <= 0 {
			n := generateRandomElements(v)
			assert.Nil(t, n)
		} else {
			n := generateRandomElements(v)
			assert.NotNil(t, n)
		}
	}
}

func TestMaximum(t *testing.T) {
	var slice []int
	slicesLen := []int{0, 1, 2, 10, 100}
	for _, v := range slicesLen {
		if v < 1 {
			var slice []int
			max := maximum(slice)
			assert.Equal(t, 0, max)
		} else if v == 1 {
			slice := append(slice, v)
			max := maximum(slice)
			assert.Equal(t, 0, max)
		} else {
			var slice []int
			for range v {
				slice = append(slice, int(rand.Intn(v+1)))
			}
			max := maximum(slice)
			assert.InDelta(t, v, max, 10)
		}
	}
}
