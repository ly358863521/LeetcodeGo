func randInt(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return min + rand.Intn(max-min)
}

func randomPartition(list []int, begin, end int) int {
    // 生成真随机数
    i := randInt(begin, end)
    // 下面这行是核心部分，随机选择主元， 如果没有此次交换，就是普通快排
    list[i], list[begin] = list[begin], list[i]
    return partition(list, begin, end)
}
func partition(list []int, begin, end int) (i int) {
    cValue := list[begin]
    i = begin
    for j := i + 1; j < end; j++ {
        if list[j] > cValue {
            i++
            list[j], list[i] = list[i], list[j]
        }
    }
    list[i], list[begin] = list[begin], list[i]
    return i
}
func find(List []int,start,end,k int)int{
    if start ==end{
        return List[start]
    }
    mid := randomPartition(List, start, end)
    if mid ==k-1{
        return List[mid]
    }
    if mid <k-1{
        return find(List,mid+1,end,k)
    }else{
        return find(List,start,mid,k)
    }
}
func findKthLargest(nums []int, k int) int {
    if nums==nil || len(nums)<k{
        return 0
    }
    return find(nums,0,len(nums),k)
}