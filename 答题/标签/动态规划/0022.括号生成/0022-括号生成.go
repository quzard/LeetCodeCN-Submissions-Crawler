func generateParenthesis(n int) []string {
    res := []string{}
    word := []byte{}
    var dfs func(l, r int)
    dfs = func(l, r int) {
        if len(word) == n*2 {
            res = append(res, string(word))
            return
        }
        if l < n {
            word = append(word, '(')
            l++
            dfs(l, r)
            l--
            word = word[:len(word)-1]
        }
        if l > r {
            word = append(word, ')')
            r++
            dfs(l, r)
            r--
            word = word[:len(word)-1]
        }
    }
    dfs(0, 0)
    return res
}