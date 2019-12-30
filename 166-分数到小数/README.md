```go
import "strconv"
func fractionToDecimal(numerator int, denominator int) string {
    var res string
    if numerator*denominator<0{
        res +="-"
    }
    if numerator<0{
        numerator = -numerator
    }
    if denominator<0{
        denominator = -denominator
    }
    res += strconv.Itoa(numerator/denominator)
    numerator %= denominator
    if numerator ==0{
        return res
    }
    res +="."
    loop_int :=-1
    m :=make(map[int]int)
    for numerator!=0{
        if m[numerator]!=0{
            loop_int = m[numerator]
            break
        }else{
            m[numerator] = len(res)
        }
        numerator *=10
        res +=strconv.Itoa(numerator/denominator)
        numerator %=denominator
    }
    if numerator==0{
        return res
    }
    return res[:loop_int]+"("+res[loop_int:]+")"

}
```

