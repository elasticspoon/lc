/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 */
package main

import (
	"math"
)

// @lc code=start
func kthGrammar(n int, k int) int {
	l := int(math.Pow(2, float64(n-1)))
	return recKth(n, k, l)
}

func recKth(n int, k int, l int) int {
	if n == 1 {
		return 0
	}

	half := l / 2
	if k <= half {
		return recKth(n-1, k, half)
	}

	if n%2 == 0 {
		return 1 - recKth(n-1, k-half, half)
	} else {
		return recKth(n-1, l-k+1, half)
	}
}

// @lc code=end
