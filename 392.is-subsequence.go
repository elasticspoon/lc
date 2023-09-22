/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

package main

// @lc code=start
func isSubsequence(s string, t string) bool {
	for subI, mainI, found := 0, 0, false; subI < len(s); subI++ {
		found = false
		for mainI < len(t) {
			if s[subI] == t[mainI] {
				mainI++
				found = true
				break
			} else {
				mainI++
			}
		}
		if !found {
			return false
		}
	}

	return true
}

// @lc code=end
