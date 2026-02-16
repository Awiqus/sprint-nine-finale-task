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
	if size < 1 {
		return nil
	}

	slice := make([]int, 0, size)
	for range size {
		slice = append(slice, rand.Int())
	}

	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	var max int

	if len(data) < 1 {
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
	var wg sync.WaitGroup
	var max int
	sliceMax := make([]int, CHUNKS)

	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	chunkLen := len(data) / CHUNKS
	for i := 0; i < CHUNKS; i++ {
		startIndx := chunkLen * i
		lastIndx := startIndx + chunkLen
		if i == CHUNKS-1 {
			lastIndx = len(data)
		}
		sliceChunk := data[startIndx:lastIndx]
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sliceMax[i] = maximum(sliceChunk)
		}(i)
	}

	wg.Wait()
	max = maximum(sliceMax)
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
