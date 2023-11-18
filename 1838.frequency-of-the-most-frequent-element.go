/*
 * @lc app=leetcode id=1838 lang=golang
 *
 * [1838] Frequency of the Most Frequent Element
 */

package main

import (
	"sort"
)

// @lc code=start
func maxFrequency(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	left, right := 0, 1
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
	}

	for right < len(nums) {
		curr := prefix[right] - prefix[left] + nums[left]
		goal := nums[right] * (right - left + 1)
		if goal-curr > k {
			left++
		}
		right++
	}

	return right - left
}

// @lc code=end
