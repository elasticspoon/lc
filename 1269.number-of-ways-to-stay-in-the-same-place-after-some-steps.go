/*
 * @lc app=leetcode id=1269 lang=golang
 *
 * [1269] Number of Ways to Stay in the Same Place After Some Steps
 */
package main

// @lc code=start
func numWays(steps int, arrLen int) int {
	const mod = 1e9 + 7
	dp := make([][]int, steps)
	var maxDist int

	if steps/2+1 > arrLen {
		maxDist = arrLen
	} else {
		maxDist = steps/2 + 1
	}

	for i := range dp {
		dp[i] = make([]int, maxDist)
	}

	dp[0][0] = 1
	if maxDist > 1 {
		dp[0][1] = 1
	}

	for s := 1; s < len(dp); s++ {
		for l := 0; l < maxDist; l++ {
			temp := dp[s-1][l]
			if l > 0 {
				temp += dp[s-1][l-1]
			}
			if l < maxDist-1 {
				temp += dp[s-1][l+1]
			}
			dp[s][l] = temp % mod
		}
	}

	return dp[len(dp)-1][0]
}

// @lc code=end
