/*
 * @lc app=leetcode id=135 lang=golang
 *
 * [135] Candy
 */

package main

// @lc code=start

func candy(ratings []int) int {
	totalCandy := 1

	peak, consDec, consInc := 0, 0, 0

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			consInc++
			consDec = 0
			peak = consInc
			totalCandy += consInc + 1
		} else if ratings[i] == ratings[i-1] {
			consInc = 0
			consDec = 0
			peak = 0
			totalCandy++
		} else {
			consDec++
			consInc = 0
			if peak >= consDec {
				totalCandy += consDec
			} else {
				totalCandy += consDec + 1
			}
		}
	}

	return totalCandy
}

// @lc code=end
