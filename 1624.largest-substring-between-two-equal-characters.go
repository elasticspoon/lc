/*
 * @lc app=leetcode id=1624 lang=golang
 *
 * [1624] Largest Substring Between Two Equal Characters
 */
package main

// @lc code=start
func maxLengthBetweenEqualCharacters(s string) int {
	charLocs := make(map[rune]int)

	best := -1

	for i, c := range s {
		if loc, ok := charLocs[c]; ok {
			if i-loc-1 > best {
				best = i - loc - 1
			}
		} else {
			charLocs[c] = i
		}
	}

	return best
}

// @lc code=end
