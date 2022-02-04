var res []int
// 桶排序
func getLeastNumbers1(arr []int, k int) []int {
    if k == 0 {
        return []int{}
    }
    res = make([]int, 0, k)
    for _, num := range arr {
        if len(res) < k {
            res = append(res, num)
            swim(len(res) - 1)
            sink(0)
        } else if res[0] > num{
            res[0] = num
            sink(0)
        }
    }
    return res
}

func less(a, b int) bool {
    return res[a] < res[b]
}

func swim(k int) {
    for k > 1 {
        // k / 2 < k
        if less(k / 2, k) {
            res[k], res[k/2] = res[k/2], res[k]
            k = k / 2
        } else {
            break
        }
    }
}

func sink(k int) {
    for 2*k < len(res) - 1{
        j := 2 * k
        // j < j + 1 选择较大的那个
        if less(j, j + 1) {
            j++
        }
        // j <= k
        if !less(k, j) {
            break
        }
        res[k], res[j] = res[j], res[k]
        k = j
    }
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

//Less  小于就是小跟堆，大于号就是大根堆
func (h IntHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{}{
    num := (*h)[h.Len()-1]
    (*h) = (*h)[:h.Len()-1]
    return num
}

func (h *IntHeap) Push(n interface{}) {
    *h = append(*h, n.(int))
}

func (h IntHeap) Peek() int {
    return h[0]
}
// 堆排序
func getLeastNumbers(arr []int, k int) []int {
    if k == 0 {
        return []int{}
    }
    h := &IntHeap{}
    heap.Init(h)
    for _, num := range arr {
        if h.Len()< k {
            heap.Push(h, num)
        }else if h.Peek() > num {
            heap.Pop(h)
            heap.Push(h, num)
        }
    }
    res = make([]int, k)
    for i := 0; i < h.Len(); i++{
        res[i] = (*h)[i]
    }
    return res
}