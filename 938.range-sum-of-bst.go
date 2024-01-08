/*
 * @lc app=leetcode id=938 lang=golang
 *
 * [938] Range Sum of BST
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	sum := 0
	if low <= root.Val && root.Val <= high {
		sum += root.Val
	}

	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	if low < root.Val {
		sum += rangeSumBST(root.Left, low, high)
	}

	return sum
}

// @lc code=end
