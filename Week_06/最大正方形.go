func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0{
        return 0
    }
    m, n, ret := len(matrix), len(matrix[0]), 0
    dp := make([][]int, m + 1)
    for i := 0;i <= m;i++{
        dp[i] = make([]int, n + 1)
    }
    max := func(a, b int)int{
        if a > b{
            return a
        }
        return b
    }
    min := func(a, b int)int{
        if a > b{
            return b
        }
        return a
    }
    for i := 1;i <= m;i++{
        for j := 1;j <= n;j++{
            if matrix[i-1][j-1] == '1'{
                dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
                ret = max(ret, dp[i][j] * dp[i][j])
            }
        }
    }
    return ret
}