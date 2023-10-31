/*
 * @lc app=leetcode id=1356 lang=golang
 *
 * [1356] Sort Integers by The Number of 1 Bits
 */
package main

import "sort"

// @lc code=start
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		bitsA := numOneBits(arr[i])
		bitsB := numOneBits(arr[j])
		if bitsA == bitsB {
			return arr[i] < arr[j]
		} else {
			return bitsA < bitsB
		}
	})
	return arr
}

func numOneBits(v int) int {
	count := 0
	for v != 0 {
		v = v & (v - 1)
		count++
	}
	return count
}

// @lc code=end
