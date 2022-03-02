func exist(board [][]byte, word string) bool {
    res := false
    var dfs func(i, j, k int) 
    dfs = func(i, j, k int) {
        if k == len(word) {
            res = true
            return
        }
        if res || i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[k] || board[i][j] == ' ' {
            return
        }
        temp := board[i][j]
        board[i][j] = ' '
        dfs(i-1, j, k+1)
        dfs(i+1, j, k+1)
        dfs(i, j-1, k+1)
        dfs(i, j+1, k+1)
        board[i][j] = temp
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if res {
                return true
            }
            if board[i][j] == word[0] {
                dfs(i, j, 0)
            }
        }
    }
    return res
}