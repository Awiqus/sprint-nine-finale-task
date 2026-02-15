package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	var tests = []int{-1, 0, 1, 10, SIZE}

	for i, v := range tests {
		if v <= 1 {
			n := generateRandomElements(v)
			assert.Nil(t, n)
		} else {
			n := generateRandomElements(v)
			assert.Equal(t, tests[i], len(n))
		}
	}
}

func TestMaximum(t *testing.T) {
	testsRand := []int{
		0,
		1,
		2,
		10,
		100,
		SIZE,
	}

	testsDeter := map[int][]int{
		1:       {0, 1},
		37:      {5, 6, 9, 37},
		1900880: {0, 56289, 1900880, 67, 27, 21, 69},
	}

	for _, v := range testsRand {
		if v <= 1 {
			slice := make([]int, v)
			max := maximum(slice)
			assert.Equal(t, 0, max)
		} else {
			var slice []int
			for range v {
				slice = append(slice, rand.Intn(v))
			}
			max := maximum(slice)

			sort.Ints(slice)
			x := slice[len(slice)-1]

			assert.Equal(t, x, max)
		}
	}

	for i, v := range testsDeter {
		max := maximum(v)
		assert.Equal(t, i, max)
	}
}
