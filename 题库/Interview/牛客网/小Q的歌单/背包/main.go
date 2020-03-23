package main

import (
	"fmt"
)

const mod int = 1000000007

func main() {
	var k int
	var x, xn, y, yn int
	fmt.Scanf("%d\n", &k)
	fmt.Scanf("%d %d %d %d", &x, &xn, &y, &yn)
	p := make([]int, xn+yn+1)
	dp := make([][]int, xn+yn+1)
	for i := 0; i < xn+yn+1; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= xn; i++ {
		p[i] = x
	}
	for i := xn + 1; i <= xn+yn; i++ {
		p[i] = y
	}
	for i := 1; i <= xn+yn; i++ {
		for j := 0; j <= k; j++ {
			if j >= p[i] {
				dp[i][j] = (dp[i-1][j]%mod + dp[i-1][j-p[i]]%mod) % mod
			} else {
				dp[i][j] = dp[i-1][j] % mod
				// dp[i][0]=1
			}
		}
	}
	fmt.Println(dp[xn+yn][k] % mod)
}
