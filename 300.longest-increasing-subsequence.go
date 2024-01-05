/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

package main

// @lc code=start
func lengthOfLIS(nums []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(nums)
	memo := make([]int, n)

	var rec func(i int) int
	rec = func(i int) int {
		if i >= n {
			return 0
		} else if memo[i] > 0 {
			return memo[i]
		}

		best := 1
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] {
				best = max(best, 1+rec(j))
			}
		}
		memo[i] = best
		return best
	}

	best := 1
	for i := 0; i < n; i++ {
		best = max(best, rec(i))
	}
	return best
}

// @lc code=end
