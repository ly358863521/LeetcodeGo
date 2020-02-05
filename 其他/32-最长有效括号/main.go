func longestValidParentheses(s string) int {
    var left,right,maxans int
    for _,v :=range s{
        if v =='('{
            left++
        }else{
            right++
        }
        if left==right&&maxans<2*left{
            maxans = 2*left
        }else if right >left{
            left = 0
            right =0
        }
    }
    left = 0
    right =0
    for i:=len(s)-1;i>=0;i--{
        if s[i] ==')'{
            right++
        }else{
            left++
        }
        if left ==right&&maxans<2*left{
            maxans = 2*left
        }else if left>right{
            left = 0
            right = 0
        }
    }
    return maxans
}