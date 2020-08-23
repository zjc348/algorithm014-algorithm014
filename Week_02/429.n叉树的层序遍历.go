/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var res [][]int

func levelOrder(root *Node) [][]int {
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *Node, level int) {
	if root == nil {
		return
	}
	if len(res) == level {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	for _, v := range root.Children {
		dfs(v, level+1)
	}
}

// @lc code=end
