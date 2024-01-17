/*
 * @lc app=leetcode id=1207 lang=golang
 *
 * [1207] Unique Number of Occurrences
 */
package main

// @lc code=start
func uniqueOccurrences(arr []int) bool {
	numOcc := make(map[int]int)
	for _, v := range arr {
		numOcc[v]++
	}

	uniqOcc := make(map[int]bool)

	for _, v := range numOcc {
		if _, ok := uniqOcc[v]; ok {
			return false
		} else {
			uniqOcc[v] = true
		}
	}
	return true
}

// @lc code=end
