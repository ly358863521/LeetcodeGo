package main

import (
	"fmt"
)

func main() {
	input := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &input[i])
	}
	fmt.Println(input)
	ans := 10000
	var dfs func(dp []int, i, res int)
	dfs = func(dp []int, i, res int) {
		if i > 9 {
			if res < ans {
				ans = res
			}
			return
		}
		if dp[i] == 0 {
			dfs(dp, i+1, res)
		} else {
			if i <= 7 && (dp[i] >= 2 && dp[i+1] >= 2 && dp[i+2] >= 2) {
				dp[i], dp[i+1], dp[i+2] = dp[i]-2, dp[i+1]-2, dp[i+2]-2
				dfs(dp, i, res+1)
				dp[i], dp[i+1], dp[i+2] = dp[i]+2, dp[i+1]+2, dp[i+2]+2
			}
			if i <= 5 && (dp[i] >= 1 && dp[i+1] >= 1 && dp[i+2] >= 1 && dp[i+3] >= 1 && dp[i+4] >= 1) {
				dp[i], dp[i+1], dp[i+2], dp[i+3], dp[i+4] = dp[i]-1, dp[i+1]-1, dp[i+2]-1, dp[i+3]-1, dp[i+4]-1
				dfs(dp, i, res+1)
				dp[i], dp[i+1], dp[i+2], dp[i+3], dp[i+4] = dp[i]+1, dp[i+1]+1, dp[i+2]+1, dp[i+3]+1, dp[i+4]+1
			}
			if dp[i] >= 2 {
				dp[i] = dp[i] - 2
				dfs(dp, i, res+1)
				dp[i] = dp[i] + 2
			}
			if dp[i] >= 1 {
				dp[i]--
				dfs(dp, i, res+1)
				dp[i]++
			}
		}
	}
	dfs(input, 0, 0)
	fmt.Println(ans)
}
