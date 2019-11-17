package main

import (
	"fmt"
	"math/bits"
)
func hammingDistance(x int, y int) int {
    return bits.OnesCount(uint(x) ^ uint(y))
}

func main(){
	var val uint
	val = 9
	//1001
	fmt.Println(hammingDistance(1,4))
	fmt.Println(bits.UintSize)
	fmt.Println("OnesCount",bits.OnesCount(val))
	fmt.Println("Len",bits.Len(val))
	fmt.Println("Reverse",bits.Reverse(bits.Reverse(uint(9))))
	fmt.Println("Reverse",bits.Reverse(uint(9)))
	fmt.Println("Reverse",bits.ReverseBytes(uint(9)))
}

// 2
// 64
// OnesCount 2
// Len 4
// Reverse 9
// Reverse 10376293541461622784
// Reverse 648518346341351424