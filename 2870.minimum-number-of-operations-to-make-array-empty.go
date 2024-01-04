/*
 * @lc app=leetcode id=2870 lang=golang
 *
 * [2870] Minimum Number of Operations to Make Array Empty
 */

package main

// @lc code=start
func minOperations(nums []int) int {
	digitCount := make(map[int]int)

	for _, num := range nums {
		digitCount[num]++
	}

	minOps := 0

	for _, count := range digitCount {
		if count == 1 {
			return -1
		}

		minOps += count / 3
		if count%3 != 0 {
			minOps++
		}

	}

	return minOps
}

// @lc code=end
