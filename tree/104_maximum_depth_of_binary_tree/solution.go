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

// maxDepth returns the maximum depth of a binary tree.
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func main() {
	// Build test trees
	// case1: nil tree -> depth 0
	var t1 *TreeNode = nil

	// case2: [3,9,20,null,null,15,7] -> depth 3
	t2 := &TreeNode{Val: 3,
		Left:  &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
	}

	// case3: [1,null,2] -> depth 2
	t3 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	// case4: skewed left [1,2,3,4,5] -> depth 5
	t4 := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left: &TreeNode{Val: 3,
				Left: &TreeNode{Val: 4,
					Left: &TreeNode{Val: 5},
				},
			},
		},
	}

	// case5: single node [0] -> depth 1
	t5 := &TreeNode{Val: 0}

	// case6: perfect tree depth 3: [1,2,3,4,5,6,7]
	t6 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}},
	}

	testCases := []struct {
		name   string
		root   *TreeNode
		result int
	}{
		{
			name:   "nil tree",
			root:   t1,
			result: 0,
		},
		{
			name:   "[3,9,20,null,null,15,7]",
			root:   t2,
			result: 3,
		},
		{
			name:   "[1,null,2]",
			root:   t3,
			result: 2,
		},
		{
			name:   "skewed left depth=5",
			root:   t4,
			result: 5,
		},
		{
			name:   "single node [0]",
			root:   t5,
			result: 1,
		},
		{
			name:   "perfect tree depth=3",
			root:   t6,
			result: 3,
		},
	}

	for i, tc := range testCases {
		got := maxDepth(tc.root)
		if got != tc.result {
			fmt.Printf("Test %d failed ❌:\n  Case: %s\n  Expected: %d\n  Got: %d\n", i+1, tc.name, tc.result, got)
		} else {
			fmt.Printf("Test %d passed ✅\n", i+1)
		}
	}
}
