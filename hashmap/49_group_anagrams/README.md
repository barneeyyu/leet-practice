# 49. Group Anagrams

## 題目連結
[LeetCode](https://leetcode.com/problems/group-anagrams/)

## 題目簡述
給一個字串陣列，將所有的字母異位詞(anagram)分組。字母異位詞是指由相同的字母以不同順序組成的單詞。

## 解法
1. 建立一個 map 來儲存分組結果
2. 對每個字串進行排序，排序後的字串作為 key
3. 將原始字串加入對應的 key 的陣列中
4. 最後將 map 的值轉換成陣列回傳

## 時間 / 空間複雜度
- 時間：O(n * k * log k)，其中 n 是字串數量，k 是最長字串的長度
- 空間：O(n * k)，需要儲存所有字串

## 心得
一開始想了蠻久，後來看解析才意識到，把anagram的詞排序過後就會是一樣的單詞，把這一樣的單詞當成key存到hasp map，再把原本的單詞append到value之中，這樣就簡單了，最後再遍歷map，也因為map裡面已經是[]string了，再把[]string直接append到result的string裡面，這樣就大功告成。
