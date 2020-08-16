/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	len := len(height)
	res := 0
	for i := 1; i < len-1; i++ {
		leftMax := 0
		for j := 0; j <= i; j++ {
			leftMax = max(leftMax, height[j])
		}
		rightMax := 0
		for k := i; k < len; k++ {
			rightMax = max(rightMax, height[k])
		}
		res += min(leftMax, rightMax) - height[i]
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
