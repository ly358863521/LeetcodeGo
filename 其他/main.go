package main

import "fmt"

func ptrify(v, p interface{}) {
	switch v.(type) {
	case []int:
		p = &v
	}
}

func main() {
	// iden := make(map[byte]*[]int)
	// for i := byte('A'); i < 'X'; i++ {
	// 	iden[i] = &[]int{}
	// }
	// c := iden['A']
	// *c = append(*c, 100)
	// fmt.Println(iden['A'])

	c := make(map[int]int)
	// c[1] = 1
	// p1 := &c
	// c[1] = 2
	// fmt.Println(c, *p1)
	var p *[]int
	ptrify(c, p)
	fmt.Println(p)
	// p := ptrify(c)
	// p[1] = 2
}
