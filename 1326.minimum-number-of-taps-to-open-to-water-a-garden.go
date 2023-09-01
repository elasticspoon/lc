/*
 * @lc app=leetcode id=1326 lang=golang
 *
 * [1326] Minimum Number of Taps to Open to Water a Garden
 *
 * https://leetcode.com/problems/minimum-number-of-taps-to-open-to-water-a-garden/description/
 *
 * algorithms
 * Hard (47.34%)
 * Likes:    2985
 * Dislikes: 161
 * Total Accepted:    109.1K
 * Total Submissions: 212K
 * Testcase Example:  '5\n[3,4,1,1,0,0]'
 *
 * There is a one-dimensional garden on the x-axis. The garden starts at the
 * point 0 and ends at the point n. (i.e The length of the garden is n).
 *
 * There are n + 1 taps located at points [0, 1, ..., n] in the garden.
 *
 * Given an integer n and an integer array ranges of length n + 1 where
 * ranges[i] (0-indexed) means the i-th tap can water the area [i - ranges[i],
 * i + ranges[i]] if it was open.
 *
 * Return the minimum number of taps that should be open to water the whole
 * garden, If the garden cannot be watered return -1.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 5, ranges = [3,4,1,1,0,0]
 * Output: 1
 * Explanation: The tap at point 0 can cover the interval [-3,3]
 * The tap at point 1 can cover the interval [-3,5]
 * The tap at point 2 can cover the interval [1,3]
 * The tap at point 3 can cover the interval [2,4]
 * The tap at point 4 can cover the interval [4,4]
 * The tap at point 5 can cover the interval [5,5]
 * Opening Only the second tap will water the whole garden [0,5]
 *
 *
 * Example 2:
 *
 *
 * Input: n = 3, ranges = [0,0,0,0]
 * Output: -1
 * Explanation: Even if you activate all the four taps you cannot water the
 * whole garden.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= n <= 10^4
 * ranges.length == n + 1
 * 0 <= ranges[i] <= 100
 *
 *
 */

package main

import (
	"sort"
)

// @lc code=start
func minTaps(n int, ranges []int) int {
	var intervals [][2]int
	for i, v := range ranges {
		if v == 0 {
			continue
		}
		intervals = append(intervals, [2]int{i - v, i + v})
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var taps int
	var gSpot int

	for gSpot < n {
		validT := [][2]int{}
		for _, v := range intervals {
			if contains(gSpot, v) && v[1] > gSpot {
				validT = append(validT, v)
			}
		}
		if len(validT) == 0 {
			return -1
		}
		sort.Slice(validT, func(i, j int) bool {
			return validT[i][1] > validT[j][1]
		})
		taps++
		gSpot = validT[0][1]
	}

	return taps
}

func contains(v int, r [2]int) bool {
	return r[0] <= v && v <= r[1]
}

// @lc code=end
