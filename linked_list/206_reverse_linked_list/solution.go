package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

// Helper function to create linked list from slice
func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, v := range values[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// Helper function to convert linked list to slice
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3, 4, 5}, expected: []int{5, 4, 3, 2, 1}},
		{input: []int{1, 2}, expected: []int{2, 1}},
		{input: []int{1}, expected: []int{1}},
		{input: []int{}, expected: []int{}},
	}

	for i, test := range tests {
		head := createList(test.input)
		reversed := reverseList(head)
		result := listToSlice(reversed)

		match := len(result) == len(test.expected)
		for j := range result {
			if result[j] != test.expected[j] {
				match = false
				break
			}
		}

		if !match {
			panic(fmt.Sprintf("Example %d failed ❌: input: %v | Expected: %v | Got: %v",
				i+1, test.input, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
