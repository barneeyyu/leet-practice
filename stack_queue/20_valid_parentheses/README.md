# 20. Valid Parentheses

## 題目連結
[LeetCode](https://leetcode.com/problems/valid-parentheses/)

## 題目簡述
給一個字串，只包含括號，判斷是否為有效的括號(有效的括號，代表括號的順序是正確的，只要有一個左括號，就必須有一個右括號)


## 解法
因為括號勢必要對稱，所以用stack來解，遇到左括號就push，遇到右括號就pop，如果pop出來的括號不是對應的左括號，就return false

## 時間 / 空間複雜度
- 時間：O(n) stack的pop和push都是O(1)
- 空間：O(n) 因為需要一個stack來存括號


## 心得
其實這題簡單來說就是會先想到stack的特性，後面放進去的括號，就要搭配對稱的括號能先被拿出來，但在golang裡面沒有stack的資料結構，所以需要自己實作一個stack，如果實作stack的話，也可以直接操作slice來實現stack，這樣就不需要自己實作stack了。