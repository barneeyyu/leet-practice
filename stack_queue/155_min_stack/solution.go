package main

import "fmt"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
	if len(ms.minStack) == 0 || val <= ms.minStack[len(ms.minStack)-1] {
		ms.minStack = append(ms.minStack, val)
	}
}

func (ms *MinStack) Pop() {
	size := len(ms.stack)
	s := ms.stack[size-1]
	ms.stack = ms.stack[:size-1]

	if len(ms.minStack) > 0 && s == ms.minStack[len(ms.minStack)-1] {
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	type operation struct {
		op       string
		value    int
		expected int
	}

	tests := []struct {
		ops []operation
	}{
		{
			ops: []operation{
				{op: "push", value: -2},
				{op: "push", value: 0},
				{op: "push", value: -3},
				{op: "getMin", expected: -3},
				{op: "pop"},
				{op: "top", expected: 0},
				{op: "getMin", expected: -2},
			},
		},
	}

	for i, test := range tests {
		stack := Constructor()
		for j, op := range test.ops {
			switch op.op {
			case "push":
				stack.Push(op.value)
			case "pop":
				stack.Pop()
			case "top":
				result := stack.Top()
				if result != op.expected {
					panic(fmt.Sprintf("Test %d.%d: Top() Expected %d, got %d", i+1, j+1, op.expected, result))
				}
				fmt.Printf("Test %d.%d Top() ✅\n", i+1, j+1)
			case "getMin":
				result := stack.GetMin()
				if result != op.expected {
					panic(fmt.Sprintf("Test %d.%d: GetMin() Expected %d, got %d", i+1, j+1, op.expected, result))
				}
				fmt.Printf("Test %d.%d GetMin() ✅\n", i+1, j+1)
			}
		}
	}
}
