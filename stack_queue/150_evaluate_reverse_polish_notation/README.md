# 150. Evaluate Reverse Polish Notation

## 題目連結
[LeetCode](https://leetcode.com/problems/evaluate-reverse-polish-notation/)

## 題目簡述
給定一個「逆波蘭表示法」的運算式（例如 ["2", "1", "+", "3", "*"]），請你回傳計算後的結果。逆波蘭表示法是一種不需要括號的運算方式，運算子會放在操作數的後面。

**Example 1:**  
Input: tokens = ["2","1","+","3","*"]  
Output: 9  
Explanation: ((2 + 1) * 3) = 9

**Example 2:**  
Input: tokens = ["4","13","5","/","+"]  
Output: 6  
Explanation: (4 + (13 / 5)) = 6

## 解法


## 時間 / 空間複雜度
- 時間：O(n) 
- 空間：O(n) 


## 心得
這題其實就是stack的應用，遇到數字就push，遇到運算子就pop兩個數字來做運算，運算完再把結果push回去stack，最後stack裡面剩下的數字就是答案。 
也因為stack的特性，所以在操作陣列裡面的元素時不管拿出來或放進去都很方便