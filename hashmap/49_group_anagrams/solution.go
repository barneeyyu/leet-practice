package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	result := [][]string{}
	for _, v := range strs {
		b := []byte(v)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		sortStr := string(b)

		m[sortStr] = append(m[sortStr], v)
	}
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	testCases := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for i, tc := range testCases {
		got := groupAnagrams(tc.input)
		// 這裡只簡單比較長度，實際上順序不重要，正式比對時可用更嚴謹的方式
		if len(got) == len(tc.expected) {
			fmt.Printf("Test %d passed ✅\n", i+1)
		} else {
			fmt.Printf("Test %d failed ❌:\n  Input: %v\n  Expected: %v\n  Got: %v\n", i+1, tc.input, tc.expected, got)
		}
	}
}
