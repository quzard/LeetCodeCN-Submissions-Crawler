func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    n := len(matrix) * len(matrix[0])
    res := make([]int, 0, n)
    cnt := 0
    top, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
    for cnt < n {
        for i := left; i <= right && cnt < n; i++ {
            res = append(res, matrix[top][i])
            cnt++
        }
        for i := top+1; i <= down && cnt < n; i++ {
            res = append(res, matrix[i][right])
            cnt++
        }
        for i := right-1; i >= left && cnt < n; i-- {
            res = append(res, matrix[down][i])
            cnt++
        }
        for i := down-1; i >= top+1 && cnt < n; i-- {
            res = append(res, matrix[i][left])
            cnt++
        }
        top++
        down--
        left++
        right--
    } 
    return res
}