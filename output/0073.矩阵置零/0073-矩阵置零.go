func setZeroes1(matrix [][]int)  {
    selectedX := []int{}
    selectedY := []int{}
    for x := 0; x < len(matrix); x++{
        for y := 0; y < len(matrix[0]); y++{
            if matrix[x][y] == 0{
                selectedX = append(selectedX, x)
                selectedY = append(selectedY, y)
            }
        }
    }
    for i := 0; i < len(selectedX); i++{
        x := selectedX[i]
        for y := 0; y < len(matrix[0]); y++{
            matrix[x][y] = 0
        }
    }

    for i := 0; i < len(selectedY); i++{
        y := selectedY[i]
        for x := 0; x < len(matrix); x++{
            matrix[x][y] = 0
        }
    }
}

func setZeroes(matrix [][]int) {
    n, m := len(matrix), len(matrix[0])
    col0 := false
    for _, r := range matrix {
        if r[0] == 0 {
            col0 = true
        }
        for j := 1; j < m; j++ {
            if r[j] == 0 {
                r[0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := n - 1; i >= 0; i-- {
        for j := 1; j < m; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
        if col0 {
            matrix[i][0] = 0
        }
    }
}
