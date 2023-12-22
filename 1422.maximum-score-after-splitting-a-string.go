/*
 * @lc app=leetcode id=1422 lang=golang
 *
 * [1422] Maximum Score After Splitting a String
 */

package main

// @lc code=start
func maxScore(s string) int {
	zeros, ones := 0, 0
	for _, c := range s {
		if c == '0' {
			zeros++
		} else {
			ones++
		}
	}

	left, right, res := 0, ones, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			left++
		} else {
			right--
		}

		if sum := left + right; sum > res {
			res = sum
		}
	}

	return res
}

// @lc code=end
