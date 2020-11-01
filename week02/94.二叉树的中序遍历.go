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
var ans []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	inorderTraversal(root.Left)
	ans = append(ans, root.Val)
	inorderTraversal(root.Right)
	return ans
}

// @lc code=end

