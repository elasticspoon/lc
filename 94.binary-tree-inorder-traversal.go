/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
package main

type TreeNode94 struct {
	Val   int
	Left  *TreeNode94
	Right *TreeNode94
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode94) []int {
	if root == nil {
		return []int{}
	}
	res := append([]int{}, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

// @lc code=end
