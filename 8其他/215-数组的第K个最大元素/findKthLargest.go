package findkthlargest

//最小堆
func heapify(nums []int, i int) {
	n := len(nums)
	left, right := i*2+1, i*2+2 //左右节点
	min := i
	if left < n && nums[min] > nums[left] {
		min = left
	}
	if right < n && nums[min] > nums[right] {
		min = right
	}
	if min != i {
		nums[min], nums[i] = nums[i], nums[min]
		heapify(nums, min)
	}
}

func buildHeap(nums []int) {
	n := len(nums)
	for i := n / 2; i >= 0; i-- {
		heapify(nums, i)
	}
}

func findKthLargest(nums []int, k int) int {
	heap := nums[:k]
	buildHeap(heap)
	for i := k; i < len(nums); i++ {
		if nums[i] > heap[0] {
			heap = heap[1:]
			heap = append(heap, nums[i])
			buildHeap(heap)
		}
	}
	return heap[0]
}
