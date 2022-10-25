func mySqrt(x int) int {
    l, r := 0, x
    mid := 0
    for l <= r {
        mid = l + (r-l) / 2
        if x >= mid*mid && x < (mid+1)*(mid+1) {
            return mid
        }
        if x > mid*mid {
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return mid
}