package main

import (
	"fmt"
	"reflect"
)

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for numbers[left]+numbers[right] != target {
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}
	return []int{left + 1, right + 1}
}

func main() {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			numbers:  []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			numbers:  []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
	}

	for i, test := range tests {
		result := twoSum(test.numbers, test.target)
		if !reflect.DeepEqual(result, test.expected) {
			panic(fmt.Sprintf("Example %d: numbers: %v, target: %v | Expected: %v | Got: %v",
				i+1, test.numbers, test.target, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
