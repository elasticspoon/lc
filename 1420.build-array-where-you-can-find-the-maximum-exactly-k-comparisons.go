/*
 * @lc app=leetcode id=1420 lang=golang
 *
 * [1420] Build Array Where You Can Find The Maximum Exactly K Comparisons
 */

package main

import "fmt"

// @lc code=start
func numOfArrays(n int, m int, k int) int {
	const mod = int(1e9 + 7)

	dp := make([][]int, m+1)
	prefix := make([][]int, m+1)
	prevDp := make([][]int, m+1)
	prevPrefix := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, k+1)
		prefix[i] = make([]int, k+1)
		prevDp[i] = make([]int, k+1)
		prevPrefix[i] = make([]int, k+1)
	}

	for i := 1; i <= m; i++ {
		prevDp[i][1] = 1
		prevPrefix[i][1] = i
	}

	for length := 2; length <= n; length++ {
		for maxN := 1; maxN <= m; maxN++ {
			for maxK := 1; maxK <= k; maxK++ {
				dp[maxN][maxK] = (maxN * prevDp[maxN][maxK]) % mod

				if maxN > 1 && maxK > 1 {
					dp[maxN][maxK] = (dp[maxN][maxK] + prevPrefix[maxN-1][maxK-1]) % mod
				}

				prefix[maxN][maxK] = (prefix[maxN-1][maxK] + dp[maxN][maxK]) % mod
			}
		}
		for i := 1; i <= m; i++ {
			copy(prevDp[i], dp[i])
			copy(prevPrefix[i], prefix[i])
		}
	}

	return prefix[m][k]
}

// @lc code=end
func main1420() {
	fmt.Println(numOfArrays(2, 3, 1) == 6)
	fmt.Println(numOfArrays(5, 2, 3) == 0)
	fmt.Println(numOfArrays(9, 1, 1) == 1)
	fmt.Println(numOfArrays(50, 100, 25) == 34549172)
}
