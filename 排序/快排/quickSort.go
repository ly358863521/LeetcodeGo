package sort

import "math/rand"

func partition(arr []int) (primeIdx int) {
	primeIdx = 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[len(arr)-1] {
			arr[i], arr[primeIdx] = arr[primeIdx], arr[i]
			primeIdx++
		}
	}
	arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
	return
}

func quickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := partition(arr)
		quickSort(arr[:primeIdx])
		quickSort(arr[primeIdx+1:])
	}
}

func randomQuickSort(arr []int) {
	if len(arr) > 1 {
		primeIdx := rand.Intn(len(arr))
		arr[primeIdx], arr[len(arr)-1] = arr[len(arr)-1], arr[primeIdx]
		primeIdx = partition(arr)
		randomQuickSort(arr[:primeIdx])
		randomQuickSort(arr[primeIdx+1:])
	}
}

func quickSortTail(arr []int) {
	for len(arr) > 1 {
		primeIdx := partition(arr)
		if primeIdx < len(arr)/2 {
			quickSortTail(arr[:primeIdx])
			arr = arr[primeIdx+1:]
		} else {
			quickSortTail(arr[primeIdx+1:])
			arr = arr[:primeIdx]
		}
	}
}
