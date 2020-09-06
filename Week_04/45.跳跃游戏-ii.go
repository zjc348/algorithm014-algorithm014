/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	m, end, count := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		m = max(m, i+nums[i])
		if i == end {
			end = m
			count++
		}
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
