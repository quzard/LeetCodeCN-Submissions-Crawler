func solve1(board [][]byte) {
    visited := make([][]bool, len(board))
    for i := 0; i < len(board); i++ {
        visited[i] = make([]bool, len(board[0]))
    }

    var dfs func(i, j int, filled bool)
    dfs = func(i, j int, filled bool) {
        if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] == 'X' || visited[i][j] {
            return
        }
        visited[i][j] = true
        dfs(i+1, j, filled)
        dfs(i-1, j, filled)
        dfs(i, j-1, filled)
        dfs(i, j+1, filled)
        if filled {
            board[i][j] = 'X'
        }
    }
    for i := 0; i < len(board); i++ {
        j := 0
        if board[i][j] == 'O' && !visited[i][j] {
            dfs(i, j, false)
        }
        j = len(board[0]) - 1
        if board[i][j] == 'O' && !visited[i][j] {
            dfs(i, j, false)
        }
    }
    for j := 0; j < len(board[0]); j++ {
        i := 0
        if board[i][j] == 'O' && !visited[i][j] {
            dfs(i, j, false)
        }
        i = len(board) - 1
        if board[i][j] == 'O' && !visited[i][j] {
            dfs(i, j, false)
        }
    }

    for i := 1; i < len(board)-1; i++ {
        for j := 1; j < len(board[0])-1; j++ {
            if board[i][j] == 'O' && !visited[i][j] {
                dfs(i, j, true)
            }
        }
    }

}

func solve(board [][]byte) {
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != 'O' {
            return
        }
        board[i][j] = 'A'
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j-1)
        dfs(i, j+1)
    }
    for i := 0; i < len(board); i++ {
        j := 0
        if board[i][j] == 'O' {
            dfs(i, j)
        }
        j = len(board[0]) - 1
        if board[i][j] == 'O' {
            dfs(i, j)
        }
    }
    for j := 0; j < len(board[0]); j++ {
        i := 0
        if board[i][j] == 'O' {
            dfs(i, j)
        }
        i = len(board) - 1
        if board[i][j] == 'O' {
            dfs(i, j)
        }
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == 'A' {
                board[i][j] = 'O'
            } else if board[i][j] == 'O' {
                board[i][j] = 'X'
            }
        }
    }

}