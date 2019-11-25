import "sort"
type Freq struct {
	key   int
	value int
}
type FreqList []Freq
func (p FreqList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p FreqList) Len() int           { return len(p) }
func (p FreqList) Less(i, j int) bool { return p[i].value > p[j].value }
func topKFrequent(nums []int, k int) []int {
    dict :=make(map[int]int)
    for _,i:= range nums{
        dict[i]++
    }
    p:=make(FreqList,len(dict))
    i:=0
    for k, v := range dict {
        p[i].key = k
        p[i].value = v
        i++
    }
    sort.Sort(p)
    res :=make([]int,k)
    for i:=0;i<k;i++{
        res[i] = p[i].key
    }
    return res
}
