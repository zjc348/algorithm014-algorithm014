func leastInterval(tasks []byte, n int) int {
    char_s := [26]int{}
    max := 0
    count := 0
    for i := 0;i < len(tasks);i++{
        temp := tasks[i] - 'A'
        char_s[temp]++
        if max < char_s[temp]{
            max = char_s[temp]
            count = 1
        }else if char_s[temp] == max{
            count++
        }
    }
    if n == 0 || (max - 1) * (n + 1) + count < len(tasks){
        return len(tasks)
    }
    return (max - 1) * (n + 1) + count
}