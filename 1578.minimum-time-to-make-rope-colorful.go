/*
 * @lc app=leetcode id=1578 lang=golang
 *
 * [1578] Minimum Time to Make Rope Colorful
 */

package main

// @lc code=start
func minCost(colors string, neededTime []int) int {
	biggestRemoval := neededTime[0]
	removalCost := neededTime[0]

	for i := 1; i < len(colors); i++ {
		if colors[i] != colors[i-1] {
			removalCost -= biggestRemoval
			biggestRemoval = neededTime[i]
		} else if neededTime[i] > biggestRemoval {
			biggestRemoval = neededTime[i]
		}
		removalCost += neededTime[i]
	}

	removalCost -= biggestRemoval

	return removalCost
}

// @lc code=end
