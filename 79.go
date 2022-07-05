func dfs(board [][]byte, r int, c int, word string) bool {

    if r < 0 || r == len(board) || c < 0 || c == len(board[0]) {
        return false
    }

    if word[0] != board[r][c] {
        return false
    }

    if len(word) == 1 {
        return true
    }
    
    ch := board[r][c]
    board[r][c] = 0
    
    if dfs(board, r-1, c, word[1:]) {
        return true
    }
    if dfs(board, r+1, c, word[1:]) {
        return true
    }
    if dfs(board, r, c-1, word[1:]) {
        return true
    }
    if dfs(board, r, c+1, word[1:]) {
        return true
    }

    board[r][c] = ch
    
    return false
}

func exist(board [][]byte, word string) bool {
    
    for i := range(board) {
        for j := range(board[0]) {
            if dfs(board, i, j, word) {
                return true
            }
        }
    }

    return false
}
