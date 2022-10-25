import (
    "container/heap"
)

// 堆排序
var res []int
var hashTabe map[int]int
var selected map[int]int

func topKFrequent(nums []int, k int) []int {
    res = make([]int, 1)
    hashTabe = map[int]int{}
    selected = map[int]int{}
    for i := 0; i < len(nums); i++ {
        hashTabe[nums[i]]++
    }
    for key, value := range hashTabe {
        if len(res) < k+1 {
            res = append(res, key)
            selected[key] = len(res) - 1
            swim(len(res) - 1)
        } else {
            if value >= hashTabe[res[1]] {
                selected[res[1]] = 0
                res[1] = key
                selected[res[1]] = 1
                sink(1)
            }
        }
    }
    return res[1:]
}

// 父比子大
func less(a, b int) bool {
    if hashTabe[res[a]] >= hashTabe[res[b]] {
        return true
    }
    return false
}

// 从下往上浮 越上越小
func swim(k int) {
    for k > 1 && less(int(k/2), k) {
        selected[res[int(k/2)]], selected[res[k]] = k, int(k/2)
        res[int(k/2)], res[k] = res[k], res[int(k/2)]
        k = k / 2
    }
}

// 从上往下沉 越下越大
func sink(k int) {
    for 2*k < len(res) {
        j := 2 * k
        if j < len(res)-1 && less(j, j+1) {
            // 取j 和 j+1 小的那个
            j++
        }
        if !less(k, j) {
            // k < j
            break
        }
        selected[res[j]], selected[res[k]] = k, j
        res[j], res[k] = res[k], res[j]
        k = j
    }
}

func topKFrequent1(nums []int, k int) []int {
    occurrences := map[int]int{}
    for _, num := range nums {
        occurrences[num]++
    }
    h := &IHeap{}
    heap.Init(h)
    for key, value := range occurrences {
        heap.Push(h, [2]int{key, value})
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    ret := make([]int, k)
    for i := 0; i < k; i++ {
        ret[k-i-1] = heap.Pop(h).([2]int)[0]
    }
    return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
    *h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}