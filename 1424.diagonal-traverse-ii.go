/*
 * @lc app=leetcode id=1424 lang=golang
 *
 * [1424] Diagonal Traverse II
 */

package main

// @lc code=start
func findDiagonalOrder(nums [][]int) []int {
	diagonals := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			diagonals[i+j] = append(diagonals[i+j], nums[i][j])
		}
	}

	result := []int{}
	for i := 0; i < len(diagonals); i++ {
		v := diagonals[i]
		for j := len(v) - 1; j >= 0; j-- {
			result = append(result, v[j])
		}
	}

	return result
}

// @lc code=end
