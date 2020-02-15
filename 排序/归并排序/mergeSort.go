package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(nums []int) []int {
	tmp := make([]int, len(nums), len(nums))
	var sort func(L, R int)
	var merge func(L, M, R int)
	sort = func(L, R int) {
		if L < R {
			M := (L + R) / 2
			sort(L, M)
			sort(M+1, R)
			merge(L, M, R)
		}
	}
	merge = func(L, M, R int) {
		i, j, t := L, M+1, L
		for i <= M && j <= R {
			if nums[i] <= nums[j] {
				tmp[t] = nums[i]
				i++
			} else {
				tmp[t] = nums[j]
				j++
			}
			t++
		}
		for i <= M {
			tmp[t] = nums[i]
			i++
			t++
		}
		for j <= R {
			tmp[t] = nums[j]
			j++
			t++
		}
		for k := L; k <= R; k++ {
			nums[k] = tmp[k]
		}
	}
	sort(0, len(nums)-1)
	return nums
}

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100)
	nums := make([]int, n, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(100)
	}
	fmt.Println(nums)
	fmt.Println(mergeSort(nums))
}
