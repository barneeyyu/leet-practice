# 155. Min Stack

## 題目連結
[LeetCode](https://leetcode.com/problems/min-stack/)

## 題目簡述
實作一個MinStack，包含push, pop, top, getMin四個方法

## 解法
用golang的slice來實作stack，也因為slice本身就像stack的特性一樣，先進後出，所以可以很方便的實作stack。而且不論是push, pop, top都是O(1)，getMin的話，原本的解法是O(n)，後來看了別人的解法，發現可以用兩個stack來解，一個stack存數字，一個stack存最小值，這樣就可以把`getMin`降到O(1)

## 時間 / 空間複雜度
- 時間：O(n) 
- 空間：O(n)

## 心得
我原本的`getMin`是O(n)，後來看了別人的解法，發現可以用兩個stack來解，一個stack存數字，一個stack存最小值，這樣就可以把`getMin`降到O(1)

原本的： 這樣的時間複雜度是O(n)
```go
func (this *MinStack) GetMin() int {
	tempMin := this.minStack[0]
	for _, v := range this.minStack {
		if v < tempMin {
			tempMin = v
		}
	}
	return tempMin
}
```
