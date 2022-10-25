func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    for l := len(temperatures) - 2; l >= 0; l-- {
        for r := l + 1; r < len(temperatures); {
            if temperatures[r] <= temperatures[l] {
                next := r + res[r]
                if next == r {
                    break
                }
                r = next
            } else {
                res[l] = r - l
                break
            }
        }
    }
    return res
}