/*
 * @lc app=leetcode id=1347 lang=golang
 *
 * [1347] Minimum Number of Steps to Make Two Strings Anagram
 */

package main

// @lc code=start
func minSteps(s string, t string) int {
	sChars := make([]int, 26)
	tChars := make([]int, 26)

	for i := 0; i < len(s); i++ {
		sChars[s[i]%26]++
		tChars[t[i]%26]++
	}

	numDiffs := 0
	for i, v := range sChars {
		if v > tChars[i] {
			numDiffs += v - tChars[i]
		}
	}

	return numDiffs
}

// @lc code=end
