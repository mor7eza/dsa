package main

func bubbleSort(arr []int) (swaps, iterations int) {
	for i := range SIZE - 1 {
		isSorted := true
		for j := range SIZE - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps++
				isSorted = false
			}
			iterations++
		}
		// return when array is sorted
		if isSorted {
			break
		}
	}
	return
}
