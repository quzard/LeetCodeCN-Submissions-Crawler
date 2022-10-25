func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    i, j := len(matrix)-1, 0
    for i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[0]) {
        if matrix[i][j] == target {
            return true
        }
        if matrix[i][j] < target {
            j++
            continue
        }
        if matrix[i][j] > target {
            i--
            continue
        }
    }


    return false
}