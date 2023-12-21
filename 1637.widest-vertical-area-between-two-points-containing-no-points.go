/*
 * @lc app=leetcode id=1637 lang=golang
 *
 * [1637] Widest Vertical Area Between Two Points Containing No Points
 */

package main

import (
	"sort"
)

// @lc code=start
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	max := 0
	for i := 1; i < len(points); i++ {
		if diff := points[i][0] - points[i-1][0]; diff > max {
			max = diff
		}
	}

	return max
}

// @lc code=end
