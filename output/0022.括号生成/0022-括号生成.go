var res []string

func dfs(s string, left, right , n int){
    if left == right && left == n{
        res = append(res, s)
    }
    if left < n{
        dfs(s + "(", left + 1, right, n)
    }
    if left > right{
        dfs(s + ")", left, right + 1, n)
    }
}

func generateParenthesis(n int) []string {
    res = []string{}
    dfs("(", 1, 0, n) 
    return res
}