### 双指针
```golang
func strStr(haystack string, needle string) int {
    m,n :=len(haystack),len(needle)
    for i:=0;i<m-n+1;i++{
        j:=0
        for ;j<n;j++{
            if haystack[i+j]!=needle[j]{
                break
            }
        }
        if j==n{
            return i
        }
    }
    return -1
}
```
### KMP
- 通过pat去构建next数组，通过dp数组去匹配string
- 匹配过程search()
  - txt指针i，pat指针j
  - 通过j = next[j]确定j的下个位置
  - if j==-1 || s[i]==p[j]
    - i++,j++
  - else
    - j = next[j]
  - if j==len(pat) return i-j
- 构建过程getnext()
  - 根据前后缀最大公共元素长度表，失配时，pat像右移动的位数 = 已匹配位数 - 上一位字符对应的最大长度值
  - next[0] = -1,k==-1
    - if k==-1||p[j]==p[k]
      - i++,k++
      - next[j]=k
    - else
      - k = next[k]
```golang
func getnext(pat string) []int {
	next := make([]int, len(pat), len(pat))
	next[0] = -1
	k, j := -1, 0
	for j < len(pat)-1 {
		if k == -1 || pat[j] == pat[k] {
			j++
			k++
			next[j] = k
			// if pat[j] != pat[k] {
			// 	next[j] = k
			// } else {
			// 	next[j] = next[k]
			// }
		} else {
			k = next[k]
		}
	}
	return next
}

func search(s string, p string, next []int) int {
	i, j, sLen, pLen := 0, 0, len(s), len(p)
	for i < sLen && j < pLen {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	} else {
		return -1
	}
}
```
### Sunday