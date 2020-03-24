package main

import "fmt"

// 同01背包，区别只是每个物品可以重复取
func main() {
	var n, m int // n物品个数，m背包大小
	fmt.Scanf("%d %d\n", &n, &m)
	w := make([]int, n)
	v := make([]int, n)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &w[i], &v[i])
	}
	// 核心代码与01背包的区别只是第二层循环时，正向循环
	// 正序的时候即包括了当前容量下，i物品重复利用的情况
	for i := 0; i < n; i++ {
		for j := w[i]; j <= m; j++ {
			if f[j-w[i]]+v[i] > f[j] {
				f[j] = f[j-w[i]] + v[i]
			}
		}
	}
	fmt.Println(f[m])
}

// input
// 4 6
// 1 4
// 2 6
// 3 12
// 2 7
