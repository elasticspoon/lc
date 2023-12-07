/*
 * @lc app=leetcode id=1903 lang=golang
 *
 * [1903] Largest Odd Number in String
 */
package main

// @lc code=start
func largestOddNumber(num string) string {
	if len(num) == 0 {
		return ""
	}

	for i := len(num) - 1; i >= 0; i-- {
		t := num[i]
		if t == '1' || t == '3' || t == '5' || t == '7' || t == '9' {
			return num[:i+1]
		}
	}

	return ""
}

// @lc code=end
