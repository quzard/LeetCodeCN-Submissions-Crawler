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

func numIslands1(grid [][]byte) int {
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

func find(fa []int, i int) int {
    if fa[i] == i {
        return i
    }
    fa[i] = find(fa, fa[i])
    return fa[i]
}

func union(fa []int, i, j int, cnt *int) {
    x, y := find(fa, i), find(fa, j)
    if x != y {
        fa[y] = x
        *cnt = *cnt - 1
    }
}

// 并查集
func numIslands(grid [][]byte) int {
    x, y := len(grid), len(grid[0])
    fa := make([]int, x * y)
    cnt := 0
    for i := 0; i < x; i++ {
        for j := 0; j < y; j++{
            if grid[i][j] == '1' {
                cnt++
            }
            fa[i*y+j] = i*y+j
        }
    }
    dirs := [][]int{
        []int{1, 0},
        []int{-1, 0},
        []int{0, 1},
        []int{0, -1},
    }
    for i := 0; i < x; i++ {
        for j := 0; j < y; j++ {
            if grid[i][j] == '1' {
                for _, dir := range dirs {
                    grid[i][j] = '0'
                    ii := i + dir[0]
                    jj := j + dir[1]
                    if ii >= 0 && ii < x && jj >= 0 && jj < y && grid[ii][jj] == '1'{
                        union(fa, i*y+j, ii*y+jj, &cnt)
                    }
                }
            }
        }
    }
    return cnt
}