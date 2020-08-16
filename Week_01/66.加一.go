/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	len := len(digits)
	for i := len - 1; i >= 0; i-- {
		if i == len-1 {
			digits[len-1]++
		}
		if digits[i] >= 10 {
			digits[i] = digits[i] - 10
			if i != 0 {
				digits[i-1]++
			} else {
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}

// @lc code=end
