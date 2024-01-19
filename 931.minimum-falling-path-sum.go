/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */

package main

import (
	"math"
)

// @lc code=start
func minFallingPathSum(matrix [][]int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	memo := make([][]int, len(matrix))
	for i := range memo {
		memo[i] = make([]int, len(matrix[0]))
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if i == len(matrix)-1 {
				memo[i][j] = matrix[i][j]
				continue
			}
			best := memo[i+1][j]

			if j > 0 {
				best = min(best, memo[i+1][j-1])
			}
			if j < len(matrix[0])-1 {
				best = min(best, memo[i+1][j+1])
			}
			memo[i][j] = matrix[i][j] + best
		}
	}

	best := math.MaxInt32
	for _, v := range memo[0] {
		best = min(best, v)
	}
	return best
}

// @lc code=end
