/*
 * @lc app=leetcode id=1582 lang=golang
 *
 * [1582] Special Positions in a Binary Matrix
 */

package main

// @lc code=start
func numSpecial(mat [][]int) int {
	sumCols := make([]int, len(mat[0]))
	sumRows := make([]int, len(mat))

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				sumRows[i]++
				sumCols[j]++
			}
		}
	}
	res := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 && sumRows[i] == 1 && sumCols[j] == 1 {
				res++
			}
		}
	}

	return res
}

// @lc code=end
