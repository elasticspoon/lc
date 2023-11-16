/*
 * @lc app=leetcode id=1980 lang=golang
 *
 * [1980] Find Unique Binary String
 */
package main

import "fmt"

// @lc code=start
func findDifferentBinaryString(nums []string) string {
	present := make(map[string]bool)
	l := len(nums[0])

	for _, num := range nums {
		present[num] = true
	}

	for i := 0; i < 30; i++ {
		s := fmt.Sprintf("%0*b", l, i)
		if !present[s] {
			return s
		}
	}

	return ""
}

// @lc code=end
