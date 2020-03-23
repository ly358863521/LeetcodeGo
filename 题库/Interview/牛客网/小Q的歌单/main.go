package main

import (
	"fmt"
)

func combination(n, m int) int {
	nu, de := 1, 1
	for i := 1; i <= m; i++ {
		de *= i
		nu *= n
		n--
	}
	return nu / de
}
func main() {
	var k int
	var x, xn, y, yn int
	fmt.Scanf("%d\n", &k)
	fmt.Scanf("%d %d %d %d", &x, &xn, &y, &yn)
	// fmt.Println(k, x, xn, y, yn)
	var res int
	for i := 0; i <= xn; i++ {
		if i*x > k {
			break
		}
		for j := 0; j <= yn; j++ {
			if i*x+j*y > k {
				break
			}
			if i*x+j*y == k {
				res += combination(xn, i) * combination(yn, j)
				res %= 1000000007
			}
		}
	}
	fmt.Println(res)
}
