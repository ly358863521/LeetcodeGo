package main

import "fmt"

// 有N件物品和一个容量为V的背包。第i件物品的重量是c[i]，价值是w[i]。求解将哪些物品装入背包可使这些物品的重量总和不超过背包容量，且价值总和最大。

// 状态转移方程f(i) = max(f(i),f(i-wi)+vi)
// 此处用一维滚动数组优化了二维数组
// f(i)表示处理到当前物品时，背包容量为i的最大价值
func main() {
	var n, m int // n物品个数，m背包大小
	fmt.Scanf("%d %d\n", &n, &m)
	w := make([]int, n)
	v := make([]int, n)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d\n", &w[i], &v[i])
	}
	// 核心代码f(i) = max(f(i),f(i-wi)+vi)
	for i := 0; i < n; i++ {
		for j := m; j >= w[i]; j-- {
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
