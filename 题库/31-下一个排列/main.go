package main
import "fmt"
func nextPermutation(nums []int)  {
	flag :=true
	for i:=len(nums)-1;i>0 ;i--{
		if nums[i] >nums[i-1]{
			for j:=i;j<len(nums);j++{
				if nums[j]<=nums[i-1]{
					nums[j-1],nums[i-1] = nums[i-1],nums[j-1]
					break
				}else if j==len(nums)-1{
					nums[j],nums[i-1] = nums[i-1],nums[j]
					fmt.Println("abc")
				}
			}
			reverse(nums,i)
			flag = false
			break
		}
	}
	if flag{
		reverse(nums,0)
	}
}

func reverse(nums []int,i int){
	for from, to := i, len(nums)-1; from < to; from, to = from+1, to-1 {
		nums[from], nums[to] = nums[to], nums[from]
	}
}
func main(){
	var a =[]int{1,2,3}
	nextPermutation(a)
	fmt.Println(a)
}