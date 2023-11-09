/*
 * @lc app=leetcode id=1759 lang=golang
 *
 * [1759] Count Number of Homogenous Substrings
 */
package main

// @lc code=start

const mod1759 = 1e9 + 7

func countHomogenous(s string) int {
	count := 0
	currCount := 1

	for i, numCons := 1, 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			numCons++
			currCount += numCons
			currCount %= mod1759
			// fmt.Println("cons letter", s[i], "number", numCons, "curr count", currCount)
		} else {
			count += currCount
			currCount %= mod1759
			numCons = 1
			// fmt.Println("non-cons letter", s[i], "number", numCons, "count", count)
			currCount = 1
		}
	}

	return count + currCount
}

// @lc code=end
