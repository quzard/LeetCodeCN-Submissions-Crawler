func isValidSudoku1(board [][]byte) bool {
    n := 9
    for i:=0; i<n; i++{
        hashX := map[byte]bool{}
        hashY := map[byte]bool{}
        hashN := map[byte]bool{}
        for j:=0; j<n; j++{
            num_x := board[i][j]
            if num_x != '.' && hashX[num_x] == true{
                return false
            }else{
                hashX[num_x] = true
            }
            num_y := board[j][i]
            if num_y != '.' && hashY[num_y] == true{
                return false
            }else{
                hashY[num_y] = true
            }


            x, y := 0, 0
            if j > 5{
                x = j % 6
                y = 2
            }else if j > 2{
                x = j % 3
                y = 1
            }else{
                x = j
                y = 0
            }

            if i > 5{
                y += 6
                x += i % 6 * 3
            }else if i > 2{
                y += 3
                x += i % 3 * 3
            }else{
                x += i * 3
            }
            
            num_N := board[y][x]
            if num_N != '.' && hashN[num_N] == true{
                
                return false
            }else{
                hashN[num_N] = true
            }
        }
    }
    return true
}

func isValidSudoku(board [][]byte) bool {
	var line [9][9]int
	var column [9][9]int
	var cell [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - '0' -1
			k := i/3*3+j/3
			if line[i][num] != 0 || column[j][num] != 0 || cell[k][num] != 0 {
				return false
			}
			line[i][num] = 1
			column[j][num] = 1
			cell[k][num] = 1
		}
	}
	return true
}