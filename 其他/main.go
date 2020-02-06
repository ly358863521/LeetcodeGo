package main

import "fmt"

func main() {
	iden := make(map[byte]*[]int)
	for i := byte('A'); i < 'X'; i++ {
		iden[i] = &[]int{}
	}
	c := iden['A']
	*c = append(*c, 100)
	fmt.Println(iden['A'])
}
