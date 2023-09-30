/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */

package main

import "math"

// @lc code=start
func find132pattern(nums []int) bool {
	stack := []int{}
	jValue := math.MinInt64

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < jValue {
			return true
		}

		for len(stack) > 0 && stack[len(stack)-1] < nums[i] {
			jValue = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return false
}

// @lc code=end
