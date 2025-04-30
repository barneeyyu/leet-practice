# 121. Best Time to Buy and Sell Stock

## 題目連結
[LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock)

## 題目簡述
給定一個整數陣列，代表股票在每一天的價格，請找出在哪一天買進，哪一天賣出，可以獲得最大的利潤

## 解法
用sliding window來解，因為我們只需要找到最大的利潤，所以可以用一個變數紀錄最小價格，並且用另一個變數來記錄最大的利潤

## 時間 / 空間複雜度
- 時間：O(n) 其中n是prices的長度
- 空間：O(1) 

## 心得
這題每次寫都要看解答，還不太確定什麼情境下要用sliding window來解，多練習