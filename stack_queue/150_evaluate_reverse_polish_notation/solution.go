package main

import (
	"fmt"
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	s := Constructor()
	for _, v := range tokens {
		if v != "+" && v != "-" && v != "*" && v != "/" {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			s.Push(int(num))
		} else {
			secondValue := s.Pop()
			firstValue := s.Pop()
			var tempResult int
			switch v {
			case "+":
				tempResult = firstValue + secondValue
			case "-":
				tempResult = firstValue - secondValue
			case "*":
				tempResult = firstValue * secondValue
			case "/":
				tempResult = firstValue / secondValue
			}
			s.Push(tempResult)
		}
	}
	return s.Pop()
}

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
	}
}

func (ms *MinStack) Push(val int) {
	ms.stack = append(ms.stack, val)
}

func (ms *MinStack) Pop() int {
	size := len(ms.stack)
	topValue := ms.stack[size-1]
	ms.stack = ms.stack[:size-1]
	return topValue
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func main() {
	tests := []struct {
		tokens   []string
		expected int
	}{
		{[]string{"2", "1", "+", "3", "*"}, 9},
		{[]string{"4", "13", "5", "/", "+"}, 6},
		{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
	}

	for i, test := range tests {
		result := evalRPN(test.tokens)
		if result != test.expected {
			fmt.Printf("Test %d failed: input %v | Expected %d, got %d\n", i+1, test.tokens, test.expected, result)
		} else {
			fmt.Printf("Test %d passed ✅ Result: %d\n", i+1, result)
		}
	}
}
