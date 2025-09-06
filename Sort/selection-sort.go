package main

func selectionSort(arr []int) (int, int) {
	var swaps, iterations int
	for i := range SIZE {
		minIndex := i
		for j := i + 1; j < SIZE; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			iterations++
		}
		if i != minIndex {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
			swaps++
		}
	}
	return swaps, iterations
}
