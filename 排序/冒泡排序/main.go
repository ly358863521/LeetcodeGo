package main

import (
	"fmt"
)

func bubbleSort(a []int) {
	// 冒泡排序是稳定排序
	// 有序情况复杂度O(n)
	// 最坏情况和平均情况都是O(n^2)
	for flag := true; flag; {
		flag = false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				flag = true
			}
		}
	}
}
func main() {
	list := []int{4, 6, 71, 58, 3, 61}
	bubbleSort(list)
	fmt.Println(list)
}
