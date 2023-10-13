/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

package main

// @lc code=start
func minCostClimbingStairs(costToUse []int) int {
	costToReach := make([]int, len(costToUse)+1)
	costToReach[0] = 0
	costToReach[1] = 0

	for i := 2; i < len(costToReach); i++ {
		a := costToReach[i-1] + costToUse[i-1]
		b := costToReach[i-2] + costToUse[i-2]

		costToReach[i] = min746(a, b)
	}

	return costToReach[len(costToReach)-1]
}

func min746(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end
