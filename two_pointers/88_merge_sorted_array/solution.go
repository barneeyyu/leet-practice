package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := n + m - 1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

func main() {
	testCases := []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			result: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			result: []int{1},
		},
		{
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			result: []int{1},
		},
	}

	for i, tc := range testCases {
		original := append([]int(nil), tc.nums1...) // make a copy for debugging
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		pass := true
		for j := range tc.result {
			if tc.nums1[j] != tc.result[j] {
				pass = false
				break
			}
		}
		if !pass {
			panic(fmt.Sprintf("Test %d failed ❌:\n  Input nums1: %v\n  Input nums2: %v\n  Expected: %v\n  Got: %v\n",
				i+1, original, tc.nums2, tc.result, tc.nums1))
		} else {
			fmt.Printf("Test %d passed ✅\n", i+1)
		}
	}
}
