/*
 * @lc app=leetcode id=1758 lang=golang
 *
 * [1758] Minimum Changes To Make Alternating Binary String
 */

package main

// @lc code=start
func minOperations(s string) int {
	zeroStart, oneStart := '0', '1'
	diffZero, diffOne := 0, 0

	for _, c := range s {
		if c != zeroStart {
			diffZero++
		}
		if c != oneStart {
			diffOne++
		}

		zeroStart, oneStart = oneStart, zeroStart
	}

	if diffZero > diffOne {
		return diffOne
	}
	return diffZero
}

// @lc code=end
