/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	temp := []int{}
	for i, _ := range matrix {
		for _, v := range matrix[i] {
			temp = append(temp, v)
		}
	}
	left, right := 0, len(temp)-1
	for left < right {
		mid := (left + right) >> 1
		if temp[mid] == target {
			return true
		}
		if temp[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left == right && temp[left] == target
}

// @lc code=end
