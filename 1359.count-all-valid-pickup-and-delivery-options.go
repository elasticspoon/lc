/*
 * @lc app=leetcode id=1359 lang=golang
 *
 * [1359] Count All Valid Pickup and Delivery Options
 */

package main

// import "fmt"
//
// func debugDP(dp [][]int) {
// 	fmt.Println("")
// 	for d := len(dp) - 1; d >= 0; d-- {
// 		line := ""
// 		for p := 0; p < len(dp); p++ {
// 			if line == "" {
// 				line = fmt.Sprintf("%d", dp[p][d])
// 			} else {
// 				line = fmt.Sprintf("%s, %d", line, dp[p][d])
// 			}
// 		}
// 		fmt.Printf("[ %s ]\n", line)
// 	}
// }

func countOrdersDP(n int) int {
	dp := make([][]int, n+1)
	const MOD int = 1_000_000_007

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][1] = 1

	for d := 1; d <= n; d++ {
		for p := 0; p <= d; p++ {
			var num int

			if p > 0 {
				num += (p * dp[p-1][d])
			}
			if d > 0 {
				num += ((d - p) * dp[p][d-1])
			}

			dp[p][d] += num % MOD
		}
	}

	return dp[n][n]
}

// @lc code=start

func countOrders(n int) int {
	const MOD int = 1_000_000_007
	res := 1

	for i := 1; i <= n; i++ {
		res = (res * (i*2 - 1) * i) % MOD
	}

	return res
}

// @lc code=end
