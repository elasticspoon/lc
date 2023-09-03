/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 *
 * https://leetcode.com/problems/unique-paths/description/
 *
 * algorithms
 * Medium (63.02%)
 * Likes:    15443
 * Dislikes: 410
 * Total Accepted:    1.6M
 * Total Submissions: 2.5M
 * Testcase Example:  '3\n7'
 *
 * There is a robot on an m x n grid. The robot is initially located at the
 * top-left corner (i.e., grid[0][0]). The robot tries to move to the
 * bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move
 * either down or right at any point in time.
 *
 * Given the two integers m and n, return the number of possible unique paths
 * that the robot can take to reach the bottom-right corner.
 *
 * The test cases are generated so that the answer will be less than or equal
 * to 2 * 10^9.
 *
 *
 * Example 1:
 *
 *
 * Input: m = 3, n = 7
 * Output: 28
 *
 *
 * Example 2:
 *
 *
 * Input: m = 3, n = 2
 * Output: 3
 * Explanation: From the top-left corner, there are a total of 3 ways to reach
 * the bottom-right corner:
 * 1. Right -> Down -> Down
 * 2. Down -> Down -> Right
 * 3. Down -> Right -> Down
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= m, n <= 100
 *
 *
 */

package main

import (
	"math/big"
)

func lessVarsUniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	num, denom := big.NewInt(1), big.NewInt(1)

	if m > n {
		for i := m; i <= (m + n - 2); i++ {
			num.Mul(num, big.NewInt(int64(i)))
		}

		for i := 1; i <= n-1; i++ {
			denom.Mul(denom, big.NewInt(int64(i)))
		}
	} else {
		for i := n; i <= (m + n - 2); i++ {
			num.Mul(num, big.NewInt(int64(i)))
		}

		for i := 1; i <= m-1; i++ {
			denom.Mul(denom, big.NewInt(int64(i)))
		}
	}

	return int(num.Div(num, denom).Int64())
}

func myUniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	var a, b, c int
	num, denom := big.NewInt(1), big.NewInt(1)

	if m > n {
		a = m
		b = n - 1
		c = m + n - 2
	} else {
		a = n
		b = m - 1
		c = m + n - 2
	}

	for i := a; i <= c; i++ {
		num.Mul(num, big.NewInt(int64(i)))
	}

	for i := 1; i <= b; i++ {
		denom.Mul(denom, big.NewInt(int64(i)))
	}

	return int(num.Div(num, denom).Int64())
}

// @lc code=start
// I do not understand why this works.
func uniquePaths(m int, n int) int {
	ans := 1

	for i, j := (m + n - 2), 1; i >= maxV(m, n); i, j = i-1, j+1 {
		// fmt.Printf("%v * %v / %v\n", ans, i, j)
		ans = (ans * i) / j
	}

	return ans
}

func maxV(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
