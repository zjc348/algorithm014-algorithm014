/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	ret := [][]string{}
	m := map[[26]int][]string{}
	for i := 0; i < len(strs); i++ {
		temp := isana(strs[i])
		if _, ok := m[temp]; ok {
			m[temp] = append(m[temp], strs[i])
		} else {
			m[temp] = []string{strs[i]}
		}
	}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func isana(str string) [26]int {
	string_s := [26]int{}
	for i := 0; i < len(str); i++ {
		index := str[i] - 'a'
		string_s[index]++
	}
	return string_s
}

// @lc code=end
