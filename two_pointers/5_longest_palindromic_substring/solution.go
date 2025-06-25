package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := calculatePalindromeLong(s, i, i)
		l2, r2 := calculatePalindromeLong(s, i, i+1)

		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

func calculatePalindromeLong(s string, l, r int) (int, int) {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return l + 1, r - 1
}

func main() {
	testCases := []struct {
		input  string
		result string
	}{
		{
			input:  "babad",
			result: "bab", // or "aba" are both acceptable
		},
		{
			input:  "cbbd",
			result: "bb",
		},
		{
			input:  "a",
			result: "a",
		},
		{
			input:  "ac",
			result: "a", // or "c"
		},
		{
			input:  "forgeeksskeegfor",
			result: "geeksskeeg",
		},
		{
			input:  "abccba",
			result: "abccba",
		},
	}

	for i, tc := range testCases {
		got := longestPalindrome(tc.input)
		if got != tc.result {
			fmt.Printf("Test %d failed ❌:\n  Input: %s\n  Expected: %s\n  Got: %s\n", i+1, tc.input, tc.result, got)
		} else {
			fmt.Printf("Test %d passed ✅\n", i+1)
		}
	}
}
