/*
 * @lc app=leetcode id=2225 lang=golang
 *
 * [2225] Find Players With Zero or One Losses
 */

package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func findWinners(matches [][]int) [][]int {
	numLosses := make(map[int]int)

	for i := 0; i < len(matches); i++ {
		if _, ok := numLosses[matches[i][0]]; !ok {
			numLosses[matches[i][0]] = 0
		}
		numLosses[matches[i][1]]++
	}
	fmt.Println(numLosses)

	noLosses := make([]int, 0)
	oneLoss := make([]int, 0)
	for k, v := range numLosses {
		switch v {
		case 0:
			noLosses = append(noLosses, k)
		case 1:
			oneLoss = append(oneLoss, k)
		}
	}

	sort.Ints(noLosses)
	sort.Ints(oneLoss)

	return [][]int{noLosses, oneLoss}
}

// @lc code=end
