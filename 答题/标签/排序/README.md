# 排序

## 快速排序

算法描述：是对插入算法的一种优化，利用对问题的二分化，实现递归完成快速排序 ，在所有算法中二分化是最常用的方式，将问题尽量的分成两种情况加以分析， 最终以形成类似树的方式加以利用，因为在比较模型中的算法中，最快的排序时间 负载度为 **O(nlgn)**.

#### 算法步骤

- 将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于
- 将两边数据进行递归调用步骤1
- 将所有数据合并

```go
func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	length := end - start + 1
	index := start + rand.Intn(length)%length
	pivot := nums[index]
    
	nums[end], nums[index] = nums[index], nums[end]
	left, right := start-1, start
	for right <= end {
		if nums[right] > pivot {
			right++
			continue
		}
		left++
		nums[left], nums[right] = nums[right], nums[left]
		right++
	}
	quickSort(nums, start, left-1)
	quickSort(nums, left+1, end)
}
```

## 堆排序

### 模板1

算法描述：首先建一个堆，然后调整堆，调整过程是将节点和子节点进行比较，将 其中最大的值变为父节点，递归调整调整次数lgn,最后将根节点和尾节点交换再n次 调整**O(nlgn)**.

```go
func sortArray(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	end := len(nums) - 1
	buildMaxHeap(nums, end)
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		maxHeapify(nums, 0, end)
	}
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

### 模板2
```go
import "container/heap"

type IntHeap struct {
	arr []int
}

func (h *IntHeap) Len() int {
	return len(h.arr)
}

func (h *IntHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *IntHeap) Less(i, j int) bool {
	return h.arr[i] < h.arr[j]
}

func (h *IntHeap) Push(num interface{}) {
	h.arr = append(h.arr, num.(int))
}

func (h *IntHeap) Pop() interface{} {
	num := h.arr[h.Len()-1]
	h.arr = h.arr[:h.Len()-1]
	return num
}

func SortHeap(arr []int) []int {
	h := &IntHeap{arr: arr}
	heap.Init(h)
	res := make([]int, len(arr))
	for i := 0; i < len(res); i++ {
		res[i] = heap.Pop(h).(int)
	}
	return res
}
```

## 冒泡排序

算法描述：冒泡算法，数组中前一个元素和后一个元素进行比较如果大于或者小于 前者就进行交换，最终返回最大或者最小都冒到数组的最后序列时间复杂度为 **O(n^2)**.

#### 算法步骤

- 从数组开头选择相邻两个元素进行比较，并进行交换
- 不停向后移动

```go
//冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

//冒泡排序获取最大值
func GetMax(arr []int) int {
	for j := 1; j < len(arr); j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr[len(arr)-1]
}
```

## 选择排序

算法描述：从未排序数据中选择最大或者最小的值和当前值交换 **O(n^2)**.

#### 算法步骤

- 选择一个数当最小值或者最大值，进行比较然后交换
- 循环向后查进行

```go
//选择排序
func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 1; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
```

## 基数排序

算法描述：基数排序类似计数排序，需要额外的空间来记录对应的基数内的数据 额外的空间是有序的，最终时间复杂度**O(nlogrm)**,r是基数，r^m=n.当给定 特定的范围，计数排序又可以叫桶排序，当以10进制为基数时就是简单的桶排序

#### 算法步骤

- 从个位开始排序，从低到高进行递推
- 比较过程中如果遇到高位相同时，顺序不变

#### 算法分两类

1. 低位排序LSD
2. 高位排序MSD

```go

```

## 拓扑排序
### 定义

对于一些有前后依赖关系的排序算法，是利用有向无环图进行实现，通过局部依赖关系确定全局顺序的算法

### 应用场景

- 编译有序依赖的文件

### 模板1

#### Kahn算法

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

### 模板2

#### DFS算法

1.使用深度算法，产生逆向邻接表先输出其他依赖，最后输出自己。


## 桶排序
```go

```

## 归并排序

```go
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	l, r, i := 0, 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	copy(result[i:], left[l:])
	copy(result[i+len(left)-l:], right[r:])
	return result
}
```