//×ÖµäÊ÷
func findKthNumber(n int, k int) int {
    var getCount func(pre int) int
    getCount = func(pre int) int {
        next := pre + 1
        cnt := 0
        for pre <= n {
            cnt += min(n+1, next) - pre
            pre *= 10
            next *= 10
        }
        return cnt
    }
    cnt := 1
    pre := 1
    for cnt < k {
        c := getCount(pre)
        if c+cnt > k {
            pre *= 10
            cnt++
        } else {
            pre++
            cnt += c
        }
    }
    return pre
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}