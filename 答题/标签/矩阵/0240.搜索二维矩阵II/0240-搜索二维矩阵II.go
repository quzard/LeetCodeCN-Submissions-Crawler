func searchMatrix(matrix [][]int, target int) bool {
       x, y := 0, len(matrix[0]) - 1
       for x < len(matrix) && y >= 0{
           if matrix[x][y] == target{
               return true
           }
           if matrix[x][y] < target{
               x++
           }else if matrix[x][y] > target{
               y--
           }
       }
       return false
}