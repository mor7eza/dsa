package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE = 100_000

func main() {
	arr := make([]int, SIZE)
	for i := range SIZE {
		arr[i] = rand.Intn(SIZE)
	}

	startTime := time.Now()
	bubbleSort(&arr)
	fmt.Println("Execution time:", time.Since(startTime).Milliseconds(), "ms")
}

func bubbleSort(arr *[]int) {
	for i := range SIZE - 1 {
		for j := range SIZE - i - 1 {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
