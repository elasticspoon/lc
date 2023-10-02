/*
 * @lc app=leetcode id=2038 lang=golang
 *
 * [2038] Remove Colored Pieces if Both Neighbors are the Same Color
 */
package main

// @lc code=start
func winnerOfGame(colors string) bool {
	var numA, numB, consA, consB int

	for _, v := range colors {
		if v == 'A' {
			consA++
			consB = 0
			if consA >= 3 {
				numA += 1
			}
		} else {
			consB++
			consA = 0
			if consB >= 3 {
				numB += 1
			}
		}
	}
	return numA-numB > 0
}

// @lc code=end
//
