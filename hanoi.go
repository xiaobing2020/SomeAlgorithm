package main

import (
	"fmt"
)

type Stack struct {
	datas []int
	size  int
}

func (stack *Stack) Push(num int) {
	stack.datas = append(stack.datas, num)
	stack.size++
}

func (stack *Stack) Pop() int {
	lastNum := stack.datas[stack.size-1]
	stack.datas = stack.datas[0 : stack.size-1]
	stack.size--
	return lastNum
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) Get(index int) int {
	return stack.datas[index]
}

func (stack *Stack) Show() {
	for i := 0; i < stack.size; i++ {
		fmt.Print(stack.datas[i], " ")
	}
	fmt.Println()
}

func NewStack(nums ...int) *Stack {
	if len(nums) == 0 {
		return &Stack{nil, 0}
	} else {
		return &Stack{nums, len(nums)}
	}
}

var temp = NewStack()
var dest = NewStack()

func hanoi(source *Stack, temp *Stack, dest *Stack, n int) {
	if source.Size() < n || n < 2 {
		return
	}

	if n == 2 {
		t := source.Pop()
		temp.Push(t)
		t = source.Pop()
		dest.Push(t)
		t = temp.Pop()
		dest.Push(t)
	} else {
		hanoi(source, dest, temp, n-1)
		t := source.Pop()
		dest.Push(t)
		hanoi(temp, source, dest, n-1)
	}
}

func main() {
	// 例子
	s := NewStack()
	for i := 10; i > 0; i-- {
		s.Push(i)
	}
	hanoi(s, temp, dest, 10)
}
