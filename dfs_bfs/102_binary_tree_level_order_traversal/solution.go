package main

import (
	"fmt"
	"reflect"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelArray := make([]int, 0, levelSize) // 會這樣寫是因為已知陣列的長度，就避免了這種「擴容」的額外開銷，因為一開始就把空間開到剛好可以裝下這層所有節點

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelArray = append(levelArray, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelArray)
	}
	return res
}

// TreeNode 定義
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 測資建樹小工具
func buildTree() *TreeNode {
	// 建立 [3,9,20,null,null,15,7]
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func main() {
	testCases := []struct {
		input    *TreeNode
		expected [][]int
	}{
		{
			input: buildTree(),
			expected: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
		{
			input:    &TreeNode{Val: 1}, // 單節點
			expected: [][]int{{1}},
		},
		{
			input:    nil, // 空樹
			expected: [][]int{},
		},
	}

	for i, tc := range testCases {
		got := levelOrder(tc.input)
		if reflect.DeepEqual(got, tc.expected) {
			fmt.Printf("Test %d passed ✅\n", i+1)
		} else {
			fmt.Printf("Test %d failed ❌:\n  Expected: %v\n  Got: %v\n", i+1, tc.expected, got)
		}
	}
}
