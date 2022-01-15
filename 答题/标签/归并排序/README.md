# 归并排序


```go
func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	return merge(left, right)
}

func merge(l, r []int) []int {
	i, j := 0, 0
	m, n := len(r), len(r)
	res := []int{}
	for i < m && j < n {
		if l[i] < r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, l[j])
			j++
		}
	}
	res = append(res, l[i:]...)
	res = append(res, l[j:]...)
	return res
}
```

## 案例

| 序号 | 题目                                                         |
| ---- | ------------------------------------------------------------ |
| 1    | [23. 合并K个升序链表(困难)](https://leetcode-cn.com/problems/merge-k-sorted-lists) |
