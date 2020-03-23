package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	dp := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &dp[i])
	}
	// sort.Sort(sort.Reverse(sort.IntSlice(dp)))
	sort.Slice(dp, func(i, j int) bool {
		return dp[i] > dp[j]
	})
	var res int
	fmt.Println(dp)
	for i := 0; i < n; i += 2 {
		res += dp[i]
	}
	for i := 1; i < n; i += 2 {
		res -= dp[i]
	}
	fmt.Println(res)
}
