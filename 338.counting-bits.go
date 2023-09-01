/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 *
 * https://leetcode.com/problems/counting-bits/description/
 *
 * algorithms
 * Easy (76.32%)
 * Likes:    10027
 * Dislikes: 454
 * Total Accepted:    889.8K
 * Total Submissions: 1.2M
 * Testcase Example:  '2'
 *
 * Given an integer n, return an array ans of length n + 1 such that for each i
 * (0 <= i <= n), ans[i] is the number of 1's in the binary representation of
 * i.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 2
 * Output: [0,1,1]
 * Explanation:
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 *
 *
 * Example 2:
 *
 *
 * Input: n = 5
 * Output: [0,1,1,2,1,2]
 * Explanation:
 * 0 --> 0
 * 1 --> 1
 * 2 --> 10
 * 3 --> 11
 * 4 --> 100
 * 5 --> 101
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= n <= 10^5
 *
 *
 *
 * Follow up:
 *
 *
 * It is very easy to come up with a solution with a runtime of O(n log n). Can
 * you do it in linear time O(n) and possibly in a single pass?
 * Can you do it without using any built-in function (i.e., like
 * __builtin_popcount in C++)?
 *
 *
 */
package main

// func oldCountBits(n int) []int {
// 	// cBits := make([]int, n+1)
// 	cBits := []int{0}
// 	for i := 1; i <= n; i++ {
// 		// bits := 0
// 		// j := i
// 		// for j > 0 {
// 		// 	bits += j % 2
// 		// 	j = j >> 1
// 		// }
// 		// cBits = append(cBits, bits)
// 		// cBits[i] = bits
// 		cBits = append(cBits, oldNOneBits(i))
// 	}
//
// 	return cBits
// }
//
// func oldNOneBits(i int) int {
// 	bits := 0
// 	for i > 0 {
// 		bits += i % 2
// 		i = i >> 1
// 	}
// 	return bits
// }
//
// func nOneBits(i int, bits []int) int {
// 	odd := (i%2 == 1)
// 	prev := bits[i/2]
// 	if odd {
// 		return prev + 1
// 	}
// 	return prev
// }
//
// func countBitsSlice(n int) []int {
// 	// cBits := make([]int, n+1)
// 	cBits := []int{0}
// 	for i := 1; i <= n; i++ {
// 		cBits = append(cBits, cBits[i/2]+i%2)
// 	}
//
// 	return cBits
// }
//
// func countBitsDoubleLoop(n int) []int {
// 	cBits := make([]int, n+2)
// 	// doubled up the loop to try to save on executions. does not seem to have much effect
// 	for i := 0; i <= n; i += 2 {
// 		v := cBits[i/2]
// 		// cBits[i] = cBits[i/2] + i%2
// 		cBits[i] = v
// 		cBits[i+1] = v + 1
// 	}
//
// 	return cBits[:n+1]
// }

// @lc code=start

func countBits(n int) []int {
	cBits := make([]int, n+1)
	for i := 0; i <= n; i++ {
		cBits[i] = cBits[i/2] + i%2
		// cBits[i] = cBits[i/2] + i&1
	}
	return cBits
}

// @lc code=end
