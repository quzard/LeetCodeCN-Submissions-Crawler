func reverse(x int) int {
    abs := false
    if x < 0 {
        abs = true
        x = -x
    }
    res := 0
    for x > 0 {
        res = res * 10 + x % 10
        x /= 10
    }
    if abs {
        res = -res
    }
    if res > math.MaxInt32 || res < math.MinInt32 {
        return 0
    }
    return res
}