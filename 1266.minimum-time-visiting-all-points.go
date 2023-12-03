/*
 * @lc app=leetcode id=1266 lang=golang
 *
 * [1266] Minimum Time Visiting All Points
 */
package main

// @lc code=start
func minTimeToVisitAllPoints(points [][]int) int {
	total := 0

	for i := 0; i < len(points)-1; i++ {
		total += minTime(points[i], points[i+1])
	}
	return total
}

func minTime(a []int, b []int) int {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	xDiff := abs(a[0] - b[0])
	yDiff := abs(a[1] - b[1])

	if xDiff > yDiff {
		return xDiff
	}
	return yDiff
}

// @lc code=end
