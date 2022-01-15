## 滚动哈希

把数组等的定长一部分作为哈希表的key

```go
// 源自187. 重复的DNA序列(中等)
const L = 10

func findRepeatedDnaSequences1(s string) (ans []string) {
	cnt := map[string]int{}
	for i := 0; i <= len(s)-L; i++ {
		sub := s[i : i+L]
		cnt[sub]++
		if cnt[sub] == 2 {
			ans = append(ans, sub)
		}
	}
	return
}
```

## 案例

| 序号 | 题目                                                         |
| ---- | ------------------------------------------------------------ |
| 1    | [187. 重复的DNA序列(中等)](https://leetcode-cn.com/problems/repeated-dna-sequences) |
