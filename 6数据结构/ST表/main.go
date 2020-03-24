package main

import (
	"fmt"
	"math/rand"
)

const (
	maxn int = 2e5 + 5
	lgn  int = 25
)

var (
	a     []int   = make([]int, maxn)
	lg    []int   = make([]int, maxn)
	stmax [][]int = make([][]int, maxn)
	stmin [][]int = make([][]int, maxn)
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func new() {
	lg[0] = -1
	for i := 1; i < maxn; i++ {
		lg[i] = lg[i/2] + 1
		stmax[i] = make([]int, lgn)
		stmax[i][0] = a[i]
		stmin[i] = make([]int, lgn)
		stmin[i][0] = a[i]
	}
	for i := 1; i < lg[maxn-1]; i++ {
		for j := 1; j+(1<<i)-1 < maxn; j++ {
			stmax[j][i] = max(stmax[j][i-1], stmax[j+(1<<(i-1))][i-1])
			stmin[j][i] = min(stmin[j][i-1], stmin[j+(1<<(i-1))][i-1])
		}
	}
}
func qmax(l, r int) int {
	s := lg[r-l+1]
	return max(stmax[l][s], stmax[r-(1<<s)+1][s])
}
func qmin(l, r int) int {
	s := lg[r-l+1]
	return min(stmin[l][s], stmin[r-(1<<s)+1][s])
}
func main() {
	for i := 0; i < maxn; i++ {
		a[i] = rand.Intn(maxn)
	}
	new()
	fmt.Println(qmax(10, 100))
}
