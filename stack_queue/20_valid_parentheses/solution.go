package main

import "fmt"

func isValid(s string) bool {
	var stack Stack
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack.Push(byte(r))
		case ')':
			top, ok := stack.Pop()
			if !ok || top != '(' {
				return false
			}
		case ']':
			top, ok := stack.Pop()
			if !ok || top != '[' {
				return false
			}
		case '}':
			top, ok := stack.Pop()
			if !ok || top != '{' {
				return false
			}
		}
	}
	return len(stack.items) == 0
}

type Stack struct {
	items []byte
}

func (s *Stack) Push(item byte) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (byte, bool) {
	if len(s.items) == 0 {
		return 0, false
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, true
}

func main() {
	tests := []struct {
		s        string
		expected bool
	}{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		},
		{
			s:        "(]",
			expected: false,
		},
		{
			s:        "([)]",
			expected: false,
		},
		{
			s:        "{[]}",
			expected: true,
		},
	}

	for i, test := range tests {
		result := isValid(test.s)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: s: %v | Expected: %v | Got: %v",
				i+1, test.s, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
