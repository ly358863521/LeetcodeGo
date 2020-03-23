package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	var sum func(s int) int
	sum = func(s int) int {
		var sum int
		for i := 0; i < n; i++ {
			sum += s
			s = (s + 1) >> 1
		}
		return sum
	}
	if n == 1 {
		fmt.Println(m)
		return
	}
	low, high := 1, m
	for low < high {
		mid := (low + high + 1) >> 1
		if s := sum(mid); s == m {
			fmt.Println(mid)
			return
		} else if s < m {
			low = mid
		} else {
			high = mid - 1
		}
	}
	fmt.Println(high)
	return
}
