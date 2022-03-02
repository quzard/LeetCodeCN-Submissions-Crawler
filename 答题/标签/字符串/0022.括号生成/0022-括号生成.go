func generateParenthesis(n int) []string {
    res := []string{}
    s := []byte{}
    var dfs func(l, r int)
    dfs = func(l, r int) {
        if len(s) == n * 2 {
            res = append(res, string(s))
            return
        }
        if l < n {
            s = append(s, '(')
            l++
            dfs(l, r)
            l--
            s = s[:len(s)-1]
        }
        if l > r {
            s = append(s, ')')
            r++
            dfs(l, r)
            r--
            s = s[:len(s)-1]
        }
    }
    dfs(0, 0)
    return res
}