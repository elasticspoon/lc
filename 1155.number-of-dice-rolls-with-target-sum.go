/*
 * @lc app=leetcode id=1155 lang=golang
 *
 * [1155] Number of Dice Rolls With Target Sum
 */

package main

// @lc code=start
const mod1155 = 1e9 + 7

func numRollsToTarget(n int, k int, target int) int {
	memo := make(map[[2]int]int)
	return recNumRollsToTarget(n, k, target, memo)
}

func recNumRollsToTarget(n int, k int, target int, memo map[[2]int]int) int {
	v := [2]int{n, target}

	if ways, ok := memo[v]; ok {
		return ways
	}

	if target < n || target > n*k {
		memo[v] = 0
		return 0
	}

	if target == 0 {
		memo[v] = 1
		return 1
	}

	ways := 0

	for i := 1; i <= k; i++ {
		ways += recNumRollsToTarget(n-1, k, target-i, memo)
		ways %= mod1155
	}

	memo[v] = ways
	return ways
}

// @lc code=end
