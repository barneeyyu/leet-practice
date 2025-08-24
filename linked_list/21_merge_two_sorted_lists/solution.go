package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
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
		l1       []int
		l2       []int
		expected []int
	}{
		{l1: []int{1, 2, 4}, l2: []int{1, 3, 4}, expected: []int{1, 1, 2, 3, 4, 4}},
		{l1: []int{}, l2: []int{}, expected: []int{}},
		{l1: []int{}, l2: []int{0}, expected: []int{0}},
	}

	for i, test := range tests {
		l1 := createList(test.l1)
		l2 := createList(test.l2)
		merged := mergeTwoLists(l1, l2)
		result := listToSlice(merged)

		// This is a placeholder check and will fail until the solution is implemented
		if len(result) != len(test.expected) {
			fmt.Printf("Example %d placeholder check failed. Implement the solution.\n", i+1)
		} else {
			fmt.Printf("Example %d placeholder check passed.\n", i+1)
		}
	}
}
