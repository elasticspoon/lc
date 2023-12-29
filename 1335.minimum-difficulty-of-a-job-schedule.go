/*
 * @lc app=leetcode id=1335 lang=golang
 *
 * [1335] Minimum Difficulty of a Job Schedule
 */

package main

// @lc code=start
func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, d)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	prefix := make([]int, n)
	prefix[0] = jobDifficulty[0]
	for i := 1; i < n; i++ {
		prefix[i] = max1335(prefix[i-1], jobDifficulty[i])
	}

	var dfs func(i, j int) int
	dfs = func(index, days int) int {
		if days == 0 {
			v := jobDifficulty[index]
			for i := index + 1; i < n; i++ {
				v = max1335(v, jobDifficulty[i])
			}
			return v
		}
		if dp[index][days] != -1 {
			return dp[index][days]
		}

		minV := 1<<31 - 1
		maxV := -1 << 31
		for i := index; i < n-days; i++ {
			maxV = max1335(maxV, jobDifficulty[i])
			minV = min1335(minV, maxV+dfs(i+1, days-1))
		}

		dp[index][days] = minV
		return minV
	}

	return dfs(0, d-1)
}

func min1335(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1335(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
