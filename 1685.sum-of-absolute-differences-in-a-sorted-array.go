/*
 * @lc app=leetcode id=1685 lang=golang
 *
 * [1685] Sum of Absolute Differences in a Sorted Array
 */

package main

// @lc code=start
func getSumAbsoluteDifferences(nums []int) []int {
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prev := (i+1)*nums[i] - prefixSum[i]
		next := prefixSum[len(nums)-1] - prefixSum[i] - (len(nums)-1-i)*nums[i]
		result[i] = prev + next
	}

	return result
}

// @lc code=end
