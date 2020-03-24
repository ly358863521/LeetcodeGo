package main

import (
	"fmt"
)

func longestWPI(hours []int) int {
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
	}
	presum := make([]int, len(hours)+1, len(hours)+1)
	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + hours[i-1]
	}
	stack := []int{0}
	ans := 0
	for i := 1; i < len(presum); i++ {
		if presum[i] < presum[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	for i := len(presum) - 1; i > ans; i-- {
		for e := stack[len(stack)-1]; len(stack) != 0 && presum[e] < presum[i]; e = stack[len(stack)-1] {
			if i-e > ans {
				ans = i - e
			}
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
		}
	}
	return ans
}
func maxWidthRamp(A []int) int {
	stack := []int{0}
	for i := 1; i < len(A); i++ {
		if A[i] <= A[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	var max int
	for i := len(A) - 1; i >= 0; i-- {
		for len(stack) != 0 && A[i] >= A[stack[len(stack)-1]] {
			if i-stack[len(stack)-1] > max {
				max = i - stack[len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}
	return max
}

func main() {
	// fmt.Println(maxWidthRamp([]int{5, 4, 3, 6, 7}))
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9}))
}
