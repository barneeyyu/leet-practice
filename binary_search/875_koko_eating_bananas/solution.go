package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, maxInPile(piles)
	result := right // the worst situation
	for left <= right {
		mid := (left + right) / 2
		hours := calculateHours(piles, mid)
		if hours > h {
			left = mid + 1
		} else {
			result = mid
			right = mid - 1
		}
	}
	return result
}

func calculateHours(piles []int, k int) int {
	var hours int
	for _, v := range piles {
		if v < k {
			hours++
		} else {
			hours += (v + k - 1) / k //這要注意
		}
	}
	return hours
}

func maxInPile(piles []int) int {
	max := piles[0]
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	tests := []struct {
		piles    []int
		h        int
		expected int
	}{
		{
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
	}

	for i, test := range tests {
		result := minEatingSpeed(test.piles, test.h)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: piles: %v, h: %v | Expected: %v | Got: %v",
				i+1, test.piles, test.h, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
