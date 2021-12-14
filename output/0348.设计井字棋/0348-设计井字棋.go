type TicTacToe struct {
    x [][]int
    y [][]int
    z1 []int
    z2 []int
    xFailed []bool
    yFailed []bool
    z1Failed bool
    z2Failed bool
    n int
}


func Constructor(n int) TicTacToe {
    // n 为边长
    return TicTacToe{
        x: make([][]int, n),
        y: make([][]int, n),
        z1: []int{},
        z2: []int{},
        xFailed : make([]bool, n),
        yFailed : make([]bool, n),
        z1Failed: false,
        z2Failed: false,
        n: n,
    }
}


func (this *TicTacToe) Move(row int, col int, player int) int {
    if len(this.x[row]) == 0{
        this.x[row] = append(this.x[row], player)
    }else if this.xFailed[row] == false{
        if this.x[row][0] != player{
            this.xFailed[row] = true
        }else{
            this.x[row] = append(this.x[row], player)
        }
    }
    if len(this.x[row]) == this.n{
        return player
    }
    
    if len(this.y[col]) == 0{
        this.y[col] = append(this.y[col], player)
    }else if this.yFailed[col] == false{
        if this.y[col][0] != player{
            this.yFailed[col] = true
        }else{
            this.y[col] = append(this.y[col], player)
        }
    }
    if len(this.y[col]) == this.n{
        return player
    }

    if col == row{
        if len(this.z1) == 0{
            this.z1 = append(this.z1, player)
        }else if this.z1Failed == false{
            if this.z1[0] != player{
                this.z1Failed = true
            }else{
                this.z1 = append(this.z1, player)
            }
        }
    }
    if len(this.z1) == this.n{
        return player
    }

    if col + row == this.n - 1{
        if len(this.z2) == 0{
            this.z2 = append(this.z2, player)
        }else if this.z2Failed == false{
            if this.z2[0] != player{
                this.z2Failed = true
            }else{
                this.z2 = append(this.z2, player)
            }
        }
    }

    if len(this.z2) == this.n{
        return player
    }


    return 0
    
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */