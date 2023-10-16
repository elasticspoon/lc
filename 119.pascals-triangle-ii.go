/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
package main

// @lc code=start
func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1

	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if j > 0 {
				result[j] += result[j-1]
			}
		}
	}

	return result
}

// @lc code=end
