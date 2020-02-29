package main

import "fmt"

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
func main() {
	fmt.Println(hammingWeight(11))
}
