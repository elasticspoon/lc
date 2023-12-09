/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
 */

package main

import (
	"strconv"
)

type TreeNode606 struct {
	Val   int
	Left  *TreeNode606
	Right *TreeNode606
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
func recTree2str(root *TreeNode606) string {
	if root == nil {
		return "()"
	}

	l := recTree2str(root.Left)
	r := recTree2str(root.Right)

	if r == "()" && l == "()" {
		return "(" + strconv.Itoa(root.Val) + ")"
	} else if r == "()" {
		return "(" + strconv.Itoa(root.Val) + l + ")"
	}

	return "(" + strconv.Itoa(root.Val) + l + r + ")"
}

func tree2str(root *TreeNode606) string {
	l := recTree2str(root.Left)
	r := recTree2str(root.Right)

	if r == "()" && l == "()" {
		return strconv.Itoa(root.Val)
	} else if r == "()" {
		return strconv.Itoa(root.Val) + l
	}

	return strconv.Itoa(root.Val) + l + r
}

// @lc code=end
