func countSubstrings(s string) int {
    check := func(i int, j int) (cnt int) {
        cnt = 0
        for i >= 0 && j < len(s) && s[i] == s[j] {
            cnt += 1
            i -= 1
            j += 1
        }
        return
    }

    ans := 0
    for i := 0; i < len(s); i++ {
        ans += check(i, i)
        ans += check(i, i + 1)
    }
    return ans
}