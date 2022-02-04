# 堆(优先队列)

## 优先队列算法

优先队列与普通队列的区别：普通队列遵循先进先出的原则。优先队列的出队顺序与入队顺序无关，与优先级相关，它支持插入和删除最小值操作（返回并删除最小元素）或删除最大值操作（返回并删除最大元素）。在有些情况下，可能需要找到元素集合中的最小或者最大元素，可以利用优先队列来完成操作。

### 优先队列的主要操作 

优先队列是元素的容器，每个元素有一个相关的键值；

- `insert(key, data)`：插入键值为key的数据到优先队列中，元素以其key进行排序；
- `deleteMin/deleteMax`：删除并返回最小/最大键值的元素；
- `getMinimum/getMaximum`：返回最小/最大剑指的元素，但不删除它；

### 优先队列的辅助操作

- `第k最小/第k最大`：返回优先队列中键值为第k个最小/最大的元素；
- `大小（size）`：返回优先队列中的元素个数；
- `堆排序（Heap Sort）`：基于键值的优先级将优先队列中的元素进行排序；

### 优先队列的应用

- 数据压缩：赫夫曼编码算法；
- 最短路径算法：Dijkstra算法；
- 最小生成树算法：Prim算法；
- 事件驱动仿真：顾客排队算法；
- 选择问题：查找第k个最小元素；

## 堆

**堆是一颗具有特定性质的二叉树**，堆的基本要求就是堆中所有结点的值必须大于或等于（或小于或等于）其子结点的值，这也称为堆的性质

**堆并不一定是完全二叉树**，平时使用完全二叉树的原因是易于存储，并且便于索引。

![v2-4ff9b2dda8a84334d35174519cd05b60_1440w](https://pic1.zhimg.com/80/v2-4ff9b2dda8a84334d35174519cd05b60_1440w.jpg)

### 二叉堆

在二叉堆中，每个结点最多有两个孩子结点，在实际应用中，二叉堆已经足够满足需求，因此接下来主要讨论二叉最小堆和二叉最大堆；

**堆的表示：**在描述堆的操作前，首先来看堆是怎样表示的，一种可能的方法就是使用数组，因为堆在形式上是一颗完全二叉树，用数组来存储它不会浪费任何空间

![v2-b6c5427d867469e50d68b703754a432a_1440w](https://pic3.zhimg.com/80/v2-b6c5427d867469e50d68b703754a432a_1440w.jpg)

用数组来表示堆不仅不会浪费空间还具有一定的优势：

- 每个结点的左孩子为下标i的2倍：`left child(i) = i * 2`；每个结点的右孩子为下标i的2倍加1：`right child(i) = i * 2 + 1`
- 每个结点的父亲结点为下标的二分之一：`parent(i) = i / 2`，注意这里是整数除，2和3除以2都为1，大家可以验证一下；
- **注意：**这里是把下标为0的地方空出来了的，主要是为了方便理解，如果0不空出来只需要在计算的时候把i值往右偏移一个位置就行了（也就是加1，大家可以试试，下面的演示也采取这样的方式）；

## 模板

### 模板1 自己实现

算法描述：首先建一个堆，然后调整堆，调整过程是将节点和子节点进行比较，将 其中最大的值变为父节点，递归调整调整次数lgn,最后将根节点和尾节点交换再n次 调整**O(nlgn)**.

#### 算法步骤

- 创建最大堆或者最小堆（下面的实现是最小堆）
- 调整堆
- 交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)

```go
package main

import "fmt"

//堆排序
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	heapSize := len(arr)
	buildMaxHeap(arr, heapSize)
	fmt.Println(arr)
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		// 保证a[i] >= a[i*2+1] 左子节点 && a[i] >= a[i*2+2] 右子节点
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l <= heapSize && a[l] > a[largest] {
		largest = l
	}
	if r <= heapSize && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

```

 



### 模板2 使用`container/heap`

1. `import  "container/heap"`
2. 定义一个数据结构 `type IntHeap ****`
3. 定义
   - `func (h IntHeap) Len() int `
   - `func (h IntHeap) Less(i, j int) bool`
   - `func (h IntHeap) Swap(i, j int)`
   - `func (h *IntHeap) Push(x interface{})`
   - `func (h *IntHeap) Pop() interface{}`
4. 初始化 
   1. `h := &IntHeap{}`
   2. `heap.Init(h)`
5. pop: `heap.Pop(&intHeap)`
6. push: `heap.Push(&intHeap, num)`

```go
package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	num := (*h)[len((*h))-1]
	*h = (*h)[:len(*h)-1]
	return num
}

```

```go
import "sort"

type IntHeap struct {
	sort.IntSlice
}

func (h IntHeap) Len() int {
	return len(heap.IntSlice)
}

func (h IntHeap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h IntHeap) Swap(i, j int) {
	h.IntSlice[i], h.IntSlice[j] = h.IntSlice[j], h.IntSlice[i]
}

func (h *IntHeap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	num := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return num
}
```



## 案例