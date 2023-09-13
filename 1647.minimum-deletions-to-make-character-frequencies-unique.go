/*
 * @lc app=leetcode id=1647 lang=golang
 *
 * [1647] Minimum Deletions to Make Character Frequencies Unique
 */

package main

// @lc code=start
import "sort"

func minDeletions(s string) int {
	hist := make([]int, 26)

	for i := range s {
		hist[s[i]-'a']++
	}

	sort.Slice(hist, func(i, j int) bool {
		return hist[i] < hist[j]
	})

	subtractions := 0
	for i := 24; i >= 0; i-- {
		if hist[i] == 0 {
			break
		}

		if hist[i] >= hist[i+1] {

			prev := hist[i]
			if hist[i+1]-1 > 0 {
				hist[i] = hist[i+1] - 1
			} else {
				hist[i] = 0
			}
			subtractions += prev - hist[i]
		}
	}

	return subtractions
}

// @lc code=end
