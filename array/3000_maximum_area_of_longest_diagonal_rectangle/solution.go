package main

import (
	"fmt"
	"reflect"
)

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDiag := 0
	maxArea := 0
	for _, v := range dimensions {
		tempDiag := v[0]*v[0] + v[1]*v[1]
		tempArea := v[0] * v[1]
		if tempDiag > maxDiag || (tempDiag == maxDiag) && tempArea > maxArea {
			maxDiag = tempDiag
			maxArea = tempArea
		}
	}
	return maxArea
}

func main() {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input:    [][]int{{9, 3}, {8, 6}},
			expected: 48,
		},
		{
			input:    [][]int{{3, 4}, {4, 3}},
			expected: 12,
		},
		{
			input:    [][]int{{4, 7}, {8, 9}, {5, 3}, {6, 10}},
			expected: 72, // [8,9] 對角線最長
		},
		{
			input:    [][]int{{25, 60}, {39, 52}, {16, 63}, {33, 56}},
			expected: 2028, // [39,52] 與其他對角線一樣長，但面積最大
		},
	}

	for i, tc := range testCases {
		got := areaOfMaxDiagonal(tc.input)

		if reflect.DeepEqual(got, tc.expected) {
			fmt.Printf("Test %d passed ✅\n", i+1)
		} else {
			fmt.Printf("Test %d failed ❌:\n  Input: %v\n  Expected: %v\n  Got: %v\n", i+1, tc.input, tc.expected, got)
		}
	}
}
