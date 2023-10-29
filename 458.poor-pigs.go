/*
 * @lc app=leetcode id=458 lang=golang
 *
 * [458] Poor Pigs
 */
package main

// @lc code=start

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	ops := minutesToTest / minutesToDie
	if minutesToTest%minutesToDie != 0 {
		ops += 1
	}

	for accum, i := 1, 1; i < buckets; i++ {
		accum *= ops + 1
		if accum >= buckets {
			return i
		}
	}

	return 0
}

// @lc code=end
