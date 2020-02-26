package main

import "fmt"

func main() {
	var casenum int
	fmt.Scanf("%d", &casenum)
	// fmt.Println(casenum)
	for t := 0; t < casenum; t++ {
		var n int
		var score int
		var citation [1e5]int
		fmt.Scanf("%d", &n)
		// fmt.Println("n:", n)
		fmt.Printf("Case #%d:", t+1)
		for i := 0; i < n; i++ {
			var ai int
			fmt.Scanf("%d", &ai)
			for j := ai; j > score; j-- {
				citation[j]++
				if citation[j] >= j {
					score = j
					break
				}
			}
			fmt.Printf(" %d", score)
		}
		if t < casenum-1 {
			fmt.Printf("\n")
		}
	}
}
