/*
 * @lc app=leetcode id=501 lang=golang
 *
 * [501] Find Mode in Binary Search Tree
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (t *TreeNode) Insert(v int) {
	if v < t.Val {
		if t.Left == nil {
			t.Left = &TreeNode{Val: v}
		} else {
			t.Left.Insert(v)
		}
	} else {
		if t.Right == nil {
			t.Right = &TreeNode{Val: v}
		} else {
			t.Right.Insert(v)
		}
	}
}

// @lc code=start
func findMode(root *TreeNode) []int {
	queue := []*TreeNode{root}
	var modes []int
	nodeCount := make(map[int]int)
	maxCount := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		nodeCount[node.Val]++
		if nodeCount[node.Val] > maxCount {
			maxCount = nodeCount[node.Val]
			modes = []int{node.Val}
		} else if nodeCount[node.Val] == maxCount {
			modes = append(modes, node.Val)
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return modes
}

// @lc code=end
