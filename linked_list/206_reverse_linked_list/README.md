# 206. Reverse Linked List

## 題目連結
[LeetCode - Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/)

## 題目簡述
給定一個單向鏈結串列的頭節點 head，請你反轉這個鏈結串列，並回傳反轉後的頭節點。

## 解法
使用雙指針解法，透過三個指標：
- `prev`: 指向已反轉部分的尾端（初始為 `null`）
- `curr`: 當前正在處理的節點
- `temp`: 暫存 `curr.next`，以防止指標斷掉

每次迴圈都將 `curr.next` 指向 `prev`，然後更新 `prev` 和 `curr` 的位置。

這題還有**遞迴**寫法，不過主要是雙指針要先搞懂，概念差不多！

## 時間 / 空間複雜度
- 時間：O(n)，每個節點都會被訪問一次
- 空間：O(1)，只使用常數額外空間

## 心得
這題是經典的鏈結串列反轉問題，一開始會對指標操作感到困惑，但畫出流程圖會清楚很多。
熟練這題有助於處理所有鏈結串列相關的進階題目。