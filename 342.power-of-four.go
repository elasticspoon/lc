/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */
package main

// @lc code=start
// var valMap = map[int]bool{1: true, 1 << 2: true, 1 << 4: true, 1 << 6: true, 1 << 8: true, 1 << 10: true, 1 << 12: true, 1 << 14: true, 1 << 16: true, 1 << 18: true, 1 << 20: true, 1 << 22: true, 1 << 24: true, 1 << 26: true, 1 << 28: true, 1 << 30: true}

const mask = 0x55555555

//	func isPowerOfFour(n int) bool {
//		_, ok := valMap[n]
//		return ok
//	}
func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&mask != 0
}

// @lc code=end
