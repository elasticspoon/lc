/*
 * @lc app=leetcode id=1727 lang=golang
 *
 * [1727] Largest Submatrix With Rearrangements
 */

package main

import (
	"sort"
)

// @lc code=start
func largestSubmatrix(matrix [][]int) int {
	for col := 0; col < len(matrix[0]); col++ {
		for row := 1; row < len(matrix); row++ {
			if matrix[row][col] > 0 && matrix[row-1][col] > 0 {
				matrix[row][col] += matrix[row-1][col]
			}
		}
	}

	for _, row := range matrix {
		sort.Slice(row, func(i, j int) bool {
			return row[i] < row[j]
		})
	}

	best := 0
	for _, row := range matrix {
		for i, val := range row {
			potential := val * (len(row) - i)
			if potential > best {
				best = potential
			}
		}
	}

	// fmt.Println(matrix)

	return best
}

// @lc code=end
