/*
 * @lc app=leetcode id=1897 lang=golang
 *
 * [1897] Redistribute Characters to Make All Strings Equal
 */

package main

// @lc code=start
func makeEqual(words []string) bool {
	letters := make([]int, 26)
	for _, word := range words {
		for _, c := range word {
			letters[c-'a']++
		}
	}

	for _, letterCount := range letters {
		if letterCount%len(words) != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
