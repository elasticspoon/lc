/*
 * @lc app=leetcode id=1611 lang=golang
 *
 * [1611] Minimum One Bit Operations to Make Integers Zero
 */
package main

// @lc code=start
func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}
	val, tog, sum := 1, 1, 0

	for ; n > 1; n, val = n/2, val*2 {
		if n&1 == 1 {
			sum += ((val * 2) - 1) * tog
			tog *= -1
		}
	}

	if sum > 0 {
		return val<<1 - sum - 1
	}
	return val<<1 + sum - 1
}

// @lc code=end
