```golang
import (
    "strconv"
    "strings"
)
type sum struct{
    r int
    c byte
    strs []string
}

type Excel struct {
    H int
    W byte
    table [][]int
    sumStrs  []sum
}


func Constructor(H int, W byte) Excel {
    table := make([][]int,H)
    for i:=0;i<len(table);i++{
        table[i] = make([]int,W-'A'+1)
    }
    return Excel{H:H,W:W,table:table,sumStrs:nil}
}

func (this *Excel) Set(r int, c byte, v int)  {
    this.table[r-1][c-'A'] = v
    if this.sumStrs ==nil{
        return
    }else{
        for _,strs := range this.sumStrs{
            if strs.r == r &&strs.c ==c{
                return 
            }else{
                this.Sum(strs.r,strs.c,strs.strs)
            }
        }
        for _,strs := range this.sumStrs{
                this.Sum(strs.r,strs.c,strs.strs)
        }
    }

}
func (this *Excel) Get(r int, c byte) int {
    return this.table[r-1][c-'A']
}


func (this *Excel) Sum(r int, c byte, strs []string) int {
    if this.sumStrs ==nil{
        this.sumStrs = append(this.sumStrs,sum{r:r,c:c,strs:strs})
    }else{
        flag := true
        for i,s:= range this.sumStrs{
            if s.r ==r&&s.c==c{
                this.sumStrs[i].strs = strs
                flag = false
                break
            }
        }
        if flag{
            this.sumStrs = append(this.sumStrs,sum{r:r,c:c,strs:strs})
        }
    }
    var res int
    for _,s := range strs{
        if len(s)>3{
            area := strings.Split(s,":")
            a,_ :=strconv.Atoi(area[0][1:])
            b,_ :=strconv.Atoi(area[1][1:])
            for i:=a-1;i<b;i++{
                for j:= area[0][0];j<=area[1][0];j++{
                    res += this.table[i][j-'A']
                }
            }
        }else{
            a,_ := strconv.Atoi(s[1:])
            res += this.table[a-1][s[0]-'A']
        }
    }
    this.table[r-1][c-'A'] = res
    return res
}


/**
 * Your Excel object will be instantiated and called as such:
 * obj := Constructor(H, W);
 * obj.Set(r,c,v);
 * param_2 := obj.Get(r,c);
 * param_3 := obj.Sum(r,c,strs);
 */
```