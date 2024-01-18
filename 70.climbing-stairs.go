/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

package main

// @lc code=start
func climbStairs(n int) int {
	memo := make([]int, n+1)
	memo[0] = 1
	memo[1] = 1

	var fib func(n int) int
	fib = func(n int) int {
		if memo[n] != 0 {
			return memo[n]
		}
		res := fib(n-1) + fib(n-2)
		memo[n] = res
		return res
	}

	return fib(n)
}

// @lc code=end
