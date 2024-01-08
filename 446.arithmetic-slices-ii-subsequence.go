/*
 * @lc app=leetcode id=446 lang=golang
 *
 * [446] Arithmetic Slices II - Subsequence
 */

package main

import "fmt"

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	counted := make(map[int]int)
	numSubseqs := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := 1; j < len(nums)-1; j += j {

			diff := nums[i+j] - nums[i]
			fmt.Println("i", i, "diff:", diff)
			if index, ok := counted[diff]; ok && index > i {
				continue
			}

			prev := nums[i+1]
			seqLen := 2
			for k := i + 2; k < len(nums); k++ {
				if nums[k]-prev == diff {
					seqLen++
					prev = nums[k]
				}
			}

			if seqLen < 3 {
				continue
			}

			counted[diff] = i + seqLen - 1
			fmt.Println("counted:", counted)
			fmt.Println("seqLen:", seqLen)

			if diff == 0 {
				numSubseqs += zeroDiffSubseqs(seqLen)
			} else {
				numSubseqs += (seqLen - 2) * (seqLen - 1) / 2
			}
		}
	}

	return numSubseqs
}

func zeroDiffSubseqs(n int) int {
	memo := make([]int, n+1)

	var fact func(int) int
	fact = func(start int) int {
		if memo[start] > 0 {
			return memo[start]
		}
		if start == 0 {
			return 1
		}
		res := start * fact(start-1)
		memo[start] = res
		return res
	}

	numSubseqs := 0
	for i := 3; i <= n; i++ {
		numSubseqs += fact(n) / (fact(i) * fact(n-i))
	}

	return numSubseqs
}

// @lc code=end
