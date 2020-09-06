/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for row, _ := range grid {
		for col, _ := range grid[row] {
			if grid[row][col] == '1' {
				dfs(grid, row, col)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] != '1' {
		return
	}
	grid[row][col] = '0'
	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)
}

// @lc code=end
