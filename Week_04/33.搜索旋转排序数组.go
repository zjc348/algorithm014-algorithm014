/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		a, b, c := 0, 0, 0
		if nums[0] > target {
			a = 1
		}
		if nums[0] > nums[mid] {
			b = 1
		}
		if target > nums[mid] {
			c = 1
		}
		if a^b^c != 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if left == right && nums[left] == target {
		return left
	}
	return -1
}

// @lc code=end
