/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
    if root == nil{
        return []int{}
    }
    ret := []int{}
    stack := []*TreeNode{root}
    m := map[*TreeNode]int{root:1}
    for len(stack) > 0{
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        if v,ok := m[node];ok&&v==1{
            if node.Right != nil{
                m[node.Right] = 1
                stack = append(stack, node.Right)
            }
            m[node] = 2
            stack = append(stack, node)
            if node.Left != nil{
                m[node.Left] = 1
                stack = append(stack, node.Left)
            }
        }else{
            ret = append(ret, node.Val)
        }
    }
    return ret
}
// @lc code=end

