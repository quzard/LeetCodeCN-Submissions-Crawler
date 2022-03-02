var res []int
func findKthLargest(nums []int, k int) int {
    if k == 0 || k > len(nums){
        return 0
    }
    res = make([]int, 0, k)
    for _, num := range nums {
        if len(res) < k {
            res = append(res, num)
            swim(len(res)-1)
            sink(0)
        } else if res[0] < num {
            res[0] = num
            sink(0)
        }
    }
    return res[0]
}

func sink (i int) {
    for i * 2 < len(res) - 1 {
        j := i * 2
        if res[j+1] < res[j] {
            j++
        }
        if res[i] <= res[j] {
            return
        }
        res[i], res[j] = res[j], res[i]
        i = j
    }
}

func swim(i int) {
    for i > 0 {
        j := i / 2
        if res[i] >= res[j] {
            return
        }
        res[i], res[j] = res[j], res[i]
        i = j
    }
}