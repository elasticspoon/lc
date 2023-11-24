/*
 * @lc app=leetcode id=1630 lang=golang
 *
 * [1630] Arithmetic Subarrays
 */

package main

import (
	"sort"
)

// @lc code=start
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		vals := nums[l[i] : r[i]+1]
		result[i] = isArithmetic(vals)
	}

	return result
}

func isArithmetic(in []int) bool {
	nums := append([]int{}, in...)
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	diff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			return false
		}
	}
	return true
}

// @lc code=end
