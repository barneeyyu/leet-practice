package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	var filteredString []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			filteredString = append(filteredString, unicode.ToLower(r))
		}
	}
	left := 0
	right := len(filteredString) - 1

	if right <= 0 {
		return true
	}

	for left < right {
		if filteredString[left] != filteredString[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	tests := []struct {
		s        string
		expected bool
	}{
		{
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			s:        "race a car",
			expected: false,
		},
		{
			s:        " ",
			expected: true,
		},
	}

	for i, test := range tests {
		result := isPalindrome(test.s)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: s: %v | Expected: %v | Got: %v", i+1, test.s, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
