/*
 * @lc app=leetcode id=1877 lang=golang
 *
 * [1877] Minimize Maximum Pair Sum in Array
 */
package main

import "sort"

// @lc code=start
func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	max := nums[0] + nums[len(nums)-1]

	for i := 1; i < len(nums)/2; i++ {
		sum := nums[i] + nums[len(nums)-1-i]
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end
