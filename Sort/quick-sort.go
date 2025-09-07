package main

func quickSort(arr []int, low, high int) (swaps, iterations int) {
	if low >= high {
		return
	}

	pivotValue := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		iterations++
		if arr[j] <= pivotValue {
			i++
			arr[i], arr[j] = arr[j], arr[i]
			swaps++
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	s1, i1 := quickSort(arr, low, i-1)
	s2, i2 := quickSort(arr, i+1, high)
	swaps += s1 + s2
	iterations += i1 + i2
	return
}
