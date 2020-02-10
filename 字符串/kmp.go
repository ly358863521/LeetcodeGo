package main

func getnext1(p string) []int {
	next := make([]int, len(p), len(p))
	i, j := 0, 1
	for j < len(p) {
		if p[j] == p[i] {
			next[j] = i + 1
			i++
			j++
		} else {
			if i == 0 {
				next[j] = 0
				j++
				continue
			}
			i = next[i-1]
		}
	}
	return next
}
func getnext(pat string) []int {
	next := make([]int, len(pat), len(pat))
	next[0] = -1
	k, j := -1, 0
	for j < len(pat)-1 {
		if k == -1 || pat[j] == pat[k] {
			j++
			k++
			next[j] = k
			// if pat[j] != pat[k] {
			// 	next[j] = k
			// } else {
			// 	next[j] = next[k]
			// }
		} else {
			k = next[k]
		}
	}
	return next
}

func search(s string, p string, next []int) int {
	i, j, sLen, pLen := 0, 0, len(s), len(p)
	for i < sLen && j < pLen {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}
}
