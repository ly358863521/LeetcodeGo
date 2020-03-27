package main

import (
	"fmt"
)

func main() {
	var buf = make([]int, 10)
	s := 15
	for i := len(buf) - 1; i >= 0; i-- {
		buf[i] = s & 1
		s = s >> 1
	}
	fmt.Println(buf)
}
