/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

package main

// @lc code=start
func findDuplicate(nums []int) int {
	fast := nums[0]
	slow := nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := nums[0]
	for slow != slow2 {
		slow = nums[slow]
		slow2 = nums[slow2]
	}
	return slow
}

// func findDuplicate(nums []int) int {
// 	array := make([]bool, len(nums))
// 	for _, num := range nums {
// 		if array[num] {
// 			return num
// 		} else {
// 			array[num] = true
// 		}
// 	}
// 	return -1
// }

// @lc code=end
