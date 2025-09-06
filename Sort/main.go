package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

const SIZE = 10_000

func main() {
	random := make([]int, SIZE)
	bestCase := make([]int, SIZE)
	worstCase := make([]int, SIZE)
	for i := range SIZE {
		random[i] = rand.Intn(SIZE)
	}
	copy(bestCase, random)
	slices.Sort(bestCase)
	copy(worstCase, bestCase)
	slices.Reverse(worstCase)

	fmt.Printf("%-15s %-10s %-15s %-15s %-15s\n", "Algorithm", "Case", "Elasped(ms)", "Swaps", "Iterations")
	fmt.Println(strings.Repeat("-", 70))
	measureExecution(bubbleSort, slices.Clone(random), "Bubble", "Random")
	measureExecution(bubbleSort, slices.Clone(bestCase), "Bubble", "Best")
	measureExecution(bubbleSort, slices.Clone(worstCase), "Bubble", "Worst")

	measureExecution(selectionSort, slices.Clone(random), "Selection", "Random")
	measureExecution(selectionSort, slices.Clone(bestCase), "Selection", "Best")
	measureExecution(selectionSort, slices.Clone(worstCase), "Selection", "Worst")
}

func measureExecution(fn func([]int) (int, int), array []int, name, ratio string) {
	start := time.Now()
	swaps, iterations := fn(array)
	elasped := time.Since(start).Milliseconds()
	fmt.Printf("%-15s %-10s %-15d %-15d %-15d\n", name, ratio, elasped, swaps, iterations)
}
