package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, value := range nums {
		if m[value] {
			return true
		}
		m[value] = true
	}
	return false
}

func main() {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for i, test := range tests {
		result := containsDuplicate(test.input)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: Input: %v | Expected: %v | Got: %v", i+1, test.input, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
