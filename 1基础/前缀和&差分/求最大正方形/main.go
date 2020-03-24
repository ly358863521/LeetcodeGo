package main

import (
	"fmt"
	"math/rand"
	"time"
)

func min(i, j, k int) int {
	m := i
	if j < m {
		m = j
	}
	if k < m {
		m = k
	}
	return m
}
func maximalSquare(matrix [][]int) int {
	n := len(matrix)
	// 求二维的前缀和
	ans := 1
	b := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		b[i] = make([]int, n+1)
	}
	b[1][1] = matrix[0][0]
	for i := 2; i <= n; i++ {
		b[1][i] = b[1][i-1] + matrix[0][i-1]
		b[i][1] = b[i-1][1] + matrix[i-1][0]
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= n; j++ {
			b[i][j] = b[i-1][j] + b[i][j-1] + matrix[i-1][j-1] - b[i-1][j-1]
		}
	}
	// 打印
	for i := 0; i <= n; i++ {
		fmt.Println(b[i])
	}

	// 求最大面积

	for l := 2; l <= n; l++ {
		for i := l; i <= n; i++ {
			for j := l; j <= n; j++ {
				if b[i][j]-b[i-l][j]-b[i][j-l]+b[i-l][j-l] == l*l {
					if l*l > ans {
						ans = l * l
					}
				}
			}
		}
	}
	return ans
}

// 动态规划求最大面积
func dpmaxArea(matrix [][]int) int {
	b := make([][]int, len(matrix))
	copy(b, matrix)
	var maxarea int
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if b[i][j] == 1 {
				b[i][j] = min(b[i-1][j-1], b[i-1][j], b[i][j-1]) + 1
				if b[i][j] > maxarea {
					maxarea = b[i][j]
				}
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		fmt.Println(b[i])
	}
	return maxarea * maxarea
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([][]int, n)
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if rand.Float64() < 0.7 {
				a[i][j] = 1
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
	fmt.Println()
	area := maximalSquare(a)
	fmt.Println(area)
	area = dpmaxArea(a)
	fmt.Println(area)
}
