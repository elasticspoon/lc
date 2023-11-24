/*
 * @lc app=leetcode id=1561 lang=golang
 *
 * [1561] Maximum Number of Coins You Can Get
 */

package main

import "sort"

// @lc code=start
func maxCoins(piles []int) int {
	total := 0

	sort.Slice(piles, func(i, j int) bool { return piles[i] > piles[j] })

	for i := 1; i < len(piles)/3*2; i += 2 {
		total += piles[i]
	}

	return total
}

// @lc code=end
