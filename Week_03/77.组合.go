/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
var ret [][]int

func combine(n int, k int) [][]int {
	ret = [][]int{}
	doCombine(n, k, 1, []int{})
	return ret
}

func doCombine(n int, k int, begin int, stack []int) {
	if len(stack) == k {
		temp := make([]int, len(stack))
		copy(temp, stack)
		ret = append(ret, temp)
	}
	for i := begin; i <= n; i++ {
		stack = append(stack, i)
		doCombine(n, k, i+1, stack)
		stack = stack[:len(stack)-1]
	}
}

// @lc code=end
