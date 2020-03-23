![](https://i.loli.net/2019/11/26/grHlx5MASQdymUF.png)

```go
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
```

