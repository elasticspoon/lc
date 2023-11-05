/*
 * @lc app=leetcode id=2265 lang=golang
 *
 * [2265] Count Nodes Equal to Average of Subtree
 */

package main

type TreeNode2265 struct {
	Left  *TreeNode2265
	Right *TreeNode2265
	Val   int
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

func averageOfSubtree(root *TreeNode2265) int {
	_, _, count := recCountStuff(root)
	return count
}

func recCountStuff(node *TreeNode2265) (int, int, int) {
	if node == nil {
		return 0, 0, 0
	}
	lsum, lnodes, lcount := recCountStuff(node.Left)
	rsum, rnodes, rcount := recCountStuff(node.Right)

	sum := lsum + rsum + node.Val
	nodes := lnodes + rnodes + 1
	count := lcount + rcount
	if sum/nodes == node.Val {
		count += 1
	}

	// fmt.Printf("node: %d, sum: %d, nodes: %d, count: %d\n", node.Val, sum, nodes, count)

	return sum, nodes, count
}

// @lc code=end
