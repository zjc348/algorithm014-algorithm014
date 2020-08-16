/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if val, ok := m[target-v]; ok {
			return []int{val, i}
		}
		m[v] = i
	}
	return []int{}
}

// @lc code=end
