package main

import (
	"fmt"
	"sort"
)

type m struct {
	time, grade int
}
type mlist []*m

func (l mlist) Len() int { return len(l) }
func (l mlist) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l mlist) Less(i, j int) bool {
	if l[i].time == l[j].time {
		return l[i].grade > l[j].grade
	}
	return l[i].time > l[j].time
}
func main() {
	var numM, numT int
	fmt.Scanf("%d %d\n", &numM, &numT)
	M := make(mlist, numM)
	for i := 0; i < numM; i++ {
		M[i] = &m{}
	}
	T := make(mlist, numT)
	for i := 0; i < numT; i++ {
		T[i] = &m{}
	}
	for i := 0; i < numM; i++ {
		fmt.Scanf("%d %d\n", &M[i].time, &M[i].grade)
	}
	for i := 0; i < numT; i++ {
		fmt.Scanf("%d %d\n", &T[i].time, &T[i].grade)
	}
	sort.Sort(M)
	sort.Sort(T)
	cnt := make([]int, 101)
	var num, ans, j, k int
	for i := 0; i < numT; i++ {
		for j < numM && M[j].time >= T[i].time {
			cnt[M[j].grade]++
			j++
		}
		for k = T[i].grade; k <= 100; k++ {
			if cnt[k] > 0 {
				num++
				cnt[k]--
				ans += 200*T[i].time + 3*T[i].grade
				break
			}
		}
	}
	fmt.Println(num, ans)

}
