# 拓扑排序

## 定义

对于一些有前后依赖关系的排序算法，是利用有向无环图进行实现，通过局部依赖关系确定全局顺序的算法

## 应用场景

- 编译有序依赖的文件

## 模板1

### Kahn算法

- 算法逻辑
- 利用贪心算法，如果两个顶点，顶点b依赖于顶点a,就将a指向b,将a的入度加1。当一个顶点的入度为0，将这个顶点就是最优排序点， 并且将顶点从图中移除，将可达顶点的入度减一。

```go
// 来自1136. 平行课程
// 将 1 到 N，改成 0 到 N - 1，方便数组索引
func minimumSemesters(N int, relations [][]int) int {
	preClassCount := map[int]int{}
	nextClasses := make([][]int, N)
	for i := 0; i < N; i++ {
		preClassCount[i] = 0
	}
	for _, r := range relations {
		nextClasses[r[0]-1] = append(nextClasses[r[0]-1], r[1]-1) // 记录后置课程
		preClassCount[r[1]-1]++                                   // 入度加1
	}
	term := 0
	for len(preClassCount) > 0 {
		term++
		learn := []int{}
		for class, count := range preClassCount {
			if count == 0 {
				learn = append(learn, class) // 入度为 0，可以学习
			}
		}
		if len(learn) == 0 { // 没有课程可以学了，说明有循环
			return -1
		}
		for _, l := range learn {
			delete(preClassCount, l)
			for _, class := range nextClasses[l] {
				preClassCount[class]--
			}
		}
	}
	return term
}
```

## 模板2

### DFS算法

1.使用深度算法，产生逆向邻接表先输出其他依赖，最后输出自己。

