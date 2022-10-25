func findNumberIn2DArray(matrix [][]int, target int) bool {
    i, j := len(matrix)- 1, 0
    for i >= 0 && i < len(matrix) && j>=0 && j < len(matrix[0]) {
        if matrix[i][j] == target {
            return true
        }
        if matrix[i][j] > target {
            i--
        } else if matrix[i][j] < target {
            j++
        }
    }
    return false
}