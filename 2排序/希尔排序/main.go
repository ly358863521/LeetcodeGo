package main

import "fmt"

func shellSort(a []int) {
	for d := len(a) / 2; d > 0; d /= 2 {
		// 确定增量d
		for i := d; i < len(a); i++ {
			// 插入排序
			for j := i - d; j >= 0; j -= d {
				if a[j] > a[j+d] {
					a[j], a[j+d] = a[j+d], a[j]
				}
			}
		}
	}
}
func main() {
	list := []int{4, 6, 71, 58, 3, 61}
	shellSort(list)
	fmt.Println(list)
}
