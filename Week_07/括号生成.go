var ret []string
func generateParenthesis(n int) []string {
    ret = []string{}
    if n == 0{
        return []string{}
    }
    dfs(0, 0, n, "")
    return ret
}

func dfs(left, right, n int, s string){
    if left == n && right == n{
        ret = append(ret, s)
    }
    if left < n{
        dfs(left+1, right, n, s + "(")
    }
    if right < left{
        dfs(left, right+1, n, s + ")")
    }
}