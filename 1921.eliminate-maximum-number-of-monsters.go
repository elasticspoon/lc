/*
 * @lc app=leetcode id=1921 lang=golang
 *
 * [1921] Eliminate Maximum Number of Monsters
 */
package main

import (
	"math"
	"sort"
)

// @lc code=start
func eliminateMaximum(dist []int, speed []int) int {
	time := make([]int, len(dist))
	for i := 0; i < len(dist); i++ {
		time[i] = int(math.Ceil(float64(dist[i]) / float64(speed[i])))
	}

	sort.Slice(time, func(i, j int) bool {
		return time[i] < time[j]
	})

	for i := 1; i < len(time); i++ {
		if time[i] < i+1 {
			return i
		}
	}

	return len(time)
}

// @lc code=end
