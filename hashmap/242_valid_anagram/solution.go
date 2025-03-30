package main

import "fmt"

func isAnagram(s string, t string) bool {
	dic := make(map[rune]int)
	for _, ch := range t {
		dic[ch]++
	}
	for _, ch := range s {
		if dic[ch] != 0 {
			dic[ch]--
		} else {
			return false
		}
	}
	for _, v := range dic {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
		},
		{
			s:        "a",
			t:        "ab",
			expected: false,
		},
	}

	for i, test := range tests {
		result := isAnagram(test.s, test.t)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: s: %v, t: %v | Expected: %v | Got: %v", i+1, test.s, test.t, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
