package main

func countingSort(arr []int) (swaps, iterations int) {
	countingArray := make([]int, SIZE)
	for _, v := range arr {
		countingArray[v]++
		iterations++
	}

	counter := 0
	for i, v := range countingArray {
		for range v {
			arr[counter] = i
			counter++
			iterations++
			swaps++
		}
	}
	return
}
