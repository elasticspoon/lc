/*
 * @lc app=leetcode id=1535 lang=golang
 *
 * [1535] Find the Winner of an Array Game
 */
package main

// @lc code=start
func getWinner(arr []int, k int) int {
	currentWinner := arr[0]
	winCount := 0

	for i := 1; i < len(arr); i++ {

		if arr[i] > currentWinner {
			currentWinner = arr[i]
			winCount = 1
		} else {
			winCount++
		}
		if winCount == k {
			return currentWinner
		}
	}
	return currentWinner
}

// @lc code=end
