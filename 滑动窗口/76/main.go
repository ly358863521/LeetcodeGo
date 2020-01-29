package main

func minWindow(s string, t string) string {
	var (
		start, left, right, match int
		minlen                    = 100000
		win                       [256]int
		need                      [256]int
		needs                     int
	)
	for _, i := range t {
		if need[i] == 0 {
			needs++
		}
		need[i]++
	}
	for right < len(s) {
		c1 := s[right]
		if need[c1] > 0 {
			win[c1]++
			if win[c1] == need[c1] {
				match++
			}
		}
		right++
		for match == needs {
			if right-left < minlen {
				start = left
				minlen = right - left
			}
			c2 := s[left]
			if need[c2] > 0 {
				win[c2]--
				if win[c2] < need[c2] {
					match--
				}
			}
			left++
		}
	}
	if minlen == 100000 {
		return ""
	}
	return s[start : start+minlen]
}
func main() {

}
