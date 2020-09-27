func ladderLength(beginWord string, endWord string, wordList []string) int {
    if idxOf(endWord, wordList) == -1{
        return 0
    }
    startQueue := []string{beginWord}
    endQueue := []string{endWord}
    used := make([]bool, len(wordList))
    step := 0
    for len(startQueue) > 0{
        step++
        ll := len(startQueue)
        for i := 0;i < ll;i++{
            for k := 0;k < len(endQueue);k++{
                if hasOneDiff(startQueue[i], endQueue[k]){
                    return step + 1
                }
            }
            for j, v := range wordList{
                if !used[j] && hasOneDiff(v, startQueue[i]){
                    startQueue = append(startQueue, v)
                    used[j] = true
                }
            }
        }
        startQueue = startQueue[ll:]
        if len(startQueue) > len(endQueue){
            startQueue, endQueue = endQueue, startQueue
        }
    }
    return 0
}

func idxOf(str string, bank []string)int{
    for i, v := range bank{
        if v == str{
            return i
        }
    }
    return -1
}

func hasOneDiff(one, two string)bool{
    count := 0
    for i := 0;i < len(one);i++{
        if one[i] != two[i]{
            count++
        }
        if count > 1{
            return false
        }
    }
    if count == 1{
        return true
    }
    return false
}