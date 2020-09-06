/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}
	if bills[0] > 5 {
		return false
	}
	five, ten := 0, 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			five++
		} else if bills[i] == 10 {
			ten++
			five--
			if five < 0 {
				return false
			}
		} else {
			if ten > 0 {
				ten--
				five--
				if ten < 0 {
					return false
				}
				if five < 0 {
					return false
				}
			} else {
				five -= 3
				if five < 0 {
					return false
				}
			}
		}
	}
	return true
}

// @lc code=end
