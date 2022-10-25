func findCircleNum(isConnected [][]int) int {
    visited := make([]bool, len(isConnected))

    var dfs func(i int)
    dfs = func(i int) {
        visited[i] = true
        for j := 0; j < len(isConnected[i]); j++ {
            if isConnected[i][j] == 1 && !visited[j] {
                dfs(j)
            }
        }
    }

    res := 0
    for i := 0; i < len(visited); i++ {
        if !visited[i] {
            res++
            dfs(i)
        }
    }
    return res
}