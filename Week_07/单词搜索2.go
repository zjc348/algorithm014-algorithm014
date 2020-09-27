type TrieNode struct{
    word string
    next [26]*TrieNode
}
func findWords(board [][]byte, words []string) []string {
    root := &TrieNode{}
    for _, w := range words{
        node := root
        for _, v := range w{
            if node.next[v - 'a'] == nil{
                node.next[v - 'a'] = &TrieNode{}
            }
            node = node.next[v - 'a']
        }
        node.word = w
    }
    result := make([]string, 0)
    for  i := 0;i < len(board);i++{
        for j := 0;j < len(board[0]);j++{
            dfs(i, j, board, root, &result)
        }
    }
    return result
}

func dfs(i, j int, board [][] byte, node *TrieNode, result *[]string){
    if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
        return 
    }
    c := board[i][j]
    if c == '#' || node.next[c - 'a'] == nil{
        return
    }
    node = node.next[c - 'a']
    if node.word != ""{
        *result = append(*result, node.word)
        node.word = ""
    }
    board[i][j] = '#'
    dfs(i, j+1, board, node,result)
    dfs(i+1, j, board, node,result)
    dfs(i, j-1, board, node,result)
    dfs(i-1, j, board, node,result)
    board[i][j] = c
}