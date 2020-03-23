package main

import (
	"fmt"
)

func selectionSort(a []int) {
	// 选择排序是不稳定排序，即有可能改变想等元素的顺序
	// 时间复杂度O(n^2)
	n := len(a)
	for i := 0; i < n; i++ {
		ith := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[ith] {
				ith = j
			}
		}
		a[i], a[ith] = a[ith], a[i]
	}
}
func main() {
	list := []int{4, 6, 71, 58, 3, 61}
	selectionSort(list)
	fmt.Println(list)
}
