import "sort"

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	if len(g) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	for i := 0; i < len(s) && j < len(g); i++ {
		if s[i] >= g[j] {
			j++
		}
	}
	return j
}

// @lc code=end
