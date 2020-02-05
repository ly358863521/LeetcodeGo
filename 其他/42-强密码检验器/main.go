package main

import (
	"fmt"
	"unicode"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func btoi(b bool) int {
	if b {
		return 0
	}
	return 1
}
func strongPasswordChecker(s string) int {
	var has_digit, has_lower, has_upper bool
	cnt_mod := []int{0, 0, 0}
	var i, n_modify int
	S := []rune(s)
	var c rune
	for {
		if i == len(S) {
			break
		}
		c = S[i]
		if unicode.IsDigit(c) {
			has_digit = true
		} else if unicode.IsLower(c) {
			has_lower = true
		} else if unicode.IsUpper(c) {
			has_upper = true
		}
		j := i
		for S[j] == c {
			//fmt.Println(j)
			j++
			if j == len(S) {
				break
			}
		}
		fmt.Println(j, i)
		if j-i >= 3 {
			n_modify += (j - i) / 3
			cnt_mod[(j-i)%3]++
		}
		i = j
	}
	n_miss_ctype := btoi(has_digit) + btoi(has_lower) + btoi(has_upper)
	fmt.Println(n_miss_ctype, cnt_mod, n_modify)
	if i < 6 {
		return max(6-i, n_miss_ctype)
	} else if i <= 20 {
		return max(n_modify, n_miss_ctype)
	} else {
		n_remove := i - 20
		if n_remove < cnt_mod[0] {
			return max(n_modify-n_remove, n_miss_ctype) + i - 20
		}
		n_remove -= cnt_mod[0]
		n_modify -= cnt_mod[0]
		if n_remove < cnt_mod[1]*2 {
			return max(n_modify-n_remove/2, n_miss_ctype) + i - 20
		}
		n_remove -= cnt_mod[1] * 2
		n_modify -= cnt_mod[1]
		return max(n_modify-n_remove/3, n_miss_ctype) + i - 20

	}
}
func main() {
	fmt.Println(strongPasswordChecker("1111111111"))
}
