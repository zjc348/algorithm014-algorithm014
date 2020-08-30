/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
var ret [][]int

func permute(nums []int) [][]int {
	ret = [][]int{}
	used := make([]bool, len(nums))
	doPermute(nums, []int{}, used)
	return ret
}

func doPermute(nums []int, stack []int, used []bool) {
	if len(stack) == len(nums) {
		temp := make([]int, len(stack))
		copy(temp, stack)
		ret = append(ret, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !used[i] {
			stack = append(stack, nums[i])
			used[i] = true
			doPermute(nums, stack, used)
			stack = stack[:len(stack)-1]
			used[i] = false
		}
	}
}

// @lc code=end
