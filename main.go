package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}

	var slice []int
	for range size {
		slice = append(slice, int(rand.Int63()))
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	var max int

	if len(data) <= 1 {
		return 0
	}

	for _, v := range data {
		if v > max {
			max = v
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}

	var max int
	var sliceMax []int

	if len(data) <= 1 {
		return 0
	}

	for i := 0; i <= 7; i++ {
		chunk := len(data) / CHUNKS
		startIndx := chunk * i
		lastIndx := startIndx + chunk
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			for _, v := range data[startIndx:lastIndx] {
				if v > max {
					max = v
				}
				sliceMax = append(sliceMax, max)
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	for _, v := range sliceMax {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %s ms\n", max, elapsed.String())

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(start)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %s ms\n", max, elapsed.String())
}
