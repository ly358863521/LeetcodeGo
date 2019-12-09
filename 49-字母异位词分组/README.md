## python

### frozenset

```python
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        count = {}
        for stri in strs:
            countstri = {}
            for j in stri:
                countstri[j] = countstri.get(j,0)+1
            a = frozenset(countstri.items())
            if a in count:
                count[a].append(stri)
            else:
                count[a] = [stri]
        return count.values()
```

### defaultdict+sorted

```python
from collections import defaultdict
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        x = defaultdict(list)
        for _ in strs:
            x[''.join(sorted(_))].append(_)
        return x.values()
```

### go

```go
func groupAnagrams(strs []string) [][]string {
    res :=make([][]string,0)
    x :=make(map[string][]string)

    for _,stri :=range strs{
        strList :=strings.Split(stri,"")
        sort.Strings(strList)
        sorted := strings.Join(strList,"")
        if _,ok :=x[sorted];!ok{
            x[sorted] = []string{stri}
        }else{
            x[sorted] = append(x[sorted],stri)
        }
    }

    for _,v :=range x{
        res = append(res,v)
    }
    return res
}
```

```go
func groupAnagrams(strs []string) [][]string {
    result := [][]string{}
    mark := make(map[[26]byte]int)
    
    for _, str := range strs {
        key := [26]byte{}
        for _, r := range str {
            key[r-'a']++
        }

        i, ok := mark[key]
        if ok {
            result[i] = append(result[i], str)
        } else {
            result = append(result, []string{str})
            mark[key] = len(mark) //实现字典value值累加
        }
    }

    return result
}

```

