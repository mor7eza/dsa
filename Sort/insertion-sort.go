package main

func insertionSort(arr []int) (swaps, iterations int) {
	for i := 1; i < SIZE; i++ {
		insertIndex := i
		currentValue := arr[i]
		for j := i - 1; j > -1; j-- {
			iterations++
			if arr[j] > currentValue {
				arr[j+1] = arr[j]
				insertIndex = j
				swaps++
			} else {
				break
			}
		}
		arr[insertIndex] = currentValue
	}
	return
}
