func findIsland(grid [][]byte, x, y int){
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) || grid[x][y] != '1'{
        return
    }
    grid[x][y] = '2'
    findIsland(grid, x, y + 1)
    findIsland(grid, x, y - 1)
    findIsland(grid, x + 1, y)
    findIsland(grid, x - 1, y)
}

func numIslands(grid [][]byte) int {
    res := 0
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j]=='1'{
                res++
                findIsland(grid,i,j)
            }
        }
    }
    return res
}