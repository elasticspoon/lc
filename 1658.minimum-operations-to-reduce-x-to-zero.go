/*
 * @lc app=leetcode id=1658 lang=golang
 *
 * [1658] Minimum Operations to Reduce X to Zero
 */

package main

// @lc code=start
func minOperations1658(nums []int, x int) int {
	target, n := -x, len(nums)
	for _, v := range nums {
		target += v
	}

	if target == 0 {
		return n
	}

	maxLen, curSum, left := 0, 0, 0

	for right, v := range nums {
		curSum += v
		for left <= right && curSum > target {
			curSum -= nums[left]
			left++
		}
		if curSum == target {
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}
		}

	}

	if maxLen != 0 {
		return n - maxLen
	} else {
		return -1
	}
}

// func minOperations(nums []int, x int) int {
// 	reverseMap := make(map[int]int)
// 	reverseMap[0] = 0
//
// 	for sum, i := 0, len(nums)-1; i >= 0; i-- {
// 		sum += nums[i]
// 		if sum > x {
// 			break
// 		}
// 		reverseMap[sum] = len(nums) - i
// 	}
//
// 	minOps := 1<<31 - 1
// 	forwardOps := 0
// 	sum := 0
//
// 	for sum <= x && forwardOps < len(nums) {
// 		reverseOps, ok := reverseMap[x-sum]
// 		if ok && reverseOps+forwardOps < minOps && forwardOps+reverseOps <= len(nums) {
// 			minOps = forwardOps + reverseOps
// 		}
//
// 		sum += nums[forwardOps]
// 		forwardOps++
// 	}
//
// 	if minOps == 1<<31-1 {
// 		return -1
// 	} else {
// 		return minOps
// 	}
// }

// @lc code=end
