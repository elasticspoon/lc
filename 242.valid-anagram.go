/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

package main

// @lc code=start
func isAnagram(s string, t string) bool {
	sLetters := [26]int{}
	for i := 0; i < len(s); i++ {
		sLetters[s[i]%26]++
	}

	for i := 0; i < len(t); i++ {
		sLetters[t[i]%26]--
		if sLetters[t[i]%26] < 0 {
			return false
		}
	}

	for _, v := range sLetters {
		if v != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
