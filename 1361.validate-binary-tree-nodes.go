/*
 * @lc app=leetcode id=1361 lang=golang
 *
 * [1361] Validate Binary Tree Nodes
 */
package main

// @lc code=start
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	for i := 0; i < n; i++ {
		visited := make(map[int]bool)
		stack := []int{}
		stack = append(stack, i)
		for len(stack) > 0 {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if _, ok := visited[current]; !ok {
				visited[current] = true
				if leftChild[current] != -1 {
					stack = append(stack, leftChild[current])
				}
				if rightChild[current] != -1 {
					stack = append(stack, rightChild[current])
				}
			} else {
				return false
			}
		}

		if len(visited) == n {
			return true
		}
	}

	return false
}

// @lc code=end
