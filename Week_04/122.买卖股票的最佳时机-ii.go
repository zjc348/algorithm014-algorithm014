/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	ret := 0
	for i := 0; i < len(prices)-1; i++ {
		temp := prices[i+1] - prices[i]
		if temp > 0 {
			ret += temp
		}
	}
	return ret
}

// @lc code=end
