/*
 * @lc app=leetcode id=2125 lang=golang
 *
 * [2125] Number of Laser Beams in a Bank
 */

package main

// @lc code=start
func numberOfBeams(bank []string) int {
	beams := 0
	currRow, prevRow := 0, 0

	for i := 0; i < len(bank); i++ {
		for j := 0; j < len(bank[i]); j++ {
			if bank[i][j] == '1' {
				currRow++
			}
		}

		if currRow > 0 {
			beams += currRow * prevRow
			prevRow = currRow
			currRow = 0
		}
	}

	return beams
}

// @lc code=end
