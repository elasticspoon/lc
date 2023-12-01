/*
 * @lc app=leetcode id=1662 lang=golang
 *
 * [1662] Check If Two String Arrays are Equivalent
 */

package main

// @lc code=start
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	a, b := [2]int{0, 0}, [2]int{0, 0}
	var next1, next2 bool

	for next1, next2 = true, true; next1 && next2; {
		if word1[a[0]][a[1]] != word2[b[0]][b[1]] {
			return false
		}

		if a[1]+1 < len(word1[a[0]]) {
			a[1] = a[1] + 1
		} else if a[0]+1 < len(word1) {
			a[0], a[1] = a[0]+1, 0
		} else {
			next1 = false
		}
		if b[1]+1 < len(word2[b[0]]) {
			b[1] = b[1] + 1
		} else if b[0]+1 < len(word2) {
			b[0], b[1] = b[0]+1, 0
		} else {
			next2 = false
		}

	}
	if !next1 && !next2 {
		return true
	}
	return false
}

// @lc code=end
