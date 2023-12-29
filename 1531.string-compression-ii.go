/*
 * @lc app=leetcode id=1531 lang=golang
 *
* [1531] String Compression II
*/

package main

// @lc code=start
const MAX = 100

func getLengthOfOptimalCompression(s string, k int) int {
	n := len(s)
	dp := make([][]int, n)

	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = MAX + 1
		}
	}

	var dfs func(i, k int) int
	dfs = func(i, k int) int {
		if i+k >= n {
			return 0
		} else if k < 0 {
			return MAX + 1
		} else if dp[i][k] != MAX+1 {
			return dp[i][k]
		}

		res := dfs(i+1, k-1)
		diff, same, length := 0, 0, 0

		for j := i; j < n; j++ {
			if k < diff {
				break
			}
			if s[j] == s[i] {
				same++
				if same == 1 || same == 2 || same == 10 || same == 100 {
					length++
				}
			} else {
				diff++
			}
			res = min1531(res, length+dfs(j+1, k-diff))

		}

		dp[i][k] = res
		return res
	}

	return dfs(0, k)
}

func min1531(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
