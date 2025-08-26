# 3000 Maximum Area of Longest Diagonal Rectangle

## 題目連結
[LeetCode](https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle)

## 題目簡述
給你一個二維整數陣列 dimensions，其中：
- dimensions[i][0] 表示第 i 個矩形的長度
- dimensions[i][1] 表示第 i 個矩形的寬度

請你回傳 擁有最長對角線的矩形的面積。
如果有多個矩形的對角線同樣最長，則回傳其中 面積最大的那個矩形的面積。

## 解法
遍歷每個矩形，計算其對角線長度的平方（避免使用平方根以節省計算資源），同時記錄最大對角線長度的平方和對應的最大面積。最後返回最大面積。

## 時間 / 空間複雜度
- 時間：O(n)，其中n是陣列的長度
- 空間：O(1)

## 心得
偏簡單的題目，但要注意的就是同樣對角線長度但面積不同的情況；也要注意對角線更長但面積更小的情況，所以寫判斷式的時候要分兩種獨立情境，只要滿足其中一種即可。
