# 单调栈

单调栈实际上就是栈，只是限制要比普通的栈更严格而已了。要求是每次入栈的元素必须要有序（如果新元素入栈不符合要求，则将之前的元素出栈，直到符合要求再入栈），使之形成单调递增/单调递减的一个栈。

- 单调递增栈：只有比栈顶小的才能入栈，否则就把栈顶出栈后，再入栈。出栈时可能会有一些计算。适用于求解第一个大于该位置元素的数。
- 单调递减栈：与单调递增栈相反。适用于求解第一个小于该位置元素的数。

## 如何判断

单调递增/递减栈是根据出栈后的顺序来决定的。例如，栈内顺序[1, 2, 6]，出栈后顺序[6, 2, 1]，这就是单调递减栈。

## 适用场景

单调栈一般用于解决 第一个大于 xxx 或者 第一个小于 xxx 这种题目。

## 哨兵技巧

对于有些时候，如果会用到数组的全部元素，即栈中的元素最后都要出栈，那么很可能因为没有考虑边界而无法通过。所以我们可以使用 **哨兵法** 。

例如在 {1, 3, 4, 5, 2, 9, 6} 末尾添加一个 -1 作为哨兵，变成了 {1, 3, 4, 5, 2, 9, 6, -1} ，这种技巧可以简化代码逻辑。

## 模板

```go
package main

func trap(arr []int){
	n := len(arr)
	stack := []int{}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
    // arr[stack[i]] 单调递减
		for len(stack) > 0 && arr[i] > arr[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[i] = i - top
		}
		stack = append(stack, i)
	}
	return ans[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```



## 案例

| 序号 | 题目                                                         |
| ---- | ------------------------------------------------------------ |
| 1    | [42. 接雨水（困难）](https://leetcode-cn.com/problems/trapping-rain-water/) |
| 2    | [84. 柱状图中最大的矩形（困难）](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/) |
| 3    | [739. 每日温度（中等）](https://leetcode-cn.com/problems/daily-temperatures/) |
| 4    | [496. 下一个更大元素 I（简单）](https://leetcode-cn.com/problems/next-greater-element-i/) |
| 5    | [316. 去除重复字母（困难）](https://leetcode-cn.com/problems/remove-duplicate-letters/) |
| 6    | [901. 股票价格跨度（中等）](https://leetcode-cn.com/problems/online-stock-span/) |
| 7    | [402. 移掉K位数字](https://leetcode-cn.com/problems/remove-k-digits/) |
| 8    | [581. 最短无序连续子数组](https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/) |
