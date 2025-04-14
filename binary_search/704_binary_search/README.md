# 704. Binary Search

## 題目連結
[LeetCode](https://leetcode.com/problems/binary-search/)

## 題目簡述
給一個已經排序好的陣列，判斷target是否在陣列中，如果target在陣列中，回傳target在陣列中的index，如果target不在陣列中，回傳-1

## 解法
使用二分搜尋法，每次都取中間的值，如果中間的值大於target，就往左邊找，如果中間的值小於target，就往右邊找，直到找到target為止

## 時間 / 空間複雜度
- 時間：O(log n) 
- 空間：O(1) 

## 心得
這題切記有“左閉右開”跟“左閉右閉”的觀念，這兩種觀念在寫二分搜尋法的時候，會影響到while的條件，以及left和right的邊界，所以切記要分清楚