/*
 * @lc app=leetcode id=2849 lang=golang
 *
 * [2849] Determine if a Cell Is Reachable at a Given Time
 */
package main

// @lc code=start
func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	if t == 1 && (sx == fx && sy == fy) {
		return false
	}
	return abs2849(sx-fx) <= t && abs2849(sy-fy) <= t
	// (t-abs2849(sx-fx)-abs2849(sy-fy))%2 == 0
}

func abs2849(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
