func exist(board [][]byte, word string) bool {
    if len(word) == 0 || len(board) == 0 || len(board[0]) == 0 || len(board)*len(board[0]) < len(word) {
        return false
    }

    var dfs func(i, j int, k int) bool
    dfs = func(i, j int, k int) bool {
        if k == len(word) {
            return true
        }
        if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] == ' '{
            return false
        }
        if board[i][j] != word[k] {
            return false
        }
        temp := board[i][j]
        board[i][j] = ' '
        if dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1) {
            board[i][j] = temp
            return true
        } 
        board[i][j] = temp
        return false

    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}