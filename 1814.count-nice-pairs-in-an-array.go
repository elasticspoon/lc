/*
 * @lc app=leetcode id=1814 lang=golang
 *
 * [1814] Count Nice Pairs in an Array
 */

package main

// @lc code=start
const mod1814 = 1000000007

func countNicePairs(nums []int) int {
	buckets := make(map[int]int)
	for _, num := range nums {
		buckets[num-reverseInt(num)]++
	}

	numPairs := 0
	for _, v := range buckets {
		if v > 1 {
			// C(n, 2) = n * (n - 1) / 2
			numPairs += (v * (v - 1) / 2) % mod1814
		}
	}
	return numPairs % mod1814
}

func reverseInt(value int) int {
	sum := 0
	for ; value > 0; value /= 10 {
		sum *= 10
		sum += value % 10
	}

	return sum
}

// @lc code=end
