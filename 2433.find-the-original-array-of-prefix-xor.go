/*
 * @lc app=leetcode id=2433 lang=golang
 *
 * [2433] Find The Original Array of Prefix Xor
 */

package main

// @lc code=start
func findArray(pref []int) []int {
	// result := make([]int, len(pref))
	// result[0] = pref[0]

	for i, prev := 0, 0; i < len(pref); i++ {
		// result[i] = pref[i] ^ pref[i-1]
		pref[i] ^= prev
		prev ^= pref[i]
	}

	return pref
}

// @lc code=end
