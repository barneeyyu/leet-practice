# 3. Longest Substring Without Repeating Characters

## 題目連結
[LeetCode](https://leetcode.com/problems/longest-substring-without-repeating-characters)

## 題目簡述
給定一個字串，請找出最長的子字串，且其中不包含重複的字元

## 解法
用sliding window搭配map來解，map主要在記字元出現過與否以及紀錄字元出現的index，好讓我們去知道如有重複字元那我們的left要移動到哪

## 時間 / 空間複雜度
- 時間：O(n)
- 空間：O(n) 

## 心得
現在已經知道在這種一定的字串或陣列中找尋最長的子字串或子陣列，可以考慮用sliding window來解，利用動態的left和right來找到最長的子字串或子陣列，一開始只想說用map紀錄字元出現過與否，但沒想到還可以紀錄字元出現的index，這樣就可以知道如遇到重複字元時，left要移動到哪，這樣就可以避免重複計算了
