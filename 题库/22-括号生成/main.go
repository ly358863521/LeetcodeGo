package main

import "fmt"

func generateParenthesis(n int) []string {
	var a []string
	var f func(s string, left, right int)
	f = func(s string, left, right int) {
		if len(s) == 2*n {
			a = append(a, s)
			return
		}
		if left > right {
			f(s+")", left, right+1)
		}
		if left < n {
			f(s+"(", left+1, right)
		}
	}
	f("", 0, 0)
	return a
}
func main() {
	fmt.Println(len("a"))
	fmt.Println(generateParenthesis(3))
}
