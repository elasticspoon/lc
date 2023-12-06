/*
 * @lc app=leetcode id=1716 lang=golang
 *
 * [1716] Calculate Money in Leetcode Bank
 */
package main

// @lc code=start
func totalMoney(n int) int {
	total := 0
	for i := 0; i < n; i++ {
		total += (i % 7) + 1 + (i / 7)
	}
	return total
}

// @lc code=end
