/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	index1, index2, index3 := 0, 0, 0
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = min(dp[index1]*2, min(dp[index2]*3, dp[index3]*5))
		if dp[i] == dp[index1]*2 {
			index1++
		}
		if dp[i] == dp[index2]*3 {
			index2++
		}
		if dp[i] == dp[index3]*5 {
			index3++
		}
	}
	return dp[n-1]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
