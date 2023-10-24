/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
 */

package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func (t *TreeNode) String() string {
	queue := []*TreeNode{t}
	result := ""
	v := 1
	count := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		count++
		if node == nil {
			result += "null "
			continue
		}
		result += fmt.Sprintf("%d ", node.Val)
		if count == v {
			result += "\n"
			v *= 2
			count = 0
		}
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}
	return result
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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := []int{}

	for len(queue) > 0 {
		currLevelSize := len(queue)
		minSize := -1 << 32

		for i := 0; i < currLevelSize; i++ {
			val := queue[0]
			queue = queue[1:]

			if val.Val > minSize {
				minSize = val.Val
			}

			if val.Left != nil {
				queue = append(queue, val.Left)
			}
			if val.Right != nil {
				queue = append(queue, val.Right)
			}
		}

		result = append(result, minSize)

	}
	return result
}

// @lc code=end
