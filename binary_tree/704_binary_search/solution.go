package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := right + left/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 0, -1},
	}

	for i, test := range tests {
		result := search(test.nums, test.target)
		if result != test.expected {
			fmt.Printf("Test %d failed ❌: Expected %d, got %d\n", i+1, test.expected, result)
		} else {
			fmt.Printf("Test %d passed ✅\n", i+1)
		}
	}
}
