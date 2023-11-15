/*
 * @lc app=leetcode id=1846 lang=golang
 *
 * [1846] Maximum Element After Decreasing and Rearranging
 */
package main

import (
	"sort"
)

// @lc code=start
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	l := len(arr)

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	arr[0] = 1

	for i := 1; i < l; i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}

	return arr[l-1]
}

// @lc code=end
