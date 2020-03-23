package main

import "fmt"

func countSort(a []int, w int) {
	// 计数排序适用于值域很小的情况
	n := len(a)
	cnt := make([]int, w+1)
	b := make([]int, n)

	// 计算前缀和
	for i := 0; i < n; i++ {
		cnt[a[i]]++
	}
	for i := 1; i < w+1; i++ {
		cnt[i] += cnt[i-1]
	}

	// 后序遍历定位元素位置
	for i := n - 1; i >= 0; i-- {
		b[cnt[a[i]]-1] = a[i]
		cnt[a[i]]--
	}

	for i := 0; i < n; i++ {
		a[i] = b[i]
	}
}
func main() {
	list := []int{4, 6, 71, 58, 3, 61}
	countSort(list, 100)
	fmt.Println(list)
}
