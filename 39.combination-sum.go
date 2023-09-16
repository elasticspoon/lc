/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

package main

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	dp := make([][][]int, target+1)
	for i := range dp {
		dp[i] = [][]int{}
	}

	for i := 1; i <= target; i++ {
		for _, candidate := range candidates {
			if candidate <= i {
				dp[i] = append(dp[i], dp[i-candidate]...)
				for j := 0; j < len(dp[i]); j++ {
					dp[i][j] = append(dp[i][j], candidate)
				}

			}
		}
	}

	return dp[target]
}

// @lc code=end
