func isHappy(n int) bool {
    selected := map[int]bool{}
    var dfs func(n int) bool
    dfs = func(n int) bool {
        if n == 1 {
            return true
        }
        if selected[n] == false {
            selected[n] = true
        } else {
            return false
        }
        sum := 0
        for n > 0 {
            sum += (n % 10) * (n % 10)
            n = int(n / 10)
        }
        return dfs(sum)
    }
    return dfs(n)
}