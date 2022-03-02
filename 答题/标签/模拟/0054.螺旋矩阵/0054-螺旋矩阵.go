func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    res := make([]int, 0, len(matrix) * len(matrix[0]))
    cnt := len(matrix) * len(matrix[0])
    top, down, left, right := 0, len(matrix) - 1, 0, len(matrix[0]) - 1
    for cnt > 0 {
        for i := left; i <= right; i++ {
            res = append(res, matrix[top][i])
            cnt--
        }
        for i := top + 1; i <= down; i++ {
            res = append(res, matrix[i][right])
            cnt--
        }
        if top < down && left < right {
            for i := right - 1; i >= left; i-- {
                res = append(res, matrix[down][i])
                cnt--
            }
            for i := down - 1; i > top; i-- {
                res = append(res, matrix[i][left])
                cnt--
            }
        }
        top++
        left++
        right--
        down--
    }
    return res
}