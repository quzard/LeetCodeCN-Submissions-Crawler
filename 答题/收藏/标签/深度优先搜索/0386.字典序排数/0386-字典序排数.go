func lexicalOrder(n int) []int {
    res := make([]int,0,n)
    var dfs func(cur int) 
    dfs = func(cur int) {
        if cur > n {
            return
        }
        for i := 0; i <= 9; i++ {
            if cur + i > 0 && cur + i <= n{
                res = append(res, cur + i)
                dfs((cur + i) * 10)
            }
        }
    }
    dfs(0)
    return res
}