/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := []int{}
	stack := []*TreeNode{root}
	m := map[*TreeNode]int{root: 1}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if v, ok := m[node]; ok && v == 1 {
			if node.Right != nil {
				m[node.Right] = 1
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				m[node.Left] = 1
				stack = append(stack, node.Left)
			}
			m[node] = 2
			stack = append(stack, node)
		} else {
			ret = append(ret, node.Val)
		}
	}
	return ret
}

// @lc code=end
