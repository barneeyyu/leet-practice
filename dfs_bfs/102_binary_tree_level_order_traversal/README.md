# 102. Binary Tree Level Order Traversal

## 題目連結
[LeetCode](https://leetcode.com/problems/binary-tree-level-order-traversal/)

## 題目簡述
給一個二元樹，請以層序遍歷的方式回傳每一層的節點值，用二元陣列的方式回傳，外層 slice 代表「第幾層」，內層 slice 存該層的所有節點值。

## 解法
1. 使用 Queue 的特性來進行層序遍歷
2. 初始化一個 Queue，將根節點加入 Queue
3. 當 Queue 不為空時，進行以下操作：
   - 取得當前 Queue 的長度，這代表當前層的節點數量
   - 初始化一個空的 slice 來存放當前層的節點值
   - 遍歷當前層的所有節點，將節點值加入 slice，並將該節點的左子節點和右子節點（如果存在）加入 Queue
   - 將當前層的節點值 slice 加入結果 slice
4. 返回結果 slice

## 時間 / 空間複雜度
- 時間：O(n)，其中 n 是二元樹的節點數量，每個節點會被訪問一次
- 空間：O(m)，其中 m 是二元樹的最大寬度

## 心得
一開始沒想過要用queue去儲存每一層的節點，後來發現原來把根節點放進去後，queue的長度就代表這一層的節點數量，然後每次從queue裡面pop出來的節點再把它的子節點放進去queue，這樣就能一層一層的遍歷整棵樹了，整個過程很直觀也很容易理解。