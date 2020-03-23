```golang
type MinStack struct {
    stack []int
    min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{stack:[]int{},min:[]int{}}
}


func (this *MinStack) Push(x int)  {
    this.stack = append(this.stack,x)
    if len(this.min)==0 || this.min[len(this.min)-1]>=x{
        this.min = append(this.min,x)
    }
}


func (this *MinStack) Pop()  {
    if len(this.stack)==0{
        return
    }
    x := this.stack[len(this.stack)-1]
    this.stack = this.stack[:len(this.stack)-1]
    if x == this.min[len(this.min)-1]{
        this.min = this.min[:len(this.min)-1]
    }
}

func (this *MinStack) Top() int {
    if len(this.stack)==0{
        return 0
    }
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) Min() int {
    if len(this.min)==0{
        return 0
    }
    return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
```