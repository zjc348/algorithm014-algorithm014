/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	len := len(nums)
	j := 0
	for i := 0; i < len; i++ {
		if nums[i] != nums[j] {
			nums[j+1] = nums[i]
			j++
		}
	}
	return j + 1
}

// @lc code=end
