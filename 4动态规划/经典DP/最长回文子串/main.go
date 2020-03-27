package main

import "fmt"

// 利用动态规划
// 时间O(n^2)，空间O(n^2)
func dpPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	maxlen, start := 1, 0
	for r := 1; r < len(s); r++ {
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1]) {
				dp[l][r] = true
				if r-l+1 > maxlen {
					maxlen = r - l + 1
					start = l
				}
			}
		}
	}
	return s[start : start+maxlen]
}

// 中心扩散
// 时间O(n^2)，空间O(1)
func inPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var expand func(l, r int) int
	expand = func(l, r int) int {
		for l >= 0 && r < len(s) && (s[l] == s[r]) {
			l--
			r++
		}
		return r - l - 1
	}
	end, start := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expand(i, i)
		len2 := expand(i, i+1)
		if len1 < len2 {
			len1 = len2
		}
		if len1 > end-start {
			start = i - (len1-1)/2
			end = i + (len1 / 2)
		}
	}
	return s[start : end+1]
}

// Manacher 算法
// mx = id+p[id]
// j = 2*id-i是i关于id的对称点
func manacher(s string) (int, int) {
	min := func(i, j int) int {
		if i > j {
			return j
		}
		return i
	}
	ss := make([]byte, len(s)*2+3)
	ss[0] = '$'
	ss[1] = '#'
	for i, j := 0, 2; i < len(s); i++ {
		ss[j] = s[i]
		j++
		ss[j] = '#'
		j++
	}
	ss[len(ss)-1] = '&'
	fmt.Println(string(ss))
	l := len(ss)
	maxLen := -1
	id, mx := 0, 0
	maxID := 0
	p := make([]int, l)
	for i := 1; i < l-1; i++ {
		if i < mx {
			p[i] = min(p[2*id-i], mx-i)
		} else {
			p[i] = 1
		}
		for ss[i-p[i]] == ss[i+p[i]] { // 由于左侧有$右侧有&，不需要边界判断
			p[i]++
		}
		if mx < i+p[i] {
			id = i
			mx = i + p[i]
		}
		if p[i]-1 > maxLen {
			maxLen = p[i] - 1
			maxID = i
		}
	}
	return (maxID-1)/2 - maxLen/2, maxLen
}
func main() {
	s := "aabccba"
	fmt.Println(dpPalindrome(s))
	fmt.Println(inPalindrome(s))
	fmt.Println("-------------Manacher----------------")
	id, maxLen := manacher(s)
	fmt.Println("最长回文长度为:", maxLen)
	fmt.Println(s[id : id+maxLen])
}
