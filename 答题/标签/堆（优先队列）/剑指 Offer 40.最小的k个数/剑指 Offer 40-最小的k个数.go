var res []int
func getLeastNumbers(arr []int, k int) []int {
    res = make([]int, 0, k)
    if k == 0 {
        return res
    }
    for i := 0; i < len(arr); i++ {
        if len(res) < k {
            res = append(res, arr[i])
            swim(len(res)-1)
        } else if len(res) == k && res[0] > arr[i] {
            res[0] = arr[i]
            sink(0)
        }
    }
    return res
}

func sink(i int) {
    for {
        j1 := i * 2 + 1
        if j1 >= len(res) || j1 < 0 {
            break
        }
        j := j1
        
        j2 := j1 + 1
        if j2 < len(res) && j2 >= 0 && res[j2] > res[j1] {
            j = j2
        }
        if res[i] >= res[j] {
            break
        }
        res[i], res[j] = res[j], res[i]
        i = j
    }
}

func swim(i int) {
    for {
        j := (i-1) / 2 
        if i == j || res[j] >= res[i] {
            break
        }
        res[i], res[j] = res[j], res[i]
        i = j
    }
}