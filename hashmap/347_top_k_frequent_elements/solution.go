package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}

	type pair struct {
		num  int
		freq int
	}
	pairs := []pair{}
	for num, freq := range freqMap {
		pairs = append(pairs, pair{num, freq})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].freq > pairs[j].freq
	})

	if len(pairs) < k {
		return []int{-1}
	}
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, pairs[i].num)
	}

	return result
}
func main() {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
	}

	for i, test := range tests {
		result := topKFrequent(test.nums, test.k)
		if !equalSlices(result, test.expected) {
			panic(fmt.Sprintf("Example %d: nums: %v, k: %v | Expected: %v | Got: %v", i+1, test.nums, test.k, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}

// equalSlices compares two int slices for equality
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
