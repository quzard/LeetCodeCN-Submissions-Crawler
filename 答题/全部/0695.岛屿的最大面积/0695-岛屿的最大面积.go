func maxAreaOfIsland(grid [][]int) int {
    res := 0
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
            return 0
        }
        grid[i][j] = 0
        sum := 1
        sum += dfs(i-1, j)
        sum += dfs(i+1, j)
        sum += dfs(i, j-1)
        sum += dfs(i, j+1)
        return sum
    }
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                res = max(res, dfs(i, j))
            }
        }
    }


    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}