func permutation(s string) []string {
    w := []byte(s)
    sort.Slice(w, func(a, b int) bool {
        return w[a] < w[b]
    })
    n := len(w)
    res := []string{}
    words := make([]byte, 0, n)
    selected := make([]bool, n)
    var dfs func(i int)
    dfs = func(i int) {
        if i == n {
            res = append(res, string(words))
            return
        }
        for j := 0; j < len(w); j++ {
            if selected[j] == true || j > 0 && !selected[j-1] && w[j] == w[j-1] {
                continue
            }
            selected[j] = true
            words = append(words, w[j])
            dfs(i + 1)
            words = words[:len(words)-1]
            selected[j] = false
        }
    }
    dfs(0)
    return res
}