package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var perm func(nums []int, l, r int)
	perm = func(nums []int, l, r int) {
		if l == r {
			tmpnums := make([]int, r)
			copy(tmpnums, nums)
			res = append(res, tmpnums)
		} else {
			for i := l; i < r; i++ {
				nums[i], nums[l] = nums[l], nums[i]
				// fmt.Println(2, nums)
				perm(nums, l+1, r)
				nums[i], nums[l] = nums[l], nums[i]
			}
		}
	}
	perm(nums, 0, len(nums))
	return res
}
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
