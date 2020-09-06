/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
    if board[click[0]][click[1]] == 'M'{
        board[click[0]][click[1]] = 'X'
    }else{
        dfs(board, click[0], click[1])
    }
    return board
}

func dfs(board [][] byte, row, col int){
    if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]){
        return
    }
    if board[row][col] == 'B'{
        return
    }
    var touch func(board [][] byte, row, col int)int
    touch = func(board [][]byte, row, col int)int{
        if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]){
            return 0
        }
        if board[row][col] == 'M'{
            return 1
        }
        return 0
    }
    
    count := 0
    count += touch(board, row-1, col)
    count += touch(board, row+1, col)
    count += touch(board, row, col-1)
    count += touch(board, row, col+1)
    count += touch(board, row - 1, col-1)
    count += touch(board, row-1, col+1)
    count += touch(board, row+1, col-1)
    count += touch(board, row+1, col+1)
    if count > 0{
        board[row][col] = byte(count + '0')
    }else{
        board[row][col] = 'B'
        dfs(board, row-1, col)
        dfs(board, row+1, col)
        dfs(board, row, col-1)
        dfs(board, row, col+1)
        dfs(board, row - 1, col-1)
        dfs(board, row-1, col+1)
        dfs(board, row+1, col-1)
        dfs(board, row+1, col+1)
    }
}
// @lc code=end

