/*
 * @lc app=leetcode id=2742 lang=golang
 *
 * [2742] Painting the Walls
 */
package main

// @lc code=start

func paintWalls(cost []int, time []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = 1e9
	}

	for i := 0; i < n; i++ {
		for j := n; j > 0; j-- {
			curr := dp[j]
			potential := dp[max2742(j-time[i]-1, 0)] + cost[i]
			dp[j] = min2742(curr, potential)
		}
	}

	return dp[len(cost)]
}

func min2742(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max2742(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
