# 242. Valid Anagram

## 題目連結
[LeetCode](https://leetcode.com/problems/valid-anagram)

## 題目簡述
判斷一個字串s是否包含另一個字串t的所有字元

## 解法
- 使用 Go 的 map[rune]int先計算t中每個字元的出現次數，再遍歷s，將s中每個字元的出現次數減一，如果減到0以下，則返回false。

## 時間 / 空間複雜度
- 時間：O(n+m) 其中n是s的長度，m是t的長度
- 空間：O(k) 因為k是map的裡不同字元的個數

## 心得
這題用map效率最高，但需要考慮到字元是unicode，所以需要使用rune來表示字元。
如果是先排序兩個字串，再比較，時間複雜度會是O(nlogn+mlogm)，空間複雜度是O(1)。