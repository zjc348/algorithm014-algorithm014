import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
var ret [][]int

func permuteUnique(nums []int) [][]int {
	ret = [][]int{}
	used := make([]bool, len(nums))
	sort.Ints(nums)
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
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			stack = append(stack, nums[i])
			doPermute(nums, stack, used)
			stack = stack[:len(stack)-1]
			used[i] = false
		}
	}
}

// @lc code=end
