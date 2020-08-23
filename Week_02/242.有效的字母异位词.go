/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	string_s := [26]int{}
	string_t := [26]int{}
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		string_s[index]++
	}
	for i := 0; i < len(t); i++ {
		index := t[i] - 'a'
		string_t[index]++
	}
	return string_s == string_t
}

// @lc code=end
