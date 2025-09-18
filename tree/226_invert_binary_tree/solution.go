package main

import (
	"fmt"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree inverts a binary tree and returns its root.
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// equal checks structural and value equality between two trees.
func equal(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	if a.Val != b.Val {
		return false
	}
	return equal(a.Left, b.Left) && equal(a.Right, b.Right)
}

func main() {
	// === Test cases for 226. Invert Binary Tree ===
	// Keep 3~4 representative cases to be concise and cover edge + typical.

	// case1: nil tree -> nil
	var t1 *TreeNode = nil
	var e1 *TreeNode = nil

	// case2: single node [1] -> [1]
	t2 := &TreeNode{Val: 1}
	e2 := &TreeNode{Val: 1}

	// case3: balanced tree [4,2,7,1,3,6,9] -> [4,7,2,9,6,3,1]
	t3 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}},
	}
	e3 := &TreeNode{Val: 4,
		Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
	}

	// case4: skewed-left [1,2,3,4] -> skewed-right [1,null,2,null,3,null,4]
	t4 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 4},
			},
		},
	}
	e4 := &TreeNode{Val: 1,
		Right: &TreeNode{Val: 2,
			Right: &TreeNode{Val: 3,
				Right: &TreeNode{Val: 4},
			},
		},
	}

	testCases := []struct {
		name     string
		input    *TreeNode
		expected *TreeNode
	}{
		{"nil tree", t1, e1},
		{"single node", t2, e2},
		{"balanced sample", t3, e3},
		{"skewed-left -> right", t4, e4},
	}

	for i, tc := range testCases {
		// Deep-copy the input if you need to preserve original; here we mutate in-place and compare.
		got := invertTree(tc.input)
		if !equal(got, tc.expected) {
			fmt.Printf("Test %d failed ❌: %s\n", i+1, tc.name)
		} else {
			fmt.Printf("Test %d passed ✅\n", i+1)
		}
	}
}
