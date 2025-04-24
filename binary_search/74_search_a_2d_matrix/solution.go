package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	left, right := 0, row*col-1
	for left <= right {
		mid := (left + right) / 2
		mid_row := mid / col
		mid_col := mid % col
		if matrix[mid_row][mid_col] == target {
			return true
		} else if matrix[mid_row][mid_col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	tests := []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   3,
			expected: true,
		},
		{
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   13,
			expected: false,
		},
		{
			matrix:   [][]int{},
			target:   0,
			expected: false,
		},
	}

	for i, test := range tests {
		result := searchMatrix(test.matrix, test.target)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: matrix: %v, target: %v | Expected: %v | Got: %v",
				i+1, test.matrix, test.target, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
