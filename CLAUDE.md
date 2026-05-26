# CLAUDE.md — leet-practice

## 這個 Repo 在做什麼

這是 Barney 的 LeetCode 刷題紀錄，用 **Go 語言**實作解法，並按照題型分類整理。
每次新增題目後，GitHub Actions 會自動觸發 `scripts/update_readmes.py`，呼叫 LeetCode API 查難度，並更新各題型的 `README.md` 以及根目錄的難度統計徽章。

## 目錄結構

```
<題型>/
  <題號>_<題目名稱>/
    solution.go   ← Go 解法（含 main 做 test）
    README.md     ← 題目筆記
  README.md       ← 題型總表（自動生成，勿手動修改）
README.md         ← 根目錄總覽（難度徽章自動更新）
scripts/
  update_readmes.py  ← 自動生成 README 的腳本
.github/workflows/
  update-readmes.yaml  ← GitHub Actions 設定
```

支援的題型資料夾：`array`, `binary_search`, `dfs_bfs`, `dynamic_programming`, `greedy`, `hashmap`, `sliding_window`, `stack_queue`, `string`, `two_pointers`, `linked_list`, `tree`

## 新增題目

當使用者說「新增題目」或「add problem」，依照以下步驟建立檔案：

1. 判斷題型，確認放在哪個資料夾
2. 建立資料夾：`<題型>/<題號>_<題目名稱小寫底線>/`
3. 建立 `solution.go`（見下方 template）
4. 建立 `README.md`（見下方 template）

### solution.go template

```go
package main

import (
	"fmt"
)

// TODO: 加上題目需要的 struct（例如 TreeNode、ListNode）

// <functionName> TODO: 補上函式簽名
func <functionName>(<params>) <returnType> {
	// TODO: 實作
}

func main() {
	tests := []struct {
		input    <inputType>
		expected <returnType>
	}{
		// TODO: 填入 test cases（3~4 個，涵蓋 edge case + 一般 case）
	}

	for i, test := range tests {
		result := <functionName>(test.input)
		if result != test.expected {
			panic(fmt.Sprintf("Example %d: Input: %v | Expected: %v | Got: %v", i+1, test.input, test.expected, result))
		} else {
			fmt.Printf("Example %d passed ✅\n", i+1)
		}
	}
}
```

### README.md template

```markdown
# <題號>. <題目名稱>

## 題目連結
[LeetCode](https://leetcode.com/problems/<題目-slug>/)

## 題目簡述
<!-- TODO: 用一兩句話描述題目 -->

## 解法
<!-- TODO: 描述解法思路 -->

## 時間 / 空間複雜度
- 時間：O(?)
- 空間：O(?)

## 心得
<!-- TODO: 寫下解題過程的想法或學到的東西 -->
```

## 注意事項

- 各題型的 `README.md`（總表）由腳本自動生成，**不要手動修改**
- solution.go 用 `main` package，在 `main()` 裡做測試，不需要引入測試框架
- test case 格式統一用 table-driven（struct slice），3~4 個 case 即可
- 題目資料夾命名格式：`<題號>_<題目名稱全小寫底線>`，例如 `226_invert_binary_tree`
