# 226. Invert Binary Tree

## 題目連結
[LeetCode](https://leetcode.com/problems/invert-binary-tree)

## 題目簡述
給定一個二元樹，請反轉它的結構（invert the tree），並回傳反轉後的樹。

## 解法
用遞迴（DFS）解，對當前節點進行左右子樹交換，再遞迴處理左右子樹。
Base case：若節點為 nil，直接回傳。


## 時間 / 空間複雜度
- 時間： O(n)，每個節點都要訪問一次，並交換一次左右子樹
- 空間： O(h)，遞迴呼叫堆疊的空間複雜度，h 是樹的高度

## 心得
用遞迴解這題很直觀，但如果要用queue來解的話，要從root開始一層一層的交換左右子樹，比較複雜。我覺得用遞迴比較簡潔。