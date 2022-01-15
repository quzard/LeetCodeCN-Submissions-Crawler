# 队列

先进先出



## 模板
```go
// 创建队列
queue := make([]int, 0)
// enqueue入队
queue = append(queue, 10)
// dequeue出队
v := queue[0]
queue = queue[1:]
// 长度0为空
len(queue) == 0
```

## 案例

| 序号 | 题目                                                         |
| ---- | ------------------------------------------------------------ |
| 1    | [387. 字符串中的第一个唯一字符(简单)](https://leetcode-cn.com/problems/first-unique-character-in-a-string) |
| 2    | [剑指 Offer 09. 用两个栈实现队列(简单)](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof) |

