/*
 * @lc app=leetcode id=1793 lang=golang
 *
 * [1793] Maximum Score of a Good Subarray
 */

package main

// @lc code=start
func maximumScore(nums []int, k int) int {
	left, right := k, k
	minVal := nums[k]
	maxScore := nums[k]

	for left > 0 || right < len(nums)-1 {
		if left > 0 && right < len(nums)-1 {
			if nums[left-1] > nums[right+1] {
				left--
			} else {
				right++
			}
		} else if left > 0 {
			left--
		} else {
			right++
		}

		if nums[left] < minVal {
			minVal = nums[left]
		} else if nums[right] < minVal {
			minVal = nums[right]
		}

		if minVal*(right-left+1) > maxScore {
			maxScore = minVal * (right - left + 1)
		}
	}

	return maxScore
}

// @lc code=end
