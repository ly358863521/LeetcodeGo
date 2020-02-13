package main

import "fmt"

func heapify(nums []int, i, n int) {
	min, left, right := i, 2*i+1, 2*i+2
	if left < n && nums[min] > nums[left] {
		min = left
	}
	if right < n && nums[min] > nums[right] {
		min = right
	}
	if min != i {
		nums[min], nums[i] = nums[i], nums[min]
		heapify(nums, min, n)
	}
}
func buildheap(nums []int, n int) {
	for i := n / 2; i >= 0; i-- {
		heapify(nums, i, n)
	}
}
func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) int {
	var res int
	heap := make([]int, maxOnes, maxOnes)
	for i := 0; i < sideLength; i++ {
		for j := 0; j < sideLength; j++ {
			v := ((width-j-1)/sideLength + 1) * ((height-i-1)/sideLength + 1)
			if v > heap[0] {
				heap = heap[1:]
				heap = append(heap, v)
				fmt.Println(i, j, v, heap)
				buildheap(heap, maxOnes)
				fmt.Println(i, j, v, heap)
			}
		}
	}
	for _, v := range heap {
		res += v
	}
	return res
}
func main() {
	fmt.Println(maximumNumberOfOnes(86, 32, 3, 4))
	// h := []int{308, 319, 319, 319}
	// buildheap(h, 4)
	// fmt.Println(h)
}
