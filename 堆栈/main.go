package main

import (
	"errors"
	"fmt"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack *Stack) Push(value interface{}) {
	*stack = append(*stack, value)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}
func main() {
	var stack Stack
	// for i := 0; i < 10; i++ {
	// 	s.Push(i)
	// }
	// fmt.Println(s)
	// for i := 0; i < 10; i++ {
	// 	t, _ := s.Pop()
	// 	fmt.Println(t.(int) - i)
	// }
	s := "(()"
	res := 0
	stack.Push(-1)
	for i, v := range s {
		if v == '(' {
			stack.Push(i)
		} else {
			stack.Pop()
			if stack.IsEmpty() {
				stack.Push(i)
			} else {
				t, _ := stack.Top()
				if i-t.(int) > res {
					res = i - t.(int)
				}
			}
		}
	}
	fmt.Println(res)
}
