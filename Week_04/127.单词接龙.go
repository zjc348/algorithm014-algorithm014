/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func idxOf(str string, bank []string) int {
	for i := 0; i < len(bank); i++ {
		if bank[i] == str {
			return i
		}
	}
	return -1
}

func hasOneDiff(one string, two string) bool {
	count := 0
	for i := 0; i < len(one); i++ {
		if one[i] != two[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count == 1 {
		return true
	}
	return false
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if idxOf(endWord, wordList) == -1 {
		return 0
	}
	step := 0
	startQueue := []string{beginWord}
	used := make([]bool, len(wordList))
	for len(startQueue) > 0 {
		step++
		ll := len(startQueue)
		for i := 0; i < ll; i++ {
			if startQueue[i] == endWord {
				return step
			}
			for j, v := range wordList {
				if !used[j] && hasOneDiff(v, startQueue[i]) {
					startQueue = append(startQueue, v)
					used[j] = true
				}
			}
		}
		startQueue = startQueue[ll:]
	}
	return 0
}

// @lc code=end
