package main

import "fmt"

// 多重背包是在01背包基础上，每个物品有一个k数量
// 简单的可以将其转换为k个物品，即变成01背包问题
// 可以用二进制分组优化
func main() {
	var n, m int // n物品个数，m背包大小
	fmt.Scanf("%d %d\n", &n, &m)
	var wlist, vlist []int
	var w, v, k int
	// k := make([]int, n)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d\n", &w, &v, &k)
		c := 1
		for k-c > 0 {
			k -= c
			wlist = append(wlist, w*c)
			vlist = append(vlist, v*c)
			c *= 2
		}
		wlist = append(wlist, w*k)
		vlist = append(vlist, v*k)
	}
	// 核心代码f(i) = max(f(i),f(i-wi)+vi)
	for i := 0; i < len(wlist); i++ {
		for j := m; j >= wlist[i]; j-- {
			if f[j-wlist[i]]+vlist[i] > f[j] {
				f[j] = f[j-wlist[i]] + vlist[i]
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
