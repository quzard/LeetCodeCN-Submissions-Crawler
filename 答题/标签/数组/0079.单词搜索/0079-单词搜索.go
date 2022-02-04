func exist(board [][]byte, word string) bool {
    if len(word) == 0{
        return true
    }
    res := false
    var dfs func(i, j, index int)
    dfs = func(i, j, index int){
        if index == len(word){
            res = true
            return
        }
        if res || i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || (board[i][j] == ' '){
            return
        }
        if board[i][j] == word[index]{
            temp := board[i][j]
            board[i][j] = ' '
            dfs(i - 1, j, index + 1)
            dfs(i + 1, j, index + 1)
            dfs(i, j - 1, index + 1)
            dfs(i, j + 1, index + 1)
            board[i][j] = temp
        }
    }
    for i := 0; i < len(board); i++{
        for j := 0; j < len(board[i]); j++{
            if board[i][j] == word[0]{
                temp := board[i][j]
                board[i][j] = ' '
                dfs(i - 1, j, 1)
                dfs(i + 1, j, 1)
                dfs(i, j - 1, 1)
                dfs(i, j + 1, 1)
                if res{
                    return true
                }
                board[i][j] = temp
            }
        }
    }
    return res
}


type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist1(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
