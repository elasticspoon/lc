/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

package main

// @lc code=start
func generate(numRows int) [][]int {
	list := make([][]int, numRows)
	list[0] = []int{1}

	for i := 1; i < numRows; i++ {
		list[i] = make([]int, i+1)
		list[i][0] = 1
		list[i][i] = 1

		for j := 1; j < i; j++ {
			list[i][j] = list[i-1][j] + list[i-1][j-1]
		}
	}

	return list
}

// @lc code=end
