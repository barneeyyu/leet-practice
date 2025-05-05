package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[rune]int)
	maxLen := 0
	left := 0

	for right, char := range s {
		if prevIndex, ok := charMap[char]; ok && prevIndex >= left {
			left = prevIndex + 1
		}
		charMap[char] = right
		maxLen = max(maxLen, right-left+1)
		fmt.Println(int(char), charMap[char], right, left, maxLen)
	}

	return maxLen
}

func main() {
	tests := []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		{
			s:        "",
			expected: 0,
		},
	}

	for i, test := range tests {
		result := lengthOfLongestSubstring(test.s)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: s: %v | Expected: %v | Got: %v",
				i+1, test.s, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
