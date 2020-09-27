type Trie struct {
    isEnd bool
    next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node := this
    for _, v := range word{
        if node.next[v - 'a'] == nil{
            node.next[v - 'a'] = &Trie{}
        }
        node = node.next[v - 'a']
    }
    node.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this
    for _, v := range word{
        if node = node.next[v - 'a'];node == nil{
            return false
        }
    }
    return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this
    for _, v := range prefix{
        if node = node.next[v - 'a'];node == nil{
            return false
        }
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */