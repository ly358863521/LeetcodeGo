package main

func findAnagrams(s string, p string) []int {
	var (
		res                = make([]int, 0)
		match, left, right int
		win                [26]int
		need               [26]int
	)
	for _, i := range p {
		need[i-'a']++
	}
	needs := 0
	for _, i := range need {
		if i != 0 {
			needs++
		}
	}
	for right < len(s) {
		c1 := s[right] - 'a'
		if need[c1] > 0 {
			win[c1]++
			if win[c1] == need[c1] {
				match++
			}
		}
		right++
		for match == needs {
			if right-left == len(p) {
				res = append(res, left)
			}
			c2 := s[left] - 'a'
			if need[c2] > 0 {
				win[c2]--
				if win[c2] < need[c2] {
					match--
				}
			}
			left++
		}
	}
	return res
}
func main() {
	findAnagrams("cbaedbaec", "ac")
}
