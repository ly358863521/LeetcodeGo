package main

import "fmt"

const bucketNum int = 5

func insertSort(a []int) {
	// 插入排序是稳定排序，复杂度O(n^2),在几乎有序状态下效率很高
	for i := 1; i < len(a); i++ {
		key, j := a[i], i-1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

func bucketSort(a []int, w int) {
	w = w / bucketNum
	bucket := make([][]int, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucket[i] = make([]int, 0)
	}
	for _, v := range a {
		bucket[v/w] = append(bucket[v/w], v)
	}
	for i := 0; i < bucketNum; i++ {
		insertSort(bucket[i])
	}
	fmt.Println(bucket)
}
func main() {
	list := []int{4, 6, 71, 58, 3, 61}
	bucketSort(list, 100)
}
