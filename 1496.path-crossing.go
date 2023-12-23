/*
 * @lc app=leetcode id=1496 lang=golang
 *
 * [1496] Path Crossing
 */

package main

// @lc code=start

func isPathCrossing(path string) bool {
	pathMap := make(map[[2]int]bool)
	curr := [2]int{0, 0}
	pathMap[curr] = true

	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			curr[1]++
		case 'S':
			curr[1]--
		case 'E':
			curr[0]++
		case 'W':
			curr[0]--
		}

		if pathMap[curr] {
			return true
		} else {
			pathMap[curr] = true
		}
	}
	return false
}

// @lc code=end
