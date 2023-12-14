/*
 * @lc app=leetcode id=2482 lang=golang
 *
 * [2482] Difference Between Ones and Zeros in Row and Column
 */

package main

// @lc code=start
func onesMinusZeros(grid [][]int) [][]int {
	colLen := len(grid[0])
	colsOnes := make([]int, len(grid[0]))

	rowLen := len(grid)
	rowsOnes := make([]int, len(grid))

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			colsOnes[j] += grid[i][j]
			rowsOnes[i] += grid[i][j]
		}
	}

	for i := 0; i < rowLen; i++ {
		for j := 0; j < colLen; j++ {
			grid[i][j] = 2*rowsOnes[i] - rowLen + 2*colsOnes[j] - colLen
		}
	}

	return grid
}

// @lc code=end
