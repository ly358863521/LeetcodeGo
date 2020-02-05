package main

func longestConsecutive(nums []int) int {
	M := make(map[int]bool)
	for _, i := range nums {
		if _, ok := M[i]; !ok {
			M[i] = true
		}
	}
	var res int
	// _, ok := M[100]
	// fmt.Println(M, M[100], ok)
	for _, x := range nums {
		if _, ok := M[x-1]; !ok {
			k := 0
			for {
				_, ok := M[x]
				if !ok {
					break
				}
				k++
				x++
			}
			if k > res {
				res = k
			}
		}
	}
	return res
}

func main() {
	longestConsecutive([]int{100, 4, 200, 1, 3, 2})
}
