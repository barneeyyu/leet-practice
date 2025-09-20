# 104. Maximum Depth of Binary Tree

## 題目連結
[LeetCode](https://leetcode.com/problems/maximum-depth-of-binary-tree)

## 題目簡述
給定一個二元樹，請找出它的最大深度（maximum depth）。最大深度是從根節點到最遠葉節點的最長路徑上的節點數量。

## 解法
用後序遍歷的方式：左 -> 右 -> 根
1. 如果節點為 nil，回傳 0
2. 遞迴計算左子樹的深度 leftDepth
3. 遞迴計算右子樹的深度 rightDepth
4. 回傳 max(leftDepth, rightDepth) + 1


## 時間 / 空間複雜度
- 時間： O(n)
- 空間： O(h)，h 為樹的高度，因為遞迴的過程中，呼叫棧（call stack）會佔用額外空間。
在最壞情況下（例如：樹退化成一條鏈），遞迴深度會等於樹的高度 h。

## 心得
之前覺得這題很難理解，後來發現
求高度 -> 用後序遍歷（左右中）
求深度 -> 用前序遍歷（中左右）
剛好這題要求最大深度就是根節點的高度，所以用後序方便
然後原本是
```go
if root == nil {
    return 0
}

left := maxDepth(root.Left)
right := maxDepth(root.Right)
height := 1 + max(left, right)

return height
```
最後可以濃縮成一行！