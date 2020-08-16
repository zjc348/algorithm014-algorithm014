/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	len := len(nums)
	k = k % len
	copy(nums, append(nums[len-k:], nums[:len-k]...))
}

// @lc code=end
