func reverse(x int) int {
    f := false
    if x < 0 {
        f = true
        x = -x
    }
    r := 0
    for x > 0 {
        if r * 10 + x % 10 > math.MaxInt32 {
            return 0
        }
        r = r * 10 + x % 10
        x /= 10
    }

    if f {
        if -r < math.MinInt32{
            return 0
        }
        return -r
    }
    return r
}