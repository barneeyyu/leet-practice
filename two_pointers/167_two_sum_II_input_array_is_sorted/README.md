# 167. Two Sum II - Input Array Is Sorted

## 題目連結
[LeetCode](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

## 題目簡述
跟Two sum I 題目一樣，但這題的數列已經排序過了

## 解法
跟Two sum I 題目一樣，可以用map來解，但這樣空間複雜度會是O(n)，這題已經排序過了，所以可以用雙指針來解，同樣的時間複雜度，空間複雜度會是O(1)

## 時間 / 空間複雜度
- 時間：O(n)
- 空間：O(1)

## 心得
這題一開始想很久，也知道可以用雙指針，但不知道怎麼寫，後來才知道原來就是判斷最左跟最右加起來跟target比較大小，因為有排序過，所以小則進、大則退這麼簡單！
這題似乎也可以用二分搜尋法解，但我看時間複雜度為O(nlogn)，就沒有實際研究了