package main

import (
	"bytes"
	"fmt"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1, n2 := len(num1), len(num2)
	result := make([]int, n1+n2)
	var sum func(value, index int)
	sum = func(value, index int) {
		tmpindex := n1 + n2 - 1 - index
		tmp := result[tmpindex] + value
		if tmp > 9 {
			result[tmpindex] = tmp % 10
			sum(tmp/10, index+1)
		} else {
			result[tmpindex] = tmp
		}
	}
	for i := 0; i < n1; i++ {
		a := int(num1[n1-1-i] - '0')
		for j := 0; j < n2; j++ {
			b := int(num2[n2-1-j] - '0')
			c := a * b
			sum(c, i+j)
		}
	}
	if result[0] == 0 {
		result = result[1:]
	}
	fmt.Println(byte(result[0] + '0'))
	res := bytes.Buffer{}
	for i := 0; i < len(result); i++ {
		res.WriteByte(byte(result[i] - '0'))
	}
	return res.String()
}

func isSlash(r rune) bool {
	return (r == '/' || r == '\\')
}

func main() {
	// multiply("2", "3")
	// fmt.Println(byte(6) + '0')
	// s := "12/23//2\\3///23"
	// res := strings.Split(s, "/")
	// for _, i := range res {
	// 	fmt.Println(i)
	// }
	// fmt.Println(strings.SplitN(s, "/", -1))
	// fmt.Println(strings.Fields(s))
	// fmt.Println(strings.FieldsFunc(s, isSlash))
	fmt.Println(getnext("abcdabcdabcdabcd"))
	fmt.Println(getnext1("abcdabcdabcdabcd"))
}
