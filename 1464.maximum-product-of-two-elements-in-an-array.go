/*
 * @lc app=leetcode id=1464 lang=golang
 *
 * [1464] Maximum Product of Two Elements in an Array
 */
package main

// @lc code=start
func maxProduct(nums []int) int {
	a, b := nums[0], nums[1]
	if a < b {
		a, b = b, a
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > a {
			b = a
			a = nums[i]
		} else if nums[i] > b {
			b = nums[i]
		}
	}

	return (a - 1) * (b - 1)
}

// @lc code=end
