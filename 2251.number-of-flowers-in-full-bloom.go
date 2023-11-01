/*
 * @lc app=leetcode id=2251 lang=golang
 *
 * [2251] Number of Flowers in Full Bloom
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func fullBloomFlowers(flowers [][]int, people []int) []int {
	result := make([]int, len(people))
	bloomStart := flowers
	sort.Slice(bloomStart, func(i, j int) bool {
		return bloomStart[i][0] < bloomStart[j][0]
	})

	bloomEnd := make([][]int, len(bloomStart))
	copy(bloomEnd, flowers)

	sort.Slice(bloomEnd, func(i, j int) bool {
		return bloomEnd[i][1] < bloomEnd[j][1]
	})
	fmt.Println(bloomStart, bloomEnd)

	for _, v := range people {
		// bloomsStarted := 0

		for l, r := 0, len(people)-1; l <= r; {
			m := (l + r) / 2
			if flowers[m][0] > v {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}

	return result
}

// @lc ctde=end
