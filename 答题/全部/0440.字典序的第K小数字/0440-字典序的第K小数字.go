//字典树
func findKthNumber(n int, k int) int {
    var getCount func(pre int) int
    getCount = func(pre int) int {
        count := 0
        next := pre + 1
        for pre <= n{
            count += min(n + 1, next) - pre
            pre *= 10
            next *= 10
        }
        return count
    }
    count := 1
    pre := 1
    for count < k{
        c := getCount(pre)
        if count + c > k{
            count++
            pre *= 10
        }else{
            pre++
            count += c
        }
    }
    return pre
}

func min(a, b int) int{
    if a < b{
        return a
    }
    return b
}