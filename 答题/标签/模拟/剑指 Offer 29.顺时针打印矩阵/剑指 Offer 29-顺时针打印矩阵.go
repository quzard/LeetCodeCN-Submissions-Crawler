func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0{
        return []int{}
    }
    dirs := [][]int{
        []int{0,1}, // 左上角 0 
        []int{1,0}, // 右上角 1
        []int{0,-1}, // 右下角 2
        []int{-1,0}, // 左下角 3
    }
    res := []int{}
    count := len(matrix) * len(matrix[0])
    x, y := 0, 0
    up, down, left, right := 0, len(matrix) - 1, 0, len(matrix[0]) - 1
    dir := 0
    for{
        if count == 0{
            break
        }
        count--
        
        res = append(res, matrix[x][y])
        if x == up && y == right && dir == 0{ // 右上角 1
            up++
            dir = 1
        }else if x == up && y == left && dir == 3{  // 左上角 0
            left++
            dir = 0
        }else if x == down && y == left && dir == 2{ // 左下角 3
            down--
            dir = 3
        }else if x == down && y == right && dir == 1{ // 右下角 2
            right--
            dir = 2
        }
        x += dirs[dir][0]
        y += dirs[dir][1]
    }
    return res
}