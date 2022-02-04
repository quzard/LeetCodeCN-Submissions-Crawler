func rotate(matrix [][]int)  {
    l, r, t, d := 0, len(matrix[0]), 0, len(matrix)
    for l < r && t < d{
        for i:=l; i < r -1;i++{
            temp := matrix[t][i]
            matrix[t][i] = matrix[d - 1 - (i-l)][l]
            matrix[d - 1 - (i-l)][l]  = matrix[d - 1][r-1-(i-l)]
            matrix[d - 1][r-1-(i-l)] = matrix[t + (i -l)][r - 1]
            matrix[t + (i -l)][r - 1] = temp
        }
        l++
        r--
        t++
        d--
    }
}