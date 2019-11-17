package main

import (
    "fmt"
    "math/rand"
    "time"
)



func main() {
    //origin := rand.Perm(20) //伪随机数
    origin := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    fmt.Println("origin:", origin)
    randomQuickSort(origin, 0, len(origin))
    fmt.Println("quick sort:", origin)

}

func randomQuickSort(list []int, start, end int) {
    if end-start > 1 {
        // get the pivot
        mid := randomPartition(list, start, end)
        randomQuickSort(list, start, mid)
        randomQuickSort(list, mid+1, end)
    }
}

func randomPartition(list []int, begin, end int) int {
    // 生成真随机数
    i := randInt(begin, end)
    fmt.Println("random number:", i)
    // 下面这行是核心部分，随机选择主元， 如果没有此次交换，就是普通快排
    list[i], list[begin] = list[begin], list[i]
    return partition(list, begin, end)
}

func partition(list []int, begin, end int) (i int) {
    cValue := list[begin]
    i = begin
    for j := i + 1; j < end; j++ {
        if list[j] < cValue {
            i++
            list[j], list[i] = list[i], list[j]
        }
    }
    list[i], list[begin] = list[begin], list[i]
    return i
}

// 真随机数
func randInt(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return min + rand.Intn(max-min)
}