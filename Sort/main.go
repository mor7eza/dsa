package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

const SIZE = 1_000_000

func main() {
	array := make([]int, SIZE)
	for i := range SIZE {
		array[i] = rand.Intn(SIZE)
	}

	fmt.Printf("%-15s %-15s %-15s %-15s\n", "Algorithm", "Elasped(ms)", "Swaps", "Iterations")
	fmt.Println(strings.Repeat("-", 70))

	// Bubble Sort
	measureExecution(bubbleSort, slices.Clone(array), "Bubble")

	// Selection Sort
	measureExecution(selectionSort, slices.Clone(array), "Selection")

	// Insertion Sort
	measureExecution(insertionSort, slices.Clone(array), "Insertion")

	// Quick Sort
	measureExecution(quickSort, slices.Clone(array), "Quick")

	// Counting Sort
	measureExecution(countingSort, slices.Clone(array), "Counting")
}

func measureExecution(fn any, array []int, name string) {
	var swaps, iterations int
	start := time.Now()
	switch f := fn.(type) {
	case func([]int) (int, int):
		swaps, iterations = f(array)
	case func([]int, int, int) (int, int):
		swaps, iterations = f(array, 0, SIZE-1)
	default:
		panic("unsupported function!")

	}

	elasped := time.Since(start).Milliseconds()
	fmt.Printf("%-15s %-15d %-15d %-15d\n", name, elasped, swaps, iterations)
}
