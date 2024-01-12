/*
 * @lc app=leetcode id=1704 lang=golang
 *
 * [1704] Determine if String Halves Are Alike
 */

package main

// @lc code=start
func halvesAreAlike(s string) bool {
	numVowels := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < len(s)/2 {
				numVowels++
			} else {
				numVowels--
			}
		}

		if numVowels < 0 {
			return false
		}
	}
	return numVowels == 0
}

// @lc code=end
