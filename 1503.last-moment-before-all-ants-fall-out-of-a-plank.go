/*
 * @lc app=leetcode id=1503 lang=golang
 *
 * [1503] Last Moment Before All Ants Fall Out of a Plank
 */

package main

// @lc code=start
func getLastMoment(n int, left []int, right []int) int {
	max := 0
	for _, l := range left {
		if l > max {
			max = l
		}
	}
	for _, r := range right {
		if n-r > max {
			max = n - r
		}
	}
	return max
}

// @lc code=end
