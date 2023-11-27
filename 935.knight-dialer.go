/*
 * @lc app=leetcode id=935 lang=golang
 *
 * [935] Knight Dialer
 */
package main

// @lc code=start
func knightDialer(n int) int {
	const mod = 1e9 + 7
	jumps := [][]int{{4, 6}, {6, 8}, {7, 9}, {4, 8}, {0, 3, 9}, {}, {0, 1, 7}, {2, 6}, {1, 3}, {2, 4}}

	dp := make([]int, 10)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		temp := make([]int, 10)
		for i, v := range dp {
			for _, j := range jumps[i] {
				temp[j] += v
				temp[j] %= mod
			}
		}
		dp = temp
	}

	sum := 0
	for _, v := range dp {
		sum += v
		sum %= mod
	}

	return sum
}

// @lc code=end
