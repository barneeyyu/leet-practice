package main

import (
	"fmt"
	"reflect"
)

func moveZeroes(nums []int) {
	nonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for nonZeroIndex < len(nums) {
		nums[nonZeroIndex] = 0
		nonZeroIndex++
	}
}

func main() {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			input:    []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			input:    []int{1, 0, 0, 0, 0},
			expected: []int{1, 0, 0, 0, 0},
		},
		{
			input:    []int{0, 0, 0, 0, 1},
			expected: []int{1, 0, 0, 0, 0},
		},
		{
			input:    []int{1},
			expected: []int{1},
		},
		{
			input:    []int{0},
			expected: []int{0},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{2, 1},
			expected: []int{2, 1},
		},
		{
			input:    []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
	}

	for i, tc := range testCases {
		// 創建輸入數組的副本，因為moveZeroes會修改原數組
		input := make([]int, len(tc.input))
		copy(input, tc.input)

		moveZeroes(input)

		if reflect.DeepEqual(input, tc.expected) {
			fmt.Printf("Test %d passed ✅\n", i+1)
		} else {
			fmt.Printf("Test %d failed ❌:\n  Input: %v\n  Expected: %v\n  Got: %v\n", i+1, tc.input, tc.expected, input)
		}
	}
}
