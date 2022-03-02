func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    res := make([]int, 0, len(matrix) * len(matrix[0]))
    h, e, l, r := 0, len(matrix), 0, len(matrix[0])
    for len(res) < len(matrix) * len(matrix[0]){
        for i := l; i < r; i++ {
            res = append(res, matrix[h][i])
        }
        for i := h+1; i < e; i++ {
            res = append(res, matrix[i][r-1])
        }
        if l < r-1 && h < e-1 {
            for i := r-2; i >= l; i-- {
                res = append(res, matrix[e-1][i])
            }
            for i := e-2; i > h; i-- {
                res = append(res, matrix[i][l])
            }
        }
        h++
        l++
        e--
        r--
    }
    return res
}