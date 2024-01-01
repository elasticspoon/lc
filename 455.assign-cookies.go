/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

package main

import (
	"sort"
)

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	numStatisfied := 0
	for cookie, greed := 0, 0; cookie < len(s) && greed < len(g); {
		if s[cookie] >= g[greed] {
			numStatisfied++
			cookie++
			greed++
		} else {
			cookie++
		}
	}

	return numStatisfied
}

// @lc code=end
