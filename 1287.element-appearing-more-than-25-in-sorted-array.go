/*
 * @lc app=leetcode id=1287 lang=golang
 *
 * [1287] Element Appearing More Than 25% In Sorted Array
 */
package main

// @lc code=start
func findSpecialInteger(arr []int) int {
	target := len(arr) / 4

	for i, count := 1, 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			count++
		} else {
			count = 1
		}

		if count > target {
			return arr[i]
		}

	}

	return arr[0]
}

// @lc code=end
