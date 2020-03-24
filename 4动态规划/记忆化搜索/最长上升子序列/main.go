package main

import (
	"fmt"
	"math/rand"
)

func max(i,j int)int{
	if i>j{
		return i
	}
	return j
}

// 动态规划 转移方程dp[i] = max(dp[i], dp[j] + 1) for j in [0, i)
func lengthOfLIS(nums []int) int {
	dp:= make([]int,len(nums))
	for i:= range dp{
		dp[i] = 1
	}
	for i:=0;i<len(nums);i++{
		for j:=0;j<i;j++{
			if dp[i]<dp[j]+1 && nums[i]>nums[j]{
				dp[i] = dp[j]+1
			}
		}
	}
	fmt.Println(dp)
	res := 0
	for _,v := range dp{
		if v>res{
			res = v
		}
	}
	return res
}

// 记忆化搜索
func searchLIS(nums []int)int{
	mem := make([]int,len(nums))
	for i:= range mem{
		mem[i] = -1
	}
	var dfs func(i int)int
	dfs = func(i int)int{
		if mem[i]!=-1{
			return mem[i]
		}
		ret := 1
		for j:=0;j<i;j++{
			if ll:=dfs(j)+1;nums[j]<nums[i] && ll>ret{
				ret = ll
			}
		}
		mem[i] = ret
		return ret
	}
	
	dfs(len(nums)-1)
	fmt.Println(mem)
	res := 0
	for _,v := range mem{
		if v>res{
			res = v
		}
	}
	return res
}

// 二分查找优化动态规划算法
// 将dp数组优化为升序单调数组tail
// 若大于则入栈
// 若小于则二分查找数组中首个大于该元素的值，将其替换
func binaryLIS(nums []int)int{
	var length int
    tail := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if length == 0 || tail[length-1] < nums[i] {
            tail[length] = nums[i]
            length++
        } else {
            x, y := 0, length-1
            for x < y {
                mid := x + (y - x) >> 1
                if tail[mid] < nums[i] {
                    x = mid + 1
                } else {
                    y = mid
                }
            }
            tail[x] = nums[i]
        }
    }
    return length
}


func main() {
	nums := make([]int,20)
	for i:=0;i<20;i++{
		nums[i] = rand.Intn(100)
	}
	fmt.Println(nums)
	fmt.Println(lengthOfLIS(nums))
	fmt.Println(searchLIS(nums))
	fmt.Println(binaryLIS(nums))

}