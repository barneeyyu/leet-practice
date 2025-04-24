# 74. Search a 2D Matrix

## 題目連結
[LeetCode](https://leetcode.com/problems/search-a-2d-matrix/)

## 題目簡述
給一個已經排序好的2D矩陣，判斷target是否在矩陣中，如果target在矩陣中，回傳target在矩陣中的index，如果target不在矩陣中，回傳-1

## 解法
最快的時間是用二分搜尋法，把2D矩陣看作一個1D矩陣，然後用二分搜尋法來找target，就蠻簡單的了


## 時間 / 空間複雜度
- 時間：O(log(m*n)) 
- 空間：O(1) 

## 心得
一開始用了暴力解，先找到target所在的row，再找到target所在的column，這樣時間複雜度會是O(m*n)，所以需要使用二分搜尋法來優化
```go
func searchMatrix(matrix [][]int, target int) bool {
    row := len(matrix)
    col := len(matrix[0])

    if target < matrix[0][0] {
        return false
    }

    targetRow := -1
    for i := 0; i < row; i++ {
        if target >= matrix[i][0] {
            targetRow = i
        } else {
            break
        }
    }
    if targetRow == -1 {
        return false
    }
    for j := 0; j < col; j++ {
        if matrix[targetRow][j] == target {
            return true
        }
    }
    return false
}
```
一開始用二元搜尋法的時候還想了一下，mid_row, mid_col是怎麼來的，後來想了一下，mid_row是mid除以col，mid_col是mid除以col的餘數，這樣就可以找到target所在的row和column了，然後一樣用左開右閉的觀念，就可以找到target了