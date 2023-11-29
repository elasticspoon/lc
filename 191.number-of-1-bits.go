/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

package main

// @lc code=start
func hammingWeight(num uint32) int {
	count := 0

	for ; num > 0; num >>= 1 {
		count += int(num & 1)
	}

	return count
}

// @lc code=end
