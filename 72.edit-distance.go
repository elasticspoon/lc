/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

package main

// @lc code=start
func minDistance(word1 string, word2 string) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([][]int, len(word1)+1)

	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
		if i == 0 {
			for j := 0; j <= len(word2); j++ {
				dp[0][j] = j
			}
			continue
		}
		for j := 1; j <= len(word2); j++ {
			v := min(dp[i-1][j]+1, dp[i][j-1]+1)
			if word1[i-1] != word2[j-1] {
				v = min(v, dp[i-1][j-1]+1)
			} else {
				v = min(v, dp[i-1][j-1])
			}

			dp[i][j] = v
		}
	}

	return dp[len(word1)][len(word2)]
}

// @lc code=end
