package main

import (
	"fmt"
	"math/rand"
	"time"
)

func down(h *[]int, i, n int) {
	old := *h
	for {
		left, right := i*2+1, i*2+2
		min := i
		if left < n && old[min] > old[left] {
			min = left
		}
		if right < n && old[min] > old[right] {
			min = right
		}
		if min != i {
			old[min], old[i] = old[i], old[min]
		} else {
			*h = old
			break
		}
	}
}

func up(h *[]int, j int) {
	old := *h
	for {
		i := (j - 1) / 2
		if i == j || old[i] <= old[j] {
			*h = old
			break
		}
		old[i], old[j] = old[j], old[i]
		j = i
	}
}

func buildheap(h *[]int) {
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func push(h *[]int, x int) {
	old := *h
	old = append(old, x)
	*h = old
	up(h, len(old)-1)
}

func pop(h *[]int) int {
	old := *h
	n := len(old) - 1
	old[0], old[n] = old[n], old[0]
	*h = old
	down(h, 0, n)
	x := old[n]
	old = old[:n]
	*h = old
	return x
}
func main() {
	var nums []int
	h := &nums
	var cur int
	start := time.Now().Unix()
	for i := 0; i < 100000000; i++ {
		ai := rand.Intn(100000000)
		// ai := 100000000
		if ai > cur {
			push(h, ai)
		}
		for len(*h) > cur {
			cur++
			for len(*h) > 0 && (*h)[0] == cur {
				pop(h)
			}
		}
	}
	fmt.Println(time.Now().Unix() - start)
}
